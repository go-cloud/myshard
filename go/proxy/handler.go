package proxy

import (
	"fmt"
	. "github.com/siddontang/go-mysql/mysql"
)

type Handler struct {
}

func (h *Handler) UseDB(dbName string) error {
	return nil
}

func (h *Handler) HandleQuery(query string) (*Result, error) {
	return nil, fmt.Errorf("not supported now")
}

func (h *Handler) HandleFieldList(table string, fieldWildcard string) ([]*Field, error) {
	return nil, fmt.Errorf("not supported now")
}

func (h *Handler) HandleStmtPreprare(query string) (params int, columns int, context interface{}, err error) {
	err = fmt.Errorf("not supported now")
	return
}

func (h *Handler) HandleStmtExecute(context interface{}, query string, args []interface{}) (*Result, error) {
	return nil, fmt.Errorf("not supported now")
}
