import React from 'react';
import StatusLed from '../StatusLed';
import TargetDataCell from './TargetDataCell'

class TargetDataRow extends React.Component {
    render() {

    	if(this.props.target != undefined) {
    		var dataCell = this.props.target;   	
        return	<tr>
        			<TargetDataCell id={dataCell['id']}/>
        			<TargetDataCell destination={dataCell['destination']}/>
        			<TargetDataCell status={dataCell['status']}/>
        		</tr>
        }
        else{
        	return <div>Loadingsdf</div>
        }
    }
}

export default TargetDataRow;
