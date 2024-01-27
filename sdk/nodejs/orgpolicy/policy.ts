// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * An organization policy gives you programmatic control over your organization's cloud resources.  Using Organization Policies, you will be able to configure constraints across your entire resource hierarchy.
 *
 * For more information, see:
 * * [Understanding Org Policy concepts](https://cloud.google.com/resource-manager/docs/organization-policy/overview)
 * * [The resource hierarchy](https://cloud.google.com/resource-manager/docs/cloud-platform-resource-hierarchy)
 * * [All valid constraints](https://cloud.google.com/resource-manager/docs/organization-policy/org-policy-constraints)
 * ## Example Usage
 * ### Enforce_policy
 * A test of an enforce orgpolicy policy for a project
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const basic = new gcp.organizations.Project("basic", {orgId: "123456789"});
 * const primary = new gcp.orgpolicy.Policy("primary", {
 *     parent: pulumi.interpolate`projects/${basic.name}`,
 *     spec: {
 *         rules: [{
 *             enforce: "FALSE",
 *         }],
 *     },
 * });
 * ```
 * ### Folder_policy
 * A test of an orgpolicy policy for a folder
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const basic = new gcp.organizations.Folder("basic", {
 *     parent: "organizations/123456789",
 *     displayName: "folder",
 * });
 * const primary = new gcp.orgpolicy.Policy("primary", {
 *     parent: basic.name,
 *     spec: {
 *         inheritFromParent: true,
 *         rules: [{
 *             denyAll: "TRUE",
 *         }],
 *     },
 * });
 * ```
 * ### Organization_policy
 * A test of an orgpolicy policy for an organization
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const primary = new gcp.orgpolicy.Policy("primary", {
 *     parent: "organizations/123456789",
 *     spec: {
 *         reset: true,
 *     },
 * });
 * ```
 * ### Project_policy
 * A test of an orgpolicy policy for a project
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const basic = new gcp.organizations.Project("basic", {orgId: "123456789"});
 * const primary = new gcp.orgpolicy.Policy("primary", {
 *     parent: pulumi.interpolate`projects/${basic.name}`,
 *     spec: {
 *         rules: [
 *             {
 *                 condition: {
 *                     description: "A sample condition for the policy",
 *                     expression: "resource.matchLabels('labelKeys/123', 'labelValues/345')",
 *                     location: "sample-location.log",
 *                     title: "sample-condition",
 *                 },
 *                 values: {
 *                     allowedValues: ["projects/allowed-project"],
 *                     deniedValues: ["projects/denied-project"],
 *                 },
 *             },
 *             {
 *                 allowAll: "TRUE",
 *             },
 *         ],
 *     },
 * });
 * ```
 * ### Dry_run_spec
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const constraint = new gcp.orgpolicy.CustomConstraint("constraint", {
 *     actionType: "ALLOW",
 *     condition: "resource.management.autoUpgrade == false",
 *     description: "Only allow GKE NodePool resource to be created or updated if AutoUpgrade is not enabled where this custom constraint is enforced.",
 *     displayName: "Disable GKE auto upgrade",
 *     methodTypes: ["CREATE"],
 *     parent: "organizations/123456789",
 *     resourceTypes: ["container.googleapis.com/NodePool"],
 * });
 * const primary = new gcp.orgpolicy.Policy("primary", {
 *     dryRunSpec: {
 *         inheritFromParent: false,
 *         reset: false,
 *         rules: [{
 *             enforce: "FALSE",
 *         }],
 *     },
 *     parent: "organizations/123456789",
 *     spec: {
 *         rules: [{
 *             enforce: "FALSE",
 *         }],
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Policy can be imported using any of these accepted formats* `{{parent}}/policies/{{name}}`
 *
 * When using the `pulumi import` command, Policy can be imported using one of the formats above. For example:
 *
 * ```sh
 *  $ pulumi import gcp:orgpolicy/policy:Policy default {{parent}}/policies/{{name}}
 * ```
 */
export class Policy extends pulumi.CustomResource {
    /**
     * Get an existing Policy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: PolicyState, opts?: pulumi.CustomResourceOptions): Policy {
        return new Policy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:orgpolicy/policy:Policy';

    /**
     * Returns true if the given object is an instance of Policy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Policy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Policy.__pulumiType;
    }

    /**
     * Dry-run policy. Audit-only policy, can be used to monitor how the policy would have impacted the existing and future resources if it's enforced.
     */
    public readonly dryRunSpec!: pulumi.Output<outputs.orgpolicy.PolicyDryRunSpec | undefined>;
    /**
     * An opaque tag indicating the current version of the policy, used for concurrency control. This field is ignored if used in a `CreatePolicy` request. When the policy` is returned from either a `GetPolicy` or a `ListPolicies` request, this `etag` indicates the version of the current policy to use when executing a read-modify-write loop. When the policy is returned from a `GetEffectivePolicy` request, the `etag` will be unset.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Immutable. The resource name of the Policy. Must be one of the following forms, where constraintName is the name of the constraint which this Policy configures: * `projects/{project_number}/policies/{constraint_name}` * `folders/{folder_id}/policies/{constraint_name}` * `organizations/{organization_id}/policies/{constraint_name}` For example, "projects/123/policies/compute.disableSerialPortAccess". Note: `projects/{project_id}/policies/{constraint_name}` is also an acceptable name for API requests, but responses will return the name using the equivalent project number.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The parent of the resource.
     *
     *
     *
     * - - -
     */
    public readonly parent!: pulumi.Output<string>;
    /**
     * Basic information about the Organization Policy.
     */
    public readonly spec!: pulumi.Output<outputs.orgpolicy.PolicySpec | undefined>;

    /**
     * Create a Policy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PolicyArgs | PolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as PolicyState | undefined;
            resourceInputs["dryRunSpec"] = state ? state.dryRunSpec : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["parent"] = state ? state.parent : undefined;
            resourceInputs["spec"] = state ? state.spec : undefined;
        } else {
            const args = argsOrState as PolicyArgs | undefined;
            if ((!args || args.parent === undefined) && !opts.urn) {
                throw new Error("Missing required property 'parent'");
            }
            resourceInputs["dryRunSpec"] = args ? args.dryRunSpec : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["parent"] = args ? args.parent : undefined;
            resourceInputs["spec"] = args ? args.spec : undefined;
            resourceInputs["etag"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Policy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Policy resources.
 */
export interface PolicyState {
    /**
     * Dry-run policy. Audit-only policy, can be used to monitor how the policy would have impacted the existing and future resources if it's enforced.
     */
    dryRunSpec?: pulumi.Input<inputs.orgpolicy.PolicyDryRunSpec>;
    /**
     * An opaque tag indicating the current version of the policy, used for concurrency control. This field is ignored if used in a `CreatePolicy` request. When the policy` is returned from either a `GetPolicy` or a `ListPolicies` request, this `etag` indicates the version of the current policy to use when executing a read-modify-write loop. When the policy is returned from a `GetEffectivePolicy` request, the `etag` will be unset.
     */
    etag?: pulumi.Input<string>;
    /**
     * Immutable. The resource name of the Policy. Must be one of the following forms, where constraintName is the name of the constraint which this Policy configures: * `projects/{project_number}/policies/{constraint_name}` * `folders/{folder_id}/policies/{constraint_name}` * `organizations/{organization_id}/policies/{constraint_name}` For example, "projects/123/policies/compute.disableSerialPortAccess". Note: `projects/{project_id}/policies/{constraint_name}` is also an acceptable name for API requests, but responses will return the name using the equivalent project number.
     */
    name?: pulumi.Input<string>;
    /**
     * The parent of the resource.
     *
     *
     *
     * - - -
     */
    parent?: pulumi.Input<string>;
    /**
     * Basic information about the Organization Policy.
     */
    spec?: pulumi.Input<inputs.orgpolicy.PolicySpec>;
}

/**
 * The set of arguments for constructing a Policy resource.
 */
export interface PolicyArgs {
    /**
     * Dry-run policy. Audit-only policy, can be used to monitor how the policy would have impacted the existing and future resources if it's enforced.
     */
    dryRunSpec?: pulumi.Input<inputs.orgpolicy.PolicyDryRunSpec>;
    /**
     * Immutable. The resource name of the Policy. Must be one of the following forms, where constraintName is the name of the constraint which this Policy configures: * `projects/{project_number}/policies/{constraint_name}` * `folders/{folder_id}/policies/{constraint_name}` * `organizations/{organization_id}/policies/{constraint_name}` For example, "projects/123/policies/compute.disableSerialPortAccess". Note: `projects/{project_id}/policies/{constraint_name}` is also an acceptable name for API requests, but responses will return the name using the equivalent project number.
     */
    name?: pulumi.Input<string>;
    /**
     * The parent of the resource.
     *
     *
     *
     * - - -
     */
    parent: pulumi.Input<string>;
    /**
     * Basic information about the Organization Policy.
     */
    spec?: pulumi.Input<inputs.orgpolicy.PolicySpec>;
}
