package sql_test

import (
	"fmt"
	. "github.com/Masterminds/squirrel"
	"testing"
)

type TestBuilder struct {
	selector SelectBuilder
}

func (t *TestBuilder) BuildName() {
	t.selector = t.selector.Where(Or{Eq{"name": 5}})
}

func (t *TestBuilder) BuildAge() {
	t.selector = t.selector.Where(map[string]interface{}{
		"age":   []int{20, 30},
		"hello": "world",
	})
}

func CountColumn(builder SelectBuilder) SelectBuilder {
	builder = builder.Column("count(1) as count")
	return builder
}

func QueryColumn(builder SelectBuilder) SelectBuilder {
	builder = builder.Column("*")
	return builder
}

func TestSqu(t *testing.T) {
	builder := Select("*")
	builder = builder.Where(Eq{"a": 2})
	fmt.Println(builder.ToSql())
}

func TestInsert(t *testing.T) {
	var (
		name string = "hell"
		age  int    = 20
		sex  int    = 0
		//status int    = 1
		//substatus int    = 2
	)
	//s := Select("*").From("user")
	//s := SelectBuilder(builder.EmptyBuilder)
	s := StatementBuilder

	if name != "" {
		//conds = append(conds, Eq{"name": name})
		s = s.Where(Eq{"name": name})
	}

	if age > 0 {
		//conds = append(conds, Eq{"age": age})
		s = s.Where(Eq{"age": age})
	}

	if sex > 0 {
		//conds = append(conds, Eq{"sex": sex})
		s = s.Where(Eq{"sex": sex})
	}

	/*
		if status > 0 {
			//conds = append(conds, Eq{"status": status})
			statusExpr := Expr(" status = ? ", status)
			//s = s.Where(Eq{"status": status})
			if status == 1 {
				ageExpr := Expr(" or age = ? ", 20)
				statusExpr = ConcatExpr(statusExpr, ageExpr)
				//s = s.Where(Eq{"age": 20}) // or age = 20
				//conds = append(conds, Or{Eq{"age": 20}})
			}

			sql, args, err := statusExpr.ToSql()
			if nil != err {
				fmt.Println(err)
			}
			fmt.Println(sql, args, err)
			s = s.Where(sql, args...)
		}

		//if substatus > 0 {
		//	s = s.Where(Eq{"substatus": substatus})
		//}

		//SELECT * FROM user WHERE name = ?  AND (status = 1 or age = 20) AND substatus = ?
	*/

	fmt.Println(s.Select("a").ToSql())

	//sql, args, err := s.ToSql()
	//fmt.Println(sql, args, err)
}
func TestConcatExpr(t *testing.T) {
	b := ConcatExpr("COALESCE(name,", Expr("CONCAT(?,' ',?)", "f", "l"), ")")
	b = ConcatExpr(b, Expr(" or age = ? ", 20))
	fmt.Println(b.ToSql())
	//assert.NoError(t, err)
	//
	//expectedSql := "COALESCE(name,CONCAT(?,' ',?))"
	//assert.Equal(t, expectedSql, sql)
	//
	//expectedArgs := []interface{}{"f", "l"}
	//assert.Equal(t, expectedArgs, args)
}

func TestSqlizer(t *testing.T) {
	s := []Sqlizer{
		Eq{"a": 10},
		GtOrEq{"b": 20},
	}
	c := ConcatExpr(s[0].ToSql())
	fmt.Println(c)
	for _, temp := range s {
		str, i, _ := temp.ToSql()
		c = ConcatExpr(c, str, i)
	}
	fmt.Println(c.ToSql())

	builder := SelectBuilder{}
	builder = builder.Where(s).Columns("*")
	fmt.Println(builder.ToSql())

}

func TestStatement(t *testing.T) {
	nestedBuilder := StatementBuilder.
		Select("*").Columns("a", "b", "C").
		//Prefix("NOT EXISTS (").
		From("bar").Where("y = ?", 42).Suffix(")")
	fmt.Println(nestedBuilder.ToSql())
}

func insert(v interface{}) {
	fmt.Println(v)
}

func Convert(value ...interface{}) {
	for _, v := range value {
		insert(v)
	}
}

func TestUpdate(t *testing.T) {
	values := []int{1, 2, 3}
	Convert(values)
}

func updateHelper(updater UpdateBuilder) UpdateBuilder {
	return updater.Table("clue_info")
}

func TestSelect(t *testing.T) {
	s := Select("").Column("a").From("b")
	fmt.Println(s.ToSql())
}

func TestMultiUpdate(t *testing.T) {
	update := Update("").
		Set("updatecnt", Expr("updatecnt + 1")).
		Set("phone", Case("phone").When("10", "20")).Where(Eq{"clueid": 4})
	update = updateHelper(update)

	fmt.Println(update.ToSql())
}

type ITest interface {
	do()
}

type ITest2 interface {
	ITest
	dodo()
}

type TestDoubleInter struct {
}

func (t TestDoubleInter) do() {
	fmt.Println("do")
}

func (t TestDoubleInter) dodo() {
	fmt.Println("dodo")
}

func intertest(test ITest) {
	test.do()
}

func TestInterface(t *testing.T) {
	var i ITest2 = &TestDoubleInter{}
	intertest(i)
}

func SqlUpdateFunc(builder *UpdateBuilder) {
	*builder = builder.Set("key", 10).Set("key1", 20)
	fmt.Println(builder.ToSql())
}

func TestUpdateFunc(t *testing.T) {
	updater := Update("test")
	SqlUpdateFunc(&updater)
	fmt.Println(updater.ToSql())
}

func TestEmptyWhere(t *testing.T) {
	//values := make([]uint32, 0)
	//updater := Update("test").Where(Eq{"id": values}).Set("name", "hello")
	//fmt.Println(updater.ToSql())
	//

	updater := Update("users").Set("name", Case("id").
		When("1", Expr(fmt.Sprintf("%+v", "anme100"))).
		When("2", Expr(fmt.Sprintf("%+v", "name200")))).Where(Eq{"id": []int{1, 2}})
	fmt.Println(updater.ToSql())
}

func TestGroup(t *testing.T) {
	selector := StatementBuilder.Select("*").Column("b").Columns("count(1) as count").
		GroupBy("a", "b").From("test").Where(Eq{"A": 20})
	fmt.Println(selector.ToSql())
}

func TestSliceAppend(t *testing.T) {
	var s []int
	s = append(s, 10)
	fmt.Println(s)
}

func TestElseIf(t *testing.T)  {
	a := 10
	b := 20

	if a == 10 {
		fmt.Println(a)
	} else if b == 20 {
		fmt.Println(b)
	} else {
		fmt.Println("no")
	}


	switch  {
	case true:
		fmt.Println(a)
	case a == 10:
		fmt.Println(b)
	default:
		fmt.Println("no ")
	}

}
