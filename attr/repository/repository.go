package repository

import "github.com/chengtiesheng/ContainerAnalyzer/attr"

type RepositoryBackend struct {
	repoData          *RepoData
	username          string
	password          string
	insecure          bool
	hostsV2Support    map[string]bool
	hostsV2AuthTokens map[string]map[string]string
	imageManifests    map[attr.ParsedDockerURL]v2Manifest
}

func NewRepositoryBackend(username string, password string, insecure bool) *RepositoryBackend {
	return &RepositoryBackend{
		username:          username,
		password:          password,
		insecure:          insecure,
		hostsV2Support:    make(map[string]bool),
		hostsV2AuthTokens: make(map[string]map[string]string),
		imageManifests:    make(map[attr.ParsedDockerURL]v2Manifest),
	}
}

func (rb *RepositoryBackend) GetImageInfo(url string) ([]string, *attr.ParsedDockerURL, error) {
	dockerURL := attr.ParseDockerURL(url)

	var supportsV2, ok bool
	if supportsV2, ok = rb.hostsV2Support[dockerURL.IndexURL]; !ok {
		var err error
		supportsV2, err = rb.supportsV2(dockerURL.IndexURL)
		if err != nil {
			return nil, nil, err
		}
		rb.hostsV2Support[dockerURL.IndexURL] = supportsV2
	}

	if supportsV2 {
		return rb.getImageInfoV2(dockerURL)
	} else {
		return rb.getImageInfoV1(dockerURL)
	}
}

func (rb *RepositoryBackend) protocol() string {
	if rb.insecure {
		return "http://"
	} else {
		return "https://"
	}
}

func (rb *RepositoryBackend) GetLayerInfo(layerID string, dockerURL *attr.ParsedDockerURL) (*attr.DockerImageData, error) {
	if rb.hostsV2Support[dockerURL.IndexURL] {
		return rb.getLayerInfoV2(layerID, dockerURL)
	} else {
		return rb.getLayerInfoV1(layerID, dockerURL)
	}
}
