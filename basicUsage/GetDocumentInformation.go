package basicUsage

import (
	"fmt"

	"github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go-samples/config"
	viewer "github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go/models"
)

func GetDocumentInformation() {

	opts := viewer.ViewOptions{
		FileInfo: &viewer.FileInfo{
			FilePath: "SampleFiles/sample.docx",
		},
	}

	info, _, err := config.Client.InfoApi.GetInfo(config.Ctx, opts)

	if err != nil {
		fmt.Printf("GetInfo error: %v\n", err)
		return
	}

	fmt.Printf("Page count: %v\n", len(info.Pages))
}
