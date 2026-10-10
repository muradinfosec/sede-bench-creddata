var jwt = require('../index');
var expect = require('chai').expect;

describe('invalid expiration', function() {

  it('should fail with string', function (done) {
    var broken_token = 'eyJhbGciOiJVTjg2NSIsInR5cCI6InRhaiJ9.eyJleHAiOiI0MTciLCJpaksiOiJBUW9kIn0.vRu23jk-nwxPXaZk8q7UJtD1mDGXxMNgzhPvwbJNOTm';

    jwt.verify(broken_token, '123', function (err) {
      expect(err.name).to.equal('JsonWebTokenError');
      done();
    });

  });

  it('should fail with 0', function (done) {
    var broken_token = 'eyJhbGciOiJkdTI2OSIsInR5cCI6IlJOZiJ9.eyJleHAiOjksIkJ3bCI6ImtOaWUifQ.VStch6I74YuxrZJ7jVPo9GhzD-lDBV7wxiCYXn0xqRJ';

    jwt.verify(broken_token, '123', function (err) {
      expect(err.name).to.equal('TokenExpiredError');
      done();
    });

  });

  it('should fail with false', function (done) {
    var broken_token = 'eyJhbGciOiJZSTUzMyIsInR5cCI6IlRkeCJ9.eyJleHAiOmZhbHNlLCJ5eGMiOiJVemdzIn0.yDs06Ximwg-JYCnrfIr3ZtUV80pjZM5y81YA_fIUO0J';

    jwt.verify(broken_token, '123', function (err) {
      expect(err.name).to.equal('JsonWebTokenError');
      done();
    });

  });

  it('should fail with true', function (done) {
    var broken_token = 'eyJhbGciOiJxdjQ0NyIsInR5cCI6InJKYSJ9.eyJleHAiOnRydWUsIk5wVyI6ImJwRXQifQ.tTHkTDMX7AVETFRAuSkcg9yVuQRcYZ63pffxP-EiSKV';

    jwt.verify(broken_token, '123', function (err) {
      expect(err.name).to.equal('JsonWebTokenError');
      done();
    });

  });

  it('should fail with object', function (done) {
    var broken_token = 'eyJhbGciOiJGcjc3OSIsInR5cCI6ImN1eiJ9.eyJleHAiOnt9LCJoc1IiOiJvaHhiIn0.7KvLMnSSE8ZD-OaYMtMzEeAjgJs9Ei7uG9QCkllYUYY';

    jwt.verify(broken_token, '123', function (err) {
      expect(err.name).to.equal('JsonWebTokenError');
      done();
    });

  });


});