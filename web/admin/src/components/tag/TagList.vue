<template>
    <div>
        <div class="pageTitle">标签列表</div>
        <a-card>
            <a-row :gutter="20">
                <a-col :span="4">
                    <a-button type="primary" @click="addTagVisible = true">新增标签</a-button>
                </a-col>
            </a-row>
            <!-- 标签列表表格 -->
            <a-table 
                class="antTable"
                rowKey="id" 
                :columns="columns" 
                :dataSource="Taglist" 
                :pagination="pagination" 
                @change="handleTableChange"
                bordered
            >
                <span slot="action" class="actionSlot" slot-scope="Tag">
                        <a-button type="primary" @click="editTag(Tag.id)">编辑</a-button>
                        <a-button type="danger" @click="delTag(Tag.id)">删除</a-button>
                </span>
            </a-table>
        </a-card>
        <!-- 新增标签对话框 -->
        <a-modal
            title="新增标签"
            :visible="addTagVisible"
            @ok="addTagOk"
            @cancel="addTagCancel"
        >
            <a-form-model :model="newTag" :rules="tagRules" ref="addTagRef">
                <a-form-model-item :hasFeedback="true" label="标签名" prop="name">
                    <a-input v-model="newTag.name"></a-input>
                </a-form-model-item>
            </a-form-model>
        </a-modal>

        <!-- 编辑标签对话框 -->
        <a-modal
            title="编辑标签"
            :visible="editTagVisible"
            @ok="editTagOk"
            @cancel="editTagCancel"
        >
            <a-form-model :model="TagInfo" :rules="tagRules" ref="editTagRef">
                <a-form-model-item :hasFeedback="true" label="标签名" prop="name">
                    <a-input v-model="TagInfo.name"></a-input>
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
            title: '标签名',
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
                Taglist: [],
                columns,
                addTagVisible: false,
                TagInfo: {
                    id: 0,
                    name: '',
                },
                newTag: {
                    name: '',
                },
                tagRules: {
                    name: [
                        { pattern: /(^\S)((.)*\S)?(\S*$)/, message: '首尾不能有空格' },
                        { required: true, message: '请输入标签名', trigger: 'change' },
                        { max: 20, message: '标签名不超过20个字符', trigger: 'change' },
                    ],
                },
                editTagVisible: false,
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
                this.getTagList()
            },
            // 查询标签列表
            async getTagList() {
                const { data : res } = await this.$http.get('tags', {
                    params: {
                        pagesize: this.queryParam.pagesize,
                        pagenum: this.queryParam.pagenum,
                    },
                })
                if (res.status != 200) return this.$message.error(res.msg)
                this.Taglist = res.data
                this.pagination.total = res.total
            },
            // 删除标签
            delTag(id) {
                this.$confirm({
                    title: '确定要删除该标签吗?',
                    content: '删除后不可恢复',
                    okText: '确定',
                    okType: 'danger',
                    cancelText: '取消',
                    onOk: async () => {
                        const { data : res } = await this.$http.delete(`tag/${id}`)
                        if (res.status != 200){
                            this.$message.error(res.msg)
                            return
                        }
                        this.$message.success('删除成功')
                        this.getTagList()
                    },
                    onCancel() {},
                })
            },
            // 新增标签
            addTagOk() {
                this.$refs.addTagRef.validate(async valid => {
                    if (valid) {
                        const { data:res } = await this.$http.post('tag/add', {
                            name: this.newTag.name,
                        })
                        if (res.status !== 200) return this.$message.error(res.msg)
                        this.addTagVisible = false
                        this.$message.success('新增标签成功')
                        this.$refs.addTagRef.resetFields() // 重置表单
                        this.getTagList() // 刷新标签列表
                    } else {
                        return this.$message.error('参数不符合要求，请检查')
                    }
                });
            },
            addTagCancel() {
                this.$refs.addTagRef.resetFields()
                this.addTagVisible = false
            },
            isAdmin(value) {
                this.newTag.role = value
            },
            // 编辑标签
            async editTag(id){
                this.editTagVisible = true
                const { data : res} = await this.$http.get(`tag/${ id }`)
                this.TagInfo.name = res.data.name
                this.TagInfo.id = id
            },
            editTagOk() {
                this.$refs.editTagRef.validate(async valid => {
                    if (valid) {
                        const { data:res } = await this.$http.put(`tag/${ this.TagInfo.id }`, {
                            name: this.TagInfo.name,
                        })
                        if (res.status !== 200) return this.$message.error(res.msg)
                        this.editTagVisible = false
                        this.$message.success('更新标签成功')
                        this.$refs.editTagRef.resetFields() // 重置表单
                        this.getTagList() // 刷新标签列表
                    } else {
                        return this.$message.error('参数不符合要求，请检查')
                    }
                });
            },
            editTagCancel() {
                this.$refs.editTagRef.resetFields()
                this.editTagVisible = false
            },
        },
        created() {
            this.getTagList()
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