// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.diagflow;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.diagflow.FulfillmentArgs;
import com.pulumi.gcp.diagflow.inputs.FulfillmentState;
import com.pulumi.gcp.diagflow.outputs.FulfillmentFeature;
import com.pulumi.gcp.diagflow.outputs.FulfillmentGenericWebService;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * By default, your agent responds to a matched intent with a static response. If you&#39;re using one of the integration options, you can provide a more dynamic response by using fulfillment. When you enable fulfillment for an intent, Dialogflow responds to that intent by calling a service that you define. For example, if an end-user wants to schedule a haircut on Friday, your service can check your database and respond to the end-user with availability information for Friday.
 * 
 * To get more information about Fulfillment, see:
 * 
 * * [API documentation](https://cloud.google.com/dialogflow/es/docs/reference/rest/v2/projects.agent/getFulfillment)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/dialogflow/es/docs/fulfillment-overview)
 * 
 * ## Example Usage
 * ### Dialogflow Fulfillment Basic
 * ```java
 * package generated_program;
 * 
 * import java.util.*;
 * import java.io.*;
 * import java.nio.*;
 * import com.pulumi.*;
 * import com.pulumi.resources.CustomResourceOptions;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var basicAgent = new Agent(&#34;basicAgent&#34;, AgentArgs.builder()        
 *             .displayName(&#34;example_agent&#34;)
 *             .defaultLanguageCode(&#34;en&#34;)
 *             .timeZone(&#34;America/New_York&#34;)
 *             .build());
 * 
 *         var basicFulfillment = new Fulfillment(&#34;basicFulfillment&#34;, FulfillmentArgs.builder()        
 *             .displayName(&#34;basic-fulfillment&#34;)
 *             .enabled(true)
 *             .genericWebService(FulfillmentGenericWebServiceArgs.builder()
 *                 .uri(&#34;https://google.com&#34;)
 *                 .username(&#34;admin&#34;)
 *                 .password(&#34;password&#34;)
 *                 .requestHeaders(Map.of(&#34;name&#34;, &#34;wrench&#34;))
 *                 .build())
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(basicAgent)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Fulfillment can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:diagflow/fulfillment:Fulfillment default {{name}}
 * ```
 * 
 */
@ResourceType(type="gcp:diagflow/fulfillment:Fulfillment")
public class Fulfillment extends com.pulumi.resources.CustomResource {
    /**
     * The human-readable name of the fulfillment, unique within the agent.
     * 
     */
    @Export(name="displayName", type=String.class, parameters={})
    private Output<String> displayName;

    /**
     * @return The human-readable name of the fulfillment, unique within the agent.
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }
    /**
     * Whether fulfillment is enabled.
     * 
     */
    @Export(name="enabled", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> enabled;

    /**
     * @return Whether fulfillment is enabled.
     * 
     */
    public Output<Optional<Boolean>> enabled() {
        return Codegen.optional(this.enabled);
    }
    /**
     * The field defines whether the fulfillment is enabled for certain features.
     * Structure is documented below.
     * 
     */
    @Export(name="features", type=List.class, parameters={FulfillmentFeature.class})
    private Output</* @Nullable */ List<FulfillmentFeature>> features;

    /**
     * @return The field defines whether the fulfillment is enabled for certain features.
     * Structure is documented below.
     * 
     */
    public Output<Optional<List<FulfillmentFeature>>> features() {
        return Codegen.optional(this.features);
    }
    /**
     * Represents configuration for a generic web service. Dialogflow supports two mechanisms for authentications: - Basic authentication with username and password. - Authentication with additional authentication headers.
     * Structure is documented below.
     * 
     */
    @Export(name="genericWebService", type=FulfillmentGenericWebService.class, parameters={})
    private Output</* @Nullable */ FulfillmentGenericWebService> genericWebService;

    /**
     * @return Represents configuration for a generic web service. Dialogflow supports two mechanisms for authentications: - Basic authentication with username and password. - Authentication with additional authentication headers.
     * Structure is documented below.
     * 
     */
    public Output<Optional<FulfillmentGenericWebService>> genericWebService() {
        return Codegen.optional(this.genericWebService);
    }
    /**
     * The unique identifier of the fulfillment. Format: projects/&lt;Project ID&gt;/agent/fulfillment - projects/&lt;Project
     * ID&gt;/locations/&lt;Location ID&gt;/agent/fulfillment
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The unique identifier of the fulfillment. Format: projects/&lt;Project ID&gt;/agent/fulfillment - projects/&lt;Project
     * ID&gt;/locations/&lt;Location ID&gt;/agent/fulfillment
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
    @Export(name="project", type=String.class, parameters={})
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
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Fulfillment(String name) {
        this(name, FulfillmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Fulfillment(String name, FulfillmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Fulfillment(String name, FulfillmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:diagflow/fulfillment:Fulfillment", name, args == null ? FulfillmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Fulfillment(String name, Output<String> id, @Nullable FulfillmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:diagflow/fulfillment:Fulfillment", name, state, makeResourceOptions(options, id));
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
    public static Fulfillment get(String name, Output<String> id, @Nullable FulfillmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Fulfillment(name, id, state, options);
    }
}
