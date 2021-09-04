package unpacking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnpackString(t *testing.T) {
	var tests = []struct {
		got  string
		want string
		err  error
	}{
		{got: "a4bc2d5e",
			want: "aaaabccddddde",
			err:  nil},
		{got: "abcd",
			want: "abcd",
			err:  nil},
		{got: "45",
			want: "",
			err:  &ErrInvalidString{}},
		{got: "",
			want: "",
			err:  nil},
		{got: "qwe\\4\\5",
			want: "qwe45",
			err:  nil},
		{got: "qwe\\45",
			want: "qwe44444",
			err:  nil},
		{got: "qwe\\\\5",
			want: `qwe\\\\\`,
			err:  nil},
	}

	for _, tc := range tests {
		gotval, err := UnpackString(tc.got)
		assert.IsType(t, tc.err, err)
		assert.Equal(t, tc.want, gotval)
	}
}
