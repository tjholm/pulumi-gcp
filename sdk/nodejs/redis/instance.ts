// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A Google Cloud Redis instance.
 *
 * To get more information about Instance, see:
 *
 * * [API documentation](https://cloud.google.com/memorystore/docs/redis/reference/rest/v1/projects.locations.instances)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/memorystore/docs/redis/)
 *
 * ## Example Usage
 * ### Redis Instance Full
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const redis-network = gcp.compute.getNetwork({
 *     name: "redis-test-network",
 * });
 * const cache = new gcp.redis.Instance("cache", {
 *     tier: "STANDARD_HA",
 *     memorySizeGb: 1,
 *     locationId: "us-central1-a",
 *     alternativeLocationId: "us-central1-f",
 *     authorizedNetwork: redis_network.then(redis_network => redis_network.id),
 *     redisVersion: "REDIS_4_0",
 *     displayName: "Test Instance",
 *     reservedIpRange: "192.168.0.0/29",
 *     labels: {
 *         my_key: "my_val",
 *         other_key: "other_val",
 *     },
 *     maintenancePolicy: {
 *         weeklyMaintenanceWindows: [{
 *             day: "TUESDAY",
 *             startTime: {
 *                 hours: 0,
 *                 minutes: 30,
 *                 seconds: 0,
 *                 nanos: 0,
 *             },
 *         }],
 *     },
 * });
 * ```
 * ### Redis Instance Private Service
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * // This example assumes this network already exists.
 * // The API creates a tenant network per network authorized for a
 * // Redis instance and that network is not deleted when the user-created
 * // network (authorized_network) is deleted, so this prevents issues
 * // with tenant network quota.
 * // If this network hasn't been created and you are using this example in your
 * // config, add an additional network resource or change
 * // this from "data"to "resource"
 * const redis_network = new gcp.compute.Network("redis-network", {});
 * const serviceRange = new gcp.compute.GlobalAddress("serviceRange", {
 *     purpose: "VPC_PEERING",
 *     addressType: "INTERNAL",
 *     prefixLength: 16,
 *     network: redis_network.id,
 * });
 * const privateServiceConnection = new gcp.servicenetworking.Connection("privateServiceConnection", {
 *     network: redis_network.id,
 *     service: "servicenetworking.googleapis.com",
 *     reservedPeeringRanges: [serviceRange.name],
 * });
 * const cache = new gcp.redis.Instance("cache", {
 *     tier: "STANDARD_HA",
 *     memorySizeGb: 1,
 *     locationId: "us-central1-a",
 *     alternativeLocationId: "us-central1-f",
 *     authorizedNetwork: redis_network.id,
 *     connectMode: "PRIVATE_SERVICE_ACCESS",
 *     redisVersion: "REDIS_4_0",
 *     displayName: "Test Instance",
 * }, {
 *     dependsOn: [privateServiceConnection],
 * });
 * ```
 * ### Redis Instance Mrr
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const redis-network = gcp.compute.getNetwork({
 *     name: "redis-test-network",
 * });
 * const cache = new gcp.redis.Instance("cache", {
 *     tier: "STANDARD_HA",
 *     memorySizeGb: 5,
 *     locationId: "us-central1-a",
 *     alternativeLocationId: "us-central1-f",
 *     authorizedNetwork: redis_network.then(redis_network => redis_network.id),
 *     redisVersion: "REDIS_6_X",
 *     displayName: "Terraform Test Instance",
 *     reservedIpRange: "192.168.0.0/28",
 *     replicaCount: 5,
 *     readReplicasMode: "READ_REPLICAS_ENABLED",
 *     labels: {
 *         my_key: "my_val",
 *         other_key: "other_val",
 *     },
 * });
 * ```
 * ### Redis Instance Cmek
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const redisKeyring = new gcp.kms.KeyRing("redisKeyring", {location: "us-central1"});
 * const redisKey = new gcp.kms.CryptoKey("redisKey", {keyRing: redisKeyring.id});
 * const redis-network = gcp.compute.getNetwork({
 *     name: "redis-test-network",
 * });
 * const cache = new gcp.redis.Instance("cache", {
 *     tier: "STANDARD_HA",
 *     memorySizeGb: 1,
 *     locationId: "us-central1-a",
 *     alternativeLocationId: "us-central1-f",
 *     authorizedNetwork: redis_network.then(redis_network => redis_network.id),
 *     redisVersion: "REDIS_6_X",
 *     displayName: "Terraform Test Instance",
 *     reservedIpRange: "192.168.0.0/29",
 *     labels: {
 *         my_key: "my_val",
 *         other_key: "other_val",
 *     },
 *     customerManagedKey: redisKey.id,
 * });
 * ```
 *
 * ## Import
 *
 * Instance can be imported using any of these accepted formats* `projects/{{project}}/locations/{{region}}/instances/{{name}}` * `{{project}}/{{region}}/{{name}}` * `{{region}}/{{name}}` * `{{name}}` When using the `pulumi import` command, Instance can be imported using one of the formats above. For example
 *
 * ```sh
 *  $ pulumi import gcp:redis/instance:Instance default projects/{{project}}/locations/{{region}}/instances/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:redis/instance:Instance default {{project}}/{{region}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:redis/instance:Instance default {{region}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:redis/instance:Instance default {{name}}
 * ```
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceState, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:redis/instance:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * Only applicable to STANDARD_HA tier which protects the instance
     * against zonal failures by provisioning it across two zones.
     * If provided, it must be a different zone from the one provided in
     * [locationId].
     */
    public readonly alternativeLocationId!: pulumi.Output<string>;
    /**
     * Optional. Indicates whether OSS Redis AUTH is enabled for the
     * instance. If set to "true" AUTH is enabled on the instance.
     * Default value is "false" meaning AUTH is disabled.
     */
    public readonly authEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * AUTH String set on the instance. This field will only be populated if auth_enabled is true.
     */
    public /*out*/ readonly authString!: pulumi.Output<string>;
    /**
     * The full name of the Google Compute Engine network to which the
     * instance is connected. If left unspecified, the default network
     * will be used.
     */
    public readonly authorizedNetwork!: pulumi.Output<string>;
    /**
     * The connection mode of the Redis instance.
     * Default value is `DIRECT_PEERING`.
     * Possible values are: `DIRECT_PEERING`, `PRIVATE_SERVICE_ACCESS`.
     */
    public readonly connectMode!: pulumi.Output<string | undefined>;
    /**
     * (Output)
     * Output only. The time when the policy was created.
     * A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
     * resolution and up to nine fractional digits.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The current zone where the Redis endpoint is placed.
     * For Basic Tier instances, this will always be the same as the
     * [locationId] provided by the user at creation time. For Standard Tier
     * instances, this can be either [locationId] or [alternativeLocationId]
     * and can change after a failover event.
     */
    public /*out*/ readonly currentLocationId!: pulumi.Output<string>;
    /**
     * Optional. The KMS key reference that you want to use to encrypt the data at rest for this Redis
     * instance. If this is provided, CMEK is enabled.
     */
    public readonly customerManagedKey!: pulumi.Output<string | undefined>;
    /**
     * An arbitrary and optional user-provided name for the instance.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     */
    public /*out*/ readonly effectiveLabels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Hostname or IP address of the exposed Redis endpoint used by clients
     * to connect to the service.
     */
    public /*out*/ readonly host!: pulumi.Output<string>;
    /**
     * Resource labels to represent user provided metadata.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The zone where the instance will be provisioned. If not provided,
     * the service will choose a zone for the instance. For STANDARD_HA tier,
     * instances will be created across two zones for protection against
     * zonal failures. If [alternativeLocationId] is also provided, it must
     * be different from [locationId].
     */
    public readonly locationId!: pulumi.Output<string>;
    /**
     * Maintenance policy for an instance.
     * Structure is documented below.
     */
    public readonly maintenancePolicy!: pulumi.Output<outputs.redis.InstanceMaintenancePolicy | undefined>;
    /**
     * Upcoming maintenance schedule.
     * Structure is documented below.
     */
    public /*out*/ readonly maintenanceSchedules!: pulumi.Output<outputs.redis.InstanceMaintenanceSchedule[]>;
    /**
     * Redis memory size in GiB.
     *
     *
     * - - -
     */
    public readonly memorySizeGb!: pulumi.Output<number>;
    /**
     * The ID of the instance or a fully qualified identifier for the instance.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Output only. Info per node.
     * Structure is documented below.
     */
    public /*out*/ readonly nodes!: pulumi.Output<outputs.redis.InstanceNode[]>;
    /**
     * Persistence configuration for an instance.
     * Structure is documented below.
     */
    public readonly persistenceConfig!: pulumi.Output<outputs.redis.InstancePersistenceConfig>;
    /**
     * Output only. Cloud IAM identity used by import / export operations
     * to transfer data to/from Cloud Storage. Format is "serviceAccount:".
     * The value may change over time for a given instance so should be
     * checked before each import/export operation.
     */
    public /*out*/ readonly persistenceIamIdentity!: pulumi.Output<string>;
    /**
     * The port number of the exposed Redis endpoint.
     */
    public /*out*/ readonly port!: pulumi.Output<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     */
    public /*out*/ readonly pulumiLabels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Output only. Hostname or IP address of the exposed readonly Redis endpoint. Standard tier only.
     * Targets all healthy replica nodes in instance. Replication is asynchronous and replica nodes
     * will exhibit some lag behind the primary. Write requests must target 'host'.
     */
    public /*out*/ readonly readEndpoint!: pulumi.Output<string>;
    /**
     * Output only. The port number of the exposed readonly redis endpoint. Standard tier only.
     * Write requests should target 'port'.
     */
    public /*out*/ readonly readEndpointPort!: pulumi.Output<number>;
    /**
     * Optional. Read replica mode. Can only be specified when trying to create the instance.
     * If not set, Memorystore Redis backend will default to READ_REPLICAS_DISABLED.
     * - READ_REPLICAS_DISABLED: If disabled, read endpoint will not be provided and the
     * instance cannot scale up or down the number of replicas.
     * - READ_REPLICAS_ENABLED: If enabled, read endpoint will be provided and the instance
     * can scale up and down the number of replicas.
     * Possible values are: `READ_REPLICAS_DISABLED`, `READ_REPLICAS_ENABLED`.
     */
    public readonly readReplicasMode!: pulumi.Output<string>;
    /**
     * Redis configuration parameters, according to http://redis.io/topics/config.
     * Please check Memorystore documentation for the list of supported parameters:
     * https://cloud.google.com/memorystore/docs/redis/reference/rest/v1/projects.locations.instances#Instance.FIELDS.redis_configs
     */
    public readonly redisConfigs!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The version of Redis software. If not provided, latest supported
     * version will be used. Please check the API documentation linked
     * at the top for the latest valid values.
     */
    public readonly redisVersion!: pulumi.Output<string>;
    /**
     * The name of the Redis region of the instance.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * Optional. The number of replica nodes. The valid range for the Standard Tier with
     * read replicas enabled is [1-5] and defaults to 2. If read replicas are not enabled
     * for a Standard Tier instance, the only valid value is 1 and the default is 1.
     * The valid value for basic tier is 0 and the default is also 0.
     */
    public readonly replicaCount!: pulumi.Output<number>;
    /**
     * The CIDR range of internal addresses that are reserved for this
     * instance. If not provided, the service will choose an unused /29
     * block, for example, 10.0.0.0/29 or 192.168.0.0/29. Ranges must be
     * unique and non-overlapping with existing subnets in an authorized
     * network.
     */
    public readonly reservedIpRange!: pulumi.Output<string>;
    /**
     * Optional. Additional IP range for node placement. Required when enabling read replicas on
     * an existing instance. For DIRECT_PEERING mode value must be a CIDR range of size /28, or
     * "auto". For PRIVATE_SERVICE_ACCESS mode value must be the name of an allocated address
     * range associated with the private service access connection, or "auto".
     */
    public readonly secondaryIpRange!: pulumi.Output<string>;
    /**
     * List of server CA certificates for the instance.
     * Structure is documented below.
     */
    public /*out*/ readonly serverCaCerts!: pulumi.Output<outputs.redis.InstanceServerCaCert[]>;
    /**
     * The service tier of the instance. Must be one of these values:
     * - BASIC: standalone instance
     * - STANDARD_HA: highly available primary/replica instances
     * Default value is `BASIC`.
     * Possible values are: `BASIC`, `STANDARD_HA`.
     */
    public readonly tier!: pulumi.Output<string | undefined>;
    /**
     * The TLS mode of the Redis instance, If not provided, TLS is disabled for the instance.
     * - SERVER_AUTHENTICATION: Client to Server traffic encryption enabled with server authentication
     * Default value is `DISABLED`.
     * Possible values are: `SERVER_AUTHENTICATION`, `DISABLED`.
     */
    public readonly transitEncryptionMode!: pulumi.Output<string | undefined>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceArgs | InstanceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceState | undefined;
            resourceInputs["alternativeLocationId"] = state ? state.alternativeLocationId : undefined;
            resourceInputs["authEnabled"] = state ? state.authEnabled : undefined;
            resourceInputs["authString"] = state ? state.authString : undefined;
            resourceInputs["authorizedNetwork"] = state ? state.authorizedNetwork : undefined;
            resourceInputs["connectMode"] = state ? state.connectMode : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["currentLocationId"] = state ? state.currentLocationId : undefined;
            resourceInputs["customerManagedKey"] = state ? state.customerManagedKey : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["effectiveLabels"] = state ? state.effectiveLabels : undefined;
            resourceInputs["host"] = state ? state.host : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["locationId"] = state ? state.locationId : undefined;
            resourceInputs["maintenancePolicy"] = state ? state.maintenancePolicy : undefined;
            resourceInputs["maintenanceSchedules"] = state ? state.maintenanceSchedules : undefined;
            resourceInputs["memorySizeGb"] = state ? state.memorySizeGb : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nodes"] = state ? state.nodes : undefined;
            resourceInputs["persistenceConfig"] = state ? state.persistenceConfig : undefined;
            resourceInputs["persistenceIamIdentity"] = state ? state.persistenceIamIdentity : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["pulumiLabels"] = state ? state.pulumiLabels : undefined;
            resourceInputs["readEndpoint"] = state ? state.readEndpoint : undefined;
            resourceInputs["readEndpointPort"] = state ? state.readEndpointPort : undefined;
            resourceInputs["readReplicasMode"] = state ? state.readReplicasMode : undefined;
            resourceInputs["redisConfigs"] = state ? state.redisConfigs : undefined;
            resourceInputs["redisVersion"] = state ? state.redisVersion : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["replicaCount"] = state ? state.replicaCount : undefined;
            resourceInputs["reservedIpRange"] = state ? state.reservedIpRange : undefined;
            resourceInputs["secondaryIpRange"] = state ? state.secondaryIpRange : undefined;
            resourceInputs["serverCaCerts"] = state ? state.serverCaCerts : undefined;
            resourceInputs["tier"] = state ? state.tier : undefined;
            resourceInputs["transitEncryptionMode"] = state ? state.transitEncryptionMode : undefined;
        } else {
            const args = argsOrState as InstanceArgs | undefined;
            if ((!args || args.memorySizeGb === undefined) && !opts.urn) {
                throw new Error("Missing required property 'memorySizeGb'");
            }
            resourceInputs["alternativeLocationId"] = args ? args.alternativeLocationId : undefined;
            resourceInputs["authEnabled"] = args ? args.authEnabled : undefined;
            resourceInputs["authorizedNetwork"] = args ? args.authorizedNetwork : undefined;
            resourceInputs["connectMode"] = args ? args.connectMode : undefined;
            resourceInputs["customerManagedKey"] = args ? args.customerManagedKey : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["locationId"] = args ? args.locationId : undefined;
            resourceInputs["maintenancePolicy"] = args ? args.maintenancePolicy : undefined;
            resourceInputs["memorySizeGb"] = args ? args.memorySizeGb : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["persistenceConfig"] = args ? args.persistenceConfig : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["readReplicasMode"] = args ? args.readReplicasMode : undefined;
            resourceInputs["redisConfigs"] = args ? args.redisConfigs : undefined;
            resourceInputs["redisVersion"] = args ? args.redisVersion : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["replicaCount"] = args ? args.replicaCount : undefined;
            resourceInputs["reservedIpRange"] = args ? args.reservedIpRange : undefined;
            resourceInputs["secondaryIpRange"] = args ? args.secondaryIpRange : undefined;
            resourceInputs["tier"] = args ? args.tier : undefined;
            resourceInputs["transitEncryptionMode"] = args ? args.transitEncryptionMode : undefined;
            resourceInputs["authString"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["currentLocationId"] = undefined /*out*/;
            resourceInputs["effectiveLabels"] = undefined /*out*/;
            resourceInputs["host"] = undefined /*out*/;
            resourceInputs["maintenanceSchedules"] = undefined /*out*/;
            resourceInputs["nodes"] = undefined /*out*/;
            resourceInputs["persistenceIamIdentity"] = undefined /*out*/;
            resourceInputs["port"] = undefined /*out*/;
            resourceInputs["pulumiLabels"] = undefined /*out*/;
            resourceInputs["readEndpoint"] = undefined /*out*/;
            resourceInputs["readEndpointPort"] = undefined /*out*/;
            resourceInputs["serverCaCerts"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["authString", "effectiveLabels", "pulumiLabels"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Instance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Instance resources.
 */
export interface InstanceState {
    /**
     * Only applicable to STANDARD_HA tier which protects the instance
     * against zonal failures by provisioning it across two zones.
     * If provided, it must be a different zone from the one provided in
     * [locationId].
     */
    alternativeLocationId?: pulumi.Input<string>;
    /**
     * Optional. Indicates whether OSS Redis AUTH is enabled for the
     * instance. If set to "true" AUTH is enabled on the instance.
     * Default value is "false" meaning AUTH is disabled.
     */
    authEnabled?: pulumi.Input<boolean>;
    /**
     * AUTH String set on the instance. This field will only be populated if auth_enabled is true.
     */
    authString?: pulumi.Input<string>;
    /**
     * The full name of the Google Compute Engine network to which the
     * instance is connected. If left unspecified, the default network
     * will be used.
     */
    authorizedNetwork?: pulumi.Input<string>;
    /**
     * The connection mode of the Redis instance.
     * Default value is `DIRECT_PEERING`.
     * Possible values are: `DIRECT_PEERING`, `PRIVATE_SERVICE_ACCESS`.
     */
    connectMode?: pulumi.Input<string>;
    /**
     * (Output)
     * Output only. The time when the policy was created.
     * A timestamp in RFC3339 UTC "Zulu" format, with nanosecond
     * resolution and up to nine fractional digits.
     */
    createTime?: pulumi.Input<string>;
    /**
     * The current zone where the Redis endpoint is placed.
     * For Basic Tier instances, this will always be the same as the
     * [locationId] provided by the user at creation time. For Standard Tier
     * instances, this can be either [locationId] or [alternativeLocationId]
     * and can change after a failover event.
     */
    currentLocationId?: pulumi.Input<string>;
    /**
     * Optional. The KMS key reference that you want to use to encrypt the data at rest for this Redis
     * instance. If this is provided, CMEK is enabled.
     */
    customerManagedKey?: pulumi.Input<string>;
    /**
     * An arbitrary and optional user-provided name for the instance.
     */
    displayName?: pulumi.Input<string>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     */
    effectiveLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Hostname or IP address of the exposed Redis endpoint used by clients
     * to connect to the service.
     */
    host?: pulumi.Input<string>;
    /**
     * Resource labels to represent user provided metadata.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The zone where the instance will be provisioned. If not provided,
     * the service will choose a zone for the instance. For STANDARD_HA tier,
     * instances will be created across two zones for protection against
     * zonal failures. If [alternativeLocationId] is also provided, it must
     * be different from [locationId].
     */
    locationId?: pulumi.Input<string>;
    /**
     * Maintenance policy for an instance.
     * Structure is documented below.
     */
    maintenancePolicy?: pulumi.Input<inputs.redis.InstanceMaintenancePolicy>;
    /**
     * Upcoming maintenance schedule.
     * Structure is documented below.
     */
    maintenanceSchedules?: pulumi.Input<pulumi.Input<inputs.redis.InstanceMaintenanceSchedule>[]>;
    /**
     * Redis memory size in GiB.
     *
     *
     * - - -
     */
    memorySizeGb?: pulumi.Input<number>;
    /**
     * The ID of the instance or a fully qualified identifier for the instance.
     */
    name?: pulumi.Input<string>;
    /**
     * Output only. Info per node.
     * Structure is documented below.
     */
    nodes?: pulumi.Input<pulumi.Input<inputs.redis.InstanceNode>[]>;
    /**
     * Persistence configuration for an instance.
     * Structure is documented below.
     */
    persistenceConfig?: pulumi.Input<inputs.redis.InstancePersistenceConfig>;
    /**
     * Output only. Cloud IAM identity used by import / export operations
     * to transfer data to/from Cloud Storage. Format is "serviceAccount:".
     * The value may change over time for a given instance so should be
     * checked before each import/export operation.
     */
    persistenceIamIdentity?: pulumi.Input<string>;
    /**
     * The port number of the exposed Redis endpoint.
     */
    port?: pulumi.Input<number>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     */
    pulumiLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Output only. Hostname or IP address of the exposed readonly Redis endpoint. Standard tier only.
     * Targets all healthy replica nodes in instance. Replication is asynchronous and replica nodes
     * will exhibit some lag behind the primary. Write requests must target 'host'.
     */
    readEndpoint?: pulumi.Input<string>;
    /**
     * Output only. The port number of the exposed readonly redis endpoint. Standard tier only.
     * Write requests should target 'port'.
     */
    readEndpointPort?: pulumi.Input<number>;
    /**
     * Optional. Read replica mode. Can only be specified when trying to create the instance.
     * If not set, Memorystore Redis backend will default to READ_REPLICAS_DISABLED.
     * - READ_REPLICAS_DISABLED: If disabled, read endpoint will not be provided and the
     * instance cannot scale up or down the number of replicas.
     * - READ_REPLICAS_ENABLED: If enabled, read endpoint will be provided and the instance
     * can scale up and down the number of replicas.
     * Possible values are: `READ_REPLICAS_DISABLED`, `READ_REPLICAS_ENABLED`.
     */
    readReplicasMode?: pulumi.Input<string>;
    /**
     * Redis configuration parameters, according to http://redis.io/topics/config.
     * Please check Memorystore documentation for the list of supported parameters:
     * https://cloud.google.com/memorystore/docs/redis/reference/rest/v1/projects.locations.instances#Instance.FIELDS.redis_configs
     */
    redisConfigs?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The version of Redis software. If not provided, latest supported
     * version will be used. Please check the API documentation linked
     * at the top for the latest valid values.
     */
    redisVersion?: pulumi.Input<string>;
    /**
     * The name of the Redis region of the instance.
     */
    region?: pulumi.Input<string>;
    /**
     * Optional. The number of replica nodes. The valid range for the Standard Tier with
     * read replicas enabled is [1-5] and defaults to 2. If read replicas are not enabled
     * for a Standard Tier instance, the only valid value is 1 and the default is 1.
     * The valid value for basic tier is 0 and the default is also 0.
     */
    replicaCount?: pulumi.Input<number>;
    /**
     * The CIDR range of internal addresses that are reserved for this
     * instance. If not provided, the service will choose an unused /29
     * block, for example, 10.0.0.0/29 or 192.168.0.0/29. Ranges must be
     * unique and non-overlapping with existing subnets in an authorized
     * network.
     */
    reservedIpRange?: pulumi.Input<string>;
    /**
     * Optional. Additional IP range for node placement. Required when enabling read replicas on
     * an existing instance. For DIRECT_PEERING mode value must be a CIDR range of size /28, or
     * "auto". For PRIVATE_SERVICE_ACCESS mode value must be the name of an allocated address
     * range associated with the private service access connection, or "auto".
     */
    secondaryIpRange?: pulumi.Input<string>;
    /**
     * List of server CA certificates for the instance.
     * Structure is documented below.
     */
    serverCaCerts?: pulumi.Input<pulumi.Input<inputs.redis.InstanceServerCaCert>[]>;
    /**
     * The service tier of the instance. Must be one of these values:
     * - BASIC: standalone instance
     * - STANDARD_HA: highly available primary/replica instances
     * Default value is `BASIC`.
     * Possible values are: `BASIC`, `STANDARD_HA`.
     */
    tier?: pulumi.Input<string>;
    /**
     * The TLS mode of the Redis instance, If not provided, TLS is disabled for the instance.
     * - SERVER_AUTHENTICATION: Client to Server traffic encryption enabled with server authentication
     * Default value is `DISABLED`.
     * Possible values are: `SERVER_AUTHENTICATION`, `DISABLED`.
     */
    transitEncryptionMode?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * Only applicable to STANDARD_HA tier which protects the instance
     * against zonal failures by provisioning it across two zones.
     * If provided, it must be a different zone from the one provided in
     * [locationId].
     */
    alternativeLocationId?: pulumi.Input<string>;
    /**
     * Optional. Indicates whether OSS Redis AUTH is enabled for the
     * instance. If set to "true" AUTH is enabled on the instance.
     * Default value is "false" meaning AUTH is disabled.
     */
    authEnabled?: pulumi.Input<boolean>;
    /**
     * The full name of the Google Compute Engine network to which the
     * instance is connected. If left unspecified, the default network
     * will be used.
     */
    authorizedNetwork?: pulumi.Input<string>;
    /**
     * The connection mode of the Redis instance.
     * Default value is `DIRECT_PEERING`.
     * Possible values are: `DIRECT_PEERING`, `PRIVATE_SERVICE_ACCESS`.
     */
    connectMode?: pulumi.Input<string>;
    /**
     * Optional. The KMS key reference that you want to use to encrypt the data at rest for this Redis
     * instance. If this is provided, CMEK is enabled.
     */
    customerManagedKey?: pulumi.Input<string>;
    /**
     * An arbitrary and optional user-provided name for the instance.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Resource labels to represent user provided metadata.
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The zone where the instance will be provisioned. If not provided,
     * the service will choose a zone for the instance. For STANDARD_HA tier,
     * instances will be created across two zones for protection against
     * zonal failures. If [alternativeLocationId] is also provided, it must
     * be different from [locationId].
     */
    locationId?: pulumi.Input<string>;
    /**
     * Maintenance policy for an instance.
     * Structure is documented below.
     */
    maintenancePolicy?: pulumi.Input<inputs.redis.InstanceMaintenancePolicy>;
    /**
     * Redis memory size in GiB.
     *
     *
     * - - -
     */
    memorySizeGb: pulumi.Input<number>;
    /**
     * The ID of the instance or a fully qualified identifier for the instance.
     */
    name?: pulumi.Input<string>;
    /**
     * Persistence configuration for an instance.
     * Structure is documented below.
     */
    persistenceConfig?: pulumi.Input<inputs.redis.InstancePersistenceConfig>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Optional. Read replica mode. Can only be specified when trying to create the instance.
     * If not set, Memorystore Redis backend will default to READ_REPLICAS_DISABLED.
     * - READ_REPLICAS_DISABLED: If disabled, read endpoint will not be provided and the
     * instance cannot scale up or down the number of replicas.
     * - READ_REPLICAS_ENABLED: If enabled, read endpoint will be provided and the instance
     * can scale up and down the number of replicas.
     * Possible values are: `READ_REPLICAS_DISABLED`, `READ_REPLICAS_ENABLED`.
     */
    readReplicasMode?: pulumi.Input<string>;
    /**
     * Redis configuration parameters, according to http://redis.io/topics/config.
     * Please check Memorystore documentation for the list of supported parameters:
     * https://cloud.google.com/memorystore/docs/redis/reference/rest/v1/projects.locations.instances#Instance.FIELDS.redis_configs
     */
    redisConfigs?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The version of Redis software. If not provided, latest supported
     * version will be used. Please check the API documentation linked
     * at the top for the latest valid values.
     */
    redisVersion?: pulumi.Input<string>;
    /**
     * The name of the Redis region of the instance.
     */
    region?: pulumi.Input<string>;
    /**
     * Optional. The number of replica nodes. The valid range for the Standard Tier with
     * read replicas enabled is [1-5] and defaults to 2. If read replicas are not enabled
     * for a Standard Tier instance, the only valid value is 1 and the default is 1.
     * The valid value for basic tier is 0 and the default is also 0.
     */
    replicaCount?: pulumi.Input<number>;
    /**
     * The CIDR range of internal addresses that are reserved for this
     * instance. If not provided, the service will choose an unused /29
     * block, for example, 10.0.0.0/29 or 192.168.0.0/29. Ranges must be
     * unique and non-overlapping with existing subnets in an authorized
     * network.
     */
    reservedIpRange?: pulumi.Input<string>;
    /**
     * Optional. Additional IP range for node placement. Required when enabling read replicas on
     * an existing instance. For DIRECT_PEERING mode value must be a CIDR range of size /28, or
     * "auto". For PRIVATE_SERVICE_ACCESS mode value must be the name of an allocated address
     * range associated with the private service access connection, or "auto".
     */
    secondaryIpRange?: pulumi.Input<string>;
    /**
     * The service tier of the instance. Must be one of these values:
     * - BASIC: standalone instance
     * - STANDARD_HA: highly available primary/replica instances
     * Default value is `BASIC`.
     * Possible values are: `BASIC`, `STANDARD_HA`.
     */
    tier?: pulumi.Input<string>;
    /**
     * The TLS mode of the Redis instance, If not provided, TLS is disabled for the instance.
     * - SERVER_AUTHENTICATION: Client to Server traffic encryption enabled with server authentication
     * Default value is `DISABLED`.
     * Possible values are: `SERVER_AUTHENTICATION`, `DISABLED`.
     */
    transitEncryptionMode?: pulumi.Input<string>;
}
