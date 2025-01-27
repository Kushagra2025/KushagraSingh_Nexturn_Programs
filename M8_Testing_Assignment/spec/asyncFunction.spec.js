// delayedGreeting.test.js
const asyncFunction = require('../src/asyncFunction');
describe('delayedGreeting function', function () {

    beforeEach(function () {
      jasmine.clock().install();  // Mocking the clock to simulate setTimeout
    });
  
    afterEach(function () {
      jasmine.clock().uninstall();  // Uninstall clock after each test
    });
  
    it('should return the correct greeting message after the delay', function (done) {
      var promise = delayedGreeting('Alice', 1000);
      jasmine.clock().tick(1000); // Simulate the passage of 1000ms
      promise.then(function (greeting) {
        expect(greeting).toBe('Hello, Alice!');
        done();  // Finish the test when the async call is done
      });
    });
  
    it('should respect the delay', function (done) {
      var delay = 1000;
      var start = Date.now();
      var promise = delayedGreeting('Bob', delay);
      jasmine.clock().tick(delay); // Simulate the delay
      promise.then(function () {
        var end = Date.now();
        expect(end - start).toBeGreaterThanOrEqual(delay);
        done();
      });
    });
  
  });
  