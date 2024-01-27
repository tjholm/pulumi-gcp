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
    /// Vertex AI Feature Online Store provides a centralized repository for serving ML features and embedding indexes at low latency. The Feature Online Store is a top-level container.
    /// 
    /// To get more information about FeatureOnlineStore, see:
    /// 
    /// * [API documentation](https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featureOnlineStores)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/vertex-ai/docs)
    /// 
    /// ## Example Usage
    /// ### Vertex Ai Feature Online Store
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var featureOnlineStore = new Gcp.Vertex.AiFeatureOnlineStore("featureOnlineStore", new()
    ///     {
    ///         Bigtable = new Gcp.Vertex.Inputs.AiFeatureOnlineStoreBigtableArgs
    ///         {
    ///             AutoScaling = new Gcp.Vertex.Inputs.AiFeatureOnlineStoreBigtableAutoScalingArgs
    ///             {
    ///                 CpuUtilizationTarget = 50,
    ///                 MaxNodeCount = 3,
    ///                 MinNodeCount = 1,
    ///             },
    ///         },
    ///         Labels = 
    ///         {
    ///             { "foo", "bar" },
    ///         },
    ///         Region = "us-central1",
    ///     });
    /// 
    /// });
    /// ```
    /// ### Vertex Ai Featureonlinestore With Beta Fields Optimized
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var project = Gcp.Organizations.GetProject.Invoke();
    /// 
    ///     var featureonlinestore = new Gcp.Vertex.AiFeatureOnlineStore("featureonlinestore", new()
    ///     {
    ///         Labels = 
    ///         {
    ///             { "foo", "bar" },
    ///         },
    ///         Region = "us-central1",
    ///         Optimized = null,
    ///         DedicatedServingEndpoint = new Gcp.Vertex.Inputs.AiFeatureOnlineStoreDedicatedServingEndpointArgs
    ///         {
    ///             PrivateServiceConnectConfig = new Gcp.Vertex.Inputs.AiFeatureOnlineStoreDedicatedServingEndpointPrivateServiceConnectConfigArgs
    ///             {
    ///                 EnablePrivateServiceConnect = true,
    ///                 ProjectAllowlists = new[]
    ///                 {
    ///                     project.Apply(getProjectResult =&gt; getProjectResult.Number),
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
    /// ### Vertex Ai Featureonlinestore With Beta Fields Bigtable
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var featureonlinestore = new Gcp.Vertex.AiFeatureOnlineStore("featureonlinestore", new()
    ///     {
    ///         Labels = 
    ///         {
    ///             { "foo", "bar" },
    ///         },
    ///         Region = "us-central1",
    ///         Bigtable = new Gcp.Vertex.Inputs.AiFeatureOnlineStoreBigtableArgs
    ///         {
    ///             AutoScaling = new Gcp.Vertex.Inputs.AiFeatureOnlineStoreBigtableAutoScalingArgs
    ///             {
    ///                 MinNodeCount = 1,
    ///                 MaxNodeCount = 2,
    ///                 CpuUtilizationTarget = 80,
    ///             },
    ///         },
    ///         EmbeddingManagement = new Gcp.Vertex.Inputs.AiFeatureOnlineStoreEmbeddingManagementArgs
    ///         {
    ///             Enabled = true,
    ///         },
    ///         ForceDestroy = true,
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    ///     var project = Gcp.Organizations.GetProject.Invoke();
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// FeatureOnlineStore can be imported using any of these accepted formats* `projects/{{project}}/locations/{{region}}/featureOnlineStores/{{name}}` * `{{project}}/{{region}}/{{name}}` * `{{region}}/{{name}}` * `{{name}}` When using the `pulumi import` command, FeatureOnlineStore can be imported using one of the formats above. For example
    /// 
    /// ```sh
    ///  $ pulumi import gcp:vertex/aiFeatureOnlineStore:AiFeatureOnlineStore default projects/{{project}}/locations/{{region}}/featureOnlineStores/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:vertex/aiFeatureOnlineStore:AiFeatureOnlineStore default {{project}}/{{region}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:vertex/aiFeatureOnlineStore:AiFeatureOnlineStore default {{region}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:vertex/aiFeatureOnlineStore:AiFeatureOnlineStore default {{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:vertex/aiFeatureOnlineStore:AiFeatureOnlineStore")]
    public partial class AiFeatureOnlineStore : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Settings for Cloud Bigtable instance that will be created to serve featureValues for all FeatureViews under this FeatureOnlineStore.
        /// Structure is documented below.
        /// </summary>
        [Output("bigtable")]
        public Output<Outputs.AiFeatureOnlineStoreBigtable?> Bigtable { get; private set; } = null!;

        /// <summary>
        /// The timestamp of when the feature online store was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// (Optional, Beta)
        /// The dedicated serving endpoint for this FeatureOnlineStore, which is different from common vertex service endpoint. Only need to set when you choose Optimized storage type or enable EmbeddingManagement. Will use public endpoint by default.
        /// Structure is documented below.
        /// </summary>
        [Output("dedicatedServingEndpoint")]
        public Output<Outputs.AiFeatureOnlineStoreDedicatedServingEndpoint> DedicatedServingEndpoint { get; private set; } = null!;

        /// <summary>
        /// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
        /// </summary>
        [Output("effectiveLabels")]
        public Output<ImmutableDictionary<string, string>> EffectiveLabels { get; private set; } = null!;

        /// <summary>
        /// (Optional, Beta)
        /// The settings for embedding management in FeatureOnlineStore. Embedding management can only be used with BigTable.
        /// Structure is documented below.
        /// </summary>
        [Output("embeddingManagement")]
        public Output<Outputs.AiFeatureOnlineStoreEmbeddingManagement?> EmbeddingManagement { get; private set; } = null!;

        /// <summary>
        /// Used to perform consistent read-modify-write updates.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// If set to true, any FeatureViews and Features for this FeatureOnlineStore will also be deleted.
        /// </summary>
        [Output("forceDestroy")]
        public Output<bool?> ForceDestroy { get; private set; } = null!;

        /// <summary>
        /// The labels with user-defined metadata to organize your feature online stores.
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// The resource name of the Feature Online Store. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// (Optional, Beta)
        /// Settings for the Optimized store that will be created to serve featureValues for all FeatureViews under this FeatureOnlineStore
        /// </summary>
        [Output("optimized")]
        public Output<Outputs.AiFeatureOnlineStoreOptimized?> Optimized { get; private set; } = null!;

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
        /// The region of feature online store. eg us-central1
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The state of the Feature Online Store. See the possible states in [this link](https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featureOnlineStores#state).
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The timestamp of when the feature online store was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a AiFeatureOnlineStore resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AiFeatureOnlineStore(string name, AiFeatureOnlineStoreArgs? args = null, CustomResourceOptions? options = null)
            : base("gcp:vertex/aiFeatureOnlineStore:AiFeatureOnlineStore", name, args ?? new AiFeatureOnlineStoreArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AiFeatureOnlineStore(string name, Input<string> id, AiFeatureOnlineStoreState? state = null, CustomResourceOptions? options = null)
            : base("gcp:vertex/aiFeatureOnlineStore:AiFeatureOnlineStore", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AiFeatureOnlineStore resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AiFeatureOnlineStore Get(string name, Input<string> id, AiFeatureOnlineStoreState? state = null, CustomResourceOptions? options = null)
        {
            return new AiFeatureOnlineStore(name, id, state, options);
        }
    }

    public sealed class AiFeatureOnlineStoreArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Settings for Cloud Bigtable instance that will be created to serve featureValues for all FeatureViews under this FeatureOnlineStore.
        /// Structure is documented below.
        /// </summary>
        [Input("bigtable")]
        public Input<Inputs.AiFeatureOnlineStoreBigtableArgs>? Bigtable { get; set; }

        /// <summary>
        /// (Optional, Beta)
        /// The dedicated serving endpoint for this FeatureOnlineStore, which is different from common vertex service endpoint. Only need to set when you choose Optimized storage type or enable EmbeddingManagement. Will use public endpoint by default.
        /// Structure is documented below.
        /// </summary>
        [Input("dedicatedServingEndpoint")]
        public Input<Inputs.AiFeatureOnlineStoreDedicatedServingEndpointArgs>? DedicatedServingEndpoint { get; set; }

        /// <summary>
        /// (Optional, Beta)
        /// The settings for embedding management in FeatureOnlineStore. Embedding management can only be used with BigTable.
        /// Structure is documented below.
        /// </summary>
        [Input("embeddingManagement")]
        public Input<Inputs.AiFeatureOnlineStoreEmbeddingManagementArgs>? EmbeddingManagement { get; set; }

        /// <summary>
        /// If set to true, any FeatureViews and Features for this FeatureOnlineStore will also be deleted.
        /// </summary>
        [Input("forceDestroy")]
        public Input<bool>? ForceDestroy { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The labels with user-defined metadata to organize your feature online stores.
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The resource name of the Feature Online Store. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// (Optional, Beta)
        /// Settings for the Optimized store that will be created to serve featureValues for all FeatureViews under this FeatureOnlineStore
        /// </summary>
        [Input("optimized")]
        public Input<Inputs.AiFeatureOnlineStoreOptimizedArgs>? Optimized { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The region of feature online store. eg us-central1
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public AiFeatureOnlineStoreArgs()
        {
        }
        public static new AiFeatureOnlineStoreArgs Empty => new AiFeatureOnlineStoreArgs();
    }

    public sealed class AiFeatureOnlineStoreState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Settings for Cloud Bigtable instance that will be created to serve featureValues for all FeatureViews under this FeatureOnlineStore.
        /// Structure is documented below.
        /// </summary>
        [Input("bigtable")]
        public Input<Inputs.AiFeatureOnlineStoreBigtableGetArgs>? Bigtable { get; set; }

        /// <summary>
        /// The timestamp of when the feature online store was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// (Optional, Beta)
        /// The dedicated serving endpoint for this FeatureOnlineStore, which is different from common vertex service endpoint. Only need to set when you choose Optimized storage type or enable EmbeddingManagement. Will use public endpoint by default.
        /// Structure is documented below.
        /// </summary>
        [Input("dedicatedServingEndpoint")]
        public Input<Inputs.AiFeatureOnlineStoreDedicatedServingEndpointGetArgs>? DedicatedServingEndpoint { get; set; }

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
        /// (Optional, Beta)
        /// The settings for embedding management in FeatureOnlineStore. Embedding management can only be used with BigTable.
        /// Structure is documented below.
        /// </summary>
        [Input("embeddingManagement")]
        public Input<Inputs.AiFeatureOnlineStoreEmbeddingManagementGetArgs>? EmbeddingManagement { get; set; }

        /// <summary>
        /// Used to perform consistent read-modify-write updates.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// If set to true, any FeatureViews and Features for this FeatureOnlineStore will also be deleted.
        /// </summary>
        [Input("forceDestroy")]
        public Input<bool>? ForceDestroy { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The labels with user-defined metadata to organize your feature online stores.
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The resource name of the Feature Online Store. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// (Optional, Beta)
        /// Settings for the Optimized store that will be created to serve featureValues for all FeatureViews under this FeatureOnlineStore
        /// </summary>
        [Input("optimized")]
        public Input<Inputs.AiFeatureOnlineStoreOptimizedGetArgs>? Optimized { get; set; }

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
        /// The region of feature online store. eg us-central1
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The state of the Feature Online Store. See the possible states in [this link](https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.featureOnlineStores#state).
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// The timestamp of when the feature online store was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public AiFeatureOnlineStoreState()
        {
        }
        public static new AiFeatureOnlineStoreState Empty => new AiFeatureOnlineStoreState();
    }
}
