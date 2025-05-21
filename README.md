# Go markdown table

With this module you can easily create tables and output them in markdown format.

## How to use

```bash
go get github.com/larsnieuwenhuizen/go-markdown-table@latest
```

### Examples

Create a table with header columns:

In your main.go file
```go
package main

import (
	"fmt"
	markdownTable "github.com/LarsNieuwenhuizen/go-markdown-table"
)

func main() {
	table := markdownTable.InitiateMarkdownTable()
	headerStringSlice := []string{
		"Name",
		"Description",
	}
	rowStringSlice := []string{
		"My name",
		"My description",
	}
	table.AddHeaderColumnsFromStringSlice(headerStringSlice)
	table.AddRowFromStringSlice(rowStringSlice)

	fmt.Println(table.ToString())
}
```

```bash
go run main.go
```

Result:

```markdown
| Name    | Description    |
|---------|----------------|
| My name | My description |
```