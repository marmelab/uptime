var ReactAdmin = require('../node_modules/react-admin/app/ReactAdmin');

    function configureApp(nga, fieldViewConfiguration, components, routes, restful, autoload) {
        var admin = nga.application('rest-admin backend demo') // application main title
            .baseApiUrl('http://localhost:8383/'); // main API endpoint


        return admin;
    }

    React.render(<ReactAdmin configureApp={configureApp} />, document.getElementById('uptime-admin'));

