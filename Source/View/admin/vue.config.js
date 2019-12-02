module.exports = {
    publicPath: './',//使用./后就不用修改那个index里面的内容了,也可以改成CDN地址
    outputDir: 'dist/admin',//build输出地址
    assetsDir: './assets/',
    lintOnSave: false,//启用eslint检查,这个鬼东西检查太严格了,关掉
    runtimeCompiler: false,//运行时编译,有性能损耗,去掉
    transpileDependencies: [],//babel-loader 默认会跳过 node_modules 依赖
    productionSourceMap: false,//生产环境是否构建生成source map
    css: {
        // 将组件内的 CSS 提取到一个单独的 CSS 文件 (只用在生产环境中)
        // 也可以是一个传递给 `extract-text-webpack-plugin` 的选项对象
        extract: false,

        // 是否开启 CSS source map
        sourceMap: false
    },
    devServer: { // 开发配置
        port: 8881, // 端口号
        disableHostCheck: true, // 禁用域名检查
        noInfo: true, // 启用 noInfo 后，诸如「启动时和每次保存之后，那些显示的 webpack 包(bundle)信息」的消息将被隐藏。错误和警告仍然会显示。
        hot: true, // 热更新
        compress: true, // 一切服务都启用 gzip 压缩
        inline: true,
        // 代理api服务
        proxy: {
            '/api/*': {
                target: 'http://localhost:8080',
                changeOrigin: true,
                secure: false,
                pathRewrite: {
                    // '^/api': '' // 重写路径：“/api” > "/"
                },
                headers: {

                }
            }
        }
    },
    // 在生产环境下为 Babel 和 TypeScript 使用 `thread-loader`
    // 在多核机器下会默认开启。
    parallel: require('os').cpus().length > 1,
};
