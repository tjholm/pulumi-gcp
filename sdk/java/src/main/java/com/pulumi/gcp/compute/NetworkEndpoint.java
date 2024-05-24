// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.compute.NetworkEndpointArgs;
import com.pulumi.gcp.compute.inputs.NetworkEndpointState;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A Network endpoint represents a IP address and port combination that is
 * part of a specific network endpoint group (NEG). NEGs are zonal
 * collections of these endpoints for GCP resources within a
 * single subnet. **NOTE**: Network endpoints cannot be created outside of a
 * network endpoint group.
 * 
 * &gt; **NOTE** In case the Endpoint&#39;s Instance is recreated, it&#39;s needed to
 * perform `apply` twice. To avoid situations like this, please use this resource
 * with the lifecycle `update_triggered_by` method, with the passed Instance&#39;s ID.
 * 
 * To get more information about NetworkEndpoint, see:
 * 
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/beta/networkEndpointGroups)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/load-balancing/docs/negs/)
 * 
 * ## Example Usage
 * 
 * ### Network Endpoint
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.ComputeFunctions;
 * import com.pulumi.gcp.compute.inputs.GetImageArgs;
 * import com.pulumi.gcp.compute.Network;
 * import com.pulumi.gcp.compute.NetworkArgs;
 * import com.pulumi.gcp.compute.Subnetwork;
 * import com.pulumi.gcp.compute.SubnetworkArgs;
 * import com.pulumi.gcp.compute.Instance;
 * import com.pulumi.gcp.compute.InstanceArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceNetworkInterfaceArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceBootDiskArgs;
 * import com.pulumi.gcp.compute.inputs.InstanceBootDiskInitializeParamsArgs;
 * import com.pulumi.gcp.compute.NetworkEndpoint;
 * import com.pulumi.gcp.compute.NetworkEndpointArgs;
 * import com.pulumi.gcp.compute.NetworkEndpointGroup;
 * import com.pulumi.gcp.compute.NetworkEndpointGroupArgs;
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
 *         final var myImage = ComputeFunctions.getImage(GetImageArgs.builder()
 *             .family("debian-11")
 *             .project("debian-cloud")
 *             .build());
 * 
 *         var default_ = new Network("default", NetworkArgs.builder()
 *             .name("neg-network")
 *             .autoCreateSubnetworks(false)
 *             .build());
 * 
 *         var defaultSubnetwork = new Subnetwork("defaultSubnetwork", SubnetworkArgs.builder()
 *             .name("neg-subnetwork")
 *             .ipCidrRange("10.0.0.1/16")
 *             .region("us-central1")
 *             .network(default_.id())
 *             .build());
 * 
 *         var endpoint_instance = new Instance("endpoint-instance", InstanceArgs.builder()
 *             .networkInterfaces(InstanceNetworkInterfaceArgs.builder()
 *                 .accessConfigs()
 *                 .subnetwork(defaultSubnetwork.id())
 *                 .build())
 *             .name("endpoint-instance")
 *             .machineType("e2-medium")
 *             .bootDisk(InstanceBootDiskArgs.builder()
 *                 .initializeParams(InstanceBootDiskInitializeParamsArgs.builder()
 *                     .image(myImage.applyValue(getImageResult -> getImageResult.selfLink()))
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *         var default_endpoint = new NetworkEndpoint("default-endpoint", NetworkEndpointArgs.builder()
 *             .networkEndpointGroup(neg.name())
 *             .instance(endpoint_instance.name())
 *             .port(neg.defaultPort())
 *             .ipAddress(endpoint_instance.networkInterfaces().applyValue(networkInterfaces -> networkInterfaces[0].networkIp()))
 *             .build());
 * 
 *         var group = new NetworkEndpointGroup("group", NetworkEndpointGroupArgs.builder()
 *             .name("my-lb-neg")
 *             .network(default_.id())
 *             .subnetwork(defaultSubnetwork.id())
 *             .defaultPort("90")
 *             .zone("us-central1-a")
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
 * NetworkEndpoint can be imported using any of these accepted formats:
 * 
 * * `projects/{{project}}/zones/{{zone}}/networkEndpointGroups/{{network_endpoint_group}}/{{instance}}/{{ip_address}}/{{port}}`
 * 
 * * `{{project}}/{{zone}}/{{network_endpoint_group}}/{{instance}}/{{ip_address}}/{{port}}`
 * 
 * * `{{zone}}/{{network_endpoint_group}}/{{instance}}/{{ip_address}}/{{port}}`
 * 
 * * `{{network_endpoint_group}}/{{instance}}/{{ip_address}}/{{port}}`
 * 
 * When using the `pulumi import` command, NetworkEndpoint can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:compute/networkEndpoint:NetworkEndpoint default projects/{{project}}/zones/{{zone}}/networkEndpointGroups/{{network_endpoint_group}}/{{instance}}/{{ip_address}}/{{port}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:compute/networkEndpoint:NetworkEndpoint default {{project}}/{{zone}}/{{network_endpoint_group}}/{{instance}}/{{ip_address}}/{{port}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:compute/networkEndpoint:NetworkEndpoint default {{zone}}/{{network_endpoint_group}}/{{instance}}/{{ip_address}}/{{port}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:compute/networkEndpoint:NetworkEndpoint default {{network_endpoint_group}}/{{instance}}/{{ip_address}}/{{port}}
 * ```
 * 
 */
@ResourceType(type="gcp:compute/networkEndpoint:NetworkEndpoint")
public class NetworkEndpoint extends com.pulumi.resources.CustomResource {
    /**
     * The name for a specific VM instance that the IP address belongs to.
     * This is required for network endpoints of type GCE_VM_IP_PORT.
     * The instance must be in the same zone of network endpoint group.
     * 
     */
    @Export(name="instance", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> instance;

    /**
     * @return The name for a specific VM instance that the IP address belongs to.
     * This is required for network endpoints of type GCE_VM_IP_PORT.
     * The instance must be in the same zone of network endpoint group.
     * 
     */
    public Output<Optional<String>> instance() {
        return Codegen.optional(this.instance);
    }
    /**
     * IPv4 address of network endpoint. The IP address must belong
     * to a VM in GCE (either the primary IP or as part of an aliased IP
     * range).
     * 
     */
    @Export(name="ipAddress", refs={String.class}, tree="[0]")
    private Output<String> ipAddress;

    /**
     * @return IPv4 address of network endpoint. The IP address must belong
     * to a VM in GCE (either the primary IP or as part of an aliased IP
     * range).
     * 
     */
    public Output<String> ipAddress() {
        return this.ipAddress;
    }
    /**
     * The network endpoint group this endpoint is part of.
     * 
     * ***
     * 
     */
    @Export(name="networkEndpointGroup", refs={String.class}, tree="[0]")
    private Output<String> networkEndpointGroup;

    /**
     * @return The network endpoint group this endpoint is part of.
     * 
     * ***
     * 
     */
    public Output<String> networkEndpointGroup() {
        return this.networkEndpointGroup;
    }
    /**
     * Port number of network endpoint.
     * **Note** `port` is required unless the Network Endpoint Group is created
     * with the type of `GCE_VM_IP`
     * 
     */
    @Export(name="port", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> port;

    /**
     * @return Port number of network endpoint.
     * **Note** `port` is required unless the Network Endpoint Group is created
     * with the type of `GCE_VM_IP`
     * 
     */
    public Output<Optional<Integer>> port() {
        return Codegen.optional(this.port);
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
     * Zone where the containing network endpoint group is located.
     * 
     */
    @Export(name="zone", refs={String.class}, tree="[0]")
    private Output<String> zone;

    /**
     * @return Zone where the containing network endpoint group is located.
     * 
     */
    public Output<String> zone() {
        return this.zone;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public NetworkEndpoint(String name) {
        this(name, NetworkEndpointArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public NetworkEndpoint(String name, NetworkEndpointArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public NetworkEndpoint(String name, NetworkEndpointArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/networkEndpoint:NetworkEndpoint", name, args == null ? NetworkEndpointArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private NetworkEndpoint(String name, Output<String> id, @Nullable NetworkEndpointState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:compute/networkEndpoint:NetworkEndpoint", name, state, makeResourceOptions(options, id));
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
    public static NetworkEndpoint get(String name, Output<String> id, @Nullable NetworkEndpointState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new NetworkEndpoint(name, id, state, options);
    }
}
