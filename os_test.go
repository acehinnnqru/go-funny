package funny

import (
	"github.com/stretchr/testify/require"
	"os"
	"path/filepath"
	"testing"
)

func TestBackTracePath(t *testing.T) {
	wd, _ := os.Getwd()

	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantPath    string
		wantFileInfo  func(os.FileInfo) bool
		err error
	}{
		{
			".git",
			args{path: ".git"},
			filepath.Join(wd, ".git"),
			func(info os.FileInfo) bool {
				return info.IsDir()
			},
			nil,
		},
		{
			"go.mod",
			args{path: "go.mod"},
			filepath.Join(wd, "go.mod"),
			func(info os.FileInfo) bool {
				return !info.IsDir()
			},
			nil,
		},
		{
			"..git",
			args{path: "..git"},
			"",
			func(info os.FileInfo) bool {
				return info == nil
			},
			os.ErrNotExist,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := BacktracePath(tt.args.path)

			require.True(t, err == tt.err)
			require.Equal(t, tt.wantPath, got)
			require.True(t, tt.wantFileInfo(got1))
		})
	}
}