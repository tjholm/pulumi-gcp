// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.BigLake
{
    /// <summary>
    /// Databases are containers of tables.
    /// 
    /// To get more information about Database, see:
    /// 
    /// * [API documentation](https://cloud.google.com/bigquery/docs/reference/biglake/rest/v1/projects.locations.catalogs.databases)
    /// * How-to Guides
    ///     * [Manage open source metadata with BigLake Metastore](https://cloud.google.com/bigquery/docs/manage-open-source-metadata#create_databases)
    /// 
    /// ## Example Usage
    /// ### Biglake Database
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var catalog = new Gcp.BigLake.Catalog("catalog", new()
    ///     {
    ///         Location = "US",
    ///     });
    /// 
    ///     var bucket = new Gcp.Storage.Bucket("bucket", new()
    ///     {
    ///         Location = "US",
    ///         ForceDestroy = true,
    ///         UniformBucketLevelAccess = true,
    ///     });
    /// 
    ///     var metadataFolder = new Gcp.Storage.BucketObject("metadataFolder", new()
    ///     {
    ///         Content = " ",
    ///         Bucket = bucket.Name,
    ///     });
    /// 
    ///     var database = new Gcp.BigLake.Database("database", new()
    ///     {
    ///         Catalog = catalog.Id,
    ///         Type = "HIVE",
    ///         HiveOptions = new Gcp.BigLake.Inputs.DatabaseHiveOptionsArgs
    ///         {
    ///             LocationUri = Output.Tuple(bucket.Name, metadataFolder.Name).Apply(values =&gt;
    ///             {
    ///                 var bucketName = values.Item1;
    ///                 var metadataFolderName = values.Item2;
    ///                 return $"gs://{bucketName}/{metadataFolderName}";
    ///             }),
    ///             Parameters = 
    ///             {
    ///                 { "owner", "John Doe" },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Database can be imported using any of these accepted formats* `{{catalog}}/databases/{{name}}` When using the `pulumi import` command, Database can be imported using one of the formats above. For example
    /// 
    /// ```sh
    ///  $ pulumi import gcp:biglake/database:Database default {{catalog}}/databases/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:biglake/database:Database")]
    public partial class Database : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The parent catalog.
        /// </summary>
        [Output("catalog")]
        public Output<string> Catalog { get; private set; } = null!;

        /// <summary>
        /// Output only. The creation time of the database. A timestamp in RFC3339
        /// UTC "Zulu" format, with nanosecond resolution and up to nine fractional
        /// digits. Examples: "2014-10-02T15:01:23Z" and
        /// "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Output only. The deletion time of the database. Only set after the
        /// database is deleted. A timestamp in RFC3339 UTC "Zulu" format, with
        /// nanosecond resolution and up to nine fractional digits. Examples:
        /// "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Output("deleteTime")]
        public Output<string> DeleteTime { get; private set; } = null!;

        /// <summary>
        /// Output only. The time when this database is considered expired. Only set
        /// after the database is deleted. A timestamp in RFC3339 UTC "Zulu" format,
        /// with nanosecond resolution and up to nine fractional digits. Examples:
        /// "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Output("expireTime")]
        public Output<string> ExpireTime { get; private set; } = null!;

        /// <summary>
        /// Options of a Hive database.
        /// Structure is documented below.
        /// </summary>
        [Output("hiveOptions")]
        public Output<Outputs.DatabaseHiveOptions> HiveOptions { get; private set; } = null!;

        /// <summary>
        /// The name of the database.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The database type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Output only. The last modification time of the database. A timestamp in
        /// RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
        /// fractional digits. Examples: "2014-10-02T15:01:23Z" and
        /// "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a Database resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Database(string name, DatabaseArgs args, CustomResourceOptions? options = null)
            : base("gcp:biglake/database:Database", name, args ?? new DatabaseArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Database(string name, Input<string> id, DatabaseState? state = null, CustomResourceOptions? options = null)
            : base("gcp:biglake/database:Database", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Database resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Database Get(string name, Input<string> id, DatabaseState? state = null, CustomResourceOptions? options = null)
        {
            return new Database(name, id, state, options);
        }
    }

    public sealed class DatabaseArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The parent catalog.
        /// </summary>
        [Input("catalog", required: true)]
        public Input<string> Catalog { get; set; } = null!;

        /// <summary>
        /// Options of a Hive database.
        /// Structure is documented below.
        /// </summary>
        [Input("hiveOptions", required: true)]
        public Input<Inputs.DatabaseHiveOptionsArgs> HiveOptions { get; set; } = null!;

        /// <summary>
        /// The name of the database.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The database type.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public DatabaseArgs()
        {
        }
        public static new DatabaseArgs Empty => new DatabaseArgs();
    }

    public sealed class DatabaseState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The parent catalog.
        /// </summary>
        [Input("catalog")]
        public Input<string>? Catalog { get; set; }

        /// <summary>
        /// Output only. The creation time of the database. A timestamp in RFC3339
        /// UTC "Zulu" format, with nanosecond resolution and up to nine fractional
        /// digits. Examples: "2014-10-02T15:01:23Z" and
        /// "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Output only. The deletion time of the database. Only set after the
        /// database is deleted. A timestamp in RFC3339 UTC "Zulu" format, with
        /// nanosecond resolution and up to nine fractional digits. Examples:
        /// "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Input("deleteTime")]
        public Input<string>? DeleteTime { get; set; }

        /// <summary>
        /// Output only. The time when this database is considered expired. Only set
        /// after the database is deleted. A timestamp in RFC3339 UTC "Zulu" format,
        /// with nanosecond resolution and up to nine fractional digits. Examples:
        /// "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Input("expireTime")]
        public Input<string>? ExpireTime { get; set; }

        /// <summary>
        /// Options of a Hive database.
        /// Structure is documented below.
        /// </summary>
        [Input("hiveOptions")]
        public Input<Inputs.DatabaseHiveOptionsGetArgs>? HiveOptions { get; set; }

        /// <summary>
        /// The name of the database.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The database type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Output only. The last modification time of the database. A timestamp in
        /// RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
        /// fractional digits. Examples: "2014-10-02T15:01:23Z" and
        /// "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public DatabaseState()
        {
        }
        public static new DatabaseState Empty => new DatabaseState();
    }
}
