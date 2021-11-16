// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dataflow

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a job on Dataflow, which is an implementation of Apache Beam running on Google Compute Engine. For more information see
// the official documentation for
// [Beam](https://beam.apache.org) and [Dataflow](https://cloud.google.com/dataflow/).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/dataflow"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := dataflow.NewJob(ctx, "bigDataJob", &dataflow.JobArgs{
// 			Parameters: pulumi.AnyMap{
// 				"baz": pulumi.Any("qux"),
// 				"foo": pulumi.Any("bar"),
// 			},
// 			TempGcsLocation: pulumi.String("gs://my-bucket/tmp_dir"),
// 			TemplateGcsPath: pulumi.String("gs://my-bucket/templates/template_file"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Streaming Job
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/dataflow"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/pubsub"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/storage"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		topic, err := pubsub.NewTopic(ctx, "topic", nil)
// 		if err != nil {
// 			return err
// 		}
// 		bucket1, err := storage.NewBucket(ctx, "bucket1", &storage.BucketArgs{
// 			Location:     pulumi.String("US"),
// 			ForceDestroy: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = storage.NewBucket(ctx, "bucket2", &storage.BucketArgs{
// 			Location:     pulumi.String("US"),
// 			ForceDestroy: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = dataflow.NewJob(ctx, "pubsubStream", &dataflow.JobArgs{
// 			TemplateGcsPath:       pulumi.String("gs://my-bucket/templates/template_file"),
// 			TempGcsLocation:       pulumi.String("gs://my-bucket/tmp_dir"),
// 			EnableStreamingEngine: pulumi.Bool(true),
// 			Parameters: pulumi.AnyMap{
// 				"inputFilePattern": bucket1.Url.ApplyT(func(url string) (string, error) {
// 					return fmt.Sprintf("%v%v", url, "/*.json"), nil
// 				}).(pulumi.StringOutput),
// 				"outputTopic": topic.ID(),
// 			},
// 			TransformNameMapping: pulumi.AnyMap{
// 				"name": pulumi.Any("test_job"),
// 				"env":  pulumi.Any("test"),
// 			},
// 			OnDelete: pulumi.String("cancel"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Note on "destroy" / "apply"
//
// There are many types of Dataflow jobs.  Some Dataflow jobs run constantly, getting new data from (e.g.) a GCS bucket, and outputting data continuously.  Some jobs process a set amount of data then terminate.  All jobs can fail while running due to programming errors or other issues.  In this way, Dataflow jobs are different from most other Google resources.
//
// The Dataflow resource is considered 'existing' while it is in a nonterminal state.  If it reaches a terminal state (e.g. 'FAILED', 'COMPLETE', 'CANCELLED'), it will be recreated on the next 'apply'.  This is as expected for jobs which run continuously, but may surprise users who use this resource for other kinds of Dataflow jobs.
//
// A Dataflow job which is 'destroyed' may be "cancelled" or "drained".  If "cancelled", the job terminates - any data written remains where it is, but no new data will be processed.  If "drained", no new data will enter the pipeline, but any data currently in the pipeline will finish being processed.  The default is "drain". When `onDelete` is set to `"drain"` in the configuration, you may experience a long wait for your `pulumi destroy` to complete.
//
// ## Import
//
// This resource does not support import.
type Job struct {
	pulumi.CustomResourceState

	// List of experiments that should be used by the job. An example value is `["enableStackdriverAgentMetrics"]`.
	AdditionalExperiments pulumi.StringArrayOutput `pulumi:"additionalExperiments"`
	// Enable/disable the use of [Streaming Engine](https://cloud.google.com/dataflow/docs/guides/deploying-a-pipeline#streaming-engine) for the job. Note that Streaming Engine is enabled by default for pipelines developed against the Beam SDK for Python v2.21.0 or later when using Python 3.
	EnableStreamingEngine pulumi.BoolPtrOutput `pulumi:"enableStreamingEngine"`
	// The configuration for VM IPs.  Options are `"WORKER_IP_PUBLIC"` or `"WORKER_IP_PRIVATE"`.
	IpConfiguration pulumi.StringPtrOutput `pulumi:"ipConfiguration"`
	// The unique ID of this job.
	JobId pulumi.StringOutput `pulumi:"jobId"`
	// The name for the Cloud KMS key for the job. Key format is: `projects/PROJECT_ID/locations/LOCATION/keyRings/KEY_RING/cryptoKeys/KEY`
	KmsKeyName pulumi.StringPtrOutput `pulumi:"kmsKeyName"`
	// User labels to be specified for the job. Keys and values should follow the restrictions
	// specified in the [labeling restrictions](https://cloud.google.com/compute/docs/labeling-resources#restrictions) page.
	// **NOTE**: Google-provided Dataflow templates often provide default labels that begin with `goog-dataflow-provided`.
	// Unless explicitly set in config, these labels will be ignored to prevent diffs on re-apply.
	Labels pulumi.MapOutput `pulumi:"labels"`
	// The machine type to use for the job.
	MachineType pulumi.StringPtrOutput `pulumi:"machineType"`
	// The number of workers permitted to work on the job.  More workers may improve processing speed at additional cost.
	MaxWorkers pulumi.IntPtrOutput `pulumi:"maxWorkers"`
	// A unique name for the resource, required by Dataflow.
	Name pulumi.StringOutput `pulumi:"name"`
	// The network to which VMs will be assigned. If it is not provided, "default" will be used.
	Network pulumi.StringPtrOutput `pulumi:"network"`
	// One of "drain" or "cancel".  Specifies behavior of deletion during `pulumi destroy`.  See above note.
	OnDelete pulumi.StringPtrOutput `pulumi:"onDelete"`
	// Key/Value pairs to be passed to the Dataflow job (as used in the template).
	Parameters pulumi.MapOutput `pulumi:"parameters"`
	// The project in which the resource belongs. If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The region in which the created job should run.
	Region pulumi.StringPtrOutput `pulumi:"region"`
	// The Service Account email used to create the job.
	ServiceAccountEmail pulumi.StringPtrOutput `pulumi:"serviceAccountEmail"`
	// The current state of the resource, selected from the [JobState enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobState)
	State pulumi.StringOutput `pulumi:"state"`
	// The subnetwork to which VMs will be assigned. Should be of the form "regions/REGION/subnetworks/SUBNETWORK".
	Subnetwork pulumi.StringPtrOutput `pulumi:"subnetwork"`
	// A writeable location on GCS for the Dataflow job to dump its temporary data.
	TempGcsLocation pulumi.StringOutput `pulumi:"tempGcsLocation"`
	// The GCS path to the Dataflow job template.
	TemplateGcsPath pulumi.StringOutput `pulumi:"templateGcsPath"`
	// Only applicable when updating a pipeline. Map of transform name prefixes of the job to be replaced with the corresponding name prefixes of the new job. This field is not used outside of update.
	TransformNameMapping pulumi.MapOutput `pulumi:"transformNameMapping"`
	// The type of this job, selected from the [JobType enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobType)
	Type pulumi.StringOutput `pulumi:"type"`
	// The zone in which the created job should run. If it is not provided, the provider zone is used.
	Zone pulumi.StringPtrOutput `pulumi:"zone"`
}

// NewJob registers a new resource with the given unique name, arguments, and options.
func NewJob(ctx *pulumi.Context,
	name string, args *JobArgs, opts ...pulumi.ResourceOption) (*Job, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.TempGcsLocation == nil {
		return nil, errors.New("invalid value for required argument 'TempGcsLocation'")
	}
	if args.TemplateGcsPath == nil {
		return nil, errors.New("invalid value for required argument 'TemplateGcsPath'")
	}
	var resource Job
	err := ctx.RegisterResource("gcp:dataflow/job:Job", name, args, &resource, opts...)
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
	err := ctx.ReadResource("gcp:dataflow/job:Job", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Job resources.
type jobState struct {
	// List of experiments that should be used by the job. An example value is `["enableStackdriverAgentMetrics"]`.
	AdditionalExperiments []string `pulumi:"additionalExperiments"`
	// Enable/disable the use of [Streaming Engine](https://cloud.google.com/dataflow/docs/guides/deploying-a-pipeline#streaming-engine) for the job. Note that Streaming Engine is enabled by default for pipelines developed against the Beam SDK for Python v2.21.0 or later when using Python 3.
	EnableStreamingEngine *bool `pulumi:"enableStreamingEngine"`
	// The configuration for VM IPs.  Options are `"WORKER_IP_PUBLIC"` or `"WORKER_IP_PRIVATE"`.
	IpConfiguration *string `pulumi:"ipConfiguration"`
	// The unique ID of this job.
	JobId *string `pulumi:"jobId"`
	// The name for the Cloud KMS key for the job. Key format is: `projects/PROJECT_ID/locations/LOCATION/keyRings/KEY_RING/cryptoKeys/KEY`
	KmsKeyName *string `pulumi:"kmsKeyName"`
	// User labels to be specified for the job. Keys and values should follow the restrictions
	// specified in the [labeling restrictions](https://cloud.google.com/compute/docs/labeling-resources#restrictions) page.
	// **NOTE**: Google-provided Dataflow templates often provide default labels that begin with `goog-dataflow-provided`.
	// Unless explicitly set in config, these labels will be ignored to prevent diffs on re-apply.
	Labels map[string]interface{} `pulumi:"labels"`
	// The machine type to use for the job.
	MachineType *string `pulumi:"machineType"`
	// The number of workers permitted to work on the job.  More workers may improve processing speed at additional cost.
	MaxWorkers *int `pulumi:"maxWorkers"`
	// A unique name for the resource, required by Dataflow.
	Name *string `pulumi:"name"`
	// The network to which VMs will be assigned. If it is not provided, "default" will be used.
	Network *string `pulumi:"network"`
	// One of "drain" or "cancel".  Specifies behavior of deletion during `pulumi destroy`.  See above note.
	OnDelete *string `pulumi:"onDelete"`
	// Key/Value pairs to be passed to the Dataflow job (as used in the template).
	Parameters map[string]interface{} `pulumi:"parameters"`
	// The project in which the resource belongs. If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region in which the created job should run.
	Region *string `pulumi:"region"`
	// The Service Account email used to create the job.
	ServiceAccountEmail *string `pulumi:"serviceAccountEmail"`
	// The current state of the resource, selected from the [JobState enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobState)
	State *string `pulumi:"state"`
	// The subnetwork to which VMs will be assigned. Should be of the form "regions/REGION/subnetworks/SUBNETWORK".
	Subnetwork *string `pulumi:"subnetwork"`
	// A writeable location on GCS for the Dataflow job to dump its temporary data.
	TempGcsLocation *string `pulumi:"tempGcsLocation"`
	// The GCS path to the Dataflow job template.
	TemplateGcsPath *string `pulumi:"templateGcsPath"`
	// Only applicable when updating a pipeline. Map of transform name prefixes of the job to be replaced with the corresponding name prefixes of the new job. This field is not used outside of update.
	TransformNameMapping map[string]interface{} `pulumi:"transformNameMapping"`
	// The type of this job, selected from the [JobType enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobType)
	Type *string `pulumi:"type"`
	// The zone in which the created job should run. If it is not provided, the provider zone is used.
	Zone *string `pulumi:"zone"`
}

type JobState struct {
	// List of experiments that should be used by the job. An example value is `["enableStackdriverAgentMetrics"]`.
	AdditionalExperiments pulumi.StringArrayInput
	// Enable/disable the use of [Streaming Engine](https://cloud.google.com/dataflow/docs/guides/deploying-a-pipeline#streaming-engine) for the job. Note that Streaming Engine is enabled by default for pipelines developed against the Beam SDK for Python v2.21.0 or later when using Python 3.
	EnableStreamingEngine pulumi.BoolPtrInput
	// The configuration for VM IPs.  Options are `"WORKER_IP_PUBLIC"` or `"WORKER_IP_PRIVATE"`.
	IpConfiguration pulumi.StringPtrInput
	// The unique ID of this job.
	JobId pulumi.StringPtrInput
	// The name for the Cloud KMS key for the job. Key format is: `projects/PROJECT_ID/locations/LOCATION/keyRings/KEY_RING/cryptoKeys/KEY`
	KmsKeyName pulumi.StringPtrInput
	// User labels to be specified for the job. Keys and values should follow the restrictions
	// specified in the [labeling restrictions](https://cloud.google.com/compute/docs/labeling-resources#restrictions) page.
	// **NOTE**: Google-provided Dataflow templates often provide default labels that begin with `goog-dataflow-provided`.
	// Unless explicitly set in config, these labels will be ignored to prevent diffs on re-apply.
	Labels pulumi.MapInput
	// The machine type to use for the job.
	MachineType pulumi.StringPtrInput
	// The number of workers permitted to work on the job.  More workers may improve processing speed at additional cost.
	MaxWorkers pulumi.IntPtrInput
	// A unique name for the resource, required by Dataflow.
	Name pulumi.StringPtrInput
	// The network to which VMs will be assigned. If it is not provided, "default" will be used.
	Network pulumi.StringPtrInput
	// One of "drain" or "cancel".  Specifies behavior of deletion during `pulumi destroy`.  See above note.
	OnDelete pulumi.StringPtrInput
	// Key/Value pairs to be passed to the Dataflow job (as used in the template).
	Parameters pulumi.MapInput
	// The project in which the resource belongs. If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The region in which the created job should run.
	Region pulumi.StringPtrInput
	// The Service Account email used to create the job.
	ServiceAccountEmail pulumi.StringPtrInput
	// The current state of the resource, selected from the [JobState enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobState)
	State pulumi.StringPtrInput
	// The subnetwork to which VMs will be assigned. Should be of the form "regions/REGION/subnetworks/SUBNETWORK".
	Subnetwork pulumi.StringPtrInput
	// A writeable location on GCS for the Dataflow job to dump its temporary data.
	TempGcsLocation pulumi.StringPtrInput
	// The GCS path to the Dataflow job template.
	TemplateGcsPath pulumi.StringPtrInput
	// Only applicable when updating a pipeline. Map of transform name prefixes of the job to be replaced with the corresponding name prefixes of the new job. This field is not used outside of update.
	TransformNameMapping pulumi.MapInput
	// The type of this job, selected from the [JobType enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobType)
	Type pulumi.StringPtrInput
	// The zone in which the created job should run. If it is not provided, the provider zone is used.
	Zone pulumi.StringPtrInput
}

func (JobState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobState)(nil)).Elem()
}

type jobArgs struct {
	// List of experiments that should be used by the job. An example value is `["enableStackdriverAgentMetrics"]`.
	AdditionalExperiments []string `pulumi:"additionalExperiments"`
	// Enable/disable the use of [Streaming Engine](https://cloud.google.com/dataflow/docs/guides/deploying-a-pipeline#streaming-engine) for the job. Note that Streaming Engine is enabled by default for pipelines developed against the Beam SDK for Python v2.21.0 or later when using Python 3.
	EnableStreamingEngine *bool `pulumi:"enableStreamingEngine"`
	// The configuration for VM IPs.  Options are `"WORKER_IP_PUBLIC"` or `"WORKER_IP_PRIVATE"`.
	IpConfiguration *string `pulumi:"ipConfiguration"`
	// The name for the Cloud KMS key for the job. Key format is: `projects/PROJECT_ID/locations/LOCATION/keyRings/KEY_RING/cryptoKeys/KEY`
	KmsKeyName *string `pulumi:"kmsKeyName"`
	// User labels to be specified for the job. Keys and values should follow the restrictions
	// specified in the [labeling restrictions](https://cloud.google.com/compute/docs/labeling-resources#restrictions) page.
	// **NOTE**: Google-provided Dataflow templates often provide default labels that begin with `goog-dataflow-provided`.
	// Unless explicitly set in config, these labels will be ignored to prevent diffs on re-apply.
	Labels map[string]interface{} `pulumi:"labels"`
	// The machine type to use for the job.
	MachineType *string `pulumi:"machineType"`
	// The number of workers permitted to work on the job.  More workers may improve processing speed at additional cost.
	MaxWorkers *int `pulumi:"maxWorkers"`
	// A unique name for the resource, required by Dataflow.
	Name *string `pulumi:"name"`
	// The network to which VMs will be assigned. If it is not provided, "default" will be used.
	Network *string `pulumi:"network"`
	// One of "drain" or "cancel".  Specifies behavior of deletion during `pulumi destroy`.  See above note.
	OnDelete *string `pulumi:"onDelete"`
	// Key/Value pairs to be passed to the Dataflow job (as used in the template).
	Parameters map[string]interface{} `pulumi:"parameters"`
	// The project in which the resource belongs. If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region in which the created job should run.
	Region *string `pulumi:"region"`
	// The Service Account email used to create the job.
	ServiceAccountEmail *string `pulumi:"serviceAccountEmail"`
	// The subnetwork to which VMs will be assigned. Should be of the form "regions/REGION/subnetworks/SUBNETWORK".
	Subnetwork *string `pulumi:"subnetwork"`
	// A writeable location on GCS for the Dataflow job to dump its temporary data.
	TempGcsLocation string `pulumi:"tempGcsLocation"`
	// The GCS path to the Dataflow job template.
	TemplateGcsPath string `pulumi:"templateGcsPath"`
	// Only applicable when updating a pipeline. Map of transform name prefixes of the job to be replaced with the corresponding name prefixes of the new job. This field is not used outside of update.
	TransformNameMapping map[string]interface{} `pulumi:"transformNameMapping"`
	// The zone in which the created job should run. If it is not provided, the provider zone is used.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a Job resource.
type JobArgs struct {
	// List of experiments that should be used by the job. An example value is `["enableStackdriverAgentMetrics"]`.
	AdditionalExperiments pulumi.StringArrayInput
	// Enable/disable the use of [Streaming Engine](https://cloud.google.com/dataflow/docs/guides/deploying-a-pipeline#streaming-engine) for the job. Note that Streaming Engine is enabled by default for pipelines developed against the Beam SDK for Python v2.21.0 or later when using Python 3.
	EnableStreamingEngine pulumi.BoolPtrInput
	// The configuration for VM IPs.  Options are `"WORKER_IP_PUBLIC"` or `"WORKER_IP_PRIVATE"`.
	IpConfiguration pulumi.StringPtrInput
	// The name for the Cloud KMS key for the job. Key format is: `projects/PROJECT_ID/locations/LOCATION/keyRings/KEY_RING/cryptoKeys/KEY`
	KmsKeyName pulumi.StringPtrInput
	// User labels to be specified for the job. Keys and values should follow the restrictions
	// specified in the [labeling restrictions](https://cloud.google.com/compute/docs/labeling-resources#restrictions) page.
	// **NOTE**: Google-provided Dataflow templates often provide default labels that begin with `goog-dataflow-provided`.
	// Unless explicitly set in config, these labels will be ignored to prevent diffs on re-apply.
	Labels pulumi.MapInput
	// The machine type to use for the job.
	MachineType pulumi.StringPtrInput
	// The number of workers permitted to work on the job.  More workers may improve processing speed at additional cost.
	MaxWorkers pulumi.IntPtrInput
	// A unique name for the resource, required by Dataflow.
	Name pulumi.StringPtrInput
	// The network to which VMs will be assigned. If it is not provided, "default" will be used.
	Network pulumi.StringPtrInput
	// One of "drain" or "cancel".  Specifies behavior of deletion during `pulumi destroy`.  See above note.
	OnDelete pulumi.StringPtrInput
	// Key/Value pairs to be passed to the Dataflow job (as used in the template).
	Parameters pulumi.MapInput
	// The project in which the resource belongs. If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The region in which the created job should run.
	Region pulumi.StringPtrInput
	// The Service Account email used to create the job.
	ServiceAccountEmail pulumi.StringPtrInput
	// The subnetwork to which VMs will be assigned. Should be of the form "regions/REGION/subnetworks/SUBNETWORK".
	Subnetwork pulumi.StringPtrInput
	// A writeable location on GCS for the Dataflow job to dump its temporary data.
	TempGcsLocation pulumi.StringInput
	// The GCS path to the Dataflow job template.
	TemplateGcsPath pulumi.StringInput
	// Only applicable when updating a pipeline. Map of transform name prefixes of the job to be replaced with the corresponding name prefixes of the new job. This field is not used outside of update.
	TransformNameMapping pulumi.MapInput
	// The zone in which the created job should run. If it is not provided, the provider zone is used.
	Zone pulumi.StringPtrInput
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
	return reflect.TypeOf((*Job)(nil))
}

func (i *Job) ToJobOutput() JobOutput {
	return i.ToJobOutputWithContext(context.Background())
}

func (i *Job) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobOutput)
}

func (i *Job) ToJobPtrOutput() JobPtrOutput {
	return i.ToJobPtrOutputWithContext(context.Background())
}

func (i *Job) ToJobPtrOutputWithContext(ctx context.Context) JobPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobPtrOutput)
}

type JobPtrInput interface {
	pulumi.Input

	ToJobPtrOutput() JobPtrOutput
	ToJobPtrOutputWithContext(ctx context.Context) JobPtrOutput
}

type jobPtrType JobArgs

func (*jobPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Job)(nil))
}

func (i *jobPtrType) ToJobPtrOutput() JobPtrOutput {
	return i.ToJobPtrOutputWithContext(context.Background())
}

func (i *jobPtrType) ToJobPtrOutputWithContext(ctx context.Context) JobPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobPtrOutput)
}

// JobArrayInput is an input type that accepts JobArray and JobArrayOutput values.
// You can construct a concrete instance of `JobArrayInput` via:
//
//          JobArray{ JobArgs{...} }
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
//          JobMap{ "key": JobArgs{...} }
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
	return reflect.TypeOf((*Job)(nil))
}

func (o JobOutput) ToJobOutput() JobOutput {
	return o
}

func (o JobOutput) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return o
}

func (o JobOutput) ToJobPtrOutput() JobPtrOutput {
	return o.ToJobPtrOutputWithContext(context.Background())
}

func (o JobOutput) ToJobPtrOutputWithContext(ctx context.Context) JobPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Job) *Job {
		return &v
	}).(JobPtrOutput)
}

type JobPtrOutput struct{ *pulumi.OutputState }

func (JobPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Job)(nil))
}

func (o JobPtrOutput) ToJobPtrOutput() JobPtrOutput {
	return o
}

func (o JobPtrOutput) ToJobPtrOutputWithContext(ctx context.Context) JobPtrOutput {
	return o
}

func (o JobPtrOutput) Elem() JobOutput {
	return o.ApplyT(func(v *Job) Job {
		if v != nil {
			return *v
		}
		var ret Job
		return ret
	}).(JobOutput)
}

type JobArrayOutput struct{ *pulumi.OutputState }

func (JobArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Job)(nil))
}

func (o JobArrayOutput) ToJobArrayOutput() JobArrayOutput {
	return o
}

func (o JobArrayOutput) ToJobArrayOutputWithContext(ctx context.Context) JobArrayOutput {
	return o
}

func (o JobArrayOutput) Index(i pulumi.IntInput) JobOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Job {
		return vs[0].([]Job)[vs[1].(int)]
	}).(JobOutput)
}

type JobMapOutput struct{ *pulumi.OutputState }

func (JobMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Job)(nil))
}

func (o JobMapOutput) ToJobMapOutput() JobMapOutput {
	return o
}

func (o JobMapOutput) ToJobMapOutputWithContext(ctx context.Context) JobMapOutput {
	return o
}

func (o JobMapOutput) MapIndex(k pulumi.StringInput) JobOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Job {
		return vs[0].(map[string]Job)[vs[1].(string)]
	}).(JobOutput)
}

func init() {
	pulumi.RegisterOutputType(JobOutput{})
	pulumi.RegisterOutputType(JobPtrOutput{})
	pulumi.RegisterOutputType(JobArrayOutput{})
	pulumi.RegisterOutputType(JobMapOutput{})
}
