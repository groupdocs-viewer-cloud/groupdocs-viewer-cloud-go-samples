package basicUsage

import (
	"fmt"

	"github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go-samples/config"
)

func GetSupportedFormats() {

	response, _, err := config.Client.InfoApi.GetSupportedFileFormats(config.Ctx)

	if err != nil {
		fmt.Printf("GetSupportedViewers error: %v\n", err)
		return
	}

	fmt.Printf("Formats num: %v\n", len(response.Formats))
}
