// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package bigquery

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Jobs are actions that BigQuery runs on your behalf to load data, export data, query data, or copy data.
// Once a BigQuery job is created, it cannot be changed or deleted.
//
// To get more information about Job, see:
//
// * [API documentation](https://cloud.google.com/bigquery/docs/reference/rest/v2/jobs)
// * How-to Guides
//     * [BigQuery Jobs Intro](https://cloud.google.com/bigquery/docs/jobs-overview)
//
// ## Example Usage
// ### Bigquery Job Query
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/bigquery"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		bar, err := bigquery.NewDataset(ctx, "bar", &bigquery.DatasetArgs{
// 			DatasetId:    pulumi.String("job_query_dataset"),
// 			FriendlyName: pulumi.String("test"),
// 			Description:  pulumi.String("This is a test description"),
// 			Location:     pulumi.String("US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		foo, err := bigquery.NewTable(ctx, "foo", &bigquery.TableArgs{
// 			DeletionProtection: pulumi.Bool(false),
// 			DatasetId:          bar.DatasetId,
// 			TableId:            pulumi.String("job_query_table"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = bigquery.NewJob(ctx, "job", &bigquery.JobArgs{
// 			JobId: pulumi.String("job_query"),
// 			Labels: pulumi.StringMap{
// 				"example-label": pulumi.String("example-value"),
// 			},
// 			Query: &bigquery.JobQueryArgs{
// 				Query: pulumi.String("SELECT state FROM [lookerdata:cdc.project_tycho_reports]"),
// 				DestinationTable: &bigquery.JobQueryDestinationTableArgs{
// 					ProjectId: foo.Project,
// 					DatasetId: foo.DatasetId,
// 					TableId:   foo.TableId,
// 				},
// 				AllowLargeResults: pulumi.Bool(true),
// 				FlattenResults:    pulumi.Bool(true),
// 				ScriptOptions: &bigquery.JobQueryScriptOptionsArgs{
// 					KeyResultStatement: pulumi.String("LAST"),
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
// ### Bigquery Job Query Table Reference
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/bigquery"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		bar, err := bigquery.NewDataset(ctx, "bar", &bigquery.DatasetArgs{
// 			DatasetId:    pulumi.String("job_query_dataset"),
// 			FriendlyName: pulumi.String("test"),
// 			Description:  pulumi.String("This is a test description"),
// 			Location:     pulumi.String("US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		foo, err := bigquery.NewTable(ctx, "foo", &bigquery.TableArgs{
// 			DeletionProtection: pulumi.Bool(false),
// 			DatasetId:          bar.DatasetId,
// 			TableId:            pulumi.String("job_query_table"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = bigquery.NewJob(ctx, "job", &bigquery.JobArgs{
// 			JobId: pulumi.String("job_query"),
// 			Labels: pulumi.StringMap{
// 				"example-label": pulumi.String("example-value"),
// 			},
// 			Query: &bigquery.JobQueryArgs{
// 				Query: pulumi.String("SELECT state FROM [lookerdata:cdc.project_tycho_reports]"),
// 				DestinationTable: &bigquery.JobQueryDestinationTableArgs{
// 					TableId: foo.ID(),
// 				},
// 				DefaultDataset: &bigquery.JobQueryDefaultDatasetArgs{
// 					DatasetId: bar.ID(),
// 				},
// 				AllowLargeResults: pulumi.Bool(true),
// 				FlattenResults:    pulumi.Bool(true),
// 				ScriptOptions: &bigquery.JobQueryScriptOptionsArgs{
// 					KeyResultStatement: pulumi.String("LAST"),
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
// ### Bigquery Job Load
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/bigquery"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		bar, err := bigquery.NewDataset(ctx, "bar", &bigquery.DatasetArgs{
// 			DatasetId:    pulumi.String("job_load_dataset"),
// 			FriendlyName: pulumi.String("test"),
// 			Description:  pulumi.String("This is a test description"),
// 			Location:     pulumi.String("US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		foo, err := bigquery.NewTable(ctx, "foo", &bigquery.TableArgs{
// 			DeletionProtection: pulumi.Bool(false),
// 			DatasetId:          bar.DatasetId,
// 			TableId:            pulumi.String("job_load_table"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = bigquery.NewJob(ctx, "job", &bigquery.JobArgs{
// 			JobId: pulumi.String("job_load"),
// 			Labels: pulumi.StringMap{
// 				"my_job": pulumi.String("load"),
// 			},
// 			Load: &bigquery.JobLoadArgs{
// 				SourceUris: pulumi.StringArray{
// 					pulumi.String("gs://cloud-samples-data/bigquery/us-states/us-states-by-date.csv"),
// 				},
// 				DestinationTable: &bigquery.JobLoadDestinationTableArgs{
// 					ProjectId: foo.Project,
// 					DatasetId: foo.DatasetId,
// 					TableId:   foo.TableId,
// 				},
// 				SkipLeadingRows: pulumi.Int(1),
// 				SchemaUpdateOptions: pulumi.StringArray{
// 					pulumi.String("ALLOW_FIELD_RELAXATION"),
// 					pulumi.String("ALLOW_FIELD_ADDITION"),
// 				},
// 				WriteDisposition: pulumi.String("WRITE_APPEND"),
// 				Autodetect:       pulumi.Bool(true),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Bigquery Job Extract
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/bigquery"
// 	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/storage"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := bigquery.NewDataset(ctx, "source-oneDataset", &bigquery.DatasetArgs{
// 			DatasetId:    pulumi.String("job_extract_dataset"),
// 			FriendlyName: pulumi.String("test"),
// 			Description:  pulumi.String("This is a test description"),
// 			Location:     pulumi.String("US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = bigquery.NewTable(ctx, "source-oneTable", &bigquery.TableArgs{
// 			DeletionProtection: pulumi.Bool(false),
// 			DatasetId:          source_oneDataset.DatasetId,
// 			TableId:            pulumi.String("job_extract_table"),
// 			Schema:             pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "[\n", "  {\n", "    \"name\": \"name\",\n", "    \"type\": \"STRING\",\n", "    \"mode\": \"NULLABLE\"\n", "  },\n", "  {\n", "    \"name\": \"post_abbr\",\n", "    \"type\": \"STRING\",\n", "    \"mode\": \"NULLABLE\"\n", "  },\n", "  {\n", "    \"name\": \"date\",\n", "    \"type\": \"DATE\",\n", "    \"mode\": \"NULLABLE\"\n", "  }\n", "]\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		dest, err := storage.NewBucket(ctx, "dest", &storage.BucketArgs{
// 			Location:     pulumi.String("US"),
// 			ForceDestroy: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = bigquery.NewJob(ctx, "job", &bigquery.JobArgs{
// 			JobId: pulumi.String("job_extract"),
// 			Extract: &bigquery.JobExtractArgs{
// 				DestinationUris: pulumi.StringArray{
// 					dest.Url.ApplyT(func(url string) (string, error) {
// 						return fmt.Sprintf("%v%v", url, "/extract"), nil
// 					}).(pulumi.StringOutput),
// 				},
// 				SourceTable: &bigquery.JobExtractSourceTableArgs{
// 					ProjectId: source_oneTable.Project,
// 					DatasetId: source_oneTable.DatasetId,
// 					TableId:   source_oneTable.TableId,
// 				},
// 				DestinationFormat: pulumi.String("NEWLINE_DELIMITED_JSON"),
// 				Compression:       pulumi.String("GZIP"),
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
// Job can be imported using any of these accepted formats
//
// ```sh
//  $ pulumi import gcp:bigquery/job:Job default projects/{{project}}/jobs/{{job_id}}/location/{{location}}
// ```
//
// ```sh
//  $ pulumi import gcp:bigquery/job:Job default projects/{{project}}/jobs/{{job_id}}
// ```
//
// ```sh
//  $ pulumi import gcp:bigquery/job:Job default {{project}}/{{job_id}}/{{location}}
// ```
//
// ```sh
//  $ pulumi import gcp:bigquery/job:Job default {{job_id}}/{{location}}
// ```
//
// ```sh
//  $ pulumi import gcp:bigquery/job:Job default {{project}}/{{job_id}}
// ```
//
// ```sh
//  $ pulumi import gcp:bigquery/job:Job default {{job_id}}
// ```
type Job struct {
	pulumi.CustomResourceState

	// Copies a table.
	// Structure is documented below.
	Copy JobCopyPtrOutput `pulumi:"copy"`
	// Configures an extract job.
	// Structure is documented below.
	Extract JobExtractPtrOutput `pulumi:"extract"`
	// The ID of the job. The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), or dashes (-). The maximum length is 1,024 characters.
	JobId pulumi.StringOutput `pulumi:"jobId"`
	// Job timeout in milliseconds. If this time limit is exceeded, BigQuery may attempt to terminate the job.
	JobTimeoutMs pulumi.StringPtrOutput `pulumi:"jobTimeoutMs"`
	// The type of the job.
	JobType pulumi.StringOutput `pulumi:"jobType"`
	// The labels associated with this job. You can use these to organize and group your jobs.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Configures a load job.
	// Structure is documented below.
	Load JobLoadPtrOutput `pulumi:"load"`
	// The geographic location of the job. The default value is US.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Configures a query job.
	// Structure is documented below.
	Query JobQueryPtrOutput `pulumi:"query"`
	// The status of this job. Examine this value when polling an asynchronous job to see if the job is complete.
	Statuses JobStatusArrayOutput `pulumi:"statuses"`
	// Email address of the user who ran the job.
	UserEmail pulumi.StringOutput `pulumi:"userEmail"`
}

// NewJob registers a new resource with the given unique name, arguments, and options.
func NewJob(ctx *pulumi.Context,
	name string, args *JobArgs, opts ...pulumi.ResourceOption) (*Job, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.JobId == nil {
		return nil, errors.New("invalid value for required argument 'JobId'")
	}
	var resource Job
	err := ctx.RegisterResource("gcp:bigquery/job:Job", name, args, &resource, opts...)
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
	err := ctx.ReadResource("gcp:bigquery/job:Job", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Job resources.
type jobState struct {
	// Copies a table.
	// Structure is documented below.
	Copy *JobCopy `pulumi:"copy"`
	// Configures an extract job.
	// Structure is documented below.
	Extract *JobExtract `pulumi:"extract"`
	// The ID of the job. The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), or dashes (-). The maximum length is 1,024 characters.
	JobId *string `pulumi:"jobId"`
	// Job timeout in milliseconds. If this time limit is exceeded, BigQuery may attempt to terminate the job.
	JobTimeoutMs *string `pulumi:"jobTimeoutMs"`
	// The type of the job.
	JobType *string `pulumi:"jobType"`
	// The labels associated with this job. You can use these to organize and group your jobs.
	Labels map[string]string `pulumi:"labels"`
	// Configures a load job.
	// Structure is documented below.
	Load *JobLoad `pulumi:"load"`
	// The geographic location of the job. The default value is US.
	Location *string `pulumi:"location"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Configures a query job.
	// Structure is documented below.
	Query *JobQuery `pulumi:"query"`
	// The status of this job. Examine this value when polling an asynchronous job to see if the job is complete.
	Statuses []JobStatus `pulumi:"statuses"`
	// Email address of the user who ran the job.
	UserEmail *string `pulumi:"userEmail"`
}

type JobState struct {
	// Copies a table.
	// Structure is documented below.
	Copy JobCopyPtrInput
	// Configures an extract job.
	// Structure is documented below.
	Extract JobExtractPtrInput
	// The ID of the job. The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), or dashes (-). The maximum length is 1,024 characters.
	JobId pulumi.StringPtrInput
	// Job timeout in milliseconds. If this time limit is exceeded, BigQuery may attempt to terminate the job.
	JobTimeoutMs pulumi.StringPtrInput
	// The type of the job.
	JobType pulumi.StringPtrInput
	// The labels associated with this job. You can use these to organize and group your jobs.
	Labels pulumi.StringMapInput
	// Configures a load job.
	// Structure is documented below.
	Load JobLoadPtrInput
	// The geographic location of the job. The default value is US.
	Location pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Configures a query job.
	// Structure is documented below.
	Query JobQueryPtrInput
	// The status of this job. Examine this value when polling an asynchronous job to see if the job is complete.
	Statuses JobStatusArrayInput
	// Email address of the user who ran the job.
	UserEmail pulumi.StringPtrInput
}

func (JobState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobState)(nil)).Elem()
}

type jobArgs struct {
	// Copies a table.
	// Structure is documented below.
	Copy *JobCopy `pulumi:"copy"`
	// Configures an extract job.
	// Structure is documented below.
	Extract *JobExtract `pulumi:"extract"`
	// The ID of the job. The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), or dashes (-). The maximum length is 1,024 characters.
	JobId string `pulumi:"jobId"`
	// Job timeout in milliseconds. If this time limit is exceeded, BigQuery may attempt to terminate the job.
	JobTimeoutMs *string `pulumi:"jobTimeoutMs"`
	// The labels associated with this job. You can use these to organize and group your jobs.
	Labels map[string]string `pulumi:"labels"`
	// Configures a load job.
	// Structure is documented below.
	Load *JobLoad `pulumi:"load"`
	// The geographic location of the job. The default value is US.
	Location *string `pulumi:"location"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Configures a query job.
	// Structure is documented below.
	Query *JobQuery `pulumi:"query"`
}

// The set of arguments for constructing a Job resource.
type JobArgs struct {
	// Copies a table.
	// Structure is documented below.
	Copy JobCopyPtrInput
	// Configures an extract job.
	// Structure is documented below.
	Extract JobExtractPtrInput
	// The ID of the job. The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), or dashes (-). The maximum length is 1,024 characters.
	JobId pulumi.StringInput
	// Job timeout in milliseconds. If this time limit is exceeded, BigQuery may attempt to terminate the job.
	JobTimeoutMs pulumi.StringPtrInput
	// The labels associated with this job. You can use these to organize and group your jobs.
	Labels pulumi.StringMapInput
	// Configures a load job.
	// Structure is documented below.
	Load JobLoadPtrInput
	// The geographic location of the job. The default value is US.
	Location pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Configures a query job.
	// Structure is documented below.
	Query JobQueryPtrInput
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
	return reflect.TypeOf((**Job)(nil)).Elem()
}

func (o JobOutput) ToJobOutput() JobOutput {
	return o
}

func (o JobOutput) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return o
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
