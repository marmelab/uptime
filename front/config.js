(function () {
    'use strict';

    function truncate(value) {
        if (!value) {
            return '';
        }

        return value.length > 50 ? value.substr(0, 50) + '...' : value;
    }

    function configureApp(nga, fieldViewConfiguration, components, routes, restful, autoload) {
        restful.addFullRequestInterceptor(function (url, params, headers, data) {
            headers['X-From'] = 'react-admin';

            return {
                headers: headers
            };
        });

        // Add custom component
        var SendEmail = React.createClass({
            render: function () {
                return <a className='btn btn-default' href='#/stats'>Show stats</a>;
            }
        });
        autoload('SendEmail', SendEmail);

        var admin = nga.application('rest-admin backend demo') // application main title
            .baseApiUrl('http://localhost:8383/'); // main API endpoint

        // define all entities at the top to allow references between them
        var targets = nga.entity('targets'); // the API endpoint for targetss will be http://localhost:8383/targetss/:id

        // set the application entities
        admin
            .addEntity(targets)

        targets.views['DashboardView'] // customize the dashboard panel for this entity
            .title('Recent targetss')
            .order(1) // display the targets panel first in the dashboard
            .limit(5) // limit the panel to the 5 latest targetss
            .fields([nga.field('title').isDetailLink(true).map(truncate)]); // fields() called with arguments add fields to the view

        targets.views['ListView']
            .title('All targetss') // default title is '[Entity_name] list'
            .description('List of targetss with infinite pagination') // description appears under the title
            .infinitePagination(true) // load pages as the user scrolls
            .fields([
                nga.field('id').label('ID'), // The default displayed name is the camelCase field name. label() overrides id
                nga.field('title'), // the default list field type is 'string', and displays as a string
                nga.field('published_at', 'date'), // Date field type allows date formatting
                nga.field('views', 'number'),
                nga.field('tags', 'reference_many') // a Reference is a particular type of field that references another entity
                    .targetEntity(tag) // the tag entity is defined later in this file
                    .targetField(nga.field('name')) // the field to be displayed in this list
            ])
            .listActions(['show', 'edit', 'delete']);

        targets.views['CreateView']
            .fields([
                nga.field('title') // the default edit field type is 'string', and displays as a text input
                    .attributes({ placeholder: 'the targets title' }) // you can add custom attributes, too
                    .validation({ required: true, minlength: 3, maxlength: 100 }), // add validation rules for fields
                nga.field('teaser', 'text'), // text field type translates to a textarea
                nga.field('body', 'wysiwyg'), // overriding the type allows rich text editing for the body
                nga.field('published_at', 'date') // Date field type translates to a datepicker
            ]);

        var subCategories = [
            { category: 'tech', label: 'Computers', value: 'computers' },
            { category: 'tech', label: 'Gadgets', value: 'gadgets' },
            { category: 'lifestyle', label: 'Travel', value: 'travel' },
            { category: 'lifestyle', label: 'Fitness', value: 'fitness' }
        ];

        targets.views['EditView']
            .title('Edit targets "{ entry.values.title }"') // title() accepts a template string, which has access to the entry
            .actions(['list', 'show', 'delete']) // choose which buttons appear in the top action bar. Show is disabled by default

        targets.views['ShowView'] // a showView displays one entry in full page - allows to display more data than in a a list
            .fields([
                nga.field('id'),
                targets.views['EditView'].fields(), // reuse fields from another view in another order
                 nga.field('custom_action', 'template')
                     .label('')
                     .template(<SendEmail targets="entry"></SendEmail>)
            ]);

        targets.views['DeleteView']
            .title('Delete targets "{ entry.values.title }"');



        // customize menu
        admin.menu(nga.menu()
                .addChild(nga.menu(targets).icon('<span class="glyphicon glyphicon-file"></span>')) // customize the entity menu icon
                .addChild(nga.menu().title('Other')
                    .addChild(nga.menu().title('Stats').icon('').link('/stats'))
                )
            );

        // Add custom route
        var ViewActions = components.ViewActions;
        var Route = ReactRouter.Route;
        var Stats = React.createClass({
            render: function () {
                return <div>
                    <ViewActions buttons={['back']} />
                    <h1>Stats</h1>
                    <p className='lead'>You can add custom pages, too</p>
                </div>;
            }
        });

        routes.props.children.push(<Route name="stats" path="/stats" handler={Stats} />);

        return admin;
    }

    React.render(<ReactAdmin configureApp={configureApp} />, document.getElementById('my-app'));
}());
