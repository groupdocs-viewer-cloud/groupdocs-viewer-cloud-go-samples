package RenderingPdfDocuments

import (
    "fmt"

    "github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go-samples/config"
    viewer "github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go/models"
)

// This example demonstrates how to convert pages to the same size as the size of pages in a source document
func RenderOriginalPageSize() {
    viewOptions := viewer.ViewOptions{
        FileInfo: &viewer.FileInfo{
            FilePath: "SampleFiles/sample.pdf",
        },
        ViewFormat: viewer.ViewFormatPng,
        RenderOptions: &viewer.ImageOptions{
            PdfDocumentOptions: &viewer.PdfDocumentOptions{
                RenderOriginalPageSize: true,
            },
        },
    }

    response, _, err := config.Client.ViewApi.CreateView(config.Ctx, viewOptions)
    if err != nil {
        fmt.Printf("Exception: %v\n", err)
        return
    }

    fmt.Printf("RenderOriginalPageSize completed: %v pages\n", len(response.Pages))
}