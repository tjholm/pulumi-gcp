// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Represents a Route resource.
//
// A route is a rule that specifies how certain packets should be handled by
// the virtual network. Routes are associated with virtual machines by tag,
// and the set of routes for a particular virtual machine is called its
// routing table. For each packet leaving a virtual machine, the system
// searches that virtual machine's routing table for a single best matching
// route.
//
// Routes match packets by destination IP address, preferring smaller or more
// specific ranges over larger ones. If there is a tie, the system selects
// the route with the smallest priority value. If there is still a tie, it
// uses the layer three and four packet headers to select just one of the
// remaining matching routes. The packet is then forwarded as specified by
// the nextHop field of the winning route -- either to another virtual
// machine destination, a virtual machine gateway or a Compute
// Engine-operated gateway. Packets that do not match any route in the
// sending virtual machine's routing table will be dropped.
//
// A Route resource must have exactly one specification of either
// nextHopGateway, nextHopInstance, nextHopIp, nextHopVpnTunnel, or
// nextHopIlb.
//
// To get more information about Route, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/routes)
// * How-to Guides
//   - [Using Routes](https://cloud.google.com/vpc/docs/using-routes)
//
// ## Example Usage
// ### Route Basic
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
//			defaultNetwork, err := compute.NewNetwork(ctx, "defaultNetwork", nil)
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewRoute(ctx, "defaultRoute", &compute.RouteArgs{
//				DestRange: pulumi.String("15.0.0.0/24"),
//				Network:   defaultNetwork.Name,
//				NextHopIp: pulumi.String("10.132.1.5"),
//				Priority:  pulumi.Int(100),
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
// # Route can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:compute/route:Route default projects/{{project}}/global/routes/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/route:Route default {{project}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/route:Route default {{name}}
//
// ```
type Route struct {
	pulumi.CustomResourceState

	// An optional description of this resource. Provide this property
	// when you create the resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The destination range of outgoing packets that this route applies to.
	// Only IPv4 is supported.
	DestRange pulumi.StringOutput `pulumi:"destRange"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the
	// last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// The network that this route applies to.
	//
	// ***
	Network pulumi.StringOutput `pulumi:"network"`
	// URL to a gateway that should handle matching packets.
	// Currently, you can only specify the internet gateway, using a full or
	// partial valid URL:
	// * `https://www.googleapis.com/compute/v1/projects/project/global/gateways/default-internet-gateway`
	// * `projects/project/global/gateways/default-internet-gateway`
	// * `global/gateways/default-internet-gateway`
	// * The string `default-internet-gateway`.
	NextHopGateway pulumi.StringPtrOutput `pulumi:"nextHopGateway"`
	// The IP address or URL to a forwarding rule of type
	// loadBalancingScheme=INTERNAL that should handle matching
	// packets.
	// With the GA provider you can only specify the forwarding
	// rule as a partial or full URL. For example, the following
	// are all valid values:
	// * 10.128.0.56
	// * https://www.googleapis.com/compute/v1/projects/project/regions/region/forwardingRules/forwardingRule
	// * regions/region/forwardingRules/forwardingRule
	//   When the beta provider, you can also specify the IP address
	//   of a forwarding rule from the same VPC or any peered VPC.
	//   Note that this can only be used when the destinationRange is
	//   a public (non-RFC 1918) IP CIDR range.
	NextHopIlb pulumi.StringPtrOutput `pulumi:"nextHopIlb"`
	// URL to an instance that should handle matching packets.
	// You can specify this as a full or partial URL. For example:
	// * `https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/instance`
	// * `projects/project/zones/zone/instances/instance`
	// * `zones/zone/instances/instance`
	// * Just the instance name, with the zone in `nextHopInstanceZone`.
	NextHopInstance pulumi.StringPtrOutput `pulumi:"nextHopInstance"`
	// (Optional when `nextHopInstance` is
	// specified)  The zone of the instance specified in
	// `nextHopInstance`.  Omit if `nextHopInstance` is specified as
	// a URL.
	NextHopInstanceZone pulumi.StringOutput `pulumi:"nextHopInstanceZone"`
	// Network IP address of an instance that should handle matching packets.
	NextHopIp pulumi.StringOutput `pulumi:"nextHopIp"`
	// URL to a Network that should handle matching packets.
	NextHopNetwork pulumi.StringOutput `pulumi:"nextHopNetwork"`
	// URL to a VpnTunnel that should handle matching packets.
	NextHopVpnTunnel pulumi.StringPtrOutput `pulumi:"nextHopVpnTunnel"`
	// The priority of this route. Priority is used to break ties in cases
	// where there is more than one matching route of equal prefix length.
	// In the case of two routes with equal prefix length, the one with the
	// lowest-numbered priority value wins.
	// Default value is 1000. Valid range is 0 through 65535.
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// A list of instance tags to which this route applies.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
}

// NewRoute registers a new resource with the given unique name, arguments, and options.
func NewRoute(ctx *pulumi.Context,
	name string, args *RouteArgs, opts ...pulumi.ResourceOption) (*Route, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DestRange == nil {
		return nil, errors.New("invalid value for required argument 'DestRange'")
	}
	if args.Network == nil {
		return nil, errors.New("invalid value for required argument 'Network'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Route
	err := ctx.RegisterResource("gcp:compute/route:Route", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoute gets an existing Route resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteState, opts ...pulumi.ResourceOption) (*Route, error) {
	var resource Route
	err := ctx.ReadResource("gcp:compute/route:Route", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Route resources.
type routeState struct {
	// An optional description of this resource. Provide this property
	// when you create the resource.
	Description *string `pulumi:"description"`
	// The destination range of outgoing packets that this route applies to.
	// Only IPv4 is supported.
	DestRange *string `pulumi:"destRange"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the
	// last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The network that this route applies to.
	//
	// ***
	Network *string `pulumi:"network"`
	// URL to a gateway that should handle matching packets.
	// Currently, you can only specify the internet gateway, using a full or
	// partial valid URL:
	// * `https://www.googleapis.com/compute/v1/projects/project/global/gateways/default-internet-gateway`
	// * `projects/project/global/gateways/default-internet-gateway`
	// * `global/gateways/default-internet-gateway`
	// * The string `default-internet-gateway`.
	NextHopGateway *string `pulumi:"nextHopGateway"`
	// The IP address or URL to a forwarding rule of type
	// loadBalancingScheme=INTERNAL that should handle matching
	// packets.
	// With the GA provider you can only specify the forwarding
	// rule as a partial or full URL. For example, the following
	// are all valid values:
	// * 10.128.0.56
	// * https://www.googleapis.com/compute/v1/projects/project/regions/region/forwardingRules/forwardingRule
	// * regions/region/forwardingRules/forwardingRule
	//   When the beta provider, you can also specify the IP address
	//   of a forwarding rule from the same VPC or any peered VPC.
	//   Note that this can only be used when the destinationRange is
	//   a public (non-RFC 1918) IP CIDR range.
	NextHopIlb *string `pulumi:"nextHopIlb"`
	// URL to an instance that should handle matching packets.
	// You can specify this as a full or partial URL. For example:
	// * `https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/instance`
	// * `projects/project/zones/zone/instances/instance`
	// * `zones/zone/instances/instance`
	// * Just the instance name, with the zone in `nextHopInstanceZone`.
	NextHopInstance *string `pulumi:"nextHopInstance"`
	// (Optional when `nextHopInstance` is
	// specified)  The zone of the instance specified in
	// `nextHopInstance`.  Omit if `nextHopInstance` is specified as
	// a URL.
	NextHopInstanceZone *string `pulumi:"nextHopInstanceZone"`
	// Network IP address of an instance that should handle matching packets.
	NextHopIp *string `pulumi:"nextHopIp"`
	// URL to a Network that should handle matching packets.
	NextHopNetwork *string `pulumi:"nextHopNetwork"`
	// URL to a VpnTunnel that should handle matching packets.
	NextHopVpnTunnel *string `pulumi:"nextHopVpnTunnel"`
	// The priority of this route. Priority is used to break ties in cases
	// where there is more than one matching route of equal prefix length.
	// In the case of two routes with equal prefix length, the one with the
	// lowest-numbered priority value wins.
	// Default value is 1000. Valid range is 0 through 65535.
	Priority *int `pulumi:"priority"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// A list of instance tags to which this route applies.
	Tags []string `pulumi:"tags"`
}

type RouteState struct {
	// An optional description of this resource. Provide this property
	// when you create the resource.
	Description pulumi.StringPtrInput
	// The destination range of outgoing packets that this route applies to.
	// Only IPv4 is supported.
	DestRange pulumi.StringPtrInput
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the
	// last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The network that this route applies to.
	//
	// ***
	Network pulumi.StringPtrInput
	// URL to a gateway that should handle matching packets.
	// Currently, you can only specify the internet gateway, using a full or
	// partial valid URL:
	// * `https://www.googleapis.com/compute/v1/projects/project/global/gateways/default-internet-gateway`
	// * `projects/project/global/gateways/default-internet-gateway`
	// * `global/gateways/default-internet-gateway`
	// * The string `default-internet-gateway`.
	NextHopGateway pulumi.StringPtrInput
	// The IP address or URL to a forwarding rule of type
	// loadBalancingScheme=INTERNAL that should handle matching
	// packets.
	// With the GA provider you can only specify the forwarding
	// rule as a partial or full URL. For example, the following
	// are all valid values:
	// * 10.128.0.56
	// * https://www.googleapis.com/compute/v1/projects/project/regions/region/forwardingRules/forwardingRule
	// * regions/region/forwardingRules/forwardingRule
	//   When the beta provider, you can also specify the IP address
	//   of a forwarding rule from the same VPC or any peered VPC.
	//   Note that this can only be used when the destinationRange is
	//   a public (non-RFC 1918) IP CIDR range.
	NextHopIlb pulumi.StringPtrInput
	// URL to an instance that should handle matching packets.
	// You can specify this as a full or partial URL. For example:
	// * `https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/instance`
	// * `projects/project/zones/zone/instances/instance`
	// * `zones/zone/instances/instance`
	// * Just the instance name, with the zone in `nextHopInstanceZone`.
	NextHopInstance pulumi.StringPtrInput
	// (Optional when `nextHopInstance` is
	// specified)  The zone of the instance specified in
	// `nextHopInstance`.  Omit if `nextHopInstance` is specified as
	// a URL.
	NextHopInstanceZone pulumi.StringPtrInput
	// Network IP address of an instance that should handle matching packets.
	NextHopIp pulumi.StringPtrInput
	// URL to a Network that should handle matching packets.
	NextHopNetwork pulumi.StringPtrInput
	// URL to a VpnTunnel that should handle matching packets.
	NextHopVpnTunnel pulumi.StringPtrInput
	// The priority of this route. Priority is used to break ties in cases
	// where there is more than one matching route of equal prefix length.
	// In the case of two routes with equal prefix length, the one with the
	// lowest-numbered priority value wins.
	// Default value is 1000. Valid range is 0 through 65535.
	Priority pulumi.IntPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// A list of instance tags to which this route applies.
	Tags pulumi.StringArrayInput
}

func (RouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeState)(nil)).Elem()
}

type routeArgs struct {
	// An optional description of this resource. Provide this property
	// when you create the resource.
	Description *string `pulumi:"description"`
	// The destination range of outgoing packets that this route applies to.
	// Only IPv4 is supported.
	DestRange string `pulumi:"destRange"`
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the
	// last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// The network that this route applies to.
	//
	// ***
	Network string `pulumi:"network"`
	// URL to a gateway that should handle matching packets.
	// Currently, you can only specify the internet gateway, using a full or
	// partial valid URL:
	// * `https://www.googleapis.com/compute/v1/projects/project/global/gateways/default-internet-gateway`
	// * `projects/project/global/gateways/default-internet-gateway`
	// * `global/gateways/default-internet-gateway`
	// * The string `default-internet-gateway`.
	NextHopGateway *string `pulumi:"nextHopGateway"`
	// The IP address or URL to a forwarding rule of type
	// loadBalancingScheme=INTERNAL that should handle matching
	// packets.
	// With the GA provider you can only specify the forwarding
	// rule as a partial or full URL. For example, the following
	// are all valid values:
	// * 10.128.0.56
	// * https://www.googleapis.com/compute/v1/projects/project/regions/region/forwardingRules/forwardingRule
	// * regions/region/forwardingRules/forwardingRule
	//   When the beta provider, you can also specify the IP address
	//   of a forwarding rule from the same VPC or any peered VPC.
	//   Note that this can only be used when the destinationRange is
	//   a public (non-RFC 1918) IP CIDR range.
	NextHopIlb *string `pulumi:"nextHopIlb"`
	// URL to an instance that should handle matching packets.
	// You can specify this as a full or partial URL. For example:
	// * `https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/instance`
	// * `projects/project/zones/zone/instances/instance`
	// * `zones/zone/instances/instance`
	// * Just the instance name, with the zone in `nextHopInstanceZone`.
	NextHopInstance *string `pulumi:"nextHopInstance"`
	// (Optional when `nextHopInstance` is
	// specified)  The zone of the instance specified in
	// `nextHopInstance`.  Omit if `nextHopInstance` is specified as
	// a URL.
	NextHopInstanceZone *string `pulumi:"nextHopInstanceZone"`
	// Network IP address of an instance that should handle matching packets.
	NextHopIp *string `pulumi:"nextHopIp"`
	// URL to a VpnTunnel that should handle matching packets.
	NextHopVpnTunnel *string `pulumi:"nextHopVpnTunnel"`
	// The priority of this route. Priority is used to break ties in cases
	// where there is more than one matching route of equal prefix length.
	// In the case of two routes with equal prefix length, the one with the
	// lowest-numbered priority value wins.
	// Default value is 1000. Valid range is 0 through 65535.
	Priority *int `pulumi:"priority"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// A list of instance tags to which this route applies.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a Route resource.
type RouteArgs struct {
	// An optional description of this resource. Provide this property
	// when you create the resource.
	Description pulumi.StringPtrInput
	// The destination range of outgoing packets that this route applies to.
	// Only IPv4 is supported.
	DestRange pulumi.StringInput
	// Name of the resource. Provided by the client when the resource is
	// created. The name must be 1-63 characters long, and comply with
	// RFC1035.  Specifically, the name must be 1-63 characters long and
	// match the regular expression `a-z?` which means
	// the first character must be a lowercase letter, and all following
	// characters must be a dash, lowercase letter, or digit, except the
	// last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// The network that this route applies to.
	//
	// ***
	Network pulumi.StringInput
	// URL to a gateway that should handle matching packets.
	// Currently, you can only specify the internet gateway, using a full or
	// partial valid URL:
	// * `https://www.googleapis.com/compute/v1/projects/project/global/gateways/default-internet-gateway`
	// * `projects/project/global/gateways/default-internet-gateway`
	// * `global/gateways/default-internet-gateway`
	// * The string `default-internet-gateway`.
	NextHopGateway pulumi.StringPtrInput
	// The IP address or URL to a forwarding rule of type
	// loadBalancingScheme=INTERNAL that should handle matching
	// packets.
	// With the GA provider you can only specify the forwarding
	// rule as a partial or full URL. For example, the following
	// are all valid values:
	// * 10.128.0.56
	// * https://www.googleapis.com/compute/v1/projects/project/regions/region/forwardingRules/forwardingRule
	// * regions/region/forwardingRules/forwardingRule
	//   When the beta provider, you can also specify the IP address
	//   of a forwarding rule from the same VPC or any peered VPC.
	//   Note that this can only be used when the destinationRange is
	//   a public (non-RFC 1918) IP CIDR range.
	NextHopIlb pulumi.StringPtrInput
	// URL to an instance that should handle matching packets.
	// You can specify this as a full or partial URL. For example:
	// * `https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/instance`
	// * `projects/project/zones/zone/instances/instance`
	// * `zones/zone/instances/instance`
	// * Just the instance name, with the zone in `nextHopInstanceZone`.
	NextHopInstance pulumi.StringPtrInput
	// (Optional when `nextHopInstance` is
	// specified)  The zone of the instance specified in
	// `nextHopInstance`.  Omit if `nextHopInstance` is specified as
	// a URL.
	NextHopInstanceZone pulumi.StringPtrInput
	// Network IP address of an instance that should handle matching packets.
	NextHopIp pulumi.StringPtrInput
	// URL to a VpnTunnel that should handle matching packets.
	NextHopVpnTunnel pulumi.StringPtrInput
	// The priority of this route. Priority is used to break ties in cases
	// where there is more than one matching route of equal prefix length.
	// In the case of two routes with equal prefix length, the one with the
	// lowest-numbered priority value wins.
	// Default value is 1000. Valid range is 0 through 65535.
	Priority pulumi.IntPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// A list of instance tags to which this route applies.
	Tags pulumi.StringArrayInput
}

func (RouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeArgs)(nil)).Elem()
}

type RouteInput interface {
	pulumi.Input

	ToRouteOutput() RouteOutput
	ToRouteOutputWithContext(ctx context.Context) RouteOutput
}

func (*Route) ElementType() reflect.Type {
	return reflect.TypeOf((**Route)(nil)).Elem()
}

func (i *Route) ToRouteOutput() RouteOutput {
	return i.ToRouteOutputWithContext(context.Background())
}

func (i *Route) ToRouteOutputWithContext(ctx context.Context) RouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteOutput)
}

func (i *Route) ToOutput(ctx context.Context) pulumix.Output[*Route] {
	return pulumix.Output[*Route]{
		OutputState: i.ToRouteOutputWithContext(ctx).OutputState,
	}
}

// RouteArrayInput is an input type that accepts RouteArray and RouteArrayOutput values.
// You can construct a concrete instance of `RouteArrayInput` via:
//
//	RouteArray{ RouteArgs{...} }
type RouteArrayInput interface {
	pulumi.Input

	ToRouteArrayOutput() RouteArrayOutput
	ToRouteArrayOutputWithContext(context.Context) RouteArrayOutput
}

type RouteArray []RouteInput

func (RouteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Route)(nil)).Elem()
}

func (i RouteArray) ToRouteArrayOutput() RouteArrayOutput {
	return i.ToRouteArrayOutputWithContext(context.Background())
}

func (i RouteArray) ToRouteArrayOutputWithContext(ctx context.Context) RouteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteArrayOutput)
}

func (i RouteArray) ToOutput(ctx context.Context) pulumix.Output[[]*Route] {
	return pulumix.Output[[]*Route]{
		OutputState: i.ToRouteArrayOutputWithContext(ctx).OutputState,
	}
}

// RouteMapInput is an input type that accepts RouteMap and RouteMapOutput values.
// You can construct a concrete instance of `RouteMapInput` via:
//
//	RouteMap{ "key": RouteArgs{...} }
type RouteMapInput interface {
	pulumi.Input

	ToRouteMapOutput() RouteMapOutput
	ToRouteMapOutputWithContext(context.Context) RouteMapOutput
}

type RouteMap map[string]RouteInput

func (RouteMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Route)(nil)).Elem()
}

func (i RouteMap) ToRouteMapOutput() RouteMapOutput {
	return i.ToRouteMapOutputWithContext(context.Background())
}

func (i RouteMap) ToRouteMapOutputWithContext(ctx context.Context) RouteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteMapOutput)
}

func (i RouteMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Route] {
	return pulumix.Output[map[string]*Route]{
		OutputState: i.ToRouteMapOutputWithContext(ctx).OutputState,
	}
}

type RouteOutput struct{ *pulumi.OutputState }

func (RouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Route)(nil)).Elem()
}

func (o RouteOutput) ToRouteOutput() RouteOutput {
	return o
}

func (o RouteOutput) ToRouteOutputWithContext(ctx context.Context) RouteOutput {
	return o
}

func (o RouteOutput) ToOutput(ctx context.Context) pulumix.Output[*Route] {
	return pulumix.Output[*Route]{
		OutputState: o.OutputState,
	}
}

// An optional description of this resource. Provide this property
// when you create the resource.
func (o RouteOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The destination range of outgoing packets that this route applies to.
// Only IPv4 is supported.
func (o RouteOutput) DestRange() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.DestRange }).(pulumi.StringOutput)
}

// Name of the resource. Provided by the client when the resource is
// created. The name must be 1-63 characters long, and comply with
// RFC1035.  Specifically, the name must be 1-63 characters long and
// match the regular expression `a-z?` which means
// the first character must be a lowercase letter, and all following
// characters must be a dash, lowercase letter, or digit, except the
// last character, which cannot be a dash.
func (o RouteOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The network that this route applies to.
//
// ***
func (o RouteOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.Network }).(pulumi.StringOutput)
}

// URL to a gateway that should handle matching packets.
// Currently, you can only specify the internet gateway, using a full or
// partial valid URL:
// * `https://www.googleapis.com/compute/v1/projects/project/global/gateways/default-internet-gateway`
// * `projects/project/global/gateways/default-internet-gateway`
// * `global/gateways/default-internet-gateway`
// * The string `default-internet-gateway`.
func (o RouteOutput) NextHopGateway() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.NextHopGateway }).(pulumi.StringPtrOutput)
}

// The IP address or URL to a forwarding rule of type
// loadBalancingScheme=INTERNAL that should handle matching
// packets.
// With the GA provider you can only specify the forwarding
// rule as a partial or full URL. For example, the following
// are all valid values:
//   - 10.128.0.56
//   - https://www.googleapis.com/compute/v1/projects/project/regions/region/forwardingRules/forwardingRule
//   - regions/region/forwardingRules/forwardingRule
//     When the beta provider, you can also specify the IP address
//     of a forwarding rule from the same VPC or any peered VPC.
//     Note that this can only be used when the destinationRange is
//     a public (non-RFC 1918) IP CIDR range.
func (o RouteOutput) NextHopIlb() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.NextHopIlb }).(pulumi.StringPtrOutput)
}

// URL to an instance that should handle matching packets.
// You can specify this as a full or partial URL. For example:
// * `https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/instance`
// * `projects/project/zones/zone/instances/instance`
// * `zones/zone/instances/instance`
// * Just the instance name, with the zone in `nextHopInstanceZone`.
func (o RouteOutput) NextHopInstance() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.NextHopInstance }).(pulumi.StringPtrOutput)
}

// (Optional when `nextHopInstance` is
// specified)  The zone of the instance specified in
// `nextHopInstance`.  Omit if `nextHopInstance` is specified as
// a URL.
func (o RouteOutput) NextHopInstanceZone() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.NextHopInstanceZone }).(pulumi.StringOutput)
}

// Network IP address of an instance that should handle matching packets.
func (o RouteOutput) NextHopIp() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.NextHopIp }).(pulumi.StringOutput)
}

// URL to a Network that should handle matching packets.
func (o RouteOutput) NextHopNetwork() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.NextHopNetwork }).(pulumi.StringOutput)
}

// URL to a VpnTunnel that should handle matching packets.
func (o RouteOutput) NextHopVpnTunnel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.StringPtrOutput { return v.NextHopVpnTunnel }).(pulumi.StringPtrOutput)
}

// The priority of this route. Priority is used to break ties in cases
// where there is more than one matching route of equal prefix length.
// In the case of two routes with equal prefix length, the one with the
// lowest-numbered priority value wins.
// Default value is 1000. Valid range is 0 through 65535.
func (o RouteOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Route) pulumi.IntPtrOutput { return v.Priority }).(pulumi.IntPtrOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o RouteOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The URI of the created resource.
func (o RouteOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *Route) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// A list of instance tags to which this route applies.
func (o RouteOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Route) pulumi.StringArrayOutput { return v.Tags }).(pulumi.StringArrayOutput)
}

type RouteArrayOutput struct{ *pulumi.OutputState }

func (RouteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Route)(nil)).Elem()
}

func (o RouteArrayOutput) ToRouteArrayOutput() RouteArrayOutput {
	return o
}

func (o RouteArrayOutput) ToRouteArrayOutputWithContext(ctx context.Context) RouteArrayOutput {
	return o
}

func (o RouteArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Route] {
	return pulumix.Output[[]*Route]{
		OutputState: o.OutputState,
	}
}

func (o RouteArrayOutput) Index(i pulumi.IntInput) RouteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Route {
		return vs[0].([]*Route)[vs[1].(int)]
	}).(RouteOutput)
}

type RouteMapOutput struct{ *pulumi.OutputState }

func (RouteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Route)(nil)).Elem()
}

func (o RouteMapOutput) ToRouteMapOutput() RouteMapOutput {
	return o
}

func (o RouteMapOutput) ToRouteMapOutputWithContext(ctx context.Context) RouteMapOutput {
	return o
}

func (o RouteMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Route] {
	return pulumix.Output[map[string]*Route]{
		OutputState: o.OutputState,
	}
}

func (o RouteMapOutput) MapIndex(k pulumi.StringInput) RouteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Route {
		return vs[0].(map[string]*Route)[vs[1].(string)]
	}).(RouteOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouteInput)(nil)).Elem(), &Route{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouteArrayInput)(nil)).Elem(), RouteArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouteMapInput)(nil)).Elem(), RouteMap{})
	pulumi.RegisterOutputType(RouteOutput{})
	pulumi.RegisterOutputType(RouteArrayOutput{})
	pulumi.RegisterOutputType(RouteMapOutput{})
}
