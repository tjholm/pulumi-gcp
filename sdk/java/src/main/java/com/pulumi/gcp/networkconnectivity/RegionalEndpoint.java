// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.networkconnectivity;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.networkconnectivity.RegionalEndpointArgs;
import com.pulumi.gcp.networkconnectivity.inputs.RegionalEndpointState;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Regional Private Service Connect (PSC) endpoint resource.
 * 
 * To get more information about RegionalEndpoint, see:
 * 
 * * [API documentation](https://cloud.google.com/network-connectivity/docs/reference/networkconnectivity/rest/v1/projects.locations.regionalEndpoints)
 * * How-to Guides
 *     * [Access regional Google APIs through endpoints](https://cloud.google.com/vpc/docs/access-regional-google-apis-endpoints)
 * 
 * ## Example Usage
 * 
 * ### Network Connectivity Regional Endpoint Regional Access
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
 * import com.pulumi.gcp.networkconnectivity.RegionalEndpoint;
 * import com.pulumi.gcp.networkconnectivity.RegionalEndpointArgs;
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
 *         var myNetwork = new Network("myNetwork", NetworkArgs.builder()
 *             .name("my-network")
 *             .autoCreateSubnetworks(false)
 *             .build());
 * 
 *         var mySubnetwork = new Subnetwork("mySubnetwork", SubnetworkArgs.builder()
 *             .name("my-subnetwork")
 *             .ipCidrRange("192.168.0.0/24")
 *             .region("us-central1")
 *             .network(myNetwork.id())
 *             .build());
 * 
 *         var default_ = new RegionalEndpoint("default", RegionalEndpointArgs.builder()
 *             .name("my-rep")
 *             .location("us-central1")
 *             .targetGoogleApi("boqcodelabjaimin-pa.us-central1.p.rep.googleapis.com")
 *             .accessType("REGIONAL")
 *             .address("192.168.0.5")
 *             .network(myNetwork.id())
 *             .subnetwork(mySubnetwork.id())
 *             .description("My RegionalEndpoint targeting Google API boqcodelabjaimin-pa.us-central1.p.rep.googleapis.com")
 *             .labels(Map.of("env", "default"))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Network Connectivity Regional Endpoint Global Access
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
 * import com.pulumi.gcp.networkconnectivity.RegionalEndpoint;
 * import com.pulumi.gcp.networkconnectivity.RegionalEndpointArgs;
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
 *         var myNetwork = new Network("myNetwork", NetworkArgs.builder()
 *             .name("my-network")
 *             .autoCreateSubnetworks(false)
 *             .build());
 * 
 *         var mySubnetwork = new Subnetwork("mySubnetwork", SubnetworkArgs.builder()
 *             .name("my-subnetwork")
 *             .ipCidrRange("192.168.0.0/24")
 *             .region("us-central1")
 *             .network(myNetwork.id())
 *             .build());
 * 
 *         var default_ = new RegionalEndpoint("default", RegionalEndpointArgs.builder()
 *             .name("my-rep")
 *             .location("us-central1")
 *             .targetGoogleApi("boqcodelabjaimin-pa.us-central1.p.rep.googleapis.com")
 *             .accessType("GLOBAL")
 *             .address("192.168.0.4")
 *             .network(myNetwork.id())
 *             .subnetwork(mySubnetwork.id())
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
 * RegionalEndpoint can be imported using any of these accepted formats:
 * 
 * * `projects/{{project}}/locations/{{location}}/regionalEndpoints/{{name}}`
 * 
 * * `{{project}}/{{location}}/{{name}}`
 * 
 * * `{{location}}/{{name}}`
 * 
 * When using the `pulumi import` command, RegionalEndpoint can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:networkconnectivity/regionalEndpoint:RegionalEndpoint default projects/{{project}}/locations/{{location}}/regionalEndpoints/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:networkconnectivity/regionalEndpoint:RegionalEndpoint default {{project}}/{{location}}/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:networkconnectivity/regionalEndpoint:RegionalEndpoint default {{location}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:networkconnectivity/regionalEndpoint:RegionalEndpoint")
public class RegionalEndpoint extends com.pulumi.resources.CustomResource {
    /**
     * The access type of this regional endpoint. This field is reflected in the PSC Forwarding Rule configuration to enable global access.
     * Possible values are: `GLOBAL`, `REGIONAL`.
     * 
     */
    @Export(name="accessType", refs={String.class}, tree="[0]")
    private Output<String> accessType;

    /**
     * @return The access type of this regional endpoint. This field is reflected in the PSC Forwarding Rule configuration to enable global access.
     * Possible values are: `GLOBAL`, `REGIONAL`.
     * 
     */
    public Output<String> accessType() {
        return this.accessType;
    }
    /**
     * The IP Address of the Regional Endpoint. When no address is provided, an IP from the subnetwork is allocated. Use one of the following formats: * IPv4 address as in `10.0.0.1` * Address resource URI as in `projects/{project}/regions/{region}/addresses/{address_name}`
     * &gt; **Note:** This field accepts both a reference to a Compute Address resource, which is the resource name of which format is given in the description, and IP literal value. If the user chooses to input a reserved address value; they need to make sure that the reserved address is in IPv4 version, its purpose is GCE_ENDPOINT, its type is INTERNAL and its status is RESERVED. If the user chooses to input an IP literal, they need to make sure that it&#39;s a valid IPv4 address (x.x.x.x) within the subnetwork.
     * 
     */
    @Export(name="address", refs={String.class}, tree="[0]")
    private Output<String> address;

    /**
     * @return The IP Address of the Regional Endpoint. When no address is provided, an IP from the subnetwork is allocated. Use one of the following formats: * IPv4 address as in `10.0.0.1` * Address resource URI as in `projects/{project}/regions/{region}/addresses/{address_name}`
     * &gt; **Note:** This field accepts both a reference to a Compute Address resource, which is the resource name of which format is given in the description, and IP literal value. If the user chooses to input a reserved address value; they need to make sure that the reserved address is in IPv4 version, its purpose is GCE_ENDPOINT, its type is INTERNAL and its status is RESERVED. If the user chooses to input an IP literal, they need to make sure that it&#39;s a valid IPv4 address (x.x.x.x) within the subnetwork.
     * 
     */
    public Output<String> address() {
        return this.address;
    }
    /**
     * Time when the RegionalEndpoint was created.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return Time when the RegionalEndpoint was created.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * A description of this resource.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return A description of this resource.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
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
     * User-defined labels.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return User-defined labels.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * The location of the RegionalEndpoint.
     * 
     * ***
     * 
     */
    @Export(name="location", refs={String.class}, tree="[0]")
    private Output<String> location;

    /**
     * @return The location of the RegionalEndpoint.
     * 
     * ***
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    /**
     * The name of the RegionalEndpoint.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the RegionalEndpoint.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The name of the VPC network for this private regional endpoint. Format: `projects/{project}/global/networks/{network}`
     * 
     */
    @Export(name="network", refs={String.class}, tree="[0]")
    private Output<String> network;

    /**
     * @return The name of the VPC network for this private regional endpoint. Format: `projects/{project}/global/networks/{network}`
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
     * The resource reference of the PSC Forwarding Rule created on behalf of the customer. Format: `//compute.googleapis.com/projects/{project}/regions/{region}/forwardingRules/{forwarding_rule_name}`
     * 
     */
    @Export(name="pscForwardingRule", refs={String.class}, tree="[0]")
    private Output<String> pscForwardingRule;

    /**
     * @return The resource reference of the PSC Forwarding Rule created on behalf of the customer. Format: `//compute.googleapis.com/projects/{project}/regions/{region}/forwardingRules/{forwarding_rule_name}`
     * 
     */
    public Output<String> pscForwardingRule() {
        return this.pscForwardingRule;
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
     * The name of the subnetwork from which the IP address will be allocated. Format: `projects/{project}/regions/{region}/subnetworks/{subnetwork}`
     * 
     */
    @Export(name="subnetwork", refs={String.class}, tree="[0]")
    private Output<String> subnetwork;

    /**
     * @return The name of the subnetwork from which the IP address will be allocated. Format: `projects/{project}/regions/{region}/subnetworks/{subnetwork}`
     * 
     */
    public Output<String> subnetwork() {
        return this.subnetwork;
    }
    /**
     * The service endpoint this private regional endpoint connects to. Format: `{apiname}.{region}.p.rep.googleapis.com` Example: \&#34;cloudkms.us-central1.p.rep.googleapis.com\&#34;.
     * 
     */
    @Export(name="targetGoogleApi", refs={String.class}, tree="[0]")
    private Output<String> targetGoogleApi;

    /**
     * @return The service endpoint this private regional endpoint connects to. Format: `{apiname}.{region}.p.rep.googleapis.com` Example: \&#34;cloudkms.us-central1.p.rep.googleapis.com\&#34;.
     * 
     */
    public Output<String> targetGoogleApi() {
        return this.targetGoogleApi;
    }
    /**
     * Time when the RegionalEndpoint was updated.
     * 
     */
    @Export(name="updateTime", refs={String.class}, tree="[0]")
    private Output<String> updateTime;

    /**
     * @return Time when the RegionalEndpoint was updated.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RegionalEndpoint(java.lang.String name) {
        this(name, RegionalEndpointArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RegionalEndpoint(java.lang.String name, RegionalEndpointArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RegionalEndpoint(java.lang.String name, RegionalEndpointArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:networkconnectivity/regionalEndpoint:RegionalEndpoint", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private RegionalEndpoint(java.lang.String name, Output<java.lang.String> id, @Nullable RegionalEndpointState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:networkconnectivity/regionalEndpoint:RegionalEndpoint", name, state, makeResourceOptions(options, id), false);
    }

    private static RegionalEndpointArgs makeArgs(RegionalEndpointArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? RegionalEndpointArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "effectiveLabels",
                "pulumiLabels"
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
    public static RegionalEndpoint get(java.lang.String name, Output<java.lang.String> id, @Nullable RegionalEndpointState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RegionalEndpoint(name, id, state, options);
    }
}
