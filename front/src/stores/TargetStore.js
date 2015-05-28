import assign from 'object-assign';
import {EventEmitter} from 'events';
import Dispatcher from '../dispatcher/Dispatcher';
import TargetActions from '../actions/TargetActions';

var CHANGE_EVENT = 'change';
var targets = [];
var targets_error = false;
var targets_loading = true;
var number_targets_up = 0;
var number_targets_down = 0;

var TargetStore = assign({}, EventEmitter.prototype, {
	getAll: function() {
		var data = {targets: targets, targets_loading: targets_loading, targets_error: targets_error}
		return data;
	},
	getNumberTargetsUp: function() {
		return {number_targets_up: number_targets_up};
	},
	getNumberTargetsDown: function() {
		return {number_targets_down: number_targets_down};
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

		case "TARGET:GET:NUMBER_TARGETS_UP":
			var up =0;
			for (var i = targets.length - 1; i >= 0; i--) {
				if(targets[i].status == true) {
					up += 1; 
				}
			};
			number_targets_up = up;
			TargetStore.emitChange();
			break;

		case "TARGET:GET:NUMBER_TARGETS_DOWN":
			var down =0;
			for (var i = targets.length - 1; i >= 0; i--) {
				if(targets[i].status == false) {
					down += 1; 
				}
			};
			number_targets_down = down;
			TargetStore.emitChange();
			break;
	}
});

module.exports = TargetStore;


