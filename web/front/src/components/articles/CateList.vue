<template>
    <a-space direction="vertical" :size="size"  v-title data-title="分类">
        <div style="width: 100%;">
            <ali-icon style="font-size: 20px; margin-right: 10px;" type="icon-FolderOpen-1" />
            <h2 style="display: inline;">分类{{ artlist.length === 0 ? '' : ' - ' + artlist[0].Category.name }}</h2>
            <hr style="width: 100%;" color="skyblue">
        </div>
        <a-list :locale="defaultText" :loading="loading" item-layout="horizontal" :bordered="false" :split="false" :pagination="pagination" :data-source="artlist">
            <a-list-item slot="renderItem" slot-scope="item, index">
                <a-card :hoverable="true">
                    <div style="width: 100%;">
                        <h2 style="display:inline-block;" @click="readArticle(item.ID)">{{ item.title }}</h2>
                    </div>
                    <div style="width: 100%;">
                        <ali-icon type="icon-date1"/>
                        <span>{{ item.CreatedAt | dateFormat }}</span>
                        <span style="margin: 0 10px;">|</span>
                        <ali-icon type="icon-folderopen"/>
                        <span>{{ item.Category.name }}</span>
                    </div>
                    <div style="width: 100%; margin-top: 10px; font-size: 16px;">
                        <p>{{ item.desc }}</p>
                    </div>
                </a-card>
            </a-list-item>
        </a-list>
    </a-space>
</template>

<script>
export default {
    props: ['cid'],
    data() {
        return {
            loading: true,
            defaultText: {
                emptyText: '暂无数据'
            },
            pagination: {
                pageSize: 10,
                total: 0,
                showSizeChanger: false,
                hideOnSinglePage: true,
                showTotal: (total) => `共${ total }条`,
                onChange: (page) => {
                    this.queryParam.pagenum = page
                    this.getArtListByCate(this.cid)
                }
            },
            queryParam: {
                pagesize: 10,
                pagenum: 1,
            },
            artlist: [],
            catename: '',
            size: 'middle',
        }
    },
    methods: {
        // 查询分类下文章列表
        async getArtListByCate(cid) {
            try{
                const { data : res } = await this.$http.get(`article/cate/${ cid }`, {
                    params: {
                        pagesize: this.queryParam.pagesize,
                        pagenum: this.queryParam.pagenum,
                    },
                })
                if (res.status != 200){
                    this.loading=false
                    this.$message.error(res.msg)
                    return
                }
                if (res.data.length === 0) {
                    this.$router.push(`/`).catch((err) => err )
                    return this.$message.warning("当前分类为空")
                }
                this.artlist = res.data
                this.pagination.total = res.total
                this.loading=false
            }catch(err){
                this.loading = false
                this.defaultText.emptyText = err.message
                return
            }
        },
        // 阅读文章
        readArticle(id) {
              this.$router.push(`/blog/article/${ id }`).catch((err) => err )
        },
    },
  
    created () {
        this.getArtListByCate(this.cid)
    },
}
</script>

<style scoped>
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
h2 {  
    color : black ; 
    transition: 0.3 s; 
    }
h2:hover {  color : rgb(30, 180, 255) ; }
</style>