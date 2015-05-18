import alt from '../alt';
import TargetActions from '../actions/TargetActions';

class TargetStore {
  constructor() {
  	this.targets = [];
  	this.targets_error = false;
  	this.targets_loading = false;

    this.bindActions(alt.getActions("TargetActions"));
  }

  setResults(targets) {
    var response = JSON.parse(targets);
    this.setState({ targets: response });
  }
  
  setLoading(loading) {
  	this.targets_loading = loading;
  }

  setError(error) {
  	this.targets_error = error;
  }
}

module.exports = alt.createStore(TargetStore, 'TargetStore');
