// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Get TensorFlow versions available for a project. For more information see the [official documentation](https://cloud.google.com/tpu/docs/) and [API](https://cloud.google.com/tpu/docs/reference/rest/v1/projects.locations.tensorflowVersions).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const available = pulumi.output(gcp.tpu.getTensorflowVersions());
 * ```
 * ### Configure Basic TPU Node With Available Version
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const available = gcp.tpu.getTensorflowVersions({});
 * const tpu = new gcp.tpu.Node("tpu", {
 *     zone: "us-central1-b",
 *     acceleratorType: "v3-8",
 *     tensorflowVersion: available.then(available => available.versions?[0]),
 *     cidrBlock: "10.2.0.0/29",
 * });
 * ```
 */
export function getTensorflowVersions(args?: GetTensorflowVersionsArgs, opts?: pulumi.InvokeOptions): Promise<GetTensorflowVersionsResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("gcp:tpu/getTensorflowVersions:getTensorflowVersions", {
        "project": args.project,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getTensorflowVersions.
 */
export interface GetTensorflowVersionsArgs {
    /**
     * The project to list versions for. If it
     * is not provided, the provider project is used.
     */
    project?: string;
    /**
     * The zone to list versions for. If it
     * is not provided, the provider zone is used.
     */
    zone?: string;
}

/**
 * A collection of values returned by getTensorflowVersions.
 */
export interface GetTensorflowVersionsResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly project: string;
    /**
     * The list of TensorFlow versions available for the given project and zone.
     */
    readonly versions: string[];
    readonly zone: string;
}

export function getTensorflowVersionsOutput(args?: GetTensorflowVersionsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTensorflowVersionsResult> {
    return pulumi.output(args).apply(a => getTensorflowVersions(a, opts))
}

/**
 * A collection of arguments for invoking getTensorflowVersions.
 */
export interface GetTensorflowVersionsOutputArgs {
    /**
     * The project to list versions for. If it
     * is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The zone to list versions for. If it
     * is not provided, the provider zone is used.
     */
    zone?: pulumi.Input<string>;
}
