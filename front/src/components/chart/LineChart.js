import React from 'react';
import d3 from 'd3';

class LineChart extends React.Component {
	constructor(props) {
		super(props);
		this.props = { results: [] };
	}

	componentDidUpdate() {
		if (!this.props.results) {
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
			<svg width={this.props.margins['width']} height={this.props.margins['height']}></svg>
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

		var targets_id = this._getTargetsId(this.props.results);
		for (var currentTargetIndex = 0; currentTargetIndex < targets_id.length; currentTargetIndex++) {
			var data = [];
			for (var i = 0; i < this.props.results.length; i++) {
				if((this.props.results[i].Target_id == targets_id[currentTargetIndex]) && (this.props.results[i].Status == "good") && (this.props.results[i].Time != -1)) {
					data.push(this.props.results[i]);
				}
			};
			d3.select(svg)
				.append('path')
				.attr('d', line(data))
				.attr('stroke', (d, i) => colors(targets_id[currentTargetIndex]))
				.attr('fill', 'none');
		};
	}


	_xScale() {
		return d3.scale.linear()
			.domain([0, 100])
			.range([this.props.margins['left'], this.props.margins['width'] - this.props.margins['right']]);
	}

	_yScale() {
		return d3.scale.linear()
			.domain([1000000, 0])
			.range([this.props.margins['top'], this.props.margins['height'] - this.props.margins['top'] - this.props.margins['bottom']]);
	}

	_drawAxes(svg, axes) {
		d3.select(svg)
			.append('g')
			.attr('class', 'x axis')
			.attr('transform', `translate(0,${this.props.margins['height'] - this.props.margins['bottom']})`)
			.call(axes[0]);
		d3.select(svg)
			.append('g')
			.attr('class', 'y axis')
			.attr('transform', `translate(${this.props.margins['left']},${this.props.margins['top']})`)
			.call(axes[1])
			.append("text")
			.attr("transform", "rotate(0)")
			.text("Time (µsec)");
	}

	_getTargetsId(data) {
		var targets_id = [];
		for (var i = 0; i < data.length; i++) {
			if(-1 == $.inArray(data[i].Target_id,targets_id)) {
				targets_id.push(data[i].Target_id);
			}
		};
		return targets_id;
	}
}

export default LineChart;
