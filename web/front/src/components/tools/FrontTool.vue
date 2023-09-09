<template>
    <a-card :hoverable="true" v-title data-title="工具">
        <div style="width: 100%; margin-bottom: 20px;">
            <span style="font-size: 25px; margin-right: 10px;"><a-icon type="tool" /></span>
            <h1 style="display: inline;">工具</h1>
        </div>       
        <!-- 列表 -->
        <div style="width: 100%;">
        <a-list item-layout="vertical" :bordered="false" :split="true" :pagination="pagination" :data-source="toollist">
            <a-list-item slot="renderItem" slot-scope="item, index">
                <div @click="openTool(item.url)">
                    <span class="toolName">{{ item.name }}</span>
                </div>
            </a-list-item>
        </a-list>
        </div>
    </a-card>
</template>

<script>
export default {
    data() {
        return {
            toollist: [],
            queryParam: {
                pagesize: 16,
                pagenum: 1,
                title: '',
            },
            pagination: {
                pageSize: 16,
                total: 0,
                showSizeChanger: false,
                hideOnSinglePage: true,
                showTotal: (total) => `共${ total }条`,
                onChange: (page) => {
                    this.queryParam.pagenum = page
                    this.getToolList()
                }
            },
        }
    },
    methods: {
        // 查询工具列表
        async getToolList() {
            const { data : res } = await this.$http.get('tools', {
                params: {
                    pagesize: this.queryParam.pagesize,
                    pagenum: this.queryParam.pagenum,
                },
            })
            if (res.status != 200) return this.$message.error(res.msg)
            this.toollist = res.data
            this.pagination.total = res.total
        },
        openTool(url) {
            window.open(url, '_blank')
        },
    },
    created() {
        this.getToolList()
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
    align-items: center;
}
.ant-list-item {
    padding: 0 0;
    margin-bottom: 20px;
}
.toolName {
    font-size: 16px; 
    font-weight: bold;
}
.toolName:hover {
    color: rgb(30, 180, 255);
}
</style>