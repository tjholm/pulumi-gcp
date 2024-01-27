// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.DataCatalog
{
    /// <summary>
    /// A collection of policy tags that classify data along a common axis.
    /// 
    /// To get more information about Taxonomy, see:
    /// 
    /// * [API documentation](https://cloud.google.com/data-catalog/docs/reference/rest/v1/projects.locations.taxonomies)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/data-catalog/docs)
    /// 
    /// ## Example Usage
    /// ### Data Catalog Taxonomy Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var basicTaxonomy = new Gcp.DataCatalog.Taxonomy("basicTaxonomy", new()
    ///     {
    ///         ActivatedPolicyTypes = new[]
    ///         {
    ///             "FINE_GRAINED_ACCESS_CONTROL",
    ///         },
    ///         Description = "A collection of policy tags",
    ///         DisplayName = "my_taxonomy",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Taxonomy can be imported using any of these accepted formats* `{{name}}` When using the `pulumi import` command, Taxonomy can be imported using one of the formats above. For example
    /// 
    /// ```sh
    ///  $ pulumi import gcp:datacatalog/taxonomy:Taxonomy default {{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:datacatalog/taxonomy:Taxonomy")]
    public partial class Taxonomy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A list of policy types that are activated for this taxonomy. If not set,
        /// defaults to an empty list.
        /// Each value may be one of: `POLICY_TYPE_UNSPECIFIED`, `FINE_GRAINED_ACCESS_CONTROL`.
        /// </summary>
        [Output("activatedPolicyTypes")]
        public Output<ImmutableArray<string>> ActivatedPolicyTypes { get; private set; } = null!;

        /// <summary>
        /// Description of this taxonomy. It must: contain only unicode characters,
        /// tabs, newlines, carriage returns and page breaks; and be at most 2000 bytes
        /// long when encoded in UTF-8. If not set, defaults to an empty description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// User defined name of this taxonomy.
        /// The taxonomy display name must be unique within an organization.
        /// It must: contain only unicode letters, numbers, underscores, dashes
        /// and spaces; not start or end with spaces; and be at most 200 bytes
        /// long when encoded in UTF-8.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Resource name of this taxonomy, whose format is:
        /// "projects/{project}/locations/{region}/taxonomies/{taxonomy}".
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
        /// Taxonomy location region.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;


        /// <summary>
        /// Create a Taxonomy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Taxonomy(string name, TaxonomyArgs args, CustomResourceOptions? options = null)
            : base("gcp:datacatalog/taxonomy:Taxonomy", name, args ?? new TaxonomyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Taxonomy(string name, Input<string> id, TaxonomyState? state = null, CustomResourceOptions? options = null)
            : base("gcp:datacatalog/taxonomy:Taxonomy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Taxonomy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Taxonomy Get(string name, Input<string> id, TaxonomyState? state = null, CustomResourceOptions? options = null)
        {
            return new Taxonomy(name, id, state, options);
        }
    }

    public sealed class TaxonomyArgs : global::Pulumi.ResourceArgs
    {
        [Input("activatedPolicyTypes")]
        private InputList<string>? _activatedPolicyTypes;

        /// <summary>
        /// A list of policy types that are activated for this taxonomy. If not set,
        /// defaults to an empty list.
        /// Each value may be one of: `POLICY_TYPE_UNSPECIFIED`, `FINE_GRAINED_ACCESS_CONTROL`.
        /// </summary>
        public InputList<string> ActivatedPolicyTypes
        {
            get => _activatedPolicyTypes ?? (_activatedPolicyTypes = new InputList<string>());
            set => _activatedPolicyTypes = value;
        }

        /// <summary>
        /// Description of this taxonomy. It must: contain only unicode characters,
        /// tabs, newlines, carriage returns and page breaks; and be at most 2000 bytes
        /// long when encoded in UTF-8. If not set, defaults to an empty description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// User defined name of this taxonomy.
        /// The taxonomy display name must be unique within an organization.
        /// It must: contain only unicode letters, numbers, underscores, dashes
        /// and spaces; not start or end with spaces; and be at most 200 bytes
        /// long when encoded in UTF-8.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Taxonomy location region.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public TaxonomyArgs()
        {
        }
        public static new TaxonomyArgs Empty => new TaxonomyArgs();
    }

    public sealed class TaxonomyState : global::Pulumi.ResourceArgs
    {
        [Input("activatedPolicyTypes")]
        private InputList<string>? _activatedPolicyTypes;

        /// <summary>
        /// A list of policy types that are activated for this taxonomy. If not set,
        /// defaults to an empty list.
        /// Each value may be one of: `POLICY_TYPE_UNSPECIFIED`, `FINE_GRAINED_ACCESS_CONTROL`.
        /// </summary>
        public InputList<string> ActivatedPolicyTypes
        {
            get => _activatedPolicyTypes ?? (_activatedPolicyTypes = new InputList<string>());
            set => _activatedPolicyTypes = value;
        }

        /// <summary>
        /// Description of this taxonomy. It must: contain only unicode characters,
        /// tabs, newlines, carriage returns and page breaks; and be at most 2000 bytes
        /// long when encoded in UTF-8. If not set, defaults to an empty description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// User defined name of this taxonomy.
        /// The taxonomy display name must be unique within an organization.
        /// It must: contain only unicode letters, numbers, underscores, dashes
        /// and spaces; not start or end with spaces; and be at most 200 bytes
        /// long when encoded in UTF-8.
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Resource name of this taxonomy, whose format is:
        /// "projects/{project}/locations/{region}/taxonomies/{taxonomy}".
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
        /// Taxonomy location region.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        public TaxonomyState()
        {
        }
        public static new TaxonomyState Empty => new TaxonomyState();
    }
}
