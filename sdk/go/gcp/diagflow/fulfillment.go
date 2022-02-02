// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package diagflow

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// By default, your agent responds to a matched intent with a static response. If you're using one of the integration options, you can provide a more dynamic response by using fulfillment. When you enable fulfillment for an intent, Dialogflow responds to that intent by calling a service that you define. For example, if an end-user wants to schedule a haircut on Friday, your service can check your database and respond to the end-user with availability information for Friday.
//
// To get more information about Fulfillment, see:
//
// * [API documentation](https://cloud.google.com/dialogflow/es/docs/reference/rest/v2/projects.agent/getFulfillment)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/dialogflow/es/docs/fulfillment-overview)
//
// ## Example Usage
// ### Dialogflow Fulfillment Basic
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
// 		_, err = diagflow.NewFulfillment(ctx, "basicFulfillment", &diagflow.FulfillmentArgs{
// 			DisplayName: pulumi.String("basic-fulfillment"),
// 			Enabled:     pulumi.Bool(true),
// 			GenericWebService: &diagflow.FulfillmentGenericWebServiceArgs{
// 				Uri:      pulumi.String("https://google.com"),
// 				Username: pulumi.String("admin"),
// 				Password: pulumi.String("password"),
// 				RequestHeaders: pulumi.StringMap{
// 					"name": pulumi.String("wrench"),
// 				},
// 			},
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
// Fulfillment can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:diagflow/fulfillment:Fulfillment default {{name}}
// ```
type Fulfillment struct {
	pulumi.CustomResourceState

	// The human-readable name of the fulfillment, unique within the agent.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Whether fulfillment is enabled.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The field defines whether the fulfillment is enabled for certain features.
	// Structure is documented below.
	Features FulfillmentFeatureArrayOutput `pulumi:"features"`
	// Represents configuration for a generic web service. Dialogflow supports two mechanisms for authentications: - Basic authentication with username and password. - Authentication with additional authentication headers.
	// Structure is documented below.
	GenericWebService FulfillmentGenericWebServicePtrOutput `pulumi:"genericWebService"`
	// The unique identifier of the fulfillment. Format: projects/<Project ID>/agent/fulfillment - projects/<Project
	// ID>/locations/<Location ID>/agent/fulfillment
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewFulfillment registers a new resource with the given unique name, arguments, and options.
func NewFulfillment(ctx *pulumi.Context,
	name string, args *FulfillmentArgs, opts ...pulumi.ResourceOption) (*Fulfillment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	var resource Fulfillment
	err := ctx.RegisterResource("gcp:diagflow/fulfillment:Fulfillment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFulfillment gets an existing Fulfillment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFulfillment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FulfillmentState, opts ...pulumi.ResourceOption) (*Fulfillment, error) {
	var resource Fulfillment
	err := ctx.ReadResource("gcp:diagflow/fulfillment:Fulfillment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Fulfillment resources.
type fulfillmentState struct {
	// The human-readable name of the fulfillment, unique within the agent.
	DisplayName *string `pulumi:"displayName"`
	// Whether fulfillment is enabled.
	Enabled *bool `pulumi:"enabled"`
	// The field defines whether the fulfillment is enabled for certain features.
	// Structure is documented below.
	Features []FulfillmentFeature `pulumi:"features"`
	// Represents configuration for a generic web service. Dialogflow supports two mechanisms for authentications: - Basic authentication with username and password. - Authentication with additional authentication headers.
	// Structure is documented below.
	GenericWebService *FulfillmentGenericWebService `pulumi:"genericWebService"`
	// The unique identifier of the fulfillment. Format: projects/<Project ID>/agent/fulfillment - projects/<Project
	// ID>/locations/<Location ID>/agent/fulfillment
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

type FulfillmentState struct {
	// The human-readable name of the fulfillment, unique within the agent.
	DisplayName pulumi.StringPtrInput
	// Whether fulfillment is enabled.
	Enabled pulumi.BoolPtrInput
	// The field defines whether the fulfillment is enabled for certain features.
	// Structure is documented below.
	Features FulfillmentFeatureArrayInput
	// Represents configuration for a generic web service. Dialogflow supports two mechanisms for authentications: - Basic authentication with username and password. - Authentication with additional authentication headers.
	// Structure is documented below.
	GenericWebService FulfillmentGenericWebServicePtrInput
	// The unique identifier of the fulfillment. Format: projects/<Project ID>/agent/fulfillment - projects/<Project
	// ID>/locations/<Location ID>/agent/fulfillment
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (FulfillmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*fulfillmentState)(nil)).Elem()
}

type fulfillmentArgs struct {
	// The human-readable name of the fulfillment, unique within the agent.
	DisplayName string `pulumi:"displayName"`
	// Whether fulfillment is enabled.
	Enabled *bool `pulumi:"enabled"`
	// The field defines whether the fulfillment is enabled for certain features.
	// Structure is documented below.
	Features []FulfillmentFeature `pulumi:"features"`
	// Represents configuration for a generic web service. Dialogflow supports two mechanisms for authentications: - Basic authentication with username and password. - Authentication with additional authentication headers.
	// Structure is documented below.
	GenericWebService *FulfillmentGenericWebService `pulumi:"genericWebService"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a Fulfillment resource.
type FulfillmentArgs struct {
	// The human-readable name of the fulfillment, unique within the agent.
	DisplayName pulumi.StringInput
	// Whether fulfillment is enabled.
	Enabled pulumi.BoolPtrInput
	// The field defines whether the fulfillment is enabled for certain features.
	// Structure is documented below.
	Features FulfillmentFeatureArrayInput
	// Represents configuration for a generic web service. Dialogflow supports two mechanisms for authentications: - Basic authentication with username and password. - Authentication with additional authentication headers.
	// Structure is documented below.
	GenericWebService FulfillmentGenericWebServicePtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
}

func (FulfillmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fulfillmentArgs)(nil)).Elem()
}

type FulfillmentInput interface {
	pulumi.Input

	ToFulfillmentOutput() FulfillmentOutput
	ToFulfillmentOutputWithContext(ctx context.Context) FulfillmentOutput
}

func (*Fulfillment) ElementType() reflect.Type {
	return reflect.TypeOf((**Fulfillment)(nil)).Elem()
}

func (i *Fulfillment) ToFulfillmentOutput() FulfillmentOutput {
	return i.ToFulfillmentOutputWithContext(context.Background())
}

func (i *Fulfillment) ToFulfillmentOutputWithContext(ctx context.Context) FulfillmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FulfillmentOutput)
}

// FulfillmentArrayInput is an input type that accepts FulfillmentArray and FulfillmentArrayOutput values.
// You can construct a concrete instance of `FulfillmentArrayInput` via:
//
//          FulfillmentArray{ FulfillmentArgs{...} }
type FulfillmentArrayInput interface {
	pulumi.Input

	ToFulfillmentArrayOutput() FulfillmentArrayOutput
	ToFulfillmentArrayOutputWithContext(context.Context) FulfillmentArrayOutput
}

type FulfillmentArray []FulfillmentInput

func (FulfillmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Fulfillment)(nil)).Elem()
}

func (i FulfillmentArray) ToFulfillmentArrayOutput() FulfillmentArrayOutput {
	return i.ToFulfillmentArrayOutputWithContext(context.Background())
}

func (i FulfillmentArray) ToFulfillmentArrayOutputWithContext(ctx context.Context) FulfillmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FulfillmentArrayOutput)
}

// FulfillmentMapInput is an input type that accepts FulfillmentMap and FulfillmentMapOutput values.
// You can construct a concrete instance of `FulfillmentMapInput` via:
//
//          FulfillmentMap{ "key": FulfillmentArgs{...} }
type FulfillmentMapInput interface {
	pulumi.Input

	ToFulfillmentMapOutput() FulfillmentMapOutput
	ToFulfillmentMapOutputWithContext(context.Context) FulfillmentMapOutput
}

type FulfillmentMap map[string]FulfillmentInput

func (FulfillmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Fulfillment)(nil)).Elem()
}

func (i FulfillmentMap) ToFulfillmentMapOutput() FulfillmentMapOutput {
	return i.ToFulfillmentMapOutputWithContext(context.Background())
}

func (i FulfillmentMap) ToFulfillmentMapOutputWithContext(ctx context.Context) FulfillmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FulfillmentMapOutput)
}

type FulfillmentOutput struct{ *pulumi.OutputState }

func (FulfillmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Fulfillment)(nil)).Elem()
}

func (o FulfillmentOutput) ToFulfillmentOutput() FulfillmentOutput {
	return o
}

func (o FulfillmentOutput) ToFulfillmentOutputWithContext(ctx context.Context) FulfillmentOutput {
	return o
}

type FulfillmentArrayOutput struct{ *pulumi.OutputState }

func (FulfillmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Fulfillment)(nil)).Elem()
}

func (o FulfillmentArrayOutput) ToFulfillmentArrayOutput() FulfillmentArrayOutput {
	return o
}

func (o FulfillmentArrayOutput) ToFulfillmentArrayOutputWithContext(ctx context.Context) FulfillmentArrayOutput {
	return o
}

func (o FulfillmentArrayOutput) Index(i pulumi.IntInput) FulfillmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Fulfillment {
		return vs[0].([]*Fulfillment)[vs[1].(int)]
	}).(FulfillmentOutput)
}

type FulfillmentMapOutput struct{ *pulumi.OutputState }

func (FulfillmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Fulfillment)(nil)).Elem()
}

func (o FulfillmentMapOutput) ToFulfillmentMapOutput() FulfillmentMapOutput {
	return o
}

func (o FulfillmentMapOutput) ToFulfillmentMapOutputWithContext(ctx context.Context) FulfillmentMapOutput {
	return o
}

func (o FulfillmentMapOutput) MapIndex(k pulumi.StringInput) FulfillmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Fulfillment {
		return vs[0].(map[string]*Fulfillment)[vs[1].(string)]
	}).(FulfillmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FulfillmentInput)(nil)).Elem(), &Fulfillment{})
	pulumi.RegisterInputType(reflect.TypeOf((*FulfillmentArrayInput)(nil)).Elem(), FulfillmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FulfillmentMapInput)(nil)).Elem(), FulfillmentMap{})
	pulumi.RegisterOutputType(FulfillmentOutput{})
	pulumi.RegisterOutputType(FulfillmentArrayOutput{})
	pulumi.RegisterOutputType(FulfillmentMapOutput{})
}
