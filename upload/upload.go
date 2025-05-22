package upload

import (
	"context"
	"mcblog/config"
	"mcblog/utils/errno"
	"mime/multipart"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

var AccessKey = config.AppConfig.QiNiu.AccessKey
var SecretKey = config.AppConfig.QiNiu.SecretKey

func UpLoadFile(file multipart.File, fileSize int64) (string, error) {
	putPolicy := storage.PutPolicy{
		Scope: config.AppConfig.QiNiu.Bucket,
	}
	mac := qbox.NewMac(AccessKey, SecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Region = getRegionByName(config.AppConfig.QiNiu.Region)
	// 是否使用https域名
	cfg.UseHTTPS = config.AppConfig.QiNiu.UseHttps
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = config.AppConfig.QiNiu.UseCdnDomains

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	putExtra := storage.PutExtra{}

	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		return "", errno.New(errno.ERROR, err)
	}

	return config.AppConfig.QiNiu.QiNiuServer + "/" + ret.Key, nil
}

// 将配置文件中空间机房位置配置做映射
// 映射 region 字段为 storage.Zone 对象
func getRegionByName(name string) *storage.Region {
	switch name {
	case "Huadong":
		return &storage.ZoneHuadong
	case "Huabei":
		return &storage.ZoneHuabei
	case "Huanan":
		return &storage.ZoneHuanan
	case "Beimei":
		return &storage.ZoneBeimei
	case "Xinjiapo":
		return &storage.ZoneXinjiapo
	case "ShouEr1":
		return &storage.ZoneShouEr1
	case "ZheJiang2":
		return &storage.ZoneHuadongZheJiang2
	default:
		return nil // 或者 return &storage.ZoneHuadong 做默认处理
	}
}
