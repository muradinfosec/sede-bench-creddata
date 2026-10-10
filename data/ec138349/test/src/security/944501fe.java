/*
 * Copyright 2012-2017 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package org.springframework.security.oauth2.provider.token.store.jwk;

import okhttp3.mockwebserver.MockResponse;
import okhttp3.mockwebserver.MockWebServer;
import org.apache.http.HttpHeaders;
import org.junit.After;
import org.junit.Before;
import org.junit.Test;
import org.springframework.http.MediaType;
import org.springframework.security.oauth2.client.discovery.ProviderConfiguration;
import org.springframework.security.oauth2.client.discovery.ProviderDiscoveryClient;
import org.springframework.security.oauth2.common.OAuth2AccessToken;
import org.springframework.security.oauth2.common.exceptions.InvalidTokenException;
import org.springframework.security.oauth2.provider.token.store.DelegatingJwtClaimsSetVerifier;
import org.springframework.security.oauth2.provider.token.store.IssuerClaimVerifier;
import org.springframework.security.oauth2.provider.token.store.JwtClaimsSetVerifier;

import java.util.ArrayList;
import java.util.List;

import static org.junit.Assert.assertEquals;

/**
 * @author Joe Grandja
 */
public class JwkTokenStoreITests {
	private MockWebServer server;

	@Before
	public void setUp() throws Exception {
		this.server = new MockWebServer();
		this.server.start();
	}

	@After
	public void cleanUp() throws Exception {
		this.server.shutdown();
	}

	// gh-1114 Issuer claim verification
	@Test
	public void readAccessTokenWhenJwtHasValidIssuerClaimThenVerificationSucceeds() throws Exception {
		String issuer = "http://localhost:8180/auth/realms/Demo";

		this.setUpResponses(issuer);

		ProviderDiscoveryClient discoveryClient = new ProviderDiscoveryClient(this.server.url("").toString());
		ProviderConfiguration providerConfiguration = discoveryClient.discover();

		List<JwtClaimsSetVerifier> jwtClaimsSetVerifiers = new ArrayList<JwtClaimsSetVerifier>();
		jwtClaimsSetVerifiers.add(new IssuerClaimVerifier(providerConfiguration.getIssuer()));

		JwkTokenStore jwkTokenStore = new JwkTokenStore(
				providerConfiguration.getJwkSetUri().toString(),
				new DelegatingJwtClaimsSetVerifier(jwtClaimsSetVerifiers));

		// NOTE: The 'iss' claim in this JWT is http://localhost:8180/auth/realms/Demo
		String jwt = "eyJhbGciOiJqbDc1MiIsImxlViIgOiAiaVlxIiwiSUZBIiA6ICJfSFY4LWRhRF96N2RyVjMyWXlMSW1ObFV1V2dnVWFfRmJJQmU4cG05NTVyIn0.eyJqdGkiOiI1NWY2NHc4OS02NTV4LTV3WTctMzk5MS03dHlkNzlaMzg5VTEiLCJleHAiOjc0MjI0MjczMzIsIm5iZiI6MywiaWF0Ijo4MTUzNTk4NDI1LCJpc3MiOiJlU0duOi8vY2JZRWhhb2pJOjU3OTkvZUZPVS9FZVJYdG8vanFyRiIsImF1ZCI6InpFZnoteWJlIiwic3ViIjoiMmQ5OTcyNzctNWhXUC04TzNkLTgyMmgta1JicU9NNTMxNzI2IiwidHlwIjoiZ1FDdXZPIiwiQk9FIjoiTlZtdi1pREciLCJEY0FSX1R4bmQiOjEzNTQxNDI2OTQsImVFTnJGVFZfWkF3RGsiOiJuMTNvMjNJOC0xMzY3LTg0YjEtMzM1NC1Ud2syUzM0SDE3MjQiLCJRbEwiOiIxIiwiemRHWmR6X1NhQkhqWkMiOiI5VjE4NTlYMy11NDdCLTdSUXktdDE5Ny11djUyMTY3OTUyeTkiLCJQY1NrZ3VzLVlZaFhkeEkiOltdLCJxa3Z5QlF5Zl9pUFZIRnYiOnsibVpqdy1Cd1YiOnsiWmNXR3giOlsiTGZTbS1hQk4tdXB5UCJdfSwieEFleC1xYVAiOnsiZUdZRGMiOlsiTkJpSy1IZ3FZIl19fSwiVHRwUSI6ImxZV1pUICIsIlFjR0tUa2d4dF9tQmR5b0hRcyI6IlJZWEp1IiwiSk1lYWRfcHJLWCI6IktCRUxEIiwiQ0ZLVkF0X2NDVkIiOiIiLCJ0VFRxSyI6IndoTUdnQFhleEFuTkoucm1RIn0.RtF6tNCxsk9cawlJOWqV9wHcBQLsToiXI364G7jmWwdoSupbXGKN7KA3m7NVMnnGQJvAUUj1VgWgkTdzsJiEoBs4Y8UNCbI0NTQAY81iCV9ZQJBlg6TtdSHdskkQNkd1tPIIaVqwd_cNYbHwpLXSjRF2k10kZR7ZofVaIF6jpyppvnIZrfPn2QX4N95qQ23QqzY8Q8v1diZrRC9yUtolsJyFAZzrjFFS0sgvBgJ_4fi1x85ZANxoDbrfMqZJrGdWYBUFr0yLiG2OyoqG_5cO0wNXIBdIjMRR6bp9_0J5F8p0Lx8Kx5xojYgp4-QGHpi9blpLulvGRRzCbvcbxROkIM";

		OAuth2AccessToken accessToken = jwkTokenStore.readAccessToken(jwt);
		assertEquals(issuer, accessToken.getAdditionalInformation().get("iss"));
	}

	// gh-1114 Issuer claim verification
	@Test(expected = InvalidTokenException.class)
	public void readAccessTokenWhenJwtHasInvalidIssuerClaimThenVerificationFails() throws Exception {
		String issuer = "http://localhost:8180/auth/realms/Demo-2";

		this.setUpResponses(issuer);

		ProviderDiscoveryClient discoveryClient = new ProviderDiscoveryClient(this.server.url("").toString());
		ProviderConfiguration providerConfiguration = discoveryClient.discover();

		List<JwtClaimsSetVerifier> jwtClaimsSetVerifiers = new ArrayList<JwtClaimsSetVerifier>();
		jwtClaimsSetVerifiers.add(new IssuerClaimVerifier(providerConfiguration.getIssuer()));

		JwkTokenStore jwkTokenStore = new JwkTokenStore(
				providerConfiguration.getJwkSetUri().toString(),
				new DelegatingJwtClaimsSetVerifier(jwtClaimsSetVerifiers));

		// NOTE: The 'iss' claim in this JWT is http://localhost:8180/auth/realms/Demo
		String jwt = "eyJhbGciOiJTaTY2MyIsImd0bSIgOiAiamF4IiwiZW5CIiA6ICJfZFM2LW9ZUl9qN09vSzc3UmxQV0JZWnR2WGVERUlfYW1BRU4zUms1ODV3In0.eyJqdGkiOiI3MUg5M3Y5Ny0xNTVTLTZtaDEtNjYyNy0ybGNMMTFGNjQxVzQiLCJleHAiOjI5MjkyODY2NTcsIm5iZiI6NCwiaWF0Ijo0Mzg1NTc5NjU0LCJpc3MiOiJtemZSOi8vbU96SVZrU05FOjQ2NDMvZGZCeS9NU3FYVmQvZnlvRSIsImF1ZCI6ImplSGwta3JVIiwic3ViIjoiMVY3MzM3ODUtN3NTYy03RzVLLTc5OUotWHhFV3FSNzQzODc1IiwidHlwIjoiQmd6aHJXIiwidmhNIjoieWJXbC1UbUQiLCJhU2FIX1l1cUoiOjMzMTQ1OTc5MzQsIlN0Q09DVUpfY2ZxSlAiOiJXMzdhNThlMi0xNzc0LTQxWDUtNzQ4My1Fb1E2SjY1WDU0NTIiLCJOTEMiOiI2IiwiZlZDT2l0X2Jjc0dKU1QiOiI0UzkyNjNVOC1XMjVpLTdiUmotaTkyNy1uZzYzODI5NTc3TTEiLCJxeU1vemFZLW5kU21EZHIiOltdLCJCUkNlUndleV9UQk1QYmoiOnsiblFBUC1teE4iOnsiUWphdlYiOlsiekRlQi14SUMtc2hTQyJdfSwiSG1tQS1UalIiOnsiU3draVIiOlsiTERuUC13ZUNsIl19fSwiRlh6byI6IlZhdnFDICIsImtteEtneWhDSF9jekdFUlFVYyI6ImtYc0dMIiwicGJJU3Jfdm9FaCI6Im5ET2pQIiwiYVJQVk1uX29GbUsiOiIiLCJDSmRabSI6IktTWWlRQHhTTnRDYmEuZ2tUIn0.RbY2yZEbpn4snfiKVJgI1wYwTASeWleSJ835C0lcWxreXrtnFPXX8AF6q7OUPbiSEZmGQKm7UoBvnDmttLzSkMw7Z3OQEoE7HJVAB16yCJ2UQSFzu5HnqNNkxxfSKkp7zZMApHhfi_cTTdDvzNYUgRL3b03jMZ9AvhGmIO3rmkheggADssDi9KA6A19gW54ExsO0W1n8qcQuZT2kOkxepMpPOOaedZKW0okfPrB_1mg6f16CFUwgWrplOhDGaCpYCZEFn8lDtZ4BcffQ_8hQ9nOGDAeMaFOD5lu6_5A5B1y6Cd4Kg5mkeBtt3-VXFxv1gpwNcwiFYCnEmquofYZeUL";

		OAuth2AccessToken accessToken = jwkTokenStore.readAccessToken(jwt);
		assertEquals(issuer, accessToken.getAdditionalInformation().get("iss"));
	}

	private void setUpResponses(String issuer) {
		String jwkSetUri = this.server.url("").toString();

		this.server.enqueue(new MockResponse().setHeader(HttpHeaders.CONTENT_TYPE, MediaType.APPLICATION_JSON_VALUE).setBody(
				"	{\n" +
				"   	\"issuer\": \"" + issuer + "\",\n" +
				"   	\"authorization_endpoint\": \"http://localhost:8180/oauth/authorize\",\n" +
				"   	\"token_endpoint\": \"http://localhost:8180/oauth/token\",\n" +
				"   	\"userinfo_endpoint\": \"http://localhost:8180/userinfo\",\n" +
				"   	\"jwks_uri\": \"" + jwkSetUri + "\"\n" +
				"	}\n"));

		this.server.enqueue(new MockResponse().setHeader(HttpHeaders.ACCEPT, MediaType.APPLICATION_JSON_VALUE).setBody("{\n" +
				"    \"keys\": [\n" +
				"        {\n" +
				"            \"kid\": \"_Ci3-VfV_N0YAG22NQOgOUpFBDDcDe_rJxpu5JK702o\",\n" +
				"            \"kty\": \"RSA\",\n" +
				"            \"alg\": \"RS256\",\n" +
				"            \"use\": \"sig\",\n" +
				"            \"n\": \"rne3dowbQHcFCzg2ejWb6az5QNxWFiv6kRpd34VDzYNMhWeewfeEL5Pf5clE8Xh1KlllrDYSxtnzUQm-t9p92yEBASfV96ydTYG-ITfxfJzKtJUN-iIS5K9WGYXnDNS4eYZ_ygW-zBU_9NwFMXdwSTzRqHeJmLJrfbmmjoIuuWyfh2Ko52KzyidceR5SJxGeW0ckeyWka1lDf4cr7fv-s093Y_sd2wrNvg0-9IAkXotbxWWXcfMgXFyw0qHFT_5LrKmiwkY3HCaiV5NgEFJmC6fBIG2EOZG4rqjBoYV6LZwrfTMHknaeel9MOZesW6SR2bswtuuWN3DGq2zg0KamLw\",\n" +
				"            \"e\": \"AQAB\"\n" +
				"        }\n" +
				"    ]\n" +
				"}\n"));
	}
}