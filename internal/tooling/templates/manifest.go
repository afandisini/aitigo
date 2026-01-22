package templates

type Manifest struct {
	ID             string        `json:"id"`
	Name           string        `json:"name"`
	Description    string        `json:"description"`
	Type           string        `json:"type"`
	PackageManager string        `json:"packageManager"`
	Init           ManifestInit  `json:"init"`
	Check          ManifestCheck `json:"check"`
}

type ManifestInit struct {
	Copy bool `json:"copy"`
}

type ManifestCheck struct {
	Commands []CheckCommand `json:"commands"`
}

type CheckCommand struct {
	Name     string `json:"name"`
	Run      string `json:"run"`
	Optional bool   `json:"optional"`
}
