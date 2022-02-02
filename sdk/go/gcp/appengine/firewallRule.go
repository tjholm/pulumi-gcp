// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package appengine

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A single firewall rule that is evaluated against incoming traffic
// and provides an action to take on matched requests.
//
// To get more information about FirewallRule, see:
//
// * [API documentation](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.firewall.ingressRules)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/appengine/docs/standard/python/creating-firewalls#creating_firewall_rules)
//
// ## Example Usage
// ### App Engine Firewall Rule Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/appengine"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		myProject, err := organizations.NewProject(ctx, "myProject", &organizations.ProjectArgs{
// 			ProjectId: pulumi.String("ae-project"),
// 			OrgId:     pulumi.String("123456789"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		app, err := appengine.NewApplication(ctx, "app", &appengine.ApplicationArgs{
// 			Project:    myProject.ProjectId,
// 			LocationId: pulumi.String("us-central"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = appengine.NewFirewallRule(ctx, "rule", &appengine.FirewallRuleArgs{
// 			Project:     app.Project,
// 			Priority:    pulumi.Int(1000),
// 			Action:      pulumi.String("ALLOW"),
// 			SourceRange: pulumi.String("*"),
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
// FirewallRule can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:appengine/firewallRule:FirewallRule default apps/{{project}}/firewall/ingressRules/{{priority}}
// ```
//
// ```sh
//  $ pulumi import gcp:appengine/firewallRule:FirewallRule default {{project}}/{{priority}}
// ```
//
// ```sh
//  $ pulumi import gcp:appengine/firewallRule:FirewallRule default {{priority}}
// ```
type FirewallRule struct {
	pulumi.CustomResourceState

	// The action to take if this rule matches.
	// Possible values are `UNSPECIFIED_ACTION`, `ALLOW`, and `DENY`.
	Action pulumi.StringOutput `pulumi:"action"`
	// An optional string description of this rule.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// A positive integer that defines the order of rule evaluation.
	// Rules with the lowest priority are evaluated first.
	// A default rule at priority Int32.MaxValue matches all IPv4 and
	// IPv6 traffic when no previous rule matches. Only the action of
	// this rule can be modified by the user.
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// IP address or range, defined using CIDR notation, of requests that this rule applies to.
	SourceRange pulumi.StringOutput `pulumi:"sourceRange"`
}

// NewFirewallRule registers a new resource with the given unique name, arguments, and options.
func NewFirewallRule(ctx *pulumi.Context,
	name string, args *FirewallRuleArgs, opts ...pulumi.ResourceOption) (*FirewallRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.SourceRange == nil {
		return nil, errors.New("invalid value for required argument 'SourceRange'")
	}
	var resource FirewallRule
	err := ctx.RegisterResource("gcp:appengine/firewallRule:FirewallRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFirewallRule gets an existing FirewallRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFirewallRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FirewallRuleState, opts ...pulumi.ResourceOption) (*FirewallRule, error) {
	var resource FirewallRule
	err := ctx.ReadResource("gcp:appengine/firewallRule:FirewallRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FirewallRule resources.
type firewallRuleState struct {
	// The action to take if this rule matches.
	// Possible values are `UNSPECIFIED_ACTION`, `ALLOW`, and `DENY`.
	Action *string `pulumi:"action"`
	// An optional string description of this rule.
	Description *string `pulumi:"description"`
	// A positive integer that defines the order of rule evaluation.
	// Rules with the lowest priority are evaluated first.
	// A default rule at priority Int32.MaxValue matches all IPv4 and
	// IPv6 traffic when no previous rule matches. Only the action of
	// this rule can be modified by the user.
	Priority *int `pulumi:"priority"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// IP address or range, defined using CIDR notation, of requests that this rule applies to.
	SourceRange *string `pulumi:"sourceRange"`
}

type FirewallRuleState struct {
	// The action to take if this rule matches.
	// Possible values are `UNSPECIFIED_ACTION`, `ALLOW`, and `DENY`.
	Action pulumi.StringPtrInput
	// An optional string description of this rule.
	Description pulumi.StringPtrInput
	// A positive integer that defines the order of rule evaluation.
	// Rules with the lowest priority are evaluated first.
	// A default rule at priority Int32.MaxValue matches all IPv4 and
	// IPv6 traffic when no previous rule matches. Only the action of
	// this rule can be modified by the user.
	Priority pulumi.IntPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// IP address or range, defined using CIDR notation, of requests that this rule applies to.
	SourceRange pulumi.StringPtrInput
}

func (FirewallRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallRuleState)(nil)).Elem()
}

type firewallRuleArgs struct {
	// The action to take if this rule matches.
	// Possible values are `UNSPECIFIED_ACTION`, `ALLOW`, and `DENY`.
	Action string `pulumi:"action"`
	// An optional string description of this rule.
	Description *string `pulumi:"description"`
	// A positive integer that defines the order of rule evaluation.
	// Rules with the lowest priority are evaluated first.
	// A default rule at priority Int32.MaxValue matches all IPv4 and
	// IPv6 traffic when no previous rule matches. Only the action of
	// this rule can be modified by the user.
	Priority *int `pulumi:"priority"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// IP address or range, defined using CIDR notation, of requests that this rule applies to.
	SourceRange string `pulumi:"sourceRange"`
}

// The set of arguments for constructing a FirewallRule resource.
type FirewallRuleArgs struct {
	// The action to take if this rule matches.
	// Possible values are `UNSPECIFIED_ACTION`, `ALLOW`, and `DENY`.
	Action pulumi.StringInput
	// An optional string description of this rule.
	Description pulumi.StringPtrInput
	// A positive integer that defines the order of rule evaluation.
	// Rules with the lowest priority are evaluated first.
	// A default rule at priority Int32.MaxValue matches all IPv4 and
	// IPv6 traffic when no previous rule matches. Only the action of
	// this rule can be modified by the user.
	Priority pulumi.IntPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// IP address or range, defined using CIDR notation, of requests that this rule applies to.
	SourceRange pulumi.StringInput
}

func (FirewallRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*firewallRuleArgs)(nil)).Elem()
}

type FirewallRuleInput interface {
	pulumi.Input

	ToFirewallRuleOutput() FirewallRuleOutput
	ToFirewallRuleOutputWithContext(ctx context.Context) FirewallRuleOutput
}

func (*FirewallRule) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallRule)(nil)).Elem()
}

func (i *FirewallRule) ToFirewallRuleOutput() FirewallRuleOutput {
	return i.ToFirewallRuleOutputWithContext(context.Background())
}

func (i *FirewallRule) ToFirewallRuleOutputWithContext(ctx context.Context) FirewallRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallRuleOutput)
}

// FirewallRuleArrayInput is an input type that accepts FirewallRuleArray and FirewallRuleArrayOutput values.
// You can construct a concrete instance of `FirewallRuleArrayInput` via:
//
//          FirewallRuleArray{ FirewallRuleArgs{...} }
type FirewallRuleArrayInput interface {
	pulumi.Input

	ToFirewallRuleArrayOutput() FirewallRuleArrayOutput
	ToFirewallRuleArrayOutputWithContext(context.Context) FirewallRuleArrayOutput
}

type FirewallRuleArray []FirewallRuleInput

func (FirewallRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallRule)(nil)).Elem()
}

func (i FirewallRuleArray) ToFirewallRuleArrayOutput() FirewallRuleArrayOutput {
	return i.ToFirewallRuleArrayOutputWithContext(context.Background())
}

func (i FirewallRuleArray) ToFirewallRuleArrayOutputWithContext(ctx context.Context) FirewallRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallRuleArrayOutput)
}

// FirewallRuleMapInput is an input type that accepts FirewallRuleMap and FirewallRuleMapOutput values.
// You can construct a concrete instance of `FirewallRuleMapInput` via:
//
//          FirewallRuleMap{ "key": FirewallRuleArgs{...} }
type FirewallRuleMapInput interface {
	pulumi.Input

	ToFirewallRuleMapOutput() FirewallRuleMapOutput
	ToFirewallRuleMapOutputWithContext(context.Context) FirewallRuleMapOutput
}

type FirewallRuleMap map[string]FirewallRuleInput

func (FirewallRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallRule)(nil)).Elem()
}

func (i FirewallRuleMap) ToFirewallRuleMapOutput() FirewallRuleMapOutput {
	return i.ToFirewallRuleMapOutputWithContext(context.Background())
}

func (i FirewallRuleMap) ToFirewallRuleMapOutputWithContext(ctx context.Context) FirewallRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FirewallRuleMapOutput)
}

type FirewallRuleOutput struct{ *pulumi.OutputState }

func (FirewallRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FirewallRule)(nil)).Elem()
}

func (o FirewallRuleOutput) ToFirewallRuleOutput() FirewallRuleOutput {
	return o
}

func (o FirewallRuleOutput) ToFirewallRuleOutputWithContext(ctx context.Context) FirewallRuleOutput {
	return o
}

type FirewallRuleArrayOutput struct{ *pulumi.OutputState }

func (FirewallRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FirewallRule)(nil)).Elem()
}

func (o FirewallRuleArrayOutput) ToFirewallRuleArrayOutput() FirewallRuleArrayOutput {
	return o
}

func (o FirewallRuleArrayOutput) ToFirewallRuleArrayOutputWithContext(ctx context.Context) FirewallRuleArrayOutput {
	return o
}

func (o FirewallRuleArrayOutput) Index(i pulumi.IntInput) FirewallRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FirewallRule {
		return vs[0].([]*FirewallRule)[vs[1].(int)]
	}).(FirewallRuleOutput)
}

type FirewallRuleMapOutput struct{ *pulumi.OutputState }

func (FirewallRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FirewallRule)(nil)).Elem()
}

func (o FirewallRuleMapOutput) ToFirewallRuleMapOutput() FirewallRuleMapOutput {
	return o
}

func (o FirewallRuleMapOutput) ToFirewallRuleMapOutputWithContext(ctx context.Context) FirewallRuleMapOutput {
	return o
}

func (o FirewallRuleMapOutput) MapIndex(k pulumi.StringInput) FirewallRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FirewallRule {
		return vs[0].(map[string]*FirewallRule)[vs[1].(string)]
	}).(FirewallRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallRuleInput)(nil)).Elem(), &FirewallRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallRuleArrayInput)(nil)).Elem(), FirewallRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FirewallRuleMapInput)(nil)).Elem(), FirewallRuleMap{})
	pulumi.RegisterOutputType(FirewallRuleOutput{})
	pulumi.RegisterOutputType(FirewallRuleArrayOutput{})
	pulumi.RegisterOutputType(FirewallRuleMapOutput{})
}
