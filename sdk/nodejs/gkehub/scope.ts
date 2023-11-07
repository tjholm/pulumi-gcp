// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Scope represents a Scope in a Fleet.
 *
 * To get more information about Scope, see:
 *
 * * [API documentation](https://cloud.google.com/anthos/fleet-management/docs/reference/rest/v1/projects.locations.scopes)
 * * How-to Guides
 *     * [Registering a Cluster](https://cloud.google.com/anthos/multicluster-management/connect/registering-a-cluster#register_cluster)
 *
 * ## Example Usage
 * ### Gkehub Scope Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const scope = new gcp.gkehub.Scope("scope", {
 *     labels: {
 *         keya: "valuea",
 *         keyb: "valueb",
 *         keyc: "valuec",
 *     },
 *     scopeId: "tf-test-scope%{random_suffix}",
 * });
 * ```
 *
 * ## Import
 *
 * Scope can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:gkehub/scope:Scope default projects/{{project}}/locations/global/scopes/{{scope_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:gkehub/scope:Scope default {{project}}/{{scope_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:gkehub/scope:Scope default {{scope_id}}
 * ```
 */
export class Scope extends pulumi.CustomResource {
    /**
     * Get an existing Scope resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ScopeState, opts?: pulumi.CustomResourceOptions): Scope {
        return new Scope(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:gkehub/scope:Scope';

    /**
     * Returns true if the given object is an instance of Scope.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Scope {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Scope.__pulumiType;
    }

    /**
     * Time the Scope was created in UTC.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Time the Scope was deleted in UTC.
     */
    public /*out*/ readonly deleteTime!: pulumi.Output<string>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
     * clients and services.
     */
    public /*out*/ readonly effectiveLabels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Labels for this Scope.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The unique identifier of the scope
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
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
     * The client-provided identifier of the scope.
     *
     *
     * - - -
     */
    public readonly scopeId!: pulumi.Output<string>;
    /**
     * State of the scope resource.
     * Structure is documented below.
     */
    public /*out*/ readonly states!: pulumi.Output<outputs.gkehub.ScopeState[]>;
    /**
     * Google-generated UUID for this resource.
     */
    public /*out*/ readonly uid!: pulumi.Output<string>;
    /**
     * Time the Scope was updated in UTC.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Scope resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ScopeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ScopeArgs | ScopeState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ScopeState | undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["deleteTime"] = state ? state.deleteTime : undefined;
            resourceInputs["effectiveLabels"] = state ? state.effectiveLabels : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["pulumiLabels"] = state ? state.pulumiLabels : undefined;
            resourceInputs["scopeId"] = state ? state.scopeId : undefined;
            resourceInputs["states"] = state ? state.states : undefined;
            resourceInputs["uid"] = state ? state.uid : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as ScopeArgs | undefined;
            if ((!args || args.scopeId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'scopeId'");
            }
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["scopeId"] = args ? args.scopeId : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["deleteTime"] = undefined /*out*/;
            resourceInputs["effectiveLabels"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["pulumiLabels"] = undefined /*out*/;
            resourceInputs["states"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["effectiveLabels", "pulumiLabels"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Scope.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Scope resources.
 */
export interface ScopeState {
    /**
     * Time the Scope was created in UTC.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Time the Scope was deleted in UTC.
     */
    deleteTime?: pulumi.Input<string>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
     * clients and services.
     */
    effectiveLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Labels for this Scope.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The unique identifier of the scope
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
     * The client-provided identifier of the scope.
     *
     *
     * - - -
     */
    scopeId?: pulumi.Input<string>;
    /**
     * State of the scope resource.
     * Structure is documented below.
     */
    states?: pulumi.Input<pulumi.Input<inputs.gkehub.ScopeState>[]>;
    /**
     * Google-generated UUID for this resource.
     */
    uid?: pulumi.Input<string>;
    /**
     * Time the Scope was updated in UTC.
     */
    updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Scope resource.
 */
export interface ScopeArgs {
    /**
     * Labels for this Scope.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The client-provided identifier of the scope.
     *
     *
     * - - -
     */
    scopeId: pulumi.Input<string>;
}
