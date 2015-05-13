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
	  	var that=this;
	  	$.ajax({
	  		url: "http://localhost:8383/ips/results",
	  		success: function(data){
	  			that.actions.setResults(data);
	  			that.actions.setLoading(false);
	  			that.actions.setError(false);
	  		},
	  		error: function(error){
	  			that.actions.setLoading(false);
	  			that.actions.setError(true);
	  		}
	  	});
	  }
}

module.exports = alt.createActions(TargetActions);

