package tests

import (
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"testing"
)

func TestProtoBufTimeStamp(t *testing.T)  {

	//timeGo := time.Now()
	timeGo, err := ptypes.Timestamp(ptypes.TimestampNow())
	fmt.Println(timeGo, err)

	timeProto, err := ptypes.TimestampProto(timeGo)
	fmt.Println(timeProto, err)

}
