var HelloWorld = React.createClass({
  render: function() {
    return (
      <h1>It works yeahhh!</h1>
    );
  }
});
React.render(
  React.createElement(HelloWorld, null),
  document.getElementById('content')
);
