<template>
    <!-- 归档时间轴区域 -->
    <a-card :hoverable="true" v-title data-title="归档">
        <a-timeline :reverse="true">
            <a-spin :spinning="loading">
                <a-timeline-item v-for="item in artInfo" :key="item.ID">
                    <span style="margin-right: 10px;">{{ item.CreatedAt | dateFormat }}</span>
                    <span class="withLink" style="font-size: 16px; font-weight: bold;" @click="readArticle(item.ID)">{{ item.title }}</span>
                </a-timeline-item>
            </a-spin>
        </a-timeline>
    </a-card>
</template>

<script>
export default {
    data() {
        return {
            loading: true,
            artInfo: [],
        }
    },
    methods: {
        async getArtInfo() {
            try{
                const { data : res } = await this.$http.get('allartinfo')
                if (res.status != 200){
                    this.loading=false
                    this.$message.error(res.msg)
                    return
                }
                this.artInfo = res.data
                this.loading = false
            }catch(err){
                this.loading = false
                return
            }
        },
        // 阅读文章
        readArticle(id) {
              this.$router.push(`/blog/article/${ id }`).catch((err) => err )
        },
    },

    created() {
        this.getArtInfo()
    }
};
</script>

<style scoped>
.ant-card {
    /* flex: auto; */
    border-radius: 10px;
    box-shadow: 2px 8px 8px -4px rgb(200, 200, 200);
}
:deep .ant-card-body {
    width: 100%;
    display: flex;
    flex-direction: column;
    align-items: flex-start;
}
.withLink {
    color : black ; 
    transition: 0.3 s;
}.withLink:hover {  color : rgb(30, 180, 255); }
a {  
    color : black ; 
    transition: 0.3 s; 
}
</style>