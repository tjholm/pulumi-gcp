// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Custom constraints are created by administrators to provide more granular and customizable control over the specific fields that are restricted by your organization policies.
 *
 * To get more information about CustomConstraint, see:
 *
 * * [API documentation](https://cloud.google.com/resource-manager/docs/reference/orgpolicy/rest/v2/organizations.constraints)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/resource-manager/docs/organization-policy/creating-managing-custom-constraints)
 *     * [Supported Services](https://cloud.google.com/resource-manager/docs/organization-policy/custom-constraint-supported-services)
 *
 * ## Example Usage
 * ### Org Policy Custom Constraint Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const constraint = new gcp.orgpolicy.CustomConstraint("constraint", {
 *     actionType: "ALLOW",
 *     condition: "resource.management.autoUpgrade == false",
 *     methodTypes: [
 *         "CREATE",
 *         "UPDATE",
 *     ],
 *     parent: "organizations/123456789",
 *     resourceTypes: ["container.googleapis.com/NodePool"],
 * });
 * ```
 * ### Org Policy Custom Constraint Full
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const constraint = new gcp.orgpolicy.CustomConstraint("constraint", {
 *     actionType: "ALLOW",
 *     condition: "resource.management.autoUpgrade == false",
 *     description: "Only allow GKE NodePool resource to be created or updated if AutoUpgrade is not enabled where this custom constraint is enforced.",
 *     displayName: "Disable GKE auto upgrade",
 *     methodTypes: [
 *         "CREATE",
 *         "UPDATE",
 *     ],
 *     parent: "organizations/123456789",
 *     resourceTypes: ["container.googleapis.com/NodePool"],
 * });
 * const bool = new gcp.orgpolicy.Policy("bool", {
 *     parent: "organizations/123456789",
 *     spec: {
 *         rules: [{
 *             enforce: "TRUE",
 *         }],
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * CustomConstraint can be imported using any of these accepted formats* `{{parent}}/customConstraints/{{name}}` When using the `pulumi import` command, CustomConstraint can be imported using one of the formats above. For example
 *
 * ```sh
 *  $ pulumi import gcp:orgpolicy/customConstraint:CustomConstraint default {{parent}}/customConstraints/{{name}}
 * ```
 */
export class CustomConstraint extends pulumi.CustomResource {
    /**
     * Get an existing CustomConstraint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CustomConstraintState, opts?: pulumi.CustomResourceOptions): CustomConstraint {
        return new CustomConstraint(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:orgpolicy/customConstraint:CustomConstraint';

    /**
     * Returns true if the given object is an instance of CustomConstraint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CustomConstraint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CustomConstraint.__pulumiType;
    }

    /**
     * The action to take if the condition is met.
     * Possible values are: `ALLOW`, `DENY`.
     */
    public readonly actionType!: pulumi.Output<string>;
    /**
     * A CEL condition that refers to a supported service resource, for example `resource.management.autoUpgrade == false`. For details about CEL usage, see [Common Expression Language](https://cloud.google.com/resource-manager/docs/organization-policy/creating-managing-custom-constraints#common_expression_language).
     */
    public readonly condition!: pulumi.Output<string>;
    /**
     * A human-friendly description of the constraint to display as an error message when the policy is violated.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A human-friendly name for the constraint.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * A list of RESTful methods for which to enforce the constraint. Can be `CREATE`, `UPDATE`, or both. Not all Google Cloud services support both methods. To see supported methods for each service, find the service in [Supported services](https://cloud.google.com/resource-manager/docs/organization-policy/custom-constraint-supported-services).
     */
    public readonly methodTypes!: pulumi.Output<string[]>;
    /**
     * Immutable. The name of the custom constraint. This is unique within the organization.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The parent of the resource, an organization. Format should be `organizations/{organization_id}`.
     *
     *
     * - - -
     */
    public readonly parent!: pulumi.Output<string>;
    /**
     * Immutable. The fully qualified name of the Google Cloud REST resource containing the object and field you want to restrict. For example, `container.googleapis.com/NodePool`.
     */
    public readonly resourceTypes!: pulumi.Output<string[]>;
    /**
     * Output only. The timestamp representing when the constraint was last updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a CustomConstraint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CustomConstraintArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CustomConstraintArgs | CustomConstraintState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CustomConstraintState | undefined;
            resourceInputs["actionType"] = state ? state.actionType : undefined;
            resourceInputs["condition"] = state ? state.condition : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["methodTypes"] = state ? state.methodTypes : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["parent"] = state ? state.parent : undefined;
            resourceInputs["resourceTypes"] = state ? state.resourceTypes : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as CustomConstraintArgs | undefined;
            if ((!args || args.actionType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'actionType'");
            }
            if ((!args || args.condition === undefined) && !opts.urn) {
                throw new Error("Missing required property 'condition'");
            }
            if ((!args || args.methodTypes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'methodTypes'");
            }
            if ((!args || args.parent === undefined) && !opts.urn) {
                throw new Error("Missing required property 'parent'");
            }
            if ((!args || args.resourceTypes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceTypes'");
            }
            resourceInputs["actionType"] = args ? args.actionType : undefined;
            resourceInputs["condition"] = args ? args.condition : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["methodTypes"] = args ? args.methodTypes : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["parent"] = args ? args.parent : undefined;
            resourceInputs["resourceTypes"] = args ? args.resourceTypes : undefined;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CustomConstraint.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CustomConstraint resources.
 */
export interface CustomConstraintState {
    /**
     * The action to take if the condition is met.
     * Possible values are: `ALLOW`, `DENY`.
     */
    actionType?: pulumi.Input<string>;
    /**
     * A CEL condition that refers to a supported service resource, for example `resource.management.autoUpgrade == false`. For details about CEL usage, see [Common Expression Language](https://cloud.google.com/resource-manager/docs/organization-policy/creating-managing-custom-constraints#common_expression_language).
     */
    condition?: pulumi.Input<string>;
    /**
     * A human-friendly description of the constraint to display as an error message when the policy is violated.
     */
    description?: pulumi.Input<string>;
    /**
     * A human-friendly name for the constraint.
     */
    displayName?: pulumi.Input<string>;
    /**
     * A list of RESTful methods for which to enforce the constraint. Can be `CREATE`, `UPDATE`, or both. Not all Google Cloud services support both methods. To see supported methods for each service, find the service in [Supported services](https://cloud.google.com/resource-manager/docs/organization-policy/custom-constraint-supported-services).
     */
    methodTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Immutable. The name of the custom constraint. This is unique within the organization.
     */
    name?: pulumi.Input<string>;
    /**
     * The parent of the resource, an organization. Format should be `organizations/{organization_id}`.
     *
     *
     * - - -
     */
    parent?: pulumi.Input<string>;
    /**
     * Immutable. The fully qualified name of the Google Cloud REST resource containing the object and field you want to restrict. For example, `container.googleapis.com/NodePool`.
     */
    resourceTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Output only. The timestamp representing when the constraint was last updated.
     */
    updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a CustomConstraint resource.
 */
export interface CustomConstraintArgs {
    /**
     * The action to take if the condition is met.
     * Possible values are: `ALLOW`, `DENY`.
     */
    actionType: pulumi.Input<string>;
    /**
     * A CEL condition that refers to a supported service resource, for example `resource.management.autoUpgrade == false`. For details about CEL usage, see [Common Expression Language](https://cloud.google.com/resource-manager/docs/organization-policy/creating-managing-custom-constraints#common_expression_language).
     */
    condition: pulumi.Input<string>;
    /**
     * A human-friendly description of the constraint to display as an error message when the policy is violated.
     */
    description?: pulumi.Input<string>;
    /**
     * A human-friendly name for the constraint.
     */
    displayName?: pulumi.Input<string>;
    /**
     * A list of RESTful methods for which to enforce the constraint. Can be `CREATE`, `UPDATE`, or both. Not all Google Cloud services support both methods. To see supported methods for each service, find the service in [Supported services](https://cloud.google.com/resource-manager/docs/organization-policy/custom-constraint-supported-services).
     */
    methodTypes: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Immutable. The name of the custom constraint. This is unique within the organization.
     */
    name?: pulumi.Input<string>;
    /**
     * The parent of the resource, an organization. Format should be `organizations/{organization_id}`.
     *
     *
     * - - -
     */
    parent: pulumi.Input<string>;
    /**
     * Immutable. The fully qualified name of the Google Cloud REST resource containing the object and field you want to restrict. For example, `container.googleapis.com/NodePool`.
     */
    resourceTypes: pulumi.Input<pulumi.Input<string>[]>;
}
