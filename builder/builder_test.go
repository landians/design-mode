package builder

import (
	"fmt"
	"testing"
)

func Test_Builder(t *testing.T) {
	builder := newSQLQueryBuilder()
	builder.WithTable("product").
		WithField("id").WithField("name").WithField("price").
		WithCondition("enable=1").
		WithOrderBy("price desc")
	query := builder.Build()
	fmt.Println(query.ToSQL())
}
