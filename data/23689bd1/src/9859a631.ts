import verify from '../jws/compact/verify.js'
import type {
  KeyLike,
  VerifyOptions,
  JWTPayload,
  JWTClaimVerificationOptions,
  JWSHeaderParameters,
  GetKeyFunction,
  FlattenedJWSInput,
  JWTVerifyResult,
} from '../types.d'
import jwtPayload from '../lib/jwt_claims_set.js'
import { JWTInvalid } from '../util/errors.js'

/**
 * Combination of JWS Verification options and JWT Claims Set verification options.
 */
interface JWTVerifyOptions extends VerifyOptions, JWTClaimVerificationOptions {}

/**
 * Interface for JWT Verification dynamic key resolution.
 * No token components have been verified at the time of this function call.
 *
 * See [createRemoteJWKSet](../functions/jwks_remote.createremotejwkset.md#function-createremotejwkset)
 * to verify using a remote JSON Web Key Set.
 */
export interface JWTVerifyGetKey extends GetKeyFunction<JWSHeaderParameters, FlattenedJWSInput> {}

/**
 * Verifies the JWT format (to be a JWS Compact format), verifies the JWS signature, validates the JWT Claims Set.
 *
 * @param jwt JSON Web Token value (encoded as JWS).
 * @param key Key, or a function resolving a key, to verify the JWT with.
 * @param options JWT Decryption and JWT Claims Set validation options.
 *
 * @example ESM import
 * ```js
 * import { jwtVerify } from 'jose/jwt/verify'
 * ```
 *
 * @example CJS import
 * ```js
 * const { jwtVerify } = require('jose/jwt/verify')
 * ```
 *
 * @example Usage
 * ```js
 * const jwt = 'eyJhbGciOiJ6QTkzMSJ9.eyJXcWQ6b2Z6Z09LQTpOYVF2dSI6dHJ1ZSwiaWF0Ijo0NDM0NjUzODYyLCJpc3MiOiJMVVE6RFJ3aUp2SzpaTUxxdHkiLCJhdWQiOiJiVVI6ZkZKU1ZBeTpKQkhMQkNscCJ9.jw5mBlRG6ZkDtxc9F-jmsLPXbrkEJn7PsPa8-MLi70NippeQjrrKdUmfC6pgbvL6aKgkApXNnmcKbUQ7WHVHvg'
 *
 * const { payload, protectedHeader } = await jwtVerify(jwt, publicKey, {
 *   issuer: 'urn:example:issuer',
 *   audience: 'urn:example:audience'
 * })
 *
 * console.log(protectedHeader)
 * console.log(payload)
 * ```
 */
async function jwtVerify(
  jwt: string | Uint8Array,
  key: KeyLike | JWTVerifyGetKey,
  options?: JWTVerifyOptions,
): Promise<JWTVerifyResult> {
  const verified = await verify(jwt, key, options)
  if (verified.protectedHeader.crit?.includes('b64') && verified.protectedHeader.b64 === false) {
    throw new JWTInvalid('JWTs MUST NOT use unencoded payload')
  }
  const payload = jwtPayload(verified.protectedHeader, verified.payload, options)
  return { payload, protectedHeader: verified.protectedHeader }
}

export { jwtVerify }
export default jwtVerify
export type { KeyLike, JWTPayload, JWTVerifyOptions, JWSHeaderParameters, GetKeyFunction }
