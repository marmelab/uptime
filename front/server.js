var express = require('express');
var config = require('config');
var swig = require('swig');
var app = express();

app.engine('html', swig.renderFile);
app.set('views engine', 'html');
app.set('views', __dirname + '/views');

app.get('/index.html', function (req, res) {
	res.render('/usr/src/client/app/index.html', {
		cdnBaseUrl : config.get('hostWebpack')
	});
});

app.use(express.static('img'));

var server = app.listen(config.get('port'), function () {
  console.log(' server listening at http://' + config.get('host') + ':' + config.get('port'));
});
