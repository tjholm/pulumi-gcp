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
    /// A Workflow Template is a reusable workflow configuration. It defines a graph of jobs with information on where to run those jobs.
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
    ///     var template = new Gcp.Dataproc.WorkflowTemplate("template", new()
    ///     {
    ///         Jobs = new[]
    ///         {
    ///             new Gcp.Dataproc.Inputs.WorkflowTemplateJobArgs
    ///             {
    ///                 SparkJob = new Gcp.Dataproc.Inputs.WorkflowTemplateJobSparkJobArgs
    ///                 {
    ///                     MainClass = "SomeClass",
    ///                 },
    ///                 StepId = "someJob",
    ///             },
    ///             new Gcp.Dataproc.Inputs.WorkflowTemplateJobArgs
    ///             {
    ///                 PrerequisiteStepIds = new[]
    ///                 {
    ///                     "someJob",
    ///                 },
    ///                 PrestoJob = new Gcp.Dataproc.Inputs.WorkflowTemplateJobPrestoJobArgs
    ///                 {
    ///                     QueryFileUri = "someuri",
    ///                 },
    ///                 StepId = "otherJob",
    ///             },
    ///         },
    ///         Location = "us-central1",
    ///         Placement = new Gcp.Dataproc.Inputs.WorkflowTemplatePlacementArgs
    ///         {
    ///             ManagedCluster = new Gcp.Dataproc.Inputs.WorkflowTemplatePlacementManagedClusterArgs
    ///             {
    ///                 ClusterName = "my-cluster",
    ///                 Config = new Gcp.Dataproc.Inputs.WorkflowTemplatePlacementManagedClusterConfigArgs
    ///                 {
    ///                     GceClusterConfig = new Gcp.Dataproc.Inputs.WorkflowTemplatePlacementManagedClusterConfigGceClusterConfigArgs
    ///                     {
    ///                         Tags = new[]
    ///                         {
    ///                             "foo",
    ///                             "bar",
    ///                         },
    ///                         Zone = "us-central1-a",
    ///                     },
    ///                     MasterConfig = new Gcp.Dataproc.Inputs.WorkflowTemplatePlacementManagedClusterConfigMasterConfigArgs
    ///                     {
    ///                         DiskConfig = new Gcp.Dataproc.Inputs.WorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfigArgs
    ///                         {
    ///                             BootDiskSizeGb = 15,
    ///                             BootDiskType = "pd-ssd",
    ///                         },
    ///                         MachineType = "n1-standard-1",
    ///                         NumInstances = 1,
    ///                     },
    ///                     SecondaryWorkerConfig = new Gcp.Dataproc.Inputs.WorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigArgs
    ///                     {
    ///                         NumInstances = 2,
    ///                     },
    ///                     SoftwareConfig = new Gcp.Dataproc.Inputs.WorkflowTemplatePlacementManagedClusterConfigSoftwareConfigArgs
    ///                     {
    ///                         ImageVersion = "2.0.35-debian10",
    ///                     },
    ///                     WorkerConfig = new Gcp.Dataproc.Inputs.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigArgs
    ///                     {
    ///                         DiskConfig = new Gcp.Dataproc.Inputs.WorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfigArgs
    ///                         {
    ///                             BootDiskSizeGb = 10,
    ///                             NumLocalSsds = 2,
    ///                         },
    ///                         MachineType = "n1-standard-2",
    ///                         NumInstances = 3,
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// WorkflowTemplate can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:dataproc/workflowTemplate:WorkflowTemplate default projects/{{project}}/locations/{{location}}/workflowTemplates/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:dataproc/workflowTemplate:WorkflowTemplate default {{project}}/{{location}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:dataproc/workflowTemplate:WorkflowTemplate default {{location}}/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:dataproc/workflowTemplate:WorkflowTemplate")]
    public partial class WorkflowTemplate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Output only. The time template was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// (Beta only) Optional. Timeout duration for the DAG of jobs. You can use "s", "m", "h", and "d" suffixes for second, minute, hour, and day duration values, respectively. The timeout duration must be from 10 minutes ("10m") to 24 hours ("24h" or "1d"). The timer begins when the first job is submitted. If the workflow is running at the end of the timeout period, any remaining jobs are cancelled, the workflow is ended, and if the workflow was running on a (/dataproc/docs/concepts/workflows/using-workflows#configuring_or_selecting_a_cluster), the cluster is deleted.
        /// </summary>
        [Output("dagTimeout")]
        public Output<string?> DagTimeout { get; private set; } = null!;

        /// <summary>
        /// Required. The Directed Acyclic Graph of Jobs to submit.
        /// </summary>
        [Output("jobs")]
        public Output<ImmutableArray<Outputs.WorkflowTemplateJob>> Jobs { get; private set; } = null!;

        /// <summary>
        /// Optional. The labels to associate with this job. Label keys must be between 1 and 63 characters long, and must conform to the following regular expression: {0,63} No more than 32 labels can be associated with a given job.
        /// (Optional)
        /// Optional. The labels to associate with this template. These labels will be propagated to all jobs and clusters created by the workflow instance. Label **keys** must contain 1 to 63 characters, and must conform to (https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be associated with a template.
        /// (Optional)
        /// Optional. The labels to associate with this cluster. Label keys must be between 1 and 63 characters long, and must conform to the following PCRE regular expression: {0,63} No more than 32 labels can be associated with a given cluster.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// The location for the resource
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Output only. The resource name of the workflow template, as described in https://cloud.google.com/apis/design/resource_names. * For `projects.regions.workflowTemplates`, the resource name of the template has the following format: `projects/{project_id}/regions/{region}/workflowTemplates/{template_id}` * For `projects.locations.workflowTemplates`, the resource name of the template has the following format: `projects/{project_id}/locations/{location}/workflowTemplates/{template_id}`
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Optional. Template parameters whose values are substituted into the template. Values for parameters must be provided when the template is instantiated.
        /// </summary>
        [Output("parameters")]
        public Output<ImmutableArray<Outputs.WorkflowTemplateParameter>> Parameters { get; private set; } = null!;

        /// <summary>
        /// Required. WorkflowTemplate scheduling information.
        /// </summary>
        [Output("placement")]
        public Output<Outputs.WorkflowTemplatePlacement> Placement { get; private set; } = null!;

        /// <summary>
        /// The project for the resource
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Output only. The time template was last updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// Optional. Used to perform a consistent read-modify-write. This field should be left blank for a `CreateWorkflowTemplate` request. It is required for an `UpdateWorkflowTemplate` request, and must match the current server version. A typical update template flow would fetch the current template with a `GetWorkflowTemplate` request, which will return the current template with the `version` field filled in with the current server version. The user updates other fields in the template, then returns it as part of the `UpdateWorkflowTemplate` request.
        /// </summary>
        [Output("version")]
        public Output<int> Version { get; private set; } = null!;


        /// <summary>
        /// Create a WorkflowTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WorkflowTemplate(string name, WorkflowTemplateArgs args, CustomResourceOptions? options = null)
            : base("gcp:dataproc/workflowTemplate:WorkflowTemplate", name, args ?? new WorkflowTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WorkflowTemplate(string name, Input<string> id, WorkflowTemplateState? state = null, CustomResourceOptions? options = null)
            : base("gcp:dataproc/workflowTemplate:WorkflowTemplate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing WorkflowTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WorkflowTemplate Get(string name, Input<string> id, WorkflowTemplateState? state = null, CustomResourceOptions? options = null)
        {
            return new WorkflowTemplate(name, id, state, options);
        }
    }

    public sealed class WorkflowTemplateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Beta only) Optional. Timeout duration for the DAG of jobs. You can use "s", "m", "h", and "d" suffixes for second, minute, hour, and day duration values, respectively. The timeout duration must be from 10 minutes ("10m") to 24 hours ("24h" or "1d"). The timer begins when the first job is submitted. If the workflow is running at the end of the timeout period, any remaining jobs are cancelled, the workflow is ended, and if the workflow was running on a (/dataproc/docs/concepts/workflows/using-workflows#configuring_or_selecting_a_cluster), the cluster is deleted.
        /// </summary>
        [Input("dagTimeout")]
        public Input<string>? DagTimeout { get; set; }

        [Input("jobs", required: true)]
        private InputList<Inputs.WorkflowTemplateJobArgs>? _jobs;

        /// <summary>
        /// Required. The Directed Acyclic Graph of Jobs to submit.
        /// </summary>
        public InputList<Inputs.WorkflowTemplateJobArgs> Jobs
        {
            get => _jobs ?? (_jobs = new InputList<Inputs.WorkflowTemplateJobArgs>());
            set => _jobs = value;
        }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Optional. The labels to associate with this job. Label keys must be between 1 and 63 characters long, and must conform to the following regular expression: {0,63} No more than 32 labels can be associated with a given job.
        /// (Optional)
        /// Optional. The labels to associate with this template. These labels will be propagated to all jobs and clusters created by the workflow instance. Label **keys** must contain 1 to 63 characters, and must conform to (https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be associated with a template.
        /// (Optional)
        /// Optional. The labels to associate with this cluster. Label keys must be between 1 and 63 characters long, and must conform to the following PCRE regular expression: {0,63} No more than 32 labels can be associated with a given cluster.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The location for the resource
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Output only. The resource name of the workflow template, as described in https://cloud.google.com/apis/design/resource_names. * For `projects.regions.workflowTemplates`, the resource name of the template has the following format: `projects/{project_id}/regions/{region}/workflowTemplates/{template_id}` * For `projects.locations.workflowTemplates`, the resource name of the template has the following format: `projects/{project_id}/locations/{location}/workflowTemplates/{template_id}`
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("parameters")]
        private InputList<Inputs.WorkflowTemplateParameterArgs>? _parameters;

        /// <summary>
        /// Optional. Template parameters whose values are substituted into the template. Values for parameters must be provided when the template is instantiated.
        /// </summary>
        public InputList<Inputs.WorkflowTemplateParameterArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.WorkflowTemplateParameterArgs>());
            set => _parameters = value;
        }

        /// <summary>
        /// Required. WorkflowTemplate scheduling information.
        /// </summary>
        [Input("placement", required: true)]
        public Input<Inputs.WorkflowTemplatePlacementArgs> Placement { get; set; } = null!;

        /// <summary>
        /// The project for the resource
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Optional. Used to perform a consistent read-modify-write. This field should be left blank for a `CreateWorkflowTemplate` request. It is required for an `UpdateWorkflowTemplate` request, and must match the current server version. A typical update template flow would fetch the current template with a `GetWorkflowTemplate` request, which will return the current template with the `version` field filled in with the current server version. The user updates other fields in the template, then returns it as part of the `UpdateWorkflowTemplate` request.
        /// </summary>
        [Input("version")]
        public Input<int>? Version { get; set; }

        public WorkflowTemplateArgs()
        {
        }
        public static new WorkflowTemplateArgs Empty => new WorkflowTemplateArgs();
    }

    public sealed class WorkflowTemplateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Output only. The time template was created.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// (Beta only) Optional. Timeout duration for the DAG of jobs. You can use "s", "m", "h", and "d" suffixes for second, minute, hour, and day duration values, respectively. The timeout duration must be from 10 minutes ("10m") to 24 hours ("24h" or "1d"). The timer begins when the first job is submitted. If the workflow is running at the end of the timeout period, any remaining jobs are cancelled, the workflow is ended, and if the workflow was running on a (/dataproc/docs/concepts/workflows/using-workflows#configuring_or_selecting_a_cluster), the cluster is deleted.
        /// </summary>
        [Input("dagTimeout")]
        public Input<string>? DagTimeout { get; set; }

        [Input("jobs")]
        private InputList<Inputs.WorkflowTemplateJobGetArgs>? _jobs;

        /// <summary>
        /// Required. The Directed Acyclic Graph of Jobs to submit.
        /// </summary>
        public InputList<Inputs.WorkflowTemplateJobGetArgs> Jobs
        {
            get => _jobs ?? (_jobs = new InputList<Inputs.WorkflowTemplateJobGetArgs>());
            set => _jobs = value;
        }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Optional. The labels to associate with this job. Label keys must be between 1 and 63 characters long, and must conform to the following regular expression: {0,63} No more than 32 labels can be associated with a given job.
        /// (Optional)
        /// Optional. The labels to associate with this template. These labels will be propagated to all jobs and clusters created by the workflow instance. Label **keys** must contain 1 to 63 characters, and must conform to (https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be associated with a template.
        /// (Optional)
        /// Optional. The labels to associate with this cluster. Label keys must be between 1 and 63 characters long, and must conform to the following PCRE regular expression: {0,63} No more than 32 labels can be associated with a given cluster.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The location for the resource
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Output only. The resource name of the workflow template, as described in https://cloud.google.com/apis/design/resource_names. * For `projects.regions.workflowTemplates`, the resource name of the template has the following format: `projects/{project_id}/regions/{region}/workflowTemplates/{template_id}` * For `projects.locations.workflowTemplates`, the resource name of the template has the following format: `projects/{project_id}/locations/{location}/workflowTemplates/{template_id}`
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("parameters")]
        private InputList<Inputs.WorkflowTemplateParameterGetArgs>? _parameters;

        /// <summary>
        /// Optional. Template parameters whose values are substituted into the template. Values for parameters must be provided when the template is instantiated.
        /// </summary>
        public InputList<Inputs.WorkflowTemplateParameterGetArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.WorkflowTemplateParameterGetArgs>());
            set => _parameters = value;
        }

        /// <summary>
        /// Required. WorkflowTemplate scheduling information.
        /// </summary>
        [Input("placement")]
        public Input<Inputs.WorkflowTemplatePlacementGetArgs>? Placement { get; set; }

        /// <summary>
        /// The project for the resource
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Output only. The time template was last updated.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        /// <summary>
        /// Optional. Used to perform a consistent read-modify-write. This field should be left blank for a `CreateWorkflowTemplate` request. It is required for an `UpdateWorkflowTemplate` request, and must match the current server version. A typical update template flow would fetch the current template with a `GetWorkflowTemplate` request, which will return the current template with the `version` field filled in with the current server version. The user updates other fields in the template, then returns it as part of the `UpdateWorkflowTemplate` request.
        /// </summary>
        [Input("version")]
        public Input<int>? Version { get; set; }

        public WorkflowTemplateState()
        {
        }
        public static new WorkflowTemplateState Empty => new WorkflowTemplateState();
    }
}
