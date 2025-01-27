// toggleVisibility.test.js
describe('toggleVisibility function', function () {

    it('should change the display style to "none" if it was "block"', function () {
      var element = { style: { display: 'block' } };
      spyOn(element.style, 'display', 'set');  // Spy on the style.display property
      toggleVisibility(element);
      expect(element.style.display).toBe('none');
    });
  
    it('should change the display style to "block" if it was "none"', function () {
      var element = { style: { display: 'none' } };
      spyOn(element.style, 'display', 'set');  // Spy on the style.display property
      toggleVisibility(element);
      expect(element.style.display).toBe('block');
    });
  
  });
  