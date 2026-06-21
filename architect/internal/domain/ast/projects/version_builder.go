package projects

import "fmt"

type versionBuilder struct {
	number          string
	serviceFiles    []string
	objectFiles     []string
	deploymentFiles []string
	nextVersions    []Version
}

func (b *versionBuilder) Number(number string) VersionBuilder {
	b.number = number
	return b
}

func (b *versionBuilder) AddServiceFile(file string) VersionBuilder {
	if file != "" {
		b.serviceFiles = append(b.serviceFiles, file)
	}

	return b
}

func (b *versionBuilder) AddObjectFile(file string) VersionBuilder {
	if file != "" {
		b.objectFiles = append(b.objectFiles, file)
	}

	return b
}

func (b *versionBuilder) AddDeploymentFile(file string) VersionBuilder {
	if file != "" {
		b.deploymentFiles = append(b.deploymentFiles, file)
	}

	return b
}

func (b *versionBuilder) AddNextVersion(version Version) VersionBuilder {
	if version != nil {
		b.nextVersions = append(b.nextVersions, version)
	}

	return b
}

func (b *versionBuilder) Build() (Version, error) {
	if b.number == "" {
		return nil, fmt.Errorf("version number is required")
	}

	serviceFiles := make([]string, len(b.serviceFiles))
	copy(serviceFiles, b.serviceFiles)

	objectFiles := make([]string, len(b.objectFiles))
	copy(objectFiles, b.objectFiles)

	deploymentFiles := make([]string, len(b.deploymentFiles))
	copy(deploymentFiles, b.deploymentFiles)

	nextVersions := make([]Version, len(b.nextVersions))
	copy(nextVersions, b.nextVersions)

	return &version{
		number:          b.number,
		serviceFiles:    serviceFiles,
		objectFiles:     objectFiles,
		deploymentFiles: deploymentFiles,
		nextVersions:    nextVersions,
	}, nil
}
