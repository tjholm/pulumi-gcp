// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package dataloss

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A job trigger configuration.
//
// To get more information about JobTrigger, see:
//
// * [API documentation](https://cloud.google.com/dlp/docs/reference/rest/v2/projects.jobTriggers)
// * How-to Guides
//     * [Official Documentation](https://cloud.google.com/dlp/docs/creating-job-triggers)
//
// ## Example Usage
// ### Dlp Job Trigger Basic
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v5/go/gcp/dataloss"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := dataloss.NewPreventionJobTrigger(ctx, "basic", &dataloss.PreventionJobTriggerArgs{
// 			Description: pulumi.String("Description"),
// 			DisplayName: pulumi.String("Displayname"),
// 			InspectJob: &dataloss.PreventionJobTriggerInspectJobArgs{
// 				Actions: dataloss.PreventionJobTriggerInspectJobActionArray{
// 					&dataloss.PreventionJobTriggerInspectJobActionArgs{
// 						SaveFindings: &dataloss.PreventionJobTriggerInspectJobActionSaveFindingsArgs{
// 							OutputConfig: &dataloss.PreventionJobTriggerInspectJobActionSaveFindingsOutputConfigArgs{
// 								Table: &dataloss.PreventionJobTriggerInspectJobActionSaveFindingsOutputConfigTableArgs{
// 									DatasetId: pulumi.String("asdf"),
// 									ProjectId: pulumi.String("asdf"),
// 								},
// 							},
// 						},
// 					},
// 				},
// 				InspectTemplateName: pulumi.String("fake"),
// 				StorageConfig: &dataloss.PreventionJobTriggerInspectJobStorageConfigArgs{
// 					CloudStorageOptions: &dataloss.PreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsArgs{
// 						FileSet: &dataloss.PreventionJobTriggerInspectJobStorageConfigCloudStorageOptionsFileSetArgs{
// 							Url: pulumi.String("gs://mybucket/directory/"),
// 						},
// 					},
// 				},
// 			},
// 			Parent: pulumi.String("projects/my-project-name"),
// 			Triggers: dataloss.PreventionJobTriggerTriggerArray{
// 				&dataloss.PreventionJobTriggerTriggerArgs{
// 					Schedule: &dataloss.PreventionJobTriggerTriggerScheduleArgs{
// 						RecurrencePeriodDuration: pulumi.String("86400s"),
// 					},
// 				},
// 			},
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
// JobTrigger can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:dataloss/preventionJobTrigger:PreventionJobTrigger default {{parent}}/jobTriggers/{{name}}
// ```
//
// ```sh
//  $ pulumi import gcp:dataloss/preventionJobTrigger:PreventionJobTrigger default {{parent}}/{{name}}
// ```
type PreventionJobTrigger struct {
	pulumi.CustomResourceState

	// A description of the job trigger.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// User set display name of the job trigger.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Controls what and how to inspect for findings.
	// Structure is documented below.
	InspectJob PreventionJobTriggerInspectJobPtrOutput `pulumi:"inspectJob"`
	// The timestamp of the last time this trigger executed.
	LastRunTime pulumi.StringOutput `pulumi:"lastRunTime"`
	// The name of the Datastore kind.
	Name pulumi.StringOutput `pulumi:"name"`
	// The parent of the trigger, either in the format `projects/{{project}}`
	// or `projects/{{project}}/locations/{{location}}`
	Parent pulumi.StringOutput `pulumi:"parent"`
	// Whether the trigger is currently active.
	// Default value is `HEALTHY`.
	// Possible values are `PAUSED`, `HEALTHY`, and `CANCELLED`.
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// What event needs to occur for a new job to be started.
	// Structure is documented below.
	Triggers PreventionJobTriggerTriggerArrayOutput `pulumi:"triggers"`
}

// NewPreventionJobTrigger registers a new resource with the given unique name, arguments, and options.
func NewPreventionJobTrigger(ctx *pulumi.Context,
	name string, args *PreventionJobTriggerArgs, opts ...pulumi.ResourceOption) (*PreventionJobTrigger, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Parent == nil {
		return nil, errors.New("invalid value for required argument 'Parent'")
	}
	if args.Triggers == nil {
		return nil, errors.New("invalid value for required argument 'Triggers'")
	}
	var resource PreventionJobTrigger
	err := ctx.RegisterResource("gcp:dataloss/preventionJobTrigger:PreventionJobTrigger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPreventionJobTrigger gets an existing PreventionJobTrigger resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPreventionJobTrigger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PreventionJobTriggerState, opts ...pulumi.ResourceOption) (*PreventionJobTrigger, error) {
	var resource PreventionJobTrigger
	err := ctx.ReadResource("gcp:dataloss/preventionJobTrigger:PreventionJobTrigger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PreventionJobTrigger resources.
type preventionJobTriggerState struct {
	// A description of the job trigger.
	Description *string `pulumi:"description"`
	// User set display name of the job trigger.
	DisplayName *string `pulumi:"displayName"`
	// Controls what and how to inspect for findings.
	// Structure is documented below.
	InspectJob *PreventionJobTriggerInspectJob `pulumi:"inspectJob"`
	// The timestamp of the last time this trigger executed.
	LastRunTime *string `pulumi:"lastRunTime"`
	// The name of the Datastore kind.
	Name *string `pulumi:"name"`
	// The parent of the trigger, either in the format `projects/{{project}}`
	// or `projects/{{project}}/locations/{{location}}`
	Parent *string `pulumi:"parent"`
	// Whether the trigger is currently active.
	// Default value is `HEALTHY`.
	// Possible values are `PAUSED`, `HEALTHY`, and `CANCELLED`.
	Status *string `pulumi:"status"`
	// What event needs to occur for a new job to be started.
	// Structure is documented below.
	Triggers []PreventionJobTriggerTrigger `pulumi:"triggers"`
}

type PreventionJobTriggerState struct {
	// A description of the job trigger.
	Description pulumi.StringPtrInput
	// User set display name of the job trigger.
	DisplayName pulumi.StringPtrInput
	// Controls what and how to inspect for findings.
	// Structure is documented below.
	InspectJob PreventionJobTriggerInspectJobPtrInput
	// The timestamp of the last time this trigger executed.
	LastRunTime pulumi.StringPtrInput
	// The name of the Datastore kind.
	Name pulumi.StringPtrInput
	// The parent of the trigger, either in the format `projects/{{project}}`
	// or `projects/{{project}}/locations/{{location}}`
	Parent pulumi.StringPtrInput
	// Whether the trigger is currently active.
	// Default value is `HEALTHY`.
	// Possible values are `PAUSED`, `HEALTHY`, and `CANCELLED`.
	Status pulumi.StringPtrInput
	// What event needs to occur for a new job to be started.
	// Structure is documented below.
	Triggers PreventionJobTriggerTriggerArrayInput
}

func (PreventionJobTriggerState) ElementType() reflect.Type {
	return reflect.TypeOf((*preventionJobTriggerState)(nil)).Elem()
}

type preventionJobTriggerArgs struct {
	// A description of the job trigger.
	Description *string `pulumi:"description"`
	// User set display name of the job trigger.
	DisplayName *string `pulumi:"displayName"`
	// Controls what and how to inspect for findings.
	// Structure is documented below.
	InspectJob *PreventionJobTriggerInspectJob `pulumi:"inspectJob"`
	// The parent of the trigger, either in the format `projects/{{project}}`
	// or `projects/{{project}}/locations/{{location}}`
	Parent string `pulumi:"parent"`
	// Whether the trigger is currently active.
	// Default value is `HEALTHY`.
	// Possible values are `PAUSED`, `HEALTHY`, and `CANCELLED`.
	Status *string `pulumi:"status"`
	// What event needs to occur for a new job to be started.
	// Structure is documented below.
	Triggers []PreventionJobTriggerTrigger `pulumi:"triggers"`
}

// The set of arguments for constructing a PreventionJobTrigger resource.
type PreventionJobTriggerArgs struct {
	// A description of the job trigger.
	Description pulumi.StringPtrInput
	// User set display name of the job trigger.
	DisplayName pulumi.StringPtrInput
	// Controls what and how to inspect for findings.
	// Structure is documented below.
	InspectJob PreventionJobTriggerInspectJobPtrInput
	// The parent of the trigger, either in the format `projects/{{project}}`
	// or `projects/{{project}}/locations/{{location}}`
	Parent pulumi.StringInput
	// Whether the trigger is currently active.
	// Default value is `HEALTHY`.
	// Possible values are `PAUSED`, `HEALTHY`, and `CANCELLED`.
	Status pulumi.StringPtrInput
	// What event needs to occur for a new job to be started.
	// Structure is documented below.
	Triggers PreventionJobTriggerTriggerArrayInput
}

func (PreventionJobTriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*preventionJobTriggerArgs)(nil)).Elem()
}

type PreventionJobTriggerInput interface {
	pulumi.Input

	ToPreventionJobTriggerOutput() PreventionJobTriggerOutput
	ToPreventionJobTriggerOutputWithContext(ctx context.Context) PreventionJobTriggerOutput
}

func (*PreventionJobTrigger) ElementType() reflect.Type {
	return reflect.TypeOf((*PreventionJobTrigger)(nil))
}

func (i *PreventionJobTrigger) ToPreventionJobTriggerOutput() PreventionJobTriggerOutput {
	return i.ToPreventionJobTriggerOutputWithContext(context.Background())
}

func (i *PreventionJobTrigger) ToPreventionJobTriggerOutputWithContext(ctx context.Context) PreventionJobTriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PreventionJobTriggerOutput)
}

func (i *PreventionJobTrigger) ToPreventionJobTriggerPtrOutput() PreventionJobTriggerPtrOutput {
	return i.ToPreventionJobTriggerPtrOutputWithContext(context.Background())
}

func (i *PreventionJobTrigger) ToPreventionJobTriggerPtrOutputWithContext(ctx context.Context) PreventionJobTriggerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PreventionJobTriggerPtrOutput)
}

type PreventionJobTriggerPtrInput interface {
	pulumi.Input

	ToPreventionJobTriggerPtrOutput() PreventionJobTriggerPtrOutput
	ToPreventionJobTriggerPtrOutputWithContext(ctx context.Context) PreventionJobTriggerPtrOutput
}

type preventionJobTriggerPtrType PreventionJobTriggerArgs

func (*preventionJobTriggerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PreventionJobTrigger)(nil))
}

func (i *preventionJobTriggerPtrType) ToPreventionJobTriggerPtrOutput() PreventionJobTriggerPtrOutput {
	return i.ToPreventionJobTriggerPtrOutputWithContext(context.Background())
}

func (i *preventionJobTriggerPtrType) ToPreventionJobTriggerPtrOutputWithContext(ctx context.Context) PreventionJobTriggerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PreventionJobTriggerPtrOutput)
}

// PreventionJobTriggerArrayInput is an input type that accepts PreventionJobTriggerArray and PreventionJobTriggerArrayOutput values.
// You can construct a concrete instance of `PreventionJobTriggerArrayInput` via:
//
//          PreventionJobTriggerArray{ PreventionJobTriggerArgs{...} }
type PreventionJobTriggerArrayInput interface {
	pulumi.Input

	ToPreventionJobTriggerArrayOutput() PreventionJobTriggerArrayOutput
	ToPreventionJobTriggerArrayOutputWithContext(context.Context) PreventionJobTriggerArrayOutput
}

type PreventionJobTriggerArray []PreventionJobTriggerInput

func (PreventionJobTriggerArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*PreventionJobTrigger)(nil))
}

func (i PreventionJobTriggerArray) ToPreventionJobTriggerArrayOutput() PreventionJobTriggerArrayOutput {
	return i.ToPreventionJobTriggerArrayOutputWithContext(context.Background())
}

func (i PreventionJobTriggerArray) ToPreventionJobTriggerArrayOutputWithContext(ctx context.Context) PreventionJobTriggerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PreventionJobTriggerArrayOutput)
}

// PreventionJobTriggerMapInput is an input type that accepts PreventionJobTriggerMap and PreventionJobTriggerMapOutput values.
// You can construct a concrete instance of `PreventionJobTriggerMapInput` via:
//
//          PreventionJobTriggerMap{ "key": PreventionJobTriggerArgs{...} }
type PreventionJobTriggerMapInput interface {
	pulumi.Input

	ToPreventionJobTriggerMapOutput() PreventionJobTriggerMapOutput
	ToPreventionJobTriggerMapOutputWithContext(context.Context) PreventionJobTriggerMapOutput
}

type PreventionJobTriggerMap map[string]PreventionJobTriggerInput

func (PreventionJobTriggerMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*PreventionJobTrigger)(nil))
}

func (i PreventionJobTriggerMap) ToPreventionJobTriggerMapOutput() PreventionJobTriggerMapOutput {
	return i.ToPreventionJobTriggerMapOutputWithContext(context.Background())
}

func (i PreventionJobTriggerMap) ToPreventionJobTriggerMapOutputWithContext(ctx context.Context) PreventionJobTriggerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PreventionJobTriggerMapOutput)
}

type PreventionJobTriggerOutput struct {
	*pulumi.OutputState
}

func (PreventionJobTriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PreventionJobTrigger)(nil))
}

func (o PreventionJobTriggerOutput) ToPreventionJobTriggerOutput() PreventionJobTriggerOutput {
	return o
}

func (o PreventionJobTriggerOutput) ToPreventionJobTriggerOutputWithContext(ctx context.Context) PreventionJobTriggerOutput {
	return o
}

func (o PreventionJobTriggerOutput) ToPreventionJobTriggerPtrOutput() PreventionJobTriggerPtrOutput {
	return o.ToPreventionJobTriggerPtrOutputWithContext(context.Background())
}

func (o PreventionJobTriggerOutput) ToPreventionJobTriggerPtrOutputWithContext(ctx context.Context) PreventionJobTriggerPtrOutput {
	return o.ApplyT(func(v PreventionJobTrigger) *PreventionJobTrigger {
		return &v
	}).(PreventionJobTriggerPtrOutput)
}

type PreventionJobTriggerPtrOutput struct {
	*pulumi.OutputState
}

func (PreventionJobTriggerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PreventionJobTrigger)(nil))
}

func (o PreventionJobTriggerPtrOutput) ToPreventionJobTriggerPtrOutput() PreventionJobTriggerPtrOutput {
	return o
}

func (o PreventionJobTriggerPtrOutput) ToPreventionJobTriggerPtrOutputWithContext(ctx context.Context) PreventionJobTriggerPtrOutput {
	return o
}

type PreventionJobTriggerArrayOutput struct{ *pulumi.OutputState }

func (PreventionJobTriggerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PreventionJobTrigger)(nil))
}

func (o PreventionJobTriggerArrayOutput) ToPreventionJobTriggerArrayOutput() PreventionJobTriggerArrayOutput {
	return o
}

func (o PreventionJobTriggerArrayOutput) ToPreventionJobTriggerArrayOutputWithContext(ctx context.Context) PreventionJobTriggerArrayOutput {
	return o
}

func (o PreventionJobTriggerArrayOutput) Index(i pulumi.IntInput) PreventionJobTriggerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PreventionJobTrigger {
		return vs[0].([]PreventionJobTrigger)[vs[1].(int)]
	}).(PreventionJobTriggerOutput)
}

type PreventionJobTriggerMapOutput struct{ *pulumi.OutputState }

func (PreventionJobTriggerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]PreventionJobTrigger)(nil))
}

func (o PreventionJobTriggerMapOutput) ToPreventionJobTriggerMapOutput() PreventionJobTriggerMapOutput {
	return o
}

func (o PreventionJobTriggerMapOutput) ToPreventionJobTriggerMapOutputWithContext(ctx context.Context) PreventionJobTriggerMapOutput {
	return o
}

func (o PreventionJobTriggerMapOutput) MapIndex(k pulumi.StringInput) PreventionJobTriggerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) PreventionJobTrigger {
		return vs[0].(map[string]PreventionJobTrigger)[vs[1].(string)]
	}).(PreventionJobTriggerOutput)
}

func init() {
	pulumi.RegisterOutputType(PreventionJobTriggerOutput{})
	pulumi.RegisterOutputType(PreventionJobTriggerPtrOutput{})
	pulumi.RegisterOutputType(PreventionJobTriggerArrayOutput{})
	pulumi.RegisterOutputType(PreventionJobTriggerMapOutput{})
}
