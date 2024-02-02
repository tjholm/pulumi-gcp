// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Firestore
{
    /// <summary>
    /// Cloud Firestore indexes enable simple and complex queries against documents in a database.
    ///  This resource manages composite indexes and not single
    /// field indexes.
    /// 
    /// To get more information about Index, see:
    /// 
    /// * [API documentation](https://cloud.google.com/firestore/docs/reference/rest/v1/projects.databases.collectionGroups.indexes)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/firestore/docs/query-data/indexing)
    /// 
    /// &gt; **Warning:** This resource creates a Firestore Index on a project that already has
    /// a Firestore database. If you haven't already created it, you may
    /// create a `gcp.firestore.Database` resource and `location_id` set
    /// to your chosen location. If you wish to use App Engine, you may
    /// instead create a `gcp.appengine.Application` resource with
    /// `database_type` set to `"CLOUD_FIRESTORE"`. Your Firestore location
    /// will be the same as the App Engine location specified.
    /// 
    /// ## Example Usage
    /// ### Firestore Index Datastore Mode
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var my_index = new Gcp.Firestore.Index("my-index", new()
    ///     {
    ///         ApiScope = "DATASTORE_MODE_API",
    ///         Collection = "chatrooms",
    ///         Database = "(default)",
    ///         Fields = new[]
    ///         {
    ///             new Gcp.Firestore.Inputs.IndexFieldArgs
    ///             {
    ///                 FieldPath = "name",
    ///                 Order = "ASCENDING",
    ///             },
    ///             new Gcp.Firestore.Inputs.IndexFieldArgs
    ///             {
    ///                 FieldPath = "description",
    ///                 Order = "DESCENDING",
    ///             },
    ///         },
    ///         Project = "my-project-name",
    ///         QueryScope = "COLLECTION_RECURSIVE",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Index can be imported using any of these accepted formats* `{{name}}` When using the `pulumi import` command, Index can be imported using one of the formats above. For example
    /// 
    /// ```sh
    ///  $ pulumi import gcp:firestore/index:Index default {{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:firestore/index:Index")]
    public partial class Index : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The API scope at which a query is run.
        /// Default value is `ANY_API`.
        /// Possible values are: `ANY_API`, `DATASTORE_MODE_API`.
        /// </summary>
        [Output("apiScope")]
        public Output<string?> ApiScope { get; private set; } = null!;

        /// <summary>
        /// The collection being indexed.
        /// </summary>
        [Output("collection")]
        public Output<string> Collection { get; private set; } = null!;

        /// <summary>
        /// The Firestore database id. Defaults to `"(default)"`.
        /// </summary>
        [Output("database")]
        public Output<string?> Database { get; private set; } = null!;

        /// <summary>
        /// The fields supported by this index. The last field entry is always for
        /// the field path `__name__`. If, on creation, `__name__` was not
        /// specified as the last field, it will be added automatically with the
        /// same direction as that of the last field defined. If the final field
        /// in a composite index is not directional, the `__name__` will be
        /// ordered `"ASCENDING"` (unless explicitly specified otherwise).
        /// Structure is documented below.
        /// </summary>
        [Output("fields")]
        public Output<ImmutableArray<Outputs.IndexField>> Fields { get; private set; } = null!;

        /// <summary>
        /// A server defined name for this index. Format:
        /// `projects/{{project}}/databases/{{database}}/collectionGroups/{{collection}}/indexes/{{server_generated_id}}`
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
        /// The scope at which a query is run.
        /// Default value is `COLLECTION`.
        /// Possible values are: `COLLECTION`, `COLLECTION_GROUP`, `COLLECTION_RECURSIVE`.
        /// </summary>
        [Output("queryScope")]
        public Output<string?> QueryScope { get; private set; } = null!;


        /// <summary>
        /// Create a Index resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Index(string name, IndexArgs args, CustomResourceOptions? options = null)
            : base("gcp:firestore/index:Index", name, args ?? new IndexArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Index(string name, Input<string> id, IndexState? state = null, CustomResourceOptions? options = null)
            : base("gcp:firestore/index:Index", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Index resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Index Get(string name, Input<string> id, IndexState? state = null, CustomResourceOptions? options = null)
        {
            return new Index(name, id, state, options);
        }
    }

    public sealed class IndexArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The API scope at which a query is run.
        /// Default value is `ANY_API`.
        /// Possible values are: `ANY_API`, `DATASTORE_MODE_API`.
        /// </summary>
        [Input("apiScope")]
        public Input<string>? ApiScope { get; set; }

        /// <summary>
        /// The collection being indexed.
        /// </summary>
        [Input("collection", required: true)]
        public Input<string> Collection { get; set; } = null!;

        /// <summary>
        /// The Firestore database id. Defaults to `"(default)"`.
        /// </summary>
        [Input("database")]
        public Input<string>? Database { get; set; }

        [Input("fields", required: true)]
        private InputList<Inputs.IndexFieldArgs>? _fields;

        /// <summary>
        /// The fields supported by this index. The last field entry is always for
        /// the field path `__name__`. If, on creation, `__name__` was not
        /// specified as the last field, it will be added automatically with the
        /// same direction as that of the last field defined. If the final field
        /// in a composite index is not directional, the `__name__` will be
        /// ordered `"ASCENDING"` (unless explicitly specified otherwise).
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.IndexFieldArgs> Fields
        {
            get => _fields ?? (_fields = new InputList<Inputs.IndexFieldArgs>());
            set => _fields = value;
        }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The scope at which a query is run.
        /// Default value is `COLLECTION`.
        /// Possible values are: `COLLECTION`, `COLLECTION_GROUP`, `COLLECTION_RECURSIVE`.
        /// </summary>
        [Input("queryScope")]
        public Input<string>? QueryScope { get; set; }

        public IndexArgs()
        {
        }
        public static new IndexArgs Empty => new IndexArgs();
    }

    public sealed class IndexState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The API scope at which a query is run.
        /// Default value is `ANY_API`.
        /// Possible values are: `ANY_API`, `DATASTORE_MODE_API`.
        /// </summary>
        [Input("apiScope")]
        public Input<string>? ApiScope { get; set; }

        /// <summary>
        /// The collection being indexed.
        /// </summary>
        [Input("collection")]
        public Input<string>? Collection { get; set; }

        /// <summary>
        /// The Firestore database id. Defaults to `"(default)"`.
        /// </summary>
        [Input("database")]
        public Input<string>? Database { get; set; }

        [Input("fields")]
        private InputList<Inputs.IndexFieldGetArgs>? _fields;

        /// <summary>
        /// The fields supported by this index. The last field entry is always for
        /// the field path `__name__`. If, on creation, `__name__` was not
        /// specified as the last field, it will be added automatically with the
        /// same direction as that of the last field defined. If the final field
        /// in a composite index is not directional, the `__name__` will be
        /// ordered `"ASCENDING"` (unless explicitly specified otherwise).
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.IndexFieldGetArgs> Fields
        {
            get => _fields ?? (_fields = new InputList<Inputs.IndexFieldGetArgs>());
            set => _fields = value;
        }

        /// <summary>
        /// A server defined name for this index. Format:
        /// `projects/{{project}}/databases/{{database}}/collectionGroups/{{collection}}/indexes/{{server_generated_id}}`
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
        /// The scope at which a query is run.
        /// Default value is `COLLECTION`.
        /// Possible values are: `COLLECTION`, `COLLECTION_GROUP`, `COLLECTION_RECURSIVE`.
        /// </summary>
        [Input("queryScope")]
        public Input<string>? QueryScope { get; set; }

        public IndexState()
        {
        }
        public static new IndexState Empty => new IndexState();
    }
}
