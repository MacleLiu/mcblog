<template>
    <div>
        <div class="pageTitle">分类列表</div>
        <a-card>
            <a-row :gutter="20">
                <a-col :span="4">
                    <a-button type="primary" @click="addCateVisible = true">新增分类</a-button>
                </a-col>
            </a-row>
            <!-- 分类列表表格 -->
            <a-table 
                class="antTable"
                rowKey="id" 
                :columns="columns" 
                :dataSource="Catelist" 
                :pagination="pagination" 
                @change="handleTableChange"
                bordered
            >
                <span slot="action" class="actionSlot" slot-scope="Cate">
                        <a-button type="primary" @click="editCate(Cate.id)">编辑</a-button>
                        <a-button type="danger" @click="delCate(Cate.id)">删除</a-button>
                </span>
            </a-table>
        </a-card>
        <!-- 新增分类对话框 -->
        <a-modal
            title="新增分类"
            :visible="addCateVisible"
            @ok="addCateOk"
            @cancel="addCateCancel"
        >
            <a-form-model :model="newCate" :rules="addCateRules" ref="addCateRef">
                <a-form-model-item :hasFeedback="true" label="分类名" prop="name">
                    <a-input v-model="newCate.name"></a-input>
                </a-form-model-item>
            </a-form-model>
        </a-modal>

        <!-- 编辑分类对话框 -->
        <a-modal
            title="编辑分类"
            :visible="editCateVisible"
            @ok="editCateOk"
            @cancel="editCateCancel"
        >
            <a-form-model :model="CateInfo" :rules="CateRules" ref="editCateRef">
                <a-form-model-item :hasFeedback="true" label="分类名" prop="name">
                    <a-input v-model="CateInfo.name"></a-input>
                </a-form-model-item>
            </a-form-model>
        </a-modal>
    </div>
</template>

<script>

    const columns = [
        {
            title: 'ID',
            dataIndex: 'id',
            width: '10%',
            key: 'id',
            align: 'center',
        },
        {
            title: '分类名',
            dataIndex: 'name',
            width: '20%',
            key: 'name',
            align: 'center',
        },
        {
            title: '操作',
            width: '30%',
            key: 'action',
            scopedSlots: {customRender: 'action'},
            align: 'center',
        },

    ]

    export default {
        data() {
            return {
                pagination: {
                    pageSizeOptions: ['5', '10', '20'],
                    pageSize: 5,
                    total: 0,
                    showSizeChanger: true,
                    showTotal: (total) => `共${ total }条`,
                },
                queryParam: {
                    pagesize: 5,
                    pagenum: 1,
                },
                Catelist: [],
                columns,
                addCateVisible: false,
                CateInfo: {
                    id: 0,
                    name: '',
                },
                newCate: {
                    name: '',
                },
                CateRules: {
                    name: [
                        { required: true, message: '请输入分类名', trigger: 'change' },
                        { max: 20, message: '分类名不超过20个字符', trigger: 'change' },
                    ],
                },
                addCateRules: {
                    name: [
                        { required: true, message: '请输入分类名', trigger: 'change' },
                        { max: 20, message: '分类名不超过20个字符', trigger: 'change' },
                    ],
                },
                editCateVisible: false,
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
                this.getCateList()
            },
            // 查询分类列表
            async getCateList() {
                const { data : res } = await this.$http.get('categories', {
                    params: {
                        pagesize: this.queryParam.pagesize,
                        pagenum: this.queryParam.pagenum,
                    },
                })
                if (res.status != 200) return this.$message.error(res.msg)
                this.Catelist = res.data
                this.pagination.total = res.total
            },
            // 删除分类
            delCate(id) {
                this.$confirm({
                    title: '确定要删除该分类吗?',
                    content: '删除后不可恢复',
                    okText: '确定',
                    okType: 'danger',
                    cancelText: '取消',
                    onOk: async () => {
                        const { data : res } = await this.$http.delete(`category/${id}`)
                        if (res.status != 200){
                            this.$message.error(res.msg)
                            return
                        }
                        this.$message.success('删除成功')
                        this.getCateList()
                    },
                    onCancel() {},
                })
            },
            // 新增分类
            addCateOk() {
                this.$refs.addCateRef.validate(async valid => {
                    if (valid) {
                        const { data:res } = await this.$http.post('category/add', {
                            name: this.newCate.name,
                        })
                        if (res.status !== 200) return this.$message.error(res.msg)
                        this.addCateVisible = false
                        this.$message.success('新增分类成功')
                        this.$refs.addCateRef.resetFields() // 重置表单
                        this.getCateList() // 刷新分类列表
                    } else {
                        return this.$message.error('参数不符合要求，请检查')
                    }
                });
            },
            addCateCancel() {
                this.$refs.addCateRef.resetFields()
                this.addCateVisible = false
            },
            isAdmin(value) {
                this.newCate.role = value
            },
            // 编辑分类
            async editCate(id){
                this.editCateVisible = true
                const { data : res} = await this.$http.get(`category/${ id }`)
                this.CateInfo.name = res.data.name
                this.CateInfo.id = id
            },
            editCateOk() {
                this.$refs.editCateRef.validate(async valid => {
                    if (valid) {
                        const { data:res } = await this.$http.put(`category/${ this.CateInfo.id }`, {
                            name: this.CateInfo.name,
                        })
                        if (res.status !== 200) return this.$message.error(res.msg)
                        this.editCateVisible = false
                        this.$message.success('更新标签成功')
                        this.$refs.editCateRef.resetFields() // 重置表单
                        this.getCateList() // 刷新标签列表
                    } else {
                        return this.$message.error('参数不符合要求，请检查')
                    }
                });
            },
            editCateCancel() {
                this.$refs.editCateRef.resetFields()
                this.editCateVisible = false
            },
        },
        created() {
            this.getCateList()
        }
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
</style>