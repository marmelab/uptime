import React from 'react';
import TargetStore from '../stores/TargetStore';
import TargetAction from '../actions/TargetActions';
import TargetDatagrid from './datagrids/TargetDatagrid';

class TargetListPage extends React.Component{
	constructor(props){
		super(props);
		this.state = TargetStore.getState();
	}

	componentDidMount() {
		TargetAction.showResults();
		TargetStore.addChangeListener();
	}

	componentWillUnmount() {
		TargetStore.removeCHangeListener();
	}

	onChange() {
		this.setState(TargetStore.getState());
	}

	render() {
		if(this.state.getTargetsLoading){
			return	<img src="../loading.gif" alt="loading" />
		}

		if(this.state.getTargetsError){
			return <h1>Error: no target found.</h1>
		}

		return <TargetDatagrid targets={this.state.getTargets} />;
	}
}

module.exports = TargetListPage;
