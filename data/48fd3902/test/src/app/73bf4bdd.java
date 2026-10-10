package org.cloudfoundry.credhub;

final public class CertificateStringConstants {

  private CertificateStringConstants() {
    super();
  }

  /**
   * .
   * openssl x509 -in <(pbpaste) -text -noout
   * <p>
   * Subject Name: O=test-org,ST=Jupiter,C=MilkyWay,CN=test-common-name,OU=test-org-unit,L=Europa
   * Duration: 30 days
   * Key Length: 4096
   * Alternative Names: SolarSystem
   * Extended Key Usage: server_auth, client_auth
   * Key Usage: digital_signature
   * Issuer: CN=foo
   */
  public static final String BIG_TEST_CERT;
  static {
    BIG_TEST_CERT =
      "-----BEGIN CERTIFICATE-----\n"
    + "MIIEbzCCA1egAwIBAgIUYE3pB+BUaAP0YHeofpCmI/xCkmYwDQYJKoZIhvcNAQEL\n"
    + "BQAwDjEMMAoGA1UEAwwDZm9vMB4XDTE3MDIwMzAxNDMzNloXDTE3MDMwNTAxNDMz\n"
    + "NlowfDERMA8GA1UECgwIdGVzdC1vcmcxEDAOBgNVBAgMB0p1cGl0ZXIxETAPBgNV\n"
    + "BAYTCE1pbGt5V2F5MRkwFwYDVQQDDBB0ZXN0LWNvbW1vbi1uYW1lMRYwFAYDVQQL\n"
    + "DA10ZXN0LW9yZy11bml0MQ8wDQYDVQQHDAZFdXJvcGEwggIiMA0GCSqGSIb3DQEB\n"
    + "AQUAA4ICDwAwggIKAoICAQCk+byx2uL5QNAQqdeEWoD0NfuXdbtf/j7orjK7TjCn\n"
    + "djM21HLnIq96hZ+/Vxg30oxjRqAMKDUIj8OTrisorgcgpLNV7NwklPG9A0gv7xdk\n"
    + "YvxhnEnyrztZYiS8sx98YDwjQpJDeA45QX6+/9k8qmXf7XRQRRTqhkG8jpkk0vvj\n"
    + "hvwwTAma+0xALWfVBnhLJz82snJI+ezM9OuwO53iOkziNHNxtuc5sq/AjuDf48O4\n"
    + "HOtxdB27WniL0T3+4Ng8ZRAgMmlSrQdFn6x/Us32VTVLTD4x1s9H4HL3c8LZJaoD\n"
    + "CEKIwFRn2lSko4b4PAUGHZ0KpfeTlur0uR464s4PHh+EV1DOm/R/1/HIQgKanr5B\n"
    + "FzLONAqFPPCMB/hTPli+Q6nez+Q2alpyxEz/QTCTNROKl+opVWJW6gAPMXAkpqc2\n"
    + "bx0O7fRwwF9evVcQ1BZdfWaG3iGqO60o0y7lEAmvlOwnw0JjSta2NDlR0nNp9frx\n"
    + "85USPdMoSBoaAGb+BehbvFsVoRTToxCo0YwCDcGjgacR4oCu5zTZd0KUVlGJ0vpu\n"
    + "OJiUALyYSD/6mN8ZIfPa/rR8PF5ju3JzGd/AEh0F8gfgardLrNHre904/0HwBqvc\n"
    + "ShdSS0XHjA7nTLAyARgLU/E0TIL9DH63tWrB2W+m7vBMkU0fuY5c40QIalT/iGbj\n"
    + "8QIDAQABo1cwVTAWBgNVHREEDzANggtTb2xhclN5c3RlbTAOBgNVHQ8BAf8EBAMC\n"
    + "B4AwHQYDVR0lBBYwFAYIKwYBBQUHAwEGCCsGAQUFBwMCMAwGA1UdEwEB/wQCMAAw\n"
    + "DQYJKoZIhvcNAQELBQADggEBAF6cvPlK7kuoC/EBmTF4OdTrqzKYa4bm47hn0e7V\n"
    + "3fmgt5YJl/HM/wQI5Dw37y5SjtCScvZUJQ7N+vKQ06bPJHea8f3+XwgsM0JpUhkV\n"
    + "djZR9L5LrJdwmhk4arbMLMKeFwB1Xir72trL1DreI/Kzsow0LbMhllLWPyRHmAhr\n"
    + "Kqu/WgGim6m3lVgZdx4o6cguGry+ceiunCwCFL36CL1AdvYL8ZnUlQDT1hNp3anE\n"
    + "QTHPRc0mETzHET0uL+9UpaUxglRPzuxVhyIYimXSiPQlk8K43gmXM8QKi85eo8xD\n"
    + "W5kgC9Eel5YQcs5wUS/1aW72x2D+7DeGxLjFwm0Sy9S8hfI=\n"
    + "-----END CERTIFICATE-----";
  }

  /**
   * Version: 3 (0x2)
   * Signature Algorithm: sha256WithRSAEncryption
   * Issuer: CN=test.example.com, OU=app:c07353f2-b2e5-3513-a5d1-482c8f653ecb, L=exampletown
   * Validity
   * Not Before: Mar 29 20:23:53 2017 GMT
   * Not After : Mar 27 20:23:53 2027 GMT
   * Subject: CN=test.example.com, OU=app:a50530d6-e7a0-8522-d8a9-384f7c403abf, L=exampletown
   * Subject Public Key Info:
   * Public Key Algorithm: rsaEncryption
   * RSA Public Key: (2048 bit)
   * X509v3 extensions:
   * X509v3 Basic Constraints: critical
   * CA:FALSE
   */
  public static final String SIMPLE_SELF_SIGNED_TEST_CERT;
  static {
    SIMPLE_SELF_SIGNED_TEST_CERT =
      "-----BEGIN CERTIFICATE-----\n"
    + "MIIDZjCCAk6gAwIBAgIUD8qb4pWPNR6dhYtZgf5Cn5DymHowDQYJKoZIhvcNAQEL\n"
    + "BQAwZDEZMBcGA1UEAwwQdGVzdC5leGFtcGxlLmNvbTExMC8GA1UECwwoYXBwOmI2\n"
    + "NzQ0NmU1LWIyYjAtNDY0OC1hMGQwLTc3MmQzZDM5OWRjYjEUMBIGA1UEBwwLZXhh\n"
    + "bXBsZXRvd24wHhcNMTcwMzI5MjAyMzUzWhcNMjcwMzI3MjAyMzUzWjBkMRkwFwYD\n"
    + "VQQDDBB0ZXN0LmV4YW1wbGUuY29tMTEwLwYDVQQLDChhcHA6YjY3NDQ2ZTUtYjJi\n"
    + "MC00NjQ4LWEwZDAtNzcyZDNkMzk5ZGNiMRQwEgYDVQQHDAtleGFtcGxldG93bjCC\n"
    + "ASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAMle5V28ZST5EvfqPLVk6Wg5\n"
    + "i7poGmtGTkFBhi4YSIhD/cKqMlc3eh7X0cm7JYbMtK7WJ/QTKEKfzrFB/NCmu1C7\n"
    + "840lo+9tA4jDuZksi6qclTcjPP5tvjeCkRhBXw1qOjHvT5doPIGsGYNzKIEjobRN\n"
    + "fPT1Z83laKO9x3JXatulpx3VrJCeAVZADEUYonCU+l5cplWxgh1Zy1rbVgx2NSoL\n"
    + "LYst208h9I6oq0X/u2VmiPIXDE8iblexvhQwrLE31R7AsGfsObGZ05KSBsZBJsyE\n"
    + "xmb5K44IuKahpZnkX9x95fuJUzYdxcjkR9KgSak2uLFY0pki87PH6bOoApDfnNkC\n"
    + "AwEAAaMQMA4wDAYDVR0TAQH/BAIwADANBgkqhkiG9w0BAQsFAAOCAQEAvwGXYcQL\n"
    + "WNzts9Hx6Rih8Vj5ibORRl/49m820dFoxZzGmxwUiHFveni8cbkW3WkFEI6iUFRl\n"
    + "cH9SilytWi+WnQr8zPUC80OgyOUHd8JRLxn84cJBxqKsMvn2vhFYhCMhJSgzyIOF\n"
    + "2hCIx4uwTEizCeq/DWZJhSSbEkdGW2CCON97VyNAhUQ+lwkJG5cEDSEDmtqT8MD+\n"
    + "fCe0Wu650ipk2gGMZnbnRCi4phxb/bx0rEyq5upuOZw5Ja6/Ulikm4pF3mwW1xuR\n"
    + "OgN3QbqNnD+GGJnsddnTY2o900nXOki+v4tcnRTlfhnkWIXn3kF+ZgmacUQDdq+G\n"
    + "zNIcsYa7oxP46A==\n"
    + "-----END CERTIFICATE-----";
  }

  /**
   * Version: 3 (0x2)
   * Signature Algorithm: sha256WithRSAEncryption
   * Issuer: CN=test.example.com, OU=app:1b4dec0c-98ac-98b1-f4a3-06bfd9bb83c5, L=exampletown
   * Validity
   * Not Before: Mar 29 20:23:06 2017 GMT
   * Not After : Mar 27 20:23:06 2027 GMT
   * Subject: CN=test.example.com, OU=app:9f1eed7e-47cf-51e1-a4f1-52ddd7da62c6, L=exampletown
   * Subject Public Key Info:
   * Public Key Algorithm: rsaEncryption
   * RSA Public Key: (2048 bit)
   * X509v3 extensions:
   * X509v3 Basic Constraints: critical
   * CA:FALSE
   */
  public static final String TEST_CERT_WITH_INVALID_UUID_IN_ORGANIZATION_UNIT;
  static {
    TEST_CERT_WITH_INVALID_UUID_IN_ORGANIZATION_UNIT =
    "-----BEGIN CERTIFICATE-----\n"
      + "MIIDZjCCAk6gAwIBAgIULvDbv0eGhurO4sQHPGQzw1JZzBswDQYJKoZIhvcNAQEL\n"
      + "BQAwZDEZMBcGA1UEAwwQdGVzdC5leGFtcGxlLmNvbTExMC8GA1UECwwoYXBwOjdl\n"
      + "MGZiZDdkLTE0YmQtMTFlNy1hOGIxLTEwZGRiMWFhNjRiMzEUMBIGA1UEBwwLZXhh\n"
      + "bXBsZXRvd24wHhcNMTcwMzI5MjAyMzA2WhcNMjcwMzI3MjAyMzA2WjBkMRkwFwYD\n"
      + "VQQDDBB0ZXN0LmV4YW1wbGUuY29tMTEwLwYDVQQLDChhcHA6N2UwZmJkN2QtMTRi\n"
      + "ZC0xMWU3LWE4YjEtMTBkZGIxYWE2NGIzMRQwEgYDVQQHDAtleGFtcGxldG93bjCC\n"
      + "ASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBAKGehnx7S8uvY3jyuWmkZgY8\n"
      + "aCYQoIs5jVevYI5ajyyFX2F6sv6+4p0FDB1+9sM2/PIbyKHrqMcvTJwdgbW01nfo\n"
      + "mRipq91K07u6J387U78RS/XIwKg9OcxSn0RH1f32oSOqiS+miWmgek7kFrGM2hsp\n"
      + "ySPmTqr36m3Tn5Z0sy9u6B/2U1kfYMNhd0KEHrhQ6zGUfxiMXwVXm5JYsKHadlzq\n"
      + "xo1VqOSKzbVWayIlo+g2c1mcv0DXeXBcoKh9APAGKyjUwXq69pFsprq3gj45lQVE\n"
      + "8dp9ukRX4eWYgPxBublToJtUGlKth8s9rfhSVjhnkp7TT9WDor9G4/iPRna/KiUC\n"
      + "AwEAAaMQMA4wDAYDVR0TAQH/BAIwADANBgkqhkiG9w0BAQsFAAOCAQEAQL89uAb8\n"
      + "wqS1ZlVNC6mGXLrrORvhCIS+2LI3Cllo87G055kATS3u5yGq/6EzAM+4W6LePg94\n"
      + "5HT3Jyo1uO/D9rNfgJ7+yxaotwr13p0L7yeKO3XT4jOaav0UIPFDq8TcMfADilOa\n"
      + "w1FHWyMxn3wSbD6KYzjEnpo932C+CtXfQPX+dtCayR/NO7DxhAbCfIhAZNIbA8xn\n"
      + "g8GslKQKmPFVpclGjRjeMOgSOdsN/Dg0OwzKBMSoWuUGLavICh0HwlccW6+cxgUs\n"
      + "Pswa7CjG14hVnFYnj3ISDrXetgwX3Gw4kbwFdWERzDrYvoNzdguKYkrop+mUGqpy\n"
      + "OeRYsf1HBP8d0w==\n"
      + "-----END CERTIFICATE-----";
  }

  /**
   * Version: 3 (0x2)
   * Signature Algorithm: sha256WithRSAEncryption
   * Issuer: CN=test.example.com, OU=4a7eff1d-23af-63c3-a8c8-80abb0da43a2, L=exampletown
   * Validity
   * Not Before: Mar 29 20:35:54 2017 GMT
   * Not After : Mar 27 20:35:54 2027 GMT
   * Subject: CN=test.example.com, OU=9a9aed7e-61ba-61f9-c0d9-87dad6ea16d4, L=exampletown
   * Subject Public Key Info:
   * Public Key Algorithm: rsaEncryption
   * RSA Public Key: (2048 bit)
   * X509v3 extensions:
   * X509v3 Basic Constraints: critical
   * CA:FALSE
   */
  public static final String TEST_CERT_WITH_INVALID_ORGANIZATION_UNIT_PREFIX;
  static {
    TEST_CERT_WITH_INVALID_ORGANIZATION_UNIT_PREFIX =
    "-----BEGIN CERTIFICATE-----\n"
      + "MIIDXjCCAkagAwIBAgIUf0A6ygebPHdmubKEJHEelx4e/SMwDQYJKoZIhvcNAQEL\n"
      + "BQAwYDEZMBcGA1UEAwwQdGVzdC5leGFtcGxlLmNvbTEtMCsGA1UECwwkYjY3NDQ2\n"
      + "ZTUtYjJiMC00NjQ4LWEwZDAtNzcyZDNkMzk5ZGNiMRQwEgYDVQQHDAtleGFtcGxl\n"
      + "dG93bjAeFw0xNzAzMjkyMDM1NTRaFw0yNzAzMjcyMDM1NTRaMGAxGTAXBgNVBAMM\n"
      + "EHRlc3QuZXhhbXBsZS5jb20xLTArBgNVBAsMJGI2NzQ0NmU1LWIyYjAtNDY0OC1h\n"
      + "MGQwLTc3MmQzZDM5OWRjYjEUMBIGA1UEBwwLZXhhbXBsZXRvd24wggEiMA0GCSqG\n"
      + "SIb3DQEBAQUAA4IBDwAwggEKAoIBAQDKKmO/4dYA19ETDA5sCl1Ai9QZ5vtyljnT\n"
      + "gPW+w/4KIx6pJ+3pEdfnh5D7SEgWqWvK9/TUh37/6gIdMIMMCNPVIV8hN4oL7AAj\n"
      + "NLdqJJR9KtwvWiHzr/2VHrwRtYReJJ2MoUKhDvz3aRykJxFL4kVoCHtKpkTyVFtK\n"
      + "WOTKoRgBvF9hz0nUe27sHABy4hRafuLXM+gG5cwTuXTRYg3mXhLN6V9e6hV1cS4C\n"
      + "cRp10N+ncE+9FmO1i0jjSv4F1d4/nE6dNicJqhpfqJhIwVCp+W+VAYyOhOiel1ae\n"
      + "BJBx6ju5Qvwp8X5l1BXTNdHU5c93llNttaFOh9bvIA0MlVWMy4t/AgMBAAGjEDAO\n"
      + "MAwGA1UdEwEB/wQCMAAwDQYJKoZIhvcNAQELBQADggEBAAPMSsGhCP18oEiteTDO\n"
      + "rFXgoXdlQi8zZuynn5cEAVgWyqV7repy3XRoq7A7n3Y7vUDNDpKixBHp15wXAP93\n"
      + "edrwfoPFowZVXEicdMjlqhf62esa/A0gNo6vQvBIsos/OZZxhpY8kxskdlKAPY4v\n"
      + "XAHiXsOKw767M9qDeBA3SFSrEL4YMfdhoMoUZr/ayK1Y/j8xPn0CL33rNWkv7SoX\n"
      + "jnfmsSqWxhg/t/mIjvCR3EfOfpSHpPE4tCCEKD2TvrRxrZO2UBOZ+Y1mc6L33GGV\n"
      + "fpgSc+FTem/FObrbCUDruX3hCVTQY3Gn1YiSvblPRay29eeR4JeIyPYnzWW06zN8\n"
      + "IMg=\n"
      + "-----END CERTIFICATE-----";
  }

  /**
   * Version: 3 (0x2)
   * Signature Algorithm: sha256WithRSAEncryption
   * Issuer: CN=test.example.com, L=exampletown
   * Validity
   * Not Before: Mar 28 23:36:57 2017 GMT
   * Not After : Mar 26 23:36:57 2027 GMT
   * Subject: CN=test.example.com, L=exampletown
   * Subject Public Key Info:
   * Public Key Algorithm: rsaEncryption
   * RSA Public Key: (2048 bit)
   * X509v3 extensions:
   * X509v3 Basic Constraints: critical
   * CA:FALSE
   */
  public static final String TEST_CERT_WITHOUT_ORGANIZATION_UNIT;
  static {
    TEST_CERT_WITHOUT_ORGANIZATION_UNIT =
      "-----BEGIN CERTIFICATE-----\n"
    + "MIIDADCCAeigAwIBAgIUEkUiS0BpLu5eITSbvH3g8HjYplEwDQYJKoZIhvcNAQEL\n"
    + "BQAwMTEZMBcGA1UEAwwQdGVzdC5leGFtcGxlLmNvbTEUMBIGA1UEBwwLZXhhbXBs\n"
    + "ZXRvd24wHhcNMTcwMzI4MjMzNjU3WhcNMjcwMzI2MjMzNjU3WjAxMRkwFwYDVQQD\n"
    + "DBB0ZXN0LmV4YW1wbGUuY29tMRQwEgYDVQQHDAtleGFtcGxldG93bjCCASIwDQYJ\n"
    + "KoZIhvcNAQEBBQADggEPADCCAQoCggEBANViOGj5z5ExS1e0BsBJ0HObG9u6c2Bx\n"
    + "VgDGTNS7UeeN9mBjI7XCvMy/en1awqTayVJeA2QPWZKnPahjvC+02/VSOizBy1Vp\n"
    + "PTeM9cetdO94dKadYQqMbMn2163Z1F/kPLjL/8FcBpl+FmrB4rPl/+Oo1TjoUvXp\n"
    + "z/Ys2625RnC4sNlNs0B2Yq3xXg4ysdvS2rD4n8Q4LHKPugkdcRRel0WqPNuMbkGY\n"
    + "IXJNc7zD8CngO7KI8TGsjPkkwSdVdyihAc/DuZRNqdm7YH4KEUeTPHlQVrjxFlR6\n"
    + "1xHATn0q4gTwRMmqqYb9Vs1HMhWz1TFaDQAbZplkapW7fcbnGXwF1o0CAwEAAaMQ\n"
    + "MA4wDAYDVR0TAQH/BAIwADANBgkqhkiG9w0BAQsFAAOCAQEAdyHwsKBuCGjRc/2i\n"
    + "I/vrNTZrn0GYUCH8zdFxs0c/h59feYM6SWjm9iBIIFxxtGP/F/xxLeblNHDPeTdp\n"
    + "O5CTU1c+mfjaExXIxJXdzKA7IyfCljc0G1tANGqL/P12/POExiDBaIBWwu+AC2ED\n"
    + "VPx5tsfOxNxQB1axMaDIGHeJ+SB2/3FhY6VcD6DoCY6TwpJ0mwGP6sgQhJr3qCbE\n"
    + "P/3asUqyGfS5qkFmvwTytnydR/69RZquVJn9m6ZWeOaEekJ7tTy3AiDKN+OYYpzT\n"
    + "ZC7Tw5ZqvuWhXUIkFmAqkTW2CFI8zqSbpLjTHvjxP14Ck9zSuV4oojhvtfjA/wOz\n"
    + "mAeI1w==\n"
    + "-----END CERTIFICATE-----";
  }

  /**
   * This cert appears to be self-signed (Issuer DN == Subject DN) but is actually
   * signed with a CA of the same name.
   * <p>
   * Version: 3 (0x2)
   * Signature Algorithm: sha256WithRSAEncryption
   * Issuer: CN=trickster
   * Validity
   * Not Before: Feb  7 19:41:38 2017 GMT
   * Not After : Feb  7 19:41:38 2018 GMT
   * Subject: CN=trickster
   * X509v3 Basic Constraints: critical
   * CA:FALSE
   */
  public static final String MISLEADING_CERT;
  static {
    MISLEADING_CERT =
      "-----BEGIN CERTIFICATE-----\n"
    + "MIICxjCCAa6gAwIBAgIUdYAQkigXgqz1FFN0Qi2T6k2dYJswDQYJKoZIhvcNAQEL\n"
    + "BQAwFDESMBAGA1UEAwwJdHJpY2tzdGVyMB4XDTE3MDIwNzE5NDEzOFoXDTE4MDIw\n"
    + "NzE5NDEzOFowFDESMBAGA1UEAwwJdHJpY2tzdGVyMIIBIjANBgkqhkiG9w0BAQEF\n"
    + "AAOCAQ8AMIIBCgKCAQEAz9r5xtR3Iy/A9mooCA+NjetzqbN4hzzj5ZcJlE/x6UGg\n"
    + "CF393EtOfOUkimDKRgdMFXUnFLBskns9TtmmuQfcGmNV5FDwQpBX2brJbEcOOE3q\n"
    + "dJgV7HDeu1bX+ZpPm0wOwtQXyxr72j03Lt6V+Htrgana3qcejGq8syRO5FEruCsj\n"
    + "vRGlWXd3CYF7/UlVhlc06ILqD+CoePArPrussxAqt4UoUSuoJeaqEvGyi8IJ4mqu\n"
    + "pqUAvWsh8u79A+T+om3JHA7RFhV2Vu83dj+CwMfMsdDiyiaRz5wyVTDpu7X2AeGi\n"
    + "Y3uuJAm/nc6qyDmuAo6BhnFux6dLLlCJ0EqyEzjAQwIDAQABoxAwDjAMBgNVHRMB\n"
    + "Af8EAjAAMA0GCSqGSIb3DQEBCwUAA4IBAQDPa6O7GIdyQOwUWYYuf7ed8Qj28Jy7\n"
    + "K8utG7SBmE6iPt978m1VW3u01YMh61yNOcs9RBGE7I2H28Zuu9ut4MY8BNBV00rE\n"
    + "sAaPHeOQmEGPPt5oZ1ztlAjSOK9YHOfHaKj4ISfMdhDQ4vAOxGbj3NJMsxiYL4oH\n"
    + "4yQp0lZ5S1TVNZXn77sr5HXuzy+2r40HDc/4Q/9iQseB6Aypjii1Q6m8eej6rcnC\n"
    + "snC8luPGbSwo31gKr9wFxv78GJcswIGt6fi4CxV7eGWn0p9EY4NsR8jdatLd/eKD\n"
    + "qA2eKfjSi415xgI1eOf89HvoYKlBGYuFxXB3YRkJfpS+khFeu7HTsyj2\n"
    + "-----END CERTIFICATE-----";
  }

  /**
   * Version: 3 (0x2)
   * Serial Number:
   * 14:7b:69:d1:56:80:9b:07:26:f6:0e:03:0a:1f:7a:ed:96:12:be:c2
   * Signature Algorithm: sha256WithRSAEncryption
   * Issuer: CN=foo.com
   * Validity
   * Not Before: Feb 10 22:38:59 2017 GMT
   * Not After : Feb 10 22:38:59 2018 GMT
   * Subject: CN=foo.com
   * X509v3 extensions:
   * X509v3 Basic Constraints: critical
   * CA:TRUE
   */
  public static final String SELF_SIGNED_CA_CERT;
  static {
    SELF_SIGNED_CA_CERT =
      "-----BEGIN CERTIFICATE-----\n"
    + "MIICxTCCAa2gAwIBAgIUFHtp0VaAmwcm9g4DCh967ZYSvsIwDQYJKoZIhvcNAQEL\n"
    + "BQAwEjEQMA4GA1UEAwwHZm9vLmNvbTAeFw0xNzAyMTAyMjM4NTlaFw0xODAyMTAy\n"
    + "MjM4NTlaMBIxEDAOBgNVBAMMB2Zvby5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IB\n"
    + "DwAwggEKAoIBAQDFxKUdiiwrJ9fmW5iD+jz4Ofi/dH1zX6hzh1yO5AEb0B34BmcX\n"
    + "kFCjdKMidXLFY7tAbpFrZHx/kwldoXN2ee9xro5/mgaTMQgDGdUrtk2Z0laUs05I\n"
    + "fiCLWlUnCvNqIEBXc2c56Qj6dAn/cWAJIw6rH6+gLuenMYr4PGiK0rzQHQ8pAW42\n"
    + "y9XOwPuMVT62z16Qbf9Np5S6sKYIjunFSbCte3MsGyYb9A+BXlCmwgKcMBgUHT+c\n"
    + "cxICkM++yPogju+dNp2RSIIk4Ji5KAxUyT3JjZDb6dyygMjftfvqH+FUYuLIiVYR\n"
    + "UhjmXGRrYm402HZXXXxT+jcHKhAPvHofeJxzAgMBAAGjEzARMA8GA1UdEwEB/wQF\n"
    + "MAMBAf8wDQYJKoZIhvcNAQELBQADggEBAJvRkqdhiX/3fNDKBly3HxIg4ACX+pF4\n"
    + "GmG5NL+WY/o6jMxEyW3BMSx67Jkgzi9BnTw8lTSkil7lXJjG9Homzi/HJuFsceH2\n"
    + "lMtAOlLxsvczA90jyp0qGdy7L7hrKPvSGZ2vUehaTIjdcteL6c7GwlotSA7w6G8i\n"
    + "Xj//23SsjYge8IalqnFonYG8u9q0u7N0U0G9C6Rdo5WzCAIuzRBkT8ka0ZAaTgSV\n"
    + "iSIanWI+olDxF/IRhnRCU1HmvAImyvDd7CEjI0Y+HSXqZfc+c1RvfbtoR9fS4n/v\n"
    + "6qAlCDS+r+SU+v19UsFW8cC31UyWQwrpaAMYPYhwqK1Z4nPN+AgAGGY=\n"
    + "-----END CERTIFICATE-----";
  }

  /**
   * //.
   * Version: 3 (0x2)
   * Serial Number:
   * 12:65:5e:4c:79:5e:9c:ac:81:b8:34:1d:98:68:ca:52:d3:47:e5:ed
   * Signature Algorithm: sha256WithRSAEncryption
   * Issuer: CN=bar
   * Validity
   * Not Before: Feb 13 20:21:34 2017 GMT
   * Not After : Feb 13 20:21:34 2018 GMT
   * Subject: CN=bar
   */
  public static final String V3_CERT_WITHOUT_BASIC_CONSTRAINTS;
  static {
    V3_CERT_WITHOUT_BASIC_CONSTRAINTS =
      "-----BEGIN CERTIFICATE-----\n"
    + "MIICqDCCAZCgAwIBAgIUEmVeTHlenKyBuDQdmGjKUtNH5e0wDQYJKoZIhvcNAQEL\n"
    + "BQAwDjEMMAoGA1UEAwwDYmFyMB4XDTE3MDIxMzIwMjEzNFoXDTE4MDIxMzIwMjEz\n"
    + "NFowDjEMMAoGA1UEAwwDYmFyMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKC\n"
    + "AQEA5liBEeqqjTWcofS3tz4xm4CST1vYgoEt2CTEf16iKHiiirK9NIRX+OB5U6nW\n"
    + "6rXP8q7PUQ6DmbDfmjNFvWrKxuuPhcDY4M+g++pJ7VZ46E00e3yNqYoAJv/8pFfz\n"
    + "yhLc/CYF7JgfIh3fPv6gjY2Q8Asi1fIzgv17LaP7NJtFTxzPzg9hxHhHZY8OyMnm\n"
    + "/+LhXd4SPETctNndMfp97g9z2epf/sYLQMUcoq/CJQf4ilNkeK07E6om9JkGxOLl\n"
    + "6mWLCIHyx2FXLIvSPE+zvCHeK+6PDX0RarBSKRNntH+7aNlo3k7VOHZYM6Xj6qmE\n"
    + "mOyWHxUlHgVy4myY5vqd8R7+PwIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQAgH22O\n"
    + "1Tn4xeVvl5fdBR77Ywr+lSLxpZUcOm2XP4WJd5a5xzZvqC9SdhFADWEeGBXc4pG3\n"
    + "12Wmp8y/9mmOu4Yc708UCL9aimKnl4+sX9fkZnnbQrIc/bODdi3B0I+17Irn6aAe\n"
    + "0mXEok5kYMFOk77ZZF3OG8bUojHGINKvW1bamMEqL4F4VU1WIfgBgDf4F1MdbCAZ\n"
    + "KtCY6DlUwxcFDWrktx1XGFZC5PSedBHWYcx0bsV+lllCfEm48PXqUnIOjpjsbVg1\n"
    + "+NHBdPGyKlnRkvI4xIodUXfn7FrS3KWL8k9NzL5FGXCvzXoGZKBvcJuu0/2XXl+3\n"
    + "Jykws+qOGWF351JZ\n"
    + "-----END CERTIFICATE-----";
  }

  /**
   * Version: 3 (0x2)
   * Serial Number:
   * c7:0b:53:ea:27:6e:9f:26
   * Signature Algorithm: sha1WithRSAEncryption
   * Issuer: C=US, ST=NY, O=Test Org, OU=app:a12345e5-b2b0-4648-a0d0-772d3d399dcb,
   * CN=example.com/emailAddress=test@example.com
   * Validity
   * Not Before: Mar 30 18:34:52 2017 GMT
   * Not After : Mar 30 18:34:52 2018 GMT
   * Subject: C=US, ST=NY, O=Test Org, OU=app:a12345e5-b2b0-4648-a0d0-772d3d399dcb,
   * CN=example.com/emailAddress=test@example.com
   * X509v3 extensions:
   * X509v3 Subject Key Identifier:
   * 01:69:9A:20:83:EB:49:5D:7C:16:68:A0:7F:B9:D6:A5:D4:5B:FC:06
   * X509v3 Authority Key Identifier:
   * keyid:01:69:9A:20:83:EB:49:5D:7C:16:68:A0:7F:B9:D6:A5:D4:5B:FC:06
   * <p>
   * X509v3 Key Usage:
   * Digital Signature, Key Encipherment
   * X509v3 Extended Key Usage:
   * TLS Web Client Authentication
   */
  public static final String SELF_SIGNED_CERT_WITH_CLIENT_AUTH_EXT;
  static {
    SELF_SIGNED_CERT_WITH_CLIENT_AUTH_EXT =
      "-----BEGIN CERTIFICATE-----\n"
    + "MIIDEjCCAnugAwIBAgIJAMcLU+onbp8mMA0GCSqGSIb3DQEBBQUAMIGXMQswCQYD\n"
    + "VQQGEwJVUzELMAkGA1UECAwCTlkxETAPBgNVBAoMCFRlc3QgT3JnMTEwLwYDVQQL\n"
    + "DChhcHA6YTEyMzQ1ZTUtYjJiMC00NjQ4LWEwZDAtNzcyZDNkMzk5ZGNiMRQwEgYD\n"
    + "VQQDDAtleGFtcGxlLmNvbTEfMB0GCSqGSIb3DQEJARYQdGVzdEBleGFtcGxlLmNv\n"
    + "bTAeFw0xNzAzMzAxODM0NTJaFw0xODAzMzAxODM0NTJaMIGXMQswCQYDVQQGEwJV\n"
    + "UzELMAkGA1UECAwCTlkxETAPBgNVBAoMCFRlc3QgT3JnMTEwLwYDVQQLDChhcHA6\n"
    + "YTEyMzQ1ZTUtYjJiMC00NjQ4LWEwZDAtNzcyZDNkMzk5ZGNiMRQwEgYDVQQDDAtl\n"
    + "eGFtcGxlLmNvbTEfMB0GCSqGSIb3DQEJARYQdGVzdEBleGFtcGxlLmNvbTCBnzAN\n"
    + "BgkqhkiG9w0BAQEFAAOBjQAwgYkCgYEAtyHjDd5AYxUr0KFvlyeiDysyFuzoIW5y\n"
    + "9AS8LblFEsYwfVF7MU8kQ4O0HzosWbCv72U6vYBz0vmolcK5F/4LBbgqCFZQRrFV\n"
    + "jzkW60ertb2te3OHeRuxKYcA7EE4UosXulQ/3lJmS1XHBPtZOJRdelTLoHyUZ37y\n"
    + "BUa6ZkGvi90CAwEAAaNkMGIwHQYDVR0OBBYEFAFpmiCD60ldfBZooH+51qXUW/wG\n"
    + "MB8GA1UdIwQYMBaAFAFpmiCD60ldfBZooH+51qXUW/wGMAsGA1UdDwQEAwIFoDAT\n"
    + "BgNVHSUEDDAKBggrBgEFBQcDAjANBgkqhkiG9w0BAQUFAAOBgQAg6sfi3lMbWMUY\n"
    + "ObbubrWB2pPWo7W6GCsem2M0Qri8RuMUqzHNvatUnaIuD9yatdP+vDuFdGbX3PTF\n"
    + "wjh44yBrxcRGcYFdzBOIos+i5MMtDWorJY/0IH8mVEd3ckMcfVgMVP25ZvlYjhdY\n"
    + "dCNw/Rku+LafRQJOlSBZqy7K7qtzbA==\n"
    + "-----END CERTIFICATE-----";
  }

  /**
   * Version: 3 (0x2)
   * Serial Number:
   * fb:bd:39:9b:59:d5:98:4d
   * Signature Algorithm: sha1WithRSAEncryption
   * Issuer: C=US, ST=NY, O=Test Org, OU=app:a12345e5-b2b0-4648-a0d0-772d3d399dcb,
   * CN=example.com/emailAddress=test@example.com
   * Validity
   * Not Before: Mar 30 19:04:56 2017 GMT
   * Not After : Mar 30 19:04:56 2018 GMT
   * Subject: C=US, ST=NY, O=Test Org, OU=app:a12345e5-b2b0-4648-a0d0-772d3d399dcb,
   * CN=example.com/emailAddress=test@example.com
   * X509v3 extensions:
   * X509v3 Subject Key Identifier:
   * 2B:77:1B:63:E2:92:6B:29:9D:31:D4:0E:E0:79:0D:5F:61:E9:77:1A
   * X509v3 Authority Key Identifier:
   * keyid:2B:77:1B:63:E2:92:6B:29:9D:31:D4:0E:E0:79:0D:5F:61:E9:77:1A
   */
  public static final String SELF_SIGNED_CERT_WITH_NO_CLIENT_AUTH_EXT;
  static {
    SELF_SIGNED_CERT_WITH_NO_CLIENT_AUTH_EXT =
      "-----BEGIN CERTIFICATE-----\n"
    + "MIIC8DCCAlmgAwIBAgIJAPu9OZtZ1ZhNMA0GCSqGSIb3DQEBBQUAMIGXMQswCQYD\n"
    + "VQQGEwJVUzELMAkGA1UECAwCTlkxETAPBgNVBAoMCFRlc3QgT3JnMTEwLwYDVQQL\n"
    + "DChhcHA6YTEyMzQ1ZTUtYjJiMC00NjQ4LWEwZDAtNzcyZDNkMzk5ZGNiMRQwEgYD\n"
    + "VQQDDAtleGFtcGxlLmNvbTEfMB0GCSqGSIb3DQEJARYQdGVzdEBleGFtcGxlLmNv\n"
    + "bTAeFw0xNzAzMzAxOTA0NTZaFw0xODAzMzAxOTA0NTZaMIGXMQswCQYDVQQGEwJV\n"
    + "UzELMAkGA1UECAwCTlkxETAPBgNVBAoMCFRlc3QgT3JnMTEwLwYDVQQLDChhcHA6\n"
    + "YTEyMzQ1ZTUtYjJiMC00NjQ4LWEwZDAtNzcyZDNkMzk5ZGNiMRQwEgYDVQQDDAtl\n"
    + "eGFtcGxlLmNvbTEfMB0GCSqGSIb3DQEJARYQdGVzdEBleGFtcGxlLmNvbTCBnzAN\n"
    + "BgkqhkiG9w0BAQEFAAOBjQAwgYkCgYEAwnMy7/SfQIzPe+9GJujQVHb2OP5dOMvz\n"
    + "LSQLBwJtHdzv2xtnDL+vFB4F1Cv6NeTbhnWi89ebyOqGgqyeX7wga8UN9Pmih16Q\n"
    + "/i8Bx//+TJjUqj1umhGZmGL4sWx8NwrP9QiuvrNf/sdq0ilbclWgn5UBuiAKty48\n"
    + "BjaEpEGf0AMCAwEAAaNCMEAwHQYDVR0OBBYEFCt3G2PikmspnTHUDuB5DV9h6Xca\n"
    + "MB8GA1UdIwQYMBaAFCt3G2PikmspnTHUDuB5DV9h6XcaMA0GCSqGSIb3DQEBBQUA\n"
    + "A4GBAJ9SipBB/J17usRtaDp59iIlTiyF1N14Qw1XuR0ZpVGXCx5r46DQVHQIdno+\n"
    + "EPa2sOCml5CIOlNko6Edr0GlWOWkhZyTyZTW6oHaDxVQXrhbVemmKOUY0LM9r2l+\n"
    + "TqtBv192PPCKbrJChCVmrDltpc5F5TeyTAWn/ElIVz6Za+y/\n"
    + "-----END CERTIFICATE-----";
  }

  /**
   * Version: 3 (0x2)
   * Serial Number: 1 (0x1)
   * Signature Algorithm: sha256WithRSAEncryption
   * Issuer: CN=testCa
   * Validity
   * Not Before: Sep 21 17:24:09 2017 GMT
   * Not After : Sep 21 17:24:15 2027 GMT
   * Subject: CN=testCa
   * Subject Public Key Info:
   * Public Key Algorithm: rsaEncryption
   * X509v3 extensions:
   * X509v3 Key Usage: critical
   * Certificate Sign, CRL Sign
   * X509v3 Basic Constraints: critical
   * CA:TRUE, pathlen:0
   * X509v3 Subject Key Identifier:
   * B2:23:50:C6:8F:7F:B4:5F:62:1A:AD:00:91:35:FA:91:A3:BD:9E:5A
   */
  public static final String CERTSTRAP_GENERATED_CA_CERTIFICATE;
  static {
    CERTSTRAP_GENERATED_CA_CERTIFICATE =
    "-----BEGIN CERTIFICATE-----\n"
    + "MIIE4jCCAsqgAwIBAgIBATANBgkqhkiG9w0BAQsFADARMQ8wDQYDVQQDEwZ0ZXN0\n"
    + "Q2EwHhcNMTcwOTIxMTcyNDA5WhcNMjcwOTIxMTcyNDE1WjARMQ8wDQYDVQQDEwZ0\n"
    + "ZXN0Q2EwggIiMA0GCSqGSIb3DQEBAQUAA4ICDwAwggIKAoICAQDYMdo2q+GvAPn2\n"
    + "kUm82ZMLMJjOVpYHv2/G7f9gBuTHb1Z9uJcnH3rUTIh0JNNAz+/IuuRxBJb5rWFW\n"
    + "K4Dp1vINRG7ltcDQXQkS1mPigt/5rSHWv3+WSZ4tZmMw6JUxDjQ+kqO9sjDkGm4Y\n"
    + "MG4Qkmmo7hjAKei+rv1beN8m6wNSBTfDKtsZZZeXe7Ru9bHCzo/1f4Q2XZ5OYoOk\n"
    + "4gznM2JRwbbhbcpRX36QuFzSUXSMcXSaCR6WfC/1R78lz5IqdjRLTLSJe7tSh7JV\n"
    + "5ZRr6UiGvcdS0DWXXG7Ers36FfSeaJUX3LLwMy2AKcsbLcou0GzY+Mt7QqK2XH6y\n"
    + "NydxqqNtDz5wZu0117DtWBLYZP/OxUySyh5rA9JrryuSCE9SD98Mrno0Ho1OqC/P\n"
    + "3rwXdoShKHmLSbsT7wo85xOi9iBsJHpG69uK9EPz8WuuUguczVoOA8L4NQ/11Imv\n"
    + "/ZliomfhICojZihnK0DhH/WBr4xgYUVc0wfJ7zAgOUmQ+pUrzNjTmTzV/zWKsT2R\n"
    + "t+s66kxJfoqNGKRfP//YVRZrlThFph8Px4Y+ZZ49f3wucANT9Pa5A4Yx+mLdSW2d\n"
    + "13YYEWwHDWobM67T+u+ttsMBpBfa86i6oVMORnneEOw9UCsiM7hvrBnZEFS0l5dI\n"
    + "2hP1ZWTp8EMkdtCFGDB1ek5SXYz4DwIDAQABo0UwQzAOBgNVHQ8BAf8EBAMCAQYw\n"
    + "EgYDVR0TAQH/BAgwBgEB/wIBADAdBgNVHQ4EFgQUsiNQxo9/tF9iGq0AkTX6kaO9\n"
    + "nlowDQYJKoZIhvcNAQELBQADggIBAEIMx9YXZhBHwsUaRyXmNV0vWty1UReboYLh\n"
    + "XXf4i273hiBFNP0bWBLE0UCJlKm8QzZ6DtgG4sp0Su+oQfMOR3OArSc8lIPOXIkh\n"
    + "MtgwsWyhLPsjkEk0B+AG0RT8LZHSzYuIQhF0sRbE1nIeIHNRjQp+jHIv8HsZI45B\n"
    + "rh3O1PKi8D2Er12omrY0pfDug7u0D0vGdfPygeXVg3aoNTwMu4P37wfcRNgtazdg\n"
    + "Y2oFEw7U6bs8qHz9CyfjSgcuUs3utOqKhCQpOUEzRdsvsgT5EVss1eaYr9SfSiZQ\n"
    + "k/Vmy1XRqLOuhHcZdsnOMmK8WxIIyDDlq0xouHezLQhA1lXvOFmlnDtl4cMQBVuv\n"
    + "MqZgst1ItyGKxJa9Pxnjpyz+T/w3i6ljb46HxLitUSc8jTYXO/LlPWX8MbFPRuZZ\n"
    + "BvgY8wdes1vrz3wGkm5mJ6LMzmSuvXhmIzaqvGNEwBdOLiOhW/QOFPIwD2jNONnq\n"
    + "FQp8lD0jLLfZG5U+G8gb9nbaxDovmJixRje6q8qN9dLQeVyCi52SNJr9rtqPhdm1\n"
    + "RI+RRCbIKBJrLRLP16fTj1f0bM3+TH2OoS5hPR1LTnYNYyDJxwE2+o1mtmKyccN1\n"
    + "7vzbwOrOcyPZ5WBEBG1H1vfMTrAefklycBWxRK/ZtpfH5cZDjBxP7x26I2LY5/ws\n"
    + "fvToHltS\n"
    + "-----END CERTIFICATE-----\n";
  }

  public static final String CERTSTRAP_GENERATED_CA_PRIVATE_KEY;
  static {
    CERTSTRAP_GENERATED_CA_PRIVATE_KEY = "-----BEGIN RSA PRIVATE KEY----- fake\n"
    + "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\n"
    + "-----END RSA PRIVATE KEY-----\n";
  }

  public static final String PRIVATE_KEY;
  static {
    PRIVATE_KEY = "-----BEGIN RSA PRIVATE KEY----- fake\n"
      + "MIIEowIBAAKCAQEAweZOJBR+ZPEqUauEw9uLvvjgR1xy4I4VhVfMsVZr41djeDP2\n"
      + "3oGMvmepvc48eLacMsR6hBGvFYG5VK4kpdA7k1jfiIsGqptlcFtJ3Wh+kvtiacmO\n"
      + "o+ozC1CuIpEqxIhb45yRng55O6uHI7SB2d6rHl6vDUsipNVKrU/rlp8qVRavcxcF\n"
      + "h0cPWmk2cKWfqMcqVw5mO/3O1TLa2AkR/hgk8HYgcP4jaDBv14nq05VgM27QRVl7\n"
      + "9AxkJy3Mhy7I6goz9QQtCwE8kWrdf+qPvXioYUGC9u1sQ7PtKDm9kzNatXbagor2\n"
      + "sFOeapsYOaHh56tI1hfPxrY49OrIUGxpUFoY+tRREYgXTdQkdGROA34uaZkG7AVf\n"
      + "eJ0/BofknOSKYxHlpuxmIDSqVgoas3QG0qKoYC1sdXgJkeFgkwSe1/X7+uAiP0hW\n"
      + "JPpSMH/di5lFbGHbBfXOQhdvpkwcNWHGFkarwwa1l/qEmMBOkhgbMh/jHRmYwnPw\n"
      + "Nmft9i8IMbnKdA6QgXPYO4Wj7+uoqdU6cwaFKu8BsmIK2rC6Na1PmC8Pm31k/a5j\n"
      + "Jgl3PNJ6updtpCVFwTy7LY/UAFUiD+RVULNk2tDazMdrO27sfgWrHUERyAc3MpZd\n"
      + "ZxYNE1SedM50a8gkKgRjEPss3qegMOW2BpTMeziI2iW4u8+Jk2SseSZBoEWs1b2W\n"
      + "lz5H5omCLGfo0Bs/9pSG6VUZQaIOmeA1oMxhhLKf/HdeKXUQkiXrXo6kPzwmwwRH\n"
      + "ok5B0raW9Aw+rS58OizWSCwoFgT42otMt7kB00g/Ez/5QEuShu8+ryE80JgnXrfV\n"
      + "PdIU1X4cxf6YgkFRdmCKtW4Un8t/r+wtY+jvIOEdTR6igWEdTpAJHwnamcsd8x/N\n"
      + "pmpyUiIaR819wg5fsFnTjaz4tMqU2/KuiBJrtH5WdXK4AL7LPBNp+C2NK6GYkjeX\n"
      + "d09b5knwSamny7ThCn/OdSKXHj9OmJBmdLU07JVKT+Ph4CIzg5fd8yMJ0vSLxnDu\n"
      + "1ttLg2atA+PkDZhmypaUNTI1UjU1viUF+6e4vhcYUFIVdOiQZWsS9xaaXIAzPxzW\n"
      + "XKsKmiKpv5DfRz6ULEaOElQ60i6Eyb6Pfj/mwIhcPr8qfis9KNeWXyj0SSgs6dD7\n"
      + "jo+YO35uj4cbItfCJnN0MHjjUe78G5gkVnEOKjcRh3sAaifALqOz5kDiFVpnmZhJ\n"
      + "RaL+aLSsmsjzBqu/85xhh0xncDJZHwFTekoL7R16ySrk3RA6hujs41T8l93akcHD\n"
      + "hJmx6pnxfZO4vPo9p/O1uPQof+IIpAz4cBuda+7K7w/TPCuKg/N7XQz0GU2K1Ae2\n"
      + "5ooT6IOoQI4vd+v9CB930Z1+Nkhf295hjESrjkB0EIsj5gManAyQMytOySZiUaGx\n"
      + "0frlrcFsiOgY6re9i5rg1MMCR92ZXoOheDlj/pJeKtqgs6a3VsiAwVxOGenb1yct\n"
      + "I8OQZoOQAKNO8X62gIW0bO3JBzXeHNgthO1Ls2mTMOxdwpyMkeg4Y6O0tA2PUHaP\n"
      + "7KO9Gl5+LP9w8qtNdtT6B02Joev0I1Z9udvDoyinH9wjpF2Xjk4e\n"
      + "-----END RSA PRIVATE KEY-----\n"
      + "\n";
  }
}
