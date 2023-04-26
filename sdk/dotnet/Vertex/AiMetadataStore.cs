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
    /// ## Example Usage
    /// ### Vertex Ai Metadata Store
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var store = new Gcp.Vertex.AiMetadataStore("store", new()
    ///     {
    ///         Description = "Store to test the terraform module",
    ///         Region = "us-central1",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// MetadataStore can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:vertex/aiMetadataStore:AiMetadataStore default projects/{{project}}/locations/{{region}}/metadataStores/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:vertex/aiMetadataStore:AiMetadataStore default {{project}}/{{region}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:vertex/aiMetadataStore:AiMetadataStore default {{region}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:vertex/aiMetadataStore:AiMetadataStore default {{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:vertex/aiMetadataStore:AiMetadataStore")]
    public partial class AiMetadataStore : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The timestamp of when the MetadataStore was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Description of the MetadataStore.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Customer-managed encryption key spec for a MetadataStore. If set, this MetadataStore and all sub-resources of this MetadataStore will be secured by this key.
        /// Structure is documented below.
        /// </summary>
        [Output("encryptionSpec")]
        public Output<Outputs.AiMetadataStoreEncryptionSpec?> EncryptionSpec { get; private set; } = null!;

        /// <summary>
        /// The name of the MetadataStore. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
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
        /// The region of the Metadata Store. eg us-central1
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// State information of the MetadataStore.
        /// Structure is documented below.
        /// </summary>
        [Output("states")]
        public Output<ImmutableArray<Outputs.AiMetadataStoreState>> States { get; private set; } = null!;

        /// <summary>
        /// The timestamp of when the MetadataStore was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a AiMetadataStore resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AiMetadataStore(string name, AiMetadataStoreArgs? args = null, CustomResourceOptions? options = null)
            : base("gcp:vertex/aiMetadataStore:AiMetadataStore", name, args ?? new AiMetadataStoreArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AiMetadataStore(string name, Input<string> id, AiMetadataStoreState? state = null, CustomResourceOptions? options = null)
            : base("gcp:vertex/aiMetadataStore:AiMetadataStore", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AiMetadataStore resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AiMetadataStore Get(string name, Input<string> id, AiMetadataStoreState? state = null, CustomResourceOptions? options = null)
        {
            return new AiMetadataStore(name, id, state, options);
        }
    }

    public sealed class AiMetadataStoreArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the MetadataStore.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Customer-managed encryption key spec for a MetadataStore. If set, this MetadataStore and all sub-resources of this MetadataStore will be secured by this key.
        /// Structure is documented below.
        /// </summary>
        [Input("encryptionSpec")]
        public Input<Inputs.AiMetadataStoreEncryptionSpecArgs>? EncryptionSpec { get; set; }

        /// <summary>
        /// The name of the MetadataStore. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
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
        /// The region of the Metadata Store. eg us-central1
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public AiMetadataStoreArgs()
        {
        }
        public static new AiMetadataStoreArgs Empty => new AiMetadataStoreArgs();
    }

    public sealed class AiMetadataStoreState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The timestamp of when the MetadataStore was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Description of the MetadataStore.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Customer-managed encryption key spec for a MetadataStore. If set, this MetadataStore and all sub-resources of this MetadataStore will be secured by this key.
        /// Structure is documented below.
        /// </summary>
        [Input("encryptionSpec")]
        public Input<Inputs.AiMetadataStoreEncryptionSpecGetArgs>? EncryptionSpec { get; set; }

        /// <summary>
        /// The name of the MetadataStore. This value may be up to 60 characters, and valid characters are [a-z0-9_]. The first character cannot be a number.
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
        /// The region of the Metadata Store. eg us-central1
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("states")]
        private InputList<Inputs.AiMetadataStoreStateGetArgs>? _states;

        /// <summary>
        /// State information of the MetadataStore.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.AiMetadataStoreStateGetArgs> States
        {
            get => _states ?? (_states = new InputList<Inputs.AiMetadataStoreStateGetArgs>());
            set => _states = value;
        }

        /// <summary>
        /// The timestamp of when the MetadataStore was last updated in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public AiMetadataStoreState()
        {
        }
        public static new AiMetadataStoreState Empty => new AiMetadataStoreState();
    }
}
