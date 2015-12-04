package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/appc/docker2aci/lib/common"
	"github.com/chengtiesheng/ContainerAnalyzer/attr"
	"github.com/chengtiesheng/ContainerAnalyzer/attr/file"
	"github.com/chengtiesheng/ContainerAnalyzer/attr/repository"
)

type AnalyseBackend interface {
	GetImageInfo(dockerUrl string) ([]string, *attr.ParsedDockerURL, error)
	GetLayerInfo(layerID string, dockerURL *attr.ParsedDockerURL) (*attr.DockerImg_Attr, error)
}

func Analyse(dockerURL string, username string, password string, insecure bool) error {
	repositoryBackend := repository.NewRepositoryBackend(username, password, insecure)
	return analyseReal(repositoryBackend, dockerURL)
}

func AnalyseFile(dockerURL string, filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer f.Close()

	fileBackend := file.NewFileBackend(f)
	return convertReal(fileBackend, dockerURL)
}

// GetIndexName returns the docker index server from a docker URL.
func GetIndexName(dockerURL string) string {
	index, _ := common.SplitReposName(dockerURL)
	return index
}

// GetDockercfgAuth reads a ~/.dockercfg file and returns the username and password
// of the given docker index server.
func GetDockercfgAuth(indexServer string) (string, string, error) {
	return common.GetAuthInfo(indexServer)
}

func analyseReal(backend AnalyseBackend, dockerURL string) error {
	fmt.Println("Getting image info...")
	ancestry, parsedDockerURL, err := backend.GetImageInfo(dockerURL)
	if err != nil {
		return nil, err
	}

	for i := len(ancestry) - 1; i >= 0; i-- {
		layerID := ancestry[i]

		layerData, _ := backend.GetLayerInfo(layerID)

		imgAttr := attr.AnalyseDockerManifest(layerData, dockerURL)

		fmt.Println(imgAttr)

	}

	return nil
}

// striplayerID strips the layer ID from an app name:
//
// myregistry.com/organization/app-name-85738f8f9a7f1b04b5329c590ebcb9e425925c6d0984089c43a022de4f19c281
// myregistry.com/organization/app-name
func stripLayerID(layerName string) string {
	n := strings.LastIndex(layerName, "-")
	return layerName[:n]
}
