// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dataplex;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.dataplex.LakeArgs;
import com.pulumi.gcp.dataplex.inputs.LakeState;
import com.pulumi.gcp.dataplex.outputs.LakeAssetStatus;
import com.pulumi.gcp.dataplex.outputs.LakeMetastore;
import com.pulumi.gcp.dataplex.outputs.LakeMetastoreStatus;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The Dataplex Lake resource
 * 
 * ## Example Usage
 * ### Basic_lake
 * A basic example of a dataplex lake
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.dataplex.Lake;
 * import com.pulumi.gcp.dataplex.LakeArgs;
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
 *         var primary = new Lake(&#34;primary&#34;, LakeArgs.builder()        
 *             .description(&#34;Lake for DCL&#34;)
 *             .displayName(&#34;Lake for DCL&#34;)
 *             .labels(Map.of(&#34;my-lake&#34;, &#34;exists&#34;))
 *             .location(&#34;us-west1&#34;)
 *             .project(&#34;my-project-name&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Lake can be imported using any of these accepted formats* `projects/{{project}}/locations/{{location}}/lakes/{{name}}` * `{{project}}/{{location}}/{{name}}` * `{{location}}/{{name}}` When using the `pulumi import` command, Lake can be imported using one of the formats above. For example
 * 
 * ```sh
 *  $ pulumi import gcp:dataplex/lake:Lake default projects/{{project}}/locations/{{location}}/lakes/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:dataplex/lake:Lake default {{project}}/{{location}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:dataplex/lake:Lake default {{location}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:dataplex/lake:Lake")
public class Lake extends com.pulumi.resources.CustomResource {
    /**
     * Output only. Aggregated status of the underlying assets of the lake.
     * 
     */
    @Export(name="assetStatuses", refs={List.class,LakeAssetStatus.class}, tree="[0,1]")
    private Output<List<LakeAssetStatus>> assetStatuses;

    /**
     * @return Output only. Aggregated status of the underlying assets of the lake.
     * 
     */
    public Output<List<LakeAssetStatus>> assetStatuses() {
        return this.assetStatuses;
    }
    /**
     * Output only. The time when the lake was created.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return Output only. The time when the lake was created.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * Optional. Description of the lake.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Optional. Description of the lake.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
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
     * Optional. User-defined labels for the lake.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return Optional. User-defined labels for the lake.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
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
     * Optional. Settings to manage lake and Dataproc Metastore service instance association.
     * 
     */
    @Export(name="metastore", refs={LakeMetastore.class}, tree="[0]")
    private Output</* @Nullable */ LakeMetastore> metastore;

    /**
     * @return Optional. Settings to manage lake and Dataproc Metastore service instance association.
     * 
     */
    public Output<Optional<LakeMetastore>> metastore() {
        return Codegen.optional(this.metastore);
    }
    /**
     * Output only. Metastore status of the lake.
     * 
     */
    @Export(name="metastoreStatuses", refs={List.class,LakeMetastoreStatus.class}, tree="[0,1]")
    private Output<List<LakeMetastoreStatus>> metastoreStatuses;

    /**
     * @return Output only. Metastore status of the lake.
     * 
     */
    public Output<List<LakeMetastoreStatus>> metastoreStatuses() {
        return this.metastoreStatuses;
    }
    /**
     * The name of the lake.
     * 
     * ***
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the lake.
     * 
     * ***
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
     * Output only. Service account associated with this lake. This service account must be authorized to access or operate on resources managed by the lake.
     * 
     */
    @Export(name="serviceAccount", refs={String.class}, tree="[0]")
    private Output<String> serviceAccount;

    /**
     * @return Output only. Service account associated with this lake. This service account must be authorized to access or operate on resources managed by the lake.
     * 
     */
    public Output<String> serviceAccount() {
        return this.serviceAccount;
    }
    /**
     * Output only. Current state of the lake. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, DELETING, ACTION_REQUIRED
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return Output only. Current state of the lake. Possible values: STATE_UNSPECIFIED, ACTIVE, CREATING, DELETING, ACTION_REQUIRED
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * Output only. System generated globally unique ID for the lake. This ID will be different if the lake is deleted and re-created with the same name.
     * 
     */
    @Export(name="uid", refs={String.class}, tree="[0]")
    private Output<String> uid;

    /**
     * @return Output only. System generated globally unique ID for the lake. This ID will be different if the lake is deleted and re-created with the same name.
     * 
     */
    public Output<String> uid() {
        return this.uid;
    }
    /**
     * Output only. The time when the lake was last updated.
     * 
     */
    @Export(name="updateTime", refs={String.class}, tree="[0]")
    private Output<String> updateTime;

    /**
     * @return Output only. The time when the lake was last updated.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Lake(String name) {
        this(name, LakeArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Lake(String name, LakeArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Lake(String name, LakeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:dataplex/lake:Lake", name, args == null ? LakeArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Lake(String name, Output<String> id, @Nullable LakeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:dataplex/lake:Lake", name, state, makeResourceOptions(options, id));
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
    public static Lake get(String name, Output<String> id, @Nullable LakeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Lake(name, id, state, options);
    }
}
