const path = require('path');
const {VueLoaderPlugin} = require('vue-loader');
const webpack = require('webpack');
const {VuetifyPlugin} = require('webpack-plugin-vuetify')
const pkg = require('./package.json')

// Plugins
let webpackPlugins = [
    new VueLoaderPlugin(),
    new webpack.HotModuleReplacementPlugin(),
    new webpack.NoEmitOnErrorsPlugin(),
    new VuetifyPlugin({autoImport: true}), // Enabled by default
    new webpack.EnvironmentPlugin(['NODE_ENV']),
    new webpack.DefinePlugin({
        __VERSION__: JSON.stringify(pkg.version),
        __BUILT__: JSON.stringify(new Date()),
        __AUTHOR__: JSON.stringify(pkg.author)
    })
];
// Entry points
let webpackEntryPoints = [
    './src/index.js',
];

if (process.env.NODE_ENV === 'production' || process.env.NODE_ENV === 'demo') {
    webpackPlugins = [
        new VueLoaderPlugin(),
        new VuetifyPlugin({autoImport: true}), // Enabled by default
        new webpack.EnvironmentPlugin(['NODE_ENV']),
        new webpack.DefinePlugin({
            __VERSION__: JSON.stringify(pkg.version),
            __BUILT__: JSON.stringify(new Date()),
            __AUTHOR__: JSON.stringify(pkg.author)
        })
    ];
}

if (process.env.NODE_ENV === 'development') {
    // Development
    webpackEntryPoints.push('webpack-hot-middleware/client');
}

module.exports = {
    mode: process.env.NODE_ENV === 'development' ? 'development' : 'production',
    /**
     * Source Maps
     */

    // https://webpack.js.org/configuration/devtool/#production
    devtool: 'source-map',
    entry: webpackEntryPoints,
    devServer: {
        static: {
            directory: path.join(__dirname, '../build'),
        },
        port: 8080,
        hot: true
    },
    output: {
        filename: process.env.NODE_ENV === 'demo' ? 'demo.[name].bundle.js' : '[name].bundle.js',
        path: path.resolve('../build'),
    },
    optimization: {
        minimize: process.env.NODE_ENV !== 'demo',
    },
    module: {
        rules: [
            {
                test: /\.js$/, use: 'babel-loader'
            },
            {
                test: /\.s(c|a)ss$/,
                use: [
                    'vue-style-loader',
                    'css-loader',
                    {
                        loader: 'sass-loader',
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
                test: /\.mjs$/,
                include: /node_modules/,
                type: "javascript/auto"
            },
            {
                test: /\.vue$/,
                loader: 'vue-loader'
            },
            {
                test: /\.proto$/,
                use: {
                    loader: 'protobufjs-loader',
                    options: {
                        /* Import paths provided to pbjs.
                         *
                         * default: webpack import paths (i.e. config.resolve.modules)
                         */
                        paths: ['./software/frontend/src/pb'],

                        /* Additional command line arguments passed to pbjs.
                         *
                         * default: []
                         */
                        pbjsArgs: ['--no-encode'],

                        /* Enable Typescript declaration file generation via pbts.
                         *
                         * Declaration files will be written every time the loader runs.
                         * They'll be saved in the same directory as the protobuf file
                         * being processed, with a `.d.ts` extension.
                         *
                         * This only works if you're using the 'static-module' target
                         * for pbjs (i.e. the default target).
                         *
                         * The value here can be a config object or a boolean; set it to
                         * true to enable pbts with default configuration.
                         *
                         * default: false
                         */
                        pbts: {
                            /* Additional command line arguments passed to pbts.
                             */
                            args: ['--no-comments'],
                        },

                        /* Set the "target" flag to pbjs.
                         *
                         * default: 'static-module'
                         */
                        target: 'json-module',
                    },
                },
            },
        ],
    },
    plugins: webpackPlugins
}