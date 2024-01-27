// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.databasemigrationservice;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.databasemigrationservice.PrivateConnectionArgs;
import com.pulumi.gcp.databasemigrationservice.inputs.PrivateConnectionState;
import com.pulumi.gcp.databasemigrationservice.outputs.PrivateConnectionError;
import com.pulumi.gcp.databasemigrationservice.outputs.PrivateConnectionVpcPeeringConfig;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * The PrivateConnection resource is used to establish private connectivity between Database Migration Service and a customer&#39;s network.
 * 
 * To get more information about PrivateConnection, see:
 * 
 * * [API documentation](https://cloud.google.com/database-migration/docs/reference/rest/v1/projects.locations.privateConnections)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/database-migration/docs/oracle-to-postgresql/create-private-connectivity-configuration)
 * 
 * ## Example Usage
 * ### Database Migration Service Private Connection
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.ComputeFunctions;
 * import com.pulumi.gcp.compute.inputs.GetNetworkArgs;
 * import com.pulumi.gcp.databasemigrationservice.PrivateConnection;
 * import com.pulumi.gcp.databasemigrationservice.PrivateConnectionArgs;
 * import com.pulumi.gcp.databasemigrationservice.inputs.PrivateConnectionVpcPeeringConfigArgs;
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
 *         final var defaultNetwork = ComputeFunctions.getNetwork(GetNetworkArgs.builder()
 *             .name(&#34;my-network&#34;)
 *             .build());
 * 
 *         var defaultPrivateConnection = new PrivateConnection(&#34;defaultPrivateConnection&#34;, PrivateConnectionArgs.builder()        
 *             .displayName(&#34;dbms_pc&#34;)
 *             .location(&#34;us-central1&#34;)
 *             .privateConnectionId(&#34;my-connection&#34;)
 *             .labels(Map.of(&#34;key&#34;, &#34;value&#34;))
 *             .vpcPeeringConfig(PrivateConnectionVpcPeeringConfigArgs.builder()
 *                 .vpcName(defaultNetwork.applyValue(getNetworkResult -&gt; getNetworkResult.id()))
 *                 .subnet(&#34;10.0.0.0/29&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * PrivateConnection can be imported using any of these accepted formats* `projects/{{project}}/locations/{{location}}/privateConnections/{{private_connection_id}}` * `{{project}}/{{location}}/{{private_connection_id}}` * `{{location}}/{{private_connection_id}}` When using the `pulumi import` command, PrivateConnection can be imported using one of the formats above. For example
 * 
 * ```sh
 *  $ pulumi import gcp:databasemigrationservice/privateConnection:PrivateConnection default projects/{{project}}/locations/{{location}}/privateConnections/{{private_connection_id}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:databasemigrationservice/privateConnection:PrivateConnection default {{project}}/{{location}}/{{private_connection_id}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:databasemigrationservice/privateConnection:PrivateConnection default {{location}}/{{private_connection_id}}
 * ```
 * 
 */
@ResourceType(type="gcp:databasemigrationservice/privateConnection:PrivateConnection")
public class PrivateConnection extends com.pulumi.resources.CustomResource {
    /**
     * Display name.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output<String> displayName;

    /**
     * @return Display name.
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
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
     * The PrivateConnection error in case of failure.
     * Structure is documented below.
     * 
     */
    @Export(name="errors", refs={List.class,PrivateConnectionError.class}, tree="[0,1]")
    private Output<List<PrivateConnectionError>> errors;

    /**
     * @return The PrivateConnection error in case of failure.
     * Structure is documented below.
     * 
     */
    public Output<List<PrivateConnectionError>> errors() {
        return this.errors;
    }
    /**
     * Labels.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return Labels.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * The name of the location this private connection is located in.
     * 
     */
    @Export(name="location", refs={String.class}, tree="[0]")
    private Output<String> location;

    /**
     * @return The name of the location this private connection is located in.
     * 
     */
    public Output<String> location() {
        return this.location;
    }
    /**
     * The resource&#39;s name.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The resource&#39;s name.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The private connectivity identifier.
     * 
     */
    @Export(name="privateConnectionId", refs={String.class}, tree="[0]")
    private Output<String> privateConnectionId;

    /**
     * @return The private connectivity identifier.
     * 
     */
    public Output<String> privateConnectionId() {
        return this.privateConnectionId;
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
     * State of the PrivateConnection.
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return State of the PrivateConnection.
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * The VPC Peering configuration is used to create VPC peering
     * between databasemigrationservice and the consumer&#39;s VPC.
     * Structure is documented below.
     * 
     */
    @Export(name="vpcPeeringConfig", refs={PrivateConnectionVpcPeeringConfig.class}, tree="[0]")
    private Output<PrivateConnectionVpcPeeringConfig> vpcPeeringConfig;

    /**
     * @return The VPC Peering configuration is used to create VPC peering
     * between databasemigrationservice and the consumer&#39;s VPC.
     * Structure is documented below.
     * 
     */
    public Output<PrivateConnectionVpcPeeringConfig> vpcPeeringConfig() {
        return this.vpcPeeringConfig;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PrivateConnection(String name) {
        this(name, PrivateConnectionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PrivateConnection(String name, PrivateConnectionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PrivateConnection(String name, PrivateConnectionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:databasemigrationservice/privateConnection:PrivateConnection", name, args == null ? PrivateConnectionArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private PrivateConnection(String name, Output<String> id, @Nullable PrivateConnectionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:databasemigrationservice/privateConnection:PrivateConnection", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
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
    public static PrivateConnection get(String name, Output<String> id, @Nullable PrivateConnectionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PrivateConnection(name, id, state, options);
    }
}
