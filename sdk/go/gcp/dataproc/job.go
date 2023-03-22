// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dataproc

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages a job resource within a Dataproc cluster within GCE. For more information see
// [the official dataproc documentation](https://cloud.google.com/dataproc/).
//
// !> **Note:** This resource does not support 'update' and changing any attributes will cause the resource to be recreated.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/dataproc"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			mycluster, err := dataproc.NewCluster(ctx, "mycluster", &dataproc.ClusterArgs{
//				Region: pulumi.String("us-central1"),
//			})
//			if err != nil {
//				return err
//			}
//			spark, err := dataproc.NewJob(ctx, "spark", &dataproc.JobArgs{
//				Region:      mycluster.Region,
//				ForceDelete: pulumi.Bool(true),
//				Placement: &dataproc.JobPlacementArgs{
//					ClusterName: mycluster.Name,
//				},
//				SparkConfig: &dataproc.JobSparkConfigArgs{
//					MainClass: pulumi.String("org.apache.spark.examples.SparkPi"),
//					JarFileUris: pulumi.StringArray{
//						pulumi.String("file:///usr/lib/spark/examples/jars/spark-examples.jar"),
//					},
//					Args: pulumi.StringArray{
//						pulumi.String("1000"),
//					},
//					Properties: pulumi.StringMap{
//						"spark.logConf": pulumi.String("true"),
//					},
//					LoggingConfig: &dataproc.JobSparkConfigLoggingConfigArgs{
//						DriverLogLevels: pulumi.StringMap{
//							"root": pulumi.String("INFO"),
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			pyspark, err := dataproc.NewJob(ctx, "pyspark", &dataproc.JobArgs{
//				Region:      mycluster.Region,
//				ForceDelete: pulumi.Bool(true),
//				Placement: &dataproc.JobPlacementArgs{
//					ClusterName: mycluster.Name,
//				},
//				PysparkConfig: &dataproc.JobPysparkConfigArgs{
//					MainPythonFileUri: pulumi.String("gs://dataproc-examples-2f10d78d114f6aaec76462e3c310f31f/src/pyspark/hello-world/hello-world.py"),
//					Properties: pulumi.StringMap{
//						"spark.logConf": pulumi.String("true"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			ctx.Export("sparkStatus", spark.Statuses.ApplyT(func(statuses []dataproc.JobStatus) (*string, error) {
//				return &statuses[0].State, nil
//			}).(pulumi.StringPtrOutput))
//			ctx.Export("pysparkStatus", pyspark.Statuses.ApplyT(func(statuses []dataproc.JobStatus) (*string, error) {
//				return &statuses[0].State, nil
//			}).(pulumi.StringPtrOutput))
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// This resource does not support import.
type Job struct {
	pulumi.CustomResourceState

	// If present, the location of miscellaneous control files which may be used as part of job setup and handling. If not present, control files may be placed in the same location as driver_output_uri.
	DriverControlsFilesUri pulumi.StringOutput `pulumi:"driverControlsFilesUri"`
	// A URI pointing to the location of the stdout of the job's driver program.
	DriverOutputResourceUri pulumi.StringOutput `pulumi:"driverOutputResourceUri"`
	// By default, you can only delete inactive jobs within
	// Dataproc. Setting this to true, and calling destroy, will ensure that the
	// job is first cancelled before issuing the delete.
	ForceDelete pulumi.BoolPtrOutput `pulumi:"forceDelete"`
	// The config of Hadoop job
	HadoopConfig JobHadoopConfigPtrOutput `pulumi:"hadoopConfig"`
	// The config of hive job
	HiveConfig JobHiveConfigPtrOutput `pulumi:"hiveConfig"`
	// The list of labels (key/value pairs) to add to the job.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The config of pag job.
	PigConfig JobPigConfigPtrOutput `pulumi:"pigConfig"`
	// The config of job placement.
	Placement JobPlacementOutput `pulumi:"placement"`
	// The config of presto job
	PrestoConfig JobPrestoConfigPtrOutput `pulumi:"prestoConfig"`
	// The project in which the `cluster` can be found and jobs
	// subsequently run against. If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The config of pySpark job.
	PysparkConfig JobPysparkConfigPtrOutput `pulumi:"pysparkConfig"`
	// The reference of the job
	Reference JobReferenceOutput `pulumi:"reference"`
	// The Cloud Dataproc region. This essentially determines which clusters are available
	// for this job to be submitted to. If not specified, defaults to `global`.
	Region pulumi.StringPtrOutput `pulumi:"region"`
	// Optional. Job scheduling configuration.
	Scheduling JobSchedulingPtrOutput `pulumi:"scheduling"`
	// The config of the Spark job.
	SparkConfig JobSparkConfigPtrOutput `pulumi:"sparkConfig"`
	// The config of SparkSql job
	SparksqlConfig JobSparksqlConfigPtrOutput `pulumi:"sparksqlConfig"`
	// The status of the job.
	Statuses JobStatusArrayOutput `pulumi:"statuses"`
}

// NewJob registers a new resource with the given unique name, arguments, and options.
func NewJob(ctx *pulumi.Context,
	name string, args *JobArgs, opts ...pulumi.ResourceOption) (*Job, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Placement == nil {
		return nil, errors.New("invalid value for required argument 'Placement'")
	}
	var resource Job
	err := ctx.RegisterResource("gcp:dataproc/job:Job", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJob gets an existing Job resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobState, opts ...pulumi.ResourceOption) (*Job, error) {
	var resource Job
	err := ctx.ReadResource("gcp:dataproc/job:Job", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Job resources.
type jobState struct {
	// If present, the location of miscellaneous control files which may be used as part of job setup and handling. If not present, control files may be placed in the same location as driver_output_uri.
	DriverControlsFilesUri *string `pulumi:"driverControlsFilesUri"`
	// A URI pointing to the location of the stdout of the job's driver program.
	DriverOutputResourceUri *string `pulumi:"driverOutputResourceUri"`
	// By default, you can only delete inactive jobs within
	// Dataproc. Setting this to true, and calling destroy, will ensure that the
	// job is first cancelled before issuing the delete.
	ForceDelete *bool `pulumi:"forceDelete"`
	// The config of Hadoop job
	HadoopConfig *JobHadoopConfig `pulumi:"hadoopConfig"`
	// The config of hive job
	HiveConfig *JobHiveConfig `pulumi:"hiveConfig"`
	// The list of labels (key/value pairs) to add to the job.
	Labels map[string]string `pulumi:"labels"`
	// The config of pag job.
	PigConfig *JobPigConfig `pulumi:"pigConfig"`
	// The config of job placement.
	Placement *JobPlacement `pulumi:"placement"`
	// The config of presto job
	PrestoConfig *JobPrestoConfig `pulumi:"prestoConfig"`
	// The project in which the `cluster` can be found and jobs
	// subsequently run against. If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The config of pySpark job.
	PysparkConfig *JobPysparkConfig `pulumi:"pysparkConfig"`
	// The reference of the job
	Reference *JobReference `pulumi:"reference"`
	// The Cloud Dataproc region. This essentially determines which clusters are available
	// for this job to be submitted to. If not specified, defaults to `global`.
	Region *string `pulumi:"region"`
	// Optional. Job scheduling configuration.
	Scheduling *JobScheduling `pulumi:"scheduling"`
	// The config of the Spark job.
	SparkConfig *JobSparkConfig `pulumi:"sparkConfig"`
	// The config of SparkSql job
	SparksqlConfig *JobSparksqlConfig `pulumi:"sparksqlConfig"`
	// The status of the job.
	Statuses []JobStatus `pulumi:"statuses"`
}

type JobState struct {
	// If present, the location of miscellaneous control files which may be used as part of job setup and handling. If not present, control files may be placed in the same location as driver_output_uri.
	DriverControlsFilesUri pulumi.StringPtrInput
	// A URI pointing to the location of the stdout of the job's driver program.
	DriverOutputResourceUri pulumi.StringPtrInput
	// By default, you can only delete inactive jobs within
	// Dataproc. Setting this to true, and calling destroy, will ensure that the
	// job is first cancelled before issuing the delete.
	ForceDelete pulumi.BoolPtrInput
	// The config of Hadoop job
	HadoopConfig JobHadoopConfigPtrInput
	// The config of hive job
	HiveConfig JobHiveConfigPtrInput
	// The list of labels (key/value pairs) to add to the job.
	Labels pulumi.StringMapInput
	// The config of pag job.
	PigConfig JobPigConfigPtrInput
	// The config of job placement.
	Placement JobPlacementPtrInput
	// The config of presto job
	PrestoConfig JobPrestoConfigPtrInput
	// The project in which the `cluster` can be found and jobs
	// subsequently run against. If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The config of pySpark job.
	PysparkConfig JobPysparkConfigPtrInput
	// The reference of the job
	Reference JobReferencePtrInput
	// The Cloud Dataproc region. This essentially determines which clusters are available
	// for this job to be submitted to. If not specified, defaults to `global`.
	Region pulumi.StringPtrInput
	// Optional. Job scheduling configuration.
	Scheduling JobSchedulingPtrInput
	// The config of the Spark job.
	SparkConfig JobSparkConfigPtrInput
	// The config of SparkSql job
	SparksqlConfig JobSparksqlConfigPtrInput
	// The status of the job.
	Statuses JobStatusArrayInput
}

func (JobState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobState)(nil)).Elem()
}

type jobArgs struct {
	// By default, you can only delete inactive jobs within
	// Dataproc. Setting this to true, and calling destroy, will ensure that the
	// job is first cancelled before issuing the delete.
	ForceDelete *bool `pulumi:"forceDelete"`
	// The config of Hadoop job
	HadoopConfig *JobHadoopConfig `pulumi:"hadoopConfig"`
	// The config of hive job
	HiveConfig *JobHiveConfig `pulumi:"hiveConfig"`
	// The list of labels (key/value pairs) to add to the job.
	Labels map[string]string `pulumi:"labels"`
	// The config of pag job.
	PigConfig *JobPigConfig `pulumi:"pigConfig"`
	// The config of job placement.
	Placement JobPlacement `pulumi:"placement"`
	// The config of presto job
	PrestoConfig *JobPrestoConfig `pulumi:"prestoConfig"`
	// The project in which the `cluster` can be found and jobs
	// subsequently run against. If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The config of pySpark job.
	PysparkConfig *JobPysparkConfig `pulumi:"pysparkConfig"`
	// The reference of the job
	Reference *JobReference `pulumi:"reference"`
	// The Cloud Dataproc region. This essentially determines which clusters are available
	// for this job to be submitted to. If not specified, defaults to `global`.
	Region *string `pulumi:"region"`
	// Optional. Job scheduling configuration.
	Scheduling *JobScheduling `pulumi:"scheduling"`
	// The config of the Spark job.
	SparkConfig *JobSparkConfig `pulumi:"sparkConfig"`
	// The config of SparkSql job
	SparksqlConfig *JobSparksqlConfig `pulumi:"sparksqlConfig"`
}

// The set of arguments for constructing a Job resource.
type JobArgs struct {
	// By default, you can only delete inactive jobs within
	// Dataproc. Setting this to true, and calling destroy, will ensure that the
	// job is first cancelled before issuing the delete.
	ForceDelete pulumi.BoolPtrInput
	// The config of Hadoop job
	HadoopConfig JobHadoopConfigPtrInput
	// The config of hive job
	HiveConfig JobHiveConfigPtrInput
	// The list of labels (key/value pairs) to add to the job.
	Labels pulumi.StringMapInput
	// The config of pag job.
	PigConfig JobPigConfigPtrInput
	// The config of job placement.
	Placement JobPlacementInput
	// The config of presto job
	PrestoConfig JobPrestoConfigPtrInput
	// The project in which the `cluster` can be found and jobs
	// subsequently run against. If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The config of pySpark job.
	PysparkConfig JobPysparkConfigPtrInput
	// The reference of the job
	Reference JobReferencePtrInput
	// The Cloud Dataproc region. This essentially determines which clusters are available
	// for this job to be submitted to. If not specified, defaults to `global`.
	Region pulumi.StringPtrInput
	// Optional. Job scheduling configuration.
	Scheduling JobSchedulingPtrInput
	// The config of the Spark job.
	SparkConfig JobSparkConfigPtrInput
	// The config of SparkSql job
	SparksqlConfig JobSparksqlConfigPtrInput
}

func (JobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobArgs)(nil)).Elem()
}

type JobInput interface {
	pulumi.Input

	ToJobOutput() JobOutput
	ToJobOutputWithContext(ctx context.Context) JobOutput
}

func (*Job) ElementType() reflect.Type {
	return reflect.TypeOf((**Job)(nil)).Elem()
}

func (i *Job) ToJobOutput() JobOutput {
	return i.ToJobOutputWithContext(context.Background())
}

func (i *Job) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobOutput)
}

// JobArrayInput is an input type that accepts JobArray and JobArrayOutput values.
// You can construct a concrete instance of `JobArrayInput` via:
//
//	JobArray{ JobArgs{...} }
type JobArrayInput interface {
	pulumi.Input

	ToJobArrayOutput() JobArrayOutput
	ToJobArrayOutputWithContext(context.Context) JobArrayOutput
}

type JobArray []JobInput

func (JobArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Job)(nil)).Elem()
}

func (i JobArray) ToJobArrayOutput() JobArrayOutput {
	return i.ToJobArrayOutputWithContext(context.Background())
}

func (i JobArray) ToJobArrayOutputWithContext(ctx context.Context) JobArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobArrayOutput)
}

// JobMapInput is an input type that accepts JobMap and JobMapOutput values.
// You can construct a concrete instance of `JobMapInput` via:
//
//	JobMap{ "key": JobArgs{...} }
type JobMapInput interface {
	pulumi.Input

	ToJobMapOutput() JobMapOutput
	ToJobMapOutputWithContext(context.Context) JobMapOutput
}

type JobMap map[string]JobInput

func (JobMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Job)(nil)).Elem()
}

func (i JobMap) ToJobMapOutput() JobMapOutput {
	return i.ToJobMapOutputWithContext(context.Background())
}

func (i JobMap) ToJobMapOutputWithContext(ctx context.Context) JobMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobMapOutput)
}

type JobOutput struct{ *pulumi.OutputState }

func (JobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Job)(nil)).Elem()
}

func (o JobOutput) ToJobOutput() JobOutput {
	return o
}

func (o JobOutput) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return o
}

// If present, the location of miscellaneous control files which may be used as part of job setup and handling. If not present, control files may be placed in the same location as driver_output_uri.
func (o JobOutput) DriverControlsFilesUri() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.DriverControlsFilesUri }).(pulumi.StringOutput)
}

// A URI pointing to the location of the stdout of the job's driver program.
func (o JobOutput) DriverOutputResourceUri() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.DriverOutputResourceUri }).(pulumi.StringOutput)
}

// By default, you can only delete inactive jobs within
// Dataproc. Setting this to true, and calling destroy, will ensure that the
// job is first cancelled before issuing the delete.
func (o JobOutput) ForceDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Job) pulumi.BoolPtrOutput { return v.ForceDelete }).(pulumi.BoolPtrOutput)
}

// The config of Hadoop job
func (o JobOutput) HadoopConfig() JobHadoopConfigPtrOutput {
	return o.ApplyT(func(v *Job) JobHadoopConfigPtrOutput { return v.HadoopConfig }).(JobHadoopConfigPtrOutput)
}

// The config of hive job
func (o JobOutput) HiveConfig() JobHiveConfigPtrOutput {
	return o.ApplyT(func(v *Job) JobHiveConfigPtrOutput { return v.HiveConfig }).(JobHiveConfigPtrOutput)
}

// The list of labels (key/value pairs) to add to the job.
func (o JobOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Job) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The config of pag job.
func (o JobOutput) PigConfig() JobPigConfigPtrOutput {
	return o.ApplyT(func(v *Job) JobPigConfigPtrOutput { return v.PigConfig }).(JobPigConfigPtrOutput)
}

// The config of job placement.
func (o JobOutput) Placement() JobPlacementOutput {
	return o.ApplyT(func(v *Job) JobPlacementOutput { return v.Placement }).(JobPlacementOutput)
}

// The config of presto job
func (o JobOutput) PrestoConfig() JobPrestoConfigPtrOutput {
	return o.ApplyT(func(v *Job) JobPrestoConfigPtrOutput { return v.PrestoConfig }).(JobPrestoConfigPtrOutput)
}

// The project in which the `cluster` can be found and jobs
// subsequently run against. If it is not provided, the provider project is used.
func (o JobOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The config of pySpark job.
func (o JobOutput) PysparkConfig() JobPysparkConfigPtrOutput {
	return o.ApplyT(func(v *Job) JobPysparkConfigPtrOutput { return v.PysparkConfig }).(JobPysparkConfigPtrOutput)
}

// The reference of the job
func (o JobOutput) Reference() JobReferenceOutput {
	return o.ApplyT(func(v *Job) JobReferenceOutput { return v.Reference }).(JobReferenceOutput)
}

// The Cloud Dataproc region. This essentially determines which clusters are available
// for this job to be submitted to. If not specified, defaults to `global`.
func (o JobOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Job) pulumi.StringPtrOutput { return v.Region }).(pulumi.StringPtrOutput)
}

// Optional. Job scheduling configuration.
func (o JobOutput) Scheduling() JobSchedulingPtrOutput {
	return o.ApplyT(func(v *Job) JobSchedulingPtrOutput { return v.Scheduling }).(JobSchedulingPtrOutput)
}

// The config of the Spark job.
func (o JobOutput) SparkConfig() JobSparkConfigPtrOutput {
	return o.ApplyT(func(v *Job) JobSparkConfigPtrOutput { return v.SparkConfig }).(JobSparkConfigPtrOutput)
}

// The config of SparkSql job
func (o JobOutput) SparksqlConfig() JobSparksqlConfigPtrOutput {
	return o.ApplyT(func(v *Job) JobSparksqlConfigPtrOutput { return v.SparksqlConfig }).(JobSparksqlConfigPtrOutput)
}

// The status of the job.
func (o JobOutput) Statuses() JobStatusArrayOutput {
	return o.ApplyT(func(v *Job) JobStatusArrayOutput { return v.Statuses }).(JobStatusArrayOutput)
}

type JobArrayOutput struct{ *pulumi.OutputState }

func (JobArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Job)(nil)).Elem()
}

func (o JobArrayOutput) ToJobArrayOutput() JobArrayOutput {
	return o
}

func (o JobArrayOutput) ToJobArrayOutputWithContext(ctx context.Context) JobArrayOutput {
	return o
}

func (o JobArrayOutput) Index(i pulumi.IntInput) JobOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Job {
		return vs[0].([]*Job)[vs[1].(int)]
	}).(JobOutput)
}

type JobMapOutput struct{ *pulumi.OutputState }

func (JobMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Job)(nil)).Elem()
}

func (o JobMapOutput) ToJobMapOutput() JobMapOutput {
	return o
}

func (o JobMapOutput) ToJobMapOutputWithContext(ctx context.Context) JobMapOutput {
	return o
}

func (o JobMapOutput) MapIndex(k pulumi.StringInput) JobOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Job {
		return vs[0].(map[string]*Job)[vs[1].(string)]
	}).(JobOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*JobInput)(nil)).Elem(), &Job{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobArrayInput)(nil)).Elem(), JobArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobMapInput)(nil)).Elem(), JobMap{})
	pulumi.RegisterOutputType(JobOutput{})
	pulumi.RegisterOutputType(JobArrayOutput{})
	pulumi.RegisterOutputType(JobMapOutput{})
}
