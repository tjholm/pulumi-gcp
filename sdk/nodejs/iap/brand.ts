// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 * ### Iap Brand
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const project = new gcp.organizations.Project("project", {orgId: "123456789"});
 * const projectService = new gcp.projects.Service("projectService", {
 *     project: project.projectId,
 *     service: "iap.googleapis.com",
 * });
 * const projectBrand = new gcp.iap.Brand("projectBrand", {
 *     supportEmail: "support@example.com",
 *     applicationTitle: "Cloud IAP protected Application",
 *     project: projectService.project,
 * });
 * ```
 *
 * ## Import
 *
 * Brand can be imported using any of these accepted formats* `projects/{{project_id}}/brands/{{brand_id}}` * `projects/{{project_number}}/brands/{{brand_id}}` * `{{project_number}}/{{brand_id}}` When using the `pulumi import` command, Brand can be imported using one of the formats above. For example
 *
 * ```sh
 *  $ pulumi import gcp:iap/brand:Brand default projects/{{project_id}}/brands/{{brand_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:iap/brand:Brand default projects/{{project_number}}/brands/{{brand_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:iap/brand:Brand default {{project_number}}/{{brand_id}}
 * ```
 */
export class Brand extends pulumi.CustomResource {
    /**
     * Get an existing Brand resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BrandState, opts?: pulumi.CustomResourceOptions): Brand {
        return new Brand(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:iap/brand:Brand';

    /**
     * Returns true if the given object is an instance of Brand.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Brand {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Brand.__pulumiType;
    }

    /**
     * Application name displayed on OAuth consent screen.
     *
     *
     * - - -
     */
    public readonly applicationTitle!: pulumi.Output<string>;
    /**
     * Output only. Identifier of the brand, in the format `projects/{project_number}/brands/{brand_id}`
     * NOTE: The name can also be expressed as `projects/{project_id}/brands/{brand_id}`, e.g. when importing.
     * NOTE: The brand identification corresponds to the project number as only one
     * brand can be created per project.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Whether the brand is only intended for usage inside the GSuite organization only.
     */
    public /*out*/ readonly orgInternalOnly!: pulumi.Output<boolean>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Support email displayed on the OAuth consent screen. Can be either a
     * user or group email. When a user email is specified, the caller must
     * be the user with the associated email address. When a group email is
     * specified, the caller can be either a user or a service account which
     * is an owner of the specified group in Cloud Identity.
     */
    public readonly supportEmail!: pulumi.Output<string>;

    /**
     * Create a Brand resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BrandArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BrandArgs | BrandState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BrandState | undefined;
            resourceInputs["applicationTitle"] = state ? state.applicationTitle : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["orgInternalOnly"] = state ? state.orgInternalOnly : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["supportEmail"] = state ? state.supportEmail : undefined;
        } else {
            const args = argsOrState as BrandArgs | undefined;
            if ((!args || args.applicationTitle === undefined) && !opts.urn) {
                throw new Error("Missing required property 'applicationTitle'");
            }
            if ((!args || args.supportEmail === undefined) && !opts.urn) {
                throw new Error("Missing required property 'supportEmail'");
            }
            resourceInputs["applicationTitle"] = args ? args.applicationTitle : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["supportEmail"] = args ? args.supportEmail : undefined;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["orgInternalOnly"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Brand.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Brand resources.
 */
export interface BrandState {
    /**
     * Application name displayed on OAuth consent screen.
     *
     *
     * - - -
     */
    applicationTitle?: pulumi.Input<string>;
    /**
     * Output only. Identifier of the brand, in the format `projects/{project_number}/brands/{brand_id}`
     * NOTE: The name can also be expressed as `projects/{project_id}/brands/{brand_id}`, e.g. when importing.
     * NOTE: The brand identification corresponds to the project number as only one
     * brand can be created per project.
     */
    name?: pulumi.Input<string>;
    /**
     * Whether the brand is only intended for usage inside the GSuite organization only.
     */
    orgInternalOnly?: pulumi.Input<boolean>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Support email displayed on the OAuth consent screen. Can be either a
     * user or group email. When a user email is specified, the caller must
     * be the user with the associated email address. When a group email is
     * specified, the caller can be either a user or a service account which
     * is an owner of the specified group in Cloud Identity.
     */
    supportEmail?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Brand resource.
 */
export interface BrandArgs {
    /**
     * Application name displayed on OAuth consent screen.
     *
     *
     * - - -
     */
    applicationTitle: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Support email displayed on the OAuth consent screen. Can be either a
     * user or group email. When a user email is specified, the caller must
     * be the user with the associated email address. When a group email is
     * specified, the caller can be either a user or a service account which
     * is an owner of the specified group in Cloud Identity.
     */
    supportEmail: pulumi.Input<string>;
}
