<template>
  <div class="frontContainer" v-title data-title="首页">
      <FrontHeader/>

      <a-row class="indexRow">
          <a-col class="introCol" :xs="24" :sm="24" :md="8" :lg="6" :xl="5">
            <LeftIntro/>
          </a-col>
          <a-col class="infoCol" :xs="24" :sm="24" :md="16" :lg="18" :xl="14">
            <ArticleList/>
          </a-col>
      </a-row>

      <FrontFooter/>
      <div>
        <a-back-top :visibilityHeight="600"/>
      </div>
  </div>
</template>
  
<script>
import FrontHeader from '../components/front/FrontHeader'
import FrontFooter from '../components/front/FrontFooter'
import LeftIntro from '../components/front/LetfIntro.vue'
import ArticleList from '../components/articles/ArticleList.vue'

export default {
  components: {FrontHeader, FrontFooter, LeftIntro, ArticleList},
}
</script>
  
<style scoped>
.frontContainer {
    display: flex;
    flex-direction: column;
    min-height: 100vh;
}
.indexRow {
    flex: 1 0 auto;
    display: flex;
    height: 100%;
    justify-content: center;
    flex-wrap: wrap;
    align-content: flex-start;
    padding: 80px 0 20px 0;
    background-color: rgb(240, 245, 245);
}
.introCol {
    /* display: flex; */
    /* flex-direction: column; */
    padding: 0 10px;
}
.infoCol {
    display: flex;
    flex-direction: column;
    padding: 0 10px;
}
.ant-tag {
    margin-bottom: 10px;
}
</style>
    