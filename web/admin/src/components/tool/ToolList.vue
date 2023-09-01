<template>
    <div>
        <div class="pageTitle">工具列表</div>
        <a-card>
            <a-row :gutter="20">
                <a-col :span="4">
                    <a-button type="primary" @click="addToolVisible = true">新增工具</a-button>
                </a-col>
            </a-row>
            <!-- 工具列表表格 -->
            <a-table 
                class="antTable"
                rowKey="id" 
                :columns="columns" 
                :dataSource="Toollist" 
                :pagination="pagination" 
                @change="handleTableChange"
                bordered
            >
                <span slot="action" class="actionSlot" slot-scope="Tool">
                        <a-button type="primary" @click="editTool(Tool.id)">编辑</a-button>
                        <a-button type="danger" @click="delTool(Tool.id)">删除</a-button>
                </span>
            </a-table>
        </a-card>
        <!-- 新增工具对话框 -->
        <a-modal
            title="新增工具"
            :visible="addToolVisible"
            @ok="addToolOk"
            @cancel="addToolCancel"
        >
            <a-form-model :model="newTool" :rules="addToolRules" ref="addToolRef">
                <a-form-model-item :hasFeedback="true" label="工具名" prop="name">
                    <a-input v-model="newTool.name"></a-input>
                </a-form-model-item>
                <a-form-model-item :hasFeedback="true" label="链接" prop="url">
                    <a-input v-model="newTool.url"></a-input>
                </a-form-model-item>
            </a-form-model>
        </a-modal>

        <!-- 编辑工具对话框 -->
        <a-modal
            title="编辑工具"
            :visible="editToolVisible"
            @ok="editToolOk"
            @cancel="editToolCancel"
        >
            <a-form-model :model="ToolInfo" :rules="ToolRules" ref="editToolRef">
                <a-form-model-item :hasFeedback="true" label="工具名" prop="name">
                    <a-input v-model="ToolInfo.name"></a-input>
                </a-form-model-item>
                <a-form-model-item :hasFeedback="true" label="链接" prop="url">
                    <a-input v-model="ToolInfo.url"></a-input>
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
            title: '工具名',
            dataIndex: 'name',
            width: '20%',
            key: 'name',
            align: 'center',
        },
        {
            title: '链接',
            dataIndex: 'url',
            width: '40%',
            key: 'url',
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
                Toollist: [],
                columns,
                addToolVisible: false,
                ToolInfo: {
                    id: 0,
                    name: '',
                    url: '',
                },
                newTool: {
                    name: '',
                    url: '',
                },
                ToolRules: {
                    name: [
                        { required: true, message: '请输入工具名', trigger: 'change' },
                        { max: 40, message: '工具名不超过40个字符', trigger: 'change' },
                    ],
                    url: [
                        { required: true, message: '请输入工具链接', trigger: 'change' },
                        { max: 100, message: '链接不超过100个字符', trigger: 'change' },
                    ],
                },
                addToolRules: {
                    name: [
                        { required: true, message: '请输入工具名', trigger: 'change' },
                        { max: 40, message: '工具名不超过40个字符', trigger: 'change' },
                    ],
                    url: [
                        { required: true, message: '请输入工具链接', trigger: 'change' },
                        { max: 100, message: '链接不超过100个字符', trigger: 'change' },
                    ],
                },
                editToolVisible: false,
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
                this.getToolList()
            },
            // 查询工具列表
            async getToolList() {
                const { data : res } = await this.$http.get('tools', {
                    params: {
                        pagesize: this.queryParam.pagesize,
                        pagenum: this.queryParam.pagenum,
                    },
                })
                if (res.status != 200) return this.$message.error(res.msg)
                this.Toollist = res.data
                this.pagination.total = res.total
            },
            // 删除工具
            delTool(id) {
                this.$confirm({
                    title: '确定要删除该工具吗?',
                    content: '删除后不可恢复',
                    okText: '确定',
                    okType: 'danger',
                    cancelText: '取消',
                    onOk: async () => {
                        const res = await this.$http.delete(`tool/${id}`)
                        if (res.status != 200) return this.$message.error(res.msg)
                        this.$message.success('删除成功')
                        this.getToolList()
                    },
                    onCancel() {},
                })
            },
            // 新增工具
            addToolOk() {
                this.$refs.addToolRef.validate(async valid => {
                    if (valid) {
                        const { data:res } = await this.$http.post('tool/add', {
                            name: this.newTool.name,
                            url: this.newTool.url,
                        })
                        if (res.status !== 200) return this.$message.error(res.msg)
                        this.addToolVisible = false
                        this.$message.success('新增工具成功')
                        this.$refs.addToolRef.resetFields() // 重置表单
                        this.getToolList() // 刷新工具列表
                    } else {
                        return this.$message.error('参数不符合要求，请检查')
                    }
                });
            },
            addToolCancel() {
                this.$refs.addToolRef.resetFields()
                this.addToolVisible = false
            },
            isAdmin(value) {
                this.newTool.role = value
            },
            // 编辑工具
            async editTool(id){
                this.editToolVisible = true
                const { data : res} = await this.$http.get(`tool/${ id }`)
                this.ToolInfo.name = res.data.name
                this.ToolInfo.url = res.data.url
                this.ToolInfo.id = id
            },
            editToolOk() {
                this.$refs.editToolRef.validate(async valid => {
                    if (valid) {
                        const { data:res } = await this.$http.put(`tool/${ this.ToolInfo.id }`, {
                            name: this.ToolInfo.name,
                            url: this.ToolInfo.url,
                        })
                        if (res.status !== 200) return this.$message.error(res.msg)
                        this.editToolVisible = false
                        this.$message.success('更新工具成功')
                        this.$refs.editToolRef.resetFields() // 重置表单
                        this.getToolList() // 刷新用户列表
                    } else {
                        return this.$message.error('参数不符合要求，请检查')
                    }
                });
            },
            editToolCancel() {
                this.$refs.editToolRef.resetFields()
                this.editToolVisible = false
            },
        },
        created() {
            this.getToolList()
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