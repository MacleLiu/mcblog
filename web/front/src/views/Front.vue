<template>
  <div class="frontContainer">
      <FrontHeader/>
      <router-view :key="$route.path"></router-view>
      <FrontFooter/>
</div>
</template>

<script>
  import FrontHeader from '../components/front/FrontHeader'
  import FrontFooter from '../components/front/FrontFooter'

  export default {
    components: {FrontHeader, FrontFooter}
  }
</script>

<style scoped>
  .frontContainer {
    height: 100%;
    display: flex;
    flex-direction: column;
  }
</style>
  