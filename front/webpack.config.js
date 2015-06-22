var webpack = require('webpack');
var ExtractTextPlugin = require("extract-text-webpack-plugin");
module.exports = {
	entry: {
		'uptime-admin.js': [
			'webpack-dev-server/client?http://localhost:8080',
			'webpack/hot/only-dev-server',
			'./src/conf.js'
		],
		'index.html': [
			'./index.html'
		]
	},
	output: {
		filename: "app/[name]",
		publicPath: "http://localhost:8080/"
	},
	plugins: [
		new webpack.ProvidePlugin({
			$: "jquery",
			jQuery: "jquery",
			"window.jQuery": "jquery",
			React: "react"
		}),
		new webpack.DefinePlugin({
			"API_BASE_URL": "'http://localhost:8383'"
		}),
		new ExtractTextPlugin('styles.css', {
			allChunks: true
		})
	],

	module: {
		loaders: [
			{ test: /\.html$/, loaders: ['html'] },
			{ test: /node_modules\/react-admin\/.*js$/, loaders: ['babel'] },
			{ test: /\.js$/, loaders: ['react-hot', 'jsx-loader', 'babel'], exclude: /node_modules/ },
			{ test: /\.scss$/, loader: ExtractTextPlugin.extract('css!sass') },
			{ test: /\.(png|woff|woff2|eot|ttf|svg)$/, loader: 'file' }
		]
	},
    node: {
        fs: 'empty'
    }
};
