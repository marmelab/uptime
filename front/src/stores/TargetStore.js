import TargetActions from '../actions/TargetActions';

class TargetStore	{
	constructor()	{
		this.targets = [];
		this.targets_error = false;
		this.targets_loading = false;

	}

	setResults(response)	{
		this.setState({ targets: response });
	}
  
	setLoading(loading)	{
		this.targets_loading = loading;
	}

	setError(error)	{
		this.targets_error = error;
	}
}

module.exports TargetStore;
