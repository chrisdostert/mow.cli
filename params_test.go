package cli

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBoolParam(t *testing.T) {
	var into bool

	param := &boolParam{into: &into}

	require.True(t, param.IsBoolFlag())

	cases := []struct {
		input  string
		err    bool
		result bool
		string string
	}{
		{"true", false, true, "true"},
		{"false", false, false, "false"},
		{"123", true, false, ""},
		{"", true, false, ""},
	}

	for _, cas := range cases {
		t.Logf("testing with %q", cas.input)

		err := param.Set(cas.input)

		if cas.err {
			require.Error(t, err, "value %q should have returned an error", cas.input)
			continue
		}

		require.Equal(t, cas.result, into)
		require.Equal(t, cas.string, param.String())
	}
}

func TestStringParam(t *testing.T) {
	var into string

	param := &stringParam{into: &into}

	cases := []struct {
		input  string
		string string
	}{
		{"a", `"a"`},
		{"", `""`},
	}

	for _, cas := range cases {
		t.Logf("testing with %q", cas.input)

		err := param.Set(cas.input)

		require.NoError(t, err)

		require.Equal(t, cas.input, into)
		require.Equal(t, cas.string, param.String())
	}
}

func TestIntParam(t *testing.T) {
	var into int

	param := &intParam{into: &into}

	cases := []struct {
		input  string
		err    bool
		result int
		string string
	}{
		{"12", false, 12, "12"},
		{"0", false, 0, "0"},
		{"01", false, 1, "1"},
		{"", true, 0, ""},
		{"abc", true, 0, ""},
	}

	for _, cas := range cases {
		t.Logf("testing with %q", cas.input)

		err := param.Set(cas.input)

		if cas.err {
			require.Error(t, err, "value %q should have returned an error", cas.input)
			continue
		}

		require.Equal(t, cas.result, into)
		require.Equal(t, cas.string, param.String())
	}
}

func TestStringsParam(t *testing.T) {
	param := &stringsParam{into: &([]string{})}

	require.True(t, param.IsMultiValued())

	param.SetMulti([]string{"a", " b "})

	require.Equal(t, []string{"a", "b"}, *param.into)

	param.Set("c")
	param.Set("d")

	require.Equal(t, []string{"a", "b", "c", "d"}, *param.into)
	require.Equal(t, `["a", "b", "c", "d"]`, param.String())
}

func TestIntsParam(t *testing.T) {
	param := &intsParam{into: &([]int{})}

	require.True(t, param.IsMultiValued())

	err := param.SetMulti([]string{"1", "a"})

	require.Error(t, err)
	require.Equal(t, []int{}, *param.into, "An error int SetMulti should not modify into")

	err = param.SetMulti([]string{"1", "  ", "2"})

	require.NoError(t, err)
	require.Equal(t, []int{1, 2}, *param.into)

	err = param.Set("c")
	require.Error(t, err)
	require.Equal(t, []int{1, 2}, *param.into)

	err = param.Set("3")
	require.NoError(t, err)
	require.Equal(t, []int{1, 2, 3}, *param.into)

	require.Equal(t, `[1, 2, 3]`, param.String())
}
