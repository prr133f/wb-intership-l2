package patterns

import (
	"fmt"
	"strings"
)

type SQLBuilder interface {
	Select(fields []string) SQLBuilder
	From(table string) SQLBuilder
	Where(conditions map[string]string) SQLBuilder
	Build() string
}

type SQL struct {
	connection string
	table      string
	fields     []string
	conditions map[string]string
}

func (s *SQL) Select(fields []string) SQLBuilder {
	s.fields = fields
	return s
}

func (s *SQL) From(table string) SQLBuilder {
	s.table = table
	return s
}

func (s *SQL) Where(conditions map[string]string) SQLBuilder {
	s.conditions = conditions
	return s
}

func (s *SQL) Build() string {
	var q strings.Builder
	q.WriteString("SELECT ")
	if len(s.fields) == 0 {
		q.WriteString("*")
	} else {
		q.WriteString(strings.Join(s.fields, ","))
	}

	q.WriteString(" FROM ")
	q.WriteString(s.table)

	if len(s.conditions) > 0 {
		var isFirst bool = true
		for k, v := range s.conditions {
			if isFirst {
				q.WriteString(fmt.Sprintf(" WHERE %s = %s", k, v))
				isFirst = false
			} else {
				q.WriteString(fmt.Sprintf(" AND %s = %s", k, v))
			}
		}
	}
	return q.String()
}
