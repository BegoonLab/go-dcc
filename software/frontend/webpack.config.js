const path = require('path');
const { VueLoaderPlugin } = require('vue-loader');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const webpack = require('webpack');

// Plugins
let webpackPlugins = [
  new VueLoaderPlugin(),
  new webpack.HotModuleReplacementPlugin(),
  new webpack.NoEmitOnErrorsPlugin(),
];
// Entry points
let webpackEntryPoints = [
  './src/index.js',
];

if (process.env.NODE_ENV === 'production') {

  webpackPlugins = [
      new VueLoaderPlugin()
  ];

}else{
  // Development
  webpackEntryPoints.push('webpack-hot-middleware/client');
}

module.exports = {
    mode: process.env.NODE_ENV === 'development' ? 'development' : 'production',
    entry: webpackEntryPoints,
    devServer: {
      hot: true
    },
    devServer: {
      static: {
        directory: path.join(__dirname, '../build'),
      },
      port: 8080,
      hot: true
    },
    output: {
        filename: '[name].bundle.js',
        path: path.resolve('../build'),
    },
    module: {
      rules: [
        { test: /\.js$/, use: 'babel-loader' },
        {
          test: /\.s(c|a)ss$/,
          use: [
            'vue-style-loader',
            'css-loader',
            {
              loader: 'sass-loader',
              // Requires sass-loader@^7.0.0
              options: {
                implementation: require('sass'),
                indentedSyntax: true // optional
              },
              // Requires >= sass-loader@^8.0.0
              options: {
                implementation: require('sass'),
                sassOptions: {
                  indentedSyntax: true // optional
                },
              },
            },
          ],
        },
        {
            test: /.css$/,
            use: [
              'vue-style-loader',
              'css-loader',
            ]
        },
        {
            test: /\.vue$/,
            loader: 'vue-loader'
        }
      ],
    },
    plugins: webpackPlugins
  }