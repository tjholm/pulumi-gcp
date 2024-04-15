// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for Identity-Aware Proxy TunnelDestGroup. Each of these resources serves a different use case:
 *
 * * `gcp.iap.TunnelDestGroupIamPolicy`: Authoritative. Sets the IAM policy for the tunneldestgroup and replaces any existing policy already attached.
 * * `gcp.iap.TunnelDestGroupIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the tunneldestgroup are preserved.
 * * `gcp.iap.TunnelDestGroupIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the tunneldestgroup are preserved.
 *
 * A data source can be used to retrieve policy data in advent you do not need creation
 *
 * * `gcp.iap.TunnelDestGroupIamPolicy`: Retrieves the IAM policy for the tunneldestgroup
 *
 * > **Note:** `gcp.iap.TunnelDestGroupIamPolicy` **cannot** be used in conjunction with `gcp.iap.TunnelDestGroupIamBinding` and `gcp.iap.TunnelDestGroupIamMember` or they will fight over what your policy should be.
 *
 * > **Note:** `gcp.iap.TunnelDestGroupIamBinding` resources **can be** used in conjunction with `gcp.iap.TunnelDestGroupIamMember` resources **only if** they do not grant privilege to the same role.
 *
 * > **Note:**  This resource supports IAM Conditions but they have some known limitations which can be found [here](https://cloud.google.com/iam/docs/conditions-overview#limitations). Please review this article if you are having issues with IAM Conditions.
 *
 * ## google\_iap\_tunnel\_dest\_group\_iam\_policy
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const admin = gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         role: "roles/iap.tunnelResourceAccessor",
 *         members: ["user:jane@example.com"],
 *     }],
 * });
 * const policy = new gcp.iap.TunnelDestGroupIamPolicy("policy", {
 *     project: destGroup.project,
 *     region: destGroup.region,
 *     destGroup: destGroup.groupName,
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * With IAM Conditions:
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const admin = gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         role: "roles/iap.tunnelResourceAccessor",
 *         members: ["user:jane@example.com"],
 *         condition: {
 *             title: "expires_after_2019_12_31",
 *             description: "Expiring at midnight of 2019-12-31",
 *             expression: "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
 *         },
 *     }],
 * });
 * const policy = new gcp.iap.TunnelDestGroupIamPolicy("policy", {
 *     project: destGroup.project,
 *     region: destGroup.region,
 *     destGroup: destGroup.groupName,
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 * ## google\_iap\_tunnel\_dest\_group\_iam\_binding
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const binding = new gcp.iap.TunnelDestGroupIamBinding("binding", {
 *     project: destGroup.project,
 *     region: destGroup.region,
 *     destGroup: destGroup.groupName,
 *     role: "roles/iap.tunnelResourceAccessor",
 *     members: ["user:jane@example.com"],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * With IAM Conditions:
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const binding = new gcp.iap.TunnelDestGroupIamBinding("binding", {
 *     project: destGroup.project,
 *     region: destGroup.region,
 *     destGroup: destGroup.groupName,
 *     role: "roles/iap.tunnelResourceAccessor",
 *     members: ["user:jane@example.com"],
 *     condition: {
 *         title: "expires_after_2019_12_31",
 *         description: "Expiring at midnight of 2019-12-31",
 *         expression: "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
 *     },
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 * ## google\_iap\_tunnel\_dest\_group\_iam\_member
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const member = new gcp.iap.TunnelDestGroupIamMember("member", {
 *     project: destGroup.project,
 *     region: destGroup.region,
 *     destGroup: destGroup.groupName,
 *     role: "roles/iap.tunnelResourceAccessor",
 *     member: "user:jane@example.com",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * With IAM Conditions:
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const member = new gcp.iap.TunnelDestGroupIamMember("member", {
 *     project: destGroup.project,
 *     region: destGroup.region,
 *     destGroup: destGroup.groupName,
 *     role: "roles/iap.tunnelResourceAccessor",
 *     member: "user:jane@example.com",
 *     condition: {
 *         title: "expires_after_2019_12_31",
 *         description: "Expiring at midnight of 2019-12-31",
 *         expression: "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
 *     },
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## google\_iap\_tunnel\_dest\_group\_iam\_policy
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const admin = gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         role: "roles/iap.tunnelResourceAccessor",
 *         members: ["user:jane@example.com"],
 *     }],
 * });
 * const policy = new gcp.iap.TunnelDestGroupIamPolicy("policy", {
 *     project: destGroup.project,
 *     region: destGroup.region,
 *     destGroup: destGroup.groupName,
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * With IAM Conditions:
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const admin = gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         role: "roles/iap.tunnelResourceAccessor",
 *         members: ["user:jane@example.com"],
 *         condition: {
 *             title: "expires_after_2019_12_31",
 *             description: "Expiring at midnight of 2019-12-31",
 *             expression: "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
 *         },
 *     }],
 * });
 * const policy = new gcp.iap.TunnelDestGroupIamPolicy("policy", {
 *     project: destGroup.project,
 *     region: destGroup.region,
 *     destGroup: destGroup.groupName,
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 * ## google\_iap\_tunnel\_dest\_group\_iam\_binding
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const binding = new gcp.iap.TunnelDestGroupIamBinding("binding", {
 *     project: destGroup.project,
 *     region: destGroup.region,
 *     destGroup: destGroup.groupName,
 *     role: "roles/iap.tunnelResourceAccessor",
 *     members: ["user:jane@example.com"],
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * With IAM Conditions:
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const binding = new gcp.iap.TunnelDestGroupIamBinding("binding", {
 *     project: destGroup.project,
 *     region: destGroup.region,
 *     destGroup: destGroup.groupName,
 *     role: "roles/iap.tunnelResourceAccessor",
 *     members: ["user:jane@example.com"],
 *     condition: {
 *         title: "expires_after_2019_12_31",
 *         description: "Expiring at midnight of 2019-12-31",
 *         expression: "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
 *     },
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 * ## google\_iap\_tunnel\_dest\_group\_iam\_member
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const member = new gcp.iap.TunnelDestGroupIamMember("member", {
 *     project: destGroup.project,
 *     region: destGroup.region,
 *     destGroup: destGroup.groupName,
 *     role: "roles/iap.tunnelResourceAccessor",
 *     member: "user:jane@example.com",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * With IAM Conditions:
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const member = new gcp.iap.TunnelDestGroupIamMember("member", {
 *     project: destGroup.project,
 *     region: destGroup.region,
 *     destGroup: destGroup.groupName,
 *     role: "roles/iap.tunnelResourceAccessor",
 *     member: "user:jane@example.com",
 *     condition: {
 *         title: "expires_after_2019_12_31",
 *         description: "Expiring at midnight of 2019-12-31",
 *         expression: "request.time < timestamp(\"2020-01-01T00:00:00Z\")",
 *     },
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * For all import syntaxes, the "resource in question" can take any of the following forms:
 *
 * * projects/{{project}}/iap_tunnel/locations/{{region}}/destGroups/{{dest_group}}
 *
 * * {{project}}/iap_tunnel/locations/{{region}}/destGroups/{{dest_group}}
 *
 * * {{project}}/{{region}}/{{dest_group}}
 *
 * * {{region}}/{{dest_group}}
 *
 * * {{dest_group}}
 *
 * Any variables not passed in the import command will be taken from the provider configuration.
 *
 * Identity-Aware Proxy tunneldestgroup IAM resources can be imported using the resource identifiers, role, and member.
 *
 * IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
 *
 * ```sh
 * $ pulumi import gcp:iap/tunnelDestGroupIamPolicy:TunnelDestGroupIamPolicy editor "projects/{{project}}/iap_tunnel/locations/{{region}}/destGroups/{{dest_group}} roles/iap.tunnelResourceAccessor user:jane@example.com"
 * ```
 *
 * IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
 *
 * ```sh
 * $ pulumi import gcp:iap/tunnelDestGroupIamPolicy:TunnelDestGroupIamPolicy editor "projects/{{project}}/iap_tunnel/locations/{{region}}/destGroups/{{dest_group}} roles/iap.tunnelResourceAccessor"
 * ```
 *
 * IAM policy imports use the identifier of the resource in question, e.g.
 *
 * ```sh
 * $ pulumi import gcp:iap/tunnelDestGroupIamPolicy:TunnelDestGroupIamPolicy editor projects/{{project}}/iap_tunnel/locations/{{region}}/destGroups/{{dest_group}}
 * ```
 *
 * -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
 *
 *  full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 */
export class TunnelDestGroupIamPolicy extends pulumi.CustomResource {
    /**
     * Get an existing TunnelDestGroupIamPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TunnelDestGroupIamPolicyState, opts?: pulumi.CustomResourceOptions): TunnelDestGroupIamPolicy {
        return new TunnelDestGroupIamPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:iap/tunnelDestGroupIamPolicy:TunnelDestGroupIamPolicy';

    /**
     * Returns true if the given object is an instance of TunnelDestGroupIamPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TunnelDestGroupIamPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TunnelDestGroupIamPolicy.__pulumiType;
    }

    public readonly destGroup!: pulumi.Output<string>;
    /**
     * (Computed) The etag of the IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     */
    public readonly policyData!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The region of the tunnel group. Must be the same as the network resources in the group.
     * Used to find the parent resource to bind the IAM policy to. If not specified,
     * the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
     * region is specified, it is taken from the provider configuration.
     */
    public readonly region!: pulumi.Output<string>;

    /**
     * Create a TunnelDestGroupIamPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TunnelDestGroupIamPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TunnelDestGroupIamPolicyArgs | TunnelDestGroupIamPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TunnelDestGroupIamPolicyState | undefined;
            resourceInputs["destGroup"] = state ? state.destGroup : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["policyData"] = state ? state.policyData : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
        } else {
            const args = argsOrState as TunnelDestGroupIamPolicyArgs | undefined;
            if ((!args || args.destGroup === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destGroup'");
            }
            if ((!args || args.policyData === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyData'");
            }
            resourceInputs["destGroup"] = args ? args.destGroup : undefined;
            resourceInputs["policyData"] = args ? args.policyData : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["etag"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TunnelDestGroupIamPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TunnelDestGroupIamPolicy resources.
 */
export interface TunnelDestGroupIamPolicyState {
    destGroup?: pulumi.Input<string>;
    /**
     * (Computed) The etag of the IAM policy.
     */
    etag?: pulumi.Input<string>;
    /**
     * The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     */
    policyData?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The region of the tunnel group. Must be the same as the network resources in the group.
     * Used to find the parent resource to bind the IAM policy to. If not specified,
     * the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
     * region is specified, it is taken from the provider configuration.
     */
    region?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TunnelDestGroupIamPolicy resource.
 */
export interface TunnelDestGroupIamPolicyArgs {
    destGroup: pulumi.Input<string>;
    /**
     * The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     */
    policyData: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The region of the tunnel group. Must be the same as the network resources in the group.
     * Used to find the parent resource to bind the IAM policy to. If not specified,
     * the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
     * region is specified, it is taken from the provider configuration.
     */
    region?: pulumi.Input<string>;
}
