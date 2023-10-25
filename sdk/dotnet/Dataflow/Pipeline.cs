// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Dataflow
{
    /// <summary>
    /// The main pipeline entity and all the necessary metadata for launching and managing linked jobs.
    /// 
    /// To get more information about Pipeline, see:
    /// 
    /// * [API documentation](https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/dataflow)
    /// 
    /// ## Example Usage
    /// ### Data Pipeline Pipeline
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var serviceAccount = new Gcp.ServiceAccount.Account("serviceAccount", new()
    ///     {
    ///         AccountId = "my-account",
    ///         DisplayName = "Service Account",
    ///     });
    /// 
    ///     var primary = new Gcp.Dataflow.Pipeline("primary", new()
    ///     {
    ///         DisplayName = "my-pipeline",
    ///         Type = "PIPELINE_TYPE_BATCH",
    ///         State = "STATE_ACTIVE",
    ///         Region = "us-central1",
    ///         Workload = new Gcp.Dataflow.Inputs.PipelineWorkloadArgs
    ///         {
    ///             DataflowLaunchTemplateRequest = new Gcp.Dataflow.Inputs.PipelineWorkloadDataflowLaunchTemplateRequestArgs
    ///             {
    ///                 ProjectId = "my-project",
    ///                 GcsPath = "gs://my-bucket/path",
    ///                 LaunchParameters = new Gcp.Dataflow.Inputs.PipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersArgs
    ///                 {
    ///                     JobName = "my-job",
    ///                     Parameters = 
    ///                     {
    ///                         { "name", "wrench" },
    ///                     },
    ///                     Environment = new Gcp.Dataflow.Inputs.PipelineWorkloadDataflowLaunchTemplateRequestLaunchParametersEnvironmentArgs
    ///                     {
    ///                         NumWorkers = 5,
    ///                         MaxWorkers = 5,
    ///                         Zone = "us-centra1-a",
    ///                         ServiceAccountEmail = serviceAccount.Email,
    ///                         Network = "default",
    ///                         TempLocation = "gs://my-bucket/tmp_dir",
    ///                         BypassTempDirValidation = false,
    ///                         MachineType = "E2",
    ///                         AdditionalUserLabels = 
    ///                         {
    ///                             { "context", "test" },
    ///                         },
    ///                         WorkerRegion = "us-central1",
    ///                         WorkerZone = "us-central1-a",
    ///                         EnableStreamingEngine = false,
    ///                     },
    ///                     Update = false,
    ///                     TransformNameMapping = 
    ///                     {
    ///                         { "name", "wrench" },
    ///                     },
    ///                 },
    ///                 Location = "us-central1",
    ///             },
    ///         },
    ///         ScheduleInfo = new Gcp.Dataflow.Inputs.PipelineScheduleInfoArgs
    ///         {
    ///             Schedule = "* */2 * * *",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Pipeline can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:dataflow/pipeline:Pipeline default projects/{{project}}/locations/{{region}}/pipelines/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:dataflow/pipeline:Pipeline default {{project}}/{{region}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:dataflow/pipeline:Pipeline default {{region}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:dataflow/pipeline:Pipeline default {{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:dataflow/pipeline:Pipeline")]
    public partial class Pipeline : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The timestamp when the pipeline was initially created. Set by the Data Pipelines service.
        /// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The display name of the pipeline. It can contain only letters ([A-Za-z]), numbers ([0-9]), hyphens (-), and underscores (_).
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Number of jobs.
        /// </summary>
        [Output("jobCount")]
        public Output<int> JobCount { get; private set; } = null!;

        /// <summary>
        /// The timestamp when the pipeline was last modified. Set by the Data Pipelines service.
        /// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Output("lastUpdateTime")]
        public Output<string> LastUpdateTime { get; private set; } = null!;

        /// <summary>
        /// "The pipeline name. For example': 'projects/PROJECT_ID/locations/LOCATION_ID/pipelines/PIPELINE_ID."
        /// "- PROJECT_ID can contain letters ([A-Za-z]), numbers ([0-9]), hyphens (-), colons (:), and periods (.). For more information, see Identifying projects."
        /// "LOCATION_ID is the canonical ID for the pipeline's location. The list of available locations can be obtained by calling google.cloud.location.Locations.ListLocations. Note that the Data Pipelines service is not available in all regions. It depends on Cloud Scheduler, an App Engine application, so it's only available in App Engine regions."
        /// "PIPELINE_ID is the ID of the pipeline. Must be unique for the selected project and location."
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The sources of the pipeline (for example, Dataplex). The keys and values are set by the corresponding sources during pipeline creation.
        /// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        /// </summary>
        [Output("pipelineSources")]
        public Output<ImmutableDictionary<string, string>?> PipelineSources { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// A reference to the region
        /// </summary>
        [Output("region")]
        public Output<string?> Region { get; private set; } = null!;

        /// <summary>
        /// Internal scheduling information for a pipeline. If this information is provided, periodic jobs will be created per the schedule. If not, users are responsible for creating jobs externally.
        /// https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#schedulespec
        /// Structure is documented below.
        /// </summary>
        [Output("scheduleInfo")]
        public Output<Outputs.PipelineScheduleInfo?> ScheduleInfo { get; private set; } = null!;

        /// <summary>
        /// Optional. A service account email to be used with the Cloud Scheduler job. If not specified, the default compute engine service account will be used.
        /// </summary>
        [Output("schedulerServiceAccountEmail")]
        public Output<string?> SchedulerServiceAccountEmail { get; private set; } = null!;

        /// <summary>
        /// The state of the pipeline. When the pipeline is created, the state is set to 'PIPELINE_STATE_ACTIVE' by default. State changes can be requested by setting the state to stopping, paused, or resuming. State cannot be changed through pipelines.patch requests.
        /// https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#state
        /// Possible values are: `STATE_UNSPECIFIED`, `STATE_RESUMING`, `STATE_ACTIVE`, `STATE_STOPPING`, `STATE_ARCHIVED`, `STATE_PAUSED`.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The type of the pipeline. This field affects the scheduling of the pipeline and the type of metrics to show for the pipeline.
        /// https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#pipelinetype
        /// Possible values are: `PIPELINE_TYPE_UNSPECIFIED`, `PIPELINE_TYPE_BATCH`, `PIPELINE_TYPE_STREAMING`.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Workload information for creating new jobs.
        /// https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#workload
        /// Structure is documented below.
        /// </summary>
        [Output("workload")]
        public Output<Outputs.PipelineWorkload?> Workload { get; private set; } = null!;


        /// <summary>
        /// Create a Pipeline resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Pipeline(string name, PipelineArgs args, CustomResourceOptions? options = null)
            : base("gcp:dataflow/pipeline:Pipeline", name, args ?? new PipelineArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Pipeline(string name, Input<string> id, PipelineState? state = null, CustomResourceOptions? options = null)
            : base("gcp:dataflow/pipeline:Pipeline", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Pipeline resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Pipeline Get(string name, Input<string> id, PipelineState? state = null, CustomResourceOptions? options = null)
        {
            return new Pipeline(name, id, state, options);
        }
    }

    public sealed class PipelineArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The display name of the pipeline. It can contain only letters ([A-Za-z]), numbers ([0-9]), hyphens (-), and underscores (_).
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// "The pipeline name. For example': 'projects/PROJECT_ID/locations/LOCATION_ID/pipelines/PIPELINE_ID."
        /// "- PROJECT_ID can contain letters ([A-Za-z]), numbers ([0-9]), hyphens (-), colons (:), and periods (.). For more information, see Identifying projects."
        /// "LOCATION_ID is the canonical ID for the pipeline's location. The list of available locations can be obtained by calling google.cloud.location.Locations.ListLocations. Note that the Data Pipelines service is not available in all regions. It depends on Cloud Scheduler, an App Engine application, so it's only available in App Engine regions."
        /// "PIPELINE_ID is the ID of the pipeline. Must be unique for the selected project and location."
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("pipelineSources")]
        private InputMap<string>? _pipelineSources;

        /// <summary>
        /// The sources of the pipeline (for example, Dataplex). The keys and values are set by the corresponding sources during pipeline creation.
        /// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        /// </summary>
        public InputMap<string> PipelineSources
        {
            get => _pipelineSources ?? (_pipelineSources = new InputMap<string>());
            set => _pipelineSources = value;
        }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// A reference to the region
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Internal scheduling information for a pipeline. If this information is provided, periodic jobs will be created per the schedule. If not, users are responsible for creating jobs externally.
        /// https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#schedulespec
        /// Structure is documented below.
        /// </summary>
        [Input("scheduleInfo")]
        public Input<Inputs.PipelineScheduleInfoArgs>? ScheduleInfo { get; set; }

        /// <summary>
        /// Optional. A service account email to be used with the Cloud Scheduler job. If not specified, the default compute engine service account will be used.
        /// </summary>
        [Input("schedulerServiceAccountEmail")]
        public Input<string>? SchedulerServiceAccountEmail { get; set; }

        /// <summary>
        /// The state of the pipeline. When the pipeline is created, the state is set to 'PIPELINE_STATE_ACTIVE' by default. State changes can be requested by setting the state to stopping, paused, or resuming. State cannot be changed through pipelines.patch requests.
        /// https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#state
        /// Possible values are: `STATE_UNSPECIFIED`, `STATE_RESUMING`, `STATE_ACTIVE`, `STATE_STOPPING`, `STATE_ARCHIVED`, `STATE_PAUSED`.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("state", required: true)]
        public Input<string> State { get; set; } = null!;

        /// <summary>
        /// The type of the pipeline. This field affects the scheduling of the pipeline and the type of metrics to show for the pipeline.
        /// https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#pipelinetype
        /// Possible values are: `PIPELINE_TYPE_UNSPECIFIED`, `PIPELINE_TYPE_BATCH`, `PIPELINE_TYPE_STREAMING`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Workload information for creating new jobs.
        /// https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#workload
        /// Structure is documented below.
        /// </summary>
        [Input("workload")]
        public Input<Inputs.PipelineWorkloadArgs>? Workload { get; set; }

        public PipelineArgs()
        {
        }
        public static new PipelineArgs Empty => new PipelineArgs();
    }

    public sealed class PipelineState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The timestamp when the pipeline was initially created. Set by the Data Pipelines service.
        /// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// The display name of the pipeline. It can contain only letters ([A-Za-z]), numbers ([0-9]), hyphens (-), and underscores (_).
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Number of jobs.
        /// </summary>
        [Input("jobCount")]
        public Input<int>? JobCount { get; set; }

        /// <summary>
        /// The timestamp when the pipeline was last modified. Set by the Data Pipelines service.
        /// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Input("lastUpdateTime")]
        public Input<string>? LastUpdateTime { get; set; }

        /// <summary>
        /// "The pipeline name. For example': 'projects/PROJECT_ID/locations/LOCATION_ID/pipelines/PIPELINE_ID."
        /// "- PROJECT_ID can contain letters ([A-Za-z]), numbers ([0-9]), hyphens (-), colons (:), and periods (.). For more information, see Identifying projects."
        /// "LOCATION_ID is the canonical ID for the pipeline's location. The list of available locations can be obtained by calling google.cloud.location.Locations.ListLocations. Note that the Data Pipelines service is not available in all regions. It depends on Cloud Scheduler, an App Engine application, so it's only available in App Engine regions."
        /// "PIPELINE_ID is the ID of the pipeline. Must be unique for the selected project and location."
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("pipelineSources")]
        private InputMap<string>? _pipelineSources;

        /// <summary>
        /// The sources of the pipeline (for example, Dataplex). The keys and values are set by the corresponding sources during pipeline creation.
        /// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
        /// </summary>
        public InputMap<string> PipelineSources
        {
            get => _pipelineSources ?? (_pipelineSources = new InputMap<string>());
            set => _pipelineSources = value;
        }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// A reference to the region
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Internal scheduling information for a pipeline. If this information is provided, periodic jobs will be created per the schedule. If not, users are responsible for creating jobs externally.
        /// https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#schedulespec
        /// Structure is documented below.
        /// </summary>
        [Input("scheduleInfo")]
        public Input<Inputs.PipelineScheduleInfoGetArgs>? ScheduleInfo { get; set; }

        /// <summary>
        /// Optional. A service account email to be used with the Cloud Scheduler job. If not specified, the default compute engine service account will be used.
        /// </summary>
        [Input("schedulerServiceAccountEmail")]
        public Input<string>? SchedulerServiceAccountEmail { get; set; }

        /// <summary>
        /// The state of the pipeline. When the pipeline is created, the state is set to 'PIPELINE_STATE_ACTIVE' by default. State changes can be requested by setting the state to stopping, paused, or resuming. State cannot be changed through pipelines.patch requests.
        /// https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#state
        /// Possible values are: `STATE_UNSPECIFIED`, `STATE_RESUMING`, `STATE_ACTIVE`, `STATE_STOPPING`, `STATE_ARCHIVED`, `STATE_PAUSED`.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The type of the pipeline. This field affects the scheduling of the pipeline and the type of metrics to show for the pipeline.
        /// https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#pipelinetype
        /// Possible values are: `PIPELINE_TYPE_UNSPECIFIED`, `PIPELINE_TYPE_BATCH`, `PIPELINE_TYPE_STREAMING`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Workload information for creating new jobs.
        /// https://cloud.google.com/dataflow/docs/reference/data-pipelines/rest/v1/projects.locations.pipelines#workload
        /// Structure is documented below.
        /// </summary>
        [Input("workload")]
        public Input<Inputs.PipelineWorkloadGetArgs>? Workload { get; set; }

        public PipelineState()
        {
        }
        public static new PipelineState Empty => new PipelineState();
    }
}
