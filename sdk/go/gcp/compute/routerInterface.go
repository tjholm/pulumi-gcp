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

// Manages a Cloud Router interface. For more information see
// [the official documentation](https://cloud.google.com/compute/docs/cloudrouter)
// and
// [API](https://cloud.google.com/compute/docs/reference/latest/routers).
//
// ## Example Usage
//
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
//			_, err := compute.NewRouterInterface(ctx, "foobar", &compute.RouterInterfaceArgs{
//				IpRange:   pulumi.String("169.254.1.1/30"),
//				Region:    pulumi.String("us-central1"),
//				Router:    pulumi.String("router-1"),
//				VpnTunnel: pulumi.String("tunnel-1"),
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
// Router interfaces can be imported using the `project` (optional), `region`, `router`, and `name`, e.g. * `{{project_id}}/{{region}}/{{router}}/{{name}}` * `{{region}}/{{router}}/{{name}}` When using the `pulumi import` command, router interfaces can be imported using one of the formats above. For example
//
// ```sh
//
//	$ pulumi import gcp:compute/routerInterface:RouterInterface default {{project_id}}/{{region}}/{{router}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/routerInterface:RouterInterface default {{region}}/{{router}}/{{name}}
//
// ```
type RouterInterface struct {
	pulumi.CustomResourceState

	// The name or resource link to the
	// VLAN interconnect for this interface. Changing this forces a new interface to
	// be created. Only one of `vpnTunnel`, `interconnectAttachment` or `subnetwork` can be specified.
	InterconnectAttachment pulumi.StringPtrOutput `pulumi:"interconnectAttachment"`
	// IP address and range of the interface. The IP range must be
	// in the RFC3927 link-local IP space. Changing this forces a new interface to be created.
	IpRange pulumi.StringOutput `pulumi:"ipRange"`
	// A unique name for the interface, required by GCE. Changing
	// this forces a new interface to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The regional private internal IP address that is used
	// to establish BGP sessions to a VM instance acting as a third-party Router Appliance. Changing this forces a new interface to be created.
	PrivateIpAddress pulumi.StringPtrOutput `pulumi:"privateIpAddress"`
	// The ID of the project in which this interface's routerbelongs.
	// If it is not provided, the provider project is used. Changing this forces a new interface to be created.
	Project pulumi.StringOutput `pulumi:"project"`
	// The name of the interface that is redundant to
	// this interface. Changing this forces a new interface to be created.
	RedundantInterface pulumi.StringOutput `pulumi:"redundantInterface"`
	// The region this interface's router sits in.
	// If not specified, the project region will be used. Changing this forces a new interface to be created.
	Region pulumi.StringOutput `pulumi:"region"`
	// The name of the router this interface will be attached to.
	// Changing this forces a new interface to be created.
	//
	// In addition to the above required fields, a router interface must have specified either `ipRange` or exactly one of `vpnTunnel`, `interconnectAttachment` or `subnetwork`, or both.
	//
	// ***
	Router pulumi.StringOutput `pulumi:"router"`
	// The URI of the subnetwork resource that this interface
	// belongs to, which must be in the same region as the Cloud Router. When you establish a BGP session to a VM instance using this interface, the VM instance must belong to the same subnetwork as the subnetwork specified here. Changing this forces a new interface to be created. Only one of `vpnTunnel`, `interconnectAttachment` or `subnetwork` can be specified.
	Subnetwork pulumi.StringPtrOutput `pulumi:"subnetwork"`
	// The name or resource link to the VPN tunnel this
	// interface will be linked to. Changing this forces a new interface to be created. Only
	// one of `vpnTunnel`, `interconnectAttachment` or `subnetwork` can be specified.
	VpnTunnel pulumi.StringPtrOutput `pulumi:"vpnTunnel"`
}

// NewRouterInterface registers a new resource with the given unique name, arguments, and options.
func NewRouterInterface(ctx *pulumi.Context,
	name string, args *RouterInterfaceArgs, opts ...pulumi.ResourceOption) (*RouterInterface, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Router == nil {
		return nil, errors.New("invalid value for required argument 'Router'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RouterInterface
	err := ctx.RegisterResource("gcp:compute/routerInterface:RouterInterface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouterInterface gets an existing RouterInterface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouterInterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouterInterfaceState, opts ...pulumi.ResourceOption) (*RouterInterface, error) {
	var resource RouterInterface
	err := ctx.ReadResource("gcp:compute/routerInterface:RouterInterface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RouterInterface resources.
type routerInterfaceState struct {
	// The name or resource link to the
	// VLAN interconnect for this interface. Changing this forces a new interface to
	// be created. Only one of `vpnTunnel`, `interconnectAttachment` or `subnetwork` can be specified.
	InterconnectAttachment *string `pulumi:"interconnectAttachment"`
	// IP address and range of the interface. The IP range must be
	// in the RFC3927 link-local IP space. Changing this forces a new interface to be created.
	IpRange *string `pulumi:"ipRange"`
	// A unique name for the interface, required by GCE. Changing
	// this forces a new interface to be created.
	Name *string `pulumi:"name"`
	// The regional private internal IP address that is used
	// to establish BGP sessions to a VM instance acting as a third-party Router Appliance. Changing this forces a new interface to be created.
	PrivateIpAddress *string `pulumi:"privateIpAddress"`
	// The ID of the project in which this interface's routerbelongs.
	// If it is not provided, the provider project is used. Changing this forces a new interface to be created.
	Project *string `pulumi:"project"`
	// The name of the interface that is redundant to
	// this interface. Changing this forces a new interface to be created.
	RedundantInterface *string `pulumi:"redundantInterface"`
	// The region this interface's router sits in.
	// If not specified, the project region will be used. Changing this forces a new interface to be created.
	Region *string `pulumi:"region"`
	// The name of the router this interface will be attached to.
	// Changing this forces a new interface to be created.
	//
	// In addition to the above required fields, a router interface must have specified either `ipRange` or exactly one of `vpnTunnel`, `interconnectAttachment` or `subnetwork`, or both.
	//
	// ***
	Router *string `pulumi:"router"`
	// The URI of the subnetwork resource that this interface
	// belongs to, which must be in the same region as the Cloud Router. When you establish a BGP session to a VM instance using this interface, the VM instance must belong to the same subnetwork as the subnetwork specified here. Changing this forces a new interface to be created. Only one of `vpnTunnel`, `interconnectAttachment` or `subnetwork` can be specified.
	Subnetwork *string `pulumi:"subnetwork"`
	// The name or resource link to the VPN tunnel this
	// interface will be linked to. Changing this forces a new interface to be created. Only
	// one of `vpnTunnel`, `interconnectAttachment` or `subnetwork` can be specified.
	VpnTunnel *string `pulumi:"vpnTunnel"`
}

type RouterInterfaceState struct {
	// The name or resource link to the
	// VLAN interconnect for this interface. Changing this forces a new interface to
	// be created. Only one of `vpnTunnel`, `interconnectAttachment` or `subnetwork` can be specified.
	InterconnectAttachment pulumi.StringPtrInput
	// IP address and range of the interface. The IP range must be
	// in the RFC3927 link-local IP space. Changing this forces a new interface to be created.
	IpRange pulumi.StringPtrInput
	// A unique name for the interface, required by GCE. Changing
	// this forces a new interface to be created.
	Name pulumi.StringPtrInput
	// The regional private internal IP address that is used
	// to establish BGP sessions to a VM instance acting as a third-party Router Appliance. Changing this forces a new interface to be created.
	PrivateIpAddress pulumi.StringPtrInput
	// The ID of the project in which this interface's routerbelongs.
	// If it is not provided, the provider project is used. Changing this forces a new interface to be created.
	Project pulumi.StringPtrInput
	// The name of the interface that is redundant to
	// this interface. Changing this forces a new interface to be created.
	RedundantInterface pulumi.StringPtrInput
	// The region this interface's router sits in.
	// If not specified, the project region will be used. Changing this forces a new interface to be created.
	Region pulumi.StringPtrInput
	// The name of the router this interface will be attached to.
	// Changing this forces a new interface to be created.
	//
	// In addition to the above required fields, a router interface must have specified either `ipRange` or exactly one of `vpnTunnel`, `interconnectAttachment` or `subnetwork`, or both.
	//
	// ***
	Router pulumi.StringPtrInput
	// The URI of the subnetwork resource that this interface
	// belongs to, which must be in the same region as the Cloud Router. When you establish a BGP session to a VM instance using this interface, the VM instance must belong to the same subnetwork as the subnetwork specified here. Changing this forces a new interface to be created. Only one of `vpnTunnel`, `interconnectAttachment` or `subnetwork` can be specified.
	Subnetwork pulumi.StringPtrInput
	// The name or resource link to the VPN tunnel this
	// interface will be linked to. Changing this forces a new interface to be created. Only
	// one of `vpnTunnel`, `interconnectAttachment` or `subnetwork` can be specified.
	VpnTunnel pulumi.StringPtrInput
}

func (RouterInterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*routerInterfaceState)(nil)).Elem()
}

type routerInterfaceArgs struct {
	// The name or resource link to the
	// VLAN interconnect for this interface. Changing this forces a new interface to
	// be created. Only one of `vpnTunnel`, `interconnectAttachment` or `subnetwork` can be specified.
	InterconnectAttachment *string `pulumi:"interconnectAttachment"`
	// IP address and range of the interface. The IP range must be
	// in the RFC3927 link-local IP space. Changing this forces a new interface to be created.
	IpRange *string `pulumi:"ipRange"`
	// A unique name for the interface, required by GCE. Changing
	// this forces a new interface to be created.
	Name *string `pulumi:"name"`
	// The regional private internal IP address that is used
	// to establish BGP sessions to a VM instance acting as a third-party Router Appliance. Changing this forces a new interface to be created.
	PrivateIpAddress *string `pulumi:"privateIpAddress"`
	// The ID of the project in which this interface's routerbelongs.
	// If it is not provided, the provider project is used. Changing this forces a new interface to be created.
	Project *string `pulumi:"project"`
	// The name of the interface that is redundant to
	// this interface. Changing this forces a new interface to be created.
	RedundantInterface *string `pulumi:"redundantInterface"`
	// The region this interface's router sits in.
	// If not specified, the project region will be used. Changing this forces a new interface to be created.
	Region *string `pulumi:"region"`
	// The name of the router this interface will be attached to.
	// Changing this forces a new interface to be created.
	//
	// In addition to the above required fields, a router interface must have specified either `ipRange` or exactly one of `vpnTunnel`, `interconnectAttachment` or `subnetwork`, or both.
	//
	// ***
	Router string `pulumi:"router"`
	// The URI of the subnetwork resource that this interface
	// belongs to, which must be in the same region as the Cloud Router. When you establish a BGP session to a VM instance using this interface, the VM instance must belong to the same subnetwork as the subnetwork specified here. Changing this forces a new interface to be created. Only one of `vpnTunnel`, `interconnectAttachment` or `subnetwork` can be specified.
	Subnetwork *string `pulumi:"subnetwork"`
	// The name or resource link to the VPN tunnel this
	// interface will be linked to. Changing this forces a new interface to be created. Only
	// one of `vpnTunnel`, `interconnectAttachment` or `subnetwork` can be specified.
	VpnTunnel *string `pulumi:"vpnTunnel"`
}

// The set of arguments for constructing a RouterInterface resource.
type RouterInterfaceArgs struct {
	// The name or resource link to the
	// VLAN interconnect for this interface. Changing this forces a new interface to
	// be created. Only one of `vpnTunnel`, `interconnectAttachment` or `subnetwork` can be specified.
	InterconnectAttachment pulumi.StringPtrInput
	// IP address and range of the interface. The IP range must be
	// in the RFC3927 link-local IP space. Changing this forces a new interface to be created.
	IpRange pulumi.StringPtrInput
	// A unique name for the interface, required by GCE. Changing
	// this forces a new interface to be created.
	Name pulumi.StringPtrInput
	// The regional private internal IP address that is used
	// to establish BGP sessions to a VM instance acting as a third-party Router Appliance. Changing this forces a new interface to be created.
	PrivateIpAddress pulumi.StringPtrInput
	// The ID of the project in which this interface's routerbelongs.
	// If it is not provided, the provider project is used. Changing this forces a new interface to be created.
	Project pulumi.StringPtrInput
	// The name of the interface that is redundant to
	// this interface. Changing this forces a new interface to be created.
	RedundantInterface pulumi.StringPtrInput
	// The region this interface's router sits in.
	// If not specified, the project region will be used. Changing this forces a new interface to be created.
	Region pulumi.StringPtrInput
	// The name of the router this interface will be attached to.
	// Changing this forces a new interface to be created.
	//
	// In addition to the above required fields, a router interface must have specified either `ipRange` or exactly one of `vpnTunnel`, `interconnectAttachment` or `subnetwork`, or both.
	//
	// ***
	Router pulumi.StringInput
	// The URI of the subnetwork resource that this interface
	// belongs to, which must be in the same region as the Cloud Router. When you establish a BGP session to a VM instance using this interface, the VM instance must belong to the same subnetwork as the subnetwork specified here. Changing this forces a new interface to be created. Only one of `vpnTunnel`, `interconnectAttachment` or `subnetwork` can be specified.
	Subnetwork pulumi.StringPtrInput
	// The name or resource link to the VPN tunnel this
	// interface will be linked to. Changing this forces a new interface to be created. Only
	// one of `vpnTunnel`, `interconnectAttachment` or `subnetwork` can be specified.
	VpnTunnel pulumi.StringPtrInput
}

func (RouterInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routerInterfaceArgs)(nil)).Elem()
}

type RouterInterfaceInput interface {
	pulumi.Input

	ToRouterInterfaceOutput() RouterInterfaceOutput
	ToRouterInterfaceOutputWithContext(ctx context.Context) RouterInterfaceOutput
}

func (*RouterInterface) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterInterface)(nil)).Elem()
}

func (i *RouterInterface) ToRouterInterfaceOutput() RouterInterfaceOutput {
	return i.ToRouterInterfaceOutputWithContext(context.Background())
}

func (i *RouterInterface) ToRouterInterfaceOutputWithContext(ctx context.Context) RouterInterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterInterfaceOutput)
}

// RouterInterfaceArrayInput is an input type that accepts RouterInterfaceArray and RouterInterfaceArrayOutput values.
// You can construct a concrete instance of `RouterInterfaceArrayInput` via:
//
//	RouterInterfaceArray{ RouterInterfaceArgs{...} }
type RouterInterfaceArrayInput interface {
	pulumi.Input

	ToRouterInterfaceArrayOutput() RouterInterfaceArrayOutput
	ToRouterInterfaceArrayOutputWithContext(context.Context) RouterInterfaceArrayOutput
}

type RouterInterfaceArray []RouterInterfaceInput

func (RouterInterfaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterInterface)(nil)).Elem()
}

func (i RouterInterfaceArray) ToRouterInterfaceArrayOutput() RouterInterfaceArrayOutput {
	return i.ToRouterInterfaceArrayOutputWithContext(context.Background())
}

func (i RouterInterfaceArray) ToRouterInterfaceArrayOutputWithContext(ctx context.Context) RouterInterfaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterInterfaceArrayOutput)
}

// RouterInterfaceMapInput is an input type that accepts RouterInterfaceMap and RouterInterfaceMapOutput values.
// You can construct a concrete instance of `RouterInterfaceMapInput` via:
//
//	RouterInterfaceMap{ "key": RouterInterfaceArgs{...} }
type RouterInterfaceMapInput interface {
	pulumi.Input

	ToRouterInterfaceMapOutput() RouterInterfaceMapOutput
	ToRouterInterfaceMapOutputWithContext(context.Context) RouterInterfaceMapOutput
}

type RouterInterfaceMap map[string]RouterInterfaceInput

func (RouterInterfaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterInterface)(nil)).Elem()
}

func (i RouterInterfaceMap) ToRouterInterfaceMapOutput() RouterInterfaceMapOutput {
	return i.ToRouterInterfaceMapOutputWithContext(context.Background())
}

func (i RouterInterfaceMap) ToRouterInterfaceMapOutputWithContext(ctx context.Context) RouterInterfaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterInterfaceMapOutput)
}

type RouterInterfaceOutput struct{ *pulumi.OutputState }

func (RouterInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RouterInterface)(nil)).Elem()
}

func (o RouterInterfaceOutput) ToRouterInterfaceOutput() RouterInterfaceOutput {
	return o
}

func (o RouterInterfaceOutput) ToRouterInterfaceOutputWithContext(ctx context.Context) RouterInterfaceOutput {
	return o
}

// The name or resource link to the
// VLAN interconnect for this interface. Changing this forces a new interface to
// be created. Only one of `vpnTunnel`, `interconnectAttachment` or `subnetwork` can be specified.
func (o RouterInterfaceOutput) InterconnectAttachment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouterInterface) pulumi.StringPtrOutput { return v.InterconnectAttachment }).(pulumi.StringPtrOutput)
}

// IP address and range of the interface. The IP range must be
// in the RFC3927 link-local IP space. Changing this forces a new interface to be created.
func (o RouterInterfaceOutput) IpRange() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterInterface) pulumi.StringOutput { return v.IpRange }).(pulumi.StringOutput)
}

// A unique name for the interface, required by GCE. Changing
// this forces a new interface to be created.
func (o RouterInterfaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterInterface) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The regional private internal IP address that is used
// to establish BGP sessions to a VM instance acting as a third-party Router Appliance. Changing this forces a new interface to be created.
func (o RouterInterfaceOutput) PrivateIpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouterInterface) pulumi.StringPtrOutput { return v.PrivateIpAddress }).(pulumi.StringPtrOutput)
}

// The ID of the project in which this interface's routerbelongs.
// If it is not provided, the provider project is used. Changing this forces a new interface to be created.
func (o RouterInterfaceOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterInterface) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The name of the interface that is redundant to
// this interface. Changing this forces a new interface to be created.
func (o RouterInterfaceOutput) RedundantInterface() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterInterface) pulumi.StringOutput { return v.RedundantInterface }).(pulumi.StringOutput)
}

// The region this interface's router sits in.
// If not specified, the project region will be used. Changing this forces a new interface to be created.
func (o RouterInterfaceOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterInterface) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The name of the router this interface will be attached to.
// Changing this forces a new interface to be created.
//
// In addition to the above required fields, a router interface must have specified either `ipRange` or exactly one of `vpnTunnel`, `interconnectAttachment` or `subnetwork`, or both.
//
// ***
func (o RouterInterfaceOutput) Router() pulumi.StringOutput {
	return o.ApplyT(func(v *RouterInterface) pulumi.StringOutput { return v.Router }).(pulumi.StringOutput)
}

// The URI of the subnetwork resource that this interface
// belongs to, which must be in the same region as the Cloud Router. When you establish a BGP session to a VM instance using this interface, the VM instance must belong to the same subnetwork as the subnetwork specified here. Changing this forces a new interface to be created. Only one of `vpnTunnel`, `interconnectAttachment` or `subnetwork` can be specified.
func (o RouterInterfaceOutput) Subnetwork() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouterInterface) pulumi.StringPtrOutput { return v.Subnetwork }).(pulumi.StringPtrOutput)
}

// The name or resource link to the VPN tunnel this
// interface will be linked to. Changing this forces a new interface to be created. Only
// one of `vpnTunnel`, `interconnectAttachment` or `subnetwork` can be specified.
func (o RouterInterfaceOutput) VpnTunnel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RouterInterface) pulumi.StringPtrOutput { return v.VpnTunnel }).(pulumi.StringPtrOutput)
}

type RouterInterfaceArrayOutput struct{ *pulumi.OutputState }

func (RouterInterfaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RouterInterface)(nil)).Elem()
}

func (o RouterInterfaceArrayOutput) ToRouterInterfaceArrayOutput() RouterInterfaceArrayOutput {
	return o
}

func (o RouterInterfaceArrayOutput) ToRouterInterfaceArrayOutputWithContext(ctx context.Context) RouterInterfaceArrayOutput {
	return o
}

func (o RouterInterfaceArrayOutput) Index(i pulumi.IntInput) RouterInterfaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RouterInterface {
		return vs[0].([]*RouterInterface)[vs[1].(int)]
	}).(RouterInterfaceOutput)
}

type RouterInterfaceMapOutput struct{ *pulumi.OutputState }

func (RouterInterfaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RouterInterface)(nil)).Elem()
}

func (o RouterInterfaceMapOutput) ToRouterInterfaceMapOutput() RouterInterfaceMapOutput {
	return o
}

func (o RouterInterfaceMapOutput) ToRouterInterfaceMapOutputWithContext(ctx context.Context) RouterInterfaceMapOutput {
	return o
}

func (o RouterInterfaceMapOutput) MapIndex(k pulumi.StringInput) RouterInterfaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RouterInterface {
		return vs[0].(map[string]*RouterInterface)[vs[1].(string)]
	}).(RouterInterfaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouterInterfaceInput)(nil)).Elem(), &RouterInterface{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterInterfaceArrayInput)(nil)).Elem(), RouterInterfaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterInterfaceMapInput)(nil)).Elem(), RouterInterfaceMap{})
	pulumi.RegisterOutputType(RouterInterfaceOutput{})
	pulumi.RegisterOutputType(RouterInterfaceArrayOutput{})
	pulumi.RegisterOutputType(RouterInterfaceMapOutput{})
}
