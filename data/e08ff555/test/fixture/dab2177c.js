module.exports = {
    v1: {
        id: '179a5ee2-07d5-fb28-51d9-452c42f763b8',
        name: 'nested-entities',
        description: 'A simple collection to elucidate nested auth helpers, variables, and scripts',
        order: [
            'c965e827-5ef1-e05a-d33f-c468b2591cb1'
        ],
        variables: [{
            id: 'collection',
            key: 'collection',
            value: 'variable',
            type: 'string',
            description: 'This is a collection level variable'
        }],
        folders: [
            {
                name: 'F1',
                order: [
                    'c331446f-cb81-23f4-2a9a-5e2a1b7cca35'
                ],
                variables: [{
                    id: 'folder',
                    key: 'folder',
                    value: 'variable',
                    type: 'string',
                    description: 'This is a folder level variable'
                }],
                events: [{
                    listen: 'test',
                    script: {
                        type: 'text/javascript',
                        exec: ['console.log("Folder level test script");']
                    }
                }, {
                    listen: 'prerequest',
                    script: {
                        type: 'text/javascript',
                        exec: ['console.log("Folder level pre request script");']
                    }
                }],
                auth: {
                    type: 'digest',
                    digest: [
                        {
                            key: 'algorithm',
                            value: 'MD5',
                            type: 'string'
                        },
                        {
                            key: 'username',
                            value: 'postman',
                            type: 'string'
                        },
                        {
                            key: 'realm',
                            value: '{{echo_digest_realm}}',
                            type: 'string'
                        },
                        {
                            key: 'password',
                            value: 'password',
                            type: 'string'
                        },
                        {
                            key: 'nonce',
                            value: '{{echo_digest_nonce}}',
                            type: 'string'
                        },
                        {
                            key: 'nonceCount',
                            value: '',
                            type: 'string'
                        },
                        {
                            key: 'clientNonce',
                            value: '',
                            type: 'string'
                        },
                        {
                            key: 'opaque',
                            value: '',
                            type: 'string'
                        },
                        {
                            key: 'qop',
                            value: '',
                            type: 'string'
                        }
                    ]
                },
                folders_order: [],
                id: '5ce892b1-8259-7f27-9369-9482927e4566'
            }
        ],
        folders_order: [
            '3ed392b3-6247-4f72-3334-8677876f4727'
        ],
        events: [
            {
                listen: 'test',
                script: {
                    type: 'text/javascript',
                    exec: [
                        'console.log("Collection level test script");'
                    ]
                }
            },
            {
                listen: 'prerequest',
                script: {
                    type: 'text/javascript',
                    exec: [
                        'console.log("Collection level pre request script");'
                    ]
                }
            }
        ],
        auth: {
            type: 'awsv4',
            awsv4: [
                {
                    key: 'accessKey',
                    value: '{{id}}',
                    type: 'string'
                },
                {
                    key: 'id',
                    value: 'awsSigV4',
                    type: 'string'
                },
                {
                    key: 'region',
                    value: 'us-east-1',
                    type: 'string'
                },
                {
                    key: 'secretKey',
                    value: '{{key}}',
                    type: 'string'
                },
                {
                    key: 'service',
                    value: 'iam',
                    type: 'string'
                },
                {
                    key: 'sessionToken',
                    value: '',
                    type: 'string'
                }
            ]
        },
        requests: [
            {
                id: 'c136916e-fb93-50e5-4e8e-4a4c9e9dbb17',
                variables: [{
                    id: 'request',
                    key: 'request',
                    value: 'variable',
                    type: 'string',
                    description: 'This is a request level variable'
                }],
                headers: '',
                headerData: [],
                url: 'https://postman-echo.com/get',
                queryParams: [],
                events: [
                    {
                        listen: 'test',
                        script: {
                            type: 'text/javascript',
                            exec: [
                                'console.log("Request level test script");'
                            ]
                        }
                    },
                    {
                        listen: 'prerequest',
                        script: {
                            type: 'text/javascript',
                            exec: [
                                'console.log("Request level pre request script");'
                            ]
                        }
                    }
                ],
                preRequestScript: 'console.log("Request level pre request script");',
                pathVariableData: [],
                rawModeData: '',
                method: 'GET',
                data: [],
                dataMode: 'raw',
                tests: 'console.log("Request level test script");',
                auth: {
                    type: 'basic',
                    basic: [
                        {
                            key: 'username',
                            value: '{{username}}',
                            type: 'string'
                        },
                        {
                            key: 'password',
                            value: '{{token}}',
                            type: 'string'
                        }
                    ]
                },
                currentHelper: 'basicAuth',
                helperAttributes: {
                    id: 'basic',
                    username: '{{username}}',
                    password: '{{token}}'
                },
                name: 'Basic',
                collectionId: '058b2fa8-14a0-ca98-97a6-425f27f510f8',
                responses: []
            },
            {
                id: 'c244d127-4ae6-f02d-f40c-a502e0912ae7',
                variables: [{
                    id: 'request',
                    key: 'request',
                    value: 'variable',
                    type: 'string',
                    description: 'This is a request level variable'
                }],
                headers: '',
                headerData: [],
                url: 'https://postman-echo.com/get',
                queryParams: [],
                events: [
                    {
                        listen: 'test',
                        script: {
                            type: 'text/javascript',
                            exec: [
                                'console.log("Request level test script");'
                            ]
                        }
                    },
                    {
                        listen: 'prerequest',
                        script: {
                            type: 'text/javascript',
                            exec: [
                                'console.log("Request level pre request script");'
                            ]
                        }
                    }
                ],
                preRequestScript: 'console.log("Request level pre request script");',
                pathVariableData: [],
                method: 'GET',
                data: [],
                dataMode: 'raw',
                rawModeData: '',
                tests: 'console.log("Request level test script");',
                auth: {
                    type: 'oauth1',
                    oauth1: [
                        {
                            key: 'consumerKey',
                            value: 'RKCGzna7bv9YD57c',
                            type: 'string'
                        },
                        {
                            key: 'consumerSecret',
                            value: 'D+EdQ-gs$-%@2Nu7',
                            type: 'string'
                        },
                        {
                            key: 'token',
                            value: '',
                            type: 'string'
                        },
                        {
                            key: 'tokenSecret',
                            value: '',
                            type: 'string'
                        },
                        {
                            key: 'signatureMethod',
                            value: 'HMAC-SHA1',
                            type: 'string'
                        },
                        {
                            key: 'timestamp',
                            value: '1500452534',
                            type: 'string'
                        },
                        {
                            key: 'nonce',
                            value: 'S0kXloMHurS',
                            type: 'string'
                        },
                        {
                            key: 'version',
                            value: '1.0',
                            type: 'string'
                        },
                        {
                            key: 'realm',
                            value: '',
                            type: 'string'
                        }
                    ]
                },
                currentHelper: 'oAuth1',
                helperAttributes: {
                    id: 'oAuth1',
                    consumerKey: 'JIZPmkq9jj2XN34i',
                    consumerSecret: 'U+AeN-ju$-%@2Jm6',
                    token: '',
                    tokenSecret: '',
                    signatureMethod: 'HMAC-SHA1',
                    timestamp: '1500452534',
                    nonce: 'F7uWhaQBavN',
                    version: '1.0',
                    realm: ''
                },
                name: 'OAuth1',
                collectionId: '247c9af0-72e7-ac93-01d5-808a98c377b2',
                responses: []
            }
        ]
    },
    v2: {
        info: {
            name: 'nested-entities',
            _postman_id: '780b3ab4-78a2-fa81-28b0-228c45e761b7',
            description: 'A simple collection to elucidate nested auth helpers, variables, and scripts',
            schema: 'https://schema.getpostman.com/json/collection/v2.0.0/collection.json'
        },
        variable: [{
            id: 'collection',
            key: 'collection',
            value: 'variable',
            type: 'string',
            description: 'This is a collection level variable'
        }],
        auth: {
            type: 'awsv4',
            awsv4: {
                accessKey: '{{id}}',
                id: 'awsSigV4',
                region: 'us-east-1',
                secretKey: '{{key}}',
                service: 'iam',
                sessionToken: ''
            }
        },
        event: [
            {
                listen: 'test',
                script: {
                    type: 'text/javascript',
                    exec: [
                        'console.log("Collection level test script");'
                    ]
                }
            },
            {
                listen: 'prerequest',
                script: {
                    type: 'text/javascript',
                    exec: [
                        'console.log("Collection level pre request script");'
                    ]
                }
            }
        ],
        item: [
            {
                _postman_id: '9ea794a9-0215-2d62-0429-4483925e2174',
                name: 'F1',
                variable: [{
                    id: 'folder',
                    key: 'folder',
                    value: 'variable',
                    type: 'string',
                    description: 'This is a folder level variable'
                }],
                auth: {
                    type: 'digest',
                    digest: {
                        algorithm: 'MD5',
                        username: 'postman',
                        realm: '{{echo_digest_realm}}',
                        password: 'password',
                        nonce: '{{echo_digest_nonce}}',
                        nonceCount: '',
                        clientNonce: '',
                        opaque: '',
                        qop: ''
                    }
                },
                event: [
                    {
                        listen: 'test',
                        script: {
                            type: 'text/javascript',
                            exec: [
                                'console.log("Folder level test script");'
                            ]
                        }
                    },
                    {
                        listen: 'prerequest',
                        script: {
                            type: 'text/javascript',
                            exec: [
                                'console.log("Folder level pre request script");'
                            ]
                        }
                    }
                ],
                item: [
                    {
                        _postman_id: 'd547171c-fe63-41f9-6a9f-9e1b3e2ddc69',
                        name: 'Basic',
                        variable: [{
                            id: 'request',
                            key: 'request',
                            value: 'variable',
                            type: 'string',
                            description: 'This is a request level variable'
                        }],
                        event: [
                            {
                                listen: 'test',
                                script: {
                                    type: 'text/javascript',
                                    exec: [
                                        'console.log("Request level test script");'
                                    ]
                                }
                            },
                            {
                                listen: 'prerequest',
                                script: {
                                    type: 'text/javascript',
                                    exec: [
                                        'console.log("Request level pre request script");'
                                    ]
                                }
                            }
                        ],
                        request: {
                            auth: {
                                type: 'basic',
                                basic: {
                                    username: '{{username}}',
                                    password: '{{token}}'
                                }
                            },
                            method: 'GET',
                            header: [],
                            body: {
                                mode: 'raw',
                                raw: ''
                            },
                            url: 'https://postman-echo.com/get'
                        },
                        response: []
                    }
                ]
            },
            {
                _postman_id: 'c724c912-5cf6-f48f-a49e-a262e1806fa1',
                name: 'OAuth1',
                variable: [{
                    id: 'request',
                    key: 'request',
                    value: 'variable',
                    type: 'string',
                    description: 'This is a request level variable'
                }],
                event: [
                    {
                        listen: 'test',
                        script: {
                            type: 'text/javascript',
                            exec: [
                                'console.log("Request level test script");'
                            ]
                        }
                    },
                    {
                        listen: 'prerequest',
                        script: {
                            type: 'text/javascript',
                            exec: [
                                'console.log("Request level pre request script");'
                            ]
                        }
                    }
                ],
                request: {
                    auth: {
                        type: 'oauth1',
                        oauth1: {
                            consumerKey: 'XSCMjta2no3GN27g',
                            consumerSecret: 'I+EiQ-gq$-%@2Ts4',
                            token: '',
                            tokenSecret: '',
                            signatureMethod: 'HMAC-SHA1',
                            timestamp: '1500452534',
                            nonce: 'U2tGovXVvmI',
                            version: '1.0',
                            realm: ''
                        }
                    },
                    method: 'GET',
                    header: [],
                    body: {
                        mode: 'raw',
                        raw: ''
                    },
                    url: 'https://postman-echo.com/get'
                },
                response: []
            }
        ]
    },
    v21: {
        info: {
            name: 'nested-entities',
            _postman_id: '812e0da4-98e6-ed13-96e8-994c00d816b2',
            description: 'A simple collection to elucidate nested auth helpers, variables, and scripts',
            schema: 'https://schema.getpostman.com/json/collection/v2.1.0/collection.json'
        },
        variable: [{
            id: 'collection',
            key: 'collection',
            value: 'variable',
            type: 'string',
            description: 'This is a collection level variable'
        }],
        auth: {
            type: 'awsv4',
            awsv4: [
                {
                    key: 'accessKey',
                    value: '{{id}}',
                    type: 'string'
                },
                {
                    key: 'id',
                    value: 'awsSigV4',
                    type: 'string'
                },
                {
                    key: 'region',
                    value: 'us-east-1',
                    type: 'string'
                },
                {
                    key: 'secretKey',
                    value: '{{key}}',
                    type: 'string'
                },
                {
                    key: 'service',
                    value: 'iam',
                    type: 'string'
                },
                {
                    key: 'sessionToken',
                    value: '',
                    type: 'string'
                }
            ]
        },
        event: [
            {
                listen: 'test',
                script: {
                    type: 'text/javascript',
                    exec: [
                        'console.log("Collection level test script");'
                    ]
                }
            },
            {
                listen: 'prerequest',
                script: {
                    type: 'text/javascript',
                    exec: [
                        'console.log("Collection level pre request script");'
                    ]
                }
            }
        ],
        item: [
            {
                _postman_id: '6aa473d4-4909-7d40-2559-6709849b9976',
                name: 'F1',
                variable: [{
                    id: 'folder',
                    key: 'folder',
                    value: 'variable',
                    type: 'string',
                    description: 'This is a folder level variable'
                }],
                auth: {
                    type: 'digest',
                    digest: [
                        {
                            key: 'algorithm',
                            value: 'MD5',
                            type: 'string'
                        },
                        {
                            key: 'username',
                            value: 'postman',
                            type: 'string'
                        },
                        {
                            key: 'realm',
                            value: '{{echo_digest_realm}}',
                            type: 'string'
                        },
                        {
                            key: 'password',
                            value: 'password',
                            type: 'string'
                        },
                        {
                            key: 'nonce',
                            value: '{{echo_digest_nonce}}',
                            type: 'string'
                        },
                        {
                            key: 'nonceCount',
                            value: '',
                            type: 'string'
                        },
                        {
                            key: 'clientNonce',
                            value: '',
                            type: 'string'
                        },
                        {
                            key: 'opaque',
                            value: '',
                            type: 'string'
                        },
                        {
                            key: 'qop',
                            value: '',
                            type: 'string'
                        }
                    ]
                },
                event: [
                    {
                        listen: 'test',
                        script: {
                            type: 'text/javascript',
                            exec: [
                                'console.log("Folder level test script");'
                            ]
                        }
                    },
                    {
                        listen: 'prerequest',
                        script: {
                            type: 'text/javascript',
                            exec: [
                                'console.log("Folder level pre request script");'
                            ]
                        }
                    }
                ],
                item: [
                    {
                        _postman_id: 'd460728d-fb38-94f8-5d8d-7e9e2a7cfc43',
                        variable: [{
                            id: 'request',
                            key: 'request',
                            value: 'variable',
                            type: 'string',
                            description: 'This is a request level variable'
                        }],
                        name: 'Basic',
                        event: [
                            {
                                listen: 'test',
                                script: {
                                    type: 'text/javascript',
                                    exec: [
                                        'console.log("Request level test script");'
                                    ]
                                }
                            },
                            {
                                listen: 'prerequest',
                                script: {
                                    type: 'text/javascript',
                                    exec: [
                                        'console.log("Request level pre request script");'
                                    ]
                                }
                            }
                        ],
                        request: {
                            auth: {
                                type: 'basic',
                                basic: [
                                    {
                                        key: 'username',
                                        value: '{{username}}',
                                        type: 'string'
                                    },
                                    {
                                        key: 'password',
                                        value: '{{token}}',
                                        type: 'string'
                                    }
                                ]
                            },
                            method: 'GET',
                            header: [],
                            body: {
                                mode: 'raw',
                                raw: ''
                            },
                            url: {
                                raw: 'https://postman-echo.com/get',
                                protocol: 'https',
                                host: ['postman-echo', 'com'],
                                path: ['get']
                            }
                        },
                        response: []
                    }
                ]
            },
            {
                _postman_id: 'd399c332-9ee9-e05d-b24b-e124a0557db3',
                variable: [{
                    id: 'request',
                    key: 'request',
                    value: 'variable',
                    type: 'string',
                    description: 'This is a request level variable'
                }],
                name: 'OAuth1',
                event: [
                    {
                        listen: 'test',
                        script: {
                            type: 'text/javascript',
                            exec: [
                                'console.log("Request level test script");'
                            ]
                        }
                    },
                    {
                        listen: 'prerequest',
                        script: {
                            type: 'text/javascript',
                            exec: [
                                'console.log("Request level pre request script");'
                            ]
                        }
                    }
                ],
                request: {
                    auth: {
                        type: 'oauth1',
                        oauth1: [
                            {
                                key: 'consumerKey',
                                value: 'RKCGzna7bv9YD57c',
                                type: 'string'
                            },
                            {
                                key: 'consumerSecret',
                                value: 'D+EdQ-gs$-%@2Nu7',
                                type: 'string'
                            },
                            {
                                key: 'token',
                                value: '',
                                type: 'string'
                            },
                            {
                                key: 'tokenSecret',
                                value: '',
                                type: 'string'
                            },
                            {
                                key: 'signatureMethod',
                                value: 'HMAC-SHA1',
                                type: 'string'
                            },
                            {
                                key: 'timestamp',
                                value: '1500452534',
                                type: 'string'
                            },
                            {
                                key: 'nonce',
                                value: 'S0kXloMHurS',
                                type: 'string'
                            },
                            {
                                key: 'version',
                                value: '1.0',
                                type: 'string'
                            },
                            {
                                key: 'realm',
                                value: '',
                                type: 'string'
                            }
                        ]
                    },
                    method: 'GET',
                    header: [],
                    body: {
                        mode: 'raw',
                        raw: ''
                    },
                    url: {
                        raw: 'https://postman-echo.com/get',
                        protocol: 'https',
                        host: ['postman-echo', 'com'],
                        path: ['get']
                    }
                },
                response: []
            }
        ]
    }
};
