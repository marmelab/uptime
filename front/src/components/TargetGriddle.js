import React from 'react';
import TargetStore from '../stores/TargetStore';
import TargetAction from '../actions/TargetActions';
import Target from 'griddle-react';

class TargetDatagrid extends React.Component {
  constructor(props){
    super(props);
    this.state =  { data: TargetStore.getState() };
  }

  componentDidMount() {
    TargetStore.listen(this.onChange);
  }

  componentWillUnmount() {
    TargetStore.unlisten(this.onChange);
  }

  onChange(targets) {
    this.setState(this.state.targets);
  }

  render(){
    return <Target results={this.state.data.targets} />
  }
}

module.exports = TargetDatagrid
