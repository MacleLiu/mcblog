<template>
    <div>
        <div class="pageTitle">用户列表</div>
        <a-card>
            <a-row :gutter="20">
                <a-col :span="6">
                    <a-input-search 
                        v-model="queryParam.username" 
                        placeholder="输入用户名查找" 
                        enter-button 
                        @search="getUserList"
                        allowClear 
                    /></a-col>
                <a-col :span="4">
                    <a-button type="primary" @click="addUserVisible = true">新增</a-button>
                </a-col>
            </a-row>
            <!-- 用户列表表格 -->
            <a-table 
                class="antTable"
                rowKey="ID" 
                :columns="columns" 
                :dataSource="userlist" 
                :pagination="pagination" 
                @change="handleTableChange"
                bordered
            >
                <span slot="role" slot-scope="role">{{ role === 1 ? '管理员' : '读者' }}</span>
                <span slot="action" class="actionSlot" slot-scope="user">
                        <a-button type="primary" @click="editUser(user.ID)">编辑</a-button>
                        <a-button type="danger" @click="delUser(user.ID)">删除</a-button>
                </span>
            </a-table>
        </a-card>
        <!-- 新增用户对话框 -->
        <a-modal
            title="新增用户"
            :visible="addUserVisible"
            @ok="addUserOk"
            @cancel="addUserCancel"
        >
            <a-form-model :model="newUser" :rules="addUserRules" ref="addUserRef">
                <a-form-model-item :hasFeedback="true" label="用户名" prop="username">
                    <a-input v-model="newUser.username"></a-input>
                </a-form-model-item>
                <a-form-model-item :hasFeedback="true" label="密码" prop="password">
                    <a-input-password v-model="newUser.password"></a-input-password>
                </a-form-model-item>
                <a-form-model-item :hasFeedback="true" label="确认密码" prop="checkpassword">
                    <a-input-password v-model="newUser.checkpassword"></a-input-password>
                </a-form-model-item>
                <a-form-model-item label="是否为管理员">
                    <a-select v-model="newUser.role" :defaultValue="2" @change="isAdmin">
                        <a-select-option :value="1">是</a-select-option>
                        <a-select-option :value="2">否</a-select-option>
                    </a-select>
                </a-form-model-item>
            </a-form-model>
        </a-modal>

        <!-- 编辑用户对话框 -->
        <a-modal
            title="编辑用户"
            :visible="editUserVisible"
            @ok="editUserOk"
            @cancel="editUserCancel"
        >
            <a-form-model :model="userInfo" :rules="userRules" ref="editUserRef">
                <a-form-model-item :hasFeedback="true" label="用户名" prop="username">
                    <a-input v-model="userInfo.username"></a-input>
                </a-form-model-item>
                <a-form-model-item label="是否为管理员">
                    <a-select v-model="userInfo.role" :defaultValue="userInfo.role" @change="editAdmin">
                        <a-select-option :value="1">是</a-select-option>
                        <a-select-option :value="2">否</a-select-option>
                    </a-select>
                </a-form-model-item>
            </a-form-model>
        </a-modal>
    </div>
</template>

<script>

    const columns = [
        {
            title: 'ID',
            dataIndex: 'ID',
            width: '10%',
            key: 'id',
            align: 'center',
        },
        {
            title: '用户名',
            dataIndex: 'username',
            width: '20%',
            key: 'username',
            align: 'center',
        },
        {
            title: '角色',
            dataIndex: 'role',
            width: '20%',
            key: 'role',
            scopedSlots: {customRender: 'role'},
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
                    username: '',
                },
                userlist: [],
                columns,
                addUserVisible: false,
                userInfo: {
                    id: 0,
                    username: '',
                    role: 2,
                },
                newUser: {
                    id: 0,
                    username: '',
                    password: '',
                    checkpassword: '',
                    role: 2,
                },
                userRules: {
                    username: [
                        { required: true, message: '请输入用户名', trigger: 'change' },
                        { min: 4, max: 12, message: '用户名在4到12个字符之间', trigger: 'change' },
                    ],
                    password: [
                        {
                            required: true,
                            validator: (_, value, callback) => {
                                if (value === '') {
                                    callback(new Error('请输入密码'));
                                } 
                                if ([...value].length < 6 || [...value].length > 20) {
                                    callback(new Error('密码在6到20个字符之间'))
                                } else {
                                    callback()
                                }
                            },
                            trigger: 'change',
                        }
                    ],
                    checkpassword: [
                        {
                            required: true,
                            validator: (_, value, callback) => {
                                if (value === '') {
                                    callback(new Error('请输入密码'));
                                } 
                                if (value !== this.userInfo.password) {
                                    callback(new Error('密码不一致，请检查后重新输入'))
                                } else {
                                    callback()
                                }
                            },
                            trigger: 'change',
                        }
                    ],
                    
                },
                addUserRules: {
                    username: [
                        { required: true, message: '请输入用户名', trigger: 'change' },
                        { min: 4, max: 12, message: '用户名在4到12个字符之间', trigger: 'change' },
                    ],
                },
                editUserVisible: false,
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
                this.getUserList()
            },
            // 查询用户列表
            async getUserList() {
                const { data : res } = await this.$http.get('users', {
                    params: {
                        pagesize: this.queryParam.pagesize,
                        pagenum: this.queryParam.pagenum,
                        username: this.queryParam.username,
                    },
                })
                if (res.status != 200) return this.$message.error(res.msg)
                this.userlist = res.data
                this.pagination.total = res.total
            },
            // 删除用户
            delUser(id) {
                this.$confirm({
                    title: '确定要删除该用户吗?',
                    content: '删除后不可恢复',
                    okText: '确定',
                    okType: 'danger',
                    cancelText: '取消',
                    onOk: async () => {
                        const res = await this.$http.delete(`user/${id}`)
                        if (res.status != 200) return this.$message.error(res.msg)
                        this.$message.success('删除成功')
                        this.getUserList()
                    },
                    onCancel() {},
                })
            },
            // 新增用户
            addUserOk() {
                this.$refs.addUserRef.validate(async valid => {
                    if (valid) {
                        const { data:res } = await this.$http.post('user/add', {
                            username: this.newUser.username,
                            password: this.newUser.password,
                            role: this.newUser.role,
                        })
                        if (res.status !== 200) return this.$message.error(res.msg)
                        this.addUserVisible = false
                        this.$message.success('新增用户成功')
                        this.$refs.addUserRef.resetFields() // 重置表单
                        this.getUserList() // 刷新用户列表
                    } else {
                        return this.$message.error('参数不符合要求，请检查')
                    }
                });
            },
            addUserCancel() {
                this.$refs.addUserRef.resetFields()
                this.addUserVisible = false
            },
            isAdmin(value) {
                this.newUser.role = value
            },
            // 编辑用户
            async editUser(id){
                this.editUserVisible = true
                const { data : res} = await this.$http.get(`user/${ id }`)
                this.userInfo.username = res.data.username
                this.userInfo.role = res.data.role
                this.userInfo.id = id
            },
            editAdmin(value) {
                this.userInfo.role = value
            },
            editUserOk() {
                this.$refs.editUserRef.validate(async valid => {
                    if (valid) {
                        const { data:res } = await this.$http.put(`user/${ this.userInfo.id }`, {
                            username: this.userInfo.username,
                            role: this.userInfo.role,
                        })
                        if (res.status !== 200) return this.$message.error(res.msg)
                        this.editUserVisible = false
                        this.$message.success('更新用户成功')
                        this.$refs.editUserRef.resetFields() // 重置表单
                        this.getUserList() // 刷新用户列表
                    } else {
                        return this.$message.error('参数不符合要求，请检查')
                    }
                });
            },
            editUserCancel() {
                this.$refs.editUserRef.resetFields()
                this.editUserVisible = false
            },
        },
        created() {
            this.getUserList()
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