// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.vmwareengine;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.vmwareengine.ClusterArgs;
import com.pulumi.gcp.vmwareengine.inputs.ClusterState;
import com.pulumi.gcp.vmwareengine.outputs.ClusterNodeTypeConfig;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A cluster in a private cloud.
 * 
 * To get more information about Cluster, see:
 * 
 * * [API documentation](https://cloud.google.com/vmware-engine/docs/reference/rest/v1/projects.locations.privateClouds.clusters)
 * 
 * ## Example Usage
 * 
 * ### Vmware Engine Cluster Basic
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.vmwareengine.Network;
 * import com.pulumi.gcp.vmwareengine.NetworkArgs;
 * import com.pulumi.gcp.vmwareengine.PrivateCloud;
 * import com.pulumi.gcp.vmwareengine.PrivateCloudArgs;
 * import com.pulumi.gcp.vmwareengine.inputs.PrivateCloudNetworkConfigArgs;
 * import com.pulumi.gcp.vmwareengine.inputs.PrivateCloudManagementClusterArgs;
 * import com.pulumi.gcp.vmwareengine.Cluster;
 * import com.pulumi.gcp.vmwareengine.ClusterArgs;
 * import com.pulumi.gcp.vmwareengine.inputs.ClusterNodeTypeConfigArgs;
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
 *         var cluster_nw = new Network("cluster-nw", NetworkArgs.builder()
 *             .name("pc-nw")
 *             .type("STANDARD")
 *             .location("global")
 *             .description("PC network description.")
 *             .build());
 * 
 *         var cluster_pc = new PrivateCloud("cluster-pc", PrivateCloudArgs.builder()
 *             .location("us-west1-a")
 *             .name("sample-pc")
 *             .description("Sample test PC.")
 *             .networkConfig(PrivateCloudNetworkConfigArgs.builder()
 *                 .managementCidr("192.168.30.0/24")
 *                 .vmwareEngineNetwork(cluster_nw.id())
 *                 .build())
 *             .managementCluster(PrivateCloudManagementClusterArgs.builder()
 *                 .clusterId("sample-mgmt-cluster")
 *                 .nodeTypeConfigs(PrivateCloudManagementClusterNodeTypeConfigArgs.builder()
 *                     .nodeTypeId("standard-72")
 *                     .nodeCount(3)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *         var vmw_engine_ext_cluster = new Cluster("vmw-engine-ext-cluster", ClusterArgs.builder()
 *             .name("ext-cluster")
 *             .parent(cluster_pc.id())
 *             .nodeTypeConfigs(ClusterNodeTypeConfigArgs.builder()
 *                 .nodeTypeId("standard-72")
 *                 .nodeCount(3)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Vmware Engine Cluster Full
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.vmwareengine.Network;
 * import com.pulumi.gcp.vmwareengine.NetworkArgs;
 * import com.pulumi.gcp.vmwareengine.PrivateCloud;
 * import com.pulumi.gcp.vmwareengine.PrivateCloudArgs;
 * import com.pulumi.gcp.vmwareengine.inputs.PrivateCloudNetworkConfigArgs;
 * import com.pulumi.gcp.vmwareengine.inputs.PrivateCloudManagementClusterArgs;
 * import com.pulumi.gcp.vmwareengine.Cluster;
 * import com.pulumi.gcp.vmwareengine.ClusterArgs;
 * import com.pulumi.gcp.vmwareengine.inputs.ClusterNodeTypeConfigArgs;
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
 *         var cluster_nw = new Network("cluster-nw", NetworkArgs.builder()
 *             .name("pc-nw")
 *             .type("STANDARD")
 *             .location("global")
 *             .description("PC network description.")
 *             .build());
 * 
 *         var cluster_pc = new PrivateCloud("cluster-pc", PrivateCloudArgs.builder()
 *             .location("us-west1-a")
 *             .name("sample-pc")
 *             .description("Sample test PC.")
 *             .networkConfig(PrivateCloudNetworkConfigArgs.builder()
 *                 .managementCidr("192.168.30.0/24")
 *                 .vmwareEngineNetwork(cluster_nw.id())
 *                 .build())
 *             .managementCluster(PrivateCloudManagementClusterArgs.builder()
 *                 .clusterId("sample-mgmt-cluster")
 *                 .nodeTypeConfigs(PrivateCloudManagementClusterNodeTypeConfigArgs.builder()
 *                     .nodeTypeId("standard-72")
 *                     .nodeCount(3)
 *                     .customCoreCount(32)
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *         var vmw_ext_cluster = new Cluster("vmw-ext-cluster", ClusterArgs.builder()
 *             .name("ext-cluster")
 *             .parent(cluster_pc.id())
 *             .nodeTypeConfigs(ClusterNodeTypeConfigArgs.builder()
 *                 .nodeTypeId("standard-72")
 *                 .nodeCount(3)
 *                 .customCoreCount(32)
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
 * Cluster can be imported using any of these accepted formats:
 * 
 * * `{{parent}}/clusters/{{name}}`
 * 
 * When using the `pulumi import` command, Cluster can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:vmwareengine/cluster:Cluster default {{parent}}/clusters/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:vmwareengine/cluster:Cluster")
public class Cluster extends com.pulumi.resources.CustomResource {
    /**
     * True if the cluster is a management cluster; false otherwise.
     * There can only be one management cluster in a private cloud and it has to be the first one.
     * 
     */
    @Export(name="management", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> management;

    /**
     * @return True if the cluster is a management cluster; false otherwise.
     * There can only be one management cluster in a private cloud and it has to be the first one.
     * 
     */
    public Output<Boolean> management() {
        return this.management;
    }
    /**
     * The ID of the Cluster.
     * 
     * ***
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The ID of the Cluster.
     * 
     * ***
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The map of cluster node types in this cluster,
     * where the key is canonical identifier of the node type (corresponds to the NodeType).
     * Structure is documented below.
     * 
     */
    @Export(name="nodeTypeConfigs", refs={List.class,ClusterNodeTypeConfig.class}, tree="[0,1]")
    private Output</* @Nullable */ List<ClusterNodeTypeConfig>> nodeTypeConfigs;

    /**
     * @return The map of cluster node types in this cluster,
     * where the key is canonical identifier of the node type (corresponds to the NodeType).
     * Structure is documented below.
     * 
     */
    public Output<Optional<List<ClusterNodeTypeConfig>>> nodeTypeConfigs() {
        return Codegen.optional(this.nodeTypeConfigs);
    }
    /**
     * The resource name of the private cloud to create a new cluster in.
     * Resource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names.
     * For example: projects/my-project/locations/us-west1-a/privateClouds/my-cloud
     * 
     */
    @Export(name="parent", refs={String.class}, tree="[0]")
    private Output<String> parent;

    /**
     * @return The resource name of the private cloud to create a new cluster in.
     * Resource names are schemeless URIs that follow the conventions in https://cloud.google.com/apis/design/resource_names.
     * For example: projects/my-project/locations/us-west1-a/privateClouds/my-cloud
     * 
     */
    public Output<String> parent() {
        return this.parent;
    }
    /**
     * State of the Cluster.
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return State of the Cluster.
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * System-generated unique identifier for the resource.
     * 
     */
    @Export(name="uid", refs={String.class}, tree="[0]")
    private Output<String> uid;

    /**
     * @return System-generated unique identifier for the resource.
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
        super("gcp:vmwareengine/cluster:Cluster", name, args == null ? ClusterArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Cluster(String name, Output<String> id, @Nullable ClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:vmwareengine/cluster:Cluster", name, state, makeResourceOptions(options, id));
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
    public static Cluster get(String name, Output<String> id, @Nullable ClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Cluster(name, id, state, options);
    }
}
