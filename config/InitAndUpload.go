package config

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	viewer "github.com/groupdocs-viewer-cloud/groupdocs-viewer-cloud-go"
	"golang.org/x/oauth2/clientcredentials"
)

var Ctx context.Context
var Client *viewer.APIClient

func InitAndUpload(clientId string, clientSecret string, upload bool) {

	conf := &clientcredentials.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		TokenURL:     "https://api.groupdocs.cloud/connect/token",
		Scopes:       []string{},
	}

	// Request the token
	token, err := conf.Token(context.Background())
	if err != nil {
		fmt.Printf("Unable to retrieve token: %v\n", err)
	}

	// Output the access token
	fmt.Printf("Access Token: %s\n", token.AccessToken)

	cfg := viewer.NewConfiguration()
	Client = viewer.NewAPIClient(cfg)
	Ctx = context.WithValue(context.Background(), viewer.ContextAccessToken, token.AccessToken)

	if upload {
		// Define the path to the "TestData" folder
		resourcesFolder := "./Resources"

		// Walk through all files in the directory and subdirectories
		err = filepath.Walk(resourcesFolder, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				fmt.Printf("Error walking the path %s: %v\n", path, err)
				return err
			}

			// Only process regular files (not directories)
			if !info.IsDir() {
				// Get the relative path from the "TestData" folder
				relativePath, err := filepath.Rel(resourcesFolder, path)
				if err != nil {
					fmt.Printf("Failed to get relative path for file %s: %v\n", path, err)
					return nil // Continue with the next file
				}

				// Ensure the path has the leading slash (e.g., "/Note/sample.one")
				relativePath = "/" + strings.TrimPrefix(relativePath, "/")

				// Check if the file exists in the cloud storage using the relative path
				objectExistsResp, _, err := Client.StorageApi.ObjectExists(Ctx, relativePath, nil)
				if err != nil {
					fmt.Printf("Failed to check if file %s exists: %v\n", relativePath, err)
					return nil // Continue with the next file
				}

				// If the file does not exist in cloud storage, upload it
				if !objectExistsResp.Exists {
					// Open the local file
					localFile, err := os.Open(path)
					if err != nil {
						fmt.Printf("Failed to open file %s: %v\n", path, err)
						return nil // Continue with the next file
					}
					defer localFile.Close()

					// Upload the file to cloud storage with the relative path
					_, _, err = Client.FileApi.UploadFile(Ctx, relativePath, localFile, nil)
					if err != nil {
						fmt.Printf("Failed to upload file %s: %v\n", relativePath, err)
					} else {
						fmt.Printf("File %s uploaded successfully.\n", relativePath)
					}
				} else {
					fmt.Printf("File %s already exists in cloud storage.\n", relativePath)
				}
			}
			return nil
		})

		if err != nil {
			fmt.Printf("Error walking the directory: %v\n", err)
			os.Exit(1)
		}
	}
}
