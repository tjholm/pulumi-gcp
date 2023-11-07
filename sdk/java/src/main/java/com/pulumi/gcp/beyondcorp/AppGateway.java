// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.beyondcorp;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.beyondcorp.AppGatewayArgs;
import com.pulumi.gcp.beyondcorp.inputs.AppGatewayState;
import com.pulumi.gcp.beyondcorp.outputs.AppGatewayAllocatedConnection;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A BeyondCorp AppGateway resource represents a BeyondCorp protected AppGateway to a remote application. It creates
 * all the necessary GCP components needed for creating a BeyondCorp protected AppGateway. Multiple connectors can be
 * authorised for a single AppGateway.
 * 
 * To get more information about AppGateway, see:
 * 
 * * [API documentation](https://cloud.google.com/beyondcorp/docs/reference/rest#rest-resource:-v1.projects.locations.appgateways)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/beyondcorp-enterprise/docs/enable-app-connector)
 * 
 * ## Example Usage
 * ### Beyondcorp App Gateway Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.beyondcorp.AppGateway;
 * import com.pulumi.gcp.beyondcorp.AppGatewayArgs;
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
 *         var appGateway = new AppGateway(&#34;appGateway&#34;, AppGatewayArgs.builder()        
 *             .hostType(&#34;GCP_REGIONAL_MIG&#34;)
 *             .region(&#34;us-central1&#34;)
 *             .type(&#34;TCP_PROXY&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Beyondcorp App Gateway Full
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.beyondcorp.AppGateway;
 * import com.pulumi.gcp.beyondcorp.AppGatewayArgs;
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
 *         var appGateway = new AppGateway(&#34;appGateway&#34;, AppGatewayArgs.builder()        
 *             .displayName(&#34;some display name&#34;)
 *             .hostType(&#34;GCP_REGIONAL_MIG&#34;)
 *             .labels(Map.ofEntries(
 *                 Map.entry(&#34;bar&#34;, &#34;baz&#34;),
 *                 Map.entry(&#34;foo&#34;, &#34;bar&#34;)
 *             ))
 *             .region(&#34;us-central1&#34;)
 *             .type(&#34;TCP_PROXY&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * AppGateway can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:beyondcorp/appGateway:AppGateway default projects/{{project}}/locations/{{region}}/appGateways/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:beyondcorp/appGateway:AppGateway default {{project}}/{{region}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:beyondcorp/appGateway:AppGateway default {{region}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:beyondcorp/appGateway:AppGateway default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:beyondcorp/appGateway:AppGateway")
public class AppGateway extends com.pulumi.resources.CustomResource {
    /**
     * A list of connections allocated for the Gateway.
     * Structure is documented below.
     * 
     */
    @Export(name="allocatedConnections", refs={List.class,AppGatewayAllocatedConnection.class}, tree="[0,1]")
    private Output<List<AppGatewayAllocatedConnection>> allocatedConnections;

    /**
     * @return A list of connections allocated for the Gateway.
     * Structure is documented below.
     * 
     */
    public Output<List<AppGatewayAllocatedConnection>> allocatedConnections() {
        return this.allocatedConnections;
    }
    /**
     * An arbitrary user-provided name for the AppGateway.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> displayName;

    /**
     * @return An arbitrary user-provided name for the AppGateway.
     * 
     */
    public Output<Optional<String>> displayName() {
        return Codegen.optional(this.displayName);
    }
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
     * clients and services.
     * 
     */
    @Export(name="effectiveLabels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output<Map<String,String>> effectiveLabels;

    /**
     * @return All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
     * clients and services.
     * 
     */
    public Output<Map<String,String>> effectiveLabels() {
        return this.effectiveLabels;
    }
    /**
     * The type of hosting used by the AppGateway.
     * Default value is `HOST_TYPE_UNSPECIFIED`.
     * Possible values are: `HOST_TYPE_UNSPECIFIED`, `GCP_REGIONAL_MIG`.
     * 
     */
    @Export(name="hostType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> hostType;

    /**
     * @return The type of hosting used by the AppGateway.
     * Default value is `HOST_TYPE_UNSPECIFIED`.
     * Possible values are: `HOST_TYPE_UNSPECIFIED`, `GCP_REGIONAL_MIG`.
     * 
     */
    public Output<Optional<String>> hostType() {
        return Codegen.optional(this.hostType);
    }
    /**
     * Resource labels to represent user provided metadata.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    @Export(name="labels", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return Resource labels to represent user provided metadata.
     * 
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effective_labels` for all of the labels present on the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * ID of the AppGateway.
     * 
     * ***
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return ID of the AppGateway.
     * 
     * ***
     * 
     */
    public Output<String> name() {
        return this.name;
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
     * The region of the AppGateway.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> region;

    /**
     * @return The region of the AppGateway.
     * 
     */
    public Output<Optional<String>> region() {
        return Codegen.optional(this.region);
    }
    /**
     * Represents the different states of a AppGateway.
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return Represents the different states of a AppGateway.
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * The type of network connectivity used by the AppGateway.
     * Default value is `TYPE_UNSPECIFIED`.
     * Possible values are: `TYPE_UNSPECIFIED`, `TCP_PROXY`.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> type;

    /**
     * @return The type of network connectivity used by the AppGateway.
     * Default value is `TYPE_UNSPECIFIED`.
     * Possible values are: `TYPE_UNSPECIFIED`, `TCP_PROXY`.
     * 
     */
    public Output<Optional<String>> type() {
        return Codegen.optional(this.type);
    }
    /**
     * Server-defined URI for this resource.
     * 
     */
    @Export(name="uri", refs={String.class}, tree="[0]")
    private Output<String> uri;

    /**
     * @return Server-defined URI for this resource.
     * 
     */
    public Output<String> uri() {
        return this.uri;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AppGateway(String name) {
        this(name, AppGatewayArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AppGateway(String name, @Nullable AppGatewayArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AppGateway(String name, @Nullable AppGatewayArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:beyondcorp/appGateway:AppGateway", name, args == null ? AppGatewayArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AppGateway(String name, Output<String> id, @Nullable AppGatewayState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:beyondcorp/appGateway:AppGateway", name, state, makeResourceOptions(options, id));
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
    public static AppGateway get(String name, Output<String> id, @Nullable AppGatewayState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AppGateway(name, id, state, options);
    }
}
