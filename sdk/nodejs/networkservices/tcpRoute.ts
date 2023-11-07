// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ## Import
 *
 * TcpRoute can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:networkservices/tcpRoute:TcpRoute default projects/{{project}}/locations/global/tcpRoutes/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:networkservices/tcpRoute:TcpRoute default {{project}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:networkservices/tcpRoute:TcpRoute default {{name}}
 * ```
 */
export class TcpRoute extends pulumi.CustomResource {
    /**
     * Get an existing TcpRoute resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TcpRouteState, opts?: pulumi.CustomResourceOptions): TcpRoute {
        return new TcpRoute(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:networkservices/tcpRoute:TcpRoute';

    /**
     * Returns true if the given object is an instance of TcpRoute.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TcpRoute {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TcpRoute.__pulumiType;
    }

    /**
     * Time the TcpRoute was created in UTC.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * A free-text description of the resource. Max length 1024 characters.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
     * clients and services.
     */
    public /*out*/ readonly effectiveLabels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Gateways defines a list of gateways this TcpRoute is attached to, as one of the routing rules to route the requests served by the gateway.
     * Each gateway reference should match the pattern: projects/*&#47;locations/global/gateways/<gateway_name>
     */
    public readonly gateways!: pulumi.Output<string[] | undefined>;
    /**
     * Set of label tags associated with the TcpRoute resource.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Meshes defines a list of meshes this TcpRoute is attached to, as one of the routing rules to route the requests served by the mesh.
     * Each mesh reference should match the pattern: projects/*&#47;locations/global/meshes/<mesh_name>
     * The attached Mesh should be of a type SIDECAR
     */
    public readonly meshes!: pulumi.Output<string[] | undefined>;
    /**
     * Name of the TcpRoute resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     */
    public /*out*/ readonly pulumiLabels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Rules that define how traffic is routed and handled. At least one RouteRule must be supplied.
     * If there are multiple rules then the action taken will be the first rule to match.
     * Structure is documented below.
     */
    public readonly rules!: pulumi.Output<outputs.networkservices.TcpRouteRule[]>;
    /**
     * Server-defined URL of this resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * Time the TcpRoute was updated in UTC.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a TcpRoute resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TcpRouteArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TcpRouteArgs | TcpRouteState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TcpRouteState | undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["effectiveLabels"] = state ? state.effectiveLabels : undefined;
            resourceInputs["gateways"] = state ? state.gateways : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["meshes"] = state ? state.meshes : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["pulumiLabels"] = state ? state.pulumiLabels : undefined;
            resourceInputs["rules"] = state ? state.rules : undefined;
            resourceInputs["selfLink"] = state ? state.selfLink : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as TcpRouteArgs | undefined;
            if ((!args || args.rules === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rules'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["gateways"] = args ? args.gateways : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["meshes"] = args ? args.meshes : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["rules"] = args ? args.rules : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["effectiveLabels"] = undefined /*out*/;
            resourceInputs["pulumiLabels"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["effectiveLabels", "pulumiLabels"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(TcpRoute.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TcpRoute resources.
 */
export interface TcpRouteState {
    /**
     * Time the TcpRoute was created in UTC.
     */
    createTime?: pulumi.Input<string>;
    /**
     * A free-text description of the resource. Max length 1024 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
     * clients and services.
     */
    effectiveLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Gateways defines a list of gateways this TcpRoute is attached to, as one of the routing rules to route the requests served by the gateway.
     * Each gateway reference should match the pattern: projects/*&#47;locations/global/gateways/<gateway_name>
     */
    gateways?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Set of label tags associated with the TcpRoute resource.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Meshes defines a list of meshes this TcpRoute is attached to, as one of the routing rules to route the requests served by the mesh.
     * Each mesh reference should match the pattern: projects/*&#47;locations/global/meshes/<mesh_name>
     * The attached Mesh should be of a type SIDECAR
     */
    meshes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the TcpRoute resource.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     */
    pulumiLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Rules that define how traffic is routed and handled. At least one RouteRule must be supplied.
     * If there are multiple rules then the action taken will be the first rule to match.
     * Structure is documented below.
     */
    rules?: pulumi.Input<pulumi.Input<inputs.networkservices.TcpRouteRule>[]>;
    /**
     * Server-defined URL of this resource.
     */
    selfLink?: pulumi.Input<string>;
    /**
     * Time the TcpRoute was updated in UTC.
     */
    updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TcpRoute resource.
 */
export interface TcpRouteArgs {
    /**
     * A free-text description of the resource. Max length 1024 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Gateways defines a list of gateways this TcpRoute is attached to, as one of the routing rules to route the requests served by the gateway.
     * Each gateway reference should match the pattern: projects/*&#47;locations/global/gateways/<gateway_name>
     */
    gateways?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Set of label tags associated with the TcpRoute resource.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Meshes defines a list of meshes this TcpRoute is attached to, as one of the routing rules to route the requests served by the mesh.
     * Each mesh reference should match the pattern: projects/*&#47;locations/global/meshes/<mesh_name>
     * The attached Mesh should be of a type SIDECAR
     */
    meshes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the TcpRoute resource.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Rules that define how traffic is routed and handled. At least one RouteRule must be supplied.
     * If there are multiple rules then the action taken will be the first rule to match.
     * Structure is documented below.
     */
    rules: pulumi.Input<pulumi.Input<inputs.networkservices.TcpRouteRule>[]>;
}
