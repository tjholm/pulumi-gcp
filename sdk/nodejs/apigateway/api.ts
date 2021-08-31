// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

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
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const api = new gcp.apigateway.Api("api", {apiId: "api"}, {
 *     provider: google_beta,
 * });
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
 */
export class Api extends pulumi.CustomResource {
    /**
     * Get an existing Api resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApiState, opts?: pulumi.CustomResourceOptions): Api {
        return new Api(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:apigateway/api:Api';

    /**
     * Returns true if the given object is an instance of Api.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Api {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Api.__pulumiType;
    }

    /**
     * Identifier to assign to the API. Must be unique within scope of the parent resource(project)
     */
    public readonly apiId!: pulumi.Output<string>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * A user-visible name for the API.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Resource labels to represent user-provided metadata.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Immutable. The name of a Google Managed Service ( https://cloud.google.com/service-infrastructure/docs/glossary#managed).
     * If not specified, a new Service will automatically be created in the same project as this API.
     */
    public readonly managedService!: pulumi.Output<string>;
    /**
     * The resource name of the API. Format 'projects/{{project}}/locations/global/apis/{{apiId}}'
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a Api resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApiArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApiArgs | ApiState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApiState | undefined;
            inputs["apiId"] = state ? state.apiId : undefined;
            inputs["createTime"] = state ? state.createTime : undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["managedService"] = state ? state.managedService : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as ApiArgs | undefined;
            if ((!args || args.apiId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'apiId'");
            }
            inputs["apiId"] = args ? args.apiId : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["managedService"] = args ? args.managedService : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["createTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Api.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Api resources.
 */
export interface ApiState {
    /**
     * Identifier to assign to the API. Must be unique within scope of the parent resource(project)
     */
    apiId?: pulumi.Input<string>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    createTime?: pulumi.Input<string>;
    /**
     * A user-visible name for the API.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Resource labels to represent user-provided metadata.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Immutable. The name of a Google Managed Service ( https://cloud.google.com/service-infrastructure/docs/glossary#managed).
     * If not specified, a new Service will automatically be created in the same project as this API.
     */
    managedService?: pulumi.Input<string>;
    /**
     * The resource name of the API. Format 'projects/{{project}}/locations/global/apis/{{apiId}}'
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Api resource.
 */
export interface ApiArgs {
    /**
     * Identifier to assign to the API. Must be unique within scope of the parent resource(project)
     */
    apiId: pulumi.Input<string>;
    /**
     * A user-visible name for the API.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Resource labels to represent user-provided metadata.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Immutable. The name of a Google Managed Service ( https://cloud.google.com/service-infrastructure/docs/glossary#managed).
     * If not specified, a new Service will automatically be created in the same project as this API.
     */
    managedService?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
}
