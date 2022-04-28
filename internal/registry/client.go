package exporter

import (
	"fmt"
	"io"
	"net/http"
)

// Health check registry server
func Ping(url string) (bool, error) {
	srv := fmt.Sprintf("http://%s/v2", url)
	resp, err := http.Get(srv)

	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return false, err
		}
		bodyString := string(bodyBytes)
		fmt.Printf("Ping Check: %s\n", bodyString)
	}
	return true, nil
}

// image list(no have image tags)
func Catalog(url string) (string, error) {
	srv := fmt.Sprintf("http://%s/v2/_catalog", url)
	resp, err := http.Get(srv)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		bodyString := string(bodyBytes)
		fmt.Printf("Catalogs: %s\n", bodyString)
		return bodyString, nil
	}
	return "", nil
}

// image list tags
func ListTags(url string, image string) (string, error) {
	srv := fmt.Sprintf("http://%s/v2/%s/tags/list", url, image)
	resp, err := http.Get(srv)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		bodyString := string(bodyBytes)
		// fmt.Printf("Image Lists: %s", bodyString)
		return bodyString, nil
	}
	return "", nil
}
