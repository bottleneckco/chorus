const path = require('path');
const webpack = require('webpack');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const CleanWebpackPlugin = require('clean-webpack-plugin');
const ExtractTextPlugin = require('extract-text-webpack-plugin');

const pathsToClean = ['dist'];

const config = {
  entry: {
    main: [
      './src/index.js'
    ],
    // entry point of the app

    vendor: [
      'react',
      'react-dom',
      'react-router-dom',
      'prop-types'
    ]
    // global dependencies
  },

  output: {
    filename: '[name].[chunkhash:8].js',

    path: path.resolve(__dirname, 'dist'),

    publicPath: ''
    // necessary for HMR to know where to load the hot update chunks
  },

  module: {
    rules: [
      {
        test: /\.js$/,
        use: ['babel-loader'],
        exclude: /node_modules/
      }, {
        test: /\.(css|scss)$/,
        use: ExtractTextPlugin.extract({
          fallback: 'style-loader',
          use: ['css-loader', 'sass-loader', 'postcss-loader']
        })
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
    new webpack.optimize.CommonsChunkPlugin({
      names: ['vendor', 'manifest']
    }),
    // split global dependencies into a separate 'vendor' file.
    // 'manifest' file extracts webpack's runtime code from the 'vendor'
    // file. this prevents the 'vendor' file's chunkhash from changing
    // every build.

    new webpack.HashedModuleIdsPlugin(),
    // ensures the 'vendor' file's chunkhash stays the same when code
    // is modified

    new ExtractTextPlugin('styles.[contentHash:8].css'),
    // split css into separate file
    // note: do not use this for dev as it does not work with HMR

    new HtmlWebpackPlugin({ template: 'index.html' }),
    // generates 'index.html' and handles 'script' and 'link' tags

    new CleanWebpackPlugin(pathsToClean, { verbose: true })
    // removes 'dist' folder before building
  ]
};

module.exports = config;
