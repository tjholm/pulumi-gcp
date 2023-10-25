// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Retrieves the current IAM policy data for runtime
 *
 * ## example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const policy = gcp.notebooks.getRuntimeIamPolicy({
 *     project: google_notebooks_runtime.runtime.project,
 *     location: google_notebooks_runtime.runtime.location,
 *     runtimeName: google_notebooks_runtime.runtime.name,
 * });
 * ```
 */
export function getRuntimeIamPolicy(args: GetRuntimeIamPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetRuntimeIamPolicyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gcp:notebooks/getRuntimeIamPolicy:getRuntimeIamPolicy", {
        "location": args.location,
        "project": args.project,
        "runtimeName": args.runtimeName,
    }, opts);
}

/**
 * A collection of arguments for invoking getRuntimeIamPolicy.
 */
export interface GetRuntimeIamPolicyArgs {
    /**
     * A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to
     */
    location?: string;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    project?: string;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    runtimeName: string;
}

/**
 * A collection of values returned by getRuntimeIamPolicy.
 */
export interface GetRuntimeIamPolicyResult {
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
     * (Required only by `gcp.notebooks.RuntimeIamPolicy`) The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     */
    readonly policyData: string;
    readonly project: string;
    readonly runtimeName: string;
}
/**
 * Retrieves the current IAM policy data for runtime
 *
 * ## example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const policy = gcp.notebooks.getRuntimeIamPolicy({
 *     project: google_notebooks_runtime.runtime.project,
 *     location: google_notebooks_runtime.runtime.location,
 *     runtimeName: google_notebooks_runtime.runtime.name,
 * });
 * ```
 */
export function getRuntimeIamPolicyOutput(args: GetRuntimeIamPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRuntimeIamPolicyResult> {
    return pulumi.output(args).apply((a: any) => getRuntimeIamPolicy(a, opts))
}

/**
 * A collection of arguments for invoking getRuntimeIamPolicy.
 */
export interface GetRuntimeIamPolicyOutputArgs {
    /**
     * A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to
     */
    location?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    runtimeName: pulumi.Input<string>;
}
