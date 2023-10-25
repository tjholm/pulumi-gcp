// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Get a Serverless VPC Access connector.
 *
 * To get more information about Connector, see:
 *
 * * [API documentation](https://cloud.google.com/vpc/docs/reference/vpcaccess/rest/v1/projects.locations.connectors)
 * * How-to Guides
 *     * [Configuring Serverless VPC Access](https://cloud.google.com/vpc/docs/configure-serverless-vpc-access)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const sample = gcp.vpcaccess.getConnector({
 *     name: "vpc-con",
 * });
 * const connector = new gcp.vpcaccess.Connector("connector", {
 *     ipCidrRange: "10.8.0.0/28",
 *     network: "default",
 *     region: "us-central1",
 * });
 * ```
 */
export function getConnector(args: GetConnectorArgs, opts?: pulumi.InvokeOptions): Promise<GetConnectorResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gcp:vpcaccess/getConnector:getConnector", {
        "name": args.name,
        "project": args.project,
        "region": args.region,
    }, opts);
}

/**
 * A collection of arguments for invoking getConnector.
 */
export interface GetConnectorArgs {
    /**
     * Name of the resource.
     *
     * - - -
     */
    name: string;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    project?: string;
    /**
     * The region in which the resource belongs. If it
     * is not provided, the provider region is used.
     */
    region?: string;
}

/**
 * A collection of values returned by getConnector.
 */
export interface GetConnectorResult {
    readonly connectedProjects: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ipCidrRange: string;
    readonly machineType: string;
    readonly maxInstances: number;
    readonly maxThroughput: number;
    readonly minInstances: number;
    readonly minThroughput: number;
    readonly name: string;
    readonly network: string;
    readonly project?: string;
    readonly region?: string;
    readonly selfLink: string;
    readonly state: string;
    readonly subnets: outputs.vpcaccess.GetConnectorSubnet[];
}
/**
 * Get a Serverless VPC Access connector.
 *
 * To get more information about Connector, see:
 *
 * * [API documentation](https://cloud.google.com/vpc/docs/reference/vpcaccess/rest/v1/projects.locations.connectors)
 * * How-to Guides
 *     * [Configuring Serverless VPC Access](https://cloud.google.com/vpc/docs/configure-serverless-vpc-access)
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const sample = gcp.vpcaccess.getConnector({
 *     name: "vpc-con",
 * });
 * const connector = new gcp.vpcaccess.Connector("connector", {
 *     ipCidrRange: "10.8.0.0/28",
 *     network: "default",
 *     region: "us-central1",
 * });
 * ```
 */
export function getConnectorOutput(args: GetConnectorOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetConnectorResult> {
    return pulumi.output(args).apply((a: any) => getConnector(a, opts))
}

/**
 * A collection of arguments for invoking getConnector.
 */
export interface GetConnectorOutputArgs {
    /**
     * Name of the resource.
     *
     * - - -
     */
    name: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The region in which the resource belongs. If it
     * is not provided, the provider region is used.
     */
    region?: pulumi.Input<string>;
}
