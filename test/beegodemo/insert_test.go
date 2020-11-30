package beegodemo

import "testing"

func initMysql(){
	err := models.InitMysql("root:123456@tcp(172.16.31.217:3306)/uclass_pd?charset=utf8")
	if nil != err {
		panic(fmt.Sprintf("init mysql error, error = %+v", err))
	}
}

func TestInsert(t *testing.T)  {

}
