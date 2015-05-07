var webpack = require('webpack');
var path = require('path');
module.exports = {
	entry: {
		App: [
			'webpack-dev-server/client?http://localhost:8282',
			'webpack/hot/only-dev-server',
			'./src/main.js'
		]
	},

	output: {
		path: __dirname,
		filename: "app/bundle.js",
		publicPath: "http://localhost:8282/"
	},

	plugins: [
		new webpack.HotModuleReplacementPlugin(),
		new webpack.NoErrorsPlugin()
	],

	module:{
		loaders: [
			{ test: /\.js$/, loaders:['jsx-loader']},
			{ test: /\.jsx?$/, loaders:['react-hot'], exclude: /node_modules/}
		]
	}
};
