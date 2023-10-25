// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Retrieves the current IAM policy data for function
 *
 * ## example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const policy = gcp.cloudfunctionsv2.getFunctionIamPolicy({
 *     project: google_cloudfunctions2_function["function"].project,
 *     location: google_cloudfunctions2_function["function"].location,
 *     cloudFunction: google_cloudfunctions2_function["function"].name,
 * });
 * ```
 */
export function getFunctionIamPolicy(args: GetFunctionIamPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetFunctionIamPolicyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gcp:cloudfunctionsv2/getFunctionIamPolicy:getFunctionIamPolicy", {
        "cloudFunction": args.cloudFunction,
        "location": args.location,
        "project": args.project,
    }, opts);
}

/**
 * A collection of arguments for invoking getFunctionIamPolicy.
 */
export interface GetFunctionIamPolicyArgs {
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    cloudFunction: string;
    /**
     * The location of this cloud function. Used to find the parent resource to bind the IAM policy to
     */
    location?: string;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    project?: string;
}

/**
 * A collection of values returned by getFunctionIamPolicy.
 */
export interface GetFunctionIamPolicyResult {
    readonly cloudFunction: string;
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
     * (Required only by `gcp.cloudfunctionsv2.FunctionIamPolicy`) The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     */
    readonly policyData: string;
    readonly project: string;
}
/**
 * Retrieves the current IAM policy data for function
 *
 * ## example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const policy = gcp.cloudfunctionsv2.getFunctionIamPolicy({
 *     project: google_cloudfunctions2_function["function"].project,
 *     location: google_cloudfunctions2_function["function"].location,
 *     cloudFunction: google_cloudfunctions2_function["function"].name,
 * });
 * ```
 */
export function getFunctionIamPolicyOutput(args: GetFunctionIamPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFunctionIamPolicyResult> {
    return pulumi.output(args).apply((a: any) => getFunctionIamPolicy(a, opts))
}

/**
 * A collection of arguments for invoking getFunctionIamPolicy.
 */
export interface GetFunctionIamPolicyOutputArgs {
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    cloudFunction: pulumi.Input<string>;
    /**
     * The location of this cloud function. Used to find the parent resource to bind the IAM policy to
     */
    location?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    project?: pulumi.Input<string>;
}
