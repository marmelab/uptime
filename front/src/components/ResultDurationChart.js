import LineChart from './chart/LineChart';
import TargetStore from '../stores/TargetStore';
import TargetAction from '../actions/TargetActions';

class ResultDurationChart extends LineChart {
  constructor(props) {
    super(props);
    this.state = TargetStore.getResults();
  }

  componentDidMount() {
    TargetAction.fetchResults();
    TargetStore.addChangeListener(this.onChange.bind(this));
  }

  componentWillUnmount() {
    TargetStore.removeChangeListener(this.onChange.bind(this));
  }
  //@TODO retrieve data
  onChange() {
    var data = [{"Destination":"google.fr","Status":"good","Time":2115,"Created_at":"2015-06-03T07:21:26.614191Z","Target_id":1},
      {"Destination":"google.fr","Status":"good","Time":5832,"Created_at":"2015-06-03T07:21:26.614191Z","Target_id":1},
      {"Destination":"google.fr","Status":"good","Time":220,"Created_at":"2015-06-03T07:21:26.614191Z","Target_id":1},
      {"Destination":"google.fr","Status":"good","Time":2378,"Created_at":"2015-06-03T07:21:26.614191Z","Target_id":1},
      {"Destination":"google.fr","Status":"good","Time":888,"Created_at":"2015-06-03T07:21:26.614191Z","Target_id":1},
      {"Destination":"google.fr","Status":"good","Time":999,"Created_at":"2015-06-03T07:21:26.614191Z","Target_id":1},
      {"Destination":"google.fr","Status":"good","Time":3202,"Created_at":"2015-06-03T07:21:26.614191Z","Target_id":1},
      {"Destination":"google.fr","Status":"good","Time":952,"Created_at":"2015-06-03T07:21:26.614191Z","Target_id":1},
      {"Destination":"google.fr","Status":"good","Time":7777,"Created_at":"2015-06-03T07:21:26.614191Z","Target_id":1},
      {"Destination":"google.fr","Status":"good","Time":2165,"Created_at":"2015-06-03T07:21:26.614191Z","Target_id":1},
      {"Destination":"google.fr","Status":"good","Time":500,"Created_at":"2015-06-03T07:21:26.614191Z","Target_id":1},
      {"Destination":"google.fr","Status":"good","Time":1000,"Created_at":"2015-06-03T07:21:26.614191Z","Target_id":1},
      {"Destination":"google.fr","Status":"good","Time":2444,"Created_at":"2015-06-03T07:21:26.614191Z","Target_id":1},
      {"Destination":"google.fr","Status":"good","Time":3500,"Created_at":"2015-06-03T07:21:26.614191Z","Target_id":1},
      {"Destination":"google.fr","Status":"good","Time":1203,"Created_at":"2015-06-03T07:21:26.614191Z","Target_id":1},
      {"Destination":"google.fr","Status":"good","Time":855,"Created_at":"2015-06-03T07:21:26.614191Z","Target_id":1},
      {"Destination":"google.fr","Status":"good","Time":4000,"Created_at":"2015-06-03T07:21:26.614191Z","Target_id":1},
      {"Destination":"google.fr","Status":"good","Time":1500,"Created_at":"2015-06-03T07:21:26.614191Z","Target_id":1},
      {"Destination":"google.fr","Status":"good","Time":600,"Created_at":"2015-06-03T07:21:26.614191Z","Target_id":1},
      {"Destination":"google.fr","Status":"good","Time":300,"Created_at":"2015-06-03T07:21:26.614191Z","Target_id":1},
      {"Destination":"google.fr","Status":"good","Time":250,"Created_at":"2015-06-03T07:21:26.614191Z","Target_id":1},
      {"Destination":"google.fr","Status":"good","Time":1089,"Created_at":"2015-06-03T07:21:26.614191Z","Target_id":1},
      {"Destination":"google.fr","Status":"good","Time":1720,"Created_at":"2015-06-03T07:21:26.614191Z","Target_id":1},
      {"Destination":"google.fr","Status":"good","Time":3844,"Created_at":"2015-06-03T07:21:26.614191Z","Target_id":1},
      {"Destination":"google.fr","Status":"good","Time":755,"Created_at":"2015-06-03T07:21:26.614191Z","Target_id":1},
      {"Destination":"google.fr","Status":"good","Time":9898,"Created_at":"2015-06-03T07:21:26.614191Z","Target_id":1},
      {"Destination":"google.fr","Status":"good","Time":3599,"Created_at":"2015-06-03T07:21:26.614191Z","Target_id":1},
      {"Destination":"google.fr","Status":"good","Time":5555,"Created_at":"2015-06-03T07:21:26.614191Z","Target_id":1},
      {"Destination":"google.fr","Status":"good","Time":2144,"Created_at":"2015-06-03T07:21:26.614191Z","Target_id":1}];
      console.log(TargetStore.getResults());
    this.setState({ results: data });
  }
}

export default ResultDurationChart;
