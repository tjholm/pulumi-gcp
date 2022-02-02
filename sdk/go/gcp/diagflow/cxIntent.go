// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package diagflow

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An intent represents a user's intent to interact with a conversational agent.
//
// To get more information about Intent, see:
//
// * [API documentation](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.intents)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/dialogflow/cx/docs)
//
// ## Example Usage
// ### Dialogflowcx Intent Full
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/diagflow"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		agent, err := diagflow.NewCxAgent(ctx, "agent", &diagflow.CxAgentArgs{
// 			DisplayName:         pulumi.String("dialogflowcx-agent"),
// 			Location:            pulumi.String("global"),
// 			DefaultLanguageCode: pulumi.String("en"),
// 			SupportedLanguageCodes: pulumi.StringArray{
// 				pulumi.String("fr"),
// 				pulumi.String("de"),
// 				pulumi.String("es"),
// 			},
// 			TimeZone:                 pulumi.String("America/New_York"),
// 			Description:              pulumi.String("Example description."),
// 			AvatarUri:                pulumi.String("https://cloud.google.com/_static/images/cloud/icons/favicons/onecloud/super_cloud.png"),
// 			EnableStackdriverLogging: pulumi.Bool(true),
// 			EnableSpellCorrection:    pulumi.Bool(true),
// 			SpeechToTextSettings: &diagflow.CxAgentSpeechToTextSettingsArgs{
// 				EnableSpeechAdaptation: pulumi.Bool(true),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = diagflow.NewCxIntent(ctx, "basicIntent", &diagflow.CxIntentArgs{
// 			Parent:      agent.ID(),
// 			DisplayName: pulumi.String("Example"),
// 			Priority:    pulumi.Int(1),
// 			Description: pulumi.String("Intent example"),
// 			TrainingPhrases: diagflow.CxIntentTrainingPhraseArray{
// 				&diagflow.CxIntentTrainingPhraseArgs{
// 					Parts: diagflow.CxIntentTrainingPhrasePartArray{
// 						&diagflow.CxIntentTrainingPhrasePartArgs{
// 							Text: pulumi.String("training"),
// 						},
// 						&diagflow.CxIntentTrainingPhrasePartArgs{
// 							Text: pulumi.String("phrase"),
// 						},
// 						&diagflow.CxIntentTrainingPhrasePartArgs{
// 							Text: pulumi.String("example"),
// 						},
// 					},
// 					RepeatCount: pulumi.Int(1),
// 				},
// 			},
// 			Parameters: diagflow.CxIntentParameterArray{
// 				&diagflow.CxIntentParameterArgs{
// 					Id:         pulumi.String("param1"),
// 					EntityType: pulumi.String("projects/-/locations/-/agents/-/entityTypes/sys.date"),
// 				},
// 			},
// 			Labels: pulumi.StringMap{
// 				"label1": pulumi.String("value1"),
// 				"label2": pulumi.String("value2"),
// 			},
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
// Intent can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:diagflow/cxIntent:CxIntent default {{parent}}/intents/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:diagflow/cxIntent:CxIntent default {{parent}}/{{name}}
// ```
type CxIntent struct {
	pulumi.CustomResourceState

	// Human readable description for better understanding an intent like its scope, content, result etc. Maximum character limit: 140 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The human-readable name of the intent, unique within the agent.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Indicates whether this is a fallback intent. Currently only default fallback intent is allowed in the agent, which is added upon agent creation.
	// Adding training phrases to fallback intent is useful in the case of requests that are mistakenly matched, since training phrases assigned to fallback intents act as negative examples that triggers no-match event.
	IsFallback pulumi.BoolPtrOutput `pulumi:"isFallback"`
	// The key/value metadata to label an intent. Labels can contain lowercase letters, digits and the symbols '-' and '_'. International characters are allowed, including letters from unicase alphabets. Keys must start with a letter. Keys and values can be no longer than 63 characters and no more than 128 bytes.
	// Prefix "sys-" is reserved for Dialogflow defined labels. Currently allowed Dialogflow defined labels include: * sys-head * sys-contextual The above labels do not require value. "sys-head" means the intent is a head intent. "sys.contextual" means the intent is a contextual intent.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The language of the following fields in intent:
	// Intent.training_phrases.parts.text
	// If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
	LanguageCode pulumi.StringPtrOutput `pulumi:"languageCode"`
	// The unique identifier of the intent. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent
	// ID>/intents/<Intent ID>.
	Name pulumi.StringOutput `pulumi:"name"`
	// The collection of parameters associated with the intent.
	// Structure is documented below.
	Parameters CxIntentParameterArrayOutput `pulumi:"parameters"`
	// The agent to create an intent for.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	Parent pulumi.StringPtrOutput `pulumi:"parent"`
	// The priority of this intent. Higher numbers represent higher priorities.
	// If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the Normal priority in the console.
	// If the supplied value is negative, the intent is ignored in runtime detect intent requests.
	Priority pulumi.IntPtrOutput `pulumi:"priority"`
	// The collection of training phrases the agent is trained on to identify the intent.
	// Structure is documented below.
	TrainingPhrases CxIntentTrainingPhraseArrayOutput `pulumi:"trainingPhrases"`
}

// NewCxIntent registers a new resource with the given unique name, arguments, and options.
func NewCxIntent(ctx *pulumi.Context,
	name string, args *CxIntentArgs, opts ...pulumi.ResourceOption) (*CxIntent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	var resource CxIntent
	err := ctx.RegisterResource("gcp:diagflow/cxIntent:CxIntent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCxIntent gets an existing CxIntent resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCxIntent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CxIntentState, opts ...pulumi.ResourceOption) (*CxIntent, error) {
	var resource CxIntent
	err := ctx.ReadResource("gcp:diagflow/cxIntent:CxIntent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CxIntent resources.
type cxIntentState struct {
	// Human readable description for better understanding an intent like its scope, content, result etc. Maximum character limit: 140 characters.
	Description *string `pulumi:"description"`
	// The human-readable name of the intent, unique within the agent.
	DisplayName *string `pulumi:"displayName"`
	// Indicates whether this is a fallback intent. Currently only default fallback intent is allowed in the agent, which is added upon agent creation.
	// Adding training phrases to fallback intent is useful in the case of requests that are mistakenly matched, since training phrases assigned to fallback intents act as negative examples that triggers no-match event.
	IsFallback *bool `pulumi:"isFallback"`
	// The key/value metadata to label an intent. Labels can contain lowercase letters, digits and the symbols '-' and '_'. International characters are allowed, including letters from unicase alphabets. Keys must start with a letter. Keys and values can be no longer than 63 characters and no more than 128 bytes.
	// Prefix "sys-" is reserved for Dialogflow defined labels. Currently allowed Dialogflow defined labels include: * sys-head * sys-contextual The above labels do not require value. "sys-head" means the intent is a head intent. "sys.contextual" means the intent is a contextual intent.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels map[string]string `pulumi:"labels"`
	// The language of the following fields in intent:
	// Intent.training_phrases.parts.text
	// If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
	LanguageCode *string `pulumi:"languageCode"`
	// The unique identifier of the intent. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent
	// ID>/intents/<Intent ID>.
	Name *string `pulumi:"name"`
	// The collection of parameters associated with the intent.
	// Structure is documented below.
	Parameters []CxIntentParameter `pulumi:"parameters"`
	// The agent to create an intent for.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	Parent *string `pulumi:"parent"`
	// The priority of this intent. Higher numbers represent higher priorities.
	// If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the Normal priority in the console.
	// If the supplied value is negative, the intent is ignored in runtime detect intent requests.
	Priority *int `pulumi:"priority"`
	// The collection of training phrases the agent is trained on to identify the intent.
	// Structure is documented below.
	TrainingPhrases []CxIntentTrainingPhrase `pulumi:"trainingPhrases"`
}

type CxIntentState struct {
	// Human readable description for better understanding an intent like its scope, content, result etc. Maximum character limit: 140 characters.
	Description pulumi.StringPtrInput
	// The human-readable name of the intent, unique within the agent.
	DisplayName pulumi.StringPtrInput
	// Indicates whether this is a fallback intent. Currently only default fallback intent is allowed in the agent, which is added upon agent creation.
	// Adding training phrases to fallback intent is useful in the case of requests that are mistakenly matched, since training phrases assigned to fallback intents act as negative examples that triggers no-match event.
	IsFallback pulumi.BoolPtrInput
	// The key/value metadata to label an intent. Labels can contain lowercase letters, digits and the symbols '-' and '_'. International characters are allowed, including letters from unicase alphabets. Keys must start with a letter. Keys and values can be no longer than 63 characters and no more than 128 bytes.
	// Prefix "sys-" is reserved for Dialogflow defined labels. Currently allowed Dialogflow defined labels include: * sys-head * sys-contextual The above labels do not require value. "sys-head" means the intent is a head intent. "sys.contextual" means the intent is a contextual intent.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels pulumi.StringMapInput
	// The language of the following fields in intent:
	// Intent.training_phrases.parts.text
	// If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
	LanguageCode pulumi.StringPtrInput
	// The unique identifier of the intent. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent
	// ID>/intents/<Intent ID>.
	Name pulumi.StringPtrInput
	// The collection of parameters associated with the intent.
	// Structure is documented below.
	Parameters CxIntentParameterArrayInput
	// The agent to create an intent for.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	Parent pulumi.StringPtrInput
	// The priority of this intent. Higher numbers represent higher priorities.
	// If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the Normal priority in the console.
	// If the supplied value is negative, the intent is ignored in runtime detect intent requests.
	Priority pulumi.IntPtrInput
	// The collection of training phrases the agent is trained on to identify the intent.
	// Structure is documented below.
	TrainingPhrases CxIntentTrainingPhraseArrayInput
}

func (CxIntentState) ElementType() reflect.Type {
	return reflect.TypeOf((*cxIntentState)(nil)).Elem()
}

type cxIntentArgs struct {
	// Human readable description for better understanding an intent like its scope, content, result etc. Maximum character limit: 140 characters.
	Description *string `pulumi:"description"`
	// The human-readable name of the intent, unique within the agent.
	DisplayName string `pulumi:"displayName"`
	// Indicates whether this is a fallback intent. Currently only default fallback intent is allowed in the agent, which is added upon agent creation.
	// Adding training phrases to fallback intent is useful in the case of requests that are mistakenly matched, since training phrases assigned to fallback intents act as negative examples that triggers no-match event.
	IsFallback *bool `pulumi:"isFallback"`
	// The key/value metadata to label an intent. Labels can contain lowercase letters, digits and the symbols '-' and '_'. International characters are allowed, including letters from unicase alphabets. Keys must start with a letter. Keys and values can be no longer than 63 characters and no more than 128 bytes.
	// Prefix "sys-" is reserved for Dialogflow defined labels. Currently allowed Dialogflow defined labels include: * sys-head * sys-contextual The above labels do not require value. "sys-head" means the intent is a head intent. "sys.contextual" means the intent is a contextual intent.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels map[string]string `pulumi:"labels"`
	// The language of the following fields in intent:
	// Intent.training_phrases.parts.text
	// If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
	LanguageCode *string `pulumi:"languageCode"`
	// The collection of parameters associated with the intent.
	// Structure is documented below.
	Parameters []CxIntentParameter `pulumi:"parameters"`
	// The agent to create an intent for.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	Parent *string `pulumi:"parent"`
	// The priority of this intent. Higher numbers represent higher priorities.
	// If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the Normal priority in the console.
	// If the supplied value is negative, the intent is ignored in runtime detect intent requests.
	Priority *int `pulumi:"priority"`
	// The collection of training phrases the agent is trained on to identify the intent.
	// Structure is documented below.
	TrainingPhrases []CxIntentTrainingPhrase `pulumi:"trainingPhrases"`
}

// The set of arguments for constructing a CxIntent resource.
type CxIntentArgs struct {
	// Human readable description for better understanding an intent like its scope, content, result etc. Maximum character limit: 140 characters.
	Description pulumi.StringPtrInput
	// The human-readable name of the intent, unique within the agent.
	DisplayName pulumi.StringInput
	// Indicates whether this is a fallback intent. Currently only default fallback intent is allowed in the agent, which is added upon agent creation.
	// Adding training phrases to fallback intent is useful in the case of requests that are mistakenly matched, since training phrases assigned to fallback intents act as negative examples that triggers no-match event.
	IsFallback pulumi.BoolPtrInput
	// The key/value metadata to label an intent. Labels can contain lowercase letters, digits and the symbols '-' and '_'. International characters are allowed, including letters from unicase alphabets. Keys must start with a letter. Keys and values can be no longer than 63 characters and no more than 128 bytes.
	// Prefix "sys-" is reserved for Dialogflow defined labels. Currently allowed Dialogflow defined labels include: * sys-head * sys-contextual The above labels do not require value. "sys-head" means the intent is a head intent. "sys.contextual" means the intent is a contextual intent.
	// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
	Labels pulumi.StringMapInput
	// The language of the following fields in intent:
	// Intent.training_phrases.parts.text
	// If not specified, the agent's default language is used. Many languages are supported. Note: languages must be enabled in the agent before they can be used.
	LanguageCode pulumi.StringPtrInput
	// The collection of parameters associated with the intent.
	// Structure is documented below.
	Parameters CxIntentParameterArrayInput
	// The agent to create an intent for.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	Parent pulumi.StringPtrInput
	// The priority of this intent. Higher numbers represent higher priorities.
	// If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the Normal priority in the console.
	// If the supplied value is negative, the intent is ignored in runtime detect intent requests.
	Priority pulumi.IntPtrInput
	// The collection of training phrases the agent is trained on to identify the intent.
	// Structure is documented below.
	TrainingPhrases CxIntentTrainingPhraseArrayInput
}

func (CxIntentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cxIntentArgs)(nil)).Elem()
}

type CxIntentInput interface {
	pulumi.Input

	ToCxIntentOutput() CxIntentOutput
	ToCxIntentOutputWithContext(ctx context.Context) CxIntentOutput
}

func (*CxIntent) ElementType() reflect.Type {
	return reflect.TypeOf((**CxIntent)(nil)).Elem()
}

func (i *CxIntent) ToCxIntentOutput() CxIntentOutput {
	return i.ToCxIntentOutputWithContext(context.Background())
}

func (i *CxIntent) ToCxIntentOutputWithContext(ctx context.Context) CxIntentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CxIntentOutput)
}

// CxIntentArrayInput is an input type that accepts CxIntentArray and CxIntentArrayOutput values.
// You can construct a concrete instance of `CxIntentArrayInput` via:
//
//          CxIntentArray{ CxIntentArgs{...} }
type CxIntentArrayInput interface {
	pulumi.Input

	ToCxIntentArrayOutput() CxIntentArrayOutput
	ToCxIntentArrayOutputWithContext(context.Context) CxIntentArrayOutput
}

type CxIntentArray []CxIntentInput

func (CxIntentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CxIntent)(nil)).Elem()
}

func (i CxIntentArray) ToCxIntentArrayOutput() CxIntentArrayOutput {
	return i.ToCxIntentArrayOutputWithContext(context.Background())
}

func (i CxIntentArray) ToCxIntentArrayOutputWithContext(ctx context.Context) CxIntentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CxIntentArrayOutput)
}

// CxIntentMapInput is an input type that accepts CxIntentMap and CxIntentMapOutput values.
// You can construct a concrete instance of `CxIntentMapInput` via:
//
//          CxIntentMap{ "key": CxIntentArgs{...} }
type CxIntentMapInput interface {
	pulumi.Input

	ToCxIntentMapOutput() CxIntentMapOutput
	ToCxIntentMapOutputWithContext(context.Context) CxIntentMapOutput
}

type CxIntentMap map[string]CxIntentInput

func (CxIntentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CxIntent)(nil)).Elem()
}

func (i CxIntentMap) ToCxIntentMapOutput() CxIntentMapOutput {
	return i.ToCxIntentMapOutputWithContext(context.Background())
}

func (i CxIntentMap) ToCxIntentMapOutputWithContext(ctx context.Context) CxIntentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CxIntentMapOutput)
}

type CxIntentOutput struct{ *pulumi.OutputState }

func (CxIntentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CxIntent)(nil)).Elem()
}

func (o CxIntentOutput) ToCxIntentOutput() CxIntentOutput {
	return o
}

func (o CxIntentOutput) ToCxIntentOutputWithContext(ctx context.Context) CxIntentOutput {
	return o
}

type CxIntentArrayOutput struct{ *pulumi.OutputState }

func (CxIntentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CxIntent)(nil)).Elem()
}

func (o CxIntentArrayOutput) ToCxIntentArrayOutput() CxIntentArrayOutput {
	return o
}

func (o CxIntentArrayOutput) ToCxIntentArrayOutputWithContext(ctx context.Context) CxIntentArrayOutput {
	return o
}

func (o CxIntentArrayOutput) Index(i pulumi.IntInput) CxIntentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CxIntent {
		return vs[0].([]*CxIntent)[vs[1].(int)]
	}).(CxIntentOutput)
}

type CxIntentMapOutput struct{ *pulumi.OutputState }

func (CxIntentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CxIntent)(nil)).Elem()
}

func (o CxIntentMapOutput) ToCxIntentMapOutput() CxIntentMapOutput {
	return o
}

func (o CxIntentMapOutput) ToCxIntentMapOutputWithContext(ctx context.Context) CxIntentMapOutput {
	return o
}

func (o CxIntentMapOutput) MapIndex(k pulumi.StringInput) CxIntentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CxIntent {
		return vs[0].(map[string]*CxIntent)[vs[1].(string)]
	}).(CxIntentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CxIntentInput)(nil)).Elem(), &CxIntent{})
	pulumi.RegisterInputType(reflect.TypeOf((*CxIntentArrayInput)(nil)).Elem(), CxIntentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CxIntentMapInput)(nil)).Elem(), CxIntentMap{})
	pulumi.RegisterOutputType(CxIntentOutput{})
	pulumi.RegisterOutputType(CxIntentArrayOutput{})
	pulumi.RegisterOutputType(CxIntentMapOutput{})
}
