<?php
/**
 * @author    Jim Wigginton <terrafrost@php.net>
 * @copyright 2013 Jim Wigginton
 * @license   http://www.opensource.org/licenses/mit-license.html  MIT License
 */

use phpseclib3\Crypt\PublicKeyLoader;
use phpseclib3\Crypt\DSA\PrivateKey;
use phpseclib3\Crypt\DSA\PublicKey;
use phpseclib3\Crypt\DSA\Parameters;
use phpseclib3\Crypt\DSA\Formats\Keys\PKCS1;
use phpseclib3\Crypt\DSA\Formats\Keys\PKCS8;
use phpseclib3\Crypt\DSA\Formats\Keys\PuTTY;
use phpseclib3\Math\BigInteger;
use phpseclib3\Exception\NoKeyLoadedException;

class Unit_Crypt_DSA_LoadDSAKeyTest extends PhpseclibTestCase
{
    public function testBadKey()
    {
        $this->expectException(NoKeyLoadedException::class);

        $key = 'zzzzzzzzzzzzzz';
        PublicKeyLoader::load($key);
    }

    public function testPuTTYKey()
    {
        $key = 'PuTTY-User-Key-File-2: ssh-dss
Encryption: none
Comment: dsa-key-20161223
Public-Lines: 18
AAAAB3NzaC1kc3MAAAEBAIsH1A8S7goEEneW6D/zJXh1zypZRlKlw7vNaJ7yXi9v
qUkKnTSXK9lR4/nYy4SVTcVApY0sEU2RSDridLTy40ZeMPArH+sxR7lCSZ2AuNq9
sQxu6VtYg1LKGekKYWC+r4TrZoTz2PUyrUd6sYBsYIX4ozbH2ITgYmYL9ONCrLbt
4KsKO2EUE3xN3RHv8CkAp5BraCMw5z4vC43t0bcW63RN2mEvqWOqBFx5qe7uck4j
pH7AvLyvwEye7t5KwJ8P1SMN72u4AWqyXqzLK3Ye0B7hQSFnbz0od4ps3EN+Irsu
Kf20nkGgHKYJwenpIQwHz3xhBhtgiYCdQXb3ril1kd8AAAAVAJmhajsmmqwQqNkd
k4JSJ/Y1SE11AAABADYHKKmAZ0OvCZ58m5CGPfuoX4h4nc5L0p3e0Sozcc9cJ5h3
PvD18ggAf4cLoTpxETnXTFg+30shpKbr2WsE0Jrswe6V2bMaP7Hil6q3WWahLZx6
9bEClnYCTlJkungzFjJmW+M7yd8xBHCBM6r83FKlLJjolJhHDL5DqX/uHP/9H0sl
wncwALro1jTo3JU5oRdzMuLWf9P2MB7sEsoeWk2BOiil1BtKnItuObJgXUCCuaWB
dPJyk4ZVMZZqr+vCnjocRrlQUwlMUG73Ze6KJKM1OgOKMt4PwaqD9oUaKq/gc+Eb
UtCwVR6TEeIEkmazguvbAb6dSJ18TMZ07H6DJngAAAEAYE/7gKVb8P23zYzcfYOv
3zG713/i/LAJ+dW3lQupyWkUnE8wDIht7DwKgcA/Vs73jG7SlqPjcMrNzBHjQE+y
FVCP4xO+wDsh2wZ+KVppVkMljfGpp/0mQ0mQbrmTkaCNZc0ogXwtaI4r7Su2cCq0
HMwrFJFFXXLaTZY49+lkrtP6q8WFOYJGK3WnrWvXyOe9Xnh2o8E1dQJ0RnshdOft
j1KZ4JhOHnGKoz8+sKuchGVhs5VaMbJUur3cC8INaeKvtrEpwW/Ety5iy1iPyps9
S9PlwN0KyVFGWdP2B4Gyj85IXCm3r+JNVoV5tVX9IUBTXnUor7YfWNncwWn56Lc+
RQ==
Private-Lines: 1
AAAAFFMy7BG9rPXwzqZzIY/lqsHEILNf
Private-MAC: 62b92ddd8b341b9414d640c24ba6ae929a78e039
';

        $dsa = PublicKeyLoader::load($key);

        $this->assertInstanceOf(PrivateKey::class, $dsa);
        $this->assertIsString("$dsa");
        $this->assertIsString($dsa->getPublicKey()->toString('PuTTY'));
        $this->assertIsString($dsa->getParameters()->toString('PuTTY'));

        $dsa = $dsa->withPassword('password');
        $this->assertGreaterThan(0, strlen("$dsa"));
    }

    public function testPKCS1Key()
    {
        $key = '-----BEGIN DSA PRIVATE KEY-----
MIIDPQIBAAKCAQEAiwfUDxLuCgQSd5boP/MleHXPKllGUqXDu81onvJeL2+pSQqd
xCQj4lZd+toJgqAuaZDcbSQNXPrZexj5nDkGLB1z4Mpu4qboHyuOOag04L1pTb7W
B3cwcVsb3VaIsg0LzToolymC8VUaL8gFASdVMPcAHDruIPpowJf461MuVB3nDOh8
FPqkwX8mdx/MgQFEebULrfwuvJ8SaM5OedBnBx0tWq+tE4LpKRDv9P3eeYtfycT6
wW/Utv0K7Ziwcr/Bzw9du3QogXzFTXqCJl9YPsNJhBYAipB2iEyXS17bZk6t/qsq
THPjPLPD8kmDBTNTiJmQu8JVrP4RjQFgRosd4LaAAEVzFyXcQnjnvHwBV5xYa/D2
mg78ZDsDuSoappeii2ogUL03a4DDUePIR9H0Ks1h6P4K7WrGhi1bf6J6qZP77Tah
o0CmsIceClbETHF+26lLWTAh7cwn8UPZCp9Q0CZyK8oia9m6eELzZSM46olZxEqY
pdSwzKaolcSwQ+v2FR8xtTnxm3p79FIQiDzTsZWiWd0Eqj/wSm/4i4AqCOCfbqtZ
1RIY3qz9UOACmpaAd9m5Hl8IMrdEBb0fplLb9IFYECttBeDRnBuWmcKhHEbzl5IB
lwcpq+IbnghkWKKSxVbCEk91bu0FXwc0WwviHq5VTNSl5eNcXR/bf+NStEfeqD3X
vkHZmjJRvAoQHF1qUY06Tiu43n8nPfSPZRoKcF/4yyJW9L07xXFzzZSi7An821/H
/itx+PT0jgrWFRjOzG3Ymplx8EUVrEX/Xa95St0dOgUMQZMcFCoNgH+qbRec2yL+
KDuj3yT+PRgYAyRAHJIak/1mt2bkelfQwMvRcQ8UCujCFi7X1lK0XTN5LFKYVoMz
gfFdKdL77+prbFv6c7WdYydLR5jMbUPFxnU8KeB1h0v9bUW0wPmVAlsdn8hH1kjx
jPdFfP2+pnwxVFRok2SBtVpzDY4Op1SfnUiALuRcST/kmI2MQ9RGHMQ1B6yUZz7L
yeZLgBS8l2DiE26LBqy9H+iKNSd5Zsd6GMoNJPzmD7KUHXZTHUv66cd+NKVAzQzZ
cs7d0Lhzryidz+WOqxwQR35=
-----END DSA PRIVATE KEY-----';

        $dsa = PublicKeyLoader::load($key);

        $this->assertInstanceOf(PrivateKey::class, $dsa);
        $this->assertIsString("$dsa");
        $this->assertIsString($dsa->getPublicKey()->toString('PKCS1'));
        $this->assertIsString((string) $dsa->getParameters());
    }

    public function testParameters()
    {
        $key = '-----BEGIN DSA PARAMETERS-----
MIIBHgKBgQDandMycPZNOEwDXpIDSdFODWOQVO5tlnt38wK0X33TJh4wQdqOSiVF
I+g+X8reP43ag3TEHu5bstrk6Znm7y1htTTvXQVTEwp6X3YHXbJG4Faul3g08Vud
3gzV841wToVCMUinl0EOxMYP/CO9/Kvf66KACtqWITzJYBpwAeUKfwIVAM8e3xO8
aityXVRiQRWeZtOI1yq9AoGAbmA+RzIZrtPx1mC5KzrpwgwgNHbbQBT83qeNKjEh
N+S6A47iI5OVvpxd/ZwjdXoYo7D6RxR+3LNcT64DyYrBEZuzQzHeifaO6lBvDSNf
L1cwyXx0KMaaampd34MzOIHbC44SHY+cE3aVVUsnmt6Ur1nQaVYVszl+AO6m8bPm
4Vg=
-----END DSA PARAMETERS-----';
        $key = str_replace(["\n", "\r"], '', $key);
        $dsa = PublicKeyLoader::load($key);

        $this->assertInstanceOf(Parameters::class, $dsa);
        $this->assertSame($key, str_replace(["\n", "\r"], '', "$dsa"));
        $this->assertSame($key, str_replace(["\n", "\r"], '', (string) $dsa->getParameters()));
    }

    public function testPKCS8Public()
    {
        $key = '-----BEGIN PUBLIC KEY-----
MIIBtjCCASsGByqGSM44BAEwggEeAoGBANqd0zJw9k04TANekgNJ0U4NY5BU7m2W
e3fzArRffdMmHjBB2o5KJUUj6D5fyt4/jdqDdMQe7luy2uTpmebvLWG1NO9dBVMT
CnpfdgddskbgVq6XeDTxW53eDNXzjXBOhUIxSKeXQQ7Exg/8I738q9/rooAK2pYh
PMlgGnAB5Qp/AhUAzx7fE7xqK3JdVGJBFZ5m04jXKr0CgYBuYD5HMhmu0/HWYLkr
OunCDCA0dttAFPzep40qMSE35LoDjuIjk5W+nF39nCN1ehijsPpHFH7cs1xPrgPJ
isERm7NDMd6J9o7qUG8NI18vVzDJfHQoxppqal3fgzM4gdsLjhIdj5wTdpVVSyea
3pSvWdBpVhWzOX4A7qbxs+bhWAOBhAACgYBTpSKcKoVKw+hglVClqvqQdNKGC4a+
XC4lOh2221ZrTgy/sN92vT7cdBn4ydHoth6/bD236L/FYfiX4S4mOczPhrv/l/2u
ZpmyOpXM/0opRMIRdmqVW4ardBFNokmlqngwcbaptfRnk9W2cQtx0lmKy6X/vnis
3AElwP86TYgBhw==
-----END PUBLIC KEY-----';

        $dsa = PublicKeyLoader::load($key);

        $this->assertInstanceOf(PublicKey::class, $dsa);
        $this->assertIsString("$dsa");
    }

    public function testPKCS8Private()
    {
        $key = '-----BEGIN PRIVATE KEY-----
MIIBSgIBADCCASsGByqGSM44BAEwggEeAoGBANqd0zJw9k04TANekgNJ0U4NY5BU
2Y0lm1aTnxKobmixgymR2j1LjGpu2U9VAo2/KMOLyLht3MgI5tDFgGSRWqT7Fm6g
vlyyGKFIQlvHkQxpWw6kKvXhq65RvtyhhjgKqrOAgCAfwp4ERD/1f021S7/mohTa
0zjfyyVaixkT5xH/JElXjt0BN2QeY0LtFdZrXu0A05Laof1gDirERM2LToRa9/mP
JAAbKxJAplN3fLAYtvjLP41GGrZ97ubPUqfaE5g+mu09yeg9eSUYAyuTmV9EQ4dk
KFbSbkLkL6OPRG4U4b7tpp9ps07OyKfVIsWgmKDWsV8sBfM6TPRkktlUq2qhwkSj
NEum4rHqpxHoZyuQqz0x2fmMU+JYigBZvGtDO9dHrzUz1XhTFv9PM0SLi9sFqy==
-----END PRIVATE KEY-----';

        $dsa = PublicKeyLoader::load($key);

        $this->assertInstanceOf(PrivateKey::class, $dsa);
        $this->assertIsString("$dsa");
        $this->assertInstanceOf(PublicKey::class, $dsa->getPublicKey());
        $this->assertInstanceOf(Parameters::class, $dsa->getParameters());
    }

    public function testPuTTYBadMAC()
    {
        $this->expectException(NoKeyLoadedException::class);

        $key = 'PuTTY-User-Key-File-2: ssh-dss
Encryption: none
Comment: dsa-key-20161223
Public-Lines: 18
AAAAB3NzaC1kc3MAAAEBAIsH1A8S7goEEneW6D/zJXh1zypZRlKlw7vNaJ7yXi9v
qUkKnTSXK9lR4/nYy4SVTcVApY0sEU2RSDridLTy40ZeMPArH+sxR7lCSZ2AuNq9
sQxu6VtYg1LKGekKYWC+r4TrZoTz2PUyrUd6sYBsYIX4ozbH2ITgYmYL9ONCrLbt
4KsKO2EUE3xN3RHv8CkAp5BraCMw5z4vC43t0bcW63RN2mEvqWOqBFx5qe7uck4j
pH7AvLyvwEye7t5KwJ8P1SMN72u4AWqyXqzLK3Ye0B7hQSFnbz0od4ps3EN+Irsu
Kf20nkGgHKYJwenpIQwHz3xhBhtgiYCdQXb3ril1kd8AAAAVAJmhajsmmqwQqNkd
k4JSJ/Y1SE11AAABADYHKKmAZ0OvCZ58m5CGPfuoX4h4nc5L0p3e0Sozcc9cJ5h3
PvD18ggAf4cLoTpxETnXTFg+30shpKbr2WsE0Jrswe6V2bMaP7Hil6q3WWahLZx6
9bEClnYCTlJkungzFjJmW+M7yd8xBHCBM6r83FKlLJjolJhHDL5DqX/uHP/9H0sl
wncwALro1jTo3JU5oRdzMuLWf9P2MB7sEsoeWk2BOiil1BtKnItuObJgXUCCuaWB
dPJyk4ZVMZZqr+vCnjocRrlQUwlMUG73Ze6KJKM1OgOKMt4PwaqD9oUaKq/gc+Eb
UtCwVR6TEeIEkmazguvbAb6dSJ18TMZ07H6DJngAAAEAYE/7gKVb8P23zYzcfYOv
3zG713/i/LAJ+dW3lQupyWkUnE8wDIht7DwKgcA/Vs73jG7SlqPjcMrNzBHjQE+y
FVCP4xO+wDsh2wZ+KVppVkMljfGpp/0mQ0mQbrmTkaCNZc0ogXwtaI4r7Su2cCq0
HMwrFJFFXXLaTZY49+lkrtP6q8WFOYJGK3WnrWvXyOe9Xnh2o8E1dQJ0RnshdOft
j1KZ4JhOHnGKoz8+sKuchGVhs5VaMbJUur3cC8INaeKvtrEpwW/Ety5iy1iPyps9
S9PlwN0KyVFGWdP2B4Gyj85IXCm3r+JNVoV5tVX9IUBTXnUor7YfWNncwWn56Lc+
RQ==
Private-Lines: 1
AAAAFFMy7BG9rPXwzqZzIY/lqsHEILNf
Private-MAC: aaaaaadd8b341b9414d640c24ba6ae929a78e039
';

        PublicKeyLoader::load($key);
    }

    public function testXML()
    {
        $key = '-----BEGIN PUBLIC KEY-----
MIIBtjCCASsGByqGSM44BAEwggEeAoGBANqd0zJw9k04TANekgNJ0U4NY5BU7m2W
e3fzArRffdMmHjBB2o5KJUUj6D5fyt4/jdqDdMQe7luy2uTpmebvLWG1NO9dBVMT
CnpfdgddskbgVq6XeDTxW53eDNXzjXBOhUIxSKeXQQ7Exg/8I738q9/rooAK2pYh
PMlgGnAB5Qp/AhUAzx7fE7xqK3JdVGJBFZ5m04jXKr0CgYBuYD5HMhmu0/HWYLkr
OunCDCA0dttAFPzep40qMSE35LoDjuIjk5W+nF39nCN1ehijsPpHFH7cs1xPrgPJ
isERm7NDMd6J9o7qUG8NI18vVzDJfHQoxppqal3fgzM4gdsLjhIdj5wTdpVVSyea
3pSvWdBpVhWzOX4A7qbxs+bhWAOBhAACgYBTpSKcKoVKw+hglVClqvqQdNKGC4a+
XC4lOh2221ZrTgy/sN92vT7cdBn4ydHoth6/bD236L/FYfiX4S4mOczPhrv/l/2u
ZpmyOpXM/0opRMIRdmqVW4ardBFNokmlqngwcbaptfRnk9W2cQtx0lmKy6X/vnis
3AElwP86TYgBhw==
-----END PUBLIC KEY-----';

        $dsa = PublicKeyLoader::load($key);
        $xml = $dsa->toString('XML');
        $this->assertStringContainsString('DSAKeyValue', $xml);

        $dsa = PublicKeyLoader::load($xml);
        $pkcs8 = $dsa->toString('PKCS8');

        $this->assertSame(
            strtolower(preg_replace('#\s#', '', $pkcs8)),
            strtolower(preg_replace('#\s#', '', $key))
        );
    }

    public function testOpenSSHPrivate()
    {
        $key = '-----BEGIN OPENSSH PRIVATE KEY-----
b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAABswAAAAdzb2TrJz
XAKIPGKkYoH7/43m4PJYxVVvDQWmNjk1fwYUWX9Bi6/ctILOTrJ5II2QjqciYxK16cP5o4
cKNB7xHgLgeTMuDu/wyAYzUR/dC/Q+P8Ic36Ai3sqslq5LjeBOJn2TUjowon6lEq0Gk8rW
w7bsbMd253GoVbU5JNqIDnm7Z5us2FWlX+nFDIcCMINile28LCRTSz8kvDrTb2m8MZpeDV
+3WNfnmHZSpW5XfpJTFG1Z5PG1K3oe7vvrOti417xdH+okKjGZSKl0MsKZUo1/Riu810Kx
gpoKRiS+owR3a5U1mKXYvfjF3xuA4oNK645Tls/GpF5lmqOS3kViloHkLqrkzwL62dLEoi
3oNrWtHj4NNY7ITAuXlfviABF3okskVDMZsaDtJrvczvgXFs5LPdOMlXm+geNeezaOTy3M
WZAT8iGJ2bdsZhGFvSkCxT2Mka5dwMqwabID7sVPn/+oXrQCTDZqo1+m9pRvHSA9Y2+mob
6RperP4OP7RW/xvGV21/VbX1irNmOiziLbbvM43swJWNpcFo3lm/5vPgpNLeLnpky3TZa4
JD6Nc23tlOAcYgiMMMvsayYKsLJvpii5vIuZEyxoXakRfTI0/88m9UbUdSpdNPqNpxp4mp
CxzH3HM2/tmUPZxKc3os9kQqXIkxk08vr7k7bMhO4gXEOVVSJvBm/BigsWuK/JN/S+F6Pk
41Qo9jUFBJ1MwvJJpa9CPgPwMl1qSy4SI3FuK9Hqoab203OgaWI1YKKhxkW3J6EL8fOkJ+
MvDMtbyuuGcg76TvLcms4XePTQC0b7qRdJoW+1SGSuKwmPCa7SfFYUNn1M7bw6A6ol3PpC
PDy445Dmg+Mtwhcqntr6uXQFIL7/dBX013qGndgMGFK+lgo7i8e7rlSKtndG8xrb0fIl14
0Jzn/PzO8YjZvr8mhDyNusumceXwW43yMtEp5BBnfBlN7qxD3zaRyvBTyWnlT1iLwZqrlg
fgWROUbPsHqmSO7jOKVgMUr+NGJpgaEffM4HQeQO5ZxK4bSeSaKHCHNaJs7sqc5AbkJhCq
Sk5Geib/+mvvfDySXVv5+B7RyiJBU4S9+CBN2jznoY5VM8OC/ZxlE26/tHF4igHAJjTlfT
kYm24NsGtoeeIu6iN/4AOSrgWeQIZMx0kNR1TT5dQ47oskTNZBQrhJYpHl7O9fAmBPkm3X
1UXgYqOZGlvNL166xJJln3xqlWYCVazIBi==
-----END OPENSSH PRIVATE KEY-----';

        $key = PublicKeyLoader::load($key);

        $key2 = PublicKeyLoader::load($key->toString('OpenSSH'));
        $this->assertInstanceOf(PrivateKey::class, $key2);

        $sig = $key->sign('zzz');

        $key = 'ssh-dss AAAAB3NzaC1kc3MAAACBAOkTX/vVXq65p4SptoDOgSwDWR0kFmH0iff9WUJfC+UnrPPwpLNtDGsLjmw7vL18zAPlDJYSmFV4hc781J9xmEr9IP8P43VnHnYq7YoQVNvd0CF80Z73KGysBUjlbITiVLap57l8s2eHXfyFgB03sQ1GhQBDcjoeT0cAif6W0cyZAAAAFQD3mI93EcveKSSZl3VflaFAVMj72wAAAIEAxyfU91Z1SYTwPgo/qDSurgoASVlz3TiZhf5sdA0xc42bgGuxJyff+I5/fvQRqJSCSiaj4JI/YfxLeksE1uJBbsDNL1MZPvzYWUb9td/yRIwVvsjMC4WZAx6i4tmjjQDAcS/s+l3OtezoZFDS/AmU2iqhQBSL0wyF7EPOkYxs2OoAAACBAJTUm4lkPOmT5wGAR1FNEbHqYlmQ/28c/ty1SGAplwA9LJ3yR43yBAKm+Zm4rUAO13/6lcHFtAEA9vX6nzACCoNTibz4sV/aKAq1vyZ/yqn+CHFXTz9PN/lZ9wNY/E56x2U3jyottKZSCw7uCr/W+h15sxIdx25XglOryLHkTLzt root@vagrant';
        $key = PublicKeyLoader::load($key);

        $this->assertTrue($key->verify('zzz', $sig));
    }
}
