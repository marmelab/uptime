import LineChart from './chart/LineChart';
import TargetStore from '../stores/TargetStore';
import TargetAction from '../actions/TargetActions';

class ResultDurationChart extends LineChart {
 	constructor(props) {
		super(props);
	}

	componentDidMount() {
		TargetAction.fetchResults();
		TargetStore.addChangeListener(this.onChange.bind(this));
	}

	componentWillUnmount() {
		super.componentWillUnmount();
		TargetStore.removeChangeListener(this.onChange.bind(this));
	}
	onChange() {
		this.setState(TargetStore.getResults());
	}
}

export default ResultDurationChart;
