package sequence

import (
	"testing"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

func TestMySQLSequence_Next(t *testing.T) {
	type fields struct {
		conn       sqlx.SqlConn
		replaceSQL string
	}
	tests := []struct {
		name    string
		fields  fields
		want    uint64
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "demo",
			want: 0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MySQLSequence{
				conn:       tt.fields.conn,
				replaceSQL: tt.fields.replaceSQL,
			}
			got, err := m.Next()
			if (err != nil) != tt.wantErr {
				t.Errorf("MySQLSequence.Next() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MySQLSequence.Next() = %v, want %v", got, tt.want)
			}
		})
	}
}
