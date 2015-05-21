import React from 'react';
import TargetDatagrid from "./components/TargetDatagrid.js";
import TargetListgrid from "./components/TargetListgrid.js";
import AddTarget from "./components/AddTarget.js";
import RemoveTarget from "./components/RemoveTarget.js";

var Main = React.createClass({
	render: function(){
		return(
			<div>
				<h1> Ping Results </h1>
				<AddTarget />
				<RemoveTarget />
				<TargetDatagrid />
				<TargetListgrid />
			</div>
		);
	}
});

React.render(<Main/>, document.getElementById('content'));
