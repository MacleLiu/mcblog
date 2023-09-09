<template>
    <div>
        <Editor :init="init" v-model="content"></Editor>
    </div>
</template>

<script>
    import Editor from '@tinymce/tinymce-vue'
    import './tinymce.min.js'
    import './icons/default/icons.min.js'
    import './themes/silver/theme.min.js'
    import './models/dom/model.min.js'
    import './langs/zh-Hans'

    // 注册插件
    import './plugins/preview/plugin.min.js'
    import './plugins/code/plugin.min.js'
    import './plugins/codesample/plugin.min.js'
    import './plugins/wordcount/plugin.min.js'
    import './plugins/image/plugin.min.js'
    import './plugins/link/plugin.min.js'
    import './plugins/anchor/plugin.min.js'

    export default {
        components: { Editor },
        props: {
            value: {
                type: String,
                default: '',
            },
        },
        data() {
            return {
                init: {
                    language: 'zh-Hans',
                    height: 800,
                    plugins: 'preview code codesample wordcount image link anchor',
                    codesample_languages: [
                        { text: 'HTML/XML', value: 'markup' },
                        { text: 'JavaScript', value: 'javascript' },
                        { text: 'CSS', value: 'css' },
                        { text: 'PHP', value: 'php' },
                        { text: 'Ruby', value: 'ruby' },
                        { text: 'Python', value: 'python' },
                        { text: 'Java', value: 'java' },
                        { text: 'C', value: 'c' },
                        { text: 'C#', value: 'csharp' },
                        { text: 'C++', value: 'cpp' },
                        { text: 'SQL', value: 'sql' },
                        { text: 'Go', value: 'go' },
                        { text: 'Shell', value: 'shell' },
                        { text: 'ProtoBuf', value: 'protobuf' },
                    ],
                    toolbar1: 'undo redo | styles bold italic | alignleft aligncenter alignright alignjustify | outdent indent | preview code',
                    toolbar2: 'paste copy cut | codesample image | link anchor',
                    // 上传图片
                    images_upload_handler: (blobInfo, progress) => new Promise(async (resolve, reject) => {
                        let formData = new FormData()
                        formData.append('file', blobInfo.blob(), blobInfo.filename())
                        const { data : res } = await this.$http.post('upload', formData)
                        if (res.status !== 200) {
                            reject({ message: '图片上传失败：' + res.msg, remove: true })
                            return
                        }
                        resolve(res.url)
                    }),
                },
                content: this.value,
            }
        },
        watch: {
            value(newVal) {
                this.content = newVal
            },
            content(newVal) {
                this.$emit('input', newVal)
            },
        },
    }
</script>

<style>
    /* @import './skins/ui/oxide/skin.min.css'; */
</style>