<?php
/**
 * FecShop file.
 *
 * @link http://www.fecshop.com/
 * @copyright Copyright (c) 2016 FecShop Software LLC
 * @license http://www.fecshop.com/license/
 */

namespace fecshop\app\appserver\modules\Customer\controllers;

use fecshop\app\appserver\modules\AppserverController;
use Yii;
use \Firebase\JWT\JWT;

/**
 * @author Terry Zhao <2358269014@qq.com>
 * @since 1.0
 */
class JwttestController extends AppserverController
{
    
    public function init()
    {
        parent::init();
    }
    
    public function actionIndex(){
        if(Yii::$app->request->getMethod() === 'OPTIONS'){
            return [];
        }
        $privateKey = '
-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQC8kGa1pSjbSYZVebtTRBLxBz5H4i2p/llLCrEeQhta5kaQu/Rn
iSCH3R6NGG7+8GvTeU8DruFGstiCMYZxzD+46y9o3zAbTkH4XuA+DphjBfylcLC6
1+TWjGOWEAqMnJ4F3lrWlpsidUm2Yf06BqatEId7VfpN/imyq1hLu5tglmYrMDaE
yYSYZ/Phi22NzLVTpwMI1OmmmAyWnl/WYBarKCjnfjOacbawEMH0//3bYpF2pnwF
ZHC/1/HMTzaddFwZy2mbEMa05ru67CWDUnEs16IyJfwQo+4iRzS/IftRyy+OyO4a
GlH1NZre5VaNXXErfj+QmCJpIGa2dKYaTiAqdP6P7BFk0K0BVhJOzGuoXB27AhG9
yr02HjQFKqBxb4XcpvU4R+qHBS2zBvarTmvRCv22jmB9B+HymQiwvvmC+U4rdVO5
7zQke3XfWSqPUmkZly0OsDzj/K6DQW6k9ZK0vLk15lvWhX/x29UUy3h3JaYxNvwS
PAEZL1wm4WenB0R7Jw1MgjBU+m5zyZXtebpKjd2utWK90px1hwRR9sv56hqxzdEH
b78ZvqotxdbsFvKvBJTu5fdloKwsbEbyTRFE8ruj5bzcWI8ZHX1tBDo/6ZiSurET
tvpFxBaS2xCORpNoKeLuXB0yvUnFGfkLG0VnuAEEm978YYEG2oPBVlnXQm1ug21x
MWy9qShR3hUTIDcThzVMXwNk1FVPD8toSvyxOWgXe10oO+F1ZzwxfkbyhSjrLLUr
k6tabAfwFRCsPaxYarFSNiT5ImgI+yKVffEQt8Rquny=
-----END RSA PRIVATE KEY-----
';

$publicKey = '
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC8kGa1pSjbSYZVebtTRBLxBz5H
4i2p/llLCrEeQhta5kaQu/RnvuER4W8oDH3+3iuIYW4VQAzyqFpwuzjkDI+17t5t
0tyazyZ8JXw+KgXTxldMPEL95+qVhgXvwtihXC1c5oGbRlEDvDF6Sa53rcFVsYJ4
ehde/zUxo6UvS7UrBQIDAQAB
-----END PUBLIC KEY-----
';

        $token = array(
            "iss" => "example.org",
            "aud" => "example.com",
            "iat" => 1356999524,
            "nbf" => 1357000000
        );

        $jwt = JWT::encode($token, $privateKey, 'RS256');
        echo "Encode:\n" . print_r($jwt, true) . "\n";

        $decoded = JWT::decode($jwt, $publicKey, array('RS256'));

        /*
         NOTE: This will now be an object instead of an associative array. To get
         an associative array, you will need to cast it as such:
        */

        $decoded_array = (array) $decoded;
        echo "Decode:\n" . print_r($decoded_array, true) . "\n";
    }
    
    
    
}