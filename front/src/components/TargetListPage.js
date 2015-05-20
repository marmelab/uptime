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
		TargetStore.listen(this.onChange.bind(this));
	}

	componentWillUnmount() {
		TargetStore.unlisten(this.onChange.bind(this));
	}

	onChange() {
		this.setState(TargetStore.getState());
	}

	render() {
		if(this.state.targets_loading){
			return	<img src="../loading.gif" alt="loading" />
		}

		if(this.state.targets_errors){
			return <h1>Error: no target found.</h1>
		}

		return <TargetDatagrid targets={this.state.targets} />;
	}
}

module.exports = TargetListPage;
