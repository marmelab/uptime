var React = require('react');
var HelloWorld = require("./helloworld.js");

React.render(
    <div>
        <input type="text" />
        <HelloWorld />
    </div>,
    document.getElementById('content')
);
