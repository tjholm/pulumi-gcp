// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Retrieves the current IAM policy data for customtargettype
 *
 * ## example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const policy = gcp.clouddeploy.getCustomTargetTypeIamPolicy({
 *     project: custom_target_type.project,
 *     location: custom_target_type.location,
 *     name: custom_target_type.name,
 * });
 * ```
 */
export function getCustomTargetTypeIamPolicy(args: GetCustomTargetTypeIamPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetCustomTargetTypeIamPolicyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gcp:clouddeploy/getCustomTargetTypeIamPolicy:getCustomTargetTypeIamPolicy", {
        "location": args.location,
        "name": args.name,
        "project": args.project,
    }, opts);
}

/**
 * A collection of arguments for invoking getCustomTargetTypeIamPolicy.
 */
export interface GetCustomTargetTypeIamPolicyArgs {
    /**
     * The location of the source. Used to find the parent resource to bind the IAM policy to. If not specified,
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
 * A collection of values returned by getCustomTargetTypeIamPolicy.
 */
export interface GetCustomTargetTypeIamPolicyResult {
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
     * (Required only by `gcp.clouddeploy.CustomTargetTypeIamPolicy`) The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     */
    readonly policyData: string;
    readonly project: string;
}
/**
 * Retrieves the current IAM policy data for customtargettype
 *
 * ## example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const policy = gcp.clouddeploy.getCustomTargetTypeIamPolicy({
 *     project: custom_target_type.project,
 *     location: custom_target_type.location,
 *     name: custom_target_type.name,
 * });
 * ```
 */
export function getCustomTargetTypeIamPolicyOutput(args: GetCustomTargetTypeIamPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCustomTargetTypeIamPolicyResult> {
    return pulumi.output(args).apply((a: any) => getCustomTargetTypeIamPolicy(a, opts))
}

/**
 * A collection of arguments for invoking getCustomTargetTypeIamPolicy.
 */
export interface GetCustomTargetTypeIamPolicyOutputArgs {
    /**
     * The location of the source. Used to find the parent resource to bind the IAM policy to. If not specified,
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
