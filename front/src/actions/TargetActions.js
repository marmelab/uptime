import Dispatcher from '../dispatcher/Dispatcher';

module.exports = {

	fetchTargets: function() {
		Dispatcher.dispatch({
			actionType: "TARGET:FETCH:LOADING"
		});
		var url = API_BASE_URL + "/ips/";
		fetch(url, {
			method: 'get'
		})
		.then(function(targets) {
			if(targets.status == 200) {
				return targets.json()
			}
		})
		.then(function(data) {
			Dispatcher.dispatch({
				actionType: "TARGET:FETCH:SUCCESS",
				content: data
			});
			this.getNumberTargetsUp();
			this.getNumberTargetsDown();
		}.bind(this))
		.catch(function(error) {
			Dispatcher.dispatch({
				actionType: "TARGET:FETCH:ERROR"
			});
		})
	},
	getNumberTargetsUp: function() {
		Dispatcher.dispatch({
			actionType: "TARGET:GET:NUMBER_TARGETS_UP"
		});
	},
	getNumberTargetsDown: function() {
		Dispatcher.dispatch({
			actionType: "TARGET:GET:NUMBER_TARGETS_DOWN"
		});
	}
};


