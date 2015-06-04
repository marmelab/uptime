import React from 'react';

class LineChart extends React.Component {
	constructor(props) {
		super(props);
		this.state = { results: [] };

		/** @TODO: use constants */
		this.margins = {
			top: 10,
			left: 50,
			right: 20,
			bottom: 30
		};
		this.width = 400;
		this.height = 300.
	}

	/** @TODO: move state.results into props.results */
	componentDidUpdate() {
		if (!this.state.results) {
			return;
		}

        this.drawChart();

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
		console.log(this.state.results);
		var line = d3.svg.line()
			.x((d, i) => xScale(i))
			.y(d => yScale(d.Time));
		d3.select(svg)
			.append('path')
			.attr('d', line(this.state.results))
			.attr('stroke', (d, i) => colors(i))
			.attr('fill', 'none');
	}

	_xScale() {
		return d3.scale.linear()
			.domain([0, 14])
			.range([this.margins.left, this.width - this.margins.right]);
	}

	_yScale() {
		return d3.scale.linear()
			.domain([10000, 0])
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
			.call(axes[1]);
	}
}

export default LineChart;
