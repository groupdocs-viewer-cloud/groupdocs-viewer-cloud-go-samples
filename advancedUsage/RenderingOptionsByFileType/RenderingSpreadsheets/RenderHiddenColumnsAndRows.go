package RenderingSpreadsheets

import (
    "fmt"

    "github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go-samples/config"
    viewer "github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go/models"
)

// This example demonstrates how to enable rendering of hidden rows and columns
func RenderHiddenColumnsAndRows() {
    viewOptions := viewer.ViewOptions{
        FileInfo: &viewer.FileInfo{
            FilePath: "SampleFiles/with_hidden_row_and_column.xlsx",
        },
        ViewFormat: viewer.ViewFormatHtml,
        RenderOptions: &viewer.HtmlOptions{
            SpreadsheetOptions: &viewer.SpreadsheetOptions{
                RenderHiddenColumns: true,
                RenderHiddenRows:    true,
            },
        },
    }

    response, _, err := config.Client.ViewApi.CreateView(config.Ctx, viewOptions)
    if err != nil {
        fmt.Printf("Exception: %v\n", err)
        return
    }

    fmt.Printf("RenderHiddenColumnsAndRows completed: %v pages\n", len(response.Pages))
}