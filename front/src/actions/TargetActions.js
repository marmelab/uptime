import Dispatcher from '../dispatcher/TargetDispatcher';

module.exports = {

	showTargetsData : function() {
		Dispatcher.dispatch({
			actionType: "TARGET_DATA_LOADING"
		});
		var url = API_BASE_URL + "/ips/"
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
				actionType: "TARGET_DATA_LOADED",
				content: data
			});
		}.bind(this))
		.catch(function(error) {
			Dispatcher.dispatch({
				actionType: "TARGET_DATA_ERROR"
			});
		})
	}
};


