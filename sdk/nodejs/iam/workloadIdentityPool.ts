// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Represents a collection of external workload identities. You can define IAM policies to
 * grant these identities access to Google Cloud resources.
 *
 * To get more information about WorkloadIdentityPool, see:
 *
 * * [API documentation](https://cloud.google.com/iam/docs/reference/rest/v1beta/projects.locations.workloadIdentityPools)
 * * How-to Guides
 *     * [Managing workload identity pools](https://cloud.google.com/iam/docs/manage-workload-identity-pools-providers#pools)
 *
 * ## Example Usage
 * ### Iam Workload Identity Pool Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const example = new gcp.iam.WorkloadIdentityPool("example", {workloadIdentityPoolId: "example-pool"}, {
 *     provider: google_beta,
 * });
 * ```
 * ### Iam Workload Identity Pool Full
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const example = new gcp.iam.WorkloadIdentityPool("example", {
 *     workloadIdentityPoolId: "example-pool",
 *     displayName: "Name of pool",
 *     description: "Identity pool for automated test",
 *     disabled: true,
 * }, {
 *     provider: google_beta,
 * });
 * ```
 *
 * ## Import
 *
 * WorkloadIdentityPool can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:iam/workloadIdentityPool:WorkloadIdentityPool default projects/{{project}}/locations/global/workloadIdentityPools/{{workload_identity_pool_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:iam/workloadIdentityPool:WorkloadIdentityPool default {{project}}/{{workload_identity_pool_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:iam/workloadIdentityPool:WorkloadIdentityPool default {{workload_identity_pool_id}}
 * ```
 */
export class WorkloadIdentityPool extends pulumi.CustomResource {
    /**
     * Get an existing WorkloadIdentityPool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WorkloadIdentityPoolState, opts?: pulumi.CustomResourceOptions): WorkloadIdentityPool {
        return new WorkloadIdentityPool(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:iam/workloadIdentityPool:WorkloadIdentityPool';

    /**
     * Returns true if the given object is an instance of WorkloadIdentityPool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WorkloadIdentityPool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WorkloadIdentityPool.__pulumiType;
    }

    /**
     * A description of the pool. Cannot exceed 256 characters.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Whether the pool is disabled. You cannot use a disabled pool to exchange tokens, or use
     * existing tokens to access resources. If the pool is re-enabled, existing tokens grant
     * access again.
     */
    public readonly disabled!: pulumi.Output<boolean | undefined>;
    /**
     * A display name for the pool. Cannot exceed 32 characters.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * The resource name of the pool as
     * 'projects/{project_number}/locations/global/workloadIdentityPools/{workload_identity_pool_id}'.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The state of the pool. * STATE_UNSPECIFIED: State unspecified. * ACTIVE: The pool is active, and may be used in Google
     * Cloud policies. * DELETED: The pool is soft-deleted. Soft-deleted pools are permanently deleted after approximately 30
     * days. You can restore a soft-deleted pool using UndeleteWorkloadIdentityPool. You cannot reuse the ID of a soft-deleted
     * pool until it is permanently deleted. While a pool is deleted, you cannot use it to exchange tokens, or use existing
     * tokens to access resources. If the pool is undeleted, existing tokens grant access again.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The ID to use for the pool, which becomes the final component of the resource name. This
     * value should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
     * `gcp-` is reserved for use by Google, and may not be specified.
     */
    public readonly workloadIdentityPoolId!: pulumi.Output<string>;

    /**
     * Create a WorkloadIdentityPool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WorkloadIdentityPoolArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WorkloadIdentityPoolArgs | WorkloadIdentityPoolState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WorkloadIdentityPoolState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["disabled"] = state ? state.disabled : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["workloadIdentityPoolId"] = state ? state.workloadIdentityPoolId : undefined;
        } else {
            const args = argsOrState as WorkloadIdentityPoolArgs | undefined;
            if ((!args || args.workloadIdentityPoolId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workloadIdentityPoolId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["disabled"] = args ? args.disabled : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["workloadIdentityPoolId"] = args ? args.workloadIdentityPoolId : undefined;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WorkloadIdentityPool.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WorkloadIdentityPool resources.
 */
export interface WorkloadIdentityPoolState {
    /**
     * A description of the pool. Cannot exceed 256 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether the pool is disabled. You cannot use a disabled pool to exchange tokens, or use
     * existing tokens to access resources. If the pool is re-enabled, existing tokens grant
     * access again.
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * A display name for the pool. Cannot exceed 32 characters.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The resource name of the pool as
     * 'projects/{project_number}/locations/global/workloadIdentityPools/{workload_identity_pool_id}'.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The state of the pool. * STATE_UNSPECIFIED: State unspecified. * ACTIVE: The pool is active, and may be used in Google
     * Cloud policies. * DELETED: The pool is soft-deleted. Soft-deleted pools are permanently deleted after approximately 30
     * days. You can restore a soft-deleted pool using UndeleteWorkloadIdentityPool. You cannot reuse the ID of a soft-deleted
     * pool until it is permanently deleted. While a pool is deleted, you cannot use it to exchange tokens, or use existing
     * tokens to access resources. If the pool is undeleted, existing tokens grant access again.
     */
    state?: pulumi.Input<string>;
    /**
     * The ID to use for the pool, which becomes the final component of the resource name. This
     * value should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
     * `gcp-` is reserved for use by Google, and may not be specified.
     */
    workloadIdentityPoolId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WorkloadIdentityPool resource.
 */
export interface WorkloadIdentityPoolArgs {
    /**
     * A description of the pool. Cannot exceed 256 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether the pool is disabled. You cannot use a disabled pool to exchange tokens, or use
     * existing tokens to access resources. If the pool is re-enabled, existing tokens grant
     * access again.
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * A display name for the pool. Cannot exceed 32 characters.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The ID to use for the pool, which becomes the final component of the resource name. This
     * value should be 4-32 characters, and may contain the characters [a-z0-9-]. The prefix
     * `gcp-` is reserved for use by Google, and may not be specified.
     */
    workloadIdentityPoolId: pulumi.Input<string>;
}
