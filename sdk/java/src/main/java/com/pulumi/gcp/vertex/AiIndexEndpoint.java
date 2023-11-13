// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.vertex;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.vertex.AiIndexEndpointArgs;
import com.pulumi.gcp.vertex.inputs.AiIndexEndpointState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * An endpoint indexes are deployed into. An index endpoint can have multiple deployed indexes.
 * 
 * To get more information about IndexEndpoint, see:
 * 
 * * [API documentation](https://cloud.google.com/vertex-ai/docs/reference/rest/v1/projects.locations.indexEndpoints/)
 * 
 * ## Example Usage
 * ### Vertex Ai Index Endpoint
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.compute.Network;
 * import com.pulumi.gcp.compute.GlobalAddress;
 * import com.pulumi.gcp.compute.GlobalAddressArgs;
 * import com.pulumi.gcp.servicenetworking.Connection;
 * import com.pulumi.gcp.servicenetworking.ConnectionArgs;
 * import com.pulumi.gcp.organizations.OrganizationsFunctions;
 * import com.pulumi.gcp.organizations.inputs.GetProjectArgs;
 * import com.pulumi.gcp.vertex.AiIndexEndpoint;
 * import com.pulumi.gcp.vertex.AiIndexEndpointArgs;
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
 *         var vertexNetwork = new Network(&#34;vertexNetwork&#34;);
 * 
 *         var vertexRange = new GlobalAddress(&#34;vertexRange&#34;, GlobalAddressArgs.builder()        
 *             .purpose(&#34;VPC_PEERING&#34;)
 *             .addressType(&#34;INTERNAL&#34;)
 *             .prefixLength(24)
 *             .network(vertexNetwork.id())
 *             .build());
 * 
 *         var vertexVpcConnection = new Connection(&#34;vertexVpcConnection&#34;, ConnectionArgs.builder()        
 *             .network(vertexNetwork.id())
 *             .service(&#34;servicenetworking.googleapis.com&#34;)
 *             .reservedPeeringRanges(vertexRange.name())
 *             .build());
 * 
 *         final var project = OrganizationsFunctions.getProject();
 * 
 *         var indexEndpoint = new AiIndexEndpoint(&#34;indexEndpoint&#34;, AiIndexEndpointArgs.builder()        
 *             .displayName(&#34;sample-endpoint&#34;)
 *             .description(&#34;A sample vertex endpoint&#34;)
 *             .region(&#34;us-central1&#34;)
 *             .labels(Map.of(&#34;label-one&#34;, &#34;value-one&#34;))
 *             .network(vertexNetwork.name().applyValue(name -&gt; String.format(&#34;projects/%s/global/networks/%s&#34;, project.applyValue(getProjectResult -&gt; getProjectResult.number()),name)))
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(vertexVpcConnection)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * ### Vertex Ai Index Endpoint With Public Endpoint
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.vertex.AiIndexEndpoint;
 * import com.pulumi.gcp.vertex.AiIndexEndpointArgs;
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
 *         var indexEndpoint = new AiIndexEndpoint(&#34;indexEndpoint&#34;, AiIndexEndpointArgs.builder()        
 *             .description(&#34;A sample vertex endpoint with an public endpoint&#34;)
 *             .displayName(&#34;sample-endpoint&#34;)
 *             .labels(Map.of(&#34;label-one&#34;, &#34;value-one&#34;))
 *             .publicEndpointEnabled(true)
 *             .region(&#34;us-central1&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * IndexEndpoint can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:vertex/aiIndexEndpoint:AiIndexEndpoint default projects/{{project}}/locations/{{region}}/indexEndpoints/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:vertex/aiIndexEndpoint:AiIndexEndpoint default {{project}}/{{region}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:vertex/aiIndexEndpoint:AiIndexEndpoint default {{region}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:vertex/aiIndexEndpoint:AiIndexEndpoint default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:vertex/aiIndexEndpoint:AiIndexEndpoint")
public class AiIndexEndpoint extends com.pulumi.resources.CustomResource {
    /**
     * The timestamp of when the Index was created in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return The timestamp of when the Index was created in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * The description of the Index.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the Index.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The display name of the Index. The name can be up to 128 characters long and can consist of any UTF-8 characters.
     * 
     * ***
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output<String> displayName;

    /**
     * @return The display name of the Index. The name can be up to 128 characters long and can consist of any UTF-8 characters.
     * 
     * ***
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
     * Used to perform consistent read-modify-write updates.
     * 
     */
    @Export(name="etag", refs={String.class}, tree="[0]")
    private Output<String> etag;

    /**
     * @return Used to perform consistent read-modify-write updates.
     * 
     */
    public Output<String> etag() {
        return this.etag;
    }
    /**
     * The labels with user-defined metadata to organize your Indexes.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return The labels with user-defined metadata to organize your Indexes.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * The resource name of the Index.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The resource name of the Index.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The full name of the Google Compute Engine [network](https://cloud.google.com//compute/docs/networks-and-firewalls#networks) to which the index endpoint should be peered.
     * Private services access must already be configured for the network. If left unspecified, the index endpoint is not peered with any network.
     * [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/insert): `projects/{project}/global/networks/{network}`.
     * Where `{project}` is a project number, as in `12345`, and `{network}` is network name.
     * 
     */
    @Export(name="network", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> network;

    /**
     * @return The full name of the Google Compute Engine [network](https://cloud.google.com//compute/docs/networks-and-firewalls#networks) to which the index endpoint should be peered.
     * Private services access must already be configured for the network. If left unspecified, the index endpoint is not peered with any network.
     * [Format](https://cloud.google.com/compute/docs/reference/rest/v1/networks/insert): `projects/{project}/global/networks/{network}`.
     * Where `{project}` is a project number, as in `12345`, and `{network}` is network name.
     * 
     */
    public Output<Optional<String>> network() {
        return Codegen.optional(this.network);
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
     * If publicEndpointEnabled is true, this field will be populated with the domain name to use for this index endpoint.
     * 
     */
    @Export(name="publicEndpointDomainName", refs={String.class}, tree="[0]")
    private Output<String> publicEndpointDomainName;

    /**
     * @return If publicEndpointEnabled is true, this field will be populated with the domain name to use for this index endpoint.
     * 
     */
    public Output<String> publicEndpointDomainName() {
        return this.publicEndpointDomainName;
    }
    /**
     * If true, the deployed index will be accessible through public endpoint.
     * 
     */
    @Export(name="publicEndpointEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> publicEndpointEnabled;

    /**
     * @return If true, the deployed index will be accessible through public endpoint.
     * 
     */
    public Output<Optional<Boolean>> publicEndpointEnabled() {
        return Codegen.optional(this.publicEndpointEnabled);
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
     * The region of the index endpoint. eg us-central1
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> region;

    /**
     * @return The region of the index endpoint. eg us-central1
     * 
     */
    public Output<Optional<String>> region() {
        return Codegen.optional(this.region);
    }
    /**
     * The timestamp of when the Index was last updated in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
     * 
     */
    @Export(name="updateTime", refs={String.class}, tree="[0]")
    private Output<String> updateTime;

    /**
     * @return The timestamp of when the Index was last updated in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond resolution and up to nine fractional digits.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AiIndexEndpoint(String name) {
        this(name, AiIndexEndpointArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AiIndexEndpoint(String name, AiIndexEndpointArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AiIndexEndpoint(String name, AiIndexEndpointArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:vertex/aiIndexEndpoint:AiIndexEndpoint", name, args == null ? AiIndexEndpointArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AiIndexEndpoint(String name, Output<String> id, @Nullable AiIndexEndpointState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:vertex/aiIndexEndpoint:AiIndexEndpoint", name, state, makeResourceOptions(options, id));
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
    public static AiIndexEndpoint get(String name, Output<String> id, @Nullable AiIndexEndpointState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AiIndexEndpoint(name, id, state, options);
    }
}
