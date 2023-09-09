<template>
    <div class="container" v-title data-title="文章">
        <FrontHeader/>
        <a-row class="readRow">
            <a-col class="introCol" :xs="0" :sm="0" :md="6" :lg="6" :xl="4">
                <a-affix :offset-top="top">
                    <a-space direction="vertical" :size="size">
                        <!-- 摘要卡片 -->
                        <a-card :hoverable="true">
                            <div style="width: 100%; margin-bottom: 20px;">
                                <h2 style="display: inline;">摘要</h2>
                            </div>
                                <!-- 目录 -->
                            <div style="width: 100%;">
                                <p>{{ artInfo.desc }}</p>
                            </div>
                        </a-card>
                        <!-- 目录卡片 -->
                        <a-card :hoverable="true">
                            <div style="width: 100%; margin-bottom: 20px;">
                                <h2 style="display: inline;">目录</h2>
                            </div>
                                <!-- 目录 -->
                            <div style="width: 100%;">
                                <p>开发中......</p>
                            </div>
                        </a-card>
                    </a-space>
                </a-affix>
            </a-col>
            
            <a-col class="infoCol" :xs="24" :sm="24" :md="18" :lg="18" :xl="12">
                <a-space direction="vertical" :size="size">
                    <a-card :hoverable="true">
                        <!-- 文章标题 -->
                        <h1>{{ artInfo.title }}</h1>
                        <hr style="width: 100%;" color="#DDDDDD">
                        <div style="width: 100%; margin-bottom: 20px;">
                            <ali-icon type="icon-date1"/>
                            <span>发表于{{ artInfo.CreatedAt | dateFormat }}</span>
                            <span style="margin: 0 10px;">|</span>
                            <ali-icon type="icon-folderopen"/>
                            <span>最后更新于{{ artInfo.UpdatedAt| dateFormat }}</span>
                            <span style="margin: 0 10px;">|</span>
                            <ali-icon type="icon-folderopen"/>
                            <span>{{ artInfo.Category.name }}</span>
                        </div>
                        <!-- 文章内容 -->
                        <div v-html="artInfo.content" style="width: 100%; font-size: 16px;"></div>
                        <hr style="width: 100%; border: 1px dashed skyblue;">
                        <!-- 文章标签 -->
                        <div style="width: 100%;">
                            <a-tag v-for="item in taglist" :key="item.id" color="blue">
                                {{ item.name }}
                            </a-tag>
                        </div>
                    </a-card>
                </a-space>
            </a-col>
        </a-row>
        <FrontFooter/>
        <div>
            <a-back-top :visibilityHeight="600"/>
        </div>
    </div>
</template>

<script>
import FrontHeader from '../front/FrontHeader'
import FrontFooter from '../front/FrontFooter'

import Prism from "prismjs";

export default {
    components: {FrontHeader, FrontFooter},
    props: ['id'],

    data() {
        return {
            artInfo :{
                id :0,
                title: '',
                cid: 0,
                Category:{
                    id: 0,
                    name: '',
                },
                desc: '',
                content: '',
                img: '',
            },
            taglist: [],
            size: 'middle',
            top: 10,
        }
    },

    methods: {
        // 查询文章信息
        async getArtInfo(id) {
            const { data : res } = await this.$http.get(`article/${ id }`)
            if (res.status != 200) return this.$message.error(res.msg)
            this.artInfo = res.data
            this.artInfo.id = res.data.ID
        },
        // 查询文章标签
        async getArtTags(id) {
            const { data : res } = await this.$http.get(`tags/${ id }`)
            if (res.status != 200) return this.$message.error(res.msg)
            this.taglist = res.data
        },
    },

    created() {
        Promise.all([this.getArtInfo(this.id), this.getArtTags(this.id)]).then(()=>{
            Prism.highlightAll()  //  代码高亮渲染函数
        })
        // this.getArtInfo(this.id)
        // this.getArtTags(this.id)
    },
}
</script>

<style scoped>
.container {
    height: 100%;
    display: flex;
    flex-direction: column;
}
.readRow {
    flex: 1 0 auto;
    display: flex;
    justify-content: center;
    flex-wrap: wrap;
    align-content: flex-start;
    margin: 40px 0;
}
.introCol {
    /* display: flex; */
    flex-direction: column;
    padding: 0 10px;
}
.infoCol {
    display: flex;
    flex-direction: column;
    padding: 0 10px;
}
.ant-card {
    flex: auto;
    border-radius: 10px;
    box-shadow: 2px 8px 8px -4px rgb(200, 200, 200);
}
:deep .ant-card-body {
    width: 100%;
    display: flex;
    flex-direction: column;
    align-items: center;
}
.ant-tag {
    margin-bottom: 10px;
}
:deep img {
    width: 100%;
    height: auto;
}
</style>