// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.container;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.container.NodePoolArgs;
import com.pulumi.gcp.container.inputs.NodePoolState;
import com.pulumi.gcp.container.outputs.NodePoolAutoscaling;
import com.pulumi.gcp.container.outputs.NodePoolManagement;
import com.pulumi.gcp.container.outputs.NodePoolNetworkConfig;
import com.pulumi.gcp.container.outputs.NodePoolNodeConfig;
import com.pulumi.gcp.container.outputs.NodePoolPlacementPolicy;
import com.pulumi.gcp.container.outputs.NodePoolQueuedProvisioning;
import com.pulumi.gcp.container.outputs.NodePoolUpgradeSettings;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a node pool in a Google Kubernetes Engine (GKE) cluster separately from
 * the cluster control plane. For more information see [the official documentation](https://cloud.google.com/container-engine/docs/node-pools)
 * and [the API reference](https://cloud.google.com/kubernetes-engine/docs/reference/rest/v1beta1/projects.locations.clusters.nodePools).
 * 
 * ## Example Usage
 * ### Using A Separately Managed Node Pool (Recommended)
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.serviceaccount.Account;
 * import com.pulumi.gcp.serviceaccount.AccountArgs;
 * import com.pulumi.gcp.container.Cluster;
 * import com.pulumi.gcp.container.ClusterArgs;
 * import com.pulumi.gcp.container.NodePool;
 * import com.pulumi.gcp.container.NodePoolArgs;
 * import com.pulumi.gcp.container.inputs.NodePoolNodeConfigArgs;
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
 *         var default_ = new Account(&#34;default&#34;, AccountArgs.builder()        
 *             .accountId(&#34;service-account-id&#34;)
 *             .displayName(&#34;Service Account&#34;)
 *             .build());
 * 
 *         var primary = new Cluster(&#34;primary&#34;, ClusterArgs.builder()        
 *             .location(&#34;us-central1&#34;)
 *             .removeDefaultNodePool(true)
 *             .initialNodeCount(1)
 *             .build());
 * 
 *         var primaryPreemptibleNodes = new NodePool(&#34;primaryPreemptibleNodes&#34;, NodePoolArgs.builder()        
 *             .cluster(primary.id())
 *             .nodeCount(1)
 *             .nodeConfig(NodePoolNodeConfigArgs.builder()
 *                 .preemptible(true)
 *                 .machineType(&#34;e2-medium&#34;)
 *                 .serviceAccount(default_.email())
 *                 .oauthScopes(&#34;https://www.googleapis.com/auth/cloud-platform&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### 2 Node Pools, 1 Separately Managed + The Default Node Pool
 * 
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.serviceaccount.Account;
 * import com.pulumi.gcp.serviceaccount.AccountArgs;
 * import com.pulumi.gcp.container.Cluster;
 * import com.pulumi.gcp.container.ClusterArgs;
 * import com.pulumi.gcp.container.inputs.ClusterNodeConfigArgs;
 * import com.pulumi.gcp.container.NodePool;
 * import com.pulumi.gcp.container.NodePoolArgs;
 * import com.pulumi.gcp.container.inputs.NodePoolNodeConfigArgs;
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
 *         var default_ = new Account(&#34;default&#34;, AccountArgs.builder()        
 *             .accountId(&#34;service-account-id&#34;)
 *             .displayName(&#34;Service Account&#34;)
 *             .build());
 * 
 *         var primary = new Cluster(&#34;primary&#34;, ClusterArgs.builder()        
 *             .location(&#34;us-central1-a&#34;)
 *             .initialNodeCount(3)
 *             .nodeLocations(&#34;us-central1-c&#34;)
 *             .nodeConfig(ClusterNodeConfigArgs.builder()
 *                 .serviceAccount(default_.email())
 *                 .oauthScopes(&#34;https://www.googleapis.com/auth/cloud-platform&#34;)
 *                 .guestAccelerators(ClusterNodeConfigGuestAcceleratorArgs.builder()
 *                     .type(&#34;nvidia-tesla-k80&#34;)
 *                     .count(1)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *         var np = new NodePool(&#34;np&#34;, NodePoolArgs.builder()        
 *             .cluster(primary.id())
 *             .nodeConfig(NodePoolNodeConfigArgs.builder()
 *                 .machineType(&#34;e2-medium&#34;)
 *                 .serviceAccount(default_.email())
 *                 .oauthScopes(&#34;https://www.googleapis.com/auth/cloud-platform&#34;)
 *                 .build())
 *             .timeouts(%!v(PANIC=Format method: runtime error: invalid memory address or nil pointer dereference))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Node pools can be imported using the `project`, `location`, `cluster` and `name`. If the project is omitted, the project value in the provider configuration will be used. Examples* `{{project_id}}/{{location}}/{{cluster_id}}/{{pool_id}}` * `{{location}}/{{cluster_id}}/{{pool_id}}` When using the `pulumi import` command, node pools can be imported using one of the formats above. For example
 * 
 * ```sh
 *  $ pulumi import gcp:container/nodePool:NodePool default {{project_id}}/{{location}}/{{cluster_id}}/{{pool_id}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:container/nodePool:NodePool default {{location}}/{{cluster_id}}/{{pool_id}}
 * ```
 * 
 */
@ResourceType(type="gcp:container/nodePool:NodePool")
public class NodePool extends com.pulumi.resources.CustomResource {
    /**
     * Configuration required by cluster autoscaler to adjust
     * the size of the node pool to the current cluster usage. Structure is documented below.
     * 
     */
    @Export(name="autoscaling", refs={NodePoolAutoscaling.class}, tree="[0]")
    private Output</* @Nullable */ NodePoolAutoscaling> autoscaling;

    /**
     * @return Configuration required by cluster autoscaler to adjust
     * the size of the node pool to the current cluster usage. Structure is documented below.
     * 
     */
    public Output<Optional<NodePoolAutoscaling>> autoscaling() {
        return Codegen.optional(this.autoscaling);
    }
    /**
     * The cluster to create the node pool for. Cluster must be present in `location` provided for clusters. May be specified in the format `projects/{{project}}/locations/{{location}}/clusters/{{cluster}}` or as just the name of the cluster.
     * 
     * ***
     * 
     */
    @Export(name="cluster", refs={String.class}, tree="[0]")
    private Output<String> cluster;

    /**
     * @return The cluster to create the node pool for. Cluster must be present in `location` provided for clusters. May be specified in the format `projects/{{project}}/locations/{{location}}/clusters/{{cluster}}` or as just the name of the cluster.
     * 
     * ***
     * 
     */
    public Output<String> cluster() {
        return this.cluster;
    }
    /**
     * The initial number of nodes for the pool. In
     * regional or multi-zonal clusters, this is the number of nodes per zone. Changing
     * this will force recreation of the resource. WARNING: Resizing your node pool manually
     * may change this value in your existing cluster, which will trigger destruction
     * and recreation on the next provider run (to rectify the discrepancy).  If you don&#39;t
     * need this value, don&#39;t set it.  If you do need it, you can use a lifecycle block to
     * ignore subsqeuent changes to this field.
     * 
     */
    @Export(name="initialNodeCount", refs={Integer.class}, tree="[0]")
    private Output<Integer> initialNodeCount;

    /**
     * @return The initial number of nodes for the pool. In
     * regional or multi-zonal clusters, this is the number of nodes per zone. Changing
     * this will force recreation of the resource. WARNING: Resizing your node pool manually
     * may change this value in your existing cluster, which will trigger destruction
     * and recreation on the next provider run (to rectify the discrepancy).  If you don&#39;t
     * need this value, don&#39;t set it.  If you do need it, you can use a lifecycle block to
     * ignore subsqeuent changes to this field.
     * 
     */
    public Output<Integer> initialNodeCount() {
        return this.initialNodeCount;
    }
    /**
     * The resource URLs of the managed instance groups associated with this node pool.
     * 
     */
    @Export(name="instanceGroupUrls", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> instanceGroupUrls;

    /**
     * @return The resource URLs of the managed instance groups associated with this node pool.
     * 
     */
    public Output<List<String>> instanceGroupUrls() {
        return this.instanceGroupUrls;
    }
    /**
     * The location (region or zone) of the cluster.
     * 
     * ***
     * 
     */
    @Export(name="location", refs={String.class}, tree="[0]")
    private Output<String> location;

    /**
     * @return The location (region or zone) of the cluster.
     * 
     * ***
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    /**
     * List of instance group URLs which have been assigned to this node pool.
     * 
     */
    @Export(name="managedInstanceGroupUrls", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> managedInstanceGroupUrls;

    /**
     * @return List of instance group URLs which have been assigned to this node pool.
     * 
     */
    public Output<List<String>> managedInstanceGroupUrls() {
        return this.managedInstanceGroupUrls;
    }
    /**
     * Node management configuration, wherein auto-repair and
     * auto-upgrade is configured. Structure is documented below.
     * 
     */
    @Export(name="management", refs={NodePoolManagement.class}, tree="[0]")
    private Output<NodePoolManagement> management;

    /**
     * @return Node management configuration, wherein auto-repair and
     * auto-upgrade is configured. Structure is documented below.
     * 
     */
    public Output<NodePoolManagement> management() {
        return this.management;
    }
    /**
     * The maximum number of pods per node in this node pool.
     * Note that this does not work on node pools which are &#34;route-based&#34; - that is, node
     * pools belonging to clusters that do not have IP Aliasing enabled.
     * See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/flexible-pod-cidr)
     * for more information.
     * 
     */
    @Export(name="maxPodsPerNode", refs={Integer.class}, tree="[0]")
    private Output<Integer> maxPodsPerNode;

    /**
     * @return The maximum number of pods per node in this node pool.
     * Note that this does not work on node pools which are &#34;route-based&#34; - that is, node
     * pools belonging to clusters that do not have IP Aliasing enabled.
     * See the [official documentation](https://cloud.google.com/kubernetes-engine/docs/how-to/flexible-pod-cidr)
     * for more information.
     * 
     */
    public Output<Integer> maxPodsPerNode() {
        return this.maxPodsPerNode;
    }
    /**
     * The name of the node pool. If left blank, the provider will
     * auto-generate a unique name.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the node pool. If left blank, the provider will
     * auto-generate a unique name.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Creates a unique name for the node pool beginning
     * with the specified prefix. Conflicts with `name`.
     * 
     */
    @Export(name="namePrefix", refs={String.class}, tree="[0]")
    private Output<String> namePrefix;

    /**
     * @return Creates a unique name for the node pool beginning
     * with the specified prefix. Conflicts with `name`.
     * 
     */
    public Output<String> namePrefix() {
        return this.namePrefix;
    }
    /**
     * The network configuration of the pool. Such as
     * configuration for [Adding Pod IP address ranges](https://cloud.google.com/kubernetes-engine/docs/how-to/multi-pod-cidr)) to the node pool. Or enabling private nodes. Structure is
     * documented below
     * 
     */
    @Export(name="networkConfig", refs={NodePoolNetworkConfig.class}, tree="[0]")
    private Output<NodePoolNetworkConfig> networkConfig;

    /**
     * @return The network configuration of the pool. Such as
     * configuration for [Adding Pod IP address ranges](https://cloud.google.com/kubernetes-engine/docs/how-to/multi-pod-cidr)) to the node pool. Or enabling private nodes. Structure is
     * documented below
     * 
     */
    public Output<NodePoolNetworkConfig> networkConfig() {
        return this.networkConfig;
    }
    /**
     * Parameters used in creating the node pool. See
     * gcp.container.Cluster for schema.
     * 
     */
    @Export(name="nodeConfig", refs={NodePoolNodeConfig.class}, tree="[0]")
    private Output<NodePoolNodeConfig> nodeConfig;

    /**
     * @return Parameters used in creating the node pool. See
     * gcp.container.Cluster for schema.
     * 
     */
    public Output<NodePoolNodeConfig> nodeConfig() {
        return this.nodeConfig;
    }
    /**
     * The number of nodes per instance group. This field can be used to
     * update the number of nodes per instance group but should not be used alongside `autoscaling`.
     * 
     */
    @Export(name="nodeCount", refs={Integer.class}, tree="[0]")
    private Output<Integer> nodeCount;

    /**
     * @return The number of nodes per instance group. This field can be used to
     * update the number of nodes per instance group but should not be used alongside `autoscaling`.
     * 
     */
    public Output<Integer> nodeCount() {
        return this.nodeCount;
    }
    /**
     * The list of zones in which the node pool&#39;s nodes should be located. Nodes must
     * be in the region of their regional cluster or in the same region as their
     * cluster&#39;s zone for zonal clusters. If unspecified, the cluster-level
     * `node_locations` will be used.
     * 
     * &gt; Note: `node_locations` will not revert to the cluster&#39;s default set of zones
     * upon being unset. You must manually reconcile the list of zones with your
     * cluster.
     * 
     */
    @Export(name="nodeLocations", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> nodeLocations;

    /**
     * @return The list of zones in which the node pool&#39;s nodes should be located. Nodes must
     * be in the region of their regional cluster or in the same region as their
     * cluster&#39;s zone for zonal clusters. If unspecified, the cluster-level
     * `node_locations` will be used.
     * 
     * &gt; Note: `node_locations` will not revert to the cluster&#39;s default set of zones
     * upon being unset. You must manually reconcile the list of zones with your
     * cluster.
     * 
     */
    public Output<List<String>> nodeLocations() {
        return this.nodeLocations;
    }
    @Export(name="operation", refs={String.class}, tree="[0]")
    private Output<String> operation;

    public Output<String> operation() {
        return this.operation;
    }
    /**
     * Specifies a custom placement policy for the
     * nodes.
     * 
     */
    @Export(name="placementPolicy", refs={NodePoolPlacementPolicy.class}, tree="[0]")
    private Output</* @Nullable */ NodePoolPlacementPolicy> placementPolicy;

    /**
     * @return Specifies a custom placement policy for the
     * nodes.
     * 
     */
    public Output<Optional<NodePoolPlacementPolicy>> placementPolicy() {
        return Codegen.optional(this.placementPolicy);
    }
    /**
     * The ID of the project in which to create the node pool. If blank,
     * the provider-configured project will be used.
     * 
     */
    @Export(name="project", refs={String.class}, tree="[0]")
    private Output<String> project;

    /**
     * @return The ID of the project in which to create the node pool. If blank,
     * the provider-configured project will be used.
     * 
     */
    public Output<String> project() {
        return this.project;
    }
    /**
     * Specifies node pool-level settings of queued provisioning.
     * Structure is documented below.
     * 
     * &lt;a name=&#34;nested_autoscaling&#34;&gt;&lt;/a&gt;The `autoscaling` block supports (either total or per zone limits are required):
     * 
     */
    @Export(name="queuedProvisioning", refs={NodePoolQueuedProvisioning.class}, tree="[0]")
    private Output</* @Nullable */ NodePoolQueuedProvisioning> queuedProvisioning;

    /**
     * @return Specifies node pool-level settings of queued provisioning.
     * Structure is documented below.
     * 
     * &lt;a name=&#34;nested_autoscaling&#34;&gt;&lt;/a&gt;The `autoscaling` block supports (either total or per zone limits are required):
     * 
     */
    public Output<Optional<NodePoolQueuedProvisioning>> queuedProvisioning() {
        return Codegen.optional(this.queuedProvisioning);
    }
    /**
     * Specify node upgrade settings to change how GKE upgrades nodes.
     * The maximum number of nodes upgraded simultaneously is limited to 20. Structure is documented below.
     * 
     */
    @Export(name="upgradeSettings", refs={NodePoolUpgradeSettings.class}, tree="[0]")
    private Output<NodePoolUpgradeSettings> upgradeSettings;

    /**
     * @return Specify node upgrade settings to change how GKE upgrades nodes.
     * The maximum number of nodes upgraded simultaneously is limited to 20. Structure is documented below.
     * 
     */
    public Output<NodePoolUpgradeSettings> upgradeSettings() {
        return this.upgradeSettings;
    }
    /**
     * The Kubernetes version for the nodes in this pool. Note that if this field
     * and `auto_upgrade` are both specified, they will fight each other for what the node version should
     * be, so setting both is highly discouraged. While a fuzzy version can be specified, it&#39;s
     * recommended that you specify explicit versions as the provider will see spurious diffs
     * when fuzzy versions are used. See the `gcp.container.getEngineVersions` data source&#39;s
     * `version_prefix` field to approximate fuzzy versions in a provider-compatible way.
     * 
     */
    @Export(name="version", refs={String.class}, tree="[0]")
    private Output<String> version;

    /**
     * @return The Kubernetes version for the nodes in this pool. Note that if this field
     * and `auto_upgrade` are both specified, they will fight each other for what the node version should
     * be, so setting both is highly discouraged. While a fuzzy version can be specified, it&#39;s
     * recommended that you specify explicit versions as the provider will see spurious diffs
     * when fuzzy versions are used. See the `gcp.container.getEngineVersions` data source&#39;s
     * `version_prefix` field to approximate fuzzy versions in a provider-compatible way.
     * 
     */
    public Output<String> version() {
        return this.version;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public NodePool(String name) {
        this(name, NodePoolArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public NodePool(String name, NodePoolArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public NodePool(String name, NodePoolArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:container/nodePool:NodePool", name, args == null ? NodePoolArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private NodePool(String name, Output<String> id, @Nullable NodePoolState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:container/nodePool:NodePool", name, state, makeResourceOptions(options, id));
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
    public static NodePool get(String name, Output<String> id, @Nullable NodePoolState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new NodePool(name, id, state, options);
    }
}
