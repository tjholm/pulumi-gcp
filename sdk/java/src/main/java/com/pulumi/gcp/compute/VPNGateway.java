// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.compute.VPNGatewayArgs;
import com.pulumi.gcp.compute.inputs.VPNGatewayState;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Represents a VPN gateway running in GCP. This virtual device is managed
 * by Google, but used only by you.
 * 
 * To get more information about VpnGateway, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/targetVpnGateways)
 * 
 * &gt; **Warning:** Classic VPN is deprecating certain functionality on October 31, 2021. For more information,
 * see the [Classic VPN partial deprecation page](https://cloud.google.com/network-connectivity/docs/vpn/deprecations/classic-vpn-deprecation).
 * 
 * ## Example Usage
 * ### Target Vpn Gateway Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.Network;
 * import com.pulumi.gcp.compute.VPNGateway;
 * import com.pulumi.gcp.compute.VPNGatewayArgs;
 * import com.pulumi.gcp.compute.Address;
 * import com.pulumi.gcp.compute.ForwardingRule;
 * import com.pulumi.gcp.compute.ForwardingRuleArgs;
 * import com.pulumi.gcp.compute.VPNTunnel;
 * import com.pulumi.gcp.compute.VPNTunnelArgs;
 * import com.pulumi.gcp.compute.Route;
 * import com.pulumi.gcp.compute.RouteArgs;
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
 *         var network1 = new Network(&#34;network1&#34;);
 * 
 *         var targetGateway = new VPNGateway(&#34;targetGateway&#34;, VPNGatewayArgs.builder()        
 *             .network(network1.id())
 *             .build());
 * 
 *         var vpnStaticIp = new Address(&#34;vpnStaticIp&#34;);
 * 
 *         var frEsp = new ForwardingRule(&#34;frEsp&#34;, ForwardingRuleArgs.builder()        
 *             .ipProtocol(&#34;ESP&#34;)
 *             .ipAddress(vpnStaticIp.address())
 *             .target(targetGateway.id())
 *             .build());
 * 
 *         var frUdp500 = new ForwardingRule(&#34;frUdp500&#34;, ForwardingRuleArgs.builder()        
 *             .ipProtocol(&#34;UDP&#34;)
 *             .portRange(&#34;500&#34;)
 *             .ipAddress(vpnStaticIp.address())
 *             .target(targetGateway.id())
 *             .build());
 * 
 *         var frUdp4500 = new ForwardingRule(&#34;frUdp4500&#34;, ForwardingRuleArgs.builder()        
 *             .ipProtocol(&#34;UDP&#34;)
 *             .portRange(&#34;4500&#34;)
 *             .ipAddress(vpnStaticIp.address())
 *             .target(targetGateway.id())
 *             .build());
 * 
 *         var tunnel1 = new VPNTunnel(&#34;tunnel1&#34;, VPNTunnelArgs.builder()        
 *             .peerIp(&#34;15.0.0.120&#34;)
 *             .sharedSecret(&#34;a secret message&#34;)
 *             .targetVpnGateway(targetGateway.id())
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(                
 *                     frEsp,
 *                     frUdp500,
 *                     frUdp4500)
 *                 .build());
 * 
 *         var route1 = new Route(&#34;route1&#34;, RouteArgs.builder()        
 *             .network(network1.name())
 *             .destRange(&#34;15.0.0.0/24&#34;)
 *             .priority(1000)
 *             .nextHopVpnTunnel(tunnel1.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * VpnGateway can be imported using any of these accepted formats* `projects/{{project}}/regions/{{region}}/targetVpnGateways/{{name}}` * `{{project}}/{{region}}/{{name}}` * `{{region}}/{{name}}` * `{{name}}` When using the `pulumi import` command, VpnGateway can be imported using one of the formats above. For example
 * 
 * ```sh
 *  $ pulumi import gcp:compute/vPNGateway:VPNGateway default projects/{{project}}/regions/{{region}}/targetVpnGateways/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/vPNGateway:VPNGateway default {{project}}/{{region}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/vPNGateway:VPNGateway default {{region}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/vPNGateway:VPNGateway default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:compute/vPNGateway:VPNGateway")
public class VPNGateway extends com.pulumi.resources.CustomResource {
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
     * The unique identifier for the resource.
     * 
     */
    @Export(name="gatewayId", refs={Integer.class}, tree="[0]")
    private Output<Integer> gatewayId;

    /**
     * @return The unique identifier for the resource.
     * 
     */
    public Output<Integer> gatewayId() {
        return this.gatewayId;
    }
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035.  Specifically, the name must be 1-63 characters long and
     * match the regular expression `a-z?` which means
     * the first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035.  Specifically, the name must be 1-63 characters long and
     * match the regular expression `a-z?` which means
     * the first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The network this VPN gateway is accepting traffic for.
     * 
     * ***
     * 
     */
    @Export(name="network", refs={String.class}, tree="[0]")
    private Output<String> network;

    /**
     * @return The network this VPN gateway is accepting traffic for.
     * 
     * ***
     * 
     */
    public Output<String> network() {
        return this.network;
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
     * The region this gateway should sit in.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region this gateway should sit in.
     * 
     */
    public Output<String> region() {
        return this.region;
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VPNGateway(String name) {
        this(name, VPNGatewayArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VPNGateway(String name, VPNGatewayArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VPNGateway(String name, VPNGatewayArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/vPNGateway:VPNGateway", name, args == null ? VPNGatewayArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private VPNGateway(String name, Output<String> id, @Nullable VPNGatewayState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/vPNGateway:VPNGateway", name, state, makeResourceOptions(options, id));
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
    public static VPNGateway get(String name, Output<String> id, @Nullable VPNGatewayState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VPNGateway(name, id, state, options);
    }
}
