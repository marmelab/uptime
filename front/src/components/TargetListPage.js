import React from 'react';
import LineChart from './chart/LineChart';
import TargetsCounter from "./TargetsCounter"
import TargetStore from '../stores/TargetStore';
import TargetAction from '../actions/TargetActions';
import TargetDatagrid from './datagrids/TargetDatagrid';

const margins = {
		top: 10,
		left: 100,
		right: 20,
		bottom: 30,
};
const width = 600;
const height = 400;

class TargetListPage extends React.Component{
	constructor(props){
		super(props);
		this.state = TargetStore.getAll();
	}
	componentDidMount() {
		TargetAction.fetchTargets();
		TargetAction.fetchResults();
		TargetStore.addChangeListener(this.onChange.bind(this));
	}

	componentWillUnmount() {
		TargetStore.removeChangeListener(this.onChange.bind(this));
	}

	onChange() {
		this.setState(TargetStore.getAll());
	}

	render() {
		if (this.state.targets_loading) {
			return <img src="../loading.gif" alt="loading" />
		}
		if (this.state.targets_error){
			return <h1>Error: no target found.</h1>
		}
		return(	
			<div>
				<h1>Global view of targets status :</h1>
				<TargetDatagrid id="content" targets={this.state.targets} />
				<TargetsCounter id="targets_counter" className="nav navbar-nav navbar-right"targets={this.state.targets} />
				<LineChart id="results_chart" results={this.state.results} margins={margins} width={width} height={height}/>
			</div>
		);
	}
}

export default TargetListPage;
