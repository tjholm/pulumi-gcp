// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";

/**
 * Creates a new Google SQL Database Instance. For more information, see the [official documentation](https://cloud.google.com/sql/),
 * or the [JSON API](https://cloud.google.com/sql/docs/admin-api/v1beta4/instances).
 * 
 * ~> **NOTE on `google_sql_database_instance`:** - Second-generation instances include a
 * default 'root'@'%' user with no password. This user will be deleted by Terraform on
 * instance creation. You should use `google_sql_user` to define a custom user with
 * a restricted host and strong password.
 */
export class DatabaseInstance extends pulumi.CustomResource {
    /**
     * Get an existing DatabaseInstance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatabaseInstanceState): DatabaseInstance {
        return new DatabaseInstance(name, <any>state, { id });
    }

    /**
     * The connection name of the instance to be used in connection strings.
     */
    public /*out*/ readonly connectionName: pulumi.Output<string>;
    /**
     * The MySQL version to
     * use. Can be `MYSQL_5_6`, `MYSQL_5_7` or `POSTGRES_9_6` for second-generation
     * instances, or `MYSQL_5_5` or `MYSQL_5_6` for first-generation instances.
     * See [Second Generation Capabilities](https://cloud.google.com/sql/docs/1st-2nd-gen-differences)
     * for more information. `POSTGRES_9_6` support is in [Beta](/docs/providers/google/index.html#beta-features).
     */
    public readonly databaseVersion: pulumi.Output<string | undefined>;
    /**
     * The first IPv4 address of the addresses assigned. This is
     * is to support accessing the [first address in the list in a terraform output](https://github.com/terraform-providers/terraform-provider-google/issues/912)
     * when the resource is configured with a `count`.
     */
    public /*out*/ readonly firstIpAddress: pulumi.Output<string>;
    public /*out*/ readonly ipAddresses: pulumi.Output<{ ipAddress: string, timeToRetire: string }[]>;
    /**
     * The name of the instance that will act as
     * the master in the replication setup. Note, this requires the master to have
     * `binary_log_enabled` set, as well as existing backups.
     */
    public readonly masterInstanceName: pulumi.Output<string>;
    /**
     * A name for this whitelist entry.
     */
    public readonly name: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project: pulumi.Output<string>;
    /**
     * The region the instance will sit in. Note, first-generation Cloud SQL instance
     * regions do not line up with the Google Compute Engine (GCE) regions, and Cloud SQL is not
     * available in all regions - choose from one of the options listed [here](https://cloud.google.com/sql/docs/mysql/instance-locations).
     * A valid region must be provided to use this resource. If a region is not provided in the resource definition,
     * the provider region will be used instead, but this will be an apply-time error for all first-generation
     * instances *and* for second-generation instances if the provider region is not supported with Cloud SQL.
     * If you choose not to provide the `region` argument for this resource, make sure you understand this.
     */
    public readonly region: pulumi.Output<string | undefined>;
    /**
     * The configuration for replication. The
     * configuration is detailed below.
     */
    public readonly replicaConfiguration: pulumi.Output<{ caCertificate?: string, clientCertificate?: string, clientKey?: string, connectRetryInterval?: number, dumpFilePath?: string, failoverTarget?: boolean, masterHeartbeatPeriod?: number, password?: string, sslCipher?: string, username?: string, verifyServerCertificate?: boolean }>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink: pulumi.Output<string>;
    public /*out*/ readonly serverCaCert: pulumi.Output<{ cert: string, commonName: string, createTime: string, expirationTime: string, sha1Fingerprint: string }>;
    /**
     * The settings to use for the database. The
     * configuration is detailed below.
     */
    public readonly settings: pulumi.Output<{ activationPolicy: string, authorizedGaeApplications?: string[], availabilityType: string, backupConfiguration: { binaryLogEnabled?: boolean, enabled?: boolean, startTime: string }, crashSafeReplication: boolean, databaseFlags?: { name?: string, value?: string }[], diskAutoresize?: boolean, diskSize: number, diskType: string, ipConfiguration: { authorizedNetworks?: { expirationTime?: string, name?: string, value?: string }[], ipv4Enabled: boolean, requireSsl?: boolean }, locationPreference: { followGaeApplication?: string, zone?: string }, maintenanceWindow?: { day?: number, hour?: number, updateTrack?: string }, pricingPlan?: string, replicationType?: string, tier: string, userLabels?: {[key: string]: string}, version: number }>;

    /**
     * Create a DatabaseInstance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatabaseInstanceArgs, opts?: pulumi.ResourceOptions)
    constructor(name: string, argsOrState?: DatabaseInstanceArgs | DatabaseInstanceState, opts?: pulumi.ResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state: DatabaseInstanceState = argsOrState as DatabaseInstanceState | undefined;
            inputs["connectionName"] = state ? state.connectionName : undefined;
            inputs["databaseVersion"] = state ? state.databaseVersion : undefined;
            inputs["firstIpAddress"] = state ? state.firstIpAddress : undefined;
            inputs["ipAddresses"] = state ? state.ipAddresses : undefined;
            inputs["masterInstanceName"] = state ? state.masterInstanceName : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["replicaConfiguration"] = state ? state.replicaConfiguration : undefined;
            inputs["selfLink"] = state ? state.selfLink : undefined;
            inputs["serverCaCert"] = state ? state.serverCaCert : undefined;
            inputs["settings"] = state ? state.settings : undefined;
        } else {
            const args = argsOrState as DatabaseInstanceArgs | undefined;
            if (!args || args.settings === undefined) {
                throw new Error("Missing required property 'settings'");
            }
            inputs["databaseVersion"] = args ? args.databaseVersion : undefined;
            inputs["masterInstanceName"] = args ? args.masterInstanceName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["replicaConfiguration"] = args ? args.replicaConfiguration : undefined;
            inputs["settings"] = args ? args.settings : undefined;
            inputs["connectionName"] = undefined /*out*/;
            inputs["firstIpAddress"] = undefined /*out*/;
            inputs["ipAddresses"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["serverCaCert"] = undefined /*out*/;
        }
        super("gcp:sql/databaseInstance:DatabaseInstance", name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DatabaseInstance resources.
 */
export interface DatabaseInstanceState {
    /**
     * The connection name of the instance to be used in connection strings.
     */
    readonly connectionName?: pulumi.Input<string>;
    /**
     * The MySQL version to
     * use. Can be `MYSQL_5_6`, `MYSQL_5_7` or `POSTGRES_9_6` for second-generation
     * instances, or `MYSQL_5_5` or `MYSQL_5_6` for first-generation instances.
     * See [Second Generation Capabilities](https://cloud.google.com/sql/docs/1st-2nd-gen-differences)
     * for more information. `POSTGRES_9_6` support is in [Beta](/docs/providers/google/index.html#beta-features).
     */
    readonly databaseVersion?: pulumi.Input<string>;
    /**
     * The first IPv4 address of the addresses assigned. This is
     * is to support accessing the [first address in the list in a terraform output](https://github.com/terraform-providers/terraform-provider-google/issues/912)
     * when the resource is configured with a `count`.
     */
    readonly firstIpAddress?: pulumi.Input<string>;
    readonly ipAddresses?: pulumi.Input<pulumi.Input<{ ipAddress?: pulumi.Input<string>, timeToRetire?: pulumi.Input<string> }>[]>;
    /**
     * The name of the instance that will act as
     * the master in the replication setup. Note, this requires the master to have
     * `binary_log_enabled` set, as well as existing backups.
     */
    readonly masterInstanceName?: pulumi.Input<string>;
    /**
     * A name for this whitelist entry.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The region the instance will sit in. Note, first-generation Cloud SQL instance
     * regions do not line up with the Google Compute Engine (GCE) regions, and Cloud SQL is not
     * available in all regions - choose from one of the options listed [here](https://cloud.google.com/sql/docs/mysql/instance-locations).
     * A valid region must be provided to use this resource. If a region is not provided in the resource definition,
     * the provider region will be used instead, but this will be an apply-time error for all first-generation
     * instances *and* for second-generation instances if the provider region is not supported with Cloud SQL.
     * If you choose not to provide the `region` argument for this resource, make sure you understand this.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The configuration for replication. The
     * configuration is detailed below.
     */
    readonly replicaConfiguration?: pulumi.Input<{ caCertificate?: pulumi.Input<string>, clientCertificate?: pulumi.Input<string>, clientKey?: pulumi.Input<string>, connectRetryInterval?: pulumi.Input<number>, dumpFilePath?: pulumi.Input<string>, failoverTarget?: pulumi.Input<boolean>, masterHeartbeatPeriod?: pulumi.Input<number>, password?: pulumi.Input<string>, sslCipher?: pulumi.Input<string>, username?: pulumi.Input<string>, verifyServerCertificate?: pulumi.Input<boolean> }>;
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: pulumi.Input<string>;
    readonly serverCaCert?: pulumi.Input<{ cert?: pulumi.Input<string>, commonName?: pulumi.Input<string>, createTime?: pulumi.Input<string>, expirationTime?: pulumi.Input<string>, sha1Fingerprint?: pulumi.Input<string> }>;
    /**
     * The settings to use for the database. The
     * configuration is detailed below.
     */
    readonly settings?: pulumi.Input<{ activationPolicy?: pulumi.Input<string>, authorizedGaeApplications?: pulumi.Input<pulumi.Input<string>[]>, availabilityType?: pulumi.Input<string>, backupConfiguration?: pulumi.Input<{ binaryLogEnabled?: pulumi.Input<boolean>, enabled?: pulumi.Input<boolean>, startTime?: pulumi.Input<string> }>, crashSafeReplication?: pulumi.Input<boolean>, databaseFlags?: pulumi.Input<pulumi.Input<{ name?: pulumi.Input<string>, value?: pulumi.Input<string> }>[]>, diskAutoresize?: pulumi.Input<boolean>, diskSize?: pulumi.Input<number>, diskType?: pulumi.Input<string>, ipConfiguration?: pulumi.Input<{ authorizedNetworks?: pulumi.Input<pulumi.Input<{ expirationTime?: pulumi.Input<string>, name?: pulumi.Input<string>, value?: pulumi.Input<string> }>[]>, ipv4Enabled?: pulumi.Input<boolean>, requireSsl?: pulumi.Input<boolean> }>, locationPreference?: pulumi.Input<{ followGaeApplication?: pulumi.Input<string>, zone?: pulumi.Input<string> }>, maintenanceWindow?: pulumi.Input<{ day?: pulumi.Input<number>, hour?: pulumi.Input<number>, updateTrack?: pulumi.Input<string> }>, pricingPlan?: pulumi.Input<string>, replicationType?: pulumi.Input<string>, tier: pulumi.Input<string>, userLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, version?: pulumi.Input<number> }>;
}

/**
 * The set of arguments for constructing a DatabaseInstance resource.
 */
export interface DatabaseInstanceArgs {
    /**
     * The MySQL version to
     * use. Can be `MYSQL_5_6`, `MYSQL_5_7` or `POSTGRES_9_6` for second-generation
     * instances, or `MYSQL_5_5` or `MYSQL_5_6` for first-generation instances.
     * See [Second Generation Capabilities](https://cloud.google.com/sql/docs/1st-2nd-gen-differences)
     * for more information. `POSTGRES_9_6` support is in [Beta](/docs/providers/google/index.html#beta-features).
     */
    readonly databaseVersion?: pulumi.Input<string>;
    /**
     * The name of the instance that will act as
     * the master in the replication setup. Note, this requires the master to have
     * `binary_log_enabled` set, as well as existing backups.
     */
    readonly masterInstanceName?: pulumi.Input<string>;
    /**
     * A name for this whitelist entry.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    readonly project?: pulumi.Input<string>;
    /**
     * The region the instance will sit in. Note, first-generation Cloud SQL instance
     * regions do not line up with the Google Compute Engine (GCE) regions, and Cloud SQL is not
     * available in all regions - choose from one of the options listed [here](https://cloud.google.com/sql/docs/mysql/instance-locations).
     * A valid region must be provided to use this resource. If a region is not provided in the resource definition,
     * the provider region will be used instead, but this will be an apply-time error for all first-generation
     * instances *and* for second-generation instances if the provider region is not supported with Cloud SQL.
     * If you choose not to provide the `region` argument for this resource, make sure you understand this.
     */
    readonly region?: pulumi.Input<string>;
    /**
     * The configuration for replication. The
     * configuration is detailed below.
     */
    readonly replicaConfiguration?: pulumi.Input<{ caCertificate?: pulumi.Input<string>, clientCertificate?: pulumi.Input<string>, clientKey?: pulumi.Input<string>, connectRetryInterval?: pulumi.Input<number>, dumpFilePath?: pulumi.Input<string>, failoverTarget?: pulumi.Input<boolean>, masterHeartbeatPeriod?: pulumi.Input<number>, password?: pulumi.Input<string>, sslCipher?: pulumi.Input<string>, username?: pulumi.Input<string>, verifyServerCertificate?: pulumi.Input<boolean> }>;
    /**
     * The settings to use for the database. The
     * configuration is detailed below.
     */
    readonly settings: pulumi.Input<{ activationPolicy?: pulumi.Input<string>, authorizedGaeApplications?: pulumi.Input<pulumi.Input<string>[]>, availabilityType?: pulumi.Input<string>, backupConfiguration?: pulumi.Input<{ binaryLogEnabled?: pulumi.Input<boolean>, enabled?: pulumi.Input<boolean>, startTime?: pulumi.Input<string> }>, crashSafeReplication?: pulumi.Input<boolean>, databaseFlags?: pulumi.Input<pulumi.Input<{ name?: pulumi.Input<string>, value?: pulumi.Input<string> }>[]>, diskAutoresize?: pulumi.Input<boolean>, diskSize?: pulumi.Input<number>, diskType?: pulumi.Input<string>, ipConfiguration?: pulumi.Input<{ authorizedNetworks?: pulumi.Input<pulumi.Input<{ expirationTime?: pulumi.Input<string>, name?: pulumi.Input<string>, value?: pulumi.Input<string> }>[]>, ipv4Enabled?: pulumi.Input<boolean>, requireSsl?: pulumi.Input<boolean> }>, locationPreference?: pulumi.Input<{ followGaeApplication?: pulumi.Input<string>, zone?: pulumi.Input<string> }>, maintenanceWindow?: pulumi.Input<{ day?: pulumi.Input<number>, hour?: pulumi.Input<number>, updateTrack?: pulumi.Input<string> }>, pricingPlan?: pulumi.Input<string>, replicationType?: pulumi.Input<string>, tier: pulumi.Input<string>, userLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>, version?: pulumi.Input<number> }>;
}
