// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package container

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An Anthos cluster running on AWS.
//
// For more information, see:
// * [Multicloud overview](https://cloud.google.com/anthos/clusters/docs/multi-cloud)
// ## Example Usage
// ### Basic_aws_cluster
// A basic example of a containeraws cluster
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/container"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "us-west1"
// 		opt1 := "my-project-name"
// 		versions, err := container.GetAwsVersions(ctx, &container.GetAwsVersionsArgs{
// 			Location: &opt0,
// 			Project:  &opt1,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = container.NewAwsCluster(ctx, "primary", &container.AwsClusterArgs{
// 			Annotations: pulumi.StringMap{
// 				"label-one": pulumi.String("value-one"),
// 			},
// 			Authorization: &container.AwsClusterAuthorizationArgs{
// 				AdminUsers: container.AwsClusterAuthorizationAdminUserArray{
// 					&container.AwsClusterAuthorizationAdminUserArgs{
// 						Username: pulumi.String("emailAddress:my@service-account.com"),
// 					},
// 				},
// 			},
// 			AwsRegion: pulumi.String("my-aws-region"),
// 			ControlPlane: &container.AwsClusterControlPlaneArgs{
// 				AwsServicesAuthentication: &container.AwsClusterControlPlaneAwsServicesAuthenticationArgs{
// 					RoleArn:         pulumi.String("arn:aws:iam::012345678910:role/my--1p-dev-oneplatform"),
// 					RoleSessionName: pulumi.String("my--1p-dev-session"),
// 				},
// 				ConfigEncryption: &container.AwsClusterControlPlaneConfigEncryptionArgs{
// 					KmsKeyArn: pulumi.String("arn:aws:kms:my-aws-region:012345678910:key/12345678-1234-1234-1234-123456789111"),
// 				},
// 				DatabaseEncryption: &container.AwsClusterControlPlaneDatabaseEncryptionArgs{
// 					KmsKeyArn: pulumi.String("arn:aws:kms:my-aws-region:012345678910:key/12345678-1234-1234-1234-123456789111"),
// 				},
// 				IamInstanceProfile: pulumi.String("my--1p-dev-controlplane"),
// 				InstanceType:       pulumi.String("t3.medium"),
// 				MainVolume: &container.AwsClusterControlPlaneMainVolumeArgs{
// 					Iops:       pulumi.Int(3000),
// 					KmsKeyArn:  pulumi.String("arn:aws:kms:my-aws-region:012345678910:key/12345678-1234-1234-1234-123456789111"),
// 					SizeGib:    pulumi.Int(10),
// 					VolumeType: pulumi.String("GP3"),
// 				},
// 				ProxyConfig: &container.AwsClusterControlPlaneProxyConfigArgs{
// 					SecretArn:     pulumi.String("arn:aws:secretsmanager:us-west-2:126285863215:secret:proxy_config20210824150329476300000001-ABCDEF"),
// 					SecretVersion: pulumi.String("12345678-ABCD-EFGH-IJKL-987654321098"),
// 				},
// 				RootVolume: &container.AwsClusterControlPlaneRootVolumeArgs{
// 					Iops:       pulumi.Int(3000),
// 					KmsKeyArn:  pulumi.String("arn:aws:kms:my-aws-region:012345678910:key/12345678-1234-1234-1234-123456789111"),
// 					SizeGib:    pulumi.Int(10),
// 					VolumeType: pulumi.String("GP3"),
// 				},
// 				SecurityGroupIds: pulumi.StringArray{
// 					pulumi.String("sg-00000000000000000"),
// 				},
// 				SshConfig: &container.AwsClusterControlPlaneSshConfigArgs{
// 					Ec2KeyPair: pulumi.String("my--1p-dev-ssh"),
// 				},
// 				SubnetIds: pulumi.StringArray{
// 					pulumi.String("subnet-00000000000000000"),
// 				},
// 				Tags: pulumi.StringMap{
// 					"owner": pulumi.String("emailAddress:my@service-account.com"),
// 				},
// 				Version: pulumi.String(versions.ValidVersions[0]),
// 			},
// 			Description: pulumi.String("A sample aws cluster"),
// 			Fleet: &container.AwsClusterFleetArgs{
// 				Project: pulumi.String("my-project-number"),
// 			},
// 			Location: pulumi.String("us-west1"),
// 			Networking: &container.AwsClusterNetworkingArgs{
// 				PodAddressCidrBlocks: pulumi.StringArray{
// 					pulumi.String("10.2.0.0/16"),
// 				},
// 				ServiceAddressCidrBlocks: pulumi.StringArray{
// 					pulumi.String("10.1.0.0/16"),
// 				},
// 				VpcId: pulumi.String("vpc-00000000000000000"),
// 			},
// 			Project: pulumi.String("my-project-name"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Cluster can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:container/awsCluster:AwsCluster default projects/{{project}}/locations/{{location}}/awsClusters/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:container/awsCluster:AwsCluster default {{project}}/{{location}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:container/awsCluster:AwsCluster default {{location}}/{{name}}
// ```
type AwsCluster struct {
	pulumi.CustomResourceState

	// Optional. Annotations on the cluster. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Key can have 2 segments: prefix (optional) and name (required), separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.
	Annotations pulumi.StringMapOutput `pulumi:"annotations"`
	// Required. Configuration related to the cluster RBAC settings.
	Authorization AwsClusterAuthorizationOutput `pulumi:"authorization"`
	// Required. The AWS region where the cluster runs. Each Google Cloud region supports a subset of nearby AWS regions. You can call to list all supported AWS regions within a given Google Cloud region.
	AwsRegion pulumi.StringOutput `pulumi:"awsRegion"`
	// Required. Configuration related to the cluster control plane.
	ControlPlane AwsClusterControlPlaneOutput `pulumi:"controlPlane"`
	// Output only. The time at which this cluster was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Optional. A human readable description of this cluster. Cannot be longer than 255 UTF-8 encoded bytes.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Output only. The endpoint of the cluster's API server.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// Allows clients to perform consistent read-modify-writes through optimistic concurrency control. May be sent on update
	// and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Fleet configuration.
	Fleet AwsClusterFleetOutput `pulumi:"fleet"`
	// The location for the resource
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of this resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Required. Cluster-wide networking configuration.
	Networking AwsClusterNetworkingOutput `pulumi:"networking"`
	// The project for the resource
	Project pulumi.StringOutput `pulumi:"project"`
	// Output only. If set, there are currently changes in flight to the cluster.
	Reconciling pulumi.BoolOutput `pulumi:"reconciling"`
	// Output only. The current state of the cluster. Possible values: STATE_UNSPECIFIED, PROVISIONING, RUNNING, RECONCILING,
	// STOPPING, ERROR, DEGRADED
	State pulumi.StringOutput `pulumi:"state"`
	// Output only. A globally unique identifier for the cluster.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// Output only. The time at which this cluster was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// Output only. Workload Identity settings.
	WorkloadIdentityConfigs AwsClusterWorkloadIdentityConfigArrayOutput `pulumi:"workloadIdentityConfigs"`
}

// NewAwsCluster registers a new resource with the given unique name, arguments, and options.
func NewAwsCluster(ctx *pulumi.Context,
	name string, args *AwsClusterArgs, opts ...pulumi.ResourceOption) (*AwsCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Authorization == nil {
		return nil, errors.New("invalid value for required argument 'Authorization'")
	}
	if args.AwsRegion == nil {
		return nil, errors.New("invalid value for required argument 'AwsRegion'")
	}
	if args.ControlPlane == nil {
		return nil, errors.New("invalid value for required argument 'ControlPlane'")
	}
	if args.Fleet == nil {
		return nil, errors.New("invalid value for required argument 'Fleet'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Networking == nil {
		return nil, errors.New("invalid value for required argument 'Networking'")
	}
	var resource AwsCluster
	err := ctx.RegisterResource("gcp:container/awsCluster:AwsCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAwsCluster gets an existing AwsCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAwsCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AwsClusterState, opts ...pulumi.ResourceOption) (*AwsCluster, error) {
	var resource AwsCluster
	err := ctx.ReadResource("gcp:container/awsCluster:AwsCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AwsCluster resources.
type awsClusterState struct {
	// Optional. Annotations on the cluster. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Key can have 2 segments: prefix (optional) and name (required), separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.
	Annotations map[string]string `pulumi:"annotations"`
	// Required. Configuration related to the cluster RBAC settings.
	Authorization *AwsClusterAuthorization `pulumi:"authorization"`
	// Required. The AWS region where the cluster runs. Each Google Cloud region supports a subset of nearby AWS regions. You can call to list all supported AWS regions within a given Google Cloud region.
	AwsRegion *string `pulumi:"awsRegion"`
	// Required. Configuration related to the cluster control plane.
	ControlPlane *AwsClusterControlPlane `pulumi:"controlPlane"`
	// Output only. The time at which this cluster was created.
	CreateTime *string `pulumi:"createTime"`
	// Optional. A human readable description of this cluster. Cannot be longer than 255 UTF-8 encoded bytes.
	Description *string `pulumi:"description"`
	// Output only. The endpoint of the cluster's API server.
	Endpoint *string `pulumi:"endpoint"`
	// Allows clients to perform consistent read-modify-writes through optimistic concurrency control. May be sent on update
	// and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag *string `pulumi:"etag"`
	// Fleet configuration.
	Fleet *AwsClusterFleet `pulumi:"fleet"`
	// The location for the resource
	Location *string `pulumi:"location"`
	// The name of this resource.
	Name *string `pulumi:"name"`
	// Required. Cluster-wide networking configuration.
	Networking *AwsClusterNetworking `pulumi:"networking"`
	// The project for the resource
	Project *string `pulumi:"project"`
	// Output only. If set, there are currently changes in flight to the cluster.
	Reconciling *bool `pulumi:"reconciling"`
	// Output only. The current state of the cluster. Possible values: STATE_UNSPECIFIED, PROVISIONING, RUNNING, RECONCILING,
	// STOPPING, ERROR, DEGRADED
	State *string `pulumi:"state"`
	// Output only. A globally unique identifier for the cluster.
	Uid *string `pulumi:"uid"`
	// Output only. The time at which this cluster was last updated.
	UpdateTime *string `pulumi:"updateTime"`
	// Output only. Workload Identity settings.
	WorkloadIdentityConfigs []AwsClusterWorkloadIdentityConfig `pulumi:"workloadIdentityConfigs"`
}

type AwsClusterState struct {
	// Optional. Annotations on the cluster. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Key can have 2 segments: prefix (optional) and name (required), separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.
	Annotations pulumi.StringMapInput
	// Required. Configuration related to the cluster RBAC settings.
	Authorization AwsClusterAuthorizationPtrInput
	// Required. The AWS region where the cluster runs. Each Google Cloud region supports a subset of nearby AWS regions. You can call to list all supported AWS regions within a given Google Cloud region.
	AwsRegion pulumi.StringPtrInput
	// Required. Configuration related to the cluster control plane.
	ControlPlane AwsClusterControlPlanePtrInput
	// Output only. The time at which this cluster was created.
	CreateTime pulumi.StringPtrInput
	// Optional. A human readable description of this cluster. Cannot be longer than 255 UTF-8 encoded bytes.
	Description pulumi.StringPtrInput
	// Output only. The endpoint of the cluster's API server.
	Endpoint pulumi.StringPtrInput
	// Allows clients to perform consistent read-modify-writes through optimistic concurrency control. May be sent on update
	// and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag pulumi.StringPtrInput
	// Fleet configuration.
	Fleet AwsClusterFleetPtrInput
	// The location for the resource
	Location pulumi.StringPtrInput
	// The name of this resource.
	Name pulumi.StringPtrInput
	// Required. Cluster-wide networking configuration.
	Networking AwsClusterNetworkingPtrInput
	// The project for the resource
	Project pulumi.StringPtrInput
	// Output only. If set, there are currently changes in flight to the cluster.
	Reconciling pulumi.BoolPtrInput
	// Output only. The current state of the cluster. Possible values: STATE_UNSPECIFIED, PROVISIONING, RUNNING, RECONCILING,
	// STOPPING, ERROR, DEGRADED
	State pulumi.StringPtrInput
	// Output only. A globally unique identifier for the cluster.
	Uid pulumi.StringPtrInput
	// Output only. The time at which this cluster was last updated.
	UpdateTime pulumi.StringPtrInput
	// Output only. Workload Identity settings.
	WorkloadIdentityConfigs AwsClusterWorkloadIdentityConfigArrayInput
}

func (AwsClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*awsClusterState)(nil)).Elem()
}

type awsClusterArgs struct {
	// Optional. Annotations on the cluster. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Key can have 2 segments: prefix (optional) and name (required), separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.
	Annotations map[string]string `pulumi:"annotations"`
	// Required. Configuration related to the cluster RBAC settings.
	Authorization AwsClusterAuthorization `pulumi:"authorization"`
	// Required. The AWS region where the cluster runs. Each Google Cloud region supports a subset of nearby AWS regions. You can call to list all supported AWS regions within a given Google Cloud region.
	AwsRegion string `pulumi:"awsRegion"`
	// Required. Configuration related to the cluster control plane.
	ControlPlane AwsClusterControlPlane `pulumi:"controlPlane"`
	// Optional. A human readable description of this cluster. Cannot be longer than 255 UTF-8 encoded bytes.
	Description *string `pulumi:"description"`
	// Fleet configuration.
	Fleet AwsClusterFleet `pulumi:"fleet"`
	// The location for the resource
	Location string `pulumi:"location"`
	// The name of this resource.
	Name *string `pulumi:"name"`
	// Required. Cluster-wide networking configuration.
	Networking AwsClusterNetworking `pulumi:"networking"`
	// The project for the resource
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a AwsCluster resource.
type AwsClusterArgs struct {
	// Optional. Annotations on the cluster. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Key can have 2 segments: prefix (optional) and name (required), separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.
	Annotations pulumi.StringMapInput
	// Required. Configuration related to the cluster RBAC settings.
	Authorization AwsClusterAuthorizationInput
	// Required. The AWS region where the cluster runs. Each Google Cloud region supports a subset of nearby AWS regions. You can call to list all supported AWS regions within a given Google Cloud region.
	AwsRegion pulumi.StringInput
	// Required. Configuration related to the cluster control plane.
	ControlPlane AwsClusterControlPlaneInput
	// Optional. A human readable description of this cluster. Cannot be longer than 255 UTF-8 encoded bytes.
	Description pulumi.StringPtrInput
	// Fleet configuration.
	Fleet AwsClusterFleetInput
	// The location for the resource
	Location pulumi.StringInput
	// The name of this resource.
	Name pulumi.StringPtrInput
	// Required. Cluster-wide networking configuration.
	Networking AwsClusterNetworkingInput
	// The project for the resource
	Project pulumi.StringPtrInput
}

func (AwsClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*awsClusterArgs)(nil)).Elem()
}

type AwsClusterInput interface {
	pulumi.Input

	ToAwsClusterOutput() AwsClusterOutput
	ToAwsClusterOutputWithContext(ctx context.Context) AwsClusterOutput
}

func (*AwsCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**AwsCluster)(nil)).Elem()
}

func (i *AwsCluster) ToAwsClusterOutput() AwsClusterOutput {
	return i.ToAwsClusterOutputWithContext(context.Background())
}

func (i *AwsCluster) ToAwsClusterOutputWithContext(ctx context.Context) AwsClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsClusterOutput)
}

// AwsClusterArrayInput is an input type that accepts AwsClusterArray and AwsClusterArrayOutput values.
// You can construct a concrete instance of `AwsClusterArrayInput` via:
//
//          AwsClusterArray{ AwsClusterArgs{...} }
type AwsClusterArrayInput interface {
	pulumi.Input

	ToAwsClusterArrayOutput() AwsClusterArrayOutput
	ToAwsClusterArrayOutputWithContext(context.Context) AwsClusterArrayOutput
}

type AwsClusterArray []AwsClusterInput

func (AwsClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AwsCluster)(nil)).Elem()
}

func (i AwsClusterArray) ToAwsClusterArrayOutput() AwsClusterArrayOutput {
	return i.ToAwsClusterArrayOutputWithContext(context.Background())
}

func (i AwsClusterArray) ToAwsClusterArrayOutputWithContext(ctx context.Context) AwsClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsClusterArrayOutput)
}

// AwsClusterMapInput is an input type that accepts AwsClusterMap and AwsClusterMapOutput values.
// You can construct a concrete instance of `AwsClusterMapInput` via:
//
//          AwsClusterMap{ "key": AwsClusterArgs{...} }
type AwsClusterMapInput interface {
	pulumi.Input

	ToAwsClusterMapOutput() AwsClusterMapOutput
	ToAwsClusterMapOutputWithContext(context.Context) AwsClusterMapOutput
}

type AwsClusterMap map[string]AwsClusterInput

func (AwsClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AwsCluster)(nil)).Elem()
}

func (i AwsClusterMap) ToAwsClusterMapOutput() AwsClusterMapOutput {
	return i.ToAwsClusterMapOutputWithContext(context.Background())
}

func (i AwsClusterMap) ToAwsClusterMapOutputWithContext(ctx context.Context) AwsClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsClusterMapOutput)
}

type AwsClusterOutput struct{ *pulumi.OutputState }

func (AwsClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AwsCluster)(nil)).Elem()
}

func (o AwsClusterOutput) ToAwsClusterOutput() AwsClusterOutput {
	return o
}

func (o AwsClusterOutput) ToAwsClusterOutputWithContext(ctx context.Context) AwsClusterOutput {
	return o
}

type AwsClusterArrayOutput struct{ *pulumi.OutputState }

func (AwsClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AwsCluster)(nil)).Elem()
}

func (o AwsClusterArrayOutput) ToAwsClusterArrayOutput() AwsClusterArrayOutput {
	return o
}

func (o AwsClusterArrayOutput) ToAwsClusterArrayOutputWithContext(ctx context.Context) AwsClusterArrayOutput {
	return o
}

func (o AwsClusterArrayOutput) Index(i pulumi.IntInput) AwsClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AwsCluster {
		return vs[0].([]*AwsCluster)[vs[1].(int)]
	}).(AwsClusterOutput)
}

type AwsClusterMapOutput struct{ *pulumi.OutputState }

func (AwsClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AwsCluster)(nil)).Elem()
}

func (o AwsClusterMapOutput) ToAwsClusterMapOutput() AwsClusterMapOutput {
	return o
}

func (o AwsClusterMapOutput) ToAwsClusterMapOutputWithContext(ctx context.Context) AwsClusterMapOutput {
	return o
}

func (o AwsClusterMapOutput) MapIndex(k pulumi.StringInput) AwsClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AwsCluster {
		return vs[0].(map[string]*AwsCluster)[vs[1].(string)]
	}).(AwsClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AwsClusterInput)(nil)).Elem(), &AwsCluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*AwsClusterArrayInput)(nil)).Elem(), AwsClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AwsClusterMapInput)(nil)).Elem(), AwsClusterMap{})
	pulumi.RegisterOutputType(AwsClusterOutput{})
	pulumi.RegisterOutputType(AwsClusterArrayOutput{})
	pulumi.RegisterOutputType(AwsClusterMapOutput{})
}
