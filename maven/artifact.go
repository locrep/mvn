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

type ArtifactRepos []ArtifactRepo

type ArtifactRepo struct {
	Repo  string   `json:"repo"`
	Files []string `json:"files"`
}
