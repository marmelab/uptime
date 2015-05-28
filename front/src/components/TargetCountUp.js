import React from 'react';
import TargetStore from '../stores/TargetStore';
import TargetAction from '../actions/TargetActions';

class TargetCountUp extends React.Component {
	constructor(props){
		super(props);
		this.state = TargetStore.getNumberTargetsUp();
	}
	componentDidMount() {
		TargetAction.getNumberTargetsUp();
		TargetStore.addChangeListener(this.onChange.bind(this));
	}

	componentWillUnmount() {
		TargetStore.removeChangeListener(this.onChange.bind(this));
	}

	onChange() {
		this.setState(TargetStore.getNumberTargetsUp());
	}

    render() {
        return (	        
			<div>
	        	<p>{this.state.number_targets_up}</p>
	        </div>
	        );
    }
}

module.exports = TargetCountUp;
