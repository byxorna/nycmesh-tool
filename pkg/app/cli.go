package app

// cli.go implements CLI coloring, formatting of structured data, etc
//

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func (a *App) ConsoleWriteObject(obj string) error {
	fmt.Printf("%+v\n", obj)
	return nil
}

func (a *App) Tableify(headers []string, data [][]string) error {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(headers)
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t")
	table.SetNoWhiteSpace(true)
	table.AppendBulk(data)
	table.Render()
	return nil
}
