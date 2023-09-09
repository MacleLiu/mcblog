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
	cfg.Region = &storage.ZoneBeimei
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	putExtra := storage.PutExtra{}

	err := formUploader.PutWithoutKey(context.Background(), &ret, upToken, file, fileSize, &putExtra)
	if err != nil {
		return "", errno.New(errno.ERROR, err)
	}

	return config.AppConfig.QiNiu.QiNiuServer + "/" + ret.Key, nil
}
