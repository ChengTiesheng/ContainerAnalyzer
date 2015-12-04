package attr

import (
	"encoding/json"
	"strings"
	"time"
)

type DockerImg_Attr struct {
	Type    string `json:"type"` // required= "dockerimage"
	Layer   string `json:"layer"`
	Name    string `json:"name"`
	Version string `json:"version"`
	OS      string `json:"os"`
	Arch    string `json:"arch"`
	Author  string `json:"author"`
	Epoch   string `json:"epoch"`
	Comment string `json:"comment"`
	Parent  string `json:"parent"`
	App     App    `json:"app"`
}

func (imgAttr *DockerImg_Attr) GetAttr(imgData *DockerImageData) error {
	//imgAttr.Type = "dockerimage"
	//imgAttr.Id = imgData.ID

	return nil
}

func (imgAttr *DockerImg_Attr) GetAttrFromJson(j []byte) error {
	layerData := DockerImageData{}
	if err := json.Unmarshal(j, &layerData); err != nil {
		return err
	}

	//imgAttr.Id = layerData.ID
	imgAttr.Name = layerData.Author

	return nil
}

func getDockerImgAttr(imgData *DockerImageData, imgAttr *DockerImg_Attr) error {
	return nil
}

type Exec []string

type Port struct {
	Name     string `json:"name"`
	Protocol string `json:"protocol"`
	Port     uint   `json:"port"`
}

type ParsedDockerURL struct {
	IndexURL  string
	ImageName string
	Tag       string
}

type EventHandler struct {
	Name string `json:"name"`
	Exec Exec   `json:"exec"`
}

type Hash struct {
	typ string
	Val string
}

type Labels []Label

type labels Labels

type Label struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Dependencies []Dependency

type Dependency struct {
	ImageName string `json:"imageName"`
	ImageID   *Hash  `json:"imageID,omitempty"`
	Labels    Labels `json:"labels,omitempty"`
	Size      uint   `json:"size,omitempty"`
}

type App struct {
	Exec              Exec           `json:"exec"`
	EventHandlers     []EventHandler `json:"eventHandlers,omitempty"`
	User              string         `json:"user"`
	Group             string         `json:"group"`
	SupplementaryGIDs []int          `json:"supplementaryGIDs,omitempty"`
	WorkingDirectory  string         `json:"workingDirectory,omitempty"`
	Environment       Environment    `json:"environment,omitempty"`
	MountPoints       []MountPoint   `json:"mountPoints,omitempty"`
	Ports             []Port         `json:"ports,omitempty"`
}

type MountPoint struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	ReadOnly bool   `json:"readOnly,omitempty"`
}

type Environment []EnvironmentVariable

type EnvironmentVariable struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (e *Environment) Set(name string, value string) {
	for i, env := range *e {
		if env.Name == name {
			(*e)[i] = EnvironmentVariable{
				Name:  name,
				Value: value,
			}
			return
		}
	}
	env := EnvironmentVariable{
		Name:  name,
		Value: value,
	}
	*e = append(*e, env)
}

func AnalyseDockerManifest(layerData DockerImageData, dockerURL *ParsedDockerURL) (*DockerImg_Attr, error) {
	dockerConfig := layerData.Config
	imgAttr := &DockerImg_Attr{}

	imgAttr.Type = "dockerimage"

	imgAttr.Layer = layerData.ID

	url := dockerURL.IndexURL + "/"
	url += dockerURL.ImageName + "-" + layerData.ID
	imgAttr.Name = url

	tag := dockerURL.Tag
	imgAttr.Version = tag

	if layerData.OS != "" {
		os := layerData.OS
		imgAttr.OS = os
		if layerData.Architecture != "" {
			arch := layerData.Architecture
			imgAttr.Arch = arch
		}
	}

	if layerData.Author != "" {
		author := layerData.Author
		imgAttr.Author = author
	}

	epoch := time.Unix(0, 0)
	if !layerData.Created.Equal(epoch) {
		createdTime := layerData.Created.Format(time.RFC3339)
		imgAttr.Epoch = createdTime
	}

	if layerData.Comment != "" {
		comment := layerData.Comment
		imgAttr.Comment = comment
	}

	if dockerConfig != nil {
		exec := getExecCommand(dockerConfig.Entrypoint, dockerConfig.Cmd)
		if exec != nil {
			user, group := parseDockerUser(dockerConfig.User)
			var env Environment
			for _, v := range dockerConfig.Env {
				parts := strings.SplitN(v, "=", 2)
				env.Set(parts[0], parts[1])
			}
			app := &App{
				Exec:             exec,
				User:             user,
				Group:            group,
				Environment:      env,
				WorkingDirectory: dockerConfig.WorkingDir,
			}
			app.MountPoints, _ = convertVolumesToMPs(dockerConfig.Volumes)
			app.Ports, _ = getPorts(dockerConfig.ExposedPorts, dockerConfig.PortSpecs)

			imgAttr.App = *app
		}
	}

	if layerData.Parent != "" {
		indexPrefix := ""
		indexPrefix = dockerURL.IndexURL + "/"
		parentImageName := indexPrefix + dockerURL.ImageName + "-" + layerData.Parent
		imgAttr.Parent = parentImageName
	}

	return imgAttr, nil
}