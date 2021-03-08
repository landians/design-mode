package builder

import "bytes"

// SQL 查询表达式接口, 这是表示过程
type ISQLQuery interface {
	ToSQL() string
}

// SQL 查询表达式的建造者接口，这是创建过程
type ISQLQueryBuilder interface {
	WithTable(table string) ISQLQueryBuilder
	WithField(field string) ISQLQueryBuilder
	WithCondition(condition string) ISQLQueryBuilder
	WithOrderBy(orderBy string) ISQLQueryBuilder
	Build() ISQLQuery
}

// 实现 ISQLQuery 接口
type SQLQuery struct {
	table      string
	fields     []string
	conditions []string
	orderBy    string
}

func newSQLQuery() *SQLQuery {
	return &SQLQuery{
		table:      "",
		fields:     make([]string, 0),
		conditions: make([]string, 0),
		orderBy:    "",
	}
}

func (s *SQLQuery) ToSQL() string {
	b := bytes.Buffer{}

	b.WriteString("SELECT ")

	for i, it := range s.fields {
		if i > 0 {
			b.WriteString(",")
		}
		b.WriteString(it)
	}

	b.WriteString(" FROM ")
	b.WriteString(s.table)

	if len(s.conditions) > 0 {
		b.WriteString(" WHERE ")
		for i, it := range s.conditions {
			if i > 0 {
				b.WriteString(" AND ")
			}
			b.WriteString(it)
		}
	}

	if len(s.orderBy) > 0 {
		b.WriteString(" ORDER BY ")
		b.WriteString(s.orderBy)
	}

	return b.String()
}

type SQLQueryBuilder struct {
	query *SQLQuery
}

func newSQLQueryBuilder() ISQLQueryBuilder {
	return &SQLQueryBuilder{
		query: newSQLQuery(),
	}
}

func (s *SQLQueryBuilder) WithTable(table string) ISQLQueryBuilder {
	s.query.table = table
	return s
}

func (s *SQLQueryBuilder) WithField(field string) ISQLQueryBuilder {
	s.query.fields = append(s.query.fields, field)
	return s
}

func (s *SQLQueryBuilder) WithCondition(condition string) ISQLQueryBuilder {
	s.query.conditions = append(s.query.conditions, condition)
	return s
}

func (s *SQLQueryBuilder) WithOrderBy(orderBy string) ISQLQueryBuilder {
	s.query.orderBy = orderBy
	return s
}

func (s *SQLQueryBuilder) Build() ISQLQuery {
	return s.query
}
