import React from 'react';
import RemoveStore from '../stores/TargetListStore';
import RemoveTargets from '../actions/AddAndRemoveTargetActions';
import Target from 'griddle-react';

var RemoveTarget = React.createClass({

  getInitialState() {
    return RemoveStore.getState();
  },

  componentDidMount() {
    RemoveStore.listen(this.onChange);
  },

  componentWillUnmount() {
    RemoveStore.unlisten(this.onChange);
  },

  onChange() {
    this.setState(this.getInitialState());
  },

  onClickRemoveTarget(){
    RemoveTargets.removeTarget(document.getElementById('targetNameForRemove').value);
  },

  render(){
    return (
      <div>
          <input id="targetNameForRemove" type="text" />
          <button type="submit" onClick={this.onClickRemoveTarget}>
            Remove a target to ping
          </button>
      </div>

    )
  }
});

module.exports = RemoveTarget;
