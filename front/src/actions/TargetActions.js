import alt from'../alt';

class TargetActions {
	constructor(){
		this.generateActions(
			'setLoading',
			'setError',
			'setResults',
			'setTargetsList'
			);		
	}

	showResults() {
		this.dispatch("SHOW_RESULTS");

		this.actions.setLoading(true);
		this.actions.setError(false);

		$.ajax({
			url: API_BASE_URL + "/ips/",
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
		});
	}

	showTargetsList() {
		this.dispatch("SHOW_TARGETS_LIST");

		this.actions.setLoading(true);
		this.actions.setError(false);

		$.ajax({
			url: API_BASE_URL + "/ips/",
			complete: function() {
				this.actions.setLoading(false);
			}.bind(this),
			success: function(data){
				console.log(data);
				this.actions.setResults(data);
				this.actions.setError(false);
			}.bind(this),
			error: function(error){
				this.actions.setError(true);
			}.bind(this)
		});
	}
}

module.exports = alt.createActions(TargetActions);

