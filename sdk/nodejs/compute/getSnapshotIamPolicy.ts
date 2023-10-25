// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Retrieves the current IAM policy data for snapshot
 *
 * ## example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const policy = gcp.compute.getSnapshotIamPolicy({
 *     project: google_compute_snapshot.snapshot.project,
 *     name: google_compute_snapshot.snapshot.name,
 * });
 * ```
 */
export function getSnapshotIamPolicy(args: GetSnapshotIamPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetSnapshotIamPolicyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gcp:compute/getSnapshotIamPolicy:getSnapshotIamPolicy", {
        "name": args.name,
        "project": args.project,
    }, opts);
}

/**
 * A collection of arguments for invoking getSnapshotIamPolicy.
 */
export interface GetSnapshotIamPolicyArgs {
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    name: string;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    project?: string;
}

/**
 * A collection of values returned by getSnapshotIamPolicy.
 */
export interface GetSnapshotIamPolicyResult {
    /**
     * (Computed) The etag of the IAM policy.
     */
    readonly etag: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    /**
     * (Required only by `gcp.compute.SnapshotIamPolicy`) The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     */
    readonly policyData: string;
    readonly project: string;
}
/**
 * Retrieves the current IAM policy data for snapshot
 *
 * ## example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const policy = gcp.compute.getSnapshotIamPolicy({
 *     project: google_compute_snapshot.snapshot.project,
 *     name: google_compute_snapshot.snapshot.name,
 * });
 * ```
 */
export function getSnapshotIamPolicyOutput(args: GetSnapshotIamPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSnapshotIamPolicyResult> {
    return pulumi.output(args).apply((a: any) => getSnapshotIamPolicy(a, opts))
}

/**
 * A collection of arguments for invoking getSnapshotIamPolicy.
 */
export interface GetSnapshotIamPolicyOutputArgs {
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    name: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    project?: pulumi.Input<string>;
}
