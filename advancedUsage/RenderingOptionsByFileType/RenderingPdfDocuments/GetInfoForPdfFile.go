package RenderingPdfDocuments

import (
    "fmt"

    "github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go-samples/config"
    viewer "github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go/models"
)

// This example demonstrates how to retrieve view information for a PDF file
func GetInfoForPdfFile() {
    viewOptions := viewer.ViewOptions{
        FileInfo: &viewer.FileInfo{
            FilePath: "SampleFiles/sample.pdf",
        },
    }

    response, _, err := config.Client.InfoApi.GetInfo(config.Ctx, viewOptions)
    if err != nil {
        fmt.Printf("Exception while calling InfoApi: %v\n", err)
        return
    }

    fmt.Printf("PrintingAllowed: %v\n", response.PdfViewInfo.PrintingAllowed)
    fmt.Printf("GetInfoForPdfFile completed: %v pages\n", len(response.Pages))
}