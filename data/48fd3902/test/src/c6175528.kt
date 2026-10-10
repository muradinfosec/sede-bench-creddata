package org.cloudfoundry.credhub.endToEnd.v1.credentials

import org.apache.commons.lang3.StringUtils
import org.cloudfoundry.credhub.CredhubTestApp
import org.cloudfoundry.credhub.helpers.RequestHelper.generatePassword
import org.cloudfoundry.credhub.helpers.RequestHelper.setPassword
import org.cloudfoundry.credhub.utils.AuthConstants
import org.cloudfoundry.credhub.utils.AuthConstants.Companion.ALL_PERMISSIONS_TOKEN
import org.cloudfoundry.credhub.utils.DatabaseProfileResolver
import org.hamcrest.CoreMatchers.containsString
import org.hamcrest.CoreMatchers.equalTo
import org.hamcrest.CoreMatchers.`is`
import org.hamcrest.MatcherAssert.assertThat
import org.json.JSONObject
import org.junit.Before
import org.junit.Test
import org.junit.runner.RunWith
import org.skyscreamer.jsonassert.JSONAssert
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.boot.test.context.SpringBootTest
import org.springframework.http.MediaType.APPLICATION_JSON
import org.springframework.security.test.web.servlet.setup.SecurityMockMvcConfigurers.springSecurity
import org.springframework.test.context.ActiveProfiles
import org.springframework.test.context.junit4.SpringRunner
import org.springframework.test.web.servlet.MockMvc
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders.get
import org.springframework.test.web.servlet.request.MockMvcRequestBuilders.put
import org.springframework.test.web.servlet.result.MockMvcResultHandlers.print
import org.springframework.test.web.servlet.result.MockMvcResultMatchers.status
import org.springframework.test.web.servlet.setup.DefaultMockMvcBuilder
import org.springframework.test.web.servlet.setup.MockMvcBuilders
import org.springframework.transaction.annotation.Transactional
import org.springframework.web.context.WebApplicationContext

@RunWith(SpringRunner::class)
@ActiveProfiles(value = ["unit-test"], resolver = DatabaseProfileResolver::class)
@SpringBootTest(classes = [CredhubTestApp::class])
@Transactional
class CredentialSetEndToEndTest {

    @Autowired
    private lateinit var webApplicationContext: WebApplicationContext

    private lateinit var mockMvc: MockMvc

    companion object {
        private const val CREDENTIAL_NAME = "/set_credential"
        private val CREDENTIAL_NAME_1024_CHARACTERS = StringUtils.rightPad(
            "/",
            1024,
            'a'
        )
    }

    @Before
    fun setUp() {
        mockMvc = MockMvcBuilders
            .webAppContextSetup(webApplicationContext)
            .apply<DefaultMockMvcBuilder>(springSecurity())
            .build()
    }

    @Test
    fun `rsa credential can be set without private key`() {
        val setRsaRequest = put("/api/v1/data")
            .header("Authorization", "Bearer ${AuthConstants.ALL_PERMISSIONS_TOKEN}")
            .accept(APPLICATION_JSON)
            .contentType(APPLICATION_JSON)
            // language=JSON
            .content(
                """
                {
                    "name": "$CREDENTIAL_NAME",
                    "type": "rsa",
                    "value": {
                      "public_key": "a_certain_public_key",
                      "private_key": ""
                    }
                }
                """.trimIndent()
            )

        this.mockMvc
            .perform(setRsaRequest)
            .andDo(print())
            .andExpect(status().isOk)
            .andReturn().response
            .contentAsString
    }

    @Test
    fun `user credential returns null username when set with blank string as username`() {
        val setUserRequest = put("/api/v1/data")
            .header("Authorization", "Bearer ${AuthConstants.ALL_PERMISSIONS_TOKEN}")
            .accept(APPLICATION_JSON)
            .contentType(APPLICATION_JSON)
            // language=JSON
            .content(
                """
              {
                "name": "$CREDENTIAL_NAME",
                "type": "user",
                "value": {
                  "username": "",
                  "password": "some_silly_password"
                }
              }
                """.trimIndent()
            )

        val response = this.mockMvc
            .perform(setUserRequest)
            .andDo(print())
            .andExpect(status().isOk)
            .andReturn().response
            .contentAsString

        assertThat(response, containsString("\"username\":null"))
    }

    @Test
    fun `credential should always be overwritten in set request`() {
        setPassword(
            mockMvc,
            CREDENTIAL_NAME,
            "original-password",
            ALL_PERMISSIONS_TOKEN
        )

        val secondResponse = setPassword(
            mockMvc,
            CREDENTIAL_NAME,
            "new-password",
            ALL_PERMISSIONS_TOKEN
        )
        val updatedPassword = JSONObject(secondResponse).getString("value")

        assertThat(updatedPassword, equalTo("new-password"))
    }

    @Test
    fun `credential names can have a length of 1024 characters`() {
        assertThat(CREDENTIAL_NAME_1024_CHARACTERS.length, `is`(equalTo(1024)))

        val setResponse = setPassword(
            mockMvc,
            CREDENTIAL_NAME_1024_CHARACTERS,
            "foobar",
            ALL_PERMISSIONS_TOKEN
        )
        val setPassword = JSONObject(setResponse).getString("value")

        assertThat(setPassword, equalTo("foobar"))

        val getResponse = generatePassword(
            mockMvc,
            CREDENTIAL_NAME_1024_CHARACTERS,
            true,
            14,
            ALL_PERMISSIONS_TOKEN
        )
        val getPassword = JSONObject(getResponse).getString("value")
        assertThat(getPassword.length, equalTo(14))
    }

    @Test
    fun `credential names that exceed the maximum length should result in 400`() {
        val name1025 = CREDENTIAL_NAME_1024_CHARACTERS + "a"
        assertThat(name1025.length, `is`(equalTo(1025)))

        setPassword(
            mockMvc,
            name1025,
            "foobar",
            ALL_PERMISSIONS_TOKEN
        )
        generatePassword(
            mockMvc,
            name1025,
            false,
            10,
            ALL_PERMISSIONS_TOKEN
        )
    }

    @Test
    fun `malformed private key should result in 400`() {

        val certificate =
            """
            -----BEGIN CERTIFICATE----- fake
            MIIDPjCCAiagAwIBAgIUIgg7xZVYF3qFsUVAhAFldTvCDJ4wDQYJKoZIhvcNAQEL
            BQAwFjEUMBIGA1UEAxMLZXhhbXBsZS5jb20wHhcNMTgwNjE1MTUwMDU3WhcNMTkw
            NjE1MTUwMDU3WjAWMRQwEgYDVQQDEwtleGFtcGxlLmNvbTCCASIwDQYJKoZIhvcN
            AQEBBQADggEPADCCAQoCggEBAM6z9Y/odS4pldElmK3syIbxhy5gPR5yvRIpEE89
            yXEkAJjyW8+zIjZM6/bIEIkAOAObXWLbcqI/Wv+FSxsUq55IYIZlaBpoHjl5rsvv
            inBbsKBChAPLuLBNNR8NJ/8gkZkeBsobBkkhTjZl1f6+GGAnLazqLxl8tyxwhNBe
            dlONwozUuJ1Vlve65L+cuapnKlmYz+ZYd4f75mJcs2OPUmXhbhTK+RI0gtZC84Qg
            0+pPheXjde/E8f0HrW2cO0wewxdAPnzD5MvQCZdc1ndpp2df4DZgLtxXozpLCSHF
            LxhnOkEGjtmxHG8YelrXZ0QbsZOumuvbWmK71PTalOKSe4cCAwEAAaOBgzCBgDAd
            BgNVHQ4EFgQUJbJRTUNhGiVXo/ELta+dlRCALwswUQYDVR0jBEowSIAUJbJRTUNh
            GiVXo/ELta+dlRCALwuhGqQYMBYxFDASBgNVBAMTC2V4YW1wbGUuY29tghQiCDvF
            lVgXeoWxRUCEAWV1O8IMnjAMBgNVHRMBAf8EAjAAMA0GCSqGSIb3DQEBCwUAA4IB
            AQAUM7zOD09vxMMGELbm3m+DgJOIhWm6zkibpzn1P1e7Pi7BOQ+2GvXBmn030yQU
            O5rKLNv49up9XGViKsPfVjbmxWp9WbElNPW+dJyO3zLMMkFtm/1T39Y+/A1LH3ww
            HSOnT3s54pSI66L9Mpiq+V2VmiOKEvoxy2mGQteMkXWSX31p0PlKMToV34TDIk9M
            9XyxHVWTf5NLe/gUEIoZatdvMmANKmKBiUWI5Aqnh93a2TXDu2Q8WXc0U0W8hsbD
            Wv7ec0Gguo4GtOomkmFIgXBLZd0ZqWywEjSGRy4us/71gBioTgCBMw8g75SzxX5u
            hQHS5//LiA50aEI4X0k5TDQp
            -----END CERTIFICATE-----
            """.trimIndent().replace("\n", "\\n")

        val invalidPrivateKey =
            """
            -----BEGIN RSA PRIVATE KEY----- fake
            MIIEpQIBAAKCAQEAwqIrV8HpCuPyuJ6VvyG7gVhYJGAOX4zhclxkQgVe7haQ2OVX
            175sqdZwmX+uTw+HBooRXWPBTBvVJCHl4AW8Wn2tdXhXQ03CIRYGic57WNhfOOql
            TOUPQ/1lokNemf7HlFK02oEAwUcKLdP8pqMI/24Ehb8duJcLNq6ft9h0cqjUGkKO
            bTOfjYMCzVPFlZTxY1OsoPlAe18ooNqF29lHYgKFINXOws6FoaK0A0AHa4m0OAM5
            PxV5CuucgyWsZC9txYVn5zCOPZevcYZWOGYJMiGtJ4DsRif1BsIQST3E965ZnL3A
            s2xUlssbcR9EeCIwXHUtLkjOfiYTn6Woffc8dPREQZqiSNKxAazIOmVsRINxegmo
            2FZsRtnm4dBYBzWjp9Yjj7IDTrEXHw5i7TneW0zpIuedHGvWL0NWPZIOJZuAC3Q/
            XiM/gNjthdAGjo/wLapN1coDRPaHGNHnHznOM7F0tYbKlpP3pytOiq/c3Xv81b6T
            BvtWNuuPw44TavbodI1dOukeRe/t3/hnZkRmS1mcBTCwQK7A5ajyv40JNcif0rxP
            00oRTVLaOZgzpNAtYI9fsSWiocPr1oIFLkVDXAa7WkJ7VDEiSMyoMX/CSEGIx4Xq
            GFjc8nHe7fZNk/FQZMd6+ZlhHhSqgepgBBJrpOHIfpAyUMlWkipX7Xz1icCa70fr
            cj3urQBwbtDZ/WLg97yfxqjgj11sSb/lsSM4FxfP0v4VAqdwmghnfIlfNMsgm1RD
            U8lVPorzC715wtkmTCXIg1PYnuxEl4peSQj+kk+APBU3lVz8XGJLowYUHt8lcA+U
            3B5/QQ/Y4GSpZyI+J8QhYwvFJCSIP0FhxZhgZJX7KKoI4jhVGmb7uJXRkOKvX8iP
            ZZRwisJ5hDNUhSmshdutDfIs1mCQULkCYoQsyY7cDgg51UBlfqd2JBm8leB/QLfg
            8AmReDS1iVl//4FDYHVXU/Si+1TBnZdgTz/7xu551VamugDI7JAR1GipXGihwadv
            E4BfyWiaxP5Xum/W479wk30uQf8yDczHu7yIthpjvLUC6X4ZuaqN77gB+mBmzayY
            PvSJv2QT/9uf7E1VcMdYqy5HD2p7tfQ+kNjvMRF3Szcn1hE8OJgBrqdlEp1zaJaj
            aoAavy9toKhhH8rn6csA6sUzHCwuhpGkkRhx0uEQFjUkFS0e2LYS35PAQIjX/lRW
            sX7sXLt6a/rgpVEd+2w7xvKorHPHzVhQqTuYxGO9PZvTY7kPV829B8qKkq3GdJl3
            lXzVUm5RVc82aikMu2mMHpRRWqlxQ39IRgH9WatswBrsouUf8AqYEdBIcVqe+GoS
            QEY/oUJxY68ZTtmm+J9bIPGC7PS2ERwei9XTUuLCqDFnbyGScV6eASQUezfWoRsk
            ++2WvvaVjanjdhK7cIOzh9nrMfGNl9ivS1qZa4gHlr+D1xyUkOw9vv/xrvRPiUr5
            NksZzzcoi1JdTKL7Sw3j3qNK4CiVpR9mCcmPamE+Hpug80RBEBcZCg3dTclcZgj4
            6dJyv8AeXYzvddxM6KEK33Jrp9BuDByz5bMZHY4+RJytokTK0O2/zEJE=
            -----END RSA PRIVATE KEY-----
            """.trimIndent().replace("\n", "\\n")

        val request = put("/api/v1/data")
            .header("Authorization", "Bearer ${AuthConstants.ALL_PERMISSIONS_TOKEN}")
            .accept(APPLICATION_JSON)
            .contentType(APPLICATION_JSON)
            // language=json
            .content(
                """
                  {
                    "name": "some-cert",
                    "type": "certificate",
                    "value": {
                      "ca": "$certificate",
                      "certificate": "$certificate",
                      "private_key": "$invalidPrivateKey"
                    }
                  }
                """.trimIndent()
            )

        val responseBody = this.mockMvc
            .perform(request)
            .andDo(print())
            .andExpect(status().isBadRequest)
            .andReturn().response
            .contentAsString

        JSONAssert.assertEquals(
            responseBody,
            """
                {
                  "error": "Private key is malformed. Keys must be PEM-encoded PKCS#1 or unencrypted PKCS#8 keys."
                }
            """.trimIndent(),
            true
        )
    }

    @Test
    fun `credential contains metadata`() {
        val setUserRequest = put("/api/v1/data")
            .header("Authorization", "Bearer ${AuthConstants.ALL_PERMISSIONS_TOKEN}")
            .accept(APPLICATION_JSON)
            .contentType(APPLICATION_JSON)
            // language=JSON
            .content(
                """
              {
                "name": "$CREDENTIAL_NAME",
                "type": "user",
                "metadata": { "description": "example metadata" },
                "value": {
                  "username": "some_silly_username",
                  "password": "some_silly_password"
                }
              }
                """.trimIndent()
            )

        val setResponse = this.mockMvc
            .perform(setUserRequest)
            .andDo(print())
            .andExpect(status().isOk)
            .andReturn().response
            .contentAsString

        assertThat(setResponse, containsString("\"metadata\":{\"description\":\"example metadata\"}"))

        val getUserRequest = get("/api/v1/data?name=$CREDENTIAL_NAME")
            .header("Authorization", "Bearer ${AuthConstants.ALL_PERMISSIONS_TOKEN}")
            .accept(APPLICATION_JSON)

        val getResponse = this.mockMvc
            .perform(getUserRequest)
            .andDo(print())
            .andExpect(status().isOk)
            .andReturn().response
            .contentAsString

        assertThat(getResponse, containsString("\"metadata\":{\"description\":\"example metadata\"}"))
    }
}
