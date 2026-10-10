+++
title = "Data source HTTP API "
description = "Grafana Data source HTTP API"
keywords = ["grafana", "http", "documentation", "api", "data source"]
aliases = ["/http_api/datasource/"]
type = "docs"
[menu.docs]
name = "Data source"
parent = "http_api"
+++


# Data source API

## Get all datasources

`GET /api/datasources`

**Example Request**:

```http
GET /api/datasources HTTP/1.1
Accept: application/json
Content-Type: application/json
Authorization: Bearer eyJrIjoieUllcUR1Q2M4TkVpWDF6YjJJNHFnc3d0V05PNHF6Uw
```

**Example Response**:

```http
HTTP/1.1 200
Content-Type: application/json

[
  {
    "id":1,
    "orgId":1,
    "name":"datasource_elastic",
    "type":"elasticsearch",
    "access":"proxy",
    "url":"http://mydatasource.com",
    "password":"",
    "user":"",
    "database":"grafana-dash",
    "basicAuth":false,
    "basicAuthUser":"",
    "basicAuthPassword":"",
    "isDefault":false,
    "jsonData":null
  }
]
```

## Get a single data sources by Id

`GET /api/datasources/:datasourceId`

**Example Request**:

```http
GET /api/datasources/1 HTTP/1.1
Accept: application/json
Content-Type: application/json
Authorization: Bearer eyJrIjoibVZ0RXd0bXYzWW9tTzNJdDhKOEl6and3SmdIM1pzRw
```

**Example Response**:

```http
HTTP/1.1 200
Content-Type: application/json

{
  "id":1,
  "orgId":1,
  "name":"test_datasource",
  "type":"graphite",
  "access":"proxy",
  "url":"http://mydatasource.com",
  "password":"",
  "user":"",
  "database":"",
  "basicAuth":false,
  "basicAuthUser":"",
  "basicAuthPassword":"",
  "isDefault":false,
  "jsonData":null
}
```

## Get a single data source by Name

`GET /api/datasources/name/:name`

**Example Request**:

```http
GET /api/datasources/name/test_datasource HTTP/1.1
Accept: application/json
Content-Type: application/json
Authorization: Bearer eyJrIjoienZQbmluQ2U2bENUYjduczdqMmpiTW1GdXl6OXVYTA
```

**Example Response**:

```http
HTTP/1.1 200
Content-Type: application/json

{
  "id":1,
  "orgId":1,
  "name":"test_datasource",
  "type":"graphite",
  "access":"proxy",
  "url":"http://mydatasource.com",
  "password":"",
  "user":"",
  "database":"",
  "basicAuth":false,
  "basicAuthUser":"",
  "basicAuthPassword":"",
  "isDefault":false,
  "jsonData":null
}
```

## Get data source Id by Name

`GET /api/datasources/id/:name`

**Example Request**:

```http
GET /api/datasources/id/test_datasource HTTP/1.1
Accept: application/json
Content-Type: application/json
Authorization: Bearer eyJrIjoiYm9ZcE9XcWMyb1pUeDlqdzVBMVRJbm5PdEF3OUVRdQ
```

**Example Response**:

```http
HTTP/1.1 200
Content-Type: application/json

{
  "id":1
}
```

## Create data source

`POST /api/datasources`

**Example Graphite Request**:

```http
POST /api/datasources HTTP/1.1
Accept: application/json
Content-Type: application/json
Authorization: Bearer eyJrIjoiT0dqV3FBZmY0d3R2aTZSQTJFMVJySUtqY3pIM0tDaA

{
  "name":"test_datasource",
  "type":"graphite",
  "url":"http://mydatasource.com",
  "access":"proxy",
  "basicAuth":false
}
```

**Example CloudWatch Request**:

```http
POST /api/datasources HTTP/1.1
Accept: application/json
Content-Type: application/json
Authorization: Bearer eyJrIjoiRkZ5cUhOYkwxa014eDdqVTlJMVJiUHZKRHlmNHBiaA

{
  "name": "test_datasource",
  "type": "cloudwatch",
  "url": "http://monitoring.us-west-1.amazonaws.com",
  "access": "proxy",
  "jsonData": {
    "authType": "keys",
    "defaultRegion": "us-west-1"
  },
  "secureJsonData": {
    "accessKey": "Aa5xCHqdZXD3VvryAf4v", //should not be encoded
    "secretKey": "nSFebGUtPHudKrurNILpYCUjcxk5BSV6MCOu" //should be Base-64 encoded
  }
}
```

**Example Response**:

```http
HTTP/1.1 200
Content-Type: application/json

{"id":1,"message":"Datasource added", "name": "test_datasource"}
```

## Update an existing data source

`PUT /api/datasources/:datasourceId`

**Example Request**:

```http
PUT /api/datasources/1 HTTP/1.1
Accept: application/json
Content-Type: application/json
Authorization: Bearer eyJrIjoidGxjT3BmZ3QxQm96cDhNaDlEM0RDR1lBY1RIN0xwcA

{
  "id":1,
  "orgId":1,
  "name":"test_datasource",
  "type":"graphite",
  "access":"proxy",
  "url":"http://mydatasource.com",
  "password":"",
  "user":"",
  "database":"",
  "basicAuth":true,
  "basicAuthUser":"basicuser",
  "basicAuthPassword":"xehuddgrx",
  "isDefault":false,
  "jsonData":null
}
```

**Example Response**:

```http
HTTP/1.1 200
Content-Type: application/json

{"message":"Datasource updated", "id": 1, "name": "test_datasource"}
```

## Delete an existing data source by id

`DELETE /api/datasources/:datasourceId`

**Example Request**:

```http
DELETE /api/datasources/1 HTTP/1.1
Accept: application/json
Content-Type: application/json
Authorization: Bearer eyJrIjoiekdxUXZWUFo2V2ZPaTJycjZXOVpLY0xQSGVYNHZQYQ
```

**Example Response**:

```http
HTTP/1.1 200
Content-Type: application/json

{"message":"Data source deleted"}
```

## Delete an existing data source by name

`DELETE /api/datasources/name/:datasourceName`

**Example Request**:

```http
DELETE /api/datasources/name/test_datasource HTTP/1.1
Accept: application/json
Content-Type: application/json
Authorization: Bearer eyJrIjoiZmZZT2tTRVcyc2ZyTzREbTlwMXhmQU5FR0FkM3JDeQ
```

**Example Response**:

```http
HTTP/1.1 200
Content-Type: application/json

{"message":"Data source deleted"}
```

## Data source proxy calls

`GET /api/datasources/proxy/:datasourceId/*`

Proxies all calls to the actual datasource.
