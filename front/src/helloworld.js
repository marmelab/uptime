var React = require('react');
var Griddle = require('griddle-react');

var data = [
	{
		"id":"1",
		"destination":'google.fr',
		"date":"15/05/15",
		"status":"good",
		"time":240
	},
	{
		"id":"2",
		"destination":'rhfjfj',
		"date":"15/05/15",
		"status":"failed",
		"time":240
	}
];

$.get("http://api:8383/ips/results",function(d){
	 data=d;
	 console.log("je suis la");
});

var HelloWorld = React.createClass({
  render: function() {
    return (
        <Griddle results={data} />
    );
  }
});

module.exports = HelloWorld;
 
