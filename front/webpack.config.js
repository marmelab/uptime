var webpack = require('webpack');
module.exports = {
	entry: {
		App: [
			'webpack-dev-server/client?http://localhost:8080',
			'webpack/hot/only-dev-server',
			'./src/main.js'
		]
	},
	output: {
		filename: "app/bundle.js",
		publicPath: "http://localhost:8080/"
	},

	plugins: [
		new webpack.ProvidePlugin({
			$: "jquery",
			jQuery: "jquery",
			"window.jQuery": "jquery"
		})
	],
	
	module:{
		loaders: [
			{ test: /\.js$/, loaders: ['react-hot', 'jsx-loader', 'babel-loader'], exclude: /node_modules/ }
		]
	}
};
