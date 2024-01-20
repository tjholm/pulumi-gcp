// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vmwareengine

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
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
	case "gcp:vmwareengine/cluster:Cluster":
		r = &Cluster{}
	case "gcp:vmwareengine/externalAccessRule:ExternalAccessRule":
		r = &ExternalAccessRule{}
	case "gcp:vmwareengine/externalAddress:ExternalAddress":
		r = &ExternalAddress{}
	case "gcp:vmwareengine/network:Network":
		r = &Network{}
	case "gcp:vmwareengine/networkPeering:NetworkPeering":
		r = &NetworkPeering{}
	case "gcp:vmwareengine/networkPolicy:NetworkPolicy":
		r = &NetworkPolicy{}
	case "gcp:vmwareengine/privateCloud:PrivateCloud":
		r = &PrivateCloud{}
	case "gcp:vmwareengine/subnet:Subnet":
		r = &Subnet{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"gcp",
		"vmwareengine/cluster",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"vmwareengine/externalAccessRule",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"vmwareengine/externalAddress",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"vmwareengine/network",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"vmwareengine/networkPeering",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"vmwareengine/networkPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"vmwareengine/privateCloud",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"vmwareengine/subnet",
		&module{version},
	)
}
