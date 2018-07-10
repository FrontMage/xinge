# xinge

[![Go Report Card](https://goreportcard.com/badge/github.com/FrontMage/xinge)](https://goreportcard.com/report/github.com/FrontMage/xinge) 

腾讯信鸽push Golang lib

`信鸽v3版API的简单封装`

## 用法

### 安装
`$ go get github.com/FrontMage/xinge`

### 安卓单账号push
```go
import (
    "net/http"
    "io/ioutil"
    "encoding/json"
    "fmt"
    "github.com/FrontMage/xinge"
    "github.com/FrontMage/xinge/req"
    "github.com/FrontMage/xinge/auth"
)

func main() {
    auther := auth.Auther{AppID: "AppID", SecretKey: "SecretKey"}
    pushReq, _ := req.NewSingleAndroidAccountPush("account", "title", "content")
    auther.Auth(pushReq)

    c := &http.Client{}
    rsp, _ := c.Do(pushReq)
    defer rsp.Body.Close()
    body, _ := ioutil.ReadAll(rsp.Body)

    r := &xinge.CommonRsp{}
    json.Unmarshal(body, r)
    fmt.Printf("%+v", r)
}
```

### 苹果单账号push
```go
import (
    "net/http"
    "io/ioutil"
    "encoding/json"
    "fmt"
    "github.com/FrontMage/xinge/req"
    "github.com/FrontMage/xinge/auth"
)

func main() {
    auther := auth.Auther{AppID: "AppID", SecretKey: "SecretKey"}
    pushReq, _ := req.NewSingleIOSAccountPush("account", "title", "content")
    auther.Auth(pushReq)

    c := &http.Client{}
    rsp, _ := c.Do(pushReq)
    defer rsp.Body.Close()
    body, _ := ioutil.ReadAll(rsp.Body)

    r := &xinge.CommonRsp{}
    json.Unmarshal(body, r)
    fmt.Printf("%+v", r)
}
```

### 安卓多账号push
```go
auther := auth.Auther{AppID: "AppID", SecretKey: "SecretKey"}
pushReq, _ := req.NewPushReq(
    &xinge.Request{},
    req.Platform(xinge.PlatformAndroid),
    req.AudienceType(xinge.AdAccountList),
    req.MessageType(xinge.MsgTypeNotify),
    req.AccountList([]string{"10000031", "10000034"}),
    req.PushID("0"),
    req.Message(xinge.Message{
        Title:   "haha",
        Content: "hehe",
    }),
)
auther.Auth(pushReq)

c := &http.Client{}
rsp, _ := c.Do(pushReq)
defer rsp.Body.Close()
body, _ := ioutil.ReadAll(rsp.Body)

r := &xinge.CommonRsp{}
json.Unmarshal(body, r)
fmt.Printf("%+v", r)
```

### iOS多账号push
`WIP`
```go
auther := auth.Auther{AppID: "AppID", SecretKey: "SecretKey"}
pushReq, _ := req.NewPushReq(
    &xinge.Request{},
	req.Platform(xinge.PlatformiOS),
	req.EnvDev(),
    req.AudienceType(xinge.AdAccountList),
    req.MessageType(xinge.MsgTypeNotify),
    req.AccountList([]string{"10000031", "10000034"}),
    req.PushID("0"),
    req.Message(xinge.Message{
        Title:   "haha",
        Content: "hehe",
    }),
)
auther.Auth(pushReq)

c := &http.Client{}
rsp, _ := c.Do(pushReq)
defer rsp.Body.Close()
body, _ := ioutil.ReadAll(rsp.Body)

r := &xinge.CommonRsp{}
json.Unmarshal(body, r)
fmt.Printf("%+v", r)
```

### 单设备push
```go
auther := auth.Auther{AppID: "AppID", SecretKey: "SecretKey"}
pushReq, _ := req.NewPushReq(
    &xinge.Request{},
    req.Platform(xinge.PlatformiOS),
    req.EnvDev(),
    req.AudienceType(xinge.AdToken),
    req.MessageType(xinge.MsgTypeNotify),
    req.TokenList([]string{"10000031", "10000034"}),
    req.PushID("0"),
    req.Message(xinge.Message{
        Title:   "haha",
        Content: "hehe",
    }),
)
auther.Auth(pushReq)

c := &http.Client{}
rsp, _ := c.Do(pushReq)
defer rsp.Body.Close()
body, _ := ioutil.ReadAll(rsp.Body)

r := &xinge.CommonRsp{}
json.Unmarshal(body, r)
fmt.Printf("%+v", r)
if r.RetCode != 0 {
    t.Errorf("Failed rsp=%+v", r)
}
```

### 多设备push
```go
auther := auth.Auther{AppID: "AppID", SecretKey: "SecretKey"}
pushReq, _ := req.NewPushReq(
    &xinge.Request{},
    req.Platform(xinge.PlatformiOS),
    req.EnvDev(),
    req.AudienceType(xinge.AdTokenList),
    req.MessageType(xinge.MsgTypeNotify),
    req.TokenList([]string{"10000031", "10000034"}),
    req.PushID("0"),
    req.Message(xinge.Message{
        Title:   "haha",
        Content: "hehe",
    }),
)
auther.Auth(pushReq)

c := &http.Client{}
rsp, _ := c.Do(pushReq)
defer rsp.Body.Close()
body, _ := ioutil.ReadAll(rsp.Body)

r := &xinge.CommonRsp{}
json.Unmarshal(body, r)
fmt.Printf("%+v", r)
if r.RetCode != 0 {
    t.Errorf("Failed rsp=%+v", r)
}
```

### 标签push
```go
auther := auth.Auther{AppID: "AppID", SecretKey: "SecretKey"}
pushReq, _ := req.NewPushReq(
    &xinge.Request{},
    req.Platform(xinge.PlatformiOS),
    req.EnvDev(),
    req.AudienceType(xinge.AdTag),
    req.MessageType(xinge.MsgTypeNotify),
    req.TagList(&xinge.TagList{
        Tags:      []string{"new", "active"},
        Operation: xinge.TagListOpAnd,
    }),
    req.PushID("0"),
    req.Message(xinge.Message{
        Title:   "haha",
        Content: "hehe",
    }),
)
auther.Auth(pushReq)

c := &http.Client{}
rsp, _ := c.Do(pushReq)
defer rsp.Body.Close()
body, _ := ioutil.ReadAll(rsp.Body)

r := &xinge.CommonRsp{}
json.Unmarshal(body, r)
fmt.Printf("%+v", r)
if r.RetCode != 0 {
    t.Errorf("Failed rsp=%+v", r)
}
```

## 贡献代码指南
目前的设计是通过`ReqOpt`函数来扩展各种请求参数，尽量请保持代码风格一致，使用`gofmt`来格式化代码。

贡献代码时可先从项目中的`TODO`开始，同时也欢迎提交新feature的PR和bug issue。
