import React from 'react';
import TargetStore from '../stores/TargetStore';
import TargetAction from '../actions/TargetActions';
import Target from 'griddle-react';

var TargetGriddle = React.createClass({
  getInitialState() {
    return TargetStore.getState();
  },

  componentDidMount() {
    TargetAction.getGriddle(this.getInitialState());
    TargetStore.listen(this.onChange);
  },

  componentWillUnmount() {
    TargetStore.unlisten(this.onChange);
  },

  onChange(griddle) {
    this.setState(this.getInitialState());
  },

  render(){
    return <Target results={this.state.griddle} />
  }
});

module.exports = TargetGriddle;
