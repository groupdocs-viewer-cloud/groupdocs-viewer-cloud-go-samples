package CommonRenderingOptions

import (
    "fmt"

    "github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go-samples/config"
    viewer "github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go/models"
)

// This example demonstrates how to set default font name
func ReplaceMissingFont() {
    viewOptions := viewer.ViewOptions{
        FileInfo: &viewer.FileInfo{
            FilePath: "SampleFiles/with_missing_font.pptx",
        },
        ViewFormat: viewer.ViewFormatHtml,
        RenderOptions: &viewer.RenderOptions{
            DefaultFontName: "Courier New",
        },
    }

    response, _, err := config.Client.ViewApi.CreateView(config.Ctx, viewOptions)
    if err != nil {
        fmt.Printf("Exception: %v\n", err)
        return
    }

    fmt.Printf("ReplaceMissingFont completed: %v pages\n", len(response.Pages))
}