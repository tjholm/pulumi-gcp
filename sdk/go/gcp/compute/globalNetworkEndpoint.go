// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Global Network endpoint represents a IP address and port combination that exists outside of GCP.
// **NOTE**: Global network endpoints cannot be created outside of a
// global network endpoint group.
//
// To get more information about GlobalNetworkEndpoint, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/networkEndpointGroups)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/load-balancing/docs/negs/)
//
// ## Example Usage
// ### Global Network Endpoint
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		neg, err := compute.NewGlobalNetworkEndpointGroup(ctx, "neg", &compute.GlobalNetworkEndpointGroupArgs{
// 			DefaultPort:         pulumi.Int(90),
// 			NetworkEndpointType: pulumi.String("INTERNET_FQDN_PORT"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = compute.NewGlobalNetworkEndpoint(ctx, "default-endpoint", &compute.GlobalNetworkEndpointArgs{
// 			GlobalNetworkEndpointGroup: neg.Name,
// 			Fqdn:                       pulumi.String("www.example.com"),
// 			Port:                       pulumi.Int(90),
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
// GlobalNetworkEndpoint can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:compute/globalNetworkEndpoint:GlobalNetworkEndpoint default projects/{{project}}/global/networkEndpointGroups/{{global_network_endpoint_group}}/{{ip_address}}/{{fqdn}}/{{port}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/globalNetworkEndpoint:GlobalNetworkEndpoint default {{project}}/{{global_network_endpoint_group}}/{{ip_address}}/{{fqdn}}/{{port}}
// ```
//
// ```sh
//  $ pulumi import gcp:compute/globalNetworkEndpoint:GlobalNetworkEndpoint default {{global_network_endpoint_group}}/{{ip_address}}/{{fqdn}}/{{port}}
// ```
type GlobalNetworkEndpoint struct {
	pulumi.CustomResourceState

	// Fully qualified domain name of network endpoint.
	// This can only be specified when networkEndpointType of the NEG is INTERNET_FQDN_PORT.
	Fqdn pulumi.StringPtrOutput `pulumi:"fqdn"`
	// The global network endpoint group this endpoint is part of.
	GlobalNetworkEndpointGroup pulumi.StringOutput `pulumi:"globalNetworkEndpointGroup"`
	// IPv4 address external endpoint.
	IpAddress pulumi.StringPtrOutput `pulumi:"ipAddress"`
	// Port number of the external endpoint.
	Port pulumi.IntOutput `pulumi:"port"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewGlobalNetworkEndpoint registers a new resource with the given unique name, arguments, and options.
func NewGlobalNetworkEndpoint(ctx *pulumi.Context,
	name string, args *GlobalNetworkEndpointArgs, opts ...pulumi.ResourceOption) (*GlobalNetworkEndpoint, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GlobalNetworkEndpointGroup == nil {
		return nil, errors.New("invalid value for required argument 'GlobalNetworkEndpointGroup'")
	}
	if args.Port == nil {
		return nil, errors.New("invalid value for required argument 'Port'")
	}
	var resource GlobalNetworkEndpoint
	err := ctx.RegisterResource("gcp:compute/globalNetworkEndpoint:GlobalNetworkEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGlobalNetworkEndpoint gets an existing GlobalNetworkEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGlobalNetworkEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GlobalNetworkEndpointState, opts ...pulumi.ResourceOption) (*GlobalNetworkEndpoint, error) {
	var resource GlobalNetworkEndpoint
	err := ctx.ReadResource("gcp:compute/globalNetworkEndpoint:GlobalNetworkEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GlobalNetworkEndpoint resources.
type globalNetworkEndpointState struct {
	// Fully qualified domain name of network endpoint.
	// This can only be specified when networkEndpointType of the NEG is INTERNET_FQDN_PORT.
	Fqdn *string `pulumi:"fqdn"`
	// The global network endpoint group this endpoint is part of.
	GlobalNetworkEndpointGroup *string `pulumi:"globalNetworkEndpointGroup"`
	// IPv4 address external endpoint.
	IpAddress *string `pulumi:"ipAddress"`
	// Port number of the external endpoint.
	Port *int `pulumi:"port"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type GlobalNetworkEndpointState struct {
	// Fully qualified domain name of network endpoint.
	// This can only be specified when networkEndpointType of the NEG is INTERNET_FQDN_PORT.
	Fqdn pulumi.StringPtrInput
	// The global network endpoint group this endpoint is part of.
	GlobalNetworkEndpointGroup pulumi.StringPtrInput
	// IPv4 address external endpoint.
	IpAddress pulumi.StringPtrInput
	// Port number of the external endpoint.
	Port pulumi.IntPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (GlobalNetworkEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*globalNetworkEndpointState)(nil)).Elem()
}

type globalNetworkEndpointArgs struct {
	// Fully qualified domain name of network endpoint.
	// This can only be specified when networkEndpointType of the NEG is INTERNET_FQDN_PORT.
	Fqdn *string `pulumi:"fqdn"`
	// The global network endpoint group this endpoint is part of.
	GlobalNetworkEndpointGroup string `pulumi:"globalNetworkEndpointGroup"`
	// IPv4 address external endpoint.
	IpAddress *string `pulumi:"ipAddress"`
	// Port number of the external endpoint.
	Port int `pulumi:"port"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a GlobalNetworkEndpoint resource.
type GlobalNetworkEndpointArgs struct {
	// Fully qualified domain name of network endpoint.
	// This can only be specified when networkEndpointType of the NEG is INTERNET_FQDN_PORT.
	Fqdn pulumi.StringPtrInput
	// The global network endpoint group this endpoint is part of.
	GlobalNetworkEndpointGroup pulumi.StringInput
	// IPv4 address external endpoint.
	IpAddress pulumi.StringPtrInput
	// Port number of the external endpoint.
	Port pulumi.IntInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (GlobalNetworkEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*globalNetworkEndpointArgs)(nil)).Elem()
}

type GlobalNetworkEndpointInput interface {
	pulumi.Input

	ToGlobalNetworkEndpointOutput() GlobalNetworkEndpointOutput
	ToGlobalNetworkEndpointOutputWithContext(ctx context.Context) GlobalNetworkEndpointOutput
}

func (*GlobalNetworkEndpoint) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalNetworkEndpoint)(nil)).Elem()
}

func (i *GlobalNetworkEndpoint) ToGlobalNetworkEndpointOutput() GlobalNetworkEndpointOutput {
	return i.ToGlobalNetworkEndpointOutputWithContext(context.Background())
}

func (i *GlobalNetworkEndpoint) ToGlobalNetworkEndpointOutputWithContext(ctx context.Context) GlobalNetworkEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalNetworkEndpointOutput)
}

// GlobalNetworkEndpointArrayInput is an input type that accepts GlobalNetworkEndpointArray and GlobalNetworkEndpointArrayOutput values.
// You can construct a concrete instance of `GlobalNetworkEndpointArrayInput` via:
//
//          GlobalNetworkEndpointArray{ GlobalNetworkEndpointArgs{...} }
type GlobalNetworkEndpointArrayInput interface {
	pulumi.Input

	ToGlobalNetworkEndpointArrayOutput() GlobalNetworkEndpointArrayOutput
	ToGlobalNetworkEndpointArrayOutputWithContext(context.Context) GlobalNetworkEndpointArrayOutput
}

type GlobalNetworkEndpointArray []GlobalNetworkEndpointInput

func (GlobalNetworkEndpointArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GlobalNetworkEndpoint)(nil)).Elem()
}

func (i GlobalNetworkEndpointArray) ToGlobalNetworkEndpointArrayOutput() GlobalNetworkEndpointArrayOutput {
	return i.ToGlobalNetworkEndpointArrayOutputWithContext(context.Background())
}

func (i GlobalNetworkEndpointArray) ToGlobalNetworkEndpointArrayOutputWithContext(ctx context.Context) GlobalNetworkEndpointArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalNetworkEndpointArrayOutput)
}

// GlobalNetworkEndpointMapInput is an input type that accepts GlobalNetworkEndpointMap and GlobalNetworkEndpointMapOutput values.
// You can construct a concrete instance of `GlobalNetworkEndpointMapInput` via:
//
//          GlobalNetworkEndpointMap{ "key": GlobalNetworkEndpointArgs{...} }
type GlobalNetworkEndpointMapInput interface {
	pulumi.Input

	ToGlobalNetworkEndpointMapOutput() GlobalNetworkEndpointMapOutput
	ToGlobalNetworkEndpointMapOutputWithContext(context.Context) GlobalNetworkEndpointMapOutput
}

type GlobalNetworkEndpointMap map[string]GlobalNetworkEndpointInput

func (GlobalNetworkEndpointMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GlobalNetworkEndpoint)(nil)).Elem()
}

func (i GlobalNetworkEndpointMap) ToGlobalNetworkEndpointMapOutput() GlobalNetworkEndpointMapOutput {
	return i.ToGlobalNetworkEndpointMapOutputWithContext(context.Background())
}

func (i GlobalNetworkEndpointMap) ToGlobalNetworkEndpointMapOutputWithContext(ctx context.Context) GlobalNetworkEndpointMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalNetworkEndpointMapOutput)
}

type GlobalNetworkEndpointOutput struct{ *pulumi.OutputState }

func (GlobalNetworkEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalNetworkEndpoint)(nil)).Elem()
}

func (o GlobalNetworkEndpointOutput) ToGlobalNetworkEndpointOutput() GlobalNetworkEndpointOutput {
	return o
}

func (o GlobalNetworkEndpointOutput) ToGlobalNetworkEndpointOutputWithContext(ctx context.Context) GlobalNetworkEndpointOutput {
	return o
}

type GlobalNetworkEndpointArrayOutput struct{ *pulumi.OutputState }

func (GlobalNetworkEndpointArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GlobalNetworkEndpoint)(nil)).Elem()
}

func (o GlobalNetworkEndpointArrayOutput) ToGlobalNetworkEndpointArrayOutput() GlobalNetworkEndpointArrayOutput {
	return o
}

func (o GlobalNetworkEndpointArrayOutput) ToGlobalNetworkEndpointArrayOutputWithContext(ctx context.Context) GlobalNetworkEndpointArrayOutput {
	return o
}

func (o GlobalNetworkEndpointArrayOutput) Index(i pulumi.IntInput) GlobalNetworkEndpointOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GlobalNetworkEndpoint {
		return vs[0].([]*GlobalNetworkEndpoint)[vs[1].(int)]
	}).(GlobalNetworkEndpointOutput)
}

type GlobalNetworkEndpointMapOutput struct{ *pulumi.OutputState }

func (GlobalNetworkEndpointMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GlobalNetworkEndpoint)(nil)).Elem()
}

func (o GlobalNetworkEndpointMapOutput) ToGlobalNetworkEndpointMapOutput() GlobalNetworkEndpointMapOutput {
	return o
}

func (o GlobalNetworkEndpointMapOutput) ToGlobalNetworkEndpointMapOutputWithContext(ctx context.Context) GlobalNetworkEndpointMapOutput {
	return o
}

func (o GlobalNetworkEndpointMapOutput) MapIndex(k pulumi.StringInput) GlobalNetworkEndpointOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GlobalNetworkEndpoint {
		return vs[0].(map[string]*GlobalNetworkEndpoint)[vs[1].(string)]
	}).(GlobalNetworkEndpointOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GlobalNetworkEndpointInput)(nil)).Elem(), &GlobalNetworkEndpoint{})
	pulumi.RegisterInputType(reflect.TypeOf((*GlobalNetworkEndpointArrayInput)(nil)).Elem(), GlobalNetworkEndpointArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GlobalNetworkEndpointMapInput)(nil)).Elem(), GlobalNetworkEndpointMap{})
	pulumi.RegisterOutputType(GlobalNetworkEndpointOutput{})
	pulumi.RegisterOutputType(GlobalNetworkEndpointArrayOutput{})
	pulumi.RegisterOutputType(GlobalNetworkEndpointMapOutput{})
}
