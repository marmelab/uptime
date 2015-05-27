import assign from 'object-assign';
import {EventEmitter} from 'events';
import Dispatcher from '../dispatcher/Dispatcher';
import TargetActions from '../actions/TargetActions';

var CHANGE_EVENT = 'change';
var targets = [];
var targets_error = false;
var targets_loading = true; 

var TargetStore = assign({}, EventEmitter.prototype, {
	getAll: function() {
		var data = {targets: targets, targets_loading: targets_loading, targets_error: targets_error}
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
		case "TARGET:FETCH:LOADING":
			targets_loading = true;
			TargetStore.emitChange();
			break;

		case "TARGET:FETCH:SUCCESS":
			targets = action.content;
			targets_error = false;
			targets_loading = false;
			TargetStore.emitChange();
			break;

		case "TARGET:FETCH:ERROR":
			targets_error = true;
			TargetStore.emitChange();
			break;

	}
});

module.exports = TargetStore;


