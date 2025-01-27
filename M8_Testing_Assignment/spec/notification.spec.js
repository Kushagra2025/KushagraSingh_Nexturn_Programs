// sendNotification.test.js
describe('sendNotification function', function () {

    it('should return "Notification Sent" when notification is successfully sent', function () {
      var mockService = {
        send: function () {
          return true;
        }
      };
      var result = sendNotification(mockService, 'Test message');
      expect(result).toBe('Notification Sent');
    });
  
    it('should return "Failed to Send" when notification fails to send', function () {
      var mockService = {
        send: function () {
          return false;
        }
      };
      var result = sendNotification(mockService, 'Test message');
      expect(result).toBe('Failed to Send');
    });
  
  });
  