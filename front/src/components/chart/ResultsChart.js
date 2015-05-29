// d3Chart.js

var d3Chart = {};

d3Chart.create = function(el, props, state) {
  var svg = d3.select(el).append('svg')
      .attr('class', 'd3')
      .attr('width', props.width)
      .attr('height', props.height);

  svg.append('g')
      .attr('class', 'd3-points');

  this.update(el, state);
};

d3Chart.update = function(el, state) {
  // initialiser/mettre a jour les données du chart ?
};

d3Chart.destroy = function(el) {
  // vider le chart ?
};

