// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Retrieves the current IAM policy data for entrygroup
 *
 * ## example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const policy = gcp.datacatalog.getEntryGroupIamPolicy({
 *     entryGroup: google_data_catalog_entry_group.basic_entry_group.name,
 * });
 * ```
 */
export function getEntryGroupIamPolicy(args: GetEntryGroupIamPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetEntryGroupIamPolicyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gcp:datacatalog/getEntryGroupIamPolicy:getEntryGroupIamPolicy", {
        "entryGroup": args.entryGroup,
        "project": args.project,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getEntryGroupIamPolicy.
 */
export interface GetEntryGroupIamPolicyArgs {
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    entryGroup: string;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    project?: string;
    region?: string;
}

/**
 * A collection of values returned by getEntryGroupIamPolicy.
 */
export interface GetEntryGroupIamPolicyResult {
    readonly entryGroup: string;
    /**
     * (Computed) The etag of the IAM policy.
     */
    readonly etag: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * (Required only by `gcp.datacatalog.EntryGroupIamPolicy`) The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     */
    readonly policyData: string;
    readonly project: string;
    readonly region: string;
}
/**
 * Retrieves the current IAM policy data for entrygroup
 *
 * ## example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const policy = gcp.datacatalog.getEntryGroupIamPolicy({
 *     entryGroup: google_data_catalog_entry_group.basic_entry_group.name,
 * });
 * ```
 */
export function getEntryGroupIamPolicyOutput(args: GetEntryGroupIamPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEntryGroupIamPolicyResult> {
    return pulumi.output(args).apply((a: any) => getEntryGroupIamPolicy(a, opts))
}

/**
 * A collection of arguments for invoking getEntryGroupIamPolicy.
 */
export interface GetEntryGroupIamPolicyOutputArgs {
    /**
     * Used to find the parent resource to bind the IAM policy to
     */
    entryGroup: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    project?: pulumi.Input<string>;
    region?: pulumi.Input<string>;
}
