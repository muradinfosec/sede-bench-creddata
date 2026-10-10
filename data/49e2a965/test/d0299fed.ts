import SonicBoom from "sonic-boom";
import { createProbot, Probot } from "../src";

const env = {
  APP_ID: "1",
  PRIVATE_KEY: `-----BEGIN RSA PRIVATE KEY-----
MIIBOQIBAAJBAIILhiN9IFpaE0pUXsesuuoaj6eeDiAqCiE49WB1tMB8ZMhC37kY
ex86ITQbewy1ghs4EW2e7OvF9jF68b68ezOHRRHnMBGza74zvYfoKumUVnp80iVG
/MyqKSu7d4qXZwl6N1OB8Aqd3B846sSlbU0BBrMYxBgW80bwyo+mmq9nuTV4W5lO
DRbbEy1h/ioyWvodkaXBi1RvM3qqaidm/y2tQdBx/7kibypRQnSvlXkv0vbohmwt
T3CfED0V8Lb+oAgW+gxgbibPNYivN2vHJhHGYa0ehU2kos6lMxPUQHM5yLXVgdSt
SjPGJyfcvUmEMKQ7sksnnizG9DpJsTnNLRuoMHI8N44eoCTdOAEASOnFiULjlZ2K
x1fcnfGNXps2jH3YFrhkU5ypdfsibPII5dE67To=
-----END RSA PRIVATE KEY-----`,
  WEBHOOK_SECRET: "secret",
};
// tslint:disable:no-empty
describe("createProbot", () => {
  test("createProbot()", () => {
    const probot = createProbot({ env });
    expect(probot).toBeInstanceOf(Probot);
  });

  test("defaults, env", () => {
    const probot = createProbot({
      env: {
        ...env,
        LOG_LEVEL: "debug",
      },
      defaults: { logLevel: "trace" },
    });
    expect(probot.log.level).toEqual("debug");
  });

  test("defaults, overrides", () => {
    const probot = createProbot({
      env,
      defaults: { logLevel: "debug" },
      overrides: { logLevel: "trace" },
    });
    expect(probot.log.level).toEqual("trace");
  });

  test("env, overrides", () => {
    const probot = createProbot({
      env: {
        ...env,
        LOG_LEVEL: "fatal",
      },
      overrides: { logLevel: "trace" },
    });
    expect(probot.log.level).toEqual("trace");
  });

  test("defaults, env, overrides", () => {
    const probot = createProbot({
      env: {
        ...env,
        LOG_LEVEL: "fatal",
      },
      defaults: { logLevel: "debug" },
      overrides: { logLevel: "trace" },
    });
    expect(probot.log.level).toEqual("trace");
  });

  test("env, logger message key", () => {
    let outputData = "";

    const sbWrite = SonicBoom.prototype.write;
    SonicBoom.prototype.write = function (data) {
      outputData += data;
    };

    const probot = createProbot({
      env: {
        ...env,
        LOG_LEVEL: "info",
        LOG_FORMAT: "json",
        LOG_MESSAGE_KEY: "myMessage",
      },
      defaults: { logLevel: "trace" },
    });

    probot.log.info("Ciao");

    try {
      expect(JSON.parse(outputData).myMessage).toEqual("Ciao");
    } finally {
      SonicBoom.prototype.write = sbWrite;
    }
  });
});
