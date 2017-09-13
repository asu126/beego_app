var htmlWebpackPlugin = require('html-webpack-plugin');
var path = require('path');
var ROOT_PATH = path.resolve(__dirname, '..');

module.exports = {
  devtool: 'eval-source-map',
  context: ROOT_PATH,
  // entry:
  entry:  {
    index:   "./static/js/index.js",
    main:    "./static/js/main.js"//已多次提及的唯一入口文件
  },
  output: {
    path: path.join(ROOT_PATH, 'public'),//打包后的文件存放的地方
    publicPath: '//localhost:8081/public',
    filename: '[name].bundle.js'//打包后输出文件的文件名
  },

  // vue
  resolve: {
    alias: {
      'vue$': 'vue/dist/vue.esm.js' // 'vue/dist/vue.common.js' for webpack 1
    }
  },

  module: {
        loaders: [
            {
                test: /\.vue$/,
                loader: 'vue-loader'
            },
            {
                test: /\.js$/,
                loader: 'babel-loader',
                query:{
                  presets: 'es2015'  
                },
                exclude: /node_modules/
            }
        ]
  },
  // plugins: [
  //       new htmlWebpackPlugin({
  //           template: __dirname + "/views/vue.html",
  //           hash: true
  //       })
  // ],

  devServer: {
    // contentBase: "./public",//本地服务器所加载的页面所在的目录
    historyApiFallback: true,//不跳转
    port:8081,
    inline: true, //实时刷新
    hot: true,
    headers: { 'Access-Control-Allow-Origin': '*' }
  } 
}
