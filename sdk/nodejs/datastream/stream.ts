// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A resource representing streaming data from a source to a destination.
 *
 * To get more information about Stream, see:
 *
 * * [API documentation](https://cloud.google.com/datastream/docs/reference/rest/v1/projects.locations.streams)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/datastream/docs/create-a-stream)
 *
 * ## Example Usage
 * ### Datastream Stream Full
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * import * as random from "@pulumi/random";
 *
 * const project = gcp.organizations.getProject({});
 * const instance = new gcp.sql.DatabaseInstance("instance", {
 *     databaseVersion: "MYSQL_8_0",
 *     region: "us-central1",
 *     settings: {
 *         tier: "db-f1-micro",
 *         backupConfiguration: {
 *             enabled: true,
 *             binaryLogEnabled: true,
 *         },
 *         ipConfiguration: {
 *             authorizedNetworks: [
 *                 {
 *                     value: "34.71.242.81",
 *                 },
 *                 {
 *                     value: "34.72.28.29",
 *                 },
 *                 {
 *                     value: "34.67.6.157",
 *                 },
 *                 {
 *                     value: "34.67.234.134",
 *                 },
 *                 {
 *                     value: "34.72.239.218",
 *                 },
 *             ],
 *         },
 *     },
 *     deletionProtection: true,
 * });
 * const db = new gcp.sql.Database("db", {instance: instance.name});
 * const pwd = new random.RandomPassword("pwd", {
 *     length: 16,
 *     special: false,
 * });
 * const user = new gcp.sql.User("user", {
 *     instance: instance.name,
 *     host: "%",
 *     password: pwd.result,
 * });
 * const sourceConnectionProfile = new gcp.datastream.ConnectionProfile("sourceConnectionProfile", {
 *     displayName: "Source connection profile",
 *     location: "us-central1",
 *     connectionProfileId: "source-profile",
 *     mysqlProfile: {
 *         hostname: instance.publicIpAddress,
 *         username: user.name,
 *         password: user.password,
 *     },
 * });
 * const bucket = new gcp.storage.Bucket("bucket", {
 *     location: "US",
 *     uniformBucketLevelAccess: true,
 * });
 * const viewer = new gcp.storage.BucketIAMMember("viewer", {
 *     bucket: bucket.name,
 *     role: "roles/storage.objectViewer",
 *     member: project.then(project => `serviceAccount:service-${project.number}@gcp-sa-datastream.iam.gserviceaccount.com`),
 * });
 * const creator = new gcp.storage.BucketIAMMember("creator", {
 *     bucket: bucket.name,
 *     role: "roles/storage.objectCreator",
 *     member: project.then(project => `serviceAccount:service-${project.number}@gcp-sa-datastream.iam.gserviceaccount.com`),
 * });
 * const reader = new gcp.storage.BucketIAMMember("reader", {
 *     bucket: bucket.name,
 *     role: "roles/storage.legacyBucketReader",
 *     member: project.then(project => `serviceAccount:service-${project.number}@gcp-sa-datastream.iam.gserviceaccount.com`),
 * });
 * const keyUser = new gcp.kms.CryptoKeyIAMMember("keyUser", {
 *     cryptoKeyId: "kms-name",
 *     role: "roles/cloudkms.cryptoKeyEncrypterDecrypter",
 *     member: project.then(project => `serviceAccount:service-${project.number}@gcp-sa-datastream.iam.gserviceaccount.com`),
 * });
 * const destinationConnectionProfile = new gcp.datastream.ConnectionProfile("destinationConnectionProfile", {
 *     displayName: "Connection profile",
 *     location: "us-central1",
 *     connectionProfileId: "destination-profile",
 *     gcsProfile: {
 *         bucket: bucket.name,
 *         rootPath: "/path",
 *     },
 * });
 * const _default = new gcp.datastream.Stream("default", {
 *     streamId: "my-stream",
 *     desiredState: "NOT_STARTED",
 *     location: "us-central1",
 *     displayName: "my stream",
 *     labels: {
 *         key: "value",
 *     },
 *     sourceConfig: {
 *         sourceConnectionProfile: sourceConnectionProfile.id,
 *         mysqlSourceConfig: {
 *             includeObjects: {
 *                 mysqlDatabases: [{
 *                     database: "my-database",
 *                     mysqlTables: [{
 *                         table: "includedTable",
 *                         mysqlColumns: [{
 *                             column: "includedColumn",
 *                             dataType: "VARCHAR",
 *                             collation: "utf8mb4",
 *                             primaryKey: false,
 *                             nullable: false,
 *                             ordinalPosition: 0,
 *                         }],
 *                     }],
 *                 }],
 *             },
 *             excludeObjects: {
 *                 mysqlDatabases: [{
 *                     database: "my-database",
 *                     mysqlTables: [{
 *                         table: "excludedTable",
 *                         mysqlColumns: [{
 *                             column: "excludedColumn",
 *                             dataType: "VARCHAR",
 *                             collation: "utf8mb4",
 *                             primaryKey: false,
 *                             nullable: false,
 *                             ordinalPosition: 0,
 *                         }],
 *                     }],
 *                 }],
 *             },
 *             maxConcurrentCdcTasks: 5,
 *         },
 *     },
 *     destinationConfig: {
 *         destinationConnectionProfile: destinationConnectionProfile.id,
 *         gcsDestinationConfig: {
 *             path: "mydata",
 *             fileRotationMb: 200,
 *             fileRotationInterval: "900s",
 *             jsonFileFormat: {
 *                 schemaFileFormat: "NO_SCHEMA_FILE",
 *                 compression: "GZIP",
 *             },
 *         },
 *     },
 *     backfillAll: {
 *         mysqlExcludedObjects: {
 *             mysqlDatabases: [{
 *                 database: "my-database",
 *                 mysqlTables: [{
 *                     table: "excludedTable",
 *                     mysqlColumns: [{
 *                         column: "excludedColumn",
 *                         dataType: "VARCHAR",
 *                         collation: "utf8mb4",
 *                         primaryKey: false,
 *                         nullable: false,
 *                         ordinalPosition: 0,
 *                     }],
 *                 }],
 *             }],
 *         },
 *     },
 *     customerManagedEncryptionKey: "kms-name",
 * }, {
 *     dependsOn: [keyUser],
 * });
 * ```
 * ### Datastream Stream Bigquery
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 * import * as random from "@pulumi/random";
 *
 * const project = gcp.organizations.getProject({});
 * const instance = new gcp.sql.DatabaseInstance("instance", {
 *     databaseVersion: "MYSQL_8_0",
 *     region: "us-central1",
 *     settings: {
 *         tier: "db-f1-micro",
 *         backupConfiguration: {
 *             enabled: true,
 *             binaryLogEnabled: true,
 *         },
 *         ipConfiguration: {
 *             authorizedNetworks: [
 *                 {
 *                     value: "34.71.242.81",
 *                 },
 *                 {
 *                     value: "34.72.28.29",
 *                 },
 *                 {
 *                     value: "34.67.6.157",
 *                 },
 *                 {
 *                     value: "34.67.234.134",
 *                 },
 *                 {
 *                     value: "34.72.239.218",
 *                 },
 *             ],
 *         },
 *     },
 *     deletionProtection: true,
 * });
 * const db = new gcp.sql.Database("db", {instance: instance.name});
 * const pwd = new random.RandomPassword("pwd", {
 *     length: 16,
 *     special: false,
 * });
 * const user = new gcp.sql.User("user", {
 *     instance: instance.name,
 *     host: "%",
 *     password: pwd.result,
 * });
 * const sourceConnectionProfile = new gcp.datastream.ConnectionProfile("sourceConnectionProfile", {
 *     displayName: "Source connection profile",
 *     location: "us-central1",
 *     connectionProfileId: "source-profile",
 *     mysqlProfile: {
 *         hostname: instance.publicIpAddress,
 *         username: user.name,
 *         password: user.password,
 *     },
 * });
 * const bqSa = gcp.bigquery.getDefaultServiceAccount({});
 * const bigqueryKeyUser = new gcp.kms.CryptoKeyIAMMember("bigqueryKeyUser", {
 *     cryptoKeyId: "bigquery-kms-name",
 *     role: "roles/cloudkms.cryptoKeyEncrypterDecrypter",
 *     member: bqSa.then(bqSa => `serviceAccount:${bqSa.email}`),
 * });
 * const destinationConnectionProfile = new gcp.datastream.ConnectionProfile("destinationConnectionProfile", {
 *     displayName: "Connection profile",
 *     location: "us-central1",
 *     connectionProfileId: "destination-profile",
 *     bigqueryProfile: {},
 * });
 * const _default = new gcp.datastream.Stream("default", {
 *     streamId: "my-stream",
 *     location: "us-central1",
 *     displayName: "my stream",
 *     sourceConfig: {
 *         sourceConnectionProfile: sourceConnectionProfile.id,
 *         mysqlSourceConfig: {},
 *     },
 *     destinationConfig: {
 *         destinationConnectionProfile: destinationConnectionProfile.id,
 *         bigqueryDestinationConfig: {
 *             sourceHierarchyDatasets: {
 *                 datasetTemplate: {
 *                     location: "us-central1",
 *                     kmsKeyName: "bigquery-kms-name",
 *                 },
 *             },
 *         },
 *     },
 *     backfillNone: {},
 * }, {
 *     dependsOn: [bigqueryKeyUser],
 * });
 * ```
 *
 * ## Import
 *
 * Stream can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:datastream/stream:Stream default projects/{{project}}/locations/{{location}}/streams/{{stream_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:datastream/stream:Stream default {{project}}/{{location}}/{{stream_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:datastream/stream:Stream default {{location}}/{{stream_id}}
 * ```
 */
export class Stream extends pulumi.CustomResource {
    /**
     * Get an existing Stream resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StreamState, opts?: pulumi.CustomResourceOptions): Stream {
        return new Stream(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:datastream/stream:Stream';

    /**
     * Returns true if the given object is an instance of Stream.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Stream {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Stream.__pulumiType;
    }

    /**
     * Backfill strategy to automatically backfill the Stream's objects. Specific objects can be excluded.
     * Structure is documented below.
     */
    public readonly backfillAll!: pulumi.Output<outputs.datastream.StreamBackfillAll | undefined>;
    /**
     * Backfill strategy to disable automatic backfill for the Stream's objects.
     */
    public readonly backfillNone!: pulumi.Output<outputs.datastream.StreamBackfillNone | undefined>;
    /**
     * A reference to a KMS encryption key. If provided, it will be used to encrypt the data. If left blank, data
     * will be encrypted using an internal Stream-specific encryption key provisioned through KMS.
     */
    public readonly customerManagedEncryptionKey!: pulumi.Output<string | undefined>;
    /**
     * Desired state of the Stream. Set this field to `RUNNING` to start the stream, and `PAUSED` to pause the stream.
     */
    public readonly desiredState!: pulumi.Output<string | undefined>;
    /**
     * Destination connection profile configuration.
     * Structure is documented below.
     */
    public readonly destinationConfig!: pulumi.Output<outputs.datastream.StreamDestinationConfig>;
    /**
     * Display name.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Labels.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The name of the location this stream is located in.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The stream's name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Source connection profile configuration.
     * Structure is documented below.
     */
    public readonly sourceConfig!: pulumi.Output<outputs.datastream.StreamSourceConfig>;
    /**
     * The state of the stream.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The stream identifier.
     */
    public readonly streamId!: pulumi.Output<string>;

    /**
     * Create a Stream resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StreamArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StreamArgs | StreamState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StreamState | undefined;
            resourceInputs["backfillAll"] = state ? state.backfillAll : undefined;
            resourceInputs["backfillNone"] = state ? state.backfillNone : undefined;
            resourceInputs["customerManagedEncryptionKey"] = state ? state.customerManagedEncryptionKey : undefined;
            resourceInputs["desiredState"] = state ? state.desiredState : undefined;
            resourceInputs["destinationConfig"] = state ? state.destinationConfig : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["sourceConfig"] = state ? state.sourceConfig : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["streamId"] = state ? state.streamId : undefined;
        } else {
            const args = argsOrState as StreamArgs | undefined;
            if ((!args || args.destinationConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destinationConfig'");
            }
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.sourceConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceConfig'");
            }
            if ((!args || args.streamId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'streamId'");
            }
            resourceInputs["backfillAll"] = args ? args.backfillAll : undefined;
            resourceInputs["backfillNone"] = args ? args.backfillNone : undefined;
            resourceInputs["customerManagedEncryptionKey"] = args ? args.customerManagedEncryptionKey : undefined;
            resourceInputs["desiredState"] = args ? args.desiredState : undefined;
            resourceInputs["destinationConfig"] = args ? args.destinationConfig : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["sourceConfig"] = args ? args.sourceConfig : undefined;
            resourceInputs["streamId"] = args ? args.streamId : undefined;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Stream.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Stream resources.
 */
export interface StreamState {
    /**
     * Backfill strategy to automatically backfill the Stream's objects. Specific objects can be excluded.
     * Structure is documented below.
     */
    backfillAll?: pulumi.Input<inputs.datastream.StreamBackfillAll>;
    /**
     * Backfill strategy to disable automatic backfill for the Stream's objects.
     */
    backfillNone?: pulumi.Input<inputs.datastream.StreamBackfillNone>;
    /**
     * A reference to a KMS encryption key. If provided, it will be used to encrypt the data. If left blank, data
     * will be encrypted using an internal Stream-specific encryption key provisioned through KMS.
     */
    customerManagedEncryptionKey?: pulumi.Input<string>;
    /**
     * Desired state of the Stream. Set this field to `RUNNING` to start the stream, and `PAUSED` to pause the stream.
     */
    desiredState?: pulumi.Input<string>;
    /**
     * Destination connection profile configuration.
     * Structure is documented below.
     */
    destinationConfig?: pulumi.Input<inputs.datastream.StreamDestinationConfig>;
    /**
     * Display name.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Labels.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the location this stream is located in.
     */
    location?: pulumi.Input<string>;
    /**
     * The stream's name.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Source connection profile configuration.
     * Structure is documented below.
     */
    sourceConfig?: pulumi.Input<inputs.datastream.StreamSourceConfig>;
    /**
     * The state of the stream.
     */
    state?: pulumi.Input<string>;
    /**
     * The stream identifier.
     */
    streamId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Stream resource.
 */
export interface StreamArgs {
    /**
     * Backfill strategy to automatically backfill the Stream's objects. Specific objects can be excluded.
     * Structure is documented below.
     */
    backfillAll?: pulumi.Input<inputs.datastream.StreamBackfillAll>;
    /**
     * Backfill strategy to disable automatic backfill for the Stream's objects.
     */
    backfillNone?: pulumi.Input<inputs.datastream.StreamBackfillNone>;
    /**
     * A reference to a KMS encryption key. If provided, it will be used to encrypt the data. If left blank, data
     * will be encrypted using an internal Stream-specific encryption key provisioned through KMS.
     */
    customerManagedEncryptionKey?: pulumi.Input<string>;
    /**
     * Desired state of the Stream. Set this field to `RUNNING` to start the stream, and `PAUSED` to pause the stream.
     */
    desiredState?: pulumi.Input<string>;
    /**
     * Destination connection profile configuration.
     * Structure is documented below.
     */
    destinationConfig: pulumi.Input<inputs.datastream.StreamDestinationConfig>;
    /**
     * Display name.
     */
    displayName: pulumi.Input<string>;
    /**
     * Labels.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the location this stream is located in.
     */
    location: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Source connection profile configuration.
     * Structure is documented below.
     */
    sourceConfig: pulumi.Input<inputs.datastream.StreamSourceConfig>;
    /**
     * The stream identifier.
     */
    streamId: pulumi.Input<string>;
}
