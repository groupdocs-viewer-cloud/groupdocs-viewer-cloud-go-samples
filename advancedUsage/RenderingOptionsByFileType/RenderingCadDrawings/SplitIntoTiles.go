package RenderingCadDrawings

import (
    "fmt"

    "github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go-samples/config"
    viewer "github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go/models"
)

// This example demonstrates how to render DWG drawing into an image by dividing it into four equal parts
func SplitIntoTiles() {
    // Get document information
    viewOptions := viewer.ViewOptions{
        FileInfo: &viewer.FileInfo{
            FilePath: "SampleFiles/with_layers_and_layouts.dwg",
        },
    }

    infoResponse, _, err := config.Client.InfoApi.GetInfo(config.Ctx, viewOptions)
    if err != nil {
        fmt.Printf("Exception while calling InfoApi: %v\n", err)
        return
    }

    // Get width and height
    width := infoResponse.Pages[0].Width
    height := infoResponse.Pages[0].Height

    // Set tile width and height as a half of image total width
    tileWidth := width / 2
    tileHeight := height / 2
    pointX := int32(0)
    pointY := int32(0)

    // Create image options and add four tiles, one for each quarter
    viewOptions = viewer.ViewOptions{
        FileInfo: &viewer.FileInfo{
            FilePath: "SampleFiles/with_layers_and_layouts.dwg",
        },
        ViewFormat: viewer.ViewFormatHtml,
        RenderOptions: &viewer.HtmlOptions{
            CadOptions: &viewer.CadOptions{
                Tiles: []viewer.Tile{
                    {StartPointX: pointX, StartPointY: pointY, Width: tileWidth, Height: tileHeight},
                    {StartPointX: pointX + tileWidth, StartPointY: pointY, Width: tileWidth, Height: tileHeight},
                    {StartPointX: pointX, StartPointY: pointY + tileHeight, Width: tileWidth, Height: tileHeight},
                    {StartPointX: pointX + tileWidth, StartPointY: pointY + tileHeight, Width: tileWidth, Height: tileHeight},
                },
            },
        },
    }

    response, _, err := config.Client.ViewApi.CreateView(config.Ctx, viewOptions)
    if err != nil {
        fmt.Printf("Exception: %v\n", err)
        return
    }

    fmt.Printf("SplitIntoTiles completed: %v pages\n", len(response.Pages))
}