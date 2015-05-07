var React = require('react');

var HelloWorld = React.createClass({
  render: function() {
    return (
      <h1>It works yeahhhh!</h1>
    );
  }
});
React.render(
  React.createElement(HelloWorld, null),
  document.getElementById('content')
);
