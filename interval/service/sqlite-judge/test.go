package sqlite_judge

import (
	"database/sql"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	mapset "github.com/deckarep/golang-set"
	"github.com/infinity-oj/judger-sql/interval/basestruct/report"
	"github.com/infinity-oj/judger-sql/interval/basestruct/sqlcache"
	"github.com/infinity-oj/judger-sql/interval/service/sample"
	"github.com/infinity-oj/judger-sql/interval/util"
	_ "github.com/mattn/go-sqlite3"
)

var sqlCache = sqlcache.New()

type SQLiteTest struct {
	input  string
	sample sample.Sample
}

func (t *SQLiteTest) Init(s sample.Sample, input string) error {
	if s.Spec.Lang == sample.SQLite {
		sourceDB, err := os.Open(s.Spec.Database)
		if err != nil {
			return err
		}
		tmpDB, err := ioutil.TempFile("/tmp", "judge")
		if err != nil {
			return err
		}
		_, err = io.Copy(tmpDB, sourceDB)
		if err != nil {
			return err
		}
		s.DB, err = sql.Open("sqlite3", tmpDB.Name())
		s.TmpFile = tmpDB.Name()
		if err != nil {
			log.Fatal(err)
		}
		*t = SQLiteTest{input, s}
	}
	return nil
}

func (t *SQLiteTest) Run(reportChan chan report.Report) {
	s := t.sample
	var standardSlice []interface{}
	r := &report.Report{}
	res, ok := sqlCache.Get(s.Name, s.SQL)
	if !ok {
		standardRows, err := s.DB.Query(s.SQL)
		if err != nil {
			log.Println(s.SQL)
			log.Fatal(err)
		}
		standardSlice, _, err = util.ScanInterface(standardRows)
		if err != nil {
			log.Fatal(err)
		}
		sqlCache.Set(s.Name, s.SQL, standardSlice)
	} else {
		standardSlice = res
	}
	userRows, err := s.DB.Query(t.input)
	if err != nil {
		r.Grade = 0
		r.Summary = err.Error() + "\n"
		goto SEND
	} else {
		userSlice, _, err := util.ScanInterface(userRows)
		if err != nil {
			r.Grade = 0
			r.Summary = err.Error() + "\n"
			goto SEND

		}
		if len(userSlice) != len(standardSlice) {
			r.Grade = 0
			r.Summary = fmt.Sprintf("%s is wrong, row number is wrong\n", s.Name)
			goto SEND
		}
		if s.Spec.IsSet {
			s1 := mapset.NewSetFromSlice(standardSlice)
			s2 := mapset.NewSetFromSlice(userSlice)
			if !s1.Equal(s2) {
				r.Grade = 0
				r.Summary = fmt.Sprintf("%s is wrong\n", s.Name)
				goto SEND
			}
		} else {
			for i, s1 := range standardSlice {
				s2 := userSlice[i]
				if s1 != s2 {
					r.Grade = 0
					r.Summary = fmt.Sprintf("%s is wrong\n", s.Name)
					goto SEND
				}
			}
		}
	}
	r.Grade = s.Value
	r.Summary = fmt.Sprintf("%s is correct\n", s.Name)
SEND:
	reportChan <- *r
}

func (t SQLiteTest) Close() {
	err := os.Remove(t.sample.TmpFile)
	if err != nil {
		fmt.Println(err)
	}
}
