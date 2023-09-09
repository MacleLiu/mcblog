<template>
    <a-space direction="vertical" :size="size"  v-title data-title="标签">
        <div style="width: 100%;">
            <ali-icon style="font-size: 20px; margin-right: 10px;" type="icon-24gf-tags4" />
            <h2 style="display: inline;">标签{{ tag.name === '' ? '' : ' - ' + tag.name }}</h2>
            <hr style="width: 100%;" color="skyblue">
        </div>
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
    props: ['tid'],
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
            },
            artlist: [],
            tag: {
                id: 0,
                name: '',
            },
            size: 'middle',
        }
    },
    methods: {
        // 查询标签下文章列表
        async getArtListByTag(tid) {
            const { data : res } = await this.$http.get(`article/tag/${ tid }`, {
                params: {
                    pagesize: this.queryParam.pagesize,
                    pagenum: this.queryParam.pagenum,
                },
            })
            if (res.status != 200) return this.$message.error(res.msg)
            if (res.data.length === 0) {
                this.$router.push(`/`).catch((err) => err )
                return this.$message.warning("当前标签为空")
            }
            this.artlist = res.data
            this.pagination.total = res.total
        },
        // 获取标签信息
        async getTag(tid) {
            const { data : res } = await this.$http.get(`tag/${ tid }`)
            if (res.status != 200) return this.$message.error(res.msg)
            this.tag = res.data
        },
        // 阅读文章
        readArticle(id) {
              this.$router.push(`/blog/article/${ id }`).catch((err) => err )
        },
    },
  
    created () {
        this.getArtListByTag(this.tid)
        this.getTag(this.tid)
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