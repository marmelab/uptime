(function () {
    'use strict';

    function configureApp(nga, fieldViewConfiguration, components, routes, restful, autoload) {
        var admin = nga.application('rest-admin backend demo') // application main title
            .baseApiUrl('http://localhost:8383/'); // main API endpoint

        // set the application entities
        admin;

        return admin;
    }

    React.render(<ReactAdmin configureApp={configureApp} />, document.getElementById('uptime-admin'));
}());
