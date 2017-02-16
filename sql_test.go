package sqlxplus

import (
	_ "github.com/go-sql-driver/mysql"
	"testing"
	"fmt"
)





func TestBatchInsertBegin(t *testing.T){

 a :=	BatchInsertBegin(nil, "insert into test_table(id, name)", 5, "hello")

	fmt.Println(a)
}



func TestBatchInsertAdd(t *testing.T){

 a :=	BatchInsertBegin(nil, "insert into test_table(id, name)", 5, "hello")
	a.Add(6 , "world")
	fmt.Println(a)
}



func TestBatchInsertExec(t *testing.T){

 a :=	BatchInsertBegin(nil, "insert into test_table(id, name)", 5, "hello")
	a.Add(6 , "world")
	a.Exec()
}
