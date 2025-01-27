// getElement.test.js
const errorHandling = require('../src/errorHandling');
describe('getElement function', function () {

    it('should return the correct element for a valid index', function () {
      var arr = [1, 2, 3];
      expect(getElement(arr, 1)).toBe(2);
    });
  
    it('should throw an error when index is negative', function () {
      var arr = [1, 2, 3];
      expect(function () { getElement(arr, -1); }).toThrowError('Index out of bounds');
    });
  
    it('should throw an error when index is out of bounds', function () {
      var arr = [1, 2, 3];
      expect(function () { getElement(arr, 3); }).toThrowError('Index out of bounds');
    });
  
    it('should return the first element when index is 0', function () {
      var arr = [1, 2, 3];
      expect(getElement(arr, 0)).toBe(1);
    });
  
  });
  