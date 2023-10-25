// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get information about the supported alloydb database flags in a location.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const qa = gcp.alloydb.getSupportedDatabaseFlags({
 *     location: "us-central1",
 * });
 * ```
 */
export function getSupportedDatabaseFlags(args: GetSupportedDatabaseFlagsArgs, opts?: pulumi.InvokeOptions): Promise<GetSupportedDatabaseFlagsResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gcp:alloydb/getSupportedDatabaseFlags:getSupportedDatabaseFlags", {
        "location": args.location,
        "project": args.project,
    }, opts);
}

/**
 * A collection of arguments for invoking getSupportedDatabaseFlags.
 */
export interface GetSupportedDatabaseFlagsArgs {
    /**
     * The canonical id of the location. For example: `us-east1`.
     */
    location: string;
    /**
     * The ID of the project.
     */
    project?: string;
}

/**
 * A collection of values returned by getSupportedDatabaseFlags.
 */
export interface GetSupportedDatabaseFlagsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly location: string;
    readonly project?: string;
    /**
     * Contains a list of `flag`, which contains the details about a particular flag.
     */
    readonly supportedDatabaseFlags: outputs.alloydb.GetSupportedDatabaseFlagsSupportedDatabaseFlag[];
}
/**
 * Use this data source to get information about the supported alloydb database flags in a location.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const qa = gcp.alloydb.getSupportedDatabaseFlags({
 *     location: "us-central1",
 * });
 * ```
 */
export function getSupportedDatabaseFlagsOutput(args: GetSupportedDatabaseFlagsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSupportedDatabaseFlagsResult> {
    return pulumi.output(args).apply((a: any) => getSupportedDatabaseFlags(a, opts))
}

/**
 * A collection of arguments for invoking getSupportedDatabaseFlags.
 */
export interface GetSupportedDatabaseFlagsOutputArgs {
    /**
     * The canonical id of the location. For example: `us-east1`.
     */
    location: pulumi.Input<string>;
    /**
     * The ID of the project.
     */
    project?: pulumi.Input<string>;
}
