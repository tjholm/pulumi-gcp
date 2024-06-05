// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package logging

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
	case "gcp:logging/billingAccountBucketConfig:BillingAccountBucketConfig":
		r = &BillingAccountBucketConfig{}
	case "gcp:logging/billingAccountExclusion:BillingAccountExclusion":
		r = &BillingAccountExclusion{}
	case "gcp:logging/billingAccountSink:BillingAccountSink":
		r = &BillingAccountSink{}
	case "gcp:logging/folderBucketConfig:FolderBucketConfig":
		r = &FolderBucketConfig{}
	case "gcp:logging/folderExclusion:FolderExclusion":
		r = &FolderExclusion{}
	case "gcp:logging/folderSettings:FolderSettings":
		r = &FolderSettings{}
	case "gcp:logging/folderSink:FolderSink":
		r = &FolderSink{}
	case "gcp:logging/linkedDataset:LinkedDataset":
		r = &LinkedDataset{}
	case "gcp:logging/logView:LogView":
		r = &LogView{}
	case "gcp:logging/logViewIamBinding:LogViewIamBinding":
		r = &LogViewIamBinding{}
	case "gcp:logging/logViewIamMember:LogViewIamMember":
		r = &LogViewIamMember{}
	case "gcp:logging/logViewIamPolicy:LogViewIamPolicy":
		r = &LogViewIamPolicy{}
	case "gcp:logging/metric:Metric":
		r = &Metric{}
	case "gcp:logging/organizationBucketConfig:OrganizationBucketConfig":
		r = &OrganizationBucketConfig{}
	case "gcp:logging/organizationExclusion:OrganizationExclusion":
		r = &OrganizationExclusion{}
	case "gcp:logging/organizationSettings:OrganizationSettings":
		r = &OrganizationSettings{}
	case "gcp:logging/organizationSink:OrganizationSink":
		r = &OrganizationSink{}
	case "gcp:logging/projectBucketConfig:ProjectBucketConfig":
		r = &ProjectBucketConfig{}
	case "gcp:logging/projectExclusion:ProjectExclusion":
		r = &ProjectExclusion{}
	case "gcp:logging/projectSink:ProjectSink":
		r = &ProjectSink{}
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
		"logging/billingAccountBucketConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"logging/billingAccountExclusion",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"logging/billingAccountSink",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"logging/folderBucketConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"logging/folderExclusion",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"logging/folderSettings",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"logging/folderSink",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"logging/linkedDataset",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"logging/logView",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"logging/logViewIamBinding",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"logging/logViewIamMember",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"logging/logViewIamPolicy",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"logging/metric",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"logging/organizationBucketConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"logging/organizationExclusion",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"logging/organizationSettings",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"logging/organizationSink",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"logging/projectBucketConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"logging/projectExclusion",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"gcp",
		"logging/projectSink",
		&module{version},
	)
}
