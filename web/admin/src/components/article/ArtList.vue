<template>
    <div>
        <div class="pageTitle">文章列表</div>
        <a-card>
            <a-row :gutter="20">
                <a-col :span="6">
                    <a-input-search 
                        v-model="queryParam.title" 
                        placeholder="输入标题查找" 
                        enter-button 
                        @search="getArtList"
                        allowClear 
                    /></a-col>
                <a-col :span="4">
                    <a-button type="primary" @click="addArt">新增</a-button>
                </a-col>

                <a-col :span="6" :offset="4">
                    <a-select placeholder="请选择分类" :allowClear="true" style="width: 200px;" @change="cateChange">
                        <a-select-option v-for="item in catelist"  :key="item.id" :value="item.id">{{ item.name }}</a-select-option>
                    </a-select>
                </a-col>
            </a-row>
            <!-- 文章列表表格 -->
            <a-table 
                class="antTable"
                rowKey="ID" 
                :columns="columns" 
                :dataSource="artlist" 
                :pagination="pagination" 
                @change="handleTableChange"
                bordered
            >
                <span slot="winnow" slot-scope="art">
                    <a-switch v-model="art.winnow" @change="setWinnow(art.ID)"/>
                </span>
                <span slot="img" class="imgSlot" slot-scope="img">
                    <img :src="img">
                </span>
                <span slot="action" class="actionSlot" slot-scope="art">
                        <a-button type="primary" @click="editArt(art.ID)">编辑</a-button>
                        <a-button type="danger" @click="delArt(art.ID)">删除</a-button>
                </span>
            </a-table>
        </a-card>
    </div>
</template>

<script>
    const columns = [
        {
            title: 'ID',
            dataIndex: 'ID',
            width: '5%',
            key: 'id',
            align: 'center',
        },
        {
            title: '分类',
            dataIndex: 'Category.name',
            width: '10%',
            key: 'name',
            align: 'center',
        },
        {
            title: '文章标题',
            dataIndex: 'title',
            width: '15%',
            key: 'title',
            align: 'center',
        },
        {
            title: '文章描述',
            dataIndex: 'desc',
            width: '20%',
            key: 'desc',
            align: 'center',
        },
        {
            title: '缩略图',
            dataIndex: 'img',
            width: '10%',
            key: 'img',
            scopedSlots: {customRender: 'img'},
            align: 'center',
        },
        {
            title: '精选文章',
            width: '10%',
            key: 'winnow',
            scopedSlots: {customRender: 'winnow'},
            align: 'center',
        },
        {
            title: '操作',
            width: '10%',
            key: 'action',
            scopedSlots: {customRender: 'action'},
            align: 'center',
        },

    ]

    export default {
        data() {
            return {
                pagination: {
                    pageSizeOptions: ['10', '15', '20'],
                    pageSize: 10,
                    total: 0,
                    showSizeChanger: true,
                    showTotal: (total) => `共${ total }条`,
                },
                queryParam: {
                    pagesize: 10,
                    pagenum: 1,
                    title: '',
                },
                artlist: [],
                catelist: [],
                columns,
            }
        },
        methods: {
            // 分页查询
            handleTableChange(pagination) {
                let pager = { ...this.pagination };
                pager.current = pagination.current
                pager.pageSize = pagination.pageSize

                this.queryParam.pagesize = pagination.pageSize
                this.queryParam.pagenum = pagination.current

                if (pagination.pageSize != this.pagination.pageSize){
                    this.queryParam.pagenum = 1
                    pager.current = 1
                }
                this.pagination = pager
                this.getArtList()
            },
            // 获取分类列表
            async getCateList() {
                const { data : res } = await this.$http.get('categories')
                if (res.status != 200) return this.$message.error(res.msg)
                this.catelist = res.data
            },
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
            // 新增文章
            addArt() {
                this.$router.push(`/admin/addart`)
            },
            // 删除文章标签（删除关联关系）
            async delArticleTags(id) {
                const { data : res } = await this.$http.delete(`tag/article/${ id }`)
                if (res.status != 200) return this.$message.error(res.msg)
            },
            // 删除文章
            delArt(id) {
                this.$confirm({
                    title: '确定要删除该篇文章吗?',
                    content: '删除后不可恢复',
                    okText: '确定',
                    okType: 'danger',
                    cancelText: '取消',
                    onOk: async () => {
                        const { data : res } = await this.$http.delete(`article/${id}`)
                        if (res.status != 200) return this.$message.error(res.msg)
                        this.delArticleTags(id)
                        this.$message.success('删除成功')
                        this.getArtList()
                    },
                    onCancel() {},
                })
            },
            // 编辑文章
            editArt(id){
                this.$router.push(`/admin/addart/${ id }`)
            },
            // 查询分类下文章
            cateChange(value) {
                console.log(value)
                if (typeof(value) == "undefined") {
                    this.getArtList()
                } else {
                    this.getArtByCate(value)
                }
            },
            async getArtByCate(id) {
                const { data : res } = await this.$http.get(`article/cate/${ id }`, {
                    params: {
                        pagesize: this.queryParam.pagesize,
                        pagenum: this.queryParam.pagenum,
                    },
                })
                if (res.status != 200) return this.$message.error(res.msg)
                this.artlist = res.data
                this.pagination.total = res.total
            },
            // 设置精选文章
            async setWinnow(id){
                const { data : res } = await this.$http.put(`article/winnow/${ id }`)
                if (res.status != 200) return this.$message.error(res.msg)
                this.$message.success("修改成功")
            }
        },

        created() {
            this.getArtList()
            this.getCateList()
        },
    }
</script>

<style scoped>
    .pageTitle {
        text-align: center;
        font-size: 30px;
        font-weight: bold;
        letter-spacing: 20px;
        color: black;
        font-family: Serif;
    }
    .actionSlot {
        display: flex;
        justify-content: center;
    }
    .actionSlot button {
        margin-right: 5px;
        margin-left: 5px;
    }
    .antTable {
        margin-top: 10px;
    }
    .imgSlot img {
        width: 100px;
        height: 60px;
    }
</style>