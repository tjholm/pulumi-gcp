// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.GkeHub
{
    /// <summary>
    /// ## Example Usage
    /// ### Config Management
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var cluster = new Gcp.Container.Cluster("cluster", new()
    ///     {
    ///         Location = "us-central1-a",
    ///         InitialNodeCount = 1,
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    ///     var membership = new Gcp.GkeHub.Membership("membership", new()
    ///     {
    ///         MembershipId = "my-membership",
    ///         Endpoint = new Gcp.GkeHub.Inputs.MembershipEndpointArgs
    ///         {
    ///             GkeCluster = new Gcp.GkeHub.Inputs.MembershipEndpointGkeClusterArgs
    ///             {
    ///                 ResourceLink = cluster.Id.Apply(id =&gt; $"//container.googleapis.com/{id}"),
    ///             },
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    ///     var feature = new Gcp.GkeHub.Feature("feature", new()
    ///     {
    ///         Location = "global",
    ///         Labels = 
    ///         {
    ///             { "foo", "bar" },
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    ///     var featureMember = new Gcp.GkeHub.FeatureMembership("featureMember", new()
    ///     {
    ///         Location = "global",
    ///         Feature = feature.Name,
    ///         Membership = membership.MembershipId,
    ///         Configmanagement = new Gcp.GkeHub.Inputs.FeatureMembershipConfigmanagementArgs
    ///         {
    ///             Version = "1.6.2",
    ///             ConfigSync = new Gcp.GkeHub.Inputs.FeatureMembershipConfigmanagementConfigSyncArgs
    ///             {
    ///                 Git = new Gcp.GkeHub.Inputs.FeatureMembershipConfigmanagementConfigSyncGitArgs
    ///                 {
    ///                     SyncRepo = "https://github.com/hashicorp/terraform",
    ///                 },
    ///             },
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    /// });
    /// ```
    /// ### Config Management With OCI
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var cluster = new Gcp.Container.Cluster("cluster", new()
    ///     {
    ///         Location = "us-central1-a",
    ///         InitialNodeCount = 1,
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    ///     var membership = new Gcp.GkeHub.Membership("membership", new()
    ///     {
    ///         MembershipId = "my-membership",
    ///         Endpoint = new Gcp.GkeHub.Inputs.MembershipEndpointArgs
    ///         {
    ///             GkeCluster = new Gcp.GkeHub.Inputs.MembershipEndpointGkeClusterArgs
    ///             {
    ///                 ResourceLink = cluster.Id.Apply(id =&gt; $"//container.googleapis.com/{id}"),
    ///             },
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    ///     var feature = new Gcp.GkeHub.Feature("feature", new()
    ///     {
    ///         Location = "global",
    ///         Labels = 
    ///         {
    ///             { "foo", "bar" },
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    ///     var featureMember = new Gcp.GkeHub.FeatureMembership("featureMember", new()
    ///     {
    ///         Location = "global",
    ///         Feature = feature.Name,
    ///         Membership = membership.MembershipId,
    ///         Configmanagement = new Gcp.GkeHub.Inputs.FeatureMembershipConfigmanagementArgs
    ///         {
    ///             Version = "1.12.0",
    ///             ConfigSync = new Gcp.GkeHub.Inputs.FeatureMembershipConfigmanagementConfigSyncArgs
    ///             {
    ///                 Oci = new Gcp.GkeHub.Inputs.FeatureMembershipConfigmanagementConfigSyncOciArgs
    ///                 {
    ///                     SyncRepo = "us-central1-docker.pkg.dev/sample-project/config-repo/config-sync-gke:latest",
    ///                     PolicyDir = "config-connector",
    ///                     SyncWaitSecs = "20",
    ///                     SecretType = "gcpserviceaccount",
    ///                     GcpServiceAccountEmail = "sa@project-id.iam.gserviceaccount.com",
    ///                 },
    ///             },
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    /// });
    /// ```
    /// ### Multi Cluster Service Discovery
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var feature = new Gcp.GkeHub.Feature("feature", new()
    ///     {
    ///         Location = "global",
    ///         Labels = 
    ///         {
    ///             { "foo", "bar" },
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    /// });
    /// ```
    /// ### Service Mesh
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var cluster = new Gcp.Container.Cluster("cluster", new()
    ///     {
    ///         Location = "us-central1-a",
    ///         InitialNodeCount = 1,
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    ///     var membership = new Gcp.GkeHub.Membership("membership", new()
    ///     {
    ///         MembershipId = "my-membership",
    ///         Endpoint = new Gcp.GkeHub.Inputs.MembershipEndpointArgs
    ///         {
    ///             GkeCluster = new Gcp.GkeHub.Inputs.MembershipEndpointGkeClusterArgs
    ///             {
    ///                 ResourceLink = cluster.Id.Apply(id =&gt; $"//container.googleapis.com/{id}"),
    ///             },
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    ///     var feature = new Gcp.GkeHub.Feature("feature", new()
    ///     {
    ///         Location = "global",
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    ///     var featureMember = new Gcp.GkeHub.FeatureMembership("featureMember", new()
    ///     {
    ///         Location = "global",
    ///         Feature = feature.Name,
    ///         Membership = membership.MembershipId,
    ///         Mesh = new Gcp.GkeHub.Inputs.FeatureMembershipMeshArgs
    ///         {
    ///             Management = "MANAGEMENT_AUTOMATIC",
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// FeatureMembership can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:gkehub/featureMembership:FeatureMembership default projects/{{project}}/locations/{{location}}/features/{{feature}}/membershipId/{{membership}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:gkehub/featureMembership:FeatureMembership default {{project}}/{{location}}/{{feature}}/{{membership}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:gkehub/featureMembership:FeatureMembership default {{location}}/{{feature}}/{{membership}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:gkehub/featureMembership:FeatureMembership")]
    public partial class FeatureMembership : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Config Management-specific spec. Structure is documented below.
        /// </summary>
        [Output("configmanagement")]
        public Output<Outputs.FeatureMembershipConfigmanagement?> Configmanagement { get; private set; } = null!;

        /// <summary>
        /// The name of the feature
        /// </summary>
        [Output("feature")]
        public Output<string> Feature { get; private set; } = null!;

        /// <summary>
        /// The location of the feature
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the membership
        /// </summary>
        [Output("membership")]
        public Output<string> Membership { get; private set; } = null!;

        /// <summary>
        /// Service mesh specific spec. Structure is documented below.
        /// </summary>
        [Output("mesh")]
        public Output<Outputs.FeatureMembershipMesh?> Mesh { get; private set; } = null!;

        /// <summary>
        /// The project of the feature
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;


        /// <summary>
        /// Create a FeatureMembership resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FeatureMembership(string name, FeatureMembershipArgs args, CustomResourceOptions? options = null)
            : base("gcp:gkehub/featureMembership:FeatureMembership", name, args ?? new FeatureMembershipArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FeatureMembership(string name, Input<string> id, FeatureMembershipState? state = null, CustomResourceOptions? options = null)
            : base("gcp:gkehub/featureMembership:FeatureMembership", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FeatureMembership resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FeatureMembership Get(string name, Input<string> id, FeatureMembershipState? state = null, CustomResourceOptions? options = null)
        {
            return new FeatureMembership(name, id, state, options);
        }
    }

    public sealed class FeatureMembershipArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Config Management-specific spec. Structure is documented below.
        /// </summary>
        [Input("configmanagement")]
        public Input<Inputs.FeatureMembershipConfigmanagementArgs>? Configmanagement { get; set; }

        /// <summary>
        /// The name of the feature
        /// </summary>
        [Input("feature", required: true)]
        public Input<string> Feature { get; set; } = null!;

        /// <summary>
        /// The location of the feature
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The name of the membership
        /// </summary>
        [Input("membership", required: true)]
        public Input<string> Membership { get; set; } = null!;

        /// <summary>
        /// Service mesh specific spec. Structure is documented below.
        /// </summary>
        [Input("mesh")]
        public Input<Inputs.FeatureMembershipMeshArgs>? Mesh { get; set; }

        /// <summary>
        /// The project of the feature
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public FeatureMembershipArgs()
        {
        }
        public static new FeatureMembershipArgs Empty => new FeatureMembershipArgs();
    }

    public sealed class FeatureMembershipState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Config Management-specific spec. Structure is documented below.
        /// </summary>
        [Input("configmanagement")]
        public Input<Inputs.FeatureMembershipConfigmanagementGetArgs>? Configmanagement { get; set; }

        /// <summary>
        /// The name of the feature
        /// </summary>
        [Input("feature")]
        public Input<string>? Feature { get; set; }

        /// <summary>
        /// The location of the feature
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the membership
        /// </summary>
        [Input("membership")]
        public Input<string>? Membership { get; set; }

        /// <summary>
        /// Service mesh specific spec. Structure is documented below.
        /// </summary>
        [Input("mesh")]
        public Input<Inputs.FeatureMembershipMeshGetArgs>? Mesh { get; set; }

        /// <summary>
        /// The project of the feature
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public FeatureMembershipState()
        {
        }
        public static new FeatureMembershipState Empty => new FeatureMembershipState();
    }
}
