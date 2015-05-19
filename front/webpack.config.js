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
	}),
	new webpack.DefinePlugin({
		"API_BASE_URL": "new String('http://localhost:8383')"
	})
	],
	
	module:{
		loaders: [
		{ test: /\.js$/, loaders: ['react-hot', 'jsx-loader', 'babel-loader'], exclude: /node_modules/ }
		]
	}
};
