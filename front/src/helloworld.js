var HelloWorld = React.createClass({
  render: function() {
    return (
      <h1>It works yeah!</h1>
    );
  }
});
React.render(
  React.createElement(HelloWorld, null),
  document.getElementById('content')
);
