// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows associating hierarchical firewall policies with the target where they are applied. This allows creating policies and rules in a different location than they are applied.
//
// For more information on applying hierarchical firewall policies see the [official documentation](https://cloud.google.com/vpc/docs/firewall-policies#managing_hierarchical_firewall_policy_resources)
//
// ## Example Usage
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
//			defaultFirewallPolicy, err := compute.NewFirewallPolicy(ctx, "defaultFirewallPolicy", &compute.FirewallPolicyArgs{
//				Parent:      pulumi.String("organizations/12345"),
//				ShortName:   pulumi.String("my-policy"),
//				Description: pulumi.String("Example Resource"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewFirewallPolicyAssociation(ctx, "defaultFirewallPolicyAssociation", &compute.FirewallPolicyAssociationArgs{
//				FirewallPolicy:   defaultFirewallPolicy.ID(),
//				AttachmentTarget: pulumi.Any(google_folder.Folder.Name),
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
// # FirewallPolicyAssociation can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:compute/firewallPolicyAssociation:FirewallPolicyAssociation default locations/global/firewallPolicies/{{firewall_policy}}/associations/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/firewallPolicyAssociation:FirewallPolicyAssociation default {{firewall_policy}}/{{name}}
//
// ```
type FirewallPolicyAssociation struct {
	pulumi.CustomResourceState

	// The target that the firewall policy is attached to.
	AttachmentTarget pulumi.StringOutput `pulumi:"attachmentTarget"`
	// The firewall policy ID of the association.
	FirewallPolicy pulumi.StringOutput `pulumi:"firewallPolicy"`
	// The name for an association.
	Name pulumi.StringOutput `pulumi:"name"`
	// The short name of the firewall policy of the association.
	ShortName pulumi.StringOutput `pulumi:"shortName"`
}

// NewFirewallPolicyAssociation registers a new resource with the given unique name, arguments, and options.
func NewFirewallPolicyAssociation(ctx *pulumi.Context,
	name string, args *FirewallPolicyAssociationArgs, opts ...pulumi.ResourceOption) (*FirewallPolicyAssociation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AttachmentTarget == nil {
		return nil, errors.New("invalid value for required argument 'AttachmentTarget'")
	}
	if args.FirewallPolicy == nil {
		return nil, errors.New("invalid value for required argument 'FirewallPolicy'")
	}
	var resource FirewallPolicyAssociation
	err := ctx.RegisterResource("gcp:compute/firewallPolicyAssociation:FirewallPolicyAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallPolicyAssociation gets an existing FirewallPolicyAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallPolicyAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallPolicyAssociationState, opts ...pulumi.ResourceOption) (*FirewallPolicyAssociation, error) {
	var resource FirewallPolicyAssociation
	err := ctx.ReadResource("gcp:compute/firewallPolicyAssociation:FirewallPolicyAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallPolicyAssociation resources.
type firewallPolicyAssociationState struct {
	// The target that the firewall policy is attached to.
	AttachmentTarget *string `pulumi:"attachmentTarget"`
	// The firewall policy ID of the association.
	FirewallPolicy *string `pulumi:"firewallPolicy"`
	// The name for an association.
	Name *string `pulumi:"name"`
	// The short name of the firewall policy of the association.
	ShortName *string `pulumi:"shortName"`
}

type FirewallPolicyAssociationState struct {
	// The target that the firewall policy is attached to.
	AttachmentTarget pulumi.StringPtrInput
	// The firewall policy ID of the association.
	FirewallPolicy pulumi.StringPtrInput
	// The name for an association.
	Name pulumi.StringPtrInput
	// The short name of the firewall policy of the association.
	ShortName pulumi.StringPtrInput
}

func (FirewallPolicyAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicyAssociationState)(nil)).Elem()
}

type firewallPolicyAssociationArgs struct {
	// The target that the firewall policy is attached to.
	AttachmentTarget string `pulumi:"attachmentTarget"`
	// The firewall policy ID of the association.
	FirewallPolicy string `pulumi:"firewallPolicy"`
	// The name for an association.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a FirewallPolicyAssociation resource.
type FirewallPolicyAssociationArgs struct {
	// The target that the firewall policy is attached to.
	AttachmentTarget pulumi.StringInput
	// The firewall policy ID of the association.
	FirewallPolicy pulumi.StringInput
	// The name for an association.
	Name pulumi.StringPtrInput
}

func (FirewallPolicyAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicyAssociationArgs)(nil)).Elem()
}

type FirewallPolicyAssociationInput interface {
	pulumi.Input

	ToFirewallPolicyAssociationOutput() FirewallPolicyAssociationOutput
	ToFirewallPolicyAssociationOutputWithContext(ctx context.Context) FirewallPolicyAssociationOutput
}

func (*FirewallPolicyAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPolicyAssociation)(nil)).Elem()
}

func (i *FirewallPolicyAssociation) ToFirewallPolicyAssociationOutput() FirewallPolicyAssociationOutput {
	return i.ToFirewallPolicyAssociationOutputWithContext(context.Background())
}

func (i *FirewallPolicyAssociation) ToFirewallPolicyAssociationOutputWithContext(ctx context.Context) FirewallPolicyAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyAssociationOutput)
}

// FirewallPolicyAssociationArrayInput is an input type that accepts FirewallPolicyAssociationArray and FirewallPolicyAssociationArrayOutput values.
// You can construct a concrete instance of `FirewallPolicyAssociationArrayInput` via:
//
//	FirewallPolicyAssociationArray{ FirewallPolicyAssociationArgs{...} }
type FirewallPolicyAssociationArrayInput interface {
	pulumi.Input

	ToFirewallPolicyAssociationArrayOutput() FirewallPolicyAssociationArrayOutput
	ToFirewallPolicyAssociationArrayOutputWithContext(context.Context) FirewallPolicyAssociationArrayOutput
}

type FirewallPolicyAssociationArray []FirewallPolicyAssociationInput

func (FirewallPolicyAssociationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallPolicyAssociation)(nil)).Elem()
}

func (i FirewallPolicyAssociationArray) ToFirewallPolicyAssociationArrayOutput() FirewallPolicyAssociationArrayOutput {
	return i.ToFirewallPolicyAssociationArrayOutputWithContext(context.Background())
}

func (i FirewallPolicyAssociationArray) ToFirewallPolicyAssociationArrayOutputWithContext(ctx context.Context) FirewallPolicyAssociationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyAssociationArrayOutput)
}

// FirewallPolicyAssociationMapInput is an input type that accepts FirewallPolicyAssociationMap and FirewallPolicyAssociationMapOutput values.
// You can construct a concrete instance of `FirewallPolicyAssociationMapInput` via:
//
//	FirewallPolicyAssociationMap{ "key": FirewallPolicyAssociationArgs{...} }
type FirewallPolicyAssociationMapInput interface {
	pulumi.Input

	ToFirewallPolicyAssociationMapOutput() FirewallPolicyAssociationMapOutput
	ToFirewallPolicyAssociationMapOutputWithContext(context.Context) FirewallPolicyAssociationMapOutput
}

type FirewallPolicyAssociationMap map[string]FirewallPolicyAssociationInput

func (FirewallPolicyAssociationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallPolicyAssociation)(nil)).Elem()
}

func (i FirewallPolicyAssociationMap) ToFirewallPolicyAssociationMapOutput() FirewallPolicyAssociationMapOutput {
	return i.ToFirewallPolicyAssociationMapOutputWithContext(context.Background())
}

func (i FirewallPolicyAssociationMap) ToFirewallPolicyAssociationMapOutputWithContext(ctx context.Context) FirewallPolicyAssociationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyAssociationMapOutput)
}

type FirewallPolicyAssociationOutput struct{ *pulumi.OutputState }

func (FirewallPolicyAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPolicyAssociation)(nil)).Elem()
}

func (o FirewallPolicyAssociationOutput) ToFirewallPolicyAssociationOutput() FirewallPolicyAssociationOutput {
	return o
}

func (o FirewallPolicyAssociationOutput) ToFirewallPolicyAssociationOutputWithContext(ctx context.Context) FirewallPolicyAssociationOutput {
	return o
}

// The target that the firewall policy is attached to.
func (o FirewallPolicyAssociationOutput) AttachmentTarget() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicyAssociation) pulumi.StringOutput { return v.AttachmentTarget }).(pulumi.StringOutput)
}

// The firewall policy ID of the association.
func (o FirewallPolicyAssociationOutput) FirewallPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicyAssociation) pulumi.StringOutput { return v.FirewallPolicy }).(pulumi.StringOutput)
}

// The name for an association.
func (o FirewallPolicyAssociationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicyAssociation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The short name of the firewall policy of the association.
func (o FirewallPolicyAssociationOutput) ShortName() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicyAssociation) pulumi.StringOutput { return v.ShortName }).(pulumi.StringOutput)
}

type FirewallPolicyAssociationArrayOutput struct{ *pulumi.OutputState }

func (FirewallPolicyAssociationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallPolicyAssociation)(nil)).Elem()
}

func (o FirewallPolicyAssociationArrayOutput) ToFirewallPolicyAssociationArrayOutput() FirewallPolicyAssociationArrayOutput {
	return o
}

func (o FirewallPolicyAssociationArrayOutput) ToFirewallPolicyAssociationArrayOutputWithContext(ctx context.Context) FirewallPolicyAssociationArrayOutput {
	return o
}

func (o FirewallPolicyAssociationArrayOutput) Index(i pulumi.IntInput) FirewallPolicyAssociationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallPolicyAssociation {
		return vs[0].([]*FirewallPolicyAssociation)[vs[1].(int)]
	}).(FirewallPolicyAssociationOutput)
}

type FirewallPolicyAssociationMapOutput struct{ *pulumi.OutputState }

func (FirewallPolicyAssociationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallPolicyAssociation)(nil)).Elem()
}

func (o FirewallPolicyAssociationMapOutput) ToFirewallPolicyAssociationMapOutput() FirewallPolicyAssociationMapOutput {
	return o
}

func (o FirewallPolicyAssociationMapOutput) ToFirewallPolicyAssociationMapOutputWithContext(ctx context.Context) FirewallPolicyAssociationMapOutput {
	return o
}

func (o FirewallPolicyAssociationMapOutput) MapIndex(k pulumi.StringInput) FirewallPolicyAssociationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallPolicyAssociation {
		return vs[0].(map[string]*FirewallPolicyAssociation)[vs[1].(string)]
	}).(FirewallPolicyAssociationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallPolicyAssociationInput)(nil)).Elem(), &FirewallPolicyAssociation{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallPolicyAssociationArrayInput)(nil)).Elem(), FirewallPolicyAssociationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallPolicyAssociationMapInput)(nil)).Elem(), FirewallPolicyAssociationMap{})
	pulumi.RegisterOutputType(FirewallPolicyAssociationOutput{})
	pulumi.RegisterOutputType(FirewallPolicyAssociationArrayOutput{})
	pulumi.RegisterOutputType(FirewallPolicyAssociationMapOutput{})
}
