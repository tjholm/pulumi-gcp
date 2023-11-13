// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.alloydb;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.alloydb.BackupArgs;
import com.pulumi.gcp.alloydb.inputs.BackupState;
import com.pulumi.gcp.alloydb.outputs.BackupEncryptionConfig;
import com.pulumi.gcp.alloydb.outputs.BackupEncryptionInfo;
import com.pulumi.gcp.alloydb.outputs.BackupExpiryQuantity;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * An AlloyDB Backup.
 * 
 * To get more information about Backup, see:
 * 
 * * [API documentation](https://cloud.google.com/alloydb/docs/reference/rest/v1/projects.locations.backups/create)
 * * How-to Guides
 *     * [AlloyDB](https://cloud.google.com/alloydb/docs/)
 * 
 * ## Example Usage
 * ### Alloydb Backup Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.Network;
 * import com.pulumi.gcp.alloydb.Cluster;
 * import com.pulumi.gcp.alloydb.ClusterArgs;
 * import com.pulumi.gcp.compute.GlobalAddress;
 * import com.pulumi.gcp.compute.GlobalAddressArgs;
 * import com.pulumi.gcp.servicenetworking.Connection;
 * import com.pulumi.gcp.servicenetworking.ConnectionArgs;
 * import com.pulumi.gcp.alloydb.Instance;
 * import com.pulumi.gcp.alloydb.InstanceArgs;
 * import com.pulumi.gcp.alloydb.Backup;
 * import com.pulumi.gcp.alloydb.BackupArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         var defaultNetwork = new Network(&#34;defaultNetwork&#34;);
 * 
 *         var defaultCluster = new Cluster(&#34;defaultCluster&#34;, ClusterArgs.builder()        
 *             .clusterId(&#34;alloydb-cluster&#34;)
 *             .location(&#34;us-central1&#34;)
 *             .network(defaultNetwork.id())
 *             .build());
 * 
 *         var privateIpAlloc = new GlobalAddress(&#34;privateIpAlloc&#34;, GlobalAddressArgs.builder()        
 *             .addressType(&#34;INTERNAL&#34;)
 *             .purpose(&#34;VPC_PEERING&#34;)
 *             .prefixLength(16)
 *             .network(defaultNetwork.id())
 *             .build());
 * 
 *         var vpcConnection = new Connection(&#34;vpcConnection&#34;, ConnectionArgs.builder()        
 *             .network(defaultNetwork.id())
 *             .service(&#34;servicenetworking.googleapis.com&#34;)
 *             .reservedPeeringRanges(privateIpAlloc.name())
 *             .build());
 * 
 *         var defaultInstance = new Instance(&#34;defaultInstance&#34;, InstanceArgs.builder()        
 *             .cluster(defaultCluster.name())
 *             .instanceId(&#34;alloydb-instance&#34;)
 *             .instanceType(&#34;PRIMARY&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(vpcConnection)
 *                 .build());
 * 
 *         var defaultBackup = new Backup(&#34;defaultBackup&#34;, BackupArgs.builder()        
 *             .location(&#34;us-central1&#34;)
 *             .backupId(&#34;alloydb-backup&#34;)
 *             .clusterName(defaultCluster.name())
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(defaultInstance)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * ### Alloydb Backup Full
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.Network;
 * import com.pulumi.gcp.alloydb.Cluster;
 * import com.pulumi.gcp.alloydb.ClusterArgs;
 * import com.pulumi.gcp.compute.GlobalAddress;
 * import com.pulumi.gcp.compute.GlobalAddressArgs;
 * import com.pulumi.gcp.servicenetworking.Connection;
 * import com.pulumi.gcp.servicenetworking.ConnectionArgs;
 * import com.pulumi.gcp.alloydb.Instance;
 * import com.pulumi.gcp.alloydb.InstanceArgs;
 * import com.pulumi.gcp.alloydb.Backup;
 * import com.pulumi.gcp.alloydb.BackupArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         var defaultNetwork = new Network(&#34;defaultNetwork&#34;);
 * 
 *         var defaultCluster = new Cluster(&#34;defaultCluster&#34;, ClusterArgs.builder()        
 *             .clusterId(&#34;alloydb-cluster&#34;)
 *             .location(&#34;us-central1&#34;)
 *             .network(defaultNetwork.id())
 *             .build());
 * 
 *         var privateIpAlloc = new GlobalAddress(&#34;privateIpAlloc&#34;, GlobalAddressArgs.builder()        
 *             .addressType(&#34;INTERNAL&#34;)
 *             .purpose(&#34;VPC_PEERING&#34;)
 *             .prefixLength(16)
 *             .network(defaultNetwork.id())
 *             .build());
 * 
 *         var vpcConnection = new Connection(&#34;vpcConnection&#34;, ConnectionArgs.builder()        
 *             .network(defaultNetwork.id())
 *             .service(&#34;servicenetworking.googleapis.com&#34;)
 *             .reservedPeeringRanges(privateIpAlloc.name())
 *             .build());
 * 
 *         var defaultInstance = new Instance(&#34;defaultInstance&#34;, InstanceArgs.builder()        
 *             .cluster(defaultCluster.name())
 *             .instanceId(&#34;alloydb-instance&#34;)
 *             .instanceType(&#34;PRIMARY&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(vpcConnection)
 *                 .build());
 * 
 *         var defaultBackup = new Backup(&#34;defaultBackup&#34;, BackupArgs.builder()        
 *             .location(&#34;us-central1&#34;)
 *             .backupId(&#34;alloydb-backup&#34;)
 *             .clusterName(defaultCluster.name())
 *             .description(&#34;example description&#34;)
 *             .type(&#34;ON_DEMAND&#34;)
 *             .labels(Map.of(&#34;label&#34;, &#34;key&#34;))
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(defaultInstance)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Backup can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:alloydb/backup:Backup default projects/{{project}}/locations/{{location}}/backups/{{backup_id}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:alloydb/backup:Backup default {{project}}/{{location}}/{{backup_id}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:alloydb/backup:Backup default {{location}}/{{backup_id}}
 * ```
 * 
 */
@ResourceType(type="gcp:alloydb/backup:Backup")
public class Backup extends com.pulumi.resources.CustomResource {
    /**
     * Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels. https://google.aip.dev/128
     * An object containing a list of &#34;key&#34;: value pairs. Example: { &#34;name&#34;: &#34;wrench&#34;, &#34;mass&#34;: &#34;1.3kg&#34;, &#34;count&#34;: &#34;3&#34; }.
     * 
     * **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
     * Please refer to the field `effective_annotations` for all of the annotations present on the resource.
     * 
     */
    @Export(name="annotations", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> annotations;

    /**
     * @return Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels. https://google.aip.dev/128
     * An object containing a list of &#34;key&#34;: value pairs. Example: { &#34;name&#34;: &#34;wrench&#34;, &#34;mass&#34;: &#34;1.3kg&#34;, &#34;count&#34;: &#34;3&#34; }.
     * 
     * **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
     * Please refer to the field `effective_annotations` for all of the annotations present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> annotations() {
        return Codegen.optional(this.annotations);
    }
    /**
     * The ID of the alloydb backup.
     * 
     */
    @Export(name="backupId", refs={String.class}, tree="[0]")
    private Output<String> backupId;

    /**
     * @return The ID of the alloydb backup.
     * 
     */
    public Output<String> backupId() {
        return this.backupId;
    }
    /**
     * The full resource name of the backup source cluster (e.g., projects/{project}/locations/{location}/clusters/{clusterId}).
     * 
     */
    @Export(name="clusterName", refs={String.class}, tree="[0]")
    private Output<String> clusterName;

    /**
     * @return The full resource name of the backup source cluster (e.g., projects/{project}/locations/{location}/clusters/{clusterId}).
     * 
     */
    public Output<String> clusterName() {
        return this.clusterName;
    }
    /**
     * Output only. The system-generated UID of the cluster which was used to create this resource.
     * 
     */
    @Export(name="clusterUid", refs={String.class}, tree="[0]")
    private Output<String> clusterUid;

    /**
     * @return Output only. The system-generated UID of the cluster which was used to create this resource.
     * 
     */
    public Output<String> clusterUid() {
        return this.clusterUid;
    }
    /**
     * Output only. Create time stamp. A timestamp in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
     * Examples: &#34;2014-10-02T15:01:23Z&#34; and &#34;2014-10-02T15:01:23.045123456Z&#34;.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return Output only. Create time stamp. A timestamp in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
     * Examples: &#34;2014-10-02T15:01:23Z&#34; and &#34;2014-10-02T15:01:23.045123456Z&#34;.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * Output only. Delete time stamp. A timestamp in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
     * Examples: &#34;2014-10-02T15:01:23Z&#34; and &#34;2014-10-02T15:01:23.045123456Z&#34;.
     * 
     */
    @Export(name="deleteTime", refs={String.class}, tree="[0]")
    private Output<String> deleteTime;

    /**
     * @return Output only. Delete time stamp. A timestamp in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
     * Examples: &#34;2014-10-02T15:01:23Z&#34; and &#34;2014-10-02T15:01:23.045123456Z&#34;.
     * 
     */
    public Output<String> deleteTime() {
        return this.deleteTime;
    }
    /**
     * User-provided description of the backup.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return User-provided description of the backup.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * User-settable and human-readable display name for the Backup.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> displayName;

    /**
     * @return User-settable and human-readable display name for the Backup.
     * 
     */
    public Output<Optional<String>> displayName() {
        return Codegen.optional(this.displayName);
    }
    /**
     * All of annotations (key/value pairs) present on the resource in GCP, including the annotations configured through
     * Terraform, other clients and services.
     * 
     */
    @Export(name="effectiveAnnotations", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> effectiveAnnotations;

    /**
     * @return All of annotations (key/value pairs) present on the resource in GCP, including the annotations configured through
     * Terraform, other clients and services.
     * 
     */
    public Output<Map<String,String>> effectiveAnnotations() {
        return this.effectiveAnnotations;
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
     * EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK (customer-managed encryption key).
     * Structure is documented below.
     * 
     */
    @Export(name="encryptionConfig", refs={BackupEncryptionConfig.class}, tree="[0]")
    private Output</* @Nullable */ BackupEncryptionConfig> encryptionConfig;

    /**
     * @return EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK (customer-managed encryption key).
     * Structure is documented below.
     * 
     */
    public Output<Optional<BackupEncryptionConfig>> encryptionConfig() {
        return Codegen.optional(this.encryptionConfig);
    }
    /**
     * EncryptionInfo describes the encryption information of a cluster or a backup.
     * Structure is documented below.
     * 
     */
    @Export(name="encryptionInfos", refs={List.class,BackupEncryptionInfo.class}, tree="[0,1]")
    private Output<List<BackupEncryptionInfo>> encryptionInfos;

    /**
     * @return EncryptionInfo describes the encryption information of a cluster or a backup.
     * Structure is documented below.
     * 
     */
    public Output<List<BackupEncryptionInfo>> encryptionInfos() {
        return this.encryptionInfos;
    }
    /**
     * For Resource freshness validation (https://google.aip.dev/154)
     * 
     */
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output<String> etag;

    /**
     * @return For Resource freshness validation (https://google.aip.dev/154)
     * 
     */
    public Output<String> etag() {
        return this.etag;
    }
    /**
     * Output only. The QuantityBasedExpiry of the backup, specified by the backup&#39;s retention policy.
     * Once the expiry quantity is over retention, the backup is eligible to be garbage collected.
     * Structure is documented below.
     * 
     */
    @Export(name="expiryQuantities", refs={List.class,BackupExpiryQuantity.class}, tree="[0,1]")
    private Output<List<BackupExpiryQuantity>> expiryQuantities;

    /**
     * @return Output only. The QuantityBasedExpiry of the backup, specified by the backup&#39;s retention policy.
     * Once the expiry quantity is over retention, the backup is eligible to be garbage collected.
     * Structure is documented below.
     * 
     */
    public Output<List<BackupExpiryQuantity>> expiryQuantities() {
        return this.expiryQuantities;
    }
    /**
     * Output only. The time at which after the backup is eligible to be garbage collected.
     * It is the duration specified by the backup&#39;s retention policy, added to the backup&#39;s createTime.
     * 
     */
    @Export(name="expiryTime", refs={String.class}, tree="[0]")
    private Output<String> expiryTime;

    /**
     * @return Output only. The time at which after the backup is eligible to be garbage collected.
     * It is the duration specified by the backup&#39;s retention policy, added to the backup&#39;s createTime.
     * 
     */
    public Output<String> expiryTime() {
        return this.expiryTime;
    }
    /**
     * User-defined labels for the alloydb backup. An object containing a list of &#34;key&#34;: value pairs. Example: { &#34;name&#34;: &#34;wrench&#34;, &#34;mass&#34;: &#34;1.3kg&#34;, &#34;count&#34;: &#34;3&#34; }.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return User-defined labels for the alloydb backup. An object containing a list of &#34;key&#34;: value pairs. Example: { &#34;name&#34;: &#34;wrench&#34;, &#34;mass&#34;: &#34;1.3kg&#34;, &#34;count&#34;: &#34;3&#34; }.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * The location where the alloydb backup should reside.
     * 
     * ***
     * 
     */
    @Export(name="location", refs={String.class}, tree="[0]")
    private Output<String> location;

    /**
     * @return The location where the alloydb backup should reside.
     * 
     * ***
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    /**
     * Output only. The name of the backup resource with the format: * projects/{project}/locations/{region}/backups/{backupId}
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Output only. The name of the backup resource with the format: * projects/{project}/locations/{region}/backups/{backupId}
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
     * Output only. Reconciling (https://google.aip.dev/128#reconciliation), if true, indicates that the service is actively updating the resource.
     * This can happen due to user-triggered updates or system actions like failover or maintenance.
     * 
     */
    @Export(name="reconciling", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> reconciling;

    /**
     * @return Output only. Reconciling (https://google.aip.dev/128#reconciliation), if true, indicates that the service is actively updating the resource.
     * This can happen due to user-triggered updates or system actions like failover or maintenance.
     * 
     */
    public Output<Boolean> reconciling() {
        return this.reconciling;
    }
    /**
     * Output only. The size of the backup in bytes.
     * 
     */
    @Export(name="sizeBytes", refs={String.class}, tree="[0]")
    private Output<String> sizeBytes;

    /**
     * @return Output only. The size of the backup in bytes.
     * 
     */
    public Output<String> sizeBytes() {
        return this.sizeBytes;
    }
    /**
     * Output only. The current state of the backup.
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return Output only. The current state of the backup.
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * The backup type, which suggests the trigger for the backup.
     * Possible values are: `TYPE_UNSPECIFIED`, `ON_DEMAND`, `AUTOMATED`, `CONTINUOUS`.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return The backup type, which suggests the trigger for the backup.
     * Possible values are: `TYPE_UNSPECIFIED`, `ON_DEMAND`, `AUTOMATED`, `CONTINUOUS`.
     * 
     */
    public Output<String> type() {
        return this.type;
    }
    /**
     * Output only. The system-generated UID of the resource. The UID is assigned when the resource is created, and it is retained until it is deleted.
     * 
     */
    @Export(name="uid", refs={String.class}, tree="[0]")
    private Output<String> uid;

    /**
     * @return Output only. The system-generated UID of the resource. The UID is assigned when the resource is created, and it is retained until it is deleted.
     * 
     */
    public Output<String> uid() {
        return this.uid;
    }
    /**
     * Output only. Update time stamp. A timestamp in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
     * Examples: &#34;2014-10-02T15:01:23Z&#34; and &#34;2014-10-02T15:01:23.045123456Z&#34;.
     * 
     */
    @Export(name="updateTime", refs={String.class}, tree="[0]")
    private Output<String> updateTime;

    /**
     * @return Output only. Update time stamp. A timestamp in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
     * Examples: &#34;2014-10-02T15:01:23Z&#34; and &#34;2014-10-02T15:01:23.045123456Z&#34;.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
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
        super("gcp:alloydb/backup:Backup", name, args == null ? BackupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Backup(String name, Output<String> id, @Nullable BackupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:alloydb/backup:Backup", name, state, makeResourceOptions(options, id));
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
