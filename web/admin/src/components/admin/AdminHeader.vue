<template>
    <div>
        <a-popconfirm 
            title="确定要退出吗？" 
            placement="left"
            ok-text="确定" 
            cancel-text="取消" 
            @confirm="confirm"
        >
        <a-button type="danger">退出</a-button>
      </a-popconfirm>
    </div>
</template>

<script>
    export default {
        methods: {
            logout(){
                window.sessionStorage.clear('token')
                this.$router.push('login').catch(err => err)
            },
            confirm() {
                this.logout()
            },
        }
    }
</script>