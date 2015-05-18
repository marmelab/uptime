import React from 'react';
import TargetStore from '../stores/TargetStore';
import TargetAction from '../actions/TargetActions';
import Target from 'griddle-react';

class TargetDatagrid extends React.Component {
  constructor(props){
    super(props);
    this.state =  { targets: TargetStore.getState() };
  }

  componentDidMount() {
    TargetAction.getGriddle(this.state.targets);
    TargetStore.listen(this.onChange);
  }

  componentWillUnmount() {
    TargetStore.unlisten(this.onChange);
  }

  onChange(targets) {
    this.setState(this.state.targets);
  }

  render(){
    return <Target results={this.state.targets.targets} />
  }
}

module.exports = TargetDatagrid
