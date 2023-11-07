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
    /// A representation of a collection of database items organized in a way that allows for approximate nearest neighbor (a.k.a ANN) algorithms search.
    /// 
    /// To get more information about Index, see:
    /// 
    /// * [API documentation](https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.indexes/)
    /// 
    /// ## Example Usage
    /// ### Vertex Ai Index
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var bucket = new Gcp.Storage.Bucket("bucket", new()
    ///     {
    ///         Location = "us-central1",
    ///         UniformBucketLevelAccess = true,
    ///     });
    /// 
    ///     // The sample data comes from the following link:
    ///     // https://cloud.google.com/vertex-ai/docs/matching-engine/filtering#specify-namespaces-tokens
    ///     var data = new Gcp.Storage.BucketObject("data", new()
    ///     {
    ///         Bucket = bucket.Name,
    ///         Content = @"{""id"": ""42"", ""embedding"": [0.5, 1.0], ""restricts"": [{""namespace"": ""class"", ""allow"": [""cat"", ""pet""]},{""namespace"": ""category"", ""allow"": [""feline""]}]}
    /// {""id"": ""43"", ""embedding"": [0.6, 1.0], ""restricts"": [{""namespace"": ""class"", ""allow"": [""dog"", ""pet""]},{""namespace"": ""category"", ""allow"": [""canine""]}]}
    /// ",
    ///     });
    /// 
    ///     var index = new Gcp.Vertex.AiIndex("index", new()
    ///     {
    ///         Labels = 
    ///         {
    ///             { "foo", "bar" },
    ///         },
    ///         Region = "us-central1",
    ///         DisplayName = "test-index",
    ///         Description = "index for test",
    ///         Metadata = new Gcp.Vertex.Inputs.AiIndexMetadataArgs
    ///         {
    ///             ContentsDeltaUri = bucket.Name.Apply(name =&gt; $"gs://{name}/contents"),
    ///             Config = new Gcp.Vertex.Inputs.AiIndexMetadataConfigArgs
    ///             {
    ///                 Dimensions = 2,
    ///                 ApproximateNeighborsCount = 150,
    ///                 ShardSize = "SHARD_SIZE_SMALL",
    ///                 DistanceMeasureType = "DOT_PRODUCT_DISTANCE",
    ///                 AlgorithmConfig = new Gcp.Vertex.Inputs.AiIndexMetadataConfigAlgorithmConfigArgs
    ///                 {
    ///                     TreeAhConfig = new Gcp.Vertex.Inputs.AiIndexMetadataConfigAlgorithmConfigTreeAhConfigArgs
    ///                     {
    ///                         LeafNodeEmbeddingCount = 500,
    ///                         LeafNodesToSearchPercent = 7,
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///         IndexUpdateMethod = "BATCH_UPDATE",
    ///     });
    /// 
    /// });
    /// ```
    /// ### Vertex Ai Index Streaming
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var bucket = new Gcp.Storage.Bucket("bucket", new()
    ///     {
    ///         Location = "us-central1",
    ///         UniformBucketLevelAccess = true,
    ///     });
    /// 
    ///     // The sample data comes from the following link:
    ///     // https://cloud.google.com/vertex-ai/docs/matching-engine/filtering#specify-namespaces-tokens
    ///     var data = new Gcp.Storage.BucketObject("data", new()
    ///     {
    ///         Bucket = bucket.Name,
    ///         Content = @"{""id"": ""42"", ""embedding"": [0.5, 1.0], ""restricts"": [{""namespace"": ""class"", ""allow"": [""cat"", ""pet""]},{""namespace"": ""category"", ""allow"": [""feline""]}]}
    /// {""id"": ""43"", ""embedding"": [0.6, 1.0], ""restricts"": [{""namespace"": ""class"", ""allow"": [""dog"", ""pet""]},{""namespace"": ""category"", ""allow"": [""canine""]}]}
    /// ",
    ///     });
    /// 
    ///     var index = new Gcp.Vertex.AiIndex("index", new()
    ///     {
    ///         Labels = 
    ///         {
    ///             { "foo", "bar" },
    ///         },
    ///         Region = "us-central1",
    ///         DisplayName = "test-index",
    ///         Description = "index for test",
    ///         Metadata = new Gcp.Vertex.Inputs.AiIndexMetadataArgs
    ///         {
    ///             ContentsDeltaUri = bucket.Name.Apply(name =&gt; $"gs://{name}/contents"),
    ///             Config = new Gcp.Vertex.Inputs.AiIndexMetadataConfigArgs
    ///             {
    ///                 Dimensions = 2,
    ///                 ShardSize = "SHARD_SIZE_LARGE",
    ///                 DistanceMeasureType = "COSINE_DISTANCE",
    ///                 FeatureNormType = "UNIT_L2_NORM",
    ///                 AlgorithmConfig = new Gcp.Vertex.Inputs.AiIndexMetadataConfigAlgorithmConfigArgs
    ///                 {
    ///                     BruteForceConfig = null,
    ///                 },
    ///             },
    ///         },
    ///         IndexUpdateMethod = "STREAM_UPDATE",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Index can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:vertex/aiIndex:AiIndex default projects/{{project}}/locations/{{region}}/indexes/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:vertex/aiIndex:AiIndex default {{project}}/{{region}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:vertex/aiIndex:AiIndex default {{region}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:vertex/aiIndex:AiIndex default {{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:vertex/aiIndex:AiIndex")]
    public partial class AiIndex : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The timestamp of when the Index was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The pointers to DeployedIndexes created from this Index. An Index can be only deleted if all its DeployedIndexes had been undeployed first.
        /// Structure is documented below.
        /// </summary>
        [Output("deployedIndexes")]
        public Output<ImmutableArray<Outputs.AiIndexDeployedIndex>> DeployedIndexes { get; private set; } = null!;

        /// <summary>
        /// The description of the Index.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The display name of the Index. The name can be up to 128 characters long and can consist of any UTF-8 characters.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
        /// clients and services.
        /// </summary>
        [Output("effectiveLabels")]
        public Output<ImmutableDictionary<string, string>> EffectiveLabels { get; private set; } = null!;

        /// <summary>
        /// Used to perform consistent read-modify-write updates.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Stats of the index resource.
        /// Structure is documented below.
        /// </summary>
        [Output("indexStats")]
        public Output<ImmutableArray<Outputs.AiIndexIndexStat>> IndexStats { get; private set; } = null!;

        /// <summary>
        /// The update method to use with this Index. The value must be the followings. If not set, BATCH_UPDATE will be used by default.
        /// * BATCH_UPDATE: user can call indexes.patch with files on Cloud Storage of datapoints to update.
        /// * STREAM_UPDATE: user can call indexes.upsertDatapoints/DeleteDatapoints to update the Index and the updates will be applied in corresponding DeployedIndexes in nearly real-time.
        /// </summary>
        [Output("indexUpdateMethod")]
        public Output<string?> IndexUpdateMethod { get; private set; } = null!;

        /// <summary>
        /// The labels with user-defined metadata to organize your Indexes.
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// An additional information about the Index
        /// Structure is documented below.
        /// </summary>
        [Output("metadata")]
        public Output<Outputs.AiIndexMetadata?> Metadata { get; private set; } = null!;

        /// <summary>
        /// Points to a YAML file stored on Google Cloud Storage describing additional information about the Index, that is specific to it. Unset if the Index does not have any additional information.
        /// </summary>
        [Output("metadataSchemaUri")]
        public Output<string> MetadataSchemaUri { get; private set; } = null!;

        /// <summary>
        /// The resource name of the Index.
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
        /// The region of the index. eg us-central1
        /// </summary>
        [Output("region")]
        public Output<string?> Region { get; private set; } = null!;

        /// <summary>
        /// The timestamp of when the Index was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a AiIndex resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AiIndex(string name, AiIndexArgs args, CustomResourceOptions? options = null)
            : base("gcp:vertex/aiIndex:AiIndex", name, args ?? new AiIndexArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AiIndex(string name, Input<string> id, AiIndexState? state = null, CustomResourceOptions? options = null)
            : base("gcp:vertex/aiIndex:AiIndex", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AiIndex resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AiIndex Get(string name, Input<string> id, AiIndexState? state = null, CustomResourceOptions? options = null)
        {
            return new AiIndex(name, id, state, options);
        }
    }

    public sealed class AiIndexArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the Index.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The display name of the Index. The name can be up to 128 characters long and can consist of any UTF-8 characters.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// The update method to use with this Index. The value must be the followings. If not set, BATCH_UPDATE will be used by default.
        /// * BATCH_UPDATE: user can call indexes.patch with files on Cloud Storage of datapoints to update.
        /// * STREAM_UPDATE: user can call indexes.upsertDatapoints/DeleteDatapoints to update the Index and the updates will be applied in corresponding DeployedIndexes in nearly real-time.
        /// </summary>
        [Input("indexUpdateMethod")]
        public Input<string>? IndexUpdateMethod { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The labels with user-defined metadata to organize your Indexes.
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// An additional information about the Index
        /// Structure is documented below.
        /// </summary>
        [Input("metadata")]
        public Input<Inputs.AiIndexMetadataArgs>? Metadata { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The region of the index. eg us-central1
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public AiIndexArgs()
        {
        }
        public static new AiIndexArgs Empty => new AiIndexArgs();
    }

    public sealed class AiIndexState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The timestamp of when the Index was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        [Input("deployedIndexes")]
        private InputList<Inputs.AiIndexDeployedIndexGetArgs>? _deployedIndexes;

        /// <summary>
        /// The pointers to DeployedIndexes created from this Index. An Index can be only deleted if all its DeployedIndexes had been undeployed first.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.AiIndexDeployedIndexGetArgs> DeployedIndexes
        {
            get => _deployedIndexes ?? (_deployedIndexes = new InputList<Inputs.AiIndexDeployedIndexGetArgs>());
            set => _deployedIndexes = value;
        }

        /// <summary>
        /// The description of the Index.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The display name of the Index. The name can be up to 128 characters long and can consist of any UTF-8 characters.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

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

        /// <summary>
        /// Used to perform consistent read-modify-write updates.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("indexStats")]
        private InputList<Inputs.AiIndexIndexStatGetArgs>? _indexStats;

        /// <summary>
        /// Stats of the index resource.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.AiIndexIndexStatGetArgs> IndexStats
        {
            get => _indexStats ?? (_indexStats = new InputList<Inputs.AiIndexIndexStatGetArgs>());
            set => _indexStats = value;
        }

        /// <summary>
        /// The update method to use with this Index. The value must be the followings. If not set, BATCH_UPDATE will be used by default.
        /// * BATCH_UPDATE: user can call indexes.patch with files on Cloud Storage of datapoints to update.
        /// * STREAM_UPDATE: user can call indexes.upsertDatapoints/DeleteDatapoints to update the Index and the updates will be applied in corresponding DeployedIndexes in nearly real-time.
        /// </summary>
        [Input("indexUpdateMethod")]
        public Input<string>? IndexUpdateMethod { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The labels with user-defined metadata to organize your Indexes.
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// An additional information about the Index
        /// Structure is documented below.
        /// </summary>
        [Input("metadata")]
        public Input<Inputs.AiIndexMetadataGetArgs>? Metadata { get; set; }

        /// <summary>
        /// Points to a YAML file stored on Google Cloud Storage describing additional information about the Index, that is specific to it. Unset if the Index does not have any additional information.
        /// </summary>
        [Input("metadataSchemaUri")]
        public Input<string>? MetadataSchemaUri { get; set; }

        /// <summary>
        /// The resource name of the Index.
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

        /// <summary>
        /// The region of the index. eg us-central1
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The timestamp of when the Index was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public AiIndexState()
        {
        }
        public static new AiIndexState Empty => new AiIndexState();
    }
}
