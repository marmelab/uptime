var React = require('react');


var data =""; 
$.get("http://api:8383/ips/results",function(d){
	 data=d;
	 console.log("je suis la");
});
console.log("my data",data);
React.render(<Griddle results={data} />, document.getElementById('content'));


