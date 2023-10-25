// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Retrieves the current IAM policy data for secret
 *
 * ## example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const policy = gcp.secretmanager.getSecretIamPolicy({
 *     project: google_secret_manager_secret["secret-basic"].project,
 *     secretId: google_secret_manager_secret["secret-basic"].secret_id,
 * });
 * ```
 */
export function getSecretIamPolicy(args: GetSecretIamPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetSecretIamPolicyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gcp:secretmanager/getSecretIamPolicy:getSecretIamPolicy", {
        "project": args.project,
        "secretId": args.secretId,
    }, opts);
}

/**
 * A collection of arguments for invoking getSecretIamPolicy.
 */
export interface GetSecretIamPolicyArgs {
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    project?: string;
    secretId: string;
}

/**
 * A collection of values returned by getSecretIamPolicy.
 */
export interface GetSecretIamPolicyResult {
    /**
     * (Computed) The etag of the IAM policy.
     */
    readonly etag: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * (Required only by `gcp.secretmanager.SecretIamPolicy`) The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     */
    readonly policyData: string;
    readonly project: string;
    readonly secretId: string;
}
/**
 * Retrieves the current IAM policy data for secret
 *
 * ## example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const policy = gcp.secretmanager.getSecretIamPolicy({
 *     project: google_secret_manager_secret["secret-basic"].project,
 *     secretId: google_secret_manager_secret["secret-basic"].secret_id,
 * });
 * ```
 */
export function getSecretIamPolicyOutput(args: GetSecretIamPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSecretIamPolicyResult> {
    return pulumi.output(args).apply((a: any) => getSecretIamPolicy(a, opts))
}

/**
 * A collection of arguments for invoking getSecretIamPolicy.
 */
export interface GetSecretIamPolicyOutputArgs {
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    project?: pulumi.Input<string>;
    secretId: pulumi.Input<string>;
}
