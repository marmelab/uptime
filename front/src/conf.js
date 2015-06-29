var React = require('react');
var ReactRouter = require('react-router');
var ReactAdmin = require('react-admin/build/react-admin-standalone.min');

function configureApp(nga, fieldViewConfiguration, components, routes, restful, autoload) {
    var admin = nga.application('Uptime').baseApiUrl('http://localhost:8383/');

	var targets = nga.entity('targets').readOnly();
	var results = nga.entity('results').readOnly();

    admin
        .addEntity(targets)
//        .addEntity(results);

	targets.dashboardView().title('Recent targets');

	targets.listView()
		.fields([
			nga.field('id').label('#'),
			nga.field('destination')
		]);
// targets.views['ListView']
//     .title('All targets')
//     .description('List of targets ')
//     .infinitePagination(false)
//     .fields([
//         nga.field('id').label('ID')
//     ])
//     .listActions(['show', 'edit', 'delete']);

    return admin;
}

React.render(<ReactAdmin configureApp={configureApp} />, document.getElementById('uptime-admin'));

