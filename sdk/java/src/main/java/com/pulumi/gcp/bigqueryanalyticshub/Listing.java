// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.bigqueryanalyticshub;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.bigqueryanalyticshub.ListingArgs;
import com.pulumi.gcp.bigqueryanalyticshub.inputs.ListingState;
import com.pulumi.gcp.bigqueryanalyticshub.outputs.ListingBigqueryDataset;
import com.pulumi.gcp.bigqueryanalyticshub.outputs.ListingDataProvider;
import com.pulumi.gcp.bigqueryanalyticshub.outputs.ListingPublisher;
import com.pulumi.gcp.bigqueryanalyticshub.outputs.ListingRestrictedExportConfig;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A Bigquery Analytics Hub data exchange listing
 * 
 * To get more information about Listing, see:
 * 
 * * [API documentation](https://cloud.google.com/bigquery/docs/reference/analytics-hub/rest/v1/projects.locations.dataExchanges.listings)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/bigquery/docs/analytics-hub-introduction)
 * 
 * ## Example Usage
 * 
 * ### Bigquery Analyticshub Listing Basic
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.bigqueryanalyticshub.DataExchange;
 * import com.pulumi.gcp.bigqueryanalyticshub.DataExchangeArgs;
 * import com.pulumi.gcp.bigquery.Dataset;
 * import com.pulumi.gcp.bigquery.DatasetArgs;
 * import com.pulumi.gcp.bigqueryanalyticshub.Listing;
 * import com.pulumi.gcp.bigqueryanalyticshub.ListingArgs;
 * import com.pulumi.gcp.bigqueryanalyticshub.inputs.ListingBigqueryDatasetArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var listing = new DataExchange("listing", DataExchangeArgs.builder()
 *             .location("US")
 *             .dataExchangeId("my_data_exchange")
 *             .displayName("my_data_exchange")
 *             .description("example data exchange")
 *             .build());
 * 
 *         var listingDataset = new Dataset("listingDataset", DatasetArgs.builder()
 *             .datasetId("my_listing")
 *             .friendlyName("my_listing")
 *             .description("example data exchange")
 *             .location("US")
 *             .build());
 * 
 *         var listingListing = new Listing("listingListing", ListingArgs.builder()
 *             .location("US")
 *             .dataExchangeId(listing.dataExchangeId())
 *             .listingId("my_listing")
 *             .displayName("my_listing")
 *             .description("example data exchange")
 *             .bigqueryDataset(ListingBigqueryDatasetArgs.builder()
 *                 .dataset(listingDataset.id())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Bigquery Analyticshub Listing Restricted
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.bigqueryanalyticshub.DataExchange;
 * import com.pulumi.gcp.bigqueryanalyticshub.DataExchangeArgs;
 * import com.pulumi.gcp.bigquery.Dataset;
 * import com.pulumi.gcp.bigquery.DatasetArgs;
 * import com.pulumi.gcp.bigqueryanalyticshub.Listing;
 * import com.pulumi.gcp.bigqueryanalyticshub.ListingArgs;
 * import com.pulumi.gcp.bigqueryanalyticshub.inputs.ListingBigqueryDatasetArgs;
 * import com.pulumi.gcp.bigqueryanalyticshub.inputs.ListingRestrictedExportConfigArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var listing = new DataExchange("listing", DataExchangeArgs.builder()
 *             .location("US")
 *             .dataExchangeId("my_data_exchange")
 *             .displayName("my_data_exchange")
 *             .description("example data exchange")
 *             .build());
 * 
 *         var listingDataset = new Dataset("listingDataset", DatasetArgs.builder()
 *             .datasetId("my_listing")
 *             .friendlyName("my_listing")
 *             .description("example data exchange")
 *             .location("US")
 *             .build());
 * 
 *         var listingListing = new Listing("listingListing", ListingArgs.builder()
 *             .location("US")
 *             .dataExchangeId(listing.dataExchangeId())
 *             .listingId("my_listing")
 *             .displayName("my_listing")
 *             .description("example data exchange")
 *             .bigqueryDataset(ListingBigqueryDatasetArgs.builder()
 *                 .dataset(listingDataset.id())
 *                 .build())
 *             .restrictedExportConfig(ListingRestrictedExportConfigArgs.builder()
 *                 .enabled(true)
 *                 .restrictQueryResult(true)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Listing can be imported using any of these accepted formats:
 * 
 * * `projects/{{project}}/locations/{{location}}/dataExchanges/{{data_exchange_id}}/listings/{{listing_id}}`
 * 
 * * `{{project}}/{{location}}/{{data_exchange_id}}/{{listing_id}}`
 * 
 * * `{{location}}/{{data_exchange_id}}/{{listing_id}}`
 * 
 * When using the `pulumi import` command, Listing can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:bigqueryanalyticshub/listing:Listing default projects/{{project}}/locations/{{location}}/dataExchanges/{{data_exchange_id}}/listings/{{listing_id}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:bigqueryanalyticshub/listing:Listing default {{project}}/{{location}}/{{data_exchange_id}}/{{listing_id}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:bigqueryanalyticshub/listing:Listing default {{location}}/{{data_exchange_id}}/{{listing_id}}
 * ```
 * 
 */
@ResourceType(type="gcp:bigqueryanalyticshub/listing:Listing")
public class Listing extends com.pulumi.resources.CustomResource {
    /**
     * Shared dataset i.e. BigQuery dataset source.
     * Structure is documented below.
     * 
     */
    @Export(name="bigqueryDataset", refs={ListingBigqueryDataset.class}, tree="[0]")
    private Output<ListingBigqueryDataset> bigqueryDataset;

    /**
     * @return Shared dataset i.e. BigQuery dataset source.
     * Structure is documented below.
     * 
     */
    public Output<ListingBigqueryDataset> bigqueryDataset() {
        return this.bigqueryDataset;
    }
    /**
     * Categories of the listing. Up to two categories are allowed.
     * 
     */
    @Export(name="categories", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> categories;

    /**
     * @return Categories of the listing. Up to two categories are allowed.
     * 
     */
    public Output<Optional<List<String>>> categories() {
        return Codegen.optional(this.categories);
    }
    /**
     * The ID of the data exchange. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces.
     * 
     */
    @Export(name="dataExchangeId", refs={String.class}, tree="[0]")
    private Output<String> dataExchangeId;

    /**
     * @return The ID of the data exchange. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces.
     * 
     */
    public Output<String> dataExchangeId() {
        return this.dataExchangeId;
    }
    /**
     * Details of the data provider who owns the source data.
     * 
     */
    @Export(name="dataProvider", refs={ListingDataProvider.class}, tree="[0]")
    private Output</* @Nullable */ ListingDataProvider> dataProvider;

    /**
     * @return Details of the data provider who owns the source data.
     * 
     */
    public Output<Optional<ListingDataProvider>> dataProvider() {
        return Codegen.optional(this.dataProvider);
    }
    /**
     * Short description of the listing. The description must not contain Unicode non-characters and C0 and C1 control codes
     * except tabs (HT), new lines (LF), carriage returns (CR), and page breaks (FF).
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Short description of the listing. The description must not contain Unicode non-characters and C0 and C1 control codes
     * except tabs (HT), new lines (LF), carriage returns (CR), and page breaks (FF).
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Human-readable display name of the listing. The display name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), ampersands (&amp;) and can&#39;t start or end with spaces.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output<String> displayName;

    /**
     * @return Human-readable display name of the listing. The display name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), ampersands (&amp;) and can&#39;t start or end with spaces.
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }
    /**
     * Documentation describing the listing.
     * 
     */
    @Export(name="documentation", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> documentation;

    /**
     * @return Documentation describing the listing.
     * 
     */
    public Output<Optional<String>> documentation() {
        return Codegen.optional(this.documentation);
    }
    /**
     * Base64 encoded image representing the listing.
     * 
     */
    @Export(name="icon", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> icon;

    /**
     * @return Base64 encoded image representing the listing.
     * 
     */
    public Output<Optional<String>> icon() {
        return Codegen.optional(this.icon);
    }
    /**
     * The ID of the listing. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces.
     * 
     */
    @Export(name="listingId", refs={String.class}, tree="[0]")
    private Output<String> listingId;

    /**
     * @return The ID of the listing. Must contain only Unicode letters, numbers (0-9), underscores (_). Should not use characters that require URL-escaping, or characters outside of ASCII, spaces.
     * 
     */
    public Output<String> listingId() {
        return this.listingId;
    }
    /**
     * The name of the location this data exchange listing.
     * 
     */
    @Export(name="location", refs={String.class}, tree="[0]")
    private Output<String> location;

    /**
     * @return The name of the location this data exchange listing.
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    /**
     * The resource name of the listing. e.g. &#34;projects/myproject/locations/US/dataExchanges/123/listings/456&#34;
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The resource name of the listing. e.g. &#34;projects/myproject/locations/US/dataExchanges/123/listings/456&#34;
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Email or URL of the primary point of contact of the listing.
     * 
     */
    @Export(name="primaryContact", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> primaryContact;

    /**
     * @return Email or URL of the primary point of contact of the listing.
     * 
     */
    public Output<Optional<String>> primaryContact() {
        return Codegen.optional(this.primaryContact);
    }
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    public Output<String> project() {
        return this.project;
    }
    /**
     * Details of the publisher who owns the listing and who can share the source data.
     * 
     */
    @Export(name="publisher", refs={ListingPublisher.class}, tree="[0]")
    private Output</* @Nullable */ ListingPublisher> publisher;

    /**
     * @return Details of the publisher who owns the listing and who can share the source data.
     * 
     */
    public Output<Optional<ListingPublisher>> publisher() {
        return Codegen.optional(this.publisher);
    }
    /**
     * Email or URL of the request access of the listing. Subscribers can use this reference to request access.
     * 
     */
    @Export(name="requestAccess", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> requestAccess;

    /**
     * @return Email or URL of the request access of the listing. Subscribers can use this reference to request access.
     * 
     */
    public Output<Optional<String>> requestAccess() {
        return Codegen.optional(this.requestAccess);
    }
    /**
     * If set, restricted export configuration will be propagated and enforced on the linked dataset.
     * 
     */
    @Export(name="restrictedExportConfig", refs={ListingRestrictedExportConfig.class}, tree="[0]")
    private Output</* @Nullable */ ListingRestrictedExportConfig> restrictedExportConfig;

    /**
     * @return If set, restricted export configuration will be propagated and enforced on the linked dataset.
     * 
     */
    public Output<Optional<ListingRestrictedExportConfig>> restrictedExportConfig() {
        return Codegen.optional(this.restrictedExportConfig);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Listing(String name) {
        this(name, ListingArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Listing(String name, ListingArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Listing(String name, ListingArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:bigqueryanalyticshub/listing:Listing", name, args == null ? ListingArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Listing(String name, Output<String> id, @Nullable ListingState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:bigqueryanalyticshub/listing:Listing", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Listing get(String name, Output<String> id, @Nullable ListingState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Listing(name, id, state, options);
    }
}
