module.export = {
	entry: './src/main.js',
	output: {
		filename: 'bundle.js'
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
