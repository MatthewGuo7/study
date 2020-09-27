package Course

import (
	"database/sql/driver"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"time"
)

func (t *TimeStamp) Scan(value interface{}) (err error) {
	switch valueType := value.(type) {
	case time.Time:
		t.Timestamp, err = ptypes.TimestampProto(valueType)
		if nil != err {
			return err
		}
	default:
		return fmt.Errorf("timestamp error")
	}
	return nil
}


func (this *TimeStamp) Value() (driver.Value, error) {
	return ptypes.Timestamp(this.Timestamp)
}
