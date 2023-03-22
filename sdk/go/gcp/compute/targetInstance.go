// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a TargetInstance resource which defines an endpoint instance
// that terminates traffic of certain protocols. In particular, they are used
// in Protocol Forwarding, where forwarding rules can send packets to a
// non-NAT'ed target instance. Each target instance contains a single
// virtual machine instance that receives and handles traffic from the
// corresponding forwarding rules.
//
// To get more information about TargetInstance, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/v1/targetInstances)
// * How-to Guides
//   - [Using Protocol Forwarding](https://cloud.google.com/compute/docs/protocol-forwarding)
//
// ## Example Usage
// ### Target Instance Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			vmimage, err := compute.LookupImage(ctx, &compute.LookupImageArgs{
//				Family:  pulumi.StringRef("debian-11"),
//				Project: pulumi.StringRef("debian-cloud"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewInstance(ctx, "target-vm", &compute.InstanceArgs{
//				MachineType: pulumi.String("e2-medium"),
//				Zone:        pulumi.String("us-central1-a"),
//				BootDisk: &compute.InstanceBootDiskArgs{
//					InitializeParams: &compute.InstanceBootDiskInitializeParamsArgs{
//						Image: *pulumi.String(vmimage.SelfLink),
//					},
//				},
//				NetworkInterfaces: compute.InstanceNetworkInterfaceArray{
//					&compute.InstanceNetworkInterfaceArgs{
//						Network: pulumi.String("default"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewTargetInstance(ctx, "default", &compute.TargetInstanceArgs{
//				Instance: target_vm.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Target Instance Custom Network
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			target_vmNetwork, err := compute.LookupNetwork(ctx, &compute.LookupNetworkArgs{
//				Name: "default",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			vmimage, err := compute.LookupImage(ctx, &compute.LookupImageArgs{
//				Family:  pulumi.StringRef("debian-10"),
//				Project: pulumi.StringRef("debian-cloud"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewInstance(ctx, "target-vmInstance", &compute.InstanceArgs{
//				MachineType: pulumi.String("e2-medium"),
//				Zone:        pulumi.String("us-central1-a"),
//				BootDisk: &compute.InstanceBootDiskArgs{
//					InitializeParams: &compute.InstanceBootDiskInitializeParamsArgs{
//						Image: *pulumi.String(vmimage.SelfLink),
//					},
//				},
//				NetworkInterfaces: compute.InstanceNetworkInterfaceArray{
//					&compute.InstanceNetworkInterfaceArgs{
//						Network: pulumi.String("default"),
//					},
//				},
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewTargetInstance(ctx, "customNetwork", &compute.TargetInstanceArgs{
//				Instance: target_vmInstance.ID(),
//				Network:  *pulumi.String(target_vmNetwork.SelfLink),
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
// # TargetInstance can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:compute/targetInstance:TargetInstance default projects/{{project}}/zones/{{zone}}/targetInstances/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/targetInstance:TargetInstance default {{project}}/{{zone}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/targetInstance:TargetInstance default {{zone}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/targetInstance:TargetInstance default {{name}}
//
// ```
type TargetInstance struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The Compute instance VM handling traffic for this target instance.
	// Accepts the instance self-link, relative path
	// (e.g. `projects/project/zones/zone/instances/instance`) or name. If
	// name is given, the zone will default to the given zone or
	// the provider-default zone and the project will default to the
	// provider-level project.
	Instance pulumi.StringOutput `pulumi:"instance"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// NAT option controlling how IPs are NAT'ed to the instance.
	// Currently only NO_NAT (default value) is supported.
	// Default value is `NO_NAT`.
	// Possible values are `NO_NAT`.
	NatPolicy pulumi.StringPtrOutput `pulumi:"natPolicy"`
	// The URL of the network this target instance uses to forward traffic. If not specified, the traffic will be forwarded to the network that the default network interface belongs to.
	Network pulumi.StringPtrOutput `pulumi:"network"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// URL of the zone where the target instance resides.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewTargetInstance registers a new resource with the given unique name, arguments, and options.
func NewTargetInstance(ctx *pulumi.Context,
	name string, args *TargetInstanceArgs, opts ...pulumi.ResourceOption) (*TargetInstance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Instance == nil {
		return nil, errors.New("invalid value for required argument 'Instance'")
	}
	var resource TargetInstance
	err := ctx.RegisterResource("gcp:compute/targetInstance:TargetInstance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTargetInstance gets an existing TargetInstance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTargetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TargetInstanceState, opts ...pulumi.ResourceOption) (*TargetInstance, error) {
	var resource TargetInstance
	err := ctx.ReadResource("gcp:compute/targetInstance:TargetInstance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TargetInstance resources.
type targetInstanceState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// The Compute instance VM handling traffic for this target instance.
	// Accepts the instance self-link, relative path
	// (e.g. `projects/project/zones/zone/instances/instance`) or name. If
	// name is given, the zone will default to the given zone or
	// the provider-default zone and the project will default to the
	// provider-level project.
	Instance *string `pulumi:"instance"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// NAT option controlling how IPs are NAT'ed to the instance.
	// Currently only NO_NAT (default value) is supported.
	// Default value is `NO_NAT`.
	// Possible values are `NO_NAT`.
	NatPolicy *string `pulumi:"natPolicy"`
	// The URL of the network this target instance uses to forward traffic. If not specified, the traffic will be forwarded to the network that the default network interface belongs to.
	Network *string `pulumi:"network"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// URL of the zone where the target instance resides.
	Zone *string `pulumi:"zone"`
}

type TargetInstanceState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// The Compute instance VM handling traffic for this target instance.
	// Accepts the instance self-link, relative path
	// (e.g. `projects/project/zones/zone/instances/instance`) or name. If
	// name is given, the zone will default to the given zone or
	// the provider-default zone and the project will default to the
	// provider-level project.
	Instance pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// NAT option controlling how IPs are NAT'ed to the instance.
	// Currently only NO_NAT (default value) is supported.
	// Default value is `NO_NAT`.
	// Possible values are `NO_NAT`.
	NatPolicy pulumi.StringPtrInput
	// The URL of the network this target instance uses to forward traffic. If not specified, the traffic will be forwarded to the network that the default network interface belongs to.
	Network pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// URL of the zone where the target instance resides.
	Zone pulumi.StringPtrInput
}

func (TargetInstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*targetInstanceState)(nil)).Elem()
}

type targetInstanceArgs struct {
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// The Compute instance VM handling traffic for this target instance.
	// Accepts the instance self-link, relative path
	// (e.g. `projects/project/zones/zone/instances/instance`) or name. If
	// name is given, the zone will default to the given zone or
	// the provider-default zone and the project will default to the
	// provider-level project.
	Instance string `pulumi:"instance"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// NAT option controlling how IPs are NAT'ed to the instance.
	// Currently only NO_NAT (default value) is supported.
	// Default value is `NO_NAT`.
	// Possible values are `NO_NAT`.
	NatPolicy *string `pulumi:"natPolicy"`
	// The URL of the network this target instance uses to forward traffic. If not specified, the traffic will be forwarded to the network that the default network interface belongs to.
	Network *string `pulumi:"network"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// URL of the zone where the target instance resides.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a TargetInstance resource.
type TargetInstanceArgs struct {
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// The Compute instance VM handling traffic for this target instance.
	// Accepts the instance self-link, relative path
	// (e.g. `projects/project/zones/zone/instances/instance`) or name. If
	// name is given, the zone will default to the given zone or
	// the provider-default zone and the project will default to the
	// provider-level project.
	Instance pulumi.StringInput
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035. Specifically, the name must be 1-63 characters long and match
	// the regular expression `a-z?` which means the
	// first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the last
	// character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// NAT option controlling how IPs are NAT'ed to the instance.
	// Currently only NO_NAT (default value) is supported.
	// Default value is `NO_NAT`.
	// Possible values are `NO_NAT`.
	NatPolicy pulumi.StringPtrInput
	// The URL of the network this target instance uses to forward traffic. If not specified, the traffic will be forwarded to the network that the default network interface belongs to.
	Network pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// URL of the zone where the target instance resides.
	Zone pulumi.StringPtrInput
}

func (TargetInstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*targetInstanceArgs)(nil)).Elem()
}

type TargetInstanceInput interface {
	pulumi.Input

	ToTargetInstanceOutput() TargetInstanceOutput
	ToTargetInstanceOutputWithContext(ctx context.Context) TargetInstanceOutput
}

func (*TargetInstance) ElementType() reflect.Type {
	return reflect.TypeOf((**TargetInstance)(nil)).Elem()
}

func (i *TargetInstance) ToTargetInstanceOutput() TargetInstanceOutput {
	return i.ToTargetInstanceOutputWithContext(context.Background())
}

func (i *TargetInstance) ToTargetInstanceOutputWithContext(ctx context.Context) TargetInstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetInstanceOutput)
}

// TargetInstanceArrayInput is an input type that accepts TargetInstanceArray and TargetInstanceArrayOutput values.
// You can construct a concrete instance of `TargetInstanceArrayInput` via:
//
//	TargetInstanceArray{ TargetInstanceArgs{...} }
type TargetInstanceArrayInput interface {
	pulumi.Input

	ToTargetInstanceArrayOutput() TargetInstanceArrayOutput
	ToTargetInstanceArrayOutputWithContext(context.Context) TargetInstanceArrayOutput
}

type TargetInstanceArray []TargetInstanceInput

func (TargetInstanceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TargetInstance)(nil)).Elem()
}

func (i TargetInstanceArray) ToTargetInstanceArrayOutput() TargetInstanceArrayOutput {
	return i.ToTargetInstanceArrayOutputWithContext(context.Background())
}

func (i TargetInstanceArray) ToTargetInstanceArrayOutputWithContext(ctx context.Context) TargetInstanceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetInstanceArrayOutput)
}

// TargetInstanceMapInput is an input type that accepts TargetInstanceMap and TargetInstanceMapOutput values.
// You can construct a concrete instance of `TargetInstanceMapInput` via:
//
//	TargetInstanceMap{ "key": TargetInstanceArgs{...} }
type TargetInstanceMapInput interface {
	pulumi.Input

	ToTargetInstanceMapOutput() TargetInstanceMapOutput
	ToTargetInstanceMapOutputWithContext(context.Context) TargetInstanceMapOutput
}

type TargetInstanceMap map[string]TargetInstanceInput

func (TargetInstanceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TargetInstance)(nil)).Elem()
}

func (i TargetInstanceMap) ToTargetInstanceMapOutput() TargetInstanceMapOutput {
	return i.ToTargetInstanceMapOutputWithContext(context.Background())
}

func (i TargetInstanceMap) ToTargetInstanceMapOutputWithContext(ctx context.Context) TargetInstanceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetInstanceMapOutput)
}

type TargetInstanceOutput struct{ *pulumi.OutputState }

func (TargetInstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TargetInstance)(nil)).Elem()
}

func (o TargetInstanceOutput) ToTargetInstanceOutput() TargetInstanceOutput {
	return o
}

func (o TargetInstanceOutput) ToTargetInstanceOutputWithContext(ctx context.Context) TargetInstanceOutput {
	return o
}

// Creation timestamp in RFC3339 text format.
func (o TargetInstanceOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *TargetInstance) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource.
func (o TargetInstanceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TargetInstance) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The Compute instance VM handling traffic for this target instance.
// Accepts the instance self-link, relative path
// (e.g. `projects/project/zones/zone/instances/instance`) or name. If
// name is given, the zone will default to the given zone or
// the provider-default zone and the project will default to the
// provider-level project.
func (o TargetInstanceOutput) Instance() pulumi.StringOutput {
	return o.ApplyT(func(v *TargetInstance) pulumi.StringOutput { return v.Instance }).(pulumi.StringOutput)
}

// Name of the resource. Provided by the client when the resource is
// created. The name must be 1-63 characters long, and comply with
// RFC1035. Specifically, the name must be 1-63 characters long and match
// the regular expression `a-z?` which means the
// first character must be a lowercase letter, and all following
// characters must be a dash, lowercase letter, or digit, except the last
// character, which cannot be a dash.
func (o TargetInstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TargetInstance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// NAT option controlling how IPs are NAT'ed to the instance.
// Currently only NO_NAT (default value) is supported.
// Default value is `NO_NAT`.
// Possible values are `NO_NAT`.
func (o TargetInstanceOutput) NatPolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TargetInstance) pulumi.StringPtrOutput { return v.NatPolicy }).(pulumi.StringPtrOutput)
}

// The URL of the network this target instance uses to forward traffic. If not specified, the traffic will be forwarded to the network that the default network interface belongs to.
func (o TargetInstanceOutput) Network() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TargetInstance) pulumi.StringPtrOutput { return v.Network }).(pulumi.StringPtrOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o TargetInstanceOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *TargetInstance) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The URI of the created resource.
func (o TargetInstanceOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *TargetInstance) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// URL of the zone where the target instance resides.
func (o TargetInstanceOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *TargetInstance) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type TargetInstanceArrayOutput struct{ *pulumi.OutputState }

func (TargetInstanceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TargetInstance)(nil)).Elem()
}

func (o TargetInstanceArrayOutput) ToTargetInstanceArrayOutput() TargetInstanceArrayOutput {
	return o
}

func (o TargetInstanceArrayOutput) ToTargetInstanceArrayOutputWithContext(ctx context.Context) TargetInstanceArrayOutput {
	return o
}

func (o TargetInstanceArrayOutput) Index(i pulumi.IntInput) TargetInstanceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TargetInstance {
		return vs[0].([]*TargetInstance)[vs[1].(int)]
	}).(TargetInstanceOutput)
}

type TargetInstanceMapOutput struct{ *pulumi.OutputState }

func (TargetInstanceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TargetInstance)(nil)).Elem()
}

func (o TargetInstanceMapOutput) ToTargetInstanceMapOutput() TargetInstanceMapOutput {
	return o
}

func (o TargetInstanceMapOutput) ToTargetInstanceMapOutputWithContext(ctx context.Context) TargetInstanceMapOutput {
	return o
}

func (o TargetInstanceMapOutput) MapIndex(k pulumi.StringInput) TargetInstanceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TargetInstance {
		return vs[0].(map[string]*TargetInstance)[vs[1].(string)]
	}).(TargetInstanceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TargetInstanceInput)(nil)).Elem(), &TargetInstance{})
	pulumi.RegisterInputType(reflect.TypeOf((*TargetInstanceArrayInput)(nil)).Elem(), TargetInstanceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TargetInstanceMapInput)(nil)).Elem(), TargetInstanceMap{})
	pulumi.RegisterOutputType(TargetInstanceOutput{})
	pulumi.RegisterOutputType(TargetInstanceArrayOutput{})
	pulumi.RegisterOutputType(TargetInstanceMapOutput{})
}
