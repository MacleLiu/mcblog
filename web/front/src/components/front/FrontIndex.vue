<template>
        <a-row class="indexRow">
            <a-col :xs="24" :sm="24" :md="6" :lg="6" :xl="4">
                <a-space class="introCol" direction="vertical" :size="size">
                    <a-card :hoverable="true">
                        <a-avatar :size="81" src="http://rzp1xpqex.hn-bkt.clouddn.com/grw.jpg"/>
                        <h2 style="text-align: center;">Macle</h2>
                        <p style="text-align: center; font-size: 16px;">座右铭</p>
                        <a-row 
                            style="width: 100%;
                                display: flex; 
                                justify-content: space-between; 
                                text-align: center;"
                        >
                            <a-col :span="8">
                                <h3>文章</h3>
                            </a-col>
                            <a-col :span="8">
                                <h3>分类</h3>
                            </a-col>
                            <a-col :span="8">
                                <h3>标签</h3>
                            </a-col>
                        </a-row>
                        <a-row 
                            style="width: 100%;
                                display: flex; 
                                justify-content: space-between; 
                                text-align: center;"
                        >
                            <a-col :span="8">
                                <h3>{{ pagination.total }}</h3>
                            </a-col>
                            <a-col :span="8">
                                <h3>{{ catetotal }}</h3>
                            </a-col>
                            <a-col :span="8">
                                <h3>99</h3>
                            </a-col>
                        </a-row>
                        <a-button 
                            style="
                                width: 80%; 
                                margin: 10px 0;
                                display: flex;
                                justify-content: center;
                                align-items: center;" 
                            type="primary"
                        >
                            <a-icon type="github" />
                            <span>Follow Me</span>
                        </a-button>
                        <div style="margin-top: 10px; 
                                width: 80%; 
                                display: flex;
                                justify-content: space-around;"
                        >
                            <a href="https://github.com/MacleLiu" target="_blank" title="github">
                                <ali-icon style="font-size: 20px;" type="icon-GitHub" />
                            </a>
                            <a href="https://space.bilibili.com/89687310" target="_blank" title="bilibili">
                                <ali-icon style="font-size: 20px;" type="icon-bilibili" />
                            </a>
                            
                            <ali-icon style="font-size: 20px;" type="icon-weibo" />
                        </div>
                    </a-card>

                    <!-- 精选文章卡片 -->
                    <a-card :hoverable="true" style>
                        <div style="width: 100%; margin-bottom: 20px;">
                            <ali-icon style="font-size: 20px; margin-right: 10px;" type="icon-hot" />
                            <h2 style="display: inline;">精选文章</h2>
                        </div>
                        
                        <!-- 列表 -->
                        <div style="width: 100%;">
                        <a-list item-layout="vertical" :bordered="false" :split="false" :data-source="artlist">
                            <a-list-item slot="renderItem" slot-scope="item, index"  v-if="item.winnow">
                                <div style="width: 100%; font-size: 15px; margin-bottom: 2px;">
                                    <a href="" >{{ item.title }}</a>
                                </div>
                                <div style="width: 100%;">
                                    <ali-icon type="icon-date1"/>
                                    <span>{{ item.CreatedAt | dateFormat }}</span>
                                    <span style="margin: 0 10px;">|</span>
                                    <ali-icon type="icon-folderopen"/>
                                    <span>{{ getCateByCid(item.cid) }}</span>
                                </div>
                            </a-list-item>
                        </a-list>
                        </div>

                    </a-card>

                    <!-- 分类卡片 -->
                    <a-card :hoverable="true">
                        <div style="width: 100%; margin-bottom: 20px;">
                            <ali-icon style="font-size: 20px; margin-right: 10px;" type="icon-FolderOpen-1" />
                            <h2 style="display: inline;">分类</h2>
                        </div>

                        <!-- 列表 -->
                        <div style="width: 100%;">
                        <a-list item-layout="vertical" :bordered="false" :split="false" :data-source="catestat">
                            <a-list-item slot="renderItem" slot-scope="item, index">
                                <a href="" >
                                    <div style="width: 100%; font-size: 16px; display: flex; justify-content: space-between;">
                                        <span>{{ item.name }}</span>
                                        <span>{{ item.count }}</span>
                                    </div>
                                </a>
                            </a-list-item>
                        </a-list>
                        </div>
                    </a-card>

                    <!-- 标签卡片 -->
                    <a-card :hoverable="true">
                        <div style="width: 100%; margin-bottom: 20px;">
                            <ali-icon style="font-size: 20px; margin-right: 10px;" type="icon-24gf-tags4" />
                            <h2 style="display: inline;">标签</h2>
                        </div>
                        
                        <div>
                          <a-tag color="pink">
                            pink
                          </a-tag>
                          <a-tag color="red">
                            red
                          </a-tag>
                          <a-tag color="orange">
                            orange
                          </a-tag>
                          <a-tag color="green">
                            green
                          </a-tag>
                          <a-tag color="cyan">
                            cyan
                          </a-tag>
                          <a-tag color="blue">
                            blue
                          </a-tag>
                          <a-tag color="purple">
                            purple
                          </a-tag>
                        </div>
                    </a-card>
                </a-space>
            </a-col>

            <!-- 文章列表区域 -->
            <a-col :xs="24" :sm="24" :md="18" :lg="18" :xl="12">
                <a-space class="infoCol" direction="vertical" :size="size">
                    <a-list item-layout="horizontal" :bordered="false" :split="false" :pagination="pagination" :data-source="artlist">
                        <a-list-item slot="renderItem" slot-scope="item, index">
                            <a-card :hoverable="true">
                                <div style="width: 100%; font-size: 24px;">
                                    <a href="" >{{ item.title }}</a>
                                </div>
                                <div style="width: 100%;">
                                    <ali-icon type="icon-date1"/>
                                    <span>{{ item.CreatedAt | dateFormat }}</span>
                                    <span style="margin: 0 10px;">|</span>
                                    <ali-icon type="icon-folderopen"/>
                                    <span>{{ getCateByCid(item.cid) }}</span>
                                </div>
                                <div style="width: 100%; margin-top: 10px; font-size: 16px;">
                                    <p>文章描述</p>
                                </div>
                            </a-card>
                        </a-list-item>
                    </a-list>
                </a-space>
            </a-col>
        </a-row>
</template>

<script>
export default {
    data() {
        return {
            pagination: {
                pageSizeOptions: ['5', '10', '20'],
                pageSize: 10,
                total: 0,
                showSizeChanger: false,
                showTotal: (total) => `共${ total }条`,
            },
            queryParam: {
                pagesize: 5,
                pagenum: 1,
                title: '',
            },
            artlist: [],
            catelist: [],
            catestat: [],
            catetotal: 0,
            size: 'middle',
        }
    },
    methods: {
        // 查询文章列表
        async getArtList() {
            const { data : res } = await this.$http.get('articles', {
                params: {
                    pagesize: this.queryParam.pagesize,
                    pagenum: this.queryParam.pagenum,
                    title: this.queryParam.title,
                },
            })
            if (res.status != 200) return this.$message.error(res.msg)
            this.artlist = res.data
            this.pagination.total = res.total
        },
        // 获取分类列表
        async getCateList() {
            const { data : res } = await this.$http.get('categories')
            if (res.status != 200) return this.$message.error(res.msg)
            this.catelist = res.data
            this.catetotal = res.total
        },
        // 获取分类统计
        async getCateStat() {
            const { data : res } = await this.$http.get('catestat')
            if (res.status != 200) return this.$message.error(res.msg)
            this.catestat = res.data
        },
        // 根据Cid获取分类名
        getCateByCid(cid) {
            console.log(cid)
            let cate = this.catelist.filter(item => {
                return item.id === cid
            })
            console.log(cate)
            if (cate.length > 0){
                return cate[0].name
            } else {
                return '其他'
            }
        }
    },

    created() {
        this.getArtList()
        this.getCateList()
        this.getCateStat()
    },
  }
</script>

<style scoped>
    .indexRow {
        flex: 1 0 auto;
        display: flex;
        justify-content: center;
        flex-wrap: wrap;
        align-content: flex-start;
        margin: 40px 0;
    }
    .introCol {
        display: flex;
        flex-direction: column;
        margin: 0 10px;
    }
    .infoCol {
        display: flex;
        flex-direction: column;
        margin: 0 10px;
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
    .ant-list-item {
        padding: 0 0;
        margin-bottom: 20px;
    }
    .ant-tag {
        margin-bottom: 10px;
    }

    a {  
        color : black ; 
        transition: 0.3 s; 
    }
    a:hover {  color : rgb(30, 180, 255) ; }
</style>