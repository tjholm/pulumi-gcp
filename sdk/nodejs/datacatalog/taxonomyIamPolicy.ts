// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for Data catalog Taxonomy. Each of these resources serves a different use case:
 *
 * * `gcp.datacatalog.TaxonomyIamPolicy`: Authoritative. Sets the IAM policy for the taxonomy and replaces any existing policy already attached.
 * * `gcp.datacatalog.TaxonomyIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the taxonomy are preserved.
 * * `gcp.datacatalog.TaxonomyIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the taxonomy are preserved.
 *
 * A data source can be used to retrieve policy data in advent you do not need creation
 *
 * * `gcp.datacatalog.TaxonomyIamPolicy`: Retrieves the IAM policy for the taxonomy
 *
 * > **Note:** `gcp.datacatalog.TaxonomyIamPolicy` **cannot** be used in conjunction with `gcp.datacatalog.TaxonomyIamBinding` and `gcp.datacatalog.TaxonomyIamMember` or they will fight over what your policy should be.
 *
 * > **Note:** `gcp.datacatalog.TaxonomyIamBinding` resources **can be** used in conjunction with `gcp.datacatalog.TaxonomyIamMember` resources **only if** they do not grant privilege to the same role.
 *
 * ## gcp.datacatalog.TaxonomyIamPolicy
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const admin = gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         role: "roles/viewer",
 *         members: ["user:jane@example.com"],
 *     }],
 * });
 * const policy = new gcp.datacatalog.TaxonomyIamPolicy("policy", {
 *     taxonomy: basicTaxonomy.name,
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 *
 * ## gcp.datacatalog.TaxonomyIamBinding
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const binding = new gcp.datacatalog.TaxonomyIamBinding("binding", {
 *     taxonomy: basicTaxonomy.name,
 *     role: "roles/viewer",
 *     members: ["user:jane@example.com"],
 * });
 * ```
 *
 * ## gcp.datacatalog.TaxonomyIamMember
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const member = new gcp.datacatalog.TaxonomyIamMember("member", {
 *     taxonomy: basicTaxonomy.name,
 *     role: "roles/viewer",
 *     member: "user:jane@example.com",
 * });
 * ```
 *
 * ## This resource supports User Project Overrides.
 *
 * - 
 *
 * # IAM policy for Data catalog Taxonomy
 * Three different resources help you manage your IAM policy for Data catalog Taxonomy. Each of these resources serves a different use case:
 *
 * * `gcp.datacatalog.TaxonomyIamPolicy`: Authoritative. Sets the IAM policy for the taxonomy and replaces any existing policy already attached.
 * * `gcp.datacatalog.TaxonomyIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the taxonomy are preserved.
 * * `gcp.datacatalog.TaxonomyIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the taxonomy are preserved.
 *
 * A data source can be used to retrieve policy data in advent you do not need creation
 *
 * * `gcp.datacatalog.TaxonomyIamPolicy`: Retrieves the IAM policy for the taxonomy
 *
 * > **Note:** `gcp.datacatalog.TaxonomyIamPolicy` **cannot** be used in conjunction with `gcp.datacatalog.TaxonomyIamBinding` and `gcp.datacatalog.TaxonomyIamMember` or they will fight over what your policy should be.
 *
 * > **Note:** `gcp.datacatalog.TaxonomyIamBinding` resources **can be** used in conjunction with `gcp.datacatalog.TaxonomyIamMember` resources **only if** they do not grant privilege to the same role.
 *
 * ## gcp.datacatalog.TaxonomyIamPolicy
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const admin = gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         role: "roles/viewer",
 *         members: ["user:jane@example.com"],
 *     }],
 * });
 * const policy = new gcp.datacatalog.TaxonomyIamPolicy("policy", {
 *     taxonomy: basicTaxonomy.name,
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 *
 * ## gcp.datacatalog.TaxonomyIamBinding
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const binding = new gcp.datacatalog.TaxonomyIamBinding("binding", {
 *     taxonomy: basicTaxonomy.name,
 *     role: "roles/viewer",
 *     members: ["user:jane@example.com"],
 * });
 * ```
 *
 * ## gcp.datacatalog.TaxonomyIamMember
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const member = new gcp.datacatalog.TaxonomyIamMember("member", {
 *     taxonomy: basicTaxonomy.name,
 *     role: "roles/viewer",
 *     member: "user:jane@example.com",
 * });
 * ```
 *
 * ## Import
 *
 * For all import syntaxes, the "resource in question" can take any of the following forms:
 *
 * * projects/{{project}}/locations/{{region}}/taxonomies/{{taxonomy}}
 *
 * * {{project}}/{{region}}/{{taxonomy}}
 *
 * * {{region}}/{{taxonomy}}
 *
 * * {{taxonomy}}
 *
 * Any variables not passed in the import command will be taken from the provider configuration.
 *
 * Data catalog taxonomy IAM resources can be imported using the resource identifiers, role, and member.
 *
 * IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
 *
 * ```sh
 * $ pulumi import gcp:datacatalog/taxonomyIamPolicy:TaxonomyIamPolicy editor "projects/{{project}}/locations/{{region}}/taxonomies/{{taxonomy}} roles/viewer user:jane@example.com"
 * ```
 *
 * IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
 *
 * ```sh
 * $ pulumi import gcp:datacatalog/taxonomyIamPolicy:TaxonomyIamPolicy editor "projects/{{project}}/locations/{{region}}/taxonomies/{{taxonomy}} roles/viewer"
 * ```
 *
 * IAM policy imports use the identifier of the resource in question, e.g.
 *
 * ```sh
 * $ pulumi import gcp:datacatalog/taxonomyIamPolicy:TaxonomyIamPolicy editor projects/{{project}}/locations/{{region}}/taxonomies/{{taxonomy}}
 * ```
 *
 * -> **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
 *
 *  full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 */
export class TaxonomyIamPolicy extends pulumi.CustomResource {
    /**
     * Get an existing TaxonomyIamPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TaxonomyIamPolicyState, opts?: pulumi.CustomResourceOptions): TaxonomyIamPolicy {
        return new TaxonomyIamPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:datacatalog/taxonomyIamPolicy:TaxonomyIamPolicy';

    /**
     * Returns true if the given object is an instance of TaxonomyIamPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TaxonomyIamPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TaxonomyIamPolicy.__pulumiType;
    }

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
    public readonly region!: pulumi.Output<string>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    public readonly taxonomy!: pulumi.Output<string>;

    /**
     * Create a TaxonomyIamPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TaxonomyIamPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TaxonomyIamPolicyArgs | TaxonomyIamPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TaxonomyIamPolicyState | undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["policyData"] = state ? state.policyData : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["taxonomy"] = state ? state.taxonomy : undefined;
        } else {
            const args = argsOrState as TaxonomyIamPolicyArgs | undefined;
            if ((!args || args.policyData === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyData'");
            }
            if ((!args || args.taxonomy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'taxonomy'");
            }
            resourceInputs["policyData"] = args ? args.policyData : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["taxonomy"] = args ? args.taxonomy : undefined;
            resourceInputs["etag"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TaxonomyIamPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TaxonomyIamPolicy resources.
 */
export interface TaxonomyIamPolicyState {
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
    region?: pulumi.Input<string>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    taxonomy?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a TaxonomyIamPolicy resource.
 */
export interface TaxonomyIamPolicyArgs {
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
    region?: pulumi.Input<string>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    taxonomy: pulumi.Input<string>;
}
