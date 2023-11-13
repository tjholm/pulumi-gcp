// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Vertex
{
    /// <summary>
    /// Feature Metadata information that describes an attribute of an entity type. For example, apple is an entity type, and color is a feature that describes apple.
    /// 
    /// To get more information about FeaturestoreEntitytypeFeature, see:
    /// 
    /// * [API documentation](https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featurestores.entityTypes.features)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/vertex-ai/docs)
    /// 
    /// ## Example Usage
    /// ### Vertex Ai Featurestore Entitytype Feature
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var featurestore = new Gcp.Vertex.AiFeatureStore("featurestore", new()
    ///     {
    ///         Labels = 
    ///         {
    ///             { "foo", "bar" },
    ///         },
    ///         Region = "us-central1",
    ///         OnlineServingConfig = new Gcp.Vertex.Inputs.AiFeatureStoreOnlineServingConfigArgs
    ///         {
    ///             FixedNodeCount = 2,
    ///         },
    ///     });
    /// 
    ///     var entity = new Gcp.Vertex.AiFeatureStoreEntityType("entity", new()
    ///     {
    ///         Labels = 
    ///         {
    ///             { "foo", "bar" },
    ///         },
    ///         Featurestore = featurestore.Id,
    ///     });
    /// 
    ///     var feature = new Gcp.Vertex.AiFeatureStoreEntityTypeFeature("feature", new()
    ///     {
    ///         Labels = 
    ///         {
    ///             { "foo", "bar" },
    ///         },
    ///         Entitytype = entity.Id,
    ///         ValueType = "INT64_ARRAY",
    ///     });
    /// 
    /// });
    /// ```
    /// ### Vertex Ai Featurestore Entitytype Feature With Beta Fields
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var featurestore = new Gcp.Vertex.AiFeatureStore("featurestore", new()
    ///     {
    ///         Labels = 
    ///         {
    ///             { "foo", "bar" },
    ///         },
    ///         Region = "us-central1",
    ///         OnlineServingConfig = new Gcp.Vertex.Inputs.AiFeatureStoreOnlineServingConfigArgs
    ///         {
    ///             FixedNodeCount = 2,
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    ///     var entity = new Gcp.Vertex.AiFeatureStoreEntityType("entity", new()
    ///     {
    ///         Labels = 
    ///         {
    ///             { "foo", "bar" },
    ///         },
    ///         Featurestore = featurestore.Id,
    ///         MonitoringConfig = new Gcp.Vertex.Inputs.AiFeatureStoreEntityTypeMonitoringConfigArgs
    ///         {
    ///             SnapshotAnalysis = new Gcp.Vertex.Inputs.AiFeatureStoreEntityTypeMonitoringConfigSnapshotAnalysisArgs
    ///             {
    ///                 Disabled = false,
    ///                 MonitoringInterval = "86400s",
    ///             },
    ///             CategoricalThresholdConfig = new Gcp.Vertex.Inputs.AiFeatureStoreEntityTypeMonitoringConfigCategoricalThresholdConfigArgs
    ///             {
    ///                 Value = 0.3,
    ///             },
    ///             NumericalThresholdConfig = new Gcp.Vertex.Inputs.AiFeatureStoreEntityTypeMonitoringConfigNumericalThresholdConfigArgs
    ///             {
    ///                 Value = 0.3,
    ///             },
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    ///     var feature = new Gcp.Vertex.AiFeatureStoreEntityTypeFeature("feature", new()
    ///     {
    ///         Labels = 
    ///         {
    ///             { "foo", "bar" },
    ///         },
    ///         Entitytype = entity.Id,
    ///         ValueType = "INT64_ARRAY",
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
    /// FeaturestoreEntitytypeFeature can be imported using any of these accepted formats:
    /// 
    /// ```sh
    ///  $ pulumi import gcp:vertex/aiFeatureStoreEntityTypeFeature:AiFeatureStoreEntityTypeFeature default {{entitytype}}/features/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:vertex/aiFeatureStoreEntityTypeFeature:AiFeatureStoreEntityTypeFeature")]
    public partial class AiFeatureStoreEntityTypeFeature : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The timestamp of when the entity type was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Description of the feature.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
        /// </summary>
        [Output("effectiveLabels")]
        public Output<ImmutableDictionary<string, string>> EffectiveLabels { get; private set; } = null!;

        /// <summary>
        /// The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}/entityTypes/{entitytype}.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Output("entitytype")]
        public Output<string> Entitytype { get; private set; } = null!;

        /// <summary>
        /// Used to perform consistent read-modify-write updates.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// A set of key/value label pairs to assign to the feature.
        /// 
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// The name of the feature. The feature can be up to 64 characters long and can consist only of ASCII Latin letters A-Z and a-z, underscore(_), and ASCII digits 0-9 starting with a letter. The value will be unique given an entity type.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The combination of labels configured directly on the resource
        /// and default labels configured on the provider.
        /// </summary>
        [Output("pulumiLabels")]
        public Output<ImmutableDictionary<string, string>> PulumiLabels { get; private set; } = null!;

        /// <summary>
        /// The region of the feature
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The timestamp when the entity type was most recently updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// Type of Feature value. Immutable. https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featurestores.entityTypes.features#ValueType
        /// </summary>
        [Output("valueType")]
        public Output<string> ValueType { get; private set; } = null!;


        /// <summary>
        /// Create a AiFeatureStoreEntityTypeFeature resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AiFeatureStoreEntityTypeFeature(string name, AiFeatureStoreEntityTypeFeatureArgs args, CustomResourceOptions? options = null)
            : base("gcp:vertex/aiFeatureStoreEntityTypeFeature:AiFeatureStoreEntityTypeFeature", name, args ?? new AiFeatureStoreEntityTypeFeatureArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AiFeatureStoreEntityTypeFeature(string name, Input<string> id, AiFeatureStoreEntityTypeFeatureState? state = null, CustomResourceOptions? options = null)
            : base("gcp:vertex/aiFeatureStoreEntityTypeFeature:AiFeatureStoreEntityTypeFeature", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AiFeatureStoreEntityTypeFeature resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AiFeatureStoreEntityTypeFeature Get(string name, Input<string> id, AiFeatureStoreEntityTypeFeatureState? state = null, CustomResourceOptions? options = null)
        {
            return new AiFeatureStoreEntityTypeFeature(name, id, state, options);
        }
    }

    public sealed class AiFeatureStoreEntityTypeFeatureArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the feature.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}/entityTypes/{entitytype}.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("entitytype", required: true)]
        public Input<string> Entitytype { get; set; } = null!;

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// A set of key/value label pairs to assign to the feature.
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
        /// The name of the feature. The feature can be up to 64 characters long and can consist only of ASCII Latin letters A-Z and a-z, underscore(_), and ASCII digits 0-9 starting with a letter. The value will be unique given an entity type.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Type of Feature value. Immutable. https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featurestores.entityTypes.features#ValueType
        /// </summary>
        [Input("valueType", required: true)]
        public Input<string> ValueType { get; set; } = null!;

        public AiFeatureStoreEntityTypeFeatureArgs()
        {
        }
        public static new AiFeatureStoreEntityTypeFeatureArgs Empty => new AiFeatureStoreEntityTypeFeatureArgs();
    }

    public sealed class AiFeatureStoreEntityTypeFeatureState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The timestamp of when the entity type was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Description of the feature.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

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

        /// <summary>
        /// The name of the Featurestore to use, in the format projects/{project}/locations/{location}/featurestores/{featurestore}/entityTypes/{entitytype}.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("entitytype")]
        public Input<string>? Entitytype { get; set; }

        /// <summary>
        /// Used to perform consistent read-modify-write updates.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// A set of key/value label pairs to assign to the feature.
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
        /// The name of the feature. The feature can be up to 64 characters long and can consist only of ASCII Latin letters A-Z and a-z, underscore(_), and ASCII digits 0-9 starting with a letter. The value will be unique given an entity type.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

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
        /// The region of the feature
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The timestamp when the entity type was most recently updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        /// <summary>
        /// Type of Feature value. Immutable. https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featurestores.entityTypes.features#ValueType
        /// </summary>
        [Input("valueType")]
        public Input<string>? ValueType { get; set; }

        public AiFeatureStoreEntityTypeFeatureState()
        {
        }
        public static new AiFeatureStoreEntityTypeFeatureState Empty => new AiFeatureStoreEntityTypeFeatureState();
    }
}
