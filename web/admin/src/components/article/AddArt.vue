<template>
    <div>
        <div class="pageTitle">{{ id ? '编辑文章' : '新增文章' }}</div>
        <a-card>
            <a-form-model :model="artInfo" ref="artInfoRef" :rules="artInfoRules">
                <a-form-model-item label="文章标题" prop="title">
                    <a-input style="width: 300px" v-model="artInfo.title"></a-input>
                </a-form-model-item>
                <a-form-model-item label="文章分类">
                    <a-select
                        style="width: 200px"
                        placeholder="请选择分类"
                        v-model="artInfo.cid"
                        @change="cateChange"
                    >
                      <a-select-option v-for="item in catelist" :key="item.id" :value="item.id">{{ item.name }}</a-select-option>
                    </a-select>
                </a-form-model-item>
                <a-form-model-item label="文章描述" prop="desc">
                    <a-input style="width: 400px" type="textarea" v-model="artInfo.desc"></a-input>
                </a-form-model-item>
                <a-form-model-item label="文章缩略图">
                    <a-upload
                        name="file"
                        listType="picture"
                        :defaultFileList="fileList"
                        :action = "upUrl"
                        :headers = "headers"
                        @change = "upChange"
                    >
                        <a-button><a-icon  type="upload"/>点击上传</a-button>
                        <br/>
                        <span v-if="id">
                            <img :src="artInfo.img" style="width: 120px; height: 100px;">
                        </span>
                    </a-upload>
                </a-form-model-item>
                <a-form-model-item label="文章内容" prop="content">
                    <Editor v-model="artInfo.content"></Editor>
                </a-form-model-item>
                <a-form-model-item>
                    <a-button type="primary" @click="submitInfo(artInfo.id)">{{ artInfo.id ? '更新' : '提交' }}</a-button>
                </a-form-model-item>
            </a-form-model>
        </a-card>
    </div>
</template>

<script>
    import { Url } from '../../plugins/http'
    import Editor from '../editor/index'

    export default {
        components: { Editor },
        props: ['id'],
        data() {
            return {
                artInfo :{
                    id :0,
                    title: '',
                    cid: undefined,
                    desc: '',
                    content: '',
                    img: '',
                },
                catelist: [],
                upUrl: Url + '/upload',
                headers: {},
                fileList: [],
                artInfoRules: {
                    title: [{required: true, message: '请输入文章标题', trigger: 'blur'}],
                    desc: [{required: true, message: '请输入文章描述', trigger: 'blur'}, {max: 120, message: '描述内容最多120字符', trigger: "change"}],
                    content: [{required: true, message: '请输入文章内容', trigger: 'blur'}],
                },
            }
        },
        methods: {
            // 查询分类列表
            async getCateList() {
                const { data : res } = await this.$http.get('categories')
                if (res.status != 200) return this.$message.error(res.msg)
                this.catelist = res.data
            },
            // 选择分类
            cateChange(value) {
                this.artInfo.cid = value
            },
            // 查询文章信息
            async getAriInfo(id) {
                const { data : res } = await this.$http.get(`article/${ id }`)
                if (res.status != 200) return this.$message.error(res.msg)
                this.artInfo = res.data
                this.artInfo.id = res.data.ID
            },
            // 上传图片
            upChange(info) {
                if (info.file.status !== 'uploading') {
                }
                if (info.file.status === 'done' && info.file.response.status === 200) {
                    this.$message.success(`上传成功`);
                    this.artInfo.img = info.file.response.url
                } else if (info.file.status === 'error') {
                    this.$message.error(`上传失败`);
                }
            },
            // 提交或更新
            submitInfo(id){
                this.$refs.artInfoRef.validate(async (valid) => {
                    if(valid){
                        if (id === 0) {
                            const { data : res } = await this.$http.post('article/add', this.artInfo)
                            if (res.status !== 200) return this.$message.error(res.msg)
                            this.$router.push('/admin/articles')
                            this.$message.success('添加文章成功')
                        } else {
                            const { data : res } = await this.$http.put(`article/${id}`, this.artInfo)
                            if (res.status !== 200) return this.$message.error(res.msg)
                            this.$router.push('/admin/articles')
                            this.$message.success('更新文章成功')
                        }
                    } else {
                        return this.$message.error('参数不符合要求，请检查')
                    }
                })
            },
        },
        created() {
            this.getCateList()
            this.headers = { Authorization : `Bearer ${ window.sessionStorage.getItem('token') }` }
            if(this.id) {
                this.getAriInfo(this.id)
            }
        },
    }
</script>

<style scoped>
.pageTitle {
        text-align: center;
        font-size: 30px;
        font-weight: bold;
        letter-spacing: 20px;
        color: black;
        font-family: Serif;
    }
</style>