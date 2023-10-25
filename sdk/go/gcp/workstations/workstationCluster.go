// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package workstations

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// ## Example Usage
// ### Workstation Cluster Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/workstations"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			defaultNetwork, err := compute.NewNetwork(ctx, "defaultNetwork", &compute.NetworkArgs{
//				AutoCreateSubnetworks: pulumi.Bool(false),
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			defaultSubnetwork, err := compute.NewSubnetwork(ctx, "defaultSubnetwork", &compute.SubnetworkArgs{
//				IpCidrRange: pulumi.String("10.0.0.0/24"),
//				Region:      pulumi.String("us-central1"),
//				Network:     defaultNetwork.Name,
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			_, err = workstations.NewWorkstationCluster(ctx, "defaultWorkstationCluster", &workstations.WorkstationClusterArgs{
//				WorkstationClusterId: pulumi.String("workstation-cluster"),
//				Network:              defaultNetwork.ID(),
//				Subnetwork:           defaultSubnetwork.ID(),
//				Location:             pulumi.String("us-central1"),
//				Labels: pulumi.StringMap{
//					"label": pulumi.String("key"),
//				},
//				Annotations: pulumi.StringMap{
//					"label-one": pulumi.String("value-one"),
//				},
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			_, err = organizations.LookupProject(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Workstation Cluster Private
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/workstations"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			defaultNetwork, err := compute.NewNetwork(ctx, "defaultNetwork", &compute.NetworkArgs{
//				AutoCreateSubnetworks: pulumi.Bool(false),
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			defaultSubnetwork, err := compute.NewSubnetwork(ctx, "defaultSubnetwork", &compute.SubnetworkArgs{
//				IpCidrRange: pulumi.String("10.0.0.0/24"),
//				Region:      pulumi.String("us-central1"),
//				Network:     defaultNetwork.Name,
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			_, err = workstations.NewWorkstationCluster(ctx, "defaultWorkstationCluster", &workstations.WorkstationClusterArgs{
//				WorkstationClusterId: pulumi.String("workstation-cluster-private"),
//				Network:              defaultNetwork.ID(),
//				Subnetwork:           defaultSubnetwork.ID(),
//				Location:             pulumi.String("us-central1"),
//				PrivateClusterConfig: &workstations.WorkstationClusterPrivateClusterConfigArgs{
//					EnablePrivateEndpoint: pulumi.Bool(true),
//				},
//				Labels: pulumi.StringMap{
//					"label": pulumi.String("key"),
//				},
//				Annotations: pulumi.StringMap{
//					"label-one": pulumi.String("value-one"),
//				},
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			_, err = organizations.LookupProject(ctx, nil, nil)
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
// # WorkstationCluster can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:workstations/workstationCluster:WorkstationCluster default projects/{{project}}/locations/{{location}}/workstationClusters/{{workstation_cluster_id}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:workstations/workstationCluster:WorkstationCluster default {{project}}/{{location}}/{{workstation_cluster_id}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:workstations/workstationCluster:WorkstationCluster default {{location}}/{{workstation_cluster_id}}
//
// ```
type WorkstationCluster struct {
	pulumi.CustomResourceState

	// Client-specified annotations. This is distinct from labels.
	Annotations pulumi.StringMapOutput `pulumi:"annotations"`
	// Status conditions describing the current resource state.
	// Structure is documented below.
	Conditions WorkstationClusterConditionArrayOutput `pulumi:"conditions"`
	// Time when this resource was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Whether this resource is in degraded mode, in which case it may require user action to restore full functionality.
	// Details can be found in the conditions field.
	Degraded pulumi.BoolOutput `pulumi:"degraded"`
	// Human-readable name for this resource.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Checksum computed by the server.
	// May be sent on update and delete requests to ensure that the client has an up-to-date value before proceeding.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Client-specified labels that are applied to the resource and that are also propagated to the underlying Compute Engine resources.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The location where the workstation cluster should reside.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the cluster resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The relative resource name of the VPC network on which the instance can be accessed.
	// It is specified in the following form: "projects/{projectNumber}/global/networks/{network_id}".
	Network pulumi.StringOutput `pulumi:"network"`
	// Configuration for private cluster.
	// Structure is documented below.
	PrivateClusterConfig WorkstationClusterPrivateClusterConfigPtrOutput `pulumi:"privateClusterConfig"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Name of the Compute Engine subnetwork in which instances associated with this cluster will be created.
	// Must be part of the subnetwork specified for this cluster.
	Subnetwork pulumi.StringOutput `pulumi:"subnetwork"`
	// The system-generated UID of the resource.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// ID to use for the workstation cluster.
	//
	// ***
	WorkstationClusterId pulumi.StringOutput `pulumi:"workstationClusterId"`
}

// NewWorkstationCluster registers a new resource with the given unique name, arguments, and options.
func NewWorkstationCluster(ctx *pulumi.Context,
	name string, args *WorkstationClusterArgs, opts ...pulumi.ResourceOption) (*WorkstationCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Network == nil {
		return nil, errors.New("invalid value for required argument 'Network'")
	}
	if args.Subnetwork == nil {
		return nil, errors.New("invalid value for required argument 'Subnetwork'")
	}
	if args.WorkstationClusterId == nil {
		return nil, errors.New("invalid value for required argument 'WorkstationClusterId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource WorkstationCluster
	err := ctx.RegisterResource("gcp:workstations/workstationCluster:WorkstationCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkstationCluster gets an existing WorkstationCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkstationCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkstationClusterState, opts ...pulumi.ResourceOption) (*WorkstationCluster, error) {
	var resource WorkstationCluster
	err := ctx.ReadResource("gcp:workstations/workstationCluster:WorkstationCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkstationCluster resources.
type workstationClusterState struct {
	// Client-specified annotations. This is distinct from labels.
	Annotations map[string]string `pulumi:"annotations"`
	// Status conditions describing the current resource state.
	// Structure is documented below.
	Conditions []WorkstationClusterCondition `pulumi:"conditions"`
	// Time when this resource was created.
	CreateTime *string `pulumi:"createTime"`
	// Whether this resource is in degraded mode, in which case it may require user action to restore full functionality.
	// Details can be found in the conditions field.
	Degraded *bool `pulumi:"degraded"`
	// Human-readable name for this resource.
	DisplayName *string `pulumi:"displayName"`
	// Checksum computed by the server.
	// May be sent on update and delete requests to ensure that the client has an up-to-date value before proceeding.
	Etag *string `pulumi:"etag"`
	// Client-specified labels that are applied to the resource and that are also propagated to the underlying Compute Engine resources.
	Labels map[string]string `pulumi:"labels"`
	// The location where the workstation cluster should reside.
	Location *string `pulumi:"location"`
	// The name of the cluster resource.
	Name *string `pulumi:"name"`
	// The relative resource name of the VPC network on which the instance can be accessed.
	// It is specified in the following form: "projects/{projectNumber}/global/networks/{network_id}".
	Network *string `pulumi:"network"`
	// Configuration for private cluster.
	// Structure is documented below.
	PrivateClusterConfig *WorkstationClusterPrivateClusterConfig `pulumi:"privateClusterConfig"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Name of the Compute Engine subnetwork in which instances associated with this cluster will be created.
	// Must be part of the subnetwork specified for this cluster.
	Subnetwork *string `pulumi:"subnetwork"`
	// The system-generated UID of the resource.
	Uid *string `pulumi:"uid"`
	// ID to use for the workstation cluster.
	//
	// ***
	WorkstationClusterId *string `pulumi:"workstationClusterId"`
}

type WorkstationClusterState struct {
	// Client-specified annotations. This is distinct from labels.
	Annotations pulumi.StringMapInput
	// Status conditions describing the current resource state.
	// Structure is documented below.
	Conditions WorkstationClusterConditionArrayInput
	// Time when this resource was created.
	CreateTime pulumi.StringPtrInput
	// Whether this resource is in degraded mode, in which case it may require user action to restore full functionality.
	// Details can be found in the conditions field.
	Degraded pulumi.BoolPtrInput
	// Human-readable name for this resource.
	DisplayName pulumi.StringPtrInput
	// Checksum computed by the server.
	// May be sent on update and delete requests to ensure that the client has an up-to-date value before proceeding.
	Etag pulumi.StringPtrInput
	// Client-specified labels that are applied to the resource and that are also propagated to the underlying Compute Engine resources.
	Labels pulumi.StringMapInput
	// The location where the workstation cluster should reside.
	Location pulumi.StringPtrInput
	// The name of the cluster resource.
	Name pulumi.StringPtrInput
	// The relative resource name of the VPC network on which the instance can be accessed.
	// It is specified in the following form: "projects/{projectNumber}/global/networks/{network_id}".
	Network pulumi.StringPtrInput
	// Configuration for private cluster.
	// Structure is documented below.
	PrivateClusterConfig WorkstationClusterPrivateClusterConfigPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Name of the Compute Engine subnetwork in which instances associated with this cluster will be created.
	// Must be part of the subnetwork specified for this cluster.
	Subnetwork pulumi.StringPtrInput
	// The system-generated UID of the resource.
	Uid pulumi.StringPtrInput
	// ID to use for the workstation cluster.
	//
	// ***
	WorkstationClusterId pulumi.StringPtrInput
}

func (WorkstationClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*workstationClusterState)(nil)).Elem()
}

type workstationClusterArgs struct {
	// Client-specified annotations. This is distinct from labels.
	Annotations map[string]string `pulumi:"annotations"`
	// Human-readable name for this resource.
	DisplayName *string `pulumi:"displayName"`
	// Client-specified labels that are applied to the resource and that are also propagated to the underlying Compute Engine resources.
	Labels map[string]string `pulumi:"labels"`
	// The location where the workstation cluster should reside.
	Location *string `pulumi:"location"`
	// The relative resource name of the VPC network on which the instance can be accessed.
	// It is specified in the following form: "projects/{projectNumber}/global/networks/{network_id}".
	Network string `pulumi:"network"`
	// Configuration for private cluster.
	// Structure is documented below.
	PrivateClusterConfig *WorkstationClusterPrivateClusterConfig `pulumi:"privateClusterConfig"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Name of the Compute Engine subnetwork in which instances associated with this cluster will be created.
	// Must be part of the subnetwork specified for this cluster.
	Subnetwork string `pulumi:"subnetwork"`
	// ID to use for the workstation cluster.
	//
	// ***
	WorkstationClusterId string `pulumi:"workstationClusterId"`
}

// The set of arguments for constructing a WorkstationCluster resource.
type WorkstationClusterArgs struct {
	// Client-specified annotations. This is distinct from labels.
	Annotations pulumi.StringMapInput
	// Human-readable name for this resource.
	DisplayName pulumi.StringPtrInput
	// Client-specified labels that are applied to the resource and that are also propagated to the underlying Compute Engine resources.
	Labels pulumi.StringMapInput
	// The location where the workstation cluster should reside.
	Location pulumi.StringPtrInput
	// The relative resource name of the VPC network on which the instance can be accessed.
	// It is specified in the following form: "projects/{projectNumber}/global/networks/{network_id}".
	Network pulumi.StringInput
	// Configuration for private cluster.
	// Structure is documented below.
	PrivateClusterConfig WorkstationClusterPrivateClusterConfigPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Name of the Compute Engine subnetwork in which instances associated with this cluster will be created.
	// Must be part of the subnetwork specified for this cluster.
	Subnetwork pulumi.StringInput
	// ID to use for the workstation cluster.
	//
	// ***
	WorkstationClusterId pulumi.StringInput
}

func (WorkstationClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workstationClusterArgs)(nil)).Elem()
}

type WorkstationClusterInput interface {
	pulumi.Input

	ToWorkstationClusterOutput() WorkstationClusterOutput
	ToWorkstationClusterOutputWithContext(ctx context.Context) WorkstationClusterOutput
}

func (*WorkstationCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkstationCluster)(nil)).Elem()
}

func (i *WorkstationCluster) ToWorkstationClusterOutput() WorkstationClusterOutput {
	return i.ToWorkstationClusterOutputWithContext(context.Background())
}

func (i *WorkstationCluster) ToWorkstationClusterOutputWithContext(ctx context.Context) WorkstationClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkstationClusterOutput)
}

func (i *WorkstationCluster) ToOutput(ctx context.Context) pulumix.Output[*WorkstationCluster] {
	return pulumix.Output[*WorkstationCluster]{
		OutputState: i.ToWorkstationClusterOutputWithContext(ctx).OutputState,
	}
}

// WorkstationClusterArrayInput is an input type that accepts WorkstationClusterArray and WorkstationClusterArrayOutput values.
// You can construct a concrete instance of `WorkstationClusterArrayInput` via:
//
//	WorkstationClusterArray{ WorkstationClusterArgs{...} }
type WorkstationClusterArrayInput interface {
	pulumi.Input

	ToWorkstationClusterArrayOutput() WorkstationClusterArrayOutput
	ToWorkstationClusterArrayOutputWithContext(context.Context) WorkstationClusterArrayOutput
}

type WorkstationClusterArray []WorkstationClusterInput

func (WorkstationClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WorkstationCluster)(nil)).Elem()
}

func (i WorkstationClusterArray) ToWorkstationClusterArrayOutput() WorkstationClusterArrayOutput {
	return i.ToWorkstationClusterArrayOutputWithContext(context.Background())
}

func (i WorkstationClusterArray) ToWorkstationClusterArrayOutputWithContext(ctx context.Context) WorkstationClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkstationClusterArrayOutput)
}

func (i WorkstationClusterArray) ToOutput(ctx context.Context) pulumix.Output[[]*WorkstationCluster] {
	return pulumix.Output[[]*WorkstationCluster]{
		OutputState: i.ToWorkstationClusterArrayOutputWithContext(ctx).OutputState,
	}
}

// WorkstationClusterMapInput is an input type that accepts WorkstationClusterMap and WorkstationClusterMapOutput values.
// You can construct a concrete instance of `WorkstationClusterMapInput` via:
//
//	WorkstationClusterMap{ "key": WorkstationClusterArgs{...} }
type WorkstationClusterMapInput interface {
	pulumi.Input

	ToWorkstationClusterMapOutput() WorkstationClusterMapOutput
	ToWorkstationClusterMapOutputWithContext(context.Context) WorkstationClusterMapOutput
}

type WorkstationClusterMap map[string]WorkstationClusterInput

func (WorkstationClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WorkstationCluster)(nil)).Elem()
}

func (i WorkstationClusterMap) ToWorkstationClusterMapOutput() WorkstationClusterMapOutput {
	return i.ToWorkstationClusterMapOutputWithContext(context.Background())
}

func (i WorkstationClusterMap) ToWorkstationClusterMapOutputWithContext(ctx context.Context) WorkstationClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkstationClusterMapOutput)
}

func (i WorkstationClusterMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*WorkstationCluster] {
	return pulumix.Output[map[string]*WorkstationCluster]{
		OutputState: i.ToWorkstationClusterMapOutputWithContext(ctx).OutputState,
	}
}

type WorkstationClusterOutput struct{ *pulumi.OutputState }

func (WorkstationClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkstationCluster)(nil)).Elem()
}

func (o WorkstationClusterOutput) ToWorkstationClusterOutput() WorkstationClusterOutput {
	return o
}

func (o WorkstationClusterOutput) ToWorkstationClusterOutputWithContext(ctx context.Context) WorkstationClusterOutput {
	return o
}

func (o WorkstationClusterOutput) ToOutput(ctx context.Context) pulumix.Output[*WorkstationCluster] {
	return pulumix.Output[*WorkstationCluster]{
		OutputState: o.OutputState,
	}
}

// Client-specified annotations. This is distinct from labels.
func (o WorkstationClusterOutput) Annotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v *WorkstationCluster) pulumi.StringMapOutput { return v.Annotations }).(pulumi.StringMapOutput)
}

// Status conditions describing the current resource state.
// Structure is documented below.
func (o WorkstationClusterOutput) Conditions() WorkstationClusterConditionArrayOutput {
	return o.ApplyT(func(v *WorkstationCluster) WorkstationClusterConditionArrayOutput { return v.Conditions }).(WorkstationClusterConditionArrayOutput)
}

// Time when this resource was created.
func (o WorkstationClusterOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationCluster) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Whether this resource is in degraded mode, in which case it may require user action to restore full functionality.
// Details can be found in the conditions field.
func (o WorkstationClusterOutput) Degraded() pulumi.BoolOutput {
	return o.ApplyT(func(v *WorkstationCluster) pulumi.BoolOutput { return v.Degraded }).(pulumi.BoolOutput)
}

// Human-readable name for this resource.
func (o WorkstationClusterOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkstationCluster) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// Checksum computed by the server.
// May be sent on update and delete requests to ensure that the client has an up-to-date value before proceeding.
func (o WorkstationClusterOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationCluster) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Client-specified labels that are applied to the resource and that are also propagated to the underlying Compute Engine resources.
func (o WorkstationClusterOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *WorkstationCluster) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The location where the workstation cluster should reside.
func (o WorkstationClusterOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkstationCluster) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The name of the cluster resource.
func (o WorkstationClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The relative resource name of the VPC network on which the instance can be accessed.
// It is specified in the following form: "projects/{projectNumber}/global/networks/{network_id}".
func (o WorkstationClusterOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationCluster) pulumi.StringOutput { return v.Network }).(pulumi.StringOutput)
}

// Configuration for private cluster.
// Structure is documented below.
func (o WorkstationClusterOutput) PrivateClusterConfig() WorkstationClusterPrivateClusterConfigPtrOutput {
	return o.ApplyT(func(v *WorkstationCluster) WorkstationClusterPrivateClusterConfigPtrOutput {
		return v.PrivateClusterConfig
	}).(WorkstationClusterPrivateClusterConfigPtrOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o WorkstationClusterOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationCluster) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Name of the Compute Engine subnetwork in which instances associated with this cluster will be created.
// Must be part of the subnetwork specified for this cluster.
func (o WorkstationClusterOutput) Subnetwork() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationCluster) pulumi.StringOutput { return v.Subnetwork }).(pulumi.StringOutput)
}

// The system-generated UID of the resource.
func (o WorkstationClusterOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationCluster) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// ID to use for the workstation cluster.
//
// ***
func (o WorkstationClusterOutput) WorkstationClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkstationCluster) pulumi.StringOutput { return v.WorkstationClusterId }).(pulumi.StringOutput)
}

type WorkstationClusterArrayOutput struct{ *pulumi.OutputState }

func (WorkstationClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WorkstationCluster)(nil)).Elem()
}

func (o WorkstationClusterArrayOutput) ToWorkstationClusterArrayOutput() WorkstationClusterArrayOutput {
	return o
}

func (o WorkstationClusterArrayOutput) ToWorkstationClusterArrayOutputWithContext(ctx context.Context) WorkstationClusterArrayOutput {
	return o
}

func (o WorkstationClusterArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*WorkstationCluster] {
	return pulumix.Output[[]*WorkstationCluster]{
		OutputState: o.OutputState,
	}
}

func (o WorkstationClusterArrayOutput) Index(i pulumi.IntInput) WorkstationClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WorkstationCluster {
		return vs[0].([]*WorkstationCluster)[vs[1].(int)]
	}).(WorkstationClusterOutput)
}

type WorkstationClusterMapOutput struct{ *pulumi.OutputState }

func (WorkstationClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WorkstationCluster)(nil)).Elem()
}

func (o WorkstationClusterMapOutput) ToWorkstationClusterMapOutput() WorkstationClusterMapOutput {
	return o
}

func (o WorkstationClusterMapOutput) ToWorkstationClusterMapOutputWithContext(ctx context.Context) WorkstationClusterMapOutput {
	return o
}

func (o WorkstationClusterMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*WorkstationCluster] {
	return pulumix.Output[map[string]*WorkstationCluster]{
		OutputState: o.OutputState,
	}
}

func (o WorkstationClusterMapOutput) MapIndex(k pulumi.StringInput) WorkstationClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WorkstationCluster {
		return vs[0].(map[string]*WorkstationCluster)[vs[1].(string)]
	}).(WorkstationClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkstationClusterInput)(nil)).Elem(), &WorkstationCluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkstationClusterArrayInput)(nil)).Elem(), WorkstationClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkstationClusterMapInput)(nil)).Elem(), WorkstationClusterMap{})
	pulumi.RegisterOutputType(WorkstationClusterOutput{})
	pulumi.RegisterOutputType(WorkstationClusterArrayOutput{})
	pulumi.RegisterOutputType(WorkstationClusterMapOutput{})
}
