// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dataloss

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An inspect job template.
//
// To get more information about InspectTemplate, see:
//
// * [API documentation](https://cloud.google.com/dlp/docs/reference/rest/v2/projects.inspectTemplates)
// * How-to Guides
//   - [Official Documentation](https://cloud.google.com/dlp/docs/creating-templates-inspect)
//
// ## Example Usage
// ### Dlp Inspect Template Custom Type
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/dataloss"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dataloss.NewPreventionInspectTemplate(ctx, "custom", &dataloss.PreventionInspectTemplateArgs{
//				Description: pulumi.String("My description"),
//				DisplayName: pulumi.String("display_name"),
//				InspectConfig: &dataloss.PreventionInspectTemplateInspectConfigArgs{
//					CustomInfoTypes: dataloss.PreventionInspectTemplateInspectConfigCustomInfoTypeArray{
//						&dataloss.PreventionInspectTemplateInspectConfigCustomInfoTypeArgs{
//							InfoType: &dataloss.PreventionInspectTemplateInspectConfigCustomInfoTypeInfoTypeArgs{
//								Name: pulumi.String("MY_CUSTOM_TYPE"),
//							},
//							Likelihood: pulumi.String("UNLIKELY"),
//							Regex: &dataloss.PreventionInspectTemplateInspectConfigCustomInfoTypeRegexArgs{
//								Pattern: pulumi.String("test*"),
//							},
//						},
//					},
//					InfoTypes: dataloss.PreventionInspectTemplateInspectConfigInfoTypeArray{
//						&dataloss.PreventionInspectTemplateInspectConfigInfoTypeArgs{
//							Name: pulumi.String("EMAIL_ADDRESS"),
//						},
//					},
//					Limits: &dataloss.PreventionInspectTemplateInspectConfigLimitsArgs{
//						MaxFindingsPerItem:    pulumi.Int(10),
//						MaxFindingsPerRequest: pulumi.Int(50),
//					},
//					MinLikelihood: pulumi.String("UNLIKELY"),
//					RuleSets: dataloss.PreventionInspectTemplateInspectConfigRuleSetArray{
//						&dataloss.PreventionInspectTemplateInspectConfigRuleSetArgs{
//							InfoTypes: dataloss.PreventionInspectTemplateInspectConfigRuleSetInfoTypeArray{
//								&dataloss.PreventionInspectTemplateInspectConfigRuleSetInfoTypeArgs{
//									Name: pulumi.String("EMAIL_ADDRESS"),
//								},
//							},
//							Rules: dataloss.PreventionInspectTemplateInspectConfigRuleSetRuleArray{
//								&dataloss.PreventionInspectTemplateInspectConfigRuleSetRuleArgs{
//									ExclusionRule: &dataloss.PreventionInspectTemplateInspectConfigRuleSetRuleExclusionRuleArgs{
//										MatchingType: pulumi.String("MATCHING_TYPE_FULL_MATCH"),
//										Regex: &dataloss.PreventionInspectTemplateInspectConfigRuleSetRuleExclusionRuleRegexArgs{
//											Pattern: pulumi.String(".+@example.com"),
//										},
//									},
//								},
//							},
//						},
//						&dataloss.PreventionInspectTemplateInspectConfigRuleSetArgs{
//							InfoTypes: dataloss.PreventionInspectTemplateInspectConfigRuleSetInfoTypeArray{
//								&dataloss.PreventionInspectTemplateInspectConfigRuleSetInfoTypeArgs{
//									Name: pulumi.String("MY_CUSTOM_TYPE"),
//								},
//							},
//							Rules: dataloss.PreventionInspectTemplateInspectConfigRuleSetRuleArray{
//								&dataloss.PreventionInspectTemplateInspectConfigRuleSetRuleArgs{
//									HotwordRule: &dataloss.PreventionInspectTemplateInspectConfigRuleSetRuleHotwordRuleArgs{
//										HotwordRegex: &dataloss.PreventionInspectTemplateInspectConfigRuleSetRuleHotwordRuleHotwordRegexArgs{
//											Pattern: pulumi.String("example*"),
//										},
//										LikelihoodAdjustment: &dataloss.PreventionInspectTemplateInspectConfigRuleSetRuleHotwordRuleLikelihoodAdjustmentArgs{
//											FixedLikelihood: pulumi.String("VERY_LIKELY"),
//										},
//										Proximity: &dataloss.PreventionInspectTemplateInspectConfigRuleSetRuleHotwordRuleProximityArgs{
//											WindowBefore: pulumi.Int(50),
//										},
//									},
//								},
//							},
//						},
//					},
//				},
//				Parent: pulumi.String("projects/my-project-name"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Dlp Inspect Template Custom Type Surrogate
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/dataloss"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dataloss.NewPreventionInspectTemplate(ctx, "customTypeSurrogate", &dataloss.PreventionInspectTemplateArgs{
//				Description: pulumi.String("My description"),
//				DisplayName: pulumi.String("display_name"),
//				InspectConfig: &dataloss.PreventionInspectTemplateInspectConfigArgs{
//					CustomInfoTypes: dataloss.PreventionInspectTemplateInspectConfigCustomInfoTypeArray{
//						&dataloss.PreventionInspectTemplateInspectConfigCustomInfoTypeArgs{
//							InfoType: &dataloss.PreventionInspectTemplateInspectConfigCustomInfoTypeInfoTypeArgs{
//								Name: pulumi.String("MY_CUSTOM_TYPE"),
//							},
//							Likelihood:    pulumi.String("UNLIKELY"),
//							SurrogateType: nil,
//						},
//					},
//					InfoTypes: dataloss.PreventionInspectTemplateInspectConfigInfoTypeArray{
//						&dataloss.PreventionInspectTemplateInspectConfigInfoTypeArgs{
//							Name: pulumi.String("EMAIL_ADDRESS"),
//						},
//					},
//					Limits: &dataloss.PreventionInspectTemplateInspectConfigLimitsArgs{
//						MaxFindingsPerItem:    pulumi.Int(10),
//						MaxFindingsPerRequest: pulumi.Int(50),
//					},
//					MinLikelihood: pulumi.String("UNLIKELY"),
//					RuleSets: dataloss.PreventionInspectTemplateInspectConfigRuleSetArray{
//						&dataloss.PreventionInspectTemplateInspectConfigRuleSetArgs{
//							InfoTypes: dataloss.PreventionInspectTemplateInspectConfigRuleSetInfoTypeArray{
//								&dataloss.PreventionInspectTemplateInspectConfigRuleSetInfoTypeArgs{
//									Name: pulumi.String("EMAIL_ADDRESS"),
//								},
//							},
//							Rules: dataloss.PreventionInspectTemplateInspectConfigRuleSetRuleArray{
//								&dataloss.PreventionInspectTemplateInspectConfigRuleSetRuleArgs{
//									ExclusionRule: &dataloss.PreventionInspectTemplateInspectConfigRuleSetRuleExclusionRuleArgs{
//										MatchingType: pulumi.String("MATCHING_TYPE_FULL_MATCH"),
//										Regex: &dataloss.PreventionInspectTemplateInspectConfigRuleSetRuleExclusionRuleRegexArgs{
//											Pattern: pulumi.String(".+@example.com"),
//										},
//									},
//								},
//							},
//						},
//						&dataloss.PreventionInspectTemplateInspectConfigRuleSetArgs{
//							InfoTypes: dataloss.PreventionInspectTemplateInspectConfigRuleSetInfoTypeArray{
//								&dataloss.PreventionInspectTemplateInspectConfigRuleSetInfoTypeArgs{
//									Name: pulumi.String("MY_CUSTOM_TYPE"),
//								},
//							},
//							Rules: dataloss.PreventionInspectTemplateInspectConfigRuleSetRuleArray{
//								&dataloss.PreventionInspectTemplateInspectConfigRuleSetRuleArgs{
//									HotwordRule: &dataloss.PreventionInspectTemplateInspectConfigRuleSetRuleHotwordRuleArgs{
//										HotwordRegex: &dataloss.PreventionInspectTemplateInspectConfigRuleSetRuleHotwordRuleHotwordRegexArgs{
//											Pattern: pulumi.String("example*"),
//										},
//										LikelihoodAdjustment: &dataloss.PreventionInspectTemplateInspectConfigRuleSetRuleHotwordRuleLikelihoodAdjustmentArgs{
//											FixedLikelihood: pulumi.String("VERY_LIKELY"),
//										},
//										Proximity: &dataloss.PreventionInspectTemplateInspectConfigRuleSetRuleHotwordRuleProximityArgs{
//											WindowBefore: pulumi.Int(50),
//										},
//									},
//								},
//							},
//						},
//					},
//				},
//				Parent: pulumi.String("projects/my-project-name"),
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
// # InspectTemplate can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:dataloss/preventionInspectTemplate:PreventionInspectTemplate default {{parent}}/inspectTemplates/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:dataloss/preventionInspectTemplate:PreventionInspectTemplate default {{parent}}/{{name}}
//
// ```
type PreventionInspectTemplate struct {
	pulumi.CustomResourceState

	// A description of the inspect template.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// User set display name of the inspect template.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The core content of the template.
	// Structure is documented below.
	InspectConfig PreventionInspectTemplateInspectConfigPtrOutput `pulumi:"inspectConfig"`
	// Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed
	// at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
	// (Required)
	// Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed
	// at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
	// (Required)
	// Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed
	// at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
	// (Required)
	// Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed
	// at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
	// (Required)
	// Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names
	// listed at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
	// (Required)
	// Resource name of the requested StoredInfoType, for example `organizations/433245324/storedInfoTypes/432452342`
	// or `projects/project-id/storedInfoTypes/432452342`.
	Name pulumi.StringOutput `pulumi:"name"`
	// The parent of the inspect template in any of the following formats:
	// * `projects/{{project}}`
	// * `projects/{{project}}/locations/{{location}}`
	// * `organizations/{{organization_id}}`
	// * `organizations/{{organization_id}}/locations/{{location}}`
	Parent pulumi.StringOutput `pulumi:"parent"`
}

// NewPreventionInspectTemplate registers a new resource with the given unique name, arguments, and options.
func NewPreventionInspectTemplate(ctx *pulumi.Context,
	name string, args *PreventionInspectTemplateArgs, opts ...pulumi.ResourceOption) (*PreventionInspectTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Parent == nil {
		return nil, errors.New("invalid value for required argument 'Parent'")
	}
	var resource PreventionInspectTemplate
	err := ctx.RegisterResource("gcp:dataloss/preventionInspectTemplate:PreventionInspectTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPreventionInspectTemplate gets an existing PreventionInspectTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPreventionInspectTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PreventionInspectTemplateState, opts ...pulumi.ResourceOption) (*PreventionInspectTemplate, error) {
	var resource PreventionInspectTemplate
	err := ctx.ReadResource("gcp:dataloss/preventionInspectTemplate:PreventionInspectTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PreventionInspectTemplate resources.
type preventionInspectTemplateState struct {
	// A description of the inspect template.
	Description *string `pulumi:"description"`
	// User set display name of the inspect template.
	DisplayName *string `pulumi:"displayName"`
	// The core content of the template.
	// Structure is documented below.
	InspectConfig *PreventionInspectTemplateInspectConfig `pulumi:"inspectConfig"`
	// Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed
	// at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
	// (Required)
	// Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed
	// at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
	// (Required)
	// Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed
	// at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
	// (Required)
	// Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed
	// at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
	// (Required)
	// Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names
	// listed at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
	// (Required)
	// Resource name of the requested StoredInfoType, for example `organizations/433245324/storedInfoTypes/432452342`
	// or `projects/project-id/storedInfoTypes/432452342`.
	Name *string `pulumi:"name"`
	// The parent of the inspect template in any of the following formats:
	// * `projects/{{project}}`
	// * `projects/{{project}}/locations/{{location}}`
	// * `organizations/{{organization_id}}`
	// * `organizations/{{organization_id}}/locations/{{location}}`
	Parent *string `pulumi:"parent"`
}

type PreventionInspectTemplateState struct {
	// A description of the inspect template.
	Description pulumi.StringPtrInput
	// User set display name of the inspect template.
	DisplayName pulumi.StringPtrInput
	// The core content of the template.
	// Structure is documented below.
	InspectConfig PreventionInspectTemplateInspectConfigPtrInput
	// Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed
	// at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
	// (Required)
	// Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed
	// at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
	// (Required)
	// Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed
	// at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
	// (Required)
	// Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed
	// at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
	// (Required)
	// Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names
	// listed at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
	// (Required)
	// Resource name of the requested StoredInfoType, for example `organizations/433245324/storedInfoTypes/432452342`
	// or `projects/project-id/storedInfoTypes/432452342`.
	Name pulumi.StringPtrInput
	// The parent of the inspect template in any of the following formats:
	// * `projects/{{project}}`
	// * `projects/{{project}}/locations/{{location}}`
	// * `organizations/{{organization_id}}`
	// * `organizations/{{organization_id}}/locations/{{location}}`
	Parent pulumi.StringPtrInput
}

func (PreventionInspectTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*preventionInspectTemplateState)(nil)).Elem()
}

type preventionInspectTemplateArgs struct {
	// A description of the inspect template.
	Description *string `pulumi:"description"`
	// User set display name of the inspect template.
	DisplayName *string `pulumi:"displayName"`
	// The core content of the template.
	// Structure is documented below.
	InspectConfig *PreventionInspectTemplateInspectConfig `pulumi:"inspectConfig"`
	// The parent of the inspect template in any of the following formats:
	// * `projects/{{project}}`
	// * `projects/{{project}}/locations/{{location}}`
	// * `organizations/{{organization_id}}`
	// * `organizations/{{organization_id}}/locations/{{location}}`
	Parent string `pulumi:"parent"`
}

// The set of arguments for constructing a PreventionInspectTemplate resource.
type PreventionInspectTemplateArgs struct {
	// A description of the inspect template.
	Description pulumi.StringPtrInput
	// User set display name of the inspect template.
	DisplayName pulumi.StringPtrInput
	// The core content of the template.
	// Structure is documented below.
	InspectConfig PreventionInspectTemplateInspectConfigPtrInput
	// The parent of the inspect template in any of the following formats:
	// * `projects/{{project}}`
	// * `projects/{{project}}/locations/{{location}}`
	// * `organizations/{{organization_id}}`
	// * `organizations/{{organization_id}}/locations/{{location}}`
	Parent pulumi.StringInput
}

func (PreventionInspectTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*preventionInspectTemplateArgs)(nil)).Elem()
}

type PreventionInspectTemplateInput interface {
	pulumi.Input

	ToPreventionInspectTemplateOutput() PreventionInspectTemplateOutput
	ToPreventionInspectTemplateOutputWithContext(ctx context.Context) PreventionInspectTemplateOutput
}

func (*PreventionInspectTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**PreventionInspectTemplate)(nil)).Elem()
}

func (i *PreventionInspectTemplate) ToPreventionInspectTemplateOutput() PreventionInspectTemplateOutput {
	return i.ToPreventionInspectTemplateOutputWithContext(context.Background())
}

func (i *PreventionInspectTemplate) ToPreventionInspectTemplateOutputWithContext(ctx context.Context) PreventionInspectTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PreventionInspectTemplateOutput)
}

// PreventionInspectTemplateArrayInput is an input type that accepts PreventionInspectTemplateArray and PreventionInspectTemplateArrayOutput values.
// You can construct a concrete instance of `PreventionInspectTemplateArrayInput` via:
//
//	PreventionInspectTemplateArray{ PreventionInspectTemplateArgs{...} }
type PreventionInspectTemplateArrayInput interface {
	pulumi.Input

	ToPreventionInspectTemplateArrayOutput() PreventionInspectTemplateArrayOutput
	ToPreventionInspectTemplateArrayOutputWithContext(context.Context) PreventionInspectTemplateArrayOutput
}

type PreventionInspectTemplateArray []PreventionInspectTemplateInput

func (PreventionInspectTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PreventionInspectTemplate)(nil)).Elem()
}

func (i PreventionInspectTemplateArray) ToPreventionInspectTemplateArrayOutput() PreventionInspectTemplateArrayOutput {
	return i.ToPreventionInspectTemplateArrayOutputWithContext(context.Background())
}

func (i PreventionInspectTemplateArray) ToPreventionInspectTemplateArrayOutputWithContext(ctx context.Context) PreventionInspectTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PreventionInspectTemplateArrayOutput)
}

// PreventionInspectTemplateMapInput is an input type that accepts PreventionInspectTemplateMap and PreventionInspectTemplateMapOutput values.
// You can construct a concrete instance of `PreventionInspectTemplateMapInput` via:
//
//	PreventionInspectTemplateMap{ "key": PreventionInspectTemplateArgs{...} }
type PreventionInspectTemplateMapInput interface {
	pulumi.Input

	ToPreventionInspectTemplateMapOutput() PreventionInspectTemplateMapOutput
	ToPreventionInspectTemplateMapOutputWithContext(context.Context) PreventionInspectTemplateMapOutput
}

type PreventionInspectTemplateMap map[string]PreventionInspectTemplateInput

func (PreventionInspectTemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PreventionInspectTemplate)(nil)).Elem()
}

func (i PreventionInspectTemplateMap) ToPreventionInspectTemplateMapOutput() PreventionInspectTemplateMapOutput {
	return i.ToPreventionInspectTemplateMapOutputWithContext(context.Background())
}

func (i PreventionInspectTemplateMap) ToPreventionInspectTemplateMapOutputWithContext(ctx context.Context) PreventionInspectTemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PreventionInspectTemplateMapOutput)
}

type PreventionInspectTemplateOutput struct{ *pulumi.OutputState }

func (PreventionInspectTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PreventionInspectTemplate)(nil)).Elem()
}

func (o PreventionInspectTemplateOutput) ToPreventionInspectTemplateOutput() PreventionInspectTemplateOutput {
	return o
}

func (o PreventionInspectTemplateOutput) ToPreventionInspectTemplateOutputWithContext(ctx context.Context) PreventionInspectTemplateOutput {
	return o
}

// A description of the inspect template.
func (o PreventionInspectTemplateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PreventionInspectTemplate) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// User set display name of the inspect template.
func (o PreventionInspectTemplateOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PreventionInspectTemplate) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The core content of the template.
// Structure is documented below.
func (o PreventionInspectTemplateOutput) InspectConfig() PreventionInspectTemplateInspectConfigPtrOutput {
	return o.ApplyT(func(v *PreventionInspectTemplate) PreventionInspectTemplateInspectConfigPtrOutput {
		return v.InspectConfig
	}).(PreventionInspectTemplateInspectConfigPtrOutput)
}

// Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed
// at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
// (Required)
// Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed
// at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
// (Required)
// Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed
// at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
// (Required)
// Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names listed
// at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
// (Required)
// Name of the information type. Either a name of your choosing when creating a CustomInfoType, or one of the names
// listed at https://cloud.google.com/dlp/docs/infotypes-reference when specifying a built-in type.
// (Required)
// Resource name of the requested StoredInfoType, for example `organizations/433245324/storedInfoTypes/432452342`
// or `projects/project-id/storedInfoTypes/432452342`.
func (o PreventionInspectTemplateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PreventionInspectTemplate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The parent of the inspect template in any of the following formats:
// * `projects/{{project}}`
// * `projects/{{project}}/locations/{{location}}`
// * `organizations/{{organization_id}}`
// * `organizations/{{organization_id}}/locations/{{location}}`
func (o PreventionInspectTemplateOutput) Parent() pulumi.StringOutput {
	return o.ApplyT(func(v *PreventionInspectTemplate) pulumi.StringOutput { return v.Parent }).(pulumi.StringOutput)
}

type PreventionInspectTemplateArrayOutput struct{ *pulumi.OutputState }

func (PreventionInspectTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*PreventionInspectTemplate)(nil)).Elem()
}

func (o PreventionInspectTemplateArrayOutput) ToPreventionInspectTemplateArrayOutput() PreventionInspectTemplateArrayOutput {
	return o
}

func (o PreventionInspectTemplateArrayOutput) ToPreventionInspectTemplateArrayOutputWithContext(ctx context.Context) PreventionInspectTemplateArrayOutput {
	return o
}

func (o PreventionInspectTemplateArrayOutput) Index(i pulumi.IntInput) PreventionInspectTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *PreventionInspectTemplate {
		return vs[0].([]*PreventionInspectTemplate)[vs[1].(int)]
	}).(PreventionInspectTemplateOutput)
}

type PreventionInspectTemplateMapOutput struct{ *pulumi.OutputState }

func (PreventionInspectTemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*PreventionInspectTemplate)(nil)).Elem()
}

func (o PreventionInspectTemplateMapOutput) ToPreventionInspectTemplateMapOutput() PreventionInspectTemplateMapOutput {
	return o
}

func (o PreventionInspectTemplateMapOutput) ToPreventionInspectTemplateMapOutputWithContext(ctx context.Context) PreventionInspectTemplateMapOutput {
	return o
}

func (o PreventionInspectTemplateMapOutput) MapIndex(k pulumi.StringInput) PreventionInspectTemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *PreventionInspectTemplate {
		return vs[0].(map[string]*PreventionInspectTemplate)[vs[1].(string)]
	}).(PreventionInspectTemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*PreventionInspectTemplateInput)(nil)).Elem(), &PreventionInspectTemplate{})
	pulumi.RegisterInputType(reflect.TypeOf((*PreventionInspectTemplateArrayInput)(nil)).Elem(), PreventionInspectTemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PreventionInspectTemplateMapInput)(nil)).Elem(), PreventionInspectTemplateMap{})
	pulumi.RegisterOutputType(PreventionInspectTemplateOutput{})
	pulumi.RegisterOutputType(PreventionInspectTemplateArrayOutput{})
	pulumi.RegisterOutputType(PreventionInspectTemplateMapOutput{})
}
