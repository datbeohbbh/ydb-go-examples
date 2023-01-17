module github.com/ydb-platform/ydb-go-examples

go 1.19

require (
	github.com/google/uuid v1.3.0
	github.com/gorilla/mux v1.8.0
	github.com/prometheus/client_golang v1.13.0
	github.com/rs/zerolog v1.27.0
	github.com/ydb-platform/ydb-go-sdk-auth-environ v0.1.2
	github.com/ydb-platform/ydb-go-sdk-prometheus v0.11.10
	github.com/ydb-platform/ydb-go-sdk-zerolog v0.12.2
	github.com/ydb-platform/ydb-go-sdk/v3 v3.42.4
	github.com/ydb-platform/ydb-go-yc v0.9.1
	google.golang.org/genproto v0.0.0-20221207170731-23e4bf6bdc37
	xorm.io/builder v0.3.11-0.20220531020008-1bd24a7dc978
	xorm.io/xorm v1.3.2
)

require (
	cloud.google.com/go/firestore v1.9.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/goccy/go-json v0.8.1 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/golang-jwt/jwt/v4 v4.4.3 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/jonboulle/clockwork v0.3.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.37.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	github.com/syndtr/goleveldb v1.0.0 // indirect
	github.com/yandex-cloud/go-genproto v0.0.0-20220815090733-4c139c0154e2 // indirect
	github.com/ydb-platform/ydb-go-genproto v0.0.0-20221215182650-986f9d10542f // indirect
	github.com/ydb-platform/ydb-go-sdk-metrics v0.16.3 // indirect
	github.com/ydb-platform/ydb-go-yc-metadata v0.5.3 // indirect
	golang.org/x/net v0.4.0 // indirect
	golang.org/x/sync v0.1.0 // indirect
	golang.org/x/sys v0.3.0 // indirect
	golang.org/x/text v0.5.0 // indirect
	google.golang.org/grpc v1.51.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)

replace xorm.io/xorm => ./../xorm