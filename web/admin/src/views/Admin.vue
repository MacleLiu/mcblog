<template>
    <a-layout class="admin">
      <a-layout-sider breakpoint="lg">
          <AdminSider></AdminSider>
      </a-layout-sider>

      <a-layout>
        <a-layout-header class="header">
          <AdminHeader></AdminHeader>
        </a-layout-header>

        <a-layout-content>
          <router-view></router-view>
        </a-layout-content>

        <a-layout-footer class="footer">
          <AdminFooter></AdminFooter>
        </a-layout-footer>  
      </a-layout>
      
    </a-layout>
</template>

<script>
  import AdminFooter from '../components/admin/AdminFooter'
  import AdminHeader from '../components/admin/AdminHeader'
  import AdminSider from '../components/admin/AdminSider'

  export default {
    components: { AdminFooter, AdminHeader, AdminSider }
  }
</script>

<style scoped>
  .admin {
    height: 100%
  }
  .header {
    display: flex;
    justify-content: flex-end;
    align-items: center;
  }
  .footer{
    text-align: center;
    background-color: #eee;
    font-size: 20px;
  }
</style>