var express = require('express');
var config = require('config');
var swig = require('swig');
var app = express();
var host = config.get('host');
var port = config.get('port');

app.engine('html',swig.renderFile);

app.set('views engine','html');
app.set('views', __dirname+'/views');

app.get('/index.html', function (req, res) {
	res.render('/usr/src/client/app/index.html',{
		path : config.get('hostWebpack')
	});
});

var server = app.listen(port, function () {
  console.log(' server listening at http://%s:%s', host, port);
});
