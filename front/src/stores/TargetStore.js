import alt from '../alt';
import TargetActions from '../actions/TargetActions';

class TargetStore {
  constructor() {
  	this.results = [
  		{
        "Destination": "google.fr",
        "Status": "good",
        "Time": 214,
        "Key": ""
  		}
  	];
  	this.results_error = false;
  	this.results_loading = false;
    this.bindActions(alt.getActions("TargetActions"));
  }
  setResults(data) {
    var response = JSON.parse(data);
    console.log(response)
    this.results = response;
  }
  
  setLoading(loading) {
  	this.results_loading = loading;
  }

  setError(error) {
  	this.results_error = error;
  }
}

module.exports = alt.createStore(TargetStore, 'TargetStore');
