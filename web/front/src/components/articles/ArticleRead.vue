<template>
    <div class="container" v-title data-title="文章">
        <FrontHeader/>
        <!-- <div style="background-color: rgb(240, 245, 245);"> -->
        <a-row class="readRow">
            <a-col class="introCol" :xs="0" :sm="0" :md="6" :lg="6" :xl="5">
                    <a-space direction="vertical" :size="size" style="width: 100%;">
                        <!-- 摘要卡片 -->
                        <a-card :hoverable="true">
                            <div style="width: 100%; margin-bottom: 20px;">
                                <h2 style="display: inline;">摘要</h2>
                            </div>
                            <!-- 摘要内容 -->
                            <a-spin :spinning="art_loading">
                                <div style="width: 100%;">
                                    <p>{{ artInfo.desc }}</p>
                                </div>
                            </a-spin>
                        </a-card>
                        <!-- 目录卡片 -->
                        <a-affix :offset-top="80">
                        <a-card :hoverable="true">
                            <div style="width: 100%; margin-bottom: 20px;">
                                <h2 style="display: inline;">目录</h2>
                            </div>
                            <!-- 目录 -->
                            <a-spin :spinning="catalog_loading" style="width: 100%;">
                                <a-anchor style="width: 100%;" @click="handleAnchorClick">
                                    <a-anchor-link v-for="item in catalog" :key="item.id" :href="'#'+item.id" :title="item.title">
                                    </a-anchor-link>
                                </a-anchor>
                            </a-spin>
                        </a-card>
                        </a-affix>
                    </a-space>
            </a-col>
            
            <a-col class="infoCol" :xs="24" :sm="24" :md="18" :lg="18" :xl="14">
                <a-space direction="vertical" :size="size">
                    <a-card :hoverable="true">
                        <a-spin :spinning="art_loading">
                            <!-- 文章标题 -->
                            <h1>{{ artInfo.title }}</h1>
                        </a-spin>
                        <hr style="width: 100%;" color="#DDDDDD">
                        <div v-if="!art_loading" style="width: 100%; margin-bottom: 20px;">
                            <ali-icon type="icon-date1"/>
                            <span>发表于{{ artInfo.CreatedAt | dateFormat }}</span>
                            <span style="margin: 0 10px;">|</span>
                            <ali-icon type="icon-date1"/>
                            <span>最后更新于{{ artInfo.UpdatedAt| dateFormat }}</span>
                            <span style="margin: 0 10px;">|</span>
                            <ali-icon type="icon-folderopen"/>
                            <span>{{ artInfo.Category.name }}</span>
                        </div>
                        <!-- 文章内容 -->
                        <div id="articleContent" ref="content" v-html="artInfo.content" style="width: 100%; font-size: 16px; color: black;"></div>
                        <hr style="width: 100%; border: 1px dashed skyblue;">
                        <!-- 文章标签 -->
                        <div style="width: 100%;">
                            <a-spin :spinning="tag_loading" size="small">
                                <a-tag v-for="item in taglist" :key="item.id" color="blue">
                                    {{ item.name }}
                                </a-tag>
                            </a-spin>
                        </div>
                        <hr style="width: 100%;" color="skyblue">
                        <!-- 评论区 -->
                        <div style="width: 100%;">
                            <a-comment>
                                <a-avatar 
                                  size="large"
                                  slot="avatar"
                                  style="color: white; background-color: skyblue; font-size: 25px;"
                                  alt="头像"
                                >
                                    {{ new_comment.name ? new_comment.name.charAt(0) : "U" }}
                                </a-avatar>
                                <!-- 评论框 -->
                                <div slot="content">
                                    <a-form-model :model="new_comment" ref="commentNew" :rules="commentRules">
                                        <a-form-model-item prop="name">
                                            <a-input v-model="new_comment.name" placeholder="请输入昵称" style="width: 200px" />
                                        </a-form-model-item>
                                        <a-form-model-item prop="content">
                                            <a-textarea :rows="4" placeholder="交流一下" v-model="new_comment.content" />
                                        </a-form-model-item>
                                        <a-form-model-item>
                                            <a-button html-type="submit" :loading="submitting" type="primary" @click="handleSubmit">
                                                评论
                                            </a-button>
                                        </a-form-model-item>
                                    </a-form-model>
                                </div>
                            </a-comment>
                            <!-- 评论列表 -->
                            <a-list
                              :data-source="comments"
                              :header="`${total} 条评论`"
                              item-layout="vertical"
                              :locale="comment_defaultText"
                              :loading="comment_loading"
                              style="font-weight: bolder;"
                            >
                              <a-list-item slot="renderItem" slot-scope="item, index">
                                <a-comment
                                  :author="item.name"
                                  :content="item.content"
                                  :datetime="item.CreatedAt | timeFormat"
                                >
                                    <span slot="actions" @click="handleReply(item.ID, item.name)">回复</span>
                                    <a-avatar 
                                      slot="avatar"
                                      style="color: white; background-color: skyblue; font-size: 25px;"
                                      alt="头像"
                                    > 
                                        {{ item.name ? item.name.charAt(0) : "U" }}
                                    </a-avatar>
                                    <!-- 子评论 -->
                                    <a-comment
                                      v-if="item.children.length > 0"
                                      v-for="child in item.children"
                                      :key="child.ID"
                                      :author="child.name === child.toname ? child.name : child.name + ' 回复 ' + child.toname"
                                      :content="child.content"
                                      :datetime="child.CreatedAt | timeFormat"
                                    >
                                        <span slot="actions" @click="handleReply(item.ID, child.name)">回复</span>
                                        <a-avatar 
                                          slot="avatar"
                                          style="color: white; background-color: skyblue; font-size: 25px;"
                                          alt="头像"
                                        > 
                                            {{ child.name ? child.name.charAt(0) : "U" }}
                                        </a-avatar>
                                    </a-comment>
                                </a-comment>
                              </a-list-item>
                            </a-list>
                        </div>
                    </a-card>
                </a-space>
            </a-col>
        </a-row>
        <!-- 回复评论对话框 -->
        <a-modal
            title="新增用户"
            :visible="replyVisible"
            :confirmLoading="replying"
            @ok="replyOk"
            @cancel="replyCancel"
        >
            <a-comment>
                <a-avatar 
                  size="large"
                  slot="avatar"
                  style="color: white; background-color: skyblue; font-size: 25px;"
                  alt="头像"
                >
                    {{ reply_comment.name ? reply_comment.name.charAt(0) : "U" }}
                </a-avatar>
                <!-- 评论框 -->
                <div slot="content">
                    <a-form-model :model="reply_comment" ref="commentRep" :rules="commentRules">
                        <a-form-model-item prop="name">
                            <a-input v-model="reply_comment.name" placeholder="请输入昵称" style="width: 200px" />
                        </a-form-model-item>
                        <a-form-model-item prop="content">
                            <a-textarea :rows="4" placeholder="交流一下" v-model="reply_comment.content" />
                        </a-form-model-item>
                    </a-form-model>
                </div>
            </a-comment>
        </a-modal>
        <!-- </div> -->
        <FrontFooter/>
        <div>
            <a-back-top :visibilityHeight="100"/>
        </div>
    </div>
</template>

<script>
import FrontHeader from '../front/FrontHeader'
import FrontFooter from '../front/FrontFooter'

import Moment from "moment"
import Prism from "prismjs";

export default {
    components: {FrontHeader, FrontFooter},
    props: ['id'],

    data() {
        return {
            art_loading: true,
            tag_loading: true,
            catalog_loading: true,
            comment_loading: true,
            comment_defaultText:{
                emptyText: '快来发布第一条评论吧'
            },
            artInfo :{
                id :0,
                title: '',
                cid: 0,
                Category: {
                    id: 0,
                    name: '',
                },
                desc: '',
                content: '',
                img: '',
            },
            catalog: [],
            taglist: [],
            size: 'middle',
            top: 10,
            comments: [],
            total: 0,
            submitting: false,
            new_comment: {
                name: '',  // 所属用户昵称
                content: '',  // 评论内容
                artid: 0,  // 所属文章id
                toid: 0,  // 一级评论的id
                toname: '',  // 回复目标的昵称
            },
            replying: false,
            reply_comment: {
                name: '',  // 所属用户昵称
                content: '',  // 评论内容
                artid: 0,  // 所属文章id
                toid: 0,  // 一级评论的id
                toname: '',  // 回复目标的昵称
            },
            commentRules: {
                name: [
                    { pattern: /(^\S)((.)*\S)?(\S*$)/, message: '首尾不能有空格' },
                    {required: true, message: '请输入昵称', trigger: 'blur'},
                    {max: 20, message: '昵称最大长度20字符', trigger: "change"},
                ],
                content: [{required: true, message: '请输入评论内容', trigger: 'blur'}],
            },
            replyVisible: false,
        }
    },
    

    methods: {
        // 查询文章信息
        async getArtInfo(id) {
            try{
                const { data : res } = await this.$http.get(`article/${ id }`)
                if (res.status != 200){
                    this.art_loading = false
                    this.$router.push(`/`).catch((err) => err )
                    this.$message.error(res.msg)
                    return
                }
                this.artInfo = res.data
                this.artInfo.id = res.data.ID
                this.art_loading = false
            }catch(err){
                this.art_loading = false
                return
            }
        },
        // 查询文章标签
        async getArtTags(id) {
            try{
                const { data : res } = await this.$http.get(`tags/${ id }`)
                if (res.status != 200){
                    this.tag_loading = false
                    this.$message.error(res.msg)
                    return
                }
                this.taglist = res.data
                this.tag_loading = false
            }catch(err){
                this.tag_loading = false
                return
            }
        },
        // 从本地存储获取用户名
        getUserName() {
            this.new_comment.name = localStorage.getItem('mcblog_username')
            this.reply_comment.name = localStorage.getItem('mcblog_username')
        },
        // 获取文章评论
        async GetComment(id) {
            try{
                const { data : res } = await this.$http.get(`comment/${ id }`)
                if (res.status != 200) {
                    this.comment_loading = false
                    this.comment_defaultText.emptyText = res.msg
                    this.$message.error(res.msg)
                    return 
                }
                this.comments = res.data
                this.total = res.total
                this.comment_loading = false
            }catch(err){
                this.comment_loading = false
                return
            }
        },
        // 提交评论
        handleSubmit() {
            this.$refs.commentNew.validate(async (valid) => {
                if(valid){
                    localStorage.setItem('mcblog_username', this.new_comment.name)

                    this.submitting = true
                    this.new_comment.artid = this.artInfo.id
                    try{
                        const { data : res } = await this.$http.post('comment/add', this.new_comment)
                        if (res.status !== 200) {
                            this.submitting = false
                            this.new_comment.content = ''
                            this.$message.error(res.msg)
                            return
                        }
                        this.GetComment(this.id)
                        this.submitting = false
                        this.new_comment.content = ''
                        this.$message.success('评论成功')
                    }catch(err){
                        this.submitting = false
                        this.$message.error('评论失败')
                        return
                    }
                } else {
                    return this.$message.error('参数不符合要求，请检查')
                }
            })
        },
        // 回复评论
        handleReply(id, name) {
            console.log(id)
            this.replyVisible = true
            this.reply_comment.toid = id
            this.reply_comment.toname = name
        },
        // 回复评论对话框操作
        replyOk() {
            this.$refs.commentRep.validate(async (valid) => {
                if(valid){
                    localStorage.setItem('mcblog_username', this.reply_comment.name)

                    this.replying = true
                    this.reply_comment.artid = this.artInfo.id
                    try{
                        const { data : res } = await this.$http.post('comment/add', this.reply_comment)
                        if (res.status !== 200) {
                            this.replying = false
                            this.reply_comment.toid = 0
                            this.reply_comment.toname = ''
                            this.reply_comment.content = ''
                            this.replyVisible = false
                            this.$message.error(res.msg)
                            return 
                        }
                        this.GetComment(this.id)
                        this.replying = false
                        this.reply_comment.toid = 0
                        this.reply_comment.toname = ''
                        this.reply_comment.content = ''
                        this.replyVisible = false
                        this.$message.success('评论成功')
                    }catch(err){
                        this.replying = false
                        this.replyVisible = false
                        this.$message.error('评论失败')
                        return
                    }
                } else {
                    return this.$message.error('参数不符合要求，请检查')
                }
            })
        },
        replyCancel() {
            this.reply_comment.content = ''
            this.replyVisible = false
        },
        // 生成带锚点的文章目录
        // generateDirectory() {
        //     console.log('generateDirectory')
        //     const article_content = this.$refs.content  // 获取文章内容
        //     article_content.childNodes.forEach((e, index) => {
        //         //具体执行步骤
        //         console.log(e)
        //     })
        // },
        // 阻止点击目录的默认事件修改路由
        handleAnchorClick(e) {
          e.preventDefault()
        },
    },

    // 通过watch监听artInfo以能够成功获取v-html渲染的内容
    watch: {
        artInfo: function () {
            this.$nextTick(() => {
                const article_content = this.$refs.content
                const titleTag = ["H1", "H2", "H3"]
                let titles = []
                article_content.childNodes.forEach((e, index) => {
                    if (titleTag.includes(e.nodeName)) {
                        const id = "header-" + index
                        e.setAttribute("id", id)
                        titles.push({
                            id: id,
                            title: e.innerHTML.replace(/<\/?.+?\/?>/g,''),  // 去除标题值中包含的其他标签
                            level: Number(e.nodeName.substring(1, 2)),
                            nodeName: e.nodeName,
                        })
                    }  
                })
                this.catalog = titles
            })
            this.catalog_loading = false
        }
    },

    created() {
        this.getUserName()
        this.GetComment(this.id)
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
    min-height: 100vh;
    display: flex;
    flex-direction: column;
}
.readRow {
    flex: 1 0 auto;
    display: flex;
    justify-content: center;
    flex-wrap: wrap;
    align-content: flex-start;
    padding: 80px 0 20px 0;
    background-color: rgb(240, 245, 245);
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
:deep .ant-anchor-wrapper {
    max-height: 350px !important;
}
:deep img {
    max-width: 100%;
    height: auto;
}
:deep a {
    white-space: pre-line;
}
:deep #articleContent h1 {
    padding-top: 70px;
    margin-top: -70px;
}
:deep #articleContent h2 {
    padding-top: 70px;
    margin-top: -70px;
}
:deep #articleContent h3 {
    padding-top: 70px;
    margin-top: -70px;
}
</style>