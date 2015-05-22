import alt from'../alt';

class TargetActions {
	constructor(){
		this.generateActions(
			'setLoading',
			'setError',
			'setResults'
			);		
	}

	showResults() {
		this.dispatch("SHOW_RESULTS");
		var url = API_BASE_URL + "/ips/"
		this.actions.setLoading(true);
		this.actions.setError(false);

		fetch(url, {
			method: 'get'
		})
		.then(function(response) {
			if(response.status == 200) {
				return response.json()
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
}

module.exports = alt.createActions(TargetActions);

