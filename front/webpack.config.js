var webpack = require('webpack');
var path = require('path');
module.exports = {
	entry: {
		App: [
			'webpack-dev-server/client?http://localhost:8181',
			'webpack/hot/only-dev-server',
			'./src/main.js'
		]
	},

	output: {
		filename: "app/bundle.js",
		publicPath: "http://localhost:8181/"
	},

	plugins: [
		new webpack.HotModuleReplacementPlugin(),
		new webpack.NoErrorsPlugin()
	],

	module:{
		loaders: [
			{ test: /\.js$/, loaders: ['react-hot', 'jsx-loader'], exclude: /node_modules/ }
		]
	}
};
