// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.BigQueryAnalyticsHub
{
    /// <summary>
    /// A Bigquery Analytics Hub data exchange listing
    /// 
    /// To get more information about Listing, see:
    /// 
    /// * [API documentation](https://cloud.google.com/bigquery/docs/reference/analytics-hub/rest/v1/projects.locations.dataExchanges.listings)
    /// * How-to Guides
    ///     * [Official Documentation](https://cloud.google.com/bigquery/docs/analytics-hub-introduction)
    /// 
    /// ## Example Usage
    /// ### Bigquery Analyticshub Listing Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var listingDataExchange = new Gcp.BigQueryAnalyticsHub.DataExchange("listingDataExchange", new()
    ///     {
    ///         Location = "US",
    ///         DataExchangeId = "my_data_exchange",
    ///         DisplayName = "my_data_exchange",
    ///         Description = "example data exchange",
    ///     });
    /// 
    ///     var listingDataset = new Gcp.BigQuery.Dataset("listingDataset", new()
    ///     {
    ///         DatasetId = "my_listing",
    ///         FriendlyName = "my_listing",
    ///         Description = "example data exchange",
    ///         Location = "US",
    ///     });
    /// 
    ///     var listingListing = new Gcp.BigQueryAnalyticsHub.Listing("listingListing", new()
    ///     {
    ///         Location = "US",
    ///         DataExchangeId = listingDataExchange.DataExchangeId,
    ///         ListingId = "my_listing",
    ///         DisplayName = "my_listing",
    ///         Description = "example data exchange",
    ///         BigqueryDataset = new Gcp.BigQueryAnalyticsHub.Inputs.ListingBigqueryDatasetArgs
    ///         {
    ///             Dataset = listingDataset.Id,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Listing can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:bigqueryanalyticshub/listing:Listing default projects/{{project}}/locations/{{location}}/dataExchanges/{{data_exchange_id}}/listings/{{listing_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:bigqueryanalyticshub/listing:Listing default {{project}}/{{location}}/{{data_exchange_id}}/{{listing_id}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:bigqueryanalyticshub/listing:Listing default {{location}}/{{data_exchange_id}}/{{listing_id}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:bigqueryanalyticshub/listing:Listing")]
    public partial class Listing : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Shared dataset i.e. BigQuery dataset source.
        /// Structure is documented below.
        /// </summary>
        [Output("bigqueryDataset")]
        public Output<Outputs.ListingBigqueryDataset> BigqueryDataset { get; private set; } = null!;

        /// <summary>
        /// Categories of the listing. Up to two categories are allowed.
        /// </summary>
        [Output("categories")]
        public Output<ImmutableArray<string>> Categories { get; private set; } = null!;

        /// <summary>
        /// The ID of the data exchange. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces.
        /// </summary>
        [Output("dataExchangeId")]
        public Output<string> DataExchangeId { get; private set; } = null!;

        /// <summary>
        /// Details of the data provider who owns the source data.
        /// Structure is documented below.
        /// </summary>
        [Output("dataProvider")]
        public Output<Outputs.ListingDataProvider?> DataProvider { get; private set; } = null!;

        /// <summary>
        /// Short description of the listing. The description must not contain Unicode non-characters and C0 and C1 control codes except tabs (HT), new lines (LF), carriage returns (CR), and page breaks (FF).
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Human-readable display name of the listing. The display name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), ampersands (&amp;) and can't start or end with spaces.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Documentation describing the listing.
        /// </summary>
        [Output("documentation")]
        public Output<string?> Documentation { get; private set; } = null!;

        /// <summary>
        /// Base64 encoded image representing the listing.
        /// </summary>
        [Output("icon")]
        public Output<string?> Icon { get; private set; } = null!;

        /// <summary>
        /// The ID of the listing. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces.
        /// </summary>
        [Output("listingId")]
        public Output<string> ListingId { get; private set; } = null!;

        /// <summary>
        /// The name of the location this data exchange listing.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Name of the data provider.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Email or URL of the primary point of contact of the listing.
        /// </summary>
        [Output("primaryContact")]
        public Output<string?> PrimaryContact { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Details of the publisher who owns the listing and who can share the source data.
        /// Structure is documented below.
        /// </summary>
        [Output("publisher")]
        public Output<Outputs.ListingPublisher?> Publisher { get; private set; } = null!;

        /// <summary>
        /// Email or URL of the request access of the listing. Subscribers can use this reference to request access.
        /// </summary>
        [Output("requestAccess")]
        public Output<string?> RequestAccess { get; private set; } = null!;


        /// <summary>
        /// Create a Listing resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Listing(string name, ListingArgs args, CustomResourceOptions? options = null)
            : base("gcp:bigqueryanalyticshub/listing:Listing", name, args ?? new ListingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Listing(string name, Input<string> id, ListingState? state = null, CustomResourceOptions? options = null)
            : base("gcp:bigqueryanalyticshub/listing:Listing", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Listing resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Listing Get(string name, Input<string> id, ListingState? state = null, CustomResourceOptions? options = null)
        {
            return new Listing(name, id, state, options);
        }
    }

    public sealed class ListingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Shared dataset i.e. BigQuery dataset source.
        /// Structure is documented below.
        /// </summary>
        [Input("bigqueryDataset", required: true)]
        public Input<Inputs.ListingBigqueryDatasetArgs> BigqueryDataset { get; set; } = null!;

        [Input("categories")]
        private InputList<string>? _categories;

        /// <summary>
        /// Categories of the listing. Up to two categories are allowed.
        /// </summary>
        public InputList<string> Categories
        {
            get => _categories ?? (_categories = new InputList<string>());
            set => _categories = value;
        }

        /// <summary>
        /// The ID of the data exchange. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces.
        /// </summary>
        [Input("dataExchangeId", required: true)]
        public Input<string> DataExchangeId { get; set; } = null!;

        /// <summary>
        /// Details of the data provider who owns the source data.
        /// Structure is documented below.
        /// </summary>
        [Input("dataProvider")]
        public Input<Inputs.ListingDataProviderArgs>? DataProvider { get; set; }

        /// <summary>
        /// Short description of the listing. The description must not contain Unicode non-characters and C0 and C1 control codes except tabs (HT), new lines (LF), carriage returns (CR), and page breaks (FF).
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Human-readable display name of the listing. The display name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), ampersands (&amp;) and can't start or end with spaces.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// Documentation describing the listing.
        /// </summary>
        [Input("documentation")]
        public Input<string>? Documentation { get; set; }

        /// <summary>
        /// Base64 encoded image representing the listing.
        /// </summary>
        [Input("icon")]
        public Input<string>? Icon { get; set; }

        /// <summary>
        /// The ID of the listing. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces.
        /// </summary>
        [Input("listingId", required: true)]
        public Input<string> ListingId { get; set; } = null!;

        /// <summary>
        /// The name of the location this data exchange listing.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Email or URL of the primary point of contact of the listing.
        /// </summary>
        [Input("primaryContact")]
        public Input<string>? PrimaryContact { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Details of the publisher who owns the listing and who can share the source data.
        /// Structure is documented below.
        /// </summary>
        [Input("publisher")]
        public Input<Inputs.ListingPublisherArgs>? Publisher { get; set; }

        /// <summary>
        /// Email or URL of the request access of the listing. Subscribers can use this reference to request access.
        /// </summary>
        [Input("requestAccess")]
        public Input<string>? RequestAccess { get; set; }

        public ListingArgs()
        {
        }
        public static new ListingArgs Empty => new ListingArgs();
    }

    public sealed class ListingState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Shared dataset i.e. BigQuery dataset source.
        /// Structure is documented below.
        /// </summary>
        [Input("bigqueryDataset")]
        public Input<Inputs.ListingBigqueryDatasetGetArgs>? BigqueryDataset { get; set; }

        [Input("categories")]
        private InputList<string>? _categories;

        /// <summary>
        /// Categories of the listing. Up to two categories are allowed.
        /// </summary>
        public InputList<string> Categories
        {
            get => _categories ?? (_categories = new InputList<string>());
            set => _categories = value;
        }

        /// <summary>
        /// The ID of the data exchange. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces.
        /// </summary>
        [Input("dataExchangeId")]
        public Input<string>? DataExchangeId { get; set; }

        /// <summary>
        /// Details of the data provider who owns the source data.
        /// Structure is documented below.
        /// </summary>
        [Input("dataProvider")]
        public Input<Inputs.ListingDataProviderGetArgs>? DataProvider { get; set; }

        /// <summary>
        /// Short description of the listing. The description must not contain Unicode non-characters and C0 and C1 control codes except tabs (HT), new lines (LF), carriage returns (CR), and page breaks (FF).
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Human-readable display name of the listing. The display name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), ampersands (&amp;) and can't start or end with spaces.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Documentation describing the listing.
        /// </summary>
        [Input("documentation")]
        public Input<string>? Documentation { get; set; }

        /// <summary>
        /// Base64 encoded image representing the listing.
        /// </summary>
        [Input("icon")]
        public Input<string>? Icon { get; set; }

        /// <summary>
        /// The ID of the listing. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces.
        /// </summary>
        [Input("listingId")]
        public Input<string>? ListingId { get; set; }

        /// <summary>
        /// The name of the location this data exchange listing.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Name of the data provider.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Email or URL of the primary point of contact of the listing.
        /// </summary>
        [Input("primaryContact")]
        public Input<string>? PrimaryContact { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Details of the publisher who owns the listing and who can share the source data.
        /// Structure is documented below.
        /// </summary>
        [Input("publisher")]
        public Input<Inputs.ListingPublisherGetArgs>? Publisher { get; set; }

        /// <summary>
        /// Email or URL of the request access of the listing. Subscribers can use this reference to request access.
        /// </summary>
        [Input("requestAccess")]
        public Input<string>? RequestAccess { get; set; }

        public ListingState()
        {
        }
        public static new ListingState Empty => new ListingState();
    }
}
