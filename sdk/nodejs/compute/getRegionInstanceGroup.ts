// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Get a Compute Region Instance Group within GCE.
 * For more information, see [the official documentation](https://cloud.google.com/compute/docs/instance-groups/distributing-instances-with-regional-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/latest/regionInstanceGroups).
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const group = pulumi.output(gcp.compute.getRegionInstanceGroup({
 *     name: "instance-group-name",
 * }));
 * ```
 *
 * The most common use of this datasource will be to fetch information about the instances inside regional managed instance groups, for instance:
 */
export function getRegionInstanceGroup(args?: GetRegionInstanceGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetRegionInstanceGroupResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("gcp:compute/getRegionInstanceGroup:getRegionInstanceGroup", {
        "name": args.name,
        "project": args.project,
        "region": args.region,
        "selfLink": args.selfLink,
    }, opts);
}

/**
 * A collection of arguments for invoking getRegionInstanceGroup.
 */
export interface GetRegionInstanceGroupArgs {
    /**
     * The name of the instance group.  One of `name` or `selfLink` must be provided.
     */
    name?: string;
    /**
     * The ID of the project in which the resource belongs.
     * If `selfLink` is provided, this value is ignored.  If neither `selfLink`
     * nor `project` are provided, the provider project is used.
     */
    project?: string;
    /**
     * The region in which the resource belongs.  If `selfLink`
     * is provided, this value is ignored.  If neither `selfLink` nor `region` are
     * provided, the provider region is used.
     */
    region?: string;
    /**
     * The link to the instance group.  One of `name` or `selfLink` must be provided.
     */
    selfLink?: string;
}

/**
 * A collection of values returned by getRegionInstanceGroup.
 */
export interface GetRegionInstanceGroupResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * List of instances in the group, as a list of resources, each containing:
     */
    readonly instances: outputs.compute.GetRegionInstanceGroupInstance[];
    /**
     * String port name
     */
    readonly name: string;
    readonly project: string;
    readonly region: string;
    readonly selfLink: string;
    /**
     * The number of instances in the group.
     */
    readonly size: number;
}

export function getRegionInstanceGroupOutput(args?: GetRegionInstanceGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRegionInstanceGroupResult> {
    return pulumi.output(args).apply(a => getRegionInstanceGroup(a, opts))
}

/**
 * A collection of arguments for invoking getRegionInstanceGroup.
 */
export interface GetRegionInstanceGroupOutputArgs {
    /**
     * The name of the instance group.  One of `name` or `selfLink` must be provided.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If `selfLink` is provided, this value is ignored.  If neither `selfLink`
     * nor `project` are provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The region in which the resource belongs.  If `selfLink`
     * is provided, this value is ignored.  If neither `selfLink` nor `region` are
     * provided, the provider region is used.
     */
    region?: pulumi.Input<string>;
    /**
     * The link to the instance group.  One of `name` or `selfLink` must be provided.
     */
    selfLink?: pulumi.Input<string>;
}
