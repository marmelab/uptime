import React from 'react';
import TargetStore from '../stores/TargetStore';
import TargetAction from '../actions/TargetActions';

class TargetsCounter extends React.Component {
	getNumberTargetsUp() {
		var up =0;
		for (var i = this.props.targets.length - 1; i >= 0; i--) {
			if(this.props.targets[i].status == true) {
				up += 1; 
			}
		};
		return up;
	}

	getNumberTargetsDown() {
		var down =0;
		for (var i = this.props.targets.length - 1; i >= 0; i--) {
			if(this.props.targets[i].status == false) {
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

export default TargetsCounter;
