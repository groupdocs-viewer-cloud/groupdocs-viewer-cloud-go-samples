package RenderingMsProjectDocuments

import (
    "fmt"

    "github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go-samples/config"
    viewer "github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go/models"
)

// This example demonstrates how to retrieve view information for MS Project document
func GetInfoForProjectFile() {
    viewOptions := viewer.ViewOptions{
        FileInfo: &viewer.FileInfo{
            FilePath: "SampleFiles/sample.mpp",
        },
    }

    response, _, err := config.Client.InfoApi.GetInfo(config.Ctx, viewOptions)
    if err != nil {
        fmt.Printf("Exception while calling InfoApi: %v\n", err)
        return
    }

    fmt.Printf("Start date: %v\n", response.ProjectManagementViewInfo.StartDate)
    fmt.Printf("End date: %v\n", response.ProjectManagementViewInfo.EndDate)
    fmt.Printf("GetInfoForProjectFile completed: %v pages\n", len(response.Pages))
}