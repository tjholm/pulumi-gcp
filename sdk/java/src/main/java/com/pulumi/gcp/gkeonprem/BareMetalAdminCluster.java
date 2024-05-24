// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.gkeonprem;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.gkeonprem.BareMetalAdminClusterArgs;
import com.pulumi.gcp.gkeonprem.inputs.BareMetalAdminClusterState;
import com.pulumi.gcp.gkeonprem.outputs.BareMetalAdminClusterClusterOperations;
import com.pulumi.gcp.gkeonprem.outputs.BareMetalAdminClusterControlPlane;
import com.pulumi.gcp.gkeonprem.outputs.BareMetalAdminClusterFleet;
import com.pulumi.gcp.gkeonprem.outputs.BareMetalAdminClusterLoadBalancer;
import com.pulumi.gcp.gkeonprem.outputs.BareMetalAdminClusterMaintenanceConfig;
import com.pulumi.gcp.gkeonprem.outputs.BareMetalAdminClusterNetworkConfig;
import com.pulumi.gcp.gkeonprem.outputs.BareMetalAdminClusterNodeAccessConfig;
import com.pulumi.gcp.gkeonprem.outputs.BareMetalAdminClusterNodeConfig;
import com.pulumi.gcp.gkeonprem.outputs.BareMetalAdminClusterProxy;
import com.pulumi.gcp.gkeonprem.outputs.BareMetalAdminClusterSecurityConfig;
import com.pulumi.gcp.gkeonprem.outputs.BareMetalAdminClusterStatus;
import com.pulumi.gcp.gkeonprem.outputs.BareMetalAdminClusterStorage;
import com.pulumi.gcp.gkeonprem.outputs.BareMetalAdminClusterValidationCheck;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A Google Bare Metal Admin Cluster.
 * 
 * ## Example Usage
 * 
 * ### Gkeonprem Bare Metal Admin Cluster Basic
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.gkeonprem.BareMetalAdminCluster;
 * import com.pulumi.gcp.gkeonprem.BareMetalAdminClusterArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalAdminClusterNetworkConfigArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalAdminClusterNetworkConfigIslandModeCidrArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalAdminClusterNodeConfigArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalAdminClusterControlPlaneArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalAdminClusterLoadBalancerArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalAdminClusterLoadBalancerPortConfigArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalAdminClusterLoadBalancerVipConfigArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalAdminClusterStorageArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalAdminClusterStorageLvpShareConfigArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalAdminClusterStorageLvpShareConfigLvpConfigArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalAdminClusterStorageLvpNodeMountsConfigArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalAdminClusterNodeAccessConfigArgs;
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
 *         var admin_cluster_basic = new BareMetalAdminCluster("admin-cluster-basic", BareMetalAdminClusterArgs.builder()
 *             .name("my-cluster")
 *             .location("us-west1")
 *             .bareMetalVersion("1.13.4")
 *             .networkConfig(BareMetalAdminClusterNetworkConfigArgs.builder()
 *                 .islandModeCidr(BareMetalAdminClusterNetworkConfigIslandModeCidrArgs.builder()
 *                     .serviceAddressCidrBlocks("172.26.0.0/16")
 *                     .podAddressCidrBlocks("10.240.0.0/13")
 *                     .build())
 *                 .build())
 *             .nodeConfig(BareMetalAdminClusterNodeConfigArgs.builder()
 *                 .maxPodsPerNode(250)
 *                 .build())
 *             .controlPlane(BareMetalAdminClusterControlPlaneArgs.builder()
 *                 .controlPlaneNodePoolConfig(BareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigArgs.builder()
 *                     .nodePoolConfig(BareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigArgs.builder()
 *                         .labels()
 *                         .operatingSystem("LINUX")
 *                         .nodeConfigs(                        
 *                             BareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigNodeConfigArgs.builder()
 *                                 .labels()
 *                                 .nodeIp("10.200.0.2")
 *                                 .build(),
 *                             BareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigNodeConfigArgs.builder()
 *                                 .labels()
 *                                 .nodeIp("10.200.0.3")
 *                                 .build(),
 *                             BareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigNodeConfigArgs.builder()
 *                                 .labels()
 *                                 .nodeIp("10.200.0.4")
 *                                 .build())
 *                         .build())
 *                     .build())
 *                 .build())
 *             .loadBalancer(BareMetalAdminClusterLoadBalancerArgs.builder()
 *                 .portConfig(BareMetalAdminClusterLoadBalancerPortConfigArgs.builder()
 *                     .controlPlaneLoadBalancerPort(443)
 *                     .build())
 *                 .vipConfig(BareMetalAdminClusterLoadBalancerVipConfigArgs.builder()
 *                     .controlPlaneVip("10.200.0.5")
 *                     .build())
 *                 .build())
 *             .storage(BareMetalAdminClusterStorageArgs.builder()
 *                 .lvpShareConfig(BareMetalAdminClusterStorageLvpShareConfigArgs.builder()
 *                     .lvpConfig(BareMetalAdminClusterStorageLvpShareConfigLvpConfigArgs.builder()
 *                         .path("/mnt/localpv-share")
 *                         .storageClass("local-shared")
 *                         .build())
 *                     .sharedPathPvCount(5)
 *                     .build())
 *                 .lvpNodeMountsConfig(BareMetalAdminClusterStorageLvpNodeMountsConfigArgs.builder()
 *                     .path("/mnt/localpv-disk")
 *                     .storageClass("local-disks")
 *                     .build())
 *                 .build())
 *             .nodeAccessConfig(BareMetalAdminClusterNodeAccessConfigArgs.builder()
 *                 .loginUser("root")
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Gkeonprem Bare Metal Admin Cluster Full
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.gkeonprem.BareMetalAdminCluster;
 * import com.pulumi.gcp.gkeonprem.BareMetalAdminClusterArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalAdminClusterNetworkConfigArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalAdminClusterNetworkConfigIslandModeCidrArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalAdminClusterNodeConfigArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalAdminClusterControlPlaneArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalAdminClusterLoadBalancerArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalAdminClusterLoadBalancerPortConfigArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalAdminClusterLoadBalancerVipConfigArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalAdminClusterLoadBalancerManualLbConfigArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalAdminClusterStorageArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalAdminClusterStorageLvpShareConfigArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalAdminClusterStorageLvpShareConfigLvpConfigArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalAdminClusterStorageLvpNodeMountsConfigArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalAdminClusterNodeAccessConfigArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalAdminClusterSecurityConfigArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalAdminClusterSecurityConfigAuthorizationArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalAdminClusterMaintenanceConfigArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalAdminClusterClusterOperationsArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalAdminClusterProxyArgs;
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
 *         var admin_cluster_basic = new BareMetalAdminCluster("admin-cluster-basic", BareMetalAdminClusterArgs.builder()
 *             .name("my-cluster")
 *             .location("us-west1")
 *             .description("test description")
 *             .bareMetalVersion("1.13.4")
 *             .annotations(Map.of("env", "test"))
 *             .networkConfig(BareMetalAdminClusterNetworkConfigArgs.builder()
 *                 .islandModeCidr(BareMetalAdminClusterNetworkConfigIslandModeCidrArgs.builder()
 *                     .serviceAddressCidrBlocks("172.26.0.0/16")
 *                     .podAddressCidrBlocks("10.240.0.0/13")
 *                     .build())
 *                 .build())
 *             .nodeConfig(BareMetalAdminClusterNodeConfigArgs.builder()
 *                 .maxPodsPerNode(250)
 *                 .build())
 *             .controlPlane(BareMetalAdminClusterControlPlaneArgs.builder()
 *                 .controlPlaneNodePoolConfig(BareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigArgs.builder()
 *                     .nodePoolConfig(BareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigArgs.builder()
 *                         .labels()
 *                         .operatingSystem("LINUX")
 *                         .nodeConfigs(                        
 *                             BareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigNodeConfigArgs.builder()
 *                                 .labels()
 *                                 .nodeIp("10.200.0.2")
 *                                 .build(),
 *                             BareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigNodeConfigArgs.builder()
 *                                 .labels()
 *                                 .nodeIp("10.200.0.3")
 *                                 .build(),
 *                             BareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigNodeConfigArgs.builder()
 *                                 .labels()
 *                                 .nodeIp("10.200.0.4")
 *                                 .build())
 *                         .taints(BareMetalAdminClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigTaintArgs.builder()
 *                             .key("test-key")
 *                             .value("test-value")
 *                             .effect("NO_EXECUTE")
 *                             .build())
 *                         .build())
 *                     .build())
 *                 .apiServerArgs(BareMetalAdminClusterControlPlaneApiServerArgArgs.builder()
 *                     .argument("test argument")
 *                     .value("test value")
 *                     .build())
 *                 .build())
 *             .loadBalancer(BareMetalAdminClusterLoadBalancerArgs.builder()
 *                 .portConfig(BareMetalAdminClusterLoadBalancerPortConfigArgs.builder()
 *                     .controlPlaneLoadBalancerPort(443)
 *                     .build())
 *                 .vipConfig(BareMetalAdminClusterLoadBalancerVipConfigArgs.builder()
 *                     .controlPlaneVip("10.200.0.5")
 *                     .build())
 *                 .manualLbConfig(BareMetalAdminClusterLoadBalancerManualLbConfigArgs.builder()
 *                     .enabled(true)
 *                     .build())
 *                 .build())
 *             .storage(BareMetalAdminClusterStorageArgs.builder()
 *                 .lvpShareConfig(BareMetalAdminClusterStorageLvpShareConfigArgs.builder()
 *                     .lvpConfig(BareMetalAdminClusterStorageLvpShareConfigLvpConfigArgs.builder()
 *                         .path("/mnt/localpv-share")
 *                         .storageClass("local-shared")
 *                         .build())
 *                     .sharedPathPvCount(5)
 *                     .build())
 *                 .lvpNodeMountsConfig(BareMetalAdminClusterStorageLvpNodeMountsConfigArgs.builder()
 *                     .path("/mnt/localpv-disk")
 *                     .storageClass("local-disks")
 *                     .build())
 *                 .build())
 *             .nodeAccessConfig(BareMetalAdminClusterNodeAccessConfigArgs.builder()
 *                 .loginUser("root")
 *                 .build())
 *             .securityConfig(BareMetalAdminClusterSecurityConfigArgs.builder()
 *                 .authorization(BareMetalAdminClusterSecurityConfigAuthorizationArgs.builder()
 *                     .adminUsers(BareMetalAdminClusterSecurityConfigAuthorizationAdminUserArgs.builder()
 *                         .username("admin{@literal @}hashicorptest.com")
 *                         .build())
 *                     .build())
 *                 .build())
 *             .maintenanceConfig(BareMetalAdminClusterMaintenanceConfigArgs.builder()
 *                 .maintenanceAddressCidrBlocks(                
 *                     "10.0.0.1/32",
 *                     "10.0.0.2/32")
 *                 .build())
 *             .clusterOperations(BareMetalAdminClusterClusterOperationsArgs.builder()
 *                 .enableApplicationLogs(true)
 *                 .build())
 *             .proxy(BareMetalAdminClusterProxyArgs.builder()
 *                 .uri("test proxy uri")
 *                 .noProxies("127.0.0.1")
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
 * BareMetalAdminCluster can be imported using any of these accepted formats:
 * 
 * * `projects/{{project}}/locations/{{location}}/bareMetalAdminClusters/{{name}}`
 * 
 * * `{{project}}/{{location}}/{{name}}`
 * 
 * * `{{location}}/{{name}}`
 * 
 * When using the `pulumi import` command, BareMetalAdminCluster can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:gkeonprem/bareMetalAdminCluster:BareMetalAdminCluster default projects/{{project}}/locations/{{location}}/bareMetalAdminClusters/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:gkeonprem/bareMetalAdminCluster:BareMetalAdminCluster default {{project}}/{{location}}/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:gkeonprem/bareMetalAdminCluster:BareMetalAdminCluster default {{location}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:gkeonprem/bareMetalAdminCluster:BareMetalAdminCluster")
public class BareMetalAdminCluster extends com.pulumi.resources.CustomResource {
    /**
     * Annotations on the Bare Metal Admin Cluster.
     * This field has the same restrictions as Kubernetes annotations.
     * The total size of all keys and values combined is limited to 256k.
     * Key can have 2 segments: prefix (optional) and name (required),
     * separated by a slash (/).
     * Prefix must be a DNS subdomain.
     * Name must be 63 characters or less, begin and end with alphanumerics,
     * with dashes (-), underscores (_), dots (.), and alphanumerics between.
     * 
     * **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
     * Please refer to the field `effective_annotations` for all of the annotations present on the resource.
     * 
     */
    @Export(name="annotations", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> annotations;

    /**
     * @return Annotations on the Bare Metal Admin Cluster.
     * This field has the same restrictions as Kubernetes annotations.
     * The total size of all keys and values combined is limited to 256k.
     * Key can have 2 segments: prefix (optional) and name (required),
     * separated by a slash (/).
     * Prefix must be a DNS subdomain.
     * Name must be 63 characters or less, begin and end with alphanumerics,
     * with dashes (-), underscores (_), dots (.), and alphanumerics between.
     * 
     * **Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
     * Please refer to the field `effective_annotations` for all of the annotations present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> annotations() {
        return Codegen.optional(this.annotations);
    }
    /**
     * A human readable description of this Bare Metal Admin Cluster.
     * 
     */
    @Export(name="bareMetalVersion", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> bareMetalVersion;

    /**
     * @return A human readable description of this Bare Metal Admin Cluster.
     * 
     */
    public Output<Optional<String>> bareMetalVersion() {
        return Codegen.optional(this.bareMetalVersion);
    }
    /**
     * Specifies the Admin Cluster&#39;s observability infrastructure.
     * Structure is documented below.
     * 
     */
    @Export(name="clusterOperations", refs={BareMetalAdminClusterClusterOperations.class}, tree="[0]")
    private Output</* @Nullable */ BareMetalAdminClusterClusterOperations> clusterOperations;

    /**
     * @return Specifies the Admin Cluster&#39;s observability infrastructure.
     * Structure is documented below.
     * 
     */
    public Output<Optional<BareMetalAdminClusterClusterOperations>> clusterOperations() {
        return Codegen.optional(this.clusterOperations);
    }
    /**
     * Specifies the control plane configuration.
     * Structure is documented below.
     * 
     */
    @Export(name="controlPlane", refs={BareMetalAdminClusterControlPlane.class}, tree="[0]")
    private Output</* @Nullable */ BareMetalAdminClusterControlPlane> controlPlane;

    /**
     * @return Specifies the control plane configuration.
     * Structure is documented below.
     * 
     */
    public Output<Optional<BareMetalAdminClusterControlPlane>> controlPlane() {
        return Codegen.optional(this.controlPlane);
    }
    /**
     * The time the cluster was created, in RFC3339 text format.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return The time the cluster was created, in RFC3339 text format.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * The time the cluster was deleted, in RFC3339 text format.
     * 
     */
    @Export(name="deleteTime", refs={String.class}, tree="[0]")
    private Output<String> deleteTime;

    /**
     * @return The time the cluster was deleted, in RFC3339 text format.
     * 
     */
    public Output<String> deleteTime() {
        return this.deleteTime;
    }
    /**
     * A human readable description of this Bare Metal Admin Cluster.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A human readable description of this Bare Metal Admin Cluster.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    @Export(name="effectiveAnnotations", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> effectiveAnnotations;

    public Output<Map<String,String>> effectiveAnnotations() {
        return this.effectiveAnnotations;
    }
    /**
     * The IP address name of Bare Metal Admin Cluster&#39;s API server.
     * 
     */
    @Export(name="endpoint", refs={String.class}, tree="[0]")
    private Output<String> endpoint;

    /**
     * @return The IP address name of Bare Metal Admin Cluster&#39;s API server.
     * 
     */
    public Output<String> endpoint() {
        return this.endpoint;
    }
    /**
     * This checksum is computed by the server based on the value of other
     * fields, and may be sent on update and delete requests to ensure the
     * client has an up-to-date value before proceeding.
     * Allows clients to perform consistent read-modify-writes
     * through optimistic concurrency control.
     * 
     */
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output<String> etag;

    /**
     * @return This checksum is computed by the server based on the value of other
     * fields, and may be sent on update and delete requests to ensure the
     * client has an up-to-date value before proceeding.
     * Allows clients to perform consistent read-modify-writes
     * through optimistic concurrency control.
     * 
     */
    public Output<String> etag() {
        return this.etag;
    }
    /**
     * Fleet related configuration.
     * Fleets are a Google Cloud concept for logically organizing clusters,
     * letting you use and manage multi-cluster capabilities and apply
     * consistent policies across your systems.
     * See [Anthos Fleets](https://cloud.google.com/anthos/multicluster-management/fleets) for
     * more details on Anthos multi-cluster capabilities using Fleets.
     * Structure is documented below.
     * 
     */
    @Export(name="fleets", refs={List.class,BareMetalAdminClusterFleet.class}, tree="[0,1]")
    private Output<List<BareMetalAdminClusterFleet>> fleets;

    /**
     * @return Fleet related configuration.
     * Fleets are a Google Cloud concept for logically organizing clusters,
     * letting you use and manage multi-cluster capabilities and apply
     * consistent policies across your systems.
     * See [Anthos Fleets](https://cloud.google.com/anthos/multicluster-management/fleets) for
     * more details on Anthos multi-cluster capabilities using Fleets.
     * Structure is documented below.
     * 
     */
    public Output<List<BareMetalAdminClusterFleet>> fleets() {
        return this.fleets;
    }
    /**
     * Specifies the load balancer configuration.
     * Structure is documented below.
     * 
     */
    @Export(name="loadBalancer", refs={BareMetalAdminClusterLoadBalancer.class}, tree="[0]")
    private Output</* @Nullable */ BareMetalAdminClusterLoadBalancer> loadBalancer;

    /**
     * @return Specifies the load balancer configuration.
     * Structure is documented below.
     * 
     */
    public Output<Optional<BareMetalAdminClusterLoadBalancer>> loadBalancer() {
        return Codegen.optional(this.loadBalancer);
    }
    /**
     * The object name of the Bare Metal Admin Cluster custom resource on the
     * associated admin cluster. This field is used to support conflicting
     * names when enrolling existing clusters to the API. When used as a part of
     * cluster enrollment, this field will differ from the ID in the resource
     * name. For new clusters, this field will match the user provided cluster ID
     * and be visible in the last component of the resource name. It is not
     * modifiable.
     * All users should use this name to access their cluster using gkectl or
     * kubectl and should expect to see the local name when viewing admin
     * cluster controller logs.
     * 
     */
    @Export(name="localName", refs={String.class}, tree="[0]")
    private Output<String> localName;

    /**
     * @return The object name of the Bare Metal Admin Cluster custom resource on the
     * associated admin cluster. This field is used to support conflicting
     * names when enrolling existing clusters to the API. When used as a part of
     * cluster enrollment, this field will differ from the ID in the resource
     * name. For new clusters, this field will match the user provided cluster ID
     * and be visible in the last component of the resource name. It is not
     * modifiable.
     * All users should use this name to access their cluster using gkectl or
     * kubectl and should expect to see the local name when viewing admin
     * cluster controller logs.
     * 
     */
    public Output<String> localName() {
        return this.localName;
    }
    /**
     * The location of the resource.
     * 
     * ***
     * 
     */
    @Export(name="location", refs={String.class}, tree="[0]")
    private Output<String> location;

    /**
     * @return The location of the resource.
     * 
     * ***
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    /**
     * Specifies the workload node configurations.
     * Structure is documented below.
     * 
     */
    @Export(name="maintenanceConfig", refs={BareMetalAdminClusterMaintenanceConfig.class}, tree="[0]")
    private Output</* @Nullable */ BareMetalAdminClusterMaintenanceConfig> maintenanceConfig;

    /**
     * @return Specifies the workload node configurations.
     * Structure is documented below.
     * 
     */
    public Output<Optional<BareMetalAdminClusterMaintenanceConfig>> maintenanceConfig() {
        return Codegen.optional(this.maintenanceConfig);
    }
    /**
     * The bare metal admin cluster name.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The bare metal admin cluster name.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Network configuration.
     * Structure is documented below.
     * 
     */
    @Export(name="networkConfig", refs={BareMetalAdminClusterNetworkConfig.class}, tree="[0]")
    private Output</* @Nullable */ BareMetalAdminClusterNetworkConfig> networkConfig;

    /**
     * @return Network configuration.
     * Structure is documented below.
     * 
     */
    public Output<Optional<BareMetalAdminClusterNetworkConfig>> networkConfig() {
        return Codegen.optional(this.networkConfig);
    }
    /**
     * Specifies the node access related settings for the bare metal user cluster.
     * Structure is documented below.
     * 
     */
    @Export(name="nodeAccessConfig", refs={BareMetalAdminClusterNodeAccessConfig.class}, tree="[0]")
    private Output</* @Nullable */ BareMetalAdminClusterNodeAccessConfig> nodeAccessConfig;

    /**
     * @return Specifies the node access related settings for the bare metal user cluster.
     * Structure is documented below.
     * 
     */
    public Output<Optional<BareMetalAdminClusterNodeAccessConfig>> nodeAccessConfig() {
        return Codegen.optional(this.nodeAccessConfig);
    }
    /**
     * Specifies the workload node configurations.
     * Structure is documented below.
     * 
     */
    @Export(name="nodeConfig", refs={BareMetalAdminClusterNodeConfig.class}, tree="[0]")
    private Output</* @Nullable */ BareMetalAdminClusterNodeConfig> nodeConfig;

    /**
     * @return Specifies the workload node configurations.
     * Structure is documented below.
     * 
     */
    public Output<Optional<BareMetalAdminClusterNodeConfig>> nodeConfig() {
        return Codegen.optional(this.nodeConfig);
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
     * Specifies the cluster proxy configuration.
     * Structure is documented below.
     * 
     */
    @Export(name="proxy", refs={BareMetalAdminClusterProxy.class}, tree="[0]")
    private Output</* @Nullable */ BareMetalAdminClusterProxy> proxy;

    /**
     * @return Specifies the cluster proxy configuration.
     * Structure is documented below.
     * 
     */
    public Output<Optional<BareMetalAdminClusterProxy>> proxy() {
        return Codegen.optional(this.proxy);
    }
    /**
     * If set, there are currently changes in flight to the Bare Metal Admin Cluster.
     * 
     */
    @Export(name="reconciling", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> reconciling;

    /**
     * @return If set, there are currently changes in flight to the Bare Metal Admin Cluster.
     * 
     */
    public Output<Boolean> reconciling() {
        return this.reconciling;
    }
    /**
     * Specifies the security related settings for the Bare Metal User Cluster.
     * Structure is documented below.
     * 
     */
    @Export(name="securityConfig", refs={BareMetalAdminClusterSecurityConfig.class}, tree="[0]")
    private Output</* @Nullable */ BareMetalAdminClusterSecurityConfig> securityConfig;

    /**
     * @return Specifies the security related settings for the Bare Metal User Cluster.
     * Structure is documented below.
     * 
     */
    public Output<Optional<BareMetalAdminClusterSecurityConfig>> securityConfig() {
        return Codegen.optional(this.securityConfig);
    }
    /**
     * (Output)
     * The lifecycle state of the condition.
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return (Output)
     * The lifecycle state of the condition.
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * (Output)
     * Specifies the detailed validation check status
     * Structure is documented below.
     * 
     */
    @Export(name="statuses", refs={List.class,BareMetalAdminClusterStatus.class}, tree="[0,1]")
    private Output<List<BareMetalAdminClusterStatus>> statuses;

    /**
     * @return (Output)
     * Specifies the detailed validation check status
     * Structure is documented below.
     * 
     */
    public Output<List<BareMetalAdminClusterStatus>> statuses() {
        return this.statuses;
    }
    /**
     * Specifies the cluster storage configuration.
     * Structure is documented below.
     * 
     */
    @Export(name="storage", refs={BareMetalAdminClusterStorage.class}, tree="[0]")
    private Output</* @Nullable */ BareMetalAdminClusterStorage> storage;

    /**
     * @return Specifies the cluster storage configuration.
     * Structure is documented below.
     * 
     */
    public Output<Optional<BareMetalAdminClusterStorage>> storage() {
        return Codegen.optional(this.storage);
    }
    /**
     * The unique identifier of the Bare Metal Admin Cluster.
     * 
     */
    @Export(name="uid", refs={String.class}, tree="[0]")
    private Output<String> uid;

    /**
     * @return The unique identifier of the Bare Metal Admin Cluster.
     * 
     */
    public Output<String> uid() {
        return this.uid;
    }
    /**
     * The time the cluster was last updated, in RFC3339 text format.
     * 
     */
    @Export(name="updateTime", refs={String.class}, tree="[0]")
    private Output<String> updateTime;

    /**
     * @return The time the cluster was last updated, in RFC3339 text format.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }
    /**
     * Specifies the security related settings for the Bare Metal Admin Cluster.
     * Structure is documented below.
     * 
     */
    @Export(name="validationChecks", refs={List.class,BareMetalAdminClusterValidationCheck.class}, tree="[0,1]")
    private Output<List<BareMetalAdminClusterValidationCheck>> validationChecks;

    /**
     * @return Specifies the security related settings for the Bare Metal Admin Cluster.
     * Structure is documented below.
     * 
     */
    public Output<List<BareMetalAdminClusterValidationCheck>> validationChecks() {
        return this.validationChecks;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public BareMetalAdminCluster(String name) {
        this(name, BareMetalAdminClusterArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BareMetalAdminCluster(String name, BareMetalAdminClusterArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BareMetalAdminCluster(String name, BareMetalAdminClusterArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:gkeonprem/bareMetalAdminCluster:BareMetalAdminCluster", name, args == null ? BareMetalAdminClusterArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private BareMetalAdminCluster(String name, Output<String> id, @Nullable BareMetalAdminClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:gkeonprem/bareMetalAdminCluster:BareMetalAdminCluster", name, state, makeResourceOptions(options, id));
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
    public static BareMetalAdminCluster get(String name, Output<String> id, @Nullable BareMetalAdminClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BareMetalAdminCluster(name, id, state, options);
    }
}
