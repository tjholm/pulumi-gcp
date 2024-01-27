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

// The Compute NetworkFirewallPolicyRule resource
//
// ## Example Usage
// ### Regional
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/compute"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/networksecurity"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/tags"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			basicRegionalNetworksecurityAddressGroup, err := networksecurity.NewAddressGroup(ctx, "basicRegionalNetworksecurityAddressGroup", &networksecurity.AddressGroupArgs{
//				Parent:      pulumi.String("projects/my-project-name"),
//				Description: pulumi.String("Sample regional networksecurity_address_group"),
//				Location:    pulumi.String("us-west1"),
//				Items: pulumi.StringArray{
//					pulumi.String("208.80.154.224/32"),
//				},
//				Type:     pulumi.String("IPV4"),
//				Capacity: pulumi.Int(100),
//			})
//			if err != nil {
//				return err
//			}
//			basicRegionalNetworkFirewallPolicy, err := compute.NewRegionNetworkFirewallPolicy(ctx, "basicRegionalNetworkFirewallPolicy", &compute.RegionNetworkFirewallPolicyArgs{
//				Description: pulumi.String("Sample regional network firewall policy"),
//				Project:     pulumi.String("my-project-name"),
//				Region:      pulumi.String("us-west1"),
//			})
//			if err != nil {
//				return err
//			}
//			basicNetwork, err := compute.NewNetwork(ctx, "basicNetwork", nil)
//			if err != nil {
//				return err
//			}
//			basicKey, err := tags.NewTagKey(ctx, "basicKey", &tags.TagKeyArgs{
//				Description: pulumi.String("For keyname resources."),
//				Parent:      pulumi.String("organizations/123456789"),
//				Purpose:     pulumi.String("GCE_FIREWALL"),
//				ShortName:   pulumi.String("tagkey"),
//				PurposeData: pulumi.StringMap{
//					"network": basicNetwork.Name.ApplyT(func(name string) (string, error) {
//						return fmt.Sprintf("my-project-name/%v", name), nil
//					}).(pulumi.StringOutput),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			basicValue, err := tags.NewTagValue(ctx, "basicValue", &tags.TagValueArgs{
//				Description: pulumi.String("For valuename resources."),
//				Parent: basicKey.Name.ApplyT(func(name string) (string, error) {
//					return fmt.Sprintf("tagKeys/%v", name), nil
//				}).(pulumi.StringOutput),
//				ShortName: pulumi.String("tagvalue"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewRegionNetworkFirewallPolicyRule(ctx, "primary", &compute.RegionNetworkFirewallPolicyRuleArgs{
//				Action:         pulumi.String("allow"),
//				Description:    pulumi.String("This is a simple rule description"),
//				Direction:      pulumi.String("INGRESS"),
//				Disabled:       pulumi.Bool(false),
//				EnableLogging:  pulumi.Bool(true),
//				FirewallPolicy: basicRegionalNetworkFirewallPolicy.Name,
//				Priority:       pulumi.Int(1000),
//				Region:         pulumi.String("us-west1"),
//				RuleName:       pulumi.String("test-rule"),
//				TargetServiceAccounts: pulumi.StringArray{
//					pulumi.String("my@service-account.com"),
//				},
//				Match: &compute.RegionNetworkFirewallPolicyRuleMatchArgs{
//					SrcIpRanges: pulumi.StringArray{
//						pulumi.String("10.100.0.1/32"),
//					},
//					SrcFqdns: pulumi.StringArray{
//						pulumi.String("example.com"),
//					},
//					SrcRegionCodes: pulumi.StringArray{
//						pulumi.String("US"),
//					},
//					SrcThreatIntelligences: pulumi.StringArray{
//						pulumi.String("iplist-known-malicious-ips"),
//					},
//					Layer4Configs: compute.RegionNetworkFirewallPolicyRuleMatchLayer4ConfigArray{
//						&compute.RegionNetworkFirewallPolicyRuleMatchLayer4ConfigArgs{
//							IpProtocol: pulumi.String("all"),
//						},
//					},
//					SrcSecureTags: compute.RegionNetworkFirewallPolicyRuleMatchSrcSecureTagArray{
//						&compute.RegionNetworkFirewallPolicyRuleMatchSrcSecureTagArgs{
//							Name: basicValue.Name.ApplyT(func(name string) (string, error) {
//								return fmt.Sprintf("tagValues/%v", name), nil
//							}).(pulumi.StringOutput),
//						},
//					},
//					SrcAddressGroups: pulumi.StringArray{
//						basicRegionalNetworksecurityAddressGroup.ID(),
//					},
//				},
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
// NetworkFirewallPolicyRule can be imported using any of these accepted formats* `projects/{{project}}/regions/{{region}}/firewallPolicies/{{firewall_policy}}/{{priority}}` * `{{project}}/{{region}}/{{firewall_policy}}/{{priority}}` * `{{region}}/{{firewall_policy}}/{{priority}}` * `{{firewall_policy}}/{{priority}}` When using the `pulumi import` command, NetworkFirewallPolicyRule can be imported using one of the formats above. For example
//
// ```sh
//
//	$ pulumi import gcp:compute/regionNetworkFirewallPolicyRule:RegionNetworkFirewallPolicyRule default projects/{{project}}/regions/{{region}}/firewallPolicies/{{firewall_policy}}/{{priority}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/regionNetworkFirewallPolicyRule:RegionNetworkFirewallPolicyRule default {{project}}/{{region}}/{{firewall_policy}}/{{priority}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/regionNetworkFirewallPolicyRule:RegionNetworkFirewallPolicyRule default {{region}}/{{firewall_policy}}/{{priority}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/regionNetworkFirewallPolicyRule:RegionNetworkFirewallPolicyRule default {{firewall_policy}}/{{priority}}
//
// ```
type RegionNetworkFirewallPolicyRule struct {
	pulumi.CustomResourceState

	// The Action to perform when the client connection triggers the rule. Valid actions are "allow", "deny" and "gotoNext".
	Action pulumi.StringOutput `pulumi:"action"`
	// An optional description for this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The direction in which this rule applies. Possible values: INGRESS, EGRESS
	Direction pulumi.StringOutput `pulumi:"direction"`
	// Denotes whether the firewall policy rule is disabled. When set to true, the firewall policy rule is not enforced and traffic behaves as if it did not exist. If this is unspecified, the firewall policy rule will be enabled.
	Disabled pulumi.BoolPtrOutput `pulumi:"disabled"`
	// Denotes whether to enable logging for a particular rule. If logging is enabled, logs will be exported to the configured export destination in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you cannot enable logging on "gotoNext" rules.
	EnableLogging pulumi.BoolPtrOutput `pulumi:"enableLogging"`
	// The firewall policy of the resource.
	FirewallPolicy pulumi.StringOutput `pulumi:"firewallPolicy"`
	// Type of the resource. Always `compute#firewallPolicyRule` for firewall policy rules
	Kind pulumi.StringOutput `pulumi:"kind"`
	// A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
	Match RegionNetworkFirewallPolicyRuleMatchOutput `pulumi:"match"`
	// An integer indicating the priority of a rule in the list. The priority must be a positive value between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest prority.
	Priority pulumi.IntOutput `pulumi:"priority"`
	// The project for the resource
	Project pulumi.StringOutput `pulumi:"project"`
	// The location of this resource.
	Region pulumi.StringOutput `pulumi:"region"`
	// An optional name for the rule. This field is not a unique identifier and can be updated.
	RuleName pulumi.StringPtrOutput `pulumi:"ruleName"`
	// Calculation of the complexity of a single firewall policy rule.
	RuleTupleCount pulumi.IntOutput `pulumi:"ruleTupleCount"`
	// A list of secure tags that controls which instances the firewall rule applies to. If <code>targetSecureTag</code> are specified, then the firewall rule applies only to instances in the VPC network that have one of those EFFECTIVE secure tags, if all the targetSecureTag are in INEFFECTIVE state, then this rule will be ignored. <code>targetSecureTag</code> may not be set at the same time as <code>targetServiceAccounts</code>. If neither <code>targetServiceAccounts</code> nor <code>targetSecureTag</code> are specified, the firewall rule applies to all instances on the specified network. Maximum number of target label tags allowed is 256.
	TargetSecureTags RegionNetworkFirewallPolicyRuleTargetSecureTagArrayOutput `pulumi:"targetSecureTags"`
	// A list of service accounts indicating the sets of instances that are applied with this rule.
	TargetServiceAccounts pulumi.StringArrayOutput `pulumi:"targetServiceAccounts"`
}

// NewRegionNetworkFirewallPolicyRule registers a new resource with the given unique name, arguments, and options.
func NewRegionNetworkFirewallPolicyRule(ctx *pulumi.Context,
	name string, args *RegionNetworkFirewallPolicyRuleArgs, opts ...pulumi.ResourceOption) (*RegionNetworkFirewallPolicyRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.Direction == nil {
		return nil, errors.New("invalid value for required argument 'Direction'")
	}
	if args.FirewallPolicy == nil {
		return nil, errors.New("invalid value for required argument 'FirewallPolicy'")
	}
	if args.Match == nil {
		return nil, errors.New("invalid value for required argument 'Match'")
	}
	if args.Priority == nil {
		return nil, errors.New("invalid value for required argument 'Priority'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RegionNetworkFirewallPolicyRule
	err := ctx.RegisterResource("gcp:compute/regionNetworkFirewallPolicyRule:RegionNetworkFirewallPolicyRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegionNetworkFirewallPolicyRule gets an existing RegionNetworkFirewallPolicyRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegionNetworkFirewallPolicyRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegionNetworkFirewallPolicyRuleState, opts ...pulumi.ResourceOption) (*RegionNetworkFirewallPolicyRule, error) {
	var resource RegionNetworkFirewallPolicyRule
	err := ctx.ReadResource("gcp:compute/regionNetworkFirewallPolicyRule:RegionNetworkFirewallPolicyRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegionNetworkFirewallPolicyRule resources.
type regionNetworkFirewallPolicyRuleState struct {
	// The Action to perform when the client connection triggers the rule. Valid actions are "allow", "deny" and "gotoNext".
	Action *string `pulumi:"action"`
	// An optional description for this resource.
	Description *string `pulumi:"description"`
	// The direction in which this rule applies. Possible values: INGRESS, EGRESS
	Direction *string `pulumi:"direction"`
	// Denotes whether the firewall policy rule is disabled. When set to true, the firewall policy rule is not enforced and traffic behaves as if it did not exist. If this is unspecified, the firewall policy rule will be enabled.
	Disabled *bool `pulumi:"disabled"`
	// Denotes whether to enable logging for a particular rule. If logging is enabled, logs will be exported to the configured export destination in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you cannot enable logging on "gotoNext" rules.
	EnableLogging *bool `pulumi:"enableLogging"`
	// The firewall policy of the resource.
	FirewallPolicy *string `pulumi:"firewallPolicy"`
	// Type of the resource. Always `compute#firewallPolicyRule` for firewall policy rules
	Kind *string `pulumi:"kind"`
	// A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
	Match *RegionNetworkFirewallPolicyRuleMatch `pulumi:"match"`
	// An integer indicating the priority of a rule in the list. The priority must be a positive value between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest prority.
	Priority *int `pulumi:"priority"`
	// The project for the resource
	Project *string `pulumi:"project"`
	// The location of this resource.
	Region *string `pulumi:"region"`
	// An optional name for the rule. This field is not a unique identifier and can be updated.
	RuleName *string `pulumi:"ruleName"`
	// Calculation of the complexity of a single firewall policy rule.
	RuleTupleCount *int `pulumi:"ruleTupleCount"`
	// A list of secure tags that controls which instances the firewall rule applies to. If <code>targetSecureTag</code> are specified, then the firewall rule applies only to instances in the VPC network that have one of those EFFECTIVE secure tags, if all the targetSecureTag are in INEFFECTIVE state, then this rule will be ignored. <code>targetSecureTag</code> may not be set at the same time as <code>targetServiceAccounts</code>. If neither <code>targetServiceAccounts</code> nor <code>targetSecureTag</code> are specified, the firewall rule applies to all instances on the specified network. Maximum number of target label tags allowed is 256.
	TargetSecureTags []RegionNetworkFirewallPolicyRuleTargetSecureTag `pulumi:"targetSecureTags"`
	// A list of service accounts indicating the sets of instances that are applied with this rule.
	TargetServiceAccounts []string `pulumi:"targetServiceAccounts"`
}

type RegionNetworkFirewallPolicyRuleState struct {
	// The Action to perform when the client connection triggers the rule. Valid actions are "allow", "deny" and "gotoNext".
	Action pulumi.StringPtrInput
	// An optional description for this resource.
	Description pulumi.StringPtrInput
	// The direction in which this rule applies. Possible values: INGRESS, EGRESS
	Direction pulumi.StringPtrInput
	// Denotes whether the firewall policy rule is disabled. When set to true, the firewall policy rule is not enforced and traffic behaves as if it did not exist. If this is unspecified, the firewall policy rule will be enabled.
	Disabled pulumi.BoolPtrInput
	// Denotes whether to enable logging for a particular rule. If logging is enabled, logs will be exported to the configured export destination in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you cannot enable logging on "gotoNext" rules.
	EnableLogging pulumi.BoolPtrInput
	// The firewall policy of the resource.
	FirewallPolicy pulumi.StringPtrInput
	// Type of the resource. Always `compute#firewallPolicyRule` for firewall policy rules
	Kind pulumi.StringPtrInput
	// A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
	Match RegionNetworkFirewallPolicyRuleMatchPtrInput
	// An integer indicating the priority of a rule in the list. The priority must be a positive value between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest prority.
	Priority pulumi.IntPtrInput
	// The project for the resource
	Project pulumi.StringPtrInput
	// The location of this resource.
	Region pulumi.StringPtrInput
	// An optional name for the rule. This field is not a unique identifier and can be updated.
	RuleName pulumi.StringPtrInput
	// Calculation of the complexity of a single firewall policy rule.
	RuleTupleCount pulumi.IntPtrInput
	// A list of secure tags that controls which instances the firewall rule applies to. If <code>targetSecureTag</code> are specified, then the firewall rule applies only to instances in the VPC network that have one of those EFFECTIVE secure tags, if all the targetSecureTag are in INEFFECTIVE state, then this rule will be ignored. <code>targetSecureTag</code> may not be set at the same time as <code>targetServiceAccounts</code>. If neither <code>targetServiceAccounts</code> nor <code>targetSecureTag</code> are specified, the firewall rule applies to all instances on the specified network. Maximum number of target label tags allowed is 256.
	TargetSecureTags RegionNetworkFirewallPolicyRuleTargetSecureTagArrayInput
	// A list of service accounts indicating the sets of instances that are applied with this rule.
	TargetServiceAccounts pulumi.StringArrayInput
}

func (RegionNetworkFirewallPolicyRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*regionNetworkFirewallPolicyRuleState)(nil)).Elem()
}

type regionNetworkFirewallPolicyRuleArgs struct {
	// The Action to perform when the client connection triggers the rule. Valid actions are "allow", "deny" and "gotoNext".
	Action string `pulumi:"action"`
	// An optional description for this resource.
	Description *string `pulumi:"description"`
	// The direction in which this rule applies. Possible values: INGRESS, EGRESS
	Direction string `pulumi:"direction"`
	// Denotes whether the firewall policy rule is disabled. When set to true, the firewall policy rule is not enforced and traffic behaves as if it did not exist. If this is unspecified, the firewall policy rule will be enabled.
	Disabled *bool `pulumi:"disabled"`
	// Denotes whether to enable logging for a particular rule. If logging is enabled, logs will be exported to the configured export destination in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you cannot enable logging on "gotoNext" rules.
	EnableLogging *bool `pulumi:"enableLogging"`
	// The firewall policy of the resource.
	FirewallPolicy string `pulumi:"firewallPolicy"`
	// A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
	Match RegionNetworkFirewallPolicyRuleMatch `pulumi:"match"`
	// An integer indicating the priority of a rule in the list. The priority must be a positive value between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest prority.
	Priority int `pulumi:"priority"`
	// The project for the resource
	Project *string `pulumi:"project"`
	// The location of this resource.
	Region *string `pulumi:"region"`
	// An optional name for the rule. This field is not a unique identifier and can be updated.
	RuleName *string `pulumi:"ruleName"`
	// A list of secure tags that controls which instances the firewall rule applies to. If <code>targetSecureTag</code> are specified, then the firewall rule applies only to instances in the VPC network that have one of those EFFECTIVE secure tags, if all the targetSecureTag are in INEFFECTIVE state, then this rule will be ignored. <code>targetSecureTag</code> may not be set at the same time as <code>targetServiceAccounts</code>. If neither <code>targetServiceAccounts</code> nor <code>targetSecureTag</code> are specified, the firewall rule applies to all instances on the specified network. Maximum number of target label tags allowed is 256.
	TargetSecureTags []RegionNetworkFirewallPolicyRuleTargetSecureTag `pulumi:"targetSecureTags"`
	// A list of service accounts indicating the sets of instances that are applied with this rule.
	TargetServiceAccounts []string `pulumi:"targetServiceAccounts"`
}

// The set of arguments for constructing a RegionNetworkFirewallPolicyRule resource.
type RegionNetworkFirewallPolicyRuleArgs struct {
	// The Action to perform when the client connection triggers the rule. Valid actions are "allow", "deny" and "gotoNext".
	Action pulumi.StringInput
	// An optional description for this resource.
	Description pulumi.StringPtrInput
	// The direction in which this rule applies. Possible values: INGRESS, EGRESS
	Direction pulumi.StringInput
	// Denotes whether the firewall policy rule is disabled. When set to true, the firewall policy rule is not enforced and traffic behaves as if it did not exist. If this is unspecified, the firewall policy rule will be enabled.
	Disabled pulumi.BoolPtrInput
	// Denotes whether to enable logging for a particular rule. If logging is enabled, logs will be exported to the configured export destination in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you cannot enable logging on "gotoNext" rules.
	EnableLogging pulumi.BoolPtrInput
	// The firewall policy of the resource.
	FirewallPolicy pulumi.StringInput
	// A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
	Match RegionNetworkFirewallPolicyRuleMatchInput
	// An integer indicating the priority of a rule in the list. The priority must be a positive value between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest prority.
	Priority pulumi.IntInput
	// The project for the resource
	Project pulumi.StringPtrInput
	// The location of this resource.
	Region pulumi.StringPtrInput
	// An optional name for the rule. This field is not a unique identifier and can be updated.
	RuleName pulumi.StringPtrInput
	// A list of secure tags that controls which instances the firewall rule applies to. If <code>targetSecureTag</code> are specified, then the firewall rule applies only to instances in the VPC network that have one of those EFFECTIVE secure tags, if all the targetSecureTag are in INEFFECTIVE state, then this rule will be ignored. <code>targetSecureTag</code> may not be set at the same time as <code>targetServiceAccounts</code>. If neither <code>targetServiceAccounts</code> nor <code>targetSecureTag</code> are specified, the firewall rule applies to all instances on the specified network. Maximum number of target label tags allowed is 256.
	TargetSecureTags RegionNetworkFirewallPolicyRuleTargetSecureTagArrayInput
	// A list of service accounts indicating the sets of instances that are applied with this rule.
	TargetServiceAccounts pulumi.StringArrayInput
}

func (RegionNetworkFirewallPolicyRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*regionNetworkFirewallPolicyRuleArgs)(nil)).Elem()
}

type RegionNetworkFirewallPolicyRuleInput interface {
	pulumi.Input

	ToRegionNetworkFirewallPolicyRuleOutput() RegionNetworkFirewallPolicyRuleOutput
	ToRegionNetworkFirewallPolicyRuleOutputWithContext(ctx context.Context) RegionNetworkFirewallPolicyRuleOutput
}

func (*RegionNetworkFirewallPolicyRule) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionNetworkFirewallPolicyRule)(nil)).Elem()
}

func (i *RegionNetworkFirewallPolicyRule) ToRegionNetworkFirewallPolicyRuleOutput() RegionNetworkFirewallPolicyRuleOutput {
	return i.ToRegionNetworkFirewallPolicyRuleOutputWithContext(context.Background())
}

func (i *RegionNetworkFirewallPolicyRule) ToRegionNetworkFirewallPolicyRuleOutputWithContext(ctx context.Context) RegionNetworkFirewallPolicyRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionNetworkFirewallPolicyRuleOutput)
}

// RegionNetworkFirewallPolicyRuleArrayInput is an input type that accepts RegionNetworkFirewallPolicyRuleArray and RegionNetworkFirewallPolicyRuleArrayOutput values.
// You can construct a concrete instance of `RegionNetworkFirewallPolicyRuleArrayInput` via:
//
//	RegionNetworkFirewallPolicyRuleArray{ RegionNetworkFirewallPolicyRuleArgs{...} }
type RegionNetworkFirewallPolicyRuleArrayInput interface {
	pulumi.Input

	ToRegionNetworkFirewallPolicyRuleArrayOutput() RegionNetworkFirewallPolicyRuleArrayOutput
	ToRegionNetworkFirewallPolicyRuleArrayOutputWithContext(context.Context) RegionNetworkFirewallPolicyRuleArrayOutput
}

type RegionNetworkFirewallPolicyRuleArray []RegionNetworkFirewallPolicyRuleInput

func (RegionNetworkFirewallPolicyRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RegionNetworkFirewallPolicyRule)(nil)).Elem()
}

func (i RegionNetworkFirewallPolicyRuleArray) ToRegionNetworkFirewallPolicyRuleArrayOutput() RegionNetworkFirewallPolicyRuleArrayOutput {
	return i.ToRegionNetworkFirewallPolicyRuleArrayOutputWithContext(context.Background())
}

func (i RegionNetworkFirewallPolicyRuleArray) ToRegionNetworkFirewallPolicyRuleArrayOutputWithContext(ctx context.Context) RegionNetworkFirewallPolicyRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionNetworkFirewallPolicyRuleArrayOutput)
}

// RegionNetworkFirewallPolicyRuleMapInput is an input type that accepts RegionNetworkFirewallPolicyRuleMap and RegionNetworkFirewallPolicyRuleMapOutput values.
// You can construct a concrete instance of `RegionNetworkFirewallPolicyRuleMapInput` via:
//
//	RegionNetworkFirewallPolicyRuleMap{ "key": RegionNetworkFirewallPolicyRuleArgs{...} }
type RegionNetworkFirewallPolicyRuleMapInput interface {
	pulumi.Input

	ToRegionNetworkFirewallPolicyRuleMapOutput() RegionNetworkFirewallPolicyRuleMapOutput
	ToRegionNetworkFirewallPolicyRuleMapOutputWithContext(context.Context) RegionNetworkFirewallPolicyRuleMapOutput
}

type RegionNetworkFirewallPolicyRuleMap map[string]RegionNetworkFirewallPolicyRuleInput

func (RegionNetworkFirewallPolicyRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RegionNetworkFirewallPolicyRule)(nil)).Elem()
}

func (i RegionNetworkFirewallPolicyRuleMap) ToRegionNetworkFirewallPolicyRuleMapOutput() RegionNetworkFirewallPolicyRuleMapOutput {
	return i.ToRegionNetworkFirewallPolicyRuleMapOutputWithContext(context.Background())
}

func (i RegionNetworkFirewallPolicyRuleMap) ToRegionNetworkFirewallPolicyRuleMapOutputWithContext(ctx context.Context) RegionNetworkFirewallPolicyRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegionNetworkFirewallPolicyRuleMapOutput)
}

type RegionNetworkFirewallPolicyRuleOutput struct{ *pulumi.OutputState }

func (RegionNetworkFirewallPolicyRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegionNetworkFirewallPolicyRule)(nil)).Elem()
}

func (o RegionNetworkFirewallPolicyRuleOutput) ToRegionNetworkFirewallPolicyRuleOutput() RegionNetworkFirewallPolicyRuleOutput {
	return o
}

func (o RegionNetworkFirewallPolicyRuleOutput) ToRegionNetworkFirewallPolicyRuleOutputWithContext(ctx context.Context) RegionNetworkFirewallPolicyRuleOutput {
	return o
}

// The Action to perform when the client connection triggers the rule. Valid actions are "allow", "deny" and "gotoNext".
func (o RegionNetworkFirewallPolicyRuleOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionNetworkFirewallPolicyRule) pulumi.StringOutput { return v.Action }).(pulumi.StringOutput)
}

// An optional description for this resource.
func (o RegionNetworkFirewallPolicyRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegionNetworkFirewallPolicyRule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The direction in which this rule applies. Possible values: INGRESS, EGRESS
func (o RegionNetworkFirewallPolicyRuleOutput) Direction() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionNetworkFirewallPolicyRule) pulumi.StringOutput { return v.Direction }).(pulumi.StringOutput)
}

// Denotes whether the firewall policy rule is disabled. When set to true, the firewall policy rule is not enforced and traffic behaves as if it did not exist. If this is unspecified, the firewall policy rule will be enabled.
func (o RegionNetworkFirewallPolicyRuleOutput) Disabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RegionNetworkFirewallPolicyRule) pulumi.BoolPtrOutput { return v.Disabled }).(pulumi.BoolPtrOutput)
}

// Denotes whether to enable logging for a particular rule. If logging is enabled, logs will be exported to the configured export destination in Stackdriver. Logs may be exported to BigQuery or Pub/Sub. Note: you cannot enable logging on "gotoNext" rules.
func (o RegionNetworkFirewallPolicyRuleOutput) EnableLogging() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RegionNetworkFirewallPolicyRule) pulumi.BoolPtrOutput { return v.EnableLogging }).(pulumi.BoolPtrOutput)
}

// The firewall policy of the resource.
func (o RegionNetworkFirewallPolicyRuleOutput) FirewallPolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionNetworkFirewallPolicyRule) pulumi.StringOutput { return v.FirewallPolicy }).(pulumi.StringOutput)
}

// Type of the resource. Always `compute#firewallPolicyRule` for firewall policy rules
func (o RegionNetworkFirewallPolicyRuleOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionNetworkFirewallPolicyRule) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// A match condition that incoming traffic is evaluated against. If it evaluates to true, the corresponding 'action' is enforced.
func (o RegionNetworkFirewallPolicyRuleOutput) Match() RegionNetworkFirewallPolicyRuleMatchOutput {
	return o.ApplyT(func(v *RegionNetworkFirewallPolicyRule) RegionNetworkFirewallPolicyRuleMatchOutput { return v.Match }).(RegionNetworkFirewallPolicyRuleMatchOutput)
}

// An integer indicating the priority of a rule in the list. The priority must be a positive value between 0 and 2147483647. Rules are evaluated from highest to lowest priority where 0 is the highest priority and 2147483647 is the lowest prority.
func (o RegionNetworkFirewallPolicyRuleOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v *RegionNetworkFirewallPolicyRule) pulumi.IntOutput { return v.Priority }).(pulumi.IntOutput)
}

// The project for the resource
func (o RegionNetworkFirewallPolicyRuleOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionNetworkFirewallPolicyRule) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The location of this resource.
func (o RegionNetworkFirewallPolicyRuleOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *RegionNetworkFirewallPolicyRule) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// An optional name for the rule. This field is not a unique identifier and can be updated.
func (o RegionNetworkFirewallPolicyRuleOutput) RuleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegionNetworkFirewallPolicyRule) pulumi.StringPtrOutput { return v.RuleName }).(pulumi.StringPtrOutput)
}

// Calculation of the complexity of a single firewall policy rule.
func (o RegionNetworkFirewallPolicyRuleOutput) RuleTupleCount() pulumi.IntOutput {
	return o.ApplyT(func(v *RegionNetworkFirewallPolicyRule) pulumi.IntOutput { return v.RuleTupleCount }).(pulumi.IntOutput)
}

// A list of secure tags that controls which instances the firewall rule applies to. If <code>targetSecureTag</code> are specified, then the firewall rule applies only to instances in the VPC network that have one of those EFFECTIVE secure tags, if all the targetSecureTag are in INEFFECTIVE state, then this rule will be ignored. <code>targetSecureTag</code> may not be set at the same time as <code>targetServiceAccounts</code>. If neither <code>targetServiceAccounts</code> nor <code>targetSecureTag</code> are specified, the firewall rule applies to all instances on the specified network. Maximum number of target label tags allowed is 256.
func (o RegionNetworkFirewallPolicyRuleOutput) TargetSecureTags() RegionNetworkFirewallPolicyRuleTargetSecureTagArrayOutput {
	return o.ApplyT(func(v *RegionNetworkFirewallPolicyRule) RegionNetworkFirewallPolicyRuleTargetSecureTagArrayOutput {
		return v.TargetSecureTags
	}).(RegionNetworkFirewallPolicyRuleTargetSecureTagArrayOutput)
}

// A list of service accounts indicating the sets of instances that are applied with this rule.
func (o RegionNetworkFirewallPolicyRuleOutput) TargetServiceAccounts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RegionNetworkFirewallPolicyRule) pulumi.StringArrayOutput { return v.TargetServiceAccounts }).(pulumi.StringArrayOutput)
}

type RegionNetworkFirewallPolicyRuleArrayOutput struct{ *pulumi.OutputState }

func (RegionNetworkFirewallPolicyRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RegionNetworkFirewallPolicyRule)(nil)).Elem()
}

func (o RegionNetworkFirewallPolicyRuleArrayOutput) ToRegionNetworkFirewallPolicyRuleArrayOutput() RegionNetworkFirewallPolicyRuleArrayOutput {
	return o
}

func (o RegionNetworkFirewallPolicyRuleArrayOutput) ToRegionNetworkFirewallPolicyRuleArrayOutputWithContext(ctx context.Context) RegionNetworkFirewallPolicyRuleArrayOutput {
	return o
}

func (o RegionNetworkFirewallPolicyRuleArrayOutput) Index(i pulumi.IntInput) RegionNetworkFirewallPolicyRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RegionNetworkFirewallPolicyRule {
		return vs[0].([]*RegionNetworkFirewallPolicyRule)[vs[1].(int)]
	}).(RegionNetworkFirewallPolicyRuleOutput)
}

type RegionNetworkFirewallPolicyRuleMapOutput struct{ *pulumi.OutputState }

func (RegionNetworkFirewallPolicyRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RegionNetworkFirewallPolicyRule)(nil)).Elem()
}

func (o RegionNetworkFirewallPolicyRuleMapOutput) ToRegionNetworkFirewallPolicyRuleMapOutput() RegionNetworkFirewallPolicyRuleMapOutput {
	return o
}

func (o RegionNetworkFirewallPolicyRuleMapOutput) ToRegionNetworkFirewallPolicyRuleMapOutputWithContext(ctx context.Context) RegionNetworkFirewallPolicyRuleMapOutput {
	return o
}

func (o RegionNetworkFirewallPolicyRuleMapOutput) MapIndex(k pulumi.StringInput) RegionNetworkFirewallPolicyRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RegionNetworkFirewallPolicyRule {
		return vs[0].(map[string]*RegionNetworkFirewallPolicyRule)[vs[1].(string)]
	}).(RegionNetworkFirewallPolicyRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RegionNetworkFirewallPolicyRuleInput)(nil)).Elem(), &RegionNetworkFirewallPolicyRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegionNetworkFirewallPolicyRuleArrayInput)(nil)).Elem(), RegionNetworkFirewallPolicyRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RegionNetworkFirewallPolicyRuleMapInput)(nil)).Elem(), RegionNetworkFirewallPolicyRuleMap{})
	pulumi.RegisterOutputType(RegionNetworkFirewallPolicyRuleOutput{})
	pulumi.RegisterOutputType(RegionNetworkFirewallPolicyRuleArrayOutput{})
	pulumi.RegisterOutputType(RegionNetworkFirewallPolicyRuleMapOutput{})
}
