// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Get a Cloud Router's status within GCE from its name and region. This data source exposes the
 * routes learned by a Cloud Router via BGP peers.
 *
 * For more information see [the official documentation](https://cloud.google.com/network-connectivity/docs/router/how-to/viewing-router-details)
 * and
 * [API](https://cloud.google.com/compute/docs/reference/rest/v1/routers/getRouterStatus).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const my-router = gcp.compute.getRouterStatus({
 *     name: "myrouter",
 * });
 * ```
 */
/** @deprecated gcp.compute.RouterStatus has been deprecated in favor of gcp.compute.getRouterStatus */
export function routerStatus(args: RouterStatusArgs, opts?: pulumi.InvokeOptions): Promise<RouterStatusResult> {
    pulumi.log.warn("routerStatus is deprecated: gcp.compute.RouterStatus has been deprecated in favor of gcp.compute.getRouterStatus")

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gcp:compute/routerStatus:RouterStatus", {
        "name": args.name,
        "project": args.project,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking RouterStatus.
 */
export interface RouterStatusArgs {
    /**
     * The name of the router.
     */
    name: string;
    /**
     * The ID of the project in which the resource
     * belongs. If it is not provided, the provider project is used.
     */
    project?: string;
    /**
     * The region this router has been created in. If
     * unspecified, this defaults to the region configured in the provider.
     */
    region?: string;
}

/**
 * A collection of values returned by RouterStatus.
 */
export interface RouterStatusResult {
    /**
     * List of best `compute#routes` configurations for this router's network. See gcp.compute.Route resource for available attributes.
     */
    readonly bestRoutes: outputs.compute.RouterStatusBestRoute[];
    /**
     * List of best `compute#routes` for this specific router. See gcp.compute.Route resource for available attributes.
     */
    readonly bestRoutesForRouters: outputs.compute.RouterStatusBestRoutesForRouter[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    /**
     * The network name or resource link to the parent
     * network of this subnetwork.
     */
    readonly network: string;
    readonly project?: string;
    readonly region: string;
}
/**
 * Get a Cloud Router's status within GCE from its name and region. This data source exposes the
 * routes learned by a Cloud Router via BGP peers.
 *
 * For more information see [the official documentation](https://cloud.google.com/network-connectivity/docs/router/how-to/viewing-router-details)
 * and
 * [API](https://cloud.google.com/compute/docs/reference/rest/v1/routers/getRouterStatus).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const my-router = gcp.compute.getRouterStatus({
 *     name: "myrouter",
 * });
 * ```
 */
/** @deprecated gcp.compute.RouterStatus has been deprecated in favor of gcp.compute.getRouterStatus */
export function routerStatusOutput(args: RouterStatusOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<RouterStatusResult> {
    return pulumi.output(args).apply((a: any) => routerStatus(a, opts))
}

/**
 * A collection of arguments for invoking RouterStatus.
 */
export interface RouterStatusOutputArgs {
    /**
     * The name of the router.
     */
    name: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource
     * belongs. If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The region this router has been created in. If
     * unspecified, this defaults to the region configured in the provider.
     */
    region?: pulumi.Input<string>;
}
