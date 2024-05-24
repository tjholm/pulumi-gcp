// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.servicenetworking;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.servicenetworking.ConnectionArgs;
import com.pulumi.gcp.servicenetworking.inputs.ConnectionState;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a private VPC connection with a GCP service provider. For more information see
 * [the official documentation](https://cloud.google.com/vpc/docs/configure-private-services-access#creating-connection)
 * and
 * [API](https://cloud.google.com/service-infrastructure/docs/service-networking/reference/rest/v1/services.connections).
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
 * import com.pulumi.gcp.compute.GlobalAddress;
 * import com.pulumi.gcp.compute.GlobalAddressArgs;
 * import com.pulumi.gcp.servicenetworking.Connection;
 * import com.pulumi.gcp.servicenetworking.ConnectionArgs;
 * import com.pulumi.gcp.compute.NetworkPeeringRoutesConfig;
 * import com.pulumi.gcp.compute.NetworkPeeringRoutesConfigArgs;
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
 *         // Create a VPC network
 *         var peeringNetwork = new Network("peeringNetwork", NetworkArgs.builder()
 *             .name("peering-network")
 *             .build());
 * 
 *         // Create an IP address
 *         var privateIpAlloc = new GlobalAddress("privateIpAlloc", GlobalAddressArgs.builder()
 *             .name("private-ip-alloc")
 *             .purpose("VPC_PEERING")
 *             .addressType("INTERNAL")
 *             .prefixLength(16)
 *             .network(peeringNetwork.id())
 *             .build());
 * 
 *         // Create a private connection
 *         var default_ = new Connection("default", ConnectionArgs.builder()
 *             .network(peeringNetwork.id())
 *             .service("servicenetworking.googleapis.com")
 *             .reservedPeeringRanges(privateIpAlloc.name())
 *             .build());
 * 
 *         // (Optional) Import or export custom routes
 *         var peeringRoutes = new NetworkPeeringRoutesConfig("peeringRoutes", NetworkPeeringRoutesConfigArgs.builder()
 *             .peering(default_.peering())
 *             .network(peeringNetwork.name())
 *             .importCustomRoutes(true)
 *             .exportCustomRoutes(true)
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
 * ServiceNetworkingConnection can be imported using any of these accepted formats
 * 
 * * `{{peering-network}}:{{service}}`
 * 
 * * `projects/{{project}}/global/networks/{{peering-network}}:{{service}}`
 * 
 * When using the `pulumi import` command, NAME_HERE can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:servicenetworking/connection:Connection default {{peering-network}}:{{service}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:servicenetworking/connection:Connection default /projects/{{project}}/global/networks/{{peering-network}}:{{service}}
 * ```
 * 
 */
@ResourceType(type="gcp:servicenetworking/connection:Connection")
public class Connection extends com.pulumi.resources.CustomResource {
    @Export(name="deletionPolicy", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> deletionPolicy;

    public Output<Optional<String>> deletionPolicy() {
        return Codegen.optional(this.deletionPolicy);
    }
    /**
     * Name of VPC network connected with service producers using VPC peering.
     * 
     */
    @Export(name="network", refs={String.class}, tree="[0]")
    private Output<String> network;

    /**
     * @return Name of VPC network connected with service producers using VPC peering.
     * 
     */
    public Output<String> network() {
        return this.network;
    }
    /**
     * (Computed) The name of the VPC Network Peering connection that was created by the service producer.
     * 
     */
    @Export(name="peering", refs={String.class}, tree="[0]")
    private Output<String> peering;

    /**
     * @return (Computed) The name of the VPC Network Peering connection that was created by the service producer.
     * 
     */
    public Output<String> peering() {
        return this.peering;
    }
    /**
     * Named IP address range(s) of PEERING type reserved for
     * this service provider. Note that invoking this method with a different range when connection
     * is already established will not reallocate already provisioned service producer subnetworks.
     * 
     */
    @Export(name="reservedPeeringRanges", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> reservedPeeringRanges;

    /**
     * @return Named IP address range(s) of PEERING type reserved for
     * this service provider. Note that invoking this method with a different range when connection
     * is already established will not reallocate already provisioned service producer subnetworks.
     * 
     */
    public Output<List<String>> reservedPeeringRanges() {
        return this.reservedPeeringRanges;
    }
    /**
     * Provider peering service that is managing peering connectivity for a
     * service provider organization. For Google services that support this functionality it is
     * &#39;servicenetworking.googleapis.com&#39;.
     * 
     */
    @Export(name="service", refs={String.class}, tree="[0]")
    private Output<String> service;

    /**
     * @return Provider peering service that is managing peering connectivity for a
     * service provider organization. For Google services that support this functionality it is
     * &#39;servicenetworking.googleapis.com&#39;.
     * 
     */
    public Output<String> service() {
        return this.service;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Connection(String name) {
        this(name, ConnectionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Connection(String name, ConnectionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Connection(String name, ConnectionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:servicenetworking/connection:Connection", name, args == null ? ConnectionArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Connection(String name, Output<String> id, @Nullable ConnectionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:servicenetworking/connection:Connection", name, state, makeResourceOptions(options, id));
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
    public static Connection get(String name, Output<String> id, @Nullable ConnectionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Connection(name, id, state, options);
    }
}
