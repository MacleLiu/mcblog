<template>
    <a-space direction="vertical" :size="size">
        <a-list item-layout="horizontal" :bordered="false" :split="false" :pagination="pagination" :data-source="artlist">
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
    data() {
        return {
            pagination: {
                pageSize: 10,
                total: 0,
                showSizeChanger: false,
                hideOnSinglePage: true,
                showTotal: (total) => `共${ total }条`,
                onChange: (page) => {
                    this.queryParam.pagenum = page
                    this.getArtList()
                }
            },
            queryParam: {
                pagesize: 10,
                pagenum: 1,
                title: '',
            },
            artlist: [],
            size: 'middle',
        }
    },
    methods: {
        // 查询文章列表
        async getArtList() {
            const { data : res } = await this.$http.get('articles', {
                params: {
                    pagesize: this.queryParam.pagesize,
                    pagenum: this.queryParam.pagenum,
                    title: this.queryParam.title,
                },
            })
            if (res.status != 200) return this.$message.error(res.msg)
            this.artlist = res.data
            this.pagination.total = res.total
        },
        // 阅读文章
        readArticle(id) {
              this.$router.push(`/blog/article/${ id }`).catch((err) => err )
        },
    },
  
    created () {
        this.getArtList()
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