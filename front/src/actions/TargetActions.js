import Dispatcher from '../dispatcher/Dispatcher'

export default = {

	showResults() {
		Dispatcher.dispatch("SHOW_RESULTS");
		var url = API_BASE_URL + "/ips/"
		this.actions.setLoading(true);
		this.actions.setError(false);

		fetch(url, {
			method: 'get'
		})
		.then(function(targets) {
			if(targets.status == 200) {
				return targets.json()
			}
		})
		.then(function(data) {
			this.actions.setLoading(false);
			this.actions.setResults(data);
			this.actions.setError(false);
		}.bind(this))
		.catch(function(error) {
			this.actions.setError(true);
		})
	}
};


