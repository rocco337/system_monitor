var HtmlWebpackInlineSourcePlugin = require('webpack-inline-modern-source-plugin');
var HtmlWebpackPlugin = require('html-webpack-plugin');

module.exports = {
  configureWebpack: {
    // No need for splitting
    optimization: {
      splitChunks: false
    },
    resolve: {
      alias: {
        'vue$': 'vue/dist/vue.esm.js' // 'vue/dist/vue.common.js' for webpack 1
      }
    },
    plugins: [

    //   new HtmlWebpackPlugin({
    //     template: 'public/index.html', 
    //     inlineSource: '.(js|css|woff)$' // embed all javascript and css inline
    //   }),
    //   new HtmlWebpackInlineSourcePlugin()
    ]  
  },
  pluginOptions: {
    quasar: {
      importStrategy: 'kebab',
      rtlSupport: false
    }
  },
  transpileDependencies: [
    'quasar'
  ],
  
}
