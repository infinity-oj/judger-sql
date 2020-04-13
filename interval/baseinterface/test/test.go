package test

import (
	"log"

	"github.com/infinity-oj/judger-sql/interval/basestruct/report"
	pgsql_judge "github.com/infinity-oj/judger-sql/interval/service/pgsql-judge"
	"github.com/infinity-oj/judger-sql/interval/service/sample"
	sqlite_judge "github.com/infinity-oj/judger-sql/interval/service/sqlite-judge"
)

type Test interface {
	Run(report chan report.Report)
	Init(s sample.Sample, input string) error
	Close()
}

func SelectTest(s sample.Sample) Test {
	switch s.Spec.Lang {
	case sample.SQLite:
		return &sqlite_judge.SQLiteTest{}
	case sample.Postgres:
		return &pgsql_judge.PGSQLTest{}
	default:
		log.Fatal("no default sample type")
	}
	return nil
}
