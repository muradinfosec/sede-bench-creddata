---
title: 'Auth Session'
apiurl: '/api/v1/user/auth_session'
type: 'GET'
layout: default
---

透过Session检查去判定使用者可否存取资源

### RequestHeader
透过RequestHeader 的 Apitoken做验证
```"RequestHeader": {
  "Apitoken": "{\"euzf\":\"gpnc\",\"azi\":\"815x3269p37022b30zvu3755wz787120\"}",
  "X-Forwarded-For": " 127.0.0.1"
}```

### Response

Session 为有效
```Status: 200```
```{"message":"session is vaild!"}```

For errors responses, see the [response status codes documentation](#/response-status-codes).
