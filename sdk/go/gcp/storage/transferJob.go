// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new Transfer Job in Google Cloud Storage Transfer.
//
// To get more information about Google Cloud Storage Transfer, see:
//
// * [Overview](https://cloud.google.com/storage-transfer/docs/overview)
// * [API documentation](https://cloud.google.com/storage-transfer/docs/reference/rest/v1/transferJobs)
// * How-to Guides
//   - [Configuring Access to Data Sources and Sinks](https://cloud.google.com/storage-transfer/docs/configure-access)
//
// ## Example Usage
//
// Example creating a nightly Transfer Job from an AWS S3 Bucket to a GCS bucket.
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/storage"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_default, err := storage.GetTransferProjectServieAccount(ctx, &storage.GetTransferProjectServieAccountArgs{
//				Project: pulumi.StringRef(_var.Project),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = storage.NewBucket(ctx, "s3-backup-bucketBucket", &storage.BucketArgs{
//				StorageClass: pulumi.String("NEARLINE"),
//				Project:      pulumi.Any(_var.Project),
//				Location:     pulumi.String("US"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = storage.NewBucketIAMMember(ctx, "s3-backup-bucketBucketIAMMember", &storage.BucketIAMMemberArgs{
//				Bucket: s3_backup_bucketBucket.Name,
//				Role:   pulumi.String("roles/storage.admin"),
//				Member: pulumi.String(fmt.Sprintf("serviceAccount:%v", _default.Email)),
//			}, pulumi.DependsOn([]pulumi.Resource{
//				s3_backup_bucketBucket,
//			}))
//			if err != nil {
//				return err
//			}
//			_, err = storage.NewTransferJob(ctx, "s3-bucket-nightly-backup", &storage.TransferJobArgs{
//				Description: pulumi.String("Nightly backup of S3 bucket"),
//				Project:     pulumi.Any(_var.Project),
//				TransferSpec: &storage.TransferJobTransferSpecArgs{
//					ObjectConditions: &storage.TransferJobTransferSpecObjectConditionsArgs{
//						MaxTimeElapsedSinceLastModification: pulumi.String("600s"),
//						ExcludePrefixes: pulumi.StringArray{
//							pulumi.String("requests.gz"),
//						},
//					},
//					TransferOptions: &storage.TransferJobTransferSpecTransferOptionsArgs{
//						DeleteObjectsUniqueInSink: pulumi.Bool(false),
//					},
//					AwsS3DataSource: &storage.TransferJobTransferSpecAwsS3DataSourceArgs{
//						BucketName: pulumi.Any(_var.Aws_s3_bucket),
//						AwsAccessKey: &storage.TransferJobTransferSpecAwsS3DataSourceAwsAccessKeyArgs{
//							AccessKeyId:     pulumi.Any(_var.Aws_access_key),
//							SecretAccessKey: pulumi.Any(_var.Aws_secret_key),
//						},
//					},
//					GcsDataSink: &storage.TransferJobTransferSpecGcsDataSinkArgs{
//						BucketName: s3_backup_bucketBucket.Name,
//						Path:       pulumi.String("foo/bar/"),
//					},
//				},
//				Schedule: &storage.TransferJobScheduleArgs{
//					ScheduleStartDate: &storage.TransferJobScheduleScheduleStartDateArgs{
//						Year:  pulumi.Int(2018),
//						Month: pulumi.Int(10),
//						Day:   pulumi.Int(1),
//					},
//					ScheduleEndDate: &storage.TransferJobScheduleScheduleEndDateArgs{
//						Year:  pulumi.Int(2019),
//						Month: pulumi.Int(1),
//						Day:   pulumi.Int(15),
//					},
//					StartTimeOfDay: &storage.TransferJobScheduleStartTimeOfDayArgs{
//						Hours:   pulumi.Int(23),
//						Minutes: pulumi.Int(30),
//						Seconds: pulumi.Int(0),
//						Nanos:   pulumi.Int(0),
//					},
//					RepeatInterval: pulumi.String("604800s"),
//				},
//			}, pulumi.DependsOn([]pulumi.Resource{
//				s3_backup_bucketBucketIAMMember,
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
// Storage buckets can be imported using the Transfer Job's `project` and `name` without the `transferJob/` prefix, e.g.
//
// ```sh
//
//	$ pulumi import gcp:storage/transferJob:TransferJob nightly-backup-transfer-job my-project-1asd32/8422144862922355674
//
// ```
type TransferJob struct {
	pulumi.CustomResourceState

	// When the Transfer Job was created.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	// When the Transfer Job was deleted.
	DeletionTime pulumi.StringOutput `pulumi:"deletionTime"`
	// Unique description to identify the Transfer Job.
	Description pulumi.StringOutput `pulumi:"description"`
	// When the Transfer Job was last modified.
	LastModificationTime pulumi.StringOutput `pulumi:"lastModificationTime"`
	// The name of the Transfer Job.
	Name pulumi.StringOutput `pulumi:"name"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Schedule specification defining when the Transfer Job should be scheduled to start, end and what time to run. Structure documented below.
	Schedule TransferJobSchedulePtrOutput `pulumi:"schedule"`
	// Status of the job. Default: `ENABLED`. **NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.**
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// Transfer specification. Structure documented below.
	TransferSpec TransferJobTransferSpecOutput `pulumi:"transferSpec"`
}

// NewTransferJob registers a new resource with the given unique name, arguments, and options.
func NewTransferJob(ctx *pulumi.Context,
	name string, args *TransferJobArgs, opts ...pulumi.ResourceOption) (*TransferJob, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.TransferSpec == nil {
		return nil, errors.New("invalid value for required argument 'TransferSpec'")
	}
	var resource TransferJob
	err := ctx.RegisterResource("gcp:storage/transferJob:TransferJob", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTransferJob gets an existing TransferJob resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTransferJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TransferJobState, opts ...pulumi.ResourceOption) (*TransferJob, error) {
	var resource TransferJob
	err := ctx.ReadResource("gcp:storage/transferJob:TransferJob", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TransferJob resources.
type transferJobState struct {
	// When the Transfer Job was created.
	CreationTime *string `pulumi:"creationTime"`
	// When the Transfer Job was deleted.
	DeletionTime *string `pulumi:"deletionTime"`
	// Unique description to identify the Transfer Job.
	Description *string `pulumi:"description"`
	// When the Transfer Job was last modified.
	LastModificationTime *string `pulumi:"lastModificationTime"`
	// The name of the Transfer Job.
	Name *string `pulumi:"name"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Schedule specification defining when the Transfer Job should be scheduled to start, end and what time to run. Structure documented below.
	Schedule *TransferJobSchedule `pulumi:"schedule"`
	// Status of the job. Default: `ENABLED`. **NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.**
	Status *string `pulumi:"status"`
	// Transfer specification. Structure documented below.
	TransferSpec *TransferJobTransferSpec `pulumi:"transferSpec"`
}

type TransferJobState struct {
	// When the Transfer Job was created.
	CreationTime pulumi.StringPtrInput
	// When the Transfer Job was deleted.
	DeletionTime pulumi.StringPtrInput
	// Unique description to identify the Transfer Job.
	Description pulumi.StringPtrInput
	// When the Transfer Job was last modified.
	LastModificationTime pulumi.StringPtrInput
	// The name of the Transfer Job.
	Name pulumi.StringPtrInput
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Schedule specification defining when the Transfer Job should be scheduled to start, end and what time to run. Structure documented below.
	Schedule TransferJobSchedulePtrInput
	// Status of the job. Default: `ENABLED`. **NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.**
	Status pulumi.StringPtrInput
	// Transfer specification. Structure documented below.
	TransferSpec TransferJobTransferSpecPtrInput
}

func (TransferJobState) ElementType() reflect.Type {
	return reflect.TypeOf((*transferJobState)(nil)).Elem()
}

type transferJobArgs struct {
	// Unique description to identify the Transfer Job.
	Description string `pulumi:"description"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Schedule specification defining when the Transfer Job should be scheduled to start, end and what time to run. Structure documented below.
	Schedule *TransferJobSchedule `pulumi:"schedule"`
	// Status of the job. Default: `ENABLED`. **NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.**
	Status *string `pulumi:"status"`
	// Transfer specification. Structure documented below.
	TransferSpec TransferJobTransferSpec `pulumi:"transferSpec"`
}

// The set of arguments for constructing a TransferJob resource.
type TransferJobArgs struct {
	// Unique description to identify the Transfer Job.
	Description pulumi.StringInput
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Schedule specification defining when the Transfer Job should be scheduled to start, end and what time to run. Structure documented below.
	Schedule TransferJobSchedulePtrInput
	// Status of the job. Default: `ENABLED`. **NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.**
	Status pulumi.StringPtrInput
	// Transfer specification. Structure documented below.
	TransferSpec TransferJobTransferSpecInput
}

func (TransferJobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*transferJobArgs)(nil)).Elem()
}

type TransferJobInput interface {
	pulumi.Input

	ToTransferJobOutput() TransferJobOutput
	ToTransferJobOutputWithContext(ctx context.Context) TransferJobOutput
}

func (*TransferJob) ElementType() reflect.Type {
	return reflect.TypeOf((**TransferJob)(nil)).Elem()
}

func (i *TransferJob) ToTransferJobOutput() TransferJobOutput {
	return i.ToTransferJobOutputWithContext(context.Background())
}

func (i *TransferJob) ToTransferJobOutputWithContext(ctx context.Context) TransferJobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransferJobOutput)
}

// TransferJobArrayInput is an input type that accepts TransferJobArray and TransferJobArrayOutput values.
// You can construct a concrete instance of `TransferJobArrayInput` via:
//
//	TransferJobArray{ TransferJobArgs{...} }
type TransferJobArrayInput interface {
	pulumi.Input

	ToTransferJobArrayOutput() TransferJobArrayOutput
	ToTransferJobArrayOutputWithContext(context.Context) TransferJobArrayOutput
}

type TransferJobArray []TransferJobInput

func (TransferJobArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TransferJob)(nil)).Elem()
}

func (i TransferJobArray) ToTransferJobArrayOutput() TransferJobArrayOutput {
	return i.ToTransferJobArrayOutputWithContext(context.Background())
}

func (i TransferJobArray) ToTransferJobArrayOutputWithContext(ctx context.Context) TransferJobArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransferJobArrayOutput)
}

// TransferJobMapInput is an input type that accepts TransferJobMap and TransferJobMapOutput values.
// You can construct a concrete instance of `TransferJobMapInput` via:
//
//	TransferJobMap{ "key": TransferJobArgs{...} }
type TransferJobMapInput interface {
	pulumi.Input

	ToTransferJobMapOutput() TransferJobMapOutput
	ToTransferJobMapOutputWithContext(context.Context) TransferJobMapOutput
}

type TransferJobMap map[string]TransferJobInput

func (TransferJobMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TransferJob)(nil)).Elem()
}

func (i TransferJobMap) ToTransferJobMapOutput() TransferJobMapOutput {
	return i.ToTransferJobMapOutputWithContext(context.Background())
}

func (i TransferJobMap) ToTransferJobMapOutputWithContext(ctx context.Context) TransferJobMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransferJobMapOutput)
}

type TransferJobOutput struct{ *pulumi.OutputState }

func (TransferJobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TransferJob)(nil)).Elem()
}

func (o TransferJobOutput) ToTransferJobOutput() TransferJobOutput {
	return o
}

func (o TransferJobOutput) ToTransferJobOutputWithContext(ctx context.Context) TransferJobOutput {
	return o
}

// When the Transfer Job was created.
func (o TransferJobOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *TransferJob) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

// When the Transfer Job was deleted.
func (o TransferJobOutput) DeletionTime() pulumi.StringOutput {
	return o.ApplyT(func(v *TransferJob) pulumi.StringOutput { return v.DeletionTime }).(pulumi.StringOutput)
}

// Unique description to identify the Transfer Job.
func (o TransferJobOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *TransferJob) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// When the Transfer Job was last modified.
func (o TransferJobOutput) LastModificationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *TransferJob) pulumi.StringOutput { return v.LastModificationTime }).(pulumi.StringOutput)
}

// The name of the Transfer Job.
func (o TransferJobOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *TransferJob) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The project in which the resource belongs. If it
// is not provided, the provider project is used.
func (o TransferJobOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *TransferJob) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Schedule specification defining when the Transfer Job should be scheduled to start, end and what time to run. Structure documented below.
func (o TransferJobOutput) Schedule() TransferJobSchedulePtrOutput {
	return o.ApplyT(func(v *TransferJob) TransferJobSchedulePtrOutput { return v.Schedule }).(TransferJobSchedulePtrOutput)
}

// Status of the job. Default: `ENABLED`. **NOTE: The effect of the new job status takes place during a subsequent job run. For example, if you change the job status from ENABLED to DISABLED, and an operation spawned by the transfer is running, the status change would not affect the current operation.**
func (o TransferJobOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TransferJob) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

// Transfer specification. Structure documented below.
func (o TransferJobOutput) TransferSpec() TransferJobTransferSpecOutput {
	return o.ApplyT(func(v *TransferJob) TransferJobTransferSpecOutput { return v.TransferSpec }).(TransferJobTransferSpecOutput)
}

type TransferJobArrayOutput struct{ *pulumi.OutputState }

func (TransferJobArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TransferJob)(nil)).Elem()
}

func (o TransferJobArrayOutput) ToTransferJobArrayOutput() TransferJobArrayOutput {
	return o
}

func (o TransferJobArrayOutput) ToTransferJobArrayOutputWithContext(ctx context.Context) TransferJobArrayOutput {
	return o
}

func (o TransferJobArrayOutput) Index(i pulumi.IntInput) TransferJobOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TransferJob {
		return vs[0].([]*TransferJob)[vs[1].(int)]
	}).(TransferJobOutput)
}

type TransferJobMapOutput struct{ *pulumi.OutputState }

func (TransferJobMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TransferJob)(nil)).Elem()
}

func (o TransferJobMapOutput) ToTransferJobMapOutput() TransferJobMapOutput {
	return o
}

func (o TransferJobMapOutput) ToTransferJobMapOutputWithContext(ctx context.Context) TransferJobMapOutput {
	return o
}

func (o TransferJobMapOutput) MapIndex(k pulumi.StringInput) TransferJobOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TransferJob {
		return vs[0].(map[string]*TransferJob)[vs[1].(string)]
	}).(TransferJobOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TransferJobInput)(nil)).Elem(), &TransferJob{})
	pulumi.RegisterInputType(reflect.TypeOf((*TransferJobArrayInput)(nil)).Elem(), TransferJobArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TransferJobMapInput)(nil)).Elem(), TransferJobMap{})
	pulumi.RegisterOutputType(TransferJobOutput{})
	pulumi.RegisterOutputType(TransferJobArrayOutput{})
	pulumi.RegisterOutputType(TransferJobMapOutput{})
}
