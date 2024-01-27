// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Tenant configuration in a multi-tenant project.
 *
 * You must enable the
 * [Google Identity Platform](https://console.cloud.google.com/marketplace/details/google-cloud-platform/customer-identity) in
 * the marketplace prior to using this resource.
 *
 * You must [enable multi-tenancy](https://cloud.google.com/identity-platform/docs/multi-tenancy-quickstart) via
 * the Cloud Console prior to creating tenants.
 *
 * ## Example Usage
 * ### Identity Platform Tenant Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const tenant = new gcp.identityplatform.Tenant("tenant", {
 *     allowPasswordSignup: true,
 *     displayName: "tenant",
 * });
 * ```
 *
 * ## Import
 *
 * Tenant can be imported using any of these accepted formats* `projects/{{project}}/tenants/{{name}}` * `{{project}}/{{name}}` * `{{name}}` When using the `pulumi import` command, Tenant can be imported using one of the formats above. For example
 *
 * ```sh
 *  $ pulumi import gcp:identityplatform/tenant:Tenant default projects/{{project}}/tenants/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:identityplatform/tenant:Tenant default {{project}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:identityplatform/tenant:Tenant default {{name}}
 * ```
 */
export class Tenant extends pulumi.CustomResource {
    /**
     * Get an existing Tenant resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TenantState, opts?: pulumi.CustomResourceOptions): Tenant {
        return new Tenant(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:identityplatform/tenant:Tenant';

    /**
     * Returns true if the given object is an instance of Tenant.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Tenant {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Tenant.__pulumiType;
    }

    /**
     * Whether to allow email/password user authentication.
     */
    public readonly allowPasswordSignup!: pulumi.Output<boolean | undefined>;
    /**
     * Whether authentication is disabled for the tenant. If true, the users under
     * the disabled tenant are not allowed to sign-in. Admins of the disabled tenant
     * are not able to manage its users.
     */
    public readonly disableAuth!: pulumi.Output<boolean | undefined>;
    /**
     * Human friendly display name of the tenant.
     *
     *
     * - - -
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Whether to enable email link user authentication.
     */
    public readonly enableEmailLinkSignin!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the tenant that is generated by the server
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a Tenant resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TenantArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TenantArgs | TenantState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TenantState | undefined;
            resourceInputs["allowPasswordSignup"] = state ? state.allowPasswordSignup : undefined;
            resourceInputs["disableAuth"] = state ? state.disableAuth : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["enableEmailLinkSignin"] = state ? state.enableEmailLinkSignin : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as TenantArgs | undefined;
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            resourceInputs["allowPasswordSignup"] = args ? args.allowPasswordSignup : undefined;
            resourceInputs["disableAuth"] = args ? args.disableAuth : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["enableEmailLinkSignin"] = args ? args.enableEmailLinkSignin : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["name"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Tenant.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Tenant resources.
 */
export interface TenantState {
    /**
     * Whether to allow email/password user authentication.
     */
    allowPasswordSignup?: pulumi.Input<boolean>;
    /**
     * Whether authentication is disabled for the tenant. If true, the users under
     * the disabled tenant are not allowed to sign-in. Admins of the disabled tenant
     * are not able to manage its users.
     */
    disableAuth?: pulumi.Input<boolean>;
    /**
     * Human friendly display name of the tenant.
     *
     *
     * - - -
     */
    displayName?: pulumi.Input<string>;
    /**
     * Whether to enable email link user authentication.
     */
    enableEmailLinkSignin?: pulumi.Input<boolean>;
    /**
     * The name of the tenant that is generated by the server
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Tenant resource.
 */
export interface TenantArgs {
    /**
     * Whether to allow email/password user authentication.
     */
    allowPasswordSignup?: pulumi.Input<boolean>;
    /**
     * Whether authentication is disabled for the tenant. If true, the users under
     * the disabled tenant are not allowed to sign-in. Admins of the disabled tenant
     * are not able to manage its users.
     */
    disableAuth?: pulumi.Input<boolean>;
    /**
     * Human friendly display name of the tenant.
     *
     *
     * - - -
     */
    displayName: pulumi.Input<string>;
    /**
     * Whether to enable email link user authentication.
     */
    enableEmailLinkSignin?: pulumi.Input<boolean>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
}
