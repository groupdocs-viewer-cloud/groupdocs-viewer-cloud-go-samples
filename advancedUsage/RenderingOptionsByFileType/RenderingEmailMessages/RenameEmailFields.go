package RenderingEmailMessages

import (
    "fmt"

    "github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go-samples/config"
    viewer "github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go/models"
)

// This example demonstrates how to use custom field labels
func RenameEmailFields() {
    viewOptions := viewer.ViewOptions{
        FileInfo: &viewer.FileInfo{
            FilePath: "SampleFiles/sample.msg",
        },
        ViewFormat: viewer.ViewFormatHtml,
        RenderOptions: &viewer.HtmlOptions{
            EmailOptions: &viewer.EmailOptions{
                FieldLabels: []viewer.FieldLabel{
                    {Field: "From", Label: "Sender"},
                    {Field: "To", Label: "Receiver"},
                    {Field: "Sent", Label: "Date"},
                    {Field: "Subject", Label: "Topic"},
                },
            },
        },
    }

    response, _, err := config.Client.ViewApi.CreateView(config.Ctx, viewOptions)
    if err != nil {
        fmt.Printf("Exception: %v\n", err)
        return
    }

    fmt.Printf("RenameEmailFields completed: %v pages\n", len(response.Pages))
}