/* eslint-disable quotes, max-len */

module.exports = {
    v1: {
        id: "76cb54fb-25af-ea0e-9ef3-495a38e88e79",
        name: "Postman Echo (shamasified)",
        description: "Postman Echo is service you can use to test your REST clients and make sample API calls. It provides endpoints for `GET`, `POST`, `PUT`, various auth mechanisms and other utility endpoints.",
        order: [],
        folders: [
            {
                owner: "33232",
                lastUpdatedBy: "33232",
                lastRevision: 75111444,
                id: "ee146857-a752-4c36-e0ff-d5100ea9ec81",
                name: "Auth: Digest",
                description: "Digest authentication protects an endpoint with a username and password without actually transmitting the password over network.\nOne has to apply a hash function (like MD5, etc) to the username and password before sending them over the network.\n\n> Username: `postman`\n>\n> Password: `password`\n\nUnlike Basic-Auth, authentication happens using two consecutive requests where the first request returns `401 Unauthorised` along with `WWW-Authenticate` header containing information that needs to be used to authenticate subsequent calls.\n\nTo know more about digest authentication, refer to the [Digest Access Authentication](https://en.wikipedia.org/wiki/Digest_access_authentication) wikipedia article.\nThe article on [authentication helpers](https://www.getpostman.com/docs/helpers#digest-auth) elaborates how to use the same within the Postman app.",
                order: [
                    "f9d9531a-ee06-1b85-0813-4169236a5471",
                    "edec454a-7661-b06d-846a-b51e54f59efc"
                ]
            },
            {
                owner: "33232",
                lastUpdatedBy: "61671",
                lastRevision: 90671042,
                id: "a7130439-fec8-5ad9-b2ea-9cef7622d4e0",
                name: "Auth: Others",
                order: [
                    "fe62031a-ca79-08e6-8a17-79c51ad5b52b",
                    "fbe68c98-1636-04e0-16d5-42f53f86aded"
                ]
            },
            {
                owner: "33232",
                lastUpdatedBy: "33232",
                lastRevision: 75569578,
                id: "d0a0c859-3052-9e61-e108-85778cafe28a",
                name: "Cookies",
                description: "The cookie related endpoints allow one to get, set and delete simple cookies.\n\nCookies are small snippets of information that is stored in the browser and sent back to the server with every subsequent requests in order to store useful information between requests.\nIf you want to know more about cookies, read the [HTTP Cookie](https://en.wikipedia.org/wiki/HTTP_cookie) article on wikipedia.",
                order: [
                    "14a42438-8935-ac1a-ff0a-ec47b18f7eed",
                    "7262bc38-c650-245f-719f-0a3d508324c4",
                    "7ed9b625-7922-bc9a-1a31-0e6deefc2217"
                ]
            },
            {
                owner: "33232",
                lastUpdatedBy: "33232",
                lastRevision: 75211067,
                id: "507e1e06-33f2-57e9-1935-745e04eafa24",
                name: "Headers",
                description: "The following set of endpoints allow one to see the headers being sent as part of a request and to get a custom set of headers as part of response.\n\nHTTP header fields provide required information about the request or response, or about the object sent in the message body. Both request headers and response headers can be controlled using these endpoints.",
                order: [
                    "016f4d7e-ee5f-9502-4faf-f69c80003d84",
                    "6cf9d886-9871-e9da-fda9-2af66a77f5c8"
                ]
            },
            {
                owner: "33232",
                lastUpdatedBy: "61671",
                lastRevision: 75574910,
                id: "c1334864-05ae-43f0-9493-1fbbaf6da7fc",
                name: "Request Methods",
                description: "HTTP has multiple request \"verbs\", such as `GET`, `PUT`, `POST`, `DELETE`,\n`PATCH`, `HEAD`, etc. \n\nAn HTTP Method (verb) defines how a request should be interpreted by a server. \nThe endpoints in this section demonstrate various HTTP Verbs. Postman supports \nall the HTTP Verbs, including some rarely used ones, such as `PROPFIND`, `UNLINK`, \netc.\n\nFor details about HTTP Verbs, refer to [RFC 2616](http://www.w3.org/Protocols/rfc2616/rfc2616-sec9.html#sec9)\n",
                order: [
                    "102f23f7-2dc5-5c49-2bc8-70cabab977ff",
                    "0756d7ae-f0af-9ce4-d999-f7c42e23938e",
                    "56652707-f59d-ccdc-a721-078b341b86d9",
                    "65dc126c-f80d-eaae-2bdc-cf14fade6b50",
                    "f0def219-5fb7-75e5-5b52-a5f175c72961"
                ]
            },
            {
                owner: "33232",
                lastUpdatedBy: "33232",
                lastRevision: 75050206,
                id: "7379492f-1051-a50f-05ed-13852711580b",
                name: "Utilities",
                order: [
                    "7212bf63-76c8-6e1d-8324-331e42a4ab14",
                    "7ed4a4a9-55e2-de84-bbb2-ca9ec6bd6e78",
                    "285f4c1a-f872-48c1-c53d-0ca65df01b18",
                    "a407f293-7737-1281-ebfb-05b3dda9416b",
                    "ac6a1772-13d4-9f62-0251-bc49032fc2bc",
                    "344a07b8-15da-2a2d-fb43-a5972a739eb3"
                ]
            },
            {
                owner: "33232",
                lastUpdatedBy: "33232",
                lastRevision: 75566609,
                id: "85c2634f-3e1b-b25f-1270-a6035945e2d6",
                name: "[draft] Auth: OAuth2.0",
                description: "OAuth 2.0 allows a client to access a resource owners(user) protected resources without direct access to the users credentials.\nThe client obtains an access token which can be used to access the resource owners resources.\nOAuth 2.0 implements 4 different types of grant flows. Postman Echo has implemented the [Authorization Code](https://tools.ietf.org/html/rfc6749#section-1.3.1) grant flow.\n\nThe OAuth flow typically consists of three steps. \nIn the first step the client requests an authentication code from the server implementing the OAuth flow.\nThe authentication code along with other client credentials is then used to get the access token.\nThe access token can be used by the client to access users resources.\n",
                order: [
                    "d8eebc20-7031-e76b-41fd-4168d9352c1d",
                    "2412f47a-a661-68e8-dc9a-23400426807a",
                    "d12fa807-a18c-fdb6-ebfd-12cf05b22945"
                ]
            }
        ],
        timestamp: 0,
        owner: "33232",
        remoteLink: "",
        public: false,
        requests: [
            {
                folder: "86f8023a-5b9b-a51f-0036-b4695586a2a5",
                id: "2527b77f-b148-32d1-bd4b-51194446348c",
                name: "OAuth2.0 Get Access Token",
                dataMode: "params",
                data: [
                    {
                        key: "code",
                        value: "xWnkliVQJURqB2x1",
                        type: "text"
                    },
                    {
                        key: "grant_type",
                        value: "authorization_code",
                        type: "text"
                    },
                    {
                        key: "redirect_uri",
                        value: "https://www.getpostman.com/oauth2/callback",
                        type: "text"
                    },
                    {
                        key: "client_id",
                        value: "abc123",
                        type: "text"
                    },
                    {
                        key: "client_secret",
                        value: "ssh-secret",
                        type: "text"
                    }
                ],
                rawModeData: null,
                descriptionFormat: null,
                description: "This endpoint is used to get the `access_token`\n\nIt requires the following secret client credentials to be sent as part of the form body along with the `authentication code` obtained as part of the `redirect_uri` from the previous request.\n\n> code: xWnkliVQJURqB2x1\n>\n> grant_type: authorization_code\n>\n> redirect_uri: https://www.getpostman.com/oauth2/callback\n>\n> client_id: abc123\n>\n> client_secret: ssh-secret\n\nIf the correct credentials are not passed, the server returns with a `401 Unauthorized` response.",
                headers: "",
                method: "POST",
                pathVariables: {},
                url: "https://echo.getpostman.com/oauth2/token",
                preRequestScript: "",
                tests: "tests[\"response code is 200\"] = responseCode.code === 200;\nvar body = JSON.parse(responseBody);\ntests[\"body has access token\"] = \"access_token\" in body;\ntests[\"body has bearer type\"] = \"token_type\" in body;",
                currentHelper: "normal",
                helperAttributes: {},
                collectionId: "89af66ae-81db-fe4c-1ac8-805f37d07a28",
                responses: [
                    {
                        owner: "33232",
                        lastUpdatedBy: "33232",
                        lastRevision: 75106841,
                        request: {
                            name: "Access Token",
                            description: "Tries to get the access token by passing the clientId , clientSecret,\nauthentication code, redirect URI and grant type.\n",
                            url: "https://echo.getpostman.com/oauth2/token",
                            pathVariables: {},
                            data: [
                                {
                                    key: "code",
                                    value: "xWnkliVQJURqB2x1",
                                    type: "text"
                                },
                                {
                                    key: "grant_type",
                                    value: "authorization_code",
                                    type: "text"
                                },
                                {
                                    key: "redirect_uri",
                                    value: "https://www.getpostman.com/oauth2/callback",
                                    type: "text"
                                },
                                {
                                    key: "client_id",
                                    value: "abc123",
                                    type: "text"
                                },
                                {
                                    key: "client_secret",
                                    value: "ssh-secret",
                                    type: "text"
                                }
                            ],
                            headers: "",
                            dataMode: "params",
                            method: "POST",
                            tests: "tests[\"response code is 200\"] = responseCode.code === 200;\nvar body = JSON.parse(responseBody);\ntests[\"body has access token\"] = \"access_token\" in body;\ntests[\"body has bearer type\"] = \"token_type\" in body;",
                            isFromCollection: true,
                            write: true,
                            version: 2
                        },
                        id: "9860989b-734e-deb5-fa2a-bab263fec6d6",
                        name: "401",
                        status: "",
                        responseCode: {
                            code: 401,
                            name: "Unauthorized"
                        },
                        time: "287",
                        headers: [
                            {
                                name: "Access-Control-Allow-Credentials",
                                key: "Access-Control-Allow-Credentials",
                                value: ""
                            },
                            {
                                name: "Access-Control-Allow-Headers",
                                key: "Access-Control-Allow-Headers",
                                value: ""
                            },
                            {
                                name: "Access-Control-Allow-Methods",
                                key: "Access-Control-Allow-Methods",
                                value: ""
                            },
                            {
                                name: "Access-Control-Allow-Origin",
                                key: "Access-Control-Allow-Origin",
                                value: ""
                            },
                            {
                                name: "Connection",
                                key: "Connection",
                                value: "keep-alive"
                            },
                            {
                                name: "Content-Length",
                                key: "Content-Length",
                                value: "2"
                            },
                            {
                                name: "Content-Type",
                                key: "Content-Type",
                                value: "application/json; charset=utf-8"
                            },
                            {
                                name: "Date",
                                key: "Date",
                                value: "Sat, 31 Oct 2015 06:43:44 GMT"
                            },
                            {
                                name: "Server",
                                key: "Server",
                                value: "nginx/1.6.2"
                            },
                            {
                                name: "Vary",
                                key: "Vary",
                                value: "X-HTTP-Method-Override, Accept-Encoding"
                            },
                            {
                                name: "X-Powered-By",
                                key: "X-Powered-By",
                                value: "Sails <sailsjs.org>"
                            }
                        ],
                        cookies: [],
                        mime: "",
                        text: "{}",
                        language: "javascript",
                        rawDataType: "text",
                        state: {
                            size: "normal"
                        },
                        previewType: "html",
                        searchResultScrolledTo: "-1",
                        version: null,
                        requestObject: "{\"name\":\"Access Token\",\"description\":\"Tries to get the access token by passing the clientId , clientSecret,\\nauthentication code, redirect URI and grant type.\\n\",\"url\":\"https://echo.getpostman.com/oauth2/token\",\"pathVariables\":{},\"data\":[{\"key\":\"code\",\"value\":\"xWnkliVQJURqB2x1\",\"type\":\"text\",\"enabled\":true},{\"key\":\"grant_type\",\"value\":\"authorization_code\",\"type\":\"text\",\"enabled\":true},{\"key\":\"redirect_uri\",\"value\":\"https://www.getpostman.com/oauth2/callback\",\"type\":\"text\",\"enabled\":true},{\"key\":\"client_id\",\"value\":\"abc123\",\"type\":\"text\",\"enabled\":true},{\"key\":\"client_secret\",\"value\":\"ssh-secret\",\"type\":\"text\",\"enabled\":true}],\"headers\":\"\",\"dataMode\":\"params\",\"method\":\"POST\",\"tests\":\"tests[\\\"response code is 200\\\"] = responseCode.code === 200;\\nvar body = JSON.parse(responseBody);\\ntests[\\\"body has access token\\\"] = \\\"access_token\\\" in body;\\ntests[\\\"body has bearer type\\\"] = \\\"token_type\\\" in body;\",\"isFromCollection\":true,\"write\":true,\"version\":2}",
                        createdAt: "2015-11-02T13:11:08.000Z",
                        updatedAt: "2015-11-02T18:05:45.000Z",
                        write: true
                    },
                    {
                        owner: "33232",
                        lastUpdatedBy: "33232",
                        lastRevision: 75106842,
                        request: {
                            name: "Access Token",
                            description: "Tries to get the access token by passing the clientId , clientSecret,\nauthentication code, redirect URI and grant type.\n",
                            url: "https://echo.getpostman.com/oauth2/token",
                            pathVariables: {},
                            data: [
                                {
                                    key: "code",
                                    value: "xWnkliVQJURqB2x1",
                                    type: "text"
                                },
                                {
                                    key: "grant_type",
                                    value: "authorization_code",
                                    type: "text"
                                },
                                {
                                    key: "redirect_uri",
                                    value: "https://www.getpostman.com/oauth2/callback",
                                    type: "text"
                                },
                                {
                                    key: "client_id",
                                    value: "abc123",
                                    type: "text"
                                },
                                {
                                    key: "client_secret",
                                    value: "ssh-secret",
                                    type: "text"
                                }
                            ],
                            headers: "",
                            dataMode: "params",
                            method: "POST",
                            tests: "tests[\"response code is 200\"] = responseCode.code === 200;\nvar body = JSON.parse(responseBody);\ntests[\"body has access token\"] = \"access_token\" in body;\ntests[\"body has bearer type\"] = \"token_type\" in body;",
                            isFromCollection: true,
                            write: true,
                            version: 2
                        },
                        id: "ea4b2573-240d-502e-d9de-a64be9a36451",
                        name: "200",
                        status: "",
                        responseCode: {
                            code: 200,
                            name: "OK"
                        },
                        time: "303",
                        headers: [
                            {
                                name: "Access-Control-Allow-Credentials",
                                key: "Access-Control-Allow-Credentials",
                                value: ""
                            },
                            {
                                name: "Access-Control-Allow-Headers",
                                key: "Access-Control-Allow-Headers",
                                value: ""
                            },
                            {
                                name: "Access-Control-Allow-Methods",
                                key: "Access-Control-Allow-Methods",
                                value: ""
                            },
                            {
                                name: "Access-Control-Allow-Origin",
                                key: "Access-Control-Allow-Origin",
                                value: ""
                            },
                            {
                                name: "Connection",
                                key: "Connection",
                                value: "keep-alive"
                            },
                            {
                                name: "Content-Encoding",
                                key: "Content-Encoding",
                                value: "gzip"
                            },
                            {
                                name: "Content-Length",
                                key: "Content-Length",
                                value: "153"
                            },
                            {
                                name: "Content-Type",
                                key: "Content-Type",
                                value: "application/json; charset=utf-8"
                            },
                            {
                                name: "Date",
                                key: "Date",
                                value: "Sat, 31 Oct 2015 06:43:34 GMT"
                            },
                            {
                                name: "Server",
                                key: "Server",
                                value: "nginx/1.6.2"
                            },
                            {
                                name: "Vary",
                                key: "Vary",
                                value: "X-HTTP-Method-Override, Accept-Encoding"
                            },
                            {
                                name: "X-Powered-By",
                                key: "X-Powered-By",
                                value: "Sails <sailsjs.org>"
                            }
                        ],
                        cookies: [],
                        mime: "",
                        text: "{\"access_token\":\"dr7koAvbzbbrKDp1oKfSxs4QbJLtiRweOcJAPabNW9XGvtZE04DoLD45CA0iv8SydDOJKKmkBuMh7MwQvVRMzeomG8MFRCUk\",\"token_type\":\"Bearer\"}",
                        language: "javascript",
                        rawDataType: "text",
                        state: {
                            size: "normal"
                        },
                        previewType: "html",
                        searchResultScrolledTo: "-1",
                        version: null,
                        requestObject: "{\"name\":\"Access Token\",\"description\":\"Tries to get the access token by passing the clientId , clientSecret,\\nauthentication code, redirect URI and grant type.\\n\",\"url\":\"https://echo.getpostman.com/oauth2/token\",\"pathVariables\":{},\"data\":[{\"key\":\"code\",\"value\":\"xWnkliVQJURqB2x1\",\"type\":\"text\",\"enabled\":true},{\"key\":\"grant_type\",\"value\":\"authorization_code\",\"type\":\"text\",\"enabled\":true},{\"key\":\"redirect_uri\",\"value\":\"https://www.getpostman.com/oauth2/callback\",\"type\":\"text\",\"enabled\":true},{\"key\":\"client_id\",\"value\":\"abc123\",\"type\":\"text\",\"enabled\":true},{\"key\":\"client_secret\",\"value\":\"ssh-secret\",\"type\":\"text\",\"enabled\":true}],\"headers\":\"\",\"dataMode\":\"params\",\"method\":\"POST\",\"tests\":\"tests[\\\"response code is 200\\\"] = responseCode.code === 200;\\nvar body = JSON.parse(responseBody);\\ntests[\\\"body has access token\\\"] = \\\"access_token\\\" in body;\\ntests[\\\"body has bearer type\\\"] = \\\"token_type\\\" in body;\",\"isFromCollection\":true,\"write\":true,\"version\":2}",
                        createdAt: "2015-11-02T13:11:08.000Z",
                        updatedAt: "2015-11-02T18:05:45.000Z",
                        write: true
                    }
                ]
            },
            {
                folder: "a8707477-63db-91c9-2676-5eeddc1db1fe",
                id: "66550458-b91f-aded-c224-755d467a63d5",
                name: "PUT Request",
                dataMode: "raw",
                data: [],
                descriptionFormat: null,
                description: "The HTTP `PUT` request method is similar to HTTP `POST`. It too is meant to \ntransfer data to a server (and elicit a response). What data is returned depends on the implementation\nof the server.\n\nA `PUT` request can pass parameters to the server using \"Query String \nParameters\", as well as the Request Body. For example, in the following \nraw HTTP request,\n\n> PUT /hi/there?hand=wave\n>\n> <request-body>\n\n\n",
                headers: "",
                method: "PUT",
                pathVariables: {},
                url: "https://echo.getpostman.com/put",
                preRequestScript: "",
                tests: "var data;\n\ntry { responseJSON = JSON.parse(responseBody); }\ncatch (e) {}\n\ntests[\"Body contains files\"] = responseBody.has(\"files\");\ntests[\"Body contains args\"] = responseBody.has(\"args\");\ntests[\"Body contains form\"] = responseBody.has(\"form\");\ntests[\"Body contains headers\"] = responseBody.has(\"headers\");\ntests[\"Body contains url\"] = responseBody.has(\"url\");\n\n\n\ntests[\"Data has been passed\"] = (responseJSON && responseJSON.data && responseJSON.data.length)",
                currentHelper: "normal",
                helperAttributes: {},
                collectionId: "85ac60ef-46ce-fa4b-2cd1-594d83e16d24",
                rawModeData: "Etiam mi lacus, cursus vitae felis et, blandit pellentesque neque. Vestibulum eget nisi a tortor commodo dignissim.\nQuisque ipsum ligula, faucibus a felis a, commodo elementum nisl. Mauris vulputate sapien et tincidunt viverra. Donec vitae velit nec metus."
            },
            {
                folder: "c2e4b340-8099-1b56-e168-04053adcb66b",
                id: "0312cc68-e332-717f-966d-0d6f376581c8",
                name: "Get Cookies",
                dataMode: "params",
                data: [],
                rawModeData: null,
                descriptionFormat: null,
                description: "Use this endpoint to get a list of all cookies that are stored with respect to this domain. Whatever key-value pairs that has been previously set by calling the \"Set Cookies\" endpoint, will be returned as response JSON.",
                headers: "",
                method: "GET",
                pathVariables: {},
                url: "https://echo.getpostman.com/cookies",
                preRequestScript: "",
                tests: "tests[\"Status code is 200\"] = responseCode.code === 200;\ntests[\"Body contains cookies\"] = responseBody.has(\"cookies\");\nvar body = JSON.parse(responseBody);\ntests[\"Body contains cookie foo2\"] = 'foo2' in body.cookies;",
                currentHelper: "normal",
                helperAttributes: {},
                collectionId: "26af32ac-04ad-eb0d-4aa5-002e15b22a11",
                responses: [
                    {
                        owner: "33232",
                        lastUpdatedBy: "33232",
                        lastRevision: 74988989,
                        request: "9783ba08-e040-732e-289e-8b4c737993b8",
                        id: "e8bb91e5-4841-b9ca-cba2-5e3520d301a3",
                        name: "Cookies",
                        status: "",
                        responseCode: {
                            code: 200,
                            name: "OK"
                        },
                        time: "462",
                        headers: [
                            {
                                name: "Access-Control-Allow-Credentials",
                                key: "Access-Control-Allow-Credentials",
                                value: ""
                            },
                            {
                                name: "Access-Control-Allow-Headers",
                                key: "Access-Control-Allow-Headers",
                                value: ""
                            },
                            {
                                name: "Access-Control-Allow-Methods",
                                key: "Access-Control-Allow-Methods",
                                value: ""
                            },
                            {
                                name: "Access-Control-Allow-Origin",
                                key: "Access-Control-Allow-Origin",
                                value: ""
                            },
                            {
                                name: "Connection",
                                key: "Connection",
                                value: "keep-alive"
                            },
                            {
                                name: "Content-Encoding",
                                key: "Content-Encoding",
                                value: "gzip"
                            },
                            {
                                name: "Content-Length",
                                key: "Content-Length",
                                value: "46"
                            },
                            {
                                name: "Content-Type",
                                key: "Content-Type",
                                value: "application/json; charset=utf-8"
                            },
                            {
                                name: "Date",
                                key: "Date",
                                value: "Thu, 29 Oct 2015 06:16:29 GMT"
                            },
                            {
                                name: "Server",
                                key: "Server",
                                value: "nginx/1.6.2"
                            },
                            {
                                name: "Vary",
                                key: "Vary",
                                value: "Accept-Encoding"
                            },
                            {
                                name: "X-Powered-By",
                                key: "X-Powered-By",
                                value: "Sails <sailsjs.org>"
                            }
                        ],
                        cookies: [],
                        mime: "",
                        text: "{\"cookies\":{\"foo2\":\"bar\"}}",
                        language: "javascript",
                        rawDataType: "text",
                        state: {
                            size: "normal"
                        },
                        previewType: "html",
                        searchResultScrolledTo: "-1",
                        version: null,
                        requestObject: "\"3285cd26-b547-021f-742b-9d9d604849b8\"",
                        createdAt: "2015-11-02T13:11:11.000Z",
                        updatedAt: "2015-11-02T13:48:46.000Z",
                        write: true
                    }
                ]
            },
            {
                folder: "f0552983-25ef-38c5-0167-3fcbbf6bf9cf",
                id: "502d53a4-6ee3-0f96-5fb1-39ecafb398eb",
                name: "GET Request ",
                dataMode: "params",
                data: [
                    {
                        key: "code",
                        value: "xWnkliVQJURqB2x1",
                        type: "text"
                    },
                    {
                        key: "grant_type",
                        value: "authorization_code",
                        type: "text"
                    },
                    {
                        key: "redirect_uri",
                        value: "https://www.getpostman.com/oauth2/callback",
                        type: "text"
                    },
                    {
                        key: "client_id",
                        value: "abc123",
                        type: "text"
                    },
                    {
                        key: "client_secret",
                        value: "ssh-secret",
                        type: "text"
                    }
                ],
                rawModeData: null,
                descriptionFormat: null,
                description: "The HTTP `GET` request method is meant to retrieve data from a server. The data\nis identified by a unique URI (Uniform Resource Identifier). \n\nA `GET` request can pass parameters to the server using \"Query String \nParameters\". For example, in the following request,\n\n> http://example.com/hi/there?hand=wave\n\nThe parameter \"hand\" has the value \"wave\".\n\nThis endpoint echoes the HTTP headers, request parameters and the complete\nURI requested.",
                headers: "",
                method: "GET",
                pathVariables: {},
                url: "https://echo.getpostman.com/get?test=123",
                preRequestScript: "",
                tests: "tests[\"Body contains headers\"] = responseBody.has(\"headers\");\ntests[\"Body contains args\"] = responseBody.has(\"args\");\ntests[\"Body contains url\"] = responseBody.has(\"url\");\n\nvar data = JSON.parse(responseBody)\n\ntests[\"Args key contains argument passed as url parameter\"] = 'test' in data.args",
                currentHelper: "normal",
                helperAttributes: {},
                collectionId: "38bb42ff-62dd-fe0e-8ce7-831b69d78a26"
            },
            {
                folder: "b0e6b174-4807-1f64-d649-14501abca93c",
                id: "6df0f817-0460-ce4d-8f24-5d1debbc6606",
                name: "Delete Cookies",
                dataMode: "params",
                data: [],
                rawModeData: null,
                descriptionFormat: null,
                description: "One or more cookies that has been set for this domain can be deleted by providing the cookie names as part of the URL parameter. The response of this request is a JSON containing the list of currently set cookies.",
                headers: "",
                method: "GET",
                pathVariables: {},
                url: "https://echo.getpostman.com/cookies/delete?foo1&foo2",
                preRequestScript: "",
                tests: "tests[\"Status code is 200\"] = responseCode.code === 200;\ntests[\"Body contains key cookies\"] = responseBody.has(\"cookies\");\nvar body = JSON.parse(responseBody);\ntests[\"Body contains cookie foo2\"] = 'foo2' in body.cookies;\ntests[\"Body does not contain cookie foo2\"] = !('foo1' in body.cookies);",
                currentHelper: "normal",
                helperAttributes: {},
                collectionId: "92ad41fc-10df-da6a-1af6-304e24f57f47",
                responses: [
                    {
                        owner: "33232",
                        lastUpdatedBy: "33232",
                        lastRevision: 74988921,
                        request: "6fa0c114-0294-af4e-1b00-3c8dfded7337",
                        id: "59571678-7875-ac3a-742a-742555225d0d",
                        name: "Cookies Response",
                        status: "",
                        responseCode: {
                            code: 200,
                            name: "OK"
                        },
                        time: "1417",
                        headers: [
                            {
                                name: "Access-Control-Allow-Credentials",
                                key: "Access-Control-Allow-Credentials",
                                value: ""
                            },
                            {
                                name: "Access-Control-Allow-Headers",
                                key: "Access-Control-Allow-Headers",
                                value: ""
                            },
                            {
                                name: "Access-Control-Allow-Methods",
                                key: "Access-Control-Allow-Methods",
                                value: ""
                            },
                            {
                                name: "Access-Control-Allow-Origin",
                                key: "Access-Control-Allow-Origin",
                                value: ""
                            },
                            {
                                name: "Connection",
                                key: "Connection",
                                value: "keep-alive"
                            },
                            {
                                name: "Content-Encoding",
                                key: "Content-Encoding",
                                value: "gzip"
                            },
                            {
                                name: "Content-Length",
                                key: "Content-Length",
                                value: "46"
                            },
                            {
                                name: "Content-Type",
                                key: "Content-Type",
                                value: "application/json; charset=utf-8"
                            },
                            {
                                name: "Date",
                                key: "Date",
                                value: "Thu, 29 Oct 2015 06:16:00 GMT"
                            },
                            {
                                name: "Server",
                                key: "Server",
                                value: "nginx/1.6.2"
                            },
                            {
                                name: "Vary",
                                key: "Vary",
                                value: "Accept-Encoding"
                            },
                            {
                                name: "X-Powered-By",
                                key: "X-Powered-By",
                                value: "Sails <sailsjs.org>"
                            }
                        ],
                        cookies: [],
                        mime: "",
                        text: "{\"cookies\":{\"foo2\":\"bar\"}}",
                        language: "javascript",
                        rawDataType: "text",
                        state: {
                            size: "normal"
                        },
                        previewType: "html",
                        searchResultScrolledTo: "-1",
                        version: null,
                        requestObject: "\"9bf1c722-5521-df8d-2f03-2b2aeedd3224\"",
                        createdAt: "2015-11-02T13:11:10.000Z",
                        updatedAt: "2015-11-02T13:48:42.000Z",
                        write: true
                    }
                ]
            },
            {
                folder: "7515811b-7011-f47a-32ee-62137772097b",
                id: "978e6a5f-a532-55c1-c46f-8ca30be09b09",
                name: "Delay Response",
                dataMode: "params",
                data: [
                    {
                        key: "test",
                        value: "123",
                        type: "text"
                    }
                ],
                rawModeData: null,
                descriptionFormat: null,
                description: "Using this endpoint one can configure how long it takes for the server to come back with a response. Appending a number to the URL defines the time (in seconds) the server will wait before responding.\n\nNote that a maximum delay of 10 seconds is accepted by the server.",
                headers: "",
                method: "GET",
                pathVariables: {},
                url: "https://echo.getpostman.com/delay/3",
                preRequestScript: "",
                tests: "tests[\"response code is 200\"] = responseCode.code === 200;\n\nvar data = JSON.parse(responseBody);\n\ntests[\"response body has key delay\"] = 'delay' in data;",
                currentHelper: "normal",
                helperAttributes: {},
                collectionId: "14db36da-67fe-cf5b-6ad4-850d03f19e25"
            },
            {
                folder: "591e7e13-46b1-37c0-0238-837f23cafb66",
                id: "215a5c1d-aa6d-6109-7ccf-a89d16087e62",
                name: "Request Headers",
                dataMode: "params",
                data: [],
                rawModeData: null,
                descriptionFormat: null,
                description: "A `GET` request to this endpoint returns the list of all request headers as part of the response JSON.\nIn Postman, sending your own set of headers through the [Headers tab](https://www.getpostman.com/docs/requests#headers) will reveal the headers as part of the response.",
                headers: "[object Object]",
                method: "GET",
                pathVariables: [],
                url: "https://echo.getpostman.com/headers",
                preRequestScript: "",
                tests: "tests[\"Body contains headers\"] = responseBody.has(\"headers\");\n\nvar data = JSON.parse(responseBody).headers;\n\ntests[\"Header contains host\"] = \"host\" in data;\ntests[\"Header contains test parameter sent as part of request header\"] = \"my-sample-header\" in data;",
                currentHelper: "normal",
                helperAttributes: {},
                collectionId: "91db36fb-11fe-fd5c-7ae6-968f98c29d56"
            },
            {
                folder: "293f5d14-92b6-71b0-8334-351c50dabf53",
                id: "3eb7a559-9968-b2ca-ddd8-9ec56e92b1f1",
                name: "Response Headers",
                dataMode: "params",
                data: [
                    {
                        key: "code",
                        value: "xWnkliVQJURqB2x1",
                        type: "text"
                    },
                    {
                        key: "grant_type",
                        value: "authorization_code",
                        type: "text"
                    },
                    {
                        key: "redirect_uri",
                        value: "https://www.getpostman.com/oauth2/callback",
                        type: "text"
                    },
                    {
                        key: "client_id",
                        value: "abc123",
                        type: "text"
                    },
                    {
                        key: "client_secret",
                        value: "ssh-secret",
                        type: "text"
                    }
                ],
                rawModeData: null,
                descriptionFormat: null,
                description: "This endpoint causes the server to send custom set of response headers. Providing header values as part of the URL parameters of a `GET` request to this endpoint returns the same as part of response header.\n\nTo send your own set of headers, simply add or replace the the URL parameters with your own set.",
                headers: "",
                method: "GET",
                pathVariables: {},
                url: "https://echo.getpostman.com/response-headers?Content-Type=text/html&Server=apibin",
                preRequestScript: "",
                tests: "tests[\"Body contains Content-Type\"] = responseBody.has(\"Content-Type\");\ntests[\"Body contains Server\"] = responseBody.has(\"Server\");",
                currentHelper: "normal",
                helperAttributes: {},
                collectionId: "63fa49eb-13dd-ad3d-1cd7-473c93b73d34"
            },
            {
                folder: "d6854260-76ab-78a3-3009-0bdfbe5bc8ee",
                id: "30ab523e-b09b-cfde-5eae-de75bfec7b60",
                name: "PATCH Request",
                dataMode: "raw",
                data: [],
                descriptionFormat: null,
                description: "The HTTP `PATCH` method is used to update resources on a server. The exact\nuse of `PATCH` requests depends on the server in question. There are a number\nof server implementations which handle `PATCH` differently. Technically, \n`PATCH` supports both Query String parameters and a Request Body.\n\nThis endpoint accepts an HTTP `PATCH` request and provides debug information\nsuch as the HTTP headers, Query String arguments, and the Request Body.",
                headers: "",
                method: "PATCH",
                pathVariables: {},
                url: "https://echo.getpostman.com/patch",
                preRequestScript: "",
                tests: "tests[\"Body contains files\"] = responseBody.has(\"files\");\ntests[\"Body contains args\"] = responseBody.has(\"args\");\ntests[\"Body contains form\"] = responseBody.has(\"form\");\ntests[\"Body contains headers\"] = responseBody.has(\"headers\");\ntests[\"Body contains url\"] = responseBody.has(\"url\");\n\nvar data = JSON.parse(responseBody)\n\ntests[\"form key has data passed in as form-data\"] = 'test' in data.form",
                currentHelper: "normal",
                helperAttributes: {},
                collectionId: "78dc32cc-04ca-bb2a-5ed8-407d49b66e36",
                rawModeData: "Curabitur auctor, elit nec pulvinar porttitor, ex augue condimentum enim, eget suscipit urna felis quis neque.\nSuspendisse sit amet luctus massa, nec venenatis mi. Suspendisse tincidunt massa at nibh efficitur fringilla. Nam quis congue mi. Etiam volutpat."
            },
            {
                folder: "c6211258-16ba-73f4-0348-0fedaa8ce0cd",
                id: "6111c2ec-a3fd-3ee9-d304-c4c72c31958f",
                name: "POST Request",
                dataMode: "raw",
                data: [],
                descriptionFormat: null,
                description: "The HTTP `POST` request method is meant to transfer data to a server \n(and elicit a response). What data is returned depends on the implementation\nof the server.\n\nA `POST` request can pass parameters to the server using \"Query String \nParameters\", as well as the Request Body. For example, in the following request,\n\n> POST /hi/there?hand=wave\n>\n> <request-body>\n\nThe parameter \"hand\" has the value \"wave\". The request body can be in multiple\nformats. These formats are defined by the MIME type of the request. The MIME \nType can be set using the ``Content-Type`` HTTP header. The most commonly used \nMIME types are:\n\n* `multipart/form-data`\n* `application/x-www-form-urlencoded`\n* `application/json`\n\nThis endpoint echoes the HTTP headers, request parameters, the contents of\nthe request body and the complete URI requested.",
                headers: "Content-Type: text/plain\n",
                method: "POST",
                pathVariables: {},
                url: "https://echo.getpostman.com/post",
                preRequestScript: "",
                tests: "var responseJSON;\n\ntry { responseJSON = JSON.parse(responseBody); }\ncatch (e) { }\n\n\ntests[\"response has data\"] = responseJSON && responseJSON.data && (responseJSON.data.length === 256);\ntests[\"content-type equals text/plain\"] = responseJSON && responseJSON.headers && (responseJSON.headers[\"content-type\"] === 'text/plain');",
                currentHelper: "normal",
                helperAttributes: {},
                collectionId: "38ad50ed-87bb-fd3a-0ad0-797e67d57a63",
                rawModeData: "Duis posuere augue vel cursus pharetra. In luctus a ex nec pretium. Praesent neque quam, tincidunt nec leo eget, rutrum vehicula magna.\nMaecenas consequat elementum elit, id semper sem tristique et. Integer pulvinar enim quis consectetur interdum volutpat."
            },
            {
                folder: "4406382e-3754-e94b-81cc-75004649315f",
                id: "6559bc82-07c1-7d1a-3532-918a28c5dc14",
                name: "Response Status Code",
                dataMode: "params",
                data: [
                    {
                        key: "test",
                        value: "123",
                        type: "text"
                    }
                ],
                rawModeData: null,
                descriptionFormat: null,
                description: "This endpoint allows one to instruct the server which status code to respond with.\n\nEvery response is accompanied by a status code. The status code provides a summary of the nature of response sent by the server. For example, a status code of `200` means everything is okay with the response and a code of `404` implies that the requested URL does not exist on server. \nA list of all valid HTTP status code can be found at the [List of Status Codes](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes) wikipedia article. When using Postman, the response status code is described for easy reference.\n\nNote that if an invalid status code is requested to be sent, the server returns a status code of `400 Bad Request`.",
                headers: "",
                method: "GET",
                pathVariables: {},
                url: "https://echo.getpostman.com/status/200",
                preRequestScript: "",
                tests: "tests[\"Body contains status\"] = responseBody.has(\"status\");\n\nvar data = JSON.parse(responseBody);\n\ntests[\"Status equals 200\"] = data.status === 200;\n",
                currentHelper: "normal",
                helperAttributes: {},
                collectionId: "50bc68fb-29fe-fa7b-2db4-043f86c32a13"
            },
            {
                folder: "0221825c-1375-e98e-38bf-73924421574b",
                id: "805c93e2-26ef-8b3f-ef59-a5582e595ba5",
                name: "Deflate Compressed Response",
                dataMode: "params",
                data: [],
                rawModeData: null,
                descriptionFormat: null,
                description: "This endpoint returns the response using [deflate compression algoritm](https://en.wikipedia.org/wiki/DEFLATE). \nThe uncompressed response is a JSON string containing the details of the request sent by the client. For this endpoint to work, one should request with `Accept-encoding` header containing `deflate` as part of its value. Postman supports gzip, deflate and SDCH decoding and automatically sends them as part of the request.\n\nHTTP Compression allows the server to send responses in a compressed format, which is uncompressed by the client before processing. This reduces network bandwidth consumption at the cost of increase in CPU usage.\nTo know more about this, refer the [HTTP Compression](https://en.wikipedia.org/wiki/HTTP_compression) wikipedia article.",
                headers: "",
                method: "GET",
                pathVariables: {},
                url: "https://echo.getpostman.com/deflate",
                preRequestScript: "",
                tests: "tests[\"response code is 200\"] = responseCode.code === 200;\n\ntry {\n    var data = JSON.parse(responseBody);\n    tests[\"Body contains deflated\"] = responseBody.has(\"deflated\");\n    tests[\"Body contains headers\"] = responseBody.has(\"headers\");\n    tests[\"Body contains method\"] = responseBody.has(\"method\");\n}\ncatch(e) {\n    console.log('Cannot parse response,probably not a JSON');\n}\n",
                currentHelper: "normal",
                helperAttributes: {},
                collectionId: "16df81eb-35bf-df7a-5ce0-721b04d72f04"
            },
            {
                folder: "a4d5e400-1758-8a38-c558-84682dcaf01f",
                id: "33f01151-4837-cf7d-cb0b-cd67e20b4ebe",
                name: "Set Cookies",
                dataMode: "params",
                data: [],
                rawModeData: null,
                descriptionFormat: null,
                description: "The cookie setter endpoint accepts a list of cookies and their values as part of URL parameters of a `GET` request. These cookies are saved and can be subsequently retrieved or deleted. The response of this request returns a JSON with all cookies listed.\n\nTo set your own set of cookies, simply replace the URL parameters \"foo1=bar1&foo2=bar2\" with your own set of key-value pairs.",
                headers: "",
                method: "GET",
                pathVariables: {},
                url: "https://echo.getpostman.com/cookies/set?foo1=bar1&foo2=bar2",
                preRequestScript: "",
                tests: "tests[\"Status code is 200\"] = responseCode.code === 200;\ntests[\"Body contains cookies\"] = responseBody.has(\"cookies\");\nvar body = JSON.parse(responseBody);\ntests[\"Body contains cookie foo1\"] = 'foo1' in body.cookies;\ntests[\"Body contains cookie foo2\"] = 'foo2' in body.cookies;\n\n",
                currentHelper: "normal",
                helperAttributes: {},
                collectionId: "40ee79cd-03da-ea5b-5ea8-822a41f65b64",
                responses: [
                    {
                        owner: "33232",
                        lastUpdatedBy: "33232",
                        lastRevision: 74988885,
                        request: "58e20371-9832-ea6c-ed9a-bd24a62c8fba",
                        id: "951317c4-6c83-267b-7a7f-a152f733f35e",
                        name: "Cookies",
                        status: "",
                        responseCode: {
                            code: 200,
                            name: "OK"
                        },
                        time: "3063",
                        headers: [
                            {
                                name: "Access-Control-Allow-Credentials",
                                key: "Access-Control-Allow-Credentials",
                                value: ""
                            },
                            {
                                name: "Access-Control-Allow-Headers",
                                key: "Access-Control-Allow-Headers",
                                value: ""
                            },
                            {
                                name: "Access-Control-Allow-Methods",
                                key: "Access-Control-Allow-Methods",
                                value: ""
                            },
                            {
                                name: "Access-Control-Allow-Origin",
                                key: "Access-Control-Allow-Origin",
                                value: ""
                            },
                            {
                                name: "Connection",
                                key: "Connection",
                                value: "keep-alive"
                            },
                            {
                                name: "Content-Encoding",
                                key: "Content-Encoding",
                                value: "gzip"
                            },
                            {
                                name: "Content-Length",
                                key: "Content-Length",
                                value: "51"
                            },
                            {
                                name: "Content-Type",
                                key: "Content-Type",
                                value: "application/json; charset=utf-8"
                            },
                            {
                                name: "Date",
                                key: "Date",
                                value: "Thu, 29 Oct 2015 06:15:28 GMT"
                            },
                            {
                                name: "Server",
                                key: "Server",
                                value: "nginx/1.6.2"
                            },
                            {
                                name: "Vary",
                                key: "Vary",
                                value: "Accept-Encoding"
                            },
                            {
                                name: "X-Powered-By",
                                key: "X-Powered-By",
                                value: "Sails <sailsjs.org>"
                            }
                        ],
                        cookies: [],
                        mime: "",
                        text: "{\"cookies\":{\"foo1\":\"bar\",\"foo2\":\"bar\"}}",
                        language: "javascript",
                        rawDataType: "text",
                        state: {
                            size: "normal"
                        },
                        previewType: "html",
                        searchResultScrolledTo: "-1",
                        version: null,
                        requestObject: "\"50d23398-4930-eb0f-ed7e-ee72b51a5cab\"",
                        createdAt: "2015-11-02T13:11:12.000Z",
                        updatedAt: "2015-11-02T13:48:39.000Z",
                        write: true
                    }
                ]
            },
            {
                folder: "6784369b-4825-a62e-77ba-96321476479d",
                id: "8ed4f4d0-03a3-dc35-cec3-fc6ac1ff6e47",
                name: "Streamed Response",
                dataMode: "params",
                data: [],
                rawModeData: null,
                descriptionFormat: null,
                description: "This endpoint allows one to recieve streaming http response using [chunked transfer encoding](https://en.wikipedia.org/wiki/Chunked_transfer_encoding) of a configurable length.\n\nA streaming response does not wait for the entire response to be generated on server before flushing it out. This implies that for a fairly large response, parts of it can be streamed to the requestee as and when it is generated on server. The client can then take actions of processing this partially received data.",
                headers: "",
                method: "GET",
                pathVariables: {},
                url: "https://echo.getpostman.com/stream/10",
                preRequestScript: "",
                tests: "tests[\"response code is 200\"] = responseCode.code === 200;\n\n",
                currentHelper: "normal",
                helperAttributes: {},
                collectionId: "52cf65aa-55af-bd6e-9fc9-142b10c92e30"
            },
            {
                folder: "aa035105-e781-1a91-d9ba-e1998cf6cd40",
                id: "d9f9544b-bb01-3d00-9917-3923672f9203",
                name: "DigestAuth Request",
                dataMode: "params",
                data: [
                    {
                        key: "code",
                        value: "xWnkliVQJURqB2x1",
                        type: "text"
                    },
                    {
                        key: "grant_type",
                        value: "authorization_code",
                        type: "text"
                    },
                    {
                        key: "redirect_uri",
                        value: "https://www.getpostman.com/oauth2/callback",
                        type: "text"
                    },
                    {
                        key: "client_id",
                        value: "abc123",
                        type: "text"
                    },
                    {
                        key: "client_secret",
                        value: "ssh-secret",
                        type: "text"
                    }
                ],
                rawModeData: null,
                descriptionFormat: null,
                description: "Performing a simple `GET` request to this endpoint returns status code `401 Unauthorized` with `WWW-Authenticate` header containing information to successfully authenticate subsequent requests.\nThe `WWW-Authenticate` header must be processed to extract `realm` and `nonce` values to hash subsequent requests.\n\nWhen this request is executed within Postman, the script attached with this request does the hard work of extracting realm and nonce from the header and set it as [global variables](https://www.getpostman.com/docs/environments#global-variables) named `echo_digest_nonce` and `echo_digest_realm`.\nThese variables are re-used in subsequent request for seamless integration of the two requests.",
                headers: "",
                method: "GET",
                pathVariables: {},
                url: "https://echo.getpostman.com/digest-auth",
                preRequestScript: "",
                tests: "tests[\"response code is 401\"] = responseCode.code === 401;\ntests[\"response has WWW-Authenticate header\"] = (postman.getResponseHeader('WWW-Authenticate'));\n\nvar authenticateHeader = postman.getResponseHeader('WWW-Authenticate'),\n    realmStart = authenticateHeader.indexOf('\"',authenticateHeader.indexOf(\"realm\")) + 1 ,\n    realmEnd = authenticateHeader.indexOf('\"',realmStart),\n    realm = authenticateHeader.slice(realmStart,realmEnd),\n    nonceStart = authenticateHeader.indexOf('\"',authenticateHeader.indexOf(\"nonce\")) + 1,\n    nonceEnd = authenticateHeader.indexOf('\"',nonceStart),\n    nonce = authenticateHeader.slice(nonceStart,nonceEnd);\n    \npostman.setGlobalVariable('echo_digest_realm', realm);\npostman.setGlobalVariable('echo_digest_nonce', nonce);",
                currentHelper: "normal",
                helperAttributes: {},
                collectionId: "73af53ab-04ad-ba6e-6af4-958d58b02d69"
            },
            {
                folder: "1936775e-8896-a36b-83ef-81539816963d",
                id: "d881d187-1448-1640-caef-85f4fcb6569d",
                name: "Get UTF8 Encoded Response",
                dataMode: "params",
                data: [],
                rawModeData: null,
                descriptionFormat: null,
                description: "If a response of an endpoint requires to send data beyond the basic English / ASCII character set, the `charset` parameter in the `Content-Type` response header defines the character encoding policy.\n\nThis endpoint returns an `UTF8` character encoded response body with text in various languages such as Greek, Latin, East Asian, etc. Postman can interpret the character encoding and use appropriate methods to display the character set in responses.",
                headers: "",
                method: "GET",
                pathVariables: {},
                url: "https://echo.getpostman.com/encoding/utf8",
                preRequestScript: "",
                tests: "tests[\"response code is 200\"] = responseCode.code === 200;",
                currentHelper: "normal",
                helperAttributes: {},
                collectionId: "81af26bf-82ef-dc5f-7bb1-290b20c87a99"
            },
            {
                folder: "d4476054-09ad-34f7-6154-2addcd0fa6ad",
                id: "f8eab216-3db5-33a6-8b34-f3c904b87085",
                name: "DELETE Request",
                dataMode: "raw",
                data: [],
                descriptionFormat: null,
                description: "The HTTP `DELETE` method is used to delete resources on a server. The exact\nuse of `DELETE` requests depends on the server implementation. In general, \n`DELETE` requests support both, Query String parameters as well as a Request \nBody.\n\nThis endpoint accepts an HTTP `DELETE` request and provides debug information\nsuch as the HTTP headers, Query String arguments, and the Request Body.",
                headers: "",
                method: "DELETE",
                pathVariables: {},
                url: "https://echo.getpostman.com/delete",
                preRequestScript: "",
                tests: "tests[\"Body contains files\"] = responseBody.has(\"files\");\ntests[\"Body contains args\"] = responseBody.has(\"args\");\ntests[\"Body contains form\"] = responseBody.has(\"form\");\ntests[\"Body contains headers\"] = responseBody.has(\"headers\");\ntests[\"Body contains url\"] = responseBody.has(\"url\");\n\nvar data = JSON.parse(responseBody)\n\ntests[\"form key has data passed in as form-data\"] = 'test' in data.form",
                currentHelper: "normal",
                helperAttributes: {},
                collectionId: "55fe74af-70ab-de0e-3ed9-125c96d87d32",
                rawModeData: "Donec fermentum, nisi sed cursus eleifend, nulla tortor ultricies tellus, ut vehicula orci arcu ut velit. In volutpat egestas dapibus. \nMorbi condimentum vestibulum sapien. Etiam dignissim diam quis eros lobortis gravida vel lobortis est. Etiam gravida sed."
            },
            {
                folder: "75e0719a-6d9d-f69f-7953-d7250928c9b8",
                id: "b00fc061-e67a-ace5-cdce-41ab09a86942",
                name: "OAuth2.0 Get Resource",
                dataMode: "params",
                data: [
                    {
                        key: "code",
                        value: "xWnkliVQJURqB2x1",
                        type: "text"
                    },
                    {
                        key: "grant_type",
                        value: "authorization_code",
                        type: "text"
                    },
                    {
                        key: "redirect_uri",
                        value: "https://www.getpostman.com/oauth2/callback",
                        type: "text"
                    },
                    {
                        key: "client_id",
                        value: "abc123",
                        type: "text"
                    },
                    {
                        key: "client_secret",
                        value: "ssh-secret",
                        type: "text"
                    }
                ],
                rawModeData: null,
                descriptionFormat: null,
                description: "Once the bearer token has been obtained, it can be passed as `Authorization` header to access user's resources.\n\n> Authorization: Bearer vp7jx...\n\nIf the `Authorization` header is not passed the endpoint returns `401 Unauthorized`.\n",
                headers: "Authorization: Bearer ls9qrFzbfcccQZu8fSeThm8ItBNniDlhPtFZYgjOG1LFikLQ87BsPG62AG0gu4SsjTCZSCjcOrVx8FfYbLPRzkjaU6BKGRJx\n",
                method: "GET",
                pathVariables: {},
                url: "https://echo.getpostman.com/oauth2/user/info",
                preRequestScript: "",
                tests: "tests[\"response code is 401\"] = (responseCode.code === 200);\nvar body = JSON.parse(responseBody);\ntests[\"body has user_id\"] = \"user_id\" in body;\ntests[\"body has name\"] = \"name\" in body;",
                currentHelper: "normal",
                helperAttributes: {},
                collectionId: "45ca21bf-41ed-da9e-7cd0-532a11a51d61",
                responses: [
                    {
                        owner: "33232",
                        lastUpdatedBy: "33232",
                        lastRevision: 74989098,
                        request: "d61ed015-e01f-efd7-aace-00aa84d68425",
                        id: "52f12736-07f0-3fad-2a07-c539472ee367",
                        name: "401",
                        status: "",
                        responseCode: {
                            code: 401,
                            name: "Unauthorized"
                        },
                        time: "303",
                        headers: [
                            {
                                name: "Access-Control-Allow-Credentials",
                                key: "Access-Control-Allow-Credentials",
                                value: ""
                            },
                            {
                                name: "Access-Control-Allow-Headers",
                                key: "Access-Control-Allow-Headers",
                                value: ""
                            },
                            {
                                name: "Access-Control-Allow-Methods",
                                key: "Access-Control-Allow-Methods",
                                value: ""
                            },
                            {
                                name: "Access-Control-Allow-Origin",
                                key: "Access-Control-Allow-Origin",
                                value: ""
                            },
                            {
                                name: "Connection",
                                key: "Connection",
                                value: "keep-alive"
                            },
                            {
                                name: "Date",
                                key: "Date",
                                value: "Sat, 31 Oct 2015 06:44:16 GMT"
                            },
                            {
                                name: "Server",
                                key: "Server",
                                value: "nginx/1.6.2"
                            },
                            {
                                name: "WWW-Authenticate",
                                key: "WWW-Authenticate",
                                value: "Bearer realm=\"Users\""
                            },
                            {
                                name: "X-Powered-By",
                                key: "X-Powered-By",
                                value: "Sails <sailsjs.org>"
                            },
                            {
                                name: "transfer-encoding",
                                key: "transfer-encoding",
                                value: "chunked"
                            }
                        ],
                        cookies: [],
                        mime: "",
                        text: "Unauthorized",
                        language: "html",
                        rawDataType: "text",
                        state: {
                            size: "normal"
                        },
                        previewType: "html",
                        searchResultScrolledTo: "-1",
                        version: null,
                        requestObject: "\"b73df531-f73c-fba6-bfbf-27eb12b95577\"",
                        createdAt: "2015-11-02T13:11:09.000Z",
                        updatedAt: "2015-11-02T13:48:57.000Z",
                        write: true
                    },
                    {
                        owner: "33232",
                        lastUpdatedBy: "33232",
                        lastRevision: 74989110,
                        request: "e58cd747-b64e-eae4-ddff-92be97d39585",
                        id: "3a370531-8624-a789-1429-b63434e68f2c",
                        name: "200",
                        status: "",
                        responseCode: {
                            code: 200,
                            name: "OK"
                        },
                        time: "323",
                        headers: [
                            {
                                name: "Access-Control-Allow-Credentials",
                                key: "Access-Control-Allow-Credentials",
                                value: ""
                            },
                            {
                                name: "Access-Control-Allow-Headers",
                                key: "Access-Control-Allow-Headers",
                                value: ""
                            },
                            {
                                name: "Access-Control-Allow-Methods",
                                key: "Access-Control-Allow-Methods",
                                value: ""
                            },
                            {
                                name: "Access-Control-Allow-Origin",
                                key: "Access-Control-Allow-Origin",
                                value: ""
                            },
                            {
                                name: "Connection",
                                key: "Connection",
                                value: "keep-alive"
                            },
                            {
                                name: "Content-Encoding",
                                key: "Content-Encoding",
                                value: "gzip"
                            },
                            {
                                name: "Content-Length",
                                key: "Content-Length",
                                value: "50"
                            },
                            {
                                name: "Content-Type",
                                key: "Content-Type",
                                value: "application/json; charset=utf-8"
                            },
                            {
                                name: "Date",
                                key: "Date",
                                value: "Sat, 31 Oct 2015 06:44:49 GMT"
                            },
                            {
                                name: "Server",
                                key: "Server",
                                value: "nginx/1.6.2"
                            },
                            {
                                name: "Vary",
                                key: "Vary",
                                value: "Accept-Encoding"
                            },
                            {
                                name: "X-Powered-By",
                                key: "X-Powered-By",
                                value: "Sails <sailsjs.org>"
                            }
                        ],
                        cookies: [],
                        mime: "",
                        text: "{\"user_id\":1,\"name\":\"postman\"}",
                        language: "javascript",
                        rawDataType: "text",
                        state: {
                            size: "normal"
                        },
                        previewType: "html",
                        searchResultScrolledTo: "-1",
                        version: null,
                        requestObject: "\"d08bd623-d65e-bde1-bdcf-90cc96a79634\"",
                        createdAt: "2015-11-02T13:11:09.000Z",
                        updatedAt: "2015-11-02T13:48:58.000Z",
                        write: true
                    }
                ]
            },
            {
                folder: "c0916291-cad8-4db2-c3cf-5dfd6650a8c5",
                id: "fca93e45-3719-00e2-47a1-82a45d45edff",
                name: "OAuth1.0 Verify Signature",
                dataMode: "params",
                data: [
                    {
                        key: "code",
                        value: "xWnkliVQJURqB2x1",
                        type: "text"
                    },
                    {
                        key: "grant_type",
                        value: "authorization_code",
                        type: "text"
                    },
                    {
                        key: "redirect_uri",
                        value: "https://www.getpostman.com/oauth2/callback",
                        type: "text"
                    },
                    {
                        key: "client_id",
                        value: "abc123",
                        type: "text"
                    },
                    {
                        key: "client_secret",
                        value: "ssh-secret",
                        type: "text"
                    }
                ],
                rawModeData: null,
                descriptionFormat: null,
                description: "OAuth1.0a is a specification that defines a protocol that can be used by one\nservice to access \"protected\" resources (endpoints) on another service. A\nmajor part of OAuth1.0 is HTTP Request Signing. This endpoint allows you to \ncheck whether the request calculation works properly in the client. \n\nThe endpoint supports the HTTP ``Authorization`` header. In case the signature\nverification fails, the endpoint provides the four debug values,\n\n* ``base_uri``\n* ``normalized_param_string``\n* ``base_string``\n* ``signing_key``\n\nFor more details about these parameters, check the [OAuth1.0a Specification](http://oauth.net/core/1.0a/)\n\nIn order to use this endpoint, you can set the following values:\n\n> Consumer Key: ``BMPGxio1tv9NP54b``\n>\n> Consumer Secret: ``I+HpB-ej$-%@2Nm9``\n\nIf you are using Postman, also check the \"Add params to header\" and \n\"Auto add parameters\" boxes.",
                headers: "Authorization: OAuth oauth_consumer_key=\"ZZXGejq0au7BV25f\",oauth_signature_method=\"HMAC-SHA1\",oauth_timestamp=\"1442394747\",oauth_nonce=\"UIGipk\",oauth_version=\"1.0\",oauth_signature=\"XxsyDTw4lkt4CZb4Ili6lMbv7Hk=\"\n",
                method: "GET",
                pathVariables: {},
                url: "https://echo.getpostman.com/oauth1",
                preRequestScript: "",
                tests: "tests[\"response code is 200\"] = responseCode.code === 200;\nvar body = JSON.parse(responseBody);\ntests[\"Body contains status pass\"] = body[\"status\"] == \"pass\"",
                currentHelper: "normal",
                helperAttributes: {},
                collectionId: "88ba82ff-26ed-bc1a-4db8-297d09d14a12"
            },
            {
                folder: "17d8721e-5c3e-a66f-2485-a3715667a8f1",
                id: "c7cddc93-0380-c02f-85aa-9612a6253f1e",
                name: "OAuth2.0 Get Auth Code",
                dataMode: "params",
                data: [
                    {
                        key: "code",
                        value: "xWnkliVQJURqB2x1",
                        type: "text"
                    },
                    {
                        key: "grant_type",
                        value: "authorization_code",
                        type: "text"
                    },
                    {
                        key: "redirect_uri",
                        value: "https://www.getpostman.com/oauth2/callback",
                        type: "text"
                    },
                    {
                        key: "client_id",
                        value: "abc123",
                        type: "text"
                    },
                    {
                        key: "client_secret",
                        value: "ssh-secret",
                        type: "text"
                    }
                ],
                rawModeData: null,
                descriptionFormat: null,
                description: "In the first request for the OAuth 2.0 flow, the client tries to get an authentication code by passing in `client_id` and `response_type`.\n\n> client_id: `abc123`\n>\n> response_type: `code`\n\nThe authentication code is returned as a part of the redirect URI and can be used to get the access token.",
                headers: "",
                method: "GET",
                pathVariables: {},
                url: "https://echo.getpostman.com/oauth2/authtoken?client_id=abc123&response_type=code",
                preRequestScript: "",
                tests: "tests[\"response code is 200\"] = responseCode.code === 200;",
                currentHelper: "normal",
                helperAttributes: {},
                collectionId: "06df67ce-58cd-ce8d-8ff8-807f50b00b01",
                responses: [
                    {
                        owner: "33232",
                        lastUpdatedBy: "33232",
                        lastRevision: 74989024,
                        request: "a7afcc64-7477-e32c-48cd-3568b3697e1d",
                        id: "a9193256-7ee1-1c94-c51b-65a24412fd2b",
                        name: "401",
                        status: "",
                        responseCode: {
                            code: 401,
                            name: "Unauthorized"
                        },
                        time: "291",
                        headers: [
                            {
                                name: "Access-Control-Allow-Credentials",
                                key: "Access-Control-Allow-Credentials",
                                value: ""
                            },
                            {
                                name: "Access-Control-Allow-Headers",
                                key: "Access-Control-Allow-Headers",
                                value: ""
                            },
                            {
                                name: "Access-Control-Allow-Methods",
                                key: "Access-Control-Allow-Methods",
                                value: ""
                            },
                            {
                                name: "Access-Control-Allow-Origin",
                                key: "Access-Control-Allow-Origin",
                                value: ""
                            },
                            {
                                name: "Connection",
                                key: "Connection",
                                value: "keep-alive"
                            },
                            {
                                name: "Content-Length",
                                key: "Content-Length",
                                value: "2"
                            },
                            {
                                name: "Content-Type",
                                key: "Content-Type",
                                value: "application/json; charset=utf-8"
                            },
                            {
                                name: "Date",
                                key: "Date",
                                value: "Sat, 31 Oct 2015 06:43:23 GMT"
                            },
                            {
                                name: "ETag",
                                key: "ETag",
                                value: "W/\"2-2745614147\""
                            },
                            {
                                name: "Server",
                                key: "Server",
                                value: "nginx/1.6.2"
                            },
                            {
                                name: "Vary",
                                key: "Vary",
                                value: "Accept-Encoding"
                            },
                            {
                                name: "X-Powered-By",
                                key: "X-Powered-By",
                                value: "Sails <sailsjs.org>"
                            }
                        ],
                        cookies: [],
                        mime: "",
                        text: "{}",
                        language: "javascript",
                        rawDataType: "text",
                        state: {
                            size: "normal"
                        },
                        previewType: "html",
                        searchResultScrolledTo: "-1",
                        version: null,
                        requestObject: "\"e4ebaf60-3858-f74c-45ab-7443a0736a9a\"",
                        createdAt: "2015-11-02T13:11:11.000Z",
                        updatedAt: "2015-11-02T13:48:49.000Z",
                        write: true
                    }
                ]
            },
            {
                folder: "4722033a-9368-b01a-18ee-94080826333e",
                id: "ea4c5208-44b3-7c41-5001-cd83933ef4db",
                name: "GZip Compressed Response",
                dataMode: "params",
                data: [
                    {
                        key: "code",
                        value: "xWnkliVQJURqB2x1",
                        type: "text"
                    },
                    {
                        key: "grant_type",
                        value: "authorization_code",
                        type: "text"
                    },
                    {
                        key: "redirect_uri",
                        value: "https://www.getpostman.com/oauth2/callback",
                        type: "text"
                    },
                    {
                        key: "client_id",
                        value: "abc123",
                        type: "text"
                    },
                    {
                        key: "client_secret",
                        value: "ssh-secret",
                        type: "text"
                    }
                ],
                rawModeData: null,
                descriptionFormat: null,
                description: "This endpoint returns the response using [gzip compression algoritm](https://en.wikipedia.org/wiki/Gzip).\nThe uncompressed response is a JSON string containing the details of the request sent by the client. For this endpoint to work, one should request with `Accept-encoding` header containing `gzip` as part of its value. Postman supports gzip, deflate and SDCH decoding and automatically sends them as part of the request.\n\nHTTP Compression allows the server to send responses in a compressed format, which is uncompressed by the client before processing. This reduces network bandwidth consumption at the cost of increase in CPU usage.\nTo know more about this, refer the [HTTP Compression](https://en.wikipedia.org/wiki/HTTP_compression) wikipedia article.",
                headers: "",
                method: "GET",
                pathVariables: {},
                url: "https://echo.getpostman.com/gzip",
                preRequestScript: "",
                tests: "tests[\"response code is 200\"] = responseCode.code === 200;\n\ntry {\n    var data = JSON.parse(responseBody);\n    tests[\"Body contains gzipped\"] = responseBody.has(\"gzipped\");\n    tests[\"Body contains headers\"] = responseBody.has(\"headers\");\n    tests[\"Body contains method\"] = responseBody.has(\"method\");\n}\ncatch(e) {\n    console.log('Cannot parse response,probably not a JSON');\n}",
                currentHelper: "normal",
                helperAttributes: {},
                collectionId: "79cf13fd-28af-cf8b-9fe1-878b04c89c60"
            },
            {
                folder: "ae918543-b654-4a30-e8fc-b7243ec0ea04",
                id: "abfb218b-5849-f79b-459e-d75f78c22add",
                name: "DigestAuth Success",
                dataMode: "params",
                data: [],
                rawModeData: null,
                descriptionFormat: null,
                description: "This endpoint sends a hashed Digest Authorization header to gain access to a valid `200 Ok` response code. In Postman, it uses the stored [global variables](https://www.getpostman.com/docs/environments#gloval-variables), `echo_digest_realm` and `echo_digest_nonce`, to generate the hashed authorisation header.\n\nWithin Postman, for this request to successfully authenticate, running the previous request \"DigestAuth Request\" stores the relevant information within the global variables.",
                headers: "Authorization: Digest username=\"postman\", realm=\"Users\", nonce=\"np1ZqU5M56BXQvgjOyYCtkOnoNsA6cdz\", uri=\"/digest-auth\", response=\"254679099562cf07df9b6f5d8d15db44\", opaque=\"\"\n",
                method: "GET",
                pathVariables: {},
                url: "https://echo.getpostman.com/digest-auth",
                preRequestScript: "",
                tests: "tests[\"response code is 200\"] = responseCode.code === 200;\ntests[\"body contains authenticated\"] = responseBody.has(\"authenticated\");",
                currentHelper: "digestAuth",
                helperAttributes: {
                    id: "digest",
                    time: 1446551845396,
                    algorithm: "MD5",
                    username: "postman",
                    realm: "{{echo_digest_realm}}",
                    password: "password",
                    nonce: "{{echo_digest_nonce}}",
                    nonceCount: "",
                    clientNonce: "",
                    opaque: "",
                    qop: "",
                    disableRetryRequest: false,
                    saveToRequest: true
                },
                collectionId: "46dc89ac-28de-ec9b-7ed8-965f79b57a77",
                responses: [
                    {
                        owner: "33232",
                        lastUpdatedBy: "33232",
                        lastRevision: 74989378,
                        request: "fcbe385e-8062-a14f-400c-c81f05f21fdb",
                        id: "36cd4976-828d-b952-8a57-5f880731b0fb",
                        name: "200",
                        status: "",
                        responseCode: {
                            code: 200,
                            name: "OK"
                        },
                        time: "9843",
                        headers: [
                            {
                                name: "Access-Control-Allow-Credentials",
                                key: "Access-Control-Allow-Credentials",
                                value: ""
                            },
                            {
                                name: "Access-Control-Allow-Headers",
                                key: "Access-Control-Allow-Headers",
                                value: ""
                            },
                            {
                                name: "Access-Control-Allow-Methods",
                                key: "Access-Control-Allow-Methods",
                                value: ""
                            },
                            {
                                name: "Access-Control-Allow-Origin",
                                key: "Access-Control-Allow-Origin",
                                value: ""
                            },
                            {
                                name: "Connection",
                                key: "Connection",
                                value: "keep-alive"
                            },
                            {
                                name: "Content-Encoding",
                                key: "Content-Encoding",
                                value: "gzip"
                            },
                            {
                                name: "Content-Length",
                                key: "Content-Length",
                                value: "42"
                            },
                            {
                                name: "Content-Type",
                                key: "Content-Type",
                                value: "application/json; charset=utf-8"
                            },
                            {
                                name: "Date",
                                key: "Date",
                                value: "Thu, 29 Oct 2015 06:17:51 GMT"
                            },
                            {
                                name: "Server",
                                key: "Server",
                                value: "nginx/1.6.2"
                            },
                            {
                                name: "Vary",
                                key: "Vary",
                                value: "Accept-Encoding"
                            },
                            {
                                name: "X-Powered-By",
                                key: "X-Powered-By",
                                value: "Sails <sailsjs.org>"
                            }
                        ],
                        cookies: [],
                        mime: "",
                        text: "{\"authenticated\":true}",
                        language: "javascript",
                        rawDataType: "text",
                        state: {
                            size: "normal"
                        },
                        previewType: "html",
                        searchResultScrolledTo: "-1",
                        version: null,
                        requestObject: "\"edda763a-1892-a76c-497e-c02d64e24bfb\"",
                        createdAt: "2015-11-02T13:11:09.000Z",
                        updatedAt: "2015-11-02T13:49:19.000Z",
                        write: true
                    }
                ]
            },
            {
                folder: "e5812336-dfd7-6bf7-a0bc-9adf0648e2b2",
                id: "cd39019e-cf15-16a4-9e93-97e73ca1f34c",
                name: "Basic Auth",
                dataMode: "params",
                data: [],
                rawModeData: null,
                descriptionFormat: null,
                description: "This endpoint simulates a **basic-auth** protected endpoint. \nThe endpoint accepts a default username and password and returns a status code of `200 ok` only if the same is provided. \nOtherwise it will return a status code `401 unauthorized`.\n\n> Username: `postman`\n> \n> Password: `password`\n\nTo use this endpoint, send a request with the header `Authorization: Basic aHRnYmFmZTp1cmxtaXhnYw==`. \nThe cryptic latter half of the header value is a base64 encoded concatenation of the default username and password. \nUsing Postman, to send this request, you can simply fill in the username and password in the \"Authorization\" tab and Postman will do the rest for you.\n\nTo know more about basic authentication, refer to the [Basic Access Authentication](https://en.wikipedia.org/wiki/Basic_access_authentication) wikipedia article.\nThe article on [authentication helpers](https://www.getpostman.com/docs/helpers#basic-auth) elaborates how to use the same within the Postman app.",
                headers: "Authorization: Basic ZWN3bWNxbjptYWlxbmJ3bw==\n",
                method: "GET",
                pathVariables: {},
                url: "https://echo.getpostman.com/basic-auth",
                preRequestScript: "",
                tests: "tests[\"response code is 200\"] = responseCode.code === 200;\ntests[\"Body contains authenticated\"] = responseBody.has(\"authenticated\");",
                currentHelper: "basicAuth",
                helperAttributes: {
                    username: "postman",
                    password: "password",
                    saveToRequest: true,
                    id: "basic",
                    timestamp: 1441099408703
                },
                collectionId: "57ea20ef-54fe-cb0e-5ec5-486f60b88d05",
                responses: [
                    {
                        owner: "33232",
                        lastUpdatedBy: "33232",
                        lastRevision: 74989338,
                        request: "ff33826e-ff08-10b2-6b15-67f85fd8f47b",
                        id: "81249603-f2e8-6c28-d89e-b95a02e6331b",
                        name: "401",
                        status: "",
                        responseCode: {
                            code: 401,
                            name: "Unauthorized"
                        },
                        time: "276",
                        headers: [
                            {
                                name: "Access-Control-Allow-Credentials",
                                key: "Access-Control-Allow-Credentials",
                                value: ""
                            },
                            {
                                name: "Access-Control-Allow-Headers",
                                key: "Access-Control-Allow-Headers",
                                value: ""
                            },
                            {
                                name: "Access-Control-Allow-Methods",
                                key: "Access-Control-Allow-Methods",
                                value: ""
                            },
                            {
                                name: "Access-Control-Allow-Origin",
                                key: "Access-Control-Allow-Origin",
                                value: ""
                            },
                            {
                                name: "Connection",
                                key: "Connection",
                                value: "keep-alive"
                            },
                            {
                                name: "Date",
                                key: "Date",
                                value: "Sat, 31 Oct 2015 06:38:51 GMT"
                            },
                            {
                                name: "Server",
                                key: "Server",
                                value: "nginx/1.6.2"
                            },
                            {
                                name: "WWW-Authenticate",
                                key: "WWW-Authenticate",
                                value: "Basic realm=\"Users\""
                            },
                            {
                                name: "X-Powered-By",
                                key: "X-Powered-By",
                                value: "Sails <sailsjs.org>"
                            },
                            {
                                name: "transfer-encoding",
                                key: "transfer-encoding",
                                value: "chunked"
                            }
                        ],
                        cookies: [],
                        mime: "",
                        text: "Unauthorized",
                        language: "html",
                        rawDataType: "text",
                        state: {
                            size: "normal"
                        },
                        previewType: "html",
                        searchResultScrolledTo: "-1",
                        version: null,
                        requestObject: "\"ba79773e-bf17-75e1-3c16-46a43bf5a49a\"",
                        createdAt: "2015-11-02T13:11:09.000Z",
                        updatedAt: "2015-11-02T13:49:16.000Z",
                        write: true
                    },
                    {
                        owner: "33232",
                        lastUpdatedBy: "33232",
                        lastRevision: 74989356,
                        request: "bb93970e-fe22-08b9-5f67-38f39fc9f31f",
                        id: "670c09c1-f1a8-9092-761b-fda25e227957",
                        name: "200",
                        status: "",
                        responseCode: {
                            code: 200,
                            name: "OK"
                        },
                        time: "377",
                        headers: [
                            {
                                name: "Access-Control-Allow-Credentials",
                                key: "Access-Control-Allow-Credentials",
                                value: ""
                            },
                            {
                                name: "Access-Control-Allow-Headers",
                                key: "Access-Control-Allow-Headers",
                                value: ""
                            },
                            {
                                name: "Access-Control-Allow-Methods",
                                key: "Access-Control-Allow-Methods",
                                value: ""
                            },
                            {
                                name: "Access-Control-Allow-Origin",
                                key: "Access-Control-Allow-Origin",
                                value: ""
                            },
                            {
                                name: "Connection",
                                key: "Connection",
                                value: "keep-alive"
                            },
                            {
                                name: "Content-Encoding",
                                key: "Content-Encoding",
                                value: "gzip"
                            },
                            {
                                name: "Content-Length",
                                key: "Content-Length",
                                value: "42"
                            },
                            {
                                name: "Content-Type",
                                key: "Content-Type",
                                value: "application/json; charset=utf-8"
                            },
                            {
                                name: "Date",
                                key: "Date",
                                value: "Sat, 31 Oct 2015 06:38:25 GMT"
                            },
                            {
                                name: "Server",
                                key: "Server",
                                value: "nginx/1.6.2"
                            },
                            {
                                name: "Vary",
                                key: "Vary",
                                value: "Accept-Encoding"
                            },
                            {
                                name: "X-Powered-By",
                                key: "X-Powered-By",
                                value: "Sails <sailsjs.org>"
                            }
                        ],
                        cookies: [],
                        mime: "",
                        text: "{\"authenticated\":true}",
                        language: "javascript",
                        rawDataType: "text",
                        state: {
                            size: "normal"
                        },
                        previewType: "html",
                        searchResultScrolledTo: "-1",
                        version: null,
                        requestObject: "\"de03503e-bc21-02a3-9d36-60b30ec1c79b\"",
                        createdAt: "2015-11-02T13:11:09.000Z",
                        updatedAt: "2015-11-02T13:49:17.000Z",
                        write: true
                    }
                ]
            }
        ]
    },
    v2: {
        info: {
            name: "Postman Echo (shamasified)",
            _postman_id: "17dd06bd-52cd-fc2f-2cd8-505c76b87d73",
            description: "Postman Echo is service you can use to test your REST clients and make sample API calls. It provides endpoints for `GET`, `POST`, `PUT`, various auth mechanisms and other utility endpoints.",
            schema: "https://schema.getpostman.com/json/collection/v2.0.0/collection.json"
        },
        item: [
            {
                _postman_id: "fc659586-d507-1a34-c9fc-f4950eb4aa21",
                name: "Auth: Digest",
                description: "Digest authentication protects an endpoint with a username and password without actually transmitting the password over network.\nOne has to apply a hash function (like MD5, etc) to the username and password before sending them over the network.\n\n> Username: `postman`\n>\n> Password: `password`\n\nUnlike Basic-Auth, authentication happens using two consecutive requests where the first request returns `401 Unauthorised` along with `WWW-Authenticate` header containing information that needs to be used to authenticate subsequent calls.\n\nTo know more about digest authentication, refer to the [Digest Access Authentication](https://en.wikipedia.org/wiki/Digest_access_authentication) wikipedia article.\nThe article on [authentication helpers](https://www.getpostman.com/docs/helpers#digest-auth) elaborates how to use the same within the Postman app.",
                item: [
                    {
                        _postman_id: "f7b1872d-fc31-9e63-5905-1086309d4084",
                        name: "DigestAuth Request",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "tests[\"response code is 401\"] = responseCode.code === 401;",
                                        "tests[\"response has WWW-Authenticate header\"] = (postman.getResponseHeader('WWW-Authenticate'));",
                                        "",
                                        "var authenticateHeader = postman.getResponseHeader('WWW-Authenticate'),",
                                        "    realmStart = authenticateHeader.indexOf('\"',authenticateHeader.indexOf(\"realm\")) + 1 ,",
                                        "    realmEnd = authenticateHeader.indexOf('\"',realmStart),",
                                        "    realm = authenticateHeader.slice(realmStart,realmEnd),",
                                        "    nonceStart = authenticateHeader.indexOf('\"',authenticateHeader.indexOf(\"nonce\")) + 1,",
                                        "    nonceEnd = authenticateHeader.indexOf('\"',nonceStart),",
                                        "    nonce = authenticateHeader.slice(nonceStart,nonceEnd);",
                                        "    ",
                                        "postman.setGlobalVariable('echo_digest_realm', realm);",
                                        "postman.setGlobalVariable('echo_digest_nonce', nonce);"
                                    ]
                                }
                            }
                        ],
                        request: {
                            url: "https://echo.getpostman.com/digest-auth",
                            method: "GET",
                            header: [],
                            body: {
                                mode: "formdata",
                                formdata: [
                                    {
                                        key: "code",
                                        value: "xWnkliVQJURqB2x1",
                                        type: "text"
                                    },
                                    {
                                        key: "grant_type",
                                        value: "authorization_code",
                                        type: "text"
                                    },
                                    {
                                        key: "redirect_uri",
                                        value: "https://www.getpostman.com/oauth2/callback",
                                        type: "text"
                                    },
                                    {
                                        key: "client_id",
                                        value: "abc123",
                                        type: "text"
                                    },
                                    {
                                        key: "client_secret",
                                        value: "ssh-secret",
                                        type: "text"
                                    }
                                ]
                            },
                            description: "Performing a simple `GET` request to this endpoint returns status code `401 Unauthorized` with `WWW-Authenticate` header containing information to successfully authenticate subsequent requests.\nThe `WWW-Authenticate` header must be processed to extract `realm` and `nonce` values to hash subsequent requests.\n\nWhen this request is executed within Postman, the script attached with this request does the hard work of extracting realm and nonce from the header and set it as [global variables](https://www.getpostman.com/docs/environments#global-variables) named `echo_digest_nonce` and `echo_digest_realm`.\nThese variables are re-used in subsequent request for seamless integration of the two requests."
                        },
                        response: []
                    },
                    {
                        _postman_id: "eaca957d-6632-e74c-193f-b30a14a67efb",
                        name: "DigestAuth Success",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "tests[\"response code is 200\"] = responseCode.code === 200;",
                                        "tests[\"body contains authenticated\"] = responseBody.has(\"authenticated\");"
                                    ]
                                }
                            }
                        ],
                        request: {
                            auth: {
                                type: "digest",
                                digest: {
                                    algorithm: "MD5",
                                    username: "postman",
                                    realm: "{{echo_digest_realm}}",
                                    password: "password",
                                    nonce: "{{echo_digest_nonce}}",
                                    nonceCount: "",
                                    clientNonce: "",
                                    opaque: "",
                                    qop: "",
                                    disableRetryRequest: false
                                }
                            },
                            url: "https://echo.getpostman.com/digest-auth",
                            method: "GET",
                            header: [
                                {
                                    key: "Authorization",
                                    value: "Digest username=\"postman\", realm=\"Users\", nonce=\"is2ClT4M30EHPczfClTLyxHddDkE0ubg\", uri=\"/digest-auth\", response=\"254679099562cf07df9b6f5d8d15db44\", opaque=\"\""
                                }
                            ],
                            body: {},
                            description: "This endpoint sends a hashed Digest Authorization header to gain access to a valid `200 Ok` response code. In Postman, it uses the stored [global variables](https://www.getpostman.com/docs/environments#gloval-variables), `echo_digest_realm` and `echo_digest_nonce`, to generate the hashed authorisation header.\n\nWithin Postman, for this request to successfully authenticate, running the previous request \"DigestAuth Request\" stores the relevant information within the global variables."
                        },
                        response: [
                            {
                                id: "47ad2611-168e-c200-5c53-2e156436a6bf",
                                name: "200",
                                originalRequest: {
                                    header: [],
                                    body: {
                                        mode: "raw",
                                        raw: ""
                                    }
                                },
                                status: "OK",
                                code: 200,
                                _postman_previewlanguage: "javascript",
                                _postman_previewtype: "html",
                                header: [
                                    {
                                        name: "Access-Control-Allow-Credentials",
                                        key: "Access-Control-Allow-Credentials",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Headers",
                                        key: "Access-Control-Allow-Headers",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Methods",
                                        key: "Access-Control-Allow-Methods",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Origin",
                                        key: "Access-Control-Allow-Origin",
                                        value: ""
                                    },
                                    {
                                        name: "Connection",
                                        key: "Connection",
                                        value: "keep-alive"
                                    },
                                    {
                                        name: "Content-Encoding",
                                        key: "Content-Encoding",
                                        value: "gzip"
                                    },
                                    {
                                        name: "Content-Length",
                                        key: "Content-Length",
                                        value: "42"
                                    },
                                    {
                                        name: "Content-Type",
                                        key: "Content-Type",
                                        value: "application/json; charset=utf-8"
                                    },
                                    {
                                        name: "Date",
                                        key: "Date",
                                        value: "Thu, 29 Oct 2015 06:17:51 GMT"
                                    },
                                    {
                                        name: "Server",
                                        key: "Server",
                                        value: "nginx/1.6.2"
                                    },
                                    {
                                        name: "Vary",
                                        key: "Vary",
                                        value: "Accept-Encoding"
                                    },
                                    {
                                        name: "X-Powered-By",
                                        key: "X-Powered-By",
                                        value: "Sails <sailsjs.org>"
                                    }
                                ],
                                cookie: [],
                                responseTime: "9843",
                                body: "{\"authenticated\":true}"
                            }
                        ]
                    }
                ]
            },
            {
                _postman_id: "b6049582-cfd3-0cb3-d7cd-3efb3594c4f3",
                name: "Auth: Others",
                item: [
                    {
                        _postman_id: "df91128d-bb38-94b8-7d33-56d59dc1f87a",
                        name: "Basic Auth",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "tests[\"response code is 200\"] = responseCode.code === 200;",
                                        "tests[\"Body contains authenticated\"] = responseBody.has(\"authenticated\");"
                                    ]
                                }
                            }
                        ],
                        request: {
                            auth: {
                                type: "basic",
                                basic: {
                                    username: "postman",
                                    password: "password",
                                    saveHelperData: true,
                                    showPassword: false
                                }
                            },
                            url: "https://echo.getpostman.com/basic-auth",
                            method: "GET",
                            header: [
                                {
                                    key: "Authorization",
                                    value: "Basic a2N3aHZ6YjptZGRpdGVyZA=="
                                }
                            ],
                            body: {},
                            description: "This endpoint simulates a **basic-auth** protected endpoint. \nThe endpoint accepts a default username and password and returns a status code of `200 ok` only if the same is provided. \nOtherwise it will return a status code `401 unauthorized`.\n\n> Username: `postman`\n> \n> Password: `password`\n\nTo use this endpoint, send a request with the header `Authorization: Basic andvcGptbTpodGVnb2dzZQ==`. \nThe cryptic latter half of the header value is a base64 encoded concatenation of the default username and password. \nUsing Postman, to send this request, you can simply fill in the username and password in the \"Authorization\" tab and Postman will do the rest for you.\n\nTo know more about basic authentication, refer to the [Basic Access Authentication](https://en.wikipedia.org/wiki/Basic_access_authentication) wikipedia article.\nThe article on [authentication helpers](https://www.getpostman.com/docs/helpers#basic-auth) elaborates how to use the same within the Postman app."
                        },
                        response: [
                            {
                                id: "39637292-d9c9-7c53-f27b-f33e25e5752f",
                                name: "401",
                                originalRequest: {
                                    header: [],
                                    body: {
                                        mode: "raw",
                                        raw: ""
                                    }
                                },
                                status: "Unauthorized",
                                code: 401,
                                _postman_previewlanguage: "html",
                                _postman_previewtype: "html",
                                header: [
                                    {
                                        name: "Access-Control-Allow-Credentials",
                                        key: "Access-Control-Allow-Credentials",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Headers",
                                        key: "Access-Control-Allow-Headers",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Methods",
                                        key: "Access-Control-Allow-Methods",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Origin",
                                        key: "Access-Control-Allow-Origin",
                                        value: ""
                                    },
                                    {
                                        name: "Connection",
                                        key: "Connection",
                                        value: "keep-alive"
                                    },
                                    {
                                        name: "Date",
                                        key: "Date",
                                        value: "Sat, 31 Oct 2015 06:38:51 GMT"
                                    },
                                    {
                                        name: "Server",
                                        key: "Server",
                                        value: "nginx/1.6.2"
                                    },
                                    {
                                        name: "WWW-Authenticate",
                                        key: "WWW-Authenticate",
                                        value: "Basic realm=\"Users\""
                                    },
                                    {
                                        name: "X-Powered-By",
                                        key: "X-Powered-By",
                                        value: "Sails <sailsjs.org>"
                                    },
                                    {
                                        name: "transfer-encoding",
                                        key: "transfer-encoding",
                                        value: "chunked"
                                    }
                                ],
                                cookie: [],
                                responseTime: "276",
                                body: "Unauthorized"
                            },
                            {
                                id: "193d36b2-f7f8-9671-039d-afa44b514730",
                                name: "200",
                                originalRequest: {
                                    header: [],
                                    body: {
                                        mode: "raw",
                                        raw: ""
                                    }
                                },
                                status: "OK",
                                code: 200,
                                _postman_previewlanguage: "javascript",
                                _postman_previewtype: "html",
                                header: [
                                    {
                                        name: "Access-Control-Allow-Credentials",
                                        key: "Access-Control-Allow-Credentials",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Headers",
                                        key: "Access-Control-Allow-Headers",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Methods",
                                        key: "Access-Control-Allow-Methods",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Origin",
                                        key: "Access-Control-Allow-Origin",
                                        value: ""
                                    },
                                    {
                                        name: "Connection",
                                        key: "Connection",
                                        value: "keep-alive"
                                    },
                                    {
                                        name: "Content-Encoding",
                                        key: "Content-Encoding",
                                        value: "gzip"
                                    },
                                    {
                                        name: "Content-Length",
                                        key: "Content-Length",
                                        value: "42"
                                    },
                                    {
                                        name: "Content-Type",
                                        key: "Content-Type",
                                        value: "application/json; charset=utf-8"
                                    },
                                    {
                                        name: "Date",
                                        key: "Date",
                                        value: "Sat, 31 Oct 2015 06:38:25 GMT"
                                    },
                                    {
                                        name: "Server",
                                        key: "Server",
                                        value: "nginx/1.6.2"
                                    },
                                    {
                                        name: "Vary",
                                        key: "Vary",
                                        value: "Accept-Encoding"
                                    },
                                    {
                                        name: "X-Powered-By",
                                        key: "X-Powered-By",
                                        value: "Sails <sailsjs.org>"
                                    }
                                ],
                                cookie: [],
                                responseTime: "377",
                                body: "{\"authenticated\":true}"
                            }
                        ]
                    },
                    {
                        _postman_id: "edd40f77-3441-60e8-15b5-66a45a09ceed",
                        name: "OAuth1.0 Verify Signature",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "tests[\"response code is 200\"] = responseCode.code === 200;",
                                        "var body = JSON.parse(responseBody);",
                                        "tests[\"Body contains status pass\"] = body[\"status\"] == \"pass\""
                                    ]
                                }
                            }
                        ],
                        request: {
                            url: "https://echo.getpostman.com/oauth1",
                            method: "GET",
                            header: [
                                {
                                    key: "Authorization",
                                    value: "OAuth oauth_consumer_key=\"QVLVvsw6jj4BN90a\",oauth_signature_method=\"HMAC-SHA1\",oauth_timestamp=\"1442394747\",oauth_nonce=\"UIGipk\",oauth_version=\"1.0\",oauth_signature=\"PialPOr0mvh3AVy0Buq2fHhu3Fk=\""
                                }
                            ],
                            body: {
                                mode: "formdata",
                                formdata: [
                                    {
                                        key: "code",
                                        value: "xWnkliVQJURqB2x1",
                                        type: "text"
                                    },
                                    {
                                        key: "grant_type",
                                        value: "authorization_code",
                                        type: "text"
                                    },
                                    {
                                        key: "redirect_uri",
                                        value: "https://www.getpostman.com/oauth2/callback",
                                        type: "text"
                                    },
                                    {
                                        key: "client_id",
                                        value: "abc123",
                                        type: "text"
                                    },
                                    {
                                        key: "client_secret",
                                        value: "ssh-secret",
                                        type: "text"
                                    }
                                ]
                            },
                            description: "OAuth1.0a is a specification that defines a protocol that can be used by one\nservice to access \"protected\" resources (endpoints) on another service. A\nmajor part of OAuth1.0 is HTTP Request Signing. This endpoint allows you to \ncheck whether the request calculation works properly in the client. \n\nThe endpoint supports the HTTP ``Authorization`` header. In case the signature\nverification fails, the endpoint provides the four debug values,\n\n* ``base_uri``\n* ``normalized_param_string``\n* ``base_string``\n* ``signing_key``\n\nFor more details about these parameters, check the [OAuth1.0a Specification](http://oauth.net/core/1.0a/)\n\nIn order to use this endpoint, you can set the following values:\n\n> Consumer Key: ``FWYFbbn8ny2WX51r``\n>\n> Consumer Secret: ``G+RiV-um$-%@2Rq3``\n\nIf you are using Postman, also check the \"Add params to header\" and \n\"Auto add parameters\" boxes."
                        },
                        response: []
                    }
                ]
            },
            {
                _postman_id: "a6c7f893-3900-6e50-a353-98933eefc91e",
                name: "Cookies",
                description: "The cookie related endpoints allow one to get, set and delete simple cookies.\n\nCookies are small snippets of information that is stored in the browser and sent back to the server with every subsequent requests in order to store useful information between requests.\nIf you want to know more about cookies, read the [HTTP Cookie](https://en.wikipedia.org/wiki/HTTP_cookie) article on wikipedia.",
                item: [
                    {
                        _postman_id: "87b93478-5129-cb0d-ad0f-ac06c18a6def",
                        name: "Set Cookies",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "tests[\"Status code is 200\"] = responseCode.code === 200;",
                                        "tests[\"Body contains cookies\"] = responseBody.has(\"cookies\");",
                                        "var body = JSON.parse(responseBody);",
                                        "tests[\"Body contains cookie foo1\"] = 'foo1' in body.cookies;",
                                        "tests[\"Body contains cookie foo2\"] = 'foo2' in body.cookies;",
                                        "",
                                        ""
                                    ]
                                }
                            }
                        ],
                        request: {
                            url: {
                                raw: "https://echo.getpostman.com/cookies/set?foo1=bar1&foo2=bar2",
                                protocol: "https",
                                host: ["echo", "getpostman", "com"],
                                path: ["cookies", "set"],
                                query: [{key: "foo1", value: "bar1"}, {key: "foo2", value: "bar2"}]
                            },
                            method: "GET",
                            header: [],
                            body: {},
                            description: "The cookie setter endpoint accepts a list of cookies and their values as part of URL parameters of a `GET` request. These cookies are saved and can be subsequently retrieved or deleted. The response of this request returns a JSON with all cookies listed.\n\nTo set your own set of cookies, simply replace the URL parameters \"foo1=bar1&foo2=bar2\" with your own set of key-value pairs."
                        },
                        response: [
                            {
                                id: "460587c2-1a60-512c-6e8c-b072f296f48a",
                                name: "Cookies",
                                originalRequest: {
                                    header: [],
                                    body: {
                                        mode: "raw",
                                        raw: ""
                                    }
                                },
                                status: "OK",
                                code: 200,
                                _postman_previewlanguage: "javascript",
                                _postman_previewtype: "html",
                                header: [
                                    {
                                        name: "Access-Control-Allow-Credentials",
                                        key: "Access-Control-Allow-Credentials",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Headers",
                                        key: "Access-Control-Allow-Headers",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Methods",
                                        key: "Access-Control-Allow-Methods",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Origin",
                                        key: "Access-Control-Allow-Origin",
                                        value: ""
                                    },
                                    {
                                        name: "Connection",
                                        key: "Connection",
                                        value: "keep-alive"
                                    },
                                    {
                                        name: "Content-Encoding",
                                        key: "Content-Encoding",
                                        value: "gzip"
                                    },
                                    {
                                        name: "Content-Length",
                                        key: "Content-Length",
                                        value: "51"
                                    },
                                    {
                                        name: "Content-Type",
                                        key: "Content-Type",
                                        value: "application/json; charset=utf-8"
                                    },
                                    {
                                        name: "Date",
                                        key: "Date",
                                        value: "Thu, 29 Oct 2015 06:15:28 GMT"
                                    },
                                    {
                                        name: "Server",
                                        key: "Server",
                                        value: "nginx/1.6.2"
                                    },
                                    {
                                        name: "Vary",
                                        key: "Vary",
                                        value: "Accept-Encoding"
                                    },
                                    {
                                        name: "X-Powered-By",
                                        key: "X-Powered-By",
                                        value: "Sails <sailsjs.org>"
                                    }
                                ],
                                cookie: [],
                                responseTime: "3063",
                                body: "{\"cookies\":{\"foo1\":\"bar\",\"foo2\":\"bar\"}}"
                            }
                        ]
                    },
                    {
                        _postman_id: "7047ca38-a195-180c-780a-3a5b792859f2",
                        name: "Get Cookies",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "tests[\"Status code is 200\"] = responseCode.code === 200;",
                                        "tests[\"Body contains cookies\"] = responseBody.has(\"cookies\");",
                                        "var body = JSON.parse(responseBody);",
                                        "tests[\"Body contains cookie foo2\"] = 'foo2' in body.cookies;"
                                    ]
                                }
                            }
                        ],
                        request: {
                            url: "https://echo.getpostman.com/cookies",
                            method: "GET",
                            header: [],
                            body: {},
                            description: "Use this endpoint to get a list of all cookies that are stored with respect to this domain. Whatever key-value pairs that has been previously set by calling the \"Set Cookies\" endpoint, will be returned as response JSON."
                        },
                        response: [
                            {
                                id: "b8fc90c3-5770-e0af-acd0-8b8426f493f3",
                                name: "Cookies",
                                originalRequest: {
                                    header: [],
                                    body: {
                                        mode: "raw",
                                        raw: ""
                                    }
                                },
                                status: "OK",
                                code: 200,
                                _postman_previewlanguage: "javascript",
                                _postman_previewtype: "html",
                                header: [
                                    {
                                        name: "Access-Control-Allow-Credentials",
                                        key: "Access-Control-Allow-Credentials",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Headers",
                                        key: "Access-Control-Allow-Headers",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Methods",
                                        key: "Access-Control-Allow-Methods",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Origin",
                                        key: "Access-Control-Allow-Origin",
                                        value: ""
                                    },
                                    {
                                        name: "Connection",
                                        key: "Connection",
                                        value: "keep-alive"
                                    },
                                    {
                                        name: "Content-Encoding",
                                        key: "Content-Encoding",
                                        value: "gzip"
                                    },
                                    {
                                        name: "Content-Length",
                                        key: "Content-Length",
                                        value: "46"
                                    },
                                    {
                                        name: "Content-Type",
                                        key: "Content-Type",
                                        value: "application/json; charset=utf-8"
                                    },
                                    {
                                        name: "Date",
                                        key: "Date",
                                        value: "Thu, 29 Oct 2015 06:16:29 GMT"
                                    },
                                    {
                                        name: "Server",
                                        key: "Server",
                                        value: "nginx/1.6.2"
                                    },
                                    {
                                        name: "Vary",
                                        key: "Vary",
                                        value: "Accept-Encoding"
                                    },
                                    {
                                        name: "X-Powered-By",
                                        key: "X-Powered-By",
                                        value: "Sails <sailsjs.org>"
                                    }
                                ],
                                cookie: [],
                                responseTime: "462",
                                body: "{\"cookies\":{\"foo2\":\"bar\"}}"
                            }
                        ]
                    },
                    {
                        _postman_id: "3ed0d123-7493-be4b-4c03-6a8ceeea8123",
                        name: "Delete Cookies",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "tests[\"Status code is 200\"] = responseCode.code === 200;",
                                        "tests[\"Body contains key cookies\"] = responseBody.has(\"cookies\");",
                                        "var body = JSON.parse(responseBody);",
                                        "tests[\"Body contains cookie foo2\"] = 'foo2' in body.cookies;",
                                        "tests[\"Body does not contain cookie foo2\"] = !('foo1' in body.cookies);"
                                    ]
                                }
                            }
                        ],
                        request: {
                            url: {
                                raw: "https://echo.getpostman.com/cookies/delete?foo1&foo2",
                                protocol: "https",
                                host: ["echo", "getpostman", "com"],
                                path: ["cookies", "delete"],
                                query: [{key: "foo1", value: ""}, {key: "foo2", value: ""}]
                            },
                            method: "GET",
                            header: [],
                            body: {},
                            description: "One or more cookies that has been set for this domain can be deleted by providing the cookie names as part of the URL parameter. The response of this request is a JSON containing the list of currently set cookies."
                        },
                        response: [
                            {
                                id: "28417323-0654-ad3a-382b-339480713e3d",
                                name: "Cookies Response",
                                originalRequest: {
                                    header: [],
                                    body: {
                                        mode: "raw",
                                        raw: ""
                                    }
                                },
                                status: "OK",
                                code: 200,
                                _postman_previewlanguage: "javascript",
                                _postman_previewtype: "html",
                                header: [
                                    {
                                        name: "Access-Control-Allow-Credentials",
                                        key: "Access-Control-Allow-Credentials",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Headers",
                                        key: "Access-Control-Allow-Headers",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Methods",
                                        key: "Access-Control-Allow-Methods",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Origin",
                                        key: "Access-Control-Allow-Origin",
                                        value: ""
                                    },
                                    {
                                        name: "Connection",
                                        key: "Connection",
                                        value: "keep-alive"
                                    },
                                    {
                                        name: "Content-Encoding",
                                        key: "Content-Encoding",
                                        value: "gzip"
                                    },
                                    {
                                        name: "Content-Length",
                                        key: "Content-Length",
                                        value: "46"
                                    },
                                    {
                                        name: "Content-Type",
                                        key: "Content-Type",
                                        value: "application/json; charset=utf-8"
                                    },
                                    {
                                        name: "Date",
                                        key: "Date",
                                        value: "Thu, 29 Oct 2015 06:16:00 GMT"
                                    },
                                    {
                                        name: "Server",
                                        key: "Server",
                                        value: "nginx/1.6.2"
                                    },
                                    {
                                        name: "Vary",
                                        key: "Vary",
                                        value: "Accept-Encoding"
                                    },
                                    {
                                        name: "X-Powered-By",
                                        key: "X-Powered-By",
                                        value: "Sails <sailsjs.org>"
                                    }
                                ],
                                cookie: [],
                                responseTime: "1417",
                                body: "{\"cookies\":{\"foo2\":\"bar\"}}"
                            }
                        ]
                    }
                ]
            },
            {
                _postman_id: "943d5c25-24f8-18e1-5803-772c76ffbe26",
                name: "Headers",
                description: "The following set of endpoints allow one to see the headers being sent as part of a request and to get a custom set of headers as part of response.\n\nHTTP header fields provide required information about the request or response, or about the object sent in the message body. Both request headers and response headers can be controlled using these endpoints.",
                item: [
                    {
                        _postman_id: "169c0f7e-ae6d-9593-9cbc-d31e73060d82",
                        name: "Request Headers",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "tests[\"Body contains headers\"] = responseBody.has(\"headers\");",
                                        "",
                                        "var data = JSON.parse(responseBody).headers;",
                                        "",
                                        "tests[\"Header contains host\"] = \"host\" in data;",
                                        "tests[\"Header contains test parameter sent as part of request header\"] = \"my-sample-header\" in data;"
                                    ]
                                }
                            }
                        ],
                        request: {
                            url: "https://echo.getpostman.com/headers",
                            method: "GET",
                            header: [],
                            body: {},
                            description: "A `GET` request to this endpoint returns the list of all request headers as part of the response JSON.\nIn Postman, sending your own set of headers through the [Headers tab](https://www.getpostman.com/docs/requests#headers) will reveal the headers as part of the response."
                        },
                        response: []
                    },
                    {
                        _postman_id: "4bc7b496-6951-c8cc-eef6-2af11e38e7f2",
                        name: "Response Headers",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "tests[\"Body contains Content-Type\"] = responseBody.has(\"Content-Type\");",
                                        "tests[\"Body contains Server\"] = responseBody.has(\"Server\");"
                                    ]
                                }
                            }
                        ],
                        request: {
                            url: {
                                raw: "https://echo.getpostman.com/response-headers?Content-Type=text/html&Server=apibin",
                                protocol: "https",
                                host: ["echo", "getpostman", "com"],
                                path: ["response-headers"],
                                query: [{key: "Content-Type", value: "text/html"}, {key: "Server", value: "apibin"}]
                            },
                            method: "GET",
                            header: [],
                            body: {
                                mode: "formdata",
                                formdata: [
                                    {
                                        key: "code",
                                        value: "xWnkliVQJURqB2x1",
                                        type: "text"
                                    },
                                    {
                                        key: "grant_type",
                                        value: "authorization_code",
                                        type: "text"
                                    },
                                    {
                                        key: "redirect_uri",
                                        value: "https://www.getpostman.com/oauth2/callback",
                                        type: "text"
                                    },
                                    {
                                        key: "client_id",
                                        value: "abc123",
                                        type: "text"
                                    },
                                    {
                                        key: "client_secret",
                                        value: "ssh-secret",
                                        type: "text"
                                    }
                                ]
                            },
                            description: "This endpoint causes the server to send custom set of response headers. Providing header values as part of the URL parameters of a `GET` request to this endpoint returns the same as part of response header.\n\nTo send your own set of headers, simply add or replace the the URL parameters with your own set."
                        },
                        response: []
                    }
                ]
            },
            {
                _postman_id: "b6433359-86ff-48f8-7515-0aacda5fb3ac",
                name: "Request Methods",
                description: "HTTP has multiple request \"verbs\", such as `GET`, `PUT`, `POST`, `DELETE`,\n`PATCH`, `HEAD`, etc. \n\nAn HTTP Method (verb) defines how a request should be interpreted by a server. \nThe endpoints in this section demonstrate various HTTP Verbs. Postman supports \nall the HTTP Verbs, including some rarely used ones, such as `PROPFIND`, `UNLINK`, \netc.\n\nFor details about HTTP Verbs, refer to [RFC 2616](http://www.w3.org/Protocols/rfc2616/rfc2616-sec9.html#sec9)\n",
                item: [
                    {
                        _postman_id: "068f51d0-2ee4-6d53-7dd3-60accde901ca",
                        name: "GET Request ",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "tests[\"Body contains headers\"] = responseBody.has(\"headers\");",
                                        "tests[\"Body contains args\"] = responseBody.has(\"args\");",
                                        "tests[\"Body contains url\"] = responseBody.has(\"url\");",
                                        "",
                                        "var data = JSON.parse(responseBody)",
                                        "",
                                        "tests[\"Args key contains argument passed as url parameter\"] = 'test' in data.args"
                                    ]
                                }
                            }
                        ],
                        request: {
                            url: {
                                raw: "https://echo.getpostman.com/get?test=123",
                                protocol: "https",
                                host: ["echo", "getpostman", "com"],
                                path: ["get"],
                                query: [{key: "test", value: "123"}]
                            },
                            method: "GET",
                            header: [],
                            body: {
                                mode: "formdata",
                                formdata: [
                                    {
                                        key: "code",
                                        value: "xWnkliVQJURqB2x1",
                                        type: "text"
                                    },
                                    {
                                        key: "grant_type",
                                        value: "authorization_code",
                                        type: "text"
                                    },
                                    {
                                        key: "redirect_uri",
                                        value: "https://www.getpostman.com/oauth2/callback",
                                        type: "text"
                                    },
                                    {
                                        key: "client_id",
                                        value: "abc123",
                                        type: "text"
                                    },
                                    {
                                        key: "client_secret",
                                        value: "ssh-secret",
                                        type: "text"
                                    }
                                ]
                            },
                            description: "The HTTP `GET` request method is meant to retrieve data from a server. The data\nis identified by a unique URI (Uniform Resource Identifier). \n\nA `GET` request can pass parameters to the server using \"Query String \nParameters\". For example, in the following request,\n\n> http://example.com/hi/there?hand=wave\n\nThe parameter \"hand\" has the value \"wave\".\n\nThis endpoint echoes the HTTP headers, request parameters and the complete\nURI requested."
                        },
                        response: []
                    },
                    {
                        _postman_id: "3898b9ed-b2be-7cc6-d999-f5a22f72884e",
                        name: "POST Request",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "var responseJSON;",
                                        "",
                                        "try { responseJSON = JSON.parse(responseBody); }",
                                        "catch (e) { }",
                                        "",
                                        "",
                                        "tests[\"response has data\"] = responseJSON && responseJSON.data && (responseJSON.data.length === 256);",
                                        "tests[\"content-type equals text/plain\"] = responseJSON && responseJSON.headers && (responseJSON.headers[\"content-type\"] === 'text/plain');"
                                    ]
                                }
                            }
                        ],
                        request: {
                            url: "https://echo.getpostman.com/post",
                            method: "POST",
                            header: [
                                {
                                    key: "Content-Type",
                                    value: "text/plain"
                                }
                            ],
                            body: {
                                mode: "raw",
                                raw: "Duis posuere augue vel cursus pharetra. In luctus a ex nec pretium. Praesent neque quam, tincidunt nec leo eget, rutrum vehicula magna.\nMaecenas consequat elementum elit, id semper sem tristique et. Integer pulvinar enim quis consectetur interdum volutpat."
                            },
                            description: "The HTTP `POST` request method is meant to transfer data to a server \n(and elicit a response). What data is returned depends on the implementation\nof the server.\n\nA `POST` request can pass parameters to the server using \"Query String \nParameters\", as well as the Request Body. For example, in the following request,\n\n> POST /hi/there?hand=wave\n>\n> <request-body>\n\nThe parameter \"hand\" has the value \"wave\". The request body can be in multiple\nformats. These formats are defined by the MIME type of the request. The MIME \nType can be set using the ``Content-Type`` HTTP header. The most commonly used \nMIME types are:\n\n* `multipart/form-data`\n* `application/x-www-form-urlencoded`\n* `application/json`\n\nThis endpoint echoes the HTTP headers, request parameters, the contents of\nthe request body and the complete URI requested."
                        },
                        response: []
                    },
                    {
                        _postman_id: "77751356-d81d-eedf-e755-642f862b92b2",
                        name: "PUT Request",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "var data;",
                                        "",
                                        "try { responseJSON = JSON.parse(responseBody); }",
                                        "catch (e) {}",
                                        "",
                                        "tests[\"Body contains files\"] = responseBody.has(\"files\");",
                                        "tests[\"Body contains args\"] = responseBody.has(\"args\");",
                                        "tests[\"Body contains form\"] = responseBody.has(\"form\");",
                                        "tests[\"Body contains headers\"] = responseBody.has(\"headers\");",
                                        "tests[\"Body contains url\"] = responseBody.has(\"url\");",
                                        "",
                                        "",
                                        "",
                                        "tests[\"Data has been passed\"] = (responseJSON && responseJSON.data && responseJSON.data.length)"
                                    ]
                                }
                            }
                        ],
                        request: {
                            url: "https://echo.getpostman.com/put",
                            method: "PUT",
                            header: [],
                            body: {
                                mode: "raw",
                                raw: "Etiam mi lacus, cursus vitae felis et, blandit pellentesque neque. Vestibulum eget nisi a tortor commodo dignissim.\nQuisque ipsum ligula, faucibus a felis a, commodo elementum nisl. Mauris vulputate sapien et tincidunt viverra. Donec vitae velit nec metus."
                            },
                            description: "The HTTP `PUT` request method is similar to HTTP `POST`. It too is meant to \ntransfer data to a server (and elicit a response). What data is returned depends on the implementation\nof the server.\n\nA `PUT` request can pass parameters to the server using \"Query String \nParameters\", as well as the Request Body. For example, in the following \nraw HTTP request,\n\n> PUT /hi/there?hand=wave\n>\n> <request-body>\n\n\n"
                        },
                        response: []
                    },
                    {
                        _postman_id: "51fd673c-e51b-aecb-4fed-dc82bcdc0a55",
                        name: "PATCH Request",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "tests[\"Body contains files\"] = responseBody.has(\"files\");",
                                        "tests[\"Body contains args\"] = responseBody.has(\"args\");",
                                        "tests[\"Body contains form\"] = responseBody.has(\"form\");",
                                        "tests[\"Body contains headers\"] = responseBody.has(\"headers\");",
                                        "tests[\"Body contains url\"] = responseBody.has(\"url\");",
                                        "",
                                        "var data = JSON.parse(responseBody)",
                                        "",
                                        "tests[\"form key has data passed in as form-data\"] = 'test' in data.form"
                                    ]
                                }
                            }
                        ],
                        request: {
                            url: "https://echo.getpostman.com/patch",
                            method: "PATCH",
                            header: [],
                            body: {
                                mode: "raw",
                                raw: "Curabitur auctor, elit nec pulvinar porttitor, ex augue condimentum enim, eget suscipit urna felis quis neque.\nSuspendisse sit amet luctus massa, nec venenatis mi. Suspendisse tincidunt massa at nibh efficitur fringilla. Nam quis congue mi. Etiam volutpat."
                            },
                            description: "The HTTP `PATCH` method is used to update resources on a server. The exact\nuse of `PATCH` requests depends on the server in question. There are a number\nof server implementations which handle `PATCH` differently. Technically, \n`PATCH` supports both Query String parameters and a Request Body.\n\nThis endpoint accepts an HTTP `PATCH` request and provides debug information\nsuch as the HTTP headers, Query String arguments, and the Request Body."
                        },
                        response: []
                    },
                    {
                        _postman_id: "b0eeb482-6aa6-06c1-9f23-c6e893c15226",
                        name: "DELETE Request",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "tests[\"Body contains files\"] = responseBody.has(\"files\");",
                                        "tests[\"Body contains args\"] = responseBody.has(\"args\");",
                                        "tests[\"Body contains form\"] = responseBody.has(\"form\");",
                                        "tests[\"Body contains headers\"] = responseBody.has(\"headers\");",
                                        "tests[\"Body contains url\"] = responseBody.has(\"url\");",
                                        "",
                                        "var data = JSON.parse(responseBody)",
                                        "",
                                        "tests[\"form key has data passed in as form-data\"] = 'test' in data.form"
                                    ]
                                }
                            }
                        ],
                        request: {
                            url: "https://echo.getpostman.com/delete",
                            method: "DELETE",
                            header: [],
                            body: {
                                mode: "raw",
                                raw: "Donec fermentum, nisi sed cursus eleifend, nulla tortor ultricies tellus, ut vehicula orci arcu ut velit. In volutpat egestas dapibus. \nMorbi condimentum vestibulum sapien. Etiam dignissim diam quis eros lobortis gravida vel lobortis est. Etiam gravida sed."
                            },
                            description: "The HTTP `DELETE` method is used to delete resources on a server. The exact\nuse of `DELETE` requests depends on the server implementation. In general, \n`DELETE` requests support both, Query String parameters as well as a Request \nBody.\n\nThis endpoint accepts an HTTP `DELETE` request and provides debug information\nsuch as the HTTP headers, Query String arguments, and the Request Body."
                        },
                        response: []
                    }
                ]
            },
            {
                _postman_id: "0615548d-5422-e93d-35fa-09200252084c",
                name: "Utilities",
                item: [
                    {
                        _postman_id: "2588db73-87a1-0a5a-5396-481e26f1cf80",
                        name: "Response Status Code",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "tests[\"Body contains status\"] = responseBody.has(\"status\");",
                                        "",
                                        "var data = JSON.parse(responseBody);",
                                        "",
                                        "tests[\"Status equals 200\"] = data.status === 200;",
                                        ""
                                    ]
                                }
                            }
                        ],
                        request: {
                            url: "https://echo.getpostman.com/status/200",
                            method: "GET",
                            header: [],
                            body: {
                                mode: "formdata",
                                formdata: [
                                    {
                                        key: "test",
                                        value: "123",
                                        type: "text"
                                    }
                                ]
                            },
                            description: "This endpoint allows one to instruct the server which status code to respond with.\n\nEvery response is accompanied by a status code. The status code provides a summary of the nature of response sent by the server. For example, a status code of `200` means everything is okay with the response and a code of `404` implies that the requested URL does not exist on server. \nA list of all valid HTTP status code can be found at the [List of Status Codes](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes) wikipedia article. When using Postman, the response status code is described for easy reference.\n\nNote that if an invalid status code is requested to be sent, the server returns a status code of `400 Bad Request`."
                        },
                        response: []
                    },
                    {
                        _postman_id: "5dc4c0a2-22c5-aa32-eea1-fd2fd6ea4f99",
                        name: "Streamed Response",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "tests[\"response code is 200\"] = responseCode.code === 200;",
                                        "",
                                        ""
                                    ]
                                }
                            }
                        ],
                        request: {
                            url: "https://echo.getpostman.com/stream/10",
                            method: "GET",
                            header: [],
                            body: {},
                            description: "This endpoint allows one to recieve streaming http response using [chunked transfer encoding](https://en.wikipedia.org/wiki/Chunked_transfer_encoding) of a configurable length.\n\nA streaming response does not wait for the entire response to be generated on server before flushing it out. This implies that for a fairly large response, parts of it can be streamed to the requestee as and when it is generated on server. The client can then take actions of processing this partially received data."
                        },
                        response: []
                    },
                    {
                        _postman_id: "952d7f3c-b264-34d9-f55f-5ed32ec32f41",
                        name: "Delay Response",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "tests[\"response code is 200\"] = responseCode.code === 200;",
                                        "",
                                        "var data = JSON.parse(responseBody);",
                                        "",
                                        "tests[\"response body has key delay\"] = 'delay' in data;"
                                    ]
                                }
                            }
                        ],
                        request: {
                            url: "https://echo.getpostman.com/delay/3",
                            method: "GET",
                            header: [],
                            body: {
                                mode: "formdata",
                                formdata: [
                                    {
                                        key: "test",
                                        value: "123",
                                        type: "text"
                                    }
                                ]
                            },
                            description: "Using this endpoint one can configure how long it takes for the server to come back with a response. Appending a number to the URL defines the time (in seconds) the server will wait before responding.\n\nNote that a maximum delay of 10 seconds is accepted by the server."
                        },
                        response: []
                    },
                    {
                        _postman_id: "e148a828-3325-7206-cfde-12a5dfa1375e",
                        name: "Get UTF8 Encoded Response",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "tests[\"response code is 200\"] = responseCode.code === 200;"
                                    ]
                                }
                            }
                        ],
                        request: {
                            url: "https://echo.getpostman.com/encoding/utf8",
                            method: "GET",
                            header: [],
                            body: {},
                            description: "If a response of an endpoint requires to send data beyond the basic English / ASCII character set, the `charset` parameter in the `Content-Type` response header defines the character encoding policy.\n\nThis endpoint returns an `UTF8` character encoded response body with text in various languages such as Greek, Latin, East Asian, etc. Postman can interpret the character encoding and use appropriate methods to display the character set in responses."
                        },
                        response: []
                    },
                    {
                        _postman_id: "cc0c4761-40a0-2f02-9155-fc17173dc7ae",
                        name: "GZip Compressed Response",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "tests[\"response code is 200\"] = responseCode.code === 200;",
                                        "",
                                        "try {",
                                        "    var data = JSON.parse(responseBody);",
                                        "    tests[\"Body contains gzipped\"] = responseBody.has(\"gzipped\");",
                                        "    tests[\"Body contains headers\"] = responseBody.has(\"headers\");",
                                        "    tests[\"Body contains method\"] = responseBody.has(\"method\");",
                                        "}",
                                        "catch(e) {",
                                        "    console.log('Cannot parse response,probably not a JSON');",
                                        "}"
                                    ]
                                }
                            }
                        ],
                        request: {
                            url: "https://echo.getpostman.com/gzip",
                            method: "GET",
                            header: [],
                            body: {
                                mode: "formdata",
                                formdata: [
                                    {
                                        key: "code",
                                        value: "xWnkliVQJURqB2x1",
                                        type: "text"
                                    },
                                    {
                                        key: "grant_type",
                                        value: "authorization_code",
                                        type: "text"
                                    },
                                    {
                                        key: "redirect_uri",
                                        value: "https://www.getpostman.com/oauth2/callback",
                                        type: "text"
                                    },
                                    {
                                        key: "client_id",
                                        value: "abc123",
                                        type: "text"
                                    },
                                    {
                                        key: "client_secret",
                                        value: "ssh-secret",
                                        type: "text"
                                    }
                                ]
                            },
                            description: "This endpoint returns the response using [gzip compression algoritm](https://en.wikipedia.org/wiki/Gzip).\nThe uncompressed response is a JSON string containing the details of the request sent by the client. For this endpoint to work, one should request with `Accept-encoding` header containing `gzip` as part of its value. Postman supports gzip, deflate and SDCH decoding and automatically sends them as part of the request.\n\nHTTP Compression allows the server to send responses in a compressed format, which is uncompressed by the client before processing. This reduces network bandwidth consumption at the cost of increase in CPU usage.\nTo know more about this, refer the [HTTP Compression](https://en.wikipedia.org/wiki/HTTP_compression) wikipedia article."
                        },
                        response: []
                    },
                    {
                        _postman_id: "515b33c1-28bb-6e5f-fe79-e9482f635bf1",
                        name: "Deflate Compressed Response",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "tests[\"response code is 200\"] = responseCode.code === 200;",
                                        "",
                                        "try {",
                                        "    var data = JSON.parse(responseBody);",
                                        "    tests[\"Body contains deflated\"] = responseBody.has(\"deflated\");",
                                        "    tests[\"Body contains headers\"] = responseBody.has(\"headers\");",
                                        "    tests[\"Body contains method\"] = responseBody.has(\"method\");",
                                        "}",
                                        "catch(e) {",
                                        "    console.log('Cannot parse response,probably not a JSON');",
                                        "}",
                                        ""
                                    ]
                                }
                            }
                        ],
                        request: {
                            url: "https://echo.getpostman.com/deflate",
                            method: "GET",
                            header: [],
                            body: {},
                            description: "This endpoint returns the response using [deflate compression algoritm](https://en.wikipedia.org/wiki/DEFLATE). \nThe uncompressed response is a JSON string containing the details of the request sent by the client. For this endpoint to work, one should request with `Accept-encoding` header containing `deflate` as part of its value. Postman supports gzip, deflate and SDCH decoding and automatically sends them as part of the request.\n\nHTTP Compression allows the server to send responses in a compressed format, which is uncompressed by the client before processing. This reduces network bandwidth consumption at the cost of increase in CPU usage.\nTo know more about this, refer the [HTTP Compression](https://en.wikipedia.org/wiki/HTTP_compression) wikipedia article."
                        },
                        response: []
                    }
                ]
            },
            {
                _postman_id: "96a8035c-6b3b-f58b-0283-b5960491b8e3",
                name: "[draft] Auth: OAuth2.0",
                description: "OAuth 2.0 allows a client to access a resource owners(user) protected resources without direct access to the users credentials.\nThe client obtains an access token which can be used to access the resource owners resources.\nOAuth 2.0 implements 4 different types of grant flows. Postman Echo has implemented the [Authorization Code](https://tools.ietf.org/html/rfc6749#section-1.3.1) grant flow.\n\nThe OAuth flow typically consists of three steps. \nIn the first step the client requests an authentication code from the server implementing the OAuth flow.\nThe authentication code along with other client credentials is then used to get the access token.\nThe access token can be used by the client to access users resources.\n",
                item: [
                    {
                        _postman_id: "c6fdfd74-4083-f23a-29ad-0576e8192f2c",
                        name: "OAuth2.0 Get Auth Code",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "tests[\"response code is 200\"] = responseCode.code === 200;"
                                    ]
                                }
                            }
                        ],
                        request: {
                            url: {
                                raw: "https://echo.getpostman.com/oauth2/authtoken?client_id=abc123&response_type=code",
                                protocol: 'https',
                                host: ['echo', 'getpostman', 'com'],
                                path: ['oauth2', 'authtoken'],
                                query: [{key: 'client_id', value: 'abc123'}, {key: 'response_type', value: 'code'}]
                            },
                            method: "GET",
                            header: [],
                            body: {
                                mode: "formdata",
                                formdata: [
                                    {
                                        key: "code",
                                        value: "xWnkliVQJURqB2x1",
                                        type: "text"
                                    },
                                    {
                                        key: "grant_type",
                                        value: "authorization_code",
                                        type: "text"
                                    },
                                    {
                                        key: "redirect_uri",
                                        value: "https://www.getpostman.com/oauth2/callback",
                                        type: "text"
                                    },
                                    {
                                        key: "client_id",
                                        value: "abc123",
                                        type: "text"
                                    },
                                    {
                                        key: "client_secret",
                                        value: "ssh-secret",
                                        type: "text"
                                    }
                                ]
                            },
                            description: "In the first request for the OAuth 2.0 flow, the client tries to get an authentication code by passing in `client_id` and `response_type`.\n\n> client_id: `abc123`\n>\n> response_type: `code`\n\nThe authentication code is returned as a part of the redirect URI and can be used to get the access token."
                        },
                        response: [
                            {
                                id: "a3541175-4df0-3a76-c19b-78e48454dc9b",
                                name: "401",
                                originalRequest: {
                                    header: [],
                                    body: {
                                        mode: "raw",
                                        raw: ""
                                    }
                                },
                                status: "Unauthorized",
                                code: 401,
                                _postman_previewlanguage: "javascript",
                                _postman_previewtype: "html",
                                header: [
                                    {
                                        name: "Access-Control-Allow-Credentials",
                                        key: "Access-Control-Allow-Credentials",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Headers",
                                        key: "Access-Control-Allow-Headers",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Methods",
                                        key: "Access-Control-Allow-Methods",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Origin",
                                        key: "Access-Control-Allow-Origin",
                                        value: ""
                                    },
                                    {
                                        name: "Connection",
                                        key: "Connection",
                                        value: "keep-alive"
                                    },
                                    {
                                        name: "Content-Length",
                                        key: "Content-Length",
                                        value: "2"
                                    },
                                    {
                                        name: "Content-Type",
                                        key: "Content-Type",
                                        value: "application/json; charset=utf-8"
                                    },
                                    {
                                        name: "Date",
                                        key: "Date",
                                        value: "Sat, 31 Oct 2015 06:43:23 GMT"
                                    },
                                    {
                                        name: "ETag",
                                        key: "ETag",
                                        value: "W/\"2-2745614147\""
                                    },
                                    {
                                        name: "Server",
                                        key: "Server",
                                        value: "nginx/1.6.2"
                                    },
                                    {
                                        name: "Vary",
                                        key: "Vary",
                                        value: "Accept-Encoding"
                                    },
                                    {
                                        name: "X-Powered-By",
                                        key: "X-Powered-By",
                                        value: "Sails <sailsjs.org>"
                                    }
                                ],
                                cookie: [],
                                responseTime: "291",
                                body: "{}"
                            }
                        ]
                    },
                    {
                        _postman_id: "6325a83f-a839-24f6-fe0f-52772637915d",
                        name: "OAuth2.0 Get Access Token",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "tests[\"response code is 200\"] = responseCode.code === 200;",
                                        "var body = JSON.parse(responseBody);",
                                        "tests[\"body has access token\"] = \"access_token\" in body;",
                                        "tests[\"body has bearer type\"] = \"token_type\" in body;"
                                    ]
                                }
                            }
                        ],
                        request: {
                            url: "https://echo.getpostman.com/oauth2/token",
                            method: "POST",
                            header: [],
                            body: {
                                mode: "formdata",
                                formdata: [
                                    {
                                        key: "code",
                                        value: "xWnkliVQJURqB2x1",
                                        type: "text"
                                    },
                                    {
                                        key: "grant_type",
                                        value: "authorization_code",
                                        type: "text"
                                    },
                                    {
                                        key: "redirect_uri",
                                        value: "https://www.getpostman.com/oauth2/callback",
                                        type: "text"
                                    },
                                    {
                                        key: "client_id",
                                        value: "abc123",
                                        type: "text"
                                    },
                                    {
                                        key: "client_secret",
                                        value: "ssh-secret",
                                        type: "text"
                                    }
                                ]
                            },
                            description: "This endpoint is used to get the `access_token`\n\nIt requires the following secret client credentials to be sent as part of the form body along with the `authentication code` obtained as part of the `redirect_uri` from the previous request.\n\n> code: xWnkliVQJURqB2x1\n>\n> grant_type: authorization_code\n>\n> redirect_uri: https://www.getpostman.com/oauth2/callback\n>\n> client_id: abc123\n>\n> client_secret: ssh-secret\n\nIf the correct credentials are not passed, the server returns with a `401 Unauthorized` response."
                        },
                        response: [
                            {
                                id: "6110021b-508e-dcb4-de5a-cdf667dfe0d7",
                                name: "401",
                                originalRequest: {
                                    url: "https://echo.getpostman.com/oauth2/token",
                                    method: "POST",
                                    header: [],
                                    body: {
                                        mode: "formdata",
                                        formdata: [
                                            {
                                                key: "code",
                                                value: "xWnkliVQJURqB2x1",
                                                type: "text"
                                            },
                                            {
                                                key: "grant_type",
                                                value: "authorization_code",
                                                type: "text"
                                            },
                                            {
                                                key: "redirect_uri",
                                                value: "https://www.getpostman.com/oauth2/callback",
                                                type: "text"
                                            },
                                            {
                                                key: "client_id",
                                                value: "abc123",
                                                type: "text"
                                            },
                                            {
                                                key: "client_secret",
                                                value: "ssh-secret",
                                                type: "text"
                                            }
                                        ]
                                    },
                                    description: "Tries to get the access token by passing the clientId , clientSecret,\nauthentication code, redirect URI and grant type.\n"
                                },
                                status: "Unauthorized",
                                code: 401,
                                _postman_previewlanguage: "javascript",
                                _postman_previewtype: "html",
                                header: [
                                    {
                                        name: "Access-Control-Allow-Credentials",
                                        key: "Access-Control-Allow-Credentials",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Headers",
                                        key: "Access-Control-Allow-Headers",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Methods",
                                        key: "Access-Control-Allow-Methods",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Origin",
                                        key: "Access-Control-Allow-Origin",
                                        value: ""
                                    },
                                    {
                                        name: "Connection",
                                        key: "Connection",
                                        value: "keep-alive"
                                    },
                                    {
                                        name: "Content-Length",
                                        key: "Content-Length",
                                        value: "2"
                                    },
                                    {
                                        name: "Content-Type",
                                        key: "Content-Type",
                                        value: "application/json; charset=utf-8"
                                    },
                                    {
                                        name: "Date",
                                        key: "Date",
                                        value: "Sat, 31 Oct 2015 06:43:44 GMT"
                                    },
                                    {
                                        name: "Server",
                                        key: "Server",
                                        value: "nginx/1.6.2"
                                    },
                                    {
                                        name: "Vary",
                                        key: "Vary",
                                        value: "X-HTTP-Method-Override, Accept-Encoding"
                                    },
                                    {
                                        name: "X-Powered-By",
                                        key: "X-Powered-By",
                                        value: "Sails <sailsjs.org>"
                                    }
                                ],
                                cookie: [],
                                responseTime: "287",
                                body: "{}"
                            },
                            {
                                id: "ae2b9408-955b-556b-a8fc-c08ef6c27422",
                                name: "200",
                                originalRequest: {
                                    url: "https://echo.getpostman.com/oauth2/token",
                                    method: "POST",
                                    header: [],
                                    body: {
                                        mode: "formdata",
                                        formdata: [
                                            {
                                                key: "code",
                                                value: "xWnkliVQJURqB2x1",
                                                type: "text"
                                            },
                                            {
                                                key: "grant_type",
                                                value: "authorization_code",
                                                type: "text"
                                            },
                                            {
                                                key: "redirect_uri",
                                                value: "https://www.getpostman.com/oauth2/callback",
                                                type: "text"
                                            },
                                            {
                                                key: "client_id",
                                                value: "abc123",
                                                type: "text"
                                            },
                                            {
                                                key: "client_secret",
                                                value: "ssh-secret",
                                                type: "text"
                                            }
                                        ]
                                    },
                                    description: "Tries to get the access token by passing the clientId , clientSecret,\nauthentication code, redirect URI and grant type.\n"
                                },
                                status: "OK",
                                code: 200,
                                _postman_previewlanguage: "javascript",
                                _postman_previewtype: "html",
                                header: [
                                    {
                                        name: "Access-Control-Allow-Credentials",
                                        key: "Access-Control-Allow-Credentials",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Headers",
                                        key: "Access-Control-Allow-Headers",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Methods",
                                        key: "Access-Control-Allow-Methods",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Origin",
                                        key: "Access-Control-Allow-Origin",
                                        value: ""
                                    },
                                    {
                                        name: "Connection",
                                        key: "Connection",
                                        value: "keep-alive"
                                    },
                                    {
                                        name: "Content-Encoding",
                                        key: "Content-Encoding",
                                        value: "gzip"
                                    },
                                    {
                                        name: "Content-Length",
                                        key: "Content-Length",
                                        value: "153"
                                    },
                                    {
                                        name: "Content-Type",
                                        key: "Content-Type",
                                        value: "application/json; charset=utf-8"
                                    },
                                    {
                                        name: "Date",
                                        key: "Date",
                                        value: "Sat, 31 Oct 2015 06:43:34 GMT"
                                    },
                                    {
                                        name: "Server",
                                        key: "Server",
                                        value: "nginx/1.6.2"
                                    },
                                    {
                                        name: "Vary",
                                        key: "Vary",
                                        value: "X-HTTP-Method-Override, Accept-Encoding"
                                    },
                                    {
                                        name: "X-Powered-By",
                                        key: "X-Powered-By",
                                        value: "Sails <sailsjs.org>"
                                    }
                                ],
                                cookie: [],
                                responseTime: "303",
                                body: "{\"access_token\":\"mp9bpThovuukSPq8lGuMwa3MwSFpcSlpWbSFBarRI6PEdeZO00DzIR14UC2ya2PoqYRLWRvsXsJl7RtJgUQPpevoL8ERUNCi\",\"token_type\":\"Bearer\"}"
                            }
                        ]
                    },
                    {
                        _postman_id: "e10ff272-f57c-add2-ebdb-14ea35a85860",
                        name: "OAuth2.0 Get Resource",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "tests[\"response code is 401\"] = (responseCode.code === 200);",
                                        "var body = JSON.parse(responseBody);",
                                        "tests[\"body has user_id\"] = \"user_id\" in body;",
                                        "tests[\"body has name\"] = \"name\" in body;"
                                    ]
                                }
                            }
                        ],
                        request: {
                            url: "https://echo.getpostman.com/oauth2/user/info",
                            method: "GET",
                            header: [
                                {
                                    key: "Authorization",
                                    value: "Bearer mg4enEibzgfuVDs0qGbYao0IuYGewHuwIkHCHbkHE9EZvsVH68LgMT77LI8vd4TxsFYVVVctFtEn2AeQsHLQjcolT1YBAQIp"
                                }
                            ],
                            body: {
                                mode: "formdata",
                                formdata: [
                                    {
                                        key: "code",
                                        value: "xWnkliVQJURqB2x1",
                                        type: "text"
                                    },
                                    {
                                        key: "grant_type",
                                        value: "authorization_code",
                                        type: "text"
                                    },
                                    {
                                        key: "redirect_uri",
                                        value: "https://www.getpostman.com/oauth2/callback",
                                        type: "text"
                                    },
                                    {
                                        key: "client_id",
                                        value: "abc123",
                                        type: "text"
                                    },
                                    {
                                        key: "client_secret",
                                        value: "ssh-secret",
                                        type: "text"
                                    }
                                ]
                            },
                            description: "Once the bearer token has been obtained, it can be passed as `Authorization` header to access user's resources.\n\n> Authorization: Bearer vp7jx...\n\nIf the `Authorization` header is not passed the endpoint returns `401 Unauthorized`.\n"
                        },
                        response: [
                            {
                                id: "09c75874-75f1-0ffe-9f88-b077762cb889",
                                name: "401",
                                originalRequest: {
                                    header: [],
                                    body: {
                                        mode: "raw",
                                        raw: ""
                                    }
                                },
                                status: "Unauthorized",
                                code: 401,
                                _postman_previewlanguage: "html",
                                _postman_previewtype: "html",
                                header: [
                                    {
                                        name: "Access-Control-Allow-Credentials",
                                        key: "Access-Control-Allow-Credentials",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Headers",
                                        key: "Access-Control-Allow-Headers",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Methods",
                                        key: "Access-Control-Allow-Methods",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Origin",
                                        key: "Access-Control-Allow-Origin",
                                        value: ""
                                    },
                                    {
                                        name: "Connection",
                                        key: "Connection",
                                        value: "keep-alive"
                                    },
                                    {
                                        name: "Date",
                                        key: "Date",
                                        value: "Sat, 31 Oct 2015 06:44:16 GMT"
                                    },
                                    {
                                        name: "Server",
                                        key: "Server",
                                        value: "nginx/1.6.2"
                                    },
                                    {
                                        name: "WWW-Authenticate",
                                        key: "WWW-Authenticate",
                                        value: "Bearer realm=\"Users\""
                                    },
                                    {
                                        name: "X-Powered-By",
                                        key: "X-Powered-By",
                                        value: "Sails <sailsjs.org>"
                                    },
                                    {
                                        name: "transfer-encoding",
                                        key: "transfer-encoding",
                                        value: "chunked"
                                    }
                                ],
                                cookie: [],
                                responseTime: "303",
                                body: "Unauthorized"
                            },
                            {
                                id: "6c616341-0111-f074-7177-a08995f37c1e",
                                name: "200",
                                originalRequest: {
                                    header: [],
                                    body: {
                                        mode: "raw",
                                        raw: ""
                                    }
                                },
                                status: "OK",
                                code: 200,
                                _postman_previewlanguage: "javascript",
                                _postman_previewtype: "html",
                                header: [
                                    {
                                        name: "Access-Control-Allow-Credentials",
                                        key: "Access-Control-Allow-Credentials",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Headers",
                                        key: "Access-Control-Allow-Headers",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Methods",
                                        key: "Access-Control-Allow-Methods",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Origin",
                                        key: "Access-Control-Allow-Origin",
                                        value: ""
                                    },
                                    {
                                        name: "Connection",
                                        key: "Connection",
                                        value: "keep-alive"
                                    },
                                    {
                                        name: "Content-Encoding",
                                        key: "Content-Encoding",
                                        value: "gzip"
                                    },
                                    {
                                        name: "Content-Length",
                                        key: "Content-Length",
                                        value: "50"
                                    },
                                    {
                                        name: "Content-Type",
                                        key: "Content-Type",
                                        value: "application/json; charset=utf-8"
                                    },
                                    {
                                        name: "Date",
                                        key: "Date",
                                        value: "Sat, 31 Oct 2015 06:44:49 GMT"
                                    },
                                    {
                                        name: "Server",
                                        key: "Server",
                                        value: "nginx/1.6.2"
                                    },
                                    {
                                        name: "Vary",
                                        key: "Vary",
                                        value: "Accept-Encoding"
                                    },
                                    {
                                        name: "X-Powered-By",
                                        key: "X-Powered-By",
                                        value: "Sails <sailsjs.org>"
                                    }
                                ],
                                cookie: [],
                                responseTime: "323",
                                body: "{\"user_id\":1,\"name\":\"postman\"}"
                            }
                        ]
                    }
                ]
            }
        ]
    },
    v21: {
        info: {
            name: "Postman Echo (shamasified)",
            _postman_id: "76bb52ff-58cd-cc2f-7cc1-165f02e79a55",
            description: "Postman Echo is service you can use to test your REST clients and make sample API calls. It provides endpoints for `GET`, `POST`, `PUT`, various auth mechanisms and other utility endpoints.",
            schema: "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
        },
        item: [
            {
                _postman_id: "ce678012-b686-8f53-c9aa-a6606ad5bb44",
                name: "Auth: Digest",
                description: "Digest authentication protects an endpoint with a username and password without actually transmitting the password over network.\nOne has to apply a hash function (like MD5, etc) to the username and password before sending them over the network.\n\n> Username: `postman`\n>\n> Password: `password`\n\nUnlike Basic-Auth, authentication happens using two consecutive requests where the first request returns `401 Unauthorised` along with `WWW-Authenticate` header containing information that needs to be used to authenticate subsequent calls.\n\nTo know more about digest authentication, refer to the [Digest Access Authentication](https://en.wikipedia.org/wiki/Digest_access_authentication) wikipedia article.\nThe article on [authentication helpers](https://www.getpostman.com/docs/helpers#digest-auth) elaborates how to use the same within the Postman app.",
                item: [
                    {
                        _postman_id: "e1f6297a-cf76-6e94-3723-0461686e7491",
                        name: "DigestAuth Request",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "tests[\"response code is 401\"] = responseCode.code === 401;",
                                        "tests[\"response has WWW-Authenticate header\"] = (postman.getResponseHeader('WWW-Authenticate'));",
                                        "",
                                        "var authenticateHeader = postman.getResponseHeader('WWW-Authenticate'),",
                                        "    realmStart = authenticateHeader.indexOf('\"',authenticateHeader.indexOf(\"realm\")) + 1 ,",
                                        "    realmEnd = authenticateHeader.indexOf('\"',realmStart),",
                                        "    realm = authenticateHeader.slice(realmStart,realmEnd),",
                                        "    nonceStart = authenticateHeader.indexOf('\"',authenticateHeader.indexOf(\"nonce\")) + 1,",
                                        "    nonceEnd = authenticateHeader.indexOf('\"',nonceStart),",
                                        "    nonce = authenticateHeader.slice(nonceStart,nonceEnd);",
                                        "    ",
                                        "postman.setGlobalVariable('echo_digest_realm', realm);",
                                        "postman.setGlobalVariable('echo_digest_nonce', nonce);"
                                    ]
                                }
                            }
                        ],
                        request: {
                            url: {
                                raw: "https://echo.getpostman.com/digest-auth",
                                protocol: "https",
                                host: ["echo", "getpostman", "com"],
                                path: ["digest-auth"]
                            },
                            method: "GET",
                            header: [],
                            body: {
                                mode: "formdata",
                                formdata: [
                                    {
                                        key: "code",
                                        value: "xWnkliVQJURqB2x1",
                                        type: "text"
                                    },
                                    {
                                        key: "grant_type",
                                        value: "authorization_code",
                                        type: "text"
                                    },
                                    {
                                        key: "redirect_uri",
                                        value: "https://www.getpostman.com/oauth2/callback",
                                        type: "text"
                                    },
                                    {
                                        key: "client_id",
                                        value: "abc123",
                                        type: "text"
                                    },
                                    {
                                        key: "client_secret",
                                        value: "ssh-secret",
                                        type: "text"
                                    }
                                ]
                            },
                            description: "Performing a simple `GET` request to this endpoint returns status code `401 Unauthorized` with `WWW-Authenticate` header containing information to successfully authenticate subsequent requests.\nThe `WWW-Authenticate` header must be processed to extract `realm` and `nonce` values to hash subsequent requests.\n\nWhen this request is executed within Postman, the script attached with this request does the hard work of extracting realm and nonce from the header and set it as [global variables](https://www.getpostman.com/docs/environments#global-variables) named `echo_digest_nonce` and `echo_digest_realm`.\nThese variables are re-used in subsequent request for seamless integration of the two requests."
                        },
                        response: []
                    },
                    {
                        _postman_id: "bdfc431e-9439-c60d-657c-b66c44f07ace",
                        name: "DigestAuth Success",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "tests[\"response code is 200\"] = responseCode.code === 200;",
                                        "tests[\"body contains authenticated\"] = responseBody.has(\"authenticated\");"
                                    ]
                                }
                            }
                        ],
                        request: {
                            auth: {
                                type: "digest",
                                digest: [
                                    { key: 'algorithm', value: 'MD5', type: 'string' },
                                    { key: 'username', value: 'postman', type: 'string' },
                                    { key: 'realm', value: '{{echo_digest_realm}}', type: 'string' },
                                    { key: 'password', value: 'password', type: 'string' },
                                    { key: 'nonce', value: '{{echo_digest_nonce}}', type: 'string' },
                                    { key: 'nonceCount', value: '', type: 'string' },
                                    { key: 'clientNonce', value: '', type: 'string' },
                                    { key: 'opaque', value: '', type: 'string' },
                                    { key: 'qop', value: '', type: 'string' },
                                    { key: 'disableRetryRequest', value: false, type: 'boolean' }
                                ]
                            },
                            url: {
                                raw: "https://echo.getpostman.com/digest-auth",
                                protocol: "https",
                                host: ["echo", "getpostman", "com"],
                                path: ["digest-auth"]
                            },
                            method: "GET",
                            header: [
                                {
                                    key: "Authorization",
                                    value: "Digest username=\"postman\", realm=\"Users\", nonce=\"fa8YqG0K75WTXhhuWtHBnvCexOfE2jhj\", uri=\"/digest-auth\", response=\"254679099562cf07df9b6f5d8d15db44\", opaque=\"\""
                                }
                            ],
                            body: {},
                            description: "This endpoint sends a hashed Digest Authorization header to gain access to a valid `200 Ok` response code. In Postman, it uses the stored [global variables](https://www.getpostman.com/docs/environments#gloval-variables), `echo_digest_realm` and `echo_digest_nonce`, to generate the hashed authorisation header.\n\nWithin Postman, for this request to successfully authenticate, running the previous request \"DigestAuth Request\" stores the relevant information within the global variables."
                        },
                        response: [
                            {
                                id: "65cb7460-458a-a869-9b60-9f914016e7ff",
                                name: "200",
                                originalRequest: {
                                    header: [],
                                    body: {
                                        mode: "raw",
                                        raw: ""
                                    }
                                },
                                status: "OK",
                                code: 200,
                                _postman_previewlanguage: "javascript",
                                _postman_previewtype: "html",
                                header: [
                                    {
                                        name: "Access-Control-Allow-Credentials",
                                        key: "Access-Control-Allow-Credentials",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Headers",
                                        key: "Access-Control-Allow-Headers",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Methods",
                                        key: "Access-Control-Allow-Methods",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Origin",
                                        key: "Access-Control-Allow-Origin",
                                        value: ""
                                    },
                                    {
                                        name: "Connection",
                                        key: "Connection",
                                        value: "keep-alive"
                                    },
                                    {
                                        name: "Content-Encoding",
                                        key: "Content-Encoding",
                                        value: "gzip"
                                    },
                                    {
                                        name: "Content-Length",
                                        key: "Content-Length",
                                        value: "42"
                                    },
                                    {
                                        name: "Content-Type",
                                        key: "Content-Type",
                                        value: "application/json; charset=utf-8"
                                    },
                                    {
                                        name: "Date",
                                        key: "Date",
                                        value: "Thu, 29 Oct 2015 06:17:51 GMT"
                                    },
                                    {
                                        name: "Server",
                                        key: "Server",
                                        value: "nginx/1.6.2"
                                    },
                                    {
                                        name: "Vary",
                                        key: "Vary",
                                        value: "Accept-Encoding"
                                    },
                                    {
                                        name: "X-Powered-By",
                                        key: "X-Powered-By",
                                        value: "Sails <sailsjs.org>"
                                    }
                                ],
                                cookie: [],
                                responseTime: "9843",
                                body: "{\"authenticated\":true}"
                            }
                        ]
                    }
                ]
            },
            {
                _postman_id: "d9479565-ecd5-7ed2-e9da-1aaa8253f2c7",
                name: "Auth: Others",
                item: [
                    {
                        _postman_id: "ac44877d-ff97-44f9-3e93-97a51ab9d82b",
                        name: "Basic Auth",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "tests[\"response code is 200\"] = responseCode.code === 200;",
                                        "tests[\"Body contains authenticated\"] = responseBody.has(\"authenticated\");"
                                    ]
                                }
                            }
                        ],
                        request: {
                            auth: {
                                type: "basic",
                                basic: [
                                    { key: 'username', value: 'postman', type: 'string' },
                                    { key: 'password', value: 'password', type: 'string' },
                                    { key: 'saveHelperData', value: true, type: 'boolean' },
                                    { key: 'showPassword', value: false, type: 'boolean' }
                                ]
                            },
                            url: {
                                raw: "https://echo.getpostman.com/basic-auth",
                                protocol: "https",
                                host: ["echo", "getpostman", "com"],
                                path: ["basic-auth"]
                            },
                            method: "GET",
                            header: [
                                {
                                    key: "Authorization",
                                    value: "Basic eXlzY3FjajpvY2VrbXd1bw=="
                                }
                            ],
                            body: {},
                            description: "This endpoint simulates a **basic-auth** protected endpoint. \nThe endpoint accepts a default username and password and returns a status code of `200 ok` only if the same is provided. \nOtherwise it will return a status code `401 unauthorized`.\n\n> Username: `postman`\n> \n> Password: `password`\n\nTo use this endpoint, send a request with the header `Authorization: Basic ZWJmeWR5czp3dnB6cGZqYw==`. \nThe cryptic latter half of the header value is a base64 encoded concatenation of the default username and password. \nUsing Postman, to send this request, you can simply fill in the username and password in the \"Authorization\" tab and Postman will do the rest for you.\n\nTo know more about basic authentication, refer to the [Basic Access Authentication](https://en.wikipedia.org/wiki/Basic_access_authentication) wikipedia article.\nThe article on [authentication helpers](https://www.getpostman.com/docs/helpers#basic-auth) elaborates how to use the same within the Postman app."
                        },
                        response: [
                            {
                                id: "85758650-f8c6-1e90-e55a-c02d82d1066e",
                                name: "401",
                                originalRequest: {
                                    header: [],
                                    body: {
                                        mode: "raw",
                                        raw: ""
                                    }
                                },
                                status: "Unauthorized",
                                code: 401,
                                _postman_previewlanguage: "html",
                                _postman_previewtype: "html",
                                header: [
                                    {
                                        name: "Access-Control-Allow-Credentials",
                                        key: "Access-Control-Allow-Credentials",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Headers",
                                        key: "Access-Control-Allow-Headers",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Methods",
                                        key: "Access-Control-Allow-Methods",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Origin",
                                        key: "Access-Control-Allow-Origin",
                                        value: ""
                                    },
                                    {
                                        name: "Connection",
                                        key: "Connection",
                                        value: "keep-alive"
                                    },
                                    {
                                        name: "Date",
                                        key: "Date",
                                        value: "Sat, 31 Oct 2015 06:38:51 GMT"
                                    },
                                    {
                                        name: "Server",
                                        key: "Server",
                                        value: "nginx/1.6.2"
                                    },
                                    {
                                        name: "WWW-Authenticate",
                                        key: "WWW-Authenticate",
                                        value: "Basic realm=\"Users\""
                                    },
                                    {
                                        name: "X-Powered-By",
                                        key: "X-Powered-By",
                                        value: "Sails <sailsjs.org>"
                                    },
                                    {
                                        name: "transfer-encoding",
                                        key: "transfer-encoding",
                                        value: "chunked"
                                    }
                                ],
                                cookie: [],
                                responseTime: "276",
                                body: "Unauthorized"
                            },
                            {
                                id: "270c73a7-d8e5-0487-665d-bba67c272244",
                                name: "200",
                                originalRequest: {
                                    header: [],
                                    body: {
                                        mode: "raw",
                                        raw: ""
                                    }
                                },
                                status: "OK",
                                code: 200,
                                _postman_previewlanguage: "javascript",
                                _postman_previewtype: "html",
                                header: [
                                    {
                                        name: "Access-Control-Allow-Credentials",
                                        key: "Access-Control-Allow-Credentials",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Headers",
                                        key: "Access-Control-Allow-Headers",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Methods",
                                        key: "Access-Control-Allow-Methods",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Origin",
                                        key: "Access-Control-Allow-Origin",
                                        value: ""
                                    },
                                    {
                                        name: "Connection",
                                        key: "Connection",
                                        value: "keep-alive"
                                    },
                                    {
                                        name: "Content-Encoding",
                                        key: "Content-Encoding",
                                        value: "gzip"
                                    },
                                    {
                                        name: "Content-Length",
                                        key: "Content-Length",
                                        value: "42"
                                    },
                                    {
                                        name: "Content-Type",
                                        key: "Content-Type",
                                        value: "application/json; charset=utf-8"
                                    },
                                    {
                                        name: "Date",
                                        key: "Date",
                                        value: "Sat, 31 Oct 2015 06:38:25 GMT"
                                    },
                                    {
                                        name: "Server",
                                        key: "Server",
                                        value: "nginx/1.6.2"
                                    },
                                    {
                                        name: "Vary",
                                        key: "Vary",
                                        value: "Accept-Encoding"
                                    },
                                    {
                                        name: "X-Powered-By",
                                        key: "X-Powered-By",
                                        value: "Sails <sailsjs.org>"
                                    }
                                ],
                                cookie: [],
                                responseTime: "377",
                                body: "{\"authenticated\":true}"
                            }
                        ]
                    },
                    {
                        _postman_id: "ebd59c95-0198-62f4-27d0-80c40d29acfd",
                        name: "OAuth1.0 Verify Signature",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "tests[\"response code is 200\"] = responseCode.code === 200;",
                                        "var body = JSON.parse(responseBody);",
                                        "tests[\"Body contains status pass\"] = body[\"status\"] == \"pass\""
                                    ]
                                }
                            }
                        ],
                        request: {
                            url: {
                                raw: "https://echo.getpostman.com/oauth1",
                                protocol: "https",
                                host: ["echo", "getpostman", "com"],
                                path: ["oauth1"]
                            },
                            method: "GET",
                            header: [
                                {
                                    key: "Authorization",
                                    value: "OAuth oauth_consumer_key=\"JKVFnpx2rr6LE94z\",oauth_signature_method=\"HMAC-SHA1\",oauth_timestamp=\"1442394747\",oauth_nonce=\"UIGipk\",oauth_version=\"1.0\",oauth_signature=\"GwmqYTb2pbv5REs9Zkz9tQrl8Tx=\""
                                }
                            ],
                            body: {
                                mode: "formdata",
                                formdata: [
                                    {
                                        key: "code",
                                        value: "xWnkliVQJURqB2x1",
                                        type: "text"
                                    },
                                    {
                                        key: "grant_type",
                                        value: "authorization_code",
                                        type: "text"
                                    },
                                    {
                                        key: "redirect_uri",
                                        value: "https://www.getpostman.com/oauth2/callback",
                                        type: "text"
                                    },
                                    {
                                        key: "client_id",
                                        value: "abc123",
                                        type: "text"
                                    },
                                    {
                                        key: "client_secret",
                                        value: "ssh-secret",
                                        type: "text"
                                    }
                                ]
                            },
                            description: "OAuth1.0a is a specification that defines a protocol that can be used by one\nservice to access \"protected\" resources (endpoints) on another service. A\nmajor part of OAuth1.0 is HTTP Request Signing. This endpoint allows you to \ncheck whether the request calculation works properly in the client. \n\nThe endpoint supports the HTTP ``Authorization`` header. In case the signature\nverification fails, the endpoint provides the four debug values,\n\n* ``base_uri``\n* ``normalized_param_string``\n* ``base_string``\n* ``signing_key``\n\nFor more details about these parameters, check the [OAuth1.0a Specification](http://oauth.net/core/1.0a/)\n\nIn order to use this endpoint, you can set the following values:\n\n> Consumer Key: ``WLPCmna7qv4TJ13l``\n>\n> Consumer Secret: ``H+MrW-iq$-%@2Dx4``\n\nIf you are using Postman, also check the \"Add params to header\" and \n\"Auto add parameters\" boxes."
                        },
                        response: []
                    }
                ]
            },
            {
                _postman_id: "d2d0e427-5596-6d46-f320-28557cccd27f",
                name: "Cookies",
                description: "The cookie related endpoints allow one to get, set and delete simple cookies.\n\nCookies are small snippets of information that is stored in the browser and sent back to the server with every subsequent requests in order to store useful information between requests.\nIf you want to know more about cookies, read the [HTTP Cookie](https://en.wikipedia.org/wiki/HTTP_cookie) article on wikipedia.",
                item: [
                    {
                        _postman_id: "30d44589-7283-cc5c-ac8e-fb15e61f1cec",
                        name: "Set Cookies",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "tests[\"Status code is 200\"] = responseCode.code === 200;",
                                        "tests[\"Body contains cookies\"] = responseBody.has(\"cookies\");",
                                        "var body = JSON.parse(responseBody);",
                                        "tests[\"Body contains cookie foo1\"] = 'foo1' in body.cookies;",
                                        "tests[\"Body contains cookie foo2\"] = 'foo2' in body.cookies;",
                                        "",
                                        ""
                                    ]
                                }
                            }
                        ],
                        request: {
                            url: {
                                raw: "https://echo.getpostman.com/cookies/set?foo1=bar1&foo2=bar2",
                                protocol: "https",
                                host: ["echo", "getpostman", "com"],
                                path: ["cookies", "set"],
                                query: [{ key: "foo1", value: "bar1" }, { key: "foo2", value: "bar2" }]
                            },
                            method: "GET",
                            header: [],
                            body: {},
                            description: "The cookie setter endpoint accepts a list of cookies and their values as part of URL parameters of a `GET` request. These cookies are saved and can be subsequently retrieved or deleted. The response of this request returns a JSON with all cookies listed.\n\nTo set your own set of cookies, simply replace the URL parameters \"foo1=bar1&foo2=bar2\" with your own set of key-value pairs."
                        },
                        response: [
                            {
                                id: "027475a2-3c47-654e-9f0e-c454c699b30c",
                                name: "Cookies",
                                originalRequest: {
                                    header: [],
                                    body: {
                                        mode: "raw",
                                        raw: ""
                                    }
                                },
                                status: "OK",
                                code: 200,
                                _postman_previewlanguage: "javascript",
                                _postman_previewtype: "html",
                                header: [
                                    {
                                        name: "Access-Control-Allow-Credentials",
                                        key: "Access-Control-Allow-Credentials",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Headers",
                                        key: "Access-Control-Allow-Headers",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Methods",
                                        key: "Access-Control-Allow-Methods",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Origin",
                                        key: "Access-Control-Allow-Origin",
                                        value: ""
                                    },
                                    {
                                        name: "Connection",
                                        key: "Connection",
                                        value: "keep-alive"
                                    },
                                    {
                                        name: "Content-Encoding",
                                        key: "Content-Encoding",
                                        value: "gzip"
                                    },
                                    {
                                        name: "Content-Length",
                                        key: "Content-Length",
                                        value: "51"
                                    },
                                    {
                                        name: "Content-Type",
                                        key: "Content-Type",
                                        value: "application/json; charset=utf-8"
                                    },
                                    {
                                        name: "Date",
                                        key: "Date",
                                        value: "Thu, 29 Oct 2015 06:15:28 GMT"
                                    },
                                    {
                                        name: "Server",
                                        key: "Server",
                                        value: "nginx/1.6.2"
                                    },
                                    {
                                        name: "Vary",
                                        key: "Vary",
                                        value: "Accept-Encoding"
                                    },
                                    {
                                        name: "X-Powered-By",
                                        key: "X-Powered-By",
                                        value: "Sails <sailsjs.org>"
                                    }
                                ],
                                cookie: [],
                                responseTime: "3063",
                                body: "{\"cookies\":{\"foo1\":\"bar\",\"foo2\":\"bar\"}}"
                            }
                        ]
                    },
                    {
                        _postman_id: "9310ae39-b405-913c-606e-4f2d884785f2",
                        name: "Get Cookies",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "tests[\"Status code is 200\"] = responseCode.code === 200;",
                                        "tests[\"Body contains cookies\"] = responseBody.has(\"cookies\");",
                                        "var body = JSON.parse(responseBody);",
                                        "tests[\"Body contains cookie foo2\"] = 'foo2' in body.cookies;"
                                    ]
                                }
                            }
                        ],
                        request: {
                            url: {
                                raw: "https://echo.getpostman.com/cookies",
                                protocol: "https",
                                host: ["echo", "getpostman", "com"],
                                path: ["cookies"]
                            },
                            method: "GET",
                            header: [],
                            body: {},
                            description: "Use this endpoint to get a list of all cookies that are stored with respect to this domain. Whatever key-value pairs that has been previously set by calling the \"Set Cookies\" endpoint, will be returned as response JSON."
                        },
                        response: [
                            {
                                id: "f0da49f4-8077-b5dc-fce3-4f3582e999c5",
                                name: "Cookies",
                                originalRequest: {
                                    header: [],
                                    body: {
                                        mode: "raw",
                                        raw: ""
                                    }
                                },
                                status: "OK",
                                code: 200,
                                _postman_previewlanguage: "javascript",
                                _postman_previewtype: "html",
                                header: [
                                    {
                                        name: "Access-Control-Allow-Credentials",
                                        key: "Access-Control-Allow-Credentials",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Headers",
                                        key: "Access-Control-Allow-Headers",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Methods",
                                        key: "Access-Control-Allow-Methods",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Origin",
                                        key: "Access-Control-Allow-Origin",
                                        value: ""
                                    },
                                    {
                                        name: "Connection",
                                        key: "Connection",
                                        value: "keep-alive"
                                    },
                                    {
                                        name: "Content-Encoding",
                                        key: "Content-Encoding",
                                        value: "gzip"
                                    },
                                    {
                                        name: "Content-Length",
                                        key: "Content-Length",
                                        value: "46"
                                    },
                                    {
                                        name: "Content-Type",
                                        key: "Content-Type",
                                        value: "application/json; charset=utf-8"
                                    },
                                    {
                                        name: "Date",
                                        key: "Date",
                                        value: "Thu, 29 Oct 2015 06:16:29 GMT"
                                    },
                                    {
                                        name: "Server",
                                        key: "Server",
                                        value: "nginx/1.6.2"
                                    },
                                    {
                                        name: "Vary",
                                        key: "Vary",
                                        value: "Accept-Encoding"
                                    },
                                    {
                                        name: "X-Powered-By",
                                        key: "X-Powered-By",
                                        value: "Sails <sailsjs.org>"
                                    }
                                ],
                                cookie: [],
                                responseTime: "462",
                                body: "{\"cookies\":{\"foo2\":\"bar\"}}"
                            }
                        ]
                    },
                    {
                        _postman_id: "7eb1d053-4367-fa7c-2b61-1a1fefcd8665",
                        name: "Delete Cookies",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "tests[\"Status code is 200\"] = responseCode.code === 200;",
                                        "tests[\"Body contains key cookies\"] = responseBody.has(\"cookies\");",
                                        "var body = JSON.parse(responseBody);",
                                        "tests[\"Body contains cookie foo2\"] = 'foo2' in body.cookies;",
                                        "tests[\"Body does not contain cookie foo2\"] = !('foo1' in body.cookies);"
                                    ]
                                }
                            }
                        ],
                        request: {
                            url: {
                                raw: "https://echo.getpostman.com/cookies/delete?foo1&foo2",
                                protocol: "https",
                                host: ["echo", "getpostman", "com"],
                                path: ["cookies", "delete"],
                                query: [{ key: "foo1", value: "" }, { key: "foo2", value: "" }]
                            },
                            method: "GET",
                            header: [],
                            body: {},
                            description: "One or more cookies that has been set for this domain can be deleted by providing the cookie names as part of the URL parameter. The response of this request is a JSON containing the list of currently set cookies."
                        },
                        response: [
                            {
                                id: "06979414-5663-ba7c-828e-270228944a9f",
                                name: "Cookies Response",
                                originalRequest: {
                                    header: [],
                                    body: {
                                        mode: "raw",
                                        raw: ""
                                    }
                                },
                                status: "OK",
                                code: 200,
                                _postman_previewlanguage: "javascript",
                                _postman_previewtype: "html",
                                header: [
                                    {
                                        name: "Access-Control-Allow-Credentials",
                                        key: "Access-Control-Allow-Credentials",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Headers",
                                        key: "Access-Control-Allow-Headers",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Methods",
                                        key: "Access-Control-Allow-Methods",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Origin",
                                        key: "Access-Control-Allow-Origin",
                                        value: ""
                                    },
                                    {
                                        name: "Connection",
                                        key: "Connection",
                                        value: "keep-alive"
                                    },
                                    {
                                        name: "Content-Encoding",
                                        key: "Content-Encoding",
                                        value: "gzip"
                                    },
                                    {
                                        name: "Content-Length",
                                        key: "Content-Length",
                                        value: "46"
                                    },
                                    {
                                        name: "Content-Type",
                                        key: "Content-Type",
                                        value: "application/json; charset=utf-8"
                                    },
                                    {
                                        name: "Date",
                                        key: "Date",
                                        value: "Thu, 29 Oct 2015 06:16:00 GMT"
                                    },
                                    {
                                        name: "Server",
                                        key: "Server",
                                        value: "nginx/1.6.2"
                                    },
                                    {
                                        name: "Vary",
                                        key: "Vary",
                                        value: "Accept-Encoding"
                                    },
                                    {
                                        name: "X-Powered-By",
                                        key: "X-Powered-By",
                                        value: "Sails <sailsjs.org>"
                                    }
                                ],
                                cookie: [],
                                responseTime: "1417",
                                body: "{\"cookies\":{\"foo2\":\"bar\"}}"
                            }
                        ]
                    }
                ]
            },
            {
                _postman_id: "633c4e43-41d3-18b8-7852-620f15dcad58",
                name: "Headers",
                description: "The following set of endpoints allow one to see the headers being sent as part of a request and to get a custom set of headers as part of response.\n\nHTTP header fields provide required information about the request or response, or about the object sent in the message body. Both request headers and response headers can be controlled using these endpoints.",
                item: [
                    {
                        _postman_id: "899e1b6f-ff3b-1884-1ebb-c73d88238a24",
                        name: "Request Headers",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "tests[\"Body contains headers\"] = responseBody.has(\"headers\");",
                                        "",
                                        "var data = JSON.parse(responseBody).headers;",
                                        "",
                                        "tests[\"Header contains host\"] = \"host\" in data;",
                                        "tests[\"Header contains test parameter sent as part of request header\"] = \"my-sample-header\" in data;"
                                    ]
                                }
                            }
                        ],
                        request: {
                            url: {
                                raw: "https://echo.getpostman.com/headers",
                                protocol: "https",
                                host: ["echo", "getpostman", "com"],
                                path: ["headers"]
                            },
                            method: "GET",
                            header: [],
                            body: {},
                            description: "A `GET` request to this endpoint returns the list of all request headers as part of the response JSON.\nIn Postman, sending your own set of headers through the [Headers tab](https://www.getpostman.com/docs/requests#headers) will reveal the headers as part of the response."
                        },
                        response: []
                    },
                    {
                        _postman_id: "4ac3e329-9401-c3ca-bec5-7fb67e20a0e4",
                        name: "Response Headers",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "tests[\"Body contains Content-Type\"] = responseBody.has(\"Content-Type\");",
                                        "tests[\"Body contains Server\"] = responseBody.has(\"Server\");"
                                    ]
                                }
                            }
                        ],
                        request: {
                            url: {
                                raw: "https://echo.getpostman.com/response-headers?Content-Type=text/html&Server=apibin",
                                protocol: "https",
                                host: ["echo", "getpostman", "com"],
                                path: ["response-headers"],
                                query: [{ key: "Content-Type", value: "text/html" }, { key: "Server", value: "apibin" }]
                            },
                            method: "GET",
                            header: [],
                            body: {
                                mode: "formdata",
                                formdata: [
                                    {
                                        key: "code",
                                        value: "xWnkliVQJURqB2x1",
                                        type: "text"
                                    },
                                    {
                                        key: "grant_type",
                                        value: "authorization_code",
                                        type: "text"
                                    },
                                    {
                                        key: "redirect_uri",
                                        value: "https://www.getpostman.com/oauth2/callback",
                                        type: "text"
                                    },
                                    {
                                        key: "client_id",
                                        value: "abc123",
                                        type: "text"
                                    },
                                    {
                                        key: "client_secret",
                                        value: "ssh-secret",
                                        type: "text"
                                    }
                                ]
                            },
                            description: "This endpoint causes the server to send custom set of response headers. Providing header values as part of the URL parameters of a `GET` request to this endpoint returns the same as part of response header.\n\nTo send your own set of headers, simply add or replace the the URL parameters with your own set."
                        },
                        response: []
                    }
                ]
            },
            {
                _postman_id: "a5549727-04bd-88c4-3590-0feabc8ac1af",
                name: "Request Methods",
                description: "HTTP has multiple request \"verbs\", such as `GET`, `PUT`, `POST`, `DELETE`,\n`PATCH`, `HEAD`, etc. \n\nAn HTTP Method (verb) defines how a request should be interpreted by a server. \nThe endpoints in this section demonstrate various HTTP Verbs. Postman supports \nall the HTTP Verbs, including some rarely used ones, such as `PROPFIND`, `UNLINK`, \netc.\n\nFor details about HTTP Verbs, refer to [RFC 2616](http://www.w3.org/Protocols/rfc2616/rfc2616-sec9.html#sec9)\n",
                item: [
                    {
                        _postman_id: "666d39d8-6aa6-0e29-2cf9-59cffce396bd",
                        name: "GET Request ",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "tests[\"Body contains headers\"] = responseBody.has(\"headers\");",
                                        "tests[\"Body contains args\"] = responseBody.has(\"args\");",
                                        "tests[\"Body contains url\"] = responseBody.has(\"url\");",
                                        "",
                                        "var data = JSON.parse(responseBody)",
                                        "",
                                        "tests[\"Args key contains argument passed as url parameter\"] = 'test' in data.args"
                                    ]
                                }
                            }
                        ],
                        request: {
                            url: {
                                raw: "https://echo.getpostman.com/get?test=123",
                                protocol: "https",
                                host: ["echo", "getpostman", "com"],
                                path: ["get"],
                                query: [{ key: "test", value: "123" }]
                            },
                            method: "GET",
                            header: [],
                            body: {
                                mode: "formdata",
                                formdata: [
                                    {
                                        key: "code",
                                        value: "xWnkliVQJURqB2x1",
                                        type: "text"
                                    },
                                    {
                                        key: "grant_type",
                                        value: "authorization_code",
                                        type: "text"
                                    },
                                    {
                                        key: "redirect_uri",
                                        value: "https://www.getpostman.com/oauth2/callback",
                                        type: "text"
                                    },
                                    {
                                        key: "client_id",
                                        value: "abc123",
                                        type: "text"
                                    },
                                    {
                                        key: "client_secret",
                                        value: "ssh-secret",
                                        type: "text"
                                    }
                                ]
                            },
                            description: "The HTTP `GET` request method is meant to retrieve data from a server. The data\nis identified by a unique URI (Uniform Resource Identifier). \n\nA `GET` request can pass parameters to the server using \"Query String \nParameters\". For example, in the following request,\n\n> http://example.com/hi/there?hand=wave\n\nThe parameter \"hand\" has the value \"wave\".\n\nThis endpoint echoes the HTTP headers, request parameters and the complete\nURI requested."
                        },
                        response: []
                    },
                    {
                        _postman_id: "1172f5ad-b9ed-8fb2-e447-a5e20f08327b",
                        name: "POST Request",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "var responseJSON;",
                                        "",
                                        "try { responseJSON = JSON.parse(responseBody); }",
                                        "catch (e) { }",
                                        "",
                                        "",
                                        "tests[\"response has data\"] = responseJSON && responseJSON.data && (responseJSON.data.length === 256);",
                                        "tests[\"content-type equals text/plain\"] = responseJSON && responseJSON.headers && (responseJSON.headers[\"content-type\"] === 'text/plain');"
                                    ]
                                }
                            }
                        ],
                        request: {
                            url: {
                                raw: "https://echo.getpostman.com/post",
                                protocol: "https",
                                host: ["echo", "getpostman", "com"],
                                path: ["post"]
                            },
                            method: "POST",
                            header: [
                                {
                                    key: "Content-Type",
                                    value: "text/plain"
                                }
                            ],
                            body: {
                                mode: "raw",
                                raw: "Duis posuere augue vel cursus pharetra. In luctus a ex nec pretium. Praesent neque quam, tincidunt nec leo eget, rutrum vehicula magna.\nMaecenas consequat elementum elit, id semper sem tristique et. Integer pulvinar enim quis consectetur interdum volutpat."
                            },
                            description: "The HTTP `POST` request method is meant to transfer data to a server \n(and elicit a response). What data is returned depends on the implementation\nof the server.\n\nA `POST` request can pass parameters to the server using \"Query String \nParameters\", as well as the Request Body. For example, in the following request,\n\n> POST /hi/there?hand=wave\n>\n> <request-body>\n\nThe parameter \"hand\" has the value \"wave\". The request body can be in multiple\nformats. These formats are defined by the MIME type of the request. The MIME \nType can be set using the ``Content-Type`` HTTP header. The most commonly used \nMIME types are:\n\n* `multipart/form-data`\n* `application/x-www-form-urlencoded`\n* `application/json`\n\nThis endpoint echoes the HTTP headers, request parameters, the contents of\nthe request body and the complete URI requested."
                        },
                        response: []
                    },
                    {
                        _postman_id: "80051609-e77d-aefa-a858-215d092f61a5",
                        name: "PUT Request",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "var data;",
                                        "",
                                        "try { responseJSON = JSON.parse(responseBody); }",
                                        "catch (e) {}",
                                        "",
                                        "tests[\"Body contains files\"] = responseBody.has(\"files\");",
                                        "tests[\"Body contains args\"] = responseBody.has(\"args\");",
                                        "tests[\"Body contains form\"] = responseBody.has(\"form\");",
                                        "tests[\"Body contains headers\"] = responseBody.has(\"headers\");",
                                        "tests[\"Body contains url\"] = responseBody.has(\"url\");",
                                        "",
                                        "",
                                        "",
                                        "tests[\"Data has been passed\"] = (responseJSON && responseJSON.data && responseJSON.data.length)"
                                    ]
                                }
                            }
                        ],
                        request: {
                            url: {
                                raw: "https://echo.getpostman.com/put",
                                protocol: "https",
                                host: ["echo", "getpostman", "com"],
                                path: ["put"]
                            },
                            method: "PUT",
                            header: [],
                            body: {
                                mode: "raw",
                                raw: "Etiam mi lacus, cursus vitae felis et, blandit pellentesque neque. Vestibulum eget nisi a tortor commodo dignissim.\nQuisque ipsum ligula, faucibus a felis a, commodo elementum nisl. Mauris vulputate sapien et tincidunt viverra. Donec vitae velit nec metus."
                            },
                            description: "The HTTP `PUT` request method is similar to HTTP `POST`. It too is meant to \ntransfer data to a server (and elicit a response). What data is returned depends on the implementation\nof the server.\n\nA `PUT` request can pass parameters to the server using \"Query String \nParameters\", as well as the Request Body. For example, in the following \nraw HTTP request,\n\n> PUT /hi/there?hand=wave\n>\n> <request-body>\n\n\n"
                        },
                        response: []
                    },
                    {
                        _postman_id: "97ed164c-d02b-bead-4abd-fe31cedb5a43",
                        name: "PATCH Request",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "tests[\"Body contains files\"] = responseBody.has(\"files\");",
                                        "tests[\"Body contains args\"] = responseBody.has(\"args\");",
                                        "tests[\"Body contains form\"] = responseBody.has(\"form\");",
                                        "tests[\"Body contains headers\"] = responseBody.has(\"headers\");",
                                        "tests[\"Body contains url\"] = responseBody.has(\"url\");",
                                        "",
                                        "var data = JSON.parse(responseBody)",
                                        "",
                                        "tests[\"form key has data passed in as form-data\"] = 'test' in data.form"
                                    ]
                                }
                            }
                        ],
                        request: {
                            url: {
                                raw: "https://echo.getpostman.com/patch",
                                protocol: "https",
                                host: ["echo", "getpostman", "com"],
                                path: ["patch"]
                            },
                            method: "PATCH",
                            header: [],
                            body: {
                                mode: "raw",
                                raw: "Curabitur auctor, elit nec pulvinar porttitor, ex augue condimentum enim, eget suscipit urna felis quis neque.\nSuspendisse sit amet luctus massa, nec venenatis mi. Suspendisse tincidunt massa at nibh efficitur fringilla. Nam quis congue mi. Etiam volutpat."
                            },
                            description: "The HTTP `PATCH` method is used to update resources on a server. The exact\nuse of `PATCH` requests depends on the server in question. There are a number\nof server implementations which handle `PATCH` differently. Technically, \n`PATCH` supports both Query String parameters and a Request Body.\n\nThis endpoint accepts an HTTP `PATCH` request and provides debug information\nsuch as the HTTP headers, Query String arguments, and the Request Body."
                        },
                        response: []
                    },
                    {
                        _postman_id: "b1dbc715-5df5-10a8-5b33-b5a130e52398",
                        name: "DELETE Request",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "tests[\"Body contains files\"] = responseBody.has(\"files\");",
                                        "tests[\"Body contains args\"] = responseBody.has(\"args\");",
                                        "tests[\"Body contains form\"] = responseBody.has(\"form\");",
                                        "tests[\"Body contains headers\"] = responseBody.has(\"headers\");",
                                        "tests[\"Body contains url\"] = responseBody.has(\"url\");",
                                        "",
                                        "var data = JSON.parse(responseBody)",
                                        "",
                                        "tests[\"form key has data passed in as form-data\"] = 'test' in data.form"
                                    ]
                                }
                            }
                        ],
                        request: {
                            url: {
                                raw: "https://echo.getpostman.com/delete",
                                protocol: "https",
                                host: ["echo", "getpostman", "com"],
                                path: ["delete"]
                            },
                            method: "DELETE",
                            header: [],
                            body: {
                                mode: "raw",
                                raw: "Donec fermentum, nisi sed cursus eleifend, nulla tortor ultricies tellus, ut vehicula orci arcu ut velit. In volutpat egestas dapibus. \nMorbi condimentum vestibulum sapien. Etiam dignissim diam quis eros lobortis gravida vel lobortis est. Etiam gravida sed."
                            },
                            description: "The HTTP `DELETE` method is used to delete resources on a server. The exact\nuse of `DELETE` requests depends on the server implementation. In general, \n`DELETE` requests support both, Query String parameters as well as a Request \nBody.\n\nThis endpoint accepts an HTTP `DELETE` request and provides debug information\nsuch as the HTTP headers, Query String arguments, and the Request Body."
                        },
                        response: []
                    }
                ]
            },
            {
                _postman_id: "4818889d-9709-a48a-90db-87442297021c",
                name: "Utilities",
                item: [
                    {
                        _postman_id: "0228eb86-77b8-6e3f-9641-539b14c0da12",
                        name: "Response Status Code",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "tests[\"Body contains status\"] = responseBody.has(\"status\");",
                                        "",
                                        "var data = JSON.parse(responseBody);",
                                        "",
                                        "tests[\"Status equals 200\"] = data.status === 200;",
                                        ""
                                    ]
                                }
                            }
                        ],
                        request: {
                            url: {
                                raw: "https://echo.getpostman.com/status/200",
                                protocol: "https",
                                host: ["echo", "getpostman", "com"],
                                path: ["status", "200"]
                            },
                            method: "GET",
                            header: [],
                            body: {
                                mode: "formdata",
                                formdata: [
                                    {
                                        key: "test",
                                        value: "123",
                                        type: "text"
                                    }
                                ]
                            },
                            description: "This endpoint allows one to instruct the server which status code to respond with.\n\nEvery response is accompanied by a status code. The status code provides a summary of the nature of response sent by the server. For example, a status code of `200` means everything is okay with the response and a code of `404` implies that the requested URL does not exist on server. \nA list of all valid HTTP status code can be found at the [List of Status Codes](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes) wikipedia article. When using Postman, the response status code is described for easy reference.\n\nNote that if an invalid status code is requested to be sent, the server returns a status code of `400 Bad Request`."
                        },
                        response: []
                    },
                    {
                        _postman_id: "4bc0e2a4-79c2-ee83-aae9-bc0bb6db6f85",
                        name: "Streamed Response",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "tests[\"response code is 200\"] = responseCode.code === 200;",
                                        "",
                                        ""
                                    ]
                                }
                            }
                        ],
                        request: {
                            url: {
                                raw: "https://echo.getpostman.com/stream/10",
                                protocol: "https",
                                host: ["echo", "getpostman", "com"],
                                path: ["stream", "10"]
                            },
                            method: "GET",
                            header: [],
                            body: {},
                            description: "This endpoint allows one to recieve streaming http response using [chunked transfer encoding](https://en.wikipedia.org/wiki/Chunked_transfer_encoding) of a configurable length.\n\nA streaming response does not wait for the entire response to be generated on server before flushing it out. This implies that for a fairly large response, parts of it can be streamed to the requestee as and when it is generated on server. The client can then take actions of processing this partially received data."
                        },
                        response: []
                    },
                    {
                        _postman_id: "195c6f9c-f328-74d5-d96c-1ae49da04f72",
                        name: "Delay Response",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "tests[\"response code is 200\"] = responseCode.code === 200;",
                                        "",
                                        "var data = JSON.parse(responseBody);",
                                        "",
                                        "tests[\"response body has key delay\"] = 'delay' in data;"
                                    ]
                                }
                            }
                        ],
                        request: {
                            url: {
                                raw: "https://echo.getpostman.com/delay/3",
                                protocol: "https",
                                host: ["echo", "getpostman", "com"],
                                path: ["delay", "3"]
                            },
                            method: "GET",
                            header: [],
                            body: {
                                mode: "formdata",
                                formdata: [
                                    {
                                        key: "test",
                                        value: "123",
                                        type: "text"
                                    }
                                ]
                            },
                            description: "Using this endpoint one can configure how long it takes for the server to come back with a response. Appending a number to the URL defines the time (in seconds) the server will wait before responding.\n\nNote that a maximum delay of 10 seconds is accepted by the server."
                        },
                        response: []
                    },
                    {
                        _postman_id: "b838d874-0889-3238-ebbd-30c4dca8442a",
                        name: "Get UTF8 Encoded Response",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "tests[\"response code is 200\"] = responseCode.code === 200;"
                                    ]
                                }
                            }
                        ],
                        request: {
                            url: {
                                raw: "https://echo.getpostman.com/encoding/utf8",
                                protocol: "https",
                                host: ["echo", "getpostman", "com"],
                                path: ["encoding", "utf8"]
                            },
                            method: "GET",
                            header: [],
                            body: {},
                            description: "If a response of an endpoint requires to send data beyond the basic English / ASCII character set, the `charset` parameter in the `Content-Type` response header defines the character encoding policy.\n\nThis endpoint returns an `UTF8` character encoded response body with text in various languages such as Greek, Latin, East Asian, etc. Postman can interpret the character encoding and use appropriate methods to display the character set in responses."
                        },
                        response: []
                    },
                    {
                        _postman_id: "bc9a7247-95f6-9d29-3622-ad30943ec7bf",
                        name: "GZip Compressed Response",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "tests[\"response code is 200\"] = responseCode.code === 200;",
                                        "",
                                        "try {",
                                        "    var data = JSON.parse(responseBody);",
                                        "    tests[\"Body contains gzipped\"] = responseBody.has(\"gzipped\");",
                                        "    tests[\"Body contains headers\"] = responseBody.has(\"headers\");",
                                        "    tests[\"Body contains method\"] = responseBody.has(\"method\");",
                                        "}",
                                        "catch(e) {",
                                        "    console.log('Cannot parse response,probably not a JSON');",
                                        "}"
                                    ]
                                }
                            }
                        ],
                        request: {
                            url: {
                                raw: "https://echo.getpostman.com/gzip",
                                protocol: "https",
                                host: ["echo", "getpostman", "com"],
                                path: ["gzip"]
                            },
                            method: "GET",
                            header: [],
                            body: {
                                mode: "formdata",
                                formdata: [
                                    {
                                        key: "code",
                                        value: "xWnkliVQJURqB2x1",
                                        type: "text"
                                    },
                                    {
                                        key: "grant_type",
                                        value: "authorization_code",
                                        type: "text"
                                    },
                                    {
                                        key: "redirect_uri",
                                        value: "https://www.getpostman.com/oauth2/callback",
                                        type: "text"
                                    },
                                    {
                                        key: "client_id",
                                        value: "abc123",
                                        type: "text"
                                    },
                                    {
                                        key: "client_secret",
                                        value: "ssh-secret",
                                        type: "text"
                                    }
                                ]
                            },
                            description: "This endpoint returns the response using [gzip compression algoritm](https://en.wikipedia.org/wiki/Gzip).\nThe uncompressed response is a JSON string containing the details of the request sent by the client. For this endpoint to work, one should request with `Accept-encoding` header containing `gzip` as part of its value. Postman supports gzip, deflate and SDCH decoding and automatically sends them as part of the request.\n\nHTTP Compression allows the server to send responses in a compressed format, which is uncompressed by the client before processing. This reduces network bandwidth consumption at the cost of increase in CPU usage.\nTo know more about this, refer the [HTTP Compression](https://en.wikipedia.org/wiki/HTTP_compression) wikipedia article."
                        },
                        response: []
                    },
                    {
                        _postman_id: "286b43d5-24ac-3c1e-ce42-c6204d627fa2",
                        name: "Deflate Compressed Response",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "tests[\"response code is 200\"] = responseCode.code === 200;",
                                        "",
                                        "try {",
                                        "    var data = JSON.parse(responseBody);",
                                        "    tests[\"Body contains deflated\"] = responseBody.has(\"deflated\");",
                                        "    tests[\"Body contains headers\"] = responseBody.has(\"headers\");",
                                        "    tests[\"Body contains method\"] = responseBody.has(\"method\");",
                                        "}",
                                        "catch(e) {",
                                        "    console.log('Cannot parse response,probably not a JSON');",
                                        "}",
                                        ""
                                    ]
                                }
                            }
                        ],
                        request: {
                            url: {
                                raw: "https://echo.getpostman.com/deflate",
                                protocol: "https",
                                host: ["echo", "getpostman", "com"],
                                path: ["deflate"]
                            },
                            method: "GET",
                            header: [],
                            body: {},
                            description: "This endpoint returns the response using [deflate compression algoritm](https://en.wikipedia.org/wiki/DEFLATE). \nThe uncompressed response is a JSON string containing the details of the request sent by the client. For this endpoint to work, one should request with `Accept-encoding` header containing `deflate` as part of its value. Postman supports gzip, deflate and SDCH decoding and automatically sends them as part of the request.\n\nHTTP Compression allows the server to send responses in a compressed format, which is uncompressed by the client before processing. This reduces network bandwidth consumption at the cost of increase in CPU usage.\nTo know more about this, refer the [HTTP Compression](https://en.wikipedia.org/wiki/HTTP_compression) wikipedia article."
                        },
                        response: []
                    }
                ]
            },
            {
                _postman_id: "48b8020c-1d8c-d97c-9472-d4376812b5e2",
                name: "[draft] Auth: OAuth2.0",
                description: "OAuth 2.0 allows a client to access a resource owners(user) protected resources without direct access to the users credentials.\nThe client obtains an access token which can be used to access the resource owners resources.\nOAuth 2.0 implements 4 different types of grant flows. Postman Echo has implemented the [Authorization Code](https://tools.ietf.org/html/rfc6749#section-1.3.1) grant flow.\n\nThe OAuth flow typically consists of three steps. \nIn the first step the client requests an authentication code from the server implementing the OAuth flow.\nThe authentication code along with other client credentials is then used to get the access token.\nThe access token can be used by the client to access users resources.\n",
                item: [
                    {
                        _postman_id: "e8deea86-6869-d33e-88ee-8295a5634e2f",
                        name: "OAuth2.0 Get Auth Code",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "tests[\"response code is 200\"] = responseCode.code === 200;"
                                    ]
                                }
                            }
                        ],
                        request: {
                            url: {
                                raw: "https://echo.getpostman.com/oauth2/authtoken?client_id=abc123&response_type=code",
                                protocol: "https",
                                host: ["echo", "getpostman", "com"],
                                path: ["oauth2", "authtoken"],
                                query: [{ key: "client_id", value: "abc123" }, { key: "response_type", value: "code" }]
                            },
                            method: "GET",
                            header: [],
                            body: {
                                mode: "formdata",
                                formdata: [
                                    {
                                        key: "code",
                                        value: "xWnkliVQJURqB2x1",
                                        type: "text"
                                    },
                                    {
                                        key: "grant_type",
                                        value: "authorization_code",
                                        type: "text"
                                    },
                                    {
                                        key: "redirect_uri",
                                        value: "https://www.getpostman.com/oauth2/callback",
                                        type: "text"
                                    },
                                    {
                                        key: "client_id",
                                        value: "abc123",
                                        type: "text"
                                    },
                                    {
                                        key: "client_secret",
                                        value: "ssh-secret",
                                        type: "text"
                                    }
                                ]
                            },
                            description: "In the first request for the OAuth 2.0 flow, the client tries to get an authentication code by passing in `client_id` and `response_type`.\n\n> client_id: `abc123`\n>\n> response_type: `code`\n\nThe authentication code is returned as a part of the redirect URI and can be used to get the access token."
                        },
                        response: [
                            {
                                id: "e1034070-2cd7-4a53-f07a-16e05426cf2b",
                                name: "401",
                                originalRequest: {
                                    header: [],
                                    body: {
                                        mode: "raw",
                                        raw: ""
                                    }
                                },
                                status: "Unauthorized",
                                code: 401,
                                _postman_previewlanguage: "javascript",
                                _postman_previewtype: "html",
                                header: [
                                    {
                                        name: "Access-Control-Allow-Credentials",
                                        key: "Access-Control-Allow-Credentials",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Headers",
                                        key: "Access-Control-Allow-Headers",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Methods",
                                        key: "Access-Control-Allow-Methods",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Origin",
                                        key: "Access-Control-Allow-Origin",
                                        value: ""
                                    },
                                    {
                                        name: "Connection",
                                        key: "Connection",
                                        value: "keep-alive"
                                    },
                                    {
                                        name: "Content-Length",
                                        key: "Content-Length",
                                        value: "2"
                                    },
                                    {
                                        name: "Content-Type",
                                        key: "Content-Type",
                                        value: "application/json; charset=utf-8"
                                    },
                                    {
                                        name: "Date",
                                        key: "Date",
                                        value: "Sat, 31 Oct 2015 06:43:23 GMT"
                                    },
                                    {
                                        name: "ETag",
                                        key: "ETag",
                                        value: "W/\"2-2745614147\""
                                    },
                                    {
                                        name: "Server",
                                        key: "Server",
                                        value: "nginx/1.6.2"
                                    },
                                    {
                                        name: "Vary",
                                        key: "Vary",
                                        value: "Accept-Encoding"
                                    },
                                    {
                                        name: "X-Powered-By",
                                        key: "X-Powered-By",
                                        value: "Sails <sailsjs.org>"
                                    }
                                ],
                                cookie: [],
                                responseTime: "291",
                                body: "{}"
                            }
                        ]
                    },
                    {
                        _postman_id: "2433b80a-d144-93b9-db7b-25195999050a",
                        name: "OAuth2.0 Get Access Token",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "tests[\"response code is 200\"] = responseCode.code === 200;",
                                        "var body = JSON.parse(responseBody);",
                                        "tests[\"body has access token\"] = \"access_token\" in body;",
                                        "tests[\"body has bearer type\"] = \"token_type\" in body;"
                                    ]
                                }
                            }
                        ],
                        request: {
                            url: {
                                raw: "https://echo.getpostman.com/oauth2/token",
                                protocol: "https",
                                host: [ "echo", "getpostman", "com" ],
                                path: [ "oauth2", "token" ]
                            },
                            method: "POST",
                            header: [],
                            body: {
                                mode: "formdata",
                                formdata: [
                                    {
                                        key: "code",
                                        value: "xWnkliVQJURqB2x1",
                                        type: "text"
                                    },
                                    {
                                        key: "grant_type",
                                        value: "authorization_code",
                                        type: "text"
                                    },
                                    {
                                        key: "redirect_uri",
                                        value: "https://www.getpostman.com/oauth2/callback",
                                        type: "text"
                                    },
                                    {
                                        key: "client_id",
                                        value: "abc123",
                                        type: "text"
                                    },
                                    {
                                        key: "client_secret",
                                        value: "ssh-secret",
                                        type: "text"
                                    }
                                ]
                            },
                            description: "This endpoint is used to get the `access_token`\n\nIt requires the following secret client credentials to be sent as part of the form body along with the `authentication code` obtained as part of the `redirect_uri` from the previous request.\n\n> code: xWnkliVQJURqB2x1\n>\n> grant_type: authorization_code\n>\n> redirect_uri: https://www.getpostman.com/oauth2/callback\n>\n> client_id: abc123\n>\n> client_secret: ssh-secret\n\nIf the correct credentials are not passed, the server returns with a `401 Unauthorized` response."
                        },
                        response: [
                            {
                                id: "3008931a-685b-aae8-fa0d-efe528bae4b3",
                                name: "401",
                                originalRequest: {
                                    url: {
                                        raw: "https://echo.getpostman.com/oauth2/token",
                                        protocol: "https",
                                        host: ["echo", "getpostman", "com"],
                                        path: ["oauth2", "token"]
                                    },
                                    method: "POST",
                                    header: [],
                                    body: {
                                        mode: "formdata",
                                        formdata: [
                                            {
                                                key: "code",
                                                value: "xWnkliVQJURqB2x1",
                                                type: "text"
                                            },
                                            {
                                                key: "grant_type",
                                                value: "authorization_code",
                                                type: "text"
                                            },
                                            {
                                                key: "redirect_uri",
                                                value: "https://www.getpostman.com/oauth2/callback",
                                                type: "text"
                                            },
                                            {
                                                key: "client_id",
                                                value: "abc123",
                                                type: "text"
                                            },
                                            {
                                                key: "client_secret",
                                                value: "ssh-secret",
                                                type: "text"
                                            }
                                        ]
                                    },
                                    description: "Tries to get the access token by passing the clientId , clientSecret,\nauthentication code, redirect URI and grant type.\n"
                                },
                                status: "Unauthorized",
                                code: 401,
                                _postman_previewlanguage: "javascript",
                                _postman_previewtype: "html",
                                header: [
                                    {
                                        name: "Access-Control-Allow-Credentials",
                                        key: "Access-Control-Allow-Credentials",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Headers",
                                        key: "Access-Control-Allow-Headers",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Methods",
                                        key: "Access-Control-Allow-Methods",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Origin",
                                        key: "Access-Control-Allow-Origin",
                                        value: ""
                                    },
                                    {
                                        name: "Connection",
                                        key: "Connection",
                                        value: "keep-alive"
                                    },
                                    {
                                        name: "Content-Length",
                                        key: "Content-Length",
                                        value: "2"
                                    },
                                    {
                                        name: "Content-Type",
                                        key: "Content-Type",
                                        value: "application/json; charset=utf-8"
                                    },
                                    {
                                        name: "Date",
                                        key: "Date",
                                        value: "Sat, 31 Oct 2015 06:43:44 GMT"
                                    },
                                    {
                                        name: "Server",
                                        key: "Server",
                                        value: "nginx/1.6.2"
                                    },
                                    {
                                        name: "Vary",
                                        key: "Vary",
                                        value: "X-HTTP-Method-Override, Accept-Encoding"
                                    },
                                    {
                                        name: "X-Powered-By",
                                        key: "X-Powered-By",
                                        value: "Sails <sailsjs.org>"
                                    }
                                ],
                                cookie: [],
                                responseTime: "287",
                                body: "{}"
                            },
                            {
                                id: "bb4e9632-233f-694f-b3fa-e54bb5e18322",
                                name: "200",
                                originalRequest: {
                                    url: {
                                        raw: "https://echo.getpostman.com/oauth2/token",
                                        protocol: "https",
                                        host: ["echo", "getpostman", "com"],
                                        path: ["oauth2", "token"]
                                    },
                                    method: "POST",
                                    header: [],
                                    body: {
                                        mode: "formdata",
                                        formdata: [
                                            {
                                                key: "code",
                                                value: "xWnkliVQJURqB2x1",
                                                type: "text"
                                            },
                                            {
                                                key: "grant_type",
                                                value: "authorization_code",
                                                type: "text"
                                            },
                                            {
                                                key: "redirect_uri",
                                                value: "https://www.getpostman.com/oauth2/callback",
                                                type: "text"
                                            },
                                            {
                                                key: "client_id",
                                                value: "abc123",
                                                type: "text"
                                            },
                                            {
                                                key: "client_secret",
                                                value: "ssh-secret",
                                                type: "text"
                                            }
                                        ]
                                    },
                                    description: "Tries to get the access token by passing the clientId , clientSecret,\nauthentication code, redirect URI and grant type.\n"
                                },
                                status: "OK",
                                code: 200,
                                _postman_previewlanguage: "javascript",
                                _postman_previewtype: "html",
                                header: [
                                    {
                                        name: "Access-Control-Allow-Credentials",
                                        key: "Access-Control-Allow-Credentials",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Headers",
                                        key: "Access-Control-Allow-Headers",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Methods",
                                        key: "Access-Control-Allow-Methods",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Origin",
                                        key: "Access-Control-Allow-Origin",
                                        value: ""
                                    },
                                    {
                                        name: "Connection",
                                        key: "Connection",
                                        value: "keep-alive"
                                    },
                                    {
                                        name: "Content-Encoding",
                                        key: "Content-Encoding",
                                        value: "gzip"
                                    },
                                    {
                                        name: "Content-Length",
                                        key: "Content-Length",
                                        value: "153"
                                    },
                                    {
                                        name: "Content-Type",
                                        key: "Content-Type",
                                        value: "application/json; charset=utf-8"
                                    },
                                    {
                                        name: "Date",
                                        key: "Date",
                                        value: "Sat, 31 Oct 2015 06:43:34 GMT"
                                    },
                                    {
                                        name: "Server",
                                        key: "Server",
                                        value: "nginx/1.6.2"
                                    },
                                    {
                                        name: "Vary",
                                        key: "Vary",
                                        value: "X-HTTP-Method-Override, Accept-Encoding"
                                    },
                                    {
                                        name: "X-Powered-By",
                                        key: "X-Powered-By",
                                        value: "Sails <sailsjs.org>"
                                    }
                                ],
                                cookie: [],
                                responseTime: "303",
                                body: "{\"access_token\":\"bj1wmMssgzidKHf6oGiMqy7TkGZkmOnrKlOMZcyEL4TEjqXG91TsAB68KG0wb1NwkNBDSXmhXoUg4JiVgXSFjkfwP1UDHSXe\",\"token_type\":\"Bearer\"}"
                            }
                        ]
                    },
                    {
                        _postman_id: "d72dc612-f45c-dfb7-aaeb-47ce79e34964",
                        name: "OAuth2.0 Get Resource",
                        event: [
                            {
                                listen: "test",
                                script: {
                                    type: "text/javascript",
                                    exec: [
                                        "tests[\"response code is 401\"] = (responseCode.code === 200);",
                                        "var body = JSON.parse(responseBody);",
                                        "tests[\"body has user_id\"] = \"user_id\" in body;",
                                        "tests[\"body has name\"] = \"name\" in body;"
                                    ]
                                }
                            }
                        ],
                        request: {
                            url: {
                                raw: "https://echo.getpostman.com/oauth2/user/info",
                                protocol: "https",
                                host: ["echo", "getpostman", "com"],
                                path: ["oauth2", "user", "info"]
                            },
                            method: "GET",
                            header: [
                                {
                                    key: "Authorization",
                                    value: "Bearer wn5htMraskxuODi0oHgLgz0UoXEvdSfoQcQSCxmQS0UPnqZS66GyRS10BA9lu0TjdKUPPIzhRqJl3PfFxZLVuagfM5LKYVCr"
                                }
                            ],
                            body: {
                                mode: "formdata",
                                formdata: [
                                    {
                                        key: "code",
                                        value: "xWnkliVQJURqB2x1",
                                        type: "text"
                                    },
                                    {
                                        key: "grant_type",
                                        value: "authorization_code",
                                        type: "text"
                                    },
                                    {
                                        key: "redirect_uri",
                                        value: "https://www.getpostman.com/oauth2/callback",
                                        type: "text"
                                    },
                                    {
                                        key: "client_id",
                                        value: "abc123",
                                        type: "text"
                                    },
                                    {
                                        key: "client_secret",
                                        value: "ssh-secret",
                                        type: "text"
                                    }
                                ]
                            },
                            description: "Once the bearer token has been obtained, it can be passed as `Authorization` header to access user's resources.\n\n> Authorization: Bearer vp7jx...\n\nIf the `Authorization` header is not passed the endpoint returns `401 Unauthorized`.\n"
                        },
                        response: [
                            {
                                id: "91b24532-54f2-7dcc-5c57-a452526bf531",
                                name: "401",
                                originalRequest: {
                                    header: [],
                                    body: {
                                        mode: "raw",
                                        raw: ""
                                    }
                                },
                                status: "Unauthorized",
                                code: 401,
                                _postman_previewlanguage: "html",
                                _postman_previewtype: "html",
                                header: [
                                    {
                                        name: "Access-Control-Allow-Credentials",
                                        key: "Access-Control-Allow-Credentials",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Headers",
                                        key: "Access-Control-Allow-Headers",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Methods",
                                        key: "Access-Control-Allow-Methods",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Origin",
                                        key: "Access-Control-Allow-Origin",
                                        value: ""
                                    },
                                    {
                                        name: "Connection",
                                        key: "Connection",
                                        value: "keep-alive"
                                    },
                                    {
                                        name: "Date",
                                        key: "Date",
                                        value: "Sat, 31 Oct 2015 06:44:16 GMT"
                                    },
                                    {
                                        name: "Server",
                                        key: "Server",
                                        value: "nginx/1.6.2"
                                    },
                                    {
                                        name: "WWW-Authenticate",
                                        key: "WWW-Authenticate",
                                        value: "Bearer realm=\"Users\""
                                    },
                                    {
                                        name: "X-Powered-By",
                                        key: "X-Powered-By",
                                        value: "Sails <sailsjs.org>"
                                    },
                                    {
                                        name: "transfer-encoding",
                                        key: "transfer-encoding",
                                        value: "chunked"
                                    }
                                ],
                                cookie: [],
                                responseTime: "303",
                                body: "Unauthorized"
                            },
                            {
                                id: "1b411278-3957-b390-4390-e10782a85c6c",
                                name: "200",
                                originalRequest: {
                                    header: [],
                                    body: {
                                        mode: "raw",
                                        raw: ""
                                    }
                                },
                                status: "OK",
                                code: 200,
                                _postman_previewlanguage: "javascript",
                                _postman_previewtype: "html",
                                header: [
                                    {
                                        name: "Access-Control-Allow-Credentials",
                                        key: "Access-Control-Allow-Credentials",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Headers",
                                        key: "Access-Control-Allow-Headers",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Methods",
                                        key: "Access-Control-Allow-Methods",
                                        value: ""
                                    },
                                    {
                                        name: "Access-Control-Allow-Origin",
                                        key: "Access-Control-Allow-Origin",
                                        value: ""
                                    },
                                    {
                                        name: "Connection",
                                        key: "Connection",
                                        value: "keep-alive"
                                    },
                                    {
                                        name: "Content-Encoding",
                                        key: "Content-Encoding",
                                        value: "gzip"
                                    },
                                    {
                                        name: "Content-Length",
                                        key: "Content-Length",
                                        value: "50"
                                    },
                                    {
                                        name: "Content-Type",
                                        key: "Content-Type",
                                        value: "application/json; charset=utf-8"
                                    },
                                    {
                                        name: "Date",
                                        key: "Date",
                                        value: "Sat, 31 Oct 2015 06:44:49 GMT"
                                    },
                                    {
                                        name: "Server",
                                        key: "Server",
                                        value: "nginx/1.6.2"
                                    },
                                    {
                                        name: "Vary",
                                        key: "Vary",
                                        value: "Accept-Encoding"
                                    },
                                    {
                                        name: "X-Powered-By",
                                        key: "X-Powered-By",
                                        value: "Sails <sailsjs.org>"
                                    }
                                ],
                                cookie: [],
                                responseTime: "323",
                                body: "{\"user_id\":1,\"name\":\"postman\"}"
                            }
                        ]
                    }
                ]
            }
        ]
    }
};
