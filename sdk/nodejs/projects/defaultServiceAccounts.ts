// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Allows management of Google Cloud Platform project default service accounts.
 *
 * When certain service APIs are enabled, Google Cloud Platform automatically creates service accounts to help get started, but
 * this is not recommended for production environments as per [Google's documentation](https://cloud.google.com/iam/docs/service-accounts#default).
 * See the [Organization documentation](https://cloud.google.com/resource-manager/docs/quickstarts) for more details.
 *
 * > **WARNING** Some Google Cloud products do not work if the default service accounts are deleted so it is better to `DEPRIVILEGE` as
 * Google **CAN NOT** recover service accounts that have been deleted for more than 30 days.
 * Also Google recommends using the `constraints/iam.automaticIamGrantsForDefaultServiceAccounts` [constraint](https://www.terraform.io/docs/providers/google/r/google_organization_policy.html)
 * to disable automatic IAM Grants to default service accounts.
 *
 * > This resource works on a best-effort basis, as no API formally describes the default service accounts
 * and it is for users who are unable to use constraints. If the default service accounts change their name
 * or additional service accounts are added, this resource will need to be updated.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const myProject = new gcp.projects.DefaultServiceAccounts("myProject", {
 *     action: "DELETE",
 *     project: "my-project-id",
 * });
 * ```
 *
 * To enable the default service accounts on the resource destroy:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const myProject = new gcp.projects.DefaultServiceAccounts("myProject", {
 *     action: "DISABLE",
 *     project: "my-project-id",
 *     restorePolicy: "REVERT",
 * });
 * ```
 *
 * ## Import
 *
 * This resource does not support import
 */
export class DefaultServiceAccounts extends pulumi.CustomResource {
    /**
     * Get an existing DefaultServiceAccounts resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DefaultServiceAccountsState, opts?: pulumi.CustomResourceOptions): DefaultServiceAccounts {
        return new DefaultServiceAccounts(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:projects/defaultServiceAccounts:DefaultServiceAccounts';

    /**
     * Returns true if the given object is an instance of DefaultServiceAccounts.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DefaultServiceAccounts {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DefaultServiceAccounts.__pulumiType;
    }

    /**
     * The action to be performed in the default service accounts. Valid values are: `DEPRIVILEGE`, `DELETE`, `DISABLE`. Note that `DEPRIVILEGE` action will ignore the REVERT configuration in the restore_policy
     */
    public readonly action!: pulumi.Output<string>;
    /**
     * The project ID where service accounts are created.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The action to be performed in the default service accounts on the resource destroy.
     * Valid values are NONE, REVERT and REVERT_AND_IGNORE_FAILURE. It is applied for any action but in the DEPRIVILEGE.
     * If set to REVERT it attempts to restore all default SAs but the DEPRIVILEGE action.
     * If set to REVERT_AND_IGNORE_FAILURE it is the same behavior as REVERT but ignores errors returned by the API.
     */
    public readonly restorePolicy!: pulumi.Output<string | undefined>;
    /**
     * The Service Accounts changed by this resource. It is used for `REVERT` the `action` on the destroy.
     */
    public /*out*/ readonly serviceAccounts!: pulumi.Output<{[key: string]: any}>;

    /**
     * Create a DefaultServiceAccounts resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DefaultServiceAccountsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DefaultServiceAccountsArgs | DefaultServiceAccountsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DefaultServiceAccountsState | undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["restorePolicy"] = state ? state.restorePolicy : undefined;
            resourceInputs["serviceAccounts"] = state ? state.serviceAccounts : undefined;
        } else {
            const args = argsOrState as DefaultServiceAccountsArgs | undefined;
            if ((!args || args.action === undefined) && !opts.urn) {
                throw new Error("Missing required property 'action'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["restorePolicy"] = args ? args.restorePolicy : undefined;
            resourceInputs["serviceAccounts"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DefaultServiceAccounts.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DefaultServiceAccounts resources.
 */
export interface DefaultServiceAccountsState {
    /**
     * The action to be performed in the default service accounts. Valid values are: `DEPRIVILEGE`, `DELETE`, `DISABLE`. Note that `DEPRIVILEGE` action will ignore the REVERT configuration in the restore_policy
     */
    action?: pulumi.Input<string>;
    /**
     * The project ID where service accounts are created.
     */
    project?: pulumi.Input<string>;
    /**
     * The action to be performed in the default service accounts on the resource destroy.
     * Valid values are NONE, REVERT and REVERT_AND_IGNORE_FAILURE. It is applied for any action but in the DEPRIVILEGE.
     * If set to REVERT it attempts to restore all default SAs but the DEPRIVILEGE action.
     * If set to REVERT_AND_IGNORE_FAILURE it is the same behavior as REVERT but ignores errors returned by the API.
     */
    restorePolicy?: pulumi.Input<string>;
    /**
     * The Service Accounts changed by this resource. It is used for `REVERT` the `action` on the destroy.
     */
    serviceAccounts?: pulumi.Input<{[key: string]: any}>;
}

/**
 * The set of arguments for constructing a DefaultServiceAccounts resource.
 */
export interface DefaultServiceAccountsArgs {
    /**
     * The action to be performed in the default service accounts. Valid values are: `DEPRIVILEGE`, `DELETE`, `DISABLE`. Note that `DEPRIVILEGE` action will ignore the REVERT configuration in the restore_policy
     */
    action: pulumi.Input<string>;
    /**
     * The project ID where service accounts are created.
     */
    project: pulumi.Input<string>;
    /**
     * The action to be performed in the default service accounts on the resource destroy.
     * Valid values are NONE, REVERT and REVERT_AND_IGNORE_FAILURE. It is applied for any action but in the DEPRIVILEGE.
     * If set to REVERT it attempts to restore all default SAs but the DEPRIVILEGE action.
     * If set to REVERT_AND_IGNORE_FAILURE it is the same behavior as REVERT but ignores errors returned by the API.
     */
    restorePolicy?: pulumi.Input<string>;
}
