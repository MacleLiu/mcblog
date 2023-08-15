<template>
    <div>
        <div class="logo">
            <span>McBlog</span>
        </div>

        <a-menu theme="dark" mode="inline" @click="goToPage">
            <a-menu-item key="index"><a-icon type="dashboard" /><span>仪表盘</span></a-menu-item>

            <a-sub-menu>
                <span slot="title"><a-icon type="file" /><span>文章管理</span></span>
                <a-menu-item key="addArt"><a-icon type="form" /><span>写文章</span></a-menu-item>
                <a-menu-item key="articles"><a-icon type="unordered-list" /><span>文章列表</span></a-menu-item>
            </a-sub-menu>

            <a-menu-item key="categories"><a-icon type="book" /><span>分类列表</span></a-menu-item>

            <a-menu-item key="users"><a-icon type="user" /><span>用户列表</span></a-menu-item>
        </a-menu>
    </div>
</template>

<script>
    export default {
        methods: {
            goToPage(item) {
                this.$router.push('/admin/' + item.key).catch((err) => err )
            }
        }
    }
</script>

<style>
    .logo {
        width: 100%;
        margin: 10px;
        text-align: center;
        font-size: 25px;
        font-weight: bold;
        color: white;
    }
</style>