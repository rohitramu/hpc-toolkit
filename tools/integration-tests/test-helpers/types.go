package testhelpers

// TODO: 1. Implement deploy blueprint and ensure deployment name is correct.
// TODO: 2. Implement destroy blueprint.
// TODO: 3. Implement GetResources().
// TODO: 4. Create a library of helper methods for interacting with particular types of resources.
//          E.g. VMs, Storage buckets, VM Images
// TODO: 5. Reimplement Ansible tests with this framework.

// GhpcDeployment represents a GCP deployment using the HPC Toolkit.
type GhpcDeployment interface {
	GetName() string
	GetProject() string
	GetRegion() string
	GetZone() string
	GetBlueprint() string
	SetBlueprint(string) error
	Redeploy() error
	GetResources() ([]GcpResource, error)
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
