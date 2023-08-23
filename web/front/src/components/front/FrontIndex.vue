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
                                <h3>99</h3>
                            </a-col>
                            <a-col :span="8">
                                <h3>99</h3>
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
                                width: 60%; 
                                display: flex;
                                justify-content: space-around;"
                        >
                            <ali-icon style="font-size: 20px;" type="icon-GitHub" />
                            <ali-icon style="font-size: 20px;" type="icon-bilibili" />
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
                        <a-list item-layout="vertical" :bordered="false" :split="false" :data-source="data">
                            <a-list-item slot="renderItem" slot-scope="item, index">
                                <div style="width: 100%; font-size: 16px;">
                                    <a href="" >{{ item.title }}</a>
                                </div>
                                <div style="width: 100%;">
                                    <ali-icon type="icon-date1"/>
                                    <span>2023-08-24</span>
                                    <span style="margin: 0 10px;">|</span>
                                    <ali-icon type="icon-folderopen"/>
                                    <span>分类</span>
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
                        <a-list item-layout="vertical" :bordered="false" :split="false" :data-source="cate">
                            <a-list-item slot="renderItem" slot-scope="item, index">
                                <a href="" >
                                    <div style="width: 100%; font-size: 16px; display: flex; justify-content: space-between;">
                                        <span>{{ item.name }}</span>
                                        <span>{{ item.total }}</span>
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
                    <a-list item-layout="horizontal" :bordered="false" :split="false" :data-source="data">
                        <a-list-item slot="renderItem" slot-scope="item, index">
                            <a-card :hoverable="true">
                                <div style="width: 100%; font-size: 24px;">
                                    <a href="" >{{ item.title }}</a>
                                </div>
                                <div style="width: 100%;">
                                    <ali-icon type="icon-date1"/>
                                    <span>2023-08-24</span>
                                    <span style="margin: 0 10px;">|</span>
                                    <ali-icon type="icon-folderopen"/>
                                    <span>分类</span>
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
const data = [
  {
    title: 'Ant Design Title 1',
  },
  {
    title: 'Ant Design Title 2',
  },
  {
    title: 'Ant Design Title 3',
  },
  {
    title: 'Ant Design Title 4',
  },
];

const cate = [
  {
    name: '前端',
    total: 99,
  },
  {
    name: '后端',
    total: 99,
  },
  {
    name: '安全',
    total: 99,
  },
  {
    name: '随笔',
    total: 99,
  },
];

export default {
    data() {
        return {
            data,
            cate,
            size: 'middle',
        }
    },
    methods: {
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