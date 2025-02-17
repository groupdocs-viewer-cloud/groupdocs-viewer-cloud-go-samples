package LoadingOptions

import (
    "fmt"

    "github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go-samples/config"
    viewer "github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go/models"
)

// This example demonstrates how to specify encoding when rendering documents
func SpecifyEncoding() {
    viewOptions := viewer.ViewOptions{
        FileInfo: &viewer.FileInfo{
            FilePath: "SampleFiles/shift_jis_encoded.txt",
        },
        ViewFormat: viewer.ViewFormatHtml,
        RenderOptions: &viewer.RenderOptions{
            DefaultEncoding: "shift_jis",
        },
    }

    response, _, err := config.Client.ViewApi.CreateView(config.Ctx, viewOptions)
    if err != nil {
        fmt.Printf("Exception: %v\n", err)
        return
    }

    fmt.Printf("SpecifyEncoding completed: %v pages\n", len(response.Pages))
}