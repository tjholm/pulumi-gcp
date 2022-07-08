// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.compute.AddressArgs;
import com.pulumi.gcp.compute.inputs.AddressState;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Represents an Address resource.
 * 
 * Each virtual machine instance has an ephemeral internal IP address and,
 * optionally, an external IP address. To communicate between instances on
 * the same network, you can use an instance&#39;s internal IP address. To
 * communicate with the Internet and instances outside of the same network,
 * you must specify the instance&#39;s external IP address.
 * 
 * Internal IP addresses are ephemeral and only belong to an instance for
 * the lifetime of the instance; if the instance is deleted and recreated,
 * the instance is assigned a new internal IP address, either by Compute
 * Engine or by you. External IP addresses can be either ephemeral or
 * static.
 * 
 * To get more information about Address, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/beta/addresses)
 * * How-to Guides
 *     * [Reserving a Static External IP Address](https://cloud.google.com/compute/docs/instances-and-network)
 *     * [Reserving a Static Internal IP Address](https://cloud.google.com/compute/docs/ip-addresses/reserve-static-internal-ip-address)
 * 
 * ## Example Usage
 * ### Address Basic
 * ```java
 * package generated_program;
 * 
 * import java.util.*;
 * import java.io.*;
 * import java.nio.*;
 * import com.pulumi.*;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var ipAddress = new Address(&#34;ipAddress&#34;);
 * 
 *     }
 * }
 * ```
 * ### Address With Subnetwork
 * ```java
 * package generated_program;
 * 
 * import java.util.*;
 * import java.io.*;
 * import java.nio.*;
 * import com.pulumi.*;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var defaultNetwork = new Network(&#34;defaultNetwork&#34;);
 * 
 *         var defaultSubnetwork = new Subnetwork(&#34;defaultSubnetwork&#34;, SubnetworkArgs.builder()        
 *             .ipCidrRange(&#34;10.0.0.0/16&#34;)
 *             .region(&#34;us-central1&#34;)
 *             .network(defaultNetwork.id())
 *             .build());
 * 
 *         var internalWithSubnetAndAddress = new Address(&#34;internalWithSubnetAndAddress&#34;, AddressArgs.builder()        
 *             .subnetwork(defaultSubnetwork.id())
 *             .addressType(&#34;INTERNAL&#34;)
 *             .address(&#34;10.0.42.42&#34;)
 *             .region(&#34;us-central1&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Address With Gce Endpoint
 * ```java
 * package generated_program;
 * 
 * import java.util.*;
 * import java.io.*;
 * import java.nio.*;
 * import com.pulumi.*;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var internalWithGceEndpoint = new Address(&#34;internalWithGceEndpoint&#34;, AddressArgs.builder()        
 *             .addressType(&#34;INTERNAL&#34;)
 *             .purpose(&#34;GCE_ENDPOINT&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Instance With Ip
 * ```java
 * package generated_program;
 * 
 * import java.util.*;
 * import java.io.*;
 * import java.nio.*;
 * import com.pulumi.*;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var static_ = new Address(&#34;static&#34;);
 * 
 *         final var debianImage = Output.of(ComputeFunctions.getImage(GetImageArgs.builder()
 *             .family(&#34;debian-9&#34;)
 *             .project(&#34;debian-cloud&#34;)
 *             .build()));
 * 
 *         var instanceWithIp = new Instance(&#34;instanceWithIp&#34;, InstanceArgs.builder()        
 *             .machineType(&#34;f1-micro&#34;)
 *             .zone(&#34;us-central1-a&#34;)
 *             .bootDisk(InstanceBootDiskArgs.builder()
 *                 .initializeParams(InstanceBootDiskInitializeParamsArgs.builder()
 *                     .image(debianImage.apply(getImageResult -&gt; getImageResult.selfLink()))
 *                     .build())
 *                 .build())
 *             .networkInterfaces(InstanceNetworkInterfaceArgs.builder()
 *                 .network(&#34;default&#34;)
 *                 .accessConfigs(InstanceNetworkInterfaceAccessConfigArgs.builder()
 *                     .natIp(static_.address())
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Compute Address Ipsec Interconnect
 * ```java
 * package generated_program;
 * 
 * import java.util.*;
 * import java.io.*;
 * import java.nio.*;
 * import com.pulumi.*;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var network = new Network(&#34;network&#34;, NetworkArgs.builder()        
 *             .autoCreateSubnetworks(false)
 *             .build());
 * 
 *         var ipsec_interconnect_address = new Address(&#34;ipsec-interconnect-address&#34;, AddressArgs.builder()        
 *             .addressType(&#34;INTERNAL&#34;)
 *             .purpose(&#34;IPSEC_INTERCONNECT&#34;)
 *             .address(&#34;192.168.1.0&#34;)
 *             .prefixLength(29)
 *             .network(network.selfLink())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Address can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:compute/address:Address default projects/{{project}}/regions/{{region}}/addresses/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/address:Address default {{project}}/{{region}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/address:Address default {{region}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:compute/address:Address default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:compute/address:Address")
public class Address extends com.pulumi.resources.CustomResource {
    /**
     * The static external IP address represented by this resource. Only
     * IPv4 is supported. An address may only be specified for INTERNAL
     * address types. The IP address must be inside the specified subnetwork,
     * if any. Set by the API if undefined.
     * 
     */
    @Export(name="address", type=String.class, parameters={})
    private Output<String> address;

    /**
     * @return The static external IP address represented by this resource. Only
     * IPv4 is supported. An address may only be specified for INTERNAL
     * address types. The IP address must be inside the specified subnetwork,
     * if any. Set by the API if undefined.
     * 
     */
    public Output<String> address() {
        return this.address;
    }
    /**
     * The type of address to reserve.
     * Default value is `EXTERNAL`.
     * Possible values are `INTERNAL` and `EXTERNAL`.
     * 
     */
    @Export(name="addressType", type=String.class, parameters={})
    private Output</* @Nullable */ String> addressType;

    /**
     * @return The type of address to reserve.
     * Default value is `EXTERNAL`.
     * Possible values are `INTERNAL` and `EXTERNAL`.
     * 
     */
    public Output<Optional<String>> addressType() {
        return Codegen.optional(this.addressType);
    }
    /**
     * Creation timestamp in RFC3339 text format.
     * 
     */
    @Export(name="creationTimestamp", type=String.class, parameters={})
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
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return An optional description of this resource.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The fingerprint used for optimistic locking of this resource. Used internally during updates.
     * 
     */
    @Export(name="labelFingerprint", type=String.class, parameters={})
    private Output<String> labelFingerprint;

    /**
     * @return The fingerprint used for optimistic locking of this resource. Used internally during updates.
     * 
     */
    public Output<String> labelFingerprint() {
        return this.labelFingerprint;
    }
    /**
     * Labels to apply to this address.  A list of key-&gt;value pairs.
     * 
     */
    @Export(name="labels", type=Map.class, parameters={String.class, String.class})
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return Labels to apply to this address.  A list of key-&gt;value pairs.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * Name of the resource. The name must be 1-63 characters long, and
     * comply with RFC1035. Specifically, the name must be 1-63 characters
     * long and match the regular expression `a-z?`
     * which means the first character must be a lowercase letter, and all
     * following characters must be a dash, lowercase letter, or digit,
     * except the last character, which cannot be a dash.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return Name of the resource. The name must be 1-63 characters long, and
     * comply with RFC1035. Specifically, the name must be 1-63 characters
     * long and match the regular expression `a-z?`
     * which means the first character must be a lowercase letter, and all
     * following characters must be a dash, lowercase letter, or digit,
     * except the last character, which cannot be a dash.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The URL of the network in which to reserve the address. This field
     * can only be used with INTERNAL type with the VPC_PEERING and
     * IPSEC_INTERCONNECT purposes.
     * 
     */
    @Export(name="network", type=String.class, parameters={})
    private Output</* @Nullable */ String> network;

    /**
     * @return The URL of the network in which to reserve the address. This field
     * can only be used with INTERNAL type with the VPC_PEERING and
     * IPSEC_INTERCONNECT purposes.
     * 
     */
    public Output<Optional<String>> network() {
        return Codegen.optional(this.network);
    }
    /**
     * The networking tier used for configuring this address. If this field is not
     * specified, it is assumed to be PREMIUM.
     * Possible values are `PREMIUM` and `STANDARD`.
     * 
     */
    @Export(name="networkTier", type=String.class, parameters={})
    private Output<String> networkTier;

    /**
     * @return The networking tier used for configuring this address. If this field is not
     * specified, it is assumed to be PREMIUM.
     * Possible values are `PREMIUM` and `STANDARD`.
     * 
     */
    public Output<String> networkTier() {
        return this.networkTier;
    }
    /**
     * The prefix length if the resource represents an IP range.
     * 
     */
    @Export(name="prefixLength", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> prefixLength;

    /**
     * @return The prefix length if the resource represents an IP range.
     * 
     */
    public Output<Optional<Integer>> prefixLength() {
        return Codegen.optional(this.prefixLength);
    }
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    @Export(name="project", type=String.class, parameters={})
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
     * The purpose of this resource, which can be one of the following values:
     * * GCE_ENDPOINT for addresses that are used by VM instances, alias IP
     *   ranges, internal load balancers, and similar resources.
     * * SHARED_LOADBALANCER_VIP for an address that can be used by multiple
     *   internal load balancers.
     * * VPC_PEERING for addresses that are reserved for VPC peer networks.
     * * IPSEC_INTERCONNECT for addresses created from a private IP range
     *   that are reserved for a VLAN attachment in an IPsec-encrypted Cloud
     *   Interconnect configuration. These addresses are regional resources.
     * * PRIVATE_SERVICE_CONNECT for a private network address that is used
     *   to configure Private Service Connect. Only global internal addresses
     *   can use this purpose.
     *   This should only be set when using an Internal address.
     * 
     */
    @Export(name="purpose", type=String.class, parameters={})
    private Output<String> purpose;

    /**
     * @return The purpose of this resource, which can be one of the following values:
     * * GCE_ENDPOINT for addresses that are used by VM instances, alias IP
     *   ranges, internal load balancers, and similar resources.
     * * SHARED_LOADBALANCER_VIP for an address that can be used by multiple
     *   internal load balancers.
     * * VPC_PEERING for addresses that are reserved for VPC peer networks.
     * * IPSEC_INTERCONNECT for addresses created from a private IP range
     *   that are reserved for a VLAN attachment in an IPsec-encrypted Cloud
     *   Interconnect configuration. These addresses are regional resources.
     * * PRIVATE_SERVICE_CONNECT for a private network address that is used
     *   to configure Private Service Connect. Only global internal addresses
     *   can use this purpose.
     *   This should only be set when using an Internal address.
     * 
     */
    public Output<String> purpose() {
        return this.purpose;
    }
    /**
     * The Region in which the created address should reside.
     * If it is not provided, the provider region is used.
     * 
     */
    @Export(name="region", type=String.class, parameters={})
    private Output<String> region;

    /**
     * @return The Region in which the created address should reside.
     * If it is not provided, the provider region is used.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The URI of the created resource.
     * 
     */
    @Export(name="selfLink", type=String.class, parameters={})
    private Output<String> selfLink;

    /**
     * @return The URI of the created resource.
     * 
     */
    public Output<String> selfLink() {
        return this.selfLink;
    }
    /**
     * The URL of the subnetwork in which to reserve the address. If an IP
     * address is specified, it must be within the subnetwork&#39;s IP range.
     * This field can only be used with INTERNAL type with
     * GCE_ENDPOINT/DNS_RESOLVER purposes.
     * 
     */
    @Export(name="subnetwork", type=String.class, parameters={})
    private Output<String> subnetwork;

    /**
     * @return The URL of the subnetwork in which to reserve the address. If an IP
     * address is specified, it must be within the subnetwork&#39;s IP range.
     * This field can only be used with INTERNAL type with
     * GCE_ENDPOINT/DNS_RESOLVER purposes.
     * 
     */
    public Output<String> subnetwork() {
        return this.subnetwork;
    }
    /**
     * The URLs of the resources that are using this address.
     * 
     */
    @Export(name="users", type=List.class, parameters={String.class})
    private Output<List<String>> users;

    /**
     * @return The URLs of the resources that are using this address.
     * 
     */
    public Output<List<String>> users() {
        return this.users;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Address(String name) {
        this(name, AddressArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Address(String name, @Nullable AddressArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Address(String name, @Nullable AddressArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/address:Address", name, args == null ? AddressArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Address(String name, Output<String> id, @Nullable AddressState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/address:Address", name, state, makeResourceOptions(options, id));
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
    public static Address get(String name, Output<String> id, @Nullable AddressState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Address(name, id, state, options);
    }
}
