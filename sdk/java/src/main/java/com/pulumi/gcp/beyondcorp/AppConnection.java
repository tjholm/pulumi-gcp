// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.beyondcorp;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.beyondcorp.AppConnectionArgs;
import com.pulumi.gcp.beyondcorp.inputs.AppConnectionState;
import com.pulumi.gcp.beyondcorp.outputs.AppConnectionApplicationEndpoint;
import com.pulumi.gcp.beyondcorp.outputs.AppConnectionGateway;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A BeyondCorp AppConnection resource represents a BeyondCorp protected AppConnection to a remote application.
 * It creates all the necessary GCP components needed for creating a BeyondCorp protected AppConnection.
 * Multiple connectors can be authorised for a single AppConnection.
 * 
 * To get more information about AppConnection, see:
 * 
 * * [API documentation](https://cloud.google.com/beyondcorp/docs/reference/rest#rest-resource:-v1.projects.locations.appconnections)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/beyondcorp-enterprise/docs/enable-app-connector)
 * 
 * ## Example Usage
 * ### Beyondcorp App Connection Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.serviceaccount.Account;
 * import com.pulumi.gcp.serviceaccount.AccountArgs;
 * import com.pulumi.gcp.beyondcorp.AppConnector;
 * import com.pulumi.gcp.beyondcorp.AppConnectorArgs;
 * import com.pulumi.gcp.beyondcorp.inputs.AppConnectorPrincipalInfoArgs;
 * import com.pulumi.gcp.beyondcorp.inputs.AppConnectorPrincipalInfoServiceAccountArgs;
 * import com.pulumi.gcp.beyondcorp.AppConnection;
 * import com.pulumi.gcp.beyondcorp.AppConnectionArgs;
 * import com.pulumi.gcp.beyondcorp.inputs.AppConnectionApplicationEndpointArgs;
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
 *         var serviceAccount = new Account(&#34;serviceAccount&#34;, AccountArgs.builder()        
 *             .accountId(&#34;my-account&#34;)
 *             .displayName(&#34;Test Service Account&#34;)
 *             .build());
 * 
 *         var appConnector = new AppConnector(&#34;appConnector&#34;, AppConnectorArgs.builder()        
 *             .principalInfo(AppConnectorPrincipalInfoArgs.builder()
 *                 .serviceAccount(AppConnectorPrincipalInfoServiceAccountArgs.builder()
 *                     .email(serviceAccount.email())
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *         var appConnection = new AppConnection(&#34;appConnection&#34;, AppConnectionArgs.builder()        
 *             .type(&#34;TCP_PROXY&#34;)
 *             .applicationEndpoint(AppConnectionApplicationEndpointArgs.builder()
 *                 .host(&#34;foo-host&#34;)
 *                 .port(8080)
 *                 .build())
 *             .connectors(appConnector.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * ### Beyondcorp App Connection Full
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.serviceaccount.Account;
 * import com.pulumi.gcp.serviceaccount.AccountArgs;
 * import com.pulumi.gcp.beyondcorp.AppGateway;
 * import com.pulumi.gcp.beyondcorp.AppGatewayArgs;
 * import com.pulumi.gcp.beyondcorp.AppConnector;
 * import com.pulumi.gcp.beyondcorp.AppConnectorArgs;
 * import com.pulumi.gcp.beyondcorp.inputs.AppConnectorPrincipalInfoArgs;
 * import com.pulumi.gcp.beyondcorp.inputs.AppConnectorPrincipalInfoServiceAccountArgs;
 * import com.pulumi.gcp.beyondcorp.AppConnection;
 * import com.pulumi.gcp.beyondcorp.AppConnectionArgs;
 * import com.pulumi.gcp.beyondcorp.inputs.AppConnectionApplicationEndpointArgs;
 * import com.pulumi.gcp.beyondcorp.inputs.AppConnectionGatewayArgs;
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
 *         var serviceAccount = new Account(&#34;serviceAccount&#34;, AccountArgs.builder()        
 *             .accountId(&#34;my-account&#34;)
 *             .displayName(&#34;Test Service Account&#34;)
 *             .build());
 * 
 *         var appGateway = new AppGateway(&#34;appGateway&#34;, AppGatewayArgs.builder()        
 *             .type(&#34;TCP_PROXY&#34;)
 *             .hostType(&#34;GCP_REGIONAL_MIG&#34;)
 *             .build());
 * 
 *         var appConnector = new AppConnector(&#34;appConnector&#34;, AppConnectorArgs.builder()        
 *             .principalInfo(AppConnectorPrincipalInfoArgs.builder()
 *                 .serviceAccount(AppConnectorPrincipalInfoServiceAccountArgs.builder()
 *                     .email(serviceAccount.email())
 *                     .build())
 *                 .build())
 *             .build());
 * 
 *         var appConnection = new AppConnection(&#34;appConnection&#34;, AppConnectionArgs.builder()        
 *             .type(&#34;TCP_PROXY&#34;)
 *             .displayName(&#34;some display name&#34;)
 *             .applicationEndpoint(AppConnectionApplicationEndpointArgs.builder()
 *                 .host(&#34;foo-host&#34;)
 *                 .port(8080)
 *                 .build())
 *             .connectors(appConnector.id())
 *             .gateway(AppConnectionGatewayArgs.builder()
 *                 .appGateway(appGateway.id())
 *                 .build())
 *             .labels(Map.ofEntries(
 *                 Map.entry(&#34;foo&#34;, &#34;bar&#34;),
 *                 Map.entry(&#34;bar&#34;, &#34;baz&#34;)
 *             ))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * AppConnection can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:beyondcorp/appConnection:AppConnection default projects/{{project}}/locations/{{region}}/appConnections/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:beyondcorp/appConnection:AppConnection default {{project}}/{{region}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:beyondcorp/appConnection:AppConnection default {{region}}/{{name}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:beyondcorp/appConnection:AppConnection default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:beyondcorp/appConnection:AppConnection")
public class AppConnection extends com.pulumi.resources.CustomResource {
    /**
     * Address of the remote application endpoint for the BeyondCorp AppConnection.
     * Structure is documented below.
     * 
     */
    @Export(name="applicationEndpoint", refs={AppConnectionApplicationEndpoint.class}, tree="[0]")
    private Output<AppConnectionApplicationEndpoint> applicationEndpoint;

    /**
     * @return Address of the remote application endpoint for the BeyondCorp AppConnection.
     * Structure is documented below.
     * 
     */
    public Output<AppConnectionApplicationEndpoint> applicationEndpoint() {
        return this.applicationEndpoint;
    }
    /**
     * List of AppConnectors that are authorised to be associated with this AppConnection
     * 
     */
    @Export(name="connectors", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> connectors;

    /**
     * @return List of AppConnectors that are authorised to be associated with this AppConnection
     * 
     */
    public Output<Optional<List<String>>> connectors() {
        return Codegen.optional(this.connectors);
    }
    /**
     * An arbitrary user-provided name for the AppConnection.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> displayName;

    /**
     * @return An arbitrary user-provided name for the AppConnection.
     * 
     */
    public Output<Optional<String>> displayName() {
        return Codegen.optional(this.displayName);
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
     * Gateway used by the AppConnection.
     * Structure is documented below.
     * 
     */
    @Export(name="gateway", refs={AppConnectionGateway.class}, tree="[0]")
    private Output<AppConnectionGateway> gateway;

    /**
     * @return Gateway used by the AppConnection.
     * Structure is documented below.
     * 
     */
    public Output<AppConnectionGateway> gateway() {
        return this.gateway;
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
     * ID of the AppConnection.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return ID of the AppConnection.
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
     * The region of the AppConnection.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> region;

    /**
     * @return The region of the AppConnection.
     * 
     */
    public Output<Optional<String>> region() {
        return Codegen.optional(this.region);
    }
    /**
     * The type of network connectivity used by the AppConnection. Refer to
     * https://cloud.google.com/beyondcorp/docs/reference/rest/v1/projects.locations.appConnections#type
     * for a list of possible values.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> type;

    /**
     * @return The type of network connectivity used by the AppConnection. Refer to
     * https://cloud.google.com/beyondcorp/docs/reference/rest/v1/projects.locations.appConnections#type
     * for a list of possible values.
     * 
     */
    public Output<Optional<String>> type() {
        return Codegen.optional(this.type);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AppConnection(String name) {
        this(name, AppConnectionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AppConnection(String name, AppConnectionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AppConnection(String name, AppConnectionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:beyondcorp/appConnection:AppConnection", name, args == null ? AppConnectionArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AppConnection(String name, Output<String> id, @Nullable AppConnectionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:beyondcorp/appConnection:AppConnection", name, state, makeResourceOptions(options, id));
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
    public static AppConnection get(String name, Output<String> id, @Nullable AppConnectionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AppConnection(name, id, state, options);
    }
}
