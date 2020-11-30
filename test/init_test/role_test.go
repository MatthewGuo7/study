package initTest_test

import (
	"encoding/json"
	sq "github.com/Masterminds/squirrel"
)

import (
	"fmt"
	initTest "test/init_test"
	"testing"
)

type TestSliceJson struct {
	Values []int `json:"valeus,omitempty"`
}

func TestSlice(t *testing.T) {
	v := &TestSliceJson{
		Values: make([]int, 0),
	}
	b, _ := json.Marshal(v)
	fmt.Println(string(b))
}

const (
	DispatchTypeToDept = iota + 1
	DispatchTypeToUser
)

func TestEnum(t *testing.T) {
	fmt.Println(DispatchTypeToDept)
	fmt.Println(DispatchTypeToUser)
}

func TestRole(t *testing.T) {
	roleMange := initTest.NewRoleMange(2)
	roleMange.DoAction()
}

type TestInter interface{}

type TypeString1 struct {
	value string
}

type TestString struct {
	Value *TypeString1
}

func (t *TestString) String() string {
	return fmt.Sprintf("value = %+v", t.Value)
}

func TestInterfaceNil(t *testing.T) {
	v1 := &TypeString1{
		value: "hello",
	}
	v := &TestString{
		Value: v1,
	}

	fmt.Println(v)
}

/*
func TestBeegoOrm(t *testing.T) {
	qb, _ := orm.NewQueryBuilder("mysql")

	// 构建查询对象
	qb.Select("user.id",
		"user.name",
		"profile.age").
		From("user").
		InnerJoin("profile").On("user.profile_id = profile.id").
		Where("age > ?").
		OrderBy("name").Desc().
		Limit(10).Offset(0)

	// 导出 SQL 语句
	sql := qb.String()
	fmt.Println(sql)

}

func TestBeegoOrm1(t *testing.T) {
	qb, _ := orm.NewQueryBuilder("mysql")

	// 构建查询对象
	qb.Select("users.*").
		From("users").
		LeftJoin("user_role").
		On("users.uid = user_role.uid").
		LeftJoin("dept").
		On("users.dept_id = dept.id").
		Where("users.uid = 180481").
		Limit(10).Offset(0)

	// 导出 SQL 语句
	sql := qb.String()
	fmt.Println(sql)

}

func TestBeegoOrm2(t *testing.T) {
	qb, _ := orm.NewQueryBuilder("mysql")
	// 构建查询对象
	query := qb.Select("users.phoneno phone",
		"users.uid uid",
		"users.nickname name",
		"role.id role",
		"role.name role_name",
		"users.grade grade",
		"users.subject subject",
		"users.exten exten",
		"users.work_place work_place",
		"users.dept_id dept_id",
		"dept.name dept_name",
		"users.status status",
		"users.is_dept_manager is_dept_manager").
		From("users").
		InnerJoin("user_role").
		On("users.uid = user_role.uid").
		InnerJoin("dept").
		On("users.dept_id = dept.id").
		InnerJoin("role").
		On("user_role.role_id = role.id").
		Where("user_role.role_id").In([]string{"2","3","4"}...).
		And("users.nickname").In([]string{"''"}...)
	fmt.Println(query.String())
	fmt.Printf(" uid = %#v\n",10)
	fmt.Printf("uid = %p", []interface{}{10,20})
}
*/

type ITestImpl interface {
	JoinDept(impl ITestImpl)
	CallMe()
}

type TestImplParent struct {
}

func NewTestImplParent() *TestImplParent {
	return &TestImplParent{}
}

func (t *TestImplParent) Say() {
	fmt.Println("parent say")
}

func (t *TestImplParent) JoinDept(impl ITestImpl) {
	impl.CallMe()
	fmt.Println("join dept")
}

func (t *TestImplParent) CallMe() {
	fmt.Println("test impl parent")
}

type TestImplFirst struct {
	*TestImplParent
}

func NewTestImplFirst() *TestImplFirst {
	return &TestImplFirst{
		TestImplParent: NewTestImplParent(),
	}
}

func (t *TestImplFirst) CallMe() {
	//t.TestImplParent.CallMe()
	fmt.Println("test impl first")
}

type TestImpleSecond struct {
	*TestImplParent
}

func NewTestImplSecond() *TestImpleSecond {
	return &TestImpleSecond{
		TestImplParent: NewTestImplParent(),
	}
}

func TestParentImpl(t *testing.T) {
	var first ITestImpl = NewTestImplFirst()
	first.JoinDept(first)

	var second ITestImpl = NewTestImplSecond()
	second.JoinDept(second)

	var parent ITestImpl = NewTestImplParent()
	parent.JoinDept(parent)
}
func Test1(t *testing.T) {

}

func Test(t *testing.T) {
	users := sq.Select("*").From("users").Join("emails USING (email_id)")

	active := users.Where(sq.Eq{"username": "moe"})
	//active := users.Where(sq.Eq{"username": []string{"moe", "larry", "curly", "shemp"}})
	//active := users.Where(sq.Eq{"deleted_at": nil})

	sql, args, err := active.ToSql()

	fmt.Println(sql, args, err)

	fmt.Println(fmt.Sprintf("%+v.%+v = %+v.%+v", "usersTableName", "UID", "userRoleTableName", "def.UID"))
}

type Father struct {
}

func (f *Father) Encode() {
	fmt.Println("father")
}

func demoFunction(father *Father) {
	father.Encode()
}

type Son struct {
	*Father
}

func (s *Son) Encode() {
	fmt.Println("son")
}

func TestExtend(t *testing.T) {
}

func testDefer() {
	code := 10
	msg := "hello"
	defer func() {
		fmt.Println(code, msg)
	}()

	code = 20
	msg = "asdfa"
}

func TestDefer(t *testing.T) {
	testDefer()
}
