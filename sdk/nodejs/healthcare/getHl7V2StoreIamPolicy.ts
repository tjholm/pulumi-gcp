// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Retrieves the current IAM policy data for a Google Cloud Healthcare HL7v2 store.
 *
 * ## example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const foo = gcp.healthcare.getHl7V2StoreIamPolicy({
 *     hl7V2StoreId: google_healthcare_hl7_v2_store.hl7_v2_store.id,
 * });
 * ```
 */
export function getHl7V2StoreIamPolicy(args: GetHl7V2StoreIamPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetHl7V2StoreIamPolicyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gcp:healthcare/getHl7V2StoreIamPolicy:getHl7V2StoreIamPolicy", {
        "hl7V2StoreId": args.hl7V2StoreId,
    }, opts);
}

/**
 * A collection of arguments for invoking getHl7V2StoreIamPolicy.
 */
export interface GetHl7V2StoreIamPolicyArgs {
    /**
     * The HL7v2 store ID, in the form
     * `{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}` or
     * `{location_name}/{dataset_name}/{hl7_v2_store_name}`. In the second form, the provider's
     * project setting will be used as a fallback.
     */
    hl7V2StoreId: string;
}

/**
 * A collection of values returned by getHl7V2StoreIamPolicy.
 */
export interface GetHl7V2StoreIamPolicyResult {
    /**
     * (Computed) The etag of the IAM policy.
     */
    readonly etag: string;
    readonly hl7V2StoreId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * (Computed) The policy data
     */
    readonly policyData: string;
}
/**
 * Retrieves the current IAM policy data for a Google Cloud Healthcare HL7v2 store.
 *
 * ## example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const foo = gcp.healthcare.getHl7V2StoreIamPolicy({
 *     hl7V2StoreId: google_healthcare_hl7_v2_store.hl7_v2_store.id,
 * });
 * ```
 */
export function getHl7V2StoreIamPolicyOutput(args: GetHl7V2StoreIamPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetHl7V2StoreIamPolicyResult> {
    return pulumi.output(args).apply((a: any) => getHl7V2StoreIamPolicy(a, opts))
}

/**
 * A collection of arguments for invoking getHl7V2StoreIamPolicy.
 */
export interface GetHl7V2StoreIamPolicyOutputArgs {
    /**
     * The HL7v2 store ID, in the form
     * `{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}` or
     * `{location_name}/{dataset_name}/{hl7_v2_store_name}`. In the second form, the provider's
     * project setting will be used as a fallback.
     */
    hl7V2StoreId: pulumi.Input<string>;
}
