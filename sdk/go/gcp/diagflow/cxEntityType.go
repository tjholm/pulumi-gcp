// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package diagflow

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Entities are extracted from user input and represent parameters that are meaningful to your application.
// For example, a date range, a proper name such as a geographic location or landmark, and so on. Entities represent actionable data for your application.
//
// To get more information about EntityType, see:
//
// * [API documentation](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.entityTypes)
// * How-to Guides
//   - [Official Documentation](https://cloud.google.com/dialogflow/cx/docs)
//
// ## Example Usage
// ### Dialogflowcx Entity Type Full
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/diagflow"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			agent, err := diagflow.NewCxAgent(ctx, "agent", &diagflow.CxAgentArgs{
//				DisplayName:         pulumi.String("dialogflowcx-agent"),
//				Location:            pulumi.String("global"),
//				DefaultLanguageCode: pulumi.String("en"),
//				SupportedLanguageCodes: pulumi.StringArray{
//					pulumi.String("fr"),
//					pulumi.String("de"),
//					pulumi.String("es"),
//				},
//				TimeZone:                 pulumi.String("America/New_York"),
//				Description:              pulumi.String("Example description."),
//				AvatarUri:                pulumi.String("https://cloud.google.com/_static/images/cloud/icons/favicons/onecloud/super_cloud.png"),
//				EnableStackdriverLogging: pulumi.Bool(true),
//				EnableSpellCorrection:    pulumi.Bool(true),
//				SpeechToTextSettings: &diagflow.CxAgentSpeechToTextSettingsArgs{
//					EnableSpeechAdaptation: pulumi.Bool(true),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = diagflow.NewCxEntityType(ctx, "basicEntityType", &diagflow.CxEntityTypeArgs{
//				Parent:      agent.ID(),
//				DisplayName: pulumi.String("MyEntity"),
//				Kind:        pulumi.String("KIND_MAP"),
//				Entities: diagflow.CxEntityTypeEntityArray{
//					&diagflow.CxEntityTypeEntityArgs{
//						Value: pulumi.String("value1"),
//						Synonyms: pulumi.StringArray{
//							pulumi.String("synonym1"),
//							pulumi.String("synonym2"),
//						},
//					},
//					&diagflow.CxEntityTypeEntityArgs{
//						Value: pulumi.String("value2"),
//						Synonyms: pulumi.StringArray{
//							pulumi.String("synonym3"),
//							pulumi.String("synonym4"),
//						},
//					},
//				},
//				EnableFuzzyExtraction: pulumi.Bool(false),
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
// # EntityType can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:diagflow/cxEntityType:CxEntityType default {{parent}}/entityTypes/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:diagflow/cxEntityType:CxEntityType default {{parent}}/{{name}}
//
// ```
type CxEntityType struct {
	pulumi.CustomResourceState

	// Represents kinds of entities.
	// * AUTO_EXPANSION_MODE_UNSPECIFIED: Auto expansion disabled for the entity.
	// * AUTO_EXPANSION_MODE_DEFAULT: Allows an agent to recognize values that have not been explicitly listed in the entity.
	//   Possible values are `AUTO_EXPANSION_MODE_DEFAULT` and `AUTO_EXPANSION_MODE_UNSPECIFIED`.
	AutoExpansionMode pulumi.StringPtrOutput `pulumi:"autoExpansionMode"`
	// The human-readable name of the entity type, unique within the agent.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Enables fuzzy entity extraction during classification.
	EnableFuzzyExtraction pulumi.BoolPtrOutput `pulumi:"enableFuzzyExtraction"`
	// The collection of entity entries associated with the entity type.
	// Structure is documented below.
	Entities CxEntityTypeEntityArrayOutput `pulumi:"entities"`
	// Collection of exceptional words and phrases that shouldn't be matched. For example, if you have a size entity type with entry giant(an adjective), you might consider adding giants(a noun) as an exclusion.
	// If the kind of entity type is KIND_MAP, then the phrases specified by entities and excluded phrases should be mutually exclusive.
	// Structure is documented below.
	ExcludedPhrases CxEntityTypeExcludedPhraseArrayOutput `pulumi:"excludedPhrases"`
	// Indicates whether the entity type can be automatically expanded.
	// * KIND_MAP: Map entity types allow mapping of a group of synonyms to a canonical value.
	// * KIND_LIST: List entity types contain a set of entries that do not map to canonical values. However, list entity types can contain references to other entity types (with or without aliases).
	// * KIND_REGEXP: Regexp entity types allow to specify regular expressions in entries values.
	//   Possible values are `KIND_MAP`, `KIND_LIST`, and `KIND_REGEXP`.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The language of the following fields in entityType:
	// EntityType.entities.value
	// EntityType.entities.synonyms
	// EntityType.excluded_phrases.value
	// If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
	LanguageCode pulumi.StringPtrOutput `pulumi:"languageCode"`
	// The unique identifier of the entity type. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent
	// ID>/entityTypes/<Entity Type ID>.
	Name pulumi.StringOutput `pulumi:"name"`
	// The agent to create a entity type for.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	Parent pulumi.StringPtrOutput `pulumi:"parent"`
	// Indicates whether parameters of the entity type should be redacted in log. If redaction is enabled, page parameters and intent parameters referring to the entity type will be replaced by parameter name when logging.
	Redact pulumi.BoolPtrOutput `pulumi:"redact"`
}

// NewCxEntityType registers a new resource with the given unique name, arguments, and options.
func NewCxEntityType(ctx *pulumi.Context,
	name string, args *CxEntityTypeArgs, opts ...pulumi.ResourceOption) (*CxEntityType, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Entities == nil {
		return nil, errors.New("invalid value for required argument 'Entities'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	var resource CxEntityType
	err := ctx.RegisterResource("gcp:diagflow/cxEntityType:CxEntityType", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCxEntityType gets an existing CxEntityType resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCxEntityType(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CxEntityTypeState, opts ...pulumi.ResourceOption) (*CxEntityType, error) {
	var resource CxEntityType
	err := ctx.ReadResource("gcp:diagflow/cxEntityType:CxEntityType", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CxEntityType resources.
type cxEntityTypeState struct {
	// Represents kinds of entities.
	// * AUTO_EXPANSION_MODE_UNSPECIFIED: Auto expansion disabled for the entity.
	// * AUTO_EXPANSION_MODE_DEFAULT: Allows an agent to recognize values that have not been explicitly listed in the entity.
	//   Possible values are `AUTO_EXPANSION_MODE_DEFAULT` and `AUTO_EXPANSION_MODE_UNSPECIFIED`.
	AutoExpansionMode *string `pulumi:"autoExpansionMode"`
	// The human-readable name of the entity type, unique within the agent.
	DisplayName *string `pulumi:"displayName"`
	// Enables fuzzy entity extraction during classification.
	EnableFuzzyExtraction *bool `pulumi:"enableFuzzyExtraction"`
	// The collection of entity entries associated with the entity type.
	// Structure is documented below.
	Entities []CxEntityTypeEntity `pulumi:"entities"`
	// Collection of exceptional words and phrases that shouldn't be matched. For example, if you have a size entity type with entry giant(an adjective), you might consider adding giants(a noun) as an exclusion.
	// If the kind of entity type is KIND_MAP, then the phrases specified by entities and excluded phrases should be mutually exclusive.
	// Structure is documented below.
	ExcludedPhrases []CxEntityTypeExcludedPhrase `pulumi:"excludedPhrases"`
	// Indicates whether the entity type can be automatically expanded.
	// * KIND_MAP: Map entity types allow mapping of a group of synonyms to a canonical value.
	// * KIND_LIST: List entity types contain a set of entries that do not map to canonical values. However, list entity types can contain references to other entity types (with or without aliases).
	// * KIND_REGEXP: Regexp entity types allow to specify regular expressions in entries values.
	//   Possible values are `KIND_MAP`, `KIND_LIST`, and `KIND_REGEXP`.
	Kind *string `pulumi:"kind"`
	// The language of the following fields in entityType:
	// EntityType.entities.value
	// EntityType.entities.synonyms
	// EntityType.excluded_phrases.value
	// If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
	LanguageCode *string `pulumi:"languageCode"`
	// The unique identifier of the entity type. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent
	// ID>/entityTypes/<Entity Type ID>.
	Name *string `pulumi:"name"`
	// The agent to create a entity type for.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	Parent *string `pulumi:"parent"`
	// Indicates whether parameters of the entity type should be redacted in log. If redaction is enabled, page parameters and intent parameters referring to the entity type will be replaced by parameter name when logging.
	Redact *bool `pulumi:"redact"`
}

type CxEntityTypeState struct {
	// Represents kinds of entities.
	// * AUTO_EXPANSION_MODE_UNSPECIFIED: Auto expansion disabled for the entity.
	// * AUTO_EXPANSION_MODE_DEFAULT: Allows an agent to recognize values that have not been explicitly listed in the entity.
	//   Possible values are `AUTO_EXPANSION_MODE_DEFAULT` and `AUTO_EXPANSION_MODE_UNSPECIFIED`.
	AutoExpansionMode pulumi.StringPtrInput
	// The human-readable name of the entity type, unique within the agent.
	DisplayName pulumi.StringPtrInput
	// Enables fuzzy entity extraction during classification.
	EnableFuzzyExtraction pulumi.BoolPtrInput
	// The collection of entity entries associated with the entity type.
	// Structure is documented below.
	Entities CxEntityTypeEntityArrayInput
	// Collection of exceptional words and phrases that shouldn't be matched. For example, if you have a size entity type with entry giant(an adjective), you might consider adding giants(a noun) as an exclusion.
	// If the kind of entity type is KIND_MAP, then the phrases specified by entities and excluded phrases should be mutually exclusive.
	// Structure is documented below.
	ExcludedPhrases CxEntityTypeExcludedPhraseArrayInput
	// Indicates whether the entity type can be automatically expanded.
	// * KIND_MAP: Map entity types allow mapping of a group of synonyms to a canonical value.
	// * KIND_LIST: List entity types contain a set of entries that do not map to canonical values. However, list entity types can contain references to other entity types (with or without aliases).
	// * KIND_REGEXP: Regexp entity types allow to specify regular expressions in entries values.
	//   Possible values are `KIND_MAP`, `KIND_LIST`, and `KIND_REGEXP`.
	Kind pulumi.StringPtrInput
	// The language of the following fields in entityType:
	// EntityType.entities.value
	// EntityType.entities.synonyms
	// EntityType.excluded_phrases.value
	// If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
	LanguageCode pulumi.StringPtrInput
	// The unique identifier of the entity type. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent
	// ID>/entityTypes/<Entity Type ID>.
	Name pulumi.StringPtrInput
	// The agent to create a entity type for.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	Parent pulumi.StringPtrInput
	// Indicates whether parameters of the entity type should be redacted in log. If redaction is enabled, page parameters and intent parameters referring to the entity type will be replaced by parameter name when logging.
	Redact pulumi.BoolPtrInput
}

func (CxEntityTypeState) ElementType() reflect.Type {
	return reflect.TypeOf((*cxEntityTypeState)(nil)).Elem()
}

type cxEntityTypeArgs struct {
	// Represents kinds of entities.
	// * AUTO_EXPANSION_MODE_UNSPECIFIED: Auto expansion disabled for the entity.
	// * AUTO_EXPANSION_MODE_DEFAULT: Allows an agent to recognize values that have not been explicitly listed in the entity.
	//   Possible values are `AUTO_EXPANSION_MODE_DEFAULT` and `AUTO_EXPANSION_MODE_UNSPECIFIED`.
	AutoExpansionMode *string `pulumi:"autoExpansionMode"`
	// The human-readable name of the entity type, unique within the agent.
	DisplayName string `pulumi:"displayName"`
	// Enables fuzzy entity extraction during classification.
	EnableFuzzyExtraction *bool `pulumi:"enableFuzzyExtraction"`
	// The collection of entity entries associated with the entity type.
	// Structure is documented below.
	Entities []CxEntityTypeEntity `pulumi:"entities"`
	// Collection of exceptional words and phrases that shouldn't be matched. For example, if you have a size entity type with entry giant(an adjective), you might consider adding giants(a noun) as an exclusion.
	// If the kind of entity type is KIND_MAP, then the phrases specified by entities and excluded phrases should be mutually exclusive.
	// Structure is documented below.
	ExcludedPhrases []CxEntityTypeExcludedPhrase `pulumi:"excludedPhrases"`
	// Indicates whether the entity type can be automatically expanded.
	// * KIND_MAP: Map entity types allow mapping of a group of synonyms to a canonical value.
	// * KIND_LIST: List entity types contain a set of entries that do not map to canonical values. However, list entity types can contain references to other entity types (with or without aliases).
	// * KIND_REGEXP: Regexp entity types allow to specify regular expressions in entries values.
	//   Possible values are `KIND_MAP`, `KIND_LIST`, and `KIND_REGEXP`.
	Kind string `pulumi:"kind"`
	// The language of the following fields in entityType:
	// EntityType.entities.value
	// EntityType.entities.synonyms
	// EntityType.excluded_phrases.value
	// If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
	LanguageCode *string `pulumi:"languageCode"`
	// The agent to create a entity type for.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	Parent *string `pulumi:"parent"`
	// Indicates whether parameters of the entity type should be redacted in log. If redaction is enabled, page parameters and intent parameters referring to the entity type will be replaced by parameter name when logging.
	Redact *bool `pulumi:"redact"`
}

// The set of arguments for constructing a CxEntityType resource.
type CxEntityTypeArgs struct {
	// Represents kinds of entities.
	// * AUTO_EXPANSION_MODE_UNSPECIFIED: Auto expansion disabled for the entity.
	// * AUTO_EXPANSION_MODE_DEFAULT: Allows an agent to recognize values that have not been explicitly listed in the entity.
	//   Possible values are `AUTO_EXPANSION_MODE_DEFAULT` and `AUTO_EXPANSION_MODE_UNSPECIFIED`.
	AutoExpansionMode pulumi.StringPtrInput
	// The human-readable name of the entity type, unique within the agent.
	DisplayName pulumi.StringInput
	// Enables fuzzy entity extraction during classification.
	EnableFuzzyExtraction pulumi.BoolPtrInput
	// The collection of entity entries associated with the entity type.
	// Structure is documented below.
	Entities CxEntityTypeEntityArrayInput
	// Collection of exceptional words and phrases that shouldn't be matched. For example, if you have a size entity type with entry giant(an adjective), you might consider adding giants(a noun) as an exclusion.
	// If the kind of entity type is KIND_MAP, then the phrases specified by entities and excluded phrases should be mutually exclusive.
	// Structure is documented below.
	ExcludedPhrases CxEntityTypeExcludedPhraseArrayInput
	// Indicates whether the entity type can be automatically expanded.
	// * KIND_MAP: Map entity types allow mapping of a group of synonyms to a canonical value.
	// * KIND_LIST: List entity types contain a set of entries that do not map to canonical values. However, list entity types can contain references to other entity types (with or without aliases).
	// * KIND_REGEXP: Regexp entity types allow to specify regular expressions in entries values.
	//   Possible values are `KIND_MAP`, `KIND_LIST`, and `KIND_REGEXP`.
	Kind pulumi.StringInput
	// The language of the following fields in entityType:
	// EntityType.entities.value
	// EntityType.entities.synonyms
	// EntityType.excluded_phrases.value
	// If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
	LanguageCode pulumi.StringPtrInput
	// The agent to create a entity type for.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	Parent pulumi.StringPtrInput
	// Indicates whether parameters of the entity type should be redacted in log. If redaction is enabled, page parameters and intent parameters referring to the entity type will be replaced by parameter name when logging.
	Redact pulumi.BoolPtrInput
}

func (CxEntityTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cxEntityTypeArgs)(nil)).Elem()
}

type CxEntityTypeInput interface {
	pulumi.Input

	ToCxEntityTypeOutput() CxEntityTypeOutput
	ToCxEntityTypeOutputWithContext(ctx context.Context) CxEntityTypeOutput
}

func (*CxEntityType) ElementType() reflect.Type {
	return reflect.TypeOf((**CxEntityType)(nil)).Elem()
}

func (i *CxEntityType) ToCxEntityTypeOutput() CxEntityTypeOutput {
	return i.ToCxEntityTypeOutputWithContext(context.Background())
}

func (i *CxEntityType) ToCxEntityTypeOutputWithContext(ctx context.Context) CxEntityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CxEntityTypeOutput)
}

// CxEntityTypeArrayInput is an input type that accepts CxEntityTypeArray and CxEntityTypeArrayOutput values.
// You can construct a concrete instance of `CxEntityTypeArrayInput` via:
//
//	CxEntityTypeArray{ CxEntityTypeArgs{...} }
type CxEntityTypeArrayInput interface {
	pulumi.Input

	ToCxEntityTypeArrayOutput() CxEntityTypeArrayOutput
	ToCxEntityTypeArrayOutputWithContext(context.Context) CxEntityTypeArrayOutput
}

type CxEntityTypeArray []CxEntityTypeInput

func (CxEntityTypeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CxEntityType)(nil)).Elem()
}

func (i CxEntityTypeArray) ToCxEntityTypeArrayOutput() CxEntityTypeArrayOutput {
	return i.ToCxEntityTypeArrayOutputWithContext(context.Background())
}

func (i CxEntityTypeArray) ToCxEntityTypeArrayOutputWithContext(ctx context.Context) CxEntityTypeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CxEntityTypeArrayOutput)
}

// CxEntityTypeMapInput is an input type that accepts CxEntityTypeMap and CxEntityTypeMapOutput values.
// You can construct a concrete instance of `CxEntityTypeMapInput` via:
//
//	CxEntityTypeMap{ "key": CxEntityTypeArgs{...} }
type CxEntityTypeMapInput interface {
	pulumi.Input

	ToCxEntityTypeMapOutput() CxEntityTypeMapOutput
	ToCxEntityTypeMapOutputWithContext(context.Context) CxEntityTypeMapOutput
}

type CxEntityTypeMap map[string]CxEntityTypeInput

func (CxEntityTypeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CxEntityType)(nil)).Elem()
}

func (i CxEntityTypeMap) ToCxEntityTypeMapOutput() CxEntityTypeMapOutput {
	return i.ToCxEntityTypeMapOutputWithContext(context.Background())
}

func (i CxEntityTypeMap) ToCxEntityTypeMapOutputWithContext(ctx context.Context) CxEntityTypeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CxEntityTypeMapOutput)
}

type CxEntityTypeOutput struct{ *pulumi.OutputState }

func (CxEntityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CxEntityType)(nil)).Elem()
}

func (o CxEntityTypeOutput) ToCxEntityTypeOutput() CxEntityTypeOutput {
	return o
}

func (o CxEntityTypeOutput) ToCxEntityTypeOutputWithContext(ctx context.Context) CxEntityTypeOutput {
	return o
}

// Represents kinds of entities.
//   - AUTO_EXPANSION_MODE_UNSPECIFIED: Auto expansion disabled for the entity.
//   - AUTO_EXPANSION_MODE_DEFAULT: Allows an agent to recognize values that have not been explicitly listed in the entity.
//     Possible values are `AUTO_EXPANSION_MODE_DEFAULT` and `AUTO_EXPANSION_MODE_UNSPECIFIED`.
func (o CxEntityTypeOutput) AutoExpansionMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CxEntityType) pulumi.StringPtrOutput { return v.AutoExpansionMode }).(pulumi.StringPtrOutput)
}

// The human-readable name of the entity type, unique within the agent.
func (o CxEntityTypeOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *CxEntityType) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Enables fuzzy entity extraction during classification.
func (o CxEntityTypeOutput) EnableFuzzyExtraction() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CxEntityType) pulumi.BoolPtrOutput { return v.EnableFuzzyExtraction }).(pulumi.BoolPtrOutput)
}

// The collection of entity entries associated with the entity type.
// Structure is documented below.
func (o CxEntityTypeOutput) Entities() CxEntityTypeEntityArrayOutput {
	return o.ApplyT(func(v *CxEntityType) CxEntityTypeEntityArrayOutput { return v.Entities }).(CxEntityTypeEntityArrayOutput)
}

// Collection of exceptional words and phrases that shouldn't be matched. For example, if you have a size entity type with entry giant(an adjective), you might consider adding giants(a noun) as an exclusion.
// If the kind of entity type is KIND_MAP, then the phrases specified by entities and excluded phrases should be mutually exclusive.
// Structure is documented below.
func (o CxEntityTypeOutput) ExcludedPhrases() CxEntityTypeExcludedPhraseArrayOutput {
	return o.ApplyT(func(v *CxEntityType) CxEntityTypeExcludedPhraseArrayOutput { return v.ExcludedPhrases }).(CxEntityTypeExcludedPhraseArrayOutput)
}

// Indicates whether the entity type can be automatically expanded.
//   - KIND_MAP: Map entity types allow mapping of a group of synonyms to a canonical value.
//   - KIND_LIST: List entity types contain a set of entries that do not map to canonical values. However, list entity types can contain references to other entity types (with or without aliases).
//   - KIND_REGEXP: Regexp entity types allow to specify regular expressions in entries values.
//     Possible values are `KIND_MAP`, `KIND_LIST`, and `KIND_REGEXP`.
func (o CxEntityTypeOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *CxEntityType) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// The language of the following fields in entityType:
// EntityType.entities.value
// EntityType.entities.synonyms
// EntityType.excluded_phrases.value
// If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
func (o CxEntityTypeOutput) LanguageCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CxEntityType) pulumi.StringPtrOutput { return v.LanguageCode }).(pulumi.StringPtrOutput)
}

// The unique identifier of the entity type. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent
// ID>/entityTypes/<Entity Type ID>.
func (o CxEntityTypeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CxEntityType) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The agent to create a entity type for.
// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
func (o CxEntityTypeOutput) Parent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CxEntityType) pulumi.StringPtrOutput { return v.Parent }).(pulumi.StringPtrOutput)
}

// Indicates whether parameters of the entity type should be redacted in log. If redaction is enabled, page parameters and intent parameters referring to the entity type will be replaced by parameter name when logging.
func (o CxEntityTypeOutput) Redact() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CxEntityType) pulumi.BoolPtrOutput { return v.Redact }).(pulumi.BoolPtrOutput)
}

type CxEntityTypeArrayOutput struct{ *pulumi.OutputState }

func (CxEntityTypeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CxEntityType)(nil)).Elem()
}

func (o CxEntityTypeArrayOutput) ToCxEntityTypeArrayOutput() CxEntityTypeArrayOutput {
	return o
}

func (o CxEntityTypeArrayOutput) ToCxEntityTypeArrayOutputWithContext(ctx context.Context) CxEntityTypeArrayOutput {
	return o
}

func (o CxEntityTypeArrayOutput) Index(i pulumi.IntInput) CxEntityTypeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CxEntityType {
		return vs[0].([]*CxEntityType)[vs[1].(int)]
	}).(CxEntityTypeOutput)
}

type CxEntityTypeMapOutput struct{ *pulumi.OutputState }

func (CxEntityTypeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CxEntityType)(nil)).Elem()
}

func (o CxEntityTypeMapOutput) ToCxEntityTypeMapOutput() CxEntityTypeMapOutput {
	return o
}

func (o CxEntityTypeMapOutput) ToCxEntityTypeMapOutputWithContext(ctx context.Context) CxEntityTypeMapOutput {
	return o
}

func (o CxEntityTypeMapOutput) MapIndex(k pulumi.StringInput) CxEntityTypeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CxEntityType {
		return vs[0].(map[string]*CxEntityType)[vs[1].(string)]
	}).(CxEntityTypeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CxEntityTypeInput)(nil)).Elem(), &CxEntityType{})
	pulumi.RegisterInputType(reflect.TypeOf((*CxEntityTypeArrayInput)(nil)).Elem(), CxEntityTypeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CxEntityTypeMapInput)(nil)).Elem(), CxEntityTypeMap{})
	pulumi.RegisterOutputType(CxEntityTypeOutput{})
	pulumi.RegisterOutputType(CxEntityTypeArrayOutput{})
	pulumi.RegisterOutputType(CxEntityTypeMapOutput{})
}
