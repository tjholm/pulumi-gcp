// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 * ### Bigquery Dataset Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const bqowner = new gcp.serviceaccount.Account("bqowner", {accountId: "bqowner"});
 * const dataset = new gcp.bigquery.Dataset("dataset", {
 *     datasetId: "example_dataset",
 *     friendlyName: "test",
 *     description: "This is a test description",
 *     location: "EU",
 *     defaultTableExpirationMs: 3600000,
 *     labels: {
 *         env: "default",
 *     },
 *     accesses: [
 *         {
 *             role: "OWNER",
 *             userByEmail: bqowner.email,
 *         },
 *         {
 *             role: "READER",
 *             domain: "hashicorp.com",
 *         },
 *     ],
 * });
 * ```
 * ### Bigquery Dataset Cmek
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const keyRing = new gcp.kms.KeyRing("keyRing", {location: "us"});
 * const cryptoKey = new gcp.kms.CryptoKey("cryptoKey", {keyRing: keyRing.id});
 * const dataset = new gcp.bigquery.Dataset("dataset", {
 *     datasetId: "example_dataset",
 *     friendlyName: "test",
 *     description: "This is a test description",
 *     location: "US",
 *     defaultTableExpirationMs: 3600000,
 *     defaultEncryptionConfiguration: {
 *         kmsKeyName: cryptoKey.id,
 *     },
 * });
 * ```
 * ### Bigquery Dataset Authorized Dataset
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const bqowner = new gcp.serviceaccount.Account("bqowner", {accountId: "bqowner"});
 * const _public = new gcp.bigquery.Dataset("public", {
 *     datasetId: "public",
 *     friendlyName: "test",
 *     description: "This dataset is public",
 *     location: "EU",
 *     defaultTableExpirationMs: 3600000,
 *     labels: {
 *         env: "default",
 *     },
 *     accesses: [
 *         {
 *             role: "OWNER",
 *             userByEmail: bqowner.email,
 *         },
 *         {
 *             role: "READER",
 *             domain: "hashicorp.com",
 *         },
 *     ],
 * });
 * const dataset = new gcp.bigquery.Dataset("dataset", {
 *     datasetId: "private",
 *     friendlyName: "test",
 *     description: "This dataset is private",
 *     location: "EU",
 *     defaultTableExpirationMs: 3600000,
 *     labels: {
 *         env: "default",
 *     },
 *     accesses: [
 *         {
 *             role: "OWNER",
 *             userByEmail: bqowner.email,
 *         },
 *         {
 *             role: "READER",
 *             domain: "hashicorp.com",
 *         },
 *         {
 *             dataset: {
 *                 dataset: {
 *                     projectId: _public.project,
 *                     datasetId: _public.datasetId,
 *                 },
 *                 targetTypes: ["VIEWS"],
 *             },
 *         },
 *     ],
 * });
 * ```
 * ### Bigquery Dataset Authorized Routine
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const publicDataset = new gcp.bigquery.Dataset("publicDataset", {
 *     datasetId: "public_dataset",
 *     description: "This dataset is public",
 * });
 * const publicRoutine = new gcp.bigquery.Routine("publicRoutine", {
 *     datasetId: publicDataset.datasetId,
 *     routineId: "public_routine",
 *     routineType: "TABLE_VALUED_FUNCTION",
 *     language: "SQL",
 *     definitionBody: "SELECT 1 + value AS value\n",
 *     arguments: [{
 *         name: "value",
 *         argumentKind: "FIXED_TYPE",
 *         dataType: JSON.stringify({
 *             typeKind: "INT64",
 *         }),
 *     }],
 *     returnTableType: JSON.stringify({
 *         columns: [{
 *             name: "value",
 *             type: {
 *                 typeKind: "INT64",
 *             },
 *         }],
 *     }),
 * });
 * const _private = new gcp.bigquery.Dataset("private", {
 *     datasetId: "private_dataset",
 *     description: "This dataset is private",
 *     accesses: [
 *         {
 *             role: "OWNER",
 *             userByEmail: "my@service-account.com",
 *         },
 *         {
 *             routine: {
 *                 projectId: publicRoutine.project,
 *                 datasetId: publicRoutine.datasetId,
 *                 routineId: publicRoutine.routineId,
 *             },
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Dataset can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:bigquery/dataset:Dataset default projects/{{project}}/datasets/{{dataset_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:bigquery/dataset:Dataset default {{project}}/{{dataset_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:bigquery/dataset:Dataset default {{dataset_id}}
 * ```
 */
export class Dataset extends pulumi.CustomResource {
    /**
     * Get an existing Dataset resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatasetState, opts?: pulumi.CustomResourceOptions): Dataset {
        return new Dataset(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:bigquery/dataset:Dataset';

    /**
     * Returns true if the given object is an instance of Dataset.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Dataset {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Dataset.__pulumiType;
    }

    /**
     * An array of objects that define dataset access for one or more entities.
     * Structure is documented below.
     */
    public readonly accesses!: pulumi.Output<outputs.bigquery.DatasetAccess[]>;
    /**
     * The time when this dataset was created, in milliseconds since the
     * epoch.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<number>;
    /**
     * A unique ID for this dataset, without the project name. The ID
     * must contain only letters (a-z, A-Z), numbers (0-9), or
     * underscores (_). The maximum length is 1,024 characters.
     *
     *
     * - - -
     */
    public readonly datasetId!: pulumi.Output<string>;
    /**
     * Defines the default collation specification of future tables created
     * in the dataset. If a table is created in this dataset without table-level
     * default collation, then the table inherits the dataset default collation,
     * which is applied to the string fields that do not have explicit collation
     * specified. A change to this field affects only tables created afterwards,
     * and does not alter the existing tables.
     * The following values are supported:
     * - 'und:ci': undetermined locale, case insensitive.
     * - '': empty string. Default to case-sensitive behavior.
     */
    public readonly defaultCollation!: pulumi.Output<string>;
    /**
     * The default encryption key for all tables in the dataset. Once this property is set,
     * all newly-created partitioned tables in the dataset will have encryption key set to
     * this value, unless table creation request (or query) overrides the key.
     * Structure is documented below.
     */
    public readonly defaultEncryptionConfiguration!: pulumi.Output<outputs.bigquery.DatasetDefaultEncryptionConfiguration | undefined>;
    /**
     * The default partition expiration for all partitioned tables in
     * the dataset, in milliseconds.
     *
     * Once this property is set, all newly-created partitioned tables in
     * the dataset will have an `expirationMs` property in the `timePartitioning`
     * settings set to this value, and changing the value will only
     * affect new tables, not existing ones. The storage in a partition will
     * have an expiration time of its partition time plus this value.
     * Setting this property overrides the use of `defaultTableExpirationMs`
     * for partitioned tables: only one of `defaultTableExpirationMs` and
     * `defaultPartitionExpirationMs` will be used for any new partitioned
     * table. If you provide an explicit `timePartitioning.expirationMs` when
     * creating or updating a partitioned table, that value takes precedence
     * over the default partition expiration time indicated by this property.
     */
    public readonly defaultPartitionExpirationMs!: pulumi.Output<number | undefined>;
    /**
     * The default lifetime of all tables in the dataset, in milliseconds.
     * The minimum value is 3600000 milliseconds (one hour).
     *
     * Once this property is set, all newly-created tables in the dataset
     * will have an `expirationTime` property set to the creation time plus
     * the value in this property, and changing the value will only affect
     * new tables, not existing ones. When the `expirationTime` for a given
     * table is reached, that table will be deleted automatically.
     * If a table's `expirationTime` is modified or removed before the
     * table expires, or if you provide an explicit `expirationTime` when
     * creating a table, that value takes precedence over the default
     * expiration time indicated by this property.
     */
    public readonly defaultTableExpirationMs!: pulumi.Output<number | undefined>;
    /**
     * If set to `true`, delete all the tables in the
     * dataset when destroying the resource; otherwise,
     * destroying the resource will fail if tables are present.
     */
    public readonly deleteContentsOnDestroy!: pulumi.Output<boolean | undefined>;
    /**
     * A user-friendly description of the dataset
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     */
    public /*out*/ readonly effectiveLabels!: pulumi.Output<{[key: string]: string}>;
    /**
     * A hash of the resource.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * A descriptive name for the dataset
     */
    public readonly friendlyName!: pulumi.Output<string | undefined>;
    /**
     * TRUE if the dataset and its table names are case-insensitive, otherwise FALSE.
     * By default, this is FALSE, which means the dataset and its table names are
     * case-sensitive. This field does not affect routine references.
     */
    public readonly isCaseInsensitive!: pulumi.Output<boolean>;
    /**
     * The labels associated with this dataset. You can use these to
     * organize and group your datasets.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The date when this dataset or any of its tables was last modified, in
     * milliseconds since the epoch.
     */
    public /*out*/ readonly lastModifiedTime!: pulumi.Output<number>;
    /**
     * The geographic location where the dataset should reside.
     * See [official docs](https://cloud.google.com/bigquery/docs/dataset-locations).
     *
     * There are two types of locations, regional or multi-regional. A regional
     * location is a specific geographic place, such as Tokyo, and a multi-regional
     * location is a large geographic area, such as the United States, that
     * contains at least two geographic places.
     *
     * The default value is multi-regional location `US`.
     * Changing this forces a new resource to be created.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Defines the time travel window in hours. The value can be from 48 to 168 hours (2 to 7 days).
     */
    public readonly maxTimeTravelHours!: pulumi.Output<string>;
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
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * Specifies the storage billing model for the dataset.
     * Set this flag value to LOGICAL to use logical bytes for storage billing,
     * or to PHYSICAL to use physical bytes instead.
     * LOGICAL is the default if this flag isn't specified.
     */
    public readonly storageBillingModel!: pulumi.Output<string>;

    /**
     * Create a Dataset resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatasetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatasetArgs | DatasetState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DatasetState | undefined;
            resourceInputs["accesses"] = state ? state.accesses : undefined;
            resourceInputs["creationTime"] = state ? state.creationTime : undefined;
            resourceInputs["datasetId"] = state ? state.datasetId : undefined;
            resourceInputs["defaultCollation"] = state ? state.defaultCollation : undefined;
            resourceInputs["defaultEncryptionConfiguration"] = state ? state.defaultEncryptionConfiguration : undefined;
            resourceInputs["defaultPartitionExpirationMs"] = state ? state.defaultPartitionExpirationMs : undefined;
            resourceInputs["defaultTableExpirationMs"] = state ? state.defaultTableExpirationMs : undefined;
            resourceInputs["deleteContentsOnDestroy"] = state ? state.deleteContentsOnDestroy : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["effectiveLabels"] = state ? state.effectiveLabels : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["friendlyName"] = state ? state.friendlyName : undefined;
            resourceInputs["isCaseInsensitive"] = state ? state.isCaseInsensitive : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["lastModifiedTime"] = state ? state.lastModifiedTime : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["maxTimeTravelHours"] = state ? state.maxTimeTravelHours : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["pulumiLabels"] = state ? state.pulumiLabels : undefined;
            resourceInputs["selfLink"] = state ? state.selfLink : undefined;
            resourceInputs["storageBillingModel"] = state ? state.storageBillingModel : undefined;
        } else {
            const args = argsOrState as DatasetArgs | undefined;
            if ((!args || args.datasetId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'datasetId'");
            }
            resourceInputs["accesses"] = args ? args.accesses : undefined;
            resourceInputs["datasetId"] = args ? args.datasetId : undefined;
            resourceInputs["defaultCollation"] = args ? args.defaultCollation : undefined;
            resourceInputs["defaultEncryptionConfiguration"] = args ? args.defaultEncryptionConfiguration : undefined;
            resourceInputs["defaultPartitionExpirationMs"] = args ? args.defaultPartitionExpirationMs : undefined;
            resourceInputs["defaultTableExpirationMs"] = args ? args.defaultTableExpirationMs : undefined;
            resourceInputs["deleteContentsOnDestroy"] = args ? args.deleteContentsOnDestroy : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["friendlyName"] = args ? args.friendlyName : undefined;
            resourceInputs["isCaseInsensitive"] = args ? args.isCaseInsensitive : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["maxTimeTravelHours"] = args ? args.maxTimeTravelHours : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["storageBillingModel"] = args ? args.storageBillingModel : undefined;
            resourceInputs["creationTime"] = undefined /*out*/;
            resourceInputs["effectiveLabels"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["lastModifiedTime"] = undefined /*out*/;
            resourceInputs["pulumiLabels"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["effectiveLabels", "pulumiLabels"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Dataset.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Dataset resources.
 */
export interface DatasetState {
    /**
     * An array of objects that define dataset access for one or more entities.
     * Structure is documented below.
     */
    accesses?: pulumi.Input<pulumi.Input<inputs.bigquery.DatasetAccess>[]>;
    /**
     * The time when this dataset was created, in milliseconds since the
     * epoch.
     */
    creationTime?: pulumi.Input<number>;
    /**
     * A unique ID for this dataset, without the project name. The ID
     * must contain only letters (a-z, A-Z), numbers (0-9), or
     * underscores (_). The maximum length is 1,024 characters.
     *
     *
     * - - -
     */
    datasetId?: pulumi.Input<string>;
    /**
     * Defines the default collation specification of future tables created
     * in the dataset. If a table is created in this dataset without table-level
     * default collation, then the table inherits the dataset default collation,
     * which is applied to the string fields that do not have explicit collation
     * specified. A change to this field affects only tables created afterwards,
     * and does not alter the existing tables.
     * The following values are supported:
     * - 'und:ci': undetermined locale, case insensitive.
     * - '': empty string. Default to case-sensitive behavior.
     */
    defaultCollation?: pulumi.Input<string>;
    /**
     * The default encryption key for all tables in the dataset. Once this property is set,
     * all newly-created partitioned tables in the dataset will have encryption key set to
     * this value, unless table creation request (or query) overrides the key.
     * Structure is documented below.
     */
    defaultEncryptionConfiguration?: pulumi.Input<inputs.bigquery.DatasetDefaultEncryptionConfiguration>;
    /**
     * The default partition expiration for all partitioned tables in
     * the dataset, in milliseconds.
     *
     * Once this property is set, all newly-created partitioned tables in
     * the dataset will have an `expirationMs` property in the `timePartitioning`
     * settings set to this value, and changing the value will only
     * affect new tables, not existing ones. The storage in a partition will
     * have an expiration time of its partition time plus this value.
     * Setting this property overrides the use of `defaultTableExpirationMs`
     * for partitioned tables: only one of `defaultTableExpirationMs` and
     * `defaultPartitionExpirationMs` will be used for any new partitioned
     * table. If you provide an explicit `timePartitioning.expirationMs` when
     * creating or updating a partitioned table, that value takes precedence
     * over the default partition expiration time indicated by this property.
     */
    defaultPartitionExpirationMs?: pulumi.Input<number>;
    /**
     * The default lifetime of all tables in the dataset, in milliseconds.
     * The minimum value is 3600000 milliseconds (one hour).
     *
     * Once this property is set, all newly-created tables in the dataset
     * will have an `expirationTime` property set to the creation time plus
     * the value in this property, and changing the value will only affect
     * new tables, not existing ones. When the `expirationTime` for a given
     * table is reached, that table will be deleted automatically.
     * If a table's `expirationTime` is modified or removed before the
     * table expires, or if you provide an explicit `expirationTime` when
     * creating a table, that value takes precedence over the default
     * expiration time indicated by this property.
     */
    defaultTableExpirationMs?: pulumi.Input<number>;
    /**
     * If set to `true`, delete all the tables in the
     * dataset when destroying the resource; otherwise,
     * destroying the resource will fail if tables are present.
     */
    deleteContentsOnDestroy?: pulumi.Input<boolean>;
    /**
     * A user-friendly description of the dataset
     */
    description?: pulumi.Input<string>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     */
    effectiveLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A hash of the resource.
     */
    etag?: pulumi.Input<string>;
    /**
     * A descriptive name for the dataset
     */
    friendlyName?: pulumi.Input<string>;
    /**
     * TRUE if the dataset and its table names are case-insensitive, otherwise FALSE.
     * By default, this is FALSE, which means the dataset and its table names are
     * case-sensitive. This field does not affect routine references.
     */
    isCaseInsensitive?: pulumi.Input<boolean>;
    /**
     * The labels associated with this dataset. You can use these to
     * organize and group your datasets.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The date when this dataset or any of its tables was last modified, in
     * milliseconds since the epoch.
     */
    lastModifiedTime?: pulumi.Input<number>;
    /**
     * The geographic location where the dataset should reside.
     * See [official docs](https://cloud.google.com/bigquery/docs/dataset-locations).
     *
     * There are two types of locations, regional or multi-regional. A regional
     * location is a specific geographic place, such as Tokyo, and a multi-regional
     * location is a large geographic area, such as the United States, that
     * contains at least two geographic places.
     *
     * The default value is multi-regional location `US`.
     * Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * Defines the time travel window in hours. The value can be from 48 to 168 hours (2 to 7 days).
     */
    maxTimeTravelHours?: pulumi.Input<string>;
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
     * The URI of the created resource.
     */
    selfLink?: pulumi.Input<string>;
    /**
     * Specifies the storage billing model for the dataset.
     * Set this flag value to LOGICAL to use logical bytes for storage billing,
     * or to PHYSICAL to use physical bytes instead.
     * LOGICAL is the default if this flag isn't specified.
     */
    storageBillingModel?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Dataset resource.
 */
export interface DatasetArgs {
    /**
     * An array of objects that define dataset access for one or more entities.
     * Structure is documented below.
     */
    accesses?: pulumi.Input<pulumi.Input<inputs.bigquery.DatasetAccess>[]>;
    /**
     * A unique ID for this dataset, without the project name. The ID
     * must contain only letters (a-z, A-Z), numbers (0-9), or
     * underscores (_). The maximum length is 1,024 characters.
     *
     *
     * - - -
     */
    datasetId: pulumi.Input<string>;
    /**
     * Defines the default collation specification of future tables created
     * in the dataset. If a table is created in this dataset without table-level
     * default collation, then the table inherits the dataset default collation,
     * which is applied to the string fields that do not have explicit collation
     * specified. A change to this field affects only tables created afterwards,
     * and does not alter the existing tables.
     * The following values are supported:
     * - 'und:ci': undetermined locale, case insensitive.
     * - '': empty string. Default to case-sensitive behavior.
     */
    defaultCollation?: pulumi.Input<string>;
    /**
     * The default encryption key for all tables in the dataset. Once this property is set,
     * all newly-created partitioned tables in the dataset will have encryption key set to
     * this value, unless table creation request (or query) overrides the key.
     * Structure is documented below.
     */
    defaultEncryptionConfiguration?: pulumi.Input<inputs.bigquery.DatasetDefaultEncryptionConfiguration>;
    /**
     * The default partition expiration for all partitioned tables in
     * the dataset, in milliseconds.
     *
     * Once this property is set, all newly-created partitioned tables in
     * the dataset will have an `expirationMs` property in the `timePartitioning`
     * settings set to this value, and changing the value will only
     * affect new tables, not existing ones. The storage in a partition will
     * have an expiration time of its partition time plus this value.
     * Setting this property overrides the use of `defaultTableExpirationMs`
     * for partitioned tables: only one of `defaultTableExpirationMs` and
     * `defaultPartitionExpirationMs` will be used for any new partitioned
     * table. If you provide an explicit `timePartitioning.expirationMs` when
     * creating or updating a partitioned table, that value takes precedence
     * over the default partition expiration time indicated by this property.
     */
    defaultPartitionExpirationMs?: pulumi.Input<number>;
    /**
     * The default lifetime of all tables in the dataset, in milliseconds.
     * The minimum value is 3600000 milliseconds (one hour).
     *
     * Once this property is set, all newly-created tables in the dataset
     * will have an `expirationTime` property set to the creation time plus
     * the value in this property, and changing the value will only affect
     * new tables, not existing ones. When the `expirationTime` for a given
     * table is reached, that table will be deleted automatically.
     * If a table's `expirationTime` is modified or removed before the
     * table expires, or if you provide an explicit `expirationTime` when
     * creating a table, that value takes precedence over the default
     * expiration time indicated by this property.
     */
    defaultTableExpirationMs?: pulumi.Input<number>;
    /**
     * If set to `true`, delete all the tables in the
     * dataset when destroying the resource; otherwise,
     * destroying the resource will fail if tables are present.
     */
    deleteContentsOnDestroy?: pulumi.Input<boolean>;
    /**
     * A user-friendly description of the dataset
     */
    description?: pulumi.Input<string>;
    /**
     * A descriptive name for the dataset
     */
    friendlyName?: pulumi.Input<string>;
    /**
     * TRUE if the dataset and its table names are case-insensitive, otherwise FALSE.
     * By default, this is FALSE, which means the dataset and its table names are
     * case-sensitive. This field does not affect routine references.
     */
    isCaseInsensitive?: pulumi.Input<boolean>;
    /**
     * The labels associated with this dataset. You can use these to
     * organize and group your datasets.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The geographic location where the dataset should reside.
     * See [official docs](https://cloud.google.com/bigquery/docs/dataset-locations).
     *
     * There are two types of locations, regional or multi-regional. A regional
     * location is a specific geographic place, such as Tokyo, and a multi-regional
     * location is a large geographic area, such as the United States, that
     * contains at least two geographic places.
     *
     * The default value is multi-regional location `US`.
     * Changing this forces a new resource to be created.
     */
    location?: pulumi.Input<string>;
    /**
     * Defines the time travel window in hours. The value can be from 48 to 168 hours (2 to 7 days).
     */
    maxTimeTravelHours?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Specifies the storage billing model for the dataset.
     * Set this flag value to LOGICAL to use logical bytes for storage billing,
     * or to PHYSICAL to use physical bytes instead.
     * LOGICAL is the default if this flag isn't specified.
     */
    storageBillingModel?: pulumi.Input<string>;
}
