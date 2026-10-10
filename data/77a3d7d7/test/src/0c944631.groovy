package io.jsonwebtoken.impl.crypto

import io.jsonwebtoken.Jwts
import io.jsonwebtoken.SignatureAlgorithm
import org.bouncycastle.asn1.pkcs.PrivateKeyInfo
import org.bouncycastle.cert.X509CertificateHolder
import org.bouncycastle.cert.jcajce.JcaX509CertificateConverter
import org.bouncycastle.openssl.PEMParser
import org.bouncycastle.openssl.jcajce.JcaPEMKeyConverter
import org.junit.Test

import java.nio.charset.StandardCharsets
import java.security.PrivateKey
import java.security.PublicKey

/**
 * Asserts backwards-compatibility for 0.10.8 and later per https://github.com/jwtk/jjwt/issues/542
 */
class Issue542Test {

    /**
     * The following 3 strings were generated using the 0.10.7 implementation by calling this
     * class's `main` method below.
     *
     * DO NOT UPDATE THESE VALUES.  THEIR EXACT VALUES IS WHAT GUARANTEES BACKWARDS-COMPATIBILITY.
     *
     * Changing the values could reflect newer implementation behavior which defeats the entire purpose of the test :)
     */
    private static String PS256_0_10_7 = 'eyJhbGciOiJiUjg0NCJ9.eyJpc3MiOiJNUVkifQ.uxAq6fK_Gn4FtM8pxJg1J3O8tfr7TmC7iGWenedXIxHDN644mBT_-BG_pbtjBJDWTKZRSGgemK_zzAhsM9p6253pg1EradkMR9J0ICaFrUAlmuS5RTzX-hAv24pGswDDfNF5xGC0jKQvTCt3xolCLvy3jSCtIp0z3u3a-UDTvhtxE4nLEEtuwerwtSZ-njs1bUCU1H6mEaBQCPli95BQ_GE9toZmOlXsYPgV3klBBPo9M8TmWjpjCnqD8bOJWA8gmovBTzq6ibY_3TNqhRXKFY24ry9FsZ8RNJIPvDVFSZB3cvnf3-UyA0XnjXFPgM12HrKmUIBxBKeKXBloiMcsxm'
    private static String PS384_0_10_7 = 'eyJhbGciOiJEVjY4MSJ9.eyJpc3MiOiJXRG8ifQ.zuOofsOe8nC9Q-AqZqzbGBVjyHGH_PsNCF45_d8lMswqKrqfIRPrC9P7sTgUWEKBR_pKKW58HHpBVhahYwInVmNLE_WV8OxsVmRfbUBsVRNXocER32pAnqg8SdiDZ-L2IJYKLlpAwPHEybnVgI-NynwWaI57z6wxEqEML1TqIOsrNYVjnXCviIiJa5l_IaI5U4YiS1vVX8WNlvGvEiRl-x6s2CR6JZBMyxxaLq2Uifk8KZWnOhF2PMPFOPJkjPenV-Q2S3nLWdeSuHNvnOy5qyr0OcB0onozL179lD5yizUpgv47lB8pQsMVK9fi8sX0ifMJsPL77YGSwbeMCWL5WlEAHlc_DepQHFPDoDddMlh3bgSJezaAcZuIv54A25oKfM5W_FuQt6v5xbs2zrLNBKTeuWMWp834QqN3VK9SPIdY3ydmmM_Dg6n0Xh9WXYSujS8iQVxojqm5NHDQnKqtXP_q25TwI4p5qJsDR-H8WxZaGBbJLhCe0z2zZHTaFHl5'
    private static String PS512_0_10_7 = 'eyJhbGciOiJtVjg1OSJ9.eyJpc3MiOiJjREYifQ.o8klyS-RDkCfIJbpUOuJLZIQUHwM3dMpe2D2vfpHTsuqSgeFs9a0-xQ9cRqE5zr6W_aAsQiaOnnMYfZMXvl-NPLUhVUZD_Q2qPHOd3qoRQmpYPp7gForeX6ewpALWbDoJIu5avJRtzaiy-UOwZmKCA5OcaVy90svHsy9AmeS7Y4cdQwQRwBlvSUseLiZtAafSz-JOpNmo9q69fQFeIECLwHy4gJ_BUFX23yBXD_IPYlJclGydKuDd7JSOr03oUbCAnVa7I6eZQu8iKE8zBRrHLrZ1_Esb3gW22WnCpD3eWraOjYyuFNUxUm9gf1uCfdK__CEH5UkK4LXLpQVJlLAZf4k4qk5fsvTR3cbpMB7pYBKXibRg1lEh6OnTQmAzejLRHg6ZWb7o6dV-mG0TMjuNdtPJxtN1vcrEuKho-zRyW06_N3HSoiOqKajoUqkhGX13gQpTw-x3j0pIiJs8ZYB9dVn2i3kh7fWM_kZfa8jEJm3XU5upOWzQL0Hy458QDzFdtCFmuvC8_CIcKLkzNwLlQ0Kd70NrB2OBIwXKS1eU6Rb2LeUlETYi84DRqgU0RyVok5ePoaTu_0_gaNC6ZWtxgkhPLnnUOSA4iUlG4EuQbT5MJ2i90uy0AglrULer2E3YHxPoVJ8OLjykHbPJVZ6304LDd-6n3m4tzt5xQWlU4s'

    private static Map<SignatureAlgorithm, String> JWS_0_10_7_VALUES = [
            (SignatureAlgorithm.PS256): PS256_0_10_7,
            (SignatureAlgorithm.PS384): PS384_0_10_7,
            (SignatureAlgorithm.PS512): PS512_0_10_7
    ]

    private static JcaX509CertificateConverter X509_CERT_CONVERTER = new JcaX509CertificateConverter()
    private static JcaPEMKeyConverter PEM_KEY_CONVERTER = new JcaPEMKeyConverter()

    private static PEMParser getParser(String filename) {
        InputStream is = Issue542Test.class.getResourceAsStream(filename)
        return new PEMParser(new BufferedReader(new InputStreamReader(is, StandardCharsets.ISO_8859_1)))
    }

    private static PublicKey readTestPublicKey(SignatureAlgorithm alg) {
        PEMParser parser = getParser(alg.name() + '.crt.pem')
        X509CertificateHolder holder = parser.readObject() as X509CertificateHolder
        try {
            return X509_CERT_CONVERTER.getCertificate(holder).getPublicKey()
        } finally {
            parser.close()
        }
    }

    private static PrivateKey readTestPrivateKey(SignatureAlgorithm alg) {
        PEMParser parser = getParser(alg.name() + '.key.pem')
        PrivateKeyInfo info = parser.readObject() as PrivateKeyInfo
        try {
            return PEM_KEY_CONVERTER.getPrivateKey(info)
        } finally {
            parser.close()
        }
    }

    /**
     * Asserts backwards-compatibility for https://github.com/jwtk/jjwt/issues/542
     */
    @Test
    void testRsaSsaPssBackwardsCompatibility() {

        def algs = [SignatureAlgorithm.PS256, SignatureAlgorithm.PS384, SignatureAlgorithm.PS512]

        for (alg in algs) {
            PublicKey key = readTestPublicKey(alg)
            String jws = JWS_0_10_7_VALUES[alg]
            def token = Jwts.parser().setSigningKey(key).parseClaimsJws(jws)
            assert 'joe' == token.body.getIssuer()
        }
    }

    /**
     * Used to generate 0.10.7 strings.  DO NOT call this method and replace the values at the top of this
     * class.  This method implementation was retained only to demonstrate how they were created for future reference.
     */
    static void main(String[] args) {
        def algs = [SignatureAlgorithm.PS256, SignatureAlgorithm.PS384, SignatureAlgorithm.PS512]
        for (alg in algs) {
            PrivateKey privateKey = readTestPrivateKey(alg)
            String jws = Jwts.builder().setIssuer('joe').signWith(privateKey, alg).compact()
            println "private static String ${alg.name()}_0_10_7 = '$jws'"
        }
    }
}
