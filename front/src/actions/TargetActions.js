import alt from'../alt';

class TargetActions {
  getGriddle(griddle) {
    this.dispatch("GET_GRIDDLE");
  }
}

module.exports = alt.createActions(TargetActions);
