// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.ActiveDirectory
{
    /// <summary>
    /// ## Example Usage
    /// ### Active Directory Peering Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var source_network = new Gcp.Compute.Network("source-network", new()
    ///     {
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    ///     var ad_domain = new Gcp.ActiveDirectory.Domain("ad-domain", new()
    ///     {
    ///         DomainName = "ad.test.hashicorptest.com",
    ///         Locations = new[]
    ///         {
    ///             "us-central1",
    ///         },
    ///         ReservedIpRange = "192.168.255.0/24",
    ///         AuthorizedNetworks = new[]
    ///         {
    ///             source_network.Id,
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    ///     var peered_project = new Gcp.Organizations.Project("peered-project", new()
    ///     {
    ///         ProjectId = "my-peered-project",
    ///         OrgId = "123456789",
    ///         BillingAccount = "000000-0000000-0000000-000000",
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    ///     var compute = new Gcp.Projects.Service("compute", new()
    ///     {
    ///         Project = peered_project.ProjectId,
    ///         ServiceName = "compute.googleapis.com",
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    ///     var peered_network = new Gcp.Compute.Network("peered-network", new()
    ///     {
    ///         Project = compute.Project,
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    ///     var ad_domain_peering = new Gcp.ActiveDirectory.Peering("ad-domain-peering", new()
    ///     {
    ///         DomainResource = ad_domain.Name,
    ///         PeeringId = "ad-domain-peering",
    ///         AuthorizedNetwork = peered_network.Id,
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
    /// 
    /// ## Import
    /// 
    /// This resource does not support import.
    /// </summary>
    [GcpResourceType("gcp:activedirectory/peering:Peering")]
    public partial class Peering : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The full names of the Google Compute Engine networks to which the instance is connected. Caller needs to make sure that CIDR subnets do not overlap between networks, else peering creation will fail.
        /// </summary>
        [Output("authorizedNetwork")]
        public Output<string> AuthorizedNetwork { get; private set; } = null!;

        /// <summary>
        /// Full domain resource path for the Managed AD Domain involved in peering. The resource path should be in the form projects/{projectId}/locations/global/domains/{domainName}
        /// </summary>
        [Output("domainResource")]
        public Output<string> DomainResource { get; private set; } = null!;

        /// <summary>
        /// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
        /// </summary>
        [Output("effectiveLabels")]
        public Output<ImmutableDictionary<string, string>> EffectiveLabels { get; private set; } = null!;

        /// <summary>
        /// Resource labels that can contain user-provided metadata
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// Unique name of the peering in this scope including projects and location using the form: projects/{projectId}/locations/global/peerings/{peeringId}.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// - - -
        /// </summary>
        [Output("peeringId")]
        public Output<string> PeeringId { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The combination of labels configured directly on the resource
        /// and default labels configured on the provider.
        /// </summary>
        [Output("pulumiLabels")]
        public Output<ImmutableDictionary<string, string>> PulumiLabels { get; private set; } = null!;

        /// <summary>
        /// The current state of this Peering.
        /// </summary>
        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;

        /// <summary>
        /// Additional information about the current status of this peering, if available.
        /// </summary>
        [Output("statusMessage")]
        public Output<string?> StatusMessage { get; private set; } = null!;


        /// <summary>
        /// Create a Peering resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Peering(string name, PeeringArgs args, CustomResourceOptions? options = null)
            : base("gcp:activedirectory/peering:Peering", name, args ?? new PeeringArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Peering(string name, Input<string> id, PeeringState? state = null, CustomResourceOptions? options = null)
            : base("gcp:activedirectory/peering:Peering", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Peering resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Peering Get(string name, Input<string> id, PeeringState? state = null, CustomResourceOptions? options = null)
        {
            return new Peering(name, id, state, options);
        }
    }

    public sealed class PeeringArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The full names of the Google Compute Engine networks to which the instance is connected. Caller needs to make sure that CIDR subnets do not overlap between networks, else peering creation will fail.
        /// </summary>
        [Input("authorizedNetwork", required: true)]
        public Input<string> AuthorizedNetwork { get; set; } = null!;

        /// <summary>
        /// Full domain resource path for the Managed AD Domain involved in peering. The resource path should be in the form projects/{projectId}/locations/global/domains/{domainName}
        /// </summary>
        [Input("domainResource", required: true)]
        public Input<string> DomainResource { get; set; } = null!;

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Resource labels that can contain user-provided metadata
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// - - -
        /// </summary>
        [Input("peeringId", required: true)]
        public Input<string> PeeringId { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The current state of this Peering.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Additional information about the current status of this peering, if available.
        /// </summary>
        [Input("statusMessage")]
        public Input<string>? StatusMessage { get; set; }

        public PeeringArgs()
        {
        }
        public static new PeeringArgs Empty => new PeeringArgs();
    }

    public sealed class PeeringState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The full names of the Google Compute Engine networks to which the instance is connected. Caller needs to make sure that CIDR subnets do not overlap between networks, else peering creation will fail.
        /// </summary>
        [Input("authorizedNetwork")]
        public Input<string>? AuthorizedNetwork { get; set; }

        /// <summary>
        /// Full domain resource path for the Managed AD Domain involved in peering. The resource path should be in the form projects/{projectId}/locations/global/domains/{domainName}
        /// </summary>
        [Input("domainResource")]
        public Input<string>? DomainResource { get; set; }

        [Input("effectiveLabels")]
        private InputMap<string>? _effectiveLabels;

        /// <summary>
        /// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
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

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Resource labels that can contain user-provided metadata
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// Unique name of the peering in this scope including projects and location using the form: projects/{projectId}/locations/global/peerings/{peeringId}.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// - - -
        /// </summary>
        [Input("peeringId")]
        public Input<string>? PeeringId { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("pulumiLabels")]
        private InputMap<string>? _pulumiLabels;

        /// <summary>
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
        /// The current state of this Peering.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Additional information about the current status of this peering, if available.
        /// </summary>
        [Input("statusMessage")]
        public Input<string>? StatusMessage { get; set; }

        public PeeringState()
        {
        }
        public static new PeeringState Empty => new PeeringState();
    }
}
