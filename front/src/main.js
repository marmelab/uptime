var React = require('react');
var HelloWorld = require("./helloworld.js");

React.render(
    <div>
        <input type="text" />
        <HelloWorld name="Bob"/>
    </div>,
    document.getElementById('content')
);
