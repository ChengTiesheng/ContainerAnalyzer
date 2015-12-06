package analyse

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/chengtiesheng/ContainerAnalyzer/attr"
	"github.com/chengtiesheng/ContainerAnalyzer/attr/local"
	"github.com/chengtiesheng/ContainerAnalyzer/attr/registry"
)

type AnalyseBackend interface {
	GetImageInfo(dockerUrl string) ([]string, *attr.ParsedDockerURL, error)
	GetLayerInfo(layerID string, dockerURL *attr.ParsedDockerURL) (*attr.DockerImageData, error)
}

func AnalyseDockerImage(arg string, flagImage string, flagDebug bool, flagInsecure bool) error {
	if flagDebug {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
	}

	// try to analyse a local file
	u, err := url.Parse(arg)
	if err != nil {
		return fmt.Errorf("error parsing argument: %v", err)
	}
	if u.Scheme == "docker" {
		if flagImage != "" {
			return fmt.Errorf("flag --image works only with files.")
		}
		dockerURL := strings.TrimPrefix(arg, "docker://")

		indexServer := GetIndexName(dockerURL)

		var username, password string
		username, password, err = GetDockercfgAuth(indexServer)
		if err != nil {
			return fmt.Errorf("error reading .dockercfg file: %v", err)
		}

		err = AnalyseRegistry(dockerURL, username, password, flagInsecure)
	} else {
		err = AnalyseLocal(flagImage, arg)
	}
	if err != nil {
		return fmt.Errorf("analyse error: %v", err)
	}

	return nil
}

func AnalyseRegistry(dockerURL string, username string, password string, insecure bool) error {
	repositoryBackend := registry.NewRepositoryBackend(username, password, insecure)
	return analyseReal(repositoryBackend, dockerURL)
}

func AnalyseLocal(dockerURL string, filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer f.Close()

	fileBackend := local.NewFileBackend(f)
	return analyseReal(fileBackend, dockerURL)
}

// GetIndexName returns the docker index server from a docker URL.
func GetIndexName(dockerURL string) string {
	index, _ := attr.SplitReposName(dockerURL)
	return index
}

// GetDockercfgAuth reads a ~/.dockercfg file and returns the username and password
// of the given docker index server.
func GetDockercfgAuth(indexServer string) (string, string, error) {
	return attr.GetAuthInfo(indexServer)
}

func analyseReal(backend AnalyseBackend, dockerURL string) error {
	fmt.Println("Getting image attribution information...")
	ancestry, parsedDockerURL, err := backend.GetImageInfo(dockerURL)
	if err != nil {
		return err
	}

	var bPrinted bool
	for i := len(ancestry) - 1; i >= 0; i-- {
		layerID := ancestry[i]

		layerData, _ := backend.GetLayerInfo(layerID, parsedDockerURL)

		imgAttr, _ := attr.AnalyseDockerManifest(*layerData, parsedDockerURL)
		if !bPrinted {
			fmt.Printf("\nPrompt: Layer0 is the upper layer and layer%d is the bottom layer\n", i)
			bPrinted = true
		}

		fmt.Printf("\n============Layer%v================\n", i)
		printDockerImgAttr(imgAttr)
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

func printDockerImgAttr(imgAttr *attr.DockerImg_Attr) {
	fmt.Println("Type:     ", imgAttr.Type)
	fmt.Println("Layer:    ", imgAttr.Layer)
	fmt.Println("Name:     ", imgAttr.Name)
	fmt.Println("Version:  ", imgAttr.Version)
	fmt.Println("OS:       ", imgAttr.OS)
	fmt.Println("Arch:     ", imgAttr.Arch)
	fmt.Println("Author:   ", imgAttr.Author)
	fmt.Println("Epoch:    ", imgAttr.Epoch)
	fmt.Println("Comment:  ", imgAttr.Comment)
	fmt.Println("Parent:   ", imgAttr.Parent)
	printApp(imgAttr.App)

	return
}

func printApp(app attr.App) {
	execs := app.Exec
	if len(execs) > 0 {
		fmt.Printf("Exec:\n")
		for _, exec := range execs {
			fmt.Printf("\targs: %v\n", exec)
		}
	}

	mps := app.MountPoints
	if len(mps) > 0 {
		fmt.Printf("MountPoints:\n")
		for _, mp := range mps {
			fmt.Printf("\tname: %q, path: %q, readOnly: %v\n", mp.Name, mp.Path, mp.ReadOnly)
		}
	}

	ports := app.Ports
	if len(ports) > 0 {
		fmt.Printf("Ports:\n")
		for _, port := range ports {
			fmt.Printf("\tname: %q, protocol: %q, port: %v\n", port.Name, port.Protocol, port.Port)

		}
	}

	return
}
