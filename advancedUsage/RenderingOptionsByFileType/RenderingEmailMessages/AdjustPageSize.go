package RenderingEmailMessages

import (
	"fmt"

	"github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go-samples/config"
	viewer "github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go/models"
)

// This example demonstrates how to adjust output page size
func AdjustPageSize() {
	viewOptions := viewer.ViewOptions{
		FileInfo: &viewer.FileInfo{
			FilePath: "SampleFiles/sample.msg",
		},
		ViewFormat: viewer.ViewFormatPdf,
		RenderOptions: &viewer.PdfOptions{
			EmailOptions: &viewer.EmailOptions{
				PageSize: viewer.PageSizeA4,
			},
		},
	}

	response, _, err := config.Client.ViewApi.CreateView(config.Ctx, viewOptions)
	if err != nil {
		fmt.Printf("Exception: %v\n", err)
		return
	}

	fmt.Printf("AdjustPageSize completed: %v\n", response.File.Path)
}
