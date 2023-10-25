// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Creates a new notification configuration on a specified bucket, establishing a flow of event notifications from GCS to a Cloud Pub/Sub topic.
//
//	For more information see
//
// [the official documentation](https://cloud.google.com/storage/docs/pubsub-notifications)
// and
// [API](https://cloud.google.com/storage/docs/json_api/v1/notifications).
//
// In order to enable notifications, a special Google Cloud Storage service account unique to the project
// must exist and have the IAM permission "projects.topics.publish" for a Cloud Pub/Sub topic in the project.
// This service account is not created automatically when a project is created.
// To ensure the service account exists and obtain its email address for use in granting the correct IAM permission, use the
// [`storage.getProjectServiceAccount`](https://www.terraform.io/docs/providers/google/d/storage_project_service_account.html)
// datasource's `emailAddress` value, and see below for an example of enabling notifications by granting the correct IAM permission.
// See [the notifications documentation](https://cloud.google.com/storage/docs/gsutil/commands/notification) for more details.
//
// > **NOTE**: This resource can affect your storage IAM policy. If you are using this in the same config as your storage IAM policy resources, consider
// making this resource dependent on those IAM resources via `dependsOn`. This will safeguard against errors due to IAM race conditions.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/pubsub"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/storage"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			gcsAccount, err := storage.GetProjectServiceAccount(ctx, nil, nil)
//			if err != nil {
//				return err
//			}
//			topic, err := pubsub.NewTopic(ctx, "topic", nil)
//			if err != nil {
//				return err
//			}
//			binding, err := pubsub.NewTopicIAMBinding(ctx, "binding", &pubsub.TopicIAMBindingArgs{
//				Topic: topic.ID(),
//				Role:  pulumi.String("roles/pubsub.publisher"),
//				Members: pulumi.StringArray{
//					pulumi.String(fmt.Sprintf("serviceAccount:%v", gcsAccount.EmailAddress)),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			bucket, err := storage.NewBucket(ctx, "bucket", &storage.BucketArgs{
//				Location: pulumi.String("US"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = storage.NewNotification(ctx, "notification", &storage.NotificationArgs{
//				Bucket:        bucket.Name,
//				PayloadFormat: pulumi.String("JSON_API_V1"),
//				Topic:         topic.ID(),
//				EventTypes: pulumi.StringArray{
//					pulumi.String("OBJECT_FINALIZE"),
//					pulumi.String("OBJECT_METADATA_UPDATE"),
//				},
//				CustomAttributes: pulumi.StringMap{
//					"new-attribute": pulumi.String("new-attribute-value"),
//				},
//			}, pulumi.DependsOn([]pulumi.Resource{
//				binding,
//			}))
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
// Storage notifications can be imported using the notification `id` in the format `<bucket_name>/notificationConfigs/<id>` e.g.
//
// ```sh
//
//	$ pulumi import gcp:storage/notification:Notification notification default_bucket/notificationConfigs/102
//
// ```
type Notification struct {
	pulumi.CustomResourceState

	// The name of the bucket.
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// A set of key/value attribute pairs to attach to each Cloud PubSub message published for this notification subscription
	CustomAttributes pulumi.StringMapOutput `pulumi:"customAttributes"`
	// List of event type filters for this notification config. If not specified, Cloud Storage will send notifications for all event types. The valid types are: `"OBJECT_FINALIZE"`, `"OBJECT_METADATA_UPDATE"`, `"OBJECT_DELETE"`, `"OBJECT_ARCHIVE"`
	EventTypes pulumi.StringArrayOutput `pulumi:"eventTypes"`
	// The ID of the created notification.
	NotificationId pulumi.StringOutput `pulumi:"notificationId"`
	// Specifies a prefix path filter for this notification config. Cloud Storage will only send notifications for objects in this bucket whose names begin with the specified prefix.
	ObjectNamePrefix pulumi.StringPtrOutput `pulumi:"objectNamePrefix"`
	// The desired content of the Payload. One of `"JSON_API_V1"` or `"NONE"`.
	PayloadFormat pulumi.StringOutput `pulumi:"payloadFormat"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The Cloud PubSub topic to which this subscription publishes. Expects either the
	// topic name, assumed to belong to the default GCP provider project, or the project-level name,
	// i.e. `projects/my-gcp-project/topics/my-topic` or `my-topic`. If the project is not set in the provider,
	// you will need to use the project-level name.
	//
	// ***
	Topic pulumi.StringOutput `pulumi:"topic"`
}

// NewNotification registers a new resource with the given unique name, arguments, and options.
func NewNotification(ctx *pulumi.Context,
	name string, args *NotificationArgs, opts ...pulumi.ResourceOption) (*Notification, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.PayloadFormat == nil {
		return nil, errors.New("invalid value for required argument 'PayloadFormat'")
	}
	if args.Topic == nil {
		return nil, errors.New("invalid value for required argument 'Topic'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Notification
	err := ctx.RegisterResource("gcp:storage/notification:Notification", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNotification gets an existing Notification resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNotification(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NotificationState, opts ...pulumi.ResourceOption) (*Notification, error) {
	var resource Notification
	err := ctx.ReadResource("gcp:storage/notification:Notification", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Notification resources.
type notificationState struct {
	// The name of the bucket.
	Bucket *string `pulumi:"bucket"`
	// A set of key/value attribute pairs to attach to each Cloud PubSub message published for this notification subscription
	CustomAttributes map[string]string `pulumi:"customAttributes"`
	// List of event type filters for this notification config. If not specified, Cloud Storage will send notifications for all event types. The valid types are: `"OBJECT_FINALIZE"`, `"OBJECT_METADATA_UPDATE"`, `"OBJECT_DELETE"`, `"OBJECT_ARCHIVE"`
	EventTypes []string `pulumi:"eventTypes"`
	// The ID of the created notification.
	NotificationId *string `pulumi:"notificationId"`
	// Specifies a prefix path filter for this notification config. Cloud Storage will only send notifications for objects in this bucket whose names begin with the specified prefix.
	ObjectNamePrefix *string `pulumi:"objectNamePrefix"`
	// The desired content of the Payload. One of `"JSON_API_V1"` or `"NONE"`.
	PayloadFormat *string `pulumi:"payloadFormat"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// The Cloud PubSub topic to which this subscription publishes. Expects either the
	// topic name, assumed to belong to the default GCP provider project, or the project-level name,
	// i.e. `projects/my-gcp-project/topics/my-topic` or `my-topic`. If the project is not set in the provider,
	// you will need to use the project-level name.
	//
	// ***
	Topic *string `pulumi:"topic"`
}

type NotificationState struct {
	// The name of the bucket.
	Bucket pulumi.StringPtrInput
	// A set of key/value attribute pairs to attach to each Cloud PubSub message published for this notification subscription
	CustomAttributes pulumi.StringMapInput
	// List of event type filters for this notification config. If not specified, Cloud Storage will send notifications for all event types. The valid types are: `"OBJECT_FINALIZE"`, `"OBJECT_METADATA_UPDATE"`, `"OBJECT_DELETE"`, `"OBJECT_ARCHIVE"`
	EventTypes pulumi.StringArrayInput
	// The ID of the created notification.
	NotificationId pulumi.StringPtrInput
	// Specifies a prefix path filter for this notification config. Cloud Storage will only send notifications for objects in this bucket whose names begin with the specified prefix.
	ObjectNamePrefix pulumi.StringPtrInput
	// The desired content of the Payload. One of `"JSON_API_V1"` or `"NONE"`.
	PayloadFormat pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// The Cloud PubSub topic to which this subscription publishes. Expects either the
	// topic name, assumed to belong to the default GCP provider project, or the project-level name,
	// i.e. `projects/my-gcp-project/topics/my-topic` or `my-topic`. If the project is not set in the provider,
	// you will need to use the project-level name.
	//
	// ***
	Topic pulumi.StringPtrInput
}

func (NotificationState) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationState)(nil)).Elem()
}

type notificationArgs struct {
	// The name of the bucket.
	Bucket string `pulumi:"bucket"`
	// A set of key/value attribute pairs to attach to each Cloud PubSub message published for this notification subscription
	CustomAttributes map[string]string `pulumi:"customAttributes"`
	// List of event type filters for this notification config. If not specified, Cloud Storage will send notifications for all event types. The valid types are: `"OBJECT_FINALIZE"`, `"OBJECT_METADATA_UPDATE"`, `"OBJECT_DELETE"`, `"OBJECT_ARCHIVE"`
	EventTypes []string `pulumi:"eventTypes"`
	// Specifies a prefix path filter for this notification config. Cloud Storage will only send notifications for objects in this bucket whose names begin with the specified prefix.
	ObjectNamePrefix *string `pulumi:"objectNamePrefix"`
	// The desired content of the Payload. One of `"JSON_API_V1"` or `"NONE"`.
	PayloadFormat string `pulumi:"payloadFormat"`
	// The Cloud PubSub topic to which this subscription publishes. Expects either the
	// topic name, assumed to belong to the default GCP provider project, or the project-level name,
	// i.e. `projects/my-gcp-project/topics/my-topic` or `my-topic`. If the project is not set in the provider,
	// you will need to use the project-level name.
	//
	// ***
	Topic string `pulumi:"topic"`
}

// The set of arguments for constructing a Notification resource.
type NotificationArgs struct {
	// The name of the bucket.
	Bucket pulumi.StringInput
	// A set of key/value attribute pairs to attach to each Cloud PubSub message published for this notification subscription
	CustomAttributes pulumi.StringMapInput
	// List of event type filters for this notification config. If not specified, Cloud Storage will send notifications for all event types. The valid types are: `"OBJECT_FINALIZE"`, `"OBJECT_METADATA_UPDATE"`, `"OBJECT_DELETE"`, `"OBJECT_ARCHIVE"`
	EventTypes pulumi.StringArrayInput
	// Specifies a prefix path filter for this notification config. Cloud Storage will only send notifications for objects in this bucket whose names begin with the specified prefix.
	ObjectNamePrefix pulumi.StringPtrInput
	// The desired content of the Payload. One of `"JSON_API_V1"` or `"NONE"`.
	PayloadFormat pulumi.StringInput
	// The Cloud PubSub topic to which this subscription publishes. Expects either the
	// topic name, assumed to belong to the default GCP provider project, or the project-level name,
	// i.e. `projects/my-gcp-project/topics/my-topic` or `my-topic`. If the project is not set in the provider,
	// you will need to use the project-level name.
	//
	// ***
	Topic pulumi.StringInput
}

func (NotificationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*notificationArgs)(nil)).Elem()
}

type NotificationInput interface {
	pulumi.Input

	ToNotificationOutput() NotificationOutput
	ToNotificationOutputWithContext(ctx context.Context) NotificationOutput
}

func (*Notification) ElementType() reflect.Type {
	return reflect.TypeOf((**Notification)(nil)).Elem()
}

func (i *Notification) ToNotificationOutput() NotificationOutput {
	return i.ToNotificationOutputWithContext(context.Background())
}

func (i *Notification) ToNotificationOutputWithContext(ctx context.Context) NotificationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationOutput)
}

func (i *Notification) ToOutput(ctx context.Context) pulumix.Output[*Notification] {
	return pulumix.Output[*Notification]{
		OutputState: i.ToNotificationOutputWithContext(ctx).OutputState,
	}
}

// NotificationArrayInput is an input type that accepts NotificationArray and NotificationArrayOutput values.
// You can construct a concrete instance of `NotificationArrayInput` via:
//
//	NotificationArray{ NotificationArgs{...} }
type NotificationArrayInput interface {
	pulumi.Input

	ToNotificationArrayOutput() NotificationArrayOutput
	ToNotificationArrayOutputWithContext(context.Context) NotificationArrayOutput
}

type NotificationArray []NotificationInput

func (NotificationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Notification)(nil)).Elem()
}

func (i NotificationArray) ToNotificationArrayOutput() NotificationArrayOutput {
	return i.ToNotificationArrayOutputWithContext(context.Background())
}

func (i NotificationArray) ToNotificationArrayOutputWithContext(ctx context.Context) NotificationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationArrayOutput)
}

func (i NotificationArray) ToOutput(ctx context.Context) pulumix.Output[[]*Notification] {
	return pulumix.Output[[]*Notification]{
		OutputState: i.ToNotificationArrayOutputWithContext(ctx).OutputState,
	}
}

// NotificationMapInput is an input type that accepts NotificationMap and NotificationMapOutput values.
// You can construct a concrete instance of `NotificationMapInput` via:
//
//	NotificationMap{ "key": NotificationArgs{...} }
type NotificationMapInput interface {
	pulumi.Input

	ToNotificationMapOutput() NotificationMapOutput
	ToNotificationMapOutputWithContext(context.Context) NotificationMapOutput
}

type NotificationMap map[string]NotificationInput

func (NotificationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Notification)(nil)).Elem()
}

func (i NotificationMap) ToNotificationMapOutput() NotificationMapOutput {
	return i.ToNotificationMapOutputWithContext(context.Background())
}

func (i NotificationMap) ToNotificationMapOutputWithContext(ctx context.Context) NotificationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationMapOutput)
}

func (i NotificationMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Notification] {
	return pulumix.Output[map[string]*Notification]{
		OutputState: i.ToNotificationMapOutputWithContext(ctx).OutputState,
	}
}

type NotificationOutput struct{ *pulumi.OutputState }

func (NotificationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Notification)(nil)).Elem()
}

func (o NotificationOutput) ToNotificationOutput() NotificationOutput {
	return o
}

func (o NotificationOutput) ToNotificationOutputWithContext(ctx context.Context) NotificationOutput {
	return o
}

func (o NotificationOutput) ToOutput(ctx context.Context) pulumix.Output[*Notification] {
	return pulumix.Output[*Notification]{
		OutputState: o.OutputState,
	}
}

// The name of the bucket.
func (o NotificationOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v *Notification) pulumi.StringOutput { return v.Bucket }).(pulumi.StringOutput)
}

// A set of key/value attribute pairs to attach to each Cloud PubSub message published for this notification subscription
func (o NotificationOutput) CustomAttributes() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Notification) pulumi.StringMapOutput { return v.CustomAttributes }).(pulumi.StringMapOutput)
}

// List of event type filters for this notification config. If not specified, Cloud Storage will send notifications for all event types. The valid types are: `"OBJECT_FINALIZE"`, `"OBJECT_METADATA_UPDATE"`, `"OBJECT_DELETE"`, `"OBJECT_ARCHIVE"`
func (o NotificationOutput) EventTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Notification) pulumi.StringArrayOutput { return v.EventTypes }).(pulumi.StringArrayOutput)
}

// The ID of the created notification.
func (o NotificationOutput) NotificationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Notification) pulumi.StringOutput { return v.NotificationId }).(pulumi.StringOutput)
}

// Specifies a prefix path filter for this notification config. Cloud Storage will only send notifications for objects in this bucket whose names begin with the specified prefix.
func (o NotificationOutput) ObjectNamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Notification) pulumi.StringPtrOutput { return v.ObjectNamePrefix }).(pulumi.StringPtrOutput)
}

// The desired content of the Payload. One of `"JSON_API_V1"` or `"NONE"`.
func (o NotificationOutput) PayloadFormat() pulumi.StringOutput {
	return o.ApplyT(func(v *Notification) pulumi.StringOutput { return v.PayloadFormat }).(pulumi.StringOutput)
}

// The URI of the created resource.
func (o NotificationOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v *Notification) pulumi.StringOutput { return v.SelfLink }).(pulumi.StringOutput)
}

// The Cloud PubSub topic to which this subscription publishes. Expects either the
// topic name, assumed to belong to the default GCP provider project, or the project-level name,
// i.e. `projects/my-gcp-project/topics/my-topic` or `my-topic`. If the project is not set in the provider,
// you will need to use the project-level name.
//
// ***
func (o NotificationOutput) Topic() pulumi.StringOutput {
	return o.ApplyT(func(v *Notification) pulumi.StringOutput { return v.Topic }).(pulumi.StringOutput)
}

type NotificationArrayOutput struct{ *pulumi.OutputState }

func (NotificationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Notification)(nil)).Elem()
}

func (o NotificationArrayOutput) ToNotificationArrayOutput() NotificationArrayOutput {
	return o
}

func (o NotificationArrayOutput) ToNotificationArrayOutputWithContext(ctx context.Context) NotificationArrayOutput {
	return o
}

func (o NotificationArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Notification] {
	return pulumix.Output[[]*Notification]{
		OutputState: o.OutputState,
	}
}

func (o NotificationArrayOutput) Index(i pulumi.IntInput) NotificationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Notification {
		return vs[0].([]*Notification)[vs[1].(int)]
	}).(NotificationOutput)
}

type NotificationMapOutput struct{ *pulumi.OutputState }

func (NotificationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Notification)(nil)).Elem()
}

func (o NotificationMapOutput) ToNotificationMapOutput() NotificationMapOutput {
	return o
}

func (o NotificationMapOutput) ToNotificationMapOutputWithContext(ctx context.Context) NotificationMapOutput {
	return o
}

func (o NotificationMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Notification] {
	return pulumix.Output[map[string]*Notification]{
		OutputState: o.OutputState,
	}
}

func (o NotificationMapOutput) MapIndex(k pulumi.StringInput) NotificationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Notification {
		return vs[0].(map[string]*Notification)[vs[1].(string)]
	}).(NotificationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationInput)(nil)).Elem(), &Notification{})
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationArrayInput)(nil)).Elem(), NotificationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationMapInput)(nil)).Elem(), NotificationMap{})
	pulumi.RegisterOutputType(NotificationOutput{})
	pulumi.RegisterOutputType(NotificationArrayOutput{})
	pulumi.RegisterOutputType(NotificationMapOutput{})
}
