package executor

import (
	"context"
	"fmt"

	"github.com/pingcap/parser/model"
	"github.com/pingcap/tidb/util/chunk"
)

type SearchExecutor struct {
	baseExecutor
	ok bool

	DBName model.CIStr
	Table  *model.TableInfo
	Index  *model.IndexInfo

	Query string
	Mode  int
}

func (e *SearchExecutor) Close() error {
	return nil
}

func (e *SearchExecutor) Next(ctx context.Context, req *chunk.Chunk) error {
	fmt.Println("SearchExecutor", e.DBName, e.Table.Name, e.Index.Name, e.Query, e.Mode)
	req.Reset()
	if !e.ok {
		req.AppendBytes(0, []byte("huang"))
		e.ok = true
	}
	return nil
}

func (e *SearchExecutor) Open(ctx context.Context) error {
	return nil
}
