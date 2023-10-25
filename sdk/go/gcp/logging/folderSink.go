// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package logging

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Manages a folder-level logging sink. For more information see:
// * [API documentation](https://cloud.google.com/logging/docs/reference/v2/rest/v2/folders.sinks)
// * How-to Guides
//   - [Exporting Logs](https://cloud.google.com/logging/docs/export)
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
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/logging"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/projects"
//	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/storage"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := storage.NewBucket(ctx, "log-bucket", &storage.BucketArgs{
//				Location: pulumi.String("US"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = organizations.NewFolder(ctx, "my-folder", &organizations.FolderArgs{
//				DisplayName: pulumi.String("My folder"),
//				Parent:      pulumi.String("organizations/123456"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = logging.NewFolderSink(ctx, "my-sink", &logging.FolderSinkArgs{
//				Description: pulumi.String("some explanation on what this is"),
//				Folder:      my_folder.Name,
//				Destination: log_bucket.Name.ApplyT(func(name string) (string, error) {
//					return fmt.Sprintf("storage.googleapis.com/%v", name), nil
//				}).(pulumi.StringOutput),
//				Filter: pulumi.String("resource.type = gce_instance AND severity >= WARNING"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = projects.NewIAMBinding(ctx, "log-writer", &projects.IAMBindingArgs{
//				Project: pulumi.String("your-project-id"),
//				Role:    pulumi.String("roles/storage.objectCreator"),
//				Members: pulumi.StringArray{
//					my_sink.WriterIdentity,
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
// Folder-level logging sinks can be imported using this format:
//
// ```sh
//
//	$ pulumi import gcp:logging/folderSink:FolderSink my_sink folders/{{folder_id}}/sinks/{{name}}
//
// ```
type FolderSink struct {
	pulumi.CustomResourceState

	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions FolderSinkBigqueryOptionsOutput `pulumi:"bigqueryOptions"`
	// A description of this sink. The maximum length of the description is 8000 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The destination of the sink (or, in other words, where logs are written to). Can be a
	// Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket. Examples:
	//
	// - `storage.googleapis.com/[GCS_BUCKET]`
	// - `bigquery.googleapis.com/projects/[PROJECT_ID]/datasets/[DATASET]`
	// - `pubsub.googleapis.com/projects/[PROJECT_ID]/topics/[TOPIC_ID]`
	// - `logging.googleapis.com/projects/[PROJECT_ID]]/locations/global/buckets/[BUCKET_ID]`
	//
	// The writer associated with the sink must have access to write to the above resource.
	Destination pulumi.StringOutput `pulumi:"destination"`
	// If set to True, then this sink is disabled and it does not export any log entries.
	Disabled pulumi.BoolPtrOutput `pulumi:"disabled"`
	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both `filter` and one of `exclusions.filter`, it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
	Exclusions FolderSinkExclusionArrayOutput `pulumi:"exclusions"`
	// The filter to apply when exporting logs. Only log entries that match the filter are exported.
	// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
	// write a filter.
	Filter pulumi.StringPtrOutput `pulumi:"filter"`
	// The folder to be exported to the sink. Note that either `[FOLDER_ID]` or `folders/[FOLDER_ID]` is
	// accepted.
	Folder pulumi.StringOutput `pulumi:"folder"`
	// Whether or not to include children folders in the sink export. If true, logs
	// associated with child projects are also exported; otherwise only logs relating to the provided folder are included.
	IncludeChildren pulumi.BoolPtrOutput `pulumi:"includeChildren"`
	// The name of the logging sink.
	Name pulumi.StringOutput `pulumi:"name"`
	// The identity associated with this sink. This identity must be granted write access to the
	// configured `destination`.
	WriterIdentity pulumi.StringOutput `pulumi:"writerIdentity"`
}

// NewFolderSink registers a new resource with the given unique name, arguments, and options.
func NewFolderSink(ctx *pulumi.Context,
	name string, args *FolderSinkArgs, opts ...pulumi.ResourceOption) (*FolderSink, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Destination == nil {
		return nil, errors.New("invalid value for required argument 'Destination'")
	}
	if args.Folder == nil {
		return nil, errors.New("invalid value for required argument 'Folder'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource FolderSink
	err := ctx.RegisterResource("gcp:logging/folderSink:FolderSink", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFolderSink gets an existing FolderSink resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFolderSink(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FolderSinkState, opts ...pulumi.ResourceOption) (*FolderSink, error) {
	var resource FolderSink
	err := ctx.ReadResource("gcp:logging/folderSink:FolderSink", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FolderSink resources.
type folderSinkState struct {
	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions *FolderSinkBigqueryOptions `pulumi:"bigqueryOptions"`
	// A description of this sink. The maximum length of the description is 8000 characters.
	Description *string `pulumi:"description"`
	// The destination of the sink (or, in other words, where logs are written to). Can be a
	// Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket. Examples:
	//
	// - `storage.googleapis.com/[GCS_BUCKET]`
	// - `bigquery.googleapis.com/projects/[PROJECT_ID]/datasets/[DATASET]`
	// - `pubsub.googleapis.com/projects/[PROJECT_ID]/topics/[TOPIC_ID]`
	// - `logging.googleapis.com/projects/[PROJECT_ID]]/locations/global/buckets/[BUCKET_ID]`
	//
	// The writer associated with the sink must have access to write to the above resource.
	Destination *string `pulumi:"destination"`
	// If set to True, then this sink is disabled and it does not export any log entries.
	Disabled *bool `pulumi:"disabled"`
	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both `filter` and one of `exclusions.filter`, it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
	Exclusions []FolderSinkExclusion `pulumi:"exclusions"`
	// The filter to apply when exporting logs. Only log entries that match the filter are exported.
	// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
	// write a filter.
	Filter *string `pulumi:"filter"`
	// The folder to be exported to the sink. Note that either `[FOLDER_ID]` or `folders/[FOLDER_ID]` is
	// accepted.
	Folder *string `pulumi:"folder"`
	// Whether or not to include children folders in the sink export. If true, logs
	// associated with child projects are also exported; otherwise only logs relating to the provided folder are included.
	IncludeChildren *bool `pulumi:"includeChildren"`
	// The name of the logging sink.
	Name *string `pulumi:"name"`
	// The identity associated with this sink. This identity must be granted write access to the
	// configured `destination`.
	WriterIdentity *string `pulumi:"writerIdentity"`
}

type FolderSinkState struct {
	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions FolderSinkBigqueryOptionsPtrInput
	// A description of this sink. The maximum length of the description is 8000 characters.
	Description pulumi.StringPtrInput
	// The destination of the sink (or, in other words, where logs are written to). Can be a
	// Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket. Examples:
	//
	// - `storage.googleapis.com/[GCS_BUCKET]`
	// - `bigquery.googleapis.com/projects/[PROJECT_ID]/datasets/[DATASET]`
	// - `pubsub.googleapis.com/projects/[PROJECT_ID]/topics/[TOPIC_ID]`
	// - `logging.googleapis.com/projects/[PROJECT_ID]]/locations/global/buckets/[BUCKET_ID]`
	//
	// The writer associated with the sink must have access to write to the above resource.
	Destination pulumi.StringPtrInput
	// If set to True, then this sink is disabled and it does not export any log entries.
	Disabled pulumi.BoolPtrInput
	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both `filter` and one of `exclusions.filter`, it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
	Exclusions FolderSinkExclusionArrayInput
	// The filter to apply when exporting logs. Only log entries that match the filter are exported.
	// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
	// write a filter.
	Filter pulumi.StringPtrInput
	// The folder to be exported to the sink. Note that either `[FOLDER_ID]` or `folders/[FOLDER_ID]` is
	// accepted.
	Folder pulumi.StringPtrInput
	// Whether or not to include children folders in the sink export. If true, logs
	// associated with child projects are also exported; otherwise only logs relating to the provided folder are included.
	IncludeChildren pulumi.BoolPtrInput
	// The name of the logging sink.
	Name pulumi.StringPtrInput
	// The identity associated with this sink. This identity must be granted write access to the
	// configured `destination`.
	WriterIdentity pulumi.StringPtrInput
}

func (FolderSinkState) ElementType() reflect.Type {
	return reflect.TypeOf((*folderSinkState)(nil)).Elem()
}

type folderSinkArgs struct {
	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions *FolderSinkBigqueryOptions `pulumi:"bigqueryOptions"`
	// A description of this sink. The maximum length of the description is 8000 characters.
	Description *string `pulumi:"description"`
	// The destination of the sink (or, in other words, where logs are written to). Can be a
	// Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket. Examples:
	//
	// - `storage.googleapis.com/[GCS_BUCKET]`
	// - `bigquery.googleapis.com/projects/[PROJECT_ID]/datasets/[DATASET]`
	// - `pubsub.googleapis.com/projects/[PROJECT_ID]/topics/[TOPIC_ID]`
	// - `logging.googleapis.com/projects/[PROJECT_ID]]/locations/global/buckets/[BUCKET_ID]`
	//
	// The writer associated with the sink must have access to write to the above resource.
	Destination string `pulumi:"destination"`
	// If set to True, then this sink is disabled and it does not export any log entries.
	Disabled *bool `pulumi:"disabled"`
	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both `filter` and one of `exclusions.filter`, it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
	Exclusions []FolderSinkExclusion `pulumi:"exclusions"`
	// The filter to apply when exporting logs. Only log entries that match the filter are exported.
	// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
	// write a filter.
	Filter *string `pulumi:"filter"`
	// The folder to be exported to the sink. Note that either `[FOLDER_ID]` or `folders/[FOLDER_ID]` is
	// accepted.
	Folder string `pulumi:"folder"`
	// Whether or not to include children folders in the sink export. If true, logs
	// associated with child projects are also exported; otherwise only logs relating to the provided folder are included.
	IncludeChildren *bool `pulumi:"includeChildren"`
	// The name of the logging sink.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a FolderSink resource.
type FolderSinkArgs struct {
	// Options that affect sinks exporting data to BigQuery. Structure documented below.
	BigqueryOptions FolderSinkBigqueryOptionsPtrInput
	// A description of this sink. The maximum length of the description is 8000 characters.
	Description pulumi.StringPtrInput
	// The destination of the sink (or, in other words, where logs are written to). Can be a
	// Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket. Examples:
	//
	// - `storage.googleapis.com/[GCS_BUCKET]`
	// - `bigquery.googleapis.com/projects/[PROJECT_ID]/datasets/[DATASET]`
	// - `pubsub.googleapis.com/projects/[PROJECT_ID]/topics/[TOPIC_ID]`
	// - `logging.googleapis.com/projects/[PROJECT_ID]]/locations/global/buckets/[BUCKET_ID]`
	//
	// The writer associated with the sink must have access to write to the above resource.
	Destination pulumi.StringInput
	// If set to True, then this sink is disabled and it does not export any log entries.
	Disabled pulumi.BoolPtrInput
	// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both `filter` and one of `exclusions.filter`, it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
	Exclusions FolderSinkExclusionArrayInput
	// The filter to apply when exporting logs. Only log entries that match the filter are exported.
	// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
	// write a filter.
	Filter pulumi.StringPtrInput
	// The folder to be exported to the sink. Note that either `[FOLDER_ID]` or `folders/[FOLDER_ID]` is
	// accepted.
	Folder pulumi.StringInput
	// Whether or not to include children folders in the sink export. If true, logs
	// associated with child projects are also exported; otherwise only logs relating to the provided folder are included.
	IncludeChildren pulumi.BoolPtrInput
	// The name of the logging sink.
	Name pulumi.StringPtrInput
}

func (FolderSinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*folderSinkArgs)(nil)).Elem()
}

type FolderSinkInput interface {
	pulumi.Input

	ToFolderSinkOutput() FolderSinkOutput
	ToFolderSinkOutputWithContext(ctx context.Context) FolderSinkOutput
}

func (*FolderSink) ElementType() reflect.Type {
	return reflect.TypeOf((**FolderSink)(nil)).Elem()
}

func (i *FolderSink) ToFolderSinkOutput() FolderSinkOutput {
	return i.ToFolderSinkOutputWithContext(context.Background())
}

func (i *FolderSink) ToFolderSinkOutputWithContext(ctx context.Context) FolderSinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FolderSinkOutput)
}

func (i *FolderSink) ToOutput(ctx context.Context) pulumix.Output[*FolderSink] {
	return pulumix.Output[*FolderSink]{
		OutputState: i.ToFolderSinkOutputWithContext(ctx).OutputState,
	}
}

// FolderSinkArrayInput is an input type that accepts FolderSinkArray and FolderSinkArrayOutput values.
// You can construct a concrete instance of `FolderSinkArrayInput` via:
//
//	FolderSinkArray{ FolderSinkArgs{...} }
type FolderSinkArrayInput interface {
	pulumi.Input

	ToFolderSinkArrayOutput() FolderSinkArrayOutput
	ToFolderSinkArrayOutputWithContext(context.Context) FolderSinkArrayOutput
}

type FolderSinkArray []FolderSinkInput

func (FolderSinkArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FolderSink)(nil)).Elem()
}

func (i FolderSinkArray) ToFolderSinkArrayOutput() FolderSinkArrayOutput {
	return i.ToFolderSinkArrayOutputWithContext(context.Background())
}

func (i FolderSinkArray) ToFolderSinkArrayOutputWithContext(ctx context.Context) FolderSinkArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FolderSinkArrayOutput)
}

func (i FolderSinkArray) ToOutput(ctx context.Context) pulumix.Output[[]*FolderSink] {
	return pulumix.Output[[]*FolderSink]{
		OutputState: i.ToFolderSinkArrayOutputWithContext(ctx).OutputState,
	}
}

// FolderSinkMapInput is an input type that accepts FolderSinkMap and FolderSinkMapOutput values.
// You can construct a concrete instance of `FolderSinkMapInput` via:
//
//	FolderSinkMap{ "key": FolderSinkArgs{...} }
type FolderSinkMapInput interface {
	pulumi.Input

	ToFolderSinkMapOutput() FolderSinkMapOutput
	ToFolderSinkMapOutputWithContext(context.Context) FolderSinkMapOutput
}

type FolderSinkMap map[string]FolderSinkInput

func (FolderSinkMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FolderSink)(nil)).Elem()
}

func (i FolderSinkMap) ToFolderSinkMapOutput() FolderSinkMapOutput {
	return i.ToFolderSinkMapOutputWithContext(context.Background())
}

func (i FolderSinkMap) ToFolderSinkMapOutputWithContext(ctx context.Context) FolderSinkMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FolderSinkMapOutput)
}

func (i FolderSinkMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*FolderSink] {
	return pulumix.Output[map[string]*FolderSink]{
		OutputState: i.ToFolderSinkMapOutputWithContext(ctx).OutputState,
	}
}

type FolderSinkOutput struct{ *pulumi.OutputState }

func (FolderSinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FolderSink)(nil)).Elem()
}

func (o FolderSinkOutput) ToFolderSinkOutput() FolderSinkOutput {
	return o
}

func (o FolderSinkOutput) ToFolderSinkOutputWithContext(ctx context.Context) FolderSinkOutput {
	return o
}

func (o FolderSinkOutput) ToOutput(ctx context.Context) pulumix.Output[*FolderSink] {
	return pulumix.Output[*FolderSink]{
		OutputState: o.OutputState,
	}
}

// Options that affect sinks exporting data to BigQuery. Structure documented below.
func (o FolderSinkOutput) BigqueryOptions() FolderSinkBigqueryOptionsOutput {
	return o.ApplyT(func(v *FolderSink) FolderSinkBigqueryOptionsOutput { return v.BigqueryOptions }).(FolderSinkBigqueryOptionsOutput)
}

// A description of this sink. The maximum length of the description is 8000 characters.
func (o FolderSinkOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FolderSink) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The destination of the sink (or, in other words, where logs are written to). Can be a
// Cloud Storage bucket, a PubSub topic, a BigQuery dataset or a Cloud Logging bucket. Examples:
//
// - `storage.googleapis.com/[GCS_BUCKET]`
// - `bigquery.googleapis.com/projects/[PROJECT_ID]/datasets/[DATASET]`
// - `pubsub.googleapis.com/projects/[PROJECT_ID]/topics/[TOPIC_ID]`
// - `logging.googleapis.com/projects/[PROJECT_ID]]/locations/global/buckets/[BUCKET_ID]`
//
// The writer associated with the sink must have access to write to the above resource.
func (o FolderSinkOutput) Destination() pulumi.StringOutput {
	return o.ApplyT(func(v *FolderSink) pulumi.StringOutput { return v.Destination }).(pulumi.StringOutput)
}

// If set to True, then this sink is disabled and it does not export any log entries.
func (o FolderSinkOutput) Disabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FolderSink) pulumi.BoolPtrOutput { return v.Disabled }).(pulumi.BoolPtrOutput)
}

// Log entries that match any of the exclusion filters will not be exported. If a log entry is matched by both `filter` and one of `exclusions.filter`, it will not be exported.  Can be repeated multiple times for multiple exclusions. Structure is documented below.
func (o FolderSinkOutput) Exclusions() FolderSinkExclusionArrayOutput {
	return o.ApplyT(func(v *FolderSink) FolderSinkExclusionArrayOutput { return v.Exclusions }).(FolderSinkExclusionArrayOutput)
}

// The filter to apply when exporting logs. Only log entries that match the filter are exported.
// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced_filters) for information on how to
// write a filter.
func (o FolderSinkOutput) Filter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FolderSink) pulumi.StringPtrOutput { return v.Filter }).(pulumi.StringPtrOutput)
}

// The folder to be exported to the sink. Note that either `[FOLDER_ID]` or `folders/[FOLDER_ID]` is
// accepted.
func (o FolderSinkOutput) Folder() pulumi.StringOutput {
	return o.ApplyT(func(v *FolderSink) pulumi.StringOutput { return v.Folder }).(pulumi.StringOutput)
}

// Whether or not to include children folders in the sink export. If true, logs
// associated with child projects are also exported; otherwise only logs relating to the provided folder are included.
func (o FolderSinkOutput) IncludeChildren() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FolderSink) pulumi.BoolPtrOutput { return v.IncludeChildren }).(pulumi.BoolPtrOutput)
}

// The name of the logging sink.
func (o FolderSinkOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FolderSink) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The identity associated with this sink. This identity must be granted write access to the
// configured `destination`.
func (o FolderSinkOutput) WriterIdentity() pulumi.StringOutput {
	return o.ApplyT(func(v *FolderSink) pulumi.StringOutput { return v.WriterIdentity }).(pulumi.StringOutput)
}

type FolderSinkArrayOutput struct{ *pulumi.OutputState }

func (FolderSinkArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*FolderSink)(nil)).Elem()
}

func (o FolderSinkArrayOutput) ToFolderSinkArrayOutput() FolderSinkArrayOutput {
	return o
}

func (o FolderSinkArrayOutput) ToFolderSinkArrayOutputWithContext(ctx context.Context) FolderSinkArrayOutput {
	return o
}

func (o FolderSinkArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*FolderSink] {
	return pulumix.Output[[]*FolderSink]{
		OutputState: o.OutputState,
	}
}

func (o FolderSinkArrayOutput) Index(i pulumi.IntInput) FolderSinkOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *FolderSink {
		return vs[0].([]*FolderSink)[vs[1].(int)]
	}).(FolderSinkOutput)
}

type FolderSinkMapOutput struct{ *pulumi.OutputState }

func (FolderSinkMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*FolderSink)(nil)).Elem()
}

func (o FolderSinkMapOutput) ToFolderSinkMapOutput() FolderSinkMapOutput {
	return o
}

func (o FolderSinkMapOutput) ToFolderSinkMapOutputWithContext(ctx context.Context) FolderSinkMapOutput {
	return o
}

func (o FolderSinkMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*FolderSink] {
	return pulumix.Output[map[string]*FolderSink]{
		OutputState: o.OutputState,
	}
}

func (o FolderSinkMapOutput) MapIndex(k pulumi.StringInput) FolderSinkOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *FolderSink {
		return vs[0].(map[string]*FolderSink)[vs[1].(string)]
	}).(FolderSinkOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*FolderSinkInput)(nil)).Elem(), &FolderSink{})
	pulumi.RegisterInputType(reflect.TypeOf((*FolderSinkArrayInput)(nil)).Elem(), FolderSinkArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FolderSinkMapInput)(nil)).Elem(), FolderSinkMap{})
	pulumi.RegisterOutputType(FolderSinkOutput{})
	pulumi.RegisterOutputType(FolderSinkArrayOutput{})
	pulumi.RegisterOutputType(FolderSinkMapOutput{})
}
