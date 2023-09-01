<template>
    <a-space class="introSpace" direction="vertical" :size="size">
        <a-card :hoverable="true">
            <a-avatar :size="81" src="http://rzp1xpqex.hn-bkt.clouddn.com/grw.jpg"/>
            <h2 style="text-align: center;">Macle</h2>
            <p style="text-align: center; font-size: 16px;">
                既往不念，当下不杂，未来不迎
            </p>
            <a-row 
                style="width: 100%;
                    display: flex; 
                    justify-content: space-between; 
                    text-align: center;"
            >
                <a-col :span="8">
                    <h3>文章</h3>
                </a-col>
                <a-col :span="8">
                    <h3>分类</h3>
                </a-col>
                <a-col :span="8">
                    <h3>标签</h3>
                </a-col>
            </a-row>
            <a-row 
                style="width: 100%;
                    display: flex; 
                    justify-content: space-between; 
                    text-align: center;"
            >
                <a-col :span="8">
                    <h3>{{ arttotal }}</h3>
                </a-col>
                <a-col :span="8">
                    <h3>{{ catetotal }}</h3>
                </a-col>
                <a-col :span="8">
                    <h3>{{ tagtotal }}</h3>
                </a-col>
            </a-row>
            <a-button 
                style="
                    width: 80%; 
                    margin: 10px 0;
                    display: flex;
                    justify-content: center;
                    align-items: center;" 
                type="primary"
                @click="followMe"
            >
                <a-icon type="github" />
                <span>Follow Me</span>
            </a-button>
            <div style="margin-top: 10px; 
                    width: 80%; 
                    display: flex;
                    justify-content: space-around;"
            >
                <a href="https://github.com/MacleLiu" target="_blank" title="github">
                    <ali-icon style="font-size: 20px;" type="icon-GitHub" />
                </a>
                <a href="https://space.bilibili.com/89687310" target="_blank" title="bilibili">
                    <ali-icon style="font-size: 20px;" type="icon-bilibili" />
                </a>                            
                <ali-icon style="font-size: 20px;" type="icon-weibo" />
            </div>
        </a-card>
        <!-- 精选文章卡片 -->
        <a-card :hoverable="true">
            <div style="width: 100%; margin-bottom: 20px;">
                <ali-icon style="font-size: 20px; margin-right: 10px;" type="icon-hot" />
                <h2 style="display: inline;">精选文章</h2>
            </div>       
            <!-- 列表 -->
            <div style="width: 100%;">
            <a-list item-layout="vertical" :bordered="false" :split="false" :data-source="winnow">
                <a-list-item slot="renderItem" slot-scope="item, index">
                    <div style="width: 100%; font-size: 15px; margin-bottom: 2px;" @click="readArticle(item.ID)">
                        <span class="withLink">{{ item.title }}</span>
                    </div>
                    <div style="width: 100%;">
                        <ali-icon type="icon-date1"/>
                        <span>{{ item.CreatedAt | dateFormat }}</span>
                        <span style="margin: 0 10px;">|</span>
                        <ali-icon type="icon-folderopen"/>
                        <span>{{ item.Category.name }}</span>
                    </div>
                </a-list-item>
            </a-list>
            </div>
        </a-card>
        <!-- 分类卡片 -->
        <a-card :hoverable="true">
            <div style="width: 100%; margin-bottom: 20px;">
                <ali-icon style="font-size: 20px; margin-right: 10px;" type="icon-FolderOpen-1" />
                <h2 style="display: inline;">分类</h2>
            </div>
            <!-- 列表 -->
            <div style="width: 100%;">
            <a-list item-layout="vertical" :bordered="false" :split="false" :data-source="catestat">
                <a-list-item slot="renderItem" slot-scope="item, index">
                    <div class="withLink" style="width: 100%; font-size: 16px; display: flex; justify-content: space-between;" @click="goArtListByCate(item.id)">
                        <span>{{ item.name }}</span>
                        <span>{{ item.count }}</span>
                    </div>
                </a-list-item>
            </a-list>
            </div>
        </a-card>
        <!-- 标签卡片 -->
        <a-card :hoverable="true">
            <div style="width: 100%; margin-bottom: 20px;">
                <ali-icon style="font-size: 20px; margin-right: 10px;" type="icon-24gf-tags4" />
                <h2 style="display: inline;">标签</h2>
            </div>
            
            <div  style="width: 100%;">
              <a-tag v-for="item in taglist" color="blue" @click="goArtListByTag(item.id)">
                {{ item.name }}
              </a-tag>
            </div>
        </a-card>
    </a-space>
</template>

<script>
export default {
    data() {
        return {
            winnow: [],
            arttotal: 0,
            catetotal: 0,
            catestat: [],
            taglist: [],
            tagtotal: 0,
            size: 'middle',
        }
    },
    methods: {
        // 获取文章总数
        async getArtCount() {
            const { data : res } = await this.$http.get('articlecount', {
            })
            if (res.status != 200) return this.$message.error(res.msg)
            this.arttotal = res.data
        },
        // 获取分类统计
        async getCateStat() {
            const { data : res } = await this.$http.get('catestat')
            if (res.status != 200) return this.$message.error(res.msg)
            this.catestat = res.data
            if (res.data != null)
            this.catetotal = res.data.length
        },
        // 进入分类下的文章列表
        goArtListByCate(cid) {
              this.$router.push(`/blog/cate/${ cid }`).catch((err) => err )
        },
        // 查询精选文章列表
        async getWinnowList() {
            const { data : res } = await this.$http.get('winnow')
            if (res.status != 200) return this.$message.error(res.msg)
            this.winnow = res.data
        },
        // 获取标签列表
        async getTags() {
            const { data : res } = await this.$http.get('tags')
            if (res.status != 200) return this.$message.error(res.msg)
            this.taglist = res.data
            if (res.data != null)
            this.tagtotal = res.data.length
        },
        // 进入标签下的文章列表
        goArtListByTag(tid) {
            this.$router.push(`/blog/tag/${ tid }`).catch((err) => err )
        },
        // 阅读文章
        readArticle(id) {
              this.$router.push(`/blog/article/${ id }`).catch((err) => err )
        },
        followMe() {
            window.open('http://www.github.com/MacleLiu', '_blank')
        },
    },
  
    created () {
        this.getArtCount()
        this.getWinnowList()
        this.getCateStat()
        this.getTags()
    },
    }
</script>

<style scoped>
.introSpace {
    display: flex;
    flex-direction: column;
}
.ant-card {
    flex: auto;
    border-radius: 10px;
    box-shadow: 2px 8px 8px -4px rgb(200, 200, 200);
}
:deep .ant-card-body {
    width: 100%;
    display: flex;
    flex-direction: column;
    align-items: center;
}
.ant-list-item {
    padding: 0 0;
    margin-bottom: 20px;
}
.ant-tag {
    margin-bottom: 10px;
}
.withLink {
    color : black ; 
    transition: 0.3 s;
}.withLink:hover {  color : rgb(30, 180, 255); }
a {  
    color : black ; 
    transition: 0.3 s; 
}
a:hover {  color : rgb(30, 180, 255); }
</style>