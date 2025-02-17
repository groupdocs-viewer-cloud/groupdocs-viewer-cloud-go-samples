package RenderingCadDrawings

import (
    "fmt"

    "github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go-samples/config"
    viewer "github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go/models"
)

// This example demonstrates how to retrieve view information for CAD drawing
func GetInfoForCadDrawing() {
    viewOptions := viewer.ViewOptions{
        FileInfo: &viewer.FileInfo{
            FilePath: "SampleFiles/with_layers_and_layouts.dwg",
        },
    }

    response, _, err := config.Client.InfoApi.GetInfo(config.Ctx, viewOptions)
    if err != nil {
        fmt.Printf("Exception while calling InfoApi: %v\n", err)
        return
    }

    fmt.Printf("Layers count: %v\n", len(response.CadViewInfo.Layers))
    fmt.Printf("Layouts count: %v\n", len(response.CadViewInfo.Layouts))
    fmt.Printf("GetInfoForCadDrawing completed: %v pages\n", len(response.Pages))
}