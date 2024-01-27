// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.gkeonprem;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.gkeonprem.BareMetalNodePoolArgs;
import com.pulumi.gcp.gkeonprem.inputs.BareMetalNodePoolState;
import com.pulumi.gcp.gkeonprem.outputs.BareMetalNodePoolNodePoolConfig;
import com.pulumi.gcp.gkeonprem.outputs.BareMetalNodePoolStatus;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A Google Bare Metal Node Pool.
 * 
 * ## Example Usage
 * ### Gkeonprem Bare Metal Node Pool Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.gkeonprem.BareMetalCluster;
 * import com.pulumi.gcp.gkeonprem.BareMetalClusterArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalClusterNetworkConfigArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalClusterNetworkConfigIslandModeCidrArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalClusterControlPlaneArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalClusterControlPlaneControlPlaneNodePoolConfigArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalClusterLoadBalancerArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalClusterLoadBalancerPortConfigArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalClusterLoadBalancerVipConfigArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalClusterLoadBalancerMetalLbConfigArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalClusterStorageArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalClusterStorageLvpShareConfigArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalClusterStorageLvpShareConfigLvpConfigArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalClusterStorageLvpNodeMountsConfigArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalClusterSecurityConfigArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalClusterSecurityConfigAuthorizationArgs;
 * import com.pulumi.gcp.gkeonprem.BareMetalNodePool;
 * import com.pulumi.gcp.gkeonprem.BareMetalNodePoolArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalNodePoolNodePoolConfigArgs;
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
 *         var default_basic = new BareMetalCluster(&#34;default-basic&#34;, BareMetalClusterArgs.builder()        
 *             .location(&#34;us-west1&#34;)
 *             .adminClusterMembership(&#34;projects/870316890899/locations/global/memberships/gkeonprem-terraform-test&#34;)
 *             .bareMetalVersion(&#34;1.12.3&#34;)
 *             .networkConfig(BareMetalClusterNetworkConfigArgs.builder()
 *                 .islandModeCidr(BareMetalClusterNetworkConfigIslandModeCidrArgs.builder()
 *                     .serviceAddressCidrBlocks(&#34;172.26.0.0/16&#34;)
 *                     .podAddressCidrBlocks(&#34;10.240.0.0/13&#34;)
 *                     .build())
 *                 .build())
 *             .controlPlane(BareMetalClusterControlPlaneArgs.builder()
 *                 .controlPlaneNodePoolConfig(BareMetalClusterControlPlaneControlPlaneNodePoolConfigArgs.builder()
 *                     .nodePoolConfig(BareMetalClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigArgs.builder()
 *                         .labels()
 *                         .operatingSystem(&#34;LINUX&#34;)
 *                         .nodeConfigs(BareMetalClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigNodeConfigArgs.builder()
 *                             .labels()
 *                             .nodeIp(&#34;10.200.0.9&#34;)
 *                             .build())
 *                         .build())
 *                     .build())
 *                 .build())
 *             .loadBalancer(BareMetalClusterLoadBalancerArgs.builder()
 *                 .portConfig(BareMetalClusterLoadBalancerPortConfigArgs.builder()
 *                     .controlPlaneLoadBalancerPort(443)
 *                     .build())
 *                 .vipConfig(BareMetalClusterLoadBalancerVipConfigArgs.builder()
 *                     .controlPlaneVip(&#34;10.200.0.13&#34;)
 *                     .ingressVip(&#34;10.200.0.14&#34;)
 *                     .build())
 *                 .metalLbConfig(BareMetalClusterLoadBalancerMetalLbConfigArgs.builder()
 *                     .addressPools(BareMetalClusterLoadBalancerMetalLbConfigAddressPoolArgs.builder()
 *                         .pool(&#34;pool1&#34;)
 *                         .addresses(                        
 *                             &#34;10.200.0.14/32&#34;,
 *                             &#34;10.200.0.15/32&#34;,
 *                             &#34;10.200.0.16/32&#34;,
 *                             &#34;10.200.0.17/32&#34;,
 *                             &#34;10.200.0.18/32&#34;,
 *                             &#34;fd00:1::f/128&#34;,
 *                             &#34;fd00:1::10/128&#34;,
 *                             &#34;fd00:1::11/128&#34;,
 *                             &#34;fd00:1::12/128&#34;)
 *                         .build())
 *                     .build())
 *                 .build())
 *             .storage(BareMetalClusterStorageArgs.builder()
 *                 .lvpShareConfig(BareMetalClusterStorageLvpShareConfigArgs.builder()
 *                     .lvpConfig(BareMetalClusterStorageLvpShareConfigLvpConfigArgs.builder()
 *                         .path(&#34;/mnt/localpv-share&#34;)
 *                         .storageClass(&#34;local-shared&#34;)
 *                         .build())
 *                     .sharedPathPvCount(5)
 *                     .build())
 *                 .lvpNodeMountsConfig(BareMetalClusterStorageLvpNodeMountsConfigArgs.builder()
 *                     .path(&#34;/mnt/localpv-disk&#34;)
 *                     .storageClass(&#34;local-disks&#34;)
 *                     .build())
 *                 .build())
 *             .securityConfig(BareMetalClusterSecurityConfigArgs.builder()
 *                 .authorization(BareMetalClusterSecurityConfigAuthorizationArgs.builder()
 *                     .adminUsers(BareMetalClusterSecurityConfigAuthorizationAdminUserArgs.builder()
 *                         .username(&#34;admin@hashicorptest.com&#34;)
 *                         .build())
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *         var nodepool_basic = new BareMetalNodePool(&#34;nodepool-basic&#34;, BareMetalNodePoolArgs.builder()        
 *             .bareMetalCluster(default_basic.name())
 *             .location(&#34;us-west1&#34;)
 *             .nodePoolConfig(BareMetalNodePoolNodePoolConfigArgs.builder()
 *                 .operatingSystem(&#34;LINUX&#34;)
 *                 .nodeConfigs(BareMetalNodePoolNodePoolConfigNodeConfigArgs.builder()
 *                     .nodeIp(&#34;10.200.0.11&#34;)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Gkeonprem Bare Metal Node Pool Full
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.gkeonprem.BareMetalCluster;
 * import com.pulumi.gcp.gkeonprem.BareMetalClusterArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalClusterNetworkConfigArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalClusterNetworkConfigIslandModeCidrArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalClusterControlPlaneArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalClusterControlPlaneControlPlaneNodePoolConfigArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalClusterLoadBalancerArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalClusterLoadBalancerPortConfigArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalClusterLoadBalancerVipConfigArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalClusterLoadBalancerMetalLbConfigArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalClusterStorageArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalClusterStorageLvpShareConfigArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalClusterStorageLvpShareConfigLvpConfigArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalClusterStorageLvpNodeMountsConfigArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalClusterSecurityConfigArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalClusterSecurityConfigAuthorizationArgs;
 * import com.pulumi.gcp.gkeonprem.BareMetalNodePool;
 * import com.pulumi.gcp.gkeonprem.BareMetalNodePoolArgs;
 * import com.pulumi.gcp.gkeonprem.inputs.BareMetalNodePoolNodePoolConfigArgs;
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
 *         var default_full = new BareMetalCluster(&#34;default-full&#34;, BareMetalClusterArgs.builder()        
 *             .location(&#34;us-west1&#34;)
 *             .adminClusterMembership(&#34;projects/870316890899/locations/global/memberships/gkeonprem-terraform-test&#34;)
 *             .bareMetalVersion(&#34;1.12.3&#34;)
 *             .networkConfig(BareMetalClusterNetworkConfigArgs.builder()
 *                 .islandModeCidr(BareMetalClusterNetworkConfigIslandModeCidrArgs.builder()
 *                     .serviceAddressCidrBlocks(&#34;172.26.0.0/16&#34;)
 *                     .podAddressCidrBlocks(&#34;10.240.0.0/13&#34;)
 *                     .build())
 *                 .build())
 *             .controlPlane(BareMetalClusterControlPlaneArgs.builder()
 *                 .controlPlaneNodePoolConfig(BareMetalClusterControlPlaneControlPlaneNodePoolConfigArgs.builder()
 *                     .nodePoolConfig(BareMetalClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigArgs.builder()
 *                         .labels()
 *                         .operatingSystem(&#34;LINUX&#34;)
 *                         .nodeConfigs(BareMetalClusterControlPlaneControlPlaneNodePoolConfigNodePoolConfigNodeConfigArgs.builder()
 *                             .labels()
 *                             .nodeIp(&#34;10.200.0.9&#34;)
 *                             .build())
 *                         .build())
 *                     .build())
 *                 .build())
 *             .loadBalancer(BareMetalClusterLoadBalancerArgs.builder()
 *                 .portConfig(BareMetalClusterLoadBalancerPortConfigArgs.builder()
 *                     .controlPlaneLoadBalancerPort(443)
 *                     .build())
 *                 .vipConfig(BareMetalClusterLoadBalancerVipConfigArgs.builder()
 *                     .controlPlaneVip(&#34;10.200.0.13&#34;)
 *                     .ingressVip(&#34;10.200.0.14&#34;)
 *                     .build())
 *                 .metalLbConfig(BareMetalClusterLoadBalancerMetalLbConfigArgs.builder()
 *                     .addressPools(BareMetalClusterLoadBalancerMetalLbConfigAddressPoolArgs.builder()
 *                         .pool(&#34;pool1&#34;)
 *                         .addresses(                        
 *                             &#34;10.200.0.14/32&#34;,
 *                             &#34;10.200.0.15/32&#34;,
 *                             &#34;10.200.0.16/32&#34;,
 *                             &#34;10.200.0.17/32&#34;,
 *                             &#34;10.200.0.18/32&#34;,
 *                             &#34;fd00:1::f/128&#34;,
 *                             &#34;fd00:1::10/128&#34;,
 *                             &#34;fd00:1::11/128&#34;,
 *                             &#34;fd00:1::12/128&#34;)
 *                         .build())
 *                     .build())
 *                 .build())
 *             .storage(BareMetalClusterStorageArgs.builder()
 *                 .lvpShareConfig(BareMetalClusterStorageLvpShareConfigArgs.builder()
 *                     .lvpConfig(BareMetalClusterStorageLvpShareConfigLvpConfigArgs.builder()
 *                         .path(&#34;/mnt/localpv-share&#34;)
 *                         .storageClass(&#34;local-shared&#34;)
 *                         .build())
 *                     .sharedPathPvCount(5)
 *                     .build())
 *                 .lvpNodeMountsConfig(BareMetalClusterStorageLvpNodeMountsConfigArgs.builder()
 *                     .path(&#34;/mnt/localpv-disk&#34;)
 *                     .storageClass(&#34;local-disks&#34;)
 *                     .build())
 *                 .build())
 *             .securityConfig(BareMetalClusterSecurityConfigArgs.builder()
 *                 .authorization(BareMetalClusterSecurityConfigAuthorizationArgs.builder()
 *                     .adminUsers(BareMetalClusterSecurityConfigAuthorizationAdminUserArgs.builder()
 *                         .username(&#34;admin@hashicorptest.com&#34;)
 *                         .build())
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *         var nodepool_full = new BareMetalNodePool(&#34;nodepool-full&#34;, BareMetalNodePoolArgs.builder()        
 *             .displayName(&#34;test-name&#34;)
 *             .bareMetalCluster(default_full.name())
 *             .location(&#34;us-west1&#34;)
 *             .annotations()
 *             .nodePoolConfig(BareMetalNodePoolNodePoolConfigArgs.builder()
 *                 .operatingSystem(&#34;LINUX&#34;)
 *                 .labels()
 *                 .nodeConfigs(BareMetalNodePoolNodePoolConfigNodeConfigArgs.builder()
 *                     .nodeIp(&#34;10.200.0.11&#34;)
 *                     .labels()
 *                     .build())
 *                 .taints(BareMetalNodePoolNodePoolConfigTaintArgs.builder()
 *                     .key(&#34;test-key&#34;)
 *                     .value(&#34;test-value&#34;)
 *                     .effect(&#34;NO_EXECUTE&#34;)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * BareMetalNodePool can be imported using any of these accepted formats* `projects/{{project}}/locations/{{location}}/bareMetalClusters/{{bare_metal_cluster}}/bareMetalNodePools/{{name}}` * `{{project}}/{{location}}/{{bare_metal_cluster}}/{{name}}` * `{{location}}/{{bare_metal_cluster}}/{{name}}` When using the `pulumi import` command, BareMetalNodePool can be imported using one of the formats above. For example
 * 
 * ```sh
 *  $ pulumi import gcp:gkeonprem/bareMetalNodePool:BareMetalNodePool default projects/{{project}}/locations/{{location}}/bareMetalClusters/{{bare_metal_cluster}}/bareMetalNodePools/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:gkeonprem/bareMetalNodePool:BareMetalNodePool default {{project}}/{{location}}/{{bare_metal_cluster}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:gkeonprem/bareMetalNodePool:BareMetalNodePool default {{location}}/{{bare_metal_cluster}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:gkeonprem/bareMetalNodePool:BareMetalNodePool")
public class BareMetalNodePool extends com.pulumi.resources.CustomResource {
    /**
     * Annotations on the Bare Metal Node Pool.
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
     * @return Annotations on the Bare Metal Node Pool.
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
     * The cluster this node pool belongs to.
     * 
     */
    @Export(name="bareMetalCluster", refs={String.class}, tree="[0]")
    private Output<String> bareMetalCluster;

    /**
     * @return The cluster this node pool belongs to.
     * 
     */
    public Output<String> bareMetalCluster() {
        return this.bareMetalCluster;
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
     * The display name for the Bare Metal Node Pool.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> displayName;

    /**
     * @return The display name for the Bare Metal Node Pool.
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
     * The location of the resource.
     * 
     */
    @Export(name="location", refs={String.class}, tree="[0]")
    private Output<String> location;

    /**
     * @return The location of the resource.
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    /**
     * The bare metal node pool name.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The bare metal node pool name.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Node pool configuration.
     * Structure is documented below.
     * 
     */
    @Export(name="nodePoolConfig", refs={BareMetalNodePoolNodePoolConfig.class}, tree="[0]")
    private Output<BareMetalNodePoolNodePoolConfig> nodePoolConfig;

    /**
     * @return Node pool configuration.
     * Structure is documented below.
     * 
     */
    public Output<BareMetalNodePoolNodePoolConfig> nodePoolConfig() {
        return this.nodePoolConfig;
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
     * If set, there are currently changes in flight to the Bare Metal User Cluster.
     * 
     */
    @Export(name="reconciling", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> reconciling;

    /**
     * @return If set, there are currently changes in flight to the Bare Metal User Cluster.
     * 
     */
    public Output<Boolean> reconciling() {
        return this.reconciling;
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
     * Specifies detailed node pool status.
     * Structure is documented below.
     * 
     */
    @Export(name="statuses", refs={List.class,BareMetalNodePoolStatus.class}, tree="[0,1]")
    private Output<List<BareMetalNodePoolStatus>> statuses;

    /**
     * @return Specifies detailed node pool status.
     * Structure is documented below.
     * 
     */
    public Output<List<BareMetalNodePoolStatus>> statuses() {
        return this.statuses;
    }
    /**
     * The unique identifier of the Bare Metal Node Pool.
     * 
     */
    @Export(name="uid", refs={String.class}, tree="[0]")
    private Output<String> uid;

    /**
     * @return The unique identifier of the Bare Metal Node Pool.
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public BareMetalNodePool(String name) {
        this(name, BareMetalNodePoolArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BareMetalNodePool(String name, BareMetalNodePoolArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BareMetalNodePool(String name, BareMetalNodePoolArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:gkeonprem/bareMetalNodePool:BareMetalNodePool", name, args == null ? BareMetalNodePoolArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private BareMetalNodePool(String name, Output<String> id, @Nullable BareMetalNodePoolState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:gkeonprem/bareMetalNodePool:BareMetalNodePool", name, state, makeResourceOptions(options, id));
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
    public static BareMetalNodePool get(String name, Output<String> id, @Nullable BareMetalNodePoolState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BareMetalNodePool(name, id, state, options);
    }
}
