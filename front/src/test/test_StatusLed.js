var jsdom = require('jsdom');
var React = require('react/addons'),
    assert = require('assert'),
    StatusLed = require('../components/StatusLed'),
    TestUtils = React.addons.TestUtils;

global.document = jsdom.jsdom('<!doctype html><html><body></body></html>');
global.window = document.parentWindow;

describe('StatusLed component', function(){
  before('render and locate element', function() {
  	
    var renderedComponent = TestUtils.renderIntoDocument(
      <StatusLed data={true}/>
    );

    var StatusLedComponent = TestUtils.findRenderedDOMComponentWithTag(
      renderedComponent,
      'StatusLed'
    );

    this.StatusLedElement = StatusLedComponent.getDOMNode();
  });

  it('StatusLed background should be green"', function() {
    assert(this.StatusLedElement.getAttribute('styles').background === 'green');
  });

  /*it('StatusLed background should not be red', function() {
    assert(this.StatusLedElement.getAttribute('styles').background != 'red');
  });*/
});
