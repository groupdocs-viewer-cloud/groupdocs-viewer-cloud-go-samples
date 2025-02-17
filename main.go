//// ***********************************************************
////          GroupDocs.Viewer Cloud API Examples
//// ***********************************************************

package main

import (
	"github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go-samples/config"

	"github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go-samples/basicUsage"
	pdfViewer "github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go-samples/basicUsage/pdfViewer"
)

func main() {

	// TODO: Get your ClientId and ClientSecret at https://dashboard.groupdocs.cloud (free registration is required).
	config.InitAndUpload("Your_ClientId", "Your_ClientSecret", false)

	// Info API Examples
	basicUsage.GetDocumentInformation()
	basicUsage.GetSupportedFormats()

	// View API Examples
	pdfViewer.PdfViewerAdjustJpgQuality()
	pdfViewer.PdfViewerProtectPdf()

	// Check out other examples in this project to learn more about GroupDocs.Viewer Cloud API
}
