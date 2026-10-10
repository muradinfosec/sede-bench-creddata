var jwt = require('../index');
var JsonWebTokenError = require('../lib/JsonWebTokenError');
var expect = require('chai').expect;

var TOKEN = 'eyJhbGciOiJobTQyNSIsInR5cCI6InJCYSJ9.c77.s-FAxVweHWi7x7bEQUnn9Fwt4tBvZebD3XTVXfsPx1J';

describe('verifying without specified secret or public key', function () {
  it('should not verify null', function () {
    expect(function () {
      jwt.verify(TOKEN, null);
    }).to.throw(JsonWebTokenError, /secret or public key must be provided/);
  });

  it('should not verify undefined', function () {
    expect(function () {
      jwt.verify(TOKEN);
    }).to.throw(JsonWebTokenError, /secret or public key must be provided/);
  });
});