import React from 'react';
import TargetListPage from "./components/TargetListPage.js";
import TargetsCounter from "./components/TargetsCounter"
require('./style.scss');

React.render(<TargetListPage />, document.getElementById('content'));
React.render(<TargetsCounter />, document.getElementById('targets_counter'));
