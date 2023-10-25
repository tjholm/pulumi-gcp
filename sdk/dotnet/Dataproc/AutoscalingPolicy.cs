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
    /// Describes an autoscaling policy for Dataproc cluster autoscaler.
    /// 
    /// ## Example Usage
    /// ### Dataproc Autoscaling Policy
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var asp = new Gcp.Dataproc.AutoscalingPolicy("asp", new()
    ///     {
    ///         PolicyId = "dataproc-policy",
    ///         Location = "us-central1",
    ///         WorkerConfig = new Gcp.Dataproc.Inputs.AutoscalingPolicyWorkerConfigArgs
    ///         {
    ///             MaxInstances = 3,
    ///         },
    ///         BasicAlgorithm = new Gcp.Dataproc.Inputs.AutoscalingPolicyBasicAlgorithmArgs
    ///         {
    ///             YarnConfig = new Gcp.Dataproc.Inputs.AutoscalingPolicyBasicAlgorithmYarnConfigArgs
    ///             {
    ///                 GracefulDecommissionTimeout = "30s",
    ///                 ScaleUpFactor = 0.5,
    ///                 ScaleDownFactor = 0.5,
    ///             },
    ///         },
    ///     });
    /// 
    ///     var basic = new Gcp.Dataproc.Cluster("basic", new()
    ///     {
    ///         Region = "us-central1",
    ///         ClusterConfig = new Gcp.Dataproc.Inputs.ClusterClusterConfigArgs
    ///         {
    ///             AutoscalingConfig = new Gcp.Dataproc.Inputs.ClusterClusterConfigAutoscalingConfigArgs
    ///             {
    ///                 PolicyUri = asp.Name,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// AutoscalingPolicy can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:dataproc/autoscalingPolicy:AutoscalingPolicy default projects/{{project}}/locations/{{location}}/autoscalingPolicies/{{policy_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:dataproc/autoscalingPolicy:AutoscalingPolicy default {{project}}/{{location}}/{{policy_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:dataproc/autoscalingPolicy:AutoscalingPolicy default {{location}}/{{policy_id}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:dataproc/autoscalingPolicy:AutoscalingPolicy")]
    public partial class AutoscalingPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Basic algorithm for autoscaling.
        /// Structure is documented below.
        /// </summary>
        [Output("basicAlgorithm")]
        public Output<Outputs.AutoscalingPolicyBasicAlgorithm?> BasicAlgorithm { get; private set; } = null!;

        /// <summary>
        /// The  location where the autoscaling policy should reside.
        /// The default value is `global`.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The "resource name" of the autoscaling policy.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The policy id. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
        /// and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between
        /// 3 and 50 characters.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Output("policyId")]
        public Output<string> PolicyId { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Describes how the autoscaler will operate for secondary workers.
        /// Structure is documented below.
        /// </summary>
        [Output("secondaryWorkerConfig")]
        public Output<Outputs.AutoscalingPolicySecondaryWorkerConfig?> SecondaryWorkerConfig { get; private set; } = null!;

        /// <summary>
        /// Describes how the autoscaler will operate for primary workers.
        /// Structure is documented below.
        /// </summary>
        [Output("workerConfig")]
        public Output<Outputs.AutoscalingPolicyWorkerConfig?> WorkerConfig { get; private set; } = null!;


        /// <summary>
        /// Create a AutoscalingPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AutoscalingPolicy(string name, AutoscalingPolicyArgs args, CustomResourceOptions? options = null)
            : base("gcp:dataproc/autoscalingPolicy:AutoscalingPolicy", name, args ?? new AutoscalingPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AutoscalingPolicy(string name, Input<string> id, AutoscalingPolicyState? state = null, CustomResourceOptions? options = null)
            : base("gcp:dataproc/autoscalingPolicy:AutoscalingPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AutoscalingPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AutoscalingPolicy Get(string name, Input<string> id, AutoscalingPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new AutoscalingPolicy(name, id, state, options);
        }
    }

    public sealed class AutoscalingPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Basic algorithm for autoscaling.
        /// Structure is documented below.
        /// </summary>
        [Input("basicAlgorithm")]
        public Input<Inputs.AutoscalingPolicyBasicAlgorithmArgs>? BasicAlgorithm { get; set; }

        /// <summary>
        /// The  location where the autoscaling policy should reside.
        /// The default value is `global`.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The policy id. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
        /// and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between
        /// 3 and 50 characters.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("policyId", required: true)]
        public Input<string> PolicyId { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Describes how the autoscaler will operate for secondary workers.
        /// Structure is documented below.
        /// </summary>
        [Input("secondaryWorkerConfig")]
        public Input<Inputs.AutoscalingPolicySecondaryWorkerConfigArgs>? SecondaryWorkerConfig { get; set; }

        /// <summary>
        /// Describes how the autoscaler will operate for primary workers.
        /// Structure is documented below.
        /// </summary>
        [Input("workerConfig")]
        public Input<Inputs.AutoscalingPolicyWorkerConfigArgs>? WorkerConfig { get; set; }

        public AutoscalingPolicyArgs()
        {
        }
        public static new AutoscalingPolicyArgs Empty => new AutoscalingPolicyArgs();
    }

    public sealed class AutoscalingPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Basic algorithm for autoscaling.
        /// Structure is documented below.
        /// </summary>
        [Input("basicAlgorithm")]
        public Input<Inputs.AutoscalingPolicyBasicAlgorithmGetArgs>? BasicAlgorithm { get; set; }

        /// <summary>
        /// The  location where the autoscaling policy should reside.
        /// The default value is `global`.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The "resource name" of the autoscaling policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The policy id. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
        /// and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between
        /// 3 and 50 characters.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("policyId")]
        public Input<string>? PolicyId { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Describes how the autoscaler will operate for secondary workers.
        /// Structure is documented below.
        /// </summary>
        [Input("secondaryWorkerConfig")]
        public Input<Inputs.AutoscalingPolicySecondaryWorkerConfigGetArgs>? SecondaryWorkerConfig { get; set; }

        /// <summary>
        /// Describes how the autoscaler will operate for primary workers.
        /// Structure is documented below.
        /// </summary>
        [Input("workerConfig")]
        public Input<Inputs.AutoscalingPolicyWorkerConfigGetArgs>? WorkerConfig { get; set; }

        public AutoscalingPolicyState()
        {
        }
        public static new AutoscalingPolicyState Empty => new AutoscalingPolicyState();
    }
}
