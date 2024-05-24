// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dataplex;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.dataplex.DatascanArgs;
import com.pulumi.gcp.dataplex.inputs.DatascanState;
import com.pulumi.gcp.dataplex.outputs.DatascanData;
import com.pulumi.gcp.dataplex.outputs.DatascanDataProfileSpec;
import com.pulumi.gcp.dataplex.outputs.DatascanDataQualitySpec;
import com.pulumi.gcp.dataplex.outputs.DatascanExecutionSpec;
import com.pulumi.gcp.dataplex.outputs.DatascanExecutionStatus;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Represents a user-visible job which provides the insights for the related data source.
 * 
 * To get more information about Datascan, see:
 * 
 * * [API documentation](https://cloud.google.com/dataplex/docs/reference/rest)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/dataplex/docs)
 * 
 * ## Example Usage
 * 
 * ### Dataplex Datascan Basic Profile
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.dataplex.Datascan;
 * import com.pulumi.gcp.dataplex.DatascanArgs;
 * import com.pulumi.gcp.dataplex.inputs.DatascanDataArgs;
 * import com.pulumi.gcp.dataplex.inputs.DatascanExecutionSpecArgs;
 * import com.pulumi.gcp.dataplex.inputs.DatascanExecutionSpecTriggerArgs;
 * import com.pulumi.gcp.dataplex.inputs.DatascanExecutionSpecTriggerOnDemandArgs;
 * import com.pulumi.gcp.dataplex.inputs.DatascanDataProfileSpecArgs;
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
 *         var basicProfile = new Datascan("basicProfile", DatascanArgs.builder()
 *             .location("us-central1")
 *             .dataScanId("dataprofile-basic")
 *             .data(DatascanDataArgs.builder()
 *                 .resource("//bigquery.googleapis.com/projects/bigquery-public-data/datasets/samples/tables/shakespeare")
 *                 .build())
 *             .executionSpec(DatascanExecutionSpecArgs.builder()
 *                 .trigger(DatascanExecutionSpecTriggerArgs.builder()
 *                     .onDemand()
 *                     .build())
 *                 .build())
 *             .dataProfileSpec()
 *             .project("my-project-name")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Dataplex Datascan Full Profile
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.dataplex.Datascan;
 * import com.pulumi.gcp.dataplex.DatascanArgs;
 * import com.pulumi.gcp.dataplex.inputs.DatascanDataArgs;
 * import com.pulumi.gcp.dataplex.inputs.DatascanExecutionSpecArgs;
 * import com.pulumi.gcp.dataplex.inputs.DatascanExecutionSpecTriggerArgs;
 * import com.pulumi.gcp.dataplex.inputs.DatascanExecutionSpecTriggerScheduleArgs;
 * import com.pulumi.gcp.dataplex.inputs.DatascanDataProfileSpecArgs;
 * import com.pulumi.gcp.dataplex.inputs.DatascanDataProfileSpecIncludeFieldsArgs;
 * import com.pulumi.gcp.dataplex.inputs.DatascanDataProfileSpecExcludeFieldsArgs;
 * import com.pulumi.gcp.dataplex.inputs.DatascanDataProfileSpecPostScanActionsArgs;
 * import com.pulumi.gcp.dataplex.inputs.DatascanDataProfileSpecPostScanActionsBigqueryExportArgs;
 * import com.pulumi.gcp.bigquery.Dataset;
 * import com.pulumi.gcp.bigquery.DatasetArgs;
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
 *         var fullProfile = new Datascan("fullProfile", DatascanArgs.builder()
 *             .location("us-central1")
 *             .displayName("Full Datascan Profile")
 *             .dataScanId("dataprofile-full")
 *             .description("Example resource - Full Datascan Profile")
 *             .labels(Map.of("author", "billing"))
 *             .data(DatascanDataArgs.builder()
 *                 .resource("//bigquery.googleapis.com/projects/bigquery-public-data/datasets/samples/tables/shakespeare")
 *                 .build())
 *             .executionSpec(DatascanExecutionSpecArgs.builder()
 *                 .trigger(DatascanExecutionSpecTriggerArgs.builder()
 *                     .schedule(DatascanExecutionSpecTriggerScheduleArgs.builder()
 *                         .cron("TZ=America/New_York 1 1 * * *")
 *                         .build())
 *                     .build())
 *                 .build())
 *             .dataProfileSpec(DatascanDataProfileSpecArgs.builder()
 *                 .samplingPercent(80)
 *                 .rowFilter("word_count > 10")
 *                 .includeFields(DatascanDataProfileSpecIncludeFieldsArgs.builder()
 *                     .fieldNames("word_count")
 *                     .build())
 *                 .excludeFields(DatascanDataProfileSpecExcludeFieldsArgs.builder()
 *                     .fieldNames("property_type")
 *                     .build())
 *                 .postScanActions(DatascanDataProfileSpecPostScanActionsArgs.builder()
 *                     .bigqueryExport(DatascanDataProfileSpecPostScanActionsBigqueryExportArgs.builder()
 *                         .resultsTable("//bigquery.googleapis.com/projects/my-project-name/datasets/dataplex_dataset/tables/profile_export")
 *                         .build())
 *                     .build())
 *                 .build())
 *             .project("my-project-name")
 *             .build());
 * 
 *         var source = new Dataset("source", DatasetArgs.builder()
 *             .datasetId("dataplex_dataset")
 *             .friendlyName("test")
 *             .description("This is a test description")
 *             .location("US")
 *             .deleteContentsOnDestroy(true)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Dataplex Datascan Basic Quality
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.dataplex.Datascan;
 * import com.pulumi.gcp.dataplex.DatascanArgs;
 * import com.pulumi.gcp.dataplex.inputs.DatascanDataArgs;
 * import com.pulumi.gcp.dataplex.inputs.DatascanExecutionSpecArgs;
 * import com.pulumi.gcp.dataplex.inputs.DatascanExecutionSpecTriggerArgs;
 * import com.pulumi.gcp.dataplex.inputs.DatascanExecutionSpecTriggerOnDemandArgs;
 * import com.pulumi.gcp.dataplex.inputs.DatascanDataQualitySpecArgs;
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
 *         var basicQuality = new Datascan("basicQuality", DatascanArgs.builder()
 *             .location("us-central1")
 *             .dataScanId("dataquality-basic")
 *             .data(DatascanDataArgs.builder()
 *                 .resource("//bigquery.googleapis.com/projects/bigquery-public-data/datasets/samples/tables/shakespeare")
 *                 .build())
 *             .executionSpec(DatascanExecutionSpecArgs.builder()
 *                 .trigger(DatascanExecutionSpecTriggerArgs.builder()
 *                     .onDemand()
 *                     .build())
 *                 .build())
 *             .dataQualitySpec(DatascanDataQualitySpecArgs.builder()
 *                 .rules(DatascanDataQualitySpecRuleArgs.builder()
 *                     .dimension("VALIDITY")
 *                     .name("rule1")
 *                     .description("rule 1 for validity dimension")
 *                     .tableConditionExpectation(DatascanDataQualitySpecRuleTableConditionExpectationArgs.builder()
 *                         .sqlExpression("COUNT(*) > 0")
 *                         .build())
 *                     .build())
 *                 .build())
 *             .project("my-project-name")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Dataplex Datascan Full Quality
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.dataplex.Datascan;
 * import com.pulumi.gcp.dataplex.DatascanArgs;
 * import com.pulumi.gcp.dataplex.inputs.DatascanDataArgs;
 * import com.pulumi.gcp.dataplex.inputs.DatascanExecutionSpecArgs;
 * import com.pulumi.gcp.dataplex.inputs.DatascanExecutionSpecTriggerArgs;
 * import com.pulumi.gcp.dataplex.inputs.DatascanExecutionSpecTriggerScheduleArgs;
 * import com.pulumi.gcp.dataplex.inputs.DatascanDataQualitySpecArgs;
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
 *         var fullQuality = new Datascan("fullQuality", DatascanArgs.builder()
 *             .location("us-central1")
 *             .displayName("Full Datascan Quality")
 *             .dataScanId("dataquality-full")
 *             .description("Example resource - Full Datascan Quality")
 *             .labels(Map.of("author", "billing"))
 *             .data(DatascanDataArgs.builder()
 *                 .resource("//bigquery.googleapis.com/projects/bigquery-public-data/datasets/austin_bikeshare/tables/bikeshare_stations")
 *                 .build())
 *             .executionSpec(DatascanExecutionSpecArgs.builder()
 *                 .trigger(DatascanExecutionSpecTriggerArgs.builder()
 *                     .schedule(DatascanExecutionSpecTriggerScheduleArgs.builder()
 *                         .cron("TZ=America/New_York 1 1 * * *")
 *                         .build())
 *                     .build())
 *                 .field("modified_date")
 *                 .build())
 *             .dataQualitySpec(DatascanDataQualitySpecArgs.builder()
 *                 .samplingPercent(5)
 *                 .rowFilter("station_id > 1000")
 *                 .rules(                
 *                     DatascanDataQualitySpecRuleArgs.builder()
 *                         .column("address")
 *                         .dimension("VALIDITY")
 *                         .threshold(0.99)
 *                         .nonNullExpectation()
 *                         .build(),
 *                     DatascanDataQualitySpecRuleArgs.builder()
 *                         .column("council_district")
 *                         .dimension("VALIDITY")
 *                         .ignoreNull(true)
 *                         .threshold(0.9)
 *                         .rangeExpectation(DatascanDataQualitySpecRuleRangeExpectationArgs.builder()
 *                             .minValue(1)
 *                             .maxValue(10)
 *                             .strictMinEnabled(true)
 *                             .strictMaxEnabled(false)
 *                             .build())
 *                         .build(),
 *                     DatascanDataQualitySpecRuleArgs.builder()
 *                         .column("power_type")
 *                         .dimension("VALIDITY")
 *                         .ignoreNull(false)
 *                         .regexExpectation(DatascanDataQualitySpecRuleRegexExpectationArgs.builder()
 *                             .regex(".*solar.*")
 *                             .build())
 *                         .build(),
 *                     DatascanDataQualitySpecRuleArgs.builder()
 *                         .column("property_type")
 *                         .dimension("VALIDITY")
 *                         .ignoreNull(false)
 *                         .setExpectation(DatascanDataQualitySpecRuleSetExpectationArgs.builder()
 *                             .values(                            
 *                                 "sidewalk",
 *                                 "parkland")
 *                             .build())
 *                         .build(),
 *                     DatascanDataQualitySpecRuleArgs.builder()
 *                         .column("address")
 *                         .dimension("UNIQUENESS")
 *                         .uniquenessExpectation()
 *                         .build(),
 *                     DatascanDataQualitySpecRuleArgs.builder()
 *                         .column("number_of_docks")
 *                         .dimension("VALIDITY")
 *                         .statisticRangeExpectation(DatascanDataQualitySpecRuleStatisticRangeExpectationArgs.builder()
 *                             .statistic("MEAN")
 *                             .minValue(5)
 *                             .maxValue(15)
 *                             .strictMinEnabled(true)
 *                             .strictMaxEnabled(true)
 *                             .build())
 *                         .build(),
 *                     DatascanDataQualitySpecRuleArgs.builder()
 *                         .column("footprint_length")
 *                         .dimension("VALIDITY")
 *                         .rowConditionExpectation(DatascanDataQualitySpecRuleRowConditionExpectationArgs.builder()
 *                             .sqlExpression("footprint_length > 0 AND footprint_length <= 10")
 *                             .build())
 *                         .build(),
 *                     DatascanDataQualitySpecRuleArgs.builder()
 *                         .dimension("VALIDITY")
 *                         .tableConditionExpectation(DatascanDataQualitySpecRuleTableConditionExpectationArgs.builder()
 *                             .sqlExpression("COUNT(*) > 0")
 *                             .build())
 *                         .build())
 *                 .build())
 *             .project("my-project-name")
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
 * Datascan can be imported using any of these accepted formats:
 * 
 * * `projects/{{project}}/locations/{{location}}/dataScans/{{data_scan_id}}`
 * 
 * * `{{project}}/{{location}}/{{data_scan_id}}`
 * 
 * * `{{location}}/{{data_scan_id}}`
 * 
 * * `{{data_scan_id}}`
 * 
 * When using the `pulumi import` command, Datascan can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:dataplex/datascan:Datascan default projects/{{project}}/locations/{{location}}/dataScans/{{data_scan_id}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:dataplex/datascan:Datascan default {{project}}/{{location}}/{{data_scan_id}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:dataplex/datascan:Datascan default {{location}}/{{data_scan_id}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:dataplex/datascan:Datascan default {{data_scan_id}}
 * ```
 * 
 */
@ResourceType(type="gcp:dataplex/datascan:Datascan")
public class Datascan extends com.pulumi.resources.CustomResource {
    /**
     * The time when the scan was created.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return The time when the scan was created.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * The data source for DataScan.
     * Structure is documented below.
     * 
     */
    @Export(name="data", refs={DatascanData.class}, tree="[0]")
    private Output<DatascanData> data;

    /**
     * @return The data source for DataScan.
     * Structure is documented below.
     * 
     */
    public Output<DatascanData> data() {
        return this.data;
    }
    /**
     * DataProfileScan related setting.
     * 
     */
    @Export(name="dataProfileSpec", refs={DatascanDataProfileSpec.class}, tree="[0]")
    private Output</* @Nullable */ DatascanDataProfileSpec> dataProfileSpec;

    /**
     * @return DataProfileScan related setting.
     * 
     */
    public Output<Optional<DatascanDataProfileSpec>> dataProfileSpec() {
        return Codegen.optional(this.dataProfileSpec);
    }
    /**
     * DataQualityScan related setting.
     * 
     */
    @Export(name="dataQualitySpec", refs={DatascanDataQualitySpec.class}, tree="[0]")
    private Output</* @Nullable */ DatascanDataQualitySpec> dataQualitySpec;

    /**
     * @return DataQualityScan related setting.
     * 
     */
    public Output<Optional<DatascanDataQualitySpec>> dataQualitySpec() {
        return Codegen.optional(this.dataQualitySpec);
    }
    /**
     * DataScan identifier. Must contain only lowercase letters, numbers and hyphens. Must start with a letter. Must end with a number or a letter.
     * 
     */
    @Export(name="dataScanId", refs={String.class}, tree="[0]")
    private Output<String> dataScanId;

    /**
     * @return DataScan identifier. Must contain only lowercase letters, numbers and hyphens. Must start with a letter. Must end with a number or a letter.
     * 
     */
    public Output<String> dataScanId() {
        return this.dataScanId;
    }
    /**
     * Description of the scan.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of the scan.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * User friendly display name.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> displayName;

    /**
     * @return User friendly display name.
     * 
     */
    public Output<Optional<String>> displayName() {
        return Codegen.optional(this.displayName);
    }
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     * 
     */
    @Export(name="effectiveLabels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> effectiveLabels;

    /**
     * @return All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     * 
     */
    public Output<Map<String,String>> effectiveLabels() {
        return this.effectiveLabels;
    }
    /**
     * DataScan execution settings.
     * Structure is documented below.
     * 
     */
    @Export(name="executionSpec", refs={DatascanExecutionSpec.class}, tree="[0]")
    private Output<DatascanExecutionSpec> executionSpec;

    /**
     * @return DataScan execution settings.
     * Structure is documented below.
     * 
     */
    public Output<DatascanExecutionSpec> executionSpec() {
        return this.executionSpec;
    }
    /**
     * Status of the data scan execution.
     * Structure is documented below.
     * 
     */
    @Export(name="executionStatuses", refs={List.class,DatascanExecutionStatus.class}, tree="[0,1]")
    private Output<List<DatascanExecutionStatus>> executionStatuses;

    /**
     * @return Status of the data scan execution.
     * Structure is documented below.
     * 
     */
    public Output<List<DatascanExecutionStatus>> executionStatuses() {
        return this.executionStatuses;
    }
    /**
     * User-defined labels for the scan. A list of key-&gt;value pairs. **Note**: This field is non-authoritative, and will only
     * manage the labels present in your configuration. Please refer to the field &#39;effective_labels&#39; for all of the labels
     * present on the resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return User-defined labels for the scan. A list of key-&gt;value pairs. **Note**: This field is non-authoritative, and will only
     * manage the labels present in your configuration. Please refer to the field &#39;effective_labels&#39; for all of the labels
     * present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * The location where the data scan should reside.
     * 
     */
    @Export(name="location", refs={String.class}, tree="[0]")
    private Output<String> location;

    /**
     * @return The location where the data scan should reside.
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    /**
     * The relative resource name of the scan, of the form: projects/{project}/locations/{locationId}/dataScans/{datascan_id}, where project refers to a project_id or project_number and locationId refers to a GCP region.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The relative resource name of the scan, of the form: projects/{project}/locations/{locationId}/dataScans/{datascan_id}, where project refers to a project_id or project_number and locationId refers to a GCP region.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    public Output<String> project() {
        return this.project;
    }
    /**
     * The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     * 
     */
    @Export(name="pulumiLabels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> pulumiLabels;

    /**
     * @return The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     * 
     */
    public Output<Map<String,String>> pulumiLabels() {
        return this.pulumiLabels;
    }
    /**
     * Current state of the DataScan.
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return Current state of the DataScan.
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * The type of DataScan.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return The type of DataScan.
     * 
     */
    public Output<String> type() {
        return this.type;
    }
    /**
     * System generated globally unique ID for the scan. This ID will be different if the scan is deleted and re-created with the same name.
     * 
     */
    @Export(name="uid", refs={String.class}, tree="[0]")
    private Output<String> uid;

    /**
     * @return System generated globally unique ID for the scan. This ID will be different if the scan is deleted and re-created with the same name.
     * 
     */
    public Output<String> uid() {
        return this.uid;
    }
    /**
     * The time when the scan was last updated.
     * 
     */
    @Export(name="updateTime", refs={String.class}, tree="[0]")
    private Output<String> updateTime;

    /**
     * @return The time when the scan was last updated.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Datascan(String name) {
        this(name, DatascanArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Datascan(String name, DatascanArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Datascan(String name, DatascanArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:dataplex/datascan:Datascan", name, args == null ? DatascanArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Datascan(String name, Output<String> id, @Nullable DatascanState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:dataplex/datascan:Datascan", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "effectiveLabels",
                "pulumiLabels"
            ))
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
    public static Datascan get(String name, Output<String> id, @Nullable DatascanState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Datascan(name, id, state, options);
    }
}
