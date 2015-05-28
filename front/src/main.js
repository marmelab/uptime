import React from 'react';
import TargetListPage from "./components/TargetListPage.js";
import TargetCountUp from "./components/TargetCountUp.js";
import TargetCountDown from "./components/TargetCountDown.js";
require('./style.scss');

React.render(<TargetListPage />, document.getElementById('content'));
React.render(<TargetCountUp />, document.getElementById('right-count-up'));
React.render(<TargetCountDown />, document.getElementById('right-count-down'));
