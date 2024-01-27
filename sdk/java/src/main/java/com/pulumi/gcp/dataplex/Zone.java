// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dataplex;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.dataplex.ZoneArgs;
import com.pulumi.gcp.dataplex.inputs.ZoneState;
import com.pulumi.gcp.dataplex.outputs.ZoneAssetStatus;
import com.pulumi.gcp.dataplex.outputs.ZoneDiscoverySpec;
import com.pulumi.gcp.dataplex.outputs.ZoneResourceSpec;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The Dataplex Zone resource
 * 
 * ## Example Usage
 * ### Basic_zone
 * A basic example of a dataplex zone
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.dataplex.Lake;
 * import com.pulumi.gcp.dataplex.LakeArgs;
 * import com.pulumi.gcp.dataplex.Zone;
 * import com.pulumi.gcp.dataplex.ZoneArgs;
 * import com.pulumi.gcp.dataplex.inputs.ZoneDiscoverySpecArgs;
 * import com.pulumi.gcp.dataplex.inputs.ZoneResourceSpecArgs;
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
 *         var basic = new Lake(&#34;basic&#34;, LakeArgs.builder()        
 *             .location(&#34;us-west1&#34;)
 *             .description(&#34;Lake for DCL&#34;)
 *             .displayName(&#34;Lake for DCL&#34;)
 *             .project(&#34;my-project-name&#34;)
 *             .labels(Map.of(&#34;my-lake&#34;, &#34;exists&#34;))
 *             .build());
 * 
 *         var primary = new Zone(&#34;primary&#34;, ZoneArgs.builder()        
 *             .discoverySpec(ZoneDiscoverySpecArgs.builder()
 *                 .enabled(false)
 *                 .build())
 *             .lake(basic.name())
 *             .location(&#34;us-west1&#34;)
 *             .resourceSpec(ZoneResourceSpecArgs.builder()
 *                 .locationType(&#34;MULTI_REGION&#34;)
 *                 .build())
 *             .type(&#34;RAW&#34;)
 *             .description(&#34;Zone for DCL&#34;)
 *             .displayName(&#34;Zone for DCL&#34;)
 *             .project(&#34;my-project-name&#34;)
 *             .labels()
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Zone can be imported using any of these accepted formats* `projects/{{project}}/locations/{{location}}/lakes/{{lake}}/zones/{{name}}` * `{{project}}/{{location}}/{{lake}}/{{name}}` * `{{location}}/{{lake}}/{{name}}` When using the `pulumi import` command, Zone can be imported using one of the formats above. For example
 * 
 * ```sh
 *  $ pulumi import gcp:dataplex/zone:Zone default projects/{{project}}/locations/{{location}}/lakes/{{lake}}/zones/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:dataplex/zone:Zone default {{project}}/{{location}}/{{lake}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:dataplex/zone:Zone default {{location}}/{{lake}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:dataplex/zone:Zone")
public class Zone extends com.pulumi.resources.CustomResource {
    /**
     * Output only. Aggregated status of the underlying assets of the zone.
     * 
     */
    @Export(name="assetStatuses", refs={List.class,ZoneAssetStatus.class}, tree="[0,1]")
    private Output<List<ZoneAssetStatus>> assetStatuses;

    /**
     * @return Output only. Aggregated status of the underlying assets of the zone.
     * 
     */
    public Output<List<ZoneAssetStatus>> assetStatuses() {
        return this.assetStatuses;
    }
    /**
     * Output only. The time when the zone was created.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return Output only. The time when the zone was created.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * Optional. Description of the zone.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Optional. Description of the zone.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Required. Specification of the discovery feature applied to data in this zone.
     * 
     */
    @Export(name="discoverySpec", refs={ZoneDiscoverySpec.class}, tree="[0]")
    private Output<ZoneDiscoverySpec> discoverySpec;

    /**
     * @return Required. Specification of the discovery feature applied to data in this zone.
     * 
     */
    public Output<ZoneDiscoverySpec> discoverySpec() {
        return this.discoverySpec;
    }
    /**
     * Optional. User friendly display name.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> displayName;

    /**
     * @return Optional. User friendly display name.
     * 
     */
    public Output<Optional<String>> displayName() {
        return Codegen.optional(this.displayName);
    }
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     * 
     */
    @Export(name="effectiveLabels", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output<Map<String,Object>> effectiveLabels;

    /**
     * @return All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     * 
     */
    public Output<Map<String,Object>> effectiveLabels() {
        return this.effectiveLabels;
    }
    /**
     * Optional. User defined labels for the zone.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return Optional. User defined labels for the zone.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * The lake for the resource
     * 
     */
    @Export(name="lake", refs={String.class}, tree="[0]")
    private Output<String> lake;

    /**
     * @return The lake for the resource
     * 
     */
    public Output<String> lake() {
        return this.lake;
    }
    /**
     * The location for the resource
     * 
     */
    @Export(name="location", refs={String.class}, tree="[0]")
    private Output<String> location;

    /**
     * @return The location for the resource
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    /**
     * The name of the zone.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the zone.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The project for the resource
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The project for the resource
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * The combination of labels configured directly on the resource and default labels configured on the provider.
     * 
     */
    @Export(name="pulumiLabels", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output<Map<String,Object>> pulumiLabels;

    /**
     * @return The combination of labels configured directly on the resource and default labels configured on the provider.
     * 
     */
    public Output<Map<String,Object>> pulumiLabels() {
        return this.pulumiLabels;
    }
    /**
     * Required. Immutable. Specification of the resources that are referenced by the assets within this zone.
     * 
     */
    @Export(name="resourceSpec", refs={ZoneResourceSpec.class}, tree="[0]")
    private Output<ZoneResourceSpec> resourceSpec;

    /**
     * @return Required. Immutable. Specification of the resources that are referenced by the assets within this zone.
     * 
     */
    public Output<ZoneResourceSpec> resourceSpec() {
        return this.resourceSpec;
    }
    /**
     * Output only. Current state of the zone. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, DELETING, ACTION_REQUIRED
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return Output only. Current state of the zone. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, DELETING, ACTION_REQUIRED
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * Required. Immutable. The type of the zone. Possible values: TYPE_UNSPECIFIED, RAW, CURATED
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return Required. Immutable. The type of the zone. Possible values: TYPE_UNSPECIFIED, RAW, CURATED
     * 
     */
    public Output<String> type() {
        return this.type;
    }
    /**
     * Output only. System generated globally unique ID for the zone. This ID will be different if the zone is deleted and re-created with the same name.
     * 
     */
    @Export(name="uid", refs={String.class}, tree="[0]")
    private Output<String> uid;

    /**
     * @return Output only. System generated globally unique ID for the zone. This ID will be different if the zone is deleted and re-created with the same name.
     * 
     */
    public Output<String> uid() {
        return this.uid;
    }
    /**
     * Output only. The time when the zone was last updated.
     * 
     */
    @Export(name="updateTime", refs={String.class}, tree="[0]")
    private Output<String> updateTime;

    /**
     * @return Output only. The time when the zone was last updated.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Zone(String name) {
        this(name, ZoneArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Zone(String name, ZoneArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Zone(String name, ZoneArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:dataplex/zone:Zone", name, args == null ? ZoneArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Zone(String name, Output<String> id, @Nullable ZoneState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:dataplex/zone:Zone", name, state, makeResourceOptions(options, id));
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
    public static Zone get(String name, Output<String> id, @Nullable ZoneState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Zone(name, id, state, options);
    }
}
