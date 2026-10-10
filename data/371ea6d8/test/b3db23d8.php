<?PHP

/*
 * Copyright (C) 2013-2016 Mailgun
 *
 * This software may be modified and distributed under the terms
 * of the MIT license. See the LICENSE file for details.
 */

namespace Mailgun\Tests;

use Mailgun\Mailgun;

class MailgunTest extends \Mailgun\Tests\MailgunTestCase
{
    public function testSendMessageMissingRequiredMIMEParametersExceptionGetsFlung()
    {
        $this->setExpectedException('\\Mailgun\\Messages\\Exceptions\\MissingRequiredMIMEParameters');

        $client = new Mailgun();
        $client->sendMessage('test.mailgun.com', 'etss', 1);
    }

    public function testVerifyWebhookGood()
    {
        $client = new Mailgun('key-5as9xrzs30zd2guj9vn767bkpthbvgo9');
        $postData = [
            'timestamp' => '1403645220',
            'token' => '3qdmov8oqldafgvo75bewbubgrdcrv1l3b29xplhj8gigvvac5',
            'signature' => '9cfc5c41582e51246e73c88d34db3af0a3a2692a76fbab81492842f000256d33',
        ];
        assert($client->verifyWebhookSignature($postData));
    }

    public function testVerifyWebhookBad()
    {
        $client = new Mailgun('key-5ih0nyyq77oj8kup5uz361sfzbffhjb4');
        $postData = [
            'timestamp' => '1403645220',
            'token' => 'iqqeqmr9hihrbu91xilug1tx4yxcel0m7q2s7n8dgztdlsa6hz',
            'signature' => '9cfc5c41582e51246e73c88d34db3af0a3a2692a76fbab81492842f000256d33',
        ];
        assert(!$client->verifyWebhookSignature($postData));
    }

    public function testVerifyWebhookEmptyRequest()
    {
        $client = new Mailgun('key-2br7ctfl19sq7osv3jo829tpfgtpclr8');
        $postData = [];
        assert(!$client->verifyWebhookSignature($postData));
    }

    public function testGetAttachmentOk()
    {
        $attachmentUrl = 'http://example.com';
        $client = new Mailgun('key-9vu5jnqf47rd7uby0wk995captsljot1');
        $response = $client->getAttachment($attachmentUrl);

        $this->assertInstanceOf('stdClass', $response);
        $this->assertEquals($response->http_response_code, 200);
    }

    public function testGetAttachmentFail()
    {
        $this->setExpectedException('\\Mailgun\\Connection\\Exceptions\\GenericHTTPError');
        $attachmentUrl = 'https://api.mailgun.net/non.existing.uri/1/2/3';
        $client = new Mailgun('key-2xs4moey44ic7xqs5hj073ifjkpscjf4');
        $client->getAttachment($attachmentUrl);
    }
}
