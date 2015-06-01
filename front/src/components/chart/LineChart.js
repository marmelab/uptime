import ResultsChart from './ResultsChart';
import React from 'react';

class LineChart extends React.Component {
  constructor(props){
    super(props);
  }
  componentDidMount() {
    var el = React.findDOMNode();
    ResultsChart.create(el, {
      width: '400px',
      height: '300px'
    }, this.getChartState());
  }
  componentDidUpdate() {
    var el = React.findDOMNode(this);
    ResultsChart.update(el, this.getChartState());
  }
  getChartState() {
    return this.state;
  }
  componentWillUnmount() {
    var el = React.findDOMNode();
    ResultsChart.destroy(el);
  }
  render() {
    return (
      <div className="LineChart"></div>
    );
  }
}

export default LineChart;
