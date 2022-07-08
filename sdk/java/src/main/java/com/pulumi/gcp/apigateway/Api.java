// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.apigateway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.apigateway.ApiArgs;
import com.pulumi.gcp.apigateway.inputs.ApiState;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * A consumable API that can be used by multiple Gateways.
 * 
 * To get more information about Api, see:
 * 
 * * [API documentation](https://cloud.google.com/api-gateway/docs/reference/rest/v1beta/projects.locations.apis)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/api-gateway/docs/quickstart)
 * 
 * ## Example Usage
 * ### Apigateway Api Basic
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
 *         var api = new Api(&#34;api&#34;, ApiArgs.builder()        
 *             .apiId(&#34;api&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Api can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:apigateway/api:Api default projects/{{project}}/locations/global/apis/{{api_id}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:apigateway/api:Api default {{project}}/{{api_id}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:apigateway/api:Api default {{api_id}}
 * ```
 * 
 */
@ResourceType(type="gcp:apigateway/api:Api")
public class Api extends com.pulumi.resources.CustomResource {
    /**
     * Identifier to assign to the API. Must be unique within scope of the parent resource(project)
     * 
     */
    @Export(name="apiId", type=String.class, parameters={})
    private Output<String> apiId;

    /**
     * @return Identifier to assign to the API. Must be unique within scope of the parent resource(project)
     * 
     */
    public Output<String> apiId() {
        return this.apiId;
    }
    /**
     * Creation timestamp in RFC3339 text format.
     * 
     */
    @Export(name="createTime", type=String.class, parameters={})
    private Output<String> createTime;

    /**
     * @return Creation timestamp in RFC3339 text format.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * A user-visible name for the API.
     * 
     */
    @Export(name="displayName", type=String.class, parameters={})
    private Output<String> displayName;

    /**
     * @return A user-visible name for the API.
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }
    /**
     * Resource labels to represent user-provided metadata.
     * 
     */
    @Export(name="labels", type=Map.class, parameters={String.class, String.class})
    private Output</* @Nullable */ Map<String,String>> labels;

    /**
     * @return Resource labels to represent user-provided metadata.
     * 
     */
    public Output<Optional<Map<String,String>>> labels() {
        return Codegen.optional(this.labels);
    }
    /**
     * Immutable. The name of a Google Managed Service ( https://cloud.google.com/service-infrastructure/docs/glossary#managed).
     * If not specified, a new Service will automatically be created in the same project as this API.
     * 
     */
    @Export(name="managedService", type=String.class, parameters={})
    private Output<String> managedService;

    /**
     * @return Immutable. The name of a Google Managed Service ( https://cloud.google.com/service-infrastructure/docs/glossary#managed).
     * If not specified, a new Service will automatically be created in the same project as this API.
     * 
     */
    public Output<String> managedService() {
        return this.managedService;
    }
    /**
     * The resource name of the API. Format &#39;projects/{{project}}/locations/global/apis/{{apiId}}&#39;
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The resource name of the API. Format &#39;projects/{{project}}/locations/global/apis/{{apiId}}&#39;
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
    public Api(String name) {
        this(name, ApiArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Api(String name, ApiArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Api(String name, ApiArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:apigateway/api:Api", name, args == null ? ApiArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Api(String name, Output<String> id, @Nullable ApiState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:apigateway/api:Api", name, state, makeResourceOptions(options, id));
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
    public static Api get(String name, Output<String> id, @Nullable ApiState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Api(name, id, state, options);
    }
}
