// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gkeonprem

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// ## Example Usage
// ### Gkeonprem Vmware Node Pool Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/gkeonprem"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := gkeonprem.NewVMwareCluster(ctx, "default-basic", &gkeonprem.VMwareClusterArgs{
//				Location:               pulumi.String("us-west1"),
//				AdminClusterMembership: pulumi.String("projects/870316890899/locations/global/memberships/gkeonprem-terraform-test"),
//				Description:            pulumi.String("test cluster"),
//				OnPremVersion:          pulumi.String("1.13.1-gke.35"),
//				NetworkConfig: &gkeonprem.VMwareClusterNetworkConfigArgs{
//					ServiceAddressCidrBlocks: pulumi.StringArray{
//						pulumi.String("10.96.0.0/12"),
//					},
//					PodAddressCidrBlocks: pulumi.StringArray{
//						pulumi.String("192.168.0.0/16"),
//					},
//					DhcpIpConfig: &gkeonprem.VMwareClusterNetworkConfigDhcpIpConfigArgs{
//						Enabled: pulumi.Bool(true),
//					},
//				},
//				ControlPlaneNode: &gkeonprem.VMwareClusterControlPlaneNodeArgs{
//					Cpus:     pulumi.Int(4),
//					Memory:   pulumi.Int(8192),
//					Replicas: pulumi.Int(1),
//				},
//				LoadBalancer: &gkeonprem.VMwareClusterLoadBalancerArgs{
//					VipConfig: &gkeonprem.VMwareClusterLoadBalancerVipConfigArgs{
//						ControlPlaneVip: pulumi.String("10.251.133.5"),
//						IngressVip:      pulumi.String("10.251.135.19"),
//					},
//					MetalLbConfig: &gkeonprem.VMwareClusterLoadBalancerMetalLbConfigArgs{
//						AddressPools: gkeonprem.VMwareClusterLoadBalancerMetalLbConfigAddressPoolArray{
//							&gkeonprem.VMwareClusterLoadBalancerMetalLbConfigAddressPoolArgs{
//								Pool:         pulumi.String("ingress-ip"),
//								ManualAssign: pulumi.Bool(true),
//								Addresses: pulumi.StringArray{
//									pulumi.String("10.251.135.19"),
//								},
//							},
//							&gkeonprem.VMwareClusterLoadBalancerMetalLbConfigAddressPoolArgs{
//								Pool:         pulumi.String("lb-test-ip"),
//								ManualAssign: pulumi.Bool(true),
//								Addresses: pulumi.StringArray{
//									pulumi.String("10.251.135.19"),
//								},
//							},
//						},
//					},
//				},
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			_, err = gkeonprem.NewVMwareNodePool(ctx, "nodepool-basic", &gkeonprem.VMwareNodePoolArgs{
//				Location:      pulumi.String("us-west1"),
//				VmwareCluster: default_basic.Name,
//				Config: &gkeonprem.VMwareNodePoolConfigArgs{
//					Replicas:           pulumi.Int(3),
//					ImageType:          pulumi.String("ubuntu_containerd"),
//					EnableLoadBalancer: pulumi.Bool(true),
//				},
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Gkeonprem Vmware Node Pool Full
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/gkeonprem"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := gkeonprem.NewVMwareCluster(ctx, "default-full", &gkeonprem.VMwareClusterArgs{
//				Location:               pulumi.String("us-west1"),
//				AdminClusterMembership: pulumi.String("projects/870316890899/locations/global/memberships/gkeonprem-terraform-test"),
//				Description:            pulumi.String("test cluster"),
//				OnPremVersion:          pulumi.String("1.13.1-gke.35"),
//				NetworkConfig: &gkeonprem.VMwareClusterNetworkConfigArgs{
//					ServiceAddressCidrBlocks: pulumi.StringArray{
//						pulumi.String("10.96.0.0/12"),
//					},
//					PodAddressCidrBlocks: pulumi.StringArray{
//						pulumi.String("192.168.0.0/16"),
//					},
//					DhcpIpConfig: &gkeonprem.VMwareClusterNetworkConfigDhcpIpConfigArgs{
//						Enabled: pulumi.Bool(true),
//					},
//				},
//				ControlPlaneNode: &gkeonprem.VMwareClusterControlPlaneNodeArgs{
//					Cpus:     pulumi.Int(4),
//					Memory:   pulumi.Int(8192),
//					Replicas: pulumi.Int(1),
//				},
//				LoadBalancer: &gkeonprem.VMwareClusterLoadBalancerArgs{
//					VipConfig: &gkeonprem.VMwareClusterLoadBalancerVipConfigArgs{
//						ControlPlaneVip: pulumi.String("10.251.133.5"),
//						IngressVip:      pulumi.String("10.251.135.19"),
//					},
//					MetalLbConfig: &gkeonprem.VMwareClusterLoadBalancerMetalLbConfigArgs{
//						AddressPools: gkeonprem.VMwareClusterLoadBalancerMetalLbConfigAddressPoolArray{
//							&gkeonprem.VMwareClusterLoadBalancerMetalLbConfigAddressPoolArgs{
//								Pool:         pulumi.String("ingress-ip"),
//								ManualAssign: pulumi.Bool(true),
//								Addresses: pulumi.StringArray{
//									pulumi.String("10.251.135.19"),
//								},
//							},
//							&gkeonprem.VMwareClusterLoadBalancerMetalLbConfigAddressPoolArgs{
//								Pool:         pulumi.String("lb-test-ip"),
//								ManualAssign: pulumi.Bool(true),
//								Addresses: pulumi.StringArray{
//									pulumi.String("10.251.135.19"),
//								},
//							},
//						},
//					},
//				},
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			_, err = gkeonprem.NewVMwareNodePool(ctx, "nodepool-full", &gkeonprem.VMwareNodePoolArgs{
//				Location:      pulumi.String("us-west1"),
//				VmwareCluster: default_full.Name,
//				Annotations:   nil,
//				Config: &gkeonprem.VMwareNodePoolConfigArgs{
//					Cpus:           pulumi.Int(4),
//					MemoryMb:       pulumi.Int(8196),
//					Replicas:       pulumi.Int(3),
//					ImageType:      pulumi.String("ubuntu_containerd"),
//					Image:          pulumi.String("image"),
//					BootDiskSizeGb: pulumi.Int(10),
//					Taints: gkeonprem.VMwareNodePoolConfigTaintArray{
//						&gkeonprem.VMwareNodePoolConfigTaintArgs{
//							Key:   pulumi.String("key"),
//							Value: pulumi.String("value"),
//						},
//						&gkeonprem.VMwareNodePoolConfigTaintArgs{
//							Key:    pulumi.String("key"),
//							Value:  pulumi.String("value"),
//							Effect: pulumi.String("NO_SCHEDULE"),
//						},
//					},
//					Labels:             nil,
//					EnableLoadBalancer: pulumi.Bool(true),
//				},
//				NodePoolAutoscaling: &gkeonprem.VMwareNodePoolNodePoolAutoscalingArgs{
//					MinReplicas: pulumi.Int(1),
//					MaxReplicas: pulumi.Int(5),
//				},
//			}, pulumi.Provider(google_beta))
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
// # VmwareNodePool can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:gkeonprem/vMwareNodePool:VMwareNodePool default projects/{{project}}/locations/{{location}}/vmwareClusters/{{vmware_cluster}}/vmwareNodePools/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:gkeonprem/vMwareNodePool:VMwareNodePool default {{project}}/{{location}}/{{vmware_cluster}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:gkeonprem/vMwareNodePool:VMwareNodePool default {{location}}/{{vmware_cluster}}/{{name}}
//
// ```
type VMwareNodePool struct {
	pulumi.CustomResourceState

	// Annotations on the node Pool.
	// This field has the same restrictions as Kubernetes annotations.
	// The total size of all keys and values combined is limited to 256k.
	// Key can have 2 segments: prefix (optional) and name (required),
	// separated by a slash (/).
	// Prefix must be a DNS subdomain.
	// Name must be 63 characters or less, begin and end with alphanumerics,
	// with dashes (-), underscores (_), dots (.), and alphanumerics between.
	Annotations pulumi.StringMapOutput `pulumi:"annotations"`
	// The node configuration of the node pool.
	// Structure is documented below.
	Config VMwareNodePoolConfigOutput `pulumi:"config"`
	// The time the cluster was created, in RFC3339 text format.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The time the cluster was deleted, in RFC3339 text format.
	DeleteTime pulumi.StringOutput `pulumi:"deleteTime"`
	// The display name for the node pool.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// This checksum is computed by the server based on the value of other
	// fields, and may be sent on update and delete requests to ensure the
	// client has an up-to-date value before proceeding.
	// Allows clients to perform consistent read-modify-writes
	// through optimistic concurrency control.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The location of the resource.
	Location pulumi.StringOutput `pulumi:"location"`
	// The vmware node pool name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Node Pool autoscaling config for the node pool.
	// Structure is documented below.
	NodePoolAutoscaling VMwareNodePoolNodePoolAutoscalingPtrOutput `pulumi:"nodePoolAutoscaling"`
	// Anthos version for the node pool. Defaults to the user cluster version.
	OnPremVersion pulumi.StringOutput `pulumi:"onPremVersion"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// If set, there are currently changes in flight to the node pool.
	Reconciling pulumi.BoolOutput `pulumi:"reconciling"`
	// (Output)
	// The lifecycle state of the condition.
	State pulumi.StringOutput `pulumi:"state"`
	// ResourceStatus representing detailed cluster state.
	// Structure is documented below.
	Statuses VMwareNodePoolStatusArrayOutput `pulumi:"statuses"`
	// The unique identifier of the node pool.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// The time the cluster was last updated, in RFC3339 text format.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// The cluster this node pool belongs to.
	VmwareCluster pulumi.StringOutput `pulumi:"vmwareCluster"`
}

// NewVMwareNodePool registers a new resource with the given unique name, arguments, and options.
func NewVMwareNodePool(ctx *pulumi.Context,
	name string, args *VMwareNodePoolArgs, opts ...pulumi.ResourceOption) (*VMwareNodePool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Config == nil {
		return nil, errors.New("invalid value for required argument 'Config'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.VmwareCluster == nil {
		return nil, errors.New("invalid value for required argument 'VmwareCluster'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VMwareNodePool
	err := ctx.RegisterResource("gcp:gkeonprem/vMwareNodePool:VMwareNodePool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVMwareNodePool gets an existing VMwareNodePool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVMwareNodePool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VMwareNodePoolState, opts ...pulumi.ResourceOption) (*VMwareNodePool, error) {
	var resource VMwareNodePool
	err := ctx.ReadResource("gcp:gkeonprem/vMwareNodePool:VMwareNodePool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VMwareNodePool resources.
type vmwareNodePoolState struct {
	// Annotations on the node Pool.
	// This field has the same restrictions as Kubernetes annotations.
	// The total size of all keys and values combined is limited to 256k.
	// Key can have 2 segments: prefix (optional) and name (required),
	// separated by a slash (/).
	// Prefix must be a DNS subdomain.
	// Name must be 63 characters or less, begin and end with alphanumerics,
	// with dashes (-), underscores (_), dots (.), and alphanumerics between.
	Annotations map[string]string `pulumi:"annotations"`
	// The node configuration of the node pool.
	// Structure is documented below.
	Config *VMwareNodePoolConfig `pulumi:"config"`
	// The time the cluster was created, in RFC3339 text format.
	CreateTime *string `pulumi:"createTime"`
	// The time the cluster was deleted, in RFC3339 text format.
	DeleteTime *string `pulumi:"deleteTime"`
	// The display name for the node pool.
	DisplayName *string `pulumi:"displayName"`
	// This checksum is computed by the server based on the value of other
	// fields, and may be sent on update and delete requests to ensure the
	// client has an up-to-date value before proceeding.
	// Allows clients to perform consistent read-modify-writes
	// through optimistic concurrency control.
	Etag *string `pulumi:"etag"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The vmware node pool name.
	Name *string `pulumi:"name"`
	// Node Pool autoscaling config for the node pool.
	// Structure is documented below.
	NodePoolAutoscaling *VMwareNodePoolNodePoolAutoscaling `pulumi:"nodePoolAutoscaling"`
	// Anthos version for the node pool. Defaults to the user cluster version.
	OnPremVersion *string `pulumi:"onPremVersion"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// If set, there are currently changes in flight to the node pool.
	Reconciling *bool `pulumi:"reconciling"`
	// (Output)
	// The lifecycle state of the condition.
	State *string `pulumi:"state"`
	// ResourceStatus representing detailed cluster state.
	// Structure is documented below.
	Statuses []VMwareNodePoolStatus `pulumi:"statuses"`
	// The unique identifier of the node pool.
	Uid *string `pulumi:"uid"`
	// The time the cluster was last updated, in RFC3339 text format.
	UpdateTime *string `pulumi:"updateTime"`
	// The cluster this node pool belongs to.
	VmwareCluster *string `pulumi:"vmwareCluster"`
}

type VMwareNodePoolState struct {
	// Annotations on the node Pool.
	// This field has the same restrictions as Kubernetes annotations.
	// The total size of all keys and values combined is limited to 256k.
	// Key can have 2 segments: prefix (optional) and name (required),
	// separated by a slash (/).
	// Prefix must be a DNS subdomain.
	// Name must be 63 characters or less, begin and end with alphanumerics,
	// with dashes (-), underscores (_), dots (.), and alphanumerics between.
	Annotations pulumi.StringMapInput
	// The node configuration of the node pool.
	// Structure is documented below.
	Config VMwareNodePoolConfigPtrInput
	// The time the cluster was created, in RFC3339 text format.
	CreateTime pulumi.StringPtrInput
	// The time the cluster was deleted, in RFC3339 text format.
	DeleteTime pulumi.StringPtrInput
	// The display name for the node pool.
	DisplayName pulumi.StringPtrInput
	// This checksum is computed by the server based on the value of other
	// fields, and may be sent on update and delete requests to ensure the
	// client has an up-to-date value before proceeding.
	// Allows clients to perform consistent read-modify-writes
	// through optimistic concurrency control.
	Etag pulumi.StringPtrInput
	// The location of the resource.
	Location pulumi.StringPtrInput
	// The vmware node pool name.
	Name pulumi.StringPtrInput
	// Node Pool autoscaling config for the node pool.
	// Structure is documented below.
	NodePoolAutoscaling VMwareNodePoolNodePoolAutoscalingPtrInput
	// Anthos version for the node pool. Defaults to the user cluster version.
	OnPremVersion pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// If set, there are currently changes in flight to the node pool.
	Reconciling pulumi.BoolPtrInput
	// (Output)
	// The lifecycle state of the condition.
	State pulumi.StringPtrInput
	// ResourceStatus representing detailed cluster state.
	// Structure is documented below.
	Statuses VMwareNodePoolStatusArrayInput
	// The unique identifier of the node pool.
	Uid pulumi.StringPtrInput
	// The time the cluster was last updated, in RFC3339 text format.
	UpdateTime pulumi.StringPtrInput
	// The cluster this node pool belongs to.
	VmwareCluster pulumi.StringPtrInput
}

func (VMwareNodePoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*vmwareNodePoolState)(nil)).Elem()
}

type vmwareNodePoolArgs struct {
	// Annotations on the node Pool.
	// This field has the same restrictions as Kubernetes annotations.
	// The total size of all keys and values combined is limited to 256k.
	// Key can have 2 segments: prefix (optional) and name (required),
	// separated by a slash (/).
	// Prefix must be a DNS subdomain.
	// Name must be 63 characters or less, begin and end with alphanumerics,
	// with dashes (-), underscores (_), dots (.), and alphanumerics between.
	Annotations map[string]string `pulumi:"annotations"`
	// The node configuration of the node pool.
	// Structure is documented below.
	Config VMwareNodePoolConfig `pulumi:"config"`
	// The display name for the node pool.
	DisplayName *string `pulumi:"displayName"`
	// The location of the resource.
	Location string `pulumi:"location"`
	// The vmware node pool name.
	Name *string `pulumi:"name"`
	// Node Pool autoscaling config for the node pool.
	// Structure is documented below.
	NodePoolAutoscaling *VMwareNodePoolNodePoolAutoscaling `pulumi:"nodePoolAutoscaling"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The cluster this node pool belongs to.
	VmwareCluster string `pulumi:"vmwareCluster"`
}

// The set of arguments for constructing a VMwareNodePool resource.
type VMwareNodePoolArgs struct {
	// Annotations on the node Pool.
	// This field has the same restrictions as Kubernetes annotations.
	// The total size of all keys and values combined is limited to 256k.
	// Key can have 2 segments: prefix (optional) and name (required),
	// separated by a slash (/).
	// Prefix must be a DNS subdomain.
	// Name must be 63 characters or less, begin and end with alphanumerics,
	// with dashes (-), underscores (_), dots (.), and alphanumerics between.
	Annotations pulumi.StringMapInput
	// The node configuration of the node pool.
	// Structure is documented below.
	Config VMwareNodePoolConfigInput
	// The display name for the node pool.
	DisplayName pulumi.StringPtrInput
	// The location of the resource.
	Location pulumi.StringInput
	// The vmware node pool name.
	Name pulumi.StringPtrInput
	// Node Pool autoscaling config for the node pool.
	// Structure is documented below.
	NodePoolAutoscaling VMwareNodePoolNodePoolAutoscalingPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The cluster this node pool belongs to.
	VmwareCluster pulumi.StringInput
}

func (VMwareNodePoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vmwareNodePoolArgs)(nil)).Elem()
}

type VMwareNodePoolInput interface {
	pulumi.Input

	ToVMwareNodePoolOutput() VMwareNodePoolOutput
	ToVMwareNodePoolOutputWithContext(ctx context.Context) VMwareNodePoolOutput
}

func (*VMwareNodePool) ElementType() reflect.Type {
	return reflect.TypeOf((**VMwareNodePool)(nil)).Elem()
}

func (i *VMwareNodePool) ToVMwareNodePoolOutput() VMwareNodePoolOutput {
	return i.ToVMwareNodePoolOutputWithContext(context.Background())
}

func (i *VMwareNodePool) ToVMwareNodePoolOutputWithContext(ctx context.Context) VMwareNodePoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMwareNodePoolOutput)
}

func (i *VMwareNodePool) ToOutput(ctx context.Context) pulumix.Output[*VMwareNodePool] {
	return pulumix.Output[*VMwareNodePool]{
		OutputState: i.ToVMwareNodePoolOutputWithContext(ctx).OutputState,
	}
}

// VMwareNodePoolArrayInput is an input type that accepts VMwareNodePoolArray and VMwareNodePoolArrayOutput values.
// You can construct a concrete instance of `VMwareNodePoolArrayInput` via:
//
//	VMwareNodePoolArray{ VMwareNodePoolArgs{...} }
type VMwareNodePoolArrayInput interface {
	pulumi.Input

	ToVMwareNodePoolArrayOutput() VMwareNodePoolArrayOutput
	ToVMwareNodePoolArrayOutputWithContext(context.Context) VMwareNodePoolArrayOutput
}

type VMwareNodePoolArray []VMwareNodePoolInput

func (VMwareNodePoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VMwareNodePool)(nil)).Elem()
}

func (i VMwareNodePoolArray) ToVMwareNodePoolArrayOutput() VMwareNodePoolArrayOutput {
	return i.ToVMwareNodePoolArrayOutputWithContext(context.Background())
}

func (i VMwareNodePoolArray) ToVMwareNodePoolArrayOutputWithContext(ctx context.Context) VMwareNodePoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMwareNodePoolArrayOutput)
}

func (i VMwareNodePoolArray) ToOutput(ctx context.Context) pulumix.Output[[]*VMwareNodePool] {
	return pulumix.Output[[]*VMwareNodePool]{
		OutputState: i.ToVMwareNodePoolArrayOutputWithContext(ctx).OutputState,
	}
}

// VMwareNodePoolMapInput is an input type that accepts VMwareNodePoolMap and VMwareNodePoolMapOutput values.
// You can construct a concrete instance of `VMwareNodePoolMapInput` via:
//
//	VMwareNodePoolMap{ "key": VMwareNodePoolArgs{...} }
type VMwareNodePoolMapInput interface {
	pulumi.Input

	ToVMwareNodePoolMapOutput() VMwareNodePoolMapOutput
	ToVMwareNodePoolMapOutputWithContext(context.Context) VMwareNodePoolMapOutput
}

type VMwareNodePoolMap map[string]VMwareNodePoolInput

func (VMwareNodePoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VMwareNodePool)(nil)).Elem()
}

func (i VMwareNodePoolMap) ToVMwareNodePoolMapOutput() VMwareNodePoolMapOutput {
	return i.ToVMwareNodePoolMapOutputWithContext(context.Background())
}

func (i VMwareNodePoolMap) ToVMwareNodePoolMapOutputWithContext(ctx context.Context) VMwareNodePoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMwareNodePoolMapOutput)
}

func (i VMwareNodePoolMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*VMwareNodePool] {
	return pulumix.Output[map[string]*VMwareNodePool]{
		OutputState: i.ToVMwareNodePoolMapOutputWithContext(ctx).OutputState,
	}
}

type VMwareNodePoolOutput struct{ *pulumi.OutputState }

func (VMwareNodePoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VMwareNodePool)(nil)).Elem()
}

func (o VMwareNodePoolOutput) ToVMwareNodePoolOutput() VMwareNodePoolOutput {
	return o
}

func (o VMwareNodePoolOutput) ToVMwareNodePoolOutputWithContext(ctx context.Context) VMwareNodePoolOutput {
	return o
}

func (o VMwareNodePoolOutput) ToOutput(ctx context.Context) pulumix.Output[*VMwareNodePool] {
	return pulumix.Output[*VMwareNodePool]{
		OutputState: o.OutputState,
	}
}

// Annotations on the node Pool.
// This field has the same restrictions as Kubernetes annotations.
// The total size of all keys and values combined is limited to 256k.
// Key can have 2 segments: prefix (optional) and name (required),
// separated by a slash (/).
// Prefix must be a DNS subdomain.
// Name must be 63 characters or less, begin and end with alphanumerics,
// with dashes (-), underscores (_), dots (.), and alphanumerics between.
func (o VMwareNodePoolOutput) Annotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VMwareNodePool) pulumi.StringMapOutput { return v.Annotations }).(pulumi.StringMapOutput)
}

// The node configuration of the node pool.
// Structure is documented below.
func (o VMwareNodePoolOutput) Config() VMwareNodePoolConfigOutput {
	return o.ApplyT(func(v *VMwareNodePool) VMwareNodePoolConfigOutput { return v.Config }).(VMwareNodePoolConfigOutput)
}

// The time the cluster was created, in RFC3339 text format.
func (o VMwareNodePoolOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *VMwareNodePool) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The time the cluster was deleted, in RFC3339 text format.
func (o VMwareNodePoolOutput) DeleteTime() pulumi.StringOutput {
	return o.ApplyT(func(v *VMwareNodePool) pulumi.StringOutput { return v.DeleteTime }).(pulumi.StringOutput)
}

// The display name for the node pool.
func (o VMwareNodePoolOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMwareNodePool) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// This checksum is computed by the server based on the value of other
// fields, and may be sent on update and delete requests to ensure the
// client has an up-to-date value before proceeding.
// Allows clients to perform consistent read-modify-writes
// through optimistic concurrency control.
func (o VMwareNodePoolOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *VMwareNodePool) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// The location of the resource.
func (o VMwareNodePoolOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *VMwareNodePool) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The vmware node pool name.
func (o VMwareNodePoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VMwareNodePool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Node Pool autoscaling config for the node pool.
// Structure is documented below.
func (o VMwareNodePoolOutput) NodePoolAutoscaling() VMwareNodePoolNodePoolAutoscalingPtrOutput {
	return o.ApplyT(func(v *VMwareNodePool) VMwareNodePoolNodePoolAutoscalingPtrOutput { return v.NodePoolAutoscaling }).(VMwareNodePoolNodePoolAutoscalingPtrOutput)
}

// Anthos version for the node pool. Defaults to the user cluster version.
func (o VMwareNodePoolOutput) OnPremVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *VMwareNodePool) pulumi.StringOutput { return v.OnPremVersion }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o VMwareNodePoolOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *VMwareNodePool) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// If set, there are currently changes in flight to the node pool.
func (o VMwareNodePoolOutput) Reconciling() pulumi.BoolOutput {
	return o.ApplyT(func(v *VMwareNodePool) pulumi.BoolOutput { return v.Reconciling }).(pulumi.BoolOutput)
}

// (Output)
// The lifecycle state of the condition.
func (o VMwareNodePoolOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *VMwareNodePool) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// ResourceStatus representing detailed cluster state.
// Structure is documented below.
func (o VMwareNodePoolOutput) Statuses() VMwareNodePoolStatusArrayOutput {
	return o.ApplyT(func(v *VMwareNodePool) VMwareNodePoolStatusArrayOutput { return v.Statuses }).(VMwareNodePoolStatusArrayOutput)
}

// The unique identifier of the node pool.
func (o VMwareNodePoolOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *VMwareNodePool) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// The time the cluster was last updated, in RFC3339 text format.
func (o VMwareNodePoolOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *VMwareNodePool) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// The cluster this node pool belongs to.
func (o VMwareNodePoolOutput) VmwareCluster() pulumi.StringOutput {
	return o.ApplyT(func(v *VMwareNodePool) pulumi.StringOutput { return v.VmwareCluster }).(pulumi.StringOutput)
}

type VMwareNodePoolArrayOutput struct{ *pulumi.OutputState }

func (VMwareNodePoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VMwareNodePool)(nil)).Elem()
}

func (o VMwareNodePoolArrayOutput) ToVMwareNodePoolArrayOutput() VMwareNodePoolArrayOutput {
	return o
}

func (o VMwareNodePoolArrayOutput) ToVMwareNodePoolArrayOutputWithContext(ctx context.Context) VMwareNodePoolArrayOutput {
	return o
}

func (o VMwareNodePoolArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*VMwareNodePool] {
	return pulumix.Output[[]*VMwareNodePool]{
		OutputState: o.OutputState,
	}
}

func (o VMwareNodePoolArrayOutput) Index(i pulumi.IntInput) VMwareNodePoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VMwareNodePool {
		return vs[0].([]*VMwareNodePool)[vs[1].(int)]
	}).(VMwareNodePoolOutput)
}

type VMwareNodePoolMapOutput struct{ *pulumi.OutputState }

func (VMwareNodePoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VMwareNodePool)(nil)).Elem()
}

func (o VMwareNodePoolMapOutput) ToVMwareNodePoolMapOutput() VMwareNodePoolMapOutput {
	return o
}

func (o VMwareNodePoolMapOutput) ToVMwareNodePoolMapOutputWithContext(ctx context.Context) VMwareNodePoolMapOutput {
	return o
}

func (o VMwareNodePoolMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*VMwareNodePool] {
	return pulumix.Output[map[string]*VMwareNodePool]{
		OutputState: o.OutputState,
	}
}

func (o VMwareNodePoolMapOutput) MapIndex(k pulumi.StringInput) VMwareNodePoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VMwareNodePool {
		return vs[0].(map[string]*VMwareNodePool)[vs[1].(string)]
	}).(VMwareNodePoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VMwareNodePoolInput)(nil)).Elem(), &VMwareNodePool{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMwareNodePoolArrayInput)(nil)).Elem(), VMwareNodePoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMwareNodePoolMapInput)(nil)).Elem(), VMwareNodePoolMap{})
	pulumi.RegisterOutputType(VMwareNodePoolOutput{})
	pulumi.RegisterOutputType(VMwareNodePoolArrayOutput{})
	pulumi.RegisterOutputType(VMwareNodePoolMapOutput{})
}
