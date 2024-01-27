// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.compute.RouteArgs;
import com.pulumi.gcp.compute.inputs.RouteState;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Represents a Route resource.
 * 
 * A route is a rule that specifies how certain packets should be handled by
 * the virtual network. Routes are associated with virtual machines by tag,
 * and the set of routes for a particular virtual machine is called its
 * routing table. For each packet leaving a virtual machine, the system
 * searches that virtual machine&#39;s routing table for a single best matching
 * route.
 * 
 * Routes match packets by destination IP address, preferring smaller or more
 * specific ranges over larger ones. If there is a tie, the system selects
 * the route with the smallest priority value. If there is still a tie, it
 * uses the layer three and four packet headers to select just one of the
 * remaining matching routes. The packet is then forwarded as specified by
 * the next_hop field of the winning route -- either to another virtual
 * machine destination, a virtual machine gateway or a Compute
 * Engine-operated gateway. Packets that do not match any route in the
 * sending virtual machine&#39;s routing table will be dropped.
 * 
 * A Route resource must have exactly one specification of either
 * nextHopGateway, nextHopInstance, nextHopIp, nextHopVpnTunnel, or
 * nextHopIlb.
 * 
 * To get more information about Route, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/routes)
 * * How-to Guides
 *     * [Using Routes](https://cloud.google.com/vpc/docs/using-routes)
 * 
 * ## Example Usage
 * ### Route Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.Network;
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
 *         var defaultNetwork = new Network(&#34;defaultNetwork&#34;);
 * 
 *         var defaultRoute = new Route(&#34;defaultRoute&#34;, RouteArgs.builder()        
 *             .destRange(&#34;15.0.0.0/24&#34;)
 *             .network(defaultNetwork.name())
 *             .nextHopIp(&#34;10.132.1.5&#34;)
 *             .priority(100)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Route Ilb
 * 
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.Network;
 * import com.pulumi.gcp.compute.NetworkArgs;
 * import com.pulumi.gcp.compute.Subnetwork;
 * import com.pulumi.gcp.compute.SubnetworkArgs;
 * import com.pulumi.gcp.compute.HealthCheck;
 * import com.pulumi.gcp.compute.HealthCheckArgs;
 * import com.pulumi.gcp.compute.inputs.HealthCheckTcpHealthCheckArgs;
 * import com.pulumi.gcp.compute.RegionBackendService;
 * import com.pulumi.gcp.compute.RegionBackendServiceArgs;
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
 *         var defaultNetwork = new Network(&#34;defaultNetwork&#34;, NetworkArgs.builder()        
 *             .autoCreateSubnetworks(false)
 *             .build());
 * 
 *         var defaultSubnetwork = new Subnetwork(&#34;defaultSubnetwork&#34;, SubnetworkArgs.builder()        
 *             .ipCidrRange(&#34;10.0.1.0/24&#34;)
 *             .region(&#34;us-central1&#34;)
 *             .network(defaultNetwork.id())
 *             .build());
 * 
 *         var hc = new HealthCheck(&#34;hc&#34;, HealthCheckArgs.builder()        
 *             .checkIntervalSec(1)
 *             .timeoutSec(1)
 *             .tcpHealthCheck(HealthCheckTcpHealthCheckArgs.builder()
 *                 .port(&#34;80&#34;)
 *                 .build())
 *             .build());
 * 
 *         var backend = new RegionBackendService(&#34;backend&#34;, RegionBackendServiceArgs.builder()        
 *             .region(&#34;us-central1&#34;)
 *             .healthChecks(hc.id())
 *             .build());
 * 
 *         var defaultForwardingRule = new ForwardingRule(&#34;defaultForwardingRule&#34;, ForwardingRuleArgs.builder()        
 *             .region(&#34;us-central1&#34;)
 *             .loadBalancingScheme(&#34;INTERNAL&#34;)
 *             .backendService(backend.id())
 *             .allPorts(true)
 *             .network(defaultNetwork.name())
 *             .subnetwork(defaultSubnetwork.name())
 *             .build());
 * 
 *         var route_ilb = new Route(&#34;route-ilb&#34;, RouteArgs.builder()        
 *             .destRange(&#34;0.0.0.0/0&#34;)
 *             .network(defaultNetwork.name())
 *             .nextHopIlb(defaultForwardingRule.id())
 *             .priority(2000)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Route Ilb Vip
 * 
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.Network;
 * import com.pulumi.gcp.compute.NetworkArgs;
 * import com.pulumi.gcp.compute.Subnetwork;
 * import com.pulumi.gcp.compute.SubnetworkArgs;
 * import com.pulumi.gcp.compute.NetworkPeering;
 * import com.pulumi.gcp.compute.NetworkPeeringArgs;
 * import com.pulumi.gcp.compute.HealthCheck;
 * import com.pulumi.gcp.compute.HealthCheckArgs;
 * import com.pulumi.gcp.compute.inputs.HealthCheckTcpHealthCheckArgs;
 * import com.pulumi.gcp.compute.RegionBackendService;
 * import com.pulumi.gcp.compute.RegionBackendServiceArgs;
 * import com.pulumi.gcp.compute.ForwardingRule;
 * import com.pulumi.gcp.compute.ForwardingRuleArgs;
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
 *         var producerNetwork = new Network(&#34;producerNetwork&#34;, NetworkArgs.builder()        
 *             .autoCreateSubnetworks(false)
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *         var producerSubnetwork = new Subnetwork(&#34;producerSubnetwork&#34;, SubnetworkArgs.builder()        
 *             .ipCidrRange(&#34;10.0.1.0/24&#34;)
 *             .region(&#34;us-central1&#34;)
 *             .network(producerNetwork.id())
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *         var consumerNetwork = new Network(&#34;consumerNetwork&#34;, NetworkArgs.builder()        
 *             .autoCreateSubnetworks(false)
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *         var consumerSubnetwork = new Subnetwork(&#34;consumerSubnetwork&#34;, SubnetworkArgs.builder()        
 *             .ipCidrRange(&#34;10.0.2.0/24&#34;)
 *             .region(&#34;us-central1&#34;)
 *             .network(consumerNetwork.id())
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *         var peering1 = new NetworkPeering(&#34;peering1&#34;, NetworkPeeringArgs.builder()        
 *             .network(consumerNetwork.id())
 *             .peerNetwork(producerNetwork.id())
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *         var peering2 = new NetworkPeering(&#34;peering2&#34;, NetworkPeeringArgs.builder()        
 *             .network(producerNetwork.id())
 *             .peerNetwork(consumerNetwork.id())
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *         var hc = new HealthCheck(&#34;hc&#34;, HealthCheckArgs.builder()        
 *             .checkIntervalSec(1)
 *             .timeoutSec(1)
 *             .tcpHealthCheck(HealthCheckTcpHealthCheckArgs.builder()
 *                 .port(&#34;80&#34;)
 *                 .build())
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *         var backend = new RegionBackendService(&#34;backend&#34;, RegionBackendServiceArgs.builder()        
 *             .region(&#34;us-central1&#34;)
 *             .healthChecks(hc.id())
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *         var default_ = new ForwardingRule(&#34;default&#34;, ForwardingRuleArgs.builder()        
 *             .region(&#34;us-central1&#34;)
 *             .loadBalancingScheme(&#34;INTERNAL&#34;)
 *             .backendService(backend.id())
 *             .allPorts(true)
 *             .network(producerNetwork.name())
 *             .subnetwork(producerSubnetwork.name())
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *         var route_ilb = new Route(&#34;route-ilb&#34;, RouteArgs.builder()        
 *             .destRange(&#34;0.0.0.0/0&#34;)
 *             .network(consumerNetwork.name())
 *             .nextHopIlb(default_.ipAddress())
 *             .priority(2000)
 *             .tags(            
 *                 &#34;tag1&#34;,
 *                 &#34;tag2&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .dependsOn(                
 *                     peering1,
 *                     peering2)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Route can be imported using any of these accepted formats* `projects/{{project}}/global/routes/{{name}}` * `{{project}}/{{name}}` * `{{name}}` When using the `pulumi import` command, Route can be imported using one of the formats above. For example
 * 
 * ```sh
 *  $ pulumi import gcp:compute/route:Route default projects/{{project}}/global/routes/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/route:Route default {{project}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/route:Route default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:compute/route:Route")
public class Route extends com.pulumi.resources.CustomResource {
    /**
     * An optional description of this resource. Provide this property
     * when you create the resource.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return An optional description of this resource. Provide this property
     * when you create the resource.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The destination range of outgoing packets that this route applies to.
     * Only IPv4 is supported.
     * 
     */
    @Export(name="destRange", refs={String.class}, tree="[0]")
    private Output<String> destRange;

    /**
     * @return The destination range of outgoing packets that this route applies to.
     * Only IPv4 is supported.
     * 
     */
    public Output<String> destRange() {
        return this.destRange;
    }
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035.  Specifically, the name must be 1-63 characters long and
     * match the regular expression `a-z?` which means
     * the first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the
     * last character, which cannot be a dash.
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
     * characters must be a dash, lowercase letter, or digit, except the
     * last character, which cannot be a dash.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The network that this route applies to.
     * 
     * ***
     * 
     */
    @Export(name="network", refs={String.class}, tree="[0]")
    private Output<String> network;

    /**
     * @return The network that this route applies to.
     * 
     * ***
     * 
     */
    public Output<String> network() {
        return this.network;
    }
    /**
     * URL to a gateway that should handle matching packets.
     * Currently, you can only specify the internet gateway, using a full or
     * partial valid URL:
     * * `https://www.googleapis.com/compute/v1/projects/project/global/gateways/default-internet-gateway`
     * * `projects/project/global/gateways/default-internet-gateway`
     * * `global/gateways/default-internet-gateway`
     * * The string `default-internet-gateway`.
     * 
     */
    @Export(name="nextHopGateway", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> nextHopGateway;

    /**
     * @return URL to a gateway that should handle matching packets.
     * Currently, you can only specify the internet gateway, using a full or
     * partial valid URL:
     * * `https://www.googleapis.com/compute/v1/projects/project/global/gateways/default-internet-gateway`
     * * `projects/project/global/gateways/default-internet-gateway`
     * * `global/gateways/default-internet-gateway`
     * * The string `default-internet-gateway`.
     * 
     */
    public Output<Optional<String>> nextHopGateway() {
        return Codegen.optional(this.nextHopGateway);
    }
    /**
     * The IP address or URL to a forwarding rule of type
     * loadBalancingScheme=INTERNAL that should handle matching
     * packets.
     * With the GA provider you can only specify the forwarding
     * rule as a partial or full URL. For example, the following
     * are all valid values:
     * * 10.128.0.56
     * * https://www.googleapis.com/compute/v1/projects/project/regions/region/forwardingRules/forwardingRule
     * * regions/region/forwardingRules/forwardingRule
     *   When the beta provider, you can also specify the IP address
     *   of a forwarding rule from the same VPC or any peered VPC.
     *   Note that this can only be used when the destinationRange is
     *   a public (non-RFC 1918) IP CIDR range.
     * 
     */
    @Export(name="nextHopIlb", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> nextHopIlb;

    /**
     * @return The IP address or URL to a forwarding rule of type
     * loadBalancingScheme=INTERNAL that should handle matching
     * packets.
     * With the GA provider you can only specify the forwarding
     * rule as a partial or full URL. For example, the following
     * are all valid values:
     * * 10.128.0.56
     * * https://www.googleapis.com/compute/v1/projects/project/regions/region/forwardingRules/forwardingRule
     * * regions/region/forwardingRules/forwardingRule
     *   When the beta provider, you can also specify the IP address
     *   of a forwarding rule from the same VPC or any peered VPC.
     *   Note that this can only be used when the destinationRange is
     *   a public (non-RFC 1918) IP CIDR range.
     * 
     */
    public Output<Optional<String>> nextHopIlb() {
        return Codegen.optional(this.nextHopIlb);
    }
    /**
     * URL to an instance that should handle matching packets.
     * You can specify this as a full or partial URL. For example:
     * * `https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/instance`
     * * `projects/project/zones/zone/instances/instance`
     * * `zones/zone/instances/instance`
     * * Just the instance name, with the zone in `next_hop_instance_zone`.
     * 
     */
    @Export(name="nextHopInstance", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> nextHopInstance;

    /**
     * @return URL to an instance that should handle matching packets.
     * You can specify this as a full or partial URL. For example:
     * * `https://www.googleapis.com/compute/v1/projects/project/zones/zone/instances/instance`
     * * `projects/project/zones/zone/instances/instance`
     * * `zones/zone/instances/instance`
     * * Just the instance name, with the zone in `next_hop_instance_zone`.
     * 
     */
    public Output<Optional<String>> nextHopInstance() {
        return Codegen.optional(this.nextHopInstance);
    }
    /**
     * (Optional when `next_hop_instance` is
     * specified)  The zone of the instance specified in
     * `next_hop_instance`.  Omit if `next_hop_instance` is specified as
     * a URL.
     * 
     */
    @Export(name="nextHopInstanceZone", refs={String.class}, tree="[0]")
    private Output<String> nextHopInstanceZone;

    /**
     * @return (Optional when `next_hop_instance` is
     * specified)  The zone of the instance specified in
     * `next_hop_instance`.  Omit if `next_hop_instance` is specified as
     * a URL.
     * 
     */
    public Output<String> nextHopInstanceZone() {
        return this.nextHopInstanceZone;
    }
    /**
     * Network IP address of an instance that should handle matching packets.
     * 
     */
    @Export(name="nextHopIp", refs={String.class}, tree="[0]")
    private Output<String> nextHopIp;

    /**
     * @return Network IP address of an instance that should handle matching packets.
     * 
     */
    public Output<String> nextHopIp() {
        return this.nextHopIp;
    }
    /**
     * URL to a Network that should handle matching packets.
     * 
     */
    @Export(name="nextHopNetwork", refs={String.class}, tree="[0]")
    private Output<String> nextHopNetwork;

    /**
     * @return URL to a Network that should handle matching packets.
     * 
     */
    public Output<String> nextHopNetwork() {
        return this.nextHopNetwork;
    }
    /**
     * URL to a VpnTunnel that should handle matching packets.
     * 
     */
    @Export(name="nextHopVpnTunnel", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> nextHopVpnTunnel;

    /**
     * @return URL to a VpnTunnel that should handle matching packets.
     * 
     */
    public Output<Optional<String>> nextHopVpnTunnel() {
        return Codegen.optional(this.nextHopVpnTunnel);
    }
    /**
     * The priority of this route. Priority is used to break ties in cases
     * where there is more than one matching route of equal prefix length.
     * In the case of two routes with equal prefix length, the one with the
     * lowest-numbered priority value wins.
     * Default value is 1000. Valid range is 0 through 65535.
     * 
     */
    @Export(name="priority", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> priority;

    /**
     * @return The priority of this route. Priority is used to break ties in cases
     * where there is more than one matching route of equal prefix length.
     * In the case of two routes with equal prefix length, the one with the
     * lowest-numbered priority value wins.
     * Default value is 1000. Valid range is 0 through 65535.
     * 
     */
    public Output<Optional<Integer>> priority() {
        return Codegen.optional(this.priority);
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
     * A list of instance tags to which this route applies.
     * 
     */
    @Export(name="tags", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> tags;

    /**
     * @return A list of instance tags to which this route applies.
     * 
     */
    public Output<Optional<List<String>>> tags() {
        return Codegen.optional(this.tags);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Route(String name) {
        this(name, RouteArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Route(String name, RouteArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Route(String name, RouteArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/route:Route", name, args == null ? RouteArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Route(String name, Output<String> id, @Nullable RouteState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/route:Route", name, state, makeResourceOptions(options, id));
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
    public static Route get(String name, Output<String> id, @Nullable RouteState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Route(name, id, state, options);
    }
}
