import React from 'react';
import AddStore from '../stores/TargetListStore';
import AddTargetAction from '../actions/AddAndRemoveTargetActions';
import Target from 'griddle-react';

var AddTarget = React.createClass({

  getInitialState() {
    return AddStore.getState();
  },

  componentDidMount() {
    AddStore.listen(this.onChange);
  },

  componentWillUnmount() {
    AddStore.unlisten(this.onChange);
  },

  onChange() {
    this.setState(this.getInitialState());
  },

  onClickAddTarget(){
    AddTargetAction.addTarget(document.getElementById('targetName').value);
  },

  render(){
    return (
      <div>
          <input id="targetName" type="text" />
          <button type="submit" onClick={this.onClickAddTarget}>
            Add a target to ping
          </button>
      </div>

    )
  }
});

module.exports = AddTarget;
