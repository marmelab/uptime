var React = require('react');

var HelloWorld = React.createClass({
  render: function() {
    return (
      <h1>It works yeahhhhhhhhh!</h1>
    );
  }
});
React.render(
  React.createElement(HelloWorld, null),
  document.getElementById('content')
);
