package RenderingSpreadsheets

import (
    "fmt"

    "github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go-samples/config"
    viewer "github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go/models"
)

// This example demonstrates how to enable rendering of grid lines
func RenderSpreadsheetWithGridLines() {
    viewOptions := viewer.ViewOptions{
        FileInfo: &viewer.FileInfo{
            FilePath: "SampleFiles/sample.xlsx",
        },
        ViewFormat: viewer.ViewFormatHtml,
        RenderOptions: &viewer.HtmlOptions{
            SpreadsheetOptions: &viewer.SpreadsheetOptions{
                RenderGridLines: true,
            },
        },
    }

    response, _, err := config.Client.ViewApi.CreateView(config.Ctx, viewOptions)
    if err != nil {
        fmt.Printf("Exception: %v\n", err)
        return
    }

    fmt.Printf("RenderSpreadsheetWithGridLines completed: %v pages\n", len(response.Pages))
}