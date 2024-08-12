// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.compute.NetworkPeeringArgs;
import com.pulumi.gcp.compute.inputs.NetworkPeeringState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a network peering within GCE. For more information see
 * [the official documentation](https://cloud.google.com/compute/docs/vpc/vpc-peering)
 * and
 * [API](https://cloud.google.com/compute/docs/reference/latest/networks).
 * 
 * &gt; Both networks must create a peering with each other for the peering
 * to be functional.
 * 
 * &gt; Subnets IP ranges across peered VPC networks cannot overlap.
 * 
 * ## Example Usage
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
 * import com.pulumi.gcp.compute.NetworkPeering;
 * import com.pulumi.gcp.compute.NetworkPeeringArgs;
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
 *             .name("foobar")
 *             .autoCreateSubnetworks("false")
 *             .build());
 * 
 *         var other = new Network("other", NetworkArgs.builder()
 *             .name("other")
 *             .autoCreateSubnetworks("false")
 *             .build());
 * 
 *         var peering1 = new NetworkPeering("peering1", NetworkPeeringArgs.builder()
 *             .name("peering1")
 *             .network(default_.selfLink())
 *             .peerNetwork(other.selfLink())
 *             .build());
 * 
 *         var peering2 = new NetworkPeering("peering2", NetworkPeeringArgs.builder()
 *             .name("peering2")
 *             .network(other.selfLink())
 *             .peerNetwork(default_.selfLink())
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
 * VPC network peerings can be imported using the name and project of the primary network the peering exists in and the name of the network peering
 * 
 * * `{{project_id}}/{{network_id}}/{{peering_id}}`
 * 
 * When using the `pulumi import` command, VPC network peerings can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:compute/networkPeering:NetworkPeering default {{project_id}}/{{network_id}}/{{peering_id}}
 * ```
 * 
 */
@ResourceType(type="gcp:compute/networkPeering:NetworkPeering")
public class NetworkPeering extends com.pulumi.resources.CustomResource {
    /**
     * Whether to export the custom routes to the peer network. Defaults to `false`.
     * 
     */
    @Export(name="exportCustomRoutes", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> exportCustomRoutes;

    /**
     * @return Whether to export the custom routes to the peer network. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> exportCustomRoutes() {
        return Codegen.optional(this.exportCustomRoutes);
    }
    /**
     * Whether subnet routes with public IP range are exported. The default value is true, all subnet routes are exported. The IPv4 special-use ranges (https://en.wikipedia.org/wiki/IPv4#Special_addresses) are always exported to peers and are not controlled by this field.
     * 
     */
    @Export(name="exportSubnetRoutesWithPublicIp", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> exportSubnetRoutesWithPublicIp;

    /**
     * @return Whether subnet routes with public IP range are exported. The default value is true, all subnet routes are exported. The IPv4 special-use ranges (https://en.wikipedia.org/wiki/IPv4#Special_addresses) are always exported to peers and are not controlled by this field.
     * 
     */
    public Output<Optional<Boolean>> exportSubnetRoutesWithPublicIp() {
        return Codegen.optional(this.exportSubnetRoutesWithPublicIp);
    }
    /**
     * Whether to import the custom routes from the peer network. Defaults to `false`.
     * 
     */
    @Export(name="importCustomRoutes", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> importCustomRoutes;

    /**
     * @return Whether to import the custom routes from the peer network. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> importCustomRoutes() {
        return Codegen.optional(this.importCustomRoutes);
    }
    /**
     * Whether subnet routes with public IP range are imported. The default value is false. The IPv4 special-use ranges (https://en.wikipedia.org/wiki/IPv4#Special_addresses) are always imported from peers and are not controlled by this field.
     * 
     */
    @Export(name="importSubnetRoutesWithPublicIp", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> importSubnetRoutesWithPublicIp;

    /**
     * @return Whether subnet routes with public IP range are imported. The default value is false. The IPv4 special-use ranges (https://en.wikipedia.org/wiki/IPv4#Special_addresses) are always imported from peers and are not controlled by this field.
     * 
     */
    public Output<Optional<Boolean>> importSubnetRoutesWithPublicIp() {
        return Codegen.optional(this.importSubnetRoutesWithPublicIp);
    }
    /**
     * Name of the peering.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the peering.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The primary network of the peering.
     * 
     */
    @Export(name="network", refs={String.class}, tree="[0]")
    private Output<String> network;

    /**
     * @return The primary network of the peering.
     * 
     */
    public Output<String> network() {
        return this.network;
    }
    /**
     * The peer network in the peering. The peer network
     * may belong to a different project.
     * 
     */
    @Export(name="peerNetwork", refs={String.class}, tree="[0]")
    private Output<String> peerNetwork;

    /**
     * @return The peer network in the peering. The peer network
     * may belong to a different project.
     * 
     */
    public Output<String> peerNetwork() {
        return this.peerNetwork;
    }
    /**
     * Which IP version(s) of traffic and routes are allowed to be imported or exported between peer networks. The default value is IPV4_ONLY. Possible values: [&#34;IPV4_ONLY&#34;, &#34;IPV4_IPV6&#34;].
     * 
     */
    @Export(name="stackType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> stackType;

    /**
     * @return Which IP version(s) of traffic and routes are allowed to be imported or exported between peer networks. The default value is IPV4_ONLY. Possible values: [&#34;IPV4_ONLY&#34;, &#34;IPV4_IPV6&#34;].
     * 
     */
    public Output<Optional<String>> stackType() {
        return Codegen.optional(this.stackType);
    }
    /**
     * State for the peering, either `ACTIVE` or `INACTIVE`. The peering is
     * `ACTIVE` when there&#39;s a matching configuration in the peer network.
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return State for the peering, either `ACTIVE` or `INACTIVE`. The peering is
     * `ACTIVE` when there&#39;s a matching configuration in the peer network.
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * Details about the current state of the peering.
     * 
     */
    @Export(name="stateDetails", refs={String.class}, tree="[0]")
    private Output<String> stateDetails;

    /**
     * @return Details about the current state of the peering.
     * 
     */
    public Output<String> stateDetails() {
        return this.stateDetails;
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
        super("gcp:compute/networkPeering:NetworkPeering", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private NetworkPeering(java.lang.String name, Output<java.lang.String> id, @Nullable NetworkPeeringState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/networkPeering:NetworkPeering", name, state, makeResourceOptions(options, id), false);
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
