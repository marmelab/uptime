// d3Chart.js

var d3Chart = {};

d3Chart.create = function(el, props, state) {

var svg = d3.select(document.getElementById("results_chart"))
    .append("svg")
    .attr("width", 1000)
    .attr("height", 900)
var data = [{
    "sale": "202",
    "year": "2000"
}, {
    "sale": "215",
    "year": "2001"
}, {
    "sale": "179",
    "year": "2002"
}, {
    "sale": "199",
    "year": "2003"
}, {
    "sale": "134",
    "year": "2003"
}, {
    "sale": "176",
    "year": "2010"
}];
var data2 = [{
    "sale": "152",
    "year": "2000"
}, {
    "sale": "189",
    "year": "2002"
}, {
    "sale": "179",
    "year": "2004"
}, {
    "sale": "199",
    "year": "2006"
}, {
    "sale": "134",
    "year": "2008"
}, {
    "sale": "176",
    "year": "2010"
}];
var xScale = d3.scale.linear().range([50, 950]).domain([2000,2010]);
var yScale = d3.scale.linear().range([480, 20]).domain([134,215]);

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
    return xScale(d.year);
  })
  .y(function(d) {
    return yScale(d.sale);
  });
svg.append('svg:path')
  .attr('d', lineGen(data))
  .attr('stroke', 'green')
  .attr('stroke-width', 2)
  .attr('fill', 'none');
svg.append('svg:path')
  .attr('d', lineGen(data2))
  .attr('stroke', 'blue')
  .attr('stroke-width', 2)
  .attr('fill', 'none');
};

d3Chart.update = function(el, state) {  var g = d3.select(el).selectAll('.d3-results');
  var point = g.selectAll('.d3-results').data(state.results);
};

d3Chart.destroy = function(el) {
  // vider le chart ?
};

export default d3Chart;
