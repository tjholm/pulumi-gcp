// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.vmwareengine;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.vmwareengine.NetworkPeeringArgs;
import com.pulumi.gcp.vmwareengine.inputs.NetworkPeeringState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Represents a network peering resource. Network peerings are global resources.
 * 
 * To get more information about NetworkPeering, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/networks/addPeering)
 * 
 * ## Example Usage
 * 
 * ### Vmware Engine Network Peering Ven
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
 * import com.pulumi.gcp.vmwareengine.NetworkPeering;
 * import com.pulumi.gcp.vmwareengine.NetworkPeeringArgs;
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
 *         var network_peering_nw = new Network("network-peering-nw", NetworkArgs.builder()
 *             .name("default-np-nw")
 *             .location("global")
 *             .type("STANDARD")
 *             .build());
 * 
 *         var network_peering_peer_nw = new Network("network-peering-peer-nw", NetworkArgs.builder()
 *             .name("peer-np-nw")
 *             .location("global")
 *             .type("STANDARD")
 *             .build());
 * 
 *         var vmw_engine_network_peering = new NetworkPeering("vmw-engine-network-peering", NetworkPeeringArgs.builder()
 *             .name("sample-network-peering")
 *             .description("Sample description")
 *             .vmwareEngineNetwork(network_peering_nw.id())
 *             .peerNetwork(network_peering_peer_nw.id())
 *             .peerNetworkType("VMWARE_ENGINE_NETWORK")
 *             .exportCustomRoutes(false)
 *             .importCustomRoutes(false)
 *             .exportCustomRoutesWithPublicIp(false)
 *             .importCustomRoutesWithPublicIp(false)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Vmware Engine Network Peering Standard
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
 * import com.pulumi.gcp.vmwareengine.Network;
 * import com.pulumi.gcp.vmwareengine.NetworkArgs;
 * import com.pulumi.gcp.vmwareengine.NetworkPeering;
 * import com.pulumi.gcp.vmwareengine.NetworkPeeringArgs;
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
 *         var network_peering_vpc = new Network("network-peering-vpc", NetworkArgs.builder()
 *             .name("default-vpc")
 *             .build());
 * 
 *         var network_peering_standard_nw = new Network("network-peering-standard-nw", NetworkArgs.builder()
 *             .name("default-standard-nw-np")
 *             .location("global")
 *             .type("STANDARD")
 *             .build());
 * 
 *         var vmw_engine_network_peering = new NetworkPeering("vmw-engine-network-peering", NetworkPeeringArgs.builder()
 *             .name("sample-network-peering")
 *             .description("Sample description")
 *             .peerNetwork(network_peering_vpc.id())
 *             .peerNetworkType("STANDARD")
 *             .vmwareEngineNetwork(network_peering_standard_nw.id())
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
 * NetworkPeering can be imported using any of these accepted formats:
 * 
 * * `projects/{{project}}/locations/global/networkPeerings/{{name}}`
 * 
 * * `{{project}}/{{name}}`
 * 
 * * `{{name}}`
 * 
 * When using the `pulumi import` command, NetworkPeering can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:vmwareengine/networkPeering:NetworkPeering default projects/{{project}}/locations/global/networkPeerings/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:vmwareengine/networkPeering:NetworkPeering default {{project}}/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:vmwareengine/networkPeering:NetworkPeering default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:vmwareengine/networkPeering:NetworkPeering")
public class NetworkPeering extends com.pulumi.resources.CustomResource {
    /**
     * Creation time of this resource.
     * A timestamp in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and
     * up to nine fractional digits. Examples: &#34;2014-10-02T15:01:23Z&#34; and &#34;2014-10-02T15:01:23.045123456Z&#34;.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return Creation time of this resource.
     * A timestamp in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and
     * up to nine fractional digits. Examples: &#34;2014-10-02T15:01:23Z&#34; and &#34;2014-10-02T15:01:23.045123456Z&#34;.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * User-provided description for this network peering.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return User-provided description for this network peering.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * True if custom routes are exported to the peered network; false otherwise.
     * 
     */
    @Export(name="exportCustomRoutes", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> exportCustomRoutes;

    /**
     * @return True if custom routes are exported to the peered network; false otherwise.
     * 
     */
    public Output<Optional<Boolean>> exportCustomRoutes() {
        return Codegen.optional(this.exportCustomRoutes);
    }
    /**
     * True if all subnet routes with a public IP address range are exported; false otherwise.
     * 
     */
    @Export(name="exportCustomRoutesWithPublicIp", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> exportCustomRoutesWithPublicIp;

    /**
     * @return True if all subnet routes with a public IP address range are exported; false otherwise.
     * 
     */
    public Output<Optional<Boolean>> exportCustomRoutesWithPublicIp() {
        return Codegen.optional(this.exportCustomRoutesWithPublicIp);
    }
    /**
     * True if custom routes are imported from the peered network; false otherwise.
     * 
     */
    @Export(name="importCustomRoutes", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> importCustomRoutes;

    /**
     * @return True if custom routes are imported from the peered network; false otherwise.
     * 
     */
    public Output<Optional<Boolean>> importCustomRoutes() {
        return Codegen.optional(this.importCustomRoutes);
    }
    /**
     * True if custom routes are imported from the peered network; false otherwise.
     * 
     */
    @Export(name="importCustomRoutesWithPublicIp", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> importCustomRoutesWithPublicIp;

    /**
     * @return True if custom routes are imported from the peered network; false otherwise.
     * 
     */
    public Output<Optional<Boolean>> importCustomRoutesWithPublicIp() {
        return Codegen.optional(this.importCustomRoutesWithPublicIp);
    }
    /**
     * The ID of the Network Peering.
     * 
     * ***
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The ID of the Network Peering.
     * 
     * ***
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The relative resource name of the network to peer with a standard VMware Engine network.
     * The provided network can be a consumer VPC network or another standard VMware Engine network.
     * 
     */
    @Export(name="peerNetwork", refs={String.class}, tree="[0]")
    private Output<String> peerNetwork;

    /**
     * @return The relative resource name of the network to peer with a standard VMware Engine network.
     * The provided network can be a consumer VPC network or another standard VMware Engine network.
     * 
     */
    public Output<String> peerNetwork() {
        return this.peerNetwork;
    }
    /**
     * The type of the network to peer with the VMware Engine network.
     * Possible values are: `STANDARD`, `VMWARE_ENGINE_NETWORK`, `PRIVATE_SERVICES_ACCESS`, `NETAPP_CLOUD_VOLUMES`, `THIRD_PARTY_SERVICE`, `DELL_POWERSCALE`.
     * 
     */
    @Export(name="peerNetworkType", refs={String.class}, tree="[0]")
    private Output<String> peerNetworkType;

    /**
     * @return The type of the network to peer with the VMware Engine network.
     * Possible values are: `STANDARD`, `VMWARE_ENGINE_NETWORK`, `PRIVATE_SERVICES_ACCESS`, `NETAPP_CLOUD_VOLUMES`, `THIRD_PARTY_SERVICE`, `DELL_POWERSCALE`.
     * 
     */
    public Output<String> peerNetworkType() {
        return this.peerNetworkType;
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
     * State of the network peering.
     * This field has a value of &#39;ACTIVE&#39; when there&#39;s a matching configuration in the peer network.
     * New values may be added to this enum when appropriate.
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return State of the network peering.
     * This field has a value of &#39;ACTIVE&#39; when there&#39;s a matching configuration in the peer network.
     * New values may be added to this enum when appropriate.
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * Details about the current state of the network peering.
     * 
     */
    @Export(name="stateDetails", refs={String.class}, tree="[0]")
    private Output<String> stateDetails;

    /**
     * @return Details about the current state of the network peering.
     * 
     */
    public Output<String> stateDetails() {
        return this.stateDetails;
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
     * Last updated time of this resource.
     * A timestamp in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine
     * fractional digits. Examples: &#34;2014-10-02T15:01:23Z&#34; and &#34;2014-10-02T15:01:23.045123456Z&#34;.
     * 
     */
    @Export(name="updateTime", refs={String.class}, tree="[0]")
    private Output<String> updateTime;

    /**
     * @return Last updated time of this resource.
     * A timestamp in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine
     * fractional digits. Examples: &#34;2014-10-02T15:01:23Z&#34; and &#34;2014-10-02T15:01:23.045123456Z&#34;.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }
    /**
     * The relative resource name of the VMware Engine network. Specify the name in the following form:
     * projects/{project}/locations/{location}/vmwareEngineNetworks/{vmwareEngineNetworkId} where {project}
     * can either be a project number or a project ID.
     * 
     */
    @Export(name="vmwareEngineNetwork", refs={String.class}, tree="[0]")
    private Output<String> vmwareEngineNetwork;

    /**
     * @return The relative resource name of the VMware Engine network. Specify the name in the following form:
     * projects/{project}/locations/{location}/vmwareEngineNetworks/{vmwareEngineNetworkId} where {project}
     * can either be a project number or a project ID.
     * 
     */
    public Output<String> vmwareEngineNetwork() {
        return this.vmwareEngineNetwork;
    }
    /**
     * The canonical name of the VMware Engine network in the form:
     * projects/{project_number}/locations/{location}/vmwareEngineNetworks/{vmwareEngineNetworkId}
     * 
     */
    @Export(name="vmwareEngineNetworkCanonical", refs={String.class}, tree="[0]")
    private Output<String> vmwareEngineNetworkCanonical;

    /**
     * @return The canonical name of the VMware Engine network in the form:
     * projects/{project_number}/locations/{location}/vmwareEngineNetworks/{vmwareEngineNetworkId}
     * 
     */
    public Output<String> vmwareEngineNetworkCanonical() {
        return this.vmwareEngineNetworkCanonical;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public NetworkPeering(java.lang.String name) {
        this(name, NetworkPeeringArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public NetworkPeering(java.lang.String name, NetworkPeeringArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public NetworkPeering(java.lang.String name, NetworkPeeringArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:vmwareengine/networkPeering:NetworkPeering", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private NetworkPeering(java.lang.String name, Output<java.lang.String> id, @Nullable NetworkPeeringState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:vmwareengine/networkPeering:NetworkPeering", name, state, makeResourceOptions(options, id), false);
    }

    private static NetworkPeeringArgs makeArgs(NetworkPeeringArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? NetworkPeeringArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static NetworkPeering get(java.lang.String name, Output<java.lang.String> id, @Nullable NetworkPeeringState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new NetworkPeering(name, id, state, options);
    }
}
