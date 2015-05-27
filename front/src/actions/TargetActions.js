import Dispatcher from '../dispatcher/Dispatcher';

module.exports = {

	fetchTargets : function() {
		Dispatcher.dispatch({
			actionType: "FETCH:TARGET:LOADING"
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
				actionType: "FETCH:TARGET:SUCCESS",
				content: data
			});
		}.bind(this))
		.catch(function(error) {
			Dispatcher.dispatch({
				actionType: "FETCH:TARGET:ERROR"
			});
		})
	}
};


