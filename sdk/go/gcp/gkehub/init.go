// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gkehub

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp"
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
	case "gcp:gkehub/feature:Feature":
		r = &Feature{}
	case "gcp:gkehub/featureMembership:FeatureMembership":
		r = &FeatureMembership{}
	case "gcp:gkehub/membership:Membership":
		r = &Membership{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := gcp.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"gcp",
		"gkehub/feature",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"gkehub/featureMembership",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"gkehub/membership",
		&module{version},
	)
}
