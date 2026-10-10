/* eslint-disable quotes, max-len */

module.exports = {
    v1: {
        folder: "26e9263d-2c6b-c36c-5814-b4386573f9e4",
        id: "7987f16a-e666-97c3-dc8c-44501658698f",
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
        headers: "A:B\nC:D\n// E: F",
        method: "POST",
        pathVariables: {},
        url: "https://yo.postman.wtf/oauth2/token",
        preRequestScript: "",
        tests: "tests[\"response code is 200\"] = responseCode.code === 200;\nvar body = JSON.parse(responseBody);\ntests[\"body has access token\"] = \"access_token\" in body;\ntests[\"body has bearer type\"] = \"token_type\" in body;",
        currentHelper: "normal",
        helperAttributes: {},
        collectionId: "92de91bf-91be-fd1f-3ae7-412c86e70d54",
        responses: [
            {
                owner: "33232",
                lastUpdatedBy: "33232",
                lastRevision: 75106841,
                request: {
                    name: "Access Token",
                    description: "Tries to get the access token by passing the clientId , clientSecret,\nauthentication code, redirect URI and grant type.\n",
                    url: "https://yo.postman.wtf/oauth2/token",
                    pathVariables: {},
                    data: [
                        {
                            key: "code",
                            value: "xWnkliVQJURqB2x1",
                            type: "text",
                            enabled: true
                        },
                        {
                            key: "grant_type",
                            value: "authorization_code",
                            type: "text",
                            enabled: true
                        },
                        {
                            key: "redirect_uri",
                            value: "https://www.getpostman.com/oauth2/callback",
                            type: "text",
                            enabled: true
                        },
                        {
                            key: "client_id",
                            value: "abc123",
                            type: "text",
                            enabled: true
                        },
                        {
                            key: "client_secret",
                            value: "ssh-secret",
                            type: "text",
                            enabled: true
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
                id: "6530083f-548d-eae8-ae0f-cac696eea0b9",
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
                requestObject: "{\"name\":\"Access Token\",\"description\":\"Tries to get the access token by passing the clientId , clientSecret,\\nauthentication code, redirect URI and grant type.\\n\",\"url\":\"https://yo.postman.wtf/oauth2/token\",\"pathVariables\":{},\"data\":[{\"key\":\"code\",\"value\":\"xWnkliVQJURqB2x1\",\"type\":\"text\",\"enabled\":true},{\"key\":\"grant_type\",\"value\":\"authorization_code\",\"type\":\"text\",\"enabled\":true},{\"key\":\"redirect_uri\",\"value\":\"https://www.getpostman.com/oauth2/callback\",\"type\":\"text\",\"enabled\":true},{\"key\":\"client_id\",\"value\":\"abc123\",\"type\":\"text\",\"enabled\":true},{\"key\":\"client_secret\",\"value\":\"ssh-secret\",\"type\":\"text\",\"enabled\":true}],\"headers\":\"\",\"dataMode\":\"params\",\"method\":\"POST\",\"tests\":\"tests[\\\"response code is 200\\\"] = responseCode.code === 200;\\nvar body = JSON.parse(responseBody);\\ntests[\\\"body has access token\\\"] = \\\"access_token\\\" in body;\\ntests[\\\"body has bearer type\\\"] = \\\"token_type\\\" in body;\",\"isFromCollection\":true,\"write\":true,\"version\":2}",
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
                    url: "https://yo.postman.wtf/oauth2/token",
                    pathVariables: {},
                    data: [
                        {
                            key: "code",
                            value: "xWnkliVQJURqB2x1",
                            type: "text",
                            enabled: true
                        },
                        {
                            key: "grant_type",
                            value: "authorization_code",
                            type: "text",
                            enabled: true
                        },
                        {
                            key: "redirect_uri",
                            value: "https://www.getpostman.com/oauth2/callback",
                            type: "text",
                            enabled: true
                        },
                        {
                            key: "client_id",
                            value: "abc123",
                            type: "text",
                            enabled: true
                        },
                        {
                            key: "client_secret",
                            value: "ssh-secret",
                            type: "text",
                            enabled: true
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
                id: "cd8e9049-200f-895d-f7ec-f86da4e95807",
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
                text: "{\"access_token\":\"gz3crMtlfpbaAMy3tQxTep8ToMAcfTuvMdSZKhuKV2MLzdFC04GxLA46WS5or9TlbMEHSThxAkXu8UgOiVEIohllX5TAEQBr\",\"token_type\":\"Bearer\"}",
                language: "javascript",
                rawDataType: "text",
                state: {
                    size: "normal"
                },
                previewType: "html",
                searchResultScrolledTo: "-1",
                version: null,
                requestObject: "{\"name\":\"Access Token\",\"description\":\"Tries to get the access token by passing the clientId , clientSecret,\\nauthentication code, redirect URI and grant type.\\n\",\"url\":\"https://yo.postman.wtf/oauth2/token\",\"pathVariables\":{},\"data\":[{\"key\":\"code\",\"value\":\"xWnkliVQJURqB2x1\",\"type\":\"text\",\"enabled\":true},{\"key\":\"grant_type\",\"value\":\"authorization_code\",\"type\":\"text\",\"enabled\":true},{\"key\":\"redirect_uri\",\"value\":\"https://www.getpostman.com/oauth2/callback\",\"type\":\"text\",\"enabled\":true},{\"key\":\"client_id\",\"value\":\"abc123\",\"type\":\"text\",\"enabled\":true},{\"key\":\"client_secret\",\"value\":\"ssh-secret\",\"type\":\"text\",\"enabled\":true}],\"headers\":\"\",\"dataMode\":\"params\",\"method\":\"POST\",\"tests\":\"tests[\\\"response code is 200\\\"] = responseCode.code === 200;\\nvar body = JSON.parse(responseBody);\\ntests[\\\"body has access token\\\"] = \\\"access_token\\\" in body;\\ntests[\\\"body has bearer type\\\"] = \\\"token_type\\\" in body;\",\"isFromCollection\":true,\"write\":true,\"version\":2}",
                createdAt: "2015-11-02T13:11:08.000Z",
                updatedAt: "2015-11-02T18:05:45.000Z",
                write: true
            }
        ]
    },
    v2: {
        _postman_id: "7921f10a-c631-53b4-ca9e-98173806365e",
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
            url: 'https://yo.postman.wtf/oauth2/token',
            method: "POST",
            header: [
                { key: 'A', value: 'B' },
                { key: 'C', value: 'D' },
                { key: 'E', value: 'F', disabled: true }
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
            description: "This endpoint is used to get the `access_token`\n\nIt requires the following secret client credentials to be sent as part of the form body along with the `authentication code` obtained as part of the `redirect_uri` from the previous request.\n\n> code: xWnkliVQJURqB2x1\n>\n> grant_type: authorization_code\n>\n> redirect_uri: https://www.getpostman.com/oauth2/callback\n>\n> client_id: abc123\n>\n> client_secret: ssh-secret\n\nIf the correct credentials are not passed, the server returns with a `401 Unauthorized` response."
        },
        response: [
            {
                id: "3276699e-657f-ebd2-dc8e-bbd651aac8f5",
                name: "401",
                originalRequest: {
                    url: "https://yo.postman.wtf/oauth2/token",
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
                id: "ff4f0497-201f-332f-e4fd-b52ee8c05099",
                name: "200",
                originalRequest: {
                    url: "https://yo.postman.wtf/oauth2/token",
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
                body: "{\"access_token\":\"ee5cqBzdfjrcHWs9qEkOfj5GyXLixVdbGhSWJnlZJ2BUbhAI23TyEJ26WH2xu1XqjZQLFKqwPdYg6KqNmIIWivlcU9HCSWMj\",\"token_type\":\"Bearer\"}"
            }
        ]
    },
    v21: {
        _postman_id: "3579b88c-c243-79a8-bc6b-27640747547c",
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
                raw: 'https://yo.postman.wtf/oauth2/token',
                protocol: 'https',
                host: ['yo', 'postman', 'wtf'],
                path: ['oauth2', 'token']
            },
            method: "POST",
            header: [
                { key: 'A', value: 'B' },
                { key: 'C', value: 'D' },
                { key: 'E', value: 'F', disabled: true }
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
            description: "This endpoint is used to get the `access_token`\n\nIt requires the following secret client credentials to be sent as part of the form body along with the `authentication code` obtained as part of the `redirect_uri` from the previous request.\n\n> code: xWnkliVQJURqB2x1\n>\n> grant_type: authorization_code\n>\n> redirect_uri: https://www.getpostman.com/oauth2/callback\n>\n> client_id: abc123\n>\n> client_secret: ssh-secret\n\nIf the correct credentials are not passed, the server returns with a `401 Unauthorized` response."
        },
        response: [
            {
                id: "0659407b-426f-caa1-bc9a-bdf190abc5c0",
                name: "401",
                originalRequest: {
                    url: {
                        raw: "https://yo.postman.wtf/oauth2/token",
                        protocol: 'https',
                        host: ['yo', 'postman', 'wtf'],
                        path: ['oauth2', 'token']
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
                id: "af4f8470-873d-140a-b3de-e47ba0c68814",
                name: "200",
                originalRequest: {
                    url: {
                        raw: "https://yo.postman.wtf/oauth2/token",
                        protocol: 'https',
                        host: ['yo', 'postman', 'wtf'],
                        path: ['oauth2', 'token']
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
                body: "{\"access_token\":\"cq6pyNwzfbhfFBe6kZbLjn2XyEWyiYjnJhPUMqzJS6JVywDG79LjUL96VW7cj6PngIAZDWtpYrMi6TqXhQRGpgspF7VRMEXc\",\"token_type\":\"Bearer\"}"
            }
        ]
    }
};
