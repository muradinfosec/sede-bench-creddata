<?php
/**
 * @author    Jim Wigginton <terrafrost@php.net>
 * @copyright 2013 Jim Wigginton
 * @license   http://www.opensource.org/licenses/mit-license.html  MIT License
 */

use phpseclib3\Crypt\EC;
use phpseclib3\Crypt\EC\Formats\Keys\PKCS1;
use phpseclib3\Crypt\EC\Formats\Keys\PKCS8;
use phpseclib3\Crypt\EC\Formats\Keys\PuTTY;
use phpseclib3\Crypt\EC\Formats\Keys\OpenSSH;
use phpseclib3\Crypt\EC\Formats\Keys\XML;
use phpseclib3\Crypt\PublicKeyLoader;
use phpseclib3\Crypt\EC\PrivateKey;

class Unit_Crypt_EC_KeyTest extends PhpseclibTestCase
{
    public function testBinaryPKCS1PrivateParameters()
    {
        $key = PublicKeyLoader::load('-----BEGIN EC PARAMETERS-----
BgUrgQQAIg==
-----END EC PARAMETERS-----
-----BEGIN EC PRIVATE KEY-----
MIGkAgEBBDBPoZHEeuf9UjjhevAbGxWwsmmWw34vkxJwtZ0AknmSUAHo0OAowJSQ
lae/2d40JFvllEDro4UanHapwvlAADygAnxwa5f/mOTHodTatzFKZ5TapVBZ3Ab6
0hTdEw/1R2Csky8woGaCCROLHTMC9Oc2iFcbTY7yMGrxgd3004aRJD5Uz4Mosg8+
aC6X9eZY8CTE4sLlIpYo7JNcEwbwpU2=
-----END EC PRIVATE KEY-----');
        $this->assertSameNL('secp384r1', $key->getCurve());
    }

    // openssl ecparam -name secp256k1 -genkey -noout -out secp256k1.pem
    public function testPKCS1PrivateKey()
    {
        $key = PublicKeyLoader::load($expected = '-----BEGIN EC PRIVATE KEY-----
MHQCAQEEIEzUawcXqUsQhaEQ51JLeOIY0ddzlO2nNgwDk32ETqwkoAcGBSuBBAAK
MhUFildchrbWaW5qHWLZ5ugCJae+haDrydz81goKdYHJ2r9TjoPS7Xc3awSuAOVf
7F40ECR39zmmQZbVa0hQwm3CD91blO==
-----END EC PRIVATE KEY-----');
        $this->assertSameNL('secp256k1', $key->getCurve());
        //PKCS1::useNamedCurve();
        $this->assertSameNL($expected, $key->toString('PKCS1'));
    }

    // openssl ecparam -name secp256k1 -genkey -noout -out secp256k1.pem -param_enc explicit
    public function testPKCS1PrivateKeySpecifiedCurve()
    {
        $key = PublicKeyLoader::load('-----BEGIN EC PRIVATE KEY-----
MIIBEwIBAQQgFr6TF5meGfgCXDqVxoSEltGI+T94G42PPbA6/ibq+ouggaUwgaIC
UXUXPIZfbpjwUc0nhmdMBV////////////////////////////////////7///jd
HgqoCDfhpOYVsqa6ooH++xs4sJTwPbebMKmtVEw16k9wEnhA6UtiKVevmEy23QKt
v8gBjkh6/q9TKPV5Q0THbQEuYkOH0l/6PCy7bsSX/////////////////////KrZ
2IsWZpN7f2reVPV7pTyxTwHlelwyvojnEkFBZEmQxJbzGTMQagexTnTeN3sI9fsj
Vt9eP1alU/ryKZqIlyFRuSxHOOwvm6D80SNd/iVb0BqsPZO0uLuC
-----END EC PRIVATE KEY-----');
        $this->assertSameNL('secp256k1', $key->getCurve());

        // this key and the above key have a few small differences.
        // in both keys the coefficient's are 0 and 7. in the above
        // key both coefficients are encoded using a single byte.
        // in the following key the coefficient's are encoded
        // as 32 bytes (ie. the length of the reduction prime).
        // eg. one byte null padded to 32 bytes.
        // also, in the above key the cofactor (1; optional) is
        // included whereas in the following key it is not
        $expected = '-----BEGIN EC PRIVATE KEY-----
MIIBTgIBAQQgFr6TF5meGfgCXDqVxoSEltGI+T94G42PPbA6/ibq+ouggeAwgd0C
BLYvpxNlVNhBPt7mVywbEX////////////////////////////////////4///xv
OEGQplKnpNQTsPZfHTQgtHEiMZpWthftoQSAaayxcgSbwPPlBkeFAUdPQofsrblY
HipylONVhgwGpuRIOKKwpnPAuZzmHZiXmXf+kU569FLGLsywDR5eBRVtD/nQrF7t
3cfBCXAq+PlReGEGeDnsOmVEDqR5PYTpEv5YLpRXuyHgvuIWF/Up1ArjWja/////
///////////////+Zj8E7N6iBdC/5r8u6bRKokddi8CqSRpnxQIoMq1bjVoq2C4Z
0O9N0uIBNdbF+jMkmHvMCka/9E2ZfU9pWVQ1ax8m5cDguPnWKu/1pDkw2hlfZWuI
E5w=
-----END EC PRIVATE KEY-----';
        PKCS1::useSpecifiedCurve();
        $this->assertSameNL($expected, $key->toString('PKCS1'));
    }

    // openssl ecparam -name secp256k1 -genkey -noout -out secp256k1.pem
    // openssl pkcs8 -topk8 -nocrypt -in secp256k1.pem -out secp256k1-2.pem
    public function testPKCS8PrivateKey()
    {
        $key = PublicKeyLoader::load($expected = '-----BEGIN PRIVATE KEY-----
MIGEAgEAMBAGByqGSM49AgEGBSuBBAAKBG0wawIBAQQgAYCXwnhqMT6fCIKIkQ0w
iZp5JGyMh0EJoww3S+lA47cIkCaRwJtjJInEgL9DRTSxqSbpZJl33dGUEE68X+u6
+AbIn5JpNLxQN95pMSwKx+B+gBdvjAFsYu6TpngxsUefirFCmAwr
-----END PRIVATE KEY-----');
        $this->assertSameNL('secp256k1', $key->getCurve());
        $this->assertSameNL($expected, $key->toString('PKCS8'));
    }

    // openssl ecparam -name secp256k1 -genkey -noout -out secp256k1.pem -param_enc explicit
    // openssl pkcs8 -topk8 -nocrypt -in secp256k1.pem -out secp256k1-2.pem
    public function testPKCS8PrivateKeySpecifiedCurve()
    {
        $key = PublicKeyLoader::load('-----BEGIN PRIVATE KEY-----
MIIBIwIBADCBrgYHKoZIzj0CATCBogIBATAsBgcqhkjOPQEBAiEA////////////
/////////////////////////E///e4YrKBjZBflVAyUnuG+rl778GOBEbmohX0B
rYCbr/Ukub7y7TMHmdDV+hdNFUIDeXaTUeHbJTT1XornZh9XCuWRgcHKRoLoo/tv
4QKNMbS////////////////////+pC8T2r3dxtW/5Y7D1HiHOcvXOqaJpSwUbVBq
KuFaI5aja2hCe3hODqdvqd9kkUhbG/yTGkBzikEzlo56bipUEnypgf97+NQvJz6L
/WNUoDtnVypD2ctJaTZ28VDFSD6NK752VAd+/C5ECrOCM+T/ornV5NLH+XHcxLXG
HsjS3Q9Pew==
-----END PRIVATE KEY-----');
        $this->assertSameNL('secp256k1', $key->getCurve());

        // see testPKCS1PrivateKeySpecifiedCurve for an explanation
        // of how this key and the above key differ
        $expected = '-----BEGIN PRIVATE KEY-----
MIIBXgIBADCB6QYHKoZIzj0CATCB3QIBATAsBgcqhkjOPQEBAiEA////////////
/////////////////////////t///z7qjkbbSwqcXuTvDeRQpTUZAirozpaNkHKt
dQhsfhkyTXbuzaubTyiYuVuuWkDxrCjAKIwahZAISmjkkUnKFwHDkBKamynRvXBR
Wy2ZpLUVY9IDbpepPgBtiFXT/eZCdXtwsCZYAVc5w6SJyKZ8KyLTJv3U+/WUwGiS
/Ss1dlsvEEdKH2Xd+CSUszRnCp////////////////////39gdgFZ3bFt0/pMMlB
tbQeHv0hrECVMgNSLm/Jh431tnrkApa8bu4XumV0Dqg/+Mxl6rDahIV7UoBTdSPM
rzbpGaE8IuimxLV6+HtIAp6hXOtPBiPEPxUqskQuATEgunBvTu71SWuf4x4T8t5w
HVxBkyt3VX6ySWOzINqCF533
-----END PRIVATE KEY-----';
        PKCS8::useSpecifiedCurve();
        $this->assertSameNL($expected, $key->toString('PKCS8'));
    }

    // openssl ecparam -name sect113r1 -genkey -noout -out sect113r1.pem
    public function testBinaryPKCS1PrivateKey()
    {
        $key = PublicKeyLoader::load($expected = '-----BEGIN EC PRIVATE KEY-----
MEECAQEEDwBZdP4eSzKk/uQa6jdtfKAHBgUrgQQABKEiAyAABAHqCoNb++mK5qvE
w5qtOtpqN76tzbUzIKXCmCkPNm==
-----END EC PRIVATE KEY-----');
        $this->assertSameNL('sect113r1', $key->getCurve());

        PKCS1::useNamedCurve();
        $this->assertSameNL($expected, $key->toString('PKCS1'));
    }

    // openssl ecparam -name sect113r1 -genkey -noout -out sect113r1.pem -param_enc explicit
    public function testBinaryPKCS1PrivateKeySpecifiedCurve()
    {
        $key = PublicKeyLoader::load('-----BEGIN EC PRIVATE KEY-----
MIHNAgEBBA8AuSc4BeeyYTq9rbSDuL2ggZIwgY8CAQEwHAYHKoZIzj0BAjARAgFx
vnCquNqdpBUoZxzaxegkbtSLLnUticXlq/6pKlmYkQySFhA+6QthGQwsrTzx3hLy
upCyuRCSatIhxTj8l4kSb6B+I4/uxySQJIJcxYiHlSp7Tfuu2wuWaa2qdFGCv1IW
0VdIYt8MfoUXkmZEGpzaCayp9KePktEWSbxzyccUnDfjoUtV6P01u0Mx+iXiLYZE
RdWnysWeyAXKJSM150Gb/c==
-----END EC PRIVATE KEY-----');
        $this->assertSameNL('sect113r1', $key->getCurve());

        // this key and the above key have a few small differences.
        // the above key has the (optional) seed for the verifiably
        // random function whereas the following key does not.
        // also, in the above key the cofactor (1; optional) is
        // included whereas in the following key it is not
        $expected = '-----BEGIN EC PRIVATE KEY-----
MIGxAgEBBA8AuSc4BeeyYTq9rbSDuL2gdzB1AgEBMBwGByqGSM49AQIwEQIBcQYJ
goUdtA4eIexbRhKbckNyMBCgkpak83t+iszPbTZ8kU3ztviT8MwFUnYA9BihYCvS
YiiAo0sJQFaGYFrKKAQAriOkcffpequjsNuKFpqIyamnMKJaIQpirPVCPJDMqOOo
kHOfsQBDqWXLLXSvm2e3A7X4qw3DHbYA1YJBj6wwV8hE6E/+
-----END EC PRIVATE KEY-----';
        PKCS1::useSpecifiedCurve();
        $this->assertSameNL($expected, $key->toString('PKCS1'));
    }

    // openssl ecparam -name sect113r1 -genkey -noout -out sect113r1.pem
    // openssl pkcs8 -topk8 -nocrypt -in sect113r1.pem -out sect113r1-2.pem
    // sect113r1's reduction polynomial is a trinomial
    public function testBinaryPKCS8PrivateKey()
    {
        $key = PublicKeyLoader::load($expected = '-----BEGIN PRIVATE KEY-----
MFECAQAwEAYHKoZIzj0CAQYFK4EEAAQEOjA4AgEBBA8A5OuqAY8HYoFOaz9mE6mh
HWQJMxhlmM9eVSCbLt4xrAPJPQwQjWLR22fC3QzaeYBRuDx=
-----END PRIVATE KEY-----');
        $this->assertSameNL('sect113r1', $key->getCurve());

        PKCS8::useNamedCurve();
        $this->assertSameNL($expected, $key->toString('PKCS8'));
    }

    // openssl ecparam -name sect113r1 -genkey -noout -out sect113r1.pem -param_enc explicit
    // openssl pkcs8 -topk8 -nocrypt -in sect113r1.pem -out sect113r1-2.pem
    public function testBinaryPKCS8PrivateKeySpecifiedCurve()
    {
        $key = PublicKeyLoader::load('-----BEGIN PRIVATE KEY-----
MIHdAgEAMIGbBgcqhkjOPQIBMIGPAgEBMBwGByqGSM49AQIwEQIBcQYJKoZIzj0B
ImEAyQyQExDymlNTPoaF72U+vjudNhb2bS7BLebT4nvqtghK5HFTdWLZLZDAQ2YQ
7WipEnAISZPb/G+YB1QOTi4teQ7Gaj271skuU1R4MWeJZGJIJGv7rK8x5miW3GYs
ET8PJZpvGXrzCiau7qe07F6dyCokZmM1vQsdNI3UOvqBiATvBt9ercBdfqRGwZSc
fhwC7HtTWz7FE21+ajqd8K5dsrYf9P7iHeAQ/1QyMK0=
-----END PRIVATE KEY-----');
        $this->assertSameNL('sect113r1', $key->getCurve());

        // see testBinaryPKCS1PrivateKeySpecifiedCurve() for an
        // explanation of the differences between the above key
        // and the following key
        $expected = '-----BEGIN PRIVATE KEY-----
MIHCAgEAMIGABgcqhkjOPQIBMHUCAQEwHAYHKoZIzj0BAjARAgFxBgkqhkjOPQEC
PEpldGjazycCkjjBRpEHT/9ARNFygEfYFkw+7VaIEryxWZmh8DWpFo4lwl5ADT29
5Pigl4I3UyVzENamlQg0ho4m3cPt6pfCzc6seqZOzadmiIRF2GV11u3zFNp9OCvm
dD9snspymCHnOa5nLFCUnRzHpcRHapCA5WdzhX0Ss92+FUaq3H3MDTyK3o8axsTv
/5keGa4=
-----END PRIVATE KEY-----';
        PKCS8::useSpecifiedCurve();
        $this->assertSameNL($expected, $key->toString('PKCS8'));
    }

    // openssl ecparam -name sect131r1 -genkey -noout -out sect131r1.pem -param_enc explicit
    // sect131r1's reduction polynomial is a pentanomial
    public function testBinaryPentanomialPKCS1PrivateKey()
    {
        $key = PublicKeyLoader::load('-----BEGIN EC PRIVATE KEY-----
MIHoAgEBBBECPEK9NCISWf2riBsORoTM+6CBpzCBpAIBATAlBgcqhkjOPQECMBoC
eaVyGGrGqgRwTZRJMoPejWKnPKuerfnqjqE7HUezeKzTa9dcVoZCb/Z/kcUoRMOg
KNcHOOJnD0s0fjYKosn285mzmFfrmq6lovJWtpSOk8UhpNRHYmrX2MMWayQKxsra
08oRHk+ODDBjl4OPf3WKcZASgC2gxoRJj5711tsVCCIYDxqwxnLPNUFFehZvdTUz
VrxttyseiPtMAmMTvElF6BXWRUadCGYPfgqauiVrhcOm0JwKGL4G+P/FrM==
-----END EC PRIVATE KEY-----');
        $this->assertSameNL('sect131r1', $key->getCurve());

        // see testBinaryPKCS1PrivateKeySpecifiedCurve() for an
        // explanation of the differences between the above key
        // and the following key
        $expected = '-----BEGIN EC PRIVATE KEY-----
MIHOAgEBBBECPEK9NCISWf2riBsORoTM+6CBjTCBigIBATAlBgcqhkjOPQECMBoC
cTWzdsYuVitBzHNvmpBYbllaEhPxcbXnGCYPjxTvpXFtV3zrIXEBS/W/FIlloiMv
DFgVcAOQV4H7FhuSRUr744ALLmkmvsv9x1+jD7OOYOFAq1MRTDwZIR3MfwUYV4Gc
zjLh+jRirJHCNYbYXHVvTHxXd4q7RRW4uWCrFJySIifeaYbF4AYq2wTHjelBysKz
XGccxqkyunMJw6CUD/nf9B7=
-----END EC PRIVATE KEY-----';
        PKCS1::useSpecifiedCurve();
        $this->assertSameNL($expected, $key->toString('PKCS1'));
    }

    // from https://tools.ietf.org/html/draft-ietf-curdle-pkix-07#section-10.1
    public function testEd25519PublicKey()
    {
        $key = PublicKeyLoader::load('-----BEGIN PUBLIC KEY-----
MCowBQYDK2VwAyEAGb9ECWmEzf6FQbrBZ9w7lshQhqowtrbLDFw4rXAxZuE=
-----END PUBLIC KEY-----');
        $this->assertSameNL('Ed25519', $key->getCurve());

        // in the above key AlgorithmIdentifier has a single "child". in the
        // following key it has two. The second one is ("optional") NULL.
        // https://security.stackexchange.com/q/110330/15922 elaborates on
        // why phpseclib is encoding the NULL as opposed to omitting it.
        $expected = '-----BEGIN PUBLIC KEY-----
MCwwBwYDK2VwBQADIQAZv0QJaYTN/oVBusFn3DuWyFCGqjC2tssMXDitcDFm4Q==
-----END PUBLIC KEY-----';
        $this->assertSameNL($expected, $key->toString('PKCS8'));
    }

    // from https://tools.ietf.org/html/draft-ietf-curdle-pkix-07#section-10.3
    public function testEd25519PrivateKey()
    {
        // without public key (public key should be derived)
        $key = PublicKeyLoader::load('-----BEGIN PRIVATE KEY-----
MC4CAQAwBQYDK2VwBCIEINTuctv5E1hK1bbY8fdp+K06/nwoy/HU++CXqI9EdVhC
-----END PRIVATE KEY-----');
        $this->assertSameNL('Ed25519', $key->getCurve());
        $this->assertSameNL('Ed25519', $key->getPublicKey()->getCurve());

        // with public key
        $key = PublicKeyLoader::load('-----BEGIN PRIVATE KEY-----
MHICAQEwBQYDK2VwBCIEINTuctv5E1hK1bbY8fdp+K06/nwoy/HU++CXqI9EdVhC
Cm4ifubyNLrvLSOPxehdTMJODk8dhbXdEzIeV6fQKmrTkIzGwI1bCESyiS2ZhLYY
B8K7ekuDKVaSIgktpUg1pObHfvg=
-----END PRIVATE KEY-----');
        $this->assertSameNL('Ed25519', $key->getCurve());
        $this->assertSameNL('Ed25519', $key->getPublicKey()->getCurve());

        // the above key not only omits NULL - it also includes a
        // unstructuredName attribute with a value of "Curdle Chairs"
        // that the following key does not have
        $expected = '-----BEGIN PRIVATE KEY-----
MFMCAQEwBwYDK2VwBQAEIgQg1O5y2/kTWErVttjx92n4rTr+fCjL8dT74Jeoj0R1
EyJlhfYgF3xQKNJa/pGDQNgE7TlNIJnkvBZ1zRackNkJvrUa5y==
-----END PRIVATE KEY-----';
        $this->assertSameNL($expected, $key->toString('PKCS8'));

        $expected = EC::createKey('Ed25519')->toString('PKCS8');
        $key = PublicKeyLoader::load($expected);
        $this->assertSameNL('Ed25519', $key->getCurve());
        $this->assertSameNL('Ed25519', $key->getPublicKey()->getCurve());
    }

    public function testPuTTYnistp256()
    {
        $key = PublicKeyLoader::load($expected = 'PuTTY-User-Key-File-2: ecdsa-sha2-nistp256
Encryption: none
Comment: ecdsa-key-20181105
Public-Lines: 3
AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBJEXCsWA8s18
m25MJlVE1urbXPYFi4q8oMbb2H0kE2f5WPxizsKXRmb1J68paXQizryL9fC4FTqI
CJ1+UnaPfk0=
Private-Lines: 1
AAAAIQDwaPlajbXY1SxhuwsUqN1CEZ5g4adsbmJsKm+ZbUVm4g==
Private-MAC: b85ca0eb7c612df5d18af85128821bd53faaa3ef
');
        $this->assertSameNL('secp256r1', $key->getCurve());

        PuTTY::setComment('ecdsa-key-20181105');
        $this->assertSameNL($expected, $key->toString('PuTTY'));

        $key = PublicKeyLoader::load($expected = 'ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBJEXCsWA8s18m25MJlVE1urbXPYFi4q8oMbb2H0kE2f5WPxizsKXRmb1J68paXQizryL9fC4FTqICJ1+UnaPfk0= ecdsa-key-20181105');
        $this->assertSameNL('secp256r1', $key->getCurve());

        OpenSSH::setComment('ecdsa-key-20181105');
        $this->assertSameNL($expected, $key->toString('OpenSSH'));
    }

    public function testPuTTYnistp384()
    {
        $key = PublicKeyLoader::load($expected = 'PuTTY-User-Key-File-2: ecdsa-sha2-nistp384
Encryption: none
Comment: ecdsa-key-20181105
Public-Lines: 3
AAAAE2VjZHNhLXNoYTItbmlzdHAzODQAAAAIbmlzdHAzODQAAABhBOI53wHG3Cdc
AJZq5PXWZAEAxxsNVFQlQgOX9toWEOgqQF5LbK2nWLKRvaHMzocUXaTYZDccSS0A
TZFPT3j1Er1LU9cu4PHpyS07v262jdzkxIvKCPcAeISuV80MC7rHog==
Private-Lines: 2
AAAAMQCEMkGMDg6N7bUqdvLXe0YmY4qBSi8hmAuMvU38RDoVFVmV+R4RYmMueyrX
be9Oyus=
Private-MAC: 97a990a3d5f6b8f268d4be9c4ab9ebfd8fa79849
');
        $this->assertSameNL('secp384r1', $key->getCurve());

        PuTTY::setComment('ecdsa-key-20181105');
        $this->assertSameNL($expected, $key->toString('PuTTY'));

        $key = PublicKeyLoader::load($expected = 'ecdsa-sha2-nistp384 AAAAE2VjZHNhLXNoYTItbmlzdHAzODQAAAAIbmlzdHAzODQAAABhBOI53wHG3CdcAJZq5PXWZAEAxxsNVFQlQgOX9toWEOgqQF5LbK2nWLKRvaHMzocUXaTYZDccSS0ATZFPT3j1Er1LU9cu4PHpyS07v262jdzkxIvKCPcAeISuV80MC7rHog== ecdsa-key-20181105');
        $this->assertSameNL('secp384r1', $key->getCurve());

        OpenSSH::setComment('ecdsa-key-20181105');
        $this->assertSameNL($expected, $key->toString('OpenSSH'));

    }

    public function testPuTTYnistp521()
    {
        $key = PublicKeyLoader::load($expected = 'PuTTY-User-Key-File-2: ecdsa-sha2-nistp521
Encryption: none
Comment: ecdsa-key-20181105
Public-Lines: 4
AAAAE2VjZHNhLXNoYTItbmlzdHA1MjEAAAAIbmlzdHA1MjEAAACFBAF1Eg0MjaJw
ooFj6HCNh4RWbvmQRY+sdczJyBdT3EaTc/6IUcCfW7w7rAeRp2CDdE9RlAVD8IuL
qW7DJH06Xeov8wBO5G6jUqXu0rlHsOSiC6VcCxBJuWVNB1IorHnS7PX0f6HdLlIE
me73P77drqpn5YY0XLtP6hFrF7H5XfCxpNyaJA==
Private-Lines: 2
AAAAQgHJl8/dIArolFymdzhagXCfd2l8UF3CQXWGVGDQ0R04nnntlyztYiVdRXXK
r84NnzS7dJcAsR9YaUOZ69NRKNiUAQ==
Private-MAC: 6d49ce289b85549a43d74422dd8bb3ba8798c72c
');
        $this->assertSameNL('secp521r1', $key->getCurve());

        PuTTY::setComment('ecdsa-key-20181105');
        $this->assertSameNL($expected, $key->toString('PuTTY'));

        $key = PublicKeyLoader::load($expected = 'ecdsa-sha2-nistp521 AAAAE2VjZHNhLXNoYTItbmlzdHA1MjEAAAAIbmlzdHA1MjEAAACFBAF1Eg0MjaJwooFj6HCNh4RWbvmQRY+sdczJyBdT3EaTc/6IUcCfW7w7rAeRp2CDdE9RlAVD8IuLqW7DJH06Xeov8wBO5G6jUqXu0rlHsOSiC6VcCxBJuWVNB1IorHnS7PX0f6HdLlIEme73P77drqpn5YY0XLtP6hFrF7H5XfCxpNyaJA== ecdsa-key-20181105');
        $this->assertSameNL('secp521r1', $key->getCurve());

        OpenSSH::setComment('ecdsa-key-20181105');
        $this->assertSameNL($expected, $key->toString('OpenSSH'));
    }

    public function testPuTTYed25519()
    {
        $key = PublicKeyLoader::load($expected = 'PuTTY-User-Key-File-2: ssh-ed25519
Encryption: none
Comment: ed25519-key-20181105
Public-Lines: 2
AAAAC3NzaC1lZDI1NTE5AAAAIC6I6RyYAqtBcWXws9EDqGbhFtc5rKG4NMn/G7te
mQtu
Private-Lines: 1
AAAAIAHu1uI7dxFzo/SleEI2CekXKmgqlXwOgvfaRWxiX4Jd
Private-MAC: 8a06821a1c8b8b40fc40f876e543c4ea3fb81bb9
');
        $this->assertSameNL('Ed25519', $key->getCurve());

        PuTTY::setComment('ed25519-key-20181105');
        $this->assertSameNL($expected, $key->toString('PuTTY'));

        $key = PublicKeyLoader::load($expected = 'ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIC6I6RyYAqtBcWXws9EDqGbhFtc5rKG4NMn/G7temQtu ed25519-key-20181105');
        $this->assertSameNL('Ed25519', $key->getCurve());

        OpenSSH::setComment('ed25519-key-20181105');
        $this->assertSameNL($expected, $key->toString('OpenSSH'));
    }

    public function testlibsodium()
    {
        if (!function_exists('sodium_crypto_sign_keypair')) {
            self::markTestSkipped('libsodium extension is not available.');
        }

        $kp = sodium_crypto_sign_keypair();

        $key = EC::loadFormat('libsodium', $expected = sodium_crypto_sign_secretkey($kp));
        $this->assertSameNL('Ed25519', $key->getCurve());
        $this->assertSameNL($expected, $key->toString('libsodium'));

        $key = EC::loadFormat('libsodium', $expected = sodium_crypto_sign_publickey($kp));
        $this->assertSameNL('Ed25519', $key->getCurve());
        $this->assertSameNL($expected, $key->toString('libsodium'));
    }

    // ssh-keygen -t ed25519
    public function testOpenSSHPrivateKey()
    {
        $key = PublicKeyLoader::load('-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzs0dxYp
MYojSgqbCdkCSKF0PL7/ZPrx+Yypt6+norPPvpCf3+eOlgPgiDgBCa5oavhsV2aSu6QV6U
mUJexQOAe3qnguzDsgzXgWhCLpedo0IH8/hayh+fWLw1+GpTyyjyyW2+KfrmQoZsqSgD6P
QgHOMyifdBYhW41AS4KMzY4XQ6f5xyuOzTGpAc+68aSYf+G8VqA4wI7okWn442LND9aQe1
Hhex6N3NxRarqzCPezongBAXR1QGI0xQmXnpAQagoHzYWroxsVygWM==
-----END OPENSSH PRIVATE KEY-----');
        $this->assertSameNL('Ed25519', $key->getCurve());

        // testing this key is a little difficult because of this format's
        // two back to back checkint fields. both fields correspond to the
        // same randomly generated number. ostensibly this let's you verify
        // successful decryption of encrypted keys but phpseclib doesn't
        // support encrypted keys
        // none-the-less, because of the randomized component we can't easily
        // see if the key string is equal to another known string
        $key2 = PublicKeyLoader::load($key->toString('OpenSSH'));
        $this->assertSameNL('Ed25519', $key2->getCurve());
    }

    // from https://www.w3.org/TR/xmldsig-core/#sec-RFC4050Compat
    public function testXMLKey()
    {
        $key = PublicKeyLoader::load($orig = '<ECDSAKeyValue xmlns="http://www.w3.org/2001/04/xmldsig-more#">
<DomainParameters>
  <NamedCurve URN="urn:oid:1.2.840.10045.3.1.7" />
</DomainParameters>
<PublicKey>
  <X Value="58511060653801744393249179046482833320204931884267326155134056258624064349885" />
  <Y Value="102403352136827775240910267217779508359028642524881540878079119895764161434936" />
</PublicKey>
</ECDSAKeyValue>');
        $this->assertSameNL('secp256r1', $key->getCurve());

        XML::enableRFC4050Syntax();

        $dom = new DOMDocument();
        $dom->preserveWhiteSpace = false;
        $dom->loadXML($orig);
        $expected = $dom->C14N();

        //$dom = new DOMDocument();
        //$dom->preserveWhiteSpace = false;
        $dom->loadXML($key->toString('XML'));
        $actual = $dom->C14N();

        $this->assertSameNL($expected, $actual);
    }

    public function assertSameNL($expected, $actual, $message = '')
    {
        $expected = str_replace("\r\n", "\n", $expected);
        $actual = str_replace("\r\n", "\n", $actual);
        $this->assertSame($expected, $actual, $message);
    }

    public function testOpenSSHPrivateEC()
    {
        $key = '-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAaAAAABNlp1RVcK
7NOKyCvb5FY5THdFq3WLCpNH2UK8WLdvR2FtFzgocd9suNekcQeOrP+tPPiekBJDoMYCFF
ieEA8gl55vsXXRN4pDSm8PIPI1x6RFPF4emk9um+hnJPqFfWKkP50zEycocPviRX4F5w8L
weAhdpw0pjUiHhvraYsqnANrLmGnIKePnBZJtLHMHVAAbPcyIlggEkrYks4MpVXU/tuBe5
HlWPmjvSPHB7wNWA0c7qe8vCYqQpg/mnp3/kFD8HQDPW6FwODshWzO3dTGnTGwx0uBGWld
afchxbu1rz2+Zhi18JZMFw4wCLMEUBbFsw4m3ttgUiIFaK5RNECPZELM0lReI3heMmMV59
AQIDBA==
-----END OPENSSH PRIVATE KEY-----';

        $key = PublicKeyLoader::load($key);

        $key2 = PublicKeyLoader::load($key->toString('OpenSSH'));
        $this->assertInstanceOf(PrivateKey::class, $key2);

        $sig = $key->sign('zzz');

        $key = 'ecdsa-sha2-nistp256 AAAAE2VjZHNhLXNoYTItbmlzdHAyNTYAAAAIbmlzdHAyNTYAAABBBOTa1sOLJA/OWNH4ssiwzMmLCqR8eQZK1S3eQ7TQUoyhi/dEl8/qap0TcTek9xPlwGDjIH4rNixgPHC4FzTehBw= root@vagrant';
        $key = PublicKeyLoader::load($key);

        $this->assertTrue($key->verify('zzz', $sig));
    }

    public function testOpenSSHPrivateEd25519()
    {
        $key = '-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAMwAAAAtzC0uLjl
ChoqewVzLGcKrKPMfdsNVw66PFJlfuQfSCjZdZOFQ6ZZColv4IBkZCSCtjqurjvGHO7mBY
cvvjISZaH2EEfGYzooTZCmSMJkdveCuAekKo64XpwrPmwMVsusSFBWs1vSfypA4mIdurlz
xxicSzU/wV8Kf/5guNkvIJjNfBgvqApMQ5hWkR8m7SuuibnSbDaIqLDmTnsK0LH45AfDP8
iQVU39TgFnSaakbAjwZWmsVEGmsMD1NCWnmBQldRDbH=
-----END OPENSSH PRIVATE KEY-----';

        $key = PublicKeyLoader::load($key);
        $sig = $key->sign('zzz');
        $sig2 = $key->withSignatureFormat('SSH2')->sign('zzz');

        $key = 'ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIKGEJnCqQiHjcB9RE86BvJh5lEIq93iMVzIArjGaKrFD root@vagrant';
        $key = PublicKeyLoader::load($key);

        $this->assertTrue($key->verify('zzz', $sig));
        $this->assertTrue($key->withSignatureFormat('SSH2')->verify('zzz', $sig2));
    }
}
