// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
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
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewFirewallPolicy(ctx, "_default", &compute.FirewallPolicyArgs{
// 			Description: pulumi.String("Example Resource"),
// 			Parent:      pulumi.String("organizations/12345"),
// 			ShortName:   pulumi.String("my-policy"),
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
// FirewallPolicy can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:compute/firewallPolicy:FirewallPolicy default locations/global/firewallPolicies/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/firewallPolicy:FirewallPolicy default {{name}}
// ```
type FirewallPolicy struct {
	pulumi.CustomResourceState

	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description      pulumi.StringPtrOutput `pulumi:"description"`
	Fingerprint      pulumi.StringOutput    `pulumi:"fingerprint"`
	FirewallPolicyId pulumi.StringOutput    `pulumi:"firewallPolicyId"`
	Name             pulumi.StringOutput    `pulumi:"name"`
	// The parent of the firewall policy.
	Parent         pulumi.StringOutput `pulumi:"parent"`
	RuleTupleCount pulumi.IntOutput    `pulumi:"ruleTupleCount"`
	SelfLink       pulumi.StringOutput `pulumi:"selfLink"`
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
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description      *string `pulumi:"description"`
	Fingerprint      *string `pulumi:"fingerprint"`
	FirewallPolicyId *string `pulumi:"firewallPolicyId"`
	Name             *string `pulumi:"name"`
	// The parent of the firewall policy.
	Parent         *string `pulumi:"parent"`
	RuleTupleCount *int    `pulumi:"ruleTupleCount"`
	SelfLink       *string `pulumi:"selfLink"`
	SelfLinkWithId *string `pulumi:"selfLinkWithId"`
	// User-provided name of the Organization firewall policy. The name should be unique in the organization in which the firewall policy is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression a-z? which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	ShortName *string `pulumi:"shortName"`
}

type FirewallPolicyState struct {
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description      pulumi.StringPtrInput
	Fingerprint      pulumi.StringPtrInput
	FirewallPolicyId pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	// The parent of the firewall policy.
	Parent         pulumi.StringPtrInput
	RuleTupleCount pulumi.IntPtrInput
	SelfLink       pulumi.StringPtrInput
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
	return reflect.TypeOf((*FirewallPolicy)(nil))
}

func (i *FirewallPolicy) ToFirewallPolicyOutput() FirewallPolicyOutput {
	return i.ToFirewallPolicyOutputWithContext(context.Background())
}

func (i *FirewallPolicy) ToFirewallPolicyOutputWithContext(ctx context.Context) FirewallPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyOutput)
}

func (i *FirewallPolicy) ToFirewallPolicyPtrOutput() FirewallPolicyPtrOutput {
	return i.ToFirewallPolicyPtrOutputWithContext(context.Background())
}

func (i *FirewallPolicy) ToFirewallPolicyPtrOutputWithContext(ctx context.Context) FirewallPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyPtrOutput)
}

type FirewallPolicyPtrInput interface {
	pulumi.Input

	ToFirewallPolicyPtrOutput() FirewallPolicyPtrOutput
	ToFirewallPolicyPtrOutputWithContext(ctx context.Context) FirewallPolicyPtrOutput
}

type firewallPolicyPtrType FirewallPolicyArgs

func (*firewallPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPolicy)(nil))
}

func (i *firewallPolicyPtrType) ToFirewallPolicyPtrOutput() FirewallPolicyPtrOutput {
	return i.ToFirewallPolicyPtrOutputWithContext(context.Background())
}

func (i *firewallPolicyPtrType) ToFirewallPolicyPtrOutputWithContext(ctx context.Context) FirewallPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyPtrOutput)
}

// FirewallPolicyArrayInput is an input type that accepts FirewallPolicyArray and FirewallPolicyArrayOutput values.
// You can construct a concrete instance of `FirewallPolicyArrayInput` via:
//
//          FirewallPolicyArray{ FirewallPolicyArgs{...} }
type FirewallPolicyArrayInput interface {
	pulumi.Input

	ToFirewallPolicyArrayOutput() FirewallPolicyArrayOutput
	ToFirewallPolicyArrayOutputWithContext(context.Context) FirewallPolicyArrayOutput
}

type FirewallPolicyArray []FirewallPolicyInput

func (FirewallPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*FirewallPolicy)(nil))
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
//          FirewallPolicyMap{ "key": FirewallPolicyArgs{...} }
type FirewallPolicyMapInput interface {
	pulumi.Input

	ToFirewallPolicyMapOutput() FirewallPolicyMapOutput
	ToFirewallPolicyMapOutputWithContext(context.Context) FirewallPolicyMapOutput
}

type FirewallPolicyMap map[string]FirewallPolicyInput

func (FirewallPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*FirewallPolicy)(nil))
}

func (i FirewallPolicyMap) ToFirewallPolicyMapOutput() FirewallPolicyMapOutput {
	return i.ToFirewallPolicyMapOutputWithContext(context.Background())
}

func (i FirewallPolicyMap) ToFirewallPolicyMapOutputWithContext(ctx context.Context) FirewallPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallPolicyMapOutput)
}

type FirewallPolicyOutput struct {
	*pulumi.OutputState
}

func (FirewallPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FirewallPolicy)(nil))
}

func (o FirewallPolicyOutput) ToFirewallPolicyOutput() FirewallPolicyOutput {
	return o
}

func (o FirewallPolicyOutput) ToFirewallPolicyOutputWithContext(ctx context.Context) FirewallPolicyOutput {
	return o
}

func (o FirewallPolicyOutput) ToFirewallPolicyPtrOutput() FirewallPolicyPtrOutput {
	return o.ToFirewallPolicyPtrOutputWithContext(context.Background())
}

func (o FirewallPolicyOutput) ToFirewallPolicyPtrOutputWithContext(ctx context.Context) FirewallPolicyPtrOutput {
	return o.ApplyT(func(v FirewallPolicy) *FirewallPolicy {
		return &v
	}).(FirewallPolicyPtrOutput)
}

type FirewallPolicyPtrOutput struct {
	*pulumi.OutputState
}

func (FirewallPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallPolicy)(nil))
}

func (o FirewallPolicyPtrOutput) ToFirewallPolicyPtrOutput() FirewallPolicyPtrOutput {
	return o
}

func (o FirewallPolicyPtrOutput) ToFirewallPolicyPtrOutputWithContext(ctx context.Context) FirewallPolicyPtrOutput {
	return o
}

type FirewallPolicyArrayOutput struct{ *pulumi.OutputState }

func (FirewallPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FirewallPolicy)(nil))
}

func (o FirewallPolicyArrayOutput) ToFirewallPolicyArrayOutput() FirewallPolicyArrayOutput {
	return o
}

func (o FirewallPolicyArrayOutput) ToFirewallPolicyArrayOutputWithContext(ctx context.Context) FirewallPolicyArrayOutput {
	return o
}

func (o FirewallPolicyArrayOutput) Index(i pulumi.IntInput) FirewallPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FirewallPolicy {
		return vs[0].([]FirewallPolicy)[vs[1].(int)]
	}).(FirewallPolicyOutput)
}

type FirewallPolicyMapOutput struct{ *pulumi.OutputState }

func (FirewallPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FirewallPolicy)(nil))
}

func (o FirewallPolicyMapOutput) ToFirewallPolicyMapOutput() FirewallPolicyMapOutput {
	return o
}

func (o FirewallPolicyMapOutput) ToFirewallPolicyMapOutputWithContext(ctx context.Context) FirewallPolicyMapOutput {
	return o
}

func (o FirewallPolicyMapOutput) MapIndex(k pulumi.StringInput) FirewallPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FirewallPolicy {
		return vs[0].(map[string]FirewallPolicy)[vs[1].(string)]
	}).(FirewallPolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(FirewallPolicyOutput{})
	pulumi.RegisterOutputType(FirewallPolicyPtrOutput{})
	pulumi.RegisterOutputType(FirewallPolicyArrayOutput{})
	pulumi.RegisterOutputType(FirewallPolicyMapOutput{})
}
