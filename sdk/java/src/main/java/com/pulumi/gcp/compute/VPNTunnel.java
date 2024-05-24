// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.compute.VPNTunnelArgs;
import com.pulumi.gcp.compute.inputs.VPNTunnelState;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * VPN tunnel resource.
 * 
 * To get more information about VpnTunnel, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/vpnTunnels)
 * * How-to Guides
 *     * [Cloud VPN Overview](https://cloud.google.com/vpn/docs/concepts/overview)
 *     * [Networks and Tunnel Routing](https://cloud.google.com/vpn/docs/concepts/choosing-networks-routing)
 * 
 * ## Example Usage
 * 
 * ### Vpn Tunnel Basic
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
 * import com.pulumi.gcp.compute.VPNGateway;
 * import com.pulumi.gcp.compute.VPNGatewayArgs;
 * import com.pulumi.gcp.compute.VPNTunnel;
 * import com.pulumi.gcp.compute.VPNTunnelArgs;
 * import com.pulumi.gcp.compute.Address;
 * import com.pulumi.gcp.compute.AddressArgs;
 * import com.pulumi.gcp.compute.ForwardingRule;
 * import com.pulumi.gcp.compute.ForwardingRuleArgs;
 * import com.pulumi.gcp.compute.Route;
 * import com.pulumi.gcp.compute.RouteArgs;
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
 *         var network1 = new Network("network1", NetworkArgs.builder()
 *             .name("network-1")
 *             .build());
 * 
 *         var targetGateway = new VPNGateway("targetGateway", VPNGatewayArgs.builder()
 *             .name("vpn-1")
 *             .network(network1.id())
 *             .build());
 * 
 *         var tunnel1 = new VPNTunnel("tunnel1", VPNTunnelArgs.builder()
 *             .name("tunnel-1")
 *             .peerIp("15.0.0.120")
 *             .sharedSecret("a secret message")
 *             .targetVpnGateway(targetGateway.id())
 *             .labels(Map.of("foo", "bar"))
 *             .build());
 * 
 *         var vpnStaticIp = new Address("vpnStaticIp", AddressArgs.builder()
 *             .name("vpn-static-ip")
 *             .build());
 * 
 *         var frEsp = new ForwardingRule("frEsp", ForwardingRuleArgs.builder()
 *             .name("fr-esp")
 *             .ipProtocol("ESP")
 *             .ipAddress(vpnStaticIp.address())
 *             .target(targetGateway.id())
 *             .build());
 * 
 *         var frUdp500 = new ForwardingRule("frUdp500", ForwardingRuleArgs.builder()
 *             .name("fr-udp500")
 *             .ipProtocol("UDP")
 *             .portRange("500")
 *             .ipAddress(vpnStaticIp.address())
 *             .target(targetGateway.id())
 *             .build());
 * 
 *         var frUdp4500 = new ForwardingRule("frUdp4500", ForwardingRuleArgs.builder()
 *             .name("fr-udp4500")
 *             .ipProtocol("UDP")
 *             .portRange("4500")
 *             .ipAddress(vpnStaticIp.address())
 *             .target(targetGateway.id())
 *             .build());
 * 
 *         var route1 = new Route("route1", RouteArgs.builder()
 *             .name("route1")
 *             .network(network1.name())
 *             .destRange("15.0.0.0/24")
 *             .priority(1000)
 *             .nextHopVpnTunnel(tunnel1.id())
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
 * VpnTunnel can be imported using any of these accepted formats:
 * 
 * * `projects/{{project}}/regions/{{region}}/vpnTunnels/{{name}}`
 * 
 * * `{{project}}/{{region}}/{{name}}`
 * 
 * * `{{region}}/{{name}}`
 * 
 * * `{{name}}`
 * 
 * When using the `pulumi import` command, VpnTunnel can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:compute/vPNTunnel:VPNTunnel default projects/{{project}}/regions/{{region}}/vpnTunnels/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:compute/vPNTunnel:VPNTunnel default {{project}}/{{region}}/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:compute/vPNTunnel:VPNTunnel default {{region}}/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:compute/vPNTunnel:VPNTunnel default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:compute/vPNTunnel:VPNTunnel")
public class VPNTunnel extends com.pulumi.resources.CustomResource {
    /**
     * Creation timestamp in RFC3339 text format.
     * 
     */
    @Export(name="creationTimestamp", refs={String.class}, tree="[0]")
    private Output<String> creationTimestamp;

    /**
     * @return Creation timestamp in RFC3339 text format.
     * 
     */
    public Output<String> creationTimestamp() {
        return this.creationTimestamp;
    }
    /**
     * An optional description of this resource.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return An optional description of this resource.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Detailed status message for the VPN tunnel.
     * 
     */
    @Export(name="detailedStatus", refs={String.class}, tree="[0]")
    private Output<String> detailedStatus;

    /**
     * @return Detailed status message for the VPN tunnel.
     * 
     */
    public Output<String> detailedStatus() {
        return this.detailedStatus;
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
     * IKE protocol version to use when establishing the VPN tunnel with
     * peer VPN gateway.
     * Acceptable IKE versions are 1 or 2. Default version is 2.
     * 
     */
    @Export(name="ikeVersion", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> ikeVersion;

    /**
     * @return IKE protocol version to use when establishing the VPN tunnel with
     * peer VPN gateway.
     * Acceptable IKE versions are 1 or 2. Default version is 2.
     * 
     */
    public Output<Optional<Integer>> ikeVersion() {
        return Codegen.optional(this.ikeVersion);
    }
    /**
     * The fingerprint used for optimistic locking of this resource.  Used
     * internally during updates.
     * 
     */
    @Export(name="labelFingerprint", refs={String.class}, tree="[0]")
    private Output<String> labelFingerprint;

    /**
     * @return The fingerprint used for optimistic locking of this resource.  Used
     * internally during updates.
     * 
     */
    public Output<String> labelFingerprint() {
        return this.labelFingerprint;
    }
    /**
     * Labels to apply to this VpnTunnel.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return Labels to apply to this VpnTunnel.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * Local traffic selector to use when establishing the VPN tunnel with
     * peer VPN gateway. The value should be a CIDR formatted string,
     * for example `192.168.0.0/16`. The ranges should be disjoint.
     * Only IPv4 is supported.
     * 
     */
    @Export(name="localTrafficSelectors", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> localTrafficSelectors;

    /**
     * @return Local traffic selector to use when establishing the VPN tunnel with
     * peer VPN gateway. The value should be a CIDR formatted string,
     * for example `192.168.0.0/16`. The ranges should be disjoint.
     * Only IPv4 is supported.
     * 
     */
    public Output<List<String>> localTrafficSelectors() {
        return this.localTrafficSelectors;
    }
    /**
     * Name of the resource. The name must be 1-63 characters long, and
     * comply with RFC1035. Specifically, the name must be 1-63
     * characters long and match the regular expression
     * `a-z?` which means the first character
     * must be a lowercase letter, and all following characters must
     * be a dash, lowercase letter, or digit,
     * except the last character, which cannot be a dash.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the resource. The name must be 1-63 characters long, and
     * comply with RFC1035. Specifically, the name must be 1-63
     * characters long and match the regular expression
     * `a-z?` which means the first character
     * must be a lowercase letter, and all following characters must
     * be a dash, lowercase letter, or digit,
     * except the last character, which cannot be a dash.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * URL of the peer side external VPN gateway to which this VPN tunnel is connected.
     * 
     */
    @Export(name="peerExternalGateway", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> peerExternalGateway;

    /**
     * @return URL of the peer side external VPN gateway to which this VPN tunnel is connected.
     * 
     */
    public Output<Optional<String>> peerExternalGateway() {
        return Codegen.optional(this.peerExternalGateway);
    }
    /**
     * The interface ID of the external VPN gateway to which this VPN tunnel is connected.
     * 
     */
    @Export(name="peerExternalGatewayInterface", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> peerExternalGatewayInterface;

    /**
     * @return The interface ID of the external VPN gateway to which this VPN tunnel is connected.
     * 
     */
    public Output<Optional<Integer>> peerExternalGatewayInterface() {
        return Codegen.optional(this.peerExternalGatewayInterface);
    }
    /**
     * URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected.
     * If provided, the VPN tunnel will automatically use the same vpn_gateway_interface
     * ID in the peer GCP VPN gateway.
     * This field must reference a `gcp.compute.HaVpnGateway` resource.
     * 
     */
    @Export(name="peerGcpGateway", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> peerGcpGateway;

    /**
     * @return URL of the peer side HA GCP VPN gateway to which this VPN tunnel is connected.
     * If provided, the VPN tunnel will automatically use the same vpn_gateway_interface
     * ID in the peer GCP VPN gateway.
     * This field must reference a `gcp.compute.HaVpnGateway` resource.
     * 
     */
    public Output<Optional<String>> peerGcpGateway() {
        return Codegen.optional(this.peerGcpGateway);
    }
    /**
     * IP address of the peer VPN gateway. Only IPv4 is supported.
     * 
     */
    @Export(name="peerIp", refs={String.class}, tree="[0]")
    private Output<String> peerIp;

    /**
     * @return IP address of the peer VPN gateway. Only IPv4 is supported.
     * 
     */
    public Output<String> peerIp() {
        return this.peerIp;
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
     * The region where the tunnel is located. If unset, is set to the region of `target_vpn_gateway`.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region where the tunnel is located. If unset, is set to the region of `target_vpn_gateway`.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * Remote traffic selector to use when establishing the VPN tunnel with
     * peer VPN gateway. The value should be a CIDR formatted string,
     * for example `192.168.0.0/16`. The ranges should be disjoint.
     * Only IPv4 is supported.
     * 
     */
    @Export(name="remoteTrafficSelectors", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> remoteTrafficSelectors;

    /**
     * @return Remote traffic selector to use when establishing the VPN tunnel with
     * peer VPN gateway. The value should be a CIDR formatted string,
     * for example `192.168.0.0/16`. The ranges should be disjoint.
     * Only IPv4 is supported.
     * 
     */
    public Output<List<String>> remoteTrafficSelectors() {
        return this.remoteTrafficSelectors;
    }
    /**
     * URL of router resource to be used for dynamic routing.
     * 
     */
    @Export(name="router", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> router;

    /**
     * @return URL of router resource to be used for dynamic routing.
     * 
     */
    public Output<Optional<String>> router() {
        return Codegen.optional(this.router);
    }
    /**
     * The URI of the created resource.
     * 
     */
    @Export(name="selfLink", refs={String.class}, tree="[0]")
    private Output<String> selfLink;

    /**
     * @return The URI of the created resource.
     * 
     */
    public Output<String> selfLink() {
        return this.selfLink;
    }
    /**
     * Shared secret used to set the secure session between the Cloud VPN
     * gateway and the peer VPN gateway.
     * **Note**: This property is sensitive and will not be displayed in the plan.
     * 
     * ***
     * 
     */
    @Export(name="sharedSecret", refs={String.class}, tree="[0]")
    private Output<String> sharedSecret;

    /**
     * @return Shared secret used to set the secure session between the Cloud VPN
     * gateway and the peer VPN gateway.
     * **Note**: This property is sensitive and will not be displayed in the plan.
     * 
     * ***
     * 
     */
    public Output<String> sharedSecret() {
        return this.sharedSecret;
    }
    /**
     * Hash of the shared secret.
     * 
     */
    @Export(name="sharedSecretHash", refs={String.class}, tree="[0]")
    private Output<String> sharedSecretHash;

    /**
     * @return Hash of the shared secret.
     * 
     */
    public Output<String> sharedSecretHash() {
        return this.sharedSecretHash;
    }
    /**
     * URL of the Target VPN gateway with which this VPN tunnel is
     * associated.
     * 
     */
    @Export(name="targetVpnGateway", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> targetVpnGateway;

    /**
     * @return URL of the Target VPN gateway with which this VPN tunnel is
     * associated.
     * 
     */
    public Output<Optional<String>> targetVpnGateway() {
        return Codegen.optional(this.targetVpnGateway);
    }
    /**
     * The unique identifier for the resource. This identifier is defined by the server.
     * 
     */
    @Export(name="tunnelId", refs={String.class}, tree="[0]")
    private Output<String> tunnelId;

    /**
     * @return The unique identifier for the resource. This identifier is defined by the server.
     * 
     */
    public Output<String> tunnelId() {
        return this.tunnelId;
    }
    /**
     * URL of the VPN gateway with which this VPN tunnel is associated.
     * This must be used if a High Availability VPN gateway resource is created.
     * This field must reference a `gcp.compute.HaVpnGateway` resource.
     * 
     */
    @Export(name="vpnGateway", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> vpnGateway;

    /**
     * @return URL of the VPN gateway with which this VPN tunnel is associated.
     * This must be used if a High Availability VPN gateway resource is created.
     * This field must reference a `gcp.compute.HaVpnGateway` resource.
     * 
     */
    public Output<Optional<String>> vpnGateway() {
        return Codegen.optional(this.vpnGateway);
    }
    /**
     * The interface ID of the VPN gateway with which this VPN tunnel is associated.
     * 
     */
    @Export(name="vpnGatewayInterface", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> vpnGatewayInterface;

    /**
     * @return The interface ID of the VPN gateway with which this VPN tunnel is associated.
     * 
     */
    public Output<Optional<Integer>> vpnGatewayInterface() {
        return Codegen.optional(this.vpnGatewayInterface);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VPNTunnel(String name) {
        this(name, VPNTunnelArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VPNTunnel(String name, VPNTunnelArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VPNTunnel(String name, VPNTunnelArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/vPNTunnel:VPNTunnel", name, args == null ? VPNTunnelArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private VPNTunnel(String name, Output<String> id, @Nullable VPNTunnelState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/vPNTunnel:VPNTunnel", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "effectiveLabels",
                "pulumiLabels",
                "sharedSecret"
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
    public static VPNTunnel get(String name, Output<String> id, @Nullable VPNTunnelState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VPNTunnel(name, id, state, options);
    }
}
