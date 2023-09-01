<template>
    <div>
        <div class="pageTitle">心愿列表</div>
        <a-card>
            <a-row :gutter="20">
                <a-col :span="4">
                    <a-button type="primary" @click="addWishVisible = true">新增心愿</a-button>
                </a-col>
            </a-row>
            <!-- 心愿列表表格 -->
            <a-table 
                class="antTable"
                rowKey="id" 
                :columns="columns" 
                :dataSource="Wishlist" 
                :pagination="pagination" 
                @change="handleTableChange"
                bordered
            >
                <span slot="status" slot-scope="status">
                    {{ status ? '已完成' : '未完成' }}
                </span>
                <span slot="action" class="actionSlot" slot-scope="Wish">
                    <a-button type="primary" @click="editWish(Wish.id)">编辑</a-button>
                    <a-button type="danger" @click="delWish(Wish.id)">删除</a-button>
                </span>
            </a-table>
        </a-card>
        <!-- 新增心愿对话框 -->
        <a-modal
            title="新增心愿"
            :visible="addWishVisible"
            @ok="addWishOk"
            @cancel="addWishCancel"
        >
            <a-form-model :model="newWish" :rules="addWishRules" ref="addWishRef">
                <a-form-model-item :hasFeedback="true" label="心愿名" prop="name">
                    <a-input v-model="newWish.name"></a-input>
                </a-form-model-item>
            </a-form-model>
        </a-modal>

        <!-- 编辑心愿对话框 -->
        <a-modal
            title="编辑心愿"
            :visible="editWishVisible"
            @ok="editWishOk"
            @cancel="editWishCancel"
        >
            <a-form-model :model="WishInfo" :rules="WishRules" ref="editWishRef">
                <a-form-model-item :hasFeedback="true" label="心愿名" prop="name">
                    <a-input v-model="WishInfo.name"></a-input>
                </a-form-model-item>
                <a-form-model-item :hasFeedback="true" label="状态" prop="status">
                    <a-switch default-checked v-model="WishInfo.status" />
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
            title: '心愿',
            dataIndex: 'name',
            width: '20%',
            key: 'name',
            align: 'center',
        },
        {
            title: '状态',
            dataIndex: 'status',
            width: '20%',
            key: 'status',
            scopedSlots: {customRender: 'status'},
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
                    pageSizeOptions: ['10', '15', '20'],
                    pageSize: 10,
                    total: 0,
                    showSizeChanger: true,
                    showTotal: (total) => `共${ total }条`,
                },
                queryParam: {
                    pagesize: 10,
                    pagenum: 1,
                },
                Wishlist: [],
                columns,
                addWishVisible: false,
                WishInfo: {
                    id: 0,
                    name: '',
                    status: false,
                },
                newWish: {
                    name: '',
                    status: false,
                },
                WishRules: {
                    name: [
                        { required: true, message: '请输入心愿', trigger: 'change' },
                        { max: 40, message: '心愿名不超过40个字符', trigger: 'change' },
                    ],
                },
                addWishRules: {
                    name: [
                        { required: true, message: '请输入心愿', trigger: 'change' },
                        { max: 40, message: '心愿名不超过40个字符', trigger: 'change' },
                    ],
                },
                editWishVisible: false,
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
                this.getWishList()
            },
            // 查询心愿列表
            async getWishList() {
                const { data : res } = await this.$http.get('wishes', {
                    params: {
                        pagesize: this.queryParam.pagesize,
                        pagenum: this.queryParam.pagenum,
                    },
                })
                if (res.status != 200) return this.$message.error(res.msg)
                this.Wishlist = res.data
                this.pagination.total = res.total
            },
            // 删除心愿
            delWish(id) {
                this.$confirm({
                    title: '确定要删除该心愿吗?',
                    content: '删除后不可恢复',
                    okText: '确定',
                    okType: 'danger',
                    cancelText: '取消',
                    onOk: async () => {
                        const res = await this.$http.delete(`wish/${id}`)
                        if (res.status != 200) return this.$message.error(res.msg)
                        this.$message.success('删除成功')
                        this.getWishList()
                    },
                    onCancel() {},
                })
            },
            // 新增心愿
            addWishOk() {
                this.$refs.addWishRef.validate(async valid => {
                    if (valid) {
                        const { data:res } = await this.$http.post('wish/add', {
                            name: this.newWish.name,
                        })
                        if (res.status !== 200) return this.$message.error(res.msg)
                        this.addWishVisible = false
                        this.$message.success('新增心愿成功')
                        this.$refs.addWishRef.resetFields() // 重置表单
                        this.getWishList() // 刷新心愿列表
                    } else {
                        return this.$message.error('参数不符合要求，请检查')
                    }
                });
            },
            addWishCancel() {
                this.$refs.addWishRef.resetFields()
                this.addWishVisible = false
            },
            isAdmin(value) {
                this.newWish.role = value
            },
            // 编辑心愿
            async editWish(id){
                this.editWishVisible = true
                const { data : res} = await this.$http.get(`wish/${ id }`)
                this.WishInfo.name = res.data.name
                this.WishInfo.status = res.data.status
                this.WishInfo.id = id
            },
            editWishOk() {
                this.$refs.editWishRef.validate(async valid => {
                    if (valid) {
                        const { data:res } = await this.$http.put(`wish/${ this.WishInfo.id }`, {
                            name: this.WishInfo.name,
                            status: this.WishInfo.status,
                        })
                        if (res.status !== 200) return this.$message.error(res.msg)
                        this.editWishVisible = false
                        this.$message.success('更新心愿成功')
                        this.$refs.editWishRef.resetFields() // 重置表单
                        this.getWishList() // 刷新用户列表
                    } else {
                        return this.$message.error('参数不符合要求，请检查')
                    }
                });
            },
            editWishCancel() {
                this.$refs.editWishRef.resetFields()
                this.editWishVisible = false
            },
        },
        created() {
            this.getWishList()
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