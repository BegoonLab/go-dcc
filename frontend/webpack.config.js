const path = require('path');
const { VueLoaderPlugin } = require('vue-loader')

module.exports = {
    output: {
        filename: '[name].bundle.js',
        path: path.resolve('../build', 'dist'),
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
    plugins: [
        new VueLoaderPlugin()
    ]
  }