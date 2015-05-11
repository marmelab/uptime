var alt = require('../alt');
var LocationActions = require('../actions/LocationActions');

class LocationStore {
  constructor() {
    this.data = [
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

    this.bindListeners({
      handleUpdateData: LocationActions.UPDATE_DATA
    });
  }

  handleUpdateLocations(data) {
    this.data = data;
  }
}

module.exports = alt.createStore(LocationStore, 'LocationStore');
