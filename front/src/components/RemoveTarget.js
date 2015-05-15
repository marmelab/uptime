import React from 'react';
import RemoveStore from '../stores/TargetListStore';
import AddTarget from '../actions/AddAndRemoveTargetActions.js';
import Target from 'griddle-react';

var RemoveTarget = React.createClass({

  getInitialState() {
    return RemoveStore.getState();
  },

  componentDidMount() {
    AddTarget.addTarget();
    RemoveStore.listen(this.onChange);
  },

  componentWillUnmount() {
    RemoveStore.unlisten(this.onChange);
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

module.exports = RemoveTarget;
