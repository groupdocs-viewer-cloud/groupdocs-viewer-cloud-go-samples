package CommonRenderingOptions

import (
	"fmt"

	"github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go-samples/config"
	viewer "github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go/models"
)

// This example demonstrates how to render hidden pages
func RenderHiddenPages() {
	viewOptions := viewer.ViewOptions{
		FileInfo: &viewer.FileInfo{
			FilePath: "SampleFiles/with_hidden_page.pptx",
		},
		ViewFormat: viewer.ViewFormatHtml,
		RenderOptions: &viewer.RenderOptions{
			RenderHiddenPages: true,
		},
	}

	response, _, err := config.Client.ViewApi.CreateView(config.Ctx, viewOptions)
	if err != nil {
		fmt.Printf("Exception: %v\n", err)
		return
	}

	fmt.Printf("RenderHiddenPages completed: %v pages\n", len(response.Pages))
}
