// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    /// <summary>
    /// Configures the Google Compute Engine
    /// [Default Network Tier](https://cloud.google.com/network-tiers/docs/using-network-service-tiers#setting_the_tier_for_all_resources_in_a_project)
    /// for a project.
    /// 
    /// For more information, see,
    /// [the Project API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/projects/setDefaultNetworkTier).
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
    ///     var @default = new Gcp.Compute.ProjectDefaultNetworkTier("default", new()
    ///     {
    ///         NetworkTier = "PREMIUM",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// This resource can be imported using the project ID
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/projectDefaultNetworkTier:ProjectDefaultNetworkTier default project-id`
    /// ```
    /// </summary>
    [GcpResourceType("gcp:compute/projectDefaultNetworkTier:ProjectDefaultNetworkTier")]
    public partial class ProjectDefaultNetworkTier : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The default network tier to be configured for the project.
        /// This field can take the following values: `PREMIUM` or `STANDARD`.
        /// </summary>
        [Output("networkTier")]
        public Output<string> NetworkTier { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;


        /// <summary>
        /// Create a ProjectDefaultNetworkTier resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ProjectDefaultNetworkTier(string name, ProjectDefaultNetworkTierArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/projectDefaultNetworkTier:ProjectDefaultNetworkTier", name, args ?? new ProjectDefaultNetworkTierArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ProjectDefaultNetworkTier(string name, Input<string> id, ProjectDefaultNetworkTierState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/projectDefaultNetworkTier:ProjectDefaultNetworkTier", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ProjectDefaultNetworkTier resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ProjectDefaultNetworkTier Get(string name, Input<string> id, ProjectDefaultNetworkTierState? state = null, CustomResourceOptions? options = null)
        {
            return new ProjectDefaultNetworkTier(name, id, state, options);
        }
    }

    public sealed class ProjectDefaultNetworkTierArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The default network tier to be configured for the project.
        /// This field can take the following values: `PREMIUM` or `STANDARD`.
        /// </summary>
        [Input("networkTier", required: true)]
        public Input<string> NetworkTier { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public ProjectDefaultNetworkTierArgs()
        {
        }
        public static new ProjectDefaultNetworkTierArgs Empty => new ProjectDefaultNetworkTierArgs();
    }

    public sealed class ProjectDefaultNetworkTierState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The default network tier to be configured for the project.
        /// This field can take the following values: `PREMIUM` or `STANDARD`.
        /// </summary>
        [Input("networkTier")]
        public Input<string>? NetworkTier { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public ProjectDefaultNetworkTierState()
        {
        }
        public static new ProjectDefaultNetworkTierState Empty => new ProjectDefaultNetworkTierState();
    }
}
