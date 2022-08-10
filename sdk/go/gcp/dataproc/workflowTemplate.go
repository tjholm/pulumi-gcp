// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dataproc

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A Workflow Template is a reusable workflow configuration. It defines a graph of jobs with information on where to run those jobs.
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
//			_, err := dataproc.NewWorkflowTemplate(ctx, "template", &dataproc.WorkflowTemplateArgs{
//				Jobs: dataproc.WorkflowTemplateJobArray{
//					&dataproc.WorkflowTemplateJobArgs{
//						SparkJob: &dataproc.WorkflowTemplateJobSparkJobArgs{
//							MainClass: pulumi.String("SomeClass"),
//						},
//						StepId: pulumi.String("someJob"),
//					},
//					&dataproc.WorkflowTemplateJobArgs{
//						PrerequisiteStepIds: pulumi.StringArray{
//							pulumi.String("someJob"),
//						},
//						PrestoJob: &dataproc.WorkflowTemplateJobPrestoJobArgs{
//							QueryFileUri: pulumi.String("someuri"),
//						},
//						StepId: pulumi.String("otherJob"),
//					},
//				},
//				Location: pulumi.String("us-central1"),
//				Placement: &dataproc.WorkflowTemplatePlacementArgs{
//					ManagedCluster: &dataproc.WorkflowTemplatePlacementManagedClusterArgs{
//						ClusterName: pulumi.String("my-cluster"),
//						Config: &dataproc.WorkflowTemplatePlacementManagedClusterConfigArgs{
//							GceClusterConfig: &dataproc.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigArgs{
//								Tags: pulumi.StringArray{
//									pulumi.String("foo"),
//									pulumi.String("bar"),
//								},
//								Zone: pulumi.String("us-central1-a"),
//							},
//							MasterConfig: &dataproc.WorkflowTemplatePlacementManagedClusterConfigMasterConfigArgs{
//								DiskConfig: &dataproc.WorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfigArgs{
//									BootDiskSizeGb: pulumi.Int(15),
//									BootDiskType:   pulumi.String("pd-ssd"),
//								},
//								MachineType:  pulumi.String("n1-standard-1"),
//								NumInstances: pulumi.Int(1),
//							},
//							SecondaryWorkerConfig: &dataproc.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigArgs{
//								NumInstances: pulumi.Int(2),
//							},
//							SoftwareConfig: &dataproc.WorkflowTemplatePlacementManagedClusterConfigSoftwareConfigArgs{
//								ImageVersion: pulumi.String("2.0.35-debian10"),
//							},
//							WorkerConfig: &dataproc.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigArgs{
//								DiskConfig: &dataproc.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfigArgs{
//									BootDiskSizeGb: pulumi.Int(10),
//									NumLocalSsds:   pulumi.Int(2),
//								},
//								MachineType:  pulumi.String("n1-standard-2"),
//								NumInstances: pulumi.Int(3),
//							},
//						},
//					},
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
// # WorkflowTemplate can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:dataproc/workflowTemplate:WorkflowTemplate default projects/{{project}}/locations/{{location}}/workflowTemplates/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:dataproc/workflowTemplate:WorkflowTemplate default {{project}}/{{location}}/{{name}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:dataproc/workflowTemplate:WorkflowTemplate default {{location}}/{{name}}
//
// ```
type WorkflowTemplate struct {
	pulumi.CustomResourceState

	// Output only. The time template was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// (Beta only) Optional. Timeout duration for the DAG of jobs. You can use "s", "m", "h", and "d" suffixes for second, minute, hour, and day duration values, respectively. The timeout duration must be from 10 minutes ("10m") to 24 hours ("24h" or "1d"). The timer begins when the first job is submitted. If the workflow is running at the end of the timeout period, any remaining jobs are cancelled, the workflow is ended, and if the workflow was running on a (/dataproc/docs/concepts/workflows/using-workflows#configuring_or_selecting_a_cluster), the cluster is deleted.
	DagTimeout pulumi.StringPtrOutput `pulumi:"dagTimeout"`
	// Required. The Directed Acyclic Graph of Jobs to submit.
	Jobs WorkflowTemplateJobArrayOutput `pulumi:"jobs"`
	// Optional. The labels to associate with this cluster. Label keys must be between 1 and 63 characters long, and must conform to the following PCRE regular expression: {0,63} No more than 32 labels can be associated with a given cluster.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The location for the resource
	Location pulumi.StringOutput `pulumi:"location"`
	// Required. Parameter name. The parameter name is used as the key, and paired with the parameter value, which are passed to the template when the template is instantiated. The name must contain only capital letters (A-Z), numbers (0-9), and underscores (_), and must not start with a number. The maximum length is 40 characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// Optional. Template parameters whose values are substituted into the template. Values for parameters must be provided when the template is instantiated.
	Parameters WorkflowTemplateParameterArrayOutput `pulumi:"parameters"`
	// Required. WorkflowTemplate scheduling information.
	Placement WorkflowTemplatePlacementOutput `pulumi:"placement"`
	// The project for the resource
	Project pulumi.StringOutput `pulumi:"project"`
	// Output only. The time template was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// Optional. Used to perform a consistent read-modify-write. This field should be left blank for a `CreateWorkflowTemplate` request. It is required for an `UpdateWorkflowTemplate` request, and must match the current server version. A typical update template flow would fetch the current template with a `GetWorkflowTemplate` request, which will return the current template with the `version` field filled in with the current server version. The user updates other fields in the template, then returns it as part of the `UpdateWorkflowTemplate` request.
	//
	// Deprecated: version is not useful as a configurable field, and will be removed in the future.
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewWorkflowTemplate registers a new resource with the given unique name, arguments, and options.
func NewWorkflowTemplate(ctx *pulumi.Context,
	name string, args *WorkflowTemplateArgs, opts ...pulumi.ResourceOption) (*WorkflowTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Jobs == nil {
		return nil, errors.New("invalid value for required argument 'Jobs'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.Placement == nil {
		return nil, errors.New("invalid value for required argument 'Placement'")
	}
	var resource WorkflowTemplate
	err := ctx.RegisterResource("gcp:dataproc/workflowTemplate:WorkflowTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkflowTemplate gets an existing WorkflowTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkflowTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkflowTemplateState, opts ...pulumi.ResourceOption) (*WorkflowTemplate, error) {
	var resource WorkflowTemplate
	err := ctx.ReadResource("gcp:dataproc/workflowTemplate:WorkflowTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkflowTemplate resources.
type workflowTemplateState struct {
	// Output only. The time template was created.
	CreateTime *string `pulumi:"createTime"`
	// (Beta only) Optional. Timeout duration for the DAG of jobs. You can use "s", "m", "h", and "d" suffixes for second, minute, hour, and day duration values, respectively. The timeout duration must be from 10 minutes ("10m") to 24 hours ("24h" or "1d"). The timer begins when the first job is submitted. If the workflow is running at the end of the timeout period, any remaining jobs are cancelled, the workflow is ended, and if the workflow was running on a (/dataproc/docs/concepts/workflows/using-workflows#configuring_or_selecting_a_cluster), the cluster is deleted.
	DagTimeout *string `pulumi:"dagTimeout"`
	// Required. The Directed Acyclic Graph of Jobs to submit.
	Jobs []WorkflowTemplateJob `pulumi:"jobs"`
	// Optional. The labels to associate with this cluster. Label keys must be between 1 and 63 characters long, and must conform to the following PCRE regular expression: {0,63} No more than 32 labels can be associated with a given cluster.
	Labels map[string]string `pulumi:"labels"`
	// The location for the resource
	Location *string `pulumi:"location"`
	// Required. Parameter name. The parameter name is used as the key, and paired with the parameter value, which are passed to the template when the template is instantiated. The name must contain only capital letters (A-Z), numbers (0-9), and underscores (_), and must not start with a number. The maximum length is 40 characters.
	Name *string `pulumi:"name"`
	// Optional. Template parameters whose values are substituted into the template. Values for parameters must be provided when the template is instantiated.
	Parameters []WorkflowTemplateParameter `pulumi:"parameters"`
	// Required. WorkflowTemplate scheduling information.
	Placement *WorkflowTemplatePlacement `pulumi:"placement"`
	// The project for the resource
	Project *string `pulumi:"project"`
	// Output only. The time template was last updated.
	UpdateTime *string `pulumi:"updateTime"`
	// Optional. Used to perform a consistent read-modify-write. This field should be left blank for a `CreateWorkflowTemplate` request. It is required for an `UpdateWorkflowTemplate` request, and must match the current server version. A typical update template flow would fetch the current template with a `GetWorkflowTemplate` request, which will return the current template with the `version` field filled in with the current server version. The user updates other fields in the template, then returns it as part of the `UpdateWorkflowTemplate` request.
	//
	// Deprecated: version is not useful as a configurable field, and will be removed in the future.
	Version *int `pulumi:"version"`
}

type WorkflowTemplateState struct {
	// Output only. The time template was created.
	CreateTime pulumi.StringPtrInput
	// (Beta only) Optional. Timeout duration for the DAG of jobs. You can use "s", "m", "h", and "d" suffixes for second, minute, hour, and day duration values, respectively. The timeout duration must be from 10 minutes ("10m") to 24 hours ("24h" or "1d"). The timer begins when the first job is submitted. If the workflow is running at the end of the timeout period, any remaining jobs are cancelled, the workflow is ended, and if the workflow was running on a (/dataproc/docs/concepts/workflows/using-workflows#configuring_or_selecting_a_cluster), the cluster is deleted.
	DagTimeout pulumi.StringPtrInput
	// Required. The Directed Acyclic Graph of Jobs to submit.
	Jobs WorkflowTemplateJobArrayInput
	// Optional. The labels to associate with this cluster. Label keys must be between 1 and 63 characters long, and must conform to the following PCRE regular expression: {0,63} No more than 32 labels can be associated with a given cluster.
	Labels pulumi.StringMapInput
	// The location for the resource
	Location pulumi.StringPtrInput
	// Required. Parameter name. The parameter name is used as the key, and paired with the parameter value, which are passed to the template when the template is instantiated. The name must contain only capital letters (A-Z), numbers (0-9), and underscores (_), and must not start with a number. The maximum length is 40 characters.
	Name pulumi.StringPtrInput
	// Optional. Template parameters whose values are substituted into the template. Values for parameters must be provided when the template is instantiated.
	Parameters WorkflowTemplateParameterArrayInput
	// Required. WorkflowTemplate scheduling information.
	Placement WorkflowTemplatePlacementPtrInput
	// The project for the resource
	Project pulumi.StringPtrInput
	// Output only. The time template was last updated.
	UpdateTime pulumi.StringPtrInput
	// Optional. Used to perform a consistent read-modify-write. This field should be left blank for a `CreateWorkflowTemplate` request. It is required for an `UpdateWorkflowTemplate` request, and must match the current server version. A typical update template flow would fetch the current template with a `GetWorkflowTemplate` request, which will return the current template with the `version` field filled in with the current server version. The user updates other fields in the template, then returns it as part of the `UpdateWorkflowTemplate` request.
	//
	// Deprecated: version is not useful as a configurable field, and will be removed in the future.
	Version pulumi.IntPtrInput
}

func (WorkflowTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowTemplateState)(nil)).Elem()
}

type workflowTemplateArgs struct {
	// (Beta only) Optional. Timeout duration for the DAG of jobs. You can use "s", "m", "h", and "d" suffixes for second, minute, hour, and day duration values, respectively. The timeout duration must be from 10 minutes ("10m") to 24 hours ("24h" or "1d"). The timer begins when the first job is submitted. If the workflow is running at the end of the timeout period, any remaining jobs are cancelled, the workflow is ended, and if the workflow was running on a (/dataproc/docs/concepts/workflows/using-workflows#configuring_or_selecting_a_cluster), the cluster is deleted.
	DagTimeout *string `pulumi:"dagTimeout"`
	// Required. The Directed Acyclic Graph of Jobs to submit.
	Jobs []WorkflowTemplateJob `pulumi:"jobs"`
	// Optional. The labels to associate with this cluster. Label keys must be between 1 and 63 characters long, and must conform to the following PCRE regular expression: {0,63} No more than 32 labels can be associated with a given cluster.
	Labels map[string]string `pulumi:"labels"`
	// The location for the resource
	Location string `pulumi:"location"`
	// Required. Parameter name. The parameter name is used as the key, and paired with the parameter value, which are passed to the template when the template is instantiated. The name must contain only capital letters (A-Z), numbers (0-9), and underscores (_), and must not start with a number. The maximum length is 40 characters.
	Name *string `pulumi:"name"`
	// Optional. Template parameters whose values are substituted into the template. Values for parameters must be provided when the template is instantiated.
	Parameters []WorkflowTemplateParameter `pulumi:"parameters"`
	// Required. WorkflowTemplate scheduling information.
	Placement WorkflowTemplatePlacement `pulumi:"placement"`
	// The project for the resource
	Project *string `pulumi:"project"`
	// Optional. Used to perform a consistent read-modify-write. This field should be left blank for a `CreateWorkflowTemplate` request. It is required for an `UpdateWorkflowTemplate` request, and must match the current server version. A typical update template flow would fetch the current template with a `GetWorkflowTemplate` request, which will return the current template with the `version` field filled in with the current server version. The user updates other fields in the template, then returns it as part of the `UpdateWorkflowTemplate` request.
	//
	// Deprecated: version is not useful as a configurable field, and will be removed in the future.
	Version *int `pulumi:"version"`
}

// The set of arguments for constructing a WorkflowTemplate resource.
type WorkflowTemplateArgs struct {
	// (Beta only) Optional. Timeout duration for the DAG of jobs. You can use "s", "m", "h", and "d" suffixes for second, minute, hour, and day duration values, respectively. The timeout duration must be from 10 minutes ("10m") to 24 hours ("24h" or "1d"). The timer begins when the first job is submitted. If the workflow is running at the end of the timeout period, any remaining jobs are cancelled, the workflow is ended, and if the workflow was running on a (/dataproc/docs/concepts/workflows/using-workflows#configuring_or_selecting_a_cluster), the cluster is deleted.
	DagTimeout pulumi.StringPtrInput
	// Required. The Directed Acyclic Graph of Jobs to submit.
	Jobs WorkflowTemplateJobArrayInput
	// Optional. The labels to associate with this cluster. Label keys must be between 1 and 63 characters long, and must conform to the following PCRE regular expression: {0,63} No more than 32 labels can be associated with a given cluster.
	Labels pulumi.StringMapInput
	// The location for the resource
	Location pulumi.StringInput
	// Required. Parameter name. The parameter name is used as the key, and paired with the parameter value, which are passed to the template when the template is instantiated. The name must contain only capital letters (A-Z), numbers (0-9), and underscores (_), and must not start with a number. The maximum length is 40 characters.
	Name pulumi.StringPtrInput
	// Optional. Template parameters whose values are substituted into the template. Values for parameters must be provided when the template is instantiated.
	Parameters WorkflowTemplateParameterArrayInput
	// Required. WorkflowTemplate scheduling information.
	Placement WorkflowTemplatePlacementInput
	// The project for the resource
	Project pulumi.StringPtrInput
	// Optional. Used to perform a consistent read-modify-write. This field should be left blank for a `CreateWorkflowTemplate` request. It is required for an `UpdateWorkflowTemplate` request, and must match the current server version. A typical update template flow would fetch the current template with a `GetWorkflowTemplate` request, which will return the current template with the `version` field filled in with the current server version. The user updates other fields in the template, then returns it as part of the `UpdateWorkflowTemplate` request.
	//
	// Deprecated: version is not useful as a configurable field, and will be removed in the future.
	Version pulumi.IntPtrInput
}

func (WorkflowTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowTemplateArgs)(nil)).Elem()
}

type WorkflowTemplateInput interface {
	pulumi.Input

	ToWorkflowTemplateOutput() WorkflowTemplateOutput
	ToWorkflowTemplateOutputWithContext(ctx context.Context) WorkflowTemplateOutput
}

func (*WorkflowTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkflowTemplate)(nil)).Elem()
}

func (i *WorkflowTemplate) ToWorkflowTemplateOutput() WorkflowTemplateOutput {
	return i.ToWorkflowTemplateOutputWithContext(context.Background())
}

func (i *WorkflowTemplate) ToWorkflowTemplateOutputWithContext(ctx context.Context) WorkflowTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowTemplateOutput)
}

// WorkflowTemplateArrayInput is an input type that accepts WorkflowTemplateArray and WorkflowTemplateArrayOutput values.
// You can construct a concrete instance of `WorkflowTemplateArrayInput` via:
//
//	WorkflowTemplateArray{ WorkflowTemplateArgs{...} }
type WorkflowTemplateArrayInput interface {
	pulumi.Input

	ToWorkflowTemplateArrayOutput() WorkflowTemplateArrayOutput
	ToWorkflowTemplateArrayOutputWithContext(context.Context) WorkflowTemplateArrayOutput
}

type WorkflowTemplateArray []WorkflowTemplateInput

func (WorkflowTemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WorkflowTemplate)(nil)).Elem()
}

func (i WorkflowTemplateArray) ToWorkflowTemplateArrayOutput() WorkflowTemplateArrayOutput {
	return i.ToWorkflowTemplateArrayOutputWithContext(context.Background())
}

func (i WorkflowTemplateArray) ToWorkflowTemplateArrayOutputWithContext(ctx context.Context) WorkflowTemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowTemplateArrayOutput)
}

// WorkflowTemplateMapInput is an input type that accepts WorkflowTemplateMap and WorkflowTemplateMapOutput values.
// You can construct a concrete instance of `WorkflowTemplateMapInput` via:
//
//	WorkflowTemplateMap{ "key": WorkflowTemplateArgs{...} }
type WorkflowTemplateMapInput interface {
	pulumi.Input

	ToWorkflowTemplateMapOutput() WorkflowTemplateMapOutput
	ToWorkflowTemplateMapOutputWithContext(context.Context) WorkflowTemplateMapOutput
}

type WorkflowTemplateMap map[string]WorkflowTemplateInput

func (WorkflowTemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WorkflowTemplate)(nil)).Elem()
}

func (i WorkflowTemplateMap) ToWorkflowTemplateMapOutput() WorkflowTemplateMapOutput {
	return i.ToWorkflowTemplateMapOutputWithContext(context.Background())
}

func (i WorkflowTemplateMap) ToWorkflowTemplateMapOutputWithContext(ctx context.Context) WorkflowTemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowTemplateMapOutput)
}

type WorkflowTemplateOutput struct{ *pulumi.OutputState }

func (WorkflowTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkflowTemplate)(nil)).Elem()
}

func (o WorkflowTemplateOutput) ToWorkflowTemplateOutput() WorkflowTemplateOutput {
	return o
}

func (o WorkflowTemplateOutput) ToWorkflowTemplateOutputWithContext(ctx context.Context) WorkflowTemplateOutput {
	return o
}

// Output only. The time template was created.
func (o WorkflowTemplateOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowTemplate) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// (Beta only) Optional. Timeout duration for the DAG of jobs. You can use "s", "m", "h", and "d" suffixes for second, minute, hour, and day duration values, respectively. The timeout duration must be from 10 minutes ("10m") to 24 hours ("24h" or "1d"). The timer begins when the first job is submitted. If the workflow is running at the end of the timeout period, any remaining jobs are cancelled, the workflow is ended, and if the workflow was running on a (/dataproc/docs/concepts/workflows/using-workflows#configuring_or_selecting_a_cluster), the cluster is deleted.
func (o WorkflowTemplateOutput) DagTimeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkflowTemplate) pulumi.StringPtrOutput { return v.DagTimeout }).(pulumi.StringPtrOutput)
}

// Required. The Directed Acyclic Graph of Jobs to submit.
func (o WorkflowTemplateOutput) Jobs() WorkflowTemplateJobArrayOutput {
	return o.ApplyT(func(v *WorkflowTemplate) WorkflowTemplateJobArrayOutput { return v.Jobs }).(WorkflowTemplateJobArrayOutput)
}

// Optional. The labels to associate with this cluster. Label keys must be between 1 and 63 characters long, and must conform to the following PCRE regular expression: {0,63} No more than 32 labels can be associated with a given cluster.
func (o WorkflowTemplateOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *WorkflowTemplate) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// The location for the resource
func (o WorkflowTemplateOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowTemplate) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Required. Parameter name. The parameter name is used as the key, and paired with the parameter value, which are passed to the template when the template is instantiated. The name must contain only capital letters (A-Z), numbers (0-9), and underscores (_), and must not start with a number. The maximum length is 40 characters.
func (o WorkflowTemplateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowTemplate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional. Template parameters whose values are substituted into the template. Values for parameters must be provided when the template is instantiated.
func (o WorkflowTemplateOutput) Parameters() WorkflowTemplateParameterArrayOutput {
	return o.ApplyT(func(v *WorkflowTemplate) WorkflowTemplateParameterArrayOutput { return v.Parameters }).(WorkflowTemplateParameterArrayOutput)
}

// Required. WorkflowTemplate scheduling information.
func (o WorkflowTemplateOutput) Placement() WorkflowTemplatePlacementOutput {
	return o.ApplyT(func(v *WorkflowTemplate) WorkflowTemplatePlacementOutput { return v.Placement }).(WorkflowTemplatePlacementOutput)
}

// The project for the resource
func (o WorkflowTemplateOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowTemplate) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Output only. The time template was last updated.
func (o WorkflowTemplateOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowTemplate) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// Optional. Used to perform a consistent read-modify-write. This field should be left blank for a `CreateWorkflowTemplate` request. It is required for an `UpdateWorkflowTemplate` request, and must match the current server version. A typical update template flow would fetch the current template with a `GetWorkflowTemplate` request, which will return the current template with the `version` field filled in with the current server version. The user updates other fields in the template, then returns it as part of the `UpdateWorkflowTemplate` request.
//
// Deprecated: version is not useful as a configurable field, and will be removed in the future.
func (o WorkflowTemplateOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v *WorkflowTemplate) pulumi.IntOutput { return v.Version }).(pulumi.IntOutput)
}

type WorkflowTemplateArrayOutput struct{ *pulumi.OutputState }

func (WorkflowTemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*WorkflowTemplate)(nil)).Elem()
}

func (o WorkflowTemplateArrayOutput) ToWorkflowTemplateArrayOutput() WorkflowTemplateArrayOutput {
	return o
}

func (o WorkflowTemplateArrayOutput) ToWorkflowTemplateArrayOutputWithContext(ctx context.Context) WorkflowTemplateArrayOutput {
	return o
}

func (o WorkflowTemplateArrayOutput) Index(i pulumi.IntInput) WorkflowTemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *WorkflowTemplate {
		return vs[0].([]*WorkflowTemplate)[vs[1].(int)]
	}).(WorkflowTemplateOutput)
}

type WorkflowTemplateMapOutput struct{ *pulumi.OutputState }

func (WorkflowTemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*WorkflowTemplate)(nil)).Elem()
}

func (o WorkflowTemplateMapOutput) ToWorkflowTemplateMapOutput() WorkflowTemplateMapOutput {
	return o
}

func (o WorkflowTemplateMapOutput) ToWorkflowTemplateMapOutputWithContext(ctx context.Context) WorkflowTemplateMapOutput {
	return o
}

func (o WorkflowTemplateMapOutput) MapIndex(k pulumi.StringInput) WorkflowTemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *WorkflowTemplate {
		return vs[0].(map[string]*WorkflowTemplate)[vs[1].(string)]
	}).(WorkflowTemplateOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowTemplateInput)(nil)).Elem(), &WorkflowTemplate{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowTemplateArrayInput)(nil)).Elem(), WorkflowTemplateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowTemplateMapInput)(nil)).Elem(), WorkflowTemplateMap{})
	pulumi.RegisterOutputType(WorkflowTemplateOutput{})
	pulumi.RegisterOutputType(WorkflowTemplateArrayOutput{})
	pulumi.RegisterOutputType(WorkflowTemplateMapOutput{})
}
