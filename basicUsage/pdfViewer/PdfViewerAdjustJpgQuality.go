package pdfViewer

import (
	"fmt"

	"github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go-samples/config"
	viewer "github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go/models"
)

func PdfViewerAdjustJpgQuality() {
	viewOptions := viewer.ViewOptions{
		FileInfo: &viewer.FileInfo{
			FilePath: "SampleFiles/with_jpg_image.pptx",
		},
		ViewFormat: viewer.ViewFormatPdf,
		RenderOptions: &viewer.PdfOptions{
			PdfOptimizationOptions: &viewer.PdfOptimizationOptions{
				CompressImages: true,
				ImageQuality:   50,
			},
		},
	}

	response, _, err := config.Client.ViewApi.CreateView(config.Ctx, viewOptions)
	if err != nil {
		fmt.Printf("Exception: %v\n", err)
		return
	}

	fmt.Printf("PdfViewerProtectPdf completed: %v\n", response.File.Path)
}
