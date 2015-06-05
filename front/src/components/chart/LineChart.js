import React from 'react';

var d3 = require('d3');
d3.legend = require('d3-legend');

class LineChart extends React.Component {
	constructor(props) {
		super(props);
		this.state = { results: [] };

		/** @TODO: use constants */
		this.margins = {
			top: 10,
			left: 100,
			right: 20,
			bottom: 30
		};

		this.width = 600;
		this.height = 400;
	}

	/** @TODO: move state.results into props.results */
	componentDidUpdate() {
		if (!this.state.results) {
			return;
		}
		this.drawChart();
		return this;
	}

	componentWillUnmount() {
		var svg = React.findDOMNode(this);
		d3.select(svg).remove();
		return this;
	}

	render() {
		return (
			<svg width={this.width} height={this.height}></svg>
		);
	}

	drawChart() {
		var colors = d3.scale.category10();

		var xScale = this._xScale();
		var yScale = this._yScale();

		var xAxis = d3.svg.axis().scale(xScale);
		var yAxis = d3.svg.axis().scale(yScale).orient('left');

		var svg = React.findDOMNode(this);
		this._drawAxes(svg, [xAxis, yAxis]);
		var line = d3.svg.line()
			.x((d, i) => xScale(i))
			.y(d => yScale(d.Time));

		var targets_id = this._getTargetId(this.state.results);
		for (var j = 0; j < targets_id.length; j++) {
			var data = [];
			for (var i = 0; i < this.state.results.length; i++) {
				if((this.state.results[i].Target_id == targets_id[j]) && (this.state.results[i].Status == "good") && (this.state.results[i].Time != -1)) {
					data.push(this.state.results[i]);
				}
			};
			d3.select(svg)
				.append('path')
				.attr('d', line(data))
				.attr('stroke', (d, i) => colors(targets_id[j]))
				.attr('fill', 'none');
		};
	}


	_xScale() {
		return d3.scale.linear()
			.domain([0, 100])
			.range([this.margins.left, this.width - this.margins.right]);
	}

	_yScale() {
		return d3.scale.linear()
			.domain([1000000, 0])
			.range([this.margins.top, this.height - this.margins.top - this.margins.bottom]);
	}

	_drawAxes(svg, axes) {
		d3.select(svg)
			.append('g')
			.attr('class', 'x axis')
			.attr('transform', `translate(0,${this.height - this.margins.bottom})`)
			.call(axes[0]);
		d3.select(svg)
			.append('g')
			.attr('class', 'y axis')
			.attr('transform', `translate(${this.margins.left},${this.margins.top})`)
			.call(axes[1])
			.append("text")
			.attr("transform", "rotate(0)")
			.text("Time (µsec)");
	}

	_getTargetId(data) {
		var target_id = [];
		for (var i = 0; i < data.length; i++) {
			if(-1 == $.inArray(data[i].Target_id,target_id)) {
				target_id.push(data[i].Target_id);
			}
		};
		return target_id;
	}
}

export default LineChart;
