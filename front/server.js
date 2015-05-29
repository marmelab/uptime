var express = require('express');
var config = require('config');
var swig = require('swig');
var app = express();

app.engine('html', swig.renderFile);
app.set('views engine', 'html');
app.set('views', __dirname + '/views');

app.get('/', function (req, res) {
	res.render('layout.html', {
		cdnBaseUrl : config.get('hostWebpack')
	});
});

app.use(express.static('img'));

var server = app.listen(config.get('port'), function () {
  console.log(' server listening at http://' + config.get('host') + ':' + config.get('port'));
});
