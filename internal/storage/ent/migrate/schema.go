// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// JocksColumns holds the columns for the "jocks" table.
	JocksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString},
		{Name: "content", Type: field.TypeString, Size: 2147483647},
	}
	// JocksTable holds the schema information for the "jocks" table.
	JocksTable = &schema.Table{
		Name:       "jocks",
		Columns:    JocksColumns,
		PrimaryKey: []*schema.Column{JocksColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		JocksTable,
	}
)

func init() {
}
