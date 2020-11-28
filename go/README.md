# cclive-token-service go

- 集成方法
  - go get github.com/chuangcache/cclive-token-service
  - 在你的项目里引用 **rtctoken** 包，应该会自动添加引用 **"github.com/chuangcache/cclive-token-service/go/src/rtctoken"**
  - 传入你的appId，以及通过后台生成的token，并指定一个房间角色。**token, err = rtctoken.CreateToken(appId, appToken, role)**，即生成客户端SDK用token