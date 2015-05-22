import React from 'react';
import StatusLed from '../StatusLed';

class TargetDataRow extends React.Component {
    render() {

    	if(this.props.target != undefined) {
    		var dataCell = this.props.target;   	
        return	<tr>
        			<td>{dataCell['id']}</td>
        			<td>{dataCell['destination']}</td>
        			<td><StatusLed status={dataCell['status']}/></td>
        		</tr>
        }
        else{
        	return <div>Loadingsdf</div>
        }
    }
}

export default TargetDataRow;
