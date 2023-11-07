// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.BigQuery
{
    /// <summary>
    /// Jobs are actions that BigQuery runs on your behalf to load data, export data, query data, or copy data.
    /// Once a BigQuery job is created, it cannot be changed or deleted.
    /// 
    /// To get more information about Job, see:
    /// 
    /// * [API documentation](https://cloud.google.com/bigquery/docs/reference/rest/v2/jobs)
    /// * How-to Guides
    ///     * [BigQuery Jobs Intro](https://cloud.google.com/bigquery/docs/jobs-overview)
    /// 
    /// ## Example Usage
    /// ### Bigquery Job Query
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var bar = new Gcp.BigQuery.Dataset("bar", new()
    ///     {
    ///         DatasetId = "job_query_dataset",
    ///         FriendlyName = "test",
    ///         Description = "This is a test description",
    ///         Location = "US",
    ///     });
    /// 
    ///     var foo = new Gcp.BigQuery.Table("foo", new()
    ///     {
    ///         DeletionProtection = false,
    ///         DatasetId = bar.DatasetId,
    ///         TableId = "job_query_table",
    ///     });
    /// 
    ///     var job = new Gcp.BigQuery.Job("job", new()
    ///     {
    ///         JobId = "job_query",
    ///         Labels = 
    ///         {
    ///             { "example-label", "example-value" },
    ///         },
    ///         Query = new Gcp.BigQuery.Inputs.JobQueryArgs
    ///         {
    ///             Query = "SELECT state FROM [lookerdata:cdc.project_tycho_reports]",
    ///             DestinationTable = new Gcp.BigQuery.Inputs.JobQueryDestinationTableArgs
    ///             {
    ///                 ProjectId = foo.Project,
    ///                 DatasetId = foo.DatasetId,
    ///                 TableId = foo.TableId,
    ///             },
    ///             AllowLargeResults = true,
    ///             FlattenResults = true,
    ///             ScriptOptions = new Gcp.BigQuery.Inputs.JobQueryScriptOptionsArgs
    ///             {
    ///                 KeyResultStatement = "LAST",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Bigquery Job Query Table Reference
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var bar = new Gcp.BigQuery.Dataset("bar", new()
    ///     {
    ///         DatasetId = "job_query_dataset",
    ///         FriendlyName = "test",
    ///         Description = "This is a test description",
    ///         Location = "US",
    ///     });
    /// 
    ///     var foo = new Gcp.BigQuery.Table("foo", new()
    ///     {
    ///         DeletionProtection = false,
    ///         DatasetId = bar.DatasetId,
    ///         TableId = "job_query_table",
    ///     });
    /// 
    ///     var job = new Gcp.BigQuery.Job("job", new()
    ///     {
    ///         JobId = "job_query",
    ///         Labels = 
    ///         {
    ///             { "example-label", "example-value" },
    ///         },
    ///         Query = new Gcp.BigQuery.Inputs.JobQueryArgs
    ///         {
    ///             Query = "SELECT state FROM [lookerdata:cdc.project_tycho_reports]",
    ///             DestinationTable = new Gcp.BigQuery.Inputs.JobQueryDestinationTableArgs
    ///             {
    ///                 TableId = foo.Id,
    ///             },
    ///             DefaultDataset = new Gcp.BigQuery.Inputs.JobQueryDefaultDatasetArgs
    ///             {
    ///                 DatasetId = bar.Id,
    ///             },
    ///             AllowLargeResults = true,
    ///             FlattenResults = true,
    ///             ScriptOptions = new Gcp.BigQuery.Inputs.JobQueryScriptOptionsArgs
    ///             {
    ///                 KeyResultStatement = "LAST",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Bigquery Job Load
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var bar = new Gcp.BigQuery.Dataset("bar", new()
    ///     {
    ///         DatasetId = "job_load_dataset",
    ///         FriendlyName = "test",
    ///         Description = "This is a test description",
    ///         Location = "US",
    ///     });
    /// 
    ///     var foo = new Gcp.BigQuery.Table("foo", new()
    ///     {
    ///         DeletionProtection = false,
    ///         DatasetId = bar.DatasetId,
    ///         TableId = "job_load_table",
    ///     });
    /// 
    ///     var job = new Gcp.BigQuery.Job("job", new()
    ///     {
    ///         JobId = "job_load",
    ///         Labels = 
    ///         {
    ///             { "my_job", "load" },
    ///         },
    ///         Load = new Gcp.BigQuery.Inputs.JobLoadArgs
    ///         {
    ///             SourceUris = new[]
    ///             {
    ///                 "gs://cloud-samples-data/bigquery/us-states/us-states-by-date.csv",
    ///             },
    ///             DestinationTable = new Gcp.BigQuery.Inputs.JobLoadDestinationTableArgs
    ///             {
    ///                 ProjectId = foo.Project,
    ///                 DatasetId = foo.DatasetId,
    ///                 TableId = foo.TableId,
    ///             },
    ///             SkipLeadingRows = 1,
    ///             SchemaUpdateOptions = new[]
    ///             {
    ///                 "ALLOW_FIELD_RELAXATION",
    ///                 "ALLOW_FIELD_ADDITION",
    ///             },
    ///             WriteDisposition = "WRITE_APPEND",
    ///             Autodetect = true,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Bigquery Job Load Parquet
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var testBucket = new Gcp.Storage.Bucket("testBucket", new()
    ///     {
    ///         Location = "US",
    ///         UniformBucketLevelAccess = true,
    ///     });
    /// 
    ///     var testBucketObject = new Gcp.Storage.BucketObject("testBucketObject", new()
    ///     {
    ///         Source = new FileAsset("./test-fixtures/test.parquet.gzip"),
    ///         Bucket = testBucket.Name,
    ///     });
    /// 
    ///     var testDataset = new Gcp.BigQuery.Dataset("testDataset", new()
    ///     {
    ///         DatasetId = "job_load_dataset",
    ///         FriendlyName = "test",
    ///         Description = "This is a test description",
    ///         Location = "US",
    ///     });
    /// 
    ///     var testTable = new Gcp.BigQuery.Table("testTable", new()
    ///     {
    ///         DeletionProtection = false,
    ///         TableId = "job_load_table",
    ///         DatasetId = testDataset.DatasetId,
    ///     });
    /// 
    ///     var job = new Gcp.BigQuery.Job("job", new()
    ///     {
    ///         JobId = "job_load",
    ///         Labels = 
    ///         {
    ///             { "my_job", "load" },
    ///         },
    ///         Load = new Gcp.BigQuery.Inputs.JobLoadArgs
    ///         {
    ///             SourceUris = new[]
    ///             {
    ///                 Output.Tuple(testBucketObject.Bucket, testBucketObject.Name).Apply(values =&gt;
    ///                 {
    ///                     var bucket = values.Item1;
    ///                     var name = values.Item2;
    ///                     return $"gs://{bucket}/{name}";
    ///                 }),
    ///             },
    ///             DestinationTable = new Gcp.BigQuery.Inputs.JobLoadDestinationTableArgs
    ///             {
    ///                 ProjectId = testTable.Project,
    ///                 DatasetId = testTable.DatasetId,
    ///                 TableId = testTable.TableId,
    ///             },
    ///             SchemaUpdateOptions = new[]
    ///             {
    ///                 "ALLOW_FIELD_RELAXATION",
    ///                 "ALLOW_FIELD_ADDITION",
    ///             },
    ///             WriteDisposition = "WRITE_APPEND",
    ///             SourceFormat = "PARQUET",
    ///             Autodetect = true,
    ///             ParquetOptions = new Gcp.BigQuery.Inputs.JobLoadParquetOptionsArgs
    ///             {
    ///                 EnumAsString = true,
    ///                 EnableListInference = true,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Bigquery Job Extract
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var source_oneDataset = new Gcp.BigQuery.Dataset("source-oneDataset", new()
    ///     {
    ///         DatasetId = "job_extract_dataset",
    ///         FriendlyName = "test",
    ///         Description = "This is a test description",
    ///         Location = "US",
    ///     });
    /// 
    ///     var source_oneTable = new Gcp.BigQuery.Table("source-oneTable", new()
    ///     {
    ///         DeletionProtection = false,
    ///         DatasetId = source_oneDataset.DatasetId,
    ///         TableId = "job_extract_table",
    ///         Schema = @"[
    ///   {
    ///     ""name"": ""name"",
    ///     ""type"": ""STRING"",
    ///     ""mode"": ""NULLABLE""
    ///   },
    ///   {
    ///     ""name"": ""post_abbr"",
    ///     ""type"": ""STRING"",
    ///     ""mode"": ""NULLABLE""
    ///   },
    ///   {
    ///     ""name"": ""date"",
    ///     ""type"": ""DATE"",
    ///     ""mode"": ""NULLABLE""
    ///   }
    /// ]
    /// ",
    ///     });
    /// 
    ///     var dest = new Gcp.Storage.Bucket("dest", new()
    ///     {
    ///         Location = "US",
    ///         ForceDestroy = true,
    ///     });
    /// 
    ///     var job = new Gcp.BigQuery.Job("job", new()
    ///     {
    ///         JobId = "job_extract",
    ///         Extract = new Gcp.BigQuery.Inputs.JobExtractArgs
    ///         {
    ///             DestinationUris = new[]
    ///             {
    ///                 dest.Url.Apply(url =&gt; $"{url}/extract"),
    ///             },
    ///             SourceTable = new Gcp.BigQuery.Inputs.JobExtractSourceTableArgs
    ///             {
    ///                 ProjectId = source_oneTable.Project,
    ///                 DatasetId = source_oneTable.DatasetId,
    ///                 TableId = source_oneTable.TableId,
    ///             },
    ///             DestinationFormat = "NEWLINE_DELIMITED_JSON",
    ///             Compression = "GZIP",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Job can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:bigquery/job:Job default projects/{{project}}/jobs/{{job_id}}/location/{{location}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:bigquery/job:Job default projects/{{project}}/jobs/{{job_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:bigquery/job:Job default {{project}}/{{job_id}}/{{location}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:bigquery/job:Job default {{job_id}}/{{location}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:bigquery/job:Job default {{project}}/{{job_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:bigquery/job:Job default {{job_id}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:bigquery/job:Job")]
    public partial class Job : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Copies a table.
        /// Structure is documented below.
        /// </summary>
        [Output("copy")]
        public Output<Outputs.JobCopy?> Copy { get; private set; } = null!;

        /// <summary>
        /// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
        /// clients and services.
        /// </summary>
        [Output("effectiveLabels")]
        public Output<ImmutableDictionary<string, string>> EffectiveLabels { get; private set; } = null!;

        /// <summary>
        /// Configures an extract job.
        /// Structure is documented below.
        /// </summary>
        [Output("extract")]
        public Output<Outputs.JobExtract?> Extract { get; private set; } = null!;

        /// <summary>
        /// The ID of the job. The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), or dashes (-). The maximum length is 1,024 characters.
        /// </summary>
        [Output("jobId")]
        public Output<string> JobId { get; private set; } = null!;

        /// <summary>
        /// Job timeout in milliseconds. If this time limit is exceeded, BigQuery may attempt to terminate the job.
        /// </summary>
        [Output("jobTimeoutMs")]
        public Output<string?> JobTimeoutMs { get; private set; } = null!;

        /// <summary>
        /// (Output)
        /// The type of the job.
        /// </summary>
        [Output("jobType")]
        public Output<string> JobType { get; private set; } = null!;

        /// <summary>
        /// The labels associated with this job. You can use these to organize and group your jobs.
        /// 
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// Configures a load job.
        /// Structure is documented below.
        /// </summary>
        [Output("load")]
        public Output<Outputs.JobLoad?> Load { get; private set; } = null!;

        /// <summary>
        /// The geographic location of the job. The default value is US.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// (Output)
        /// The combination of labels configured directly on the resource
        /// and default labels configured on the provider.
        /// </summary>
        [Output("pulumiLabels")]
        public Output<ImmutableDictionary<string, string>> PulumiLabels { get; private set; } = null!;

        /// <summary>
        /// SQL query text to execute. The useLegacySql field can be used to indicate whether the query uses legacy SQL or standard SQL.
        /// *NOTE*: queries containing [DML language](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-manipulation-language)
        /// (`DELETE`, `UPDATE`, `MERGE`, `INSERT`) must specify `create_disposition = ""` and `write_disposition = ""`.
        /// </summary>
        [Output("query")]
        public Output<Outputs.JobQuery?> Query { get; private set; } = null!;

        /// <summary>
        /// The status of this job. Examine this value when polling an asynchronous job to see if the job is complete.
        /// Structure is documented below.
        /// </summary>
        [Output("statuses")]
        public Output<ImmutableArray<Outputs.JobStatus>> Statuses { get; private set; } = null!;

        /// <summary>
        /// Email address of the user who ran the job.
        /// </summary>
        [Output("userEmail")]
        public Output<string> UserEmail { get; private set; } = null!;


        /// <summary>
        /// Create a Job resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Job(string name, JobArgs args, CustomResourceOptions? options = null)
            : base("gcp:bigquery/job:Job", name, args ?? new JobArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Job(string name, Input<string> id, JobState? state = null, CustomResourceOptions? options = null)
            : base("gcp:bigquery/job:Job", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "effectiveLabels",
                    "pulumiLabels",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Job resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Job Get(string name, Input<string> id, JobState? state = null, CustomResourceOptions? options = null)
        {
            return new Job(name, id, state, options);
        }
    }

    public sealed class JobArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Copies a table.
        /// Structure is documented below.
        /// </summary>
        [Input("copy")]
        public Input<Inputs.JobCopyArgs>? Copy { get; set; }

        /// <summary>
        /// Configures an extract job.
        /// Structure is documented below.
        /// </summary>
        [Input("extract")]
        public Input<Inputs.JobExtractArgs>? Extract { get; set; }

        /// <summary>
        /// The ID of the job. The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), or dashes (-). The maximum length is 1,024 characters.
        /// </summary>
        [Input("jobId", required: true)]
        public Input<string> JobId { get; set; } = null!;

        /// <summary>
        /// Job timeout in milliseconds. If this time limit is exceeded, BigQuery may attempt to terminate the job.
        /// </summary>
        [Input("jobTimeoutMs")]
        public Input<string>? JobTimeoutMs { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The labels associated with this job. You can use these to organize and group your jobs.
        /// 
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Configures a load job.
        /// Structure is documented below.
        /// </summary>
        [Input("load")]
        public Input<Inputs.JobLoadArgs>? Load { get; set; }

        /// <summary>
        /// The geographic location of the job. The default value is US.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// SQL query text to execute. The useLegacySql field can be used to indicate whether the query uses legacy SQL or standard SQL.
        /// *NOTE*: queries containing [DML language](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-manipulation-language)
        /// (`DELETE`, `UPDATE`, `MERGE`, `INSERT`) must specify `create_disposition = ""` and `write_disposition = ""`.
        /// </summary>
        [Input("query")]
        public Input<Inputs.JobQueryArgs>? Query { get; set; }

        public JobArgs()
        {
        }
        public static new JobArgs Empty => new JobArgs();
    }

    public sealed class JobState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Copies a table.
        /// Structure is documented below.
        /// </summary>
        [Input("copy")]
        public Input<Inputs.JobCopyGetArgs>? Copy { get; set; }

        [Input("effectiveLabels")]
        private InputMap<string>? _effectiveLabels;

        /// <summary>
        /// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
        /// clients and services.
        /// </summary>
        public InputMap<string> EffectiveLabels
        {
            get => _effectiveLabels ?? (_effectiveLabels = new InputMap<string>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, string>());
                _effectiveLabels = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        /// <summary>
        /// Configures an extract job.
        /// Structure is documented below.
        /// </summary>
        [Input("extract")]
        public Input<Inputs.JobExtractGetArgs>? Extract { get; set; }

        /// <summary>
        /// The ID of the job. The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), or dashes (-). The maximum length is 1,024 characters.
        /// </summary>
        [Input("jobId")]
        public Input<string>? JobId { get; set; }

        /// <summary>
        /// Job timeout in milliseconds. If this time limit is exceeded, BigQuery may attempt to terminate the job.
        /// </summary>
        [Input("jobTimeoutMs")]
        public Input<string>? JobTimeoutMs { get; set; }

        /// <summary>
        /// (Output)
        /// The type of the job.
        /// </summary>
        [Input("jobType")]
        public Input<string>? JobType { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The labels associated with this job. You can use these to organize and group your jobs.
        /// 
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Configures a load job.
        /// Structure is documented below.
        /// </summary>
        [Input("load")]
        public Input<Inputs.JobLoadGetArgs>? Load { get; set; }

        /// <summary>
        /// The geographic location of the job. The default value is US.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("pulumiLabels")]
        private InputMap<string>? _pulumiLabels;

        /// <summary>
        /// (Output)
        /// The combination of labels configured directly on the resource
        /// and default labels configured on the provider.
        /// </summary>
        public InputMap<string> PulumiLabels
        {
            get => _pulumiLabels ?? (_pulumiLabels = new InputMap<string>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, string>());
                _pulumiLabels = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        /// <summary>
        /// SQL query text to execute. The useLegacySql field can be used to indicate whether the query uses legacy SQL or standard SQL.
        /// *NOTE*: queries containing [DML language](https://cloud.google.com/bigquery/docs/reference/standard-sql/data-manipulation-language)
        /// (`DELETE`, `UPDATE`, `MERGE`, `INSERT`) must specify `create_disposition = ""` and `write_disposition = ""`.
        /// </summary>
        [Input("query")]
        public Input<Inputs.JobQueryGetArgs>? Query { get; set; }

        [Input("statuses")]
        private InputList<Inputs.JobStatusGetArgs>? _statuses;

        /// <summary>
        /// The status of this job. Examine this value when polling an asynchronous job to see if the job is complete.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.JobStatusGetArgs> Statuses
        {
            get => _statuses ?? (_statuses = new InputList<Inputs.JobStatusGetArgs>());
            set => _statuses = value;
        }

        /// <summary>
        /// Email address of the user who ran the job.
        /// </summary>
        [Input("userEmail")]
        public Input<string>? UserEmail { get; set; }

        public JobState()
        {
        }
        public static new JobState Empty => new JobState();
    }
}
