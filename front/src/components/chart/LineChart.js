import ResultsChart from './ResultsChart';
import React from 'react';

class LineChart extends React.Component {
  constructor(props){
    super(props);
  }
  componentDidMount() {
    var el = this.getDOMNode();
    ResultsChart.create(el, {
      width: '400px',
      height: '300px'
    }, this.getChartState());
  }
  componentDidUpdate() {
    var el = this.getDOMNode();
    ResultsChart.update(el, this.getChartState());
  }
  getChartState() {
    return this.state.data;
  }
  componentWillUnmount() {
    var el = this.getDOMNode();
    ResultsChart.destroy(el);
  }
  render() {
    return (
      <div className="LineChart"></div>
    );
  }
}

export default LineChart;
