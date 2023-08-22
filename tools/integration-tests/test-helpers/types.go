package testhelpers

// GhpcDeployment represents a GCP deployment using the HPC Toolkit.
type GhpcDeployment interface {
	GetName() string
	GetProject() string
	GetRegion() string
	GetZone() string
	GetBlueprint() string
	SetBlueprint(string) error
	Redeploy() error
}

var _ GhpcDeployment = &ghpcDeployment{}

type ghpcDeployment struct {
	blueprint      string
	deploymentName string
	project        string
	region         string
	zone           string
}

// GcpResource represents a GCP resource.
type GcpResource interface {
	GetName() string
	GetType() string
	GetRegion() string
	GetZone() string
	GetTags() []string
}
