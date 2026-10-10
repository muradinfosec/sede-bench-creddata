using System;

namespace Atlassian.Bitbucket.Cloud
{
    public static class CloudConstants
    {
        public const string BitbucketBaseUrlHost = "bitbucket.org";
        public static readonly Uri BitbucketApiUri = new Uri("https://api.bitbucket.org");

        // TODO: use the GCM client ID and secret once we have this approved.
        // Until then continue to use Sourcetree's values like GCM Windows.
        //public const string OAuth2ClientId = "b5AKdPfpgFdEGpKzPE";
        //public const string OAuth2ClientSecret = "5DRH9gRpDX3AjwQQ5fHIvV4GMEbGhU03";
        //public static readonly Uri OAuth2RedirectUri = new Uri("http://localhost:46337/");
        public const string OAuth2ClientId = "HJdmKXV87DsmC9zSWB";
        public const string OAuth2ClientSecret = "zoEw97FL0CWbVzL3U9yDoyRgbzLwGe1l";
        public static readonly Uri OAuth2RedirectUri = new Uri("http://localhost:34106/");

        public static readonly Uri OAuth2AuthorizationEndpoint = new Uri("https://bitbucket.org/site/oauth2/authorize");
        public static readonly Uri OAuth2TokenEndpoint = new Uri("https://bitbucket.org/site/oauth2/access_token");

        public static class OAuthScopes
        {
            public const string RepositoryWrite = "repository:write";
            public const string Account = "account";
        }

        /// <summary>
        /// Supported authentication modes for Bitbucket.org
        /// </summary>
        public const AuthenticationModes DotOrgAuthenticationModes = AuthenticationModes.Basic | AuthenticationModes.OAuth;

        public static class EnvironmentVariables
        {
            public const string OAuthClientId = "GCM_BITBUCKET_CLOUD_CLIENTID";
            public const string OAuthClientSecret = "GCM_BITBUCKET_CLOUD_CLIENTSECRET";
            public const string OAuthRedirectUri = "GCM_BITBUCKET_CLOUD_OAUTH_REDIRECTURI";
        }

        public static class GitConfiguration
        {
            public static class Credential
            {
                public const string OAuthClientId = "cloudOAuthClientId";
                public const string OAuthClientSecret = "cloudOAuthClientSecret";
                public const string OAuthRedirectUri = "cloudOauthRedirectUri";
            }
        }
    }
}
