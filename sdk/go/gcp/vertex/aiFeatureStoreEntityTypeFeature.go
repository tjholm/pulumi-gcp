// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vertex

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Feature Metadata information that describes an attribute of an entity type. For example, apple is an entity type, and color is a feature that describes apple.
//
// To get more information about FeaturestoreEntitytypeFeature, see:
//
// * [API documentation](https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featurestores.entityTypes.features)
// * How-to Guides
//   - [Official Documentation](https://cloud.google.com/vertex-ai/docs)
//
// ## Example Usage
// ### Vertex Ai Featurestore Entitytype Feature
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/vertex"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			featurestore, err := vertex.NewAiFeatureStore(ctx, "featurestore", &vertex.AiFeatureStoreArgs{
//				Labels: pulumi.StringMap{
//					"foo": pulumi.String("bar"),
//				},
//				Region: pulumi.String("us-central1"),
//				OnlineServingConfig: &vertex.AiFeatureStoreOnlineServingConfigArgs{
//					FixedNodeCount: pulumi.Int(2),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			entity, err := vertex.NewAiFeatureStoreEntityType(ctx, "entity", &vertex.AiFeatureStoreEntityTypeArgs{
//				Labels: pulumi.StringMap{
//					"foo": pulumi.String("bar"),
//				},
//				Featurestore: featurestore.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = vertex.NewAiFeatureStoreEntityTypeFeature(ctx, "feature", &vertex.AiFeatureStoreEntityTypeFeatureArgs{
//				Labels: pulumi.StringMap{
//					"foo": pulumi.String("bar"),
//				},
//				Entitytype: entity.ID(),
//				ValueType:  pulumi.String("INT64_ARRAY"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Vertex Ai Featurestore Entitytype Feature With Beta Fields
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/vertex"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			featurestore, err := vertex.NewAiFeatureStore(ctx, "featurestore", &vertex.AiFeatureStoreArgs{
//				Labels: pulumi.StringMap{
//					"foo": pulumi.String("bar"),
//				},
//				Region: pulumi.String("us-central1"),
//				OnlineServingConfig: &vertex.AiFeatureStoreOnlineServingConfigArgs{
//					FixedNodeCount: pulumi.Int(2),
//				},
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			entity, err := vertex.NewAiFeatureStoreEntityType(ctx, "entity", &vertex.AiFeatureStoreEntityTypeArgs{
//				Labels: pulumi.StringMap{
//					"foo": pulumi.String("bar"),
//				},
//				Featurestore: featurestore.ID(),
//				MonitoringConfig: &vertex.AiFeatureStoreEntityTypeMonitoringConfigArgs{
//					SnapshotAnalysis: &vertex.AiFeatureStoreEntityTypeMonitoringConfigSnapshotAnalysisArgs{
//						Disabled:           pulumi.Bool(false),
//						MonitoringInterval: pulumi.String("86400s"),
//					},
//					CategoricalThresholdConfig: &vertex.AiFeatureStoreEntityTypeMonitoringConfigCategoricalThresholdConfigArgs{
//						Value: pulumi.Float64(0.3),
//					},
//					NumericalThresholdConfig: &vertex.AiFeatureStoreEntityTypeMonitoringConfigNumericalThresholdConfigArgs{
//						Value: pulumi.Float64(0.3),
//					},
//				},
//			}, pulumi.Provider(google_beta))
//			if err != nil {
//				return err
//			}
//			_, err = vertex.NewAiFeatureStoreEntityTypeFeature(ctx, "feature", &vertex.AiFeatureStoreEntityTypeFeatureArgs{
//				Labels: pulumi.StringMap{
//					"foo": pulumi.String("bar"),
//				},
//				Entitytype: entity.ID(),
//				ValueType:  pulumi.String("INT64_ARRAY"),
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
// FeaturestoreEntitytypeFeature can be imported using any of these accepted formats:
//
// ```sh
//
//	$ pulumi import gcp:vertex/aiFeatureStoreEntityTypeFeature:AiFeatureStoreEntityTypeFeature default {{entitytype}}/features/{{name}}
//
// ```
type AiFeatureStoreEntityTypeFeature struct {
	pulumi.CustomResourceState

	// The timestamp of when the entity type was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Description of the feature.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
	// clients and services.
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}/entityTypes/{entitytype}.
	//
	// ***
	Entitytype pulumi.StringOutput `pulumi:"entitytype"`
	// Used to perform consistent read-modify-write updates.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// A set of key/value label pairs to assign to the feature.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The name of the feature. The feature can be up to 64 characters long and can consist only of ASCII Latin letters A-Z and a-z, underscore(_), and ASCII digits 0-9 starting with a letter. The value will be unique given an entity type.
	Name pulumi.StringOutput `pulumi:"name"`
	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	PulumiLabels pulumi.StringMapOutput `pulumi:"pulumiLabels"`
	// The region of the feature
	Region pulumi.StringOutput `pulumi:"region"`
	// The timestamp when the entity type was most recently updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// Type of Feature value. Immutable. https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featurestores.entityTypes.features#ValueType
	ValueType pulumi.StringOutput `pulumi:"valueType"`
}

// NewAiFeatureStoreEntityTypeFeature registers a new resource with the given unique name, arguments, and options.
func NewAiFeatureStoreEntityTypeFeature(ctx *pulumi.Context,
	name string, args *AiFeatureStoreEntityTypeFeatureArgs, opts ...pulumi.ResourceOption) (*AiFeatureStoreEntityTypeFeature, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Entitytype == nil {
		return nil, errors.New("invalid value for required argument 'Entitytype'")
	}
	if args.ValueType == nil {
		return nil, errors.New("invalid value for required argument 'ValueType'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"effectiveLabels",
		"pulumiLabels",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AiFeatureStoreEntityTypeFeature
	err := ctx.RegisterResource("gcp:vertex/aiFeatureStoreEntityTypeFeature:AiFeatureStoreEntityTypeFeature", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAiFeatureStoreEntityTypeFeature gets an existing AiFeatureStoreEntityTypeFeature resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAiFeatureStoreEntityTypeFeature(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AiFeatureStoreEntityTypeFeatureState, opts ...pulumi.ResourceOption) (*AiFeatureStoreEntityTypeFeature, error) {
	var resource AiFeatureStoreEntityTypeFeature
	err := ctx.ReadResource("gcp:vertex/aiFeatureStoreEntityTypeFeature:AiFeatureStoreEntityTypeFeature", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AiFeatureStoreEntityTypeFeature resources.
type aiFeatureStoreEntityTypeFeatureState struct {
	// The timestamp of when the entity type was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	CreateTime *string `pulumi:"createTime"`
	// Description of the feature.
	Description *string `pulumi:"description"`
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
	// clients and services.
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}/entityTypes/{entitytype}.
	//
	// ***
	Entitytype *string `pulumi:"entitytype"`
	// Used to perform consistent read-modify-write updates.
	Etag *string `pulumi:"etag"`
	// A set of key/value label pairs to assign to the feature.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The name of the feature. The feature can be up to 64 characters long and can consist only of ASCII Latin letters A-Z and a-z, underscore(_), and ASCII digits 0-9 starting with a letter. The value will be unique given an entity type.
	Name *string `pulumi:"name"`
	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	PulumiLabels map[string]string `pulumi:"pulumiLabels"`
	// The region of the feature
	Region *string `pulumi:"region"`
	// The timestamp when the entity type was most recently updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	UpdateTime *string `pulumi:"updateTime"`
	// Type of Feature value. Immutable. https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featurestores.entityTypes.features#ValueType
	ValueType *string `pulumi:"valueType"`
}

type AiFeatureStoreEntityTypeFeatureState struct {
	// The timestamp of when the entity type was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	CreateTime pulumi.StringPtrInput
	// Description of the feature.
	Description pulumi.StringPtrInput
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
	// clients and services.
	EffectiveLabels pulumi.StringMapInput
	// The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}/entityTypes/{entitytype}.
	//
	// ***
	Entitytype pulumi.StringPtrInput
	// Used to perform consistent read-modify-write updates.
	Etag pulumi.StringPtrInput
	// A set of key/value label pairs to assign to the feature.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The name of the feature. The feature can be up to 64 characters long and can consist only of ASCII Latin letters A-Z and a-z, underscore(_), and ASCII digits 0-9 starting with a letter. The value will be unique given an entity type.
	Name pulumi.StringPtrInput
	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	PulumiLabels pulumi.StringMapInput
	// The region of the feature
	Region pulumi.StringPtrInput
	// The timestamp when the entity type was most recently updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
	UpdateTime pulumi.StringPtrInput
	// Type of Feature value. Immutable. https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featurestores.entityTypes.features#ValueType
	ValueType pulumi.StringPtrInput
}

func (AiFeatureStoreEntityTypeFeatureState) ElementType() reflect.Type {
	return reflect.TypeOf((*aiFeatureStoreEntityTypeFeatureState)(nil)).Elem()
}

type aiFeatureStoreEntityTypeFeatureArgs struct {
	// Description of the feature.
	Description *string `pulumi:"description"`
	// The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}/entityTypes/{entitytype}.
	//
	// ***
	Entitytype string `pulumi:"entitytype"`
	// A set of key/value label pairs to assign to the feature.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// The name of the feature. The feature can be up to 64 characters long and can consist only of ASCII Latin letters A-Z and a-z, underscore(_), and ASCII digits 0-9 starting with a letter. The value will be unique given an entity type.
	Name *string `pulumi:"name"`
	// Type of Feature value. Immutable. https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featurestores.entityTypes.features#ValueType
	ValueType string `pulumi:"valueType"`
}

// The set of arguments for constructing a AiFeatureStoreEntityTypeFeature resource.
type AiFeatureStoreEntityTypeFeatureArgs struct {
	// Description of the feature.
	Description pulumi.StringPtrInput
	// The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}/entityTypes/{entitytype}.
	//
	// ***
	Entitytype pulumi.StringInput
	// A set of key/value label pairs to assign to the feature.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// The name of the feature. The feature can be up to 64 characters long and can consist only of ASCII Latin letters A-Z and a-z, underscore(_), and ASCII digits 0-9 starting with a letter. The value will be unique given an entity type.
	Name pulumi.StringPtrInput
	// Type of Feature value. Immutable. https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featurestores.entityTypes.features#ValueType
	ValueType pulumi.StringInput
}

func (AiFeatureStoreEntityTypeFeatureArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aiFeatureStoreEntityTypeFeatureArgs)(nil)).Elem()
}

type AiFeatureStoreEntityTypeFeatureInput interface {
	pulumi.Input

	ToAiFeatureStoreEntityTypeFeatureOutput() AiFeatureStoreEntityTypeFeatureOutput
	ToAiFeatureStoreEntityTypeFeatureOutputWithContext(ctx context.Context) AiFeatureStoreEntityTypeFeatureOutput
}

func (*AiFeatureStoreEntityTypeFeature) ElementType() reflect.Type {
	return reflect.TypeOf((**AiFeatureStoreEntityTypeFeature)(nil)).Elem()
}

func (i *AiFeatureStoreEntityTypeFeature) ToAiFeatureStoreEntityTypeFeatureOutput() AiFeatureStoreEntityTypeFeatureOutput {
	return i.ToAiFeatureStoreEntityTypeFeatureOutputWithContext(context.Background())
}

func (i *AiFeatureStoreEntityTypeFeature) ToAiFeatureStoreEntityTypeFeatureOutputWithContext(ctx context.Context) AiFeatureStoreEntityTypeFeatureOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AiFeatureStoreEntityTypeFeatureOutput)
}

func (i *AiFeatureStoreEntityTypeFeature) ToOutput(ctx context.Context) pulumix.Output[*AiFeatureStoreEntityTypeFeature] {
	return pulumix.Output[*AiFeatureStoreEntityTypeFeature]{
		OutputState: i.ToAiFeatureStoreEntityTypeFeatureOutputWithContext(ctx).OutputState,
	}
}

// AiFeatureStoreEntityTypeFeatureArrayInput is an input type that accepts AiFeatureStoreEntityTypeFeatureArray and AiFeatureStoreEntityTypeFeatureArrayOutput values.
// You can construct a concrete instance of `AiFeatureStoreEntityTypeFeatureArrayInput` via:
//
//	AiFeatureStoreEntityTypeFeatureArray{ AiFeatureStoreEntityTypeFeatureArgs{...} }
type AiFeatureStoreEntityTypeFeatureArrayInput interface {
	pulumi.Input

	ToAiFeatureStoreEntityTypeFeatureArrayOutput() AiFeatureStoreEntityTypeFeatureArrayOutput
	ToAiFeatureStoreEntityTypeFeatureArrayOutputWithContext(context.Context) AiFeatureStoreEntityTypeFeatureArrayOutput
}

type AiFeatureStoreEntityTypeFeatureArray []AiFeatureStoreEntityTypeFeatureInput

func (AiFeatureStoreEntityTypeFeatureArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AiFeatureStoreEntityTypeFeature)(nil)).Elem()
}

func (i AiFeatureStoreEntityTypeFeatureArray) ToAiFeatureStoreEntityTypeFeatureArrayOutput() AiFeatureStoreEntityTypeFeatureArrayOutput {
	return i.ToAiFeatureStoreEntityTypeFeatureArrayOutputWithContext(context.Background())
}

func (i AiFeatureStoreEntityTypeFeatureArray) ToAiFeatureStoreEntityTypeFeatureArrayOutputWithContext(ctx context.Context) AiFeatureStoreEntityTypeFeatureArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AiFeatureStoreEntityTypeFeatureArrayOutput)
}

func (i AiFeatureStoreEntityTypeFeatureArray) ToOutput(ctx context.Context) pulumix.Output[[]*AiFeatureStoreEntityTypeFeature] {
	return pulumix.Output[[]*AiFeatureStoreEntityTypeFeature]{
		OutputState: i.ToAiFeatureStoreEntityTypeFeatureArrayOutputWithContext(ctx).OutputState,
	}
}

// AiFeatureStoreEntityTypeFeatureMapInput is an input type that accepts AiFeatureStoreEntityTypeFeatureMap and AiFeatureStoreEntityTypeFeatureMapOutput values.
// You can construct a concrete instance of `AiFeatureStoreEntityTypeFeatureMapInput` via:
//
//	AiFeatureStoreEntityTypeFeatureMap{ "key": AiFeatureStoreEntityTypeFeatureArgs{...} }
type AiFeatureStoreEntityTypeFeatureMapInput interface {
	pulumi.Input

	ToAiFeatureStoreEntityTypeFeatureMapOutput() AiFeatureStoreEntityTypeFeatureMapOutput
	ToAiFeatureStoreEntityTypeFeatureMapOutputWithContext(context.Context) AiFeatureStoreEntityTypeFeatureMapOutput
}

type AiFeatureStoreEntityTypeFeatureMap map[string]AiFeatureStoreEntityTypeFeatureInput

func (AiFeatureStoreEntityTypeFeatureMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AiFeatureStoreEntityTypeFeature)(nil)).Elem()
}

func (i AiFeatureStoreEntityTypeFeatureMap) ToAiFeatureStoreEntityTypeFeatureMapOutput() AiFeatureStoreEntityTypeFeatureMapOutput {
	return i.ToAiFeatureStoreEntityTypeFeatureMapOutputWithContext(context.Background())
}

func (i AiFeatureStoreEntityTypeFeatureMap) ToAiFeatureStoreEntityTypeFeatureMapOutputWithContext(ctx context.Context) AiFeatureStoreEntityTypeFeatureMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AiFeatureStoreEntityTypeFeatureMapOutput)
}

func (i AiFeatureStoreEntityTypeFeatureMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*AiFeatureStoreEntityTypeFeature] {
	return pulumix.Output[map[string]*AiFeatureStoreEntityTypeFeature]{
		OutputState: i.ToAiFeatureStoreEntityTypeFeatureMapOutputWithContext(ctx).OutputState,
	}
}

type AiFeatureStoreEntityTypeFeatureOutput struct{ *pulumi.OutputState }

func (AiFeatureStoreEntityTypeFeatureOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AiFeatureStoreEntityTypeFeature)(nil)).Elem()
}

func (o AiFeatureStoreEntityTypeFeatureOutput) ToAiFeatureStoreEntityTypeFeatureOutput() AiFeatureStoreEntityTypeFeatureOutput {
	return o
}

func (o AiFeatureStoreEntityTypeFeatureOutput) ToAiFeatureStoreEntityTypeFeatureOutputWithContext(ctx context.Context) AiFeatureStoreEntityTypeFeatureOutput {
	return o
}

func (o AiFeatureStoreEntityTypeFeatureOutput) ToOutput(ctx context.Context) pulumix.Output[*AiFeatureStoreEntityTypeFeature] {
	return pulumix.Output[*AiFeatureStoreEntityTypeFeature]{
		OutputState: o.OutputState,
	}
}

// The timestamp of when the entity type was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
func (o AiFeatureStoreEntityTypeFeatureOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AiFeatureStoreEntityTypeFeature) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Description of the feature.
func (o AiFeatureStoreEntityTypeFeatureOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AiFeatureStoreEntityTypeFeature) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
// clients and services.
func (o AiFeatureStoreEntityTypeFeatureOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AiFeatureStoreEntityTypeFeature) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}/entityTypes/{entitytype}.
//
// ***
func (o AiFeatureStoreEntityTypeFeatureOutput) Entitytype() pulumi.StringOutput {
	return o.ApplyT(func(v *AiFeatureStoreEntityTypeFeature) pulumi.StringOutput { return v.Entitytype }).(pulumi.StringOutput)
}

// Used to perform consistent read-modify-write updates.
func (o AiFeatureStoreEntityTypeFeatureOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *AiFeatureStoreEntityTypeFeature) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// A set of key/value label pairs to assign to the feature.
//
// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
func (o AiFeatureStoreEntityTypeFeatureOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AiFeatureStoreEntityTypeFeature) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The name of the feature. The feature can be up to 64 characters long and can consist only of ASCII Latin letters A-Z and a-z, underscore(_), and ASCII digits 0-9 starting with a letter. The value will be unique given an entity type.
func (o AiFeatureStoreEntityTypeFeatureOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AiFeatureStoreEntityTypeFeature) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The combination of labels configured directly on the resource
// and default labels configured on the provider.
func (o AiFeatureStoreEntityTypeFeatureOutput) PulumiLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AiFeatureStoreEntityTypeFeature) pulumi.StringMapOutput { return v.PulumiLabels }).(pulumi.StringMapOutput)
}

// The region of the feature
func (o AiFeatureStoreEntityTypeFeatureOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *AiFeatureStoreEntityTypeFeature) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The timestamp when the entity type was most recently updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
func (o AiFeatureStoreEntityTypeFeatureOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AiFeatureStoreEntityTypeFeature) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// Type of Feature value. Immutable. https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featurestores.entityTypes.features#ValueType
func (o AiFeatureStoreEntityTypeFeatureOutput) ValueType() pulumi.StringOutput {
	return o.ApplyT(func(v *AiFeatureStoreEntityTypeFeature) pulumi.StringOutput { return v.ValueType }).(pulumi.StringOutput)
}

type AiFeatureStoreEntityTypeFeatureArrayOutput struct{ *pulumi.OutputState }

func (AiFeatureStoreEntityTypeFeatureArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AiFeatureStoreEntityTypeFeature)(nil)).Elem()
}

func (o AiFeatureStoreEntityTypeFeatureArrayOutput) ToAiFeatureStoreEntityTypeFeatureArrayOutput() AiFeatureStoreEntityTypeFeatureArrayOutput {
	return o
}

func (o AiFeatureStoreEntityTypeFeatureArrayOutput) ToAiFeatureStoreEntityTypeFeatureArrayOutputWithContext(ctx context.Context) AiFeatureStoreEntityTypeFeatureArrayOutput {
	return o
}

func (o AiFeatureStoreEntityTypeFeatureArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*AiFeatureStoreEntityTypeFeature] {
	return pulumix.Output[[]*AiFeatureStoreEntityTypeFeature]{
		OutputState: o.OutputState,
	}
}

func (o AiFeatureStoreEntityTypeFeatureArrayOutput) Index(i pulumi.IntInput) AiFeatureStoreEntityTypeFeatureOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AiFeatureStoreEntityTypeFeature {
		return vs[0].([]*AiFeatureStoreEntityTypeFeature)[vs[1].(int)]
	}).(AiFeatureStoreEntityTypeFeatureOutput)
}

type AiFeatureStoreEntityTypeFeatureMapOutput struct{ *pulumi.OutputState }

func (AiFeatureStoreEntityTypeFeatureMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AiFeatureStoreEntityTypeFeature)(nil)).Elem()
}

func (o AiFeatureStoreEntityTypeFeatureMapOutput) ToAiFeatureStoreEntityTypeFeatureMapOutput() AiFeatureStoreEntityTypeFeatureMapOutput {
	return o
}

func (o AiFeatureStoreEntityTypeFeatureMapOutput) ToAiFeatureStoreEntityTypeFeatureMapOutputWithContext(ctx context.Context) AiFeatureStoreEntityTypeFeatureMapOutput {
	return o
}

func (o AiFeatureStoreEntityTypeFeatureMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*AiFeatureStoreEntityTypeFeature] {
	return pulumix.Output[map[string]*AiFeatureStoreEntityTypeFeature]{
		OutputState: o.OutputState,
	}
}

func (o AiFeatureStoreEntityTypeFeatureMapOutput) MapIndex(k pulumi.StringInput) AiFeatureStoreEntityTypeFeatureOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AiFeatureStoreEntityTypeFeature {
		return vs[0].(map[string]*AiFeatureStoreEntityTypeFeature)[vs[1].(string)]
	}).(AiFeatureStoreEntityTypeFeatureOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AiFeatureStoreEntityTypeFeatureInput)(nil)).Elem(), &AiFeatureStoreEntityTypeFeature{})
	pulumi.RegisterInputType(reflect.TypeOf((*AiFeatureStoreEntityTypeFeatureArrayInput)(nil)).Elem(), AiFeatureStoreEntityTypeFeatureArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AiFeatureStoreEntityTypeFeatureMapInput)(nil)).Elem(), AiFeatureStoreEntityTypeFeatureMap{})
	pulumi.RegisterOutputType(AiFeatureStoreEntityTypeFeatureOutput{})
	pulumi.RegisterOutputType(AiFeatureStoreEntityTypeFeatureArrayOutput{})
	pulumi.RegisterOutputType(AiFeatureStoreEntityTypeFeatureMapOutput{})
}
