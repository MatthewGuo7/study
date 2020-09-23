export MICRO_REGISTRY=etcd
export MICRO_REGISTRY_ADDRESS=localhost:2379
export MICRO_LOG_LEVEL=debug
echo $MICRO_REGISTRY
echo $MICRO_REGISTRY_ADDRESS
go run cmd/user_server.go --server_address :9090