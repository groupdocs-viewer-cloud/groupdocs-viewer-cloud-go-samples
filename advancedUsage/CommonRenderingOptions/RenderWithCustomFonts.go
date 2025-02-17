package CommonRenderingOptions

import (
	"fmt"

	"github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go-samples/config"
	viewer "github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go/models"
)

// This example demonstrates how to set a custom font source when rendering documents
func RenderWithCustomFonts() {
	viewOptions := viewer.ViewOptions{
		FileInfo: &viewer.FileInfo{
			FilePath: "SampleFiles/with_missing_font.odg",
		},
		ViewFormat: viewer.ViewFormatHtml,
		FontsPath:  "Fonts", // NOTE: Upload fonts to the folder using storage API before rendering
	}

	response, _, err := config.Client.ViewApi.CreateView(config.Ctx, viewOptions)
	if err != nil {
		fmt.Printf("Exception: %v\n", err)
		return
	}

	fmt.Printf("RenderWithCustomFonts completed: %v pages\n", len(response.Pages))
}
