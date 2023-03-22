// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gkehub

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
// ### Multi Cluster Ingress
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/container"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/gkehub"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cluster, err := container.NewCluster(ctx, "cluster", &container.ClusterArgs{
//				Location:         pulumi.String("us-central1-a"),
//				InitialNodeCount: pulumi.Int(1),
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			membership, err := gkehub.NewMembership(ctx, "membership", &gkehub.MembershipArgs{
//				MembershipId: pulumi.String("my-membership"),
//				Endpoint: &gkehub.MembershipEndpointArgs{
//					GkeCluster: &gkehub.MembershipEndpointGkeClusterArgs{
//						ResourceLink: cluster.ID().ApplyT(func(id string) (string, error) {
//							return fmt.Sprintf("//container.googleapis.com/%v", id), nil
//						}).(pulumi.StringOutput),
//					},
//				},
//				Description: pulumi.String("Membership"),
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			_, err = gkehub.NewFeature(ctx, "feature", &gkehub.FeatureArgs{
//				Location: pulumi.String("global"),
//				Spec: &gkehub.FeatureSpecArgs{
//					Multiclusteringress: &gkehub.FeatureSpecMulticlusteringressArgs{
//						ConfigMembership: membership.ID(),
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
// ### Multi Cluster Service Discovery
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/gkehub"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := gkehub.NewFeature(ctx, "feature", &gkehub.FeatureArgs{
//				Location: pulumi.String("global"),
//				Labels: pulumi.StringMap{
//					"foo": pulumi.String("bar"),
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
// ### Enable Anthos Service Mesh
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/gkehub"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := gkehub.NewFeature(ctx, "feature", &gkehub.FeatureArgs{
//				Location: pulumi.String("global"),
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
// # Feature can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:gkehub/feature:Feature default projects/{{project}}/locations/{{location}}/features/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:gkehub/feature:Feature default {{project}}/{{location}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:gkehub/feature:Feature default {{location}}/{{name}}
//
// ```
type Feature struct {
	pulumi.CustomResourceState

	// Output only. When the Feature resource was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Output only. When the Feature resource was deleted.
	DeleteTime pulumi.StringOutput `pulumi:"deleteTime"`
	// GCP labels for this Feature.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The location for the resource
	Location pulumi.StringOutput `pulumi:"location"`
	// The full, unique name of this Feature resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The project for the resource
	Project pulumi.StringOutput `pulumi:"project"`
	// State of the Feature resource itself.
	ResourceStates FeatureResourceStateArrayOutput `pulumi:"resourceStates"`
	// Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.
	Spec FeatureSpecPtrOutput `pulumi:"spec"`
	// Output only. The Hub-wide Feature state
	States FeatureStateTypeArrayOutput `pulumi:"states"`
	// Output only. When the Feature resource was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewFeature registers a new resource with the given unique name, arguments, and options.
func NewFeature(ctx *pulumi.Context,
	name string, args *FeatureArgs, opts ...pulumi.ResourceOption) (*Feature, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	var resource Feature
	err := ctx.RegisterResource("gcp:gkehub/feature:Feature", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFeature gets an existing Feature resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFeature(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FeatureState, opts ...pulumi.ResourceOption) (*Feature, error) {
	var resource Feature
	err := ctx.ReadResource("gcp:gkehub/feature:Feature", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Feature resources.
type featureState struct {
	// Output only. When the Feature resource was created.
	CreateTime *string `pulumi:"createTime"`
	// Output only. When the Feature resource was deleted.
	DeleteTime *string `pulumi:"deleteTime"`
	// GCP labels for this Feature.
	Labels map[string]string `pulumi:"labels"`
	// The location for the resource
	Location *string `pulumi:"location"`
	// The full, unique name of this Feature resource
	Name *string `pulumi:"name"`
	// The project for the resource
	Project *string `pulumi:"project"`
	// State of the Feature resource itself.
	ResourceStates []FeatureResourceState `pulumi:"resourceStates"`
	// Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.
	Spec *FeatureSpec `pulumi:"spec"`
	// Output only. The Hub-wide Feature state
	States []FeatureStateType `pulumi:"states"`
	// Output only. When the Feature resource was last updated.
	UpdateTime *string `pulumi:"updateTime"`
}

type FeatureState struct {
	// Output only. When the Feature resource was created.
	CreateTime pulumi.StringPtrInput
	// Output only. When the Feature resource was deleted.
	DeleteTime pulumi.StringPtrInput
	// GCP labels for this Feature.
	Labels pulumi.StringMapInput
	// The location for the resource
	Location pulumi.StringPtrInput
	// The full, unique name of this Feature resource
	Name pulumi.StringPtrInput
	// The project for the resource
	Project pulumi.StringPtrInput
	// State of the Feature resource itself.
	ResourceStates FeatureResourceStateArrayInput
	// Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.
	Spec FeatureSpecPtrInput
	// Output only. The Hub-wide Feature state
	States FeatureStateTypeArrayInput
	// Output only. When the Feature resource was last updated.
	UpdateTime pulumi.StringPtrInput
}

func (FeatureState) ElementType() reflect.Type {
	return reflect.TypeOf((*featureState)(nil)).Elem()
}

type featureArgs struct {
	// GCP labels for this Feature.
	Labels map[string]string `pulumi:"labels"`
	// The location for the resource
	Location string `pulumi:"location"`
	// The full, unique name of this Feature resource
	Name *string `pulumi:"name"`
	// The project for the resource
	Project *string `pulumi:"project"`
	// Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.
	Spec *FeatureSpec `pulumi:"spec"`
}

// The set of arguments for constructing a Feature resource.
type FeatureArgs struct {
	// GCP labels for this Feature.
	Labels pulumi.StringMapInput
	// The location for the resource
	Location pulumi.StringInput
	// The full, unique name of this Feature resource
	Name pulumi.StringPtrInput
	// The project for the resource
	Project pulumi.StringPtrInput
	// Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.
	Spec FeatureSpecPtrInput
}

func (FeatureArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*featureArgs)(nil)).Elem()
}

type FeatureInput interface {
	pulumi.Input

	ToFeatureOutput() FeatureOutput
	ToFeatureOutputWithContext(ctx context.Context) FeatureOutput
}

func (*Feature) ElementType() reflect.Type {
	return reflect.TypeOf((**Feature)(nil)).Elem()
}

func (i *Feature) ToFeatureOutput() FeatureOutput {
	return i.ToFeatureOutputWithContext(context.Background())
}

func (i *Feature) ToFeatureOutputWithContext(ctx context.Context) FeatureOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeatureOutput)
}

// FeatureArrayInput is an input type that accepts FeatureArray and FeatureArrayOutput values.
// You can construct a concrete instance of `FeatureArrayInput` via:
//
//	FeatureArray{ FeatureArgs{...} }
type FeatureArrayInput interface {
	pulumi.Input

	ToFeatureArrayOutput() FeatureArrayOutput
	ToFeatureArrayOutputWithContext(context.Context) FeatureArrayOutput
}

type FeatureArray []FeatureInput

func (FeatureArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Feature)(nil)).Elem()
}

func (i FeatureArray) ToFeatureArrayOutput() FeatureArrayOutput {
	return i.ToFeatureArrayOutputWithContext(context.Background())
}

func (i FeatureArray) ToFeatureArrayOutputWithContext(ctx context.Context) FeatureArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeatureArrayOutput)
}

// FeatureMapInput is an input type that accepts FeatureMap and FeatureMapOutput values.
// You can construct a concrete instance of `FeatureMapInput` via:
//
//	FeatureMap{ "key": FeatureArgs{...} }
type FeatureMapInput interface {
	pulumi.Input

	ToFeatureMapOutput() FeatureMapOutput
	ToFeatureMapOutputWithContext(context.Context) FeatureMapOutput
}

type FeatureMap map[string]FeatureInput

func (FeatureMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Feature)(nil)).Elem()
}

func (i FeatureMap) ToFeatureMapOutput() FeatureMapOutput {
	return i.ToFeatureMapOutputWithContext(context.Background())
}

func (i FeatureMap) ToFeatureMapOutputWithContext(ctx context.Context) FeatureMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FeatureMapOutput)
}

type FeatureOutput struct{ *pulumi.OutputState }

func (FeatureOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Feature)(nil)).Elem()
}

func (o FeatureOutput) ToFeatureOutput() FeatureOutput {
	return o
}

func (o FeatureOutput) ToFeatureOutputWithContext(ctx context.Context) FeatureOutput {
	return o
}

// Output only. When the Feature resource was created.
func (o FeatureOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Feature) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Output only. When the Feature resource was deleted.
func (o FeatureOutput) DeleteTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Feature) pulumi.StringOutput { return v.DeleteTime }).(pulumi.StringOutput)
}

// GCP labels for this Feature.
func (o FeatureOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Feature) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The location for the resource
func (o FeatureOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Feature) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The full, unique name of this Feature resource
func (o FeatureOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Feature) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project for the resource
func (o FeatureOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Feature) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// State of the Feature resource itself.
func (o FeatureOutput) ResourceStates() FeatureResourceStateArrayOutput {
	return o.ApplyT(func(v *Feature) FeatureResourceStateArrayOutput { return v.ResourceStates }).(FeatureResourceStateArrayOutput)
}

// Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.
func (o FeatureOutput) Spec() FeatureSpecPtrOutput {
	return o.ApplyT(func(v *Feature) FeatureSpecPtrOutput { return v.Spec }).(FeatureSpecPtrOutput)
}

// Output only. The Hub-wide Feature state
func (o FeatureOutput) States() FeatureStateTypeArrayOutput {
	return o.ApplyT(func(v *Feature) FeatureStateTypeArrayOutput { return v.States }).(FeatureStateTypeArrayOutput)
}

// Output only. When the Feature resource was last updated.
func (o FeatureOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Feature) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

type FeatureArrayOutput struct{ *pulumi.OutputState }

func (FeatureArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Feature)(nil)).Elem()
}

func (o FeatureArrayOutput) ToFeatureArrayOutput() FeatureArrayOutput {
	return o
}

func (o FeatureArrayOutput) ToFeatureArrayOutputWithContext(ctx context.Context) FeatureArrayOutput {
	return o
}

func (o FeatureArrayOutput) Index(i pulumi.IntInput) FeatureOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Feature {
		return vs[0].([]*Feature)[vs[1].(int)]
	}).(FeatureOutput)
}

type FeatureMapOutput struct{ *pulumi.OutputState }

func (FeatureMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Feature)(nil)).Elem()
}

func (o FeatureMapOutput) ToFeatureMapOutput() FeatureMapOutput {
	return o
}

func (o FeatureMapOutput) ToFeatureMapOutputWithContext(ctx context.Context) FeatureMapOutput {
	return o
}

func (o FeatureMapOutput) MapIndex(k pulumi.StringInput) FeatureOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Feature {
		return vs[0].(map[string]*Feature)[vs[1].(string)]
	}).(FeatureOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FeatureInput)(nil)).Elem(), &Feature{})
	pulumi.RegisterInputType(reflect.TypeOf((*FeatureArrayInput)(nil)).Elem(), FeatureArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FeatureMapInput)(nil)).Elem(), FeatureMap{})
	pulumi.RegisterOutputType(FeatureOutput{})
	pulumi.RegisterOutputType(FeatureArrayOutput{})
	pulumi.RegisterOutputType(FeatureMapOutput{})
}
