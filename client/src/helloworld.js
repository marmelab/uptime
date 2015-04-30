var HelloWorld = React.createClass({
  displayName: 'HelloWorld',
  render: function() {
    return (
      <h1>Helloworld!</h1>
    );
  }
});
React.render(
  React.createElement(HelloWorld, null),
  document.getElementById('content')
);
