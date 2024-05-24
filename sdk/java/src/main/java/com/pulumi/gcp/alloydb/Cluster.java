// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.alloydb;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.alloydb.ClusterArgs;
import com.pulumi.gcp.alloydb.inputs.ClusterState;
import com.pulumi.gcp.alloydb.outputs.ClusterAutomatedBackupPolicy;
import com.pulumi.gcp.alloydb.outputs.ClusterBackupSource;
import com.pulumi.gcp.alloydb.outputs.ClusterContinuousBackupConfig;
import com.pulumi.gcp.alloydb.outputs.ClusterContinuousBackupInfo;
import com.pulumi.gcp.alloydb.outputs.ClusterEncryptionConfig;
import com.pulumi.gcp.alloydb.outputs.ClusterEncryptionInfo;
import com.pulumi.gcp.alloydb.outputs.ClusterInitialUser;
import com.pulumi.gcp.alloydb.outputs.ClusterMaintenanceUpdatePolicy;
import com.pulumi.gcp.alloydb.outputs.ClusterMigrationSource;
import com.pulumi.gcp.alloydb.outputs.ClusterNetworkConfig;
import com.pulumi.gcp.alloydb.outputs.ClusterRestoreBackupSource;
import com.pulumi.gcp.alloydb.outputs.ClusterRestoreContinuousBackupSource;
import com.pulumi.gcp.alloydb.outputs.ClusterSecondaryConfig;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A managed alloydb cluster.
 * 
 * To get more information about Cluster, see:
 * 
 * * [API documentation](https://cloud.google.com/alloydb/docs/reference/rest/v1/projects.locations.clusters/create)
 * * How-to Guides
 *     * [AlloyDB](https://cloud.google.com/alloydb/docs/)
 * 
 * &gt; **Note:** Users can promote a secondary cluster to a primary cluster with the help of `cluster_type`.
 * To promote, users have to set the `cluster_type` property as `PRIMARY` and remove the `secondary_config` field from cluster configuration.
 * See Example.
 * 
 * ## Example Usage
 * 
 * ### Alloydb Cluster Basic
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.Network;
 * import com.pulumi.gcp.compute.NetworkArgs;
 * import com.pulumi.gcp.alloydb.Cluster;
 * import com.pulumi.gcp.alloydb.ClusterArgs;
 * import com.pulumi.gcp.organizations.OrganizationsFunctions;
 * import com.pulumi.gcp.organizations.inputs.GetProjectArgs;
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
 *         var defaultNetwork = new Network("defaultNetwork", NetworkArgs.builder()
 *             .name("alloydb-cluster")
 *             .build());
 * 
 *         var default_ = new Cluster("default", ClusterArgs.builder()
 *             .clusterId("alloydb-cluster")
 *             .location("us-central1")
 *             .network(defaultNetwork.id())
 *             .build());
 * 
 *         final var project = OrganizationsFunctions.getProject();
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Alloydb Cluster Full
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.Network;
 * import com.pulumi.gcp.compute.NetworkArgs;
 * import com.pulumi.gcp.alloydb.Cluster;
 * import com.pulumi.gcp.alloydb.ClusterArgs;
 * import com.pulumi.gcp.alloydb.inputs.ClusterInitialUserArgs;
 * import com.pulumi.gcp.alloydb.inputs.ClusterContinuousBackupConfigArgs;
 * import com.pulumi.gcp.alloydb.inputs.ClusterAutomatedBackupPolicyArgs;
 * import com.pulumi.gcp.alloydb.inputs.ClusterAutomatedBackupPolicyWeeklyScheduleArgs;
 * import com.pulumi.gcp.alloydb.inputs.ClusterAutomatedBackupPolicyQuantityBasedRetentionArgs;
 * import com.pulumi.gcp.organizations.OrganizationsFunctions;
 * import com.pulumi.gcp.organizations.inputs.GetProjectArgs;
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
 *         var default_ = new Network("default", NetworkArgs.builder()
 *             .name("alloydb-cluster-full")
 *             .build());
 * 
 *         var full = new Cluster("full", ClusterArgs.builder()
 *             .clusterId("alloydb-cluster-full")
 *             .location("us-central1")
 *             .network(default_.id())
 *             .databaseVersion("POSTGRES_15")
 *             .initialUser(ClusterInitialUserArgs.builder()
 *                 .user("alloydb-cluster-full")
 *                 .password("alloydb-cluster-full")
 *                 .build())
 *             .continuousBackupConfig(ClusterContinuousBackupConfigArgs.builder()
 *                 .enabled(true)
 *                 .recoveryWindowDays(14)
 *                 .build())
 *             .automatedBackupPolicy(ClusterAutomatedBackupPolicyArgs.builder()
 *                 .location("us-central1")
 *                 .backupWindow("1800s")
 *                 .enabled(true)
 *                 .weeklySchedule(ClusterAutomatedBackupPolicyWeeklyScheduleArgs.builder()
 *                     .daysOfWeeks("MONDAY")
 *                     .startTimes(ClusterAutomatedBackupPolicyWeeklyScheduleStartTimeArgs.builder()
 *                         .hours(23)
 *                         .minutes(0)
 *                         .seconds(0)
 *                         .nanos(0)
 *                         .build())
 *                     .build())
 *                 .quantityBasedRetention(ClusterAutomatedBackupPolicyQuantityBasedRetentionArgs.builder()
 *                     .count(1)
 *                     .build())
 *                 .labels(Map.of("test", "alloydb-cluster-full"))
 *                 .build())
 *             .labels(Map.of("test", "alloydb-cluster-full"))
 *             .build());
 * 
 *         final var project = OrganizationsFunctions.getProject();
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Alloydb Cluster Restore
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.ComputeFunctions;
 * import com.pulumi.gcp.compute.inputs.GetNetworkArgs;
 * import com.pulumi.gcp.alloydb.Cluster;
 * import com.pulumi.gcp.alloydb.ClusterArgs;
 * import com.pulumi.gcp.alloydb.inputs.ClusterInitialUserArgs;
 * import com.pulumi.gcp.alloydb.Instance;
 * import com.pulumi.gcp.alloydb.InstanceArgs;
 * import com.pulumi.gcp.alloydb.inputs.InstanceMachineConfigArgs;
 * import com.pulumi.gcp.alloydb.Backup;
 * import com.pulumi.gcp.alloydb.BackupArgs;
 * import com.pulumi.gcp.alloydb.inputs.ClusterRestoreBackupSourceArgs;
 * import com.pulumi.gcp.alloydb.inputs.ClusterRestoreContinuousBackupSourceArgs;
 * import com.pulumi.gcp.organizations.OrganizationsFunctions;
 * import com.pulumi.gcp.organizations.inputs.GetProjectArgs;
 * import com.pulumi.gcp.compute.GlobalAddress;
 * import com.pulumi.gcp.compute.GlobalAddressArgs;
 * import com.pulumi.gcp.servicenetworking.Connection;
 * import com.pulumi.gcp.servicenetworking.ConnectionArgs;
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
 *         final var default = ComputeFunctions.getNetwork(GetNetworkArgs.builder()
 *             .name("alloydb-network")
 *             .build());
 * 
 *         var source = new Cluster("source", ClusterArgs.builder()
 *             .clusterId("alloydb-source-cluster")
 *             .location("us-central1")
 *             .network(default_.id())
 *             .initialUser(ClusterInitialUserArgs.builder()
 *                 .password("alloydb-source-cluster")
 *                 .build())
 *             .build());
 * 
 *         var sourceInstance = new Instance("sourceInstance", InstanceArgs.builder()
 *             .cluster(source.name())
 *             .instanceId("alloydb-instance")
 *             .instanceType("PRIMARY")
 *             .machineConfig(InstanceMachineConfigArgs.builder()
 *                 .cpuCount(2)
 *                 .build())
 *             .build());
 * 
 *         var sourceBackup = new Backup("sourceBackup", BackupArgs.builder()
 *             .backupId("alloydb-backup")
 *             .location("us-central1")
 *             .clusterName(source.name())
 *             .build());
 * 
 *         var restoredFromBackup = new Cluster("restoredFromBackup", ClusterArgs.builder()
 *             .clusterId("alloydb-backup-restored")
 *             .location("us-central1")
 *             .network(default_.id())
 *             .restoreBackupSource(ClusterRestoreBackupSourceArgs.builder()
 *                 .backupName(sourceBackup.name())
 *                 .build())
 *             .build());
 * 
 *         var restoredViaPitr = new Cluster("restoredViaPitr", ClusterArgs.builder()
 *             .clusterId("alloydb-pitr-restored")
 *             .location("us-central1")
 *             .network(default_.id())
 *             .restoreContinuousBackupSource(ClusterRestoreContinuousBackupSourceArgs.builder()
 *                 .cluster(source.name())
 *                 .pointInTime("2023-08-03T19:19:00.094Z")
 *                 .build())
 *             .build());
 * 
 *         final var project = OrganizationsFunctions.getProject();
 * 
 *         var privateIpAlloc = new GlobalAddress("privateIpAlloc", GlobalAddressArgs.builder()
 *             .name("alloydb-source-cluster")
 *             .addressType("INTERNAL")
 *             .purpose("VPC_PEERING")
 *             .prefixLength(16)
 *             .network(default_.id())
 *             .build());
 * 
 *         var vpcConnection = new Connection("vpcConnection", ConnectionArgs.builder()
 *             .network(default_.id())
 *             .service("servicenetworking.googleapis.com")
 *             .reservedPeeringRanges(privateIpAlloc.name())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Alloydb Secondary Cluster Basic
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.Network;
 * import com.pulumi.gcp.compute.NetworkArgs;
 * import com.pulumi.gcp.alloydb.Cluster;
 * import com.pulumi.gcp.alloydb.ClusterArgs;
 * import com.pulumi.gcp.alloydb.Instance;
 * import com.pulumi.gcp.alloydb.InstanceArgs;
 * import com.pulumi.gcp.alloydb.inputs.InstanceMachineConfigArgs;
 * import com.pulumi.gcp.alloydb.inputs.ClusterContinuousBackupConfigArgs;
 * import com.pulumi.gcp.alloydb.inputs.ClusterSecondaryConfigArgs;
 * import com.pulumi.gcp.organizations.OrganizationsFunctions;
 * import com.pulumi.gcp.organizations.inputs.GetProjectArgs;
 * import com.pulumi.gcp.compute.GlobalAddress;
 * import com.pulumi.gcp.compute.GlobalAddressArgs;
 * import com.pulumi.gcp.servicenetworking.Connection;
 * import com.pulumi.gcp.servicenetworking.ConnectionArgs;
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
 *         var default_ = new Network("default", NetworkArgs.builder()
 *             .name("alloydb-secondary-cluster")
 *             .build());
 * 
 *         var primary = new Cluster("primary", ClusterArgs.builder()
 *             .clusterId("alloydb-primary-cluster")
 *             .location("us-central1")
 *             .network(default_.id())
 *             .build());
 * 
 *         var primaryInstance = new Instance("primaryInstance", InstanceArgs.builder()
 *             .cluster(primary.name())
 *             .instanceId("alloydb-primary-instance")
 *             .instanceType("PRIMARY")
 *             .machineConfig(InstanceMachineConfigArgs.builder()
 *                 .cpuCount(2)
 *                 .build())
 *             .build());
 * 
 *         var secondary = new Cluster("secondary", ClusterArgs.builder()
 *             .clusterId("alloydb-secondary-cluster")
 *             .location("us-east1")
 *             .network(default_.id())
 *             .clusterType("SECONDARY")
 *             .continuousBackupConfig(ClusterContinuousBackupConfigArgs.builder()
 *                 .enabled(false)
 *                 .build())
 *             .secondaryConfig(ClusterSecondaryConfigArgs.builder()
 *                 .primaryClusterName(primary.name())
 *                 .build())
 *             .build());
 * 
 *         final var project = OrganizationsFunctions.getProject();
 * 
 *         var privateIpAlloc = new GlobalAddress("privateIpAlloc", GlobalAddressArgs.builder()
 *             .name("alloydb-secondary-cluster")
 *             .addressType("INTERNAL")
 *             .purpose("VPC_PEERING")
 *             .prefixLength(16)
 *             .network(default_.id())
 *             .build());
 * 
 *         var vpcConnection = new Connection("vpcConnection", ConnectionArgs.builder()
 *             .network(default_.id())
 *             .service("servicenetworking.googleapis.com")
 *             .reservedPeeringRanges(privateIpAlloc.name())
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
 * Cluster can be imported using any of these accepted formats:
 * 
 * * `projects/{{project}}/locations/{{location}}/clusters/{{cluster_id}}`
 * 
 * * `{{project}}/{{location}}/{{cluster_id}}`
 * 
 * * `{{location}}/{{cluster_id}}`
 * 
 * * `{{cluster_id}}`
 * 
 * When using the `pulumi import` command, Cluster can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:alloydb/cluster:Cluster default projects/{{project}}/locations/{{location}}/clusters/{{cluster_id}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:alloydb/cluster:Cluster default {{project}}/{{location}}/{{cluster_id}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:alloydb/cluster:Cluster default {{location}}/{{cluster_id}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:alloydb/cluster:Cluster default {{cluster_id}}
 * ```
 * 
 */
@ResourceType(type="gcp:alloydb/cluster:Cluster")
public class Cluster extends com.pulumi.resources.CustomResource {
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
     * The automated backup policy for this cluster. AutomatedBackupPolicy is disabled by default.
     * Structure is documented below.
     * 
     */
    @Export(name="automatedBackupPolicy", refs={ClusterAutomatedBackupPolicy.class}, tree="[0]")
    private Output<ClusterAutomatedBackupPolicy> automatedBackupPolicy;

    /**
     * @return The automated backup policy for this cluster. AutomatedBackupPolicy is disabled by default.
     * Structure is documented below.
     * 
     */
    public Output<ClusterAutomatedBackupPolicy> automatedBackupPolicy() {
        return this.automatedBackupPolicy;
    }
    /**
     * Cluster created from backup.
     * Structure is documented below.
     * 
     */
    @Export(name="backupSources", refs={List.class,ClusterBackupSource.class}, tree="[0,1]")
    private Output<List<ClusterBackupSource>> backupSources;

    /**
     * @return Cluster created from backup.
     * Structure is documented below.
     * 
     */
    public Output<List<ClusterBackupSource>> backupSources() {
        return this.backupSources;
    }
    /**
     * The ID of the alloydb cluster.
     * 
     */
    @Export(name="clusterId", refs={String.class}, tree="[0]")
    private Output<String> clusterId;

    /**
     * @return The ID of the alloydb cluster.
     * 
     */
    public Output<String> clusterId() {
        return this.clusterId;
    }
    /**
     * The type of cluster. If not set, defaults to PRIMARY.
     * Default value is `PRIMARY`.
     * Possible values are: `PRIMARY`, `SECONDARY`.
     * 
     */
    @Export(name="clusterType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clusterType;

    /**
     * @return The type of cluster. If not set, defaults to PRIMARY.
     * Default value is `PRIMARY`.
     * Possible values are: `PRIMARY`, `SECONDARY`.
     * 
     */
    public Output<Optional<String>> clusterType() {
        return Codegen.optional(this.clusterType);
    }
    /**
     * The continuous backup config for this cluster.
     * If no policy is provided then the default policy will be used. The default policy takes one backup a day and retains backups for 14 days.
     * Structure is documented below.
     * 
     */
    @Export(name="continuousBackupConfig", refs={ClusterContinuousBackupConfig.class}, tree="[0]")
    private Output<ClusterContinuousBackupConfig> continuousBackupConfig;

    /**
     * @return The continuous backup config for this cluster.
     * If no policy is provided then the default policy will be used. The default policy takes one backup a day and retains backups for 14 days.
     * Structure is documented below.
     * 
     */
    public Output<ClusterContinuousBackupConfig> continuousBackupConfig() {
        return this.continuousBackupConfig;
    }
    /**
     * ContinuousBackupInfo describes the continuous backup properties of a cluster.
     * Structure is documented below.
     * 
     */
    @Export(name="continuousBackupInfos", refs={List.class,ClusterContinuousBackupInfo.class}, tree="[0,1]")
    private Output<List<ClusterContinuousBackupInfo>> continuousBackupInfos;

    /**
     * @return ContinuousBackupInfo describes the continuous backup properties of a cluster.
     * Structure is documented below.
     * 
     */
    public Output<List<ClusterContinuousBackupInfo>> continuousBackupInfos() {
        return this.continuousBackupInfos;
    }
    /**
     * The database engine major version. This is an optional field and it&#39;s populated at the Cluster creation time. This field cannot be changed after cluster creation.
     * 
     */
    @Export(name="databaseVersion", refs={String.class}, tree="[0]")
    private Output<String> databaseVersion;

    /**
     * @return The database engine major version. This is an optional field and it&#39;s populated at the Cluster creation time. This field cannot be changed after cluster creation.
     * 
     */
    public Output<String> databaseVersion() {
        return this.databaseVersion;
    }
    /**
     * Policy to determine if the cluster should be deleted forcefully.
     * Deleting a cluster forcefully, deletes the cluster and all its associated instances within the cluster.
     * Deleting a Secondary cluster with a secondary instance REQUIRES setting deletion_policy = &#34;FORCE&#34; otherwise an error is returned. This is needed as there is no support to delete just the secondary instance, and the only way to delete secondary instance is to delete the associated secondary cluster forcefully which also deletes the secondary instance.
     * 
     */
    @Export(name="deletionPolicy", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> deletionPolicy;

    /**
     * @return Policy to determine if the cluster should be deleted forcefully.
     * Deleting a cluster forcefully, deletes the cluster and all its associated instances within the cluster.
     * Deleting a Secondary cluster with a secondary instance REQUIRES setting deletion_policy = &#34;FORCE&#34; otherwise an error is returned. This is needed as there is no support to delete just the secondary instance, and the only way to delete secondary instance is to delete the associated secondary cluster forcefully which also deletes the secondary instance.
     * 
     */
    public Output<Optional<String>> deletionPolicy() {
        return Codegen.optional(this.deletionPolicy);
    }
    /**
     * User-settable and human-readable display name for the Cluster.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> displayName;

    /**
     * @return User-settable and human-readable display name for the Cluster.
     * 
     */
    public Output<Optional<String>> displayName() {
        return Codegen.optional(this.displayName);
    }
    @Export(name="effectiveAnnotations", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> effectiveAnnotations;

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
    @Export(name="encryptionConfig", refs={ClusterEncryptionConfig.class}, tree="[0]")
    private Output</* @Nullable */ ClusterEncryptionConfig> encryptionConfig;

    /**
     * @return EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK (customer-managed encryption key).
     * Structure is documented below.
     * 
     */
    public Output<Optional<ClusterEncryptionConfig>> encryptionConfig() {
        return Codegen.optional(this.encryptionConfig);
    }
    /**
     * (Output)
     * Output only. The encryption information for the WALs and backups required for ContinuousBackup.
     * Structure is documented below.
     * 
     */
    @Export(name="encryptionInfos", refs={List.class,ClusterEncryptionInfo.class}, tree="[0,1]")
    private Output<List<ClusterEncryptionInfo>> encryptionInfos;

    /**
     * @return (Output)
     * Output only. The encryption information for the WALs and backups required for ContinuousBackup.
     * Structure is documented below.
     * 
     */
    public Output<List<ClusterEncryptionInfo>> encryptionInfos() {
        return this.encryptionInfos;
    }
    /**
     * For Resource freshness validation (https://google.aip.dev/154)
     * 
     */
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> etag;

    /**
     * @return For Resource freshness validation (https://google.aip.dev/154)
     * 
     */
    public Output<Optional<String>> etag() {
        return Codegen.optional(this.etag);
    }
    /**
     * Initial user to setup during cluster creation.
     * Structure is documented below.
     * 
     */
    @Export(name="initialUser", refs={ClusterInitialUser.class}, tree="[0]")
    private Output</* @Nullable */ ClusterInitialUser> initialUser;

    /**
     * @return Initial user to setup during cluster creation.
     * Structure is documented below.
     * 
     */
    public Output<Optional<ClusterInitialUser>> initialUser() {
        return Codegen.optional(this.initialUser);
    }
    /**
     * User-defined labels for the alloydb cluster.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return User-defined labels for the alloydb cluster.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * The location where the alloydb cluster should reside.
     * 
     * ***
     * 
     */
    @Export(name="location", refs={String.class}, tree="[0]")
    private Output<String> location;

    /**
     * @return The location where the alloydb cluster should reside.
     * 
     * ***
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    /**
     * MaintenanceUpdatePolicy defines the policy for system updates.
     * Structure is documented below.
     * 
     */
    @Export(name="maintenanceUpdatePolicy", refs={ClusterMaintenanceUpdatePolicy.class}, tree="[0]")
    private Output</* @Nullable */ ClusterMaintenanceUpdatePolicy> maintenanceUpdatePolicy;

    /**
     * @return MaintenanceUpdatePolicy defines the policy for system updates.
     * Structure is documented below.
     * 
     */
    public Output<Optional<ClusterMaintenanceUpdatePolicy>> maintenanceUpdatePolicy() {
        return Codegen.optional(this.maintenanceUpdatePolicy);
    }
    /**
     * Cluster created via DMS migration.
     * Structure is documented below.
     * 
     */
    @Export(name="migrationSources", refs={List.class,ClusterMigrationSource.class}, tree="[0,1]")
    private Output<List<ClusterMigrationSource>> migrationSources;

    /**
     * @return Cluster created via DMS migration.
     * Structure is documented below.
     * 
     */
    public Output<List<ClusterMigrationSource>> migrationSources() {
        return this.migrationSources;
    }
    /**
     * The name of the cluster resource.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the cluster resource.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * (Optional, Deprecated)
     * The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:
     * &#34;projects/{projectNumber}/global/networks/{network_id}&#34;.
     * 
     * &gt; **Warning:** `network` is deprecated and will be removed in a future major release. Instead, use `network_config` to define the network configuration.
     * 
     * @deprecated
     * `network` is deprecated and will be removed in a future major release. Instead, use `network_config` to define the network configuration.
     * 
     */
    @Deprecated /* `network` is deprecated and will be removed in a future major release. Instead, use `network_config` to define the network configuration. */
    @Export(name="network", refs={String.class}, tree="[0]")
    private Output<String> network;

    /**
     * @return (Optional, Deprecated)
     * The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:
     * &#34;projects/{projectNumber}/global/networks/{network_id}&#34;.
     * 
     * &gt; **Warning:** `network` is deprecated and will be removed in a future major release. Instead, use `network_config` to define the network configuration.
     * 
     */
    public Output<String> network() {
        return this.network;
    }
    /**
     * Metadata related to network configuration.
     * Structure is documented below.
     * 
     */
    @Export(name="networkConfig", refs={ClusterNetworkConfig.class}, tree="[0]")
    private Output<ClusterNetworkConfig> networkConfig;

    /**
     * @return Metadata related to network configuration.
     * Structure is documented below.
     * 
     */
    public Output<ClusterNetworkConfig> networkConfig() {
        return this.networkConfig;
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
     * Output only. Reconciling (https://google.aip.dev/128#reconciliation).
     * Set to true if the current state of Cluster does not match the user&#39;s intended state, and the service is actively updating the resource to reconcile them.
     * This can happen due to user-triggered updates or system actions like failover or maintenance.
     * 
     */
    @Export(name="reconciling", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> reconciling;

    /**
     * @return Output only. Reconciling (https://google.aip.dev/128#reconciliation).
     * Set to true if the current state of Cluster does not match the user&#39;s intended state, and the service is actively updating the resource to reconcile them.
     * This can happen due to user-triggered updates or system actions like failover or maintenance.
     * 
     */
    public Output<Boolean> reconciling() {
        return this.reconciling;
    }
    /**
     * The source when restoring from a backup. Conflicts with &#39;restore_continuous_backup_source&#39;, both can&#39;t be set together.
     * Structure is documented below.
     * 
     */
    @Export(name="restoreBackupSource", refs={ClusterRestoreBackupSource.class}, tree="[0]")
    private Output</* @Nullable */ ClusterRestoreBackupSource> restoreBackupSource;

    /**
     * @return The source when restoring from a backup. Conflicts with &#39;restore_continuous_backup_source&#39;, both can&#39;t be set together.
     * Structure is documented below.
     * 
     */
    public Output<Optional<ClusterRestoreBackupSource>> restoreBackupSource() {
        return Codegen.optional(this.restoreBackupSource);
    }
    /**
     * The source when restoring via point in time recovery (PITR). Conflicts with &#39;restore_backup_source&#39;, both can&#39;t be set together.
     * Structure is documented below.
     * 
     */
    @Export(name="restoreContinuousBackupSource", refs={ClusterRestoreContinuousBackupSource.class}, tree="[0]")
    private Output</* @Nullable */ ClusterRestoreContinuousBackupSource> restoreContinuousBackupSource;

    /**
     * @return The source when restoring via point in time recovery (PITR). Conflicts with &#39;restore_backup_source&#39;, both can&#39;t be set together.
     * Structure is documented below.
     * 
     */
    public Output<Optional<ClusterRestoreContinuousBackupSource>> restoreContinuousBackupSource() {
        return Codegen.optional(this.restoreContinuousBackupSource);
    }
    /**
     * Configuration of the secondary cluster for Cross Region Replication. This should be set if and only if the cluster is of type SECONDARY.
     * Structure is documented below.
     * 
     */
    @Export(name="secondaryConfig", refs={ClusterSecondaryConfig.class}, tree="[0]")
    private Output</* @Nullable */ ClusterSecondaryConfig> secondaryConfig;

    /**
     * @return Configuration of the secondary cluster for Cross Region Replication. This should be set if and only if the cluster is of type SECONDARY.
     * Structure is documented below.
     * 
     */
    public Output<Optional<ClusterSecondaryConfig>> secondaryConfig() {
        return Codegen.optional(this.secondaryConfig);
    }
    /**
     * Output only. The current serving state of the cluster.
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return Output only. The current serving state of the cluster.
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * The system-generated UID of the resource.
     * 
     */
    @Export(name="uid", refs={String.class}, tree="[0]")
    private Output<String> uid;

    /**
     * @return The system-generated UID of the resource.
     * 
     */
    public Output<String> uid() {
        return this.uid;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Cluster(String name) {
        this(name, ClusterArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Cluster(String name, ClusterArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Cluster(String name, ClusterArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:alloydb/cluster:Cluster", name, args == null ? ClusterArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Cluster(String name, Output<String> id, @Nullable ClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:alloydb/cluster:Cluster", name, state, makeResourceOptions(options, id));
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
    public static Cluster get(String name, Output<String> id, @Nullable ClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Cluster(name, id, state, options);
    }
}
