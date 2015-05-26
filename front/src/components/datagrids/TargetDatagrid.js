import React from 'react';
import StatusLed from '../StatusLed';
import TargetDataRow from './TargetDataRow';

class TargetDatagrid extends React.Component {
    buildRow(targetsLength) {
        var rows = [];
        if (targetsLength != 0) {
            for (var i = 0; i < targetsLength; i++) {
                 rows.push(<TargetDataRow target = { this.props.targets[i]}/>);
            }

            return rows;           
        }
    }

    render() {
        return (
            <table >
                <tr>
                    <th>Id</th>
                    <th>Destination</th>
                    <th>Satus</th>
                </tr>
                {this.buildRow(this.props.targets.length)}
            </table>
        );
    }
}

export default TargetDatagrid;
