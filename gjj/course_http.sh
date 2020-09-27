#export MICRO_REGISTRY=etcd
#export MICRO_REGISTRY_ADDRESS=localhost:2379
#export MICRO_LOG_LEVEL=debug
#go run cmd/course_service.go --server_address :9091

export MICRO_REGISTRY=etcd
export MICRO_REGISTRY_ADDRESS=localhost:2379
export MICRO_LOG_LEVEL=debug
echo $MICRO_REGISTRY
echo $MICRO_REGISTRY_ADDRESS
go run cmd/course_http_server.go --server_address :9000
