import alt from '../alt';
import AddAndRemoveTargetActions from '../actions/AddAndRemoveTargetActions';

class TargetListStore {
  constructor() {
  	this.targetList = [
      {
        "Destination": ""
      }
    ];
  	this.results_error = false;
  	this.results_loading = true;
    this.bindActions(alt.getActions("AddAndRemoveTargetActions"));
  }
  setTargetList(target) {
    this.targetList = target;
  }
  
  setLoading(loading) {
  	this.results_loading = loading;
  }

  setError(error) {
  	this.results_error = error;
  }
}

module.exports = alt.createStore(TargetListStore, 'TargetListStore');
