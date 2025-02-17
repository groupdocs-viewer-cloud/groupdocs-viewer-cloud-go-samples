package RenderingArchiveFiles

import (
    "fmt"

    "github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go-samples/config"
    viewer "github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go/models"
)

// This example demonstrates how to render an archive folder
func RenderArchiveFolder() {
    viewOptions := viewer.ViewOptions{
        FileInfo: &viewer.FileInfo{
            FilePath: "SampleFiles/with_folders.zip",
        },
        RenderOptions: &viewer.RenderOptions{
            ArchiveOptions: &viewer.ArchiveOptions{
                Folder: "ThirdFolderWithItems",
            },
        },
    }

    response, _, err := config.Client.ViewApi.CreateView(config.Ctx, viewOptions)
    if err != nil {
        fmt.Printf("Exception: %v\n", err)
        return
    }

    fmt.Printf("RenderArchiveFolder completed: %v pages\n", len(response.Pages))
}