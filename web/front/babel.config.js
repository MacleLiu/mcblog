module.exports = {
  presets: [
    '@vue/cli-plugin-babel/preset'
  ],
  plugins: [
    [
      "import",
      { libraryName: "ant-design-vue", libraryDirectory: "es", style: 'css' },
    ],
    [
      "prismjs",
      {
        "languages": ["javascript", "css", "html", "go", "shell", "sql", "c", "java","csharp","cpp","python","ruby","php","protobuf"],//语言
        "plugins": ["line-numbers"],//行号插件（可以不用，还有其他插件网上查询）
        "theme": "okaidia",  //主题coy|twillight|tomorrow|solarizedlight|okaidia|funky|dark
        "css": true
      }
    ],
  ]
}