package RenderingArchiveFiles

import (
	"fmt"

	"github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go-samples/config"
	viewer "github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go/models"
)

// This example demonstrates how to print out archive contents
func GetInfoForArchiveFile() {

	opts := viewer.ViewOptions{
		FileInfo: &viewer.FileInfo{
			FilePath: "SampleFiles/with_folders.zip",
		},
	}

	response, _, err := config.Client.InfoApi.GetInfo(config.Ctx, opts)
	if err != nil {
		fmt.Printf("Exception while calling InfoApi: %v\n", err)
		return
	}

	for _, folder := range response.ArchiveViewInfo.Folders {
		fmt.Println(folder)
	}
	fmt.Printf("GetInfoForArchiveFile completed: %v pages\n", len(response.Pages))
}
