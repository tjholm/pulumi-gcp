// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Logging
{
    /// <summary>
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var my_exclusion = new Gcp.Logging.ProjectExclusion("my-exclusion", new()
    ///     {
    ///         Description = "Exclude GCE instance debug logs",
    ///         Filter = "resource.type = gce_instance AND severity &lt;= DEBUG",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Project-level logging exclusions can be imported using their URI, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:logging/projectExclusion:ProjectExclusion my_exclusion projects/my-project/exclusions/my-exclusion
    /// ```
    /// </summary>
    [GcpResourceType("gcp:logging/projectExclusion:ProjectExclusion")]
    public partial class ProjectExclusion : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A human-readable description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Whether this exclusion rule should be disabled or not. This defaults to
        /// false.
        /// </summary>
        [Output("disabled")]
        public Output<bool?> Disabled { get; private set; } = null!;

        /// <summary>
        /// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
        /// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
        /// write a filter.
        /// </summary>
        [Output("filter")]
        public Output<string> Filter { get; private set; } = null!;

        /// <summary>
        /// The name of the logging exclusion.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The project to create the exclusion in. If omitted, the project associated with the provider is
        /// used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;


        /// <summary>
        /// Create a ProjectExclusion resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProjectExclusion(string name, ProjectExclusionArgs args, CustomResourceOptions? options = null)
            : base("gcp:logging/projectExclusion:ProjectExclusion", name, args ?? new ProjectExclusionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProjectExclusion(string name, Input<string> id, ProjectExclusionState? state = null, CustomResourceOptions? options = null)
            : base("gcp:logging/projectExclusion:ProjectExclusion", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProjectExclusion resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProjectExclusion Get(string name, Input<string> id, ProjectExclusionState? state = null, CustomResourceOptions? options = null)
        {
            return new ProjectExclusion(name, id, state, options);
        }
    }

    public sealed class ProjectExclusionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A human-readable description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether this exclusion rule should be disabled or not. This defaults to
        /// false.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
        /// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
        /// write a filter.
        /// </summary>
        [Input("filter", required: true)]
        public Input<string> Filter { get; set; } = null!;

        /// <summary>
        /// The name of the logging exclusion.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The project to create the exclusion in. If omitted, the project associated with the provider is
        /// used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public ProjectExclusionArgs()
        {
        }
        public static new ProjectExclusionArgs Empty => new ProjectExclusionArgs();
    }

    public sealed class ProjectExclusionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A human-readable description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether this exclusion rule should be disabled or not. This defaults to
        /// false.
        /// </summary>
        [Input("disabled")]
        public Input<bool>? Disabled { get; set; }

        /// <summary>
        /// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
        /// See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
        /// write a filter.
        /// </summary>
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// The name of the logging exclusion.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The project to create the exclusion in. If omitted, the project associated with the provider is
        /// used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public ProjectExclusionState()
        {
        }
        public static new ProjectExclusionState Empty => new ProjectExclusionState();
    }
}
