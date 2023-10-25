// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Retrieves the current IAM policy data for bucket
 *
 * ## example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const policy = gcp.storage.getBucketIamPolicy({
 *     bucket: google_storage_bucket["default"].name,
 * });
 * ```
 */
export function getBucketIamPolicy(args: GetBucketIamPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetBucketIamPolicyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gcp:storage/getBucketIamPolicy:getBucketIamPolicy", {
        "bucket": args.bucket,
    }, opts);
}

/**
 * A collection of arguments for invoking getBucketIamPolicy.
 */
export interface GetBucketIamPolicyArgs {
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    bucket: string;
}

/**
 * A collection of values returned by getBucketIamPolicy.
 */
export interface GetBucketIamPolicyResult {
    readonly bucket: string;
    /**
     * (Computed) The etag of the IAM policy.
     */
    readonly etag: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * (Required only by `gcp.storage.BucketIAMPolicy`) The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     */
    readonly policyData: string;
}
/**
 * Retrieves the current IAM policy data for bucket
 *
 * ## example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const policy = gcp.storage.getBucketIamPolicy({
 *     bucket: google_storage_bucket["default"].name,
 * });
 * ```
 */
export function getBucketIamPolicyOutput(args: GetBucketIamPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBucketIamPolicyResult> {
    return pulumi.output(args).apply((a: any) => getBucketIamPolicy(a, opts))
}

/**
 * A collection of arguments for invoking getBucketIamPolicy.
 */
export interface GetBucketIamPolicyOutputArgs {
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    bucket: pulumi.Input<string>;
}
