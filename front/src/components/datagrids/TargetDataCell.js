import React from 'react';
import StatusLed from '../StatusLed'

class TargetDataCell extends React.Component {
  render() {   	
      if(this.props.destination != undefined) {	
        return <td>{this.props.destination}</td>
      }
      if(this.props.id != undefined) {
        return <td>{this.props.id}</td>   		
      }
      if(this.props.status != undefined) {
        return	<td>
                  <StatusLed status={this.props.status}/> 
                </td>   		
      }
      else {
        return <div></div>
      }
    }
}

export default TargetDataCell;
