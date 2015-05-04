
module.exports = {
	entry: {
		App: [
			'webpack-dev-server/client?http://localhost:8080',
			'./src/main.js'
		]
	},
	output: {
		filename: "app/bundle.js",
		publicPath: "http://localhost:8080/"
	},
	module:{
		loaders: [
			{ test: /\.js$/,loaders:[
				"jsx-loader"
				] 
			}
		]
	}
};
