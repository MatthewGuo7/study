module firstdemo

go 1.13

require (
	github.com/coreos/etcd v3.3.25+incompatible // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/register_center v0.0.0-00010101000000-000000000000
	go.etcd.io/etcd v3.3.25+incompatible // indirect
	go.uber.org/zap v1.16.0 // indirect
	google.golang.org/grpc v1.26.0 // indirect
)

replace github.com/register_center => ../register_center

replace github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.4
