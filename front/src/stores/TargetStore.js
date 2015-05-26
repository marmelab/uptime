import assign from 'object-assign';
import {EventEmitter} from 'events';
import Dispatcher from '../dispatcher/TargetDispatcher';
import TargetActions from '../actions/TargetActions';

var CHANGE_EVENT = 'change';

var targets = [];
var targets_error = false;
var targets_loading = true;

function setTargets(data) {
	targets = data;
}

function setTargetsError(error) {
	targets_error = error;
}

function setTargetsLoading(loading) {
	targets_loading = loading; 
}

var TargetStore = assign({}, EventEmitter.prototype, {
	getAll: function() {
		var data = {targets,targets_loading,targets_error}
		return data;
	},
	getTargets: function() {
		return targets;
	},
	getTargetsError: function() {
		return targets_error;
	},
	getTargetsLoading: function() {
		return targets_loading;
	},
	emitChange: function() {
		this.emit(CHANGE_EVENT);
	},
	addChangeListener: function(callback) {
		this.on(CHANGE_EVENT,callback);
	},
	removeChangeListener: function(callback) {
		this.removeListener(CHANGE_EVENT,callback);		
	}
});

Dispatcher.register(function(action) {
	switch(action.actionType) {
		case "TARGET_DATA_LOADING":
			setTargetsLoading(true);
			TargetStore.emitChange();
			break;

		case "TARGET_DATA_LOADED":
			setTargets(action.content);
			setTargetsError(false);
			setTargetsLoading(false);
			TargetStore.emitChange();
			break;

		case "TARGET_DATA_ERROR":
			setTargetsError(true);
			TargetStore.emitChange();
			break;

	}
});

module.exports = TargetStore;


