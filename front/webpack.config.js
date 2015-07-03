var webpack = require('webpack');
var ExtractTextPlugin = require("extract-text-webpack-plugin");
module.exports = {
	entry: {
		'uptime-admin': [
			'webpack-dev-server/client?http://localhost:8080',
			'webpack/hot/only-dev-server',
			'./src/conf.js',
			'react-admin/build/react-admin-standalone.min.css'
		],
		'index.html': [
			'./index.html'
		]
	},
	output: {
        filename: 'build/[name].min.js',
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
			{ test: /\.js$/, loaders: ['react-hot', 'jsx-loader', 'babel'], exclude: /node_modules/ },
			{ test: /\.scss$/, loader: ExtractTextPlugin.extract('css!sass') },
			{ test: /\.css$/, loader: ExtractTextPlugin.extract('css') },
			{ test: /\.(png|woff|woff2|eot|ttf|svg)$/, loader: 'file' },
		]
	},
    node: {
        fs: 'empty'
    }
};
