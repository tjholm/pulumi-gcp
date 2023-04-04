// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package logging

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a project-level logging bucket config. For more information see
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
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/logging"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := organizations.NewProject(ctx, "default", &organizations.ProjectArgs{
//				ProjectId: pulumi.String("your-project-id"),
//				OrgId:     pulumi.String("123456789"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = logging.NewProjectBucketConfig(ctx, "basic", &logging.ProjectBucketConfigArgs{
//				Project:       _default.ID(),
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
// # Create logging bucket with customId
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/logging"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := logging.NewProjectBucketConfig(ctx, "basic", &logging.ProjectBucketConfigArgs{
//				BucketId:      pulumi.String("custom-bucket"),
//				Location:      pulumi.String("global"),
//				Project:       pulumi.String("project_id"),
//				RetentionDays: pulumi.Int(30),
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
// # Create logging bucket with Log Analytics enabled
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/logging"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := logging.NewProjectBucketConfig(ctx, "analytics-enabled-bucket", &logging.ProjectBucketConfigArgs{
//				BucketId:        pulumi.String("custom-bucket"),
//				EnableAnalytics: pulumi.Bool(true),
//				Location:        pulumi.String("global"),
//				Project:         pulumi.String("project_id"),
//				RetentionDays:   pulumi.Int(30),
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
// # Create logging bucket with customId and cmekSettings
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/kms"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/logging"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cmekSettings, err := logging.GetProjectCmekSettings(ctx, &logging.GetProjectCmekSettingsArgs{
//				Project: "project_id",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			keyring, err := kms.NewKeyRing(ctx, "keyring", &kms.KeyRingArgs{
//				Location: pulumi.String("us-central1"),
//			})
//			if err != nil {
//				return err
//			}
//			key, err := kms.NewCryptoKey(ctx, "key", &kms.CryptoKeyArgs{
//				KeyRing:        keyring.ID(),
//				RotationPeriod: pulumi.String("100000s"),
//			})
//			if err != nil {
//				return err
//			}
//			cryptoKeyBinding, err := kms.NewCryptoKeyIAMBinding(ctx, "cryptoKeyBinding", &kms.CryptoKeyIAMBindingArgs{
//				CryptoKeyId: key.ID(),
//				Role:        pulumi.String("roles/cloudkms.cryptoKeyEncrypterDecrypter"),
//				Members: pulumi.StringArray{
//					pulumi.String(fmt.Sprintf("serviceAccount:%v", cmekSettings.ServiceAccountId)),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = logging.NewProjectBucketConfig(ctx, "example-project-bucket-cmek-settings", &logging.ProjectBucketConfigArgs{
//				Project:       pulumi.String("project_id"),
//				Location:      pulumi.String("us-central1"),
//				RetentionDays: pulumi.Int(30),
//				BucketId:      pulumi.String("custom-bucket"),
//				CmekSettings: &logging.ProjectBucketConfigCmekSettingsArgs{
//					KmsKeyName: key.ID(),
//				},
//			}, pulumi.DependsOn([]pulumi.Resource{
//				cryptoKeyBinding,
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
// # This resource can be imported using the following format
//
// ```sh
//
//	$ pulumi import gcp:logging/projectBucketConfig:ProjectBucketConfig default projects/{{project}}/locations/{{location}}/buckets/{{bucket_id}}
//
// ```
type ProjectBucketConfig struct {
	pulumi.CustomResourceState

	// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
	BucketId pulumi.StringOutput `pulumi:"bucketId"`
	// The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by updating the log bucket. Changing the KMS key is allowed. Structure is documented below.
	CmekSettings ProjectBucketConfigCmekSettingsPtrOutput `pulumi:"cmekSettings"`
	// Describes this bucket.
	Description pulumi.StringOutput `pulumi:"description"`
	// Whether or not Log Analytics is enabled. Logs for buckets with Log Analytics enabled can be queried in the **Log Analytics** page using SQL queries. Cannot be disabled once enabled.
	EnableAnalytics pulumi.BoolPtrOutput `pulumi:"enableAnalytics"`
	// The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
	LifecycleState pulumi.StringOutput `pulumi:"lifecycleState"`
	// The location of the bucket.
	Location pulumi.StringOutput `pulumi:"location"`
	// The resource name of the CMEK settings.
	Name pulumi.StringOutput `pulumi:"name"`
	// The parent resource that contains the logging bucket.
	Project pulumi.StringOutput `pulumi:"project"`
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays pulumi.IntPtrOutput `pulumi:"retentionDays"`
}

// NewProjectBucketConfig registers a new resource with the given unique name, arguments, and options.
func NewProjectBucketConfig(ctx *pulumi.Context,
	name string, args *ProjectBucketConfigArgs, opts ...pulumi.ResourceOption) (*ProjectBucketConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BucketId == nil {
		return nil, errors.New("invalid value for required argument 'BucketId'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource ProjectBucketConfig
	err := ctx.RegisterResource("gcp:logging/projectBucketConfig:ProjectBucketConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProjectBucketConfig gets an existing ProjectBucketConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProjectBucketConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProjectBucketConfigState, opts ...pulumi.ResourceOption) (*ProjectBucketConfig, error) {
	var resource ProjectBucketConfig
	err := ctx.ReadResource("gcp:logging/projectBucketConfig:ProjectBucketConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProjectBucketConfig resources.
type projectBucketConfigState struct {
	// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
	BucketId *string `pulumi:"bucketId"`
	// The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by updating the log bucket. Changing the KMS key is allowed. Structure is documented below.
	CmekSettings *ProjectBucketConfigCmekSettings `pulumi:"cmekSettings"`
	// Describes this bucket.
	Description *string `pulumi:"description"`
	// Whether or not Log Analytics is enabled. Logs for buckets with Log Analytics enabled can be queried in the **Log Analytics** page using SQL queries. Cannot be disabled once enabled.
	EnableAnalytics *bool `pulumi:"enableAnalytics"`
	// The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
	LifecycleState *string `pulumi:"lifecycleState"`
	// The location of the bucket.
	Location *string `pulumi:"location"`
	// The resource name of the CMEK settings.
	Name *string `pulumi:"name"`
	// The parent resource that contains the logging bucket.
	Project *string `pulumi:"project"`
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays *int `pulumi:"retentionDays"`
}

type ProjectBucketConfigState struct {
	// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
	BucketId pulumi.StringPtrInput
	// The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by updating the log bucket. Changing the KMS key is allowed. Structure is documented below.
	CmekSettings ProjectBucketConfigCmekSettingsPtrInput
	// Describes this bucket.
	Description pulumi.StringPtrInput
	// Whether or not Log Analytics is enabled. Logs for buckets with Log Analytics enabled can be queried in the **Log Analytics** page using SQL queries. Cannot be disabled once enabled.
	EnableAnalytics pulumi.BoolPtrInput
	// The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
	LifecycleState pulumi.StringPtrInput
	// The location of the bucket.
	Location pulumi.StringPtrInput
	// The resource name of the CMEK settings.
	Name pulumi.StringPtrInput
	// The parent resource that contains the logging bucket.
	Project pulumi.StringPtrInput
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays pulumi.IntPtrInput
}

func (ProjectBucketConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*projectBucketConfigState)(nil)).Elem()
}

type projectBucketConfigArgs struct {
	// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
	BucketId string `pulumi:"bucketId"`
	// The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by updating the log bucket. Changing the KMS key is allowed. Structure is documented below.
	CmekSettings *ProjectBucketConfigCmekSettings `pulumi:"cmekSettings"`
	// Describes this bucket.
	Description *string `pulumi:"description"`
	// Whether or not Log Analytics is enabled. Logs for buckets with Log Analytics enabled can be queried in the **Log Analytics** page using SQL queries. Cannot be disabled once enabled.
	EnableAnalytics *bool `pulumi:"enableAnalytics"`
	// The location of the bucket.
	Location string `pulumi:"location"`
	// The parent resource that contains the logging bucket.
	Project string `pulumi:"project"`
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays *int `pulumi:"retentionDays"`
}

// The set of arguments for constructing a ProjectBucketConfig resource.
type ProjectBucketConfigArgs struct {
	// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
	BucketId pulumi.StringInput
	// The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by updating the log bucket. Changing the KMS key is allowed. Structure is documented below.
	CmekSettings ProjectBucketConfigCmekSettingsPtrInput
	// Describes this bucket.
	Description pulumi.StringPtrInput
	// Whether or not Log Analytics is enabled. Logs for buckets with Log Analytics enabled can be queried in the **Log Analytics** page using SQL queries. Cannot be disabled once enabled.
	EnableAnalytics pulumi.BoolPtrInput
	// The location of the bucket.
	Location pulumi.StringInput
	// The parent resource that contains the logging bucket.
	Project pulumi.StringInput
	// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
	RetentionDays pulumi.IntPtrInput
}

func (ProjectBucketConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*projectBucketConfigArgs)(nil)).Elem()
}

type ProjectBucketConfigInput interface {
	pulumi.Input

	ToProjectBucketConfigOutput() ProjectBucketConfigOutput
	ToProjectBucketConfigOutputWithContext(ctx context.Context) ProjectBucketConfigOutput
}

func (*ProjectBucketConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectBucketConfig)(nil)).Elem()
}

func (i *ProjectBucketConfig) ToProjectBucketConfigOutput() ProjectBucketConfigOutput {
	return i.ToProjectBucketConfigOutputWithContext(context.Background())
}

func (i *ProjectBucketConfig) ToProjectBucketConfigOutputWithContext(ctx context.Context) ProjectBucketConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectBucketConfigOutput)
}

// ProjectBucketConfigArrayInput is an input type that accepts ProjectBucketConfigArray and ProjectBucketConfigArrayOutput values.
// You can construct a concrete instance of `ProjectBucketConfigArrayInput` via:
//
//	ProjectBucketConfigArray{ ProjectBucketConfigArgs{...} }
type ProjectBucketConfigArrayInput interface {
	pulumi.Input

	ToProjectBucketConfigArrayOutput() ProjectBucketConfigArrayOutput
	ToProjectBucketConfigArrayOutputWithContext(context.Context) ProjectBucketConfigArrayOutput
}

type ProjectBucketConfigArray []ProjectBucketConfigInput

func (ProjectBucketConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectBucketConfig)(nil)).Elem()
}

func (i ProjectBucketConfigArray) ToProjectBucketConfigArrayOutput() ProjectBucketConfigArrayOutput {
	return i.ToProjectBucketConfigArrayOutputWithContext(context.Background())
}

func (i ProjectBucketConfigArray) ToProjectBucketConfigArrayOutputWithContext(ctx context.Context) ProjectBucketConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectBucketConfigArrayOutput)
}

// ProjectBucketConfigMapInput is an input type that accepts ProjectBucketConfigMap and ProjectBucketConfigMapOutput values.
// You can construct a concrete instance of `ProjectBucketConfigMapInput` via:
//
//	ProjectBucketConfigMap{ "key": ProjectBucketConfigArgs{...} }
type ProjectBucketConfigMapInput interface {
	pulumi.Input

	ToProjectBucketConfigMapOutput() ProjectBucketConfigMapOutput
	ToProjectBucketConfigMapOutputWithContext(context.Context) ProjectBucketConfigMapOutput
}

type ProjectBucketConfigMap map[string]ProjectBucketConfigInput

func (ProjectBucketConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectBucketConfig)(nil)).Elem()
}

func (i ProjectBucketConfigMap) ToProjectBucketConfigMapOutput() ProjectBucketConfigMapOutput {
	return i.ToProjectBucketConfigMapOutputWithContext(context.Background())
}

func (i ProjectBucketConfigMap) ToProjectBucketConfigMapOutputWithContext(ctx context.Context) ProjectBucketConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProjectBucketConfigMapOutput)
}

type ProjectBucketConfigOutput struct{ *pulumi.OutputState }

func (ProjectBucketConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProjectBucketConfig)(nil)).Elem()
}

func (o ProjectBucketConfigOutput) ToProjectBucketConfigOutput() ProjectBucketConfigOutput {
	return o
}

func (o ProjectBucketConfigOutput) ToProjectBucketConfigOutputWithContext(ctx context.Context) ProjectBucketConfigOutput {
	return o
}

// The name of the logging bucket. Logging automatically creates two log buckets: `_Required` and `_Default`.
func (o ProjectBucketConfigOutput) BucketId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectBucketConfig) pulumi.StringOutput { return v.BucketId }).(pulumi.StringOutput)
}

// The CMEK settings of the log bucket. If present, new log entries written to this log bucket are encrypted using the CMEK key provided in this configuration. If a log bucket has CMEK settings, the CMEK settings cannot be disabled later by updating the log bucket. Changing the KMS key is allowed. Structure is documented below.
func (o ProjectBucketConfigOutput) CmekSettings() ProjectBucketConfigCmekSettingsPtrOutput {
	return o.ApplyT(func(v *ProjectBucketConfig) ProjectBucketConfigCmekSettingsPtrOutput { return v.CmekSettings }).(ProjectBucketConfigCmekSettingsPtrOutput)
}

// Describes this bucket.
func (o ProjectBucketConfigOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectBucketConfig) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Whether or not Log Analytics is enabled. Logs for buckets with Log Analytics enabled can be queried in the **Log Analytics** page using SQL queries. Cannot be disabled once enabled.
func (o ProjectBucketConfigOutput) EnableAnalytics() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ProjectBucketConfig) pulumi.BoolPtrOutput { return v.EnableAnalytics }).(pulumi.BoolPtrOutput)
}

// The bucket's lifecycle such as active or deleted. See [LifecycleState](https://cloud.google.com/logging/docs/reference/v2/rest/v2/billingAccounts.buckets#LogBucket.LifecycleState).
func (o ProjectBucketConfigOutput) LifecycleState() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectBucketConfig) pulumi.StringOutput { return v.LifecycleState }).(pulumi.StringOutput)
}

// The location of the bucket.
func (o ProjectBucketConfigOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectBucketConfig) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The resource name of the CMEK settings.
func (o ProjectBucketConfigOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectBucketConfig) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The parent resource that contains the logging bucket.
func (o ProjectBucketConfigOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *ProjectBucketConfig) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Logs will be retained by default for this amount of time, after which they will automatically be deleted. The minimum retention period is 1 day. If this value is set to zero at bucket creation time, the default time of 30 days will be used.
func (o ProjectBucketConfigOutput) RetentionDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ProjectBucketConfig) pulumi.IntPtrOutput { return v.RetentionDays }).(pulumi.IntPtrOutput)
}

type ProjectBucketConfigArrayOutput struct{ *pulumi.OutputState }

func (ProjectBucketConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProjectBucketConfig)(nil)).Elem()
}

func (o ProjectBucketConfigArrayOutput) ToProjectBucketConfigArrayOutput() ProjectBucketConfigArrayOutput {
	return o
}

func (o ProjectBucketConfigArrayOutput) ToProjectBucketConfigArrayOutputWithContext(ctx context.Context) ProjectBucketConfigArrayOutput {
	return o
}

func (o ProjectBucketConfigArrayOutput) Index(i pulumi.IntInput) ProjectBucketConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProjectBucketConfig {
		return vs[0].([]*ProjectBucketConfig)[vs[1].(int)]
	}).(ProjectBucketConfigOutput)
}

type ProjectBucketConfigMapOutput struct{ *pulumi.OutputState }

func (ProjectBucketConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProjectBucketConfig)(nil)).Elem()
}

func (o ProjectBucketConfigMapOutput) ToProjectBucketConfigMapOutput() ProjectBucketConfigMapOutput {
	return o
}

func (o ProjectBucketConfigMapOutput) ToProjectBucketConfigMapOutputWithContext(ctx context.Context) ProjectBucketConfigMapOutput {
	return o
}

func (o ProjectBucketConfigMapOutput) MapIndex(k pulumi.StringInput) ProjectBucketConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProjectBucketConfig {
		return vs[0].(map[string]*ProjectBucketConfig)[vs[1].(string)]
	}).(ProjectBucketConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectBucketConfigInput)(nil)).Elem(), &ProjectBucketConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectBucketConfigArrayInput)(nil)).Elem(), ProjectBucketConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProjectBucketConfigMapInput)(nil)).Elem(), ProjectBucketConfigMap{})
	pulumi.RegisterOutputType(ProjectBucketConfigOutput{})
	pulumi.RegisterOutputType(ProjectBucketConfigArrayOutput{})
	pulumi.RegisterOutputType(ProjectBucketConfigMapOutput{})
}
