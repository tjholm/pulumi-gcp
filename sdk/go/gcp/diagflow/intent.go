// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package diagflow

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents a Dialogflow intent. Intents convert a number of user expressions or patterns into an action. An action
// is an extraction of a user command or sentence semantics.
//
// To get more information about Intent, see:
//
// * [API documentation](https://cloud.google.com/dialogflow/docs/reference/rest/v2/projects.agent.intents)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/dialogflow/docs/)
//
// ## Example Usage
// ### Dialogflow Intent Basic
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
// 		basicAgent, err := diagflow.NewAgent(ctx, "basicAgent", &diagflow.AgentArgs{
// 			DisplayName:         pulumi.String("example_agent"),
// 			DefaultLanguageCode: pulumi.String("en"),
// 			TimeZone:            pulumi.String("America/New_York"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = diagflow.NewIntent(ctx, "basicIntent", &diagflow.IntentArgs{
// 			DisplayName: pulumi.String("basic-intent"),
// 		}, pulumi.DependsOn([]pulumi.Resource{
// 			basicAgent,
// 		}))
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
//  $ pulumi import gcp:diagflow/intent:Intent default {{name}}
// ```
type Intent struct {
	pulumi.CustomResourceState

	// The name of the action associated with the intent.
	// Note: The action name must not contain whitespaces.
	Action pulumi.StringOutput `pulumi:"action"`
	// The list of platforms for which the first responses will be copied from the messages in PLATFORM_UNSPECIFIED
	// (i.e. default platform).
	// Each value may be one of `FACEBOOK`, `SLACK`, `TELEGRAM`, `KIK`, `SKYPE`, `LINE`, `VIBER`, `ACTIONS_ON_GOOGLE`, and `GOOGLE_HANGOUTS`.
	DefaultResponsePlatforms pulumi.StringArrayOutput `pulumi:"defaultResponsePlatforms"`
	// The name of this intent to be displayed on the console.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The collection of event names that trigger the intent. If the collection of input contexts is not empty, all of
	// the contexts must be present in the active user session for an event to trigger this intent. See the
	// [events reference](https://cloud.google.com/dialogflow/docs/events-overview) for more details.
	Events pulumi.StringArrayOutput `pulumi:"events"`
	// Information about all followup intents that have this intent as a direct or indirect parent. We populate this field only
	// in the output.
	FollowupIntentInfos IntentFollowupIntentInfoArrayOutput `pulumi:"followupIntentInfos"`
	// The list of context names required for this intent to be triggered.
	// Format: projects/<Project ID>/agent/sessions/-/contexts/<Context ID>.
	InputContextNames pulumi.StringArrayOutput `pulumi:"inputContextNames"`
	// Indicates whether this is a fallback intent.
	IsFallback pulumi.BoolOutput `pulumi:"isFallback"`
	// Indicates whether Machine Learning is disabled for the intent.
	// Note: If mlDisabled setting is set to true, then this intent is not taken into account during inference in ML
	// ONLY match mode. Also, auto-markup in the UI is turned off.
	MlDisabled pulumi.BoolOutput `pulumi:"mlDisabled"`
	// The unique identifier of this intent. Format: projects/<Project ID>/agent/intents/<Intent ID>.
	Name pulumi.StringOutput `pulumi:"name"`
	// The unique identifier of the parent intent in the chain of followup intents.
	// Format: projects/<Project ID>/agent/intents/<Intent ID>.
	ParentFollowupIntentName pulumi.StringOutput `pulumi:"parentFollowupIntentName"`
	// The priority of this intent. Higher numbers represent higher priorities.
	// - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds
	//   to the Normal priority in the console.
	// - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
	Priority pulumi.IntOutput `pulumi:"priority"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Indicates whether to delete all contexts in the current session when this intent is matched.
	ResetContexts pulumi.BoolOutput `pulumi:"resetContexts"`
	// The unique identifier of the root intent in the chain of followup intents. It identifies the correct followup intents
	// chain for this intent. Format: projects/<Project ID>/agent/intents/<Intent ID>.
	RootFollowupIntentName pulumi.StringOutput `pulumi:"rootFollowupIntentName"`
	// Indicates whether webhooks are enabled for the intent.
	// * WEBHOOK_STATE_ENABLED: Webhook is enabled in the agent and in the intent.
	// * WEBHOOK_STATE_ENABLED_FOR_SLOT_FILLING: Webhook is enabled in the agent and in the intent. Also, each slot
	//   filling prompt is forwarded to the webhook.
	//   Possible values are `WEBHOOK_STATE_ENABLED` and `WEBHOOK_STATE_ENABLED_FOR_SLOT_FILLING`.
	WebhookState pulumi.StringOutput `pulumi:"webhookState"`
}

// NewIntent registers a new resource with the given unique name, arguments, and options.
func NewIntent(ctx *pulumi.Context,
	name string, args *IntentArgs, opts ...pulumi.ResourceOption) (*Intent, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	var resource Intent
	err := ctx.RegisterResource("gcp:diagflow/intent:Intent", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntent gets an existing Intent resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntentState, opts ...pulumi.ResourceOption) (*Intent, error) {
	var resource Intent
	err := ctx.ReadResource("gcp:diagflow/intent:Intent", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Intent resources.
type intentState struct {
	// The name of the action associated with the intent.
	// Note: The action name must not contain whitespaces.
	Action *string `pulumi:"action"`
	// The list of platforms for which the first responses will be copied from the messages in PLATFORM_UNSPECIFIED
	// (i.e. default platform).
	// Each value may be one of `FACEBOOK`, `SLACK`, `TELEGRAM`, `KIK`, `SKYPE`, `LINE`, `VIBER`, `ACTIONS_ON_GOOGLE`, and `GOOGLE_HANGOUTS`.
	DefaultResponsePlatforms []string `pulumi:"defaultResponsePlatforms"`
	// The name of this intent to be displayed on the console.
	DisplayName *string `pulumi:"displayName"`
	// The collection of event names that trigger the intent. If the collection of input contexts is not empty, all of
	// the contexts must be present in the active user session for an event to trigger this intent. See the
	// [events reference](https://cloud.google.com/dialogflow/docs/events-overview) for more details.
	Events []string `pulumi:"events"`
	// Information about all followup intents that have this intent as a direct or indirect parent. We populate this field only
	// in the output.
	FollowupIntentInfos []IntentFollowupIntentInfo `pulumi:"followupIntentInfos"`
	// The list of context names required for this intent to be triggered.
	// Format: projects/<Project ID>/agent/sessions/-/contexts/<Context ID>.
	InputContextNames []string `pulumi:"inputContextNames"`
	// Indicates whether this is a fallback intent.
	IsFallback *bool `pulumi:"isFallback"`
	// Indicates whether Machine Learning is disabled for the intent.
	// Note: If mlDisabled setting is set to true, then this intent is not taken into account during inference in ML
	// ONLY match mode. Also, auto-markup in the UI is turned off.
	MlDisabled *bool `pulumi:"mlDisabled"`
	// The unique identifier of this intent. Format: projects/<Project ID>/agent/intents/<Intent ID>.
	Name *string `pulumi:"name"`
	// The unique identifier of the parent intent in the chain of followup intents.
	// Format: projects/<Project ID>/agent/intents/<Intent ID>.
	ParentFollowupIntentName *string `pulumi:"parentFollowupIntentName"`
	// The priority of this intent. Higher numbers represent higher priorities.
	// - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds
	//   to the Normal priority in the console.
	// - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
	Priority *int `pulumi:"priority"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Indicates whether to delete all contexts in the current session when this intent is matched.
	ResetContexts *bool `pulumi:"resetContexts"`
	// The unique identifier of the root intent in the chain of followup intents. It identifies the correct followup intents
	// chain for this intent. Format: projects/<Project ID>/agent/intents/<Intent ID>.
	RootFollowupIntentName *string `pulumi:"rootFollowupIntentName"`
	// Indicates whether webhooks are enabled for the intent.
	// * WEBHOOK_STATE_ENABLED: Webhook is enabled in the agent and in the intent.
	// * WEBHOOK_STATE_ENABLED_FOR_SLOT_FILLING: Webhook is enabled in the agent and in the intent. Also, each slot
	//   filling prompt is forwarded to the webhook.
	//   Possible values are `WEBHOOK_STATE_ENABLED` and `WEBHOOK_STATE_ENABLED_FOR_SLOT_FILLING`.
	WebhookState *string `pulumi:"webhookState"`
}

type IntentState struct {
	// The name of the action associated with the intent.
	// Note: The action name must not contain whitespaces.
	Action pulumi.StringPtrInput
	// The list of platforms for which the first responses will be copied from the messages in PLATFORM_UNSPECIFIED
	// (i.e. default platform).
	// Each value may be one of `FACEBOOK`, `SLACK`, `TELEGRAM`, `KIK`, `SKYPE`, `LINE`, `VIBER`, `ACTIONS_ON_GOOGLE`, and `GOOGLE_HANGOUTS`.
	DefaultResponsePlatforms pulumi.StringArrayInput
	// The name of this intent to be displayed on the console.
	DisplayName pulumi.StringPtrInput
	// The collection of event names that trigger the intent. If the collection of input contexts is not empty, all of
	// the contexts must be present in the active user session for an event to trigger this intent. See the
	// [events reference](https://cloud.google.com/dialogflow/docs/events-overview) for more details.
	Events pulumi.StringArrayInput
	// Information about all followup intents that have this intent as a direct or indirect parent. We populate this field only
	// in the output.
	FollowupIntentInfos IntentFollowupIntentInfoArrayInput
	// The list of context names required for this intent to be triggered.
	// Format: projects/<Project ID>/agent/sessions/-/contexts/<Context ID>.
	InputContextNames pulumi.StringArrayInput
	// Indicates whether this is a fallback intent.
	IsFallback pulumi.BoolPtrInput
	// Indicates whether Machine Learning is disabled for the intent.
	// Note: If mlDisabled setting is set to true, then this intent is not taken into account during inference in ML
	// ONLY match mode. Also, auto-markup in the UI is turned off.
	MlDisabled pulumi.BoolPtrInput
	// The unique identifier of this intent. Format: projects/<Project ID>/agent/intents/<Intent ID>.
	Name pulumi.StringPtrInput
	// The unique identifier of the parent intent in the chain of followup intents.
	// Format: projects/<Project ID>/agent/intents/<Intent ID>.
	ParentFollowupIntentName pulumi.StringPtrInput
	// The priority of this intent. Higher numbers represent higher priorities.
	// - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds
	//   to the Normal priority in the console.
	// - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
	Priority pulumi.IntPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Indicates whether to delete all contexts in the current session when this intent is matched.
	ResetContexts pulumi.BoolPtrInput
	// The unique identifier of the root intent in the chain of followup intents. It identifies the correct followup intents
	// chain for this intent. Format: projects/<Project ID>/agent/intents/<Intent ID>.
	RootFollowupIntentName pulumi.StringPtrInput
	// Indicates whether webhooks are enabled for the intent.
	// * WEBHOOK_STATE_ENABLED: Webhook is enabled in the agent and in the intent.
	// * WEBHOOK_STATE_ENABLED_FOR_SLOT_FILLING: Webhook is enabled in the agent and in the intent. Also, each slot
	//   filling prompt is forwarded to the webhook.
	//   Possible values are `WEBHOOK_STATE_ENABLED` and `WEBHOOK_STATE_ENABLED_FOR_SLOT_FILLING`.
	WebhookState pulumi.StringPtrInput
}

func (IntentState) ElementType() reflect.Type {
	return reflect.TypeOf((*intentState)(nil)).Elem()
}

type intentArgs struct {
	// The name of the action associated with the intent.
	// Note: The action name must not contain whitespaces.
	Action *string `pulumi:"action"`
	// The list of platforms for which the first responses will be copied from the messages in PLATFORM_UNSPECIFIED
	// (i.e. default platform).
	// Each value may be one of `FACEBOOK`, `SLACK`, `TELEGRAM`, `KIK`, `SKYPE`, `LINE`, `VIBER`, `ACTIONS_ON_GOOGLE`, and `GOOGLE_HANGOUTS`.
	DefaultResponsePlatforms []string `pulumi:"defaultResponsePlatforms"`
	// The name of this intent to be displayed on the console.
	DisplayName string `pulumi:"displayName"`
	// The collection of event names that trigger the intent. If the collection of input contexts is not empty, all of
	// the contexts must be present in the active user session for an event to trigger this intent. See the
	// [events reference](https://cloud.google.com/dialogflow/docs/events-overview) for more details.
	Events []string `pulumi:"events"`
	// The list of context names required for this intent to be triggered.
	// Format: projects/<Project ID>/agent/sessions/-/contexts/<Context ID>.
	InputContextNames []string `pulumi:"inputContextNames"`
	// Indicates whether this is a fallback intent.
	IsFallback *bool `pulumi:"isFallback"`
	// Indicates whether Machine Learning is disabled for the intent.
	// Note: If mlDisabled setting is set to true, then this intent is not taken into account during inference in ML
	// ONLY match mode. Also, auto-markup in the UI is turned off.
	MlDisabled *bool `pulumi:"mlDisabled"`
	// The unique identifier of the parent intent in the chain of followup intents.
	// Format: projects/<Project ID>/agent/intents/<Intent ID>.
	ParentFollowupIntentName *string `pulumi:"parentFollowupIntentName"`
	// The priority of this intent. Higher numbers represent higher priorities.
	// - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds
	//   to the Normal priority in the console.
	// - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
	Priority *int `pulumi:"priority"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Indicates whether to delete all contexts in the current session when this intent is matched.
	ResetContexts *bool `pulumi:"resetContexts"`
	// Indicates whether webhooks are enabled for the intent.
	// * WEBHOOK_STATE_ENABLED: Webhook is enabled in the agent and in the intent.
	// * WEBHOOK_STATE_ENABLED_FOR_SLOT_FILLING: Webhook is enabled in the agent and in the intent. Also, each slot
	//   filling prompt is forwarded to the webhook.
	//   Possible values are `WEBHOOK_STATE_ENABLED` and `WEBHOOK_STATE_ENABLED_FOR_SLOT_FILLING`.
	WebhookState *string `pulumi:"webhookState"`
}

// The set of arguments for constructing a Intent resource.
type IntentArgs struct {
	// The name of the action associated with the intent.
	// Note: The action name must not contain whitespaces.
	Action pulumi.StringPtrInput
	// The list of platforms for which the first responses will be copied from the messages in PLATFORM_UNSPECIFIED
	// (i.e. default platform).
	// Each value may be one of `FACEBOOK`, `SLACK`, `TELEGRAM`, `KIK`, `SKYPE`, `LINE`, `VIBER`, `ACTIONS_ON_GOOGLE`, and `GOOGLE_HANGOUTS`.
	DefaultResponsePlatforms pulumi.StringArrayInput
	// The name of this intent to be displayed on the console.
	DisplayName pulumi.StringInput
	// The collection of event names that trigger the intent. If the collection of input contexts is not empty, all of
	// the contexts must be present in the active user session for an event to trigger this intent. See the
	// [events reference](https://cloud.google.com/dialogflow/docs/events-overview) for more details.
	Events pulumi.StringArrayInput
	// The list of context names required for this intent to be triggered.
	// Format: projects/<Project ID>/agent/sessions/-/contexts/<Context ID>.
	InputContextNames pulumi.StringArrayInput
	// Indicates whether this is a fallback intent.
	IsFallback pulumi.BoolPtrInput
	// Indicates whether Machine Learning is disabled for the intent.
	// Note: If mlDisabled setting is set to true, then this intent is not taken into account during inference in ML
	// ONLY match mode. Also, auto-markup in the UI is turned off.
	MlDisabled pulumi.BoolPtrInput
	// The unique identifier of the parent intent in the chain of followup intents.
	// Format: projects/<Project ID>/agent/intents/<Intent ID>.
	ParentFollowupIntentName pulumi.StringPtrInput
	// The priority of this intent. Higher numbers represent higher priorities.
	// - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds
	//   to the Normal priority in the console.
	// - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
	Priority pulumi.IntPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Indicates whether to delete all contexts in the current session when this intent is matched.
	ResetContexts pulumi.BoolPtrInput
	// Indicates whether webhooks are enabled for the intent.
	// * WEBHOOK_STATE_ENABLED: Webhook is enabled in the agent and in the intent.
	// * WEBHOOK_STATE_ENABLED_FOR_SLOT_FILLING: Webhook is enabled in the agent and in the intent. Also, each slot
	//   filling prompt is forwarded to the webhook.
	//   Possible values are `WEBHOOK_STATE_ENABLED` and `WEBHOOK_STATE_ENABLED_FOR_SLOT_FILLING`.
	WebhookState pulumi.StringPtrInput
}

func (IntentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*intentArgs)(nil)).Elem()
}

type IntentInput interface {
	pulumi.Input

	ToIntentOutput() IntentOutput
	ToIntentOutputWithContext(ctx context.Context) IntentOutput
}

func (*Intent) ElementType() reflect.Type {
	return reflect.TypeOf((**Intent)(nil)).Elem()
}

func (i *Intent) ToIntentOutput() IntentOutput {
	return i.ToIntentOutputWithContext(context.Background())
}

func (i *Intent) ToIntentOutputWithContext(ctx context.Context) IntentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntentOutput)
}

// IntentArrayInput is an input type that accepts IntentArray and IntentArrayOutput values.
// You can construct a concrete instance of `IntentArrayInput` via:
//
//          IntentArray{ IntentArgs{...} }
type IntentArrayInput interface {
	pulumi.Input

	ToIntentArrayOutput() IntentArrayOutput
	ToIntentArrayOutputWithContext(context.Context) IntentArrayOutput
}

type IntentArray []IntentInput

func (IntentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Intent)(nil)).Elem()
}

func (i IntentArray) ToIntentArrayOutput() IntentArrayOutput {
	return i.ToIntentArrayOutputWithContext(context.Background())
}

func (i IntentArray) ToIntentArrayOutputWithContext(ctx context.Context) IntentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntentArrayOutput)
}

// IntentMapInput is an input type that accepts IntentMap and IntentMapOutput values.
// You can construct a concrete instance of `IntentMapInput` via:
//
//          IntentMap{ "key": IntentArgs{...} }
type IntentMapInput interface {
	pulumi.Input

	ToIntentMapOutput() IntentMapOutput
	ToIntentMapOutputWithContext(context.Context) IntentMapOutput
}

type IntentMap map[string]IntentInput

func (IntentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Intent)(nil)).Elem()
}

func (i IntentMap) ToIntentMapOutput() IntentMapOutput {
	return i.ToIntentMapOutputWithContext(context.Background())
}

func (i IntentMap) ToIntentMapOutputWithContext(ctx context.Context) IntentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntentMapOutput)
}

type IntentOutput struct{ *pulumi.OutputState }

func (IntentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Intent)(nil)).Elem()
}

func (o IntentOutput) ToIntentOutput() IntentOutput {
	return o
}

func (o IntentOutput) ToIntentOutputWithContext(ctx context.Context) IntentOutput {
	return o
}

type IntentArrayOutput struct{ *pulumi.OutputState }

func (IntentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Intent)(nil)).Elem()
}

func (o IntentArrayOutput) ToIntentArrayOutput() IntentArrayOutput {
	return o
}

func (o IntentArrayOutput) ToIntentArrayOutputWithContext(ctx context.Context) IntentArrayOutput {
	return o
}

func (o IntentArrayOutput) Index(i pulumi.IntInput) IntentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Intent {
		return vs[0].([]*Intent)[vs[1].(int)]
	}).(IntentOutput)
}

type IntentMapOutput struct{ *pulumi.OutputState }

func (IntentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Intent)(nil)).Elem()
}

func (o IntentMapOutput) ToIntentMapOutput() IntentMapOutput {
	return o
}

func (o IntentMapOutput) ToIntentMapOutputWithContext(ctx context.Context) IntentMapOutput {
	return o
}

func (o IntentMapOutput) MapIndex(k pulumi.StringInput) IntentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Intent {
		return vs[0].(map[string]*Intent)[vs[1].(string)]
	}).(IntentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IntentInput)(nil)).Elem(), &Intent{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntentArrayInput)(nil)).Elem(), IntentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IntentMapInput)(nil)).Elem(), IntentMap{})
	pulumi.RegisterOutputType(IntentOutput{})
	pulumi.RegisterOutputType(IntentArrayOutput{})
	pulumi.RegisterOutputType(IntentMapOutput{})
}
