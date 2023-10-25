// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Use this data source to get information about a list of Cloud SQL instances in a project. You can also apply some filters over this list to get a more filtered list of Cloud SQL instances.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const qa = gcp.sql.getDatabaseInstances({
 *     project: "test-project",
 * });
 * ```
 */
export function getDatabaseInstances(args?: GetDatabaseInstancesArgs, opts?: pulumi.InvokeOptions): Promise<GetDatabaseInstancesResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gcp:sql/getDatabaseInstances:getDatabaseInstances", {
        "databaseVersion": args.databaseVersion,
        "project": args.project,
        "region": args.region,
        "state": args.state,
        "tier": args.tier,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getDatabaseInstances.
 */
export interface GetDatabaseInstancesArgs {
    /**
     * To filter out the Cloud SQL instances which are of the specified database version.
     */
    databaseVersion?: string;
    /**
     * The ID of the project in which the resources belong. If it is not provided, the provider project is used.
     */
    project?: string;
    /**
     * To filter out the Cloud SQL instances which are located in the specified region.
     */
    region?: string;
    /**
     * To filter out the Cloud SQL instances based on the current serving state of the database instance. Supported values include `SQL_INSTANCE_STATE_UNSPECIFIED`, `RUNNABLE`, `SUSPENDED`, `PENDING_DELETE`, `PENDING_CREATE`, `MAINTENANCE`, `FAILED`.
     */
    state?: string;
    /**
     * To filter out the Cloud SQL instances based on the tier(or machine type) of the database instances.
     */
    tier?: string;
    /**
     * To filter out the Cloud SQL instances which are located in the specified zone. This zone refers to the Compute Engine zone that the instance is currently serving from.
     */
    zone?: string;
}

/**
 * A collection of values returned by getDatabaseInstances.
 */
export interface GetDatabaseInstancesResult {
    readonly databaseVersion?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly instances: outputs.sql.GetDatabaseInstancesInstance[];
    readonly project?: string;
    readonly region?: string;
    readonly state?: string;
    readonly tier?: string;
    readonly zone?: string;
}
/**
 * Use this data source to get information about a list of Cloud SQL instances in a project. You can also apply some filters over this list to get a more filtered list of Cloud SQL instances.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const qa = gcp.sql.getDatabaseInstances({
 *     project: "test-project",
 * });
 * ```
 */
export function getDatabaseInstancesOutput(args?: GetDatabaseInstancesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDatabaseInstancesResult> {
    return pulumi.output(args).apply((a: any) => getDatabaseInstances(a, opts))
}

/**
 * A collection of arguments for invoking getDatabaseInstances.
 */
export interface GetDatabaseInstancesOutputArgs {
    /**
     * To filter out the Cloud SQL instances which are of the specified database version.
     */
    databaseVersion?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resources belong. If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * To filter out the Cloud SQL instances which are located in the specified region.
     */
    region?: pulumi.Input<string>;
    /**
     * To filter out the Cloud SQL instances based on the current serving state of the database instance. Supported values include `SQL_INSTANCE_STATE_UNSPECIFIED`, `RUNNABLE`, `SUSPENDED`, `PENDING_DELETE`, `PENDING_CREATE`, `MAINTENANCE`, `FAILED`.
     */
    state?: pulumi.Input<string>;
    /**
     * To filter out the Cloud SQL instances based on the tier(or machine type) of the database instances.
     */
    tier?: pulumi.Input<string>;
    /**
     * To filter out the Cloud SQL instances which are located in the specified zone. This zone refers to the Compute Engine zone that the instance is currently serving from.
     */
    zone?: pulumi.Input<string>;
}
