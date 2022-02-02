// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for API Gateway ApiConfig. Each of these resources serves a different use case:
 *
 * * `gcp.apigateway.ApiConfigIamPolicy`: Authoritative. Sets the IAM policy for the apiconfig and replaces any existing policy already attached.
 * * `gcp.apigateway.ApiConfigIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the apiconfig are preserved.
 * * `gcp.apigateway.ApiConfigIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the apiconfig are preserved.
 *
 * > **Note:** `gcp.apigateway.ApiConfigIamPolicy` **cannot** be used in conjunction with `gcp.apigateway.ApiConfigIamBinding` and `gcp.apigateway.ApiConfigIamMember` or they will fight over what your policy should be.
 *
 * > **Note:** `gcp.apigateway.ApiConfigIamBinding` resources **can be** used in conjunction with `gcp.apigateway.ApiConfigIamMember` resources **only if** they do not grant privilege to the same role.
 *
 * ## google\_api\_gateway\_api\_config\_iam\_policy
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
 * const policy = new gcp.apigateway.ApiConfigIamPolicy("policy", {
 *     api: google_api_gateway_api_config.api_cfg.api,
 *     apiConfig: google_api_gateway_api_config.api_cfg.api_config_id,
 *     policyData: admin.then(admin => admin.policyData),
 * }, {
 *     provider: google_beta,
 * });
 * ```
 *
 * ## google\_api\_gateway\_api\_config\_iam\_binding
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const binding = new gcp.apigateway.ApiConfigIamBinding("binding", {
 *     api: google_api_gateway_api_config.api_cfg.api,
 *     apiConfig: google_api_gateway_api_config.api_cfg.api_config_id,
 *     role: "roles/apigateway.viewer",
 *     members: ["user:jane@example.com"],
 * }, {
 *     provider: google_beta,
 * });
 * ```
 *
 * ## google\_api\_gateway\_api\_config\_iam\_member
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const member = new gcp.apigateway.ApiConfigIamMember("member", {
 *     api: google_api_gateway_api_config.api_cfg.api,
 *     apiConfig: google_api_gateway_api_config.api_cfg.api_config_id,
 *     role: "roles/apigateway.viewer",
 *     member: "user:jane@example.com",
 * }, {
 *     provider: google_beta,
 * });
 * ```
 *
 * ## Import
 *
 * For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/global/apis/{{api}}/configs/{{api_config}} * {{project}}/{{api}}/{{api_config}} * {{api}}/{{api_config}} * {{api_config}} Any variables not passed in the import command will be taken from the provider configuration. API Gateway apiconfig IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:apigateway/apiConfigIamPolicy:ApiConfigIamPolicy editor "projects/{{project}}/locations/global/apis/{{api}}/configs/{{api_config}} roles/apigateway.viewer user:jane@example.com"
 * ```
 *
 *  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:apigateway/apiConfigIamPolicy:ApiConfigIamPolicy editor "projects/{{project}}/locations/global/apis/{{api}}/configs/{{api_config}} roles/apigateway.viewer"
 * ```
 *
 *  IAM policy imports use the identifier of the resource in question, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:apigateway/apiConfigIamPolicy:ApiConfigIamPolicy editor projects/{{project}}/locations/global/apis/{{api}}/configs/{{api_config}}
 * ```
 *
 *  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
 *
 * full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
 */
export class ApiConfigIamPolicy extends pulumi.CustomResource {
    /**
     * Get an existing ApiConfigIamPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ApiConfigIamPolicyState, opts?: pulumi.CustomResourceOptions): ApiConfigIamPolicy {
        return new ApiConfigIamPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:apigateway/apiConfigIamPolicy:ApiConfigIamPolicy';

    /**
     * Returns true if the given object is an instance of ApiConfigIamPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ApiConfigIamPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ApiConfigIamPolicy.__pulumiType;
    }

    /**
     * The API to attach the config to.
     * Used to find the parent resource to bind the IAM policy to
     */
    public readonly api!: pulumi.Output<string>;
    public readonly apiConfig!: pulumi.Output<string>;
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
     * Create a ApiConfigIamPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ApiConfigIamPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ApiConfigIamPolicyArgs | ApiConfigIamPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ApiConfigIamPolicyState | undefined;
            resourceInputs["api"] = state ? state.api : undefined;
            resourceInputs["apiConfig"] = state ? state.apiConfig : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["policyData"] = state ? state.policyData : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as ApiConfigIamPolicyArgs | undefined;
            if ((!args || args.api === undefined) && !opts.urn) {
                throw new Error("Missing required property 'api'");
            }
            if ((!args || args.apiConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'apiConfig'");
            }
            if ((!args || args.policyData === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyData'");
            }
            resourceInputs["api"] = args ? args.api : undefined;
            resourceInputs["apiConfig"] = args ? args.apiConfig : undefined;
            resourceInputs["policyData"] = args ? args.policyData : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["etag"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ApiConfigIamPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ApiConfigIamPolicy resources.
 */
export interface ApiConfigIamPolicyState {
    /**
     * The API to attach the config to.
     * Used to find the parent resource to bind the IAM policy to
     */
    api?: pulumi.Input<string>;
    apiConfig?: pulumi.Input<string>;
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
 * The set of arguments for constructing a ApiConfigIamPolicy resource.
 */
export interface ApiConfigIamPolicyArgs {
    /**
     * The API to attach the config to.
     * Used to find the parent resource to bind the IAM policy to
     */
    api: pulumi.Input<string>;
    apiConfig: pulumi.Input<string>;
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
