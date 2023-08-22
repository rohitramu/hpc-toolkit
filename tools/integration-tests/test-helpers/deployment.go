package testhelpers

import "errors"

// DeployBlueprint deploys a GHPC blueprint and returns an object which can be used to interact with it.
func DeployBlueprint(
	name string,
	project string,
	region string,
	zone string,
	blueprint string,
) (GhpcDeployment, error) {
	var result = &ghpcDeployment{
		name,
		project,
		region,
		zone,
		blueprint,
	}

	var err = result.Redeploy()

	return result, err
}

func (deployment *ghpcDeployment) GetName() string {
	return deployment.deploymentName
}

func (deployment *ghpcDeployment) GetProject() string {
	return deployment.project
}

func (deployment *ghpcDeployment) GetRegion() string {
	return deployment.region
}

func (deployment *ghpcDeployment) GetZone() string {
	return deployment.zone
}

func (deployment *ghpcDeployment) GetBlueprint() string {
	return deployment.blueprint
}

func (deployment *ghpcDeployment) SetBlueprint(blueprint string) error {
	if blueprint == "" {
		return errors.New("provided blueprint is empty")
	}

	deployment.blueprint = blueprint

	return nil
}

func (deployment *ghpcDeployment) Redeploy() error {
	return errors.New("not yet implemented")
}
