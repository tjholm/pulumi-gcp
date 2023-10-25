// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Retrieves the current IAM policy data for tunnel
 *
 * ## example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const policy = gcp.iap.getTunnelIamPolicy({
 *     project: google_project_service.project_service.project,
 * });
 * ```
 */
export function getTunnelIamPolicy(args?: GetTunnelIamPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetTunnelIamPolicyResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gcp:iap/getTunnelIamPolicy:getTunnelIamPolicy", {
        "project": args.project,
    }, opts);
}

/**
 * A collection of arguments for invoking getTunnelIamPolicy.
 */
export interface GetTunnelIamPolicyArgs {
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    project?: string;
}

/**
 * A collection of values returned by getTunnelIamPolicy.
 */
export interface GetTunnelIamPolicyResult {
    /**
     * (Computed) The etag of the IAM policy.
     */
    readonly etag: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * (Required only by `gcp.iap.TunnelIamPolicy`) The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     */
    readonly policyData: string;
    readonly project: string;
}
/**
 * Retrieves the current IAM policy data for tunnel
 *
 * ## example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const policy = gcp.iap.getTunnelIamPolicy({
 *     project: google_project_service.project_service.project,
 * });
 * ```
 */
export function getTunnelIamPolicyOutput(args?: GetTunnelIamPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTunnelIamPolicyResult> {
    return pulumi.output(args).apply((a: any) => getTunnelIamPolicy(a, opts))
}

/**
 * A collection of arguments for invoking getTunnelIamPolicy.
 */
export interface GetTunnelIamPolicyOutputArgs {
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
     */
    project?: pulumi.Input<string>;
}
