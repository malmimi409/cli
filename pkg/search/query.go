package search

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/cli/cli/v2/pkg/text"
)

type Query struct {
	Keywords   []string
	Kind       string
	Limit      int
	Order      string
	Page       int
	Qualifiers Qualifiers
	Sort       string
}

type Qualifiers struct {
	Archived         *bool
	Created          string
	Followers        string
	Fork             string
	Forks            string
	GoodFirstIssues  string
	HelpWantedIssues string
	In               []string
	Is               string
	Language         []string
	License          []string
	Mirror           *bool
	Org              string
	Pushed           string
	Repo             string
	Size             string
	Stars            string
	Topic            []string
	Topics           string
	User             string
}

func (q Query) String() string {
	all := q.Keywords
	for k, v := range q.Qualifiers.Map() {
		all = append(all, fmt.Sprintf("%s:%s", k, v))
	}
	for k, v := range all {
		all[k] = quote(v)
	}
	return strings.Join(all, " ")
}

func (q Qualifiers) Map() map[string]string {
	m := map[string]string{}
	v := reflect.ValueOf(q)
	t := reflect.TypeOf(q)
	for i := 0; i < v.NumField(); i++ {
		fieldName := t.Field(i).Name
		key := text.CamelToKebab(fieldName)
		typ := v.FieldByName(fieldName).Kind()
		value := v.FieldByName(fieldName)
		switch typ {
		case reflect.Ptr:
			if value.IsNil() {
				continue
			}
			v := reflect.Indirect(value)
			m[key] = fmt.Sprintf("%v", v)
		case reflect.Slice:
			if value.IsNil() {
				continue
			}
			s := []string{}
			for i := 0; i < value.Len(); i++ {
				s = append(s, fmt.Sprintf("%v", value.Index(i)))
			}
			m[key] = strings.Join(s, ",")
		default:
			if value.IsZero() {
				continue
			}
			m[key] = fmt.Sprintf("%v", value)
		}
	}
	return m
}

func quote(k string) string {
	if strings.ContainsAny(k, " \"\t\r\n") {
		if strings.Contains(k, ":") {
			z := strings.SplitN(k, ":", 2)
			return fmt.Sprintf("%s:%q", z[0], z[1])
		}
		return fmt.Sprintf("%q", k)
	}
	return k
}
