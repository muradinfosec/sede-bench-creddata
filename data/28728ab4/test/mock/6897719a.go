package versio

const tokenResponseMock = `
{
  "access_token":"718ae1cd-e138-08b7-0ce3-1ea09fa46e5d",
  "token_type":"bearer",
  "expires_in":3600
}
`

const tokenFailToFindZoneMock = `{"error":{"code":401,"message":"ObjectDoesNotExist|Domain not found"}}`

const tokenFailToCreateTXTMock = `{"error":{"code":400,"message":"ProcessError|DNS record invalid type _acme-challenge.example.eu. TST"}}`
