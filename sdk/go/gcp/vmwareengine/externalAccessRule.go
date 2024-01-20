// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vmwareengine

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// External access firewall rules for filtering incoming traffic destined to `ExternalAddress` resources.
//
// To get more information about ExternalAccessRule, see:
//
// * [API documentation](https://cloud.google.com/vmware-engine/docs/reference/rest/v1/projects.locations.networkPolicies.externalAccessRules)
//
// ## Example Usage
// ### Vmware Engine External Access Rule Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/vmwareengine"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vmwareengine.NewNetwork(ctx, "external-access-rule-nw", &vmwareengine.NetworkArgs{
//				Location:    pulumi.String("global"),
//				Type:        pulumi.String("STANDARD"),
//				Description: pulumi.String("PC network description."),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vmwareengine.NewNetworkPolicy(ctx, "external-access-rule-np", &vmwareengine.NetworkPolicyArgs{
//				Location:            pulumi.String("us-west1"),
//				EdgeServicesCidr:    pulumi.String("192.168.30.0/26"),
//				VmwareEngineNetwork: external_access_rule_nw.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vmwareengine.NewExternalAccessRule(ctx, "vmw-engine-external-access-rule", &vmwareengine.ExternalAccessRuleArgs{
//				Parent:     external_access_rule_np.ID(),
//				Priority:   pulumi.Int(101),
//				Action:     pulumi.String("DENY"),
//				IpProtocol: pulumi.String("TCP"),
//				SourceIpRanges: vmwareengine.ExternalAccessRuleSourceIpRangeArray{
//					&vmwareengine.ExternalAccessRuleSourceIpRangeArgs{
//						IpAddressRange: pulumi.String("0.0.0.0/0"),
//					},
//				},
//				SourcePorts: pulumi.StringArray{
//					pulumi.String("80"),
//				},
//				DestinationIpRanges: vmwareengine.ExternalAccessRuleDestinationIpRangeArray{
//					&vmwareengine.ExternalAccessRuleDestinationIpRangeArgs{
//						IpAddressRange: pulumi.String("0.0.0.0/0"),
//					},
//				},
//				DestinationPorts: pulumi.StringArray{
//					pulumi.String("433"),
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
// ### Vmware Engine External Access Rule Full
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/vmwareengine"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vmwareengine.NewNetwork(ctx, "external-access-rule-nw", &vmwareengine.NetworkArgs{
//				Location:    pulumi.String("global"),
//				Type:        pulumi.String("STANDARD"),
//				Description: pulumi.String("PC network description."),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vmwareengine.NewPrivateCloud(ctx, "external-access-rule-pc", &vmwareengine.PrivateCloudArgs{
//				Location:    pulumi.String("us-west1-a"),
//				Description: pulumi.String("Sample test PC."),
//				NetworkConfig: &vmwareengine.PrivateCloudNetworkConfigArgs{
//					ManagementCidr:      pulumi.String("192.168.50.0/24"),
//					VmwareEngineNetwork: external_access_rule_nw.ID(),
//				},
//				ManagementCluster: &vmwareengine.PrivateCloudManagementClusterArgs{
//					ClusterId: pulumi.String("sample-mgmt-cluster"),
//					NodeTypeConfigs: vmwareengine.PrivateCloudManagementClusterNodeTypeConfigArray{
//						&vmwareengine.PrivateCloudManagementClusterNodeTypeConfigArgs{
//							NodeTypeId: pulumi.String("standard-72"),
//							NodeCount:  pulumi.Int(3),
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vmwareengine.NewNetworkPolicy(ctx, "external-access-rule-np", &vmwareengine.NetworkPolicyArgs{
//				Location:            pulumi.String("us-west1"),
//				EdgeServicesCidr:    pulumi.String("192.168.30.0/26"),
//				VmwareEngineNetwork: external_access_rule_nw.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vmwareengine.NewExternalAddress(ctx, "external-access-rule-ea", &vmwareengine.ExternalAddressArgs{
//				Parent:     external_access_rule_pc.ID(),
//				InternalIp: pulumi.String("192.168.0.65"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vmwareengine.NewExternalAccessRule(ctx, "vmw-engine-external-access-rule", &vmwareengine.ExternalAccessRuleArgs{
//				Parent:      external_access_rule_np.ID(),
//				Description: pulumi.String("Sample Description"),
//				Priority:    pulumi.Int(101),
//				Action:      pulumi.String("ALLOW"),
//				IpProtocol:  pulumi.String("tcp"),
//				SourceIpRanges: vmwareengine.ExternalAccessRuleSourceIpRangeArray{
//					&vmwareengine.ExternalAccessRuleSourceIpRangeArgs{
//						IpAddressRange: pulumi.String("0.0.0.0/0"),
//					},
//				},
//				SourcePorts: pulumi.StringArray{
//					pulumi.String("80"),
//				},
//				DestinationIpRanges: vmwareengine.ExternalAccessRuleDestinationIpRangeArray{
//					&vmwareengine.ExternalAccessRuleDestinationIpRangeArgs{
//						ExternalAddress: external_access_rule_ea.ID(),
//					},
//				},
//				DestinationPorts: pulumi.StringArray{
//					pulumi.String("433"),
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
// ExternalAccessRule can be imported using any of these accepted formats* `{{parent}}/externalAccessRules/{{name}}` In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import ExternalAccessRule using one of the formats above. For exampletf import {
//
//	id = "{{parent}}/externalAccessRules/{{name}}"
//
//	to = google_vmwareengine_external_access_rule.default }
//
// ```sh
//
//	$ pulumi import gcp:vmwareengine/externalAccessRule:ExternalAccessRule When using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import), ExternalAccessRule can be imported using one of the formats above. For example
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:vmwareengine/externalAccessRule:ExternalAccessRule default {{parent}}/externalAccessRules/{{name}}
//
// ```
type ExternalAccessRule struct {
	pulumi.CustomResourceState

	// The action that the external access rule performs.
	// Possible values are: `ALLOW`, `DENY`.
	Action pulumi.StringOutput `pulumi:"action"`
	// Creation time of this resource.
	// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and
	// up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// User-provided description for the external access rule.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// If destination ranges are specified, the external access rule applies only to
	// traffic that has a destination IP address in these ranges.
	// Structure is documented below.
	DestinationIpRanges ExternalAccessRuleDestinationIpRangeArrayOutput `pulumi:"destinationIpRanges"`
	// A list of destination ports to which the external access rule applies.
	DestinationPorts pulumi.StringArrayOutput `pulumi:"destinationPorts"`
	// The IP protocol to which the external access rule applies.
	IpProtocol pulumi.StringOutput `pulumi:"ipProtocol"`
	// The ID of the external access rule.
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource name of the network policy.
	// Resource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names.
	// For example: projects/my-project/locations/us-west1-a/networkPolicies/my-policy
	Parent pulumi.StringOutput `pulumi:"parent"`
	// External access rule priority, which determines the external access rule to use when multiple rules apply.
	Priority pulumi.IntOutput `pulumi:"priority"`
	// If source ranges are specified, the external access rule applies only to
	// traffic that has a source IP address in these ranges.
	// Structure is documented below.
	SourceIpRanges ExternalAccessRuleSourceIpRangeArrayOutput `pulumi:"sourceIpRanges"`
	// A list of source ports to which the external access rule applies.
	SourcePorts pulumi.StringArrayOutput `pulumi:"sourcePorts"`
	// State of the Cluster.
	State pulumi.StringOutput `pulumi:"state"`
	// System-generated unique identifier for the resource.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// Last updated time of this resource.
	// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
	// fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewExternalAccessRule registers a new resource with the given unique name, arguments, and options.
func NewExternalAccessRule(ctx *pulumi.Context,
	name string, args *ExternalAccessRuleArgs, opts ...pulumi.ResourceOption) (*ExternalAccessRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.DestinationIpRanges == nil {
		return nil, errors.New("invalid value for required argument 'DestinationIpRanges'")
	}
	if args.DestinationPorts == nil {
		return nil, errors.New("invalid value for required argument 'DestinationPorts'")
	}
	if args.IpProtocol == nil {
		return nil, errors.New("invalid value for required argument 'IpProtocol'")
	}
	if args.Parent == nil {
		return nil, errors.New("invalid value for required argument 'Parent'")
	}
	if args.Priority == nil {
		return nil, errors.New("invalid value for required argument 'Priority'")
	}
	if args.SourceIpRanges == nil {
		return nil, errors.New("invalid value for required argument 'SourceIpRanges'")
	}
	if args.SourcePorts == nil {
		return nil, errors.New("invalid value for required argument 'SourcePorts'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ExternalAccessRule
	err := ctx.RegisterResource("gcp:vmwareengine/externalAccessRule:ExternalAccessRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExternalAccessRule gets an existing ExternalAccessRule resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExternalAccessRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExternalAccessRuleState, opts ...pulumi.ResourceOption) (*ExternalAccessRule, error) {
	var resource ExternalAccessRule
	err := ctx.ReadResource("gcp:vmwareengine/externalAccessRule:ExternalAccessRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExternalAccessRule resources.
type externalAccessRuleState struct {
	// The action that the external access rule performs.
	// Possible values are: `ALLOW`, `DENY`.
	Action *string `pulumi:"action"`
	// Creation time of this resource.
	// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and
	// up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	CreateTime *string `pulumi:"createTime"`
	// User-provided description for the external access rule.
	Description *string `pulumi:"description"`
	// If destination ranges are specified, the external access rule applies only to
	// traffic that has a destination IP address in these ranges.
	// Structure is documented below.
	DestinationIpRanges []ExternalAccessRuleDestinationIpRange `pulumi:"destinationIpRanges"`
	// A list of destination ports to which the external access rule applies.
	DestinationPorts []string `pulumi:"destinationPorts"`
	// The IP protocol to which the external access rule applies.
	IpProtocol *string `pulumi:"ipProtocol"`
	// The ID of the external access rule.
	Name *string `pulumi:"name"`
	// The resource name of the network policy.
	// Resource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names.
	// For example: projects/my-project/locations/us-west1-a/networkPolicies/my-policy
	Parent *string `pulumi:"parent"`
	// External access rule priority, which determines the external access rule to use when multiple rules apply.
	Priority *int `pulumi:"priority"`
	// If source ranges are specified, the external access rule applies only to
	// traffic that has a source IP address in these ranges.
	// Structure is documented below.
	SourceIpRanges []ExternalAccessRuleSourceIpRange `pulumi:"sourceIpRanges"`
	// A list of source ports to which the external access rule applies.
	SourcePorts []string `pulumi:"sourcePorts"`
	// State of the Cluster.
	State *string `pulumi:"state"`
	// System-generated unique identifier for the resource.
	Uid *string `pulumi:"uid"`
	// Last updated time of this resource.
	// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
	// fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime *string `pulumi:"updateTime"`
}

type ExternalAccessRuleState struct {
	// The action that the external access rule performs.
	// Possible values are: `ALLOW`, `DENY`.
	Action pulumi.StringPtrInput
	// Creation time of this resource.
	// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and
	// up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	CreateTime pulumi.StringPtrInput
	// User-provided description for the external access rule.
	Description pulumi.StringPtrInput
	// If destination ranges are specified, the external access rule applies only to
	// traffic that has a destination IP address in these ranges.
	// Structure is documented below.
	DestinationIpRanges ExternalAccessRuleDestinationIpRangeArrayInput
	// A list of destination ports to which the external access rule applies.
	DestinationPorts pulumi.StringArrayInput
	// The IP protocol to which the external access rule applies.
	IpProtocol pulumi.StringPtrInput
	// The ID of the external access rule.
	Name pulumi.StringPtrInput
	// The resource name of the network policy.
	// Resource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names.
	// For example: projects/my-project/locations/us-west1-a/networkPolicies/my-policy
	Parent pulumi.StringPtrInput
	// External access rule priority, which determines the external access rule to use when multiple rules apply.
	Priority pulumi.IntPtrInput
	// If source ranges are specified, the external access rule applies only to
	// traffic that has a source IP address in these ranges.
	// Structure is documented below.
	SourceIpRanges ExternalAccessRuleSourceIpRangeArrayInput
	// A list of source ports to which the external access rule applies.
	SourcePorts pulumi.StringArrayInput
	// State of the Cluster.
	State pulumi.StringPtrInput
	// System-generated unique identifier for the resource.
	Uid pulumi.StringPtrInput
	// Last updated time of this resource.
	// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
	// fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime pulumi.StringPtrInput
}

func (ExternalAccessRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*externalAccessRuleState)(nil)).Elem()
}

type externalAccessRuleArgs struct {
	// The action that the external access rule performs.
	// Possible values are: `ALLOW`, `DENY`.
	Action string `pulumi:"action"`
	// User-provided description for the external access rule.
	Description *string `pulumi:"description"`
	// If destination ranges are specified, the external access rule applies only to
	// traffic that has a destination IP address in these ranges.
	// Structure is documented below.
	DestinationIpRanges []ExternalAccessRuleDestinationIpRange `pulumi:"destinationIpRanges"`
	// A list of destination ports to which the external access rule applies.
	DestinationPorts []string `pulumi:"destinationPorts"`
	// The IP protocol to which the external access rule applies.
	IpProtocol string `pulumi:"ipProtocol"`
	// The ID of the external access rule.
	Name *string `pulumi:"name"`
	// The resource name of the network policy.
	// Resource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names.
	// For example: projects/my-project/locations/us-west1-a/networkPolicies/my-policy
	Parent string `pulumi:"parent"`
	// External access rule priority, which determines the external access rule to use when multiple rules apply.
	Priority int `pulumi:"priority"`
	// If source ranges are specified, the external access rule applies only to
	// traffic that has a source IP address in these ranges.
	// Structure is documented below.
	SourceIpRanges []ExternalAccessRuleSourceIpRange `pulumi:"sourceIpRanges"`
	// A list of source ports to which the external access rule applies.
	SourcePorts []string `pulumi:"sourcePorts"`
}

// The set of arguments for constructing a ExternalAccessRule resource.
type ExternalAccessRuleArgs struct {
	// The action that the external access rule performs.
	// Possible values are: `ALLOW`, `DENY`.
	Action pulumi.StringInput
	// User-provided description for the external access rule.
	Description pulumi.StringPtrInput
	// If destination ranges are specified, the external access rule applies only to
	// traffic that has a destination IP address in these ranges.
	// Structure is documented below.
	DestinationIpRanges ExternalAccessRuleDestinationIpRangeArrayInput
	// A list of destination ports to which the external access rule applies.
	DestinationPorts pulumi.StringArrayInput
	// The IP protocol to which the external access rule applies.
	IpProtocol pulumi.StringInput
	// The ID of the external access rule.
	Name pulumi.StringPtrInput
	// The resource name of the network policy.
	// Resource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names.
	// For example: projects/my-project/locations/us-west1-a/networkPolicies/my-policy
	Parent pulumi.StringInput
	// External access rule priority, which determines the external access rule to use when multiple rules apply.
	Priority pulumi.IntInput
	// If source ranges are specified, the external access rule applies only to
	// traffic that has a source IP address in these ranges.
	// Structure is documented below.
	SourceIpRanges ExternalAccessRuleSourceIpRangeArrayInput
	// A list of source ports to which the external access rule applies.
	SourcePorts pulumi.StringArrayInput
}

func (ExternalAccessRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*externalAccessRuleArgs)(nil)).Elem()
}

type ExternalAccessRuleInput interface {
	pulumi.Input

	ToExternalAccessRuleOutput() ExternalAccessRuleOutput
	ToExternalAccessRuleOutputWithContext(ctx context.Context) ExternalAccessRuleOutput
}

func (*ExternalAccessRule) ElementType() reflect.Type {
	return reflect.TypeOf((**ExternalAccessRule)(nil)).Elem()
}

func (i *ExternalAccessRule) ToExternalAccessRuleOutput() ExternalAccessRuleOutput {
	return i.ToExternalAccessRuleOutputWithContext(context.Background())
}

func (i *ExternalAccessRule) ToExternalAccessRuleOutputWithContext(ctx context.Context) ExternalAccessRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalAccessRuleOutput)
}

// ExternalAccessRuleArrayInput is an input type that accepts ExternalAccessRuleArray and ExternalAccessRuleArrayOutput values.
// You can construct a concrete instance of `ExternalAccessRuleArrayInput` via:
//
//	ExternalAccessRuleArray{ ExternalAccessRuleArgs{...} }
type ExternalAccessRuleArrayInput interface {
	pulumi.Input

	ToExternalAccessRuleArrayOutput() ExternalAccessRuleArrayOutput
	ToExternalAccessRuleArrayOutputWithContext(context.Context) ExternalAccessRuleArrayOutput
}

type ExternalAccessRuleArray []ExternalAccessRuleInput

func (ExternalAccessRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExternalAccessRule)(nil)).Elem()
}

func (i ExternalAccessRuleArray) ToExternalAccessRuleArrayOutput() ExternalAccessRuleArrayOutput {
	return i.ToExternalAccessRuleArrayOutputWithContext(context.Background())
}

func (i ExternalAccessRuleArray) ToExternalAccessRuleArrayOutputWithContext(ctx context.Context) ExternalAccessRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalAccessRuleArrayOutput)
}

// ExternalAccessRuleMapInput is an input type that accepts ExternalAccessRuleMap and ExternalAccessRuleMapOutput values.
// You can construct a concrete instance of `ExternalAccessRuleMapInput` via:
//
//	ExternalAccessRuleMap{ "key": ExternalAccessRuleArgs{...} }
type ExternalAccessRuleMapInput interface {
	pulumi.Input

	ToExternalAccessRuleMapOutput() ExternalAccessRuleMapOutput
	ToExternalAccessRuleMapOutputWithContext(context.Context) ExternalAccessRuleMapOutput
}

type ExternalAccessRuleMap map[string]ExternalAccessRuleInput

func (ExternalAccessRuleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExternalAccessRule)(nil)).Elem()
}

func (i ExternalAccessRuleMap) ToExternalAccessRuleMapOutput() ExternalAccessRuleMapOutput {
	return i.ToExternalAccessRuleMapOutputWithContext(context.Background())
}

func (i ExternalAccessRuleMap) ToExternalAccessRuleMapOutputWithContext(ctx context.Context) ExternalAccessRuleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalAccessRuleMapOutput)
}

type ExternalAccessRuleOutput struct{ *pulumi.OutputState }

func (ExternalAccessRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExternalAccessRule)(nil)).Elem()
}

func (o ExternalAccessRuleOutput) ToExternalAccessRuleOutput() ExternalAccessRuleOutput {
	return o
}

func (o ExternalAccessRuleOutput) ToExternalAccessRuleOutputWithContext(ctx context.Context) ExternalAccessRuleOutput {
	return o
}

// The action that the external access rule performs.
// Possible values are: `ALLOW`, `DENY`.
func (o ExternalAccessRuleOutput) Action() pulumi.StringOutput {
	return o.ApplyT(func(v *ExternalAccessRule) pulumi.StringOutput { return v.Action }).(pulumi.StringOutput)
}

// Creation time of this resource.
// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and
// up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
func (o ExternalAccessRuleOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ExternalAccessRule) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// User-provided description for the external access rule.
func (o ExternalAccessRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExternalAccessRule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// If destination ranges are specified, the external access rule applies only to
// traffic that has a destination IP address in these ranges.
// Structure is documented below.
func (o ExternalAccessRuleOutput) DestinationIpRanges() ExternalAccessRuleDestinationIpRangeArrayOutput {
	return o.ApplyT(func(v *ExternalAccessRule) ExternalAccessRuleDestinationIpRangeArrayOutput {
		return v.DestinationIpRanges
	}).(ExternalAccessRuleDestinationIpRangeArrayOutput)
}

// A list of destination ports to which the external access rule applies.
func (o ExternalAccessRuleOutput) DestinationPorts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ExternalAccessRule) pulumi.StringArrayOutput { return v.DestinationPorts }).(pulumi.StringArrayOutput)
}

// The IP protocol to which the external access rule applies.
func (o ExternalAccessRuleOutput) IpProtocol() pulumi.StringOutput {
	return o.ApplyT(func(v *ExternalAccessRule) pulumi.StringOutput { return v.IpProtocol }).(pulumi.StringOutput)
}

// The ID of the external access rule.
func (o ExternalAccessRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ExternalAccessRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The resource name of the network policy.
// Resource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names.
// For example: projects/my-project/locations/us-west1-a/networkPolicies/my-policy
func (o ExternalAccessRuleOutput) Parent() pulumi.StringOutput {
	return o.ApplyT(func(v *ExternalAccessRule) pulumi.StringOutput { return v.Parent }).(pulumi.StringOutput)
}

// External access rule priority, which determines the external access rule to use when multiple rules apply.
func (o ExternalAccessRuleOutput) Priority() pulumi.IntOutput {
	return o.ApplyT(func(v *ExternalAccessRule) pulumi.IntOutput { return v.Priority }).(pulumi.IntOutput)
}

// If source ranges are specified, the external access rule applies only to
// traffic that has a source IP address in these ranges.
// Structure is documented below.
func (o ExternalAccessRuleOutput) SourceIpRanges() ExternalAccessRuleSourceIpRangeArrayOutput {
	return o.ApplyT(func(v *ExternalAccessRule) ExternalAccessRuleSourceIpRangeArrayOutput { return v.SourceIpRanges }).(ExternalAccessRuleSourceIpRangeArrayOutput)
}

// A list of source ports to which the external access rule applies.
func (o ExternalAccessRuleOutput) SourcePorts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ExternalAccessRule) pulumi.StringArrayOutput { return v.SourcePorts }).(pulumi.StringArrayOutput)
}

// State of the Cluster.
func (o ExternalAccessRuleOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *ExternalAccessRule) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// System-generated unique identifier for the resource.
func (o ExternalAccessRuleOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *ExternalAccessRule) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// Last updated time of this resource.
// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
// fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
func (o ExternalAccessRuleOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *ExternalAccessRule) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

type ExternalAccessRuleArrayOutput struct{ *pulumi.OutputState }

func (ExternalAccessRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExternalAccessRule)(nil)).Elem()
}

func (o ExternalAccessRuleArrayOutput) ToExternalAccessRuleArrayOutput() ExternalAccessRuleArrayOutput {
	return o
}

func (o ExternalAccessRuleArrayOutput) ToExternalAccessRuleArrayOutputWithContext(ctx context.Context) ExternalAccessRuleArrayOutput {
	return o
}

func (o ExternalAccessRuleArrayOutput) Index(i pulumi.IntInput) ExternalAccessRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ExternalAccessRule {
		return vs[0].([]*ExternalAccessRule)[vs[1].(int)]
	}).(ExternalAccessRuleOutput)
}

type ExternalAccessRuleMapOutput struct{ *pulumi.OutputState }

func (ExternalAccessRuleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExternalAccessRule)(nil)).Elem()
}

func (o ExternalAccessRuleMapOutput) ToExternalAccessRuleMapOutput() ExternalAccessRuleMapOutput {
	return o
}

func (o ExternalAccessRuleMapOutput) ToExternalAccessRuleMapOutputWithContext(ctx context.Context) ExternalAccessRuleMapOutput {
	return o
}

func (o ExternalAccessRuleMapOutput) MapIndex(k pulumi.StringInput) ExternalAccessRuleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ExternalAccessRule {
		return vs[0].(map[string]*ExternalAccessRule)[vs[1].(string)]
	}).(ExternalAccessRuleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExternalAccessRuleInput)(nil)).Elem(), &ExternalAccessRule{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExternalAccessRuleArrayInput)(nil)).Elem(), ExternalAccessRuleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExternalAccessRuleMapInput)(nil)).Elem(), ExternalAccessRuleMap{})
	pulumi.RegisterOutputType(ExternalAccessRuleOutput{})
	pulumi.RegisterOutputType(ExternalAccessRuleArrayOutput{})
	pulumi.RegisterOutputType(ExternalAccessRuleMapOutput{})
}
