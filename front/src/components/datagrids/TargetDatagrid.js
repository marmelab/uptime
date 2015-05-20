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
        console.log(this.props.targets);
        return <Griddle
            results={this.props.targets}
            showFilter={true}
            columnMetadata={metadata}
        />
    }
}

export default TargetDatagrid;
