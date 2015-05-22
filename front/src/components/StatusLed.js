import React from 'react';

class StatusLed extends React.Component {
    render() {
        // @TODO: refactor this VERY UGLY code
        var styles = {
            borderRadius: 50,
            width: 20,
            height: 20,
            backgroundColor: 'red'
        };
        if (this.props.status) {
            styles.backgroundColor = 'green';
        }

        return <div style={styles}></div>
    }
}

export default StatusLed;
