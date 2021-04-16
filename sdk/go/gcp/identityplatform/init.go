// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package identityplatform

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
	case "gcp:identityplatform/defaultSupportedIdpConfig:DefaultSupportedIdpConfig":
		r = &DefaultSupportedIdpConfig{}
	case "gcp:identityplatform/inboundSamlConfig:InboundSamlConfig":
		r = &InboundSamlConfig{}
	case "gcp:identityplatform/oauthIdpConfig:OauthIdpConfig":
		r = &OauthIdpConfig{}
	case "gcp:identityplatform/tenant:Tenant":
		r = &Tenant{}
	case "gcp:identityplatform/tenantDefaultSupportedIdpConfig:TenantDefaultSupportedIdpConfig":
		r = &TenantDefaultSupportedIdpConfig{}
	case "gcp:identityplatform/tenantInboundSamlConfig:TenantInboundSamlConfig":
		r = &TenantInboundSamlConfig{}
	case "gcp:identityplatform/tenantOauthIdpConfig:TenantOauthIdpConfig":
		r = &TenantOauthIdpConfig{}
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
		"identityplatform/defaultSupportedIdpConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"identityplatform/inboundSamlConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"identityplatform/oauthIdpConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"identityplatform/tenant",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"identityplatform/tenantDefaultSupportedIdpConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"identityplatform/tenantInboundSamlConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"identityplatform/tenantOauthIdpConfig",
		&module{version},
	)
}
