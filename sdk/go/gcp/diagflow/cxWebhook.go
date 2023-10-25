// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package diagflow

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Webhooks host the developer's business logic. During a session, webhooks allow the developer to use the data extracted by Dialogflow's natural language processing to generate dynamic responses, validate collected data, or trigger actions on the backend.
//
// To get more information about Webhook, see:
//
// * [API documentation](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.webhooks)
// * How-to Guides
//   - [Official Documentation](https://cloud.google.com/dialogflow/cx/docs)
//
// ## Example Usage
// ### Dialogflowcx Webhook Full
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
//					pulumi.String("it"),
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
//			_, err = diagflow.NewCxWebhook(ctx, "basicWebhook", &diagflow.CxWebhookArgs{
//				Parent:      agent.ID(),
//				DisplayName: pulumi.String("MyFlow"),
//				GenericWebService: &diagflow.CxWebhookGenericWebServiceArgs{
//					Uri: pulumi.String("https://example.com"),
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
// # Webhook can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:diagflow/cxWebhook:CxWebhook default {{parent}}/webhooks/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:diagflow/cxWebhook:CxWebhook default {{parent}}/{{name}}
//
// ```
type CxWebhook struct {
	pulumi.CustomResourceState

	// Indicates whether the webhook is disabled.
	Disabled pulumi.BoolPtrOutput `pulumi:"disabled"`
	// The human-readable name of the webhook, unique within the agent.
	//
	// ***
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Indicates if automatic spell correction is enabled in detect intent requests.
	EnableSpellCorrection pulumi.BoolPtrOutput `pulumi:"enableSpellCorrection"`
	// Determines whether this agent should log conversation queries.
	EnableStackdriverLogging pulumi.BoolPtrOutput `pulumi:"enableStackdriverLogging"`
	// Configuration for a generic web service.
	// Structure is documented below.
	GenericWebService CxWebhookGenericWebServicePtrOutput `pulumi:"genericWebService"`
	// The unique identifier of the webhook.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/webhooks/<Webhook ID>.
	Name pulumi.StringOutput `pulumi:"name"`
	// The agent to create a webhook for.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	Parent pulumi.StringPtrOutput `pulumi:"parent"`
	// Name of the SecuritySettings reference for the agent. Format: projects/<Project ID>/locations/<Location ID>/securitySettings/<Security Settings ID>.
	SecuritySettings pulumi.StringPtrOutput `pulumi:"securitySettings"`
	// Configuration for a Service Directory service.
	// Structure is documented below.
	ServiceDirectory CxWebhookServiceDirectoryPtrOutput `pulumi:"serviceDirectory"`
	// Name of the start flow in this agent. A start flow will be automatically created when the agent is created, and can only be deleted by deleting the agent. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.
	StartFlow pulumi.StringOutput `pulumi:"startFlow"`
	// Webhook execution timeout.
	Timeout pulumi.StringPtrOutput `pulumi:"timeout"`
}

// NewCxWebhook registers a new resource with the given unique name, arguments, and options.
func NewCxWebhook(ctx *pulumi.Context,
	name string, args *CxWebhookArgs, opts ...pulumi.ResourceOption) (*CxWebhook, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource CxWebhook
	err := ctx.RegisterResource("gcp:diagflow/cxWebhook:CxWebhook", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCxWebhook gets an existing CxWebhook resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCxWebhook(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CxWebhookState, opts ...pulumi.ResourceOption) (*CxWebhook, error) {
	var resource CxWebhook
	err := ctx.ReadResource("gcp:diagflow/cxWebhook:CxWebhook", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CxWebhook resources.
type cxWebhookState struct {
	// Indicates whether the webhook is disabled.
	Disabled *bool `pulumi:"disabled"`
	// The human-readable name of the webhook, unique within the agent.
	//
	// ***
	DisplayName *string `pulumi:"displayName"`
	// Indicates if automatic spell correction is enabled in detect intent requests.
	EnableSpellCorrection *bool `pulumi:"enableSpellCorrection"`
	// Determines whether this agent should log conversation queries.
	EnableStackdriverLogging *bool `pulumi:"enableStackdriverLogging"`
	// Configuration for a generic web service.
	// Structure is documented below.
	GenericWebService *CxWebhookGenericWebService `pulumi:"genericWebService"`
	// The unique identifier of the webhook.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/webhooks/<Webhook ID>.
	Name *string `pulumi:"name"`
	// The agent to create a webhook for.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	Parent *string `pulumi:"parent"`
	// Name of the SecuritySettings reference for the agent. Format: projects/<Project ID>/locations/<Location ID>/securitySettings/<Security Settings ID>.
	SecuritySettings *string `pulumi:"securitySettings"`
	// Configuration for a Service Directory service.
	// Structure is documented below.
	ServiceDirectory *CxWebhookServiceDirectory `pulumi:"serviceDirectory"`
	// Name of the start flow in this agent. A start flow will be automatically created when the agent is created, and can only be deleted by deleting the agent. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.
	StartFlow *string `pulumi:"startFlow"`
	// Webhook execution timeout.
	Timeout *string `pulumi:"timeout"`
}

type CxWebhookState struct {
	// Indicates whether the webhook is disabled.
	Disabled pulumi.BoolPtrInput
	// The human-readable name of the webhook, unique within the agent.
	//
	// ***
	DisplayName pulumi.StringPtrInput
	// Indicates if automatic spell correction is enabled in detect intent requests.
	EnableSpellCorrection pulumi.BoolPtrInput
	// Determines whether this agent should log conversation queries.
	EnableStackdriverLogging pulumi.BoolPtrInput
	// Configuration for a generic web service.
	// Structure is documented below.
	GenericWebService CxWebhookGenericWebServicePtrInput
	// The unique identifier of the webhook.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/webhooks/<Webhook ID>.
	Name pulumi.StringPtrInput
	// The agent to create a webhook for.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	Parent pulumi.StringPtrInput
	// Name of the SecuritySettings reference for the agent. Format: projects/<Project ID>/locations/<Location ID>/securitySettings/<Security Settings ID>.
	SecuritySettings pulumi.StringPtrInput
	// Configuration for a Service Directory service.
	// Structure is documented below.
	ServiceDirectory CxWebhookServiceDirectoryPtrInput
	// Name of the start flow in this agent. A start flow will be automatically created when the agent is created, and can only be deleted by deleting the agent. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.
	StartFlow pulumi.StringPtrInput
	// Webhook execution timeout.
	Timeout pulumi.StringPtrInput
}

func (CxWebhookState) ElementType() reflect.Type {
	return reflect.TypeOf((*cxWebhookState)(nil)).Elem()
}

type cxWebhookArgs struct {
	// Indicates whether the webhook is disabled.
	Disabled *bool `pulumi:"disabled"`
	// The human-readable name of the webhook, unique within the agent.
	//
	// ***
	DisplayName string `pulumi:"displayName"`
	// Indicates if automatic spell correction is enabled in detect intent requests.
	EnableSpellCorrection *bool `pulumi:"enableSpellCorrection"`
	// Determines whether this agent should log conversation queries.
	EnableStackdriverLogging *bool `pulumi:"enableStackdriverLogging"`
	// Configuration for a generic web service.
	// Structure is documented below.
	GenericWebService *CxWebhookGenericWebService `pulumi:"genericWebService"`
	// The agent to create a webhook for.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	Parent *string `pulumi:"parent"`
	// Name of the SecuritySettings reference for the agent. Format: projects/<Project ID>/locations/<Location ID>/securitySettings/<Security Settings ID>.
	SecuritySettings *string `pulumi:"securitySettings"`
	// Configuration for a Service Directory service.
	// Structure is documented below.
	ServiceDirectory *CxWebhookServiceDirectory `pulumi:"serviceDirectory"`
	// Webhook execution timeout.
	Timeout *string `pulumi:"timeout"`
}

// The set of arguments for constructing a CxWebhook resource.
type CxWebhookArgs struct {
	// Indicates whether the webhook is disabled.
	Disabled pulumi.BoolPtrInput
	// The human-readable name of the webhook, unique within the agent.
	//
	// ***
	DisplayName pulumi.StringInput
	// Indicates if automatic spell correction is enabled in detect intent requests.
	EnableSpellCorrection pulumi.BoolPtrInput
	// Determines whether this agent should log conversation queries.
	EnableStackdriverLogging pulumi.BoolPtrInput
	// Configuration for a generic web service.
	// Structure is documented below.
	GenericWebService CxWebhookGenericWebServicePtrInput
	// The agent to create a webhook for.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	Parent pulumi.StringPtrInput
	// Name of the SecuritySettings reference for the agent. Format: projects/<Project ID>/locations/<Location ID>/securitySettings/<Security Settings ID>.
	SecuritySettings pulumi.StringPtrInput
	// Configuration for a Service Directory service.
	// Structure is documented below.
	ServiceDirectory CxWebhookServiceDirectoryPtrInput
	// Webhook execution timeout.
	Timeout pulumi.StringPtrInput
}

func (CxWebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cxWebhookArgs)(nil)).Elem()
}

type CxWebhookInput interface {
	pulumi.Input

	ToCxWebhookOutput() CxWebhookOutput
	ToCxWebhookOutputWithContext(ctx context.Context) CxWebhookOutput
}

func (*CxWebhook) ElementType() reflect.Type {
	return reflect.TypeOf((**CxWebhook)(nil)).Elem()
}

func (i *CxWebhook) ToCxWebhookOutput() CxWebhookOutput {
	return i.ToCxWebhookOutputWithContext(context.Background())
}

func (i *CxWebhook) ToCxWebhookOutputWithContext(ctx context.Context) CxWebhookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CxWebhookOutput)
}

func (i *CxWebhook) ToOutput(ctx context.Context) pulumix.Output[*CxWebhook] {
	return pulumix.Output[*CxWebhook]{
		OutputState: i.ToCxWebhookOutputWithContext(ctx).OutputState,
	}
}

// CxWebhookArrayInput is an input type that accepts CxWebhookArray and CxWebhookArrayOutput values.
// You can construct a concrete instance of `CxWebhookArrayInput` via:
//
//	CxWebhookArray{ CxWebhookArgs{...} }
type CxWebhookArrayInput interface {
	pulumi.Input

	ToCxWebhookArrayOutput() CxWebhookArrayOutput
	ToCxWebhookArrayOutputWithContext(context.Context) CxWebhookArrayOutput
}

type CxWebhookArray []CxWebhookInput

func (CxWebhookArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CxWebhook)(nil)).Elem()
}

func (i CxWebhookArray) ToCxWebhookArrayOutput() CxWebhookArrayOutput {
	return i.ToCxWebhookArrayOutputWithContext(context.Background())
}

func (i CxWebhookArray) ToCxWebhookArrayOutputWithContext(ctx context.Context) CxWebhookArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CxWebhookArrayOutput)
}

func (i CxWebhookArray) ToOutput(ctx context.Context) pulumix.Output[[]*CxWebhook] {
	return pulumix.Output[[]*CxWebhook]{
		OutputState: i.ToCxWebhookArrayOutputWithContext(ctx).OutputState,
	}
}

// CxWebhookMapInput is an input type that accepts CxWebhookMap and CxWebhookMapOutput values.
// You can construct a concrete instance of `CxWebhookMapInput` via:
//
//	CxWebhookMap{ "key": CxWebhookArgs{...} }
type CxWebhookMapInput interface {
	pulumi.Input

	ToCxWebhookMapOutput() CxWebhookMapOutput
	ToCxWebhookMapOutputWithContext(context.Context) CxWebhookMapOutput
}

type CxWebhookMap map[string]CxWebhookInput

func (CxWebhookMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CxWebhook)(nil)).Elem()
}

func (i CxWebhookMap) ToCxWebhookMapOutput() CxWebhookMapOutput {
	return i.ToCxWebhookMapOutputWithContext(context.Background())
}

func (i CxWebhookMap) ToCxWebhookMapOutputWithContext(ctx context.Context) CxWebhookMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CxWebhookMapOutput)
}

func (i CxWebhookMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*CxWebhook] {
	return pulumix.Output[map[string]*CxWebhook]{
		OutputState: i.ToCxWebhookMapOutputWithContext(ctx).OutputState,
	}
}

type CxWebhookOutput struct{ *pulumi.OutputState }

func (CxWebhookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CxWebhook)(nil)).Elem()
}

func (o CxWebhookOutput) ToCxWebhookOutput() CxWebhookOutput {
	return o
}

func (o CxWebhookOutput) ToCxWebhookOutputWithContext(ctx context.Context) CxWebhookOutput {
	return o
}

func (o CxWebhookOutput) ToOutput(ctx context.Context) pulumix.Output[*CxWebhook] {
	return pulumix.Output[*CxWebhook]{
		OutputState: o.OutputState,
	}
}

// Indicates whether the webhook is disabled.
func (o CxWebhookOutput) Disabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CxWebhook) pulumi.BoolPtrOutput { return v.Disabled }).(pulumi.BoolPtrOutput)
}

// The human-readable name of the webhook, unique within the agent.
//
// ***
func (o CxWebhookOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *CxWebhook) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Indicates if automatic spell correction is enabled in detect intent requests.
func (o CxWebhookOutput) EnableSpellCorrection() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CxWebhook) pulumi.BoolPtrOutput { return v.EnableSpellCorrection }).(pulumi.BoolPtrOutput)
}

// Determines whether this agent should log conversation queries.
func (o CxWebhookOutput) EnableStackdriverLogging() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CxWebhook) pulumi.BoolPtrOutput { return v.EnableStackdriverLogging }).(pulumi.BoolPtrOutput)
}

// Configuration for a generic web service.
// Structure is documented below.
func (o CxWebhookOutput) GenericWebService() CxWebhookGenericWebServicePtrOutput {
	return o.ApplyT(func(v *CxWebhook) CxWebhookGenericWebServicePtrOutput { return v.GenericWebService }).(CxWebhookGenericWebServicePtrOutput)
}

// The unique identifier of the webhook.
// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/webhooks/<Webhook ID>.
func (o CxWebhookOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CxWebhook) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The agent to create a webhook for.
// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
func (o CxWebhookOutput) Parent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CxWebhook) pulumi.StringPtrOutput { return v.Parent }).(pulumi.StringPtrOutput)
}

// Name of the SecuritySettings reference for the agent. Format: projects/<Project ID>/locations/<Location ID>/securitySettings/<Security Settings ID>.
func (o CxWebhookOutput) SecuritySettings() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CxWebhook) pulumi.StringPtrOutput { return v.SecuritySettings }).(pulumi.StringPtrOutput)
}

// Configuration for a Service Directory service.
// Structure is documented below.
func (o CxWebhookOutput) ServiceDirectory() CxWebhookServiceDirectoryPtrOutput {
	return o.ApplyT(func(v *CxWebhook) CxWebhookServiceDirectoryPtrOutput { return v.ServiceDirectory }).(CxWebhookServiceDirectoryPtrOutput)
}

// Name of the start flow in this agent. A start flow will be automatically created when the agent is created, and can only be deleted by deleting the agent. Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.
func (o CxWebhookOutput) StartFlow() pulumi.StringOutput {
	return o.ApplyT(func(v *CxWebhook) pulumi.StringOutput { return v.StartFlow }).(pulumi.StringOutput)
}

// Webhook execution timeout.
func (o CxWebhookOutput) Timeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CxWebhook) pulumi.StringPtrOutput { return v.Timeout }).(pulumi.StringPtrOutput)
}

type CxWebhookArrayOutput struct{ *pulumi.OutputState }

func (CxWebhookArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CxWebhook)(nil)).Elem()
}

func (o CxWebhookArrayOutput) ToCxWebhookArrayOutput() CxWebhookArrayOutput {
	return o
}

func (o CxWebhookArrayOutput) ToCxWebhookArrayOutputWithContext(ctx context.Context) CxWebhookArrayOutput {
	return o
}

func (o CxWebhookArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*CxWebhook] {
	return pulumix.Output[[]*CxWebhook]{
		OutputState: o.OutputState,
	}
}

func (o CxWebhookArrayOutput) Index(i pulumi.IntInput) CxWebhookOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CxWebhook {
		return vs[0].([]*CxWebhook)[vs[1].(int)]
	}).(CxWebhookOutput)
}

type CxWebhookMapOutput struct{ *pulumi.OutputState }

func (CxWebhookMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CxWebhook)(nil)).Elem()
}

func (o CxWebhookMapOutput) ToCxWebhookMapOutput() CxWebhookMapOutput {
	return o
}

func (o CxWebhookMapOutput) ToCxWebhookMapOutputWithContext(ctx context.Context) CxWebhookMapOutput {
	return o
}

func (o CxWebhookMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*CxWebhook] {
	return pulumix.Output[map[string]*CxWebhook]{
		OutputState: o.OutputState,
	}
}

func (o CxWebhookMapOutput) MapIndex(k pulumi.StringInput) CxWebhookOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CxWebhook {
		return vs[0].(map[string]*CxWebhook)[vs[1].(string)]
	}).(CxWebhookOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CxWebhookInput)(nil)).Elem(), &CxWebhook{})
	pulumi.RegisterInputType(reflect.TypeOf((*CxWebhookArrayInput)(nil)).Elem(), CxWebhookArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CxWebhookMapInput)(nil)).Elem(), CxWebhookMap{})
	pulumi.RegisterOutputType(CxWebhookOutput{})
	pulumi.RegisterOutputType(CxWebhookArrayOutput{})
	pulumi.RegisterOutputType(CxWebhookMapOutput{})
}
