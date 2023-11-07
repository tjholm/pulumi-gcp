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
    /// The Dataplex Lake resource
    /// 
    /// ## Example Usage
    /// ### Basic_lake
    /// A basic example of a dataplex lake
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var primary = new Gcp.DataPlex.Lake("primary", new()
    ///     {
    ///         Description = "Lake for DCL",
    ///         DisplayName = "Lake for DCL",
    ///         Labels = 
    ///         {
    ///             { "my-lake", "exists" },
    ///         },
    ///         Location = "us-west1",
    ///         Project = "my-project-name",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Lake can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:dataplex/lake:Lake default projects/{{project}}/locations/{{location}}/lakes/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:dataplex/lake:Lake default {{project}}/{{location}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:dataplex/lake:Lake default {{location}}/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:dataplex/lake:Lake")]
    public partial class Lake : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Output only. Aggregated status of the underlying assets of the lake.
        /// </summary>
        [Output("assetStatuses")]
        public Output<ImmutableArray<Outputs.LakeAssetStatus>> AssetStatuses { get; private set; } = null!;

        /// <summary>
        /// Output only. The time when the lake was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Optional. Description of the lake.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Optional. User friendly display name.
        /// </summary>
        [Output("displayName")]
        public Output<string?> DisplayName { get; private set; } = null!;

        /// <summary>
        /// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
        /// clients and services.
        /// </summary>
        [Output("effectiveLabels")]
        public Output<ImmutableDictionary<string, object>> EffectiveLabels { get; private set; } = null!;

        /// <summary>
        /// Optional. User-defined labels for the lake.
        /// 
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// The location for the resource
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Optional. Settings to manage lake and Dataproc Metastore service instance association.
        /// </summary>
        [Output("metastore")]
        public Output<Outputs.LakeMetastore?> Metastore { get; private set; } = null!;

        /// <summary>
        /// Output only. Metastore status of the lake.
        /// </summary>
        [Output("metastoreStatuses")]
        public Output<ImmutableArray<Outputs.LakeMetastoreStatus>> MetastoreStatuses { get; private set; } = null!;

        /// <summary>
        /// The name of the lake.
        /// 
        /// 
        /// 
        /// - - -
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
        /// Output only. Service account associated with this lake. This service account must be authorized to access or operate on resources managed by the lake.
        /// </summary>
        [Output("serviceAccount")]
        public Output<string> ServiceAccount { get; private set; } = null!;

        /// <summary>
        /// Output only. Current state of the lake. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, DELETING, ACTION_REQUIRED
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Output only. System generated globally unique ID for the lake. This ID will be different if the lake is deleted and re-created with the same name.
        /// </summary>
        [Output("uid")]
        public Output<string> Uid { get; private set; } = null!;

        /// <summary>
        /// Output only. The time when the lake was last updated.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a Lake resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Lake(string name, LakeArgs args, CustomResourceOptions? options = null)
            : base("gcp:dataplex/lake:Lake", name, args ?? new LakeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Lake(string name, Input<string> id, LakeState? state = null, CustomResourceOptions? options = null)
            : base("gcp:dataplex/lake:Lake", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Lake resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Lake Get(string name, Input<string> id, LakeState? state = null, CustomResourceOptions? options = null)
        {
            return new Lake(name, id, state, options);
        }
    }

    public sealed class LakeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Description of the lake.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Optional. User friendly display name.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Optional. User-defined labels for the lake.
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
        /// The location for the resource
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Optional. Settings to manage lake and Dataproc Metastore service instance association.
        /// </summary>
        [Input("metastore")]
        public Input<Inputs.LakeMetastoreArgs>? Metastore { get; set; }

        /// <summary>
        /// The name of the lake.
        /// 
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The project for the resource
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public LakeArgs()
        {
        }
        public static new LakeArgs Empty => new LakeArgs();
    }

    public sealed class LakeState : global::Pulumi.ResourceArgs
    {
        [Input("assetStatuses")]
        private InputList<Inputs.LakeAssetStatusGetArgs>? _assetStatuses;

        /// <summary>
        /// Output only. Aggregated status of the underlying assets of the lake.
        /// </summary>
        public InputList<Inputs.LakeAssetStatusGetArgs> AssetStatuses
        {
            get => _assetStatuses ?? (_assetStatuses = new InputList<Inputs.LakeAssetStatusGetArgs>());
            set => _assetStatuses = value;
        }

        /// <summary>
        /// Output only. The time when the lake was created.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Optional. Description of the lake.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Optional. User friendly display name.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("effectiveLabels")]
        private InputMap<object>? _effectiveLabels;

        /// <summary>
        /// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
        /// clients and services.
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
        /// Optional. User-defined labels for the lake.
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
        /// The location for the resource
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Optional. Settings to manage lake and Dataproc Metastore service instance association.
        /// </summary>
        [Input("metastore")]
        public Input<Inputs.LakeMetastoreGetArgs>? Metastore { get; set; }

        [Input("metastoreStatuses")]
        private InputList<Inputs.LakeMetastoreStatusGetArgs>? _metastoreStatuses;

        /// <summary>
        /// Output only. Metastore status of the lake.
        /// </summary>
        public InputList<Inputs.LakeMetastoreStatusGetArgs> MetastoreStatuses
        {
            get => _metastoreStatuses ?? (_metastoreStatuses = new InputList<Inputs.LakeMetastoreStatusGetArgs>());
            set => _metastoreStatuses = value;
        }

        /// <summary>
        /// The name of the lake.
        /// 
        /// 
        /// 
        /// - - -
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
        /// Output only. Service account associated with this lake. This service account must be authorized to access or operate on resources managed by the lake.
        /// </summary>
        [Input("serviceAccount")]
        public Input<string>? ServiceAccount { get; set; }

        /// <summary>
        /// Output only. Current state of the lake. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, DELETING, ACTION_REQUIRED
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Output only. System generated globally unique ID for the lake. This ID will be different if the lake is deleted and re-created with the same name.
        /// </summary>
        [Input("uid")]
        public Input<string>? Uid { get; set; }

        /// <summary>
        /// Output only. The time when the lake was last updated.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public LakeState()
        {
        }
        public static new LakeState Empty => new LakeState();
    }
}
