/*
 * Copyright (c) 2014-2021 Bjoern Kimminich.
 * SPDX-License-Identifier: MIT
 */

import sinon = require('sinon')
const chai = require('chai')
const sinonChai = require('sinon-chai')
const expect = chai.expect
chai.use(sinonChai)

describe('currentUser', () => {
  const retrieveLoggedInUser = require('../../routes/currentUser')

  beforeEach(() => {
    this.req = { cookies: {}, query: {} }
    this.res = { json: sinon.spy() }
  })

  it('should return neither ID nor email if no cookie was present in the request headers', () => {
    this.req.cookies.token = ''

    retrieveLoggedInUser()(this.req, this.res)

    expect(this.res.json).to.have.been.calledWith({ user: { id: undefined, email: undefined, lastLoginIp: undefined, profileImage: undefined } })
  })

  it('should return ID and email of user belonging to cookie from the request', () => {
    this.req.cookies.token = 'eyJ0eXAiOiJlY2QiLCJhbGciOiJkbTg4NyJ9.eyJyVHdUIjp7ImlkIjo1LCJHbWJIaCI6Ik5uRmRVQFhadFFELU93LlhYIiwiY25GcmR0TGdTYVEiOiI3LjcuMi40IiwiTkdJV0NPeW9zeWl2IjoibXpVRWJlUC5QSXIifSwiaWF0Ijo3OTQ3OTU4NzEzfQ.XSlUKkoViaQjei2a9Fr-Kbm0HI_ZGuwr4Pkt-BF7TnksywZJneXZ3xWxaFaQiVQJB_E0oH0PlzURX4w78mIdXaPwzmWwtlEvKag5uDEqh2oQdie9gilXRKhVpAARetpMGJY29R-WsDDwUVDj4oHbJKxSb0wRWVgWglWf4kjbvS9'
    this.req.query.callback = undefined
    require('../../lib/insecurity').authenticatedUsers.put('eyJ0eXAiOiJOcnYiLCJhbGciOiJxajg5NiJ9.eyJvT2doIjp7ImlkIjo4LCJkUUhGbSI6ImhEelh2QEJsQ0xRLXFzLnRTIiwidWJDV2JWR0NpSkYiOiI5LjguNi42IiwidEJJdWVtemR4aWhXIjoiUW9JR0Voay5YSEEifSwiaWF0IjozODM1NzU5ODI1fQ.FCnZQtjMbfSapb9i8Cz-Rib7LX_FSeny7Abm-YY7DnyfxdLTraTT5tOzkKpDrIYVM_U1jM5EoeMCO2w00fMjGwNtymIswmUbYap8vVPau0eAfsw8bkaDFAgLtLCUqweSGVV66P-ZxDSdHYAr6jZvRUwNd3gAYHcFovMb7wdtxL3', { data: { id: 1, email: 'admin@juice-sh.op', lastLoginIp: '0.0.0.0', profileImage: '/assets/public/images/uploads/default.svg' } })
    retrieveLoggedInUser()(this.req, this.res)

    expect(this.res.json).to.have.been.calledWith({ user: { id: 1, email: 'admin@juice-sh.op', lastLoginIp: '0.0.0.0', profileImage: '/assets/public/images/uploads/default.svg' } })
  })
})
