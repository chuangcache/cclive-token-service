# cclive-token-service go

- 集成方法
  - go get github.com/chuangcache/cclive-token-service
  - 在你的项目里引用 **rtctoken** 包，应该会自动添加引用 **"github.com/chuangcache/cclive-token-service/go/src/rtctoken"**
  - 传入appId和appCertificate，通过CreateToken生成客户端SDK认证使用的token。**token, err = rtctoken.CreateToken(appId, appCertificate)**