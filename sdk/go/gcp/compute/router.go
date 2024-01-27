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

// Represents a Router resource.
//
// To get more information about Router, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/routers)
// * How-to Guides
//   - [Google Cloud Router](https://cloud.google.com/router/docs/)
//
// ## Example Usage
// ### Router Basic
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
//			foobarNetwork, err := compute.NewNetwork(ctx, "foobarNetwork", &compute.NetworkArgs{
//				AutoCreateSubnetworks: pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewRouter(ctx, "foobarRouter", &compute.RouterArgs{
//				Network: foobarNetwork.Name,
//				Bgp: &compute.RouterBgpArgs{
//					Asn:           pulumi.Int(64514),
//					AdvertiseMode: pulumi.String("CUSTOM"),
//					AdvertisedGroups: pulumi.StringArray{
//						pulumi.String("ALL_SUBNETS"),
//					},
//					AdvertisedIpRanges: compute.RouterBgpAdvertisedIpRangeArray{
//						&compute.RouterBgpAdvertisedIpRangeArgs{
//							Range: pulumi.String("1.2.3.4"),
//						},
//						&compute.RouterBgpAdvertisedIpRangeArgs{
//							Range: pulumi.String("6.7.0.0/16"),
//						},
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
// ### Compute Router Encrypted Interconnect
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
//			network, err := compute.NewNetwork(ctx, "network", &compute.NetworkArgs{
//				AutoCreateSubnetworks: pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = compute.NewRouter(ctx, "encrypted-interconnect-router", &compute.RouterArgs{
//				Network:                     network.Name,
//				EncryptedInterconnectRouter: pulumi.Bool(true),
//				Bgp: &compute.RouterBgpArgs{
//					Asn: pulumi.Int(64514),
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
// Router can be imported using any of these accepted formats* `projects/{{project}}/regions/{{region}}/routers/{{name}}` * `{{project}}/{{region}}/{{name}}` * `{{region}}/{{name}}` * `{{name}}` When using the `pulumi import` command, Router can be imported using one of the formats above. For example
//
// ```sh
//
//	$ pulumi import gcp:compute/router:Router default projects/{{project}}/regions/{{region}}/routers/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/router:Router default {{project}}/{{region}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/router:Router default {{region}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:compute/router:Router default {{name}}
//
// ```
type Router struct {
	pulumi.CustomResourceState

	// BGP information specific to this router.
	// Structure is documented below.
	Bgp RouterBgpPtrOutput `pulumi:"bgp"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Indicates if a router is dedicated for use with encrypted VLAN
	// attachments (interconnectAttachments).
	EncryptedInterconnectRouter pulumi.BoolPtrOutput `pulumi:"encryptedInterconnectRouter"`
	// Name of the resource. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters
	// long and match the regular expression `a-z?`
	// which means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// A reference to the network to which this router belongs.
	//
	// ***
	Network pulumi.StringOutput `pulumi:"network"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Region where the router resides.
	Region pulumi.StringOutput `pulumi:"region"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
}

// NewRouter registers a new resource with the given unique name, arguments, and options.
func NewRouter(ctx *pulumi.Context,
	name string, args *RouterArgs, opts ...pulumi.ResourceOption) (*Router, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Network == nil {
		return nil, errors.New("invalid value for required argument 'Network'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Router
	err := ctx.RegisterResource("gcp:compute/router:Router", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRouter gets an existing Router resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRouter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouterState, opts ...pulumi.ResourceOption) (*Router, error) {
	var resource Router
	err := ctx.ReadResource("gcp:compute/router:Router", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Router resources.
type routerState struct {
	// BGP information specific to this router.
	// Structure is documented below.
	Bgp *RouterBgp `pulumi:"bgp"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Indicates if a router is dedicated for use with encrypted VLAN
	// attachments (interconnectAttachments).
	EncryptedInterconnectRouter *bool `pulumi:"encryptedInterconnectRouter"`
	// Name of the resource. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters
	// long and match the regular expression `a-z?`
	// which means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// A reference to the network to which this router belongs.
	//
	// ***
	Network *string `pulumi:"network"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Region where the router resides.
	Region *string `pulumi:"region"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
}

type RouterState struct {
	// BGP information specific to this router.
	// Structure is documented below.
	Bgp RouterBgpPtrInput
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Indicates if a router is dedicated for use with encrypted VLAN
	// attachments (interconnectAttachments).
	EncryptedInterconnectRouter pulumi.BoolPtrInput
	// Name of the resource. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters
	// long and match the regular expression `a-z?`
	// which means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// A reference to the network to which this router belongs.
	//
	// ***
	Network pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Region where the router resides.
	Region pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
}

func (RouterState) ElementType() reflect.Type {
	return reflect.TypeOf((*routerState)(nil)).Elem()
}

type routerArgs struct {
	// BGP information specific to this router.
	// Structure is documented below.
	Bgp *RouterBgp `pulumi:"bgp"`
	// An optional description of this resource.
	Description *string `pulumi:"description"`
	// Indicates if a router is dedicated for use with encrypted VLAN
	// attachments (interconnectAttachments).
	EncryptedInterconnectRouter *bool `pulumi:"encryptedInterconnectRouter"`
	// Name of the resource. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters
	// long and match the regular expression `a-z?`
	// which means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// A reference to the network to which this router belongs.
	//
	// ***
	Network string `pulumi:"network"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Region where the router resides.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a Router resource.
type RouterArgs struct {
	// BGP information specific to this router.
	// Structure is documented below.
	Bgp RouterBgpPtrInput
	// An optional description of this resource.
	Description pulumi.StringPtrInput
	// Indicates if a router is dedicated for use with encrypted VLAN
	// attachments (interconnectAttachments).
	EncryptedInterconnectRouter pulumi.BoolPtrInput
	// Name of the resource. The name must be 1-63 characters long, and
	// comply with RFC1035. Specifically, the name must be 1-63 characters
	// long and match the regular expression `a-z?`
	// which means the first character must be a lowercase letter, and all
	// following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// A reference to the network to which this router belongs.
	//
	// ***
	Network pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Region where the router resides.
	Region pulumi.StringPtrInput
}

func (RouterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routerArgs)(nil)).Elem()
}

type RouterInput interface {
	pulumi.Input

	ToRouterOutput() RouterOutput
	ToRouterOutputWithContext(ctx context.Context) RouterOutput
}

func (*Router) ElementType() reflect.Type {
	return reflect.TypeOf((**Router)(nil)).Elem()
}

func (i *Router) ToRouterOutput() RouterOutput {
	return i.ToRouterOutputWithContext(context.Background())
}

func (i *Router) ToRouterOutputWithContext(ctx context.Context) RouterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterOutput)
}

// RouterArrayInput is an input type that accepts RouterArray and RouterArrayOutput values.
// You can construct a concrete instance of `RouterArrayInput` via:
//
//	RouterArray{ RouterArgs{...} }
type RouterArrayInput interface {
	pulumi.Input

	ToRouterArrayOutput() RouterArrayOutput
	ToRouterArrayOutputWithContext(context.Context) RouterArrayOutput
}

type RouterArray []RouterInput

func (RouterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Router)(nil)).Elem()
}

func (i RouterArray) ToRouterArrayOutput() RouterArrayOutput {
	return i.ToRouterArrayOutputWithContext(context.Background())
}

func (i RouterArray) ToRouterArrayOutputWithContext(ctx context.Context) RouterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterArrayOutput)
}

// RouterMapInput is an input type that accepts RouterMap and RouterMapOutput values.
// You can construct a concrete instance of `RouterMapInput` via:
//
//	RouterMap{ "key": RouterArgs{...} }
type RouterMapInput interface {
	pulumi.Input

	ToRouterMapOutput() RouterMapOutput
	ToRouterMapOutputWithContext(context.Context) RouterMapOutput
}

type RouterMap map[string]RouterInput

func (RouterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Router)(nil)).Elem()
}

func (i RouterMap) ToRouterMapOutput() RouterMapOutput {
	return i.ToRouterMapOutputWithContext(context.Background())
}

func (i RouterMap) ToRouterMapOutputWithContext(ctx context.Context) RouterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouterMapOutput)
}

type RouterOutput struct{ *pulumi.OutputState }

func (RouterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Router)(nil)).Elem()
}

func (o RouterOutput) ToRouterOutput() RouterOutput {
	return o
}

func (o RouterOutput) ToRouterOutputWithContext(ctx context.Context) RouterOutput {
	return o
}

// BGP information specific to this router.
// Structure is documented below.
func (o RouterOutput) Bgp() RouterBgpPtrOutput {
	return o.ApplyT(func(v *Router) RouterBgpPtrOutput { return v.Bgp }).(RouterBgpPtrOutput)
}

// Creation timestamp in RFC3339 text format.
func (o RouterOutput) CreationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v *Router) pulumi.StringOutput { return v.CreationTimestamp }).(pulumi.StringOutput)
}

// An optional description of this resource.
func (o RouterOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Router) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Indicates if a router is dedicated for use with encrypted VLAN
// attachments (interconnectAttachments).
func (o RouterOutput) EncryptedInterconnectRouter() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Router) pulumi.BoolPtrOutput { return v.EncryptedInterconnectRouter }).(pulumi.BoolPtrOutput)
}

// Name of the resource. The name must be 1-63 characters long, and
// comply with RFC1035. Specifically, the name must be 1-63 characters
// long and match the regular expression `a-z?`
// which means the first character must be a lowercase letter, and all
// following characters must be a dash, lowercase letter, or digit,
// except the last character, which cannot be a dash.
func (o RouterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Router) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A reference to the network to which this router belongs.
//
// ***
func (o RouterOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v *Router) pulumi.StringOutput { return v.Network }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o RouterOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Router) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Region where the router resides.
func (o RouterOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Router) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The URI of the created resource.
func (o RouterOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *Router) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

type RouterArrayOutput struct{ *pulumi.OutputState }

func (RouterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Router)(nil)).Elem()
}

func (o RouterArrayOutput) ToRouterArrayOutput() RouterArrayOutput {
	return o
}

func (o RouterArrayOutput) ToRouterArrayOutputWithContext(ctx context.Context) RouterArrayOutput {
	return o
}

func (o RouterArrayOutput) Index(i pulumi.IntInput) RouterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Router {
		return vs[0].([]*Router)[vs[1].(int)]
	}).(RouterOutput)
}

type RouterMapOutput struct{ *pulumi.OutputState }

func (RouterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Router)(nil)).Elem()
}

func (o RouterMapOutput) ToRouterMapOutput() RouterMapOutput {
	return o
}

func (o RouterMapOutput) ToRouterMapOutputWithContext(ctx context.Context) RouterMapOutput {
	return o
}

func (o RouterMapOutput) MapIndex(k pulumi.StringInput) RouterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Router {
		return vs[0].(map[string]*Router)[vs[1].(string)]
	}).(RouterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RouterInput)(nil)).Elem(), &Router{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterArrayInput)(nil)).Elem(), RouterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RouterMapInput)(nil)).Elem(), RouterMap{})
	pulumi.RegisterOutputType(RouterOutput{})
	pulumi.RegisterOutputType(RouterArrayOutput{})
	pulumi.RegisterOutputType(RouterMapOutput{})
}
