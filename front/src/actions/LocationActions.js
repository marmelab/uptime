var alt = require('../alt');

class LocationActions {
  updateData(data) {
    this.dispatch(data);
  }
}

module.exports = alt.createActions(LocationActions);
