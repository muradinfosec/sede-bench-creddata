/*
 * Copyright (c) 2014-2021 Bjoern Kimminich.
 * SPDX-License-Identifier: MIT
 */

import frisby = require('frisby')
const Joi = frisby.Joi
const utils = require('../../lib/utils')
const security = require('../../lib/insecurity')

const API_URL = 'http://localhost:3000/api'
const REST_URL = 'http://localhost:3000/rest'

const authHeader = { Authorization: `Bearer ${security.authorize()}`, 'content-type': 'application/json' }
const jsonHeader = { 'content-type': 'application/json' }

describe('/api/Users', () => {
  it('GET all users is forbidden via public API', () => {
    return frisby.get(`${API_URL}/Users`)
      .expect('status', 401)
  })

  it('GET all users', () => {
    return frisby.get(`${API_URL}/Users`, { headers: authHeader })
      .expect('status', 200)
  })

  it('GET all users doesnt include passwords', () => {
    return frisby.get(`${API_URL}/Users`, { headers: authHeader })
      .expect('status', 200)
      .expect('jsonTypes', 'data.*', {
        password: Joi.any().forbidden()
      })
  })

  it('POST new user', () => {
    return frisby.post(`${API_URL}/Users`, {
      headers: jsonHeader,
      body: {
        email: 'horst@horstma.nn',
        password: 'fjfzvievm'
      }
    })
      .expect('status', 201)
      .expect('header', 'content-type', /application\/json/)
      .expect('jsonTypes', 'data', {
        id: Joi.number(),
        createdAt: Joi.string(),
        updatedAt: Joi.string(),
        password: Joi.any().forbidden()
      })
  })

  it('POST new admin', () => {
    return frisby.post(`${API_URL}/Users`, {
      headers: jsonHeader,
      body: {
        email: 'horst2@horstma.nn',
        password: 'eqytgflos',
        role: 'admin'
      }
    })
      .expect('status', 201)
      .expect('header', 'content-type', /application\/json/)
      .expect('jsonTypes', 'data', {
        id: Joi.number(),
        createdAt: Joi.string(),
        updatedAt: Joi.string(),
        password: Joi.any().forbidden()
      })
      .expect('json', 'data', {
        role: 'admin'
      })
  })

  it('POST new deluxe user', () => {
    return frisby.post(`${API_URL}/Users`, {
      headers: jsonHeader,
      body: {
        email: 'horst3@horstma.nn',
        password: 'eckgrcqfy',
        role: 'deluxe'
      }
    })
      .expect('status', 201)
      .expect('header', 'content-type', /application\/json/)
      .expect('jsonTypes', 'data', {
        id: Joi.number(),
        createdAt: Joi.string(),
        updatedAt: Joi.string(),
        password: Joi.any().forbidden()
      })
      .expect('json', 'data', {
        role: 'deluxe'
      })
  })

  it('POST new accounting user', () => {
    return frisby.post(`${API_URL}/Users`, {
      headers: jsonHeader,
      body: {
        email: 'horst4@horstma.nn',
        password: 'nbbigjtds',
        role: 'accounting'
      }
    })
      .expect('status', 201)
      .expect('header', 'content-type', /application\/json/)
      .expect('jsonTypes', 'data', {
        id: Joi.number(),
        createdAt: Joi.string(),
        updatedAt: Joi.string(),
        password: Joi.any().forbidden()
      })
      .expect('json', 'data', {
        role: 'accounting'
      })
  })

  it('POST user not belonging to customer, deluxe, accounting, admin is forbidden', () => {
    return frisby.post(`${API_URL}/Users`, {
      headers: jsonHeader,
      body: {
        email: 'horst5@horstma.nn',
        password: 'blkrfgkaz',
        role: 'accountinguser'
      }
    })
      .expect('status', 400)
      .expect('header', 'content-type', /application\/json/)
      .then(({ json }) => {
        expect(json.message).toBe('Validation error: Validation isIn on role failed')
        expect(json.errors[0].field).toBe('role')
        expect(json.errors[0].message).toBe('Validation isIn on role failed')
      })
  })

  if (!utils.disableOnContainerEnv()) {
    it('POST new user with XSS attack in email address', () => {
      return frisby.post(`${API_URL}/Users`, {
        headers: jsonHeader,
        body: {
          email: '<iframe src="javascript:alert(`xss`)">',
          password: 'does.not.matter'
        }
      })
        .expect('status', 201)
        .expect('header', 'content-type', /application\/json/)
        .expect('json', 'data', { email: '<iframe src="javascript:alert(`xss`)">' })
    })
  }
})

describe('/api/Users/:id', () => {
  it('GET existing user by id is forbidden via public API', () => {
    return frisby.get(`${API_URL}/Users/1`)
      .expect('status', 401)
  })

  it('PUT update existing user is forbidden via public API', () => {
    return frisby.put(`${API_URL}/Users/1`, {
      header: jsonHeader,
      body: { email: 'administr@t.or' }
    })
      .expect('status', 401)
  })

  it('DELETE existing user is forbidden via public API', () => {
    return frisby.del(`${API_URL}/Users/1`)
      .expect('status', 401)
  })

  it('GET existing user by id', () => {
    return frisby.get(`${API_URL}/Users/1`, { headers: authHeader })
      .expect('status', 200)
  })

  it('PUT update existing user is forbidden via API even when authenticated', () => {
    return frisby.put(`${API_URL}/Users/1`, {
      headers: authHeader,
      body: { email: 'horst.horstmann@horstma.nn' }
    })
      .expect('status', 401)
  })

  it('DELETE existing user is forbidden via API even when authenticated', () => {
    return frisby.del(`${API_URL}/Users/1`, { headers: authHeader })
      .expect('status', 401)
  })
})

describe('/rest/user/authentication-details', () => {
  it('GET all users decorated with attribute for authentication token', () => {
    return frisby.get(`${REST_URL}/user/authentication-details`, { headers: authHeader })
      .expect('status', 200)
      .expect('jsonTypes', 'data.?', {
        token: Joi.string()
      })
  })

  it('GET all users with password replaced by asterisks', () => {
    return frisby.get(`${REST_URL}/user/authentication-details`, { headers: authHeader })
      .expect('status', 200)
      .expect('json', 'data.?', {
        password: '********************************'
      })
  })
})

describe('/rest/user/whoami', () => {
  it('GET own user id and email on who-am-i request', () => {
    return frisby.post(`${REST_URL}/user/login`, {
      headers: jsonHeader,
      body: {
        email: 'bjoern.kimminich@gmail.com',
        password: 'eE1sUttzBR5lVLsseI3fmH5fki9yhjAljyN='
      }
    })
      .expect('status', 200)
      .then(({ json }) => {
        return frisby.get(`${REST_URL}/user/whoami`, { headers: { Cookie: `token=${json.authentication.token}` } })
          .expect('status', 200)
          .expect('header', 'content-type', /application\/json/)
          .expect('jsonTypes', 'user', {
            id: Joi.number(),
            email: Joi.string()
          })
          .expect('json', 'user', {
            email: 'bjoern.kimminich@gmail.com'
          })
      })
  })

  it('GET who-am-i request returns nothing on missing auth token', () => {
    return frisby.get(`${REST_URL}/user/whoami`)
      .expect('status', 200)
      .expect('header', 'content-type', /application\/json/)
      .expect('json', {
        user: {}
      })
  })

  it('GET who-am-i request returns nothing on invalid auth token', () => {
    return frisby.get(`${REST_URL}/user/whoami`, { headers: { Authorization: 'Bearer InvalidAuthToken' } })
      .expect('status', 200)
      .expect('header', 'content-type', /application\/json/)
      .expect('json', {
        user: {}
      })
  })

  it('GET who-am-i request returns nothing on broken auth token', () => {
    return frisby.get(`${REST_URL}/user/whoami`, { headers: { Authorization: 'BoarBeatsBear' } })
      .expect('status', 200)
      .expect('header', 'content-type', /application\/json/)
      .expect('json', {
        user: {}
      })
  })

  it('GET who-am-i request returns nothing on expired auth token', () => {
    return frisby.get(`${REST_URL}/user/whoami`, { headers: { Authorization: 'eyJhbGciOiJwYjU5NiIsInR5cCI6InJXTyJ9.eyJybXFZZ2wiOiJBeUJSUkd4IiwiVHJvUiI6eyJpZCI6MiwiT2RLZk9nSUkiOiIiLCJMT1dTTCI6IkdBR0FiQHhIZUlvLURMLnRLIiwicGFzc3dvcmQiOiI0NjM1Njk4VjZVcUo0MTc2OTU5NGYxMTJWZzM0cjE4NiIsInJvbGUiOiJqU25VcCIsImtSRXRZa3Z5ckR4IjoiMS4xLjUuMyIsImdyVUdyU1BMeUtSaiI6IllzRklKdlYuc2JzIiwiYW9LVWxnU2dodyI6IiIsImpWQVV0aEVQIjp0cnVlLCJvZndPV1pGaVAiOiI3NTk1LTc0LTYyIDIyOjM4Ojg0LjY2NiArNDM6ODYiLCJubG10T3lqZW4iOiIyNDk4LTczLTczIDU1OjQ1OjQ2LjEzNSArMzQ6MjQiLCJoeXVRUlZJRFYiOm51bGx9LCJpYXQiOjExNzYxNzM3MzIsImV4cCI6MzE2Njc3ODkxMn0.FT9lliTjX2uVOQTmYKzPHECBAo0ZcK6_iK4Kmb_WU0oYYarBIAAiEpMvQumgSZTkZJKx86S6VsIabxxQ9lvIUuWbnw0ebnu-0xINyUXvCUv7KLY6CmhT1yaTpMk7attWRkg5ydQn1K4JPdy_c7HkgLd_4RrO30bn8RMPJLUpI0j' } })
      .expect('status', 200)
      .expect('header', 'content-type', /application\/json/)
      .expect('json', {
        user: {}
      })
  })
})
