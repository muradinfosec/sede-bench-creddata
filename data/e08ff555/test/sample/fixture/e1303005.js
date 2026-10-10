module.exports = {
    v1: {
        id: 'e323a78f-c333-0a18-120e-e7c2fe03fc84',
        name: 'AuthTest',
        description: 'Sample auth tests',
        order: [
            '45f6bdc1-aa7c-1bf5-4aab-a976f1960a37',
            '868fb2a7-f9c2-3e44-9b69-7d8350b7496a',
            '951fc3e8-c6b6-5c19-9f69-4e7499b3127g'
        ],
        folders: [],
        folders_order: [],
        requests: [
            {
                id: '18e2eca9-fc6c-8ba2-9eec-b440e9173e78',
                headers: '',
                url: 'httpbin.org/post',
                preRequestScript: '',
                queryParams: [],
                rawModeData: '',
                pathVariableData: [],
                method: 'POST',
                data: [],
                dataMode: 'raw',
                headerData: [],
                events: [{
                    listen: 'test',
                    script: {
                        type: 'text/javascript',
                        exec: [
                            'var data = JSON.parse(responseBody);',
                            'var oldNonce = environment.oldNonce;',
                            'var newNonce = data.form.oauth_nonce;',
                            'tests["NewNonce"]=(oldNonce != newNonce);',
                            'postman.setEnvironmentVariable("oldNonce", newNonce);',
                            // eslint-disable-next-line max-len
                            'tests["HasNonce"]=data.form.oauth_nonce;console.log("oldNonce: "+oldNonce+", newNonce="+newNonce);console.log("Signature: "+data.form.oauth_signature);'
                        ]
                    }
                }],
                // eslint-disable-next-line max-len
                tests: 'var data = JSON.parse(responseBody);\nvar oldNonce = environment.oldNonce;\nvar newNonce = data.form.oauth_nonce;\ntests["NewNonce"]=(oldNonce != newNonce);\npostman.setEnvironmentVariable("oldNonce", newNonce);\ntests["HasNonce"]=data.form.oauth_nonce;console.log("oldNonce: "+oldNonce+", newNonce="+newNonce);console.log("Signature: "+data.form.oauth_signature);',
                auth: {
                    type: 'oauth2',
                    oauth2: [
                        { key: 'accessToken', value: 'secretToken', type: 'string' },
                        { key: 'addTokenTo', value: 'header', type: 'string' },
                        { key: 'callBackUrl', value: 'https://foo.com/cb', type: 'string' },
                        { key: 'authUrl', value: 'https://foo.com/au', type: 'string' },
                        { key: 'accessTokenUrl', value: 'https://foo.com/at', type: 'string' },
                        { key: 'clientId', value: 'uniqueClientIdentifier', type: 'string' },
                        { key: 'clientSecret', value: 'secretClientValue', type: 'string' },
                        { key: 'clientAuth', value: 'body', type: 'string' },
                        { key: 'grantType', value: 'implicit', type: 'string' },
                        { key: 'scope', value: 'all', type: 'string' },
                        { key: 'username', value: 'postman', type: 'string' },
                        { key: 'password', value: 'randomSecretString', type: 'string' },
                        { key: 'tokenType', value: 'bearer', type: 'string' },
                        { key: 'redirectUri', value: 'https://foo.com/rd', type: 'string' },
                        { key: 'refreshToken', value: 'refreshToken', type: 'string' }
                    ]
                },
                currentHelper: 'oAuth2',
                helperAttributes: {
                    id: 'oAuth2',
                    accessToken: 'secretToken',
                    addTokenTo: 'header',
                    callBackUrl: 'https://foo.com/cb',
                    authUrl: 'https://foo.com/au',
                    accessTokenUrl: 'https://foo.com/at',
                    clientId: 'uniqueClientIdentifier',
                    clientSecret: 'secretClientValue',
                    clientAuth: 'body',
                    grantType: 'implicit',
                    scope: 'all',
                    username: 'postman',
                    password: 'randomSecretString',
                    tokenType: 'bearer',
                    redirectUri: 'https://foo.com/rd',
                    refreshToken: 'refreshToken'
                },
                name: 'OAuth2',
                collectionId: 'c458e40d-f821-8f18-425e-c6e4da46db34',
                responses: []
            },
            {
                id: '304da6b6-d3f2-9d99-1b38-0b4762c9218b',
                headers: 'Authorization: Bearer flcpknfymdtyy8n3',
                url: 'http://echo.getpostman.com/auth/bearer',
                preRequestScript: '',
                method: 'GET',
                data: [],
                queryParams: [],
                rawModeData: '',
                headerData: [{
                    key: 'Authorization',
                    value: 'Bearer wkjehbxoqnunc2k3'
                }],
                dataMode: 'raw',
                pathVariableData: [],
                events: [{
                    listen: 'test',
                    script: {
                        type: 'text/javascript',
                        // eslint-disable-next-line max-len
                        exec: ['var response = JSON.parse(responseBody); tests["Bearer auth should pass"] = response.status === "pass";']
                    }
                }],
                // eslint-disable-next-line max-len
                tests: 'var response = JSON.parse(responseBody); tests["Bearer auth should pass"] = response.status === "pass";',
                auth: {
                    type: 'bearer',
                    bearer: [{ key: 'token', value: 'wkjehbxoqnunc2k3', type: 'string' }]
                },
                currentHelper: 'bearerAuth',
                helperAttributes: {
                    id: 'bearer',
                    token: 'dsirxvfyzxstq2u0'
                },
                name: 'test bearer auth success',
                collectionId: 'c430f20e-e678-7e29-058e-b1c1dd07aa64',
                responses: []
            },
            {
                id: '951fc3e8-c6b6-5c19-9f69-4e7499b3127g',
                // eslint-disable-next-line max-len
                headers: 'Authorization: NTLM TlRMTVNTUAADAAAAGAAYAFYAAAAYABgAbgAAAAAAAABIAAAADgAOAEgAAAAAAAAAVgAAAAAAAACGAAAABYKIogUBKAoAAAAPcABvAHMAdABtAGEAbgCDJfCgEQC+3wAAAAAAAAAAAAAAAAAAAAB6O2T1blvWpb/pqHtSdqcZ/A34nPBZe20=',
                url: 'http://echo.getpostman.com/auth/ntlm',
                preRequestScript: '',
                method: 'GET',
                data: [],
                queryParams: [],
                rawModeData: '',
                headerData: [{
                    key: 'Authorization',
                    // eslint-disable-next-line max-len
                    value: 'NTLM TlRMTVNTUAADAAAAGAAYAFYAAAAYABgAbgAAAAAAAABIAAAADgAOAEgAAAAAAAAAVgAAAAAAAACGAAAABYKIogUBKAoAAAAPcABvAHMAdABtAGEAbgCDJfCgEQC+3wAAAAAAAAAAAAAAAAAAAAB6O2T1blvWpb/pqHtSdqcZ/A34nPBZe20='
                }],
                dataMode: 'raw',
                pathVariableData: [],
                events: [{
                    listen: 'test',
                    script: {
                        type: 'text/javascript',
                        // eslint-disable-next-line max-len
                        exec: ['var response = JSON.parse(responseBody); tests["NTLM auth should pass"] = response.status === "pass";']
                    }
                }],
                // eslint-disable-next-line max-len
                tests: 'var response = JSON.parse(responseBody); tests["NTLM auth should pass"] = response.status === "pass";',
                auth: {
                    type: 'ntlm',
                    ntlm: [
                        { key: 'username', value: 'foo', type: 'string' },
                        { key: 'password', value: 'password', type: 'string' },
                        { key: 'domain', value: 'domain', type: 'string' },
                        { key: 'workstation', value: 'workstation', type: 'string' },
                        { key: 'disableRetryRequest', value: false, type: 'boolean' }
                    ]
                },
                currentHelper: 'ntlmAuth',
                helperAttributes: {
                    id: 'ntlm',
                    username: 'foo',
                    password: 'password',
                    domain: 'domain',
                    workstation: 'workstation',
                    disableRetryRequest: false
                },
                name: 'test ntlm auth success',
                collectionId: 'c673b26c-f706-9e20-427b-d9c9be77ec47',
                responses: []
            }
        ]
    },
    v2: {
        info: {
            name: 'AuthTest',
            _postman_id: 'f451e76d-b869-6c50-047c-e7c4bb61fc18',
            description: 'Sample auth tests',
            schema: 'https://schema.getpostman.com/json/collection/v2.0.0/collection.json'
        },
        item: [
            {
                _postman_id: '86a9efe1-fe0f-8fd8-1cee-e320f0502b34',
                name: 'OAuth2',
                event: [
                    {
                        listen: 'test',
                        script: {
                            type: 'text/javascript',
                            exec: [
                                'var data = JSON.parse(responseBody);',
                                'var oldNonce = environment.oldNonce;',
                                'var newNonce = data.form.oauth_nonce;',
                                'tests["NewNonce"]=(oldNonce != newNonce);',
                                'postman.setEnvironmentVariable("oldNonce", newNonce);',
                                // eslint-disable-next-line max-len
                                'tests["HasNonce"]=data.form.oauth_nonce;console.log("oldNonce: "+oldNonce+", newNonce="+newNonce);console.log("Signature: "+data.form.oauth_signature);'
                            ]
                        }
                    }
                ],
                request: {
                    auth: {
                        type: 'oauth2',
                        oauth2: {
                            accessToken: 'secretToken',
                            addTokenTo: 'header',
                            callBackUrl: 'https://foo.com/cb',
                            authUrl: 'https://foo.com/au',
                            accessTokenUrl: 'https://foo.com/at',
                            clientId: 'uniqueClientIdentifier',
                            clientSecret: 'secretClientValue',
                            clientAuth: 'body',
                            grantType: 'implicit',
                            scope: 'all',
                            username: 'postman',
                            password: 'randomSecretString',
                            tokenType: 'bearer',
                            redirectUri: 'https://foo.com/rd',
                            refreshToken: 'refreshToken'
                        }
                    },
                    url: 'httpbin.org/post',
                    method: 'POST',
                    header: [],
                    body: {
                        mode: 'raw',
                        raw: ''
                    }
                },
                response: []
            },
            {
                _postman_id: '440af1b5-a4e6-9e00-1a41-9d0505d0491f',
                name: 'test bearer auth success',
                event: [
                    {
                        listen: 'test',
                        script: {
                            type: 'text/javascript',
                            exec: [
                                // eslint-disable-next-line max-len
                                'var response = JSON.parse(responseBody); tests["Bearer auth should pass"] = response.status === "pass";'
                            ]
                        }
                    }
                ],
                request: {
                    auth: {
                        type: 'bearer',
                        bearer: {
                            token: 'nbovrofuhnnbe4t8'
                        }
                    },
                    url: 'http://echo.getpostman.com/auth/bearer',
                    method: 'GET',
                    header: [
                        {
                            key: 'Authorization',
                            value: 'Bearer wkjehbxoqnunc2k3'
                        }
                    ],
                    body: {
                        mode: 'raw',
                        raw: ''
                    }
                },
                response: []
            },
            {
                _postman_id: '951fc3e8-c6b6-5c19-9f69-4e7499b3127g',
                name: 'test ntlm auth success',
                event: [
                    {
                        listen: 'test',
                        script: {
                            type: 'text/javascript',
                            exec: [
                                // eslint-disable-next-line max-len
                                'var response = JSON.parse(responseBody); tests["NTLM auth should pass"] = response.status === "pass";'
                            ]
                        }
                    }
                ],
                request: {
                    auth: {
                        type: 'ntlm',
                        ntlm: {
                            username: 'foo',
                            password: 'password',
                            domain: 'domain',
                            workstation: 'workstation',
                            disableRetryRequest: false
                        }
                    },
                    url: 'http://echo.getpostman.com/auth/ntlm',
                    method: 'GET',
                    header: [
                        {
                            key: 'Authorization',
                            // eslint-disable-next-line max-len
                            value: 'NTLM TlRMTVNTUAADAAAAGAAYAFYAAAAYABgAbgAAAAAAAABIAAAADgAOAEgAAAAAAAAAVgAAAAAAAACGAAAABYKIogUBKAoAAAAPcABvAHMAdABtAGEAbgCDJfCgEQC+3wAAAAAAAAAAAAAAAAAAAAB6O2T1blvWpb/pqHtSdqcZ/A34nPBZe20='
                        }
                    ],
                    body: {
                        mode: 'raw',
                        raw: ''
                    }
                },
                response: []
            }
        ]
    },
    v21: {
        info: {
            name: 'AuthTest',
            _postman_id: 'a082d69d-d097-5d04-696b-d2e1fd97ed05',
            description: 'Sample auth tests',
            schema: 'https://schema.getpostman.com/json/collection/v2.1.0/collection.json'
        },
        item: [
            {
                _postman_id: '85a8aad5-bf5a-3fc0-3bee-e339f0786c01',
                name: 'OAuth2',
                event: [
                    {
                        listen: 'test',
                        script: {
                            type: 'text/javascript',
                            exec: [
                                'var data = JSON.parse(responseBody);',
                                'var oldNonce = environment.oldNonce;',
                                'var newNonce = data.form.oauth_nonce;',
                                'tests["NewNonce"]=(oldNonce != newNonce);',
                                'postman.setEnvironmentVariable("oldNonce", newNonce);',
                                // eslint-disable-next-line max-len
                                'tests["HasNonce"]=data.form.oauth_nonce;console.log("oldNonce: "+oldNonce+", newNonce="+newNonce);console.log("Signature: "+data.form.oauth_signature);'
                            ]
                        }
                    }
                ],
                request: {
                    auth: {
                        type: 'oauth2',
                        oauth2: [
                            {
                                key: 'accessToken',
                                value: 'secretToken',
                                type: 'string'
                            },
                            {
                                key: 'addTokenTo',
                                value: 'header',
                                type: 'string'
                            },
                            {
                                key: 'callBackUrl',
                                value: 'https://foo.com/cb',
                                type: 'string'
                            },
                            {
                                key: 'authUrl',
                                value: 'https://foo.com/au',
                                type: 'string'
                            },
                            {
                                key: 'accessTokenUrl',
                                value: 'https://foo.com/at',
                                type: 'string'
                            },
                            {
                                key: 'clientId',
                                value: 'uniqueClientIdentifier',
                                type: 'string'
                            },
                            {
                                key: 'clientSecret',
                                value: 'secretClientValue',
                                type: 'string'
                            },
                            {
                                key: 'clientAuth',
                                value: 'body',
                                type: 'string'
                            },
                            {
                                key: 'grantType',
                                value: 'implicit',
                                type: 'string'
                            },
                            {
                                key: 'scope',
                                value: 'all',
                                type: 'string'
                            },
                            {
                                key: 'username',
                                value: 'postman',
                                type: 'string'
                            },
                            {
                                key: 'password',
                                value: 'randomSecretString',
                                type: 'string'
                            },
                            {
                                key: 'tokenType',
                                value: 'bearer',
                                type: 'string'
                            },
                            {
                                key: 'redirectUri',
                                value: 'https://foo.com/rd',
                                type: 'string'
                            },
                            {
                                key: 'refreshToken',
                                value: 'refreshToken',
                                type: 'string'
                            }
                        ]
                    },
                    url: {
                        raw: 'httpbin.org/post',
                        host: ['httpbin', 'org'],
                        path: ['post']
                    },
                    method: 'POST',
                    header: [],
                    body: {
                        mode: 'raw',
                        raw: ''
                    }
                },
                response: []
            },
            {
                _postman_id: '843eb9a6-a3e1-3b09-1f40-5d7571e2805d',
                name: 'test bearer auth success',
                event: [
                    {
                        listen: 'test',
                        script: {
                            type: 'text/javascript',
                            exec: [
                                // eslint-disable-next-line max-len
                                'var response = JSON.parse(responseBody); tests["Bearer auth should pass"] = response.status === "pass";'
                            ]
                        }
                    }
                ],
                request: {
                    auth: {
                        type: 'bearer',
                        bearer: [
                            {
                                key: 'token',
                                value: 'wkjehbxoqnunc2k3',
                                type: 'string'
                            }
                        ]
                    },
                    url: {
                        raw: 'http://echo.getpostman.com/auth/bearer',
                        protocol: 'http',
                        host: ['echo', 'getpostman', 'com'],
                        path: ['auth', 'bearer']
                    },
                    method: 'GET',
                    header: [
                        {
                            key: 'Authorization',
                            value: 'Bearer wkjehbxoqnunc2k3'
                        }
                    ],
                    body: {
                        mode: 'raw',
                        raw: ''
                    }
                },
                response: []
            },
            {
                _postman_id: '951fc3e8-c6b6-5c19-9f69-4e7499b3127g',
                name: 'test ntlm auth success',
                event: [
                    {
                        listen: 'test',
                        script: {
                            type: 'text/javascript',
                            exec: [
                                // eslint-disable-next-line max-len
                                'var response = JSON.parse(responseBody); tests["NTLM auth should pass"] = response.status === "pass";'
                            ]
                        }
                    }
                ],
                request: {
                    auth: {
                        type: 'ntlm',
                        ntlm: [
                            {
                                key: 'username',
                                value: 'foo',
                                type: 'string'
                            },
                            {
                                key: 'password',
                                value: 'password',
                                type: 'string'
                            },
                            {
                                key: 'domain',
                                value: 'domain',
                                type: 'string'
                            },
                            {
                                key: 'workstation',
                                value: 'workstation',
                                type: 'string'
                            },
                            {
                                key: 'disableRetryRequest',
                                value: false,
                                type: 'boolean'
                            }
                        ]
                    },
                    url: {
                        raw: 'http://echo.getpostman.com/auth/ntlm',
                        protocol: 'http',
                        host: ['echo', 'getpostman', 'com'],
                        path: ['auth', 'ntlm']
                    },
                    method: 'GET',
                    header: [
                        {
                            key: 'Authorization',
                            // eslint-disable-next-line max-len
                            value: 'NTLM TlRMTVNTUAADAAAAGAAYAFYAAAAYABgAbgAAAAAAAABIAAAADgAOAEgAAAAAAAAAVgAAAAAAAACGAAAABYKIogUBKAoAAAAPcABvAHMAdABtAGEAbgCDJfCgEQC+3wAAAAAAAAAAAAAAAAAAAAB6O2T1blvWpb/pqHtSdqcZ/A34nPBZe20='
                        }
                    ],
                    body: {
                        mode: 'raw',
                        raw: ''
                    }
                },
                response: []
            }
        ]
    }
};
