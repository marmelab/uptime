import alt from '../alt';
import TargetActions from '../actions/TargetActions';

class TargetStore {
  constructor() {
    this.targets = [
		{
			"id":"1",
			"destination":'google.fr',
			"date":"15/05/15",
			"status":"good",
			"time":240
		},
		{
			"id":"2",
			"destination":'rhfjfj',
			"date":"15/05/15",
			"status":"failed",
			"time":240
		}
		];
	}
}

module.exports = alt.createStore(TargetStore, 'TargetStore');
