package sqlparser

import (
	"testing"
)

func testParse(t *testing.T, sql string) {
	_, err := Parse(sql)
	if err != nil {
		t.Fatal(err)
	}

}

func TestSet(t *testing.T) {
	sql := "set names gbk"
	testParse(t, sql)
}

func TestSimpleSelect(t *testing.T) {
	sql := "select last_insert_id() as a"
	testParse(t, sql)
}
