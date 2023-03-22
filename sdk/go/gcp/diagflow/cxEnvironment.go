// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package diagflow

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Represents an environment for an agent. You can create multiple versions of your agent and publish them to separate environments.
// When you edit an agent, you are editing the draft agent. At any point, you can save the draft agent as an agent version, which is an immutable snapshot of your agent.
// When you save the draft agent, it is published to the default environment. When you create agent versions, you can publish them to custom environments. You can create a variety of custom environments for testing, development, production, etc.
//
// To get more information about Environment, see:
//
// * [API documentation](https://cloud.google.com/dialogflow/cx/docs/reference/rest/v3/projects.locations.agents.environments)
// * How-to Guides
//   - [Official Documentation](https://cloud.google.com/dialogflow/cx/docs)
//
// ## Example Usage
// ### Dialogflowcx Environment Full
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
//			version1, err := diagflow.NewCxVersion(ctx, "version1", &diagflow.CxVersionArgs{
//				Parent:      agent.StartFlow,
//				DisplayName: pulumi.String("1.0.0"),
//				Description: pulumi.String("version 1.0.0"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = diagflow.NewCxEnvironment(ctx, "development", &diagflow.CxEnvironmentArgs{
//				Parent:      agent.ID(),
//				DisplayName: pulumi.String("Development"),
//				Description: pulumi.String("Development Environment"),
//				VersionConfigs: diagflow.CxEnvironmentVersionConfigArray{
//					&diagflow.CxEnvironmentVersionConfigArgs{
//						Version: version1.ID(),
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
// # Environment can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:diagflow/cxEnvironment:CxEnvironment default {{parent}}/environments/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:diagflow/cxEnvironment:CxEnvironment default {{parent}}/{{name}}
//
// ```
type CxEnvironment struct {
	pulumi.CustomResourceState

	// The human-readable description of the environment. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The human-readable name of the environment (unique in an agent). Limit of 64 characters.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The name of the environment.
	Name pulumi.StringOutput `pulumi:"name"`
	// The Agent to create an Environment for.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	Parent pulumi.StringPtrOutput `pulumi:"parent"`
	// Update time of this environment. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// A list of configurations for flow versions. You should include version configs for all flows that are reachable from [Start Flow][Agent.start_flow] in the agent. Otherwise, an error will be returned.
	// Structure is documented below.
	VersionConfigs CxEnvironmentVersionConfigArrayOutput `pulumi:"versionConfigs"`
}

// NewCxEnvironment registers a new resource with the given unique name, arguments, and options.
func NewCxEnvironment(ctx *pulumi.Context,
	name string, args *CxEnvironmentArgs, opts ...pulumi.ResourceOption) (*CxEnvironment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.VersionConfigs == nil {
		return nil, errors.New("invalid value for required argument 'VersionConfigs'")
	}
	var resource CxEnvironment
	err := ctx.RegisterResource("gcp:diagflow/cxEnvironment:CxEnvironment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCxEnvironment gets an existing CxEnvironment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCxEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CxEnvironmentState, opts ...pulumi.ResourceOption) (*CxEnvironment, error) {
	var resource CxEnvironment
	err := ctx.ReadResource("gcp:diagflow/cxEnvironment:CxEnvironment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CxEnvironment resources.
type cxEnvironmentState struct {
	// The human-readable description of the environment. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description *string `pulumi:"description"`
	// The human-readable name of the environment (unique in an agent). Limit of 64 characters.
	DisplayName *string `pulumi:"displayName"`
	// The name of the environment.
	Name *string `pulumi:"name"`
	// The Agent to create an Environment for.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	Parent *string `pulumi:"parent"`
	// Update time of this environment. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime *string `pulumi:"updateTime"`
	// A list of configurations for flow versions. You should include version configs for all flows that are reachable from [Start Flow][Agent.start_flow] in the agent. Otherwise, an error will be returned.
	// Structure is documented below.
	VersionConfigs []CxEnvironmentVersionConfig `pulumi:"versionConfigs"`
}

type CxEnvironmentState struct {
	// The human-readable description of the environment. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description pulumi.StringPtrInput
	// The human-readable name of the environment (unique in an agent). Limit of 64 characters.
	DisplayName pulumi.StringPtrInput
	// The name of the environment.
	Name pulumi.StringPtrInput
	// The Agent to create an Environment for.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	Parent pulumi.StringPtrInput
	// Update time of this environment. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
	UpdateTime pulumi.StringPtrInput
	// A list of configurations for flow versions. You should include version configs for all flows that are reachable from [Start Flow][Agent.start_flow] in the agent. Otherwise, an error will be returned.
	// Structure is documented below.
	VersionConfigs CxEnvironmentVersionConfigArrayInput
}

func (CxEnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*cxEnvironmentState)(nil)).Elem()
}

type cxEnvironmentArgs struct {
	// The human-readable description of the environment. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description *string `pulumi:"description"`
	// The human-readable name of the environment (unique in an agent). Limit of 64 characters.
	DisplayName string `pulumi:"displayName"`
	// The Agent to create an Environment for.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	Parent *string `pulumi:"parent"`
	// A list of configurations for flow versions. You should include version configs for all flows that are reachable from [Start Flow][Agent.start_flow] in the agent. Otherwise, an error will be returned.
	// Structure is documented below.
	VersionConfigs []CxEnvironmentVersionConfig `pulumi:"versionConfigs"`
}

// The set of arguments for constructing a CxEnvironment resource.
type CxEnvironmentArgs struct {
	// The human-readable description of the environment. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description pulumi.StringPtrInput
	// The human-readable name of the environment (unique in an agent). Limit of 64 characters.
	DisplayName pulumi.StringInput
	// The Agent to create an Environment for.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
	Parent pulumi.StringPtrInput
	// A list of configurations for flow versions. You should include version configs for all flows that are reachable from [Start Flow][Agent.start_flow] in the agent. Otherwise, an error will be returned.
	// Structure is documented below.
	VersionConfigs CxEnvironmentVersionConfigArrayInput
}

func (CxEnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*cxEnvironmentArgs)(nil)).Elem()
}

type CxEnvironmentInput interface {
	pulumi.Input

	ToCxEnvironmentOutput() CxEnvironmentOutput
	ToCxEnvironmentOutputWithContext(ctx context.Context) CxEnvironmentOutput
}

func (*CxEnvironment) ElementType() reflect.Type {
	return reflect.TypeOf((**CxEnvironment)(nil)).Elem()
}

func (i *CxEnvironment) ToCxEnvironmentOutput() CxEnvironmentOutput {
	return i.ToCxEnvironmentOutputWithContext(context.Background())
}

func (i *CxEnvironment) ToCxEnvironmentOutputWithContext(ctx context.Context) CxEnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CxEnvironmentOutput)
}

// CxEnvironmentArrayInput is an input type that accepts CxEnvironmentArray and CxEnvironmentArrayOutput values.
// You can construct a concrete instance of `CxEnvironmentArrayInput` via:
//
//	CxEnvironmentArray{ CxEnvironmentArgs{...} }
type CxEnvironmentArrayInput interface {
	pulumi.Input

	ToCxEnvironmentArrayOutput() CxEnvironmentArrayOutput
	ToCxEnvironmentArrayOutputWithContext(context.Context) CxEnvironmentArrayOutput
}

type CxEnvironmentArray []CxEnvironmentInput

func (CxEnvironmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CxEnvironment)(nil)).Elem()
}

func (i CxEnvironmentArray) ToCxEnvironmentArrayOutput() CxEnvironmentArrayOutput {
	return i.ToCxEnvironmentArrayOutputWithContext(context.Background())
}

func (i CxEnvironmentArray) ToCxEnvironmentArrayOutputWithContext(ctx context.Context) CxEnvironmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CxEnvironmentArrayOutput)
}

// CxEnvironmentMapInput is an input type that accepts CxEnvironmentMap and CxEnvironmentMapOutput values.
// You can construct a concrete instance of `CxEnvironmentMapInput` via:
//
//	CxEnvironmentMap{ "key": CxEnvironmentArgs{...} }
type CxEnvironmentMapInput interface {
	pulumi.Input

	ToCxEnvironmentMapOutput() CxEnvironmentMapOutput
	ToCxEnvironmentMapOutputWithContext(context.Context) CxEnvironmentMapOutput
}

type CxEnvironmentMap map[string]CxEnvironmentInput

func (CxEnvironmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CxEnvironment)(nil)).Elem()
}

func (i CxEnvironmentMap) ToCxEnvironmentMapOutput() CxEnvironmentMapOutput {
	return i.ToCxEnvironmentMapOutputWithContext(context.Background())
}

func (i CxEnvironmentMap) ToCxEnvironmentMapOutputWithContext(ctx context.Context) CxEnvironmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CxEnvironmentMapOutput)
}

type CxEnvironmentOutput struct{ *pulumi.OutputState }

func (CxEnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CxEnvironment)(nil)).Elem()
}

func (o CxEnvironmentOutput) ToCxEnvironmentOutput() CxEnvironmentOutput {
	return o
}

func (o CxEnvironmentOutput) ToCxEnvironmentOutputWithContext(ctx context.Context) CxEnvironmentOutput {
	return o
}

// The human-readable description of the environment. The maximum length is 500 characters. If exceeded, the request is rejected.
func (o CxEnvironmentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CxEnvironment) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The human-readable name of the environment (unique in an agent). Limit of 64 characters.
func (o CxEnvironmentOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *CxEnvironment) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The name of the environment.
func (o CxEnvironmentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CxEnvironment) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The Agent to create an Environment for.
// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>.
func (o CxEnvironmentOutput) Parent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CxEnvironment) pulumi.StringPtrOutput { return v.Parent }).(pulumi.StringPtrOutput)
}

// Update time of this environment. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
func (o CxEnvironmentOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *CxEnvironment) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// A list of configurations for flow versions. You should include version configs for all flows that are reachable from [Start Flow][Agent.start_flow] in the agent. Otherwise, an error will be returned.
// Structure is documented below.
func (o CxEnvironmentOutput) VersionConfigs() CxEnvironmentVersionConfigArrayOutput {
	return o.ApplyT(func(v *CxEnvironment) CxEnvironmentVersionConfigArrayOutput { return v.VersionConfigs }).(CxEnvironmentVersionConfigArrayOutput)
}

type CxEnvironmentArrayOutput struct{ *pulumi.OutputState }

func (CxEnvironmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CxEnvironment)(nil)).Elem()
}

func (o CxEnvironmentArrayOutput) ToCxEnvironmentArrayOutput() CxEnvironmentArrayOutput {
	return o
}

func (o CxEnvironmentArrayOutput) ToCxEnvironmentArrayOutputWithContext(ctx context.Context) CxEnvironmentArrayOutput {
	return o
}

func (o CxEnvironmentArrayOutput) Index(i pulumi.IntInput) CxEnvironmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CxEnvironment {
		return vs[0].([]*CxEnvironment)[vs[1].(int)]
	}).(CxEnvironmentOutput)
}

type CxEnvironmentMapOutput struct{ *pulumi.OutputState }

func (CxEnvironmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CxEnvironment)(nil)).Elem()
}

func (o CxEnvironmentMapOutput) ToCxEnvironmentMapOutput() CxEnvironmentMapOutput {
	return o
}

func (o CxEnvironmentMapOutput) ToCxEnvironmentMapOutputWithContext(ctx context.Context) CxEnvironmentMapOutput {
	return o
}

func (o CxEnvironmentMapOutput) MapIndex(k pulumi.StringInput) CxEnvironmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CxEnvironment {
		return vs[0].(map[string]*CxEnvironment)[vs[1].(string)]
	}).(CxEnvironmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CxEnvironmentInput)(nil)).Elem(), &CxEnvironment{})
	pulumi.RegisterInputType(reflect.TypeOf((*CxEnvironmentArrayInput)(nil)).Elem(), CxEnvironmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CxEnvironmentMapInput)(nil)).Elem(), CxEnvironmentMap{})
	pulumi.RegisterOutputType(CxEnvironmentOutput{})
	pulumi.RegisterOutputType(CxEnvironmentArrayOutput{})
	pulumi.RegisterOutputType(CxEnvironmentMapOutput{})
}
