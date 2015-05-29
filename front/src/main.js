import React from 'react';
import ResultDurationChart from './components/ResultDurationChart.js'
import TargetListPage from "./components/TargetListPage.js";
import TargetsCounter from "./components/TargetsCounter"
require('./style.scss');

React.render(<TargetListPage />, document.getElementById('content'));
React.render(<TargetsCounter />, document.getElementById('targets_counter'));
React.render(<ResultDurationChart />, document.getElementById('results_chart'));
