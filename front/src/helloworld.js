var HelloWorld = React.createClass({
  displayName: 'Helloworld !',
  render: function() {
    return (
      <h1>It works!</h1>
    );
  }
});
React.render(
  React.createElement(HelloWorld, null),
  document.getElementById('content')
);
