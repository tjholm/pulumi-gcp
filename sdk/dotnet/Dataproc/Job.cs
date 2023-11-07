// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Dataproc
{
    /// <summary>
    /// Manages a job resource within a Dataproc cluster within GCE. For more information see
    /// [the official dataproc documentation](https://cloud.google.com/dataproc/).
    /// 
    /// !&gt; **Note:** This resource does not support 'update' and changing any attributes will cause the resource to be recreated.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var mycluster = new Gcp.Dataproc.Cluster("mycluster", new()
    ///     {
    ///         Region = "us-central1",
    ///     });
    /// 
    ///     // Submit an example spark job to a dataproc cluster
    ///     var spark = new Gcp.Dataproc.Job("spark", new()
    ///     {
    ///         Region = mycluster.Region,
    ///         ForceDelete = true,
    ///         Placement = new Gcp.Dataproc.Inputs.JobPlacementArgs
    ///         {
    ///             ClusterName = mycluster.Name,
    ///         },
    ///         SparkConfig = new Gcp.Dataproc.Inputs.JobSparkConfigArgs
    ///         {
    ///             MainClass = "org.apache.spark.examples.SparkPi",
    ///             JarFileUris = new[]
    ///             {
    ///                 "file:///usr/lib/spark/examples/jars/spark-examples.jar",
    ///             },
    ///             Args = new[]
    ///             {
    ///                 "1000",
    ///             },
    ///             Properties = 
    ///             {
    ///                 { "spark.logConf", "true" },
    ///             },
    ///             LoggingConfig = new Gcp.Dataproc.Inputs.JobSparkConfigLoggingConfigArgs
    ///             {
    ///                 DriverLogLevels = 
    ///                 {
    ///                     { "root", "INFO" },
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     // Submit an example pyspark job to a dataproc cluster
    ///     var pyspark = new Gcp.Dataproc.Job("pyspark", new()
    ///     {
    ///         Region = mycluster.Region,
    ///         ForceDelete = true,
    ///         Placement = new Gcp.Dataproc.Inputs.JobPlacementArgs
    ///         {
    ///             ClusterName = mycluster.Name,
    ///         },
    ///         PysparkConfig = new Gcp.Dataproc.Inputs.JobPysparkConfigArgs
    ///         {
    ///             MainPythonFileUri = "gs://dataproc-examples-2f10d78d114f6aaec76462e3c310f31f/src/pyspark/hello-world/hello-world.py",
    ///             Properties = 
    ///             {
    ///                 { "spark.logConf", "true" },
    ///             },
    ///         },
    ///     });
    /// 
    ///     return new Dictionary&lt;string, object?&gt;
    ///     {
    ///         ["sparkStatus"] = spark.Statuses.Apply(statuses =&gt; statuses[0].State),
    ///         ["pysparkStatus"] = pyspark.Statuses.Apply(statuses =&gt; statuses[0].State),
    ///     };
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// This resource does not support import.
    /// </summary>
    [GcpResourceType("gcp:dataproc/job:Job")]
    public partial class Job : global::Pulumi.CustomResource
    {
        /// <summary>
        /// If present, the location of miscellaneous control files which may be used as part of job setup and handling. If not present, control files may be placed in the same location as driver_output_uri.
        /// </summary>
        [Output("driverControlsFilesUri")]
        public Output<string> DriverControlsFilesUri { get; private set; } = null!;

        /// <summary>
        /// A URI pointing to the location of the stdout of the job's driver program.
        /// </summary>
        [Output("driverOutputResourceUri")]
        public Output<string> DriverOutputResourceUri { get; private set; } = null!;

        /// <summary>
        /// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
        /// clients and services.
        /// </summary>
        [Output("effectiveLabels")]
        public Output<ImmutableDictionary<string, string>> EffectiveLabels { get; private set; } = null!;

        /// <summary>
        /// By default, you can only delete inactive jobs within
        /// Dataproc. Setting this to true, and calling destroy, will ensure that the
        /// job is first cancelled before issuing the delete.
        /// </summary>
        [Output("forceDelete")]
        public Output<bool?> ForceDelete { get; private set; } = null!;

        /// <summary>
        /// The config of Hadoop job
        /// </summary>
        [Output("hadoopConfig")]
        public Output<Outputs.JobHadoopConfig?> HadoopConfig { get; private set; } = null!;

        /// <summary>
        /// The config of hive job
        /// </summary>
        [Output("hiveConfig")]
        public Output<Outputs.JobHiveConfig?> HiveConfig { get; private set; } = null!;

        /// <summary>
        /// The list of labels (key/value pairs) to add to the job.
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field 'effective_labels' for all of the labels present on the resource.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// The config of pag job.
        /// </summary>
        [Output("pigConfig")]
        public Output<Outputs.JobPigConfig?> PigConfig { get; private set; } = null!;

        /// <summary>
        /// The config of job placement.
        /// </summary>
        [Output("placement")]
        public Output<Outputs.JobPlacement> Placement { get; private set; } = null!;

        /// <summary>
        /// The config of presto job
        /// </summary>
        [Output("prestoConfig")]
        public Output<Outputs.JobPrestoConfig?> PrestoConfig { get; private set; } = null!;

        /// <summary>
        /// The project in which the `cluster` can be found and jobs
        /// subsequently run against. If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The combination of labels configured directly on the resource and default labels configured on the provider.
        /// </summary>
        [Output("pulumiLabels")]
        public Output<ImmutableDictionary<string, string>> PulumiLabels { get; private set; } = null!;

        /// <summary>
        /// The config of pySpark job.
        /// </summary>
        [Output("pysparkConfig")]
        public Output<Outputs.JobPysparkConfig?> PysparkConfig { get; private set; } = null!;

        /// <summary>
        /// The reference of the job
        /// </summary>
        [Output("reference")]
        public Output<Outputs.JobReference> Reference { get; private set; } = null!;

        /// <summary>
        /// The Cloud Dataproc region. This essentially determines which clusters are available
        /// for this job to be submitted to. If not specified, defaults to `global`.
        /// </summary>
        [Output("region")]
        public Output<string?> Region { get; private set; } = null!;

        /// <summary>
        /// Optional. Job scheduling configuration.
        /// </summary>
        [Output("scheduling")]
        public Output<Outputs.JobScheduling?> Scheduling { get; private set; } = null!;

        /// <summary>
        /// The config of the Spark job.
        /// </summary>
        [Output("sparkConfig")]
        public Output<Outputs.JobSparkConfig?> SparkConfig { get; private set; } = null!;

        /// <summary>
        /// The config of SparkSql job
        /// </summary>
        [Output("sparksqlConfig")]
        public Output<Outputs.JobSparksqlConfig?> SparksqlConfig { get; private set; } = null!;

        /// <summary>
        /// The status of the job.
        /// </summary>
        [Output("statuses")]
        public Output<ImmutableArray<Outputs.JobStatus>> Statuses { get; private set; } = null!;


        /// <summary>
        /// Create a Job resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Job(string name, JobArgs args, CustomResourceOptions? options = null)
            : base("gcp:dataproc/job:Job", name, args ?? new JobArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Job(string name, Input<string> id, JobState? state = null, CustomResourceOptions? options = null)
            : base("gcp:dataproc/job:Job", name, state, MakeResourceOptions(options, id))
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
        /// By default, you can only delete inactive jobs within
        /// Dataproc. Setting this to true, and calling destroy, will ensure that the
        /// job is first cancelled before issuing the delete.
        /// </summary>
        [Input("forceDelete")]
        public Input<bool>? ForceDelete { get; set; }

        /// <summary>
        /// The config of Hadoop job
        /// </summary>
        [Input("hadoopConfig")]
        public Input<Inputs.JobHadoopConfigArgs>? HadoopConfig { get; set; }

        /// <summary>
        /// The config of hive job
        /// </summary>
        [Input("hiveConfig")]
        public Input<Inputs.JobHiveConfigArgs>? HiveConfig { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The list of labels (key/value pairs) to add to the job.
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field 'effective_labels' for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The config of pag job.
        /// </summary>
        [Input("pigConfig")]
        public Input<Inputs.JobPigConfigArgs>? PigConfig { get; set; }

        /// <summary>
        /// The config of job placement.
        /// </summary>
        [Input("placement", required: true)]
        public Input<Inputs.JobPlacementArgs> Placement { get; set; } = null!;

        /// <summary>
        /// The config of presto job
        /// </summary>
        [Input("prestoConfig")]
        public Input<Inputs.JobPrestoConfigArgs>? PrestoConfig { get; set; }

        /// <summary>
        /// The project in which the `cluster` can be found and jobs
        /// subsequently run against. If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The config of pySpark job.
        /// </summary>
        [Input("pysparkConfig")]
        public Input<Inputs.JobPysparkConfigArgs>? PysparkConfig { get; set; }

        /// <summary>
        /// The reference of the job
        /// </summary>
        [Input("reference")]
        public Input<Inputs.JobReferenceArgs>? Reference { get; set; }

        /// <summary>
        /// The Cloud Dataproc region. This essentially determines which clusters are available
        /// for this job to be submitted to. If not specified, defaults to `global`.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Optional. Job scheduling configuration.
        /// </summary>
        [Input("scheduling")]
        public Input<Inputs.JobSchedulingArgs>? Scheduling { get; set; }

        /// <summary>
        /// The config of the Spark job.
        /// </summary>
        [Input("sparkConfig")]
        public Input<Inputs.JobSparkConfigArgs>? SparkConfig { get; set; }

        /// <summary>
        /// The config of SparkSql job
        /// </summary>
        [Input("sparksqlConfig")]
        public Input<Inputs.JobSparksqlConfigArgs>? SparksqlConfig { get; set; }

        public JobArgs()
        {
        }
        public static new JobArgs Empty => new JobArgs();
    }

    public sealed class JobState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If present, the location of miscellaneous control files which may be used as part of job setup and handling. If not present, control files may be placed in the same location as driver_output_uri.
        /// </summary>
        [Input("driverControlsFilesUri")]
        public Input<string>? DriverControlsFilesUri { get; set; }

        /// <summary>
        /// A URI pointing to the location of the stdout of the job's driver program.
        /// </summary>
        [Input("driverOutputResourceUri")]
        public Input<string>? DriverOutputResourceUri { get; set; }

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
        /// By default, you can only delete inactive jobs within
        /// Dataproc. Setting this to true, and calling destroy, will ensure that the
        /// job is first cancelled before issuing the delete.
        /// </summary>
        [Input("forceDelete")]
        public Input<bool>? ForceDelete { get; set; }

        /// <summary>
        /// The config of Hadoop job
        /// </summary>
        [Input("hadoopConfig")]
        public Input<Inputs.JobHadoopConfigGetArgs>? HadoopConfig { get; set; }

        /// <summary>
        /// The config of hive job
        /// </summary>
        [Input("hiveConfig")]
        public Input<Inputs.JobHiveConfigGetArgs>? HiveConfig { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The list of labels (key/value pairs) to add to the job.
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field 'effective_labels' for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The config of pag job.
        /// </summary>
        [Input("pigConfig")]
        public Input<Inputs.JobPigConfigGetArgs>? PigConfig { get; set; }

        /// <summary>
        /// The config of job placement.
        /// </summary>
        [Input("placement")]
        public Input<Inputs.JobPlacementGetArgs>? Placement { get; set; }

        /// <summary>
        /// The config of presto job
        /// </summary>
        [Input("prestoConfig")]
        public Input<Inputs.JobPrestoConfigGetArgs>? PrestoConfig { get; set; }

        /// <summary>
        /// The project in which the `cluster` can be found and jobs
        /// subsequently run against. If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("pulumiLabels")]
        private InputMap<string>? _pulumiLabels;

        /// <summary>
        /// The combination of labels configured directly on the resource and default labels configured on the provider.
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
        /// The config of pySpark job.
        /// </summary>
        [Input("pysparkConfig")]
        public Input<Inputs.JobPysparkConfigGetArgs>? PysparkConfig { get; set; }

        /// <summary>
        /// The reference of the job
        /// </summary>
        [Input("reference")]
        public Input<Inputs.JobReferenceGetArgs>? Reference { get; set; }

        /// <summary>
        /// The Cloud Dataproc region. This essentially determines which clusters are available
        /// for this job to be submitted to. If not specified, defaults to `global`.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Optional. Job scheduling configuration.
        /// </summary>
        [Input("scheduling")]
        public Input<Inputs.JobSchedulingGetArgs>? Scheduling { get; set; }

        /// <summary>
        /// The config of the Spark job.
        /// </summary>
        [Input("sparkConfig")]
        public Input<Inputs.JobSparkConfigGetArgs>? SparkConfig { get; set; }

        /// <summary>
        /// The config of SparkSql job
        /// </summary>
        [Input("sparksqlConfig")]
        public Input<Inputs.JobSparksqlConfigGetArgs>? SparksqlConfig { get; set; }

        [Input("statuses")]
        private InputList<Inputs.JobStatusGetArgs>? _statuses;

        /// <summary>
        /// The status of the job.
        /// </summary>
        public InputList<Inputs.JobStatusGetArgs> Statuses
        {
            get => _statuses ?? (_statuses = new InputList<Inputs.JobStatusGetArgs>());
            set => _statuses = value;
        }

        public JobState()
        {
        }
        public static new JobState Empty => new JobState();
    }
}
