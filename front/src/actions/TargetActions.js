import alt from'../alt';

class TargetActions {
  getGriddle(targets) {
    this.dispatch("GET_GRIDDLE");
  }
}

module.exports = alt.createActions(TargetActions);
