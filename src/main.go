package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	reader "github.com/filipemcg/entra/pkg"
)

func getClientCredentialsToken(clientId, clientSecret, tokenURL, scope string) (string, error) {
	// Create the request body
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("client_id", clientId)
	data.Set("client_secret", clientSecret)
	data.Set("scope", scope)

	// Create the request
	req, err := http.NewRequest("POST", tokenURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to make request: %v", err)
	}
	defer resp.Body.Close()

	// Check for HTTP errors
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get token: %v", resp.Status)
	}

	// Parse the response
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to parse response: %v", err)
	}

	// Extract the access token
	token, ok := result["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("failed to get access token from response")
	}

	return token, nil
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Please provide a command")
		os.Exit(1)
	}

	switch args[1] {
	case "--app":
		credentials, err := reader.ReadEntraCredentials(os.Args[2])
		if err != nil {
			fmt.Printf("Failed to read credentials: %v\n", err)
			os.Exit(1)
		}

		tokenURL := fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/v2.0/token", credentials.TenantId)
		clientId := credentials.ClientId
		clientSecret := credentials.ClientSecret
		scope := credentials.Scope
		token, err := getClientCredentialsToken(clientId, clientSecret, tokenURL, scope)
		if err != nil {
			fmt.Printf("Failed to get token: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(token)
	default:
		fmt.Println("Invalid command")
	}
}
