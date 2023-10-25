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
    /// Catalogs are top-level containers for Databases and Tables.
    /// 
    /// To get more information about Catalog, see:
    /// 
    /// * [API documentation](https://cloud.google.com/bigquery/docs/reference/biglake/rest/v1/projects.locations.catalogs)
    /// * How-to Guides
    ///     * [Manage open source metadata with BigLake Metastore](https://cloud.google.com/bigquery/docs/manage-open-source-metadata#create_catalogs)
    /// 
    /// ## Example Usage
    /// ### Bigquery Biglake Catalog
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = new Gcp.BigLake.Catalog("default", new()
    ///     {
    ///         Location = "US",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Catalog can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:biglake/catalog:Catalog default projects/{{project}}/locations/{{location}}/catalogs/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:biglake/catalog:Catalog default {{project}}/{{location}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:biglake/catalog:Catalog default {{location}}/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:biglake/catalog:Catalog")]
    public partial class Catalog : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Output only. The creation time of the catalog. A timestamp in RFC3339 UTC
        /// "Zulu" format, with nanosecond resolution and up to nine fractional
        /// digits.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Output only. The deletion time of the catalog. Only set after the catalog
        /// is deleted. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
        /// resolution and up to nine fractional digits.
        /// </summary>
        [Output("deleteTime")]
        public Output<string> DeleteTime { get; private set; } = null!;

        /// <summary>
        /// Output only. The time when this catalog is considered expired. Only set
        /// after the catalog is deleted. Only set after the catalog is deleted.
        /// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and
        /// up to nine fractional digits.
        /// </summary>
        [Output("expireTime")]
        public Output<string> ExpireTime { get; private set; } = null!;

        /// <summary>
        /// The geographic location where the Catalog should reside.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the Catalog. Format:
        /// projects/{project_id_or_number}/locations/{locationId}/catalogs/{catalogId}
        /// 
        /// 
        /// - - -
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
        /// Output only. The last modification time of the catalog. A timestamp in
        /// RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
        /// fractional digits.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a Catalog resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Catalog(string name, CatalogArgs args, CustomResourceOptions? options = null)
            : base("gcp:biglake/catalog:Catalog", name, args ?? new CatalogArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Catalog(string name, Input<string> id, CatalogState? state = null, CustomResourceOptions? options = null)
            : base("gcp:biglake/catalog:Catalog", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Catalog resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Catalog Get(string name, Input<string> id, CatalogState? state = null, CustomResourceOptions? options = null)
        {
            return new Catalog(name, id, state, options);
        }
    }

    public sealed class CatalogArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The geographic location where the Catalog should reside.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The name of the Catalog. Format:
        /// projects/{project_id_or_number}/locations/{locationId}/catalogs/{catalogId}
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public CatalogArgs()
        {
        }
        public static new CatalogArgs Empty => new CatalogArgs();
    }

    public sealed class CatalogState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Output only. The creation time of the catalog. A timestamp in RFC3339 UTC
        /// "Zulu" format, with nanosecond resolution and up to nine fractional
        /// digits.
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Output only. The deletion time of the catalog. Only set after the catalog
        /// is deleted. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
        /// resolution and up to nine fractional digits.
        /// </summary>
        [Input("deleteTime")]
        public Input<string>? DeleteTime { get; set; }

        /// <summary>
        /// Output only. The time when this catalog is considered expired. Only set
        /// after the catalog is deleted. Only set after the catalog is deleted.
        /// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and
        /// up to nine fractional digits.
        /// </summary>
        [Input("expireTime")]
        public Input<string>? ExpireTime { get; set; }

        /// <summary>
        /// The geographic location where the Catalog should reside.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the Catalog. Format:
        /// projects/{project_id_or_number}/locations/{locationId}/catalogs/{catalogId}
        /// 
        /// 
        /// - - -
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
        /// Output only. The last modification time of the catalog. A timestamp in
        /// RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine
        /// fractional digits.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public CatalogState()
        {
        }
        public static new CatalogState Empty => new CatalogState();
    }
}
