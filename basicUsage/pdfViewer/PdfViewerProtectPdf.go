package pdfViewer

import (
	"fmt"

	"github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go-samples/config"
	viewer "github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go/models"
)

func PdfViewerProtectPdf() {
	viewOptions := viewer.ViewOptions{
		FileInfo: &viewer.FileInfo{
			FilePath: "SampleFiles/sample.docx",
		},
		ViewFormat: viewer.ViewFormatPdf,
		RenderOptions: &viewer.PdfOptions{
			Permissions:          []string{"DenyModification"},
			PermissionsPassword:  "p123",
			DocumentOpenPassword: "o123",
		},
	}

	response, _, err := config.Client.ViewApi.CreateView(config.Ctx, viewOptions)
	if err != nil {
		fmt.Printf("Exception: %v\n", err)
		return
	}

	fmt.Printf("PdfViewerProtectPdf completed: %v\n", response.File.Path)
}
