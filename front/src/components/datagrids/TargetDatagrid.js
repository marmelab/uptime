import Griddle from 'griddle-react';
import React from 'react';
import StatusLed from '../StatusLed';

class TargetDatagrid extends React.Component {
    render() {
        var metadata = [
            { columnName: 'id', order: 0, displayName: '#' },
            { columnName: 'destination', order: 1, displayName: 'Destination' },
            { columnName: 'status', order: 2, displayName: 'Status', customComponent: StatusLed }
        ];
var s = this.props.targets;
console.log(s);
        return <Griddle results={s} columnMetadata={metadata}></Griddle>
    }
}

export default TargetDatagrid;
