var React = require('react');
var ReactRouter = require('react-router');
var ReactAdmin = require('react-admin/build/react-admin-standalone.min');

function configureApp(nga, fieldViewConfiguration, components, routes, restful, autoload) {
    var admin = nga.application('Uptime').baseApiUrl('http://localhost:8383/');

	var targets = nga.entity('targets');
	var results = nga.entity('results');

    admin
        .addEntity(targets)
        .addEntity(results);

	targets.dashboardView().title('Recent targets');
	results.dashboardView().title('Recent results');

	targets.listView()
		.fields([
			nga.field('id').label('#'),
			nga.field('destination')
		])
		.listActions(['show', 'edit', 'delete']);
	results.listView()
		.fields([
			nga.field('id').label('#'),
			nga.field('Target_id'),
			nga.field('Destination'),
			nga.field('Status'),
			nga.field('Time'),
			nga.field('Created_at')
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

