// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.apigateway;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.gcp.Utilities;
import com.pulumi.gcp.apigateway.ApiConfigArgs;
import com.pulumi.gcp.apigateway.inputs.ApiConfigState;
import com.pulumi.gcp.apigateway.outputs.ApiConfigGatewayConfig;
import com.pulumi.gcp.apigateway.outputs.ApiConfigGrpcService;
import com.pulumi.gcp.apigateway.outputs.ApiConfigManagedServiceConfig;
import com.pulumi.gcp.apigateway.outputs.ApiConfigOpenapiDocument;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * An API Configuration is an association of an API Controller Config and a Gateway Config
 * 
 * To get more information about ApiConfig, see:
 * 
 * * [API documentation](https://cloud.google.com/api-gateway/docs/reference/rest/v1beta/projects.locations.apis.configs)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/api-gateway/docs/creating-api-config)
 * 
 * ## Example Usage
 * ### Apigateway Api Config Basic
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.gcp.apigateway.Api;
 * import com.pulumi.gcp.apigateway.ApiArgs;
 * import com.pulumi.gcp.apigateway.ApiConfig;
 * import com.pulumi.gcp.apigateway.ApiConfigArgs;
 * import com.pulumi.gcp.apigateway.inputs.ApiConfigOpenapiDocumentArgs;
 * import com.pulumi.gcp.apigateway.inputs.ApiConfigOpenapiDocumentDocumentArgs;
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
 *         var apiCfgApi = new Api(&#34;apiCfgApi&#34;, ApiArgs.builder()        
 *             .apiId(&#34;my-api&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(google_beta)
 *                 .build());
 * 
 *         var apiCfgApiConfig = new ApiConfig(&#34;apiCfgApiConfig&#34;, ApiConfigArgs.builder()        
 *             .api(apiCfgApi.apiId())
 *             .apiConfigId(&#34;my-config&#34;)
 *             .openapiDocuments(ApiConfigOpenapiDocumentArgs.builder()
 *                 .document(ApiConfigOpenapiDocumentDocumentArgs.builder()
 *                     .path(&#34;spec.yaml&#34;)
 *                     .contents(Base64.getEncoder().encodeToString(Files.readAllBytes(Paths.get(&#34;test-fixtures/apigateway/openapi.yaml&#34;))))
 *                     .build())
 *                 .build())
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
 * ApiConfig can be imported using any of these accepted formats
 * 
 * ```sh
 *  $ pulumi import gcp:apigateway/apiConfig:ApiConfig default projects/{{project}}/locations/global/apis/{{api}}/configs/{{api_config_id}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:apigateway/apiConfig:ApiConfig default {{project}}/{{api}}/{{api_config_id}}
 * ```
 * 
 * ```sh
 *  $ pulumi import gcp:apigateway/apiConfig:ApiConfig default {{api}}/{{api_config_id}}
 * ```
 * 
 */
@ResourceType(type="gcp:apigateway/apiConfig:ApiConfig")
public class ApiConfig extends com.pulumi.resources.CustomResource {
    /**
     * The API to attach the config to.
     * 
     */
    @Export(name="api", type=String.class, parameters={})
    private Output<String> api;

    /**
     * @return The API to attach the config to.
     * 
     */
    public Output<String> api() {
        return this.api;
    }
    /**
     * Identifier to assign to the API Config. Must be unique within scope of the parent resource(api).
     * 
     */
    @Export(name="apiConfigId", type=String.class, parameters={})
    private Output<String> apiConfigId;

    /**
     * @return Identifier to assign to the API Config. Must be unique within scope of the parent resource(api).
     * 
     */
    public Output<String> apiConfigId() {
        return this.apiConfigId;
    }
    /**
     * Creates a unique name beginning with the
     * specified prefix. If this and api_config_id are unspecified, a random value is chosen for the name.
     * 
     */
    @Export(name="apiConfigIdPrefix", type=String.class, parameters={})
    private Output<String> apiConfigIdPrefix;

    /**
     * @return Creates a unique name beginning with the
     * specified prefix. If this and api_config_id are unspecified, a random value is chosen for the name.
     * 
     */
    public Output<String> apiConfigIdPrefix() {
        return this.apiConfigIdPrefix;
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
     * Immutable. Gateway specific configuration.
     * If not specified, backend authentication will be set to use OIDC authentication using the default compute service account
     * Structure is documented below.
     * 
     */
    @Export(name="gatewayConfig", type=ApiConfigGatewayConfig.class, parameters={})
    private Output</* @Nullable */ ApiConfigGatewayConfig> gatewayConfig;

    /**
     * @return Immutable. Gateway specific configuration.
     * If not specified, backend authentication will be set to use OIDC authentication using the default compute service account
     * Structure is documented below.
     * 
     */
    public Output<Optional<ApiConfigGatewayConfig>> gatewayConfig() {
        return Codegen.optional(this.gatewayConfig);
    }
    /**
     * gRPC service definition files. If specified, openapiDocuments must not be included.
     * Structure is documented below.
     * 
     */
    @Export(name="grpcServices", type=List.class, parameters={ApiConfigGrpcService.class})
    private Output</* @Nullable */ List<ApiConfigGrpcService>> grpcServices;

    /**
     * @return gRPC service definition files. If specified, openapiDocuments must not be included.
     * Structure is documented below.
     * 
     */
    public Output<Optional<List<ApiConfigGrpcService>>> grpcServices() {
        return Codegen.optional(this.grpcServices);
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
     * Optional. Service Configuration files. At least one must be included when using gRPC service definitions. See https://cloud.google.com/endpoints/docs/grpc/grpc-service-config#service_configuration_overview for the expected file contents.
     * If multiple files are specified, the files are merged with the following rules: * All singular scalar fields are merged using &#34;last one wins&#34; semantics in the order of the files uploaded. * Repeated fields are concatenated. * Singular embedded messages are merged using these rules for nested fields.
     * Structure is documented below.
     * 
     */
    @Export(name="managedServiceConfigs", type=List.class, parameters={ApiConfigManagedServiceConfig.class})
    private Output</* @Nullable */ List<ApiConfigManagedServiceConfig>> managedServiceConfigs;

    /**
     * @return Optional. Service Configuration files. At least one must be included when using gRPC service definitions. See https://cloud.google.com/endpoints/docs/grpc/grpc-service-config#service_configuration_overview for the expected file contents.
     * If multiple files are specified, the files are merged with the following rules: * All singular scalar fields are merged using &#34;last one wins&#34; semantics in the order of the files uploaded. * Repeated fields are concatenated. * Singular embedded messages are merged using these rules for nested fields.
     * Structure is documented below.
     * 
     */
    public Output<Optional<List<ApiConfigManagedServiceConfig>>> managedServiceConfigs() {
        return Codegen.optional(this.managedServiceConfigs);
    }
    /**
     * The resource name of the API Config.
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return The resource name of the API Config.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * OpenAPI specification documents. If specified, grpcServices and managedServiceConfigs must not be included.
     * Structure is documented below.
     * 
     */
    @Export(name="openapiDocuments", type=List.class, parameters={ApiConfigOpenapiDocument.class})
    private Output</* @Nullable */ List<ApiConfigOpenapiDocument>> openapiDocuments;

    /**
     * @return OpenAPI specification documents. If specified, grpcServices and managedServiceConfigs must not be included.
     * Structure is documented below.
     * 
     */
    public Output<Optional<List<ApiConfigOpenapiDocument>>> openapiDocuments() {
        return Codegen.optional(this.openapiDocuments);
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
     * The ID of the associated Service Config (https://cloud.google.com/service-infrastructure/docs/glossary#config).
     * 
     */
    @Export(name="serviceConfigId", type=String.class, parameters={})
    private Output<String> serviceConfigId;

    /**
     * @return The ID of the associated Service Config (https://cloud.google.com/service-infrastructure/docs/glossary#config).
     * 
     */
    public Output<String> serviceConfigId() {
        return this.serviceConfigId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ApiConfig(String name) {
        this(name, ApiConfigArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ApiConfig(String name, ApiConfigArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ApiConfig(String name, ApiConfigArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:apigateway/apiConfig:ApiConfig", name, args == null ? ApiConfigArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ApiConfig(String name, Output<String> id, @Nullable ApiConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("gcp:apigateway/apiConfig:ApiConfig", name, state, makeResourceOptions(options, id));
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
    public static ApiConfig get(String name, Output<String> id, @Nullable ApiConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ApiConfig(name, id, state, options);
    }
}
