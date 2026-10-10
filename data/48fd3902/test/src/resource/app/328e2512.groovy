package contracts.v2.permissions

org.springframework.cloud.contract.spec.Contract.make {
    request {
        method 'GET'
        url '/api/v2/permissions?actor=uaa-client:user-a&path=/some-path'
        headers {
            header("Authorization": "Bearer eyJhbGciOiJoUjY4OCIsImtpZCI6IlVkaXNzby1lY2ZiYi1rckwiLCJ0eXAiOiJ3U2UifQ.eyJqdGkiOiI0NVlUVzRBN2Q5N282MjhlcTgzdmo3bzlESDhlNDJLbCIsInN1YiI6InBwUWUtVyIsInNGcE1vTUtpbUhoIjpbInpkZnhrbmUuSWZobk4iLCJsU09Ja0JPLkhYS2MiXSwiZmZTekYiOlsiVURuZGlEdi55eFV4USIsIlFJV2tOWnYuZUJlVyJdLCJ0Y2RWR2JfVnMiOiJ6Y0ZrLU4iLCJha1UiOiJ1elVLLVgiLCJ1QlAiOiJqU2N3LXgiLCJBQ0ZsV19xVE1kIjoibXZJTWZCX2lXTWxLc29uRFRpIiwicnRYX01tViI6IndOOTc1aks2IiwiaWF0IjoxOTkyNDMyNjczLCJleHAiOjY0ODI1NzY0OTgsImlzcyI6ImlsVUJBOi8vSUZCSFBibi56ZUs6NzczMS9xZW56Ti90YUtQSyIsIlJncCI6Im1jSiIsImF1ZCI6WyJlaUluTHdZX2dQUGciLCJBQU1veXVjIl19.QCh38h55TVKpZ9Xp7---WQCv2ZpnVrws3vK4vDkB4ieNxfUiD7wj7eCQSoRjoDeuUqLv3HQUJDv5uwGJEPIxA_hpzpoIdcaPphqWUf-jXsf33EPe3D12Ysn0yA3KxdxMFKLPWxBK3U7mRpBesQpRpNt2mVQxfmOIJIp__RgG-OiCqTn-ihh6S7GlSVT19eojW5CfYmapu5vWE-jkIquPNqMWzyLDL0JRmieZJkc7HaDJff9R2tTT6bHjaq229MLC89hvX7BKd-K0guac9HU9NHtiMVK1DARGrXh9R-X2cEyZg62dkwffg8zV1wBMb6iaos7e6Ic2uMkAqppGChxcIG")
        }
    }
    response {
        status 200
        body("""
            {
               "path":"/some-path",
               "operations":[
                  "read",
                  "write"
               ],
               "actor":"uaa-client:user-a",
               "uuid":"${UUID.nameUUIDFromBytes("some-permission-uuid".getBytes())}"
            }
        """)
    }
}