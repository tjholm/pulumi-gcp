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
    /// Feature represents the settings and status of any Hub Feature.
    /// 
    /// To get more information about Feature, see:
    /// 
    /// * [API documentation](https://cloud.google.com/anthos/fleet-management/docs/reference/rest/v1/projects.locations.features)
    /// * How-to Guides
    ///     * [Registering a Cluster](https://cloud.google.com/anthos/multicluster-management/connect/registering-a-cluster#register_cluster)
    /// 
    /// ## Example Usage
    /// ### Gkehub Feature Multi Cluster Ingress
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
    ///         Description = "Membership",
    ///     });
    /// 
    ///     var feature = new Gcp.GkeHub.Feature("feature", new()
    ///     {
    ///         Location = "global",
    ///         Spec = new Gcp.GkeHub.Inputs.FeatureSpecArgs
    ///         {
    ///             Multiclusteringress = new Gcp.GkeHub.Inputs.FeatureSpecMulticlusteringressArgs
    ///             {
    ///                 ConfigMembership = membership.Id,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Gkehub Feature Multi Cluster Service Discovery
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
    ///         Labels = 
    ///         {
    ///             { "foo", "bar" },
    ///         },
    ///         Location = "global",
    ///     });
    /// 
    /// });
    /// ```
    /// ### Gkehub Feature Anthos Service Mesh
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
    ///     });
    /// 
    /// });
    /// ```
    /// ### Enable Fleet Observability For Default Logs With Copy
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
    ///         Spec = new Gcp.GkeHub.Inputs.FeatureSpecArgs
    ///         {
    ///             Fleetobservability = new Gcp.GkeHub.Inputs.FeatureSpecFleetobservabilityArgs
    ///             {
    ///                 LoggingConfig = new Gcp.GkeHub.Inputs.FeatureSpecFleetobservabilityLoggingConfigArgs
    ///                 {
    ///                     DefaultConfig = new Gcp.GkeHub.Inputs.FeatureSpecFleetobservabilityLoggingConfigDefaultConfigArgs
    ///                     {
    ///                         Mode = "COPY",
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Enable Fleet Observability For Scope Logs With Move
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
    ///         Spec = new Gcp.GkeHub.Inputs.FeatureSpecArgs
    ///         {
    ///             Fleetobservability = new Gcp.GkeHub.Inputs.FeatureSpecFleetobservabilityArgs
    ///             {
    ///                 LoggingConfig = new Gcp.GkeHub.Inputs.FeatureSpecFleetobservabilityLoggingConfigArgs
    ///                 {
    ///                     FleetScopeLogsConfig = new Gcp.GkeHub.Inputs.FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigArgs
    ///                     {
    ///                         Mode = "MOVE",
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Enable Fleet Observability For Both Default And Scope Logs
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
    ///         Spec = new Gcp.GkeHub.Inputs.FeatureSpecArgs
    ///         {
    ///             Fleetobservability = new Gcp.GkeHub.Inputs.FeatureSpecFleetobservabilityArgs
    ///             {
    ///                 LoggingConfig = new Gcp.GkeHub.Inputs.FeatureSpecFleetobservabilityLoggingConfigArgs
    ///                 {
    ///                     DefaultConfig = new Gcp.GkeHub.Inputs.FeatureSpecFleetobservabilityLoggingConfigDefaultConfigArgs
    ///                     {
    ///                         Mode = "COPY",
    ///                     },
    ///                     FleetScopeLogsConfig = new Gcp.GkeHub.Inputs.FeatureSpecFleetobservabilityLoggingConfigFleetScopeLogsConfigArgs
    ///                     {
    ///                         Mode = "MOVE",
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
    /// Feature can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:gkehub/feature:Feature default projects/{{project}}/locations/{{location}}/features/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:gkehub/feature:Feature default {{project}}/{{location}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:gkehub/feature:Feature default {{location}}/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:gkehub/feature:Feature")]
    public partial class Feature : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Output only. When the Feature resource was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Output only. When the Feature resource was deleted.
        /// </summary>
        [Output("deleteTime")]
        public Output<string> DeleteTime { get; private set; } = null!;

        /// <summary>
        /// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
        /// clients and services.
        /// </summary>
        [Output("effectiveLabels")]
        public Output<ImmutableDictionary<string, string>> EffectiveLabels { get; private set; } = null!;

        /// <summary>
        /// GCP labels for this Feature.
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// The location for the resource
        /// 
        /// 
        /// - - -
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The full, unique name of this Feature resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

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
        /// State of the Feature resource itself.
        /// Structure is documented below.
        /// </summary>
        [Output("resourceStates")]
        public Output<ImmutableArray<Outputs.FeatureResourceState>> ResourceStates { get; private set; } = null!;

        /// <summary>
        /// Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.
        /// Structure is documented below.
        /// </summary>
        [Output("spec")]
        public Output<Outputs.FeatureSpec?> Spec { get; private set; } = null!;

        /// <summary>
        /// (Output)
        /// Output only. The "running state" of the Feature in this Hub.
        /// Structure is documented below.
        /// </summary>
        [Output("states")]
        public Output<ImmutableArray<Outputs.FeatureState>> States { get; private set; } = null!;

        /// <summary>
        /// (Output)
        /// The time this status and any related Feature-specific details were updated. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z"
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a Feature resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Feature(string name, FeatureArgs args, CustomResourceOptions? options = null)
            : base("gcp:gkehub/feature:Feature", name, args ?? new FeatureArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Feature(string name, Input<string> id, FeatureState? state = null, CustomResourceOptions? options = null)
            : base("gcp:gkehub/feature:Feature", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Feature resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Feature Get(string name, Input<string> id, FeatureState? state = null, CustomResourceOptions? options = null)
        {
            return new Feature(name, id, state, options);
        }
    }

    public sealed class FeatureArgs : global::Pulumi.ResourceArgs
    {
        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// GCP labels for this Feature.
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The location for the resource
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The full, unique name of this Feature resource
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.
        /// Structure is documented below.
        /// </summary>
        [Input("spec")]
        public Input<Inputs.FeatureSpecArgs>? Spec { get; set; }

        public FeatureArgs()
        {
        }
        public static new FeatureArgs Empty => new FeatureArgs();
    }

    public sealed class FeatureState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Output only. When the Feature resource was created.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Output only. When the Feature resource was deleted.
        /// </summary>
        [Input("deleteTime")]
        public Input<string>? DeleteTime { get; set; }

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

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// GCP labels for this Feature.
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The location for the resource
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The full, unique name of this Feature resource
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

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

        [Input("resourceStates")]
        private InputList<Inputs.FeatureResourceStateGetArgs>? _resourceStates;

        /// <summary>
        /// State of the Feature resource itself.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.FeatureResourceStateGetArgs> ResourceStates
        {
            get => _resourceStates ?? (_resourceStates = new InputList<Inputs.FeatureResourceStateGetArgs>());
            set => _resourceStates = value;
        }

        /// <summary>
        /// Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.
        /// Structure is documented below.
        /// </summary>
        [Input("spec")]
        public Input<Inputs.FeatureSpecGetArgs>? Spec { get; set; }

        [Input("states")]
        private InputList<Inputs.FeatureStateGetArgs>? _states;

        /// <summary>
        /// (Output)
        /// Output only. The "running state" of the Feature in this Hub.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.FeatureStateGetArgs> States
        {
            get => _states ?? (_states = new InputList<Inputs.FeatureStateGetArgs>());
            set => _states = value;
        }

        /// <summary>
        /// (Output)
        /// The time this status and any related Feature-specific details were updated. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z"
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public FeatureState()
        {
        }
        public static new FeatureState Empty => new FeatureState();
    }
}
