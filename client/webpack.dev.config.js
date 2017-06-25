const path = require('path');
const webpack = require('webpack');
const HtmlWebpackPlugin = require('html-webpack-plugin');

const config = {
  entry: {
    main: [
      'react-hot-loader/patch',
      // activate HMR for react

      'webpack-dev-server/client?http://localhost:8000',
      // bundle the client for webpack-dev-server
      // and connect to the provided endpoint

      'webpack/hot/only-dev-server',
      // bundle the client for hot reloading
      // only- means to only hot reload for successful updates

      './src/index.js'
      // entry point of the app
    ]
  },

  output: {
    filename: '[name].js',
    // Not using [chunkhash] in dev as it increases compilation time

    path: path.resolve(__dirname, 'dist'),

    publicPath: ''
    // necessary for HMR to know where to load the hot update chunks
  },

  devtool: 'inline-source-map',

  module: {
    rules: [
      {
        test: /\.js$/,
        use: ['babel-loader'],
        exclude: /node_modules/
      }, {
        test: /\.(css|scss)$/,
        use: [{
          loader: 'style-loader'
          // creates style nodes from JS strings
        }, {
          loader: 'css-loader'
          // translates CSS into CommonJS
        }, {
          loader: 'sass-loader'
          // compiles Sass to CSS
        }, {
          loader: 'postcss-loader'
        }]
      }, {
        test: /\.(jpg|png|svg|eot|ttf|woff|woff2)$/,
        use: {
          loader: 'url-loader',
          options: {
            limit: 25000,
            name: '[name].[ext]'
          }
        }
      }
    ]
  },

  plugins: [
    new webpack.HotModuleReplacementPlugin(),
    // enable HMR globally

    new webpack.NamedModulesPlugin(),
    // prints more readable module names in the browser console on HMR updates

    new HtmlWebpackPlugin({ template: 'index.html' })
    // generates 'index.html' and handles 'script' and 'link' tags
  ],

  devServer: {
    host: '0.0.0.0',
    disableHostCheck: true,

    historyApiFallback: true,
    // respond to 404s with index.html

    hot: true,
    // enable HMR on the server

    headers: {
      'Access-Control-Allow-Origin': '*'
    },

    port: 8000
  }
};

module.exports = config;
