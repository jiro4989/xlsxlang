package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadXlsx(t *testing.T) {
	tests := []struct {
		desc    string
		path    string
		want    string
		wantErr bool
	}{
		{
			desc: "正常系: 正常に読み取れる",
			path: "./testdata/01.xlsx",
			want: `
 (if
  (= 3 2)
  (+ 1 2)
  (** 2 10)
 )
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			assert := assert.New(t)

			got, err := ReadXlsx(tt.path)
			if tt.wantErr {
				assert.Error(err)
				assert.Equal("", got)
				return
			}
			assert.Equal(tt.want, got)
		})
	}
}
