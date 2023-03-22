// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Hierarchical firewall policy rules let you create and enforce a consistent firewall policy across your organization. Rules can explicitly allow or deny connections or delegate evaluation to lower level policies. Policies can be created within organizations or folders.
//
// This resource should be generally be used with `compute.FirewallPolicyAssociation` and `compute.FirewallPolicyRule`
//
// For more information see the [official documentation](https://cloud.google.com/vpc/docs/firewall-policies)
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
//			_, err := compute.NewFirewallPolicy(ctx, "default", &compute.FirewallPolicyArgs{
//				Description: pulumi.String("Example Resource"),
//				Parent:      pulumi.String("organizations/12345"),
//				ShortName:   pulumi.String("my-policy"),
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
// # FirewallPolicy can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:compute/firewallPolicy:FirewallPolicy default locations/global/firewallPolicies/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/firewallPolicy:FirewallPolicy default {{name}}
//
// ```
type FirewallPolicy struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Fingerprint of the resource. This field is used internally during updates of this resource.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// The unique identifier for the resource. This identifier is defined by the server.
	FirewallPolicyId pulumi.StringOutput `pulumi:"firewallPolicyId"`
	// Name of the resource. It is a numeric ID allocated by GCP which uniquely identifies the Firewall Policy.
	Name pulumi.StringOutput `pulumi:"name"`
	// The parent of the firewall policy.
	Parent pulumi.StringOutput `pulumi:"parent"`
	// Total count of all firewall policy rule tuples. A firewall policy can not exceed a set number of tuples.
	RuleTupleCount pulumi.IntOutput `pulumi:"ruleTupleCount"`
	// Server-defined URL for the resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithId pulumi.StringOutput `pulumi:"selfLinkWithId"`
	// User-provided name of the Organization firewall policy. The name should be unique in the organization in which the firewall policy is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression a-z? which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	ShortName pulumi.StringOutput `pulumi:"shortName"`
}

// NewFirewallPolicy registers a new resource with the given unique name, arguments, and options.
func NewFirewallPolicy(ctx *pulumi.Context,
	name string, args *FirewallPolicyArgs, opts ...pulumi.ResourceOption) (*FirewallPolicy, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Parent == nil {
		return nil, errors.New("invalid value for required argument 'Parent'")
	}
	if args.ShortName == nil {
		return nil, errors.New("invalid value for required argument 'ShortName'")
	}
	var resource FirewallPolicy
	err := ctx.RegisterResource("gcp:compute/firewallPolicy:FirewallPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallPolicy gets an existing FirewallPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallPolicyState, opts ...pulumi.ResourceOption) (*FirewallPolicy, error) {
	var resource FirewallPolicy
	err := ctx.ReadResource("gcp:compute/firewallPolicy:FirewallPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallPolicy resources.
type firewallPolicyState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Fingerprint of the resource. This field is used internally during updates of this resource.
	Fingerprint *string `pulumi:"fingerprint"`
	// The unique identifier for the resource. This identifier is defined by the server.
	FirewallPolicyId *string `pulumi:"firewallPolicyId"`
	// Name of the resource. It is a numeric ID allocated by GCP which uniquely identifies the Firewall Policy.
	Name *string `pulumi:"name"`
	// The parent of the firewall policy.
	Parent *string `pulumi:"parent"`
	// Total count of all firewall policy rule tuples. A firewall policy can not exceed a set number of tuples.
	RuleTupleCount *int `pulumi:"ruleTupleCount"`
	// Server-defined URL for the resource.
	SelfLink *string `pulumi:"selfLink"`
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithId *string `pulumi:"selfLinkWithId"`
	// User-provided name of the Organization firewall policy. The name should be unique in the organization in which the firewall policy is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression a-z? which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	ShortName *string `pulumi:"shortName"`
}

type FirewallPolicyState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Fingerprint of the resource. This field is used internally during updates of this resource.
	Fingerprint pulumi.StringPtrInput
	// The unique identifier for the resource. This identifier is defined by the server.
	FirewallPolicyId pulumi.StringPtrInput
	// Name of the resource. It is a numeric ID allocated by GCP which uniquely identifies the Firewall Policy.
	Name pulumi.StringPtrInput
	// The parent of the firewall policy.
	Parent pulumi.StringPtrInput
	// Total count of all firewall policy rule tuples. A firewall policy can not exceed a set number of tuples.
	RuleTupleCount pulumi.IntPtrInput
	// Server-defined URL for the resource.
	SelfLink pulumi.StringPtrInput
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithId pulumi.StringPtrInput
	// User-provided name of the Organization firewall policy. The name should be unique in the organization in which the firewall policy is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression a-z? which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	ShortName pulumi.StringPtrInput
}

func (FirewallPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicyState)(nil)).Elem()
}

type firewallPolicyArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// The parent of the firewall policy.
	Parent string `pulumi:"parent"`
	// User-provided name of the Organization firewall policy. The name should be unique in the organization in which the firewall policy is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression a-z? which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	ShortName string `pulumi:"shortName"`
}

// The set of arguments for constructing a FirewallPolicy resource.
type FirewallPolicyArgs struct {
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// The parent of the firewall policy.
	Parent pulumi.StringInput
	// User-provided name of the Organization firewall policy. The name should be unique in the organization in which the firewall policy is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression a-z? which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	ShortName pulumi.StringInput
}

func (FirewallPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallPolicyArgs)(nil)).Elem()
}

type FirewallPolicyInput interface {
	pulumi.Input

	ToFirewallPolicyOutput() FirewallPolicyOutput
	ToFirewallPolicyOutputWithContext(ctx context.Context) FirewallPolicyOutput
}

func (*FirewallPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPolicy)(nil)).Elem()
}

func (i *FirewallPolicy) ToFirewallPolicyOutput() FirewallPolicyOutput {
	return i.ToFirewallPolicyOutputWithContext(context.Background())
}

func (i *FirewallPolicy) ToFirewallPolicyOutputWithContext(ctx context.Context) FirewallPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyOutput)
}

// FirewallPolicyArrayInput is an input type that accepts FirewallPolicyArray and FirewallPolicyArrayOutput values.
// You can construct a concrete instance of `FirewallPolicyArrayInput` via:
//
//	FirewallPolicyArray{ FirewallPolicyArgs{...} }
type FirewallPolicyArrayInput interface {
	pulumi.Input

	ToFirewallPolicyArrayOutput() FirewallPolicyArrayOutput
	ToFirewallPolicyArrayOutputWithContext(context.Context) FirewallPolicyArrayOutput
}

type FirewallPolicyArray []FirewallPolicyInput

func (FirewallPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallPolicy)(nil)).Elem()
}

func (i FirewallPolicyArray) ToFirewallPolicyArrayOutput() FirewallPolicyArrayOutput {
	return i.ToFirewallPolicyArrayOutputWithContext(context.Background())
}

func (i FirewallPolicyArray) ToFirewallPolicyArrayOutputWithContext(ctx context.Context) FirewallPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyArrayOutput)
}

// FirewallPolicyMapInput is an input type that accepts FirewallPolicyMap and FirewallPolicyMapOutput values.
// You can construct a concrete instance of `FirewallPolicyMapInput` via:
//
//	FirewallPolicyMap{ "key": FirewallPolicyArgs{...} }
type FirewallPolicyMapInput interface {
	pulumi.Input

	ToFirewallPolicyMapOutput() FirewallPolicyMapOutput
	ToFirewallPolicyMapOutputWithContext(context.Context) FirewallPolicyMapOutput
}

type FirewallPolicyMap map[string]FirewallPolicyInput

func (FirewallPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallPolicy)(nil)).Elem()
}

func (i FirewallPolicyMap) ToFirewallPolicyMapOutput() FirewallPolicyMapOutput {
	return i.ToFirewallPolicyMapOutputWithContext(context.Background())
}

func (i FirewallPolicyMap) ToFirewallPolicyMapOutputWithContext(ctx context.Context) FirewallPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyMapOutput)
}

type FirewallPolicyOutput struct{ *pulumi.OutputState }

func (FirewallPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPolicy)(nil)).Elem()
}

func (o FirewallPolicyOutput) ToFirewallPolicyOutput() FirewallPolicyOutput {
	return o
}

func (o FirewallPolicyOutput) ToFirewallPolicyOutputWithContext(ctx context.Context) FirewallPolicyOutput {
	return o
}

// Creation timestamp in RFC3339 text format.
func (o FirewallPolicyOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicy) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource. Provide this property when you create the resource.
func (o FirewallPolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FirewallPolicy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Fingerprint of the resource. This field is used internally during updates of this resource.
func (o FirewallPolicyOutput) Fingerprint() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicy) pulumi.StringOutput { return v.Fingerprint }).(pulumi.StringOutput)
}

// The unique identifier for the resource. This identifier is defined by the server.
func (o FirewallPolicyOutput) FirewallPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicy) pulumi.StringOutput { return v.FirewallPolicyId }).(pulumi.StringOutput)
}

// Name of the resource. It is a numeric ID allocated by GCP which uniquely identifies the Firewall Policy.
func (o FirewallPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The parent of the firewall policy.
func (o FirewallPolicyOutput) Parent() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicy) pulumi.StringOutput { return v.Parent }).(pulumi.StringOutput)
}

// Total count of all firewall policy rule tuples. A firewall policy can not exceed a set number of tuples.
func (o FirewallPolicyOutput) RuleTupleCount() pulumi.IntOutput {
	return o.ApplyT(func(v *FirewallPolicy) pulumi.IntOutput { return v.RuleTupleCount }).(pulumi.IntOutput)
}

// Server-defined URL for the resource.
func (o FirewallPolicyOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicy) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// Server-defined URL for this resource with the resource id.
func (o FirewallPolicyOutput) SelfLinkWithId() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicy) pulumi.StringOutput { return v.SelfLinkWithId }).(pulumi.StringOutput)
}

// User-provided name of the Organization firewall policy. The name should be unique in the organization in which the firewall policy is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression a-z? which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
func (o FirewallPolicyOutput) ShortName() pulumi.StringOutput {
	return o.ApplyT(func(v *FirewallPolicy) pulumi.StringOutput { return v.ShortName }).(pulumi.StringOutput)
}

type FirewallPolicyArrayOutput struct{ *pulumi.OutputState }

func (FirewallPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallPolicy)(nil)).Elem()
}

func (o FirewallPolicyArrayOutput) ToFirewallPolicyArrayOutput() FirewallPolicyArrayOutput {
	return o
}

func (o FirewallPolicyArrayOutput) ToFirewallPolicyArrayOutputWithContext(ctx context.Context) FirewallPolicyArrayOutput {
	return o
}

func (o FirewallPolicyArrayOutput) Index(i pulumi.IntInput) FirewallPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallPolicy {
		return vs[0].([]*FirewallPolicy)[vs[1].(int)]
	}).(FirewallPolicyOutput)
}

type FirewallPolicyMapOutput struct{ *pulumi.OutputState }

func (FirewallPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallPolicy)(nil)).Elem()
}

func (o FirewallPolicyMapOutput) ToFirewallPolicyMapOutput() FirewallPolicyMapOutput {
	return o
}

func (o FirewallPolicyMapOutput) ToFirewallPolicyMapOutputWithContext(ctx context.Context) FirewallPolicyMapOutput {
	return o
}

func (o FirewallPolicyMapOutput) MapIndex(k pulumi.StringInput) FirewallPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallPolicy {
		return vs[0].(map[string]*FirewallPolicy)[vs[1].(string)]
	}).(FirewallPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallPolicyInput)(nil)).Elem(), &FirewallPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallPolicyArrayInput)(nil)).Elem(), FirewallPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallPolicyMapInput)(nil)).Elem(), FirewallPolicyMap{})
	pulumi.RegisterOutputType(FirewallPolicyOutput{})
	pulumi.RegisterOutputType(FirewallPolicyArrayOutput{})
	pulumi.RegisterOutputType(FirewallPolicyMapOutput{})
}
