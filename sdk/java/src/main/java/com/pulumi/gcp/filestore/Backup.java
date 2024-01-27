// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.filestore;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.filestore.BackupArgs;
import com.pulumi.gcp.filestore.inputs.BackupState;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A Google Cloud Filestore backup.
 * 
 * To get more information about Backup, see:
 * 
 * * [API documentation](https://cloud.google.com/filestore/docs/reference/rest/v1/projects.locations.instances.backups)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/filestore/docs/backups)
 *     * [Creating Backups](https://cloud.google.com/filestore/docs/create-backups)
 * 
 * ## Example Usage
 * ### Filestore Backup Basic
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
 * import com.pulumi.gcp.filestore.Backup;
 * import com.pulumi.gcp.filestore.BackupArgs;
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
 *             .location(&#34;us-central1-b&#34;)
 *             .tier(&#34;BASIC_HDD&#34;)
 *             .fileShares(InstanceFileSharesArgs.builder()
 *                 .capacityGb(1024)
 *                 .name(&#34;share1&#34;)
 *                 .build())
 *             .networks(InstanceNetworkArgs.builder()
 *                 .network(&#34;default&#34;)
 *                 .modes(&#34;MODE_IPV4&#34;)
 *                 .connectMode(&#34;DIRECT_PEERING&#34;)
 *                 .build())
 *             .build());
 * 
 *         var backup = new Backup(&#34;backup&#34;, BackupArgs.builder()        
 *             .location(&#34;us-central1&#34;)
 *             .description(&#34;This is a filestore backup for the test instance&#34;)
 *             .sourceInstance(instance.id())
 *             .sourceFileShare(&#34;share1&#34;)
 *             .labels(Map.ofEntries(
 *                 Map.entry(&#34;files&#34;, &#34;label1&#34;),
 *                 Map.entry(&#34;other-label&#34;, &#34;label2&#34;)
 *             ))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Backup can be imported using any of these accepted formats* `projects/{{project}}/locations/{{location}}/backups/{{name}}` * `{{project}}/{{location}}/{{name}}` * `{{location}}/{{name}}` When using the `pulumi import` command, Backup can be imported using one of the formats above. For example
 * 
 * ```sh
 *  $ pulumi import gcp:filestore/backup:Backup default projects/{{project}}/locations/{{location}}/backups/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:filestore/backup:Backup default {{project}}/{{location}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:filestore/backup:Backup default {{location}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:filestore/backup:Backup")
public class Backup extends com.pulumi.resources.CustomResource {
    /**
     * The amount of bytes needed to allocate a full copy of the snapshot content.
     * 
     */
    @Export(name="capacityGb", refs={String.class}, tree="[0]")
    private Output<String> capacityGb;

    /**
     * @return The amount of bytes needed to allocate a full copy of the snapshot content.
     * 
     */
    public Output<String> capacityGb() {
        return this.capacityGb;
    }
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
     * A description of the backup with 2048 characters or less. Requests with longer descriptions will be rejected.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A description of the backup with 2048 characters or less. Requests with longer descriptions will be rejected.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Amount of bytes that will be downloaded if the backup is restored.
     * 
     */
    @Export(name="downloadBytes", refs={String.class}, tree="[0]")
    private Output<String> downloadBytes;

    /**
     * @return Amount of bytes that will be downloaded if the backup is restored.
     * 
     */
    public Output<String> downloadBytes() {
        return this.downloadBytes;
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
     * KMS key name used for data encryption.
     * 
     */
    @Export(name="kmsKeyName", refs={String.class}, tree="[0]")
    private Output<String> kmsKeyName;

    /**
     * @return KMS key name used for data encryption.
     * 
     */
    public Output<String> kmsKeyName() {
        return this.kmsKeyName;
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
     * ***
     * 
     */
    @Export(name="location", refs={String.class}, tree="[0]")
    private Output<String> location;

    /**
     * @return The name of the location of the instance. This can be a region for ENTERPRISE tier instances.
     * 
     * ***
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    /**
     * The resource name of the backup. The name must be unique within the specified instance.
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
     * @return The resource name of the backup. The name must be unique within the specified instance.
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
     * Name of the file share in the source Cloud Filestore instance that the backup is created from.
     * 
     */
    @Export(name="sourceFileShare", refs={String.class}, tree="[0]")
    private Output<String> sourceFileShare;

    /**
     * @return Name of the file share in the source Cloud Filestore instance that the backup is created from.
     * 
     */
    public Output<String> sourceFileShare() {
        return this.sourceFileShare;
    }
    /**
     * The resource name of the source Cloud Filestore instance, in the format projects/{projectId}/locations/{locationId}/instances/{instanceId}, used to create this backup.
     * 
     */
    @Export(name="sourceInstance", refs={String.class}, tree="[0]")
    private Output<String> sourceInstance;

    /**
     * @return The resource name of the source Cloud Filestore instance, in the format projects/{projectId}/locations/{locationId}/instances/{instanceId}, used to create this backup.
     * 
     */
    public Output<String> sourceInstance() {
        return this.sourceInstance;
    }
    /**
     * The service tier of the source Cloud Filestore instance that this backup is created from.
     * 
     */
    @Export(name="sourceInstanceTier", refs={String.class}, tree="[0]")
    private Output<String> sourceInstanceTier;

    /**
     * @return The service tier of the source Cloud Filestore instance that this backup is created from.
     * 
     */
    public Output<String> sourceInstanceTier() {
        return this.sourceInstanceTier;
    }
    /**
     * The backup state.
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return The backup state.
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * The size of the storage used by the backup. As backups share storage, this number is expected to change with backup creation/deletion.
     * 
     */
    @Export(name="storageBytes", refs={String.class}, tree="[0]")
    private Output<String> storageBytes;

    /**
     * @return The size of the storage used by the backup. As backups share storage, this number is expected to change with backup creation/deletion.
     * 
     */
    public Output<String> storageBytes() {
        return this.storageBytes;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Backup(String name) {
        this(name, BackupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Backup(String name, BackupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Backup(String name, BackupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:filestore/backup:Backup", name, args == null ? BackupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Backup(String name, Output<String> id, @Nullable BackupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:filestore/backup:Backup", name, state, makeResourceOptions(options, id));
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
    public static Backup get(String name, Output<String> id, @Nullable BackupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Backup(name, id, state, options);
    }
}
