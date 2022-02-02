// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package logging

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a billing account level logging bucket config. For more information see
// [the official logging documentation](https://cloud.google.com/logging/docs/) and
// [Storing Logs](https://cloud.google.com/logging/docs/storage).
//
// > **Note:** Logging buckets are automatically created for a given folder, project, organization, billingAccount and cannot be deleted. Creating a resource of this type will acquire and update the resource that already exists at the desired location. These buckets cannot be removed so deleting this resource will remove the bucket config from your state but will leave the logging bucket unchanged. The buckets that are currently automatically created are "_Default" and "_Required".
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/logging"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "00AA00-000AAA-00AA0A"
// 		_default, err := organizations.GetBillingAccount(ctx, &organizations.GetBillingAccountArgs{
// 			BillingAccount: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = logging.NewBillingAccountBucketConfig(ctx, "basic", &logging.BillingAccountBucketConfigArgs{
// 			BillingAccount: pulumi.String(_default.BillingAccount),
// 			Location:       pulumi.String("global"),
// 			RetentionDays:  pulumi.Int(30),
// 			BucketId:       pulumi.String("_Default"),
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
// This resource can be imported using the following format
//
// ```sh
//  $ pulumi import gcp:logging/billingAccountBucketConfig:BillingAccountBucketConfig default billingAccounts/{{billingAccount}}/locations/{{location}}/buckets/{{bucket_id}}
// ```
type BillingAccountBucketConfig struct {
	pulumi.CustomResourceState

	// The parent resource that contains the logging bucket.
	BillingAccount pulumi.StringOutput `pulumi:"billingAccount"`
	// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
	BucketId pulumi.StringOutput `pulumi:"bucketId"`
	// Describes this bucket.
	Description pulumi.StringOutput `pulumi:"description"`
	// The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
	LifecycleState pulumi.StringOutput `pulumi:"lifecycleState"`
	// The location of the bucket.
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name of the bucket. For example: "projects/my-project-id/locations/my-location/buckets/my-bucket-id"
	Name pulumi.StringOutput `pulumi:"name"`
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used. Bucket retention can not be increased on buckets outside of projects.
	RetentionDays pulumi.IntPtrOutput `pulumi:"retentionDays"`
}

// NewBillingAccountBucketConfig registers a new resource with the given unique name, arguments, and options.
func NewBillingAccountBucketConfig(ctx *pulumi.Context,
	name string, args *BillingAccountBucketConfigArgs, opts ...pulumi.ResourceOption) (*BillingAccountBucketConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BillingAccount == nil {
		return nil, errors.New("invalid value for required argument 'BillingAccount'")
	}
	if args.BucketId == nil {
		return nil, errors.New("invalid value for required argument 'BucketId'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	var resource BillingAccountBucketConfig
	err := ctx.RegisterResource("gcp:logging/billingAccountBucketConfig:BillingAccountBucketConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBillingAccountBucketConfig gets an existing BillingAccountBucketConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBillingAccountBucketConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BillingAccountBucketConfigState, opts ...pulumi.ResourceOption) (*BillingAccountBucketConfig, error) {
	var resource BillingAccountBucketConfig
	err := ctx.ReadResource("gcp:logging/billingAccountBucketConfig:BillingAccountBucketConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BillingAccountBucketConfig resources.
type billingAccountBucketConfigState struct {
	// The parent resource that contains the logging bucket.
	BillingAccount *string `pulumi:"billingAccount"`
	// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
	BucketId *string `pulumi:"bucketId"`
	// Describes this bucket.
	Description *string `pulumi:"description"`
	// The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
	LifecycleState *string `pulumi:"lifecycleState"`
	// The location of the bucket.
	Location *string `pulumi:"location"`
	// The resource name of the bucket. For example: "projects/my-project-id/locations/my-location/buckets/my-bucket-id"
	Name *string `pulumi:"name"`
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used. Bucket retention can not be increased on buckets outside of projects.
	RetentionDays *int `pulumi:"retentionDays"`
}

type BillingAccountBucketConfigState struct {
	// The parent resource that contains the logging bucket.
	BillingAccount pulumi.StringPtrInput
	// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
	BucketId pulumi.StringPtrInput
	// Describes this bucket.
	Description pulumi.StringPtrInput
	// The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
	LifecycleState pulumi.StringPtrInput
	// The location of the bucket.
	Location pulumi.StringPtrInput
	// The resource name of the bucket. For example: "projects/my-project-id/locations/my-location/buckets/my-bucket-id"
	Name pulumi.StringPtrInput
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used. Bucket retention can not be increased on buckets outside of projects.
	RetentionDays pulumi.IntPtrInput
}

func (BillingAccountBucketConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*billingAccountBucketConfigState)(nil)).Elem()
}

type billingAccountBucketConfigArgs struct {
	// The parent resource that contains the logging bucket.
	BillingAccount string `pulumi:"billingAccount"`
	// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
	BucketId string `pulumi:"bucketId"`
	// Describes this bucket.
	Description *string `pulumi:"description"`
	// The location of the bucket.
	Location string `pulumi:"location"`
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used. Bucket retention can not be increased on buckets outside of projects.
	RetentionDays *int `pulumi:"retentionDays"`
}

// The set of arguments for constructing a BillingAccountBucketConfig resource.
type BillingAccountBucketConfigArgs struct {
	// The parent resource that contains the logging bucket.
	BillingAccount pulumi.StringInput
	// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
	BucketId pulumi.StringInput
	// Describes this bucket.
	Description pulumi.StringPtrInput
	// The location of the bucket.
	Location pulumi.StringInput
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used. Bucket retention can not be increased on buckets outside of projects.
	RetentionDays pulumi.IntPtrInput
}

func (BillingAccountBucketConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*billingAccountBucketConfigArgs)(nil)).Elem()
}

type BillingAccountBucketConfigInput interface {
	pulumi.Input

	ToBillingAccountBucketConfigOutput() BillingAccountBucketConfigOutput
	ToBillingAccountBucketConfigOutputWithContext(ctx context.Context) BillingAccountBucketConfigOutput
}

func (*BillingAccountBucketConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**BillingAccountBucketConfig)(nil)).Elem()
}

func (i *BillingAccountBucketConfig) ToBillingAccountBucketConfigOutput() BillingAccountBucketConfigOutput {
	return i.ToBillingAccountBucketConfigOutputWithContext(context.Background())
}

func (i *BillingAccountBucketConfig) ToBillingAccountBucketConfigOutputWithContext(ctx context.Context) BillingAccountBucketConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BillingAccountBucketConfigOutput)
}

// BillingAccountBucketConfigArrayInput is an input type that accepts BillingAccountBucketConfigArray and BillingAccountBucketConfigArrayOutput values.
// You can construct a concrete instance of `BillingAccountBucketConfigArrayInput` via:
//
//          BillingAccountBucketConfigArray{ BillingAccountBucketConfigArgs{...} }
type BillingAccountBucketConfigArrayInput interface {
	pulumi.Input

	ToBillingAccountBucketConfigArrayOutput() BillingAccountBucketConfigArrayOutput
	ToBillingAccountBucketConfigArrayOutputWithContext(context.Context) BillingAccountBucketConfigArrayOutput
}

type BillingAccountBucketConfigArray []BillingAccountBucketConfigInput

func (BillingAccountBucketConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BillingAccountBucketConfig)(nil)).Elem()
}

func (i BillingAccountBucketConfigArray) ToBillingAccountBucketConfigArrayOutput() BillingAccountBucketConfigArrayOutput {
	return i.ToBillingAccountBucketConfigArrayOutputWithContext(context.Background())
}

func (i BillingAccountBucketConfigArray) ToBillingAccountBucketConfigArrayOutputWithContext(ctx context.Context) BillingAccountBucketConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BillingAccountBucketConfigArrayOutput)
}

// BillingAccountBucketConfigMapInput is an input type that accepts BillingAccountBucketConfigMap and BillingAccountBucketConfigMapOutput values.
// You can construct a concrete instance of `BillingAccountBucketConfigMapInput` via:
//
//          BillingAccountBucketConfigMap{ "key": BillingAccountBucketConfigArgs{...} }
type BillingAccountBucketConfigMapInput interface {
	pulumi.Input

	ToBillingAccountBucketConfigMapOutput() BillingAccountBucketConfigMapOutput
	ToBillingAccountBucketConfigMapOutputWithContext(context.Context) BillingAccountBucketConfigMapOutput
}

type BillingAccountBucketConfigMap map[string]BillingAccountBucketConfigInput

func (BillingAccountBucketConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BillingAccountBucketConfig)(nil)).Elem()
}

func (i BillingAccountBucketConfigMap) ToBillingAccountBucketConfigMapOutput() BillingAccountBucketConfigMapOutput {
	return i.ToBillingAccountBucketConfigMapOutputWithContext(context.Background())
}

func (i BillingAccountBucketConfigMap) ToBillingAccountBucketConfigMapOutputWithContext(ctx context.Context) BillingAccountBucketConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BillingAccountBucketConfigMapOutput)
}

type BillingAccountBucketConfigOutput struct{ *pulumi.OutputState }

func (BillingAccountBucketConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BillingAccountBucketConfig)(nil)).Elem()
}

func (o BillingAccountBucketConfigOutput) ToBillingAccountBucketConfigOutput() BillingAccountBucketConfigOutput {
	return o
}

func (o BillingAccountBucketConfigOutput) ToBillingAccountBucketConfigOutputWithContext(ctx context.Context) BillingAccountBucketConfigOutput {
	return o
}

type BillingAccountBucketConfigArrayOutput struct{ *pulumi.OutputState }

func (BillingAccountBucketConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*BillingAccountBucketConfig)(nil)).Elem()
}

func (o BillingAccountBucketConfigArrayOutput) ToBillingAccountBucketConfigArrayOutput() BillingAccountBucketConfigArrayOutput {
	return o
}

func (o BillingAccountBucketConfigArrayOutput) ToBillingAccountBucketConfigArrayOutputWithContext(ctx context.Context) BillingAccountBucketConfigArrayOutput {
	return o
}

func (o BillingAccountBucketConfigArrayOutput) Index(i pulumi.IntInput) BillingAccountBucketConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *BillingAccountBucketConfig {
		return vs[0].([]*BillingAccountBucketConfig)[vs[1].(int)]
	}).(BillingAccountBucketConfigOutput)
}

type BillingAccountBucketConfigMapOutput struct{ *pulumi.OutputState }

func (BillingAccountBucketConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*BillingAccountBucketConfig)(nil)).Elem()
}

func (o BillingAccountBucketConfigMapOutput) ToBillingAccountBucketConfigMapOutput() BillingAccountBucketConfigMapOutput {
	return o
}

func (o BillingAccountBucketConfigMapOutput) ToBillingAccountBucketConfigMapOutputWithContext(ctx context.Context) BillingAccountBucketConfigMapOutput {
	return o
}

func (o BillingAccountBucketConfigMapOutput) MapIndex(k pulumi.StringInput) BillingAccountBucketConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *BillingAccountBucketConfig {
		return vs[0].(map[string]*BillingAccountBucketConfig)[vs[1].(string)]
	}).(BillingAccountBucketConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BillingAccountBucketConfigInput)(nil)).Elem(), &BillingAccountBucketConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*BillingAccountBucketConfigArrayInput)(nil)).Elem(), BillingAccountBucketConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BillingAccountBucketConfigMapInput)(nil)).Elem(), BillingAccountBucketConfigMap{})
	pulumi.RegisterOutputType(BillingAccountBucketConfigOutput{})
	pulumi.RegisterOutputType(BillingAccountBucketConfigArrayOutput{})
	pulumi.RegisterOutputType(BillingAccountBucketConfigMapOutput{})
}
