// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Retrieve a list of testable permissions for a resource. Testable permissions mean the permissions that user can add or remove in a role at a given resource. The resource can be referenced either via the full resource name or via a URI.
 *
 * ## Example Usage
 *
 * Retrieve all the supported permissions able to be set on `my-project` that are in either GA or BETA. This is useful for dynamically constructing custom roles.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const perms = gcp.iam.getTestablePermissions({
 *     fullResourceName: "//cloudresourcemanager.googleapis.com/projects/my-project",
 *     stages: [
 *         "GA",
 *         "BETA",
 *     ],
 * });
 * ```
 */
export function getTestablePermissions(args: GetTestablePermissionsArgs, opts?: pulumi.InvokeOptions): Promise<GetTestablePermissionsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gcp:iam/getTestablePermissions:getTestablePermissions", {
        "customSupportLevel": args.customSupportLevel,
        "fullResourceName": args.fullResourceName,
        "stages": args.stages,
    }, opts);
}

/**
 * A collection of arguments for invoking getTestablePermissions.
 */
export interface GetTestablePermissionsArgs {
    /**
     * The level of support for custom roles. Can be one of `"NOT_SUPPORTED"`, `"SUPPORTED"`, `"TESTING"`. Default is `"SUPPORTED"`
     */
    customSupportLevel?: string;
    /**
     * See [full resource name documentation](https://cloud.google.com/apis/design/resource_names#full_resource_name) for more detail.
     */
    fullResourceName: string;
    /**
     * The acceptable release stages of the permission in the output. Note that `BETA` does not include permissions in `GA`, but you can specify both with `["GA", "BETA"]` for example. Can be a list of `"ALPHA"`, `"BETA"`, `"GA"`, `"DEPRECATED"`. Default is `["GA"]`.
     */
    stages?: string[];
}

/**
 * A collection of values returned by getTestablePermissions.
 */
export interface GetTestablePermissionsResult {
    /**
     * The the support level of this permission for custom roles.
     */
    readonly customSupportLevel?: string;
    readonly fullResourceName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * A list of permissions matching the provided input. Structure is defined below.
     */
    readonly permissions: outputs.iam.GetTestablePermissionsPermission[];
    readonly stages?: string[];
}
/**
 * Retrieve a list of testable permissions for a resource. Testable permissions mean the permissions that user can add or remove in a role at a given resource. The resource can be referenced either via the full resource name or via a URI.
 *
 * ## Example Usage
 *
 * Retrieve all the supported permissions able to be set on `my-project` that are in either GA or BETA. This is useful for dynamically constructing custom roles.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const perms = gcp.iam.getTestablePermissions({
 *     fullResourceName: "//cloudresourcemanager.googleapis.com/projects/my-project",
 *     stages: [
 *         "GA",
 *         "BETA",
 *     ],
 * });
 * ```
 */
export function getTestablePermissionsOutput(args: GetTestablePermissionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTestablePermissionsResult> {
    return pulumi.output(args).apply((a: any) => getTestablePermissions(a, opts))
}

/**
 * A collection of arguments for invoking getTestablePermissions.
 */
export interface GetTestablePermissionsOutputArgs {
    /**
     * The level of support for custom roles. Can be one of `"NOT_SUPPORTED"`, `"SUPPORTED"`, `"TESTING"`. Default is `"SUPPORTED"`
     */
    customSupportLevel?: pulumi.Input<string>;
    /**
     * See [full resource name documentation](https://cloud.google.com/apis/design/resource_names#full_resource_name) for more detail.
     */
    fullResourceName: pulumi.Input<string>;
    /**
     * The acceptable release stages of the permission in the output. Note that `BETA` does not include permissions in `GA`, but you can specify both with `["GA", "BETA"]` for example. Can be a list of `"ALPHA"`, `"BETA"`, `"GA"`, `"DEPRECATED"`. Default is `["GA"]`.
     */
    stages?: pulumi.Input<pulumi.Input<string>[]>;
}
