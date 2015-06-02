// d3Chart.js

/**
d3Chart.create = function(el, props, state) {
};

d3Chart.update = function(el, state) {


var xScale = d3.scale.linear().range([50, 950]).domain([state.results[0].Created_at, state.results[1].Created_at]); 
var yScale = d3.scale.linear().range([480, 20]).domain([-1,10000]);  // time

var xAxis = d3.svg.axis()
	.scale(xScale);
svg.append("svg:g")
  .attr("class","axis")
  .attr("transform", "translate(0," + (480) + ")")
  .call(xAxis); 
var yAxis = d3.svg.axis()
	.scale(yScale)
	.orient("left");
svg.append("svg:g")
	.attr("class","axis")
	.attr("transform", "translate(" + (50) + ",0)")
	.call(yAxis);
var lineGen = d3.svg.line()
  .x(function(d) {
	return xScale(d.Created_at); 
  })
  .y(function(d) {
	return yScale(d.Time);
  });
var data = [];
for (var i = 0; i < state.results.length; i++) {
	if(state.results[i].Destination == "google.fr") {
		data.push({Created_at: state.results[i].Created_at,Time: state.results[i].Time});	
	}
};
svg.append('svg:path')
  .attr('d', lineGen(data))
  .attr('stroke', 'green')
  .attr('stroke-width', 2)
  .attr('fill', 'none'); 
};

d3Chart.destroy = function(el) {
  // vider le chart ?
};
**/
export default d3Chart;
