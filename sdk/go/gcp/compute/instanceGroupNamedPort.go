// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Mange the named ports setting for a managed instance group without
// managing the group as whole. This resource is primarily intended for use
// with GKE-generated groups that shouldn't otherwise be managed by other
// tools.
//
// To get more information about InstanceGroupNamedPort, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/instanceGroup)
// * How-to Guides
//   - [Official Documentation](https://cloud.google.com/compute/docs/instance-groups/)
//
// ## Example Usage
// ### Instance Group Named Port Gke
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/container"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			containerNetwork, err := compute.NewNetwork(ctx, "containerNetwork", &compute.NetworkArgs{
//				AutoCreateSubnetworks: pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			containerSubnetwork, err := compute.NewSubnetwork(ctx, "containerSubnetwork", &compute.SubnetworkArgs{
//				Region:      pulumi.String("us-central1"),
//				Network:     containerNetwork.Name,
//				IpCidrRange: pulumi.String("10.0.36.0/24"),
//			})
//			if err != nil {
//				return err
//			}
//			myCluster, err := container.NewCluster(ctx, "myCluster", &container.ClusterArgs{
//				Location:         pulumi.String("us-central1-a"),
//				InitialNodeCount: pulumi.Int(1),
//				Network:          containerNetwork.Name,
//				Subnetwork:       containerSubnetwork.Name,
//				IpAllocationPolicy: &container.ClusterIpAllocationPolicyArgs{
//					ClusterIpv4CidrBlock:  pulumi.String("/19"),
//					ServicesIpv4CidrBlock: pulumi.String("/22"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewInstanceGroupNamedPort(ctx, "myPort", &compute.InstanceGroupNamedPortArgs{
//				Group: myCluster.NodePools.ApplyT(func(nodePools []container.ClusterNodePool) (*string, error) {
//					return &nodePools[0].InstanceGroupUrls[0], nil
//				}).(pulumi.StringPtrOutput),
//				Zone: pulumi.String("us-central1-a"),
//				Port: pulumi.Int(8080),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewInstanceGroupNamedPort(ctx, "myPorts", &compute.InstanceGroupNamedPortArgs{
//				Group: myCluster.NodePools.ApplyT(func(nodePools []container.ClusterNodePool) (*string, error) {
//					return &nodePools[0].InstanceGroupUrls[0], nil
//				}).(pulumi.StringPtrOutput),
//				Zone: pulumi.String("us-central1-a"),
//				Port: pulumi.Int(4443),
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
// # InstanceGroupNamedPort can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:compute/instanceGroupNamedPort:InstanceGroupNamedPort default projects/{{project}}/zones/{{zone}}/instanceGroups/{{group}}/{{port}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/instanceGroupNamedPort:InstanceGroupNamedPort default {{project}}/{{zone}}/{{group}}/{{port}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/instanceGroupNamedPort:InstanceGroupNamedPort default {{zone}}/{{group}}/{{port}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/instanceGroupNamedPort:InstanceGroupNamedPort default {{group}}/{{port}}/{{name}}
//
// ```
type InstanceGroupNamedPort struct {
	pulumi.CustomResourceState

	// The name of the instance group.
	Group pulumi.StringOutput `pulumi:"group"`
	// The name for this named port. The name must be 1-63 characters
	// long, and comply with RFC1035.
	Name pulumi.StringOutput `pulumi:"name"`
	// The port number, which can be a value between 1 and 65535.
	Port pulumi.IntOutput `pulumi:"port"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The zone of the instance group.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewInstanceGroupNamedPort registers a new resource with the given unique name, arguments, and options.
func NewInstanceGroupNamedPort(ctx *pulumi.Context,
	name string, args *InstanceGroupNamedPortArgs, opts ...pulumi.ResourceOption) (*InstanceGroupNamedPort, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Group == nil {
		return nil, errors.New("invalid value for required argument 'Group'")
	}
	if args.Port == nil {
		return nil, errors.New("invalid value for required argument 'Port'")
	}
	var resource InstanceGroupNamedPort
	err := ctx.RegisterResource("gcp:compute/instanceGroupNamedPort:InstanceGroupNamedPort", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceGroupNamedPort gets an existing InstanceGroupNamedPort resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceGroupNamedPort(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceGroupNamedPortState, opts ...pulumi.ResourceOption) (*InstanceGroupNamedPort, error) {
	var resource InstanceGroupNamedPort
	err := ctx.ReadResource("gcp:compute/instanceGroupNamedPort:InstanceGroupNamedPort", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceGroupNamedPort resources.
type instanceGroupNamedPortState struct {
	// The name of the instance group.
	Group *string `pulumi:"group"`
	// The name for this named port. The name must be 1-63 characters
	// long, and comply with RFC1035.
	Name *string `pulumi:"name"`
	// The port number, which can be a value between 1 and 65535.
	Port *int `pulumi:"port"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The zone of the instance group.
	Zone *string `pulumi:"zone"`
}

type InstanceGroupNamedPortState struct {
	// The name of the instance group.
	Group pulumi.StringPtrInput
	// The name for this named port. The name must be 1-63 characters
	// long, and comply with RFC1035.
	Name pulumi.StringPtrInput
	// The port number, which can be a value between 1 and 65535.
	Port pulumi.IntPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The zone of the instance group.
	Zone pulumi.StringPtrInput
}

func (InstanceGroupNamedPortState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceGroupNamedPortState)(nil)).Elem()
}

type instanceGroupNamedPortArgs struct {
	// The name of the instance group.
	Group string `pulumi:"group"`
	// The name for this named port. The name must be 1-63 characters
	// long, and comply with RFC1035.
	Name *string `pulumi:"name"`
	// The port number, which can be a value between 1 and 65535.
	Port int `pulumi:"port"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The zone of the instance group.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a InstanceGroupNamedPort resource.
type InstanceGroupNamedPortArgs struct {
	// The name of the instance group.
	Group pulumi.StringInput
	// The name for this named port. The name must be 1-63 characters
	// long, and comply with RFC1035.
	Name pulumi.StringPtrInput
	// The port number, which can be a value between 1 and 65535.
	Port pulumi.IntInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The zone of the instance group.
	Zone pulumi.StringPtrInput
}

func (InstanceGroupNamedPortArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceGroupNamedPortArgs)(nil)).Elem()
}

type InstanceGroupNamedPortInput interface {
	pulumi.Input

	ToInstanceGroupNamedPortOutput() InstanceGroupNamedPortOutput
	ToInstanceGroupNamedPortOutputWithContext(ctx context.Context) InstanceGroupNamedPortOutput
}

func (*InstanceGroupNamedPort) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceGroupNamedPort)(nil)).Elem()
}

func (i *InstanceGroupNamedPort) ToInstanceGroupNamedPortOutput() InstanceGroupNamedPortOutput {
	return i.ToInstanceGroupNamedPortOutputWithContext(context.Background())
}

func (i *InstanceGroupNamedPort) ToInstanceGroupNamedPortOutputWithContext(ctx context.Context) InstanceGroupNamedPortOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceGroupNamedPortOutput)
}

// InstanceGroupNamedPortArrayInput is an input type that accepts InstanceGroupNamedPortArray and InstanceGroupNamedPortArrayOutput values.
// You can construct a concrete instance of `InstanceGroupNamedPortArrayInput` via:
//
//	InstanceGroupNamedPortArray{ InstanceGroupNamedPortArgs{...} }
type InstanceGroupNamedPortArrayInput interface {
	pulumi.Input

	ToInstanceGroupNamedPortArrayOutput() InstanceGroupNamedPortArrayOutput
	ToInstanceGroupNamedPortArrayOutputWithContext(context.Context) InstanceGroupNamedPortArrayOutput
}

type InstanceGroupNamedPortArray []InstanceGroupNamedPortInput

func (InstanceGroupNamedPortArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceGroupNamedPort)(nil)).Elem()
}

func (i InstanceGroupNamedPortArray) ToInstanceGroupNamedPortArrayOutput() InstanceGroupNamedPortArrayOutput {
	return i.ToInstanceGroupNamedPortArrayOutputWithContext(context.Background())
}

func (i InstanceGroupNamedPortArray) ToInstanceGroupNamedPortArrayOutputWithContext(ctx context.Context) InstanceGroupNamedPortArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceGroupNamedPortArrayOutput)
}

// InstanceGroupNamedPortMapInput is an input type that accepts InstanceGroupNamedPortMap and InstanceGroupNamedPortMapOutput values.
// You can construct a concrete instance of `InstanceGroupNamedPortMapInput` via:
//
//	InstanceGroupNamedPortMap{ "key": InstanceGroupNamedPortArgs{...} }
type InstanceGroupNamedPortMapInput interface {
	pulumi.Input

	ToInstanceGroupNamedPortMapOutput() InstanceGroupNamedPortMapOutput
	ToInstanceGroupNamedPortMapOutputWithContext(context.Context) InstanceGroupNamedPortMapOutput
}

type InstanceGroupNamedPortMap map[string]InstanceGroupNamedPortInput

func (InstanceGroupNamedPortMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceGroupNamedPort)(nil)).Elem()
}

func (i InstanceGroupNamedPortMap) ToInstanceGroupNamedPortMapOutput() InstanceGroupNamedPortMapOutput {
	return i.ToInstanceGroupNamedPortMapOutputWithContext(context.Background())
}

func (i InstanceGroupNamedPortMap) ToInstanceGroupNamedPortMapOutputWithContext(ctx context.Context) InstanceGroupNamedPortMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceGroupNamedPortMapOutput)
}

type InstanceGroupNamedPortOutput struct{ *pulumi.OutputState }

func (InstanceGroupNamedPortOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceGroupNamedPort)(nil)).Elem()
}

func (o InstanceGroupNamedPortOutput) ToInstanceGroupNamedPortOutput() InstanceGroupNamedPortOutput {
	return o
}

func (o InstanceGroupNamedPortOutput) ToInstanceGroupNamedPortOutputWithContext(ctx context.Context) InstanceGroupNamedPortOutput {
	return o
}

// The name of the instance group.
func (o InstanceGroupNamedPortOutput) Group() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceGroupNamedPort) pulumi.StringOutput { return v.Group }).(pulumi.StringOutput)
}

// The name for this named port. The name must be 1-63 characters
// long, and comply with RFC1035.
func (o InstanceGroupNamedPortOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceGroupNamedPort) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The port number, which can be a value between 1 and 65535.
func (o InstanceGroupNamedPortOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *InstanceGroupNamedPort) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o InstanceGroupNamedPortOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceGroupNamedPort) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The zone of the instance group.
func (o InstanceGroupNamedPortOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *InstanceGroupNamedPort) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type InstanceGroupNamedPortArrayOutput struct{ *pulumi.OutputState }

func (InstanceGroupNamedPortArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*InstanceGroupNamedPort)(nil)).Elem()
}

func (o InstanceGroupNamedPortArrayOutput) ToInstanceGroupNamedPortArrayOutput() InstanceGroupNamedPortArrayOutput {
	return o
}

func (o InstanceGroupNamedPortArrayOutput) ToInstanceGroupNamedPortArrayOutputWithContext(ctx context.Context) InstanceGroupNamedPortArrayOutput {
	return o
}

func (o InstanceGroupNamedPortArrayOutput) Index(i pulumi.IntInput) InstanceGroupNamedPortOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *InstanceGroupNamedPort {
		return vs[0].([]*InstanceGroupNamedPort)[vs[1].(int)]
	}).(InstanceGroupNamedPortOutput)
}

type InstanceGroupNamedPortMapOutput struct{ *pulumi.OutputState }

func (InstanceGroupNamedPortMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*InstanceGroupNamedPort)(nil)).Elem()
}

func (o InstanceGroupNamedPortMapOutput) ToInstanceGroupNamedPortMapOutput() InstanceGroupNamedPortMapOutput {
	return o
}

func (o InstanceGroupNamedPortMapOutput) ToInstanceGroupNamedPortMapOutputWithContext(ctx context.Context) InstanceGroupNamedPortMapOutput {
	return o
}

func (o InstanceGroupNamedPortMapOutput) MapIndex(k pulumi.StringInput) InstanceGroupNamedPortOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *InstanceGroupNamedPort {
		return vs[0].(map[string]*InstanceGroupNamedPort)[vs[1].(string)]
	}).(InstanceGroupNamedPortOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceGroupNamedPortInput)(nil)).Elem(), &InstanceGroupNamedPort{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceGroupNamedPortArrayInput)(nil)).Elem(), InstanceGroupNamedPortArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceGroupNamedPortMapInput)(nil)).Elem(), InstanceGroupNamedPortMap{})
	pulumi.RegisterOutputType(InstanceGroupNamedPortOutput{})
	pulumi.RegisterOutputType(InstanceGroupNamedPortArrayOutput{})
	pulumi.RegisterOutputType(InstanceGroupNamedPortMapOutput{})
}
