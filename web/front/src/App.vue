<template>
  <a-config-provider :locale="zh_CN">
    <div id="app">
      <router-view></router-view>
    </div>
  </a-config-provider>
</template>

<script>
import zh_CN from 'ant-design-vue/lib/locale-provider/zh_CN';
import moment from 'moment';
import 'moment/locale/zh-cn';

moment.locale('zh-cn');

export default {
  name: 'App',
  data() {
    return {
      zh_CN,
    };
  },
}
</script>

<style>
  #app {
    min-height: 100%;
  }
  
  /* 重写默认样式 */
  html, body {
    width: 100%;
    min-height: 100%;
  }
</style>
