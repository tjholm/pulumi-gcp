// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package container

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Anthos cluster running on customer owned infrastructure.
//
// To get more information about Cluster, see:
//
// * [API documentation](https://cloud.google.com/anthos/clusters/docs/multi-cloud/reference/rest)
// * How-to Guides
//   - [API reference](https://cloud.google.com/anthos/clusters/docs/multi-cloud/reference/rest/v1/projects.locations.attachedClusters)
//   - [Multicloud overview](https://cloud.google.com/anthos/clusters/docs/multi-cloud)
//
// ## Example Usage
// ### Container Attached Cluster Basic
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/container"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			project, err := organizations.LookupProject(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			versions, err := container.GetAttachedVersions(ctx, &container.GetAttachedVersionsArgs{
//				Location: "us-west1",
//				Project:  project.ProjectId,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = container.NewAttachedCluster(ctx, "primary", &container.AttachedClusterArgs{
//				Location:     pulumi.String("us-west1"),
//				Project:      *pulumi.String(project.ProjectId),
//				Description:  pulumi.String("Test cluster"),
//				Distribution: pulumi.String("aks"),
//				OidcConfig: &container.AttachedClusterOidcConfigArgs{
//					IssuerUrl: pulumi.String("https://oidc.issuer.url"),
//				},
//				PlatformVersion: *pulumi.String(versions.ValidVersions[0]),
//				Fleet: &container.AttachedClusterFleetArgs{
//					Project: pulumi.String(fmt.Sprintf("projects/%v", project.Number)),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Container Attached Cluster Ignore Errors
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/container"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			project, err := organizations.LookupProject(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			versions, err := container.GetAttachedVersions(ctx, &container.GetAttachedVersionsArgs{
//				Location: "us-west1",
//				Project:  project.ProjectId,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = container.NewAttachedCluster(ctx, "primary", &container.AttachedClusterArgs{
//				Location:     pulumi.String("us-west1"),
//				Project:      *pulumi.String(project.ProjectId),
//				Description:  pulumi.String("Test cluster"),
//				Distribution: pulumi.String("aks"),
//				OidcConfig: &container.AttachedClusterOidcConfigArgs{
//					IssuerUrl: pulumi.String("https://oidc.issuer.url"),
//				},
//				PlatformVersion: *pulumi.String(versions.ValidVersions[0]),
//				Fleet: &container.AttachedClusterFleetArgs{
//					Project: pulumi.String(fmt.Sprintf("projects/%v", project.Number)),
//				},
//				DeletionPolicy: pulumi.String("DELETE_IGNORE_ERRORS"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// # Cluster can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:container/attachedCluster:AttachedCluster default projects/{{project}}/locations/{{location}}/attachedClusters/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:container/attachedCluster:AttachedCluster default {{project}}/{{location}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:container/attachedCluster:AttachedCluster default {{location}}/{{name}}
//
// ```
type AttachedCluster struct {
	pulumi.CustomResourceState

	// Optional. Annotations on the cluster. This field has the same
	// restrictions as Kubernetes annotations. The total size of all keys and
	// values combined is limited to 256k. Key can have 2 segments: prefix (optional)
	// and name (required), separated by a slash (/). Prefix must be a DNS subdomain.
	// Name must be 63 characters or less, begin and end with alphanumerics,
	// with dashes (-), underscores (_), dots (.), and alphanumerics between.
	Annotations pulumi.StringMapOutput `pulumi:"annotations"`
	// Configuration related to the cluster RBAC settings.
	// Structure is documented below.
	Authorization AttachedClusterAuthorizationPtrOutput `pulumi:"authorization"`
	// Output only. The region where this cluster runs.
	// For EKS clusters, this is an AWS region. For AKS clusters,
	// this is an Azure region.
	ClusterRegion pulumi.StringOutput `pulumi:"clusterRegion"`
	// Output only. The time at which this cluster was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Policy to determine what flags to send on delete.
	DeletionPolicy pulumi.StringPtrOutput `pulumi:"deletionPolicy"`
	// A human readable description of this attached cluster. Cannot be longer
	// than 255 UTF-8 encoded bytes.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The Kubernetes distribution of the underlying attached cluster. Supported values:
	// "eks", "aks".
	Distribution pulumi.StringOutput `pulumi:"distribution"`
	// A set of errors found in the cluster.
	// Structure is documented below.
	Errors AttachedClusterErrorArrayOutput `pulumi:"errors"`
	// Fleet configuration.
	// Structure is documented below.
	Fleet AttachedClusterFleetOutput `pulumi:"fleet"`
	// The Kubernetes version of the cluster.
	KubernetesVersion pulumi.StringOutput `pulumi:"kubernetesVersion"`
	// The location for the resource
	Location pulumi.StringOutput `pulumi:"location"`
	// Logging configuration.
	// Structure is documented below.
	LoggingConfig AttachedClusterLoggingConfigPtrOutput `pulumi:"loggingConfig"`
	// Monitoring configuration.
	// Structure is documented below.
	MonitoringConfig AttachedClusterMonitoringConfigOutput `pulumi:"monitoringConfig"`
	// The name of this resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// OIDC discovery information of the target cluster.
	// Kubernetes Service Account (KSA) tokens are JWT tokens signed by the cluster
	// API server. This fields indicates how GCP services
	// validate KSA tokens in order to allow system workloads (such as GKE Connect
	// and telemetry agents) to authenticate back to GCP.
	// Both clusters with public and private issuer URLs are supported.
	// Clusters with public issuers only need to specify the `issuerUrl` field
	// while clusters with private issuers need to provide both
	// `issuerUrl` and `jwks`.
	// Structure is documented below.
	OidcConfig AttachedClusterOidcConfigOutput `pulumi:"oidcConfig"`
	// The platform version for the cluster (e.g. `1.23.0-gke.1`).
	PlatformVersion pulumi.StringOutput `pulumi:"platformVersion"`
	// The number of the Fleet host project where this cluster will be registered.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// If set, there are currently changes in flight to the cluster.
	Reconciling pulumi.BoolOutput `pulumi:"reconciling"`
	// The current state of the cluster. Possible values:
	// STATE_UNSPECIFIED, PROVISIONING, RUNNING, RECONCILING, STOPPING, ERROR,
	// DEGRADED
	State pulumi.StringOutput `pulumi:"state"`
	// A globally unique identifier for the cluster.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// The time at which this cluster was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// Workload Identity settings.
	// Structure is documented below.
	WorkloadIdentityConfigs AttachedClusterWorkloadIdentityConfigArrayOutput `pulumi:"workloadIdentityConfigs"`
}

// NewAttachedCluster registers a new resource with the given unique name, arguments, and options.
func NewAttachedCluster(ctx *pulumi.Context,
	name string, args *AttachedClusterArgs, opts ...pulumi.ResourceOption) (*AttachedCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Distribution == nil {
		return nil, errors.New("invalid value for required argument 'Distribution'")
	}
	if args.Fleet == nil {
		return nil, errors.New("invalid value for required argument 'Fleet'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.OidcConfig == nil {
		return nil, errors.New("invalid value for required argument 'OidcConfig'")
	}
	if args.PlatformVersion == nil {
		return nil, errors.New("invalid value for required argument 'PlatformVersion'")
	}
	var resource AttachedCluster
	err := ctx.RegisterResource("gcp:container/attachedCluster:AttachedCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAttachedCluster gets an existing AttachedCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAttachedCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AttachedClusterState, opts ...pulumi.ResourceOption) (*AttachedCluster, error) {
	var resource AttachedCluster
	err := ctx.ReadResource("gcp:container/attachedCluster:AttachedCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AttachedCluster resources.
type attachedClusterState struct {
	// Optional. Annotations on the cluster. This field has the same
	// restrictions as Kubernetes annotations. The total size of all keys and
	// values combined is limited to 256k. Key can have 2 segments: prefix (optional)
	// and name (required), separated by a slash (/). Prefix must be a DNS subdomain.
	// Name must be 63 characters or less, begin and end with alphanumerics,
	// with dashes (-), underscores (_), dots (.), and alphanumerics between.
	Annotations map[string]string `pulumi:"annotations"`
	// Configuration related to the cluster RBAC settings.
	// Structure is documented below.
	Authorization *AttachedClusterAuthorization `pulumi:"authorization"`
	// Output only. The region where this cluster runs.
	// For EKS clusters, this is an AWS region. For AKS clusters,
	// this is an Azure region.
	ClusterRegion *string `pulumi:"clusterRegion"`
	// Output only. The time at which this cluster was created.
	CreateTime *string `pulumi:"createTime"`
	// Policy to determine what flags to send on delete.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// A human readable description of this attached cluster. Cannot be longer
	// than 255 UTF-8 encoded bytes.
	Description *string `pulumi:"description"`
	// The Kubernetes distribution of the underlying attached cluster. Supported values:
	// "eks", "aks".
	Distribution *string `pulumi:"distribution"`
	// A set of errors found in the cluster.
	// Structure is documented below.
	Errors []AttachedClusterError `pulumi:"errors"`
	// Fleet configuration.
	// Structure is documented below.
	Fleet *AttachedClusterFleet `pulumi:"fleet"`
	// The Kubernetes version of the cluster.
	KubernetesVersion *string `pulumi:"kubernetesVersion"`
	// The location for the resource
	Location *string `pulumi:"location"`
	// Logging configuration.
	// Structure is documented below.
	LoggingConfig *AttachedClusterLoggingConfig `pulumi:"loggingConfig"`
	// Monitoring configuration.
	// Structure is documented below.
	MonitoringConfig *AttachedClusterMonitoringConfig `pulumi:"monitoringConfig"`
	// The name of this resource.
	Name *string `pulumi:"name"`
	// OIDC discovery information of the target cluster.
	// Kubernetes Service Account (KSA) tokens are JWT tokens signed by the cluster
	// API server. This fields indicates how GCP services
	// validate KSA tokens in order to allow system workloads (such as GKE Connect
	// and telemetry agents) to authenticate back to GCP.
	// Both clusters with public and private issuer URLs are supported.
	// Clusters with public issuers only need to specify the `issuerUrl` field
	// while clusters with private issuers need to provide both
	// `issuerUrl` and `jwks`.
	// Structure is documented below.
	OidcConfig *AttachedClusterOidcConfig `pulumi:"oidcConfig"`
	// The platform version for the cluster (e.g. `1.23.0-gke.1`).
	PlatformVersion *string `pulumi:"platformVersion"`
	// The number of the Fleet host project where this cluster will be registered.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// If set, there are currently changes in flight to the cluster.
	Reconciling *bool `pulumi:"reconciling"`
	// The current state of the cluster. Possible values:
	// STATE_UNSPECIFIED, PROVISIONING, RUNNING, RECONCILING, STOPPING, ERROR,
	// DEGRADED
	State *string `pulumi:"state"`
	// A globally unique identifier for the cluster.
	Uid *string `pulumi:"uid"`
	// The time at which this cluster was last updated.
	UpdateTime *string `pulumi:"updateTime"`
	// Workload Identity settings.
	// Structure is documented below.
	WorkloadIdentityConfigs []AttachedClusterWorkloadIdentityConfig `pulumi:"workloadIdentityConfigs"`
}

type AttachedClusterState struct {
	// Optional. Annotations on the cluster. This field has the same
	// restrictions as Kubernetes annotations. The total size of all keys and
	// values combined is limited to 256k. Key can have 2 segments: prefix (optional)
	// and name (required), separated by a slash (/). Prefix must be a DNS subdomain.
	// Name must be 63 characters or less, begin and end with alphanumerics,
	// with dashes (-), underscores (_), dots (.), and alphanumerics between.
	Annotations pulumi.StringMapInput
	// Configuration related to the cluster RBAC settings.
	// Structure is documented below.
	Authorization AttachedClusterAuthorizationPtrInput
	// Output only. The region where this cluster runs.
	// For EKS clusters, this is an AWS region. For AKS clusters,
	// this is an Azure region.
	ClusterRegion pulumi.StringPtrInput
	// Output only. The time at which this cluster was created.
	CreateTime pulumi.StringPtrInput
	// Policy to determine what flags to send on delete.
	DeletionPolicy pulumi.StringPtrInput
	// A human readable description of this attached cluster. Cannot be longer
	// than 255 UTF-8 encoded bytes.
	Description pulumi.StringPtrInput
	// The Kubernetes distribution of the underlying attached cluster. Supported values:
	// "eks", "aks".
	Distribution pulumi.StringPtrInput
	// A set of errors found in the cluster.
	// Structure is documented below.
	Errors AttachedClusterErrorArrayInput
	// Fleet configuration.
	// Structure is documented below.
	Fleet AttachedClusterFleetPtrInput
	// The Kubernetes version of the cluster.
	KubernetesVersion pulumi.StringPtrInput
	// The location for the resource
	Location pulumi.StringPtrInput
	// Logging configuration.
	// Structure is documented below.
	LoggingConfig AttachedClusterLoggingConfigPtrInput
	// Monitoring configuration.
	// Structure is documented below.
	MonitoringConfig AttachedClusterMonitoringConfigPtrInput
	// The name of this resource.
	Name pulumi.StringPtrInput
	// OIDC discovery information of the target cluster.
	// Kubernetes Service Account (KSA) tokens are JWT tokens signed by the cluster
	// API server. This fields indicates how GCP services
	// validate KSA tokens in order to allow system workloads (such as GKE Connect
	// and telemetry agents) to authenticate back to GCP.
	// Both clusters with public and private issuer URLs are supported.
	// Clusters with public issuers only need to specify the `issuerUrl` field
	// while clusters with private issuers need to provide both
	// `issuerUrl` and `jwks`.
	// Structure is documented below.
	OidcConfig AttachedClusterOidcConfigPtrInput
	// The platform version for the cluster (e.g. `1.23.0-gke.1`).
	PlatformVersion pulumi.StringPtrInput
	// The number of the Fleet host project where this cluster will be registered.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// If set, there are currently changes in flight to the cluster.
	Reconciling pulumi.BoolPtrInput
	// The current state of the cluster. Possible values:
	// STATE_UNSPECIFIED, PROVISIONING, RUNNING, RECONCILING, STOPPING, ERROR,
	// DEGRADED
	State pulumi.StringPtrInput
	// A globally unique identifier for the cluster.
	Uid pulumi.StringPtrInput
	// The time at which this cluster was last updated.
	UpdateTime pulumi.StringPtrInput
	// Workload Identity settings.
	// Structure is documented below.
	WorkloadIdentityConfigs AttachedClusterWorkloadIdentityConfigArrayInput
}

func (AttachedClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*attachedClusterState)(nil)).Elem()
}

type attachedClusterArgs struct {
	// Optional. Annotations on the cluster. This field has the same
	// restrictions as Kubernetes annotations. The total size of all keys and
	// values combined is limited to 256k. Key can have 2 segments: prefix (optional)
	// and name (required), separated by a slash (/). Prefix must be a DNS subdomain.
	// Name must be 63 characters or less, begin and end with alphanumerics,
	// with dashes (-), underscores (_), dots (.), and alphanumerics between.
	Annotations map[string]string `pulumi:"annotations"`
	// Configuration related to the cluster RBAC settings.
	// Structure is documented below.
	Authorization *AttachedClusterAuthorization `pulumi:"authorization"`
	// Policy to determine what flags to send on delete.
	DeletionPolicy *string `pulumi:"deletionPolicy"`
	// A human readable description of this attached cluster. Cannot be longer
	// than 255 UTF-8 encoded bytes.
	Description *string `pulumi:"description"`
	// The Kubernetes distribution of the underlying attached cluster. Supported values:
	// "eks", "aks".
	Distribution string `pulumi:"distribution"`
	// Fleet configuration.
	// Structure is documented below.
	Fleet AttachedClusterFleet `pulumi:"fleet"`
	// The location for the resource
	Location string `pulumi:"location"`
	// Logging configuration.
	// Structure is documented below.
	LoggingConfig *AttachedClusterLoggingConfig `pulumi:"loggingConfig"`
	// Monitoring configuration.
	// Structure is documented below.
	MonitoringConfig *AttachedClusterMonitoringConfig `pulumi:"monitoringConfig"`
	// The name of this resource.
	Name *string `pulumi:"name"`
	// OIDC discovery information of the target cluster.
	// Kubernetes Service Account (KSA) tokens are JWT tokens signed by the cluster
	// API server. This fields indicates how GCP services
	// validate KSA tokens in order to allow system workloads (such as GKE Connect
	// and telemetry agents) to authenticate back to GCP.
	// Both clusters with public and private issuer URLs are supported.
	// Clusters with public issuers only need to specify the `issuerUrl` field
	// while clusters with private issuers need to provide both
	// `issuerUrl` and `jwks`.
	// Structure is documented below.
	OidcConfig AttachedClusterOidcConfig `pulumi:"oidcConfig"`
	// The platform version for the cluster (e.g. `1.23.0-gke.1`).
	PlatformVersion string `pulumi:"platformVersion"`
	// The number of the Fleet host project where this cluster will be registered.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a AttachedCluster resource.
type AttachedClusterArgs struct {
	// Optional. Annotations on the cluster. This field has the same
	// restrictions as Kubernetes annotations. The total size of all keys and
	// values combined is limited to 256k. Key can have 2 segments: prefix (optional)
	// and name (required), separated by a slash (/). Prefix must be a DNS subdomain.
	// Name must be 63 characters or less, begin and end with alphanumerics,
	// with dashes (-), underscores (_), dots (.), and alphanumerics between.
	Annotations pulumi.StringMapInput
	// Configuration related to the cluster RBAC settings.
	// Structure is documented below.
	Authorization AttachedClusterAuthorizationPtrInput
	// Policy to determine what flags to send on delete.
	DeletionPolicy pulumi.StringPtrInput
	// A human readable description of this attached cluster. Cannot be longer
	// than 255 UTF-8 encoded bytes.
	Description pulumi.StringPtrInput
	// The Kubernetes distribution of the underlying attached cluster. Supported values:
	// "eks", "aks".
	Distribution pulumi.StringInput
	// Fleet configuration.
	// Structure is documented below.
	Fleet AttachedClusterFleetInput
	// The location for the resource
	Location pulumi.StringInput
	// Logging configuration.
	// Structure is documented below.
	LoggingConfig AttachedClusterLoggingConfigPtrInput
	// Monitoring configuration.
	// Structure is documented below.
	MonitoringConfig AttachedClusterMonitoringConfigPtrInput
	// The name of this resource.
	Name pulumi.StringPtrInput
	// OIDC discovery information of the target cluster.
	// Kubernetes Service Account (KSA) tokens are JWT tokens signed by the cluster
	// API server. This fields indicates how GCP services
	// validate KSA tokens in order to allow system workloads (such as GKE Connect
	// and telemetry agents) to authenticate back to GCP.
	// Both clusters with public and private issuer URLs are supported.
	// Clusters with public issuers only need to specify the `issuerUrl` field
	// while clusters with private issuers need to provide both
	// `issuerUrl` and `jwks`.
	// Structure is documented below.
	OidcConfig AttachedClusterOidcConfigInput
	// The platform version for the cluster (e.g. `1.23.0-gke.1`).
	PlatformVersion pulumi.StringInput
	// The number of the Fleet host project where this cluster will be registered.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (AttachedClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*attachedClusterArgs)(nil)).Elem()
}

type AttachedClusterInput interface {
	pulumi.Input

	ToAttachedClusterOutput() AttachedClusterOutput
	ToAttachedClusterOutputWithContext(ctx context.Context) AttachedClusterOutput
}

func (*AttachedCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**AttachedCluster)(nil)).Elem()
}

func (i *AttachedCluster) ToAttachedClusterOutput() AttachedClusterOutput {
	return i.ToAttachedClusterOutputWithContext(context.Background())
}

func (i *AttachedCluster) ToAttachedClusterOutputWithContext(ctx context.Context) AttachedClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttachedClusterOutput)
}

// AttachedClusterArrayInput is an input type that accepts AttachedClusterArray and AttachedClusterArrayOutput values.
// You can construct a concrete instance of `AttachedClusterArrayInput` via:
//
//	AttachedClusterArray{ AttachedClusterArgs{...} }
type AttachedClusterArrayInput interface {
	pulumi.Input

	ToAttachedClusterArrayOutput() AttachedClusterArrayOutput
	ToAttachedClusterArrayOutputWithContext(context.Context) AttachedClusterArrayOutput
}

type AttachedClusterArray []AttachedClusterInput

func (AttachedClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AttachedCluster)(nil)).Elem()
}

func (i AttachedClusterArray) ToAttachedClusterArrayOutput() AttachedClusterArrayOutput {
	return i.ToAttachedClusterArrayOutputWithContext(context.Background())
}

func (i AttachedClusterArray) ToAttachedClusterArrayOutputWithContext(ctx context.Context) AttachedClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttachedClusterArrayOutput)
}

// AttachedClusterMapInput is an input type that accepts AttachedClusterMap and AttachedClusterMapOutput values.
// You can construct a concrete instance of `AttachedClusterMapInput` via:
//
//	AttachedClusterMap{ "key": AttachedClusterArgs{...} }
type AttachedClusterMapInput interface {
	pulumi.Input

	ToAttachedClusterMapOutput() AttachedClusterMapOutput
	ToAttachedClusterMapOutputWithContext(context.Context) AttachedClusterMapOutput
}

type AttachedClusterMap map[string]AttachedClusterInput

func (AttachedClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AttachedCluster)(nil)).Elem()
}

func (i AttachedClusterMap) ToAttachedClusterMapOutput() AttachedClusterMapOutput {
	return i.ToAttachedClusterMapOutputWithContext(context.Background())
}

func (i AttachedClusterMap) ToAttachedClusterMapOutputWithContext(ctx context.Context) AttachedClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AttachedClusterMapOutput)
}

type AttachedClusterOutput struct{ *pulumi.OutputState }

func (AttachedClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AttachedCluster)(nil)).Elem()
}

func (o AttachedClusterOutput) ToAttachedClusterOutput() AttachedClusterOutput {
	return o
}

func (o AttachedClusterOutput) ToAttachedClusterOutputWithContext(ctx context.Context) AttachedClusterOutput {
	return o
}

// Optional. Annotations on the cluster. This field has the same
// restrictions as Kubernetes annotations. The total size of all keys and
// values combined is limited to 256k. Key can have 2 segments: prefix (optional)
// and name (required), separated by a slash (/). Prefix must be a DNS subdomain.
// Name must be 63 characters or less, begin and end with alphanumerics,
// with dashes (-), underscores (_), dots (.), and alphanumerics between.
func (o AttachedClusterOutput) Annotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AttachedCluster) pulumi.StringMapOutput { return v.Annotations }).(pulumi.StringMapOutput)
}

// Configuration related to the cluster RBAC settings.
// Structure is documented below.
func (o AttachedClusterOutput) Authorization() AttachedClusterAuthorizationPtrOutput {
	return o.ApplyT(func(v *AttachedCluster) AttachedClusterAuthorizationPtrOutput { return v.Authorization }).(AttachedClusterAuthorizationPtrOutput)
}

// Output only. The region where this cluster runs.
// For EKS clusters, this is an AWS region. For AKS clusters,
// this is an Azure region.
func (o AttachedClusterOutput) ClusterRegion() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedCluster) pulumi.StringOutput { return v.ClusterRegion }).(pulumi.StringOutput)
}

// Output only. The time at which this cluster was created.
func (o AttachedClusterOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedCluster) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Policy to determine what flags to send on delete.
func (o AttachedClusterOutput) DeletionPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttachedCluster) pulumi.StringPtrOutput { return v.DeletionPolicy }).(pulumi.StringPtrOutput)
}

// A human readable description of this attached cluster. Cannot be longer
// than 255 UTF-8 encoded bytes.
func (o AttachedClusterOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AttachedCluster) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The Kubernetes distribution of the underlying attached cluster. Supported values:
// "eks", "aks".
func (o AttachedClusterOutput) Distribution() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedCluster) pulumi.StringOutput { return v.Distribution }).(pulumi.StringOutput)
}

// A set of errors found in the cluster.
// Structure is documented below.
func (o AttachedClusterOutput) Errors() AttachedClusterErrorArrayOutput {
	return o.ApplyT(func(v *AttachedCluster) AttachedClusterErrorArrayOutput { return v.Errors }).(AttachedClusterErrorArrayOutput)
}

// Fleet configuration.
// Structure is documented below.
func (o AttachedClusterOutput) Fleet() AttachedClusterFleetOutput {
	return o.ApplyT(func(v *AttachedCluster) AttachedClusterFleetOutput { return v.Fleet }).(AttachedClusterFleetOutput)
}

// The Kubernetes version of the cluster.
func (o AttachedClusterOutput) KubernetesVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedCluster) pulumi.StringOutput { return v.KubernetesVersion }).(pulumi.StringOutput)
}

// The location for the resource
func (o AttachedClusterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedCluster) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Logging configuration.
// Structure is documented below.
func (o AttachedClusterOutput) LoggingConfig() AttachedClusterLoggingConfigPtrOutput {
	return o.ApplyT(func(v *AttachedCluster) AttachedClusterLoggingConfigPtrOutput { return v.LoggingConfig }).(AttachedClusterLoggingConfigPtrOutput)
}

// Monitoring configuration.
// Structure is documented below.
func (o AttachedClusterOutput) MonitoringConfig() AttachedClusterMonitoringConfigOutput {
	return o.ApplyT(func(v *AttachedCluster) AttachedClusterMonitoringConfigOutput { return v.MonitoringConfig }).(AttachedClusterMonitoringConfigOutput)
}

// The name of this resource.
func (o AttachedClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// OIDC discovery information of the target cluster.
// Kubernetes Service Account (KSA) tokens are JWT tokens signed by the cluster
// API server. This fields indicates how GCP services
// validate KSA tokens in order to allow system workloads (such as GKE Connect
// and telemetry agents) to authenticate back to GCP.
// Both clusters with public and private issuer URLs are supported.
// Clusters with public issuers only need to specify the `issuerUrl` field
// while clusters with private issuers need to provide both
// `issuerUrl` and `jwks`.
// Structure is documented below.
func (o AttachedClusterOutput) OidcConfig() AttachedClusterOidcConfigOutput {
	return o.ApplyT(func(v *AttachedCluster) AttachedClusterOidcConfigOutput { return v.OidcConfig }).(AttachedClusterOidcConfigOutput)
}

// The platform version for the cluster (e.g. `1.23.0-gke.1`).
func (o AttachedClusterOutput) PlatformVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedCluster) pulumi.StringOutput { return v.PlatformVersion }).(pulumi.StringOutput)
}

// The number of the Fleet host project where this cluster will be registered.
// If it is not provided, the provider project is used.
func (o AttachedClusterOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedCluster) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// If set, there are currently changes in flight to the cluster.
func (o AttachedClusterOutput) Reconciling() pulumi.BoolOutput {
	return o.ApplyT(func(v *AttachedCluster) pulumi.BoolOutput { return v.Reconciling }).(pulumi.BoolOutput)
}

// The current state of the cluster. Possible values:
// STATE_UNSPECIFIED, PROVISIONING, RUNNING, RECONCILING, STOPPING, ERROR,
// DEGRADED
func (o AttachedClusterOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedCluster) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// A globally unique identifier for the cluster.
func (o AttachedClusterOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedCluster) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// The time at which this cluster was last updated.
func (o AttachedClusterOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AttachedCluster) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// Workload Identity settings.
// Structure is documented below.
func (o AttachedClusterOutput) WorkloadIdentityConfigs() AttachedClusterWorkloadIdentityConfigArrayOutput {
	return o.ApplyT(func(v *AttachedCluster) AttachedClusterWorkloadIdentityConfigArrayOutput {
		return v.WorkloadIdentityConfigs
	}).(AttachedClusterWorkloadIdentityConfigArrayOutput)
}

type AttachedClusterArrayOutput struct{ *pulumi.OutputState }

func (AttachedClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AttachedCluster)(nil)).Elem()
}

func (o AttachedClusterArrayOutput) ToAttachedClusterArrayOutput() AttachedClusterArrayOutput {
	return o
}

func (o AttachedClusterArrayOutput) ToAttachedClusterArrayOutputWithContext(ctx context.Context) AttachedClusterArrayOutput {
	return o
}

func (o AttachedClusterArrayOutput) Index(i pulumi.IntInput) AttachedClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AttachedCluster {
		return vs[0].([]*AttachedCluster)[vs[1].(int)]
	}).(AttachedClusterOutput)
}

type AttachedClusterMapOutput struct{ *pulumi.OutputState }

func (AttachedClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AttachedCluster)(nil)).Elem()
}

func (o AttachedClusterMapOutput) ToAttachedClusterMapOutput() AttachedClusterMapOutput {
	return o
}

func (o AttachedClusterMapOutput) ToAttachedClusterMapOutputWithContext(ctx context.Context) AttachedClusterMapOutput {
	return o
}

func (o AttachedClusterMapOutput) MapIndex(k pulumi.StringInput) AttachedClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AttachedCluster {
		return vs[0].(map[string]*AttachedCluster)[vs[1].(string)]
	}).(AttachedClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AttachedClusterInput)(nil)).Elem(), &AttachedCluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*AttachedClusterArrayInput)(nil)).Elem(), AttachedClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AttachedClusterMapInput)(nil)).Elem(), AttachedClusterMap{})
	pulumi.RegisterOutputType(AttachedClusterOutput{})
	pulumi.RegisterOutputType(AttachedClusterArrayOutput{})
	pulumi.RegisterOutputType(AttachedClusterMapOutput{})
}
