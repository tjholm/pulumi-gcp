// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

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
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fs from "fs";
 * import * as gcp from "@pulumi/gcp";
 *
 * const apiCfgApi = new gcp.apigateway.Api("apiCfgApi", {apiId: "my-api"}, {
 *     provider: google_beta,
 * });
 * const apiCfgApiConfig = new gcp.apigateway.ApiConfig("apiCfgApiConfig", {
 *     api: apiCfgApi.apiId,
 *     apiConfigId: "my-config",
 *     openapiDocuments: [{
 *         document: {
 *             path: "spec.yaml",
 *             contents: Buffer.from(fs.readFileSync("test-fixtures/apigateway/openapi.yaml"), 'binary').toString('base64'),
 *         },
 *     }],
 * }, {
 *     provider: google_beta,
 * });
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
 */
export class ApiConfig extends pulumi.CustomResource {
    /**
     * Get an existing ApiConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApiConfigState, opts?: pulumi.CustomResourceOptions): ApiConfig {
        return new ApiConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:apigateway/apiConfig:ApiConfig';

    /**
     * Returns true if the given object is an instance of ApiConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApiConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApiConfig.__pulumiType;
    }

    /**
     * The API to attach the config to.
     */
    public readonly api!: pulumi.Output<string>;
    /**
     * Identifier to assign to the API Config. Must be unique within scope of the parent resource(api).
     */
    public readonly apiConfigId!: pulumi.Output<string>;
    /**
     * Creates a unique name beginning with the
     * specified prefix. If this and apiConfigId are unspecified, a random value is chosen for the name.
     */
    public readonly apiConfigIdPrefix!: pulumi.Output<string>;
    /**
     * A user-visible name for the API.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Immutable. Gateway specific configuration.
     * If not specified, backend authentication will be set to use OIDC authentication using the default compute service account
     * Structure is documented below.
     */
    public readonly gatewayConfig!: pulumi.Output<outputs.apigateway.ApiConfigGatewayConfig | undefined>;
    /**
     * gRPC service definition files. If specified, openapiDocuments must not be included.
     * Structure is documented below.
     */
    public readonly grpcServices!: pulumi.Output<outputs.apigateway.ApiConfigGrpcService[] | undefined>;
    /**
     * Resource labels to represent user-provided metadata.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Optional. Service Configuration files. At least one must be included when using gRPC service definitions. See https://cloud.google.com/endpoints/docs/grpc/grpc-service-config#service_configuration_overview for the expected file contents.
     * If multiple files are specified, the files are merged with the following rules: * All singular scalar fields are merged using "last one wins" semantics in the order of the files uploaded. * Repeated fields are concatenated. * Singular embedded messages are merged using these rules for nested fields.
     * Structure is documented below.
     */
    public readonly managedServiceConfigs!: pulumi.Output<outputs.apigateway.ApiConfigManagedServiceConfig[] | undefined>;
    /**
     * The resource name of the API Config.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * OpenAPI specification documents. If specified, grpcServices and managedServiceConfigs must not be included.
     * Structure is documented below.
     */
    public readonly openapiDocuments!: pulumi.Output<outputs.apigateway.ApiConfigOpenapiDocument[] | undefined>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The ID of the associated Service Config (https://cloud.google.com/service-infrastructure/docs/glossary#config).
     */
    public /*out*/ readonly serviceConfigId!: pulumi.Output<string>;

    /**
     * Create a ApiConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApiConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApiConfigArgs | ApiConfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApiConfigState | undefined;
            resourceInputs["api"] = state ? state.api : undefined;
            resourceInputs["apiConfigId"] = state ? state.apiConfigId : undefined;
            resourceInputs["apiConfigIdPrefix"] = state ? state.apiConfigIdPrefix : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["gatewayConfig"] = state ? state.gatewayConfig : undefined;
            resourceInputs["grpcServices"] = state ? state.grpcServices : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["managedServiceConfigs"] = state ? state.managedServiceConfigs : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["openapiDocuments"] = state ? state.openapiDocuments : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["serviceConfigId"] = state ? state.serviceConfigId : undefined;
        } else {
            const args = argsOrState as ApiConfigArgs | undefined;
            if ((!args || args.api === undefined) && !opts.urn) {
                throw new Error("Missing required property 'api'");
            }
            resourceInputs["api"] = args ? args.api : undefined;
            resourceInputs["apiConfigId"] = args ? args.apiConfigId : undefined;
            resourceInputs["apiConfigIdPrefix"] = args ? args.apiConfigIdPrefix : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["gatewayConfig"] = args ? args.gatewayConfig : undefined;
            resourceInputs["grpcServices"] = args ? args.grpcServices : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["managedServiceConfigs"] = args ? args.managedServiceConfigs : undefined;
            resourceInputs["openapiDocuments"] = args ? args.openapiDocuments : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["serviceConfigId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ApiConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApiConfig resources.
 */
export interface ApiConfigState {
    /**
     * The API to attach the config to.
     */
    api?: pulumi.Input<string>;
    /**
     * Identifier to assign to the API Config. Must be unique within scope of the parent resource(api).
     */
    apiConfigId?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the
     * specified prefix. If this and apiConfigId are unspecified, a random value is chosen for the name.
     */
    apiConfigIdPrefix?: pulumi.Input<string>;
    /**
     * A user-visible name for the API.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Immutable. Gateway specific configuration.
     * If not specified, backend authentication will be set to use OIDC authentication using the default compute service account
     * Structure is documented below.
     */
    gatewayConfig?: pulumi.Input<inputs.apigateway.ApiConfigGatewayConfig>;
    /**
     * gRPC service definition files. If specified, openapiDocuments must not be included.
     * Structure is documented below.
     */
    grpcServices?: pulumi.Input<pulumi.Input<inputs.apigateway.ApiConfigGrpcService>[]>;
    /**
     * Resource labels to represent user-provided metadata.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Optional. Service Configuration files. At least one must be included when using gRPC service definitions. See https://cloud.google.com/endpoints/docs/grpc/grpc-service-config#service_configuration_overview for the expected file contents.
     * If multiple files are specified, the files are merged with the following rules: * All singular scalar fields are merged using "last one wins" semantics in the order of the files uploaded. * Repeated fields are concatenated. * Singular embedded messages are merged using these rules for nested fields.
     * Structure is documented below.
     */
    managedServiceConfigs?: pulumi.Input<pulumi.Input<inputs.apigateway.ApiConfigManagedServiceConfig>[]>;
    /**
     * The resource name of the API Config.
     */
    name?: pulumi.Input<string>;
    /**
     * OpenAPI specification documents. If specified, grpcServices and managedServiceConfigs must not be included.
     * Structure is documented below.
     */
    openapiDocuments?: pulumi.Input<pulumi.Input<inputs.apigateway.ApiConfigOpenapiDocument>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The ID of the associated Service Config (https://cloud.google.com/service-infrastructure/docs/glossary#config).
     */
    serviceConfigId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ApiConfig resource.
 */
export interface ApiConfigArgs {
    /**
     * The API to attach the config to.
     */
    api: pulumi.Input<string>;
    /**
     * Identifier to assign to the API Config. Must be unique within scope of the parent resource(api).
     */
    apiConfigId?: pulumi.Input<string>;
    /**
     * Creates a unique name beginning with the
     * specified prefix. If this and apiConfigId are unspecified, a random value is chosen for the name.
     */
    apiConfigIdPrefix?: pulumi.Input<string>;
    /**
     * A user-visible name for the API.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Immutable. Gateway specific configuration.
     * If not specified, backend authentication will be set to use OIDC authentication using the default compute service account
     * Structure is documented below.
     */
    gatewayConfig?: pulumi.Input<inputs.apigateway.ApiConfigGatewayConfig>;
    /**
     * gRPC service definition files. If specified, openapiDocuments must not be included.
     * Structure is documented below.
     */
    grpcServices?: pulumi.Input<pulumi.Input<inputs.apigateway.ApiConfigGrpcService>[]>;
    /**
     * Resource labels to represent user-provided metadata.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Optional. Service Configuration files. At least one must be included when using gRPC service definitions. See https://cloud.google.com/endpoints/docs/grpc/grpc-service-config#service_configuration_overview for the expected file contents.
     * If multiple files are specified, the files are merged with the following rules: * All singular scalar fields are merged using "last one wins" semantics in the order of the files uploaded. * Repeated fields are concatenated. * Singular embedded messages are merged using these rules for nested fields.
     * Structure is documented below.
     */
    managedServiceConfigs?: pulumi.Input<pulumi.Input<inputs.apigateway.ApiConfigManagedServiceConfig>[]>;
    /**
     * OpenAPI specification documents. If specified, grpcServices and managedServiceConfigs must not be included.
     * Structure is documented below.
     */
    openapiDocuments?: pulumi.Input<pulumi.Input<inputs.apigateway.ApiConfigOpenapiDocument>[]>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
}
