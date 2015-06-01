import React from 'react';
import TargetStore from '../stores/TargetStore';
import TargetAction from '../actions/TargetActions';

class TargetsCounter extends React.Component {
	constructor(props){
		super(props);
		this.state = TargetStore.getTargets();
	}

	componentDidMount() {
		TargetAction.fetchTargets();
		TargetStore.addChangeListener(this.onChange.bind(this));
	}

	componentWillUnmount() {
		TargetStore.removeChangeListener(this.onChange.bind(this));
	}

	onChange() {
		this.setState(TargetStore.getTargets());
	}

	getNumberTargetsUp() {
		var up =0;
		for (var i = this.state.targets.length - 1; i >= 0; i--) {
			if(this.state.targets[i].status == true) {
				up += 1; 
			}
		};
		return up;
	}

	getNumberTargetsDown() {
		var down =0;
		for (var i = this.state.targets.length - 1; i >= 0; i--) {
			if(this.state.targets[i].status == false) {
				down += 1; 
			}
		};
		return down;
	}

    render() {
        return (	        
			<div className="navbar-counters">
	       		<div className="counter up">Up: {this.getNumberTargetsUp()}</div>
				<div className="counter down">Down: {this.getNumberTargetsDown()}</div>
	        </div>
	        );
    }
}

module.exports = TargetsCounter;
