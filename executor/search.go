package executor

import (
	"context"
	"fmt"

	"github.com/pingcap/parser/model"
	"github.com/pingcap/tidb/util/chunk"
	"github.com/redis-force/tisearch/storage"
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
	result, err := storage.Search(ctx, "tisearch", e.Table.Name.L, e.Query)
	if err != nil {
		return err
	}
	if !e.ok {
		for _, row := range result.Rows {
			req.AppendInt64(0, row.DocID)
			idx := 1
			for _, column := range e.Index.Columns {
				req.AppendString(idx, row.Render[column.Name.L][0])
				idx++
			}
		}
		e.ok = true
	}
	return nil
}

func (e *SearchExecutor) Open(ctx context.Context) error {
	return nil
}
