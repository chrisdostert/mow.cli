package cli

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type boolValued interface {
	flag.Value
	IsBoolFlag() bool
}

type multiValued interface {
	flag.Value
	IsMultiValued() bool
	SetMulti([]string) error
}

/******************************************************************************/
/* BOOL                                                                        */
/******************************************************************************/

type boolParam struct {
	into *bool
}

var (
	_ flag.Value = &boolParam{}
	_ boolValued = &boolParam{}
)

func (bo *boolParam) Set(s string) error {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return err
	}
	*bo.into = b
	return nil
}

func (bo *boolParam) IsBoolFlag() bool {
	return true
}

func (bo *boolParam) String() string {
	return fmt.Sprintf("%v", *bo.into)
}

/******************************************************************************/
/* STRING                                                                        */
/******************************************************************************/

type stringParam struct {
	into *string
}

var (
	_ flag.Value = &stringParam{}
)

func (sa *stringParam) Set(s string) error {
	*sa.into = s
	return nil
}

func (sa *stringParam) String() string {
	return fmt.Sprintf("%#v", *sa.into)
}

/******************************************************************************/
/* INT                                                                        */
/******************************************************************************/

type intParam struct {
	into *int
}

var (
	_ flag.Value = &intParam{}
)

func (ia *intParam) Set(s string) error {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	*ia.into = int(i)
	return nil
}

func (ia *intParam) String() string {
	return fmt.Sprintf("%v", *ia.into)
}

/******************************************************************************/
/* STRINGS                                                                    */
/******************************************************************************/

// Strings describes a string slice argument
type stringsParam struct {
	into *[]string
}

var (
	_ flag.Value  = &stringsParam{}
	_ multiValued = &stringsParam{}
)

func (sa *stringsParam) Set(s string) error {
	*sa.into = append(*sa.into, s)
	return nil
}

func (sa *stringsParam) String() string {
	res := "["
	for idx, s := range *sa.into {
		if idx > 0 {
			res += ", "
		}
		res += fmt.Sprintf("%#v", s)
	}
	return res + "]"
}

func (sa *stringsParam) IsMultiValued() bool {
	return true
}

func (sa *stringsParam) SetMulti(vs []string) error {
	newValue := make([]string, len(vs))
	for idx, v := range vs {
		v = strings.TrimSpace(v)
		newValue[idx] = v
	}
	sa.into = &newValue
	return nil
}

/******************************************************************************/
/* INTS                                                                       */
/******************************************************************************/

// Ints describes an int slice argument
type intsParam struct {
	into *[]int
}

var (
	_ flag.Value  = &intsParam{}
	_ multiValued = &intsParam{}
)

func (ia *intsParam) Set(s string) error {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	*ia.into = append(*ia.into, int(i))
	return nil
}

func (ia *intsParam) String() string {
	res := "["
	for idx, s := range *ia.into {
		if idx > 0 {
			res += ", "
		}
		res += fmt.Sprintf("%v", s)
	}
	return res + "]"
}

func (ia *intsParam) IsMultiValued() bool {
	return true
}

func (ia *intsParam) SetMulti(vs []string) error {
	newValue := []int{}
	for _, v := range vs {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}
		i, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return err
		}
		newValue = append(newValue, int(i))
	}
	ia.into = &newValue
	return nil
}
