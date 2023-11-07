// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.filestore;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.filestore.SnapshotArgs;
import com.pulumi.gcp.filestore.inputs.SnapshotState;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A Google Cloud Filestore snapshot.
 * 
 * To get more information about Snapshot, see:
 * 
 * * [API documentation](https://cloud.google.com/filestore/docs/reference/rest/v1/projects.locations.instances.snapshots)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/filestore/docs/snapshots)
 *     * [Creating Snapshots](https://cloud.google.com/filestore/docs/create-snapshots)
 * 
 * ## Example Usage
 * ### Filestore Snapshot Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.filestore.Instance;
 * import com.pulumi.gcp.filestore.InstanceArgs;
 * import com.pulumi.gcp.filestore.inputs.InstanceFileSharesArgs;
 * import com.pulumi.gcp.filestore.inputs.InstanceNetworkArgs;
 * import com.pulumi.gcp.filestore.Snapshot;
 * import com.pulumi.gcp.filestore.SnapshotArgs;
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
 *         var instance = new Instance(&#34;instance&#34;, InstanceArgs.builder()        
 *             .location(&#34;us-east1&#34;)
 *             .tier(&#34;ENTERPRISE&#34;)
 *             .fileShares(InstanceFileSharesArgs.builder()
 *                 .capacityGb(1024)
 *                 .name(&#34;share1&#34;)
 *                 .build())
 *             .networks(InstanceNetworkArgs.builder()
 *                 .network(&#34;default&#34;)
 *                 .modes(&#34;MODE_IPV4&#34;)
 *                 .build())
 *             .build());
 * 
 *         var snapshot = new Snapshot(&#34;snapshot&#34;, SnapshotArgs.builder()        
 *             .instance(instance.name())
 *             .location(&#34;us-east1&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Filestore Snapshot Full
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.filestore.Instance;
 * import com.pulumi.gcp.filestore.InstanceArgs;
 * import com.pulumi.gcp.filestore.inputs.InstanceFileSharesArgs;
 * import com.pulumi.gcp.filestore.inputs.InstanceNetworkArgs;
 * import com.pulumi.gcp.filestore.Snapshot;
 * import com.pulumi.gcp.filestore.SnapshotArgs;
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
 *         var instance = new Instance(&#34;instance&#34;, InstanceArgs.builder()        
 *             .location(&#34;us-west1&#34;)
 *             .tier(&#34;ENTERPRISE&#34;)
 *             .fileShares(InstanceFileSharesArgs.builder()
 *                 .capacityGb(1024)
 *                 .name(&#34;share1&#34;)
 *                 .build())
 *             .networks(InstanceNetworkArgs.builder()
 *                 .network(&#34;default&#34;)
 *                 .modes(&#34;MODE_IPV4&#34;)
 *                 .build())
 *             .build());
 * 
 *         var snapshot = new Snapshot(&#34;snapshot&#34;, SnapshotArgs.builder()        
 *             .instance(instance.name())
 *             .location(&#34;us-west1&#34;)
 *             .description(&#34;Snapshot of test-instance-for-snapshot&#34;)
 *             .labels(Map.of(&#34;my_label&#34;, &#34;value&#34;))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Snapshot can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:filestore/snapshot:Snapshot default projects/{{project}}/locations/{{location}}/instances/{{instance}}/snapshots/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:filestore/snapshot:Snapshot default {{project}}/{{location}}/{{instance}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:filestore/snapshot:Snapshot default {{location}}/{{instance}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:filestore/snapshot:Snapshot")
public class Snapshot extends com.pulumi.resources.CustomResource {
    /**
     * The time when the snapshot was created in RFC3339 text format.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return The time when the snapshot was created in RFC3339 text format.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * A description of the snapshot with 2048 characters or less. Requests with longer descriptions will be rejected.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A description of the snapshot with 2048 characters or less. Requests with longer descriptions will be rejected.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
     * clients and services.
     * 
     */
    @Export(name="effectiveLabels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> effectiveLabels;

    /**
     * @return All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
     * clients and services.
     * 
     */
    public Output<Map<String,String>> effectiveLabels() {
        return this.effectiveLabels;
    }
    /**
     * The amount of bytes needed to allocate a full copy of the snapshot content.
     * 
     */
    @Export(name="filesystemUsedBytes", refs={String.class}, tree="[0]")
    private Output<String> filesystemUsedBytes;

    /**
     * @return The amount of bytes needed to allocate a full copy of the snapshot content.
     * 
     */
    public Output<String> filesystemUsedBytes() {
        return this.filesystemUsedBytes;
    }
    /**
     * The resource name of the filestore instance.
     * 
     * ***
     * 
     */
    @Export(name="instance", refs={String.class}, tree="[0]")
    private Output<String> instance;

    /**
     * @return The resource name of the filestore instance.
     * 
     * ***
     * 
     */
    public Output<String> instance() {
        return this.instance;
    }
    /**
     * Resource labels to represent user-provided metadata.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return Resource labels to represent user-provided metadata.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * The name of the location of the instance. This can be a region for ENTERPRISE tier instances.
     * 
     */
    @Export(name="location", refs={String.class}, tree="[0]")
    private Output<String> location;

    /**
     * @return The name of the location of the instance. This can be a region for ENTERPRISE tier instances.
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    /**
     * The resource name of the snapshot. The name must be unique within the specified instance.
     * The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The resource name of the snapshot. The name must be unique within the specified instance.
     * The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
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
     * The snapshot state.
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return The snapshot state.
     * 
     */
    public Output<String> state() {
        return this.state;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Snapshot(String name) {
        this(name, SnapshotArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Snapshot(String name, SnapshotArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Snapshot(String name, SnapshotArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:filestore/snapshot:Snapshot", name, args == null ? SnapshotArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Snapshot(String name, Output<String> id, @Nullable SnapshotState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:filestore/snapshot:Snapshot", name, state, makeResourceOptions(options, id));
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
    public static Snapshot get(String name, Output<String> id, @Nullable SnapshotState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Snapshot(name, id, state, options);
    }
}
