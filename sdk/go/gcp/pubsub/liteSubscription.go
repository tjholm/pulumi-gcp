// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pubsub

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A named resource representing the stream of messages from a single,
// specific topic, to be delivered to the subscribing application.
//
// To get more information about Subscription, see:
//
// * [API documentation](https://cloud.google.com/pubsub/lite/docs/reference/rest/v1/admin.projects.locations.subscriptions)
// * How-to Guides
//   - [Managing Subscriptions](https://cloud.google.com/pubsub/lite/docs/subscriptions)
//
// ## Example Usage
// ### Pubsub Lite Subscription Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/organizations"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/pubsub"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			project, err := organizations.LookupProject(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			exampleLiteTopic, err := pubsub.NewLiteTopic(ctx, "exampleLiteTopic", &pubsub.LiteTopicArgs{
//				Project: *pulumi.String(project.Number),
//				PartitionConfig: &pubsub.LiteTopicPartitionConfigArgs{
//					Count: pulumi.Int(1),
//					Capacity: &pubsub.LiteTopicPartitionConfigCapacityArgs{
//						PublishMibPerSec:   pulumi.Int(4),
//						SubscribeMibPerSec: pulumi.Int(8),
//					},
//				},
//				RetentionConfig: &pubsub.LiteTopicRetentionConfigArgs{
//					PerPartitionBytes: pulumi.String("32212254720"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = pubsub.NewLiteSubscription(ctx, "exampleLiteSubscription", &pubsub.LiteSubscriptionArgs{
//				Topic: exampleLiteTopic.Name,
//				DeliveryConfig: &pubsub.LiteSubscriptionDeliveryConfigArgs{
//					DeliveryRequirement: pulumi.String("DELIVER_AFTER_STORED"),
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
// Subscription can be imported using any of these accepted formats* `projects/{{project}}/locations/{{zone}}/subscriptions/{{name}}` * `{{project}}/{{zone}}/{{name}}` * `{{zone}}/{{name}}` * `{{name}}` When using the `pulumi import` command, Subscription can be imported using one of the formats above. For example
//
// ```sh
//
//	$ pulumi import gcp:pubsub/liteSubscription:LiteSubscription default projects/{{project}}/locations/{{zone}}/subscriptions/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:pubsub/liteSubscription:LiteSubscription default {{project}}/{{zone}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:pubsub/liteSubscription:LiteSubscription default {{zone}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:pubsub/liteSubscription:LiteSubscription default {{name}}
//
// ```
type LiteSubscription struct {
	pulumi.CustomResourceState

	// The settings for this subscription's message delivery.
	// Structure is documented below.
	DeliveryConfig LiteSubscriptionDeliveryConfigPtrOutput `pulumi:"deliveryConfig"`
	// Name of the subscription.
	//
	// ***
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The region of the pubsub lite topic.
	Region pulumi.StringPtrOutput `pulumi:"region"`
	// A reference to a Topic resource.
	Topic pulumi.StringOutput `pulumi:"topic"`
	// The zone of the pubsub lite topic.
	Zone pulumi.StringPtrOutput `pulumi:"zone"`
}

// NewLiteSubscription registers a new resource with the given unique name, arguments, and options.
func NewLiteSubscription(ctx *pulumi.Context,
	name string, args *LiteSubscriptionArgs, opts ...pulumi.ResourceOption) (*LiteSubscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Topic == nil {
		return nil, errors.New("invalid value for required argument 'Topic'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LiteSubscription
	err := ctx.RegisterResource("gcp:pubsub/liteSubscription:LiteSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLiteSubscription gets an existing LiteSubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLiteSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LiteSubscriptionState, opts ...pulumi.ResourceOption) (*LiteSubscription, error) {
	var resource LiteSubscription
	err := ctx.ReadResource("gcp:pubsub/liteSubscription:LiteSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LiteSubscription resources.
type liteSubscriptionState struct {
	// The settings for this subscription's message delivery.
	// Structure is documented below.
	DeliveryConfig *LiteSubscriptionDeliveryConfig `pulumi:"deliveryConfig"`
	// Name of the subscription.
	//
	// ***
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region of the pubsub lite topic.
	Region *string `pulumi:"region"`
	// A reference to a Topic resource.
	Topic *string `pulumi:"topic"`
	// The zone of the pubsub lite topic.
	Zone *string `pulumi:"zone"`
}

type LiteSubscriptionState struct {
	// The settings for this subscription's message delivery.
	// Structure is documented below.
	DeliveryConfig LiteSubscriptionDeliveryConfigPtrInput
	// Name of the subscription.
	//
	// ***
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The region of the pubsub lite topic.
	Region pulumi.StringPtrInput
	// A reference to a Topic resource.
	Topic pulumi.StringPtrInput
	// The zone of the pubsub lite topic.
	Zone pulumi.StringPtrInput
}

func (LiteSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*liteSubscriptionState)(nil)).Elem()
}

type liteSubscriptionArgs struct {
	// The settings for this subscription's message delivery.
	// Structure is documented below.
	DeliveryConfig *LiteSubscriptionDeliveryConfig `pulumi:"deliveryConfig"`
	// Name of the subscription.
	//
	// ***
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region of the pubsub lite topic.
	Region *string `pulumi:"region"`
	// A reference to a Topic resource.
	Topic string `pulumi:"topic"`
	// The zone of the pubsub lite topic.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a LiteSubscription resource.
type LiteSubscriptionArgs struct {
	// The settings for this subscription's message delivery.
	// Structure is documented below.
	DeliveryConfig LiteSubscriptionDeliveryConfigPtrInput
	// Name of the subscription.
	//
	// ***
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The region of the pubsub lite topic.
	Region pulumi.StringPtrInput
	// A reference to a Topic resource.
	Topic pulumi.StringInput
	// The zone of the pubsub lite topic.
	Zone pulumi.StringPtrInput
}

func (LiteSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*liteSubscriptionArgs)(nil)).Elem()
}

type LiteSubscriptionInput interface {
	pulumi.Input

	ToLiteSubscriptionOutput() LiteSubscriptionOutput
	ToLiteSubscriptionOutputWithContext(ctx context.Context) LiteSubscriptionOutput
}

func (*LiteSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((**LiteSubscription)(nil)).Elem()
}

func (i *LiteSubscription) ToLiteSubscriptionOutput() LiteSubscriptionOutput {
	return i.ToLiteSubscriptionOutputWithContext(context.Background())
}

func (i *LiteSubscription) ToLiteSubscriptionOutputWithContext(ctx context.Context) LiteSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiteSubscriptionOutput)
}

// LiteSubscriptionArrayInput is an input type that accepts LiteSubscriptionArray and LiteSubscriptionArrayOutput values.
// You can construct a concrete instance of `LiteSubscriptionArrayInput` via:
//
//	LiteSubscriptionArray{ LiteSubscriptionArgs{...} }
type LiteSubscriptionArrayInput interface {
	pulumi.Input

	ToLiteSubscriptionArrayOutput() LiteSubscriptionArrayOutput
	ToLiteSubscriptionArrayOutputWithContext(context.Context) LiteSubscriptionArrayOutput
}

type LiteSubscriptionArray []LiteSubscriptionInput

func (LiteSubscriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LiteSubscription)(nil)).Elem()
}

func (i LiteSubscriptionArray) ToLiteSubscriptionArrayOutput() LiteSubscriptionArrayOutput {
	return i.ToLiteSubscriptionArrayOutputWithContext(context.Background())
}

func (i LiteSubscriptionArray) ToLiteSubscriptionArrayOutputWithContext(ctx context.Context) LiteSubscriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiteSubscriptionArrayOutput)
}

// LiteSubscriptionMapInput is an input type that accepts LiteSubscriptionMap and LiteSubscriptionMapOutput values.
// You can construct a concrete instance of `LiteSubscriptionMapInput` via:
//
//	LiteSubscriptionMap{ "key": LiteSubscriptionArgs{...} }
type LiteSubscriptionMapInput interface {
	pulumi.Input

	ToLiteSubscriptionMapOutput() LiteSubscriptionMapOutput
	ToLiteSubscriptionMapOutputWithContext(context.Context) LiteSubscriptionMapOutput
}

type LiteSubscriptionMap map[string]LiteSubscriptionInput

func (LiteSubscriptionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LiteSubscription)(nil)).Elem()
}

func (i LiteSubscriptionMap) ToLiteSubscriptionMapOutput() LiteSubscriptionMapOutput {
	return i.ToLiteSubscriptionMapOutputWithContext(context.Background())
}

func (i LiteSubscriptionMap) ToLiteSubscriptionMapOutputWithContext(ctx context.Context) LiteSubscriptionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiteSubscriptionMapOutput)
}

type LiteSubscriptionOutput struct{ *pulumi.OutputState }

func (LiteSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LiteSubscription)(nil)).Elem()
}

func (o LiteSubscriptionOutput) ToLiteSubscriptionOutput() LiteSubscriptionOutput {
	return o
}

func (o LiteSubscriptionOutput) ToLiteSubscriptionOutputWithContext(ctx context.Context) LiteSubscriptionOutput {
	return o
}

// The settings for this subscription's message delivery.
// Structure is documented below.
func (o LiteSubscriptionOutput) DeliveryConfig() LiteSubscriptionDeliveryConfigPtrOutput {
	return o.ApplyT(func(v *LiteSubscription) LiteSubscriptionDeliveryConfigPtrOutput { return v.DeliveryConfig }).(LiteSubscriptionDeliveryConfigPtrOutput)
}

// Name of the subscription.
//
// ***
func (o LiteSubscriptionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LiteSubscription) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o LiteSubscriptionOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *LiteSubscription) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The region of the pubsub lite topic.
func (o LiteSubscriptionOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiteSubscription) pulumi.StringPtrOutput { return v.Region }).(pulumi.StringPtrOutput)
}

// A reference to a Topic resource.
func (o LiteSubscriptionOutput) Topic() pulumi.StringOutput {
	return o.ApplyT(func(v *LiteSubscription) pulumi.StringOutput { return v.Topic }).(pulumi.StringOutput)
}

// The zone of the pubsub lite topic.
func (o LiteSubscriptionOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiteSubscription) pulumi.StringPtrOutput { return v.Zone }).(pulumi.StringPtrOutput)
}

type LiteSubscriptionArrayOutput struct{ *pulumi.OutputState }

func (LiteSubscriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LiteSubscription)(nil)).Elem()
}

func (o LiteSubscriptionArrayOutput) ToLiteSubscriptionArrayOutput() LiteSubscriptionArrayOutput {
	return o
}

func (o LiteSubscriptionArrayOutput) ToLiteSubscriptionArrayOutputWithContext(ctx context.Context) LiteSubscriptionArrayOutput {
	return o
}

func (o LiteSubscriptionArrayOutput) Index(i pulumi.IntInput) LiteSubscriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LiteSubscription {
		return vs[0].([]*LiteSubscription)[vs[1].(int)]
	}).(LiteSubscriptionOutput)
}

type LiteSubscriptionMapOutput struct{ *pulumi.OutputState }

func (LiteSubscriptionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LiteSubscription)(nil)).Elem()
}

func (o LiteSubscriptionMapOutput) ToLiteSubscriptionMapOutput() LiteSubscriptionMapOutput {
	return o
}

func (o LiteSubscriptionMapOutput) ToLiteSubscriptionMapOutputWithContext(ctx context.Context) LiteSubscriptionMapOutput {
	return o
}

func (o LiteSubscriptionMapOutput) MapIndex(k pulumi.StringInput) LiteSubscriptionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LiteSubscription {
		return vs[0].(map[string]*LiteSubscription)[vs[1].(string)]
	}).(LiteSubscriptionOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LiteSubscriptionInput)(nil)).Elem(), &LiteSubscription{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiteSubscriptionArrayInput)(nil)).Elem(), LiteSubscriptionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiteSubscriptionMapInput)(nil)).Elem(), LiteSubscriptionMap{})
	pulumi.RegisterOutputType(LiteSubscriptionOutput{})
	pulumi.RegisterOutputType(LiteSubscriptionArrayOutput{})
	pulumi.RegisterOutputType(LiteSubscriptionMapOutput{})
}
