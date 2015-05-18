import React from 'react';
import TargetStore from '../stores/TargetStore';
import TargetAction from '../actions/TargetActions';
import Target from 'griddle-react';

class TargetDatagrid extends React.Component{
  constructor(props){
    super(props);
    this.state = TargetStore.getState();
  }

  componentDidMount() {
    TargetAction.showResults();
    TargetStore.listen(this.onChange.bind(this));
  }

  componentWillUnmount() {
    TargetStore.unlisten(this.onChange.bind(this));
  }

  onChange() {
    this.setState(TargetStore.getState());
  }

  render(){
    if(this.state.targets_loading){
      return  <img src="loading.gif" alt="loading" />
    }

    if(this.state.targets_errors){
      return <h1>Error can not get targets </h1>
    }

    return <Target results={this.state.targets} />
  }
}

module.exports = TargetDatagrid
