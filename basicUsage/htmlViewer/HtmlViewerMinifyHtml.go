package htmlViewer

import (
	"fmt"

	"github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go-samples/config"
	viewer "github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go/models"
)

func HtmlViewerMinifyHtml() {
	viewOptions := viewer.ViewOptions{
		FileInfo: &viewer.FileInfo{
			FilePath: "SampleFiles/sample.docx",
		},
		ViewFormat: viewer.ViewFormatHtml,
		RenderOptions: &viewer.HtmlOptions{
			Minify: true,
		},
	}

	response, _, err := config.Client.ViewApi.CreateView(config.Ctx, viewOptions)
	if err != nil {
		fmt.Printf("Exception: %v\n", err)
		return
	}

	fmt.Printf("HtmlViewerMinifyHtml completed: %v\n", response.File.Path)
}
