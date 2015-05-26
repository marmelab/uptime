import React from 'react';
import TargetStore from '../stores/TargetStore';
import TargetAction from '../actions/TargetActions';
import TargetDatagrid from './datagrids/TargetDatagrid';

class TargetListPage extends React.Component{
	constructor(props){
		super(props);
		this.state = TargetStore.getAll();
	}
	componentDidMount() {
		TargetAction.showTargetsData();
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
			return	<img src="../loading.gif" alt="loading" />
		}

		if (this.state.targets_error){
			return <h1>Error: no target found.</h1>
		}

		return <TargetDatagrid targets={this.state.targets} />;
	}
}

module.exports = TargetListPage;
