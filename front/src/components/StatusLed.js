import React from 'react';

class StatusLed extends React.Component {
    render() {
        // @TODO: refactor this VERY UGLY code
        var styles = {
            borderRadius: 50,
            width: 50,
            height: 50,
            backgroundColor: 'red'
        };
        if (this.props.data) {
            styles.backgroundColor = 'green';
        }

        return <div style={styles}></div>
    }
}

export default StatusLed;
