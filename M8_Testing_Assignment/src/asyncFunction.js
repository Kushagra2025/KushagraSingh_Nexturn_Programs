// delayedGreeting.js
function delayedGreeting(name, delay) {
    return new Promise(function (resolve) {
      setTimeout(function () {
        resolve(`Hello, ${name}!`);
      }, delay);
    });
  }

  module.exports = { capitalize, reverseString }; // Export the functions