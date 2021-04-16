// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dataflow

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a [Flex Template](https://cloud.google.com/dataflow/docs/guides/templates/using-flex-templates)
// job on Dataflow, which is an implementation of Apache Beam running on Google
// Compute Engine. For more information see the official documentation for [Beam](https://beam.apache.org)
// and [Dataflow](https://cloud.google.com/dataflow/).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/dataflow"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := dataflow.NewFlexTemplateJob(ctx, "bigDataJob", &dataflow.FlexTemplateJobArgs{
// 			ContainerSpecGcsPath: pulumi.String("gs://my-bucket/templates/template.json"),
// 			Parameters: pulumi.StringMap{
// 				"inputSubscription": pulumi.String("messages"),
// 			},
// 		}, pulumi.Provider(google_beta))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ## Note on "destroy" / "apply"
//
// There are many types of Dataflow jobs.  Some Dataflow jobs run constantly,
// getting new data from (e.g.) a GCS bucket, and outputting data continuously.
// Some jobs process a set amount of data then terminate. All jobs can fail while
// running due to programming errors or other issues. In this way, Dataflow jobs
// are different from most other provider / Google resources.
//
// The Dataflow resource is considered 'existing' while it is in a nonterminal
// state.  If it reaches a terminal state (e.g. 'FAILED', 'COMPLETE',
// 'CANCELLED'), it will be recreated on the next 'apply'.  This is as expected for
// jobs which run continuously, but may surprise users who use this resource for
// other kinds of Dataflow jobs.
//
// A Dataflow job which is 'destroyed' may be "cancelled" or "drained".  If
// "cancelled", the job terminates - any data written remains where it is, but no
// new data will be processed.  If "drained", no new data will enter the pipeline,
// but any data currently in the pipeline will finish being processed.  The default
// is "cancelled", but if a user sets `onDelete` to `"drain"` in the
// configuration, you may experience a long wait for your `pulumi destroy` to
// complete.
//
// ## Import
//
// This resource does not support import.
type FlexTemplateJob struct {
	pulumi.CustomResourceState

	// The GCS path to the Dataflow job Flex
	// Template.
	ContainerSpecGcsPath pulumi.StringOutput `pulumi:"containerSpecGcsPath"`
	// The unique ID of this job.
	JobId pulumi.StringOutput `pulumi:"jobId"`
	// User labels to be specified for the job. Keys and values
	// should follow the restrictions specified in the [labeling restrictions](https://cloud.google.com/compute/docs/labeling-resources#restrictions)
	// page. **Note**: This field is marked as deprecated as the API does not currently
	// support adding labels.
	// **NOTE**: Google-provided Dataflow templates often provide default labels
	// that begin with `goog-dataflow-provided`. Unless explicitly set in config, these
	// labels will be ignored to prevent diffs on re-apply.
	//
	// Deprecated: Deprecated until the API supports this field
	Labels pulumi.MapOutput `pulumi:"labels"`
	// A unique name for the resource, required by Dataflow.
	Name pulumi.StringOutput `pulumi:"name"`
	// One of "drain" or "cancel". Specifies behavior of
	// deletion during `pulumi destroy`.  See above note.
	OnDelete pulumi.StringPtrOutput `pulumi:"onDelete"`
	// Key/Value pairs to be passed to the Dataflow job (as
	// used in the template). Additional [pipeline options](https://cloud.google.com/dataflow/docs/guides/specifying-exec-params#setting-other-cloud-dataflow-pipeline-options)
	// such as `serviceAccount`, `workerMachineType`, etc can be specified here.
	Parameters pulumi.MapOutput `pulumi:"parameters"`
	// The project in which the resource belongs. If it is not
	// provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The region in which the created job should run.
	Region pulumi.StringOutput `pulumi:"region"`
	// The current state of the resource, selected from the [JobState enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobState)
	State pulumi.StringOutput `pulumi:"state"`
}

// NewFlexTemplateJob registers a new resource with the given unique name, arguments, and options.
func NewFlexTemplateJob(ctx *pulumi.Context,
	name string, args *FlexTemplateJobArgs, opts ...pulumi.ResourceOption) (*FlexTemplateJob, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ContainerSpecGcsPath == nil {
		return nil, errors.New("invalid value for required argument 'ContainerSpecGcsPath'")
	}
	var resource FlexTemplateJob
	err := ctx.RegisterResource("gcp:dataflow/flexTemplateJob:FlexTemplateJob", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFlexTemplateJob gets an existing FlexTemplateJob resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFlexTemplateJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FlexTemplateJobState, opts ...pulumi.ResourceOption) (*FlexTemplateJob, error) {
	var resource FlexTemplateJob
	err := ctx.ReadResource("gcp:dataflow/flexTemplateJob:FlexTemplateJob", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FlexTemplateJob resources.
type flexTemplateJobState struct {
	// The GCS path to the Dataflow job Flex
	// Template.
	ContainerSpecGcsPath *string `pulumi:"containerSpecGcsPath"`
	// The unique ID of this job.
	JobId *string `pulumi:"jobId"`
	// User labels to be specified for the job. Keys and values
	// should follow the restrictions specified in the [labeling restrictions](https://cloud.google.com/compute/docs/labeling-resources#restrictions)
	// page. **Note**: This field is marked as deprecated as the API does not currently
	// support adding labels.
	// **NOTE**: Google-provided Dataflow templates often provide default labels
	// that begin with `goog-dataflow-provided`. Unless explicitly set in config, these
	// labels will be ignored to prevent diffs on re-apply.
	//
	// Deprecated: Deprecated until the API supports this field
	Labels map[string]interface{} `pulumi:"labels"`
	// A unique name for the resource, required by Dataflow.
	Name *string `pulumi:"name"`
	// One of "drain" or "cancel". Specifies behavior of
	// deletion during `pulumi destroy`.  See above note.
	OnDelete *string `pulumi:"onDelete"`
	// Key/Value pairs to be passed to the Dataflow job (as
	// used in the template). Additional [pipeline options](https://cloud.google.com/dataflow/docs/guides/specifying-exec-params#setting-other-cloud-dataflow-pipeline-options)
	// such as `serviceAccount`, `workerMachineType`, etc can be specified here.
	Parameters map[string]interface{} `pulumi:"parameters"`
	// The project in which the resource belongs. If it is not
	// provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region in which the created job should run.
	Region *string `pulumi:"region"`
	// The current state of the resource, selected from the [JobState enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobState)
	State *string `pulumi:"state"`
}

type FlexTemplateJobState struct {
	// The GCS path to the Dataflow job Flex
	// Template.
	ContainerSpecGcsPath pulumi.StringPtrInput
	// The unique ID of this job.
	JobId pulumi.StringPtrInput
	// User labels to be specified for the job. Keys and values
	// should follow the restrictions specified in the [labeling restrictions](https://cloud.google.com/compute/docs/labeling-resources#restrictions)
	// page. **Note**: This field is marked as deprecated as the API does not currently
	// support adding labels.
	// **NOTE**: Google-provided Dataflow templates often provide default labels
	// that begin with `goog-dataflow-provided`. Unless explicitly set in config, these
	// labels will be ignored to prevent diffs on re-apply.
	//
	// Deprecated: Deprecated until the API supports this field
	Labels pulumi.MapInput
	// A unique name for the resource, required by Dataflow.
	Name pulumi.StringPtrInput
	// One of "drain" or "cancel". Specifies behavior of
	// deletion during `pulumi destroy`.  See above note.
	OnDelete pulumi.StringPtrInput
	// Key/Value pairs to be passed to the Dataflow job (as
	// used in the template). Additional [pipeline options](https://cloud.google.com/dataflow/docs/guides/specifying-exec-params#setting-other-cloud-dataflow-pipeline-options)
	// such as `serviceAccount`, `workerMachineType`, etc can be specified here.
	Parameters pulumi.MapInput
	// The project in which the resource belongs. If it is not
	// provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The region in which the created job should run.
	Region pulumi.StringPtrInput
	// The current state of the resource, selected from the [JobState enum](https://cloud.google.com/dataflow/docs/reference/rest/v1b3/projects.jobs#Job.JobState)
	State pulumi.StringPtrInput
}

func (FlexTemplateJobState) ElementType() reflect.Type {
	return reflect.TypeOf((*flexTemplateJobState)(nil)).Elem()
}

type flexTemplateJobArgs struct {
	// The GCS path to the Dataflow job Flex
	// Template.
	ContainerSpecGcsPath string `pulumi:"containerSpecGcsPath"`
	// User labels to be specified for the job. Keys and values
	// should follow the restrictions specified in the [labeling restrictions](https://cloud.google.com/compute/docs/labeling-resources#restrictions)
	// page. **Note**: This field is marked as deprecated as the API does not currently
	// support adding labels.
	// **NOTE**: Google-provided Dataflow templates often provide default labels
	// that begin with `goog-dataflow-provided`. Unless explicitly set in config, these
	// labels will be ignored to prevent diffs on re-apply.
	//
	// Deprecated: Deprecated until the API supports this field
	Labels map[string]interface{} `pulumi:"labels"`
	// A unique name for the resource, required by Dataflow.
	Name *string `pulumi:"name"`
	// One of "drain" or "cancel". Specifies behavior of
	// deletion during `pulumi destroy`.  See above note.
	OnDelete *string `pulumi:"onDelete"`
	// Key/Value pairs to be passed to the Dataflow job (as
	// used in the template). Additional [pipeline options](https://cloud.google.com/dataflow/docs/guides/specifying-exec-params#setting-other-cloud-dataflow-pipeline-options)
	// such as `serviceAccount`, `workerMachineType`, etc can be specified here.
	Parameters map[string]interface{} `pulumi:"parameters"`
	// The project in which the resource belongs. If it is not
	// provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The region in which the created job should run.
	Region *string `pulumi:"region"`
}

// The set of arguments for constructing a FlexTemplateJob resource.
type FlexTemplateJobArgs struct {
	// The GCS path to the Dataflow job Flex
	// Template.
	ContainerSpecGcsPath pulumi.StringInput
	// User labels to be specified for the job. Keys and values
	// should follow the restrictions specified in the [labeling restrictions](https://cloud.google.com/compute/docs/labeling-resources#restrictions)
	// page. **Note**: This field is marked as deprecated as the API does not currently
	// support adding labels.
	// **NOTE**: Google-provided Dataflow templates often provide default labels
	// that begin with `goog-dataflow-provided`. Unless explicitly set in config, these
	// labels will be ignored to prevent diffs on re-apply.
	//
	// Deprecated: Deprecated until the API supports this field
	Labels pulumi.MapInput
	// A unique name for the resource, required by Dataflow.
	Name pulumi.StringPtrInput
	// One of "drain" or "cancel". Specifies behavior of
	// deletion during `pulumi destroy`.  See above note.
	OnDelete pulumi.StringPtrInput
	// Key/Value pairs to be passed to the Dataflow job (as
	// used in the template). Additional [pipeline options](https://cloud.google.com/dataflow/docs/guides/specifying-exec-params#setting-other-cloud-dataflow-pipeline-options)
	// such as `serviceAccount`, `workerMachineType`, etc can be specified here.
	Parameters pulumi.MapInput
	// The project in which the resource belongs. If it is not
	// provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The region in which the created job should run.
	Region pulumi.StringPtrInput
}

func (FlexTemplateJobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*flexTemplateJobArgs)(nil)).Elem()
}

type FlexTemplateJobInput interface {
	pulumi.Input

	ToFlexTemplateJobOutput() FlexTemplateJobOutput
	ToFlexTemplateJobOutputWithContext(ctx context.Context) FlexTemplateJobOutput
}

func (*FlexTemplateJob) ElementType() reflect.Type {
	return reflect.TypeOf((*FlexTemplateJob)(nil))
}

func (i *FlexTemplateJob) ToFlexTemplateJobOutput() FlexTemplateJobOutput {
	return i.ToFlexTemplateJobOutputWithContext(context.Background())
}

func (i *FlexTemplateJob) ToFlexTemplateJobOutputWithContext(ctx context.Context) FlexTemplateJobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlexTemplateJobOutput)
}

func (i *FlexTemplateJob) ToFlexTemplateJobPtrOutput() FlexTemplateJobPtrOutput {
	return i.ToFlexTemplateJobPtrOutputWithContext(context.Background())
}

func (i *FlexTemplateJob) ToFlexTemplateJobPtrOutputWithContext(ctx context.Context) FlexTemplateJobPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlexTemplateJobPtrOutput)
}

type FlexTemplateJobPtrInput interface {
	pulumi.Input

	ToFlexTemplateJobPtrOutput() FlexTemplateJobPtrOutput
	ToFlexTemplateJobPtrOutputWithContext(ctx context.Context) FlexTemplateJobPtrOutput
}

type flexTemplateJobPtrType FlexTemplateJobArgs

func (*flexTemplateJobPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FlexTemplateJob)(nil))
}

func (i *flexTemplateJobPtrType) ToFlexTemplateJobPtrOutput() FlexTemplateJobPtrOutput {
	return i.ToFlexTemplateJobPtrOutputWithContext(context.Background())
}

func (i *flexTemplateJobPtrType) ToFlexTemplateJobPtrOutputWithContext(ctx context.Context) FlexTemplateJobPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlexTemplateJobPtrOutput)
}

// FlexTemplateJobArrayInput is an input type that accepts FlexTemplateJobArray and FlexTemplateJobArrayOutput values.
// You can construct a concrete instance of `FlexTemplateJobArrayInput` via:
//
//          FlexTemplateJobArray{ FlexTemplateJobArgs{...} }
type FlexTemplateJobArrayInput interface {
	pulumi.Input

	ToFlexTemplateJobArrayOutput() FlexTemplateJobArrayOutput
	ToFlexTemplateJobArrayOutputWithContext(context.Context) FlexTemplateJobArrayOutput
}

type FlexTemplateJobArray []FlexTemplateJobInput

func (FlexTemplateJobArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*FlexTemplateJob)(nil))
}

func (i FlexTemplateJobArray) ToFlexTemplateJobArrayOutput() FlexTemplateJobArrayOutput {
	return i.ToFlexTemplateJobArrayOutputWithContext(context.Background())
}

func (i FlexTemplateJobArray) ToFlexTemplateJobArrayOutputWithContext(ctx context.Context) FlexTemplateJobArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlexTemplateJobArrayOutput)
}

// FlexTemplateJobMapInput is an input type that accepts FlexTemplateJobMap and FlexTemplateJobMapOutput values.
// You can construct a concrete instance of `FlexTemplateJobMapInput` via:
//
//          FlexTemplateJobMap{ "key": FlexTemplateJobArgs{...} }
type FlexTemplateJobMapInput interface {
	pulumi.Input

	ToFlexTemplateJobMapOutput() FlexTemplateJobMapOutput
	ToFlexTemplateJobMapOutputWithContext(context.Context) FlexTemplateJobMapOutput
}

type FlexTemplateJobMap map[string]FlexTemplateJobInput

func (FlexTemplateJobMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*FlexTemplateJob)(nil))
}

func (i FlexTemplateJobMap) ToFlexTemplateJobMapOutput() FlexTemplateJobMapOutput {
	return i.ToFlexTemplateJobMapOutputWithContext(context.Background())
}

func (i FlexTemplateJobMap) ToFlexTemplateJobMapOutputWithContext(ctx context.Context) FlexTemplateJobMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FlexTemplateJobMapOutput)
}

type FlexTemplateJobOutput struct {
	*pulumi.OutputState
}

func (FlexTemplateJobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FlexTemplateJob)(nil))
}

func (o FlexTemplateJobOutput) ToFlexTemplateJobOutput() FlexTemplateJobOutput {
	return o
}

func (o FlexTemplateJobOutput) ToFlexTemplateJobOutputWithContext(ctx context.Context) FlexTemplateJobOutput {
	return o
}

func (o FlexTemplateJobOutput) ToFlexTemplateJobPtrOutput() FlexTemplateJobPtrOutput {
	return o.ToFlexTemplateJobPtrOutputWithContext(context.Background())
}

func (o FlexTemplateJobOutput) ToFlexTemplateJobPtrOutputWithContext(ctx context.Context) FlexTemplateJobPtrOutput {
	return o.ApplyT(func(v FlexTemplateJob) *FlexTemplateJob {
		return &v
	}).(FlexTemplateJobPtrOutput)
}

type FlexTemplateJobPtrOutput struct {
	*pulumi.OutputState
}

func (FlexTemplateJobPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FlexTemplateJob)(nil))
}

func (o FlexTemplateJobPtrOutput) ToFlexTemplateJobPtrOutput() FlexTemplateJobPtrOutput {
	return o
}

func (o FlexTemplateJobPtrOutput) ToFlexTemplateJobPtrOutputWithContext(ctx context.Context) FlexTemplateJobPtrOutput {
	return o
}

type FlexTemplateJobArrayOutput struct{ *pulumi.OutputState }

func (FlexTemplateJobArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FlexTemplateJob)(nil))
}

func (o FlexTemplateJobArrayOutput) ToFlexTemplateJobArrayOutput() FlexTemplateJobArrayOutput {
	return o
}

func (o FlexTemplateJobArrayOutput) ToFlexTemplateJobArrayOutputWithContext(ctx context.Context) FlexTemplateJobArrayOutput {
	return o
}

func (o FlexTemplateJobArrayOutput) Index(i pulumi.IntInput) FlexTemplateJobOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FlexTemplateJob {
		return vs[0].([]FlexTemplateJob)[vs[1].(int)]
	}).(FlexTemplateJobOutput)
}

type FlexTemplateJobMapOutput struct{ *pulumi.OutputState }

func (FlexTemplateJobMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FlexTemplateJob)(nil))
}

func (o FlexTemplateJobMapOutput) ToFlexTemplateJobMapOutput() FlexTemplateJobMapOutput {
	return o
}

func (o FlexTemplateJobMapOutput) ToFlexTemplateJobMapOutputWithContext(ctx context.Context) FlexTemplateJobMapOutput {
	return o
}

func (o FlexTemplateJobMapOutput) MapIndex(k pulumi.StringInput) FlexTemplateJobOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FlexTemplateJob {
		return vs[0].(map[string]FlexTemplateJob)[vs[1].(string)]
	}).(FlexTemplateJobOutput)
}

func init() {
	pulumi.RegisterOutputType(FlexTemplateJobOutput{})
	pulumi.RegisterOutputType(FlexTemplateJobPtrOutput{})
	pulumi.RegisterOutputType(FlexTemplateJobArrayOutput{})
	pulumi.RegisterOutputType(FlexTemplateJobMapOutput{})
}
