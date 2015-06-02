import React from 'react';

class LineChart extends React.Component {
	constructor(props) {
		super(props);
		this.chart = this.initChart();
	}

	/** @TODO: move state.results into props.results */
	componentDidUpdate() {
		if (!this.state.results) {
			return;
		}
	}

	render() {
		return (
			<svg width="400" height="300"></svg>
		);
	}

	initChart() {
		var xScale = d3.time.scale.utc()
			.domain(d3.extent(this.state.results, function(d) { return d.created_at; }));
			console.log('ok');
	}
}

export default LineChart;
