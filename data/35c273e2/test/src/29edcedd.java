// Copyright (c) Microsoft. All rights reserved.
// Licensed under the MIT license. See License.txt in the project root.

package com.microsoft.alm.authentication;

import com.microsoft.alm.helpers.Action;
import com.microsoft.alm.helpers.NullUserAgent;
import com.microsoft.alm.helpers.StringContent;
import com.microsoft.alm.oauth2.useragent.AuthorizationException;
import com.microsoft.alm.secret.TokenPair;
import org.junit.Assert;
import org.junit.Test;

import java.net.URI;
import java.util.UUID;
import java.util.concurrent.atomic.AtomicInteger;

public class AzureAuthorityTest
{
    static final String TEST_RESOURCE = "TEST_RESOURCE";
    static final String TEST_CLIENT_ID = "a41cecfa-1ce3-8a55-fd12-55ecf8818650";
    static final String TEST_DEVICE_CODE = "06d7b9c3-f1fd-3ac0-02e4-178c4633ef6a";
    static final String TEST_USER_CODE = "DEADBEEF";
    static final URI TEST_REDIRECT_URI = URI.create("https://terminus.example.com");
    static final int TEST_EXPIRATION = 600;
    static final int TEST_INTERVAL = 5;
    static final String TEST_ACCESS_TOKEN = "aeed4a3a-87c5-7041-3799-b98ea1ae8f01";
    static final String TEST_REFRESH_TOKEN = "b7ec0b80-9f5f-473e-5373-40552678369a";

    @Test
    public void deviceFlow_success() throws Exception
    {
        final String authorityHostUrl = "https://authorization.example.com/common/";
        final URI targetUri = URI.create("https://resource.example.com/");
        final URI verificationUri = URI.create("https://authorization.example.com/oauth/device");
        final AtomicInteger requestAuthorizationCalls = new AtomicInteger(0);
        final AtomicInteger requestTokenCalls = new AtomicInteger(0);
        final AtomicInteger callbackCalls = new AtomicInteger(0);
        final AzureDeviceFlowResponse azureDeviceFlowResponse = new AzureDeviceFlowResponse(TEST_DEVICE_CODE, TEST_USER_CODE, verificationUri, TEST_EXPIRATION, TEST_INTERVAL, "message");
        final AzureDeviceFlow testDeviceFlow = new AzureDeviceFlow()
        {
            @Override public DeviceFlowResponse requestAuthorization(final URI deviceEndpoint, final String clientId, final String scope)
            {
                requestAuthorizationCalls.addAndGet(1);
                Assert.assertEquals(TEST_CLIENT_ID, clientId);
                return azureDeviceFlowResponse;
            }

            @Override public TokenPair requestToken(final URI tokenEndpoint, final String clientId, final DeviceFlowResponse deviceFlowResponse) throws AuthorizationException
            {
                requestTokenCalls.addAndGet(1);
                Assert.assertEquals(TEST_CLIENT_ID, clientId);
                Assert.assertEquals(azureDeviceFlowResponse, deviceFlowResponse);
                return new TokenPair(TEST_ACCESS_TOKEN, TEST_REFRESH_TOKEN);
            }
        };
        final Action<DeviceFlowResponse> callback = new Action<DeviceFlowResponse>()
        {
            @Override public void call(final DeviceFlowResponse deviceFlowResponse)
            {
                callbackCalls.addAndGet(1);
                Assert.assertEquals(azureDeviceFlowResponse, deviceFlowResponse);
            }
        };
        final AzureAuthority cut = new AzureAuthority(authorityHostUrl, NullUserAgent.INSTANCE, testDeviceFlow);

        final TokenPair actualTokenPair = cut.acquireToken(targetUri, TEST_CLIENT_ID, TEST_RESOURCE, TEST_REDIRECT_URI, callback);

        Assert.assertEquals(TEST_ACCESS_TOKEN, actualTokenPair.AccessToken.Value);
        Assert.assertEquals(TEST_REFRESH_TOKEN, actualTokenPair.RefreshToken.Value);
        Assert.assertEquals(TEST_RESOURCE, testDeviceFlow.getResource());
        Assert.assertEquals(TEST_REDIRECT_URI, testDeviceFlow.getRedirectUri());
        Assert.assertEquals(1, requestAuthorizationCalls.get());
        Assert.assertEquals(1, requestTokenCalls.get());
        Assert.assertEquals(1, callbackCalls.get());
    }

    @Test
    public void deviceFlow_failure() throws Exception
    {
        final String authorityHostUrl = "https://authorization.example.com/common/";
        final URI targetUri = URI.create("https://resource.example.com/");
        final URI verificationUri = URI.create("https://authorization.example.com/oauth/device");
        final AtomicInteger requestAuthorizationCalls = new AtomicInteger(0);
        final AtomicInteger requestTokenCalls = new AtomicInteger(0);
        final AtomicInteger callbackCalls = new AtomicInteger(0);
        final AzureDeviceFlowResponse azureDeviceFlowResponse = new AzureDeviceFlowResponse(TEST_DEVICE_CODE, TEST_USER_CODE, verificationUri, TEST_EXPIRATION, TEST_INTERVAL, "message");
        final AzureDeviceFlow testDeviceFlow = new AzureDeviceFlow()
        {
            @Override public DeviceFlowResponse requestAuthorization(final URI deviceEndpoint, final String clientId, final String scope)
            {
                requestAuthorizationCalls.addAndGet(1);
                Assert.assertEquals(TEST_CLIENT_ID, clientId);
                return azureDeviceFlowResponse;
            }

            @Override public TokenPair requestToken(final URI tokenEndpoint, final String clientId, final DeviceFlowResponse deviceFlowResponse) throws AuthorizationException
            {
                requestTokenCalls.addAndGet(1);
                throw new AuthorizationException("access_denied");
            }
        };
        final Action<DeviceFlowResponse> callback = new Action<DeviceFlowResponse>()
        {
            @Override public void call(final DeviceFlowResponse deviceFlowResponse)
            {
                callbackCalls.addAndGet(1);
                Assert.assertEquals(azureDeviceFlowResponse, deviceFlowResponse);
            }
        };
        final AzureAuthority cut = new AzureAuthority(authorityHostUrl, NullUserAgent.INSTANCE, testDeviceFlow);

        final TokenPair actualTokenPair = cut.acquireToken(targetUri, TEST_CLIENT_ID, TEST_RESOURCE, TEST_REDIRECT_URI, callback);

        Assert.assertEquals(null, actualTokenPair);
        Assert.assertEquals(TEST_RESOURCE, testDeviceFlow.getResource());
        Assert.assertEquals(TEST_REDIRECT_URI, testDeviceFlow.getRedirectUri());
        Assert.assertEquals(1, requestAuthorizationCalls.get());
        Assert.assertEquals(1, requestTokenCalls.get());
        Assert.assertEquals(1, callbackCalls.get());
    }

    @Test
    public void createAuthorizationEndpointUri_minimal() throws Exception
    {
        final URI redirectUri = URI.create("https://example.com");
        final URI actual = AzureAuthority.createAuthorizationEndpointUri(
                "https://login.microsoftonline.com/common",
                "e5316b9b-cf7b-2cff-a41c-6ddfea42d8cc", "f6b91cbf-a10f-2798-fd04-5baa9ddd7170",
                redirectUri,
                UserIdentifier.ANY_USER,
                null,
                PromptBehavior.AUTO,
                null);

        Assert.assertEquals("https://login.microsoftonline.com/common/oauth2/authorize?resource=e8081d1a-fc3f-5fdb-a99d-0ebdfd92c7ad&client_id=b0d96ccd-c77c-1072-ca81-0cba1cda7833&response_type=code&redirect_uri=https%3A%2F%2Fexample.com", actual.toString());
    }

    @Test
    public void createAuthorizationEndpointUri_typical() throws Exception
    {
        final URI redirectUri = URI.create("https://example.com");
        final String state = "503d6ee6-b41a-3278-0951-9a40278555d0";
        final URI actual = AzureAuthority.createAuthorizationEndpointUri(
                "https://login.microsoftonline.com/common",
                "e8375b6d-bc8d-1adc-e19b-3effeb60d9dd", "b1c65bbd-a61b-9229-da22-2fde7dfd5682",
                redirectUri,
                UserIdentifier.ANY_USER,
                state,
                PromptBehavior.ALWAYS,
                null);

        Assert.assertEquals("https://login.microsoftonline.com/common/oauth2/authorize?resource=e0789a8d-ff1a-5ecf-e24d-8daabe34e4fe&client_id=d3d30daa-d95b-3006-cd27-9ffe2fcf6270&response_type=code&redirect_uri=https%3A%2F%2Fexample.com&state=677f3ed4-e98e-7437-0900-7c65113894c5&prompt=login", actual.toString());
    }

    @Test
    public void createAuthorizationEndpointUri_extraState() throws Exception
    {
        final URI redirectUri = URI.create("https://example.com");
        final String state = "683a6ab8-e39b-3228-0642-5c73808167b7";
        final URI actual = AzureAuthority.createAuthorizationEndpointUri(
                "https://login.microsoftonline.com/common",
                "b6087c0e-bb6c-8fec-c23d-5feafa96f5ac", "e1c08dbf-c05e-0769-ea35-2dfe9bbb1553",
                redirectUri,
                UserIdentifier.ANY_USER,
                state,
                PromptBehavior.ALWAYS,
                "state=bliss");

        Assert.assertEquals("https://login.microsoftonline.com/common/oauth2/authorize?resource=d1865f9c-cc1f-0bdb-d00d-1fccec32b2ea&client_id=d2d93adc-b56d-2662-ba47-5dff3ddc9110&response_type=code&redirect_uri=https%3A%2F%2Fexample.com&state=878f7cc4-c29b-3083-9393-0c59271952b3&prompt=login&state=bliss", actual.toString());
    }

    @Test
    public void createTokenEndpointUri_typical() throws Exception
    {
        final URI actual = AzureAuthority.createTokenEndpointUri("https://login.example.com/common");

        Assert.assertEquals("https://login.example.com/common/oauth2/token", actual.toString());
    }

    @Test
    public void createTokenRequest_typical() throws Exception
    {
        final String resource = "d8011f5a-ef4e-0cab-b70d-0aecdd55d3ca";
        final String clientId = "c0b34bbc-f14f-0358-bb27-6bfb8cfc8903";
        // authorization codes can be pretty long
        final String authorizationCode =
            "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
            "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
            "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
            "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
            "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
            "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
            "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
            "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA";
        final URI redirectUri = URI.create("https://example.com");

        final StringContent actual = AzureAuthority.createTokenRequest(resource, clientId, authorizationCode, redirectUri, null);

        Assert.assertEquals(
            "resource=e3542c4c-ab4a-6ada-c50e-5adabd48b0fa" +
            "&client_id=e0a24fba-d15d-0843-ee70-5cfb8daf5094" +
            "&grant_type=authorization_code" +
            "&code=" + authorizationCode +
            "&redirect_uri=https%3A%2F%2Fexample.com", actual.getContent());
    }

    @Test
    public void createTokenRequest_withCorrelationId() throws Exception
    {
        final String resource = "e7507a0e-cf7b-3faf-d93c-5eaeca69f0eb";
        final String clientId = "c3f43ecf-f37b-7866-da47-9fbb1edf9962";
        final UUID correlationId = UUID.fromString("949e5ad0-c16c-8707-6748-1a34343752c0");
        // authorization codes can be pretty long
        final String authorizationCode =
            "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
                "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
                "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
                "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
                "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
                "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
                "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA" +
                "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA";
        final URI redirectUri = URI.create("https://example.com");

        final StringContent actual = AzureAuthority.createTokenRequest(resource, clientId, authorizationCode, redirectUri, correlationId);

        Assert.assertEquals(
            "resource=e7019f1f-ea1b-4efa-f64e-8bdebf11a7fe" +
                "&client_id=f8f65faf-f61a-2044-cf12-2faa7cbc1265" +
                "&grant_type=authorization_code" +
                "&code=" + authorizationCode +
                "&redirect_uri=https%3A%2F%2Fexample.com" +
                "&client-request-id=043f3fd6-c96a-6304-7982-9e78033730f4" +
                "&return-client-request-id=true", actual.getContent());
    }
}
