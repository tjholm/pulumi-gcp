// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This resource creates and rolls out a Cloud Endpoints service using OpenAPI or gRPC.  View the relevant docs for [OpenAPI](https://cloud.google.com/endpoints/docs/openapi/) and [gRPC](https://cloud.google.com/endpoints/docs/grpc/).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fs from "fs";
 * import * as gcp from "@pulumi/gcp";
 *
 * const openapiService = new gcp.endpoints.Service("openapiService", {
 *     serviceName: "api-name.endpoints.project-id.cloud.goog",
 *     project: "project-id",
 *     openapiConfig: fs.readFileSync("openapi_spec.yml"),
 * });
 * const grpcService = new gcp.endpoints.Service("grpcService", {
 *     serviceName: "api-name.endpoints.project-id.cloud.goog",
 *     project: "project-id",
 *     grpcConfig: fs.readFileSync("service_spec.yml"),
 *     protocOutputBase64: Buffer.from(fs.readFileSync("compiled_descriptor_file.pb"), 'binary').toString('base64'),
 * });
 * ```
 *
 * The example in `examples/endpoints_on_compute_engine` shows the API from the quickstart running on a Compute Engine VM and reachable through Cloud Endpoints, which may also be useful.
 *
 * ## Import
 *
 * This resource does not support import.
 */
export class Service extends pulumi.CustomResource {
    /**
     * Get an existing Service resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceState, opts?: pulumi.CustomResourceOptions): Service {
        return new Service(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:endpoints/service:Service';

    /**
     * Returns true if the given object is an instance of Service.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Service {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Service.__pulumiType;
    }

    /**
     * A list of API objects.
     */
    public /*out*/ readonly apis!: pulumi.Output<outputs.endpoints.ServiceApi[]>;
    /**
     * The autogenerated ID for the configuration that is rolled out as part of the creation of this resource. Must be provided
     * to compute engine instances as a tag.
     */
    public /*out*/ readonly configId!: pulumi.Output<string>;
    /**
     * The address at which the service can be found - usually the same as the service name.
     */
    public /*out*/ readonly dnsAddress!: pulumi.Output<string>;
    /**
     * A list of Endpoint objects.
     */
    public /*out*/ readonly endpoints!: pulumi.Output<outputs.endpoints.ServiceEndpoint[]>;
    /**
     * The full text of the Service Config YAML file (Example located here). If provided, must also provide
     * protoc_output_base64. open_api config must not be provided.
     */
    public readonly grpcConfig!: pulumi.Output<string | undefined>;
    /**
     * The full text of the OpenAPI YAML configuration as described here. Either this, or both of grpc_config and
     * protoc_output_base64 must be specified.
     */
    public readonly openapiConfig!: pulumi.Output<string | undefined>;
    /**
     * The project ID that the service belongs to. If not provided, provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The full contents of the Service Descriptor File generated by protoc. This should be a compiled .pb file,
     * base64-encoded.
     */
    public readonly protocOutputBase64!: pulumi.Output<string | undefined>;
    /**
     * The name of the service. Usually of the form $apiname.endpoints.$projectid.cloud.goog.
     */
    public readonly serviceName!: pulumi.Output<string>;

    /**
     * Create a Service resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceArgs | ServiceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceState | undefined;
            resourceInputs["apis"] = state ? state.apis : undefined;
            resourceInputs["configId"] = state ? state.configId : undefined;
            resourceInputs["dnsAddress"] = state ? state.dnsAddress : undefined;
            resourceInputs["endpoints"] = state ? state.endpoints : undefined;
            resourceInputs["grpcConfig"] = state ? state.grpcConfig : undefined;
            resourceInputs["openapiConfig"] = state ? state.openapiConfig : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["protocOutputBase64"] = state ? state.protocOutputBase64 : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
        } else {
            const args = argsOrState as ServiceArgs | undefined;
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            resourceInputs["grpcConfig"] = args ? args.grpcConfig : undefined;
            resourceInputs["openapiConfig"] = args ? args.openapiConfig : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["protocOutputBase64"] = args ? args.protocOutputBase64 : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["apis"] = undefined /*out*/;
            resourceInputs["configId"] = undefined /*out*/;
            resourceInputs["dnsAddress"] = undefined /*out*/;
            resourceInputs["endpoints"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Service.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Service resources.
 */
export interface ServiceState {
    /**
     * A list of API objects.
     */
    apis?: pulumi.Input<pulumi.Input<inputs.endpoints.ServiceApi>[]>;
    /**
     * The autogenerated ID for the configuration that is rolled out as part of the creation of this resource. Must be provided
     * to compute engine instances as a tag.
     */
    configId?: pulumi.Input<string>;
    /**
     * The address at which the service can be found - usually the same as the service name.
     */
    dnsAddress?: pulumi.Input<string>;
    /**
     * A list of Endpoint objects.
     */
    endpoints?: pulumi.Input<pulumi.Input<inputs.endpoints.ServiceEndpoint>[]>;
    /**
     * The full text of the Service Config YAML file (Example located here). If provided, must also provide
     * protoc_output_base64. open_api config must not be provided.
     */
    grpcConfig?: pulumi.Input<string>;
    /**
     * The full text of the OpenAPI YAML configuration as described here. Either this, or both of grpc_config and
     * protoc_output_base64 must be specified.
     */
    openapiConfig?: pulumi.Input<string>;
    /**
     * The project ID that the service belongs to. If not provided, provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The full contents of the Service Descriptor File generated by protoc. This should be a compiled .pb file,
     * base64-encoded.
     */
    protocOutputBase64?: pulumi.Input<string>;
    /**
     * The name of the service. Usually of the form $apiname.endpoints.$projectid.cloud.goog.
     */
    serviceName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Service resource.
 */
export interface ServiceArgs {
    /**
     * The full text of the Service Config YAML file (Example located here). If provided, must also provide
     * protoc_output_base64. open_api config must not be provided.
     */
    grpcConfig?: pulumi.Input<string>;
    /**
     * The full text of the OpenAPI YAML configuration as described here. Either this, or both of grpc_config and
     * protoc_output_base64 must be specified.
     */
    openapiConfig?: pulumi.Input<string>;
    /**
     * The project ID that the service belongs to. If not provided, provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The full contents of the Service Descriptor File generated by protoc. This should be a compiled .pb file,
     * base64-encoded.
     */
    protocOutputBase64?: pulumi.Input<string>;
    /**
     * The name of the service. Usually of the form $apiname.endpoints.$projectid.cloud.goog.
     */
    serviceName: pulumi.Input<string>;
}
