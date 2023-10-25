// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Get a router within GCE from its name and VPC.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const my-router = gcp.compute.getRouter({
 *     name: "myrouter-us-east1",
 *     network: "my-network",
 * });
 * ```
 */
export function getRouter(args: GetRouterArgs, opts?: pulumi.InvokeOptions): Promise<GetRouterResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gcp:compute/getRouter:getRouter", {
        "name": args.name,
        "network": args.network,
        "project": args.project,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getRouter.
 */
export interface GetRouterArgs {
    /**
     * The name of the router.
     */
    name: string;
    /**
     * The VPC network on which this router lives.
     */
    network: string;
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
 * A collection of values returned by getRouter.
 */
export interface GetRouterResult {
    readonly bgps: outputs.compute.GetRouterBgp[];
    readonly creationTimestamp: string;
    readonly description: string;
    readonly encryptedInterconnectRouter: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly network: string;
    readonly project?: string;
    readonly region?: string;
    readonly selfLink: string;
}
/**
 * Get a router within GCE from its name and VPC.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const my-router = gcp.compute.getRouter({
 *     name: "myrouter-us-east1",
 *     network: "my-network",
 * });
 * ```
 */
export function getRouterOutput(args: GetRouterOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRouterResult> {
    return pulumi.output(args).apply((a: any) => getRouter(a, opts))
}

/**
 * A collection of arguments for invoking getRouter.
 */
export interface GetRouterOutputArgs {
    /**
     * The name of the router.
     */
    name: pulumi.Input<string>;
    /**
     * The VPC network on which this router lives.
     */
    network: pulumi.Input<string>;
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
