package sqlxplus

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type BatchSqlEntity struct {
	db              *sqlx.DB
	sqlHeader       string
	sqltemplateList []string
	sqlvalList      [][]interface{}
}

func generateTemplateString(args ...interface{}) string {
	templatStr := ""
	for range args {
		templatStr += "?, "
	}
	templatStr = "(" + templatStr[0:len(templatStr)-2] + ")"
	return templatStr
}

func (entity *BatchSqlEntity) Add(args ...interface{}) {
	templateStr := generateTemplateString(args...)
	entity.sqltemplateList = append(entity.sqltemplateList, templateStr)
	entity.sqlvalList = append(entity.sqlvalList, args)
}

func (entity *BatchSqlEntity) Exec() (sql.Result, error) {
	db := entity.db
	sql := entity.sqlHeader
	for _, template := range entity.sqltemplateList {
		sql += template + ", "
	}
	var vals []interface{}
	for i := 0; i < len(entity.sqlvalList); i++ {
		vs := entity.sqlvalList[i]
		vals = append(vals, vs...)
	}
	sql = sql[0 : len(sql)-2]
	return db.Exec(sql, vals...)
}

func BatchInsertBegin(db *sqlx.DB, insertSql string, args ...interface{}) *BatchSqlEntity {
	sql := insertSql + " values "
	templateStr := generateTemplateString(args...)
	templateList := []string{templateStr}
	var sqlvalList [][]interface{}
	sqlvalList = append(sqlvalList, args)
	entity := BatchSqlEntity{db: db, sqlHeader: sql, sqltemplateList: templateList, sqlvalList: sqlvalList}
	return &entity
}

func BatchInsertInit(db *sqlx.DB, insertSql string) *BatchSqlEntity {
	sql := insertSql + " values "
	templateList := []string{}
	var sqlvalList [][]interface{}
	entity := BatchSqlEntity{db: db, sqlHeader: sql, sqltemplateList: templateList, sqlvalList: sqlvalList}
	return &entity
}
