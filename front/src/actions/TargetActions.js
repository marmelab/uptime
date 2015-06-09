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
		}.bind(this))
		.catch(function(error) {
			Dispatcher.dispatch({
				actionType: "TARGET:FETCH:ERROR"
			});
		});
	},
	fetchResults: function() {
		Dispatcher.dispatch({
			actionType: "RESULTS:FETCH:LOADING"
		});
		var url = API_BASE_URL + "/ips/results";
		fetch(url, {
			method: 'get'
		})
		.then(function(results) {
			if(results.status == 200) {
				return results.json()
			}
		})
		.then(function(data) {
			Dispatcher.dispatch({
				actionType: "RESULTS:FETCH:SUCCESS",
				content: data
			});
		}.bind(this))
		.catch(function(error) {
			Dispatcher.dispatch({
				actionType: "RESULTS:FETCH:ERROR"
			});
		})
	},
};


