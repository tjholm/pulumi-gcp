// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dns

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A policy is a collection of DNS rules applied to one or more Virtual
// Private Cloud resources.
//
// To get more information about Policy, see:
//
// * [API documentation](https://cloud.google.com/dns/docs/reference/v1beta2/policies)
// * How-to Guides
//     * [Using DNS server policies](https://cloud.google.com/dns/zones/#using-dns-server-policies)
//
// ## Example Usage
// ### Dns Policy Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/compute"
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/dns"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := compute.NewNetwork(ctx, "network_1", &compute.NetworkArgs{
// 			AutoCreateSubnetworks: pulumi.Bool(false),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewNetwork(ctx, "network_2", &compute.NetworkArgs{
// 			AutoCreateSubnetworks: pulumi.Bool(false),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = dns.NewPolicy(ctx, "example_policy", &dns.PolicyArgs{
// 			EnableInboundForwarding: pulumi.Bool(true),
// 			EnableLogging:           pulumi.Bool(true),
// 			AlternativeNameServerConfig: &dns.PolicyAlternativeNameServerConfigArgs{
// 				TargetNameServers: dns.PolicyAlternativeNameServerConfigTargetNameServerArray{
// 					&dns.PolicyAlternativeNameServerConfigTargetNameServerArgs{
// 						Ipv4Address:    pulumi.String("172.16.1.10"),
// 						ForwardingPath: pulumi.String("private"),
// 					},
// 					&dns.PolicyAlternativeNameServerConfigTargetNameServerArgs{
// 						Ipv4Address: pulumi.String("172.16.1.20"),
// 					},
// 				},
// 			},
// 			Networks: dns.PolicyNetworkArray{
// 				&dns.PolicyNetworkArgs{
// 					NetworkUrl: network_1.ID(),
// 				},
// 				&dns.PolicyNetworkArgs{
// 					NetworkUrl: network_2.ID(),
// 				},
// 			},
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
// Policy can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:dns/policy:Policy default projects/{{project}}/policies/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:dns/policy:Policy default {{project}}/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:dns/policy:Policy default {{name}}
// ```
type Policy struct {
	pulumi.CustomResourceState

	// Sets an alternative name server for the associated networks.
	// When specified, all DNS queries are forwarded to a name server that you choose.
	// Names such as .internal are not available when an alternative name server is specified.
	// Structure is documented below.
	AlternativeNameServerConfig PolicyAlternativeNameServerConfigPtrOutput `pulumi:"alternativeNameServerConfig"`
	// A textual description field. Defaults to 'Managed by Pulumi'.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Allows networks bound to this policy to receive DNS queries sent
	// by VMs or applications over VPN connections. When enabled, a
	// virtual IP address will be allocated from each of the sub-networks
	// that are bound to this policy.
	EnableInboundForwarding pulumi.BoolPtrOutput `pulumi:"enableInboundForwarding"`
	// Controls whether logging is enabled for the networks bound to this policy.
	// Defaults to no logging if not set.
	EnableLogging pulumi.BoolPtrOutput `pulumi:"enableLogging"`
	// User assigned name for this policy.
	Name pulumi.StringOutput `pulumi:"name"`
	// List of network names specifying networks to which this policy is applied.
	// Structure is documented below.
	Networks PolicyNetworkArrayOutput `pulumi:"networks"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewPolicy registers a new resource with the given unique name, arguments, and options.
func NewPolicy(ctx *pulumi.Context,
	name string, args *PolicyArgs, opts ...pulumi.ResourceOption) (*Policy, error) {
	if args == nil {
		args = &PolicyArgs{}
	}

	var resource Policy
	err := ctx.RegisterResource("gcp:dns/policy:Policy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicy gets an existing Policy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyState, opts ...pulumi.ResourceOption) (*Policy, error) {
	var resource Policy
	err := ctx.ReadResource("gcp:dns/policy:Policy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Policy resources.
type policyState struct {
	// Sets an alternative name server for the associated networks.
	// When specified, all DNS queries are forwarded to a name server that you choose.
	// Names such as .internal are not available when an alternative name server is specified.
	// Structure is documented below.
	AlternativeNameServerConfig *PolicyAlternativeNameServerConfig `pulumi:"alternativeNameServerConfig"`
	// A textual description field. Defaults to 'Managed by Pulumi'.
	Description *string `pulumi:"description"`
	// Allows networks bound to this policy to receive DNS queries sent
	// by VMs or applications over VPN connections. When enabled, a
	// virtual IP address will be allocated from each of the sub-networks
	// that are bound to this policy.
	EnableInboundForwarding *bool `pulumi:"enableInboundForwarding"`
	// Controls whether logging is enabled for the networks bound to this policy.
	// Defaults to no logging if not set.
	EnableLogging *bool `pulumi:"enableLogging"`
	// User assigned name for this policy.
	Name *string `pulumi:"name"`
	// List of network names specifying networks to which this policy is applied.
	// Structure is documented below.
	Networks []PolicyNetwork `pulumi:"networks"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type PolicyState struct {
	// Sets an alternative name server for the associated networks.
	// When specified, all DNS queries are forwarded to a name server that you choose.
	// Names such as .internal are not available when an alternative name server is specified.
	// Structure is documented below.
	AlternativeNameServerConfig PolicyAlternativeNameServerConfigPtrInput
	// A textual description field. Defaults to 'Managed by Pulumi'.
	Description pulumi.StringPtrInput
	// Allows networks bound to this policy to receive DNS queries sent
	// by VMs or applications over VPN connections. When enabled, a
	// virtual IP address will be allocated from each of the sub-networks
	// that are bound to this policy.
	EnableInboundForwarding pulumi.BoolPtrInput
	// Controls whether logging is enabled for the networks bound to this policy.
	// Defaults to no logging if not set.
	EnableLogging pulumi.BoolPtrInput
	// User assigned name for this policy.
	Name pulumi.StringPtrInput
	// List of network names specifying networks to which this policy is applied.
	// Structure is documented below.
	Networks PolicyNetworkArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (PolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyState)(nil)).Elem()
}

type policyArgs struct {
	// Sets an alternative name server for the associated networks.
	// When specified, all DNS queries are forwarded to a name server that you choose.
	// Names such as .internal are not available when an alternative name server is specified.
	// Structure is documented below.
	AlternativeNameServerConfig *PolicyAlternativeNameServerConfig `pulumi:"alternativeNameServerConfig"`
	// A textual description field. Defaults to 'Managed by Pulumi'.
	Description *string `pulumi:"description"`
	// Allows networks bound to this policy to receive DNS queries sent
	// by VMs or applications over VPN connections. When enabled, a
	// virtual IP address will be allocated from each of the sub-networks
	// that are bound to this policy.
	EnableInboundForwarding *bool `pulumi:"enableInboundForwarding"`
	// Controls whether logging is enabled for the networks bound to this policy.
	// Defaults to no logging if not set.
	EnableLogging *bool `pulumi:"enableLogging"`
	// User assigned name for this policy.
	Name *string `pulumi:"name"`
	// List of network names specifying networks to which this policy is applied.
	// Structure is documented below.
	Networks []PolicyNetwork `pulumi:"networks"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a Policy resource.
type PolicyArgs struct {
	// Sets an alternative name server for the associated networks.
	// When specified, all DNS queries are forwarded to a name server that you choose.
	// Names such as .internal are not available when an alternative name server is specified.
	// Structure is documented below.
	AlternativeNameServerConfig PolicyAlternativeNameServerConfigPtrInput
	// A textual description field. Defaults to 'Managed by Pulumi'.
	Description pulumi.StringPtrInput
	// Allows networks bound to this policy to receive DNS queries sent
	// by VMs or applications over VPN connections. When enabled, a
	// virtual IP address will be allocated from each of the sub-networks
	// that are bound to this policy.
	EnableInboundForwarding pulumi.BoolPtrInput
	// Controls whether logging is enabled for the networks bound to this policy.
	// Defaults to no logging if not set.
	EnableLogging pulumi.BoolPtrInput
	// User assigned name for this policy.
	Name pulumi.StringPtrInput
	// List of network names specifying networks to which this policy is applied.
	// Structure is documented below.
	Networks PolicyNetworkArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (PolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyArgs)(nil)).Elem()
}

type PolicyInput interface {
	pulumi.Input

	ToPolicyOutput() PolicyOutput
	ToPolicyOutputWithContext(ctx context.Context) PolicyOutput
}

func (*Policy) ElementType() reflect.Type {
	return reflect.TypeOf((*Policy)(nil))
}

func (i *Policy) ToPolicyOutput() PolicyOutput {
	return i.ToPolicyOutputWithContext(context.Background())
}

func (i *Policy) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyOutput)
}

func (i *Policy) ToPolicyPtrOutput() PolicyPtrOutput {
	return i.ToPolicyPtrOutputWithContext(context.Background())
}

func (i *Policy) ToPolicyPtrOutputWithContext(ctx context.Context) PolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyPtrOutput)
}

type PolicyPtrInput interface {
	pulumi.Input

	ToPolicyPtrOutput() PolicyPtrOutput
	ToPolicyPtrOutputWithContext(ctx context.Context) PolicyPtrOutput
}

type policyPtrType PolicyArgs

func (*policyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Policy)(nil))
}

func (i *policyPtrType) ToPolicyPtrOutput() PolicyPtrOutput {
	return i.ToPolicyPtrOutputWithContext(context.Background())
}

func (i *policyPtrType) ToPolicyPtrOutputWithContext(ctx context.Context) PolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyPtrOutput)
}

// PolicyArrayInput is an input type that accepts PolicyArray and PolicyArrayOutput values.
// You can construct a concrete instance of `PolicyArrayInput` via:
//
//          PolicyArray{ PolicyArgs{...} }
type PolicyArrayInput interface {
	pulumi.Input

	ToPolicyArrayOutput() PolicyArrayOutput
	ToPolicyArrayOutputWithContext(context.Context) PolicyArrayOutput
}

type PolicyArray []PolicyInput

func (PolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Policy)(nil))
}

func (i PolicyArray) ToPolicyArrayOutput() PolicyArrayOutput {
	return i.ToPolicyArrayOutputWithContext(context.Background())
}

func (i PolicyArray) ToPolicyArrayOutputWithContext(ctx context.Context) PolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyArrayOutput)
}

// PolicyMapInput is an input type that accepts PolicyMap and PolicyMapOutput values.
// You can construct a concrete instance of `PolicyMapInput` via:
//
//          PolicyMap{ "key": PolicyArgs{...} }
type PolicyMapInput interface {
	pulumi.Input

	ToPolicyMapOutput() PolicyMapOutput
	ToPolicyMapOutputWithContext(context.Context) PolicyMapOutput
}

type PolicyMap map[string]PolicyInput

func (PolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Policy)(nil))
}

func (i PolicyMap) ToPolicyMapOutput() PolicyMapOutput {
	return i.ToPolicyMapOutputWithContext(context.Background())
}

func (i PolicyMap) ToPolicyMapOutputWithContext(ctx context.Context) PolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyMapOutput)
}

type PolicyOutput struct {
	*pulumi.OutputState
}

func (PolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Policy)(nil))
}

func (o PolicyOutput) ToPolicyOutput() PolicyOutput {
	return o
}

func (o PolicyOutput) ToPolicyOutputWithContext(ctx context.Context) PolicyOutput {
	return o
}

func (o PolicyOutput) ToPolicyPtrOutput() PolicyPtrOutput {
	return o.ToPolicyPtrOutputWithContext(context.Background())
}

func (o PolicyOutput) ToPolicyPtrOutputWithContext(ctx context.Context) PolicyPtrOutput {
	return o.ApplyT(func(v Policy) *Policy {
		return &v
	}).(PolicyPtrOutput)
}

type PolicyPtrOutput struct {
	*pulumi.OutputState
}

func (PolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Policy)(nil))
}

func (o PolicyPtrOutput) ToPolicyPtrOutput() PolicyPtrOutput {
	return o
}

func (o PolicyPtrOutput) ToPolicyPtrOutputWithContext(ctx context.Context) PolicyPtrOutput {
	return o
}

type PolicyArrayOutput struct{ *pulumi.OutputState }

func (PolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Policy)(nil))
}

func (o PolicyArrayOutput) ToPolicyArrayOutput() PolicyArrayOutput {
	return o
}

func (o PolicyArrayOutput) ToPolicyArrayOutputWithContext(ctx context.Context) PolicyArrayOutput {
	return o
}

func (o PolicyArrayOutput) Index(i pulumi.IntInput) PolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Policy {
		return vs[0].([]Policy)[vs[1].(int)]
	}).(PolicyOutput)
}

type PolicyMapOutput struct{ *pulumi.OutputState }

func (PolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Policy)(nil))
}

func (o PolicyMapOutput) ToPolicyMapOutput() PolicyMapOutput {
	return o
}

func (o PolicyMapOutput) ToPolicyMapOutputWithContext(ctx context.Context) PolicyMapOutput {
	return o
}

func (o PolicyMapOutput) MapIndex(k pulumi.StringInput) PolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Policy {
		return vs[0].(map[string]Policy)[vs[1].(string)]
	}).(PolicyOutput)
}

func init() {
	pulumi.RegisterOutputType(PolicyOutput{})
	pulumi.RegisterOutputType(PolicyPtrOutput{})
	pulumi.RegisterOutputType(PolicyArrayOutput{})
	pulumi.RegisterOutputType(PolicyMapOutput{})
}
