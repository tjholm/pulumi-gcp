// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.compute.SubnetworkArgs;
import com.pulumi.gcp.compute.inputs.SubnetworkState;
import com.pulumi.gcp.compute.outputs.SubnetworkLogConfig;
import com.pulumi.gcp.compute.outputs.SubnetworkSecondaryIpRange;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A VPC network is a virtual version of the traditional physical networks
 * that exist within and between physical data centers. A VPC network
 * provides connectivity for your Compute Engine virtual machine (VM)
 * instances, Container Engine containers, App Engine Flex services, and
 * other network-related resources.
 * 
 * Each GCP project contains one or more VPC networks. Each VPC network is a
 * global entity spanning all GCP regions. This global VPC network allows VM
 * instances and other resources to communicate with each other via internal,
 * private IP addresses.
 * 
 * Each VPC network is subdivided into subnets, and each subnet is contained
 * within a single region. You can have more than one subnet in a region for
 * a given VPC network. Each subnet has a contiguous private RFC1918 IP
 * space. You create instances, containers, and the like in these subnets.
 * When you create an instance, you must create it in a subnet, and the
 * instance draws its internal IP address from that subnet.
 * 
 * Virtual machine (VM) instances in a VPC network can communicate with
 * instances in all other subnets of the same VPC network, regardless of
 * region, using their RFC1918 private IP addresses. You can isolate portions
 * of the network, even entire subnets, using firewall rules.
 * 
 * To get more information about Subnetwork, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/subnetworks)
 * * How-to Guides
 *     * [Cloud Networking](https://cloud.google.com/vpc/docs/using-vpc)
 *     * [Private Google Access](https://cloud.google.com/vpc/docs/configure-private-google-access)
 * 
 * ## Example Usage
 * 
 * ### Subnetwork Basic
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
 * import com.pulumi.gcp.compute.Subnetwork;
 * import com.pulumi.gcp.compute.SubnetworkArgs;
 * import com.pulumi.gcp.compute.inputs.SubnetworkSecondaryIpRangeArgs;
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
 *         var custom_test = new Network("custom-test", NetworkArgs.builder()
 *             .name("test-network")
 *             .autoCreateSubnetworks(false)
 *             .build());
 * 
 *         var network_with_private_secondary_ip_ranges = new Subnetwork("network-with-private-secondary-ip-ranges", SubnetworkArgs.builder()
 *             .name("test-subnetwork")
 *             .ipCidrRange("10.2.0.0/16")
 *             .region("us-central1")
 *             .network(custom_test.id())
 *             .secondaryIpRanges(SubnetworkSecondaryIpRangeArgs.builder()
 *                 .rangeName("tf-test-secondary-range-update1")
 *                 .ipCidrRange("192.168.10.0/24")
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Subnetwork Logging Config
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
 * import com.pulumi.gcp.compute.Subnetwork;
 * import com.pulumi.gcp.compute.SubnetworkArgs;
 * import com.pulumi.gcp.compute.inputs.SubnetworkLogConfigArgs;
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
 *         var custom_test = new Network("custom-test", NetworkArgs.builder()
 *             .name("log-test-network")
 *             .autoCreateSubnetworks(false)
 *             .build());
 * 
 *         var subnet_with_logging = new Subnetwork("subnet-with-logging", SubnetworkArgs.builder()
 *             .name("log-test-subnetwork")
 *             .ipCidrRange("10.2.0.0/16")
 *             .region("us-central1")
 *             .network(custom_test.id())
 *             .logConfig(SubnetworkLogConfigArgs.builder()
 *                 .aggregationInterval("INTERVAL_10_MIN")
 *                 .flowSampling(0.5)
 *                 .metadata("INCLUDE_ALL_METADATA")
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Subnetwork Internal L7lb
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
 * import com.pulumi.gcp.compute.Subnetwork;
 * import com.pulumi.gcp.compute.SubnetworkArgs;
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
 *         var custom_test = new Network("custom-test", NetworkArgs.builder()
 *             .name("l7lb-test-network")
 *             .autoCreateSubnetworks(false)
 *             .build());
 * 
 *         var network_for_l7lb = new Subnetwork("network-for-l7lb", SubnetworkArgs.builder()
 *             .name("l7lb-test-subnetwork")
 *             .ipCidrRange("10.0.0.0/22")
 *             .region("us-central1")
 *             .purpose("REGIONAL_MANAGED_PROXY")
 *             .role("ACTIVE")
 *             .network(custom_test.id())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Subnetwork Ipv6
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
 * import com.pulumi.gcp.compute.Subnetwork;
 * import com.pulumi.gcp.compute.SubnetworkArgs;
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
 *         var custom_test = new Network("custom-test", NetworkArgs.builder()
 *             .name("ipv6-test-network")
 *             .autoCreateSubnetworks(false)
 *             .build());
 * 
 *         var subnetwork_ipv6 = new Subnetwork("subnetwork-ipv6", SubnetworkArgs.builder()
 *             .name("ipv6-test-subnetwork")
 *             .ipCidrRange("10.0.0.0/22")
 *             .region("us-west2")
 *             .stackType("IPV4_IPV6")
 *             .ipv6AccessType("EXTERNAL")
 *             .network(custom_test.id())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Subnetwork Internal Ipv6
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
 * import com.pulumi.gcp.compute.Subnetwork;
 * import com.pulumi.gcp.compute.SubnetworkArgs;
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
 *         var custom_test = new Network("custom-test", NetworkArgs.builder()
 *             .name("internal-ipv6-test-network")
 *             .autoCreateSubnetworks(false)
 *             .enableUlaInternalIpv6(true)
 *             .build());
 * 
 *         var subnetwork_internal_ipv6 = new Subnetwork("subnetwork-internal-ipv6", SubnetworkArgs.builder()
 *             .name("internal-ipv6-test-subnetwork")
 *             .ipCidrRange("10.0.0.0/22")
 *             .region("us-west2")
 *             .stackType("IPV4_IPV6")
 *             .ipv6AccessType("INTERNAL")
 *             .network(custom_test.id())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Subnetwork Purpose Private Nat
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
 * import com.pulumi.gcp.compute.Subnetwork;
 * import com.pulumi.gcp.compute.SubnetworkArgs;
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
 *         var custom_test = new Network("custom-test", NetworkArgs.builder()
 *             .name("subnet-purpose-test-network")
 *             .autoCreateSubnetworks(false)
 *             .build());
 * 
 *         var subnetwork_purpose_private_nat = new Subnetwork("subnetwork-purpose-private-nat", SubnetworkArgs.builder()
 *             .name("subnet-purpose-test-subnetwork")
 *             .region("us-west2")
 *             .ipCidrRange("192.168.1.0/24")
 *             .purpose("PRIVATE_NAT")
 *             .network(custom_test.id())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Subnetwork Cidr Overlap
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
 * import com.pulumi.gcp.compute.Subnetwork;
 * import com.pulumi.gcp.compute.SubnetworkArgs;
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
 *         var net_cidr_overlap = new Network("net-cidr-overlap", NetworkArgs.builder()
 *             .name("net-cidr-overlap")
 *             .autoCreateSubnetworks(false)
 *             .build());
 * 
 *         var subnetwork_cidr_overlap = new Subnetwork("subnetwork-cidr-overlap", SubnetworkArgs.builder()
 *             .name("subnet-cidr-overlap")
 *             .region("us-west2")
 *             .ipCidrRange("192.168.1.0/24")
 *             .allowSubnetCidrRoutesOverlap(true)
 *             .network(net_cidr_overlap.id())
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
 * Subnetwork can be imported using any of these accepted formats:
 * 
 * * `projects/{{project}}/regions/{{region}}/subnetworks/{{name}}`
 * 
 * * `{{project}}/{{region}}/{{name}}`
 * 
 * * `{{region}}/{{name}}`
 * 
 * * `{{name}}`
 * 
 * When using the `pulumi import` command, Subnetwork can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:compute/subnetwork:Subnetwork default projects/{{project}}/regions/{{region}}/subnetworks/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:compute/subnetwork:Subnetwork default {{project}}/{{region}}/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:compute/subnetwork:Subnetwork default {{region}}/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:compute/subnetwork:Subnetwork default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:compute/subnetwork:Subnetwork")
public class Subnetwork extends com.pulumi.resources.CustomResource {
    /**
     * Typically packets destined to IPs within the subnetwork range that do not match
     * existing resources are dropped and prevented from leaving the VPC.
     * Setting this field to true will allow these packets to match dynamic routes injected
     * via BGP even if their destinations match existing subnet ranges.
     * 
     */
    @Export(name="allowSubnetCidrRoutesOverlap", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> allowSubnetCidrRoutesOverlap;

    /**
     * @return Typically packets destined to IPs within the subnetwork range that do not match
     * existing resources are dropped and prevented from leaving the VPC.
     * Setting this field to true will allow these packets to match dynamic routes injected
     * via BGP even if their destinations match existing subnet ranges.
     * 
     */
    public Output<Boolean> allowSubnetCidrRoutesOverlap() {
        return this.allowSubnetCidrRoutesOverlap;
    }
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
     * An optional description of this resource. Provide this property when
     * you create the resource. This field can be set only at resource
     * creation time.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return An optional description of this resource. Provide this property when
     * you create the resource. This field can be set only at resource
     * creation time.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The range of external IPv6 addresses that are owned by this subnetwork.
     * 
     */
    @Export(name="externalIpv6Prefix", refs={String.class}, tree="[0]")
    private Output<String> externalIpv6Prefix;

    /**
     * @return The range of external IPv6 addresses that are owned by this subnetwork.
     * 
     */
    public Output<String> externalIpv6Prefix() {
        return this.externalIpv6Prefix;
    }
    /**
     * Fingerprint of this resource. This field is used internally during updates of this resource.
     * 
     * @deprecated
     * This field is not useful for users, and has been removed as an output.
     * 
     */
    @Deprecated /* This field is not useful for users, and has been removed as an output. */
    @Export(name="fingerprint", refs={String.class}, tree="[0]")
    private Output<String> fingerprint;

    /**
     * @return Fingerprint of this resource. This field is used internally during updates of this resource.
     * 
     */
    public Output<String> fingerprint() {
        return this.fingerprint;
    }
    /**
     * The gateway address for default routes to reach destination addresses
     * outside this subnetwork.
     * 
     */
    @Export(name="gatewayAddress", refs={String.class}, tree="[0]")
    private Output<String> gatewayAddress;

    /**
     * @return The gateway address for default routes to reach destination addresses
     * outside this subnetwork.
     * 
     */
    public Output<String> gatewayAddress() {
        return this.gatewayAddress;
    }
    /**
     * The internal IPv6 address range that is assigned to this subnetwork.
     * 
     */
    @Export(name="internalIpv6Prefix", refs={String.class}, tree="[0]")
    private Output<String> internalIpv6Prefix;

    /**
     * @return The internal IPv6 address range that is assigned to this subnetwork.
     * 
     */
    public Output<String> internalIpv6Prefix() {
        return this.internalIpv6Prefix;
    }
    /**
     * The range of internal addresses that are owned by this subnetwork.
     * Provide this property when you create the subnetwork. For example,
     * 10.0.0.0/8 or 192.168.0.0/16. Ranges must be unique and
     * non-overlapping within a network. Only IPv4 is supported.
     * 
     */
    @Export(name="ipCidrRange", refs={String.class}, tree="[0]")
    private Output<String> ipCidrRange;

    /**
     * @return The range of internal addresses that are owned by this subnetwork.
     * Provide this property when you create the subnetwork. For example,
     * 10.0.0.0/8 or 192.168.0.0/16. Ranges must be unique and
     * non-overlapping within a network. Only IPv4 is supported.
     * 
     */
    public Output<String> ipCidrRange() {
        return this.ipCidrRange;
    }
    /**
     * The access type of IPv6 address this subnet holds. It&#39;s immutable and can only be specified during creation
     * or the first time the subnet is updated into IPV4_IPV6 dual stack. If the ipv6_type is EXTERNAL then this subnet
     * cannot enable direct path.
     * Possible values are: `EXTERNAL`, `INTERNAL`.
     * 
     */
    @Export(name="ipv6AccessType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ipv6AccessType;

    /**
     * @return The access type of IPv6 address this subnet holds. It&#39;s immutable and can only be specified during creation
     * or the first time the subnet is updated into IPV4_IPV6 dual stack. If the ipv6_type is EXTERNAL then this subnet
     * cannot enable direct path.
     * Possible values are: `EXTERNAL`, `INTERNAL`.
     * 
     */
    public Output<Optional<String>> ipv6AccessType() {
        return Codegen.optional(this.ipv6AccessType);
    }
    /**
     * The range of internal IPv6 addresses that are owned by this subnetwork.
     * 
     */
    @Export(name="ipv6CidrRange", refs={String.class}, tree="[0]")
    private Output<String> ipv6CidrRange;

    /**
     * @return The range of internal IPv6 addresses that are owned by this subnetwork.
     * 
     */
    public Output<String> ipv6CidrRange() {
        return this.ipv6CidrRange;
    }
    /**
     * This field denotes the VPC flow logging options for this subnetwork. If
     * logging is enabled, logs are exported to Cloud Logging. Flow logging
     * isn&#39;t supported if the subnet `purpose` field is set to subnetwork is
     * `REGIONAL_MANAGED_PROXY` or `GLOBAL_MANAGED_PROXY`.
     * Structure is documented below.
     * 
     */
    @Export(name="logConfig", refs={SubnetworkLogConfig.class}, tree="[0]")
    private Output</* @Nullable */ SubnetworkLogConfig> logConfig;

    /**
     * @return This field denotes the VPC flow logging options for this subnetwork. If
     * logging is enabled, logs are exported to Cloud Logging. Flow logging
     * isn&#39;t supported if the subnet `purpose` field is set to subnetwork is
     * `REGIONAL_MANAGED_PROXY` or `GLOBAL_MANAGED_PROXY`.
     * Structure is documented below.
     * 
     */
    public Output<Optional<SubnetworkLogConfig>> logConfig() {
        return Codegen.optional(this.logConfig);
    }
    /**
     * The name of the resource, provided by the client when initially
     * creating the resource. The name must be 1-63 characters long, and
     * comply with RFC1035. Specifically, the name must be 1-63 characters
     * long and match the regular expression `a-z?` which
     * means the first character must be a lowercase letter, and all
     * following characters must be a dash, lowercase letter, or digit,
     * except the last character, which cannot be a dash.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the resource, provided by the client when initially
     * creating the resource. The name must be 1-63 characters long, and
     * comply with RFC1035. Specifically, the name must be 1-63 characters
     * long and match the regular expression `a-z?` which
     * means the first character must be a lowercase letter, and all
     * following characters must be a dash, lowercase letter, or digit,
     * except the last character, which cannot be a dash.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The network this subnet belongs to.
     * Only networks that are in the distributed mode can have subnetworks.
     * 
     * ***
     * 
     */
    @Export(name="network", refs={String.class}, tree="[0]")
    private Output<String> network;

    /**
     * @return The network this subnet belongs to.
     * Only networks that are in the distributed mode can have subnetworks.
     * 
     * ***
     * 
     */
    public Output<String> network() {
        return this.network;
    }
    /**
     * When enabled, VMs in this subnetwork without external IP addresses can
     * access Google APIs and services by using Private Google Access.
     * 
     */
    @Export(name="privateIpGoogleAccess", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> privateIpGoogleAccess;

    /**
     * @return When enabled, VMs in this subnetwork without external IP addresses can
     * access Google APIs and services by using Private Google Access.
     * 
     */
    public Output<Boolean> privateIpGoogleAccess() {
        return this.privateIpGoogleAccess;
    }
    /**
     * The private IPv6 google access type for the VMs in this subnet.
     * 
     */
    @Export(name="privateIpv6GoogleAccess", refs={String.class}, tree="[0]")
    private Output<String> privateIpv6GoogleAccess;

    /**
     * @return The private IPv6 google access type for the VMs in this subnet.
     * 
     */
    public Output<String> privateIpv6GoogleAccess() {
        return this.privateIpv6GoogleAccess;
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
     * The purpose of the resource. This field can be either `PRIVATE_RFC_1918`, `REGIONAL_MANAGED_PROXY`, `GLOBAL_MANAGED_PROXY`, `PRIVATE_SERVICE_CONNECT` or `PRIVATE_NAT`.
     * A subnet with purpose set to `REGIONAL_MANAGED_PROXY` is a user-created subnetwork that is reserved for regional Envoy-based load balancers.
     * A subnetwork in a given region with purpose set to `GLOBAL_MANAGED_PROXY` is a proxy-only subnet and is shared between all the cross-regional Envoy-based load balancers.
     * A subnetwork with purpose set to `PRIVATE_SERVICE_CONNECT` reserves the subnet for hosting a Private Service Connect published service.
     * A subnetwork with purpose set to `PRIVATE_NAT` is used as source range for Private NAT gateways.
     * Note that `REGIONAL_MANAGED_PROXY` is the preferred setting for all regional Envoy load balancers.
     * If unspecified, the purpose defaults to `PRIVATE_RFC_1918`.
     * 
     */
    @Export(name="purpose", refs={String.class}, tree="[0]")
    private Output<String> purpose;

    /**
     * @return The purpose of the resource. This field can be either `PRIVATE_RFC_1918`, `REGIONAL_MANAGED_PROXY`, `GLOBAL_MANAGED_PROXY`, `PRIVATE_SERVICE_CONNECT` or `PRIVATE_NAT`.
     * A subnet with purpose set to `REGIONAL_MANAGED_PROXY` is a user-created subnetwork that is reserved for regional Envoy-based load balancers.
     * A subnetwork in a given region with purpose set to `GLOBAL_MANAGED_PROXY` is a proxy-only subnet and is shared between all the cross-regional Envoy-based load balancers.
     * A subnetwork with purpose set to `PRIVATE_SERVICE_CONNECT` reserves the subnet for hosting a Private Service Connect published service.
     * A subnetwork with purpose set to `PRIVATE_NAT` is used as source range for Private NAT gateways.
     * Note that `REGIONAL_MANAGED_PROXY` is the preferred setting for all regional Envoy load balancers.
     * If unspecified, the purpose defaults to `PRIVATE_RFC_1918`.
     * 
     */
    public Output<String> purpose() {
        return this.purpose;
    }
    /**
     * The GCP region for this subnetwork.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The GCP region for this subnetwork.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The role of subnetwork.
     * Currently, this field is only used when `purpose` is `REGIONAL_MANAGED_PROXY`.
     * The value can be set to `ACTIVE` or `BACKUP`.
     * An `ACTIVE` subnetwork is one that is currently being used for Envoy-based load balancers in a region.
     * A `BACKUP` subnetwork is one that is ready to be promoted to `ACTIVE` or is currently draining.
     * Possible values are: `ACTIVE`, `BACKUP`.
     * 
     */
    @Export(name="role", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> role;

    /**
     * @return The role of subnetwork.
     * Currently, this field is only used when `purpose` is `REGIONAL_MANAGED_PROXY`.
     * The value can be set to `ACTIVE` or `BACKUP`.
     * An `ACTIVE` subnetwork is one that is currently being used for Envoy-based load balancers in a region.
     * A `BACKUP` subnetwork is one that is ready to be promoted to `ACTIVE` or is currently draining.
     * Possible values are: `ACTIVE`, `BACKUP`.
     * 
     */
    public Output<Optional<String>> role() {
        return Codegen.optional(this.role);
    }
    /**
     * An array of configurations for secondary IP ranges for VM instances
     * contained in this subnetwork. The primary IP of such VM must belong
     * to the primary ipCidrRange of the subnetwork. The alias IPs may belong
     * to either primary or secondary ranges.
     * Structure is documented below.
     * 
     */
    @Export(name="secondaryIpRanges", refs={List.class,SubnetworkSecondaryIpRange.class}, tree="[0,1]")
    private Output<List<SubnetworkSecondaryIpRange>> secondaryIpRanges;

    /**
     * @return An array of configurations for secondary IP ranges for VM instances
     * contained in this subnetwork. The primary IP of such VM must belong
     * to the primary ipCidrRange of the subnetwork. The alias IPs may belong
     * to either primary or secondary ranges.
     * Structure is documented below.
     * 
     */
    public Output<List<SubnetworkSecondaryIpRange>> secondaryIpRanges() {
        return this.secondaryIpRanges;
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
     * The stack type for this subnet to identify whether the IPv6 feature is enabled or not.
     * If not specified IPV4_ONLY will be used.
     * Possible values are: `IPV4_ONLY`, `IPV4_IPV6`.
     * 
     */
    @Export(name="stackType", refs={String.class}, tree="[0]")
    private Output<String> stackType;

    /**
     * @return The stack type for this subnet to identify whether the IPv6 feature is enabled or not.
     * If not specified IPV4_ONLY will be used.
     * Possible values are: `IPV4_ONLY`, `IPV4_IPV6`.
     * 
     */
    public Output<String> stackType() {
        return this.stackType;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Subnetwork(java.lang.String name) {
        this(name, SubnetworkArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Subnetwork(java.lang.String name, SubnetworkArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Subnetwork(java.lang.String name, SubnetworkArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/subnetwork:Subnetwork", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Subnetwork(java.lang.String name, Output<java.lang.String> id, @Nullable SubnetworkState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/subnetwork:Subnetwork", name, state, makeResourceOptions(options, id), false);
    }

    private static SubnetworkArgs makeArgs(SubnetworkArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? SubnetworkArgs.Empty : args;
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
    public static Subnetwork get(java.lang.String name, Output<java.lang.String> id, @Nullable SubnetworkState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Subnetwork(name, id, state, options);
    }
}
