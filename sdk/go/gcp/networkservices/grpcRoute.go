// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package networkservices

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// ## Example Usage
// ### Network Services Grpc Route Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/networkservices"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := networkservices.NewGrpcRoute(ctx, "default", &networkservices.GrpcRouteArgs{
//				Labels: pulumi.StringMap{
//					"foo": pulumi.String("bar"),
//				},
//				Description: pulumi.String("my description"),
//				Hostnames: pulumi.StringArray{
//					pulumi.String("example"),
//				},
//				Rules: networkservices.GrpcRouteRuleArray{
//					&networkservices.GrpcRouteRuleArgs{
//						Matches: networkservices.GrpcRouteRuleMatchArray{
//							&networkservices.GrpcRouteRuleMatchArgs{
//								Headers: networkservices.GrpcRouteRuleMatchHeaderArray{
//									&networkservices.GrpcRouteRuleMatchHeaderArgs{
//										Key:   pulumi.String("key"),
//										Value: pulumi.String("value"),
//									},
//								},
//							},
//						},
//						Action: &networkservices.GrpcRouteRuleActionArgs{
//							RetryPolicy: &networkservices.GrpcRouteRuleActionRetryPolicyArgs{
//								RetryConditions: pulumi.StringArray{
//									pulumi.String("cancelled"),
//								},
//								NumRetries: pulumi.Int(1),
//							},
//						},
//					},
//				},
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Network Services Grpc Route Matches And Actions
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/networkservices"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := networkservices.NewGrpcRoute(ctx, "default", &networkservices.GrpcRouteArgs{
//				Labels: pulumi.StringMap{
//					"foo": pulumi.String("bar"),
//				},
//				Description: pulumi.String("my description"),
//				Hostnames: pulumi.StringArray{
//					pulumi.String("example"),
//				},
//				Rules: networkservices.GrpcRouteRuleArray{
//					&networkservices.GrpcRouteRuleArgs{
//						Matches: networkservices.GrpcRouteRuleMatchArray{
//							&networkservices.GrpcRouteRuleMatchArgs{
//								Headers: networkservices.GrpcRouteRuleMatchHeaderArray{
//									&networkservices.GrpcRouteRuleMatchHeaderArgs{
//										Key:   pulumi.String("key"),
//										Value: pulumi.String("value"),
//									},
//								},
//							},
//							&networkservices.GrpcRouteRuleMatchArgs{
//								Headers: networkservices.GrpcRouteRuleMatchHeaderArray{
//									&networkservices.GrpcRouteRuleMatchHeaderArgs{
//										Key:   pulumi.String("key"),
//										Value: pulumi.String("value"),
//									},
//								},
//								Method: &networkservices.GrpcRouteRuleMatchMethodArgs{
//									GrpcService:   pulumi.String("foo"),
//									GrpcMethod:    pulumi.String("bar"),
//									CaseSensitive: pulumi.Bool(true),
//								},
//							},
//						},
//						Action: &networkservices.GrpcRouteRuleActionArgs{
//							FaultInjectionPolicy: &networkservices.GrpcRouteRuleActionFaultInjectionPolicyArgs{
//								Delay: &networkservices.GrpcRouteRuleActionFaultInjectionPolicyDelayArgs{
//									FixedDelay: pulumi.String("1s"),
//									Percentage: pulumi.Int(1),
//								},
//								Abort: &networkservices.GrpcRouteRuleActionFaultInjectionPolicyAbortArgs{
//									HttpStatus: pulumi.Int(500),
//									Percentage: pulumi.Int(1),
//								},
//							},
//							RetryPolicy: &networkservices.GrpcRouteRuleActionRetryPolicyArgs{
//								RetryConditions: pulumi.StringArray{
//									pulumi.String("cancelled"),
//								},
//								NumRetries: pulumi.Int(1),
//							},
//						},
//					},
//				},
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Network Services Grpc Route Actions
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/networkservices"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := networkservices.NewGrpcRoute(ctx, "default", &networkservices.GrpcRouteArgs{
//				Labels: pulumi.StringMap{
//					"foo": pulumi.String("bar"),
//				},
//				Description: pulumi.String("my description"),
//				Hostnames: pulumi.StringArray{
//					pulumi.String("example"),
//				},
//				Rules: networkservices.GrpcRouteRuleArray{
//					&networkservices.GrpcRouteRuleArgs{
//						Action: &networkservices.GrpcRouteRuleActionArgs{
//							FaultInjectionPolicy: &networkservices.GrpcRouteRuleActionFaultInjectionPolicyArgs{
//								Delay: &networkservices.GrpcRouteRuleActionFaultInjectionPolicyDelayArgs{
//									FixedDelay: pulumi.String("1s"),
//									Percentage: pulumi.Int(1),
//								},
//								Abort: &networkservices.GrpcRouteRuleActionFaultInjectionPolicyAbortArgs{
//									HttpStatus: pulumi.Int(500),
//									Percentage: pulumi.Int(1),
//								},
//							},
//							RetryPolicy: &networkservices.GrpcRouteRuleActionRetryPolicyArgs{
//								RetryConditions: pulumi.StringArray{
//									pulumi.String("cancelled"),
//								},
//								NumRetries: pulumi.Int(1),
//							},
//						},
//					},
//				},
//			}, pulumi.Provider(google_beta))
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
// # GrpcRoute can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:networkservices/grpcRoute:GrpcRoute default projects/{{project}}/locations/global/grpcRoutes/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:networkservices/grpcRoute:GrpcRoute default {{project}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:networkservices/grpcRoute:GrpcRoute default {{name}}
//
// ```
type GrpcRoute struct {
	pulumi.CustomResourceState

	// Time the GrpcRoute was created in UTC.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// A free-text description of the resource. Max length 1024 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// List of gateways this GrpcRoute is attached to, as one of the routing rules to route the requests served by the gateway.
	Gateways pulumi.StringArrayOutput `pulumi:"gateways"`
	// Required. Service hostnames with an optional port for which this route describes traffic.
	Hostnames pulumi.StringArrayOutput `pulumi:"hostnames"`
	// Set of label tags associated with the GrpcRoute resource.
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// List of meshes this GrpcRoute is attached to, as one of the routing rules to route the requests served by the mesh.
	Meshes pulumi.StringArrayOutput `pulumi:"meshes"`
	// Name of the GrpcRoute resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	PulumiLabels pulumi.StringMapOutput `pulumi:"pulumiLabels"`
	// Rules that define how traffic is routed and handled.
	// Structure is documented below.
	Rules GrpcRouteRuleArrayOutput `pulumi:"rules"`
	// Server-defined URL of this resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Time the GrpcRoute was updated in UTC.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewGrpcRoute registers a new resource with the given unique name, arguments, and options.
func NewGrpcRoute(ctx *pulumi.Context,
	name string, args *GrpcRouteArgs, opts ...pulumi.ResourceOption) (*GrpcRoute, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Hostnames == nil {
		return nil, errors.New("invalid value for required argument 'Hostnames'")
	}
	if args.Rules == nil {
		return nil, errors.New("invalid value for required argument 'Rules'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"effectiveLabels",
		"pulumiLabels",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource GrpcRoute
	err := ctx.RegisterResource("gcp:networkservices/grpcRoute:GrpcRoute", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGrpcRoute gets an existing GrpcRoute resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGrpcRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GrpcRouteState, opts ...pulumi.ResourceOption) (*GrpcRoute, error) {
	var resource GrpcRoute
	err := ctx.ReadResource("gcp:networkservices/grpcRoute:GrpcRoute", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GrpcRoute resources.
type grpcRouteState struct {
	// Time the GrpcRoute was created in UTC.
	CreateTime *string `pulumi:"createTime"`
	// A free-text description of the resource. Max length 1024 characters.
	Description *string `pulumi:"description"`
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// List of gateways this GrpcRoute is attached to, as one of the routing rules to route the requests served by the gateway.
	Gateways []string `pulumi:"gateways"`
	// Required. Service hostnames with an optional port for which this route describes traffic.
	Hostnames []string `pulumi:"hostnames"`
	// Set of label tags associated with the GrpcRoute resource.
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// List of meshes this GrpcRoute is attached to, as one of the routing rules to route the requests served by the mesh.
	Meshes []string `pulumi:"meshes"`
	// Name of the GrpcRoute resource.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	PulumiLabels map[string]string `pulumi:"pulumiLabels"`
	// Rules that define how traffic is routed and handled.
	// Structure is documented below.
	Rules []GrpcRouteRule `pulumi:"rules"`
	// Server-defined URL of this resource.
	SelfLink *string `pulumi:"selfLink"`
	// Time the GrpcRoute was updated in UTC.
	UpdateTime *string `pulumi:"updateTime"`
}

type GrpcRouteState struct {
	// Time the GrpcRoute was created in UTC.
	CreateTime pulumi.StringPtrInput
	// A free-text description of the resource. Max length 1024 characters.
	Description pulumi.StringPtrInput
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
	EffectiveLabels pulumi.StringMapInput
	// List of gateways this GrpcRoute is attached to, as one of the routing rules to route the requests served by the gateway.
	Gateways pulumi.StringArrayInput
	// Required. Service hostnames with an optional port for which this route describes traffic.
	Hostnames pulumi.StringArrayInput
	// Set of label tags associated with the GrpcRoute resource.
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// List of meshes this GrpcRoute is attached to, as one of the routing rules to route the requests served by the mesh.
	Meshes pulumi.StringArrayInput
	// Name of the GrpcRoute resource.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	PulumiLabels pulumi.StringMapInput
	// Rules that define how traffic is routed and handled.
	// Structure is documented below.
	Rules GrpcRouteRuleArrayInput
	// Server-defined URL of this resource.
	SelfLink pulumi.StringPtrInput
	// Time the GrpcRoute was updated in UTC.
	UpdateTime pulumi.StringPtrInput
}

func (GrpcRouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*grpcRouteState)(nil)).Elem()
}

type grpcRouteArgs struct {
	// A free-text description of the resource. Max length 1024 characters.
	Description *string `pulumi:"description"`
	// List of gateways this GrpcRoute is attached to, as one of the routing rules to route the requests served by the gateway.
	Gateways []string `pulumi:"gateways"`
	// Required. Service hostnames with an optional port for which this route describes traffic.
	Hostnames []string `pulumi:"hostnames"`
	// Set of label tags associated with the GrpcRoute resource.
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// List of meshes this GrpcRoute is attached to, as one of the routing rules to route the requests served by the mesh.
	Meshes []string `pulumi:"meshes"`
	// Name of the GrpcRoute resource.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Rules that define how traffic is routed and handled.
	// Structure is documented below.
	Rules []GrpcRouteRule `pulumi:"rules"`
}

// The set of arguments for constructing a GrpcRoute resource.
type GrpcRouteArgs struct {
	// A free-text description of the resource. Max length 1024 characters.
	Description pulumi.StringPtrInput
	// List of gateways this GrpcRoute is attached to, as one of the routing rules to route the requests served by the gateway.
	Gateways pulumi.StringArrayInput
	// Required. Service hostnames with an optional port for which this route describes traffic.
	Hostnames pulumi.StringArrayInput
	// Set of label tags associated with the GrpcRoute resource.
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// List of meshes this GrpcRoute is attached to, as one of the routing rules to route the requests served by the mesh.
	Meshes pulumi.StringArrayInput
	// Name of the GrpcRoute resource.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Rules that define how traffic is routed and handled.
	// Structure is documented below.
	Rules GrpcRouteRuleArrayInput
}

func (GrpcRouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*grpcRouteArgs)(nil)).Elem()
}

type GrpcRouteInput interface {
	pulumi.Input

	ToGrpcRouteOutput() GrpcRouteOutput
	ToGrpcRouteOutputWithContext(ctx context.Context) GrpcRouteOutput
}

func (*GrpcRoute) ElementType() reflect.Type {
	return reflect.TypeOf((**GrpcRoute)(nil)).Elem()
}

func (i *GrpcRoute) ToGrpcRouteOutput() GrpcRouteOutput {
	return i.ToGrpcRouteOutputWithContext(context.Background())
}

func (i *GrpcRoute) ToGrpcRouteOutputWithContext(ctx context.Context) GrpcRouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GrpcRouteOutput)
}

func (i *GrpcRoute) ToOutput(ctx context.Context) pulumix.Output[*GrpcRoute] {
	return pulumix.Output[*GrpcRoute]{
		OutputState: i.ToGrpcRouteOutputWithContext(ctx).OutputState,
	}
}

// GrpcRouteArrayInput is an input type that accepts GrpcRouteArray and GrpcRouteArrayOutput values.
// You can construct a concrete instance of `GrpcRouteArrayInput` via:
//
//	GrpcRouteArray{ GrpcRouteArgs{...} }
type GrpcRouteArrayInput interface {
	pulumi.Input

	ToGrpcRouteArrayOutput() GrpcRouteArrayOutput
	ToGrpcRouteArrayOutputWithContext(context.Context) GrpcRouteArrayOutput
}

type GrpcRouteArray []GrpcRouteInput

func (GrpcRouteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GrpcRoute)(nil)).Elem()
}

func (i GrpcRouteArray) ToGrpcRouteArrayOutput() GrpcRouteArrayOutput {
	return i.ToGrpcRouteArrayOutputWithContext(context.Background())
}

func (i GrpcRouteArray) ToGrpcRouteArrayOutputWithContext(ctx context.Context) GrpcRouteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GrpcRouteArrayOutput)
}

func (i GrpcRouteArray) ToOutput(ctx context.Context) pulumix.Output[[]*GrpcRoute] {
	return pulumix.Output[[]*GrpcRoute]{
		OutputState: i.ToGrpcRouteArrayOutputWithContext(ctx).OutputState,
	}
}

// GrpcRouteMapInput is an input type that accepts GrpcRouteMap and GrpcRouteMapOutput values.
// You can construct a concrete instance of `GrpcRouteMapInput` via:
//
//	GrpcRouteMap{ "key": GrpcRouteArgs{...} }
type GrpcRouteMapInput interface {
	pulumi.Input

	ToGrpcRouteMapOutput() GrpcRouteMapOutput
	ToGrpcRouteMapOutputWithContext(context.Context) GrpcRouteMapOutput
}

type GrpcRouteMap map[string]GrpcRouteInput

func (GrpcRouteMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GrpcRoute)(nil)).Elem()
}

func (i GrpcRouteMap) ToGrpcRouteMapOutput() GrpcRouteMapOutput {
	return i.ToGrpcRouteMapOutputWithContext(context.Background())
}

func (i GrpcRouteMap) ToGrpcRouteMapOutputWithContext(ctx context.Context) GrpcRouteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GrpcRouteMapOutput)
}

func (i GrpcRouteMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*GrpcRoute] {
	return pulumix.Output[map[string]*GrpcRoute]{
		OutputState: i.ToGrpcRouteMapOutputWithContext(ctx).OutputState,
	}
}

type GrpcRouteOutput struct{ *pulumi.OutputState }

func (GrpcRouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GrpcRoute)(nil)).Elem()
}

func (o GrpcRouteOutput) ToGrpcRouteOutput() GrpcRouteOutput {
	return o
}

func (o GrpcRouteOutput) ToGrpcRouteOutputWithContext(ctx context.Context) GrpcRouteOutput {
	return o
}

func (o GrpcRouteOutput) ToOutput(ctx context.Context) pulumix.Output[*GrpcRoute] {
	return pulumix.Output[*GrpcRoute]{
		OutputState: o.OutputState,
	}
}

// Time the GrpcRoute was created in UTC.
func (o GrpcRouteOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *GrpcRoute) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// A free-text description of the resource. Max length 1024 characters.
func (o GrpcRouteOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GrpcRoute) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
func (o GrpcRouteOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GrpcRoute) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// List of gateways this GrpcRoute is attached to, as one of the routing rules to route the requests served by the gateway.
func (o GrpcRouteOutput) Gateways() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GrpcRoute) pulumi.StringArrayOutput { return v.Gateways }).(pulumi.StringArrayOutput)
}

// Required. Service hostnames with an optional port for which this route describes traffic.
func (o GrpcRouteOutput) Hostnames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GrpcRoute) pulumi.StringArrayOutput { return v.Hostnames }).(pulumi.StringArrayOutput)
}

// Set of label tags associated with the GrpcRoute resource.
// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
func (o GrpcRouteOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GrpcRoute) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// List of meshes this GrpcRoute is attached to, as one of the routing rules to route the requests served by the mesh.
func (o GrpcRouteOutput) Meshes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GrpcRoute) pulumi.StringArrayOutput { return v.Meshes }).(pulumi.StringArrayOutput)
}

// Name of the GrpcRoute resource.
func (o GrpcRouteOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GrpcRoute) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o GrpcRouteOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *GrpcRoute) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource
// and default labels configured on the provider.
func (o GrpcRouteOutput) PulumiLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GrpcRoute) pulumi.StringMapOutput { return v.PulumiLabels }).(pulumi.StringMapOutput)
}

// Rules that define how traffic is routed and handled.
// Structure is documented below.
func (o GrpcRouteOutput) Rules() GrpcRouteRuleArrayOutput {
	return o.ApplyT(func(v *GrpcRoute) GrpcRouteRuleArrayOutput { return v.Rules }).(GrpcRouteRuleArrayOutput)
}

// Server-defined URL of this resource.
func (o GrpcRouteOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *GrpcRoute) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// Time the GrpcRoute was updated in UTC.
func (o GrpcRouteOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *GrpcRoute) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

type GrpcRouteArrayOutput struct{ *pulumi.OutputState }

func (GrpcRouteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GrpcRoute)(nil)).Elem()
}

func (o GrpcRouteArrayOutput) ToGrpcRouteArrayOutput() GrpcRouteArrayOutput {
	return o
}

func (o GrpcRouteArrayOutput) ToGrpcRouteArrayOutputWithContext(ctx context.Context) GrpcRouteArrayOutput {
	return o
}

func (o GrpcRouteArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*GrpcRoute] {
	return pulumix.Output[[]*GrpcRoute]{
		OutputState: o.OutputState,
	}
}

func (o GrpcRouteArrayOutput) Index(i pulumi.IntInput) GrpcRouteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GrpcRoute {
		return vs[0].([]*GrpcRoute)[vs[1].(int)]
	}).(GrpcRouteOutput)
}

type GrpcRouteMapOutput struct{ *pulumi.OutputState }

func (GrpcRouteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GrpcRoute)(nil)).Elem()
}

func (o GrpcRouteMapOutput) ToGrpcRouteMapOutput() GrpcRouteMapOutput {
	return o
}

func (o GrpcRouteMapOutput) ToGrpcRouteMapOutputWithContext(ctx context.Context) GrpcRouteMapOutput {
	return o
}

func (o GrpcRouteMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*GrpcRoute] {
	return pulumix.Output[map[string]*GrpcRoute]{
		OutputState: o.OutputState,
	}
}

func (o GrpcRouteMapOutput) MapIndex(k pulumi.StringInput) GrpcRouteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GrpcRoute {
		return vs[0].(map[string]*GrpcRoute)[vs[1].(string)]
	}).(GrpcRouteOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GrpcRouteInput)(nil)).Elem(), &GrpcRoute{})
	pulumi.RegisterInputType(reflect.TypeOf((*GrpcRouteArrayInput)(nil)).Elem(), GrpcRouteArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GrpcRouteMapInput)(nil)).Elem(), GrpcRouteMap{})
	pulumi.RegisterOutputType(GrpcRouteOutput{})
	pulumi.RegisterOutputType(GrpcRouteArrayOutput{})
	pulumi.RegisterOutputType(GrpcRouteMapOutput{})
}
