module go-frame

go 1.14

replace (
	google.golang.org/api => github.com/googleapis/google-api-go-client v0.32.0
	google.golang.org/genproto => github.com/googleapis/go-genproto v0.0.0-20200904004341-0bd0a958aa1d
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)

require (
	github.com/Me1onRind/logrotate v0.0.0-20201207055048-cc28c78da981
	github.com/coreos/etcd v3.3.18+incompatible
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/validator/v10 v10.4.1
	github.com/golang/protobuf v1.4.3
	github.com/gorilla/sessions v1.2.1
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/etcd/v2 v2.9.1
	github.com/minio/minio-go v6.0.14+incompatible
	github.com/opentracing/opentracing-go v1.1.0
	github.com/robfig/cron/v3 v3.0.1
	github.com/sirupsen/logrus v1.7.0 // indirect
	github.com/spf13/viper v1.7.1
	go.elastic.co/apm v1.10.0
	go.elastic.co/apm/module/apmhttp v1.10.0
	go.elastic.co/apm/module/apmot v1.10.0
	go.opentelemetry.io/otel v0.15.0
	go.opentelemetry.io/otel/exporters/trace/jaeger v0.15.0
	go.opentelemetry.io/otel/sdk v0.15.0
	go.uber.org/zap v1.16.0
	golang.org/x/crypto v0.0.0-20200709230013-948cd5f35899 // indirect
	gopkg.in/ini.v1 v1.57.0 // indirect
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.20.9
)
