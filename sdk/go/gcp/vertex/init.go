// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

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
	case "gcp:vertex/aiEndpoint:AiEndpoint":
		r = &AiEndpoint{}
	case "gcp:vertex/aiFeatureStore:AiFeatureStore":
		r = &AiFeatureStore{}
	case "gcp:vertex/aiFeatureStoreEntityType:AiFeatureStoreEntityType":
		r = &AiFeatureStoreEntityType{}
	case "gcp:vertex/aiFeatureStoreEntityTypeFeature:AiFeatureStoreEntityTypeFeature":
		r = &AiFeatureStoreEntityTypeFeature{}
	case "gcp:vertex/aiFeatureStoreIamBinding:AiFeatureStoreIamBinding":
		r = &AiFeatureStoreIamBinding{}
	case "gcp:vertex/aiFeatureStoreIamMember:AiFeatureStoreIamMember":
		r = &AiFeatureStoreIamMember{}
	case "gcp:vertex/aiFeatureStoreIamPolicy:AiFeatureStoreIamPolicy":
		r = &AiFeatureStoreIamPolicy{}
	case "gcp:vertex/aiMetadataStore:AiMetadataStore":
		r = &AiMetadataStore{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := gcp.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"gcp",
		"vertex/aiDataset",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"vertex/aiEndpoint",
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
	pulumi.RegisterResourceModule(
		"gcp",
		"vertex/aiFeatureStoreEntityTypeFeature",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"vertex/aiFeatureStoreIamBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"vertex/aiFeatureStoreIamMember",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"vertex/aiFeatureStoreIamPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"vertex/aiMetadataStore",
		&module{version},
	)
}
