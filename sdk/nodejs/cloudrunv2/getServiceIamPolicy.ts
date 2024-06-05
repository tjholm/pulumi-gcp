// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Retrieves the current IAM policy data for service
 *
 * ## example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const policy = gcp.cloudrunv2.getServiceIamPolicy({
 *     project: _default.project,
 *     location: _default.location,
 *     name: _default.name,
 * });
 * ```
 */
export function getServiceIamPolicy(args: GetServiceIamPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetServiceIamPolicyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gcp:cloudrunv2/getServiceIamPolicy:getServiceIamPolicy", {
        "location": args.location,
        "name": args.name,
        "project": args.project,
    }, opts);
}

/**
 * A collection of arguments for invoking getServiceIamPolicy.
 */
export interface GetServiceIamPolicyArgs {
    /**
     * The location of the cloud run service Used to find the parent resource to bind the IAM policy to. If not specified,
     * the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
     * location is specified, it is taken from the provider configuration.
     */
    location?: string;
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
 * A collection of values returned by getServiceIamPolicy.
 */
export interface GetServiceIamPolicyResult {
    /**
     * (Computed) The etag of the IAM policy.
     */
    readonly etag: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly location: string;
    readonly name: string;
    /**
     * (Required only by `gcp.cloudrunv2.ServiceIamPolicy`) The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     */
    readonly policyData: string;
    readonly project: string;
}
/**
 * Retrieves the current IAM policy data for service
 *
 * ## example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const policy = gcp.cloudrunv2.getServiceIamPolicy({
 *     project: _default.project,
 *     location: _default.location,
 *     name: _default.name,
 * });
 * ```
 */
export function getServiceIamPolicyOutput(args: GetServiceIamPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetServiceIamPolicyResult> {
    return pulumi.output(args).apply((a: any) => getServiceIamPolicy(a, opts))
}

/**
 * A collection of arguments for invoking getServiceIamPolicy.
 */
export interface GetServiceIamPolicyOutputArgs {
    /**
     * The location of the cloud run service Used to find the parent resource to bind the IAM policy to. If not specified,
     * the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
     * location is specified, it is taken from the provider configuration.
     */
    location?: pulumi.Input<string>;
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
