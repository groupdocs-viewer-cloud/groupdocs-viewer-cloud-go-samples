package RenderingPdfDocuments

import (
    "fmt"

    "github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go-samples/config"
    viewer "github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go/models"
)

// This example demonstrates how to adjust the display of an outline font
func EnableFontHinting() {
    viewOptions := viewer.ViewOptions{
        FileInfo: &viewer.FileInfo{
            FilePath: "SampleFiles/sample.pdf",
        },
        ViewFormat: viewer.ViewFormatPng,
        RenderOptions: &viewer.ImageOptions{
            PdfDocumentOptions: &viewer.PdfDocumentOptions{
                EnableFontHinting: true,
            },
        },
    }

    response, _, err := config.Client.ViewApi.CreateView(config.Ctx, viewOptions)
    if err != nil {
        fmt.Printf("Exception: %v\n", err)
        return
    }

    fmt.Printf("EnableFontHinting completed: %v pages\n", len(response.Pages))
}