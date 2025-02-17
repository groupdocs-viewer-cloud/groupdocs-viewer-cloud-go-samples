package basicUsage

import (
	"fmt"
	"os"

	"github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go-samples/config"
)

func ConvertAndDownload() {
	// Open the file to be converted
	file, err := os.Open("Resources/SampleFiles/sample.docx")
	if err != nil {
		fmt.Printf("Exception: %v\n", err)
		return
	}
	defer file.Close()

	// Call the API
	response, _, err := config.Client.ViewApi.ConvertAndDownload(config.Ctx, "jpg", file, nil)
	if err != nil {
		fmt.Printf("Exception: %v\n", err)
		return
	}

	fileInfo, err := response.Stat()
	if err != nil {
		fmt.Printf("Exception: %v\n", err)
		return
	}
	fmt.Printf("ConvertAndDownload completed: %v bytes received\n", fileInfo.Size())
}
