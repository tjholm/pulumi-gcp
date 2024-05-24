// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.compute.InterconnectAttachmentArgs;
import com.pulumi.gcp.compute.inputs.InterconnectAttachmentState;
import com.pulumi.gcp.compute.outputs.InterconnectAttachmentPrivateInterconnectInfo;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Represents an InterconnectAttachment (VLAN attachment) resource. For more
 * information, see Creating VLAN Attachments.
 * 
 * ## Example Usage
 * 
 * ### Interconnect Attachment Basic
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
 * import com.pulumi.gcp.compute.Router;
 * import com.pulumi.gcp.compute.RouterArgs;
 * import com.pulumi.gcp.compute.inputs.RouterBgpArgs;
 * import com.pulumi.gcp.compute.InterconnectAttachment;
 * import com.pulumi.gcp.compute.InterconnectAttachmentArgs;
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
 *         var foobarNetwork = new Network("foobarNetwork", NetworkArgs.builder()
 *             .name("network-1")
 *             .autoCreateSubnetworks(false)
 *             .build());
 * 
 *         var foobar = new Router("foobar", RouterArgs.builder()
 *             .name("router-1")
 *             .network(foobarNetwork.name())
 *             .bgp(RouterBgpArgs.builder()
 *                 .asn(16550)
 *                 .build())
 *             .build());
 * 
 *         var onPrem = new InterconnectAttachment("onPrem", InterconnectAttachmentArgs.builder()
 *             .name("on-prem-attachment")
 *             .edgeAvailabilityDomain("AVAILABILITY_DOMAIN_1")
 *             .type("PARTNER")
 *             .router(foobar.id())
 *             .mtu(1500)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Compute Interconnect Attachment Ipsec Encryption
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
 * import com.pulumi.gcp.compute.Address;
 * import com.pulumi.gcp.compute.AddressArgs;
 * import com.pulumi.gcp.compute.Router;
 * import com.pulumi.gcp.compute.RouterArgs;
 * import com.pulumi.gcp.compute.inputs.RouterBgpArgs;
 * import com.pulumi.gcp.compute.InterconnectAttachment;
 * import com.pulumi.gcp.compute.InterconnectAttachmentArgs;
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
 *         var network = new Network("network", NetworkArgs.builder()
 *             .name("test-network")
 *             .autoCreateSubnetworks(false)
 *             .build());
 * 
 *         var address = new Address("address", AddressArgs.builder()
 *             .name("test-address")
 *             .addressType("INTERNAL")
 *             .purpose("IPSEC_INTERCONNECT")
 *             .address("192.168.1.0")
 *             .prefixLength(29)
 *             .network(network.selfLink())
 *             .build());
 * 
 *         var router = new Router("router", RouterArgs.builder()
 *             .name("test-router")
 *             .network(network.name())
 *             .encryptedInterconnectRouter(true)
 *             .bgp(RouterBgpArgs.builder()
 *                 .asn(16550)
 *                 .build())
 *             .build());
 * 
 *         var ipsec_encrypted_interconnect_attachment = new InterconnectAttachment("ipsec-encrypted-interconnect-attachment", InterconnectAttachmentArgs.builder()
 *             .name("test-interconnect-attachment")
 *             .edgeAvailabilityDomain("AVAILABILITY_DOMAIN_1")
 *             .type("PARTNER")
 *             .router(router.id())
 *             .encryption("IPSEC")
 *             .ipsecInternalAddresses(address.selfLink())
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
 * InterconnectAttachment can be imported using any of these accepted formats:
 * 
 * * `projects/{{project}}/regions/{{region}}/interconnectAttachments/{{name}}`
 * 
 * * `{{project}}/{{region}}/{{name}}`
 * 
 * * `{{region}}/{{name}}`
 * 
 * * `{{name}}`
 * 
 * When using the `pulumi import` command, InterconnectAttachment can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:compute/interconnectAttachment:InterconnectAttachment default projects/{{project}}/regions/{{region}}/interconnectAttachments/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:compute/interconnectAttachment:InterconnectAttachment default {{project}}/{{region}}/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:compute/interconnectAttachment:InterconnectAttachment default {{region}}/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:compute/interconnectAttachment:InterconnectAttachment default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:compute/interconnectAttachment:InterconnectAttachment")
public class InterconnectAttachment extends com.pulumi.resources.CustomResource {
    /**
     * Whether the VLAN attachment is enabled or disabled.  When using
     * PARTNER type this will Pre-Activate the interconnect attachment
     * 
     */
    @Export(name="adminEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> adminEnabled;

    /**
     * @return Whether the VLAN attachment is enabled or disabled.  When using
     * PARTNER type this will Pre-Activate the interconnect attachment
     * 
     */
    public Output<Optional<Boolean>> adminEnabled() {
        return Codegen.optional(this.adminEnabled);
    }
    /**
     * Provisioned bandwidth capacity for the interconnect attachment.
     * For attachments of type DEDICATED, the user can set the bandwidth.
     * For attachments of type PARTNER, the Google Partner that is operating the interconnect must set the bandwidth.
     * Output only for PARTNER type, mutable for PARTNER_PROVIDER and DEDICATED,
     * Defaults to BPS_10G
     * Possible values are: `BPS_50M`, `BPS_100M`, `BPS_200M`, `BPS_300M`, `BPS_400M`, `BPS_500M`, `BPS_1G`, `BPS_2G`, `BPS_5G`, `BPS_10G`, `BPS_20G`, `BPS_50G`.
     * 
     */
    @Export(name="bandwidth", refs={String.class}, tree="[0]")
    private Output<String> bandwidth;

    /**
     * @return Provisioned bandwidth capacity for the interconnect attachment.
     * For attachments of type DEDICATED, the user can set the bandwidth.
     * For attachments of type PARTNER, the Google Partner that is operating the interconnect must set the bandwidth.
     * Output only for PARTNER type, mutable for PARTNER_PROVIDER and DEDICATED,
     * Defaults to BPS_10G
     * Possible values are: `BPS_50M`, `BPS_100M`, `BPS_200M`, `BPS_300M`, `BPS_400M`, `BPS_500M`, `BPS_1G`, `BPS_2G`, `BPS_5G`, `BPS_10G`, `BPS_20G`, `BPS_50G`.
     * 
     */
    public Output<String> bandwidth() {
        return this.bandwidth;
    }
    /**
     * Up to 16 candidate prefixes that can be used to restrict the allocation
     * of cloudRouterIpAddress and customerRouterIpAddress for this attachment.
     * All prefixes must be within link-local address space (169.254.0.0/16)
     * and must be /29 or shorter (/28, /27, etc). Google will attempt to select
     * an unused /29 from the supplied candidate prefix(es). The request will
     * fail if all possible /29s are in use on Google&#39;s edge. If not supplied,
     * Google will randomly select an unused /29 from all of link-local space.
     * 
     */
    @Export(name="candidateSubnets", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> candidateSubnets;

    /**
     * @return Up to 16 candidate prefixes that can be used to restrict the allocation
     * of cloudRouterIpAddress and customerRouterIpAddress for this attachment.
     * All prefixes must be within link-local address space (169.254.0.0/16)
     * and must be /29 or shorter (/28, /27, etc). Google will attempt to select
     * an unused /29 from the supplied candidate prefix(es). The request will
     * fail if all possible /29s are in use on Google&#39;s edge. If not supplied,
     * Google will randomly select an unused /29 from all of link-local space.
     * 
     */
    public Output<Optional<List<String>>> candidateSubnets() {
        return Codegen.optional(this.candidateSubnets);
    }
    /**
     * IPv4 address + prefix length to be configured on Cloud Router
     * Interface for this interconnect attachment.
     * 
     */
    @Export(name="cloudRouterIpAddress", refs={String.class}, tree="[0]")
    private Output<String> cloudRouterIpAddress;

    /**
     * @return IPv4 address + prefix length to be configured on Cloud Router
     * Interface for this interconnect attachment.
     * 
     */
    public Output<String> cloudRouterIpAddress() {
        return this.cloudRouterIpAddress;
    }
    /**
     * IPv6 address + prefix length to be configured on Cloud Router
     * Interface for this interconnect attachment.
     * 
     */
    @Export(name="cloudRouterIpv6Address", refs={String.class}, tree="[0]")
    private Output<String> cloudRouterIpv6Address;

    /**
     * @return IPv6 address + prefix length to be configured on Cloud Router
     * Interface for this interconnect attachment.
     * 
     */
    public Output<String> cloudRouterIpv6Address() {
        return this.cloudRouterIpv6Address;
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
     * IPv4 address + prefix length to be configured on the customer
     * router subinterface for this interconnect attachment.
     * 
     */
    @Export(name="customerRouterIpAddress", refs={String.class}, tree="[0]")
    private Output<String> customerRouterIpAddress;

    /**
     * @return IPv4 address + prefix length to be configured on the customer
     * router subinterface for this interconnect attachment.
     * 
     */
    public Output<String> customerRouterIpAddress() {
        return this.customerRouterIpAddress;
    }
    /**
     * IPv6 address + prefix length to be configured on the customer
     * router subinterface for this interconnect attachment.
     * 
     */
    @Export(name="customerRouterIpv6Address", refs={String.class}, tree="[0]")
    private Output<String> customerRouterIpv6Address;

    /**
     * @return IPv6 address + prefix length to be configured on the customer
     * router subinterface for this interconnect attachment.
     * 
     */
    public Output<String> customerRouterIpv6Address() {
        return this.customerRouterIpv6Address;
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
     * Desired availability domain for the attachment. Only available for type
     * PARTNER, at creation time. For improved reliability, customers should
     * configure a pair of attachments with one per availability domain. The
     * selected availability domain will be provided to the Partner via the
     * pairing key so that the provisioned circuit will lie in the specified
     * domain. If not specified, the value will default to AVAILABILITY_DOMAIN_ANY.
     * 
     */
    @Export(name="edgeAvailabilityDomain", refs={String.class}, tree="[0]")
    private Output<String> edgeAvailabilityDomain;

    /**
     * @return Desired availability domain for the attachment. Only available for type
     * PARTNER, at creation time. For improved reliability, customers should
     * configure a pair of attachments with one per availability domain. The
     * selected availability domain will be provided to the Partner via the
     * pairing key so that the provisioned circuit will lie in the specified
     * domain. If not specified, the value will default to AVAILABILITY_DOMAIN_ANY.
     * 
     */
    public Output<String> edgeAvailabilityDomain() {
        return this.edgeAvailabilityDomain;
    }
    /**
     * Indicates the user-supplied encryption option of this interconnect
     * attachment. Can only be specified at attachment creation for PARTNER or
     * DEDICATED attachments.
     * * NONE - This is the default value, which means that the VLAN attachment
     *   carries unencrypted traffic. VMs are able to send traffic to, or receive
     *   traffic from, such a VLAN attachment.
     * * IPSEC - The VLAN attachment carries only encrypted traffic that is
     *   encrypted by an IPsec device, such as an HA VPN gateway or third-party
     *   IPsec VPN. VMs cannot directly send traffic to, or receive traffic from,
     *   such a VLAN attachment. To use HA VPN over Cloud Interconnect, the VLAN
     *   attachment must be created with this option.
     *   Default value is `NONE`.
     *   Possible values are: `NONE`, `IPSEC`.
     * 
     */
    @Export(name="encryption", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> encryption;

    /**
     * @return Indicates the user-supplied encryption option of this interconnect
     * attachment. Can only be specified at attachment creation for PARTNER or
     * DEDICATED attachments.
     * * NONE - This is the default value, which means that the VLAN attachment
     *   carries unencrypted traffic. VMs are able to send traffic to, or receive
     *   traffic from, such a VLAN attachment.
     * * IPSEC - The VLAN attachment carries only encrypted traffic that is
     *   encrypted by an IPsec device, such as an HA VPN gateway or third-party
     *   IPsec VPN. VMs cannot directly send traffic to, or receive traffic from,
     *   such a VLAN attachment. To use HA VPN over Cloud Interconnect, the VLAN
     *   attachment must be created with this option.
     *   Default value is `NONE`.
     *   Possible values are: `NONE`, `IPSEC`.
     * 
     */
    public Output<Optional<String>> encryption() {
        return Codegen.optional(this.encryption);
    }
    /**
     * Google reference ID, to be used when raising support tickets with
     * Google or otherwise to debug backend connectivity issues.
     * 
     */
    @Export(name="googleReferenceId", refs={String.class}, tree="[0]")
    private Output<String> googleReferenceId;

    /**
     * @return Google reference ID, to be used when raising support tickets with
     * Google or otherwise to debug backend connectivity issues.
     * 
     */
    public Output<String> googleReferenceId() {
        return this.googleReferenceId;
    }
    /**
     * URL of the underlying Interconnect object that this attachment&#39;s
     * traffic will traverse through. Required if type is DEDICATED, must not
     * be set if type is PARTNER.
     * 
     */
    @Export(name="interconnect", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> interconnect;

    /**
     * @return URL of the underlying Interconnect object that this attachment&#39;s
     * traffic will traverse through. Required if type is DEDICATED, must not
     * be set if type is PARTNER.
     * 
     */
    public Output<Optional<String>> interconnect() {
        return Codegen.optional(this.interconnect);
    }
    /**
     * URL of addresses that have been reserved for the interconnect attachment,
     * Used only for interconnect attachment that has the encryption option as
     * IPSEC.
     * The addresses must be RFC 1918 IP address ranges. When creating HA VPN
     * gateway over the interconnect attachment, if the attachment is configured
     * to use an RFC 1918 IP address, then the VPN gateway&#39;s IP address will be
     * allocated from the IP address range specified here.
     * For example, if the HA VPN gateway&#39;s interface 0 is paired to this
     * interconnect attachment, then an RFC 1918 IP address for the VPN gateway
     * interface 0 will be allocated from the IP address specified for this
     * interconnect attachment.
     * If this field is not specified for interconnect attachment that has
     * encryption option as IPSEC, later on when creating HA VPN gateway on this
     * interconnect attachment, the HA VPN gateway&#39;s IP address will be
     * allocated from regional external IP address pool.
     * 
     */
    @Export(name="ipsecInternalAddresses", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> ipsecInternalAddresses;

    /**
     * @return URL of addresses that have been reserved for the interconnect attachment,
     * Used only for interconnect attachment that has the encryption option as
     * IPSEC.
     * The addresses must be RFC 1918 IP address ranges. When creating HA VPN
     * gateway over the interconnect attachment, if the attachment is configured
     * to use an RFC 1918 IP address, then the VPN gateway&#39;s IP address will be
     * allocated from the IP address range specified here.
     * For example, if the HA VPN gateway&#39;s interface 0 is paired to this
     * interconnect attachment, then an RFC 1918 IP address for the VPN gateway
     * interface 0 will be allocated from the IP address specified for this
     * interconnect attachment.
     * If this field is not specified for interconnect attachment that has
     * encryption option as IPSEC, later on when creating HA VPN gateway on this
     * interconnect attachment, the HA VPN gateway&#39;s IP address will be
     * allocated from regional external IP address pool.
     * 
     */
    public Output<Optional<List<String>>> ipsecInternalAddresses() {
        return Codegen.optional(this.ipsecInternalAddresses);
    }
    /**
     * Maximum Transmission Unit (MTU), in bytes, of packets passing through
     * this interconnect attachment. Currently, only 1440 and 1500 are allowed. If not specified, the value will default to 1440.
     * 
     */
    @Export(name="mtu", refs={String.class}, tree="[0]")
    private Output<String> mtu;

    /**
     * @return Maximum Transmission Unit (MTU), in bytes, of packets passing through
     * this interconnect attachment. Currently, only 1440 and 1500 are allowed. If not specified, the value will default to 1440.
     * 
     */
    public Output<String> mtu() {
        return this.mtu;
    }
    /**
     * Name of the resource. Provided by the client when the resource is created. The
     * name must be 1-63 characters long, and comply with RFC1035. Specifically, the
     * name must be 1-63 characters long and match the regular expression
     * `a-z?` which means the first character must be a
     * lowercase letter, and all following characters must be a dash, lowercase
     * letter, or digit, except the last character, which cannot be a dash.
     * 
     * ***
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the resource. Provided by the client when the resource is created. The
     * name must be 1-63 characters long, and comply with RFC1035. Specifically, the
     * name must be 1-63 characters long and match the regular expression
     * `a-z?` which means the first character must be a
     * lowercase letter, and all following characters must be a dash, lowercase
     * letter, or digit, except the last character, which cannot be a dash.
     * 
     * ***
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * [Output only for type PARTNER. Not present for DEDICATED]. The opaque
     * identifier of an PARTNER attachment used to initiate provisioning with
     * a selected partner. Of the form &#34;XXXXX/region/domain&#34;
     * 
     */
    @Export(name="pairingKey", refs={String.class}, tree="[0]")
    private Output<String> pairingKey;

    /**
     * @return [Output only for type PARTNER. Not present for DEDICATED]. The opaque
     * identifier of an PARTNER attachment used to initiate provisioning with
     * a selected partner. Of the form &#34;XXXXX/region/domain&#34;
     * 
     */
    public Output<String> pairingKey() {
        return this.pairingKey;
    }
    /**
     * [Output only for type PARTNER. Not present for DEDICATED]. Optional
     * BGP ASN for the router that should be supplied by a layer 3 Partner if
     * they configured BGP on behalf of the customer.
     * 
     */
    @Export(name="partnerAsn", refs={String.class}, tree="[0]")
    private Output<String> partnerAsn;

    /**
     * @return [Output only for type PARTNER. Not present for DEDICATED]. Optional
     * BGP ASN for the router that should be supplied by a layer 3 Partner if
     * they configured BGP on behalf of the customer.
     * 
     */
    public Output<String> partnerAsn() {
        return this.partnerAsn;
    }
    /**
     * Information specific to an InterconnectAttachment. This property
     * is populated if the interconnect that this is attached to is of type DEDICATED.
     * Structure is documented below.
     * 
     */
    @Export(name="privateInterconnectInfos", refs={List.class,InterconnectAttachmentPrivateInterconnectInfo.class}, tree="[0,1]")
    private Output<List<InterconnectAttachmentPrivateInterconnectInfo>> privateInterconnectInfos;

    /**
     * @return Information specific to an InterconnectAttachment. This property
     * is populated if the interconnect that this is attached to is of type DEDICATED.
     * Structure is documented below.
     * 
     */
    public Output<List<InterconnectAttachmentPrivateInterconnectInfo>> privateInterconnectInfos() {
        return this.privateInterconnectInfos;
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
     * Region where the regional interconnect attachment resides.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return Region where the regional interconnect attachment resides.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * URL of the cloud router to be used for dynamic routing. This router must be in
     * the same region as this InterconnectAttachment. The InterconnectAttachment will
     * automatically connect the Interconnect to the network &amp; region within which the
     * Cloud Router is configured.
     * 
     */
    @Export(name="router", refs={String.class}, tree="[0]")
    private Output<String> router;

    /**
     * @return URL of the cloud router to be used for dynamic routing. This router must be in
     * the same region as this InterconnectAttachment. The InterconnectAttachment will
     * automatically connect the Interconnect to the network &amp; region within which the
     * Cloud Router is configured.
     * 
     */
    public Output<String> router() {
        return this.router;
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
     * The stack type for this interconnect attachment to identify whether the IPv6
     * feature is enabled or not. If not specified, IPV4_ONLY will be used.
     * This field can be both set at interconnect attachments creation and update
     * interconnect attachment operations.
     * Possible values are: `IPV4_IPV6`, `IPV4_ONLY`.
     * 
     */
    @Export(name="stackType", refs={String.class}, tree="[0]")
    private Output<String> stackType;

    /**
     * @return The stack type for this interconnect attachment to identify whether the IPv6
     * feature is enabled or not. If not specified, IPV4_ONLY will be used.
     * This field can be both set at interconnect attachments creation and update
     * interconnect attachment operations.
     * Possible values are: `IPV4_IPV6`, `IPV4_ONLY`.
     * 
     */
    public Output<String> stackType() {
        return this.stackType;
    }
    /**
     * [Output Only] The current state of this attachment&#39;s functionality.
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return [Output Only] The current state of this attachment&#39;s functionality.
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * The type of InterconnectAttachment you wish to create. Defaults to
     * DEDICATED.
     * Possible values are: `DEDICATED`, `PARTNER`, `PARTNER_PROVIDER`.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return The type of InterconnectAttachment you wish to create. Defaults to
     * DEDICATED.
     * Possible values are: `DEDICATED`, `PARTNER`, `PARTNER_PROVIDER`.
     * 
     */
    public Output<String> type() {
        return this.type;
    }
    /**
     * The IEEE 802.1Q VLAN tag for this attachment, in the range 2-4094. When
     * using PARTNER type this will be managed upstream.
     * 
     */
    @Export(name="vlanTag8021q", refs={Integer.class}, tree="[0]")
    private Output<Integer> vlanTag8021q;

    /**
     * @return The IEEE 802.1Q VLAN tag for this attachment, in the range 2-4094. When
     * using PARTNER type this will be managed upstream.
     * 
     */
    public Output<Integer> vlanTag8021q() {
        return this.vlanTag8021q;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public InterconnectAttachment(String name) {
        this(name, InterconnectAttachmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public InterconnectAttachment(String name, InterconnectAttachmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public InterconnectAttachment(String name, InterconnectAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/interconnectAttachment:InterconnectAttachment", name, args == null ? InterconnectAttachmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private InterconnectAttachment(String name, Output<String> id, @Nullable InterconnectAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/interconnectAttachment:InterconnectAttachment", name, state, makeResourceOptions(options, id));
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
    public static InterconnectAttachment get(String name, Output<String> id, @Nullable InterconnectAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new InterconnectAttachment(name, id, state, options);
    }
}
