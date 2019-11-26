package maven

type Artifacts []Artifact

type Artifact struct {
	GroupID    string `json:"groupId"`
	ArtifactID string `json:"artifactId"`
	Version    string `json:"version"`
}

//return artifact path like groupId/artifactId/version
func (a Artifact) String() string {
	return "/" + a.GroupID + "/" + a.ArtifactID + "/" + a.Version
}
