package sequence

import (
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

const replaceSQLTemplate = "REPLACE INTO %s (%s) VALUES ('%s')"



type MySQLSequence struct {
	conn sqlx.SqlConn
	replaceSQL string
}

func NewMySQLSequence(dsn, table, field, value string) Sequence {
	return &MySQLSequence{
		conn: sqlx.NewMysql(dsn),
		replaceSQL: fmt.Sprintf(replaceSQLTemplate, table, field, value),
	}
}

func (m *MySQLSequence) Next() (uint64, error) {
	stmt, err := m.conn.Prepare(m.replaceSQL)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec()
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(id), nil
}