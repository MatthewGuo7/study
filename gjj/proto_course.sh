

 protoc --proto_path=protos/ --micro_out=Course/ --gogo_out=Course/ common.proto
 protoc --proto_path=protos/ --micro_out=Course/ --gogo_out=Course/ course.proto
 protoc-go-inject-tag --input=Course/course.pb.go