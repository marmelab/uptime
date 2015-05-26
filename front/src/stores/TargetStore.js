import assign from 'object-assign';
import {EventEmitter} from 'events';
import {Dispatcher} from 'flux';
import TargetActions from '../actions/TargetActions';

var CHANGE_EVENT = 'change';

var targets = [];
var targets_error = false;
var targets_loading = false;

function setTargets(data)	{
	this.targets = data;
}

function setTargetsError(error)	{
	this.targets_error = error;
}

function setTargetsLoading(loading)	{
	this.targets_loading = loading; 
}

var TargetStore = assign({}, EventEmitter.prototype, {
	getTargets: function()	{
		return targets;
	},
	getTargetsError: function()	{
		return targets_error;
	},
	getTargetsLoading: function()	{
		return targets_loading;
	},
	addChangeListener: function(callback)	{
		this.on(CHANGE_EVENT,callback);
	},
	removeChangeListener: function(callback)	{
		this.removeListener(CHANGE_EVENT,callback);		
	}
});

Dispatcher.register(function(action)	{
	switch(action.ActionType)	{
		case TARGET_DATA_LOADING:
			setTargetsLoading(true);
			TargetStore.emitChange();
			break;

		case TARGET_DATA_LOADED:
			setTargets(action.content.trim());
			TargetStore.emitChange();
			break;

		case TARGET_DATA_ERROR:
			setTargetsError(true);
			TargetStore.emitChange();
			break;

	}
});

module.exports = TargetStore;


