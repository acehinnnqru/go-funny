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
		wantIsDir  bool
		err error
	}{
		{
			".git",
			args{path: ".git"},
			filepath.Join(wd, ".git"),
			true,
			nil,
		},
		{
			"go.mod",
			args{path: "go.mod"},
			filepath.Join(wd, "go.mod"),
			false,
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := BackTracePath(tt.args.path)

			require.True(t, err == tt.err)
			require.Equal(t, tt.wantPath, got)
			require.Equal(t, tt.wantIsDir, got1.IsDir())
		})
	}
}