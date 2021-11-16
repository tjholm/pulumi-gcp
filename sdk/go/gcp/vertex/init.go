// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vertex

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "gcp:vertex/aiDataset:AiDataset":
		r = &AiDataset{}
	case "gcp:vertex/aiFeatureStore:AiFeatureStore":
		r = &AiFeatureStore{}
	case "gcp:vertex/aiFeatureStoreEntityType:AiFeatureStoreEntityType":
		r = &AiFeatureStoreEntityType{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := gcp.PkgVersion()
	if err != nil {
		fmt.Printf("failed to determine package version. defaulting to v1: %v\n", err)
	}
	pulumi.RegisterResourceModule(
		"gcp",
		"vertex/aiDataset",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"vertex/aiFeatureStore",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"vertex/aiFeatureStoreEntityType",
		&module{version},
	)
}
