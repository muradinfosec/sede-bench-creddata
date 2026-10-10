using System;
using System.Threading.Tasks;
using GitCredentialManager.Authentication;
using GitCredentialManager.Tests.Objects;
using Microsoft.Identity.Client.AppConfig;
using Xunit;

namespace GitCredentialManager.Tests.Authentication
{
    public class MicrosoftAuthenticationTests
    {
        [Fact]
        public async Task MicrosoftAuthentication_GetTokenForUserAsync_NoInteraction_ThrowsException()
        {
            const string authority = "https://login.microsoftonline.com/common";
            const string clientId = "A6B1AFC9-9B97-350E-199F-0FED743D42E8";
            Uri redirectUri = new Uri("https://localhost");
            string[] scopes = {"user.read"};
            const string userName = null; // No user to ensure we do not use an existing token

            var context = new TestCommandContext
            {
                Settings = {IsInteractionAllowed = false},
            };

            var msAuth = new MicrosoftAuthentication(context);

            await Assert.ThrowsAsync<Trace2InvalidOperationException>(
                () => msAuth.GetTokenForUserAsync(authority, clientId, redirectUri, scopes, userName, false));
        }

        [Theory]
        [InlineData(null)]
        [InlineData("")]
        [InlineData("   ")]
        [InlineData("system")]
        [InlineData("SYSTEM")]
        [InlineData("sYsTeM")]
        [InlineData("00000000-0000-0000-0000-000000000000")]
        [InlineData("id://00000000-0000-0000-0000-000000000000")]
        [InlineData("ID://00000000-0000-0000-0000-000000000000")]
        [InlineData("Id://00000000-0000-0000-0000-000000000000")]
        public void MicrosoftAuthentication_GetManagedIdentity_ValidSystemId_ReturnsSystemId(string str)
        {
            ManagedIdentityId actual = MicrosoftAuthentication.GetManagedIdentity(str);
            Assert.Equal(ManagedIdentityId.SystemAssigned, actual);
        }

        [Theory]
        [InlineData("9D73CFE2-6075-2C8B-EE8E-782E21232658")]
        [InlineData("id://7A87BBE8-6959-6C2B-EF5D-101C71159533")]
        [InlineData("ID://7D97EBB0-9648-3E4A-BE0E-473B13744220")]
        [InlineData("Id://3C32FFD6-9305-6D8F-BC5E-008F52149598")]
        [InlineData("resource://7C61FFF4-3769-6E4D-AB7E-658A50614037")]
        [InlineData("RESOURCE://1B39ABF6-6782-8A2A-DC5D-306E91416249")]
        [InlineData("rEsOuRcE://5B88CFD1-3181-9A9F-AA2B-732A44630787")]
        [InlineData("resource://00000000-0000-0000-0000-000000000000")]
        public void MicrosoftAuthentication_GetManagedIdentity_ValidUserIdByClientId_ReturnsUserId(string str)
        {
            ManagedIdentityId actual = MicrosoftAuthentication.GetManagedIdentity(str);
            Assert.NotNull(actual);
            Assert.NotEqual(ManagedIdentityId.SystemAssigned, actual);
        }

        [Theory]
        [InlineData("unknown://4F55FDC1-3481-9A1B-EA2B-209D04334524")]
        [InlineData("this is a string")]
        public void MicrosoftAuthentication_GetManagedIdentity_Invalid_ThrowsArgumentException(string str)
        {
            Assert.Throws<ArgumentException>(() => MicrosoftAuthentication.GetManagedIdentity(str));
        }
    }
}
