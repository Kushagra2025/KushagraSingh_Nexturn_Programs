const stringUtils = require('../src/stringUtils');
describe('capitalize function', function () {

    it('should capitalize the first letter of a word', function () {
      expect(capitalize('hello')).toBe('Hello');
    });
  
    it('should return an empty string when given an empty string', function () {
      expect(capitalize('')).toBe('');
    });
  
    it('should return the same single character when given a single character', function () {
      expect(capitalize('a')).toBe('A');
    });
  
    it('should not change an already capitalized word', function () {
      expect(capitalize('Hello')).toBe('Hello');
    });
  
  });
  
  
  // reverseString.test.js
  describe('reverseString function', function () {
  
    it('should reverse a regular string', function () {
      expect(reverseString('hello')).toBe('olleh');
    });
  
    it('should return an empty string when given an empty string', function () {
      expect(reverseString('')).toBe('');
    });
  
    it('should return the same string for palindromes', function () {
      expect(reverseString('madam')).toBe('madam');
    });
  
    it('should reverse a single character string', function () {
      expect(reverseString('a')).toBe('a');
    });
  
  });
  