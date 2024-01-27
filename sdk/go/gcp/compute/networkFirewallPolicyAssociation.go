// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The Compute NetworkFirewallPolicyAssociation resource
//
// ## Example Usage
// ### Global
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/compute"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			networkFirewallPolicy, err := compute.NewNetworkFirewallPolicy(ctx, "networkFirewallPolicy", &compute.NetworkFirewallPolicyArgs{
//				Project:     pulumi.String("my-project-name"),
//				Description: pulumi.String("Sample global network firewall policy"),
//			})
//			if err != nil {
//				return err
//			}
//			network, err := compute.NewNetwork(ctx, "network", nil)
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewNetworkFirewallPolicyAssociation(ctx, "primary", &compute.NetworkFirewallPolicyAssociationArgs{
//				AttachmentTarget: network.ID(),
//				FirewallPolicy:   networkFirewallPolicy.Name,
//				Project:          pulumi.String("my-project-name"),
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
// NetworkFirewallPolicyAssociation can be imported using any of these accepted formats* `projects/{{project}}/global/firewallPolicies/{{firewall_policy}}/associations/{{name}}` * `{{project}}/{{firewall_policy}}/{{name}}` When using the `pulumi import` command, NetworkFirewallPolicyAssociation can be imported using one of the formats above. For example
//
// ```sh
//
//	$ pulumi import gcp:compute/networkFirewallPolicyAssociation:NetworkFirewallPolicyAssociation default projects/{{project}}/global/firewallPolicies/{{firewall_policy}}/associations/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/networkFirewallPolicyAssociation:NetworkFirewallPolicyAssociation default {{project}}/{{firewall_policy}}/{{name}}
//
// ```
type NetworkFirewallPolicyAssociation struct {
	pulumi.CustomResourceState

	// The target that the firewall policy is attached to.
	AttachmentTarget pulumi.StringOutput `pulumi:"attachmentTarget"`
	// The firewall policy ID of the association.
	FirewallPolicy pulumi.StringOutput `pulumi:"firewallPolicy"`
	// The name for an association.
	//
	// ***
	Name pulumi.StringOutput `pulumi:"name"`
	// The project for the resource
	Project pulumi.StringOutput `pulumi:"project"`
	// The short name of the firewall policy of the association.
	ShortName pulumi.StringOutput `pulumi:"shortName"`
}

// NewNetworkFirewallPolicyAssociation registers a new resource with the given unique name, arguments, and options.
func NewNetworkFirewallPolicyAssociation(ctx *pulumi.Context,
	name string, args *NetworkFirewallPolicyAssociationArgs, opts ...pulumi.ResourceOption) (*NetworkFirewallPolicyAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AttachmentTarget == nil {
		return nil, errors.New("invalid value for required argument 'AttachmentTarget'")
	}
	if args.FirewallPolicy == nil {
		return nil, errors.New("invalid value for required argument 'FirewallPolicy'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NetworkFirewallPolicyAssociation
	err := ctx.RegisterResource("gcp:compute/networkFirewallPolicyAssociation:NetworkFirewallPolicyAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkFirewallPolicyAssociation gets an existing NetworkFirewallPolicyAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkFirewallPolicyAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkFirewallPolicyAssociationState, opts ...pulumi.ResourceOption) (*NetworkFirewallPolicyAssociation, error) {
	var resource NetworkFirewallPolicyAssociation
	err := ctx.ReadResource("gcp:compute/networkFirewallPolicyAssociation:NetworkFirewallPolicyAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkFirewallPolicyAssociation resources.
type networkFirewallPolicyAssociationState struct {
	// The target that the firewall policy is attached to.
	AttachmentTarget *string `pulumi:"attachmentTarget"`
	// The firewall policy ID of the association.
	FirewallPolicy *string `pulumi:"firewallPolicy"`
	// The name for an association.
	//
	// ***
	Name *string `pulumi:"name"`
	// The project for the resource
	Project *string `pulumi:"project"`
	// The short name of the firewall policy of the association.
	ShortName *string `pulumi:"shortName"`
}

type NetworkFirewallPolicyAssociationState struct {
	// The target that the firewall policy is attached to.
	AttachmentTarget pulumi.StringPtrInput
	// The firewall policy ID of the association.
	FirewallPolicy pulumi.StringPtrInput
	// The name for an association.
	//
	// ***
	Name pulumi.StringPtrInput
	// The project for the resource
	Project pulumi.StringPtrInput
	// The short name of the firewall policy of the association.
	ShortName pulumi.StringPtrInput
}

func (NetworkFirewallPolicyAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkFirewallPolicyAssociationState)(nil)).Elem()
}

type networkFirewallPolicyAssociationArgs struct {
	// The target that the firewall policy is attached to.
	AttachmentTarget string `pulumi:"attachmentTarget"`
	// The firewall policy ID of the association.
	FirewallPolicy string `pulumi:"firewallPolicy"`
	// The name for an association.
	//
	// ***
	Name *string `pulumi:"name"`
	// The project for the resource
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a NetworkFirewallPolicyAssociation resource.
type NetworkFirewallPolicyAssociationArgs struct {
	// The target that the firewall policy is attached to.
	AttachmentTarget pulumi.StringInput
	// The firewall policy ID of the association.
	FirewallPolicy pulumi.StringInput
	// The name for an association.
	//
	// ***
	Name pulumi.StringPtrInput
	// The project for the resource
	Project pulumi.StringPtrInput
}

func (NetworkFirewallPolicyAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkFirewallPolicyAssociationArgs)(nil)).Elem()
}

type NetworkFirewallPolicyAssociationInput interface {
	pulumi.Input

	ToNetworkFirewallPolicyAssociationOutput() NetworkFirewallPolicyAssociationOutput
	ToNetworkFirewallPolicyAssociationOutputWithContext(ctx context.Context) NetworkFirewallPolicyAssociationOutput
}

func (*NetworkFirewallPolicyAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkFirewallPolicyAssociation)(nil)).Elem()
}

func (i *NetworkFirewallPolicyAssociation) ToNetworkFirewallPolicyAssociationOutput() NetworkFirewallPolicyAssociationOutput {
	return i.ToNetworkFirewallPolicyAssociationOutputWithContext(context.Background())
}

func (i *NetworkFirewallPolicyAssociation) ToNetworkFirewallPolicyAssociationOutputWithContext(ctx context.Context) NetworkFirewallPolicyAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkFirewallPolicyAssociationOutput)
}

// NetworkFirewallPolicyAssociationArrayInput is an input type that accepts NetworkFirewallPolicyAssociationArray and NetworkFirewallPolicyAssociationArrayOutput values.
// You can construct a concrete instance of `NetworkFirewallPolicyAssociationArrayInput` via:
//
//	NetworkFirewallPolicyAssociationArray{ NetworkFirewallPolicyAssociationArgs{...} }
type NetworkFirewallPolicyAssociationArrayInput interface {
	pulumi.Input

	ToNetworkFirewallPolicyAssociationArrayOutput() NetworkFirewallPolicyAssociationArrayOutput
	ToNetworkFirewallPolicyAssociationArrayOutputWithContext(context.Context) NetworkFirewallPolicyAssociationArrayOutput
}

type NetworkFirewallPolicyAssociationArray []NetworkFirewallPolicyAssociationInput

func (NetworkFirewallPolicyAssociationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkFirewallPolicyAssociation)(nil)).Elem()
}

func (i NetworkFirewallPolicyAssociationArray) ToNetworkFirewallPolicyAssociationArrayOutput() NetworkFirewallPolicyAssociationArrayOutput {
	return i.ToNetworkFirewallPolicyAssociationArrayOutputWithContext(context.Background())
}

func (i NetworkFirewallPolicyAssociationArray) ToNetworkFirewallPolicyAssociationArrayOutputWithContext(ctx context.Context) NetworkFirewallPolicyAssociationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkFirewallPolicyAssociationArrayOutput)
}

// NetworkFirewallPolicyAssociationMapInput is an input type that accepts NetworkFirewallPolicyAssociationMap and NetworkFirewallPolicyAssociationMapOutput values.
// You can construct a concrete instance of `NetworkFirewallPolicyAssociationMapInput` via:
//
//	NetworkFirewallPolicyAssociationMap{ "key": NetworkFirewallPolicyAssociationArgs{...} }
type NetworkFirewallPolicyAssociationMapInput interface {
	pulumi.Input

	ToNetworkFirewallPolicyAssociationMapOutput() NetworkFirewallPolicyAssociationMapOutput
	ToNetworkFirewallPolicyAssociationMapOutputWithContext(context.Context) NetworkFirewallPolicyAssociationMapOutput
}

type NetworkFirewallPolicyAssociationMap map[string]NetworkFirewallPolicyAssociationInput

func (NetworkFirewallPolicyAssociationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkFirewallPolicyAssociation)(nil)).Elem()
}

func (i NetworkFirewallPolicyAssociationMap) ToNetworkFirewallPolicyAssociationMapOutput() NetworkFirewallPolicyAssociationMapOutput {
	return i.ToNetworkFirewallPolicyAssociationMapOutputWithContext(context.Background())
}

func (i NetworkFirewallPolicyAssociationMap) ToNetworkFirewallPolicyAssociationMapOutputWithContext(ctx context.Context) NetworkFirewallPolicyAssociationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkFirewallPolicyAssociationMapOutput)
}

type NetworkFirewallPolicyAssociationOutput struct{ *pulumi.OutputState }

func (NetworkFirewallPolicyAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkFirewallPolicyAssociation)(nil)).Elem()
}

func (o NetworkFirewallPolicyAssociationOutput) ToNetworkFirewallPolicyAssociationOutput() NetworkFirewallPolicyAssociationOutput {
	return o
}

func (o NetworkFirewallPolicyAssociationOutput) ToNetworkFirewallPolicyAssociationOutputWithContext(ctx context.Context) NetworkFirewallPolicyAssociationOutput {
	return o
}

// The target that the firewall policy is attached to.
func (o NetworkFirewallPolicyAssociationOutput) AttachmentTarget() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkFirewallPolicyAssociation) pulumi.StringOutput { return v.AttachmentTarget }).(pulumi.StringOutput)
}

// The firewall policy ID of the association.
func (o NetworkFirewallPolicyAssociationOutput) FirewallPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkFirewallPolicyAssociation) pulumi.StringOutput { return v.FirewallPolicy }).(pulumi.StringOutput)
}

// The name for an association.
//
// ***
func (o NetworkFirewallPolicyAssociationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkFirewallPolicyAssociation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project for the resource
func (o NetworkFirewallPolicyAssociationOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkFirewallPolicyAssociation) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The short name of the firewall policy of the association.
func (o NetworkFirewallPolicyAssociationOutput) ShortName() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkFirewallPolicyAssociation) pulumi.StringOutput { return v.ShortName }).(pulumi.StringOutput)
}

type NetworkFirewallPolicyAssociationArrayOutput struct{ *pulumi.OutputState }

func (NetworkFirewallPolicyAssociationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkFirewallPolicyAssociation)(nil)).Elem()
}

func (o NetworkFirewallPolicyAssociationArrayOutput) ToNetworkFirewallPolicyAssociationArrayOutput() NetworkFirewallPolicyAssociationArrayOutput {
	return o
}

func (o NetworkFirewallPolicyAssociationArrayOutput) ToNetworkFirewallPolicyAssociationArrayOutputWithContext(ctx context.Context) NetworkFirewallPolicyAssociationArrayOutput {
	return o
}

func (o NetworkFirewallPolicyAssociationArrayOutput) Index(i pulumi.IntInput) NetworkFirewallPolicyAssociationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NetworkFirewallPolicyAssociation {
		return vs[0].([]*NetworkFirewallPolicyAssociation)[vs[1].(int)]
	}).(NetworkFirewallPolicyAssociationOutput)
}

type NetworkFirewallPolicyAssociationMapOutput struct{ *pulumi.OutputState }

func (NetworkFirewallPolicyAssociationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkFirewallPolicyAssociation)(nil)).Elem()
}

func (o NetworkFirewallPolicyAssociationMapOutput) ToNetworkFirewallPolicyAssociationMapOutput() NetworkFirewallPolicyAssociationMapOutput {
	return o
}

func (o NetworkFirewallPolicyAssociationMapOutput) ToNetworkFirewallPolicyAssociationMapOutputWithContext(ctx context.Context) NetworkFirewallPolicyAssociationMapOutput {
	return o
}

func (o NetworkFirewallPolicyAssociationMapOutput) MapIndex(k pulumi.StringInput) NetworkFirewallPolicyAssociationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NetworkFirewallPolicyAssociation {
		return vs[0].(map[string]*NetworkFirewallPolicyAssociation)[vs[1].(string)]
	}).(NetworkFirewallPolicyAssociationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkFirewallPolicyAssociationInput)(nil)).Elem(), &NetworkFirewallPolicyAssociation{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkFirewallPolicyAssociationArrayInput)(nil)).Elem(), NetworkFirewallPolicyAssociationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkFirewallPolicyAssociationMapInput)(nil)).Elem(), NetworkFirewallPolicyAssociationMap{})
	pulumi.RegisterOutputType(NetworkFirewallPolicyAssociationOutput{})
	pulumi.RegisterOutputType(NetworkFirewallPolicyAssociationArrayOutput{})
	pulumi.RegisterOutputType(NetworkFirewallPolicyAssociationMapOutput{})
}
