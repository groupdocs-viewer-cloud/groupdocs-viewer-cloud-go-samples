package RenderingVisioDocuments

import (
    "fmt"

    "github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go-samples/config"
    viewer "github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go/models"
)

// This example demonstrates how to render Visio documents figures
func RenderVisioDocumentFigures() {
    viewOptions := viewer.ViewOptions{
        FileInfo: &viewer.FileInfo{
            FilePath: "SampleFiles/sample.vssx",
        },
        ViewFormat: viewer.ViewFormatPng,
        RenderOptions: &viewer.ImageOptions{
            VisioRenderingOptions: &viewer.VisioRenderingOptions{
                RenderFiguresOnly: true,
                FigureWidth:       250,
            },
        },
    }

    response, _, err := config.Client.ViewApi.CreateView(config.Ctx, viewOptions)
    if err != nil {
        fmt.Printf("Exception: %v\n", err)
        return
    }

    fmt.Printf("RenderVisioDocumentFigures completed: %v pages\n", len(response.Pages))
}