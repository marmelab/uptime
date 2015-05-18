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

	  	this.actions.setLoading(true);
	  	this.actions.setError(false);

	  	$.ajax({
	  		url: "http://localhost:8383/ips/results",
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
}

module.exports = alt.createActions(TargetActions);

