package client

import (
    "fmt"
)

func GetDownloadLink(token string, filePath string) string {
    return fmt.Sprintf("%s/file/bot%s/%s", baseUrl, token, filePath)
}
