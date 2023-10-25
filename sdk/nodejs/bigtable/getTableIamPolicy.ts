// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Retrieves the current IAM policy data for a Bigtable Table.
 *
 * ## example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const policy = gcp.bigtable.getTableIamPolicy({
 *     instance: google_bigtable_instance.instance.name,
 *     table: google_bigtable_table.table.name,
 * });
 * ```
 */
export function getTableIamPolicy(args: GetTableIamPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetTableIamPolicyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gcp:bigtable/getTableIamPolicy:getTableIamPolicy", {
        "instance": args.instance,
        "project": args.project,
        "table": args.table,
    }, opts);
}

/**
 * A collection of arguments for invoking getTableIamPolicy.
 */
export interface GetTableIamPolicyArgs {
    /**
     * The name or relative resource id of the instance that owns the table.
     */
    instance: string;
    project?: string;
    /**
     * The name or relative resource id of the table to manage IAM policies for.
     */
    table: string;
}

/**
 * A collection of values returned by getTableIamPolicy.
 */
export interface GetTableIamPolicyResult {
    /**
     * (Computed) The etag of the IAM policy.
     */
    readonly etag: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instance: string;
    /**
     * (Computed) The policy data
     */
    readonly policyData: string;
    readonly project: string;
    readonly table: string;
}
/**
 * Retrieves the current IAM policy data for a Bigtable Table.
 *
 * ## example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const policy = gcp.bigtable.getTableIamPolicy({
 *     instance: google_bigtable_instance.instance.name,
 *     table: google_bigtable_table.table.name,
 * });
 * ```
 */
export function getTableIamPolicyOutput(args: GetTableIamPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTableIamPolicyResult> {
    return pulumi.output(args).apply((a: any) => getTableIamPolicy(a, opts))
}

/**
 * A collection of arguments for invoking getTableIamPolicy.
 */
export interface GetTableIamPolicyOutputArgs {
    /**
     * The name or relative resource id of the instance that owns the table.
     */
    instance: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * The name or relative resource id of the table to manage IAM policies for.
     */
    table: pulumi.Input<string>;
}
