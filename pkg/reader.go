package reader

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type EntraCredentials struct {
	TenantId     string
	ClientId     string
	ClientSecret string
	Scope        string
}

func ReadEntraCredentials(app string) (*EntraCredentials, error) {
	filePath, err := getCredentialsFilePath()
	if err != nil {
		return nil, fmt.Errorf("failed to get credentials file path: %v", err)
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var currentApp string
	var credentials *EntraCredentials

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			currentApp = line[1 : len(line)-1]
		} else if currentApp == app {
			if strings.HasPrefix(line, "tenantId=") {
				if credentials == nil {
					credentials = &EntraCredentials{}
				}
				credentials.TenantId = strings.TrimSpace(strings.SplitN(line, "=", 2)[1])
			} else if strings.HasPrefix(line, "clientId=") {
				if credentials == nil {
					credentials = &EntraCredentials{}
				}
				credentials.ClientId = strings.TrimSpace(strings.SplitN(line, "=", 2)[1])
			} else if strings.HasPrefix(line, "clientSecret=") {
				if credentials == nil {
					credentials = &EntraCredentials{}
				}
				credentials.ClientSecret = strings.TrimSpace(strings.SplitN(line, "=", 2)[1])
			} else if strings.HasPrefix(line, "scope=") {
				if credentials == nil {
					credentials = &EntraCredentials{}
				}
				credentials.Scope = strings.TrimSpace(strings.SplitN(line, "=", 2)[1])
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	if credentials == nil {
		return nil, fmt.Errorf("app %s not found in credentials file", app)
	}

	return credentials, nil
}

func getCredentialsFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home directory: %v", err)
	}

	credentialsFilePath := filepath.Join(homeDir, ".entra", "credentials")
	return credentialsFilePath, nil
}
