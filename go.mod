module xjosiah.com/go-gin

go 1.15

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/astaxie/beego v1.12.3
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.6.3
	github.com/go-openapi/spec v0.20.3 // indirect
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/google/go-cmp v0.5.4 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/kr/pty v1.1.5 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/lib/pq v1.2.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/satori/go.uuid v1.2.0 // indirect
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	github.com/smartystreets/goconvey v0.0.0-20190731233626-505e41936337 // indirect
	github.com/stretchr/objx v0.2.0 // indirect
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.7.0
	github.com/ugorji/go v1.2.4 // indirect
	github.com/unknwon/com v1.0.1
	github.com/urfave/cli v1.22.5 // indirect
	github.com/urfave/cli/v2 v2.3.0 // indirect
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad // indirect
	golang.org/x/sys v0.0.0-20210124154548-22da62e12c0c // indirect
	golang.org/x/tools v0.1.0 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/ini.v1 v1.62.0
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace (
	xjosiah.com/go-gin/conf => D:\go-gin/pkg/conf
	xjosiah.com/go-gin/godoc => D:\go-gin/godoc
	xjosiah.com/go-gin/middleware => D:\go-gin/middleware
	xjosiah.com/go-gin/middleware/jwt => D:\go-gin/middleware/jwt
	xjosiah.com/go-gin/models => D:\go-gin/models
	xjosiah.com/go-gin/pkg/e => D:\go-gin/pkg/e
	xjosiah.com/go-gin/pkg/logging => D:\go-gin/pkg/logging
	xjosiah.com/go-gin/pkg/setting => D:\go-gin/pkg/setting
	xjosiah.com/go-gin/pkg/upload => D:\go-gin/pkg/upload
	xjosiah.com/go-gin/pkg/util => D:\go-gin/pkg/util
	xjosiah.com/go-gin/routers => D:\go-gin/routers
	xjosiah.com/go-gin/routers/api.v1 => D:\go-gin/routers/api/v1
	xjosiah.com/go-gin/setting => D:\go-gin/setting
	xjosiah.com/go-gin/file =>  D:\go-gin/file
)
