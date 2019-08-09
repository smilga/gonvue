const path = require('path');
const webpack = require('webpack');
const { VueLoaderPlugin } = require('vue-loader');

const isProduction = process.env.NODE_ENV === 'production';

module.exports = {
    mode: isProduction ? 'production' : 'development',
    entry: ['./src/app.js'],
    output: {
        path: path.resolve(__dirname, './dist'),
        filename: 'app.js',
        publicPath: isProduction ? '/static/' : 'http://localhost:8080/',
    },
    plugins: [
        new VueLoaderPlugin(),
    ],
    resolve: {
        extensions: ['.js', '.vue'],
        alias: {
            vue$: 'vue/dist/vue.esm.js',
            '@': path.resolve(__dirname, 'src'),
        }
    },
    module: {
        rules: [
            {
                test: /\.vue$/,
                use: [{ loader: 'vue-loader' }],
            },
            {
                test: /.js$/,
                loaders: ['babel-loader'],
                exclude: [/node_modules\/(?![vue\-color])/],
            },
            {
                test: /\.(png|jpg|gif|svg)$/,
                loader: 'file-loader',
                options: {
                    name: 'img/[name].[ext]?[hash]',
                },
            },
            {
                test: /\.(eot|ttf|woff|woff2)$/,
                loader: 'file-loader',
            },
            {
                test: /\.scss$/,
                use: ['vue-style-loader', 'css-loader', 'sass-loader'],
            },
            {
                test: /\.css$/,
                use: ['vue-style-loader', 'css-loader'],
            },
        ],
    },
    devServer: {
        headers: {
            "Access-Control-Allow-Origin": "*",
            "Access-Control-Allow-Methods": "GET, POST, PUT, DELETE, PATCH, OPTIONS",
            "Access-Control-Allow-Headers": "X-Requested-With, content-type, Authorization"
        },
        hot: true,
        sockPort: 8080,
    },
};
