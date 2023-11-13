// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package gkehub

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Feature represents the settings and status of any Hub Feature.
//
// To get more information about Feature, see:
//
// * [API documentation](https://cloud.google.com/anthos/fleet-management/docs/reference/rest/v1/projects.locations.features)
// * How-to Guides
//   - [Registering a Cluster](https://cloud.google.com/anthos/multicluster-management/connect/registering-a-cluster#register_cluster)
//
// ## Example Usage
// ### Gkehub Feature Multi Cluster Ingress
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/container"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/gkehub"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cluster, err := container.NewCluster(ctx, "cluster", &container.ClusterArgs{
//				Location:         pulumi.String("us-central1-a"),
//				InitialNodeCount: pulumi.Int(1),
//			})
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
//			})
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
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Gkehub Feature Multi Cluster Service Discovery
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/gkehub"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := gkehub.NewFeature(ctx, "feature", &gkehub.FeatureArgs{
//				Labels: pulumi.StringMap{
//					"foo": pulumi.String("bar"),
//				},
//				Location: pulumi.String("global"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Gkehub Feature Anthos Service Mesh
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/gkehub"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := gkehub.NewFeature(ctx, "feature", &gkehub.FeatureArgs{
//				Location: pulumi.String("global"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Enable Fleet Observability For Default Logs With Copy
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/gkehub"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := gkehub.NewFeature(ctx, "feature", &gkehub.FeatureArgs{
//				Location: pulumi.String("global"),
//				Spec: &gkehub.FeatureSpecArgs{
//					Fleetobservability: &gkehub.FeatureSpecFleetobservabilityArgs{
//						LoggingConfig: &gkehub.FeatureSpecFleetobservabilityLoggingConfigArgs{
//							DefaultConfig: &gkehub.FeatureSpecFleetobservabilityLoggingConfigDefaultConfigArgs{
//								Mode: pulumi.String("COPY"),
//							},
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
// ### Enable Fleet Observability For Scope Logs With Move
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/gkehub"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := gkehub.NewFeature(ctx, "feature", &gkehub.FeatureArgs{
//				Location: pulumi.String("global"),
//				Spec: &gkehub.FeatureSpecArgs{
//					Fleetobservability: &gkehub.FeatureSpecFleetobservabilityArgs{
//						LoggingConfig: &gkehub.FeatureSpecFleetobservabilityLoggingConfigArgs{
//							FleetScopeLogsConfig: &gkehub.FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigArgs{
//								Mode: pulumi.String("MOVE"),
//							},
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
// ### Enable Fleet Observability For Both Default And Scope Logs
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/gkehub"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := gkehub.NewFeature(ctx, "feature", &gkehub.FeatureArgs{
//				Location: pulumi.String("global"),
//				Spec: &gkehub.FeatureSpecArgs{
//					Fleetobservability: &gkehub.FeatureSpecFleetobservabilityArgs{
//						LoggingConfig: &gkehub.FeatureSpecFleetobservabilityLoggingConfigArgs{
//							DefaultConfig: &gkehub.FeatureSpecFleetobservabilityLoggingConfigDefaultConfigArgs{
//								Mode: pulumi.String("COPY"),
//							},
//							FleetScopeLogsConfig: &gkehub.FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigArgs{
//								Mode: pulumi.String("MOVE"),
//							},
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
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// GCP labels for this Feature.
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The location for the resource
	//
	// ***
	Location pulumi.StringOutput `pulumi:"location"`
	// The full, unique name of this Feature resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	PulumiLabels pulumi.StringMapOutput `pulumi:"pulumiLabels"`
	// State of the Feature resource itself.
	// Structure is documented below.
	ResourceStates FeatureResourceStateArrayOutput `pulumi:"resourceStates"`
	// Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.
	// Structure is documented below.
	Spec FeatureSpecPtrOutput `pulumi:"spec"`
	// (Output)
	// Output only. The "running state" of the Feature in this Hub.
	// Structure is documented below.
	States FeatureStateTypeArrayOutput `pulumi:"states"`
	// (Output)
	// The time this status and any related Feature-specific details were updated. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z"
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
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"effectiveLabels",
		"pulumiLabels",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
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
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// GCP labels for this Feature.
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The location for the resource
	//
	// ***
	Location *string `pulumi:"location"`
	// The full, unique name of this Feature resource
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	PulumiLabels map[string]string `pulumi:"pulumiLabels"`
	// State of the Feature resource itself.
	// Structure is documented below.
	ResourceStates []FeatureResourceState `pulumi:"resourceStates"`
	// Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.
	// Structure is documented below.
	Spec *FeatureSpec `pulumi:"spec"`
	// (Output)
	// Output only. The "running state" of the Feature in this Hub.
	// Structure is documented below.
	States []FeatureStateType `pulumi:"states"`
	// (Output)
	// The time this status and any related Feature-specific details were updated. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z"
	UpdateTime *string `pulumi:"updateTime"`
}

type FeatureState struct {
	// Output only. When the Feature resource was created.
	CreateTime pulumi.StringPtrInput
	// Output only. When the Feature resource was deleted.
	DeleteTime pulumi.StringPtrInput
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
	EffectiveLabels pulumi.StringMapInput
	// GCP labels for this Feature.
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The location for the resource
	//
	// ***
	Location pulumi.StringPtrInput
	// The full, unique name of this Feature resource
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	PulumiLabels pulumi.StringMapInput
	// State of the Feature resource itself.
	// Structure is documented below.
	ResourceStates FeatureResourceStateArrayInput
	// Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.
	// Structure is documented below.
	Spec FeatureSpecPtrInput
	// (Output)
	// Output only. The "running state" of the Feature in this Hub.
	// Structure is documented below.
	States FeatureStateTypeArrayInput
	// (Output)
	// The time this status and any related Feature-specific details were updated. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z"
	UpdateTime pulumi.StringPtrInput
}

func (FeatureState) ElementType() reflect.Type {
	return reflect.TypeOf((*featureState)(nil)).Elem()
}

type featureArgs struct {
	// GCP labels for this Feature.
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The location for the resource
	//
	// ***
	Location string `pulumi:"location"`
	// The full, unique name of this Feature resource
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.
	// Structure is documented below.
	Spec *FeatureSpec `pulumi:"spec"`
}

// The set of arguments for constructing a Feature resource.
type FeatureArgs struct {
	// GCP labels for this Feature.
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The location for the resource
	//
	// ***
	Location pulumi.StringInput
	// The full, unique name of this Feature resource
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.
	// Structure is documented below.
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

func (i *Feature) ToOutput(ctx context.Context) pulumix.Output[*Feature] {
	return pulumix.Output[*Feature]{
		OutputState: i.ToFeatureOutputWithContext(ctx).OutputState,
	}
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

func (i FeatureArray) ToOutput(ctx context.Context) pulumix.Output[[]*Feature] {
	return pulumix.Output[[]*Feature]{
		OutputState: i.ToFeatureArrayOutputWithContext(ctx).OutputState,
	}
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

func (i FeatureMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Feature] {
	return pulumix.Output[map[string]*Feature]{
		OutputState: i.ToFeatureMapOutputWithContext(ctx).OutputState,
	}
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

func (o FeatureOutput) ToOutput(ctx context.Context) pulumix.Output[*Feature] {
	return pulumix.Output[*Feature]{
		OutputState: o.OutputState,
	}
}

// Output only. When the Feature resource was created.
func (o FeatureOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Feature) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Output only. When the Feature resource was deleted.
func (o FeatureOutput) DeleteTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Feature) pulumi.StringOutput { return v.DeleteTime }).(pulumi.StringOutput)
}

// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
func (o FeatureOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Feature) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// GCP labels for this Feature.
// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
func (o FeatureOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Feature) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The location for the resource
//
// ***
func (o FeatureOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Feature) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The full, unique name of this Feature resource
func (o FeatureOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Feature) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o FeatureOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Feature) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource
// and default labels configured on the provider.
func (o FeatureOutput) PulumiLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Feature) pulumi.StringMapOutput { return v.PulumiLabels }).(pulumi.StringMapOutput)
}

// State of the Feature resource itself.
// Structure is documented below.
func (o FeatureOutput) ResourceStates() FeatureResourceStateArrayOutput {
	return o.ApplyT(func(v *Feature) FeatureResourceStateArrayOutput { return v.ResourceStates }).(FeatureResourceStateArrayOutput)
}

// Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.
// Structure is documented below.
func (o FeatureOutput) Spec() FeatureSpecPtrOutput {
	return o.ApplyT(func(v *Feature) FeatureSpecPtrOutput { return v.Spec }).(FeatureSpecPtrOutput)
}

// (Output)
// Output only. The "running state" of the Feature in this Hub.
// Structure is documented below.
func (o FeatureOutput) States() FeatureStateTypeArrayOutput {
	return o.ApplyT(func(v *Feature) FeatureStateTypeArrayOutput { return v.States }).(FeatureStateTypeArrayOutput)
}

// (Output)
// The time this status and any related Feature-specific details were updated. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z"
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

func (o FeatureArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Feature] {
	return pulumix.Output[[]*Feature]{
		OutputState: o.OutputState,
	}
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

func (o FeatureMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Feature] {
	return pulumix.Output[map[string]*Feature]{
		OutputState: o.OutputState,
	}
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
