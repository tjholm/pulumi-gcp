// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A Posture represents a collection of policy set including its name, state, description
 * and policy sets. A policy set includes set of policies along with their definition.
 * A posture can be created at the organization level.
 * Every update to a deployed posture creates a new posture revision with an updated revision_id.
 *
 * To get more information about Posture, see:
 *
 * * How-to Guides
 *     * [Create and deploy a posture](https://cloud.google.com/security-command-center/docs/how-to-use-security-posture)
 *
 * ## Example Usage
 * ### Securityposture Posture Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const posture1 = new gcp.securityposture.Posture("posture1", {
 *     description: "a new posture",
 *     location: "global",
 *     parent: "organizations/123456789",
 *     policySets: [
 *         {
 *             description: "set of org policies",
 *             policies: [
 *                 {
 *                     constraint: {
 *                         orgPolicyConstraint: {
 *                             cannedConstraintId: "storage.uniformBucketLevelAccess",
 *                             policyRules: [{
 *                                 enforce: true,
 *                             }],
 *                         },
 *                     },
 *                     policyId: "canned_org_policy",
 *                 },
 *                 {
 *                     constraint: {
 *                         orgPolicyConstraintCustom: {
 *                             customConstraint: {
 *                                 actionType: "ALLOW",
 *                                 condition: "resource.management.autoUpgrade == false",
 *                                 description: "Only allow GKE NodePool resource to be created or updated if AutoUpgrade is not enabled where this custom constraint is enforced.",
 *                                 displayName: "Disable GKE auto upgrade",
 *                                 methodTypes: [
 *                                     "CREATE",
 *                                     "UPDATE",
 *                                 ],
 *                                 name: "organizations/123456789/customConstraints/custom.disableGkeAutoUpgrade",
 *                                 resourceTypes: ["container.googleapis.com/NodePool"],
 *                             },
 *                             policyRules: [{
 *                                 enforce: true,
 *                             }],
 *                         },
 *                     },
 *                     policyId: "custom_org_policy",
 *                 },
 *             ],
 *             policySetId: "org_policy_set",
 *         },
 *         {
 *             description: "set of sha policies",
 *             policies: [
 *                 {
 *                     constraint: {
 *                         securityHealthAnalyticsModule: {
 *                             moduleEnablementState: "ENABLED",
 *                             moduleName: "BIGQUERY_TABLE_CMEK_DISABLED",
 *                         },
 *                     },
 *                     description: "enable BIGQUERY_TABLE_CMEK_DISABLED",
 *                     policyId: "sha_builtin_module",
 *                 },
 *                 {
 *                     constraint: {
 *                         securityHealthAnalyticsCustomModule: {
 *                             config: {
 *                                 customOutput: {
 *                                     properties: [{
 *                                         name: "duration",
 *                                         valueExpression: {
 *                                             expression: "resource.rotationPeriod",
 *                                         },
 *                                     }],
 *                                 },
 *                                 description: "Custom Module",
 *                                 predicate: {
 *                                     expression: "resource.rotationPeriod > duration('2592000s')",
 *                                 },
 *                                 recommendation: "Testing custom modules",
 *                                 resourceSelector: {
 *                                     resourceTypes: ["cloudkms.googleapis.com/CryptoKey"],
 *                                 },
 *                                 severity: "LOW",
 *                             },
 *                             displayName: "custom SHA policy",
 *                             moduleEnablementState: "ENABLED",
 *                         },
 *                     },
 *                     policyId: "sha_custom_module",
 *                 },
 *             ],
 *             policySetId: "sha_policy_set",
 *         },
 *     ],
 *     postureId: "posture_1",
 *     state: "ACTIVE",
 * });
 * ```
 *
 * ## Import
 *
 * Posture can be imported using any of these accepted formats* `{{parent}}/locations/{{location}}/postures/{{posture_id}}` When using the `pulumi import` command, Posture can be imported using one of the formats above. For example
 *
 * ```sh
 *  $ pulumi import gcp:securityposture/posture:Posture default {{parent}}/locations/{{location}}/postures/{{posture_id}}
 * ```
 */
export class Posture extends pulumi.CustomResource {
    /**
     * Get an existing Posture resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PostureState, opts?: pulumi.CustomResourceOptions): Posture {
        return new Posture(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:securityposture/posture:Posture';

    /**
     * Returns true if the given object is an instance of Posture.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Posture {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Posture.__pulumiType;
    }

    /**
     * Time the Posture was created in UTC.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Description of the posture.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * For Resource freshness validation (https://google.aip.dev/154)
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Location of the resource, eg: global.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Immutable. The name of the custom constraint. This is unique within the organization.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The parent of the resource, an organization. Format should be `organizations/{organization_id}`.
     */
    public readonly parent!: pulumi.Output<string>;
    /**
     * List of policy sets for the posture.
     * Structure is documented below.
     */
    public readonly policySets!: pulumi.Output<outputs.securityposture.PosturePolicySet[] | undefined>;
    /**
     * Id of the posture. It is an immutable field.
     *
     *
     * - - -
     */
    public readonly postureId!: pulumi.Output<string>;
    /**
     * If set, there are currently changes in flight to the posture.
     */
    public /*out*/ readonly reconciling!: pulumi.Output<boolean>;
    /**
     * Revision_id of the posture.
     */
    public /*out*/ readonly revisionId!: pulumi.Output<string>;
    /**
     * State of the posture. Update to state field should not be triggered along with
     * with other field updates.
     * Possible values are: `DEPRECATED`, `DRAFT`, `ACTIVE`.
     */
    public readonly state!: pulumi.Output<string>;
    /**
     * Time the Posture was updated in UTC.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Posture resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PostureArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PostureArgs | PostureState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PostureState | undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["parent"] = state ? state.parent : undefined;
            resourceInputs["policySets"] = state ? state.policySets : undefined;
            resourceInputs["postureId"] = state ? state.postureId : undefined;
            resourceInputs["reconciling"] = state ? state.reconciling : undefined;
            resourceInputs["revisionId"] = state ? state.revisionId : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as PostureArgs | undefined;
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.parent === undefined) && !opts.urn) {
                throw new Error("Missing required property 'parent'");
            }
            if ((!args || args.postureId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'postureId'");
            }
            if ((!args || args.state === undefined) && !opts.urn) {
                throw new Error("Missing required property 'state'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["parent"] = args ? args.parent : undefined;
            resourceInputs["policySets"] = args ? args.policySets : undefined;
            resourceInputs["postureId"] = args ? args.postureId : undefined;
            resourceInputs["state"] = args ? args.state : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["reconciling"] = undefined /*out*/;
            resourceInputs["revisionId"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Posture.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Posture resources.
 */
export interface PostureState {
    /**
     * Time the Posture was created in UTC.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Description of the posture.
     */
    description?: pulumi.Input<string>;
    /**
     * For Resource freshness validation (https://google.aip.dev/154)
     */
    etag?: pulumi.Input<string>;
    /**
     * Location of the resource, eg: global.
     */
    location?: pulumi.Input<string>;
    /**
     * Immutable. The name of the custom constraint. This is unique within the organization.
     */
    name?: pulumi.Input<string>;
    /**
     * The parent of the resource, an organization. Format should be `organizations/{organization_id}`.
     */
    parent?: pulumi.Input<string>;
    /**
     * List of policy sets for the posture.
     * Structure is documented below.
     */
    policySets?: pulumi.Input<pulumi.Input<inputs.securityposture.PosturePolicySet>[]>;
    /**
     * Id of the posture. It is an immutable field.
     *
     *
     * - - -
     */
    postureId?: pulumi.Input<string>;
    /**
     * If set, there are currently changes in flight to the posture.
     */
    reconciling?: pulumi.Input<boolean>;
    /**
     * Revision_id of the posture.
     */
    revisionId?: pulumi.Input<string>;
    /**
     * State of the posture. Update to state field should not be triggered along with
     * with other field updates.
     * Possible values are: `DEPRECATED`, `DRAFT`, `ACTIVE`.
     */
    state?: pulumi.Input<string>;
    /**
     * Time the Posture was updated in UTC.
     */
    updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Posture resource.
 */
export interface PostureArgs {
    /**
     * Description of the posture.
     */
    description?: pulumi.Input<string>;
    /**
     * Location of the resource, eg: global.
     */
    location: pulumi.Input<string>;
    /**
     * The parent of the resource, an organization. Format should be `organizations/{organization_id}`.
     */
    parent: pulumi.Input<string>;
    /**
     * List of policy sets for the posture.
     * Structure is documented below.
     */
    policySets?: pulumi.Input<pulumi.Input<inputs.securityposture.PosturePolicySet>[]>;
    /**
     * Id of the posture. It is an immutable field.
     *
     *
     * - - -
     */
    postureId: pulumi.Input<string>;
    /**
     * State of the posture. Update to state field should not be triggered along with
     * with other field updates.
     * Possible values are: `DEPRECATED`, `DRAFT`, `ACTIVE`.
     */
    state: pulumi.Input<string>;
}
