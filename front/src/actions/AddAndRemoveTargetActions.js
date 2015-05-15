import alt from'../alt';

class AddAndRemoveTargetActions {
	constructor(){
		this.generateActions(
		'setLoading',
		'setError',
		'setTargetList'
		);		
	}

  	addTarget(target) {
  		var newTarget = {
  			Destination:target
  		};
  		newTarget = JSON.stringify(newTarget);
  		console.log(newTarget);
  		this.dispatch("ADD_TARGET");
	  	this.actions.setLoading(true);
	  	this.actions.setError(false);
	  	var that=this;
		  	$.ajax({
	  		url: "http://localhost:8383/ips/",
	  		type: "POST",
	  		data: newTarget,
	  		success: function(data){
	  			that.actions.setTargetList(target);
	  			that.actions.setLoading(false);
	  			that.actions.setError(false);
	  		},
	  		error: function(error){
	  			that.actions.setLoading(false);
	  			that.actions.setError(true);
	  		}
	  	});
	  }
  	removeTarget(target) {
   		var newTarget = {
  			Destination:target
  		};
  		newTarget = JSON.stringify(newTarget);
  		console.log(newTarget);
  		this.dispatch("REMOVE_TARGET");
	  	this.actions.setLoading(true);
	  	this.actions.setError(false);
	  	var that=this;
		  	$.ajax({
	  		url: "http://localhost:8383/ips",
	  		type: "DELETE",
	  		data: newTarget,
	  		success: function(data){
	  			that.actions.setTargetList(target);
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

module.exports = alt.createActions(AddAndRemoveTargetActions);

