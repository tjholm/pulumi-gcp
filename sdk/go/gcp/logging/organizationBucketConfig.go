// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package logging

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/logging"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_default, err := organizations.GetOrganization(ctx, &organizations.GetOrganizationArgs{
//				Organization: pulumi.StringRef("123456789"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = logging.NewOrganizationBucketConfig(ctx, "basic", &logging.OrganizationBucketConfigArgs{
//				Organization:  *pulumi.String(_default.Organization),
//				Location:      pulumi.String("global"),
//				RetentionDays: pulumi.Int(30),
//				BucketId:      pulumi.String("_Default"),
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
// # This resource can be imported using the following format
//
// ```sh
//
//	$ pulumi import gcp:logging/organizationBucketConfig:OrganizationBucketConfig default organizations/{{organization}}/locations/{{location}}/buckets/{{bucket_id}}
//
// ```
type OrganizationBucketConfig struct {
	pulumi.CustomResourceState

	// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
	BucketId pulumi.StringOutput `pulumi:"bucketId"`
	// The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK
	// key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by
	// updating the log bucket. Changing the KMS key is allowed.
	CmekSettings OrganizationBucketConfigCmekSettingsPtrOutput `pulumi:"cmekSettings"`
	// Describes this bucket.
	Description pulumi.StringOutput `pulumi:"description"`
	// The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
	LifecycleState pulumi.StringOutput `pulumi:"lifecycleState"`
	// The location of the bucket. The supported locations are: "global" "us-central1"
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name of the bucket. For example: "organizations/my-organization-id/locations/my-location/buckets/my-bucket-id"
	Name pulumi.StringOutput `pulumi:"name"`
	// The parent resource that contains the logging bucket.
	Organization pulumi.StringOutput `pulumi:"organization"`
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used. Bucket retention can not be increased on buckets outside of projects.
	RetentionDays pulumi.IntPtrOutput `pulumi:"retentionDays"`
}

// NewOrganizationBucketConfig registers a new resource with the given unique name, arguments, and options.
func NewOrganizationBucketConfig(ctx *pulumi.Context,
	name string, args *OrganizationBucketConfigArgs, opts ...pulumi.ResourceOption) (*OrganizationBucketConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BucketId == nil {
		return nil, errors.New("invalid value for required argument 'BucketId'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Organization == nil {
		return nil, errors.New("invalid value for required argument 'Organization'")
	}
	var resource OrganizationBucketConfig
	err := ctx.RegisterResource("gcp:logging/organizationBucketConfig:OrganizationBucketConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationBucketConfig gets an existing OrganizationBucketConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationBucketConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationBucketConfigState, opts ...pulumi.ResourceOption) (*OrganizationBucketConfig, error) {
	var resource OrganizationBucketConfig
	err := ctx.ReadResource("gcp:logging/organizationBucketConfig:OrganizationBucketConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationBucketConfig resources.
type organizationBucketConfigState struct {
	// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
	BucketId *string `pulumi:"bucketId"`
	// The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK
	// key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by
	// updating the log bucket. Changing the KMS key is allowed.
	CmekSettings *OrganizationBucketConfigCmekSettings `pulumi:"cmekSettings"`
	// Describes this bucket.
	Description *string `pulumi:"description"`
	// The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
	LifecycleState *string `pulumi:"lifecycleState"`
	// The location of the bucket. The supported locations are: "global" "us-central1"
	Location *string `pulumi:"location"`
	// The resource name of the bucket. For example: "organizations/my-organization-id/locations/my-location/buckets/my-bucket-id"
	Name *string `pulumi:"name"`
	// The parent resource that contains the logging bucket.
	Organization *string `pulumi:"organization"`
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used. Bucket retention can not be increased on buckets outside of projects.
	RetentionDays *int `pulumi:"retentionDays"`
}

type OrganizationBucketConfigState struct {
	// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
	BucketId pulumi.StringPtrInput
	// The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK
	// key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by
	// updating the log bucket. Changing the KMS key is allowed.
	CmekSettings OrganizationBucketConfigCmekSettingsPtrInput
	// Describes this bucket.
	Description pulumi.StringPtrInput
	// The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
	LifecycleState pulumi.StringPtrInput
	// The location of the bucket. The supported locations are: "global" "us-central1"
	Location pulumi.StringPtrInput
	// The resource name of the bucket. For example: "organizations/my-organization-id/locations/my-location/buckets/my-bucket-id"
	Name pulumi.StringPtrInput
	// The parent resource that contains the logging bucket.
	Organization pulumi.StringPtrInput
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used. Bucket retention can not be increased on buckets outside of projects.
	RetentionDays pulumi.IntPtrInput
}

func (OrganizationBucketConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationBucketConfigState)(nil)).Elem()
}

type organizationBucketConfigArgs struct {
	// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
	BucketId string `pulumi:"bucketId"`
	// The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK
	// key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by
	// updating the log bucket. Changing the KMS key is allowed.
	CmekSettings *OrganizationBucketConfigCmekSettings `pulumi:"cmekSettings"`
	// Describes this bucket.
	Description *string `pulumi:"description"`
	// The location of the bucket. The supported locations are: "global" "us-central1"
	Location string `pulumi:"location"`
	// The parent resource that contains the logging bucket.
	Organization string `pulumi:"organization"`
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used. Bucket retention can not be increased on buckets outside of projects.
	RetentionDays *int `pulumi:"retentionDays"`
}

// The set of arguments for constructing a OrganizationBucketConfig resource.
type OrganizationBucketConfigArgs struct {
	// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
	BucketId pulumi.StringInput
	// The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK
	// key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by
	// updating the log bucket. Changing the KMS key is allowed.
	CmekSettings OrganizationBucketConfigCmekSettingsPtrInput
	// Describes this bucket.
	Description pulumi.StringPtrInput
	// The location of the bucket. The supported locations are: "global" "us-central1"
	Location pulumi.StringInput
	// The parent resource that contains the logging bucket.
	Organization pulumi.StringInput
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used. Bucket retention can not be increased on buckets outside of projects.
	RetentionDays pulumi.IntPtrInput
}

func (OrganizationBucketConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationBucketConfigArgs)(nil)).Elem()
}

type OrganizationBucketConfigInput interface {
	pulumi.Input

	ToOrganizationBucketConfigOutput() OrganizationBucketConfigOutput
	ToOrganizationBucketConfigOutputWithContext(ctx context.Context) OrganizationBucketConfigOutput
}

func (*OrganizationBucketConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationBucketConfig)(nil)).Elem()
}

func (i *OrganizationBucketConfig) ToOrganizationBucketConfigOutput() OrganizationBucketConfigOutput {
	return i.ToOrganizationBucketConfigOutputWithContext(context.Background())
}

func (i *OrganizationBucketConfig) ToOrganizationBucketConfigOutputWithContext(ctx context.Context) OrganizationBucketConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationBucketConfigOutput)
}

// OrganizationBucketConfigArrayInput is an input type that accepts OrganizationBucketConfigArray and OrganizationBucketConfigArrayOutput values.
// You can construct a concrete instance of `OrganizationBucketConfigArrayInput` via:
//
//	OrganizationBucketConfigArray{ OrganizationBucketConfigArgs{...} }
type OrganizationBucketConfigArrayInput interface {
	pulumi.Input

	ToOrganizationBucketConfigArrayOutput() OrganizationBucketConfigArrayOutput
	ToOrganizationBucketConfigArrayOutputWithContext(context.Context) OrganizationBucketConfigArrayOutput
}

type OrganizationBucketConfigArray []OrganizationBucketConfigInput

func (OrganizationBucketConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrganizationBucketConfig)(nil)).Elem()
}

func (i OrganizationBucketConfigArray) ToOrganizationBucketConfigArrayOutput() OrganizationBucketConfigArrayOutput {
	return i.ToOrganizationBucketConfigArrayOutputWithContext(context.Background())
}

func (i OrganizationBucketConfigArray) ToOrganizationBucketConfigArrayOutputWithContext(ctx context.Context) OrganizationBucketConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationBucketConfigArrayOutput)
}

// OrganizationBucketConfigMapInput is an input type that accepts OrganizationBucketConfigMap and OrganizationBucketConfigMapOutput values.
// You can construct a concrete instance of `OrganizationBucketConfigMapInput` via:
//
//	OrganizationBucketConfigMap{ "key": OrganizationBucketConfigArgs{...} }
type OrganizationBucketConfigMapInput interface {
	pulumi.Input

	ToOrganizationBucketConfigMapOutput() OrganizationBucketConfigMapOutput
	ToOrganizationBucketConfigMapOutputWithContext(context.Context) OrganizationBucketConfigMapOutput
}

type OrganizationBucketConfigMap map[string]OrganizationBucketConfigInput

func (OrganizationBucketConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrganizationBucketConfig)(nil)).Elem()
}

func (i OrganizationBucketConfigMap) ToOrganizationBucketConfigMapOutput() OrganizationBucketConfigMapOutput {
	return i.ToOrganizationBucketConfigMapOutputWithContext(context.Background())
}

func (i OrganizationBucketConfigMap) ToOrganizationBucketConfigMapOutputWithContext(ctx context.Context) OrganizationBucketConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationBucketConfigMapOutput)
}

type OrganizationBucketConfigOutput struct{ *pulumi.OutputState }

func (OrganizationBucketConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrganizationBucketConfig)(nil)).Elem()
}

func (o OrganizationBucketConfigOutput) ToOrganizationBucketConfigOutput() OrganizationBucketConfigOutput {
	return o
}

func (o OrganizationBucketConfigOutput) ToOrganizationBucketConfigOutputWithContext(ctx context.Context) OrganizationBucketConfigOutput {
	return o
}

// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
func (o OrganizationBucketConfigOutput) BucketId() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationBucketConfig) pulumi.StringOutput { return v.BucketId }).(pulumi.StringOutput)
}

// The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK
// key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by
// updating the log bucket. Changing the KMS key is allowed.
func (o OrganizationBucketConfigOutput) CmekSettings() OrganizationBucketConfigCmekSettingsPtrOutput {
	return o.ApplyT(func(v *OrganizationBucketConfig) OrganizationBucketConfigCmekSettingsPtrOutput { return v.CmekSettings }).(OrganizationBucketConfigCmekSettingsPtrOutput)
}

// Describes this bucket.
func (o OrganizationBucketConfigOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationBucketConfig) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
func (o OrganizationBucketConfigOutput) LifecycleState() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationBucketConfig) pulumi.StringOutput { return v.LifecycleState }).(pulumi.StringOutput)
}

// The location of the bucket. The supported locations are: "global" "us-central1"
func (o OrganizationBucketConfigOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationBucketConfig) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The resource name of the bucket. For example: "organizations/my-organization-id/locations/my-location/buckets/my-bucket-id"
func (o OrganizationBucketConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationBucketConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The parent resource that contains the logging bucket.
func (o OrganizationBucketConfigOutput) Organization() pulumi.StringOutput {
	return o.ApplyT(func(v *OrganizationBucketConfig) pulumi.StringOutput { return v.Organization }).(pulumi.StringOutput)
}

// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used. Bucket retention can not be increased on buckets outside of projects.
func (o OrganizationBucketConfigOutput) RetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *OrganizationBucketConfig) pulumi.IntPtrOutput { return v.RetentionDays }).(pulumi.IntPtrOutput)
}

type OrganizationBucketConfigArrayOutput struct{ *pulumi.OutputState }

func (OrganizationBucketConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrganizationBucketConfig)(nil)).Elem()
}

func (o OrganizationBucketConfigArrayOutput) ToOrganizationBucketConfigArrayOutput() OrganizationBucketConfigArrayOutput {
	return o
}

func (o OrganizationBucketConfigArrayOutput) ToOrganizationBucketConfigArrayOutputWithContext(ctx context.Context) OrganizationBucketConfigArrayOutput {
	return o
}

func (o OrganizationBucketConfigArrayOutput) Index(i pulumi.IntInput) OrganizationBucketConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OrganizationBucketConfig {
		return vs[0].([]*OrganizationBucketConfig)[vs[1].(int)]
	}).(OrganizationBucketConfigOutput)
}

type OrganizationBucketConfigMapOutput struct{ *pulumi.OutputState }

func (OrganizationBucketConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrganizationBucketConfig)(nil)).Elem()
}

func (o OrganizationBucketConfigMapOutput) ToOrganizationBucketConfigMapOutput() OrganizationBucketConfigMapOutput {
	return o
}

func (o OrganizationBucketConfigMapOutput) ToOrganizationBucketConfigMapOutputWithContext(ctx context.Context) OrganizationBucketConfigMapOutput {
	return o
}

func (o OrganizationBucketConfigMapOutput) MapIndex(k pulumi.StringInput) OrganizationBucketConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OrganizationBucketConfig {
		return vs[0].(map[string]*OrganizationBucketConfig)[vs[1].(string)]
	}).(OrganizationBucketConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationBucketConfigInput)(nil)).Elem(), &OrganizationBucketConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationBucketConfigArrayInput)(nil)).Elem(), OrganizationBucketConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrganizationBucketConfigMapInput)(nil)).Elem(), OrganizationBucketConfigMap{})
	pulumi.RegisterOutputType(OrganizationBucketConfigOutput{})
	pulumi.RegisterOutputType(OrganizationBucketConfigArrayOutput{})
	pulumi.RegisterOutputType(OrganizationBucketConfigMapOutput{})
}
