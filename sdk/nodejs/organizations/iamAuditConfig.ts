// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.
 *
 * This member resource can be imported using the `org_id`, role, and member e.g.
 *
 * ```sh
 *  $ pulumi import gcp:organizations/iamAuditConfig:IamAuditConfig my_organization "your-orgid roles/viewer user:foo@example.com"
 * ```
 *
 *  IAM binding imports use space-delimited identifiers; the resource in question and the role.
 *
 * This binding resource can be imported using the `org_id` and role, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:organizations/iamAuditConfig:IamAuditConfig my_organization "your-org-id roles/viewer"
 * ```
 *
 *  IAM policy imports use the identifier of the resource in question.
 *
 * This policy resource can be imported using the `org_id`.
 *
 * ```sh
 *  $ pulumi import gcp:organizations/iamAuditConfig:IamAuditConfig my_organization your-org-id
 * ```
 *
 *  IAM audit config imports use the identifier of the resource in question and the service, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:organizations/iamAuditConfig:IamAuditConfig my_organization "your-organization-id foo.googleapis.com"
 * ```
 *
 *  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
 *
 * full name of the custom role, e.g. `organizations/{{org_id}}/roles/{{role_id}}`. -> **Conditional IAM Bindings**If you're importing a IAM binding with a condition block, make sure
 *
 * ```sh
 *  $ pulumi import gcp:organizations/iamAuditConfig:IamAuditConfig to include the title of condition, e.g. `google_organization_iam_binding.my_organization "your-org-id roles/{{role_id}} condition-title"`
 * ```
 */
export class IamAuditConfig extends pulumi.CustomResource {
    /**
     * Get an existing IamAuditConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IamAuditConfigState, opts?: pulumi.CustomResourceOptions): IamAuditConfig {
        return new IamAuditConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:organizations/iamAuditConfig:IamAuditConfig';

    /**
     * Returns true if the given object is an instance of IamAuditConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IamAuditConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IamAuditConfig.__pulumiType;
    }

    /**
     * The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
     */
    public readonly auditLogConfigs!: pulumi.Output<outputs.organizations.IamAuditConfigAuditLogConfig[]>;
    /**
     * (Computed) The etag of the organization's IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The organization ID. If not specified for `gcp.organizations.IAMBinding`, `gcp.organizations.IAMMember`, or `gcp.organizations.IamAuditConfig`, uses the ID of the organization configured with the provider.
     * Required for `gcp.organizations.IAMPolicy` - you must explicitly set the organization, and it
     * will not be inferred from the provider.
     */
    public readonly orgId!: pulumi.Output<string>;
    /**
     * Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_organization\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `logTypes` specified in each `auditLogConfig` are enabled, and the `exemptedMembers` in each `auditLogConfig` are exempted.
     */
    public readonly service!: pulumi.Output<string>;

    /**
     * Create a IamAuditConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IamAuditConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IamAuditConfigArgs | IamAuditConfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IamAuditConfigState | undefined;
            resourceInputs["auditLogConfigs"] = state ? state.auditLogConfigs : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
            resourceInputs["service"] = state ? state.service : undefined;
        } else {
            const args = argsOrState as IamAuditConfigArgs | undefined;
            if ((!args || args.auditLogConfigs === undefined) && !opts.urn) {
                throw new Error("Missing required property 'auditLogConfigs'");
            }
            if ((!args || args.orgId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'orgId'");
            }
            if ((!args || args.service === undefined) && !opts.urn) {
                throw new Error("Missing required property 'service'");
            }
            resourceInputs["auditLogConfigs"] = args ? args.auditLogConfigs : undefined;
            resourceInputs["orgId"] = args ? args.orgId : undefined;
            resourceInputs["service"] = args ? args.service : undefined;
            resourceInputs["etag"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IamAuditConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IamAuditConfig resources.
 */
export interface IamAuditConfigState {
    /**
     * The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
     */
    auditLogConfigs?: pulumi.Input<pulumi.Input<inputs.organizations.IamAuditConfigAuditLogConfig>[]>;
    /**
     * (Computed) The etag of the organization's IAM policy.
     */
    etag?: pulumi.Input<string>;
    /**
     * The organization ID. If not specified for `gcp.organizations.IAMBinding`, `gcp.organizations.IAMMember`, or `gcp.organizations.IamAuditConfig`, uses the ID of the organization configured with the provider.
     * Required for `gcp.organizations.IAMPolicy` - you must explicitly set the organization, and it
     * will not be inferred from the provider.
     */
    orgId?: pulumi.Input<string>;
    /**
     * Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_organization\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `logTypes` specified in each `auditLogConfig` are enabled, and the `exemptedMembers` in each `auditLogConfig` are exempted.
     */
    service?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IamAuditConfig resource.
 */
export interface IamAuditConfigArgs {
    /**
     * The configuration for logging of each type of permission.  This can be specified multiple times.  Structure is documented below.
     */
    auditLogConfigs: pulumi.Input<pulumi.Input<inputs.organizations.IamAuditConfigAuditLogConfig>[]>;
    /**
     * The organization ID. If not specified for `gcp.organizations.IAMBinding`, `gcp.organizations.IAMMember`, or `gcp.organizations.IamAuditConfig`, uses the ID of the organization configured with the provider.
     * Required for `gcp.organizations.IAMPolicy` - you must explicitly set the organization, and it
     * will not be inferred from the provider.
     */
    orgId: pulumi.Input<string>;
    /**
     * Service which will be enabled for audit logging.  The special value `allServices` covers all services.  Note that if there are google\_organization\_iam\_audit\_config resources covering both `allServices` and a specific service then the union of the two AuditConfigs is used for that service: the `logTypes` specified in each `auditLogConfig` are enabled, and the `exemptedMembers` in each `auditLogConfig` are exempted.
     */
    service: pulumi.Input<string>;
}
