<template>
    <div class="login">
        <div class="loginBox">
            <a-form-model ref="loginForm" :model="formdata" :rules="rules" class="loginForm">
                <a-form-model-item prop="username">
                    <a-input v-model="formdata.username" placeholder="请输入用户名">
                        <a-icon slot="prefix" type="user" style="color: rgba(0,0,0,.25)" />
                    </a-input>
                </a-form-model-item>

                <a-form-model-item prop="password">
                    <a-input 
                        v-model="formdata.password" 
                        placeholder="请输入密码" 
                        type="password"
                        v-on:keyup.enter="login"
                    >
                        <a-icon slot="prefix" type="lock" style="color: rgba(0,0,0,.25)" />
                    </a-input>
                </a-form-model-item>

                <a-form-model-item>
                    <a-button @click="login" type="primary" style="width: 100%;">登录</a-button>
                </a-form-model-item>
            </a-form-model>
    </div>
    </div>
</template>

<script>
    export default{
        data(){
            return {
                formdata: {
                    username:'',
                    password:'',
                },
                rules: {
                    username: [
                        { required: true, message: '请输入用户名', trigger: 'blur' },
                        { min: 4, max: 12, message: '用户名在4到12个字符之间', trigger: 'blur' },
                    ],
                    password: [
                        { required: true, message: '请输入密码', trigger: 'blur' },
                        { min: 6, max: 20, message: '密码在6到20个字符之间', trigger: 'blur' },
                    ],
                }
            }
        },
        methods: {
            login(){
                this.$refs.loginForm.validate(async valid => {
                    if (!valid) return this.$message.error('非法输入，检查后重新输入')
                    const { data:res } = await this.$http.post('login', this.formdata)
                    if(res.status != 200) return this.$message.error(res.msg)
                    window.sessionStorage.setItem('token', res.token)
                    this.$router.push('admin/index')
                })
            }
        }
    }
</script>

<style scoped>
    .login {
        height: 100%;
        background-color:black;
    }
    .loginBox{
        width: 400px;
        height: 300px;
        background-color: white;
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
        border-radius: 8px;
    }
    .loginForm{
        width: 100%;
        position: absolute;
        bottom: 10%;
        padding: 0 20px;
        box-sizing: border-box;
    }
</style>