// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for API Gateway Api. Each of these resources serves a different use case:
 *
 * * `gcp.apigateway.ApiIamPolicy`: Authoritative. Sets the IAM policy for the api and replaces any existing policy already attached.
 * * `gcp.apigateway.ApiIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the api are preserved.
 * * `gcp.apigateway.ApiIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the api are preserved.
 *
 * > **Note:** `gcp.apigateway.ApiIamPolicy` **cannot** be used in conjunction with `gcp.apigateway.ApiIamBinding` and `gcp.apigateway.ApiIamMember` or they will fight over what your policy should be.
 *
 * > **Note:** `gcp.apigateway.ApiIamBinding` resources **can be** used in conjunction with `gcp.apigateway.ApiIamMember` resources **only if** they do not grant privilege to the same role.
 *
 * ## google\_api\_gateway\_api\_iam\_policy
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const admin = gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         role: "roles/apigateway.viewer",
 *         members: ["user:jane@example.com"],
 *     }],
 * });
 * const policy = new gcp.apigateway.ApiIamPolicy("policy", {
 *     project: google_api_gateway_api.api.project,
 *     api: google_api_gateway_api.api.api_id,
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 *
 * ## google\_api\_gateway\_api\_iam\_binding
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const binding = new gcp.apigateway.ApiIamBinding("binding", {
 *     project: google_api_gateway_api.api.project,
 *     api: google_api_gateway_api.api.api_id,
 *     role: "roles/apigateway.viewer",
 *     members: ["user:jane@example.com"],
 * });
 * ```
 *
 * ## google\_api\_gateway\_api\_iam\_member
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const member = new gcp.apigateway.ApiIamMember("member", {
 *     project: google_api_gateway_api.api.project,
 *     api: google_api_gateway_api.api.api_id,
 *     role: "roles/apigateway.viewer",
 *     member: "user:jane@example.com",
 * });
 * ```
 *
 * ## Import
 *
 * For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/global/apis/{{api}} * {{project}}/{{api}} * {{api}} Any variables not passed in the import command will be taken from the provider configuration. API Gateway api IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:apigateway/apiIamPolicy:ApiIamPolicy editor "projects/{{project}}/locations/global/apis/{{api}} roles/apigateway.viewer user:jane@example.com"
 * ```
 *
 *  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:apigateway/apiIamPolicy:ApiIamPolicy editor "projects/{{project}}/locations/global/apis/{{api}} roles/apigateway.viewer"
 * ```
 *
 *  IAM policy imports use the identifier of the resource in question, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:apigateway/apiIamPolicy:ApiIamPolicy editor projects/{{project}}/locations/global/apis/{{api}}
 * ```
 *
 *  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
 *
 * full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 */
export class ApiIamPolicy extends pulumi.CustomResource {
    /**
     * Get an existing ApiIamPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApiIamPolicyState, opts?: pulumi.CustomResourceOptions): ApiIamPolicy {
        return new ApiIamPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:apigateway/apiIamPolicy:ApiIamPolicy';

    /**
     * Returns true if the given object is an instance of ApiIamPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApiIamPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApiIamPolicy.__pulumiType;
    }

    public readonly api!: pulumi.Output<string>;
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
     * Create a ApiIamPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApiIamPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApiIamPolicyArgs | ApiIamPolicyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApiIamPolicyState | undefined;
            inputs["api"] = state ? state.api : undefined;
            inputs["etag"] = state ? state.etag : undefined;
            inputs["policyData"] = state ? state.policyData : undefined;
            inputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as ApiIamPolicyArgs | undefined;
            if ((!args || args.api === undefined) && !opts.urn) {
                throw new Error("Missing required property 'api'");
            }
            if ((!args || args.policyData === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyData'");
            }
            inputs["api"] = args ? args.api : undefined;
            inputs["policyData"] = args ? args.policyData : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["etag"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ApiIamPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApiIamPolicy resources.
 */
export interface ApiIamPolicyState {
    api?: pulumi.Input<string>;
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
}

/**
 * The set of arguments for constructing a ApiIamPolicy resource.
 */
export interface ApiIamPolicyArgs {
    api: pulumi.Input<string>;
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
}
