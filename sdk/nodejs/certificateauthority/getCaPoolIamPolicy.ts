// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Retrieves the current IAM policy data for capool
 *
 * ## example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const policy = gcp.certificateauthority.getCaPoolIamPolicy({
 *     caPool: google_privateca_ca_pool["default"].id,
 * });
 * ```
 */
export function getCaPoolIamPolicy(args: GetCaPoolIamPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetCaPoolIamPolicyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gcp:certificateauthority/getCaPoolIamPolicy:getCaPoolIamPolicy", {
        "caPool": args.caPool,
        "location": args.location,
        "project": args.project,
    }, opts);
}

/**
 * A collection of arguments for invoking getCaPoolIamPolicy.
 */
export interface GetCaPoolIamPolicyArgs {
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    caPool: string;
    /**
     * Location of the CaPool. A full list of valid locations can be found by
     * running `gcloud privateca locations list`.
     * Used to find the parent resource to bind the IAM policy to
     */
    location?: string;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    project?: string;
}

/**
 * A collection of values returned by getCaPoolIamPolicy.
 */
export interface GetCaPoolIamPolicyResult {
    readonly caPool: string;
    /**
     * (Computed) The etag of the IAM policy.
     */
    readonly etag: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly location: string;
    /**
     * (Required only by `gcp.certificateauthority.CaPoolIamPolicy`) The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     */
    readonly policyData: string;
    readonly project: string;
}
/**
 * Retrieves the current IAM policy data for capool
 *
 * ## example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const policy = gcp.certificateauthority.getCaPoolIamPolicy({
 *     caPool: google_privateca_ca_pool["default"].id,
 * });
 * ```
 */
export function getCaPoolIamPolicyOutput(args: GetCaPoolIamPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCaPoolIamPolicyResult> {
    return pulumi.output(args).apply((a: any) => getCaPoolIamPolicy(a, opts))
}

/**
 * A collection of arguments for invoking getCaPoolIamPolicy.
 */
export interface GetCaPoolIamPolicyOutputArgs {
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    caPool: pulumi.Input<string>;
    /**
     * Location of the CaPool. A full list of valid locations can be found by
     * running `gcloud privateca locations list`.
     * Used to find the parent resource to bind the IAM policy to
     */
    location?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    project?: pulumi.Input<string>;
}
