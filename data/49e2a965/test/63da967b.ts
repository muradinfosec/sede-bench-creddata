import { createServer, IncomingMessage, ServerResponse } from "http";
import Stream from "stream";

import pino from "pino";
import getPort from "get-port";
import got from "got";
import { sign } from "@octokit/webhooks";

import { createNodeMiddleware, createProbot, Probot } from "../src";
import { ApplicationFunction } from "../src/types";

const APP_ID = "1";
const PRIVATE_KEY = `-----BEGIN RSA PRIVATE KEY-----
MIIBOQIBAAJBAIILhiN9IFpaE0pUXsesuuoaj6eeDiAqCiE49WB1tMB8ZMhC37kY
fJ31nsPoXHo5WEX7Zh7G2DfL7Qx49z35mmGPPrnnOMyXt91NlLMbRYnagFI11Yfb
/HvgBzg9S5yXdfw9X9ST3lsn5l271NgQCy1fmmIzYxAl96wgXL+bsF6FgYN6j7yx
qxwQnA3f/nIKTmlxUrTrN3KhT6mTgVNJ/o3vtZQR/6setnObtBeTcCyy1mNfFCfq
j4iWrN8z2Zr+iQYi+oGblSXJDYcNR5JXBxUkiL7CsK3Hqa4UwVshkBU2kzsTYNAA
fJYtJooAxfvLHwS3HVZLsBWJ1gbrJjOcbShqosA8q48QLXuBcDYDqCGWjpyBep3A
i5OXqtKKnXe6Bp9HJbQsV0euIPrwHKUa0FP32Ix=
-----END RSA PRIVATE KEY-----`;
const WEBHOOK_SECRET = "secret";
const pushEvent = require("./fixtures/webhook/push.json");

describe("createNodeMiddleware", () => {
  let output: any[];
  const streamLogsToOutput = new Stream.Writable({ objectMode: true });
  streamLogsToOutput._write = (object, encoding, done) => {
    output.push(JSON.parse(object));
    done();
  };

  beforeEach(() => {
    output = [];
  });

  test("with createProbot", async () => {
    expect.assertions(1);

    const app: ApplicationFunction = (app) => {
      app.on("push", (event) => {
        expect(event.name).toEqual("push");
      });
    };
    const middleware = createNodeMiddleware(app, {
      probot: createProbot({
        overrides: {
          log: pino(streamLogsToOutput),
        },
        env: {
          APP_ID,
          PRIVATE_KEY,
          WEBHOOK_SECRET,
        },
      }),
    });

    const server = createServer(middleware);
    const port = await getPort();
    server.listen(port);

    const body = JSON.stringify(pushEvent);

    await got.post(`http://127.0.0.1:${port}`, {
      headers: {
        "content-type": "application/json",
        "x-github-event": "push",
        "x-github-delivery": "1",
        "x-hub-signature": sign("secret", body),
      },
      body,
    });

    await new Promise((resolve) => setTimeout(resolve, 100));

    server.close();
  });

  test("loads app only once", async () => {
    let counter = 0;
    const appFn = () => {
      counter++;
    };
    const middleware = createNodeMiddleware(appFn, {
      probot: new Probot({
        appId: APP_ID,
        privateKey: PRIVATE_KEY,
      }),
    });

    middleware({} as IncomingMessage, { end() {} } as ServerResponse);
    middleware({} as IncomingMessage, { end() {} } as ServerResponse);

    expect(counter).toEqual(1);
  });
});
