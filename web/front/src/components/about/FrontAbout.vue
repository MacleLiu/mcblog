<template>
    <a-card :hoverable="true" v-title data-title="关于">
        <div style="width: 100%; margin-bottom: 20px;">
            <span style="font-size: 25px; margin-right: 10px;"><ali-icon type="icon-about" /></span>
            <h1 style="display: inline;">关于</h1>
        </div>
        <p style="font-size: 16px; color: black;">
            你好，我是Macle。<br>
            正在计算机知识的海洋里漂泊。<br>
            喜欢骑行。<br>
            喜欢火锅。<br>
            喜欢天蓝色。<br>
            ......^_^<br>
            坚持学习，不断提升。
        </p>
        <hr style="width: 100%;" color="skyblue">
        <h2>关于本站</h2>
        <p style="font-size: 16px; color: black;">
            这里是我的个人网站，是我完成的第一个完整的网站项目。<br>
            我将会在这里撰写技术博客，写游记（应该会吧）以及发布其他作品（万一以后有呢......）。<br>
            <br>
            为什么做这个网站？<br>
            学习Golang后做的实践项目，所以为此还浅尝了VUE。<br>
            还有就是开发一个属于自己的网站难道不是很酷吗？<br>
            本站使用“Gin + VUE2”完成开发。<br>
        </p>
        <hr style="width: 100%;" color="skyblue">
        <h2>心愿单</h2>
        <!-- 愿望列表 -->
        <div style="width: 100%;">
        <a-list item-layout="vertical" :bordered="false" :split="true" :pagination="pagination" :data-source="wishlist">
            <a-list-item slot="renderItem" slot-scope="item, index">
                <a-icon :class="item.status ? 'skyblue' : 'red'" style="margin-right: 10px;" type="heart" />
                <span :class="item.status ? 'realize' : 'inProgress'">{{ item.name }}</span>
            </a-list-item>
        </a-list>
        </div>
    </a-card>
</template>

<script>
export default {
    data() {
        return {
            wishlist: [],
            queryParam: {
                pagesize: 10,
                pagenum: 1,
                title: '',
            },
            pagination: {
                pageSize: 10,
                total: 0,
                showSizeChanger: false,
                hideOnSinglePage: true,
                showTotal: (total) => `共${ total }条`,
                onChange: (page) => {
                    this.queryParam.pagenum = page
                    this.getWishList()
                }
            },
        }
    },
    methods: {
        // 查询工具列表
        async getWishList() {
            const { data : res } = await this.$http.get('wishes', {
                params: {
                    pagesize: this.queryParam.pagesize,
                    pagenum: this.queryParam.pagenum,
                },
            })
            if (res.status != 200) return this.$message.error(res.msg)
            this.wishlist = res.data
            this.pagination.total = res.total
        },
    },
    created() {
        this.getWishList()
    }
}
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
    /* align-items: center; */
}
.ant-list-item {
    padding: 0 0;
    margin-bottom: 20px;
}
.skyblue {
    color: skyblue;
}
.red {
    color: red;
}
.inProgress {
    font-size: 15px; 
    font-weight: bold;
}
.realize {
    font-size: 15px; 
    font-weight: bold;
    text-decoration: line-through;
    text-decoration-thickness: 2px;
}
</style>