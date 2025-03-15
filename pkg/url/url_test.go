package url

import (
	"testing"
)

func TestGetBasePath(t *testing.T) {
	type args struct {
		inputUrl string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "normal example",
			args: args{
				inputUrl: "http://www.baidu.com/u987123xxx",
			},
			want:    "u987123xxx",
			wantErr: false,
		},
		{
			name: "error with no host",
			args: args{
				inputUrl: "/xxx/123456",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetBasePath(tt.args.inputUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBasePath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetBasePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
