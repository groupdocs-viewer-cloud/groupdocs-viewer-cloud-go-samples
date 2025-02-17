package RenderingMsProjectDocuments

import (
    "fmt"
    "time"

    "github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go-samples/config"
    viewer "github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go/models"
)

// This example demonstrates how to render MS Project document using StartDate and EndDate
func RenderTimeInterval() {
    viewOptions := viewer.ViewOptions{
        FileInfo: &viewer.FileInfo{
            FilePath: "SampleFiles/sample.mpp",
        },
        ViewFormat: viewer.ViewFormatHtml,
        RenderOptions: &viewer.HtmlOptions{
            ProjectManagementOptions: &viewer.ProjectManagementOptions{
                StartDate: time.Date(2008, 6, 1, 0, 0, 0, 0, time.UTC),
                EndDate:   time.Date(2008, 7, 1, 0, 0, 0, 0, time.UTC),
            },
        },
    }

    response, _, err := config.Client.ViewApi.CreateView(config.Ctx, viewOptions)
    if err != nil {
        fmt.Printf("Exception: %v\n", err)
        return
    }

    fmt.Printf("RenderTimeInterval completed: %v pages\n", len(response.Pages))
}