// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package diagflow

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
	case "gcp:diagflow/agent:Agent":
		r = &Agent{}
	case "gcp:diagflow/cxAgent:CxAgent":
		r = &CxAgent{}
	case "gcp:diagflow/cxEntityType:CxEntityType":
		r = &CxEntityType{}
	case "gcp:diagflow/cxEnvironment:CxEnvironment":
		r = &CxEnvironment{}
	case "gcp:diagflow/cxFlow:CxFlow":
		r = &CxFlow{}
	case "gcp:diagflow/cxIntent:CxIntent":
		r = &CxIntent{}
	case "gcp:diagflow/cxPage:CxPage":
		r = &CxPage{}
	case "gcp:diagflow/cxVersion:CxVersion":
		r = &CxVersion{}
	case "gcp:diagflow/entityType:EntityType":
		r = &EntityType{}
	case "gcp:diagflow/fulfillment:Fulfillment":
		r = &Fulfillment{}
	case "gcp:diagflow/intent:Intent":
		r = &Intent{}
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
		"diagflow/agent",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"diagflow/cxAgent",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"diagflow/cxEntityType",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"diagflow/cxEnvironment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"diagflow/cxFlow",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"diagflow/cxIntent",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"diagflow/cxPage",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"diagflow/cxVersion",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"diagflow/entityType",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"diagflow/fulfillment",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"diagflow/intent",
		&module{version},
	)
}
