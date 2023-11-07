// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package bigquery

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Jobs are actions that BigQuery runs on your behalf to load data, export data, query data, or copy data.
// Once a BigQuery job is created, it cannot be changed or deleted.
//
// To get more information about Job, see:
//
// * [API documentation](https://cloud.google.com/bigquery/docs/reference/rest/v2/jobs)
// * How-to Guides
//   - [BigQuery Jobs Intro](https://cloud.google.com/bigquery/docs/jobs-overview)
//
// ## Example Usage
// ### Bigquery Job Query
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/bigquery"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			bar, err := bigquery.NewDataset(ctx, "bar", &bigquery.DatasetArgs{
//				DatasetId:    pulumi.String("job_query_dataset"),
//				FriendlyName: pulumi.String("test"),
//				Description:  pulumi.String("This is a test description"),
//				Location:     pulumi.String("US"),
//			})
//			if err != nil {
//				return err
//			}
//			foo, err := bigquery.NewTable(ctx, "foo", &bigquery.TableArgs{
//				DeletionProtection: pulumi.Bool(false),
//				DatasetId:          bar.DatasetId,
//				TableId:            pulumi.String("job_query_table"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = bigquery.NewJob(ctx, "job", &bigquery.JobArgs{
//				JobId: pulumi.String("job_query"),
//				Labels: pulumi.StringMap{
//					"example-label": pulumi.String("example-value"),
//				},
//				Query: &bigquery.JobQueryArgs{
//					Query: pulumi.String("SELECT state FROM [lookerdata:cdc.project_tycho_reports]"),
//					DestinationTable: &bigquery.JobQueryDestinationTableArgs{
//						ProjectId: foo.Project,
//						DatasetId: foo.DatasetId,
//						TableId:   foo.TableId,
//					},
//					AllowLargeResults: pulumi.Bool(true),
//					FlattenResults:    pulumi.Bool(true),
//					ScriptOptions: &bigquery.JobQueryScriptOptionsArgs{
//						KeyResultStatement: pulumi.String("LAST"),
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
// ### Bigquery Job Query Table Reference
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/bigquery"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			bar, err := bigquery.NewDataset(ctx, "bar", &bigquery.DatasetArgs{
//				DatasetId:    pulumi.String("job_query_dataset"),
//				FriendlyName: pulumi.String("test"),
//				Description:  pulumi.String("This is a test description"),
//				Location:     pulumi.String("US"),
//			})
//			if err != nil {
//				return err
//			}
//			foo, err := bigquery.NewTable(ctx, "foo", &bigquery.TableArgs{
//				DeletionProtection: pulumi.Bool(false),
//				DatasetId:          bar.DatasetId,
//				TableId:            pulumi.String("job_query_table"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = bigquery.NewJob(ctx, "job", &bigquery.JobArgs{
//				JobId: pulumi.String("job_query"),
//				Labels: pulumi.StringMap{
//					"example-label": pulumi.String("example-value"),
//				},
//				Query: &bigquery.JobQueryArgs{
//					Query: pulumi.String("SELECT state FROM [lookerdata:cdc.project_tycho_reports]"),
//					DestinationTable: &bigquery.JobQueryDestinationTableArgs{
//						TableId: foo.ID(),
//					},
//					DefaultDataset: &bigquery.JobQueryDefaultDatasetArgs{
//						DatasetId: bar.ID(),
//					},
//					AllowLargeResults: pulumi.Bool(true),
//					FlattenResults:    pulumi.Bool(true),
//					ScriptOptions: &bigquery.JobQueryScriptOptionsArgs{
//						KeyResultStatement: pulumi.String("LAST"),
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
// ### Bigquery Job Load
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/bigquery"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			bar, err := bigquery.NewDataset(ctx, "bar", &bigquery.DatasetArgs{
//				DatasetId:    pulumi.String("job_load_dataset"),
//				FriendlyName: pulumi.String("test"),
//				Description:  pulumi.String("This is a test description"),
//				Location:     pulumi.String("US"),
//			})
//			if err != nil {
//				return err
//			}
//			foo, err := bigquery.NewTable(ctx, "foo", &bigquery.TableArgs{
//				DeletionProtection: pulumi.Bool(false),
//				DatasetId:          bar.DatasetId,
//				TableId:            pulumi.String("job_load_table"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = bigquery.NewJob(ctx, "job", &bigquery.JobArgs{
//				JobId: pulumi.String("job_load"),
//				Labels: pulumi.StringMap{
//					"my_job": pulumi.String("load"),
//				},
//				Load: &bigquery.JobLoadArgs{
//					SourceUris: pulumi.StringArray{
//						pulumi.String("gs://cloud-samples-data/bigquery/us-states/us-states-by-date.csv"),
//					},
//					DestinationTable: &bigquery.JobLoadDestinationTableArgs{
//						ProjectId: foo.Project,
//						DatasetId: foo.DatasetId,
//						TableId:   foo.TableId,
//					},
//					SkipLeadingRows: pulumi.Int(1),
//					SchemaUpdateOptions: pulumi.StringArray{
//						pulumi.String("ALLOW_FIELD_RELAXATION"),
//						pulumi.String("ALLOW_FIELD_ADDITION"),
//					},
//					WriteDisposition: pulumi.String("WRITE_APPEND"),
//					Autodetect:       pulumi.Bool(true),
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
// ### Bigquery Job Load Parquet
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/bigquery"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/storage"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			testBucket, err := storage.NewBucket(ctx, "testBucket", &storage.BucketArgs{
//				Location:                 pulumi.String("US"),
//				UniformBucketLevelAccess: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			testBucketObject, err := storage.NewBucketObject(ctx, "testBucketObject", &storage.BucketObjectArgs{
//				Source: pulumi.NewFileAsset("./test-fixtures/test.parquet.gzip"),
//				Bucket: testBucket.Name,
//			})
//			if err != nil {
//				return err
//			}
//			testDataset, err := bigquery.NewDataset(ctx, "testDataset", &bigquery.DatasetArgs{
//				DatasetId:    pulumi.String("job_load_dataset"),
//				FriendlyName: pulumi.String("test"),
//				Description:  pulumi.String("This is a test description"),
//				Location:     pulumi.String("US"),
//			})
//			if err != nil {
//				return err
//			}
//			testTable, err := bigquery.NewTable(ctx, "testTable", &bigquery.TableArgs{
//				DeletionProtection: pulumi.Bool(false),
//				TableId:            pulumi.String("job_load_table"),
//				DatasetId:          testDataset.DatasetId,
//			})
//			if err != nil {
//				return err
//			}
//			_, err = bigquery.NewJob(ctx, "job", &bigquery.JobArgs{
//				JobId: pulumi.String("job_load"),
//				Labels: pulumi.StringMap{
//					"my_job": pulumi.String("load"),
//				},
//				Load: &bigquery.JobLoadArgs{
//					SourceUris: pulumi.StringArray{
//						pulumi.All(testBucketObject.Bucket, testBucketObject.Name).ApplyT(func(_args []interface{}) (string, error) {
//							bucket := _args[0].(string)
//							name := _args[1].(string)
//							return fmt.Sprintf("gs://%v/%v", bucket, name), nil
//						}).(pulumi.StringOutput),
//					},
//					DestinationTable: &bigquery.JobLoadDestinationTableArgs{
//						ProjectId: testTable.Project,
//						DatasetId: testTable.DatasetId,
//						TableId:   testTable.TableId,
//					},
//					SchemaUpdateOptions: pulumi.StringArray{
//						pulumi.String("ALLOW_FIELD_RELAXATION"),
//						pulumi.String("ALLOW_FIELD_ADDITION"),
//					},
//					WriteDisposition: pulumi.String("WRITE_APPEND"),
//					SourceFormat:     pulumi.String("PARQUET"),
//					Autodetect:       pulumi.Bool(true),
//					ParquetOptions: &bigquery.JobLoadParquetOptionsArgs{
//						EnumAsString:        pulumi.Bool(true),
//						EnableListInference: pulumi.Bool(true),
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
// ### Bigquery Job Extract
//
// ```go
// package main
//
// import (
//
//	"fmt"
//
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/bigquery"
//	"github.com/pulumi/pulumi-gcp/sdk/v7/go/gcp/storage"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := bigquery.NewDataset(ctx, "source-oneDataset", &bigquery.DatasetArgs{
//				DatasetId:    pulumi.String("job_extract_dataset"),
//				FriendlyName: pulumi.String("test"),
//				Description:  pulumi.String("This is a test description"),
//				Location:     pulumi.String("US"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = bigquery.NewTable(ctx, "source-oneTable", &bigquery.TableArgs{
//				DeletionProtection: pulumi.Bool(false),
//				DatasetId:          source_oneDataset.DatasetId,
//				TableId:            pulumi.String("job_extract_table"),
//				Schema: pulumi.String(`[
//	  {
//	    "name": "name",
//	    "type": "STRING",
//	    "mode": "NULLABLE"
//	  },
//	  {
//	    "name": "post_abbr",
//	    "type": "STRING",
//	    "mode": "NULLABLE"
//	  },
//	  {
//	    "name": "date",
//	    "type": "DATE",
//	    "mode": "NULLABLE"
//	  }
//
// ]
// `),
//
//			})
//			if err != nil {
//				return err
//			}
//			dest, err := storage.NewBucket(ctx, "dest", &storage.BucketArgs{
//				Location:     pulumi.String("US"),
//				ForceDestroy: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = bigquery.NewJob(ctx, "job", &bigquery.JobArgs{
//				JobId: pulumi.String("job_extract"),
//				Extract: &bigquery.JobExtractArgs{
//					DestinationUris: pulumi.StringArray{
//						dest.Url.ApplyT(func(url string) (string, error) {
//							return fmt.Sprintf("%v/extract", url), nil
//						}).(pulumi.StringOutput),
//					},
//					SourceTable: &bigquery.JobExtractSourceTableArgs{
//						ProjectId: source_oneTable.Project,
//						DatasetId: source_oneTable.DatasetId,
//						TableId:   source_oneTable.TableId,
//					},
//					DestinationFormat: pulumi.String("NEWLINE_DELIMITED_JSON"),
//					Compression:       pulumi.String("GZIP"),
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
// # Job can be imported using any of these accepted formats
//
// ```sh
//
//	$ pulumi import gcp:bigquery/job:Job default projects/{{project}}/jobs/{{job_id}}/location/{{location}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:bigquery/job:Job default projects/{{project}}/jobs/{{job_id}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:bigquery/job:Job default {{project}}/{{job_id}}/{{location}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:bigquery/job:Job default {{job_id}}/{{location}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:bigquery/job:Job default {{project}}/{{job_id}}
//
// ```
//
// ```sh
//
//	$ pulumi import gcp:bigquery/job:Job default {{job_id}}
//
// ```
type Job struct {
	pulumi.CustomResourceState

	// Copies a table.
	// Structure is documented below.
	Copy JobCopyPtrOutput `pulumi:"copy"`
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
	// clients and services.
	EffectiveLabels pulumi.StringMapOutput `pulumi:"effectiveLabels"`
	// Configures an extract job.
	// Structure is documented below.
	Extract JobExtractPtrOutput `pulumi:"extract"`
	// The ID of the job. The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), or dashes (-). The maximum length is 1,024 characters.
	JobId pulumi.StringOutput `pulumi:"jobId"`
	// Job timeout in milliseconds. If this time limit is exceeded, BigQuery may attempt to terminate the job.
	JobTimeoutMs pulumi.StringPtrOutput `pulumi:"jobTimeoutMs"`
	// (Output)
	// The type of the job.
	JobType pulumi.StringOutput `pulumi:"jobType"`
	// The labels associated with this job. You can use these to organize and group your jobs.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Configures a load job.
	// Structure is documented below.
	Load JobLoadPtrOutput `pulumi:"load"`
	// The geographic location of the job. The default value is US.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// (Output)
	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	PulumiLabels pulumi.StringMapOutput `pulumi:"pulumiLabels"`
	// SQL query text to execute. The useLegacySql field can be used to indicate whether the query uses legacy SQL or standard SQL.
	// *NOTE*: queries containing [DML language](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-manipulation-language)
	// (`DELETE`, `UPDATE`, `MERGE`, `INSERT`) must specify `createDisposition = ""` and `writeDisposition = ""`.
	Query JobQueryPtrOutput `pulumi:"query"`
	// The status of this job. Examine this value when polling an asynchronous job to see if the job is complete.
	// Structure is documented below.
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
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"effectiveLabels",
		"pulumiLabels",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
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
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
	// clients and services.
	EffectiveLabels map[string]string `pulumi:"effectiveLabels"`
	// Configures an extract job.
	// Structure is documented below.
	Extract *JobExtract `pulumi:"extract"`
	// The ID of the job. The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), or dashes (-). The maximum length is 1,024 characters.
	JobId *string `pulumi:"jobId"`
	// Job timeout in milliseconds. If this time limit is exceeded, BigQuery may attempt to terminate the job.
	JobTimeoutMs *string `pulumi:"jobTimeoutMs"`
	// (Output)
	// The type of the job.
	JobType *string `pulumi:"jobType"`
	// The labels associated with this job. You can use these to organize and group your jobs.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// Configures a load job.
	// Structure is documented below.
	Load *JobLoad `pulumi:"load"`
	// The geographic location of the job. The default value is US.
	Location *string `pulumi:"location"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// (Output)
	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	PulumiLabels map[string]string `pulumi:"pulumiLabels"`
	// SQL query text to execute. The useLegacySql field can be used to indicate whether the query uses legacy SQL or standard SQL.
	// *NOTE*: queries containing [DML language](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-manipulation-language)
	// (`DELETE`, `UPDATE`, `MERGE`, `INSERT`) must specify `createDisposition = ""` and `writeDisposition = ""`.
	Query *JobQuery `pulumi:"query"`
	// The status of this job. Examine this value when polling an asynchronous job to see if the job is complete.
	// Structure is documented below.
	Statuses []JobStatus `pulumi:"statuses"`
	// Email address of the user who ran the job.
	UserEmail *string `pulumi:"userEmail"`
}

type JobState struct {
	// Copies a table.
	// Structure is documented below.
	Copy JobCopyPtrInput
	// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
	// clients and services.
	EffectiveLabels pulumi.StringMapInput
	// Configures an extract job.
	// Structure is documented below.
	Extract JobExtractPtrInput
	// The ID of the job. The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), or dashes (-). The maximum length is 1,024 characters.
	JobId pulumi.StringPtrInput
	// Job timeout in milliseconds. If this time limit is exceeded, BigQuery may attempt to terminate the job.
	JobTimeoutMs pulumi.StringPtrInput
	// (Output)
	// The type of the job.
	JobType pulumi.StringPtrInput
	// The labels associated with this job. You can use these to organize and group your jobs.
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// Configures a load job.
	// Structure is documented below.
	Load JobLoadPtrInput
	// The geographic location of the job. The default value is US.
	Location pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// (Output)
	// The combination of labels configured directly on the resource
	// and default labels configured on the provider.
	PulumiLabels pulumi.StringMapInput
	// SQL query text to execute. The useLegacySql field can be used to indicate whether the query uses legacy SQL or standard SQL.
	// *NOTE*: queries containing [DML language](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-manipulation-language)
	// (`DELETE`, `UPDATE`, `MERGE`, `INSERT`) must specify `createDisposition = ""` and `writeDisposition = ""`.
	Query JobQueryPtrInput
	// The status of this job. Examine this value when polling an asynchronous job to see if the job is complete.
	// Structure is documented below.
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
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels map[string]string `pulumi:"labels"`
	// Configures a load job.
	// Structure is documented below.
	Load *JobLoad `pulumi:"load"`
	// The geographic location of the job. The default value is US.
	Location *string `pulumi:"location"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// SQL query text to execute. The useLegacySql field can be used to indicate whether the query uses legacy SQL or standard SQL.
	// *NOTE*: queries containing [DML language](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-manipulation-language)
	// (`DELETE`, `UPDATE`, `MERGE`, `INSERT`) must specify `createDisposition = ""` and `writeDisposition = ""`.
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
	//
	// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
	// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
	Labels pulumi.StringMapInput
	// Configures a load job.
	// Structure is documented below.
	Load JobLoadPtrInput
	// The geographic location of the job. The default value is US.
	Location pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// SQL query text to execute. The useLegacySql field can be used to indicate whether the query uses legacy SQL or standard SQL.
	// *NOTE*: queries containing [DML language](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-manipulation-language)
	// (`DELETE`, `UPDATE`, `MERGE`, `INSERT`) must specify `createDisposition = ""` and `writeDisposition = ""`.
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

func (i *Job) ToOutput(ctx context.Context) pulumix.Output[*Job] {
	return pulumix.Output[*Job]{
		OutputState: i.ToJobOutputWithContext(ctx).OutputState,
	}
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

func (i JobArray) ToOutput(ctx context.Context) pulumix.Output[[]*Job] {
	return pulumix.Output[[]*Job]{
		OutputState: i.ToJobArrayOutputWithContext(ctx).OutputState,
	}
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

func (i JobMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Job] {
	return pulumix.Output[map[string]*Job]{
		OutputState: i.ToJobMapOutputWithContext(ctx).OutputState,
	}
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

func (o JobOutput) ToOutput(ctx context.Context) pulumix.Output[*Job] {
	return pulumix.Output[*Job]{
		OutputState: o.OutputState,
	}
}

// Copies a table.
// Structure is documented below.
func (o JobOutput) Copy() JobCopyPtrOutput {
	return o.ApplyT(func(v *Job) JobCopyPtrOutput { return v.Copy }).(JobCopyPtrOutput)
}

// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
// clients and services.
func (o JobOutput) EffectiveLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Job) pulumi.StringMapOutput { return v.EffectiveLabels }).(pulumi.StringMapOutput)
}

// Configures an extract job.
// Structure is documented below.
func (o JobOutput) Extract() JobExtractPtrOutput {
	return o.ApplyT(func(v *Job) JobExtractPtrOutput { return v.Extract }).(JobExtractPtrOutput)
}

// The ID of the job. The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), or dashes (-). The maximum length is 1,024 characters.
func (o JobOutput) JobId() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.JobId }).(pulumi.StringOutput)
}

// Job timeout in milliseconds. If this time limit is exceeded, BigQuery may attempt to terminate the job.
func (o JobOutput) JobTimeoutMs() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Job) pulumi.StringPtrOutput { return v.JobTimeoutMs }).(pulumi.StringPtrOutput)
}

// (Output)
// The type of the job.
func (o JobOutput) JobType() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.JobType }).(pulumi.StringOutput)
}

// The labels associated with this job. You can use these to organize and group your jobs.
//
// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
// Please refer to the field `effectiveLabels` for all of the labels present on the resource.
func (o JobOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Job) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Configures a load job.
// Structure is documented below.
func (o JobOutput) Load() JobLoadPtrOutput {
	return o.ApplyT(func(v *Job) JobLoadPtrOutput { return v.Load }).(JobLoadPtrOutput)
}

// The geographic location of the job. The default value is US.
func (o JobOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Job) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// The ID of the project in which the resource belongs.
// If it is not provided, the provider project is used.
func (o JobOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// (Output)
// The combination of labels configured directly on the resource
// and default labels configured on the provider.
func (o JobOutput) PulumiLabels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Job) pulumi.StringMapOutput { return v.PulumiLabels }).(pulumi.StringMapOutput)
}

// SQL query text to execute. The useLegacySql field can be used to indicate whether the query uses legacy SQL or standard SQL.
// *NOTE*: queries containing [DML language](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-manipulation-language)
// (`DELETE`, `UPDATE`, `MERGE`, `INSERT`) must specify `createDisposition = ""` and `writeDisposition = ""`.
func (o JobOutput) Query() JobQueryPtrOutput {
	return o.ApplyT(func(v *Job) JobQueryPtrOutput { return v.Query }).(JobQueryPtrOutput)
}

// The status of this job. Examine this value when polling an asynchronous job to see if the job is complete.
// Structure is documented below.
func (o JobOutput) Statuses() JobStatusArrayOutput {
	return o.ApplyT(func(v *Job) JobStatusArrayOutput { return v.Statuses }).(JobStatusArrayOutput)
}

// Email address of the user who ran the job.
func (o JobOutput) UserEmail() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.UserEmail }).(pulumi.StringOutput)
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

func (o JobArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Job] {
	return pulumix.Output[[]*Job]{
		OutputState: o.OutputState,
	}
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

func (o JobMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Job] {
	return pulumix.Output[map[string]*Job]{
		OutputState: o.OutputState,
	}
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
