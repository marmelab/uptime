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
		var url = API_BASE_URL + "/ips/results"
		this.actions.setLoading(true);
		this.actions.setError(false);

		fetch(url)
		.then(function(response) {
			this.actions.setLoading(false);
			this.actions.setResults(response);
			this.actions.setError(false);
		}.bind(this))
		.catch(function() {
			this.actions.setLoading(false);
			this.actions.setError(true);
		}.bind(this));


		/*$.ajax({
			url: API_BASE_URL + "/ips/results",
>>>>>>>  using fecth instead of ajax in TargetActions
			complete: function() {
				this.actions.setLoading(false);
			}.bind(this),
			success: function(data){
				this.actions.setResults(data);
				this.actions.setError(false);
			}.bind(this),
			error: function(error){
				this.actions.setError(true);
			}.bind(this)
		});*/
	}
}

module.exports = alt.createActions(TargetActions);

