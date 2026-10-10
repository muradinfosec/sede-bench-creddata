<?php

include "vendor/autoload.php";

$mailgun = \Mailgun\Mailgun::create('key-b19202952eee8ccbe95b7106b748c76f');
$bounces = $mailgun->suppressions()->bounces()->index('sandbox40c4b0745b8a4461959950ccf9c0a1ea.mailgun.org');
$x = 2;
