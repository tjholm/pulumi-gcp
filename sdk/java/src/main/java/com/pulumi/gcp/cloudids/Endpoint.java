// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.cloudids;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.cloudids.EndpointArgs;
import com.pulumi.gcp.cloudids.inputs.EndpointState;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Cloud IDS is an intrusion detection service that provides threat detection for intrusions, malware, spyware, and command-and-control attacks on your network.
 * 
 * To get more information about Endpoint, see:
 * 
 * * [API documentation](https://cloud.google.com/intrusion-detection-system/docs/configuring-ids)
 * 
 * ## Example Usage
 * 
 * ### Cloudids Endpoint
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
 * import com.pulumi.gcp.cloudids.Endpoint;
 * import com.pulumi.gcp.cloudids.EndpointArgs;
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
 *         var default_ = new Network("default", NetworkArgs.builder()
 *             .name("tf-test-my-network")
 *             .build());
 * 
 *         var serviceRange = new GlobalAddress("serviceRange", GlobalAddressArgs.builder()
 *             .name("address")
 *             .purpose("VPC_PEERING")
 *             .addressType("INTERNAL")
 *             .prefixLength(16)
 *             .network(default_.id())
 *             .build());
 * 
 *         var privateServiceConnection = new Connection("privateServiceConnection", ConnectionArgs.builder()
 *             .network(default_.id())
 *             .service("servicenetworking.googleapis.com")
 *             .reservedPeeringRanges(serviceRange.name())
 *             .build());
 * 
 *         var example_endpoint = new Endpoint("example-endpoint", EndpointArgs.builder()
 *             .name("test")
 *             .location("us-central1-f")
 *             .network(default_.id())
 *             .severity("INFORMATIONAL")
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(privateServiceConnection)
 *                 .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Endpoint can be imported using any of these accepted formats:
 * 
 * * `projects/{{project}}/locations/{{location}}/endpoints/{{name}}`
 * 
 * * `{{project}}/{{location}}/{{name}}`
 * 
 * * `{{location}}/{{name}}`
 * 
 * When using the `pulumi import` command, Endpoint can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:cloudids/endpoint:Endpoint default projects/{{project}}/locations/{{location}}/endpoints/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:cloudids/endpoint:Endpoint default {{project}}/{{location}}/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:cloudids/endpoint:Endpoint default {{location}}/{{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:cloudids/endpoint:Endpoint")
public class Endpoint extends com.pulumi.resources.CustomResource {
    /**
     * Creation timestamp in RFC 3339 text format.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return Creation timestamp in RFC 3339 text format.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * An optional description of the endpoint.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return An optional description of the endpoint.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * URL of the endpoint&#39;s network address to which traffic is to be sent by Packet Mirroring.
     * 
     */
    @Export(name="endpointForwardingRule", refs={String.class}, tree="[0]")
    private Output<String> endpointForwardingRule;

    /**
     * @return URL of the endpoint&#39;s network address to which traffic is to be sent by Packet Mirroring.
     * 
     */
    public Output<String> endpointForwardingRule() {
        return this.endpointForwardingRule;
    }
    /**
     * Internal IP address of the endpoint&#39;s network entry point.
     * 
     */
    @Export(name="endpointIp", refs={String.class}, tree="[0]")
    private Output<String> endpointIp;

    /**
     * @return Internal IP address of the endpoint&#39;s network entry point.
     * 
     */
    public Output<String> endpointIp() {
        return this.endpointIp;
    }
    /**
     * The location for the endpoint.
     * 
     * ***
     * 
     */
    @Export(name="location", refs={String.class}, tree="[0]")
    private Output<String> location;

    /**
     * @return The location for the endpoint.
     * 
     * ***
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    /**
     * Name of the endpoint in the format projects/{project_id}/locations/{locationId}/endpoints/{endpointId}.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the endpoint in the format projects/{project_id}/locations/{locationId}/endpoints/{endpointId}.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Name of the VPC network that is connected to the IDS endpoint. This can either contain the VPC network name itself (like &#34;src-net&#34;) or the full URL to the network (like &#34;projects/{project_id}/global/networks/src-net&#34;).
     * 
     */
    @Export(name="network", refs={String.class}, tree="[0]")
    private Output<String> network;

    /**
     * @return Name of the VPC network that is connected to the IDS endpoint. This can either contain the VPC network name itself (like &#34;src-net&#34;) or the full URL to the network (like &#34;projects/{project_id}/global/networks/src-net&#34;).
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
     * The minimum alert severity level that is reported by the endpoint.
     * Possible values are: `INFORMATIONAL`, `LOW`, `MEDIUM`, `HIGH`, `CRITICAL`.
     * 
     */
    @Export(name="severity", refs={String.class}, tree="[0]")
    private Output<String> severity;

    /**
     * @return The minimum alert severity level that is reported by the endpoint.
     * Possible values are: `INFORMATIONAL`, `LOW`, `MEDIUM`, `HIGH`, `CRITICAL`.
     * 
     */
    public Output<String> severity() {
        return this.severity;
    }
    /**
     * Configuration for threat IDs excluded from generating alerts. Limit: 99 IDs.
     * 
     */
    @Export(name="threatExceptions", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> threatExceptions;

    /**
     * @return Configuration for threat IDs excluded from generating alerts. Limit: 99 IDs.
     * 
     */
    public Output<Optional<List<String>>> threatExceptions() {
        return Codegen.optional(this.threatExceptions);
    }
    /**
     * Last update timestamp in RFC 3339 text format.
     * 
     */
    @Export(name="updateTime", refs={String.class}, tree="[0]")
    private Output<String> updateTime;

    /**
     * @return Last update timestamp in RFC 3339 text format.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Endpoint(java.lang.String name) {
        this(name, EndpointArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Endpoint(java.lang.String name, EndpointArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Endpoint(java.lang.String name, EndpointArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:cloudids/endpoint:Endpoint", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Endpoint(java.lang.String name, Output<java.lang.String> id, @Nullable EndpointState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:cloudids/endpoint:Endpoint", name, state, makeResourceOptions(options, id), false);
    }

    private static EndpointArgs makeArgs(EndpointArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? EndpointArgs.Empty : args;
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
    public static Endpoint get(java.lang.String name, Output<java.lang.String> id, @Nullable EndpointState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Endpoint(name, id, state, options);
    }
}
