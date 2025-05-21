package markdownTable

import (
	"fmt"
	"strings"
)

type Stringable interface {
	ToString() string
}

type MarkdownTable struct {
	HeaderColumns []HeaderColumn
	Rows          []Row
	Result        string
}

type HeaderColumn struct {
	Name string
}

type Row struct {
	Columns []Column
}

type Column struct {
	Value string
}

func (m *MarkdownTable) AddRow(row Row) {
	m.Rows = append(m.Rows, row)
}

func (m *MarkdownTable) AddHeaderColumn(column HeaderColumn) {
	m.HeaderColumns = append(m.HeaderColumns, column)
}

func (m *MarkdownTable) AddHeaderColumnsFromStringSlice(slice []string) {
	for _, column := range slice {
		m.AddHeaderColumn(HeaderColumn{Name: column})
	}
}

func (m *MarkdownTable) AddRowFromStringSlice(slice []string) {
	row := Row{}
	for _, value := range slice {
		row.Columns = append(row.Columns, Column{Value: value})
	}
	m.AddRow(row)
}

func (m *MarkdownTable) AddHeaderColumns(columns []HeaderColumn) {
	for _, column := range columns {
		m.HeaderColumns = append(m.HeaderColumns, column)
	}
}

func (m *MarkdownTable) ToString() string {
	m.BuildResult()
	return m.Result
}

func InitiateMarkdownTable() MarkdownTable {
	return MarkdownTable{
		HeaderColumns: make([]HeaderColumn, 0),
		Rows:          make([]Row, 0),
	}
}

func (m *MarkdownTable) BuildResult() MarkdownTable {
	var sb strings.Builder

	columnWidths := calculateColumnWidths(m.HeaderColumns, m.Rows)
	createHeaderRow(m.HeaderColumns, columnWidths, &sb)
	createSeparatorRow(columnWidths, &sb)
	createRows(m.Rows, columnWidths, &sb)

	m.Result = sb.String()

	return *m
}

func calculateColumnWidths(hc []HeaderColumn, rows []Row) []int {
	columnWidths := make([]int, len(hc))
	for i, header := range hc {
		columnWidths[i] = len(header.Name)
	}

	for _, row := range rows {
		for i, column := range row.Columns {
			if len(column.Value) > columnWidths[i] {
				columnWidths[i] = len(column.Value)
			}
		}
	}

	return columnWidths
}

func createHeaderRow(headerColumns []HeaderColumn, cw []int, sb *strings.Builder) *strings.Builder {
	for i, header := range headerColumns {
		sb.WriteString(
			fmt.Sprintf("| %-*s ", cw[i], header.Name),
		)
	}
	sb.WriteString("|\n")

	return sb
}

func createSeparatorRow(cw []int, sb *strings.Builder) *strings.Builder {
	for _, width := range cw {
		sb.WriteString(
			fmt.Sprintf(
				"|-%s",
				strings.Repeat("-", width+1),
			),
		)
	}
	sb.WriteString("|\n")

	return sb
}

func createRows(rows []Row, cw []int, sb *strings.Builder) *strings.Builder {
	for _, row := range rows {
		for i, column := range row.Columns {
			sb.WriteString(fmt.Sprintf("| %-*s ", cw[i], column.Value))
		}
		sb.WriteString("|\n")
	}

	return sb
}
