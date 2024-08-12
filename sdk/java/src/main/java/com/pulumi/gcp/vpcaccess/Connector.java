// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.vpcaccess;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.vpcaccess.ConnectorArgs;
import com.pulumi.gcp.vpcaccess.inputs.ConnectorState;
import com.pulumi.gcp.vpcaccess.outputs.ConnectorSubnet;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Serverless VPC Access connector resource.
 * 
 * To get more information about Connector, see:
 * 
 * * [API documentation](https://cloud.google.com/vpc/docs/reference/vpcaccess/rest/v1/projects.locations.connectors)
 * * How-to Guides
 *     * [Configuring Serverless VPC Access](https://cloud.google.com/vpc/docs/configure-serverless-vpc-access)
 * 
 * ## Example Usage
 * 
 * ### Vpc Access Connector
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.vpcaccess.Connector;
 * import com.pulumi.gcp.vpcaccess.ConnectorArgs;
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
 *         var connector = new Connector("connector", ConnectorArgs.builder()
 *             .name("vpc-con")
 *             .ipCidrRange("10.8.0.0/28")
 *             .network("default")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * ### Vpc Access Connector Shared Vpc
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.Subnetwork;
 * import com.pulumi.gcp.compute.SubnetworkArgs;
 * import com.pulumi.gcp.vpcaccess.Connector;
 * import com.pulumi.gcp.vpcaccess.ConnectorArgs;
 * import com.pulumi.gcp.vpcaccess.inputs.ConnectorSubnetArgs;
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
 *         var customTest = new Subnetwork("customTest", SubnetworkArgs.builder()
 *             .name("vpc-con")
 *             .ipCidrRange("10.2.0.0/28")
 *             .region("us-central1")
 *             .network("default")
 *             .build());
 * 
 *         var connector = new Connector("connector", ConnectorArgs.builder()
 *             .name("vpc-con")
 *             .subnet(ConnectorSubnetArgs.builder()
 *                 .name(customTest.name())
 *                 .build())
 *             .machineType("e2-standard-4")
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
 * Connector can be imported using any of these accepted formats:
 * 
 * * `projects/{{project}}/locations/{{region}}/connectors/{{name}}`
 * 
 * * `{{project}}/{{region}}/{{name}}`
 * 
 * * `{{region}}/{{name}}`
 * 
 * * `{{name}}`
 * 
 * When using the `pulumi import` command, Connector can be imported using one of the formats above. For example:
 * 
 * ```sh
 * $ pulumi import gcp:vpcaccess/connector:Connector default projects/{{project}}/locations/{{region}}/connectors/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:vpcaccess/connector:Connector default {{project}}/{{region}}/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:vpcaccess/connector:Connector default {{region}}/{{name}}
 * ```
 * 
 * ```sh
 * $ pulumi import gcp:vpcaccess/connector:Connector default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:vpcaccess/connector:Connector")
public class Connector extends com.pulumi.resources.CustomResource {
    /**
     * List of projects using the connector.
     * 
     */
    @Export(name="connectedProjects", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> connectedProjects;

    /**
     * @return List of projects using the connector.
     * 
     */
    public Output<List<String>> connectedProjects() {
        return this.connectedProjects;
    }
    /**
     * The range of internal addresses that follows RFC 4632 notation. Example: `10.132.0.0/28`.
     * 
     */
    @Export(name="ipCidrRange", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ipCidrRange;

    /**
     * @return The range of internal addresses that follows RFC 4632 notation. Example: `10.132.0.0/28`.
     * 
     */
    public Output<Optional<String>> ipCidrRange() {
        return Codegen.optional(this.ipCidrRange);
    }
    /**
     * Machine type of VM Instance underlying connector. Default is e2-micro
     * 
     */
    @Export(name="machineType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> machineType;

    /**
     * @return Machine type of VM Instance underlying connector. Default is e2-micro
     * 
     */
    public Output<Optional<String>> machineType() {
        return Codegen.optional(this.machineType);
    }
    /**
     * Maximum value of instances in autoscaling group underlying the connector. Value must be between 3 and 10, inclusive. Must be
     * higher than the value specified by min_instances.
     * 
     */
    @Export(name="maxInstances", refs={Integer.class}, tree="[0]")
    private Output<Integer> maxInstances;

    /**
     * @return Maximum value of instances in autoscaling group underlying the connector. Value must be between 3 and 10, inclusive. Must be
     * higher than the value specified by min_instances.
     * 
     */
    public Output<Integer> maxInstances() {
        return this.maxInstances;
    }
    /**
     * Maximum throughput of the connector in Mbps, must be greater than `min_throughput`. Default is 300. Refers to the expected throughput
     * when using an e2-micro machine type. Value must be a multiple of 100 from 300 through 1000. Must be higher than the value specified by
     * min_throughput. If both max_throughput and max_instances are provided, max_instances takes precedence over max_throughput. The use of
     * max_throughput is discouraged in favor of max_instances.
     * 
     */
    @Export(name="maxThroughput", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> maxThroughput;

    /**
     * @return Maximum throughput of the connector in Mbps, must be greater than `min_throughput`. Default is 300. Refers to the expected throughput
     * when using an e2-micro machine type. Value must be a multiple of 100 from 300 through 1000. Must be higher than the value specified by
     * min_throughput. If both max_throughput and max_instances are provided, max_instances takes precedence over max_throughput. The use of
     * max_throughput is discouraged in favor of max_instances.
     * 
     */
    public Output<Optional<Integer>> maxThroughput() {
        return Codegen.optional(this.maxThroughput);
    }
    /**
     * Minimum value of instances in autoscaling group underlying the connector. Value must be between 2 and 9, inclusive. Must be
     * lower than the value specified by max_instances.
     * 
     */
    @Export(name="minInstances", refs={Integer.class}, tree="[0]")
    private Output<Integer> minInstances;

    /**
     * @return Minimum value of instances in autoscaling group underlying the connector. Value must be between 2 and 9, inclusive. Must be
     * lower than the value specified by max_instances.
     * 
     */
    public Output<Integer> minInstances() {
        return this.minInstances;
    }
    /**
     * Minimum throughput of the connector in Mbps. Default and min is 200. Refers to the expected throughput when using an e2-micro machine type.
     * Value must be a multiple of 100 from 200 through 900. Must be lower than the value specified by max_throughput. If both min_throughput and
     * min_instances are provided, min_instances takes precedence over min_throughput. The use of min_throughput is discouraged in favor of min_instances.
     * 
     */
    @Export(name="minThroughput", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> minThroughput;

    /**
     * @return Minimum throughput of the connector in Mbps. Default and min is 200. Refers to the expected throughput when using an e2-micro machine type.
     * Value must be a multiple of 100 from 200 through 900. Must be lower than the value specified by max_throughput. If both min_throughput and
     * min_instances are provided, min_instances takes precedence over min_throughput. The use of min_throughput is discouraged in favor of min_instances.
     * 
     */
    public Output<Optional<Integer>> minThroughput() {
        return Codegen.optional(this.minThroughput);
    }
    /**
     * The name of the resource (Max 25 characters).
     * 
     * ***
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the resource (Max 25 characters).
     * 
     * ***
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Name or self_link of the VPC network. Required if `ip_cidr_range` is set.
     * 
     */
    @Export(name="network", refs={String.class}, tree="[0]")
    private Output<String> network;

    /**
     * @return Name or self_link of the VPC network. Required if `ip_cidr_range` is set.
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
     * Region where the VPC Access connector resides. If it is not provided, the provider region is used.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return Region where the VPC Access connector resides. If it is not provided, the provider region is used.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The fully qualified name of this VPC connector
     * 
     */
    @Export(name="selfLink", refs={String.class}, tree="[0]")
    private Output<String> selfLink;

    /**
     * @return The fully qualified name of this VPC connector
     * 
     */
    public Output<String> selfLink() {
        return this.selfLink;
    }
    /**
     * State of the VPC access connector.
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return State of the VPC access connector.
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * The subnet in which to house the connector
     * Structure is documented below.
     * 
     */
    @Export(name="subnet", refs={ConnectorSubnet.class}, tree="[0]")
    private Output</* @Nullable */ ConnectorSubnet> subnet;

    /**
     * @return The subnet in which to house the connector
     * Structure is documented below.
     * 
     */
    public Output<Optional<ConnectorSubnet>> subnet() {
        return Codegen.optional(this.subnet);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Connector(java.lang.String name) {
        this(name, ConnectorArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Connector(java.lang.String name, @Nullable ConnectorArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Connector(java.lang.String name, @Nullable ConnectorArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:vpcaccess/connector:Connector", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Connector(java.lang.String name, Output<java.lang.String> id, @Nullable ConnectorState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:vpcaccess/connector:Connector", name, state, makeResourceOptions(options, id), false);
    }

    private static ConnectorArgs makeArgs(@Nullable ConnectorArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ConnectorArgs.Empty : args;
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
    public static Connector get(java.lang.String name, Output<java.lang.String> id, @Nullable ConnectorState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Connector(name, id, state, options);
    }
}
