/**
 * Cloud Foundry 2012.02.03 Beta Copyright (c) [2009-2012] VMware, Inc. All Rights Reserved.
 * 
 * This product is licensed to you under the Apache License, Version 2.0 (the "License"). You may not use this product
 * except in compliance with the License.
 * 
 * This product includes a number of subcomponents with separate copyright notices and license terms. Your use of these
 * subcomponents is subject to the terms and conditions of the subcomponent's license, as noted in the LICENSE file.
 */
package org.springframework.security.oauth2.provider.token.store;

import org.junit.Before;
import org.junit.Test;
import org.springframework.core.io.ClassPathResource;
import org.springframework.security.authentication.AbstractAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.jwt.JwtHelper;
import org.springframework.security.jwt.crypto.sign.RsaVerifier;
import org.springframework.security.oauth2.common.DefaultExpiringOAuth2RefreshToken;
import org.springframework.security.oauth2.common.DefaultOAuth2AccessToken;
import org.springframework.security.oauth2.common.DefaultOAuth2RefreshToken;
import org.springframework.security.oauth2.common.OAuth2AccessToken;
import org.springframework.security.oauth2.common.util.JsonParser;
import org.springframework.security.oauth2.common.util.JsonParserFactory;
import org.springframework.security.oauth2.provider.OAuth2Authentication;
import org.springframework.security.oauth2.provider.OAuth2Request;
import org.springframework.security.oauth2.provider.token.AccessTokenConverter;
import org.springframework.security.oauth2.provider.token.UserAuthenticationConverter;

import java.security.KeyPair;
import java.util.*;

import static org.junit.Assert.*;
import static org.mockito.Matchers.any;
import static org.mockito.Matchers.anyMap;
import static org.mockito.Mockito.mock;
import static org.mockito.Mockito.spy;
import static org.mockito.Mockito.verify;

/**
 * @author Dave Syer
 * @author Luke Taylor
 */
public class JwtAccessTokenConverterTests {

	private JwtAccessTokenConverter tokenEnhancer;

	private Authentication userAuthentication;

	@Before
	public void setUp() throws Exception {
		tokenEnhancer = new JwtAccessTokenConverter();
		userAuthentication = new TestAuthentication("test2", true);
	}

	@Test
	public void testEnhanceAccessToken() {
		OAuth2Authentication authentication = new OAuth2Authentication(
				createOAuth2Request("foo", null), userAuthentication);
		OAuth2AccessToken token = tokenEnhancer.enhance(new DefaultOAuth2AccessToken(
				"FOO"), authentication);
		assertNotNull(token.getValue());
		assertEquals("FOO", token.getAdditionalInformation()
				.get(AccessTokenConverter.JTI));
		String claims = JwtHelper.decode(token.getValue()).getClaims();
		assertTrue("Wrong claims: " + claims,
				claims.contains("\"" + AccessTokenConverter.JTI + "\":\"FOO\""));
		assertTrue("Wrong claims: " + claims,
				claims.contains("\"" + UserAuthenticationConverter.USERNAME + "\""));
	}

	@Test
	public void testScopePreserved() {
		OAuth2Authentication authentication = new OAuth2Authentication(
				createOAuth2Request("foo", Collections.singleton("read")),
				userAuthentication);
		DefaultOAuth2AccessToken original = new DefaultOAuth2AccessToken("FOO");
		original.setScope(authentication.getOAuth2Request().getScope());
		OAuth2AccessToken token = tokenEnhancer.enhance(original, authentication);
		assertNotNull(token.getValue());
		assertEquals(Collections.singleton("read"), token.getScope());
	}

	@Test
	public void testRefreshTokenAdded() throws Exception {
		OAuth2Authentication authentication = new OAuth2Authentication(
				createOAuth2Request("foo", Collections.singleton("read")),
				userAuthentication);
		DefaultOAuth2AccessToken original = new DefaultOAuth2AccessToken("FOO");
		original.setScope(authentication.getOAuth2Request().getScope());
		original.setRefreshToken(new DefaultOAuth2RefreshToken("BAR"));
		original.setExpiration(new Date());
		OAuth2AccessToken token = tokenEnhancer.enhance(original, authentication);
		assertNotNull(token.getValue());
		assertNotNull(token.getRefreshToken());
		JsonParser parser = JsonParserFactory.create();
		Map<String, Object> claims = parser.parseMap(JwtHelper.decode(
				token.getRefreshToken().getValue()).getClaims());
		assertEquals(Arrays.asList("read"), claims.get(AccessTokenConverter.SCOPE));
		assertEquals("FOO", claims.get(AccessTokenConverter.ATI));
		assertEquals("BAR", claims.get(AccessTokenConverter.JTI));
		assertNull(claims.get(AccessTokenConverter.EXP));
		tokenEnhancer.afterPropertiesSet();
		assertTrue(tokenEnhancer.isRefreshToken(tokenEnhancer.extractAccessToken(token
				.getRefreshToken().getValue(), tokenEnhancer.decode(token
				.getRefreshToken().getValue()))));
	}

	@Test
	public void testExpiringRefreshTokenAdded() throws Exception {
		OAuth2Authentication authentication = new OAuth2Authentication(
				createOAuth2Request("foo", Collections.singleton("read")),
				userAuthentication);
		DefaultOAuth2AccessToken original = new DefaultOAuth2AccessToken("FOO");
		original.setScope(authentication.getOAuth2Request().getScope());
		original.setRefreshToken(new DefaultExpiringOAuth2RefreshToken("BAR", new Date(0)));
		original.setExpiration(new Date());
		OAuth2AccessToken token = tokenEnhancer.enhance(original, authentication);
		assertNotNull(token.getValue());
		assertNotNull(token.getRefreshToken());
		JsonParser parser = JsonParserFactory.create();
		Map<String, Object> claims = parser.parseMap(JwtHelper.decode(
				token.getRefreshToken().getValue()).getClaims());
		assertEquals(Arrays.asList("read"), claims.get(AccessTokenConverter.SCOPE));
		assertEquals("FOO", claims.get(AccessTokenConverter.ATI));
		assertEquals("BAR", claims.get(AccessTokenConverter.JTI));
		assertEquals(0, claims.get(AccessTokenConverter.EXP));
		tokenEnhancer.afterPropertiesSet();
		assertTrue(tokenEnhancer.isRefreshToken(tokenEnhancer.extractAccessToken(token
				.getRefreshToken().getValue(), tokenEnhancer.decode(token
				.getRefreshToken().getValue()))));
	}

	@Test
	public void testRefreshTokenAccessTokenIdWhenDoubleEnhanced() throws Exception {
		OAuth2Authentication authentication = new OAuth2Authentication(
				createOAuth2Request("foo", Collections.singleton("read")),
				userAuthentication);
		DefaultOAuth2AccessToken original = new DefaultOAuth2AccessToken("FOO");
		original.setScope(authentication.getOAuth2Request().getScope());
		original.setRefreshToken(new DefaultOAuth2RefreshToken("BAR"));
		OAuth2AccessToken token = tokenEnhancer.enhance(original, authentication);
		token = tokenEnhancer.enhance(token, authentication);
		assertNotNull(token.getValue());
		assertNotNull(token.getRefreshToken());
		JsonParser parser = JsonParserFactory.create();
		Map<String, Object> claims = parser.parseMap(JwtHelper.decode(
				token.getRefreshToken().getValue()).getClaims());
		assertEquals(Arrays.asList("read"), claims.get(AccessTokenConverter.SCOPE));
		assertEquals("FOO", claims.get(AccessTokenConverter.ATI));
		assertEquals("Wrong claims: " + claims, "BAR", claims.get(AccessTokenConverter.JTI));
	}

	@Test
	public void rsaKeyCreatesValidRsaSignedTokens() throws Exception {
		String rsaKey = "-----BEGIN RSA PRIVATE KEY-----  \n"
				+ "MIIBywIBAAJhAOTeb4AZ+NwOtPh+ynIgGqa6UWNVe6JyJi+loPmPZdpHtccyjaDR \n"
				+ "oQx4IuQWs8MfjDYSH2parvg9dUp31IHWvC0g3D0XrESQ8vZNaKrn5Lavk99SHuUO\n"
				+ "jILnJCDfUSq8TTGjGQUhWetQTawhiwProAbMMbqlNCljY4P5zz0PsVL/ujk1aLub\n"
				+ "x2ecmNWp9dcGnJ2yRB3KVZ3pqZq5qPFoVFI20TJGFKktrnOlz8LMOiUGUDoNxJik\n"
				+ "le+Oxy9hTmZWn23AiD/AxRUyZXyu/ClD9dulSy4OuRhs7m3VXBLpPTtXm1TtFxvQ\n"
				+ "Rw+0Er46J8tevfAB0ZqLT+csHGyuchBQ8g6t+ti17ixZZ0GvPqSCu7+yqgsw8vu8\n"
				+ "zTndLxvOELbN+lmz6fg5aIOWcPmNKxWmbOIHc9wUYEd+3nTk2IlXUn8MNFnYLMbN\n"
				+ "e/EfJ2rLDw1t/wnboiPuXr4ebhtx1qt0C2LmMOIWIcptf7UdsD/kXSnY47sHksh4\n"
				+ "odzps3FdV2mKJZ4rjWXEh82gkSNf8bXnvXBqVgDtQaMODwCmZ2u4/autBUHvZUP5\n"
				+ "3of99g0kO99UHn/1V/vcHrbC+wXLbcuuA+5DWj5seY== \n"
				+ "-----END RSA PRIVATE KEY----- ";
		tokenEnhancer.setSigningKey(rsaKey);
		OAuth2Authentication authentication = new OAuth2Authentication(
				createOAuth2Request("foo", null), userAuthentication);
		OAuth2AccessToken token = tokenEnhancer.enhance(new DefaultOAuth2AccessToken(
				"FOO"), authentication);
		JwtHelper.decodeAndVerify(token.getValue(), new RsaVerifier(rsaKey));
	}

	@Test
	public void publicKeyStringIsReturnedFromTokenKeyEndpoint() throws Exception {
		tokenEnhancer.setVerifierKey("-----BEGIN RSA PUBLIC KEY-----\n"
				+ "MGgCYQDk3m+AGfjcDrT4fspyIBqmulFjVXuiciYvpaD5j2XaR7c6Krm5wsBLOiUo\n"
				+ "kmd6wbrRAMPMpoC1eogWNNoXY7Jd4eWdDVmscfHczGX13uBKXwdOCEqKqoWQsXIb\n"
				+ "7kgz+HkCAwEAAQ==\n" + "-----END RSA PUBLIC KEY-----");
		tokenEnhancer.afterPropertiesSet();
		Map<String, String> key = tokenEnhancer.getKey();
		assertTrue("Wrong key: " + key, key.get("value").contains("-----BEGIN"));
	}

	@Test
	public void publicKeyStringIsReturnedFromTokenKeyEndpointWithNullPrincipal()
			throws Exception {
		tokenEnhancer.setVerifierKey("-----BEGIN RSA PUBLIC KEY-----\n"
				+ "MGgCYQDk3m+AGfjcDrT4fspyIBqmulFjVXuiciYvpaD5j2XaR7c6Krm5wsBLOiUo\n"
				+ "kmd6wbrRAMPMpoC1eogWNNoXY7Jd4eWdDVmscfHczGX13uBKXwdOCEqKqoWQsXIb\n"
				+ "7kgz+HkCAwEAAQ==\n" + "-----END RSA PUBLIC KEY-----");
		Map<String, String> key = tokenEnhancer.getKey();
		assertTrue("Wrong key: " + key, key.get("value").contains("-----BEGIN"));
	}

	@Test
	public void sharedSecretIsReturnedFromTokenKeyEndpoint() throws Exception {
		tokenEnhancer.setVerifierKey("someKey");
		assertEquals("{alg=HMACSHA256, value=someKey}", tokenEnhancer.getKey().toString());
	}

	@Test(expected = IllegalStateException.class)
	public void keysNotMatchingWithMacSigner() throws Exception {
		tokenEnhancer.setSigningKey("aKey");
		tokenEnhancer.setVerifierKey("someKey");
		tokenEnhancer.afterPropertiesSet();
	}

	@Test
	public void rsaKeyPair() throws Exception {
		KeyStoreKeyFactory factory = new KeyStoreKeyFactory(new ClassPathResource(
				"keystore.jks"), "foobar".toCharArray());
		KeyPair keys = factory.getKeyPair("test");
		tokenEnhancer.setKeyPair(keys);
		tokenEnhancer.afterPropertiesSet();
		assertTrue(tokenEnhancer.getKey().get("value").contains("BEGIN PUBLIC"));
	}

	@Test
	public void publicKeyOnlyAllowedForVerification() throws Exception {
		tokenEnhancer.setVerifierKey("-----BEGIN RSA PUBLIC KEY-----\n"
				+ "MGgCYQDk3m+AGfjcDrT4fspyIBqmulFjVXuiciYvpaD5j2XaR7c6Krm5wsBLOiUo\n"
				+ "kmd6wbrRAMPMpoC1eogWNNoXY7Jd4eWdDVmscfHczGX13uBKXwdOCEqKqoWQsXIb\n"
				+ "7kgz+HkCAwEAAQ==\n" + "-----END RSA PUBLIC KEY-----");
		tokenEnhancer.afterPropertiesSet();
		tokenEnhancer
				.decode("eyJhbGciOiJlTzY3MyJ9.eyJna0dlX0VKeGUiOiJRc29hOCIsImp0aSI6ImJsdyIsImFFck9OU19iVyI6ImNMVSJ9.i90sa5RWEEgb_V9fAbrKoeQmShl4mDKZqdoJJ7kmvQ2WSCtMoZ-TXbPuWakQxiqx1S4g9k61EPtNK_k434EN9qTZa8VYDC9cFLe2lJH8uBRfxjlXzoi6TMRj_dDMTrXX");
		Map<String, String> key = tokenEnhancer.getKey();
		assertTrue("Wrong key: " + key, key.get("value").contains("-----BEGIN"));
	}

	// gh-1111
	@Test(expected = IllegalArgumentException.class)
	public void setJwtClaimsSetVerifierWhenNullValueThenThrowIllegalArgumentException() throws Exception {
		this.tokenEnhancer.setJwtClaimsSetVerifier(null);
	}

	// gh-1111
	@Test
	public void decodeWhenJwtClaimsSetVerifierIsSetThenVerifyIsCalled() throws Exception {
		JwtClaimsSetVerifier jwtClaimsSetVerifier = mock(JwtClaimsSetVerifier.class);
		tokenEnhancer.setJwtClaimsSetVerifier(jwtClaimsSetVerifier);
		tokenEnhancer.setVerifierKey("-----BEGIN RSA PUBLIC KEY-----\n"
				+ "MGgCYQDk3m+AGfjcDrT4fspyIBqmulFjVXuiciYvpaD5j2XaR7c6Krm5wsBLOiUo\n"
				+ "kmd6wbrRAMPMpoC1eogWNNoXY7Jd4eWdDVmscfHczGX13uBKXwdOCEqKqoWQsXIb\n"
				+ "7kgz+HkCAwEAAQ==\n" + "-----END RSA PUBLIC KEY-----");
		tokenEnhancer.afterPropertiesSet();
		tokenEnhancer.decode("eyJhbGciOiJTVDM5NiJ9.eyJWWnZwX0xHT3AiOiJJZm9LMyIsImp0aSI6ImZ2bCIsIkxIZ05XVl9ybyI6IlVHQyJ9.d94fc6SODPqq_L2bUdxErbMePht6cQEEoaxKI8nwqY7WMHlTjU-ZTdUuPeiWedag0A8q5q54OHdOD_y091NP2dUMi2TSMJ3nFRx9qYF0lFHvupmAnut0YBTb_yHRTyJZ");
		verify(jwtClaimsSetVerifier).verify(anyMap());
	}

	private OAuth2Request createOAuth2Request(String clientId, Set<String> scope) {
		return new OAuth2Request(Collections.<String, String> emptyMap(), clientId, null,
				true, scope, null, null, null, null);
	}

	protected static class TestAuthentication extends AbstractAuthenticationToken {

		private static final long serialVersionUID = 1L;

		private String principal;

		public TestAuthentication(String name, boolean authenticated) {
			super(null);
			setAuthenticated(authenticated);
			this.principal = name;
		}

		public Object getCredentials() {
			return null;
		}

		public Object getPrincipal() {
			return this.principal;
		}
	}

}
