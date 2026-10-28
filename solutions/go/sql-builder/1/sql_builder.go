package sqlbuilder

import (
	"strings"
)

// ColumnValue represents a single database column and its corresponding value
// to be used in SQL INSERT statements.
type ColumnValue struct {
	Column string
	Value  string
}

// BuildSelectQuery generates a valid SQL SELECT query for the specified table and columns.
// If the columns slice is empty or nil, it defaults to the wildcard "*".
func BuildSelectQuery(table string, columns []string) string {
	var sb strings.Builder
	sb.WriteString("SELECT ")
	if len(columns) == 0 {
		sb.WriteString("*")
	} else {
		for i, col := range columns {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(col)
		}
	}
	sb.WriteString(" FROM ")
	sb.WriteString(table)
	sb.WriteString(";")
	return sb.String()
}

// BuildBatchInserts generates a slice of individual SQL INSERT statements for the provided rows.
// Empty rows are skipped, and values must be enclosed in single quotes.
func BuildBatchInserts(table string, rows [][]ColumnValue) []string {
	queries := make([]string, 0, len(rows))
	for _, row := range rows {
		if len(row) == 0 {
			continue
		}
		cols := make([]string, 0, len(row))
		vals := make([]string, 0, len(row))
		for _, cv := range row {
			cols = append(cols, cv.Column)
			vals = append(vals, "'"+cv.Value+"'")
		}
		var sb strings.Builder
		sb.WriteString("INSERT INTO ")
		sb.WriteString(table)
		sb.WriteString(" (")
		sb.WriteString(strings.Join(cols, ", "))
		sb.WriteString(") VALUES (")
		sb.WriteString(strings.Join(vals, ", "))
		sb.WriteString(");")
		queries = append(queries, sb.String())
	}
	return queries
}
