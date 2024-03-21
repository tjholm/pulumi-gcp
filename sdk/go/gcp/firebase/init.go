// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package firebase

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
	case "gcp:firebase/androidApp:AndroidApp":
		r = &AndroidApp{}
	case "gcp:firebase/appCheckAppAttestConfig:AppCheckAppAttestConfig":
		r = &AppCheckAppAttestConfig{}
	case "gcp:firebase/appCheckDebugToken:AppCheckDebugToken":
		r = &AppCheckDebugToken{}
	case "gcp:firebase/appCheckDeviceCheckConfig:AppCheckDeviceCheckConfig":
		r = &AppCheckDeviceCheckConfig{}
	case "gcp:firebase/appCheckPlayIntegrityConfig:AppCheckPlayIntegrityConfig":
		r = &AppCheckPlayIntegrityConfig{}
	case "gcp:firebase/appCheckRecaptchaEnterpriseConfig:AppCheckRecaptchaEnterpriseConfig":
		r = &AppCheckRecaptchaEnterpriseConfig{}
	case "gcp:firebase/appCheckRecaptchaV3Config:AppCheckRecaptchaV3Config":
		r = &AppCheckRecaptchaV3Config{}
	case "gcp:firebase/appCheckServiceConfig:AppCheckServiceConfig":
		r = &AppCheckServiceConfig{}
	case "gcp:firebase/appleApp:AppleApp":
		r = &AppleApp{}
	case "gcp:firebase/databaseInstance:DatabaseInstance":
		r = &DatabaseInstance{}
	case "gcp:firebase/extensionsInstance:ExtensionsInstance":
		r = &ExtensionsInstance{}
	case "gcp:firebase/hostingChannel:HostingChannel":
		r = &HostingChannel{}
	case "gcp:firebase/hostingCustomDomain:HostingCustomDomain":
		r = &HostingCustomDomain{}
	case "gcp:firebase/hostingRelease:HostingRelease":
		r = &HostingRelease{}
	case "gcp:firebase/hostingSite:HostingSite":
		r = &HostingSite{}
	case "gcp:firebase/hostingVersion:HostingVersion":
		r = &HostingVersion{}
	case "gcp:firebase/project:Project":
		r = &Project{}
	case "gcp:firebase/storageBucket:StorageBucket":
		r = &StorageBucket{}
	case "gcp:firebase/webApp:WebApp":
		r = &WebApp{}
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
		"firebase/androidApp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"firebase/appCheckAppAttestConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"firebase/appCheckDebugToken",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"firebase/appCheckDeviceCheckConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"firebase/appCheckPlayIntegrityConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"firebase/appCheckRecaptchaEnterpriseConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"firebase/appCheckRecaptchaV3Config",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"firebase/appCheckServiceConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"firebase/appleApp",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"firebase/databaseInstance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"firebase/extensionsInstance",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"firebase/hostingChannel",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"firebase/hostingCustomDomain",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"firebase/hostingRelease",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"firebase/hostingSite",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"firebase/hostingVersion",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"firebase/project",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"firebase/storageBucket",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"firebase/webApp",
		&module{version},
	)
}
