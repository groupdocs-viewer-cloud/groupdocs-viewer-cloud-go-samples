package CommonRenderingOptions

import (
	"fmt"

	"github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go-samples/config"
	viewer "github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go/models"
)

// This example demonstrates how to rotate output pages when viewing a document as PDF
func FlipRotatePages() {
	viewOptions := viewer.ViewOptions{
		FileInfo: &viewer.FileInfo{
			FilePath: "SampleFiles/sample.docx",
		},
		ViewFormat: viewer.ViewFormatPdf,
		RenderOptions: &viewer.PdfOptions{
			PageRotations: []viewer.PageRotation{
				{
					PageNumber:    1,
					RotationAngle: viewer.RotationAngleOn90Degree,
				},
			},
		},
	}

	response, _, err := config.Client.ViewApi.CreateView(config.Ctx, viewOptions)
	if err != nil {
		fmt.Printf("Exception: %v\n", err)
		return
	}

	fmt.Printf("FlipRotatePages completed: %v\n", response.File.Path)
}
