import React from 'react';
import TargetStore from '../stores/TargetStore';
import TargetAction from '../actions/TargetActions';

class TargetCountDown extends React.Component {
	constructor(props){
		super(props);
		this.state = TargetStore.getNumberTargetsDown();
	}
	componentDidMount() {
		TargetAction.getNumberTargetsDown();
		TargetStore.addChangeListener(this.onChange.bind(this));
	}

	componentWillUnmount() {
		TargetStore.removeChangeListener(this.onChange.bind(this));
	}

	onChange() {
		this.setState(TargetStore.getNumberTargetsDown());
	}

    render() {
        return (	        
			<div>
	        	<p>{this.state.number_targets_down}</p>
	        </div>
	        );
    }
}

module.exports = TargetCountDown;
