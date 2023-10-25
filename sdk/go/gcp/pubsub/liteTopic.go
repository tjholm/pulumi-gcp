// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package pubsub

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// A named resource to which messages are sent by publishers.
//
// To get more information about Topic, see:
//
// * [API documentation](https://cloud.google.com/pubsub/lite/docs/reference/rest/v1/admin.projects.locations.topics)
// * How-to Guides
//   - [Managing Topics](https://cloud.google.com/pubsub/lite/docs/topics)
//
// ## Example Usage
// ### Pubsub Lite Topic Basic
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/pubsub"
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
//			exampleLiteReservation, err := pubsub.NewLiteReservation(ctx, "exampleLiteReservation", &pubsub.LiteReservationArgs{
//				Project:            *pulumi.String(project.Number),
//				ThroughputCapacity: pulumi.Int(2),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = pubsub.NewLiteTopic(ctx, "exampleLiteTopic", &pubsub.LiteTopicArgs{
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
//				ReservationConfig: &pubsub.LiteTopicReservationConfigArgs{
//					ThroughputReservation: exampleLiteReservation.Name,
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
// # Topic can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:pubsub/liteTopic:LiteTopic default projects/{{project}}/locations/{{zone}}/topics/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:pubsub/liteTopic:LiteTopic default {{project}}/{{zone}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:pubsub/liteTopic:LiteTopic default {{zone}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:pubsub/liteTopic:LiteTopic default {{name}}
//
// ```
type LiteTopic struct {
	pulumi.CustomResourceState

	// Name of the topic.
	//
	// ***
	Name pulumi.StringOutput `pulumi:"name"`
	// The settings for this topic's partitions.
	// Structure is documented below.
	PartitionConfig LiteTopicPartitionConfigPtrOutput `pulumi:"partitionConfig"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The region of the pubsub lite topic.
	Region pulumi.StringPtrOutput `pulumi:"region"`
	// The settings for this topic's Reservation usage.
	// Structure is documented below.
	ReservationConfig LiteTopicReservationConfigPtrOutput `pulumi:"reservationConfig"`
	// The settings for a topic's message retention.
	// Structure is documented below.
	RetentionConfig LiteTopicRetentionConfigPtrOutput `pulumi:"retentionConfig"`
	// The zone of the pubsub lite topic.
	Zone pulumi.StringPtrOutput `pulumi:"zone"`
}

// NewLiteTopic registers a new resource with the given unique name, arguments, and options.
func NewLiteTopic(ctx *pulumi.Context,
	name string, args *LiteTopicArgs, opts ...pulumi.ResourceOption) (*LiteTopic, error) {
	if args == nil {
		args = &LiteTopicArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LiteTopic
	err := ctx.RegisterResource("gcp:pubsub/liteTopic:LiteTopic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLiteTopic gets an existing LiteTopic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLiteTopic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LiteTopicState, opts ...pulumi.ResourceOption) (*LiteTopic, error) {
	var resource LiteTopic
	err := ctx.ReadResource("gcp:pubsub/liteTopic:LiteTopic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LiteTopic resources.
type liteTopicState struct {
	// Name of the topic.
	//
	// ***
	Name *string `pulumi:"name"`
	// The settings for this topic's partitions.
	// Structure is documented below.
	PartitionConfig *LiteTopicPartitionConfig `pulumi:"partitionConfig"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region of the pubsub lite topic.
	Region *string `pulumi:"region"`
	// The settings for this topic's Reservation usage.
	// Structure is documented below.
	ReservationConfig *LiteTopicReservationConfig `pulumi:"reservationConfig"`
	// The settings for a topic's message retention.
	// Structure is documented below.
	RetentionConfig *LiteTopicRetentionConfig `pulumi:"retentionConfig"`
	// The zone of the pubsub lite topic.
	Zone *string `pulumi:"zone"`
}

type LiteTopicState struct {
	// Name of the topic.
	//
	// ***
	Name pulumi.StringPtrInput
	// The settings for this topic's partitions.
	// Structure is documented below.
	PartitionConfig LiteTopicPartitionConfigPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The region of the pubsub lite topic.
	Region pulumi.StringPtrInput
	// The settings for this topic's Reservation usage.
	// Structure is documented below.
	ReservationConfig LiteTopicReservationConfigPtrInput
	// The settings for a topic's message retention.
	// Structure is documented below.
	RetentionConfig LiteTopicRetentionConfigPtrInput
	// The zone of the pubsub lite topic.
	Zone pulumi.StringPtrInput
}

func (LiteTopicState) ElementType() reflect.Type {
	return reflect.TypeOf((*liteTopicState)(nil)).Elem()
}

type liteTopicArgs struct {
	// Name of the topic.
	//
	// ***
	Name *string `pulumi:"name"`
	// The settings for this topic's partitions.
	// Structure is documented below.
	PartitionConfig *LiteTopicPartitionConfig `pulumi:"partitionConfig"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region of the pubsub lite topic.
	Region *string `pulumi:"region"`
	// The settings for this topic's Reservation usage.
	// Structure is documented below.
	ReservationConfig *LiteTopicReservationConfig `pulumi:"reservationConfig"`
	// The settings for a topic's message retention.
	// Structure is documented below.
	RetentionConfig *LiteTopicRetentionConfig `pulumi:"retentionConfig"`
	// The zone of the pubsub lite topic.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a LiteTopic resource.
type LiteTopicArgs struct {
	// Name of the topic.
	//
	// ***
	Name pulumi.StringPtrInput
	// The settings for this topic's partitions.
	// Structure is documented below.
	PartitionConfig LiteTopicPartitionConfigPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The region of the pubsub lite topic.
	Region pulumi.StringPtrInput
	// The settings for this topic's Reservation usage.
	// Structure is documented below.
	ReservationConfig LiteTopicReservationConfigPtrInput
	// The settings for a topic's message retention.
	// Structure is documented below.
	RetentionConfig LiteTopicRetentionConfigPtrInput
	// The zone of the pubsub lite topic.
	Zone pulumi.StringPtrInput
}

func (LiteTopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*liteTopicArgs)(nil)).Elem()
}

type LiteTopicInput interface {
	pulumi.Input

	ToLiteTopicOutput() LiteTopicOutput
	ToLiteTopicOutputWithContext(ctx context.Context) LiteTopicOutput
}

func (*LiteTopic) ElementType() reflect.Type {
	return reflect.TypeOf((**LiteTopic)(nil)).Elem()
}

func (i *LiteTopic) ToLiteTopicOutput() LiteTopicOutput {
	return i.ToLiteTopicOutputWithContext(context.Background())
}

func (i *LiteTopic) ToLiteTopicOutputWithContext(ctx context.Context) LiteTopicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiteTopicOutput)
}

func (i *LiteTopic) ToOutput(ctx context.Context) pulumix.Output[*LiteTopic] {
	return pulumix.Output[*LiteTopic]{
		OutputState: i.ToLiteTopicOutputWithContext(ctx).OutputState,
	}
}

// LiteTopicArrayInput is an input type that accepts LiteTopicArray and LiteTopicArrayOutput values.
// You can construct a concrete instance of `LiteTopicArrayInput` via:
//
//	LiteTopicArray{ LiteTopicArgs{...} }
type LiteTopicArrayInput interface {
	pulumi.Input

	ToLiteTopicArrayOutput() LiteTopicArrayOutput
	ToLiteTopicArrayOutputWithContext(context.Context) LiteTopicArrayOutput
}

type LiteTopicArray []LiteTopicInput

func (LiteTopicArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LiteTopic)(nil)).Elem()
}

func (i LiteTopicArray) ToLiteTopicArrayOutput() LiteTopicArrayOutput {
	return i.ToLiteTopicArrayOutputWithContext(context.Background())
}

func (i LiteTopicArray) ToLiteTopicArrayOutputWithContext(ctx context.Context) LiteTopicArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiteTopicArrayOutput)
}

func (i LiteTopicArray) ToOutput(ctx context.Context) pulumix.Output[[]*LiteTopic] {
	return pulumix.Output[[]*LiteTopic]{
		OutputState: i.ToLiteTopicArrayOutputWithContext(ctx).OutputState,
	}
}

// LiteTopicMapInput is an input type that accepts LiteTopicMap and LiteTopicMapOutput values.
// You can construct a concrete instance of `LiteTopicMapInput` via:
//
//	LiteTopicMap{ "key": LiteTopicArgs{...} }
type LiteTopicMapInput interface {
	pulumi.Input

	ToLiteTopicMapOutput() LiteTopicMapOutput
	ToLiteTopicMapOutputWithContext(context.Context) LiteTopicMapOutput
}

type LiteTopicMap map[string]LiteTopicInput

func (LiteTopicMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LiteTopic)(nil)).Elem()
}

func (i LiteTopicMap) ToLiteTopicMapOutput() LiteTopicMapOutput {
	return i.ToLiteTopicMapOutputWithContext(context.Background())
}

func (i LiteTopicMap) ToLiteTopicMapOutputWithContext(ctx context.Context) LiteTopicMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LiteTopicMapOutput)
}

func (i LiteTopicMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*LiteTopic] {
	return pulumix.Output[map[string]*LiteTopic]{
		OutputState: i.ToLiteTopicMapOutputWithContext(ctx).OutputState,
	}
}

type LiteTopicOutput struct{ *pulumi.OutputState }

func (LiteTopicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LiteTopic)(nil)).Elem()
}

func (o LiteTopicOutput) ToLiteTopicOutput() LiteTopicOutput {
	return o
}

func (o LiteTopicOutput) ToLiteTopicOutputWithContext(ctx context.Context) LiteTopicOutput {
	return o
}

func (o LiteTopicOutput) ToOutput(ctx context.Context) pulumix.Output[*LiteTopic] {
	return pulumix.Output[*LiteTopic]{
		OutputState: o.OutputState,
	}
}

// Name of the topic.
//
// ***
func (o LiteTopicOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LiteTopic) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The settings for this topic's partitions.
// Structure is documented below.
func (o LiteTopicOutput) PartitionConfig() LiteTopicPartitionConfigPtrOutput {
	return o.ApplyT(func(v *LiteTopic) LiteTopicPartitionConfigPtrOutput { return v.PartitionConfig }).(LiteTopicPartitionConfigPtrOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o LiteTopicOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *LiteTopic) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The region of the pubsub lite topic.
func (o LiteTopicOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiteTopic) pulumi.StringPtrOutput { return v.Region }).(pulumi.StringPtrOutput)
}

// The settings for this topic's Reservation usage.
// Structure is documented below.
func (o LiteTopicOutput) ReservationConfig() LiteTopicReservationConfigPtrOutput {
	return o.ApplyT(func(v *LiteTopic) LiteTopicReservationConfigPtrOutput { return v.ReservationConfig }).(LiteTopicReservationConfigPtrOutput)
}

// The settings for a topic's message retention.
// Structure is documented below.
func (o LiteTopicOutput) RetentionConfig() LiteTopicRetentionConfigPtrOutput {
	return o.ApplyT(func(v *LiteTopic) LiteTopicRetentionConfigPtrOutput { return v.RetentionConfig }).(LiteTopicRetentionConfigPtrOutput)
}

// The zone of the pubsub lite topic.
func (o LiteTopicOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LiteTopic) pulumi.StringPtrOutput { return v.Zone }).(pulumi.StringPtrOutput)
}

type LiteTopicArrayOutput struct{ *pulumi.OutputState }

func (LiteTopicArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LiteTopic)(nil)).Elem()
}

func (o LiteTopicArrayOutput) ToLiteTopicArrayOutput() LiteTopicArrayOutput {
	return o
}

func (o LiteTopicArrayOutput) ToLiteTopicArrayOutputWithContext(ctx context.Context) LiteTopicArrayOutput {
	return o
}

func (o LiteTopicArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*LiteTopic] {
	return pulumix.Output[[]*LiteTopic]{
		OutputState: o.OutputState,
	}
}

func (o LiteTopicArrayOutput) Index(i pulumi.IntInput) LiteTopicOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LiteTopic {
		return vs[0].([]*LiteTopic)[vs[1].(int)]
	}).(LiteTopicOutput)
}

type LiteTopicMapOutput struct{ *pulumi.OutputState }

func (LiteTopicMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LiteTopic)(nil)).Elem()
}

func (o LiteTopicMapOutput) ToLiteTopicMapOutput() LiteTopicMapOutput {
	return o
}

func (o LiteTopicMapOutput) ToLiteTopicMapOutputWithContext(ctx context.Context) LiteTopicMapOutput {
	return o
}

func (o LiteTopicMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*LiteTopic] {
	return pulumix.Output[map[string]*LiteTopic]{
		OutputState: o.OutputState,
	}
}

func (o LiteTopicMapOutput) MapIndex(k pulumi.StringInput) LiteTopicOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LiteTopic {
		return vs[0].(map[string]*LiteTopic)[vs[1].(string)]
	}).(LiteTopicOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LiteTopicInput)(nil)).Elem(), &LiteTopic{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiteTopicArrayInput)(nil)).Elem(), LiteTopicArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LiteTopicMapInput)(nil)).Elem(), LiteTopicMap{})
	pulumi.RegisterOutputType(LiteTopicOutput{})
	pulumi.RegisterOutputType(LiteTopicArrayOutput{})
	pulumi.RegisterOutputType(LiteTopicMapOutput{})
}
