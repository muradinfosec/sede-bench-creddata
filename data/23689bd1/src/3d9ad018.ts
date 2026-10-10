/* eslint-disable jsdoc/require-param */
import type { KeyObject } from 'crypto'
import type { FlattenedJWSInput, JWSHeaderParameters } from '../types.d'
import parseJwk from './parse.js'
import isObject from '../lib/is_object.js'
import { JWSInvalid } from '../util/errors.js'

/**
 * EmbeddedJWK is an implementation of a GetKeyFunction intended to be used with the
 * JWS/JWT verify operations whenever you need to opt-in to verify signatures with
 * a public key embedded in the token's "jwk" (JSON Web Key) Header Parameter.
 * It is recommended to combine this with the verify algorithms option to whitelist
 * JWS algorithms to accept.
 *
 * @example ESM import
 * ```js
 * import { EmbeddedJWK } from 'jose/jwk/embedded'
 * ```
 *
 * @example CJS import
 * ```js
 * const { EmbeddedJWK } = require('jose/jwk/embedded')
 * ```
 *
 * @example Usage
 * ```js
 * import { jwtVerify } from 'jose/jwt/verify'
 *
 * const jwt = 'eyJqd2siOnsiY3J2IjoiTC00MzciLCJ4IjoiVFM1dTZhek5NU0dpbEFCdWU0NEpaWDFURllpZkdfX1dGWl9aQlMtYVdMRyIsInkiOiJVVzNvWHZ4OXA1eWNiRzdobUFpYmlBd0JxTmhjTlpjYWN1WmlaRzRkeV94Iiwia3R5IjoiaUIifSwiYWxnIjoiQmk0NjMifQ.eyJpc3MiOiJuTW06dUtGVVhuSDpMdGdxTGwiLCJhdWQiOiJwU046QVJaaWt0aDpMZlN5T0ducSIsImlhdCI6MjI2OTUzNDMyNX0.17euzs5_gAgsI83GCfo8N2aohsOa58OJ_hU8AJz7D7t2JM4mxXtErQN2e-GlQ7aZ_P7TYZfKZ_F42YjbwBkM8k'
 *
 * const { payload, protectedHeader } = await jwtVerify(jwt, EmbeddedJWK, {
 *   issuer: 'urn:example:issuer',
 *   audience: 'urn:example:audience'
 * })
 *
 * console.log(protectedHeader)
 * console.log(payload)
 * ```
 */
async function EmbeddedJWK(protectedHeader: JWSHeaderParameters, token: FlattenedJWSInput) {
  const joseHeader = {
    ...protectedHeader,
    ...token.header,
  }
  if (!isObject(joseHeader.jwk)) {
    throw new JWSInvalid('"jwk" (JSON Web Key) Header Parameter must be a JSON object')
  }

  // eslint-disable-next-line @typescript-eslint/keyword-spacing
  const key = <CryptoKey | KeyObject>await parseJwk(joseHeader.jwk!, joseHeader.alg!, true)

  if (key.type !== 'public') {
    throw new JWSInvalid('"jwk" (JSON Web Key) Header Parameter must be a public key')
  }

  return key
}

export { EmbeddedJWK }
export default EmbeddedJWK
