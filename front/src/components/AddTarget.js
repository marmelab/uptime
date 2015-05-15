import React from 'react';
import TargetStore from '../stores/TargetStore';
import TargetAction from '../actions/TargetActions';
import Target from 'griddle-react';

var TargetGriddle = React.createClass({

  getInitialState() {
    return TargetStore.getState();
  },

  componentDidMount() {
    TargetAction.showResults();
    TargetStore.listen(this.onChange);
  },

  componentWillUnmount() {
    TargetStore.unlisten(this.onChange);
  },

  onChange() {
    this.setState(this.getInitialState());
  },

  render(){
    if(this.state.results_loading){
      return  <img src="loading51.gif" alt="loading" />
    }
    if(!this.state.results_error){
      return <Target results={this.state.results} />
    }
    if(this.state.results_errors){
      return <h1>Error can not get results </h1>
    }
  }
});

module.exports = TargetGriddle;
