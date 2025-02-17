package RenderingOutlookDataFiles

import (
    "fmt"

    "github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go-samples/config"
    viewer "github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go/models"
)

// This example demonstrates how to retrieve view information for an Outlook data file
func GetInfoForOutlookDataFile() {
    viewOptions := viewer.ViewOptions{
        FileInfo: &viewer.FileInfo{
            FilePath: "SampleFiles/sample.ost",
        },
    }

    response, _, err := config.Client.InfoApi.GetInfo(config.Ctx, viewOptions)
    if err != nil {
        fmt.Printf("Exception while calling InfoApi: %v\n", err)
        return
    }

    fmt.Printf("Folders count: %v\n", len(response.OutlookViewInfo.Folders))
    fmt.Printf("GetInfoForOutlookDataFile completed: %v pages\n", len(response.Pages))
}