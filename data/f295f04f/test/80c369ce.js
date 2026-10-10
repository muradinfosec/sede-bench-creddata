'use strict'

var http = require('http')
var request = require('../index')
var httpSignature = require('http-signature')
var tape = require('tape')

var privateKeyPEMs = {}

privateKeyPEMs['key-1'] =
  '-----BEGIN RSA PRIVATE KEY-----\n' +
  'MIIEpAIBAAKCAQEAzWSJl+Z9Bqv00FVL5N3+JCUoqmQPjIlya1BbeqQroNQ5yO7s\n' +
  'aQwTloHQn8czwU2S2dCawNV57pePMrYPV5nKhepzqVgPDKOQMcK16Qg9TXutF0i1\n' +
  '4Wtrq9Z/CZ6GJbKPkS0TLczfngVuWg3gOivdLa4fQ4FyHflp5qMt4qtql0sjcINW\n' +
  'A6GbOLmlQHSKkmgjedeUY8M6Vdj7J3uNH8a151Yz6uDFCQH5b19TZi/ICiG4ESp2\n' +
  '1HeqgAsogzyA2Dy9d9yOb7WCr9+WIBQB5CJymzRW/Wh+ouHx3NXR6vznx+JBlRT7\n' +
  'kj07qnye9O6Jl3C77yVIb3PnkSMskWxy0gio5quUPImBBlSWMB1Iiq4/nYynjpSc\n' +
  'P8xXUX2nGkRkkdGPHvRaZx8uhgjvT8Lmialxysuy+2YO/dHssLr/QuzXJo2lcy56\n' +
  'PBY3r52lhR2jeJnoteSDG/JQlJxWDJIDi2jrSgmsKhxqPGBZqgMLM8etBjo3aVOe\n' +
  'fGChvOYuZ+mssLl0qzRuIAwzV2Jjax6IZHyUn/rajB7yosLT67migZ4Oj2dcXSuy\n' +
  '71rAj7f7/0K+S93ueqO7mLimmfa5cBioUv+uHfXG88PyaXLdYP9hFhAYuAU2hhkx\n' +
  '/U9F4JqgLzJ8HBPdOTQ95RpaceVSKewyYj13sWT5btiLtQys28xbl2T2o0k7Tguz\n' +
  '/LmAkvLoYwVg+gh7NJhWNaqcLqq6EmE3cqhKjdSnxsKSiAY8poEQEPwP9xP2KPE2\n' +
  'LWlyBh7pbJpPXm98oWnwtpGq5bW9S2LCd3GoeLNqtQTpqNqfs0v8mnprR8hFfaW6\n' +
  'MHg8gwTUPLCAU5Bhylr4K5rjVTM3y4TrJ8amW97PTzsSQnUHvZe3kARYDyVl3zHj\n' +
  '004U16nw/mj7PE7ZJR+s0AYvzNg+/w3A0VFBboMjenkL6WJ7flqxjBSWP3zmphOT\n' +
  'QrYyRz4/dNsVeovIzupYxG1X6ADNghVGPA5JcLB60nbST442Du6LQi+VuAMcB1NB\n' +
  'OC28oRfMmj/wmn1YG4ibkvWgp5w96OfTjU0w5JNwVqVkGt+8/AxkzpDWTTiC2Wtl\n' +
  '8vtBtaI0dkf7Q+LUbJq9WHZDQ82yr7BO45tviHSi5jFJW1AomNr0/nNwwaYamtjr\n' +
  'qBFWmMQ2NT8rmLKwMth0qSWSSwZwJsGLwiM1hconerj/hT0oIR4txF3YW3CAL3PL\n' +
  'i06iW+R52YILhauGckgPWWwhChvJIV94X/LwD/bOQOclz6kstzB8f9L9OSo4eEhX\n' +
  'Bz3SKafiVt5aXjobFVVhiHatynkV1D6sVlpfi7jrh+mNdVMVA6fNCFm6+w6Poy5B\n' +
  '9FAMHuDxOJYwtJVPZ5PDYpli/pc3zMCP1sgjqAULmXHku9eWkwyAHfK3jPh3TsN/\n' +
  'ApQuxRznNDqVmGCoyGjFRvBP3ZAWBQx2IKZBcxFvcI24ht4IdpF2Uu1NU8FM++KH\n' +
  '5xp2wLeB974yCRz87oEsHsWPmcEt3ZNqIEpswKkcLAHF51aiKXQkAZ0yJacH1/Uy\n' +
  '8Av7dwhFpa5PajmMvpkMJzGTQXynOoNmH631Lg8W8leN7Aoc99YTXt==\n' +
  '-----END RSA PRIVATE KEY-----'

var publicKeyPEMs = {}

publicKeyPEMs['key-1'] =
  '-----BEGIN PUBLIC KEY-----\n' +
  'MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAzWSJl+Z9Bqv00FVL5N3+\n' +
  'JCUoqmQPjIlya1BbeqQroNQ5yG1iVbYTTnMRa1zQtR6r2fNvWeg94DvxivxIG9di\n' +
  'DMnrzijAnYlTLOl84CK2vOxkj5b68zrLH9b/Gd6NOHsywo8IjvXvCeTfca5WUHcu\n' +
  'Vi2lT9VjygFs1ILG4RyeX1BXUumuY8fzmposxLYdMxCqUTzAn0u9Saq2H2OVj5u1\n' +
  '14wS7OQPigu6G99dpn/iPHa3zBm87baBWDbqZWRW0BP3K6eqq8sut1+NLhNW8ADP\n' +
  'TdnO/SO+kvXy7fqd8atSn+HlQcx6tW42dhXf3E9uE7K78eZtW0KvfyNGAjsI1Fft\n' +
  '2QIDAQAB\n' +
  '-----END PUBLIC KEY-----'

publicKeyPEMs['key-2'] =
  '-----BEGIN PUBLIC KEY-----\n' +
  'MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAqp04VVr9OThli9b35Omz\n' +
  'VqSfWbsoQuRrgyWsrNRn3XkFmbWw4FzZwQ42OgGMzQ84Ta4d9zGKKQyFriTiPjPf\n' +
  'xhhrsaJnDuybcpVhcr7UNKjSZ0S59tU3hpRiEz6hO+Nc/OSSLkvalG0VKrxOln7J\n' +
  'LK/h3rNS/l6wDZ5S/KqsI6CYtV2ZLpn3ahLrizvEYNY038Qcm38qMWx+VJAvZ4di\n' +
  'qqmW7RLIsLT59SWmpXdhFKnkYYGhxrk1Mwl22dBTJNY5SbriU5G3gWgzYkm8pgHr\n' +
  '6CtrXch9ciJAcDJehPrKXNvNDOdUh8EW3fekNJerF1lWcwQg44/12v8sDPyfbaKB\n' +
  'dQIDAQAB\n' +
  '-----END PUBLIC KEY-----'

var server = http.createServer(function (req, res) {
  var parsed = httpSignature.parseRequest(req)
  var publicKeyPEM = publicKeyPEMs[parsed.keyId]
  var verified = httpSignature.verifySignature(parsed, publicKeyPEM)
  res.writeHead(verified ? 200 : 400)
  res.end()
})

tape('setup', function (t) {
  server.listen(0, function () {
    server.url = 'http://localhost:' + this.address().port
    t.end()
  })
})

tape('correct key', function (t) {
  var options = {
    httpSignature: {
      keyId: 'key-1',
      key: privateKeyPEMs['key-1']
    }
  }
  request(server.url, options, function (err, res, body) {
    t.equal(err, null)
    t.equal(200, res.statusCode)
    t.end()
  })
})

tape('incorrect key', function (t) {
  var options = {
    httpSignature: {
      keyId: 'key-2',
      key: privateKeyPEMs['key-1']
    }
  }
  request(server.url, options, function (err, res, body) {
    t.equal(err, null)
    t.equal(400, res.statusCode)
    t.end()
  })
})

tape('cleanup', function (t) {
  server.close(function () {
    t.end()
  })
})
