package rltao

import (
	"testing"
)

type testcase struct {
	in   int
	want string
}

func TestGetTetragram(t *testing.T) {
	cases := []testcase{
		{1, "𝌆"},
		{81, "𝍖"},
		{15, "𝌔"},
	}
	for _, c := range cases {
		got := GetTetragram(c.in)
		if got != c.want {
			t.Errorf("GetTetragram(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
