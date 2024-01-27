// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.DataPlex
{
    /// <summary>
    /// The Dataplex Zone resource
    /// 
    /// ## Example Usage
    /// ### Basic_zone
    /// A basic example of a dataplex zone
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var basic = new Gcp.DataPlex.Lake("basic", new()
    ///     {
    ///         Location = "us-west1",
    ///         Description = "Lake for DCL",
    ///         DisplayName = "Lake for DCL",
    ///         Project = "my-project-name",
    ///         Labels = 
    ///         {
    ///             { "my-lake", "exists" },
    ///         },
    ///     });
    /// 
    ///     var primary = new Gcp.DataPlex.Zone("primary", new()
    ///     {
    ///         DiscoverySpec = new Gcp.DataPlex.Inputs.ZoneDiscoverySpecArgs
    ///         {
    ///             Enabled = false,
    ///         },
    ///         Lake = basic.Name,
    ///         Location = "us-west1",
    ///         ResourceSpec = new Gcp.DataPlex.Inputs.ZoneResourceSpecArgs
    ///         {
    ///             LocationType = "MULTI_REGION",
    ///         },
    ///         Type = "RAW",
    ///         Description = "Zone for DCL",
    ///         DisplayName = "Zone for DCL",
    ///         Project = "my-project-name",
    ///         Labels = null,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Zone can be imported using any of these accepted formats* `projects/{{project}}/locations/{{location}}/lakes/{{lake}}/zones/{{name}}` * `{{project}}/{{location}}/{{lake}}/{{name}}` * `{{location}}/{{lake}}/{{name}}` When using the `pulumi import` command, Zone can be imported using one of the formats above. For example
    /// 
    /// ```sh
    ///  $ pulumi import gcp:dataplex/zone:Zone default projects/{{project}}/locations/{{location}}/lakes/{{lake}}/zones/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:dataplex/zone:Zone default {{project}}/{{location}}/{{lake}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:dataplex/zone:Zone default {{location}}/{{lake}}/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:dataplex/zone:Zone")]
    public partial class Zone : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Output only. Aggregated status of the underlying assets of the zone.
        /// </summary>
        [Output("assetStatuses")]
        public Output<ImmutableArray<Outputs.ZoneAssetStatus>> AssetStatuses { get; private set; } = null!;

        /// <summary>
        /// Output only. The time when the zone was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Optional. Description of the zone.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Required. Specification of the discovery feature applied to data in this zone.
        /// </summary>
        [Output("discoverySpec")]
        public Output<Outputs.ZoneDiscoverySpec> DiscoverySpec { get; private set; } = null!;

        /// <summary>
        /// Optional. User friendly display name.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
        /// </summary>
        [Output("effectiveLabels")]
        public Output<ImmutableDictionary<string, object>> EffectiveLabels { get; private set; } = null!;

        /// <summary>
        /// Optional. User defined labels for the zone.
        /// 
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// The lake for the resource
        /// </summary>
        [Output("lake")]
        public Output<string> Lake { get; private set; } = null!;

        /// <summary>
        /// The location for the resource
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the zone.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The project for the resource
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The combination of labels configured directly on the resource and default labels configured on the provider.
        /// </summary>
        [Output("pulumiLabels")]
        public Output<ImmutableDictionary<string, object>> PulumiLabels { get; private set; } = null!;

        /// <summary>
        /// Required. Immutable. Specification of the resources that are referenced by the assets within this zone.
        /// </summary>
        [Output("resourceSpec")]
        public Output<Outputs.ZoneResourceSpec> ResourceSpec { get; private set; } = null!;

        /// <summary>
        /// Output only. Current state of the zone. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, DELETING, ACTION_REQUIRED
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Required. Immutable. The type of the zone. Possible values: TYPE_UNSPECIFIED, RAW, CURATED
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Output only. System generated globally unique ID for the zone. This ID will be different if the zone is deleted and re-created with the same name.
        /// </summary>
        [Output("uid")]
        public Output<string> Uid { get; private set; } = null!;

        /// <summary>
        /// Output only. The time when the zone was last updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a Zone resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Zone(string name, ZoneArgs args, CustomResourceOptions? options = null)
            : base("gcp:dataplex/zone:Zone", name, args ?? new ZoneArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Zone(string name, Input<string> id, ZoneState? state = null, CustomResourceOptions? options = null)
            : base("gcp:dataplex/zone:Zone", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Zone resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Zone Get(string name, Input<string> id, ZoneState? state = null, CustomResourceOptions? options = null)
        {
            return new Zone(name, id, state, options);
        }
    }

    public sealed class ZoneArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Description of the zone.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Required. Specification of the discovery feature applied to data in this zone.
        /// </summary>
        [Input("discoverySpec", required: true)]
        public Input<Inputs.ZoneDiscoverySpecArgs> DiscoverySpec { get; set; } = null!;

        /// <summary>
        /// Optional. User friendly display name.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Optional. User defined labels for the zone.
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
        /// The lake for the resource
        /// </summary>
        [Input("lake", required: true)]
        public Input<string> Lake { get; set; } = null!;

        /// <summary>
        /// The location for the resource
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The name of the zone.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The project for the resource
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Required. Immutable. Specification of the resources that are referenced by the assets within this zone.
        /// </summary>
        [Input("resourceSpec", required: true)]
        public Input<Inputs.ZoneResourceSpecArgs> ResourceSpec { get; set; } = null!;

        /// <summary>
        /// Required. Immutable. The type of the zone. Possible values: TYPE_UNSPECIFIED, RAW, CURATED
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ZoneArgs()
        {
        }
        public static new ZoneArgs Empty => new ZoneArgs();
    }

    public sealed class ZoneState : global::Pulumi.ResourceArgs
    {
        [Input("assetStatuses")]
        private InputList<Inputs.ZoneAssetStatusGetArgs>? _assetStatuses;

        /// <summary>
        /// Output only. Aggregated status of the underlying assets of the zone.
        /// </summary>
        public InputList<Inputs.ZoneAssetStatusGetArgs> AssetStatuses
        {
            get => _assetStatuses ?? (_assetStatuses = new InputList<Inputs.ZoneAssetStatusGetArgs>());
            set => _assetStatuses = value;
        }

        /// <summary>
        /// Output only. The time when the zone was created.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Optional. Description of the zone.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Required. Specification of the discovery feature applied to data in this zone.
        /// </summary>
        [Input("discoverySpec")]
        public Input<Inputs.ZoneDiscoverySpecGetArgs>? DiscoverySpec { get; set; }

        /// <summary>
        /// Optional. User friendly display name.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("effectiveLabels")]
        private InputMap<object>? _effectiveLabels;

        /// <summary>
        /// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
        /// </summary>
        public InputMap<object> EffectiveLabels
        {
            get => _effectiveLabels ?? (_effectiveLabels = new InputMap<object>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, object>());
                _effectiveLabels = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Optional. User defined labels for the zone.
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
        /// The lake for the resource
        /// </summary>
        [Input("lake")]
        public Input<string>? Lake { get; set; }

        /// <summary>
        /// The location for the resource
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the zone.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The project for the resource
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("pulumiLabels")]
        private InputMap<object>? _pulumiLabels;

        /// <summary>
        /// The combination of labels configured directly on the resource and default labels configured on the provider.
        /// </summary>
        public InputMap<object> PulumiLabels
        {
            get => _pulumiLabels ?? (_pulumiLabels = new InputMap<object>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, object>());
                _pulumiLabels = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        /// <summary>
        /// Required. Immutable. Specification of the resources that are referenced by the assets within this zone.
        /// </summary>
        [Input("resourceSpec")]
        public Input<Inputs.ZoneResourceSpecGetArgs>? ResourceSpec { get; set; }

        /// <summary>
        /// Output only. Current state of the zone. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, DELETING, ACTION_REQUIRED
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Required. Immutable. The type of the zone. Possible values: TYPE_UNSPECIFIED, RAW, CURATED
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Output only. System generated globally unique ID for the zone. This ID will be different if the zone is deleted and re-created with the same name.
        /// </summary>
        [Input("uid")]
        public Input<string>? Uid { get; set; }

        /// <summary>
        /// Output only. The time when the zone was last updated.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public ZoneState()
        {
        }
        public static new ZoneState Empty => new ZoneState();
    }
}
