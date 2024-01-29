// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A FhirStore is a datastore inside a Healthcare dataset that conforms to the FHIR (https://www.hl7.org/fhir/STU3/)
 * standard for Healthcare information exchange
 *
 * To get more information about FhirStore, see:
 *
 * * [API documentation](https://cloud.google.com/healthcare/docs/reference/rest/v1/projects.locations.datasets.fhirStores)
 * * How-to Guides
 *     * [Creating a FHIR store](https://cloud.google.com/healthcare/docs/how-tos/fhir)
 *
 * ## Example Usage
 * ### Healthcare Fhir Store Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const topic = new gcp.pubsub.Topic("topic", {});
 * const dataset = new gcp.healthcare.Dataset("dataset", {location: "us-central1"});
 * const _default = new gcp.healthcare.FhirStore("default", {
 *     dataset: dataset.id,
 *     version: "R4",
 *     complexDataTypeReferenceParsing: "DISABLED",
 *     enableUpdateCreate: false,
 *     disableReferentialIntegrity: false,
 *     disableResourceVersioning: false,
 *     enableHistoryImport: false,
 *     defaultSearchHandlingStrict: false,
 *     notificationConfig: {
 *         pubsubTopic: topic.id,
 *     },
 *     labels: {
 *         label1: "labelvalue1",
 *     },
 * });
 * ```
 * ### Healthcare Fhir Store Streaming Config
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const dataset = new gcp.healthcare.Dataset("dataset", {location: "us-central1"});
 * const bqDataset = new gcp.bigquery.Dataset("bqDataset", {
 *     datasetId: "bq_example_dataset",
 *     friendlyName: "test",
 *     description: "This is a test description",
 *     location: "US",
 *     deleteContentsOnDestroy: true,
 * });
 * const _default = new gcp.healthcare.FhirStore("default", {
 *     dataset: dataset.id,
 *     version: "R4",
 *     enableUpdateCreate: false,
 *     disableReferentialIntegrity: false,
 *     disableResourceVersioning: false,
 *     enableHistoryImport: false,
 *     labels: {
 *         label1: "labelvalue1",
 *     },
 *     streamConfigs: [{
 *         resourceTypes: ["Observation"],
 *         bigqueryDestination: {
 *             datasetUri: pulumi.interpolate`bq://${bqDataset.project}.${bqDataset.datasetId}`,
 *             schemaConfig: {
 *                 recursiveStructureDepth: 3,
 *                 lastUpdatedPartitionConfig: {
 *                     type: "HOUR",
 *                     expirationMs: "1000000",
 *                 },
 *             },
 *         },
 *     }],
 * });
 * const topic = new gcp.pubsub.Topic("topic", {});
 * ```
 * ### Healthcare Fhir Store Notification Config
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const topic = new gcp.pubsub.Topic("topic", {});
 * const dataset = new gcp.healthcare.Dataset("dataset", {location: "us-central1"});
 * const _default = new gcp.healthcare.FhirStore("default", {
 *     dataset: dataset.id,
 *     version: "R4",
 *     enableUpdateCreate: false,
 *     disableReferentialIntegrity: false,
 *     disableResourceVersioning: false,
 *     enableHistoryImport: false,
 *     labels: {
 *         label1: "labelvalue1",
 *     },
 *     notificationConfig: {
 *         pubsubTopic: topic.id,
 *     },
 * });
 * ```
 * ### Healthcare Fhir Store Notification Configs
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const topic = new gcp.pubsub.Topic("topic", {}, {
 *     provider: google_beta,
 * });
 * const dataset = new gcp.healthcare.Dataset("dataset", {location: "us-central1"}, {
 *     provider: google_beta,
 * });
 * const _default = new gcp.healthcare.FhirStore("default", {
 *     dataset: dataset.id,
 *     version: "R4",
 *     enableUpdateCreate: false,
 *     disableReferentialIntegrity: false,
 *     disableResourceVersioning: false,
 *     enableHistoryImport: false,
 *     enableHistoryModifications: false,
 *     labels: {
 *         label1: "labelvalue1",
 *     },
 *     notificationConfigs: [{
 *         pubsubTopic: topic.id,
 *         sendFullResource: true,
 *         sendPreviousResourceOnDelete: true,
 *     }],
 * }, {
 *     provider: google_beta,
 * });
 * ```
 *
 * ## Import
 *
 * FhirStore can be imported using any of these accepted formats* `{{dataset}}/fhirStores/{{name}}` * `{{dataset}}/{{name}}` When using the `pulumi import` command, FhirStore can be imported using one of the formats above. For example
 *
 * ```sh
 *  $ pulumi import gcp:healthcare/fhirStore:FhirStore default {{dataset}}/fhirStores/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:healthcare/fhirStore:FhirStore default {{dataset}}/{{name}}
 * ```
 */
export class FhirStore extends pulumi.CustomResource {
    /**
     * Get an existing FhirStore resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FhirStoreState, opts?: pulumi.CustomResourceOptions): FhirStore {
        return new FhirStore(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:healthcare/fhirStore:FhirStore';

    /**
     * Returns true if the given object is an instance of FhirStore.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FhirStore {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FhirStore.__pulumiType;
    }

    /**
     * Enable parsing of references within complex FHIR data types such as Extensions. If this value is set to ENABLED, then features like referential integrity and Bundle reference rewriting apply to all references. If this flag has not been specified the behavior of the FHIR store will not change, references in complex data types will not be parsed. New stores will have this value set to ENABLED by default after a notification period. Warning: turning on this flag causes processing existing resources to fail if they contain references to non-existent resources.
     * Possible values are: `COMPLEX_DATA_TYPE_REFERENCE_PARSING_UNSPECIFIED`, `DISABLED`, `ENABLED`.
     */
    public readonly complexDataTypeReferenceParsing!: pulumi.Output<string>;
    /**
     * Identifies the dataset addressed by this request. Must be in the format
     * 'projects/{project}/locations/{location}/datasets/{dataset}'
     *
     *
     * - - -
     */
    public readonly dataset!: pulumi.Output<string>;
    /**
     * If true, overrides the default search behavior for this FHIR store to handling=strict which returns an error for unrecognized search parameters.
     * If false, uses the FHIR specification default handling=lenient which ignores unrecognized search parameters.
     * The handling can always be changed from the default on an individual API call by setting the HTTP header Prefer: handling=strict or Prefer: handling=lenient.
     */
    public readonly defaultSearchHandlingStrict!: pulumi.Output<boolean | undefined>;
    /**
     * Whether to disable referential integrity in this FHIR store. This field is immutable after FHIR store
     * creation. The default value is false, meaning that the API will enforce referential integrity and fail the
     * requests that will result in inconsistent state in the FHIR store. When this field is set to true, the API
     * will skip referential integrity check. Consequently, operations that rely on references, such as
     * Patient.get$everything, will not return all the results if broken references exist.
     * ** Changing this property may recreate the FHIR store (removing all data) **
     */
    public readonly disableReferentialIntegrity!: pulumi.Output<boolean | undefined>;
    /**
     * Whether to disable resource versioning for this FHIR store. This field can not be changed after the creation
     * of FHIR store. If set to false, which is the default behavior, all write operations will cause historical
     * versions to be recorded automatically. The historical versions can be fetched through the history APIs, but
     * cannot be updated. If set to true, no historical versions will be kept. The server will send back errors for
     * attempts to read the historical versions.
     * ** Changing this property may recreate the FHIR store (removing all data) **
     */
    public readonly disableResourceVersioning!: pulumi.Output<boolean | undefined>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     */
    public /*out*/ readonly effectiveLabels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Whether to allow the bulk import API to accept history bundles and directly insert historical resource
     * versions into the FHIR store. Importing resource histories creates resource interactions that appear to have
     * occurred in the past, which clients may not want to allow. If set to false, history bundles within an import
     * will fail with an error.
     * ** Changing this property may recreate the FHIR store (removing all data) **
     * ** This property can be changed manually in the Google Cloud Healthcare admin console without recreating the FHIR store **
     */
    public readonly enableHistoryImport!: pulumi.Output<boolean | undefined>;
    /**
     * Whether to allow the ExecuteBundle API to accept history bundles, and directly insert and overwrite historical
     * resource versions into the FHIR store. If set to false, using history bundles fails with an error.
     */
    public readonly enableHistoryModifications!: pulumi.Output<boolean | undefined>;
    /**
     * Whether this FHIR store has the updateCreate capability. This determines if the client can use an Update
     * operation to create a new resource with a client-specified ID. If false, all IDs are server-assigned through
     * the Create operation and attempts to Update a non-existent resource will return errors. Please treat the audit
     * logs with appropriate levels of care if client-specified resource IDs contain sensitive data such as patient
     * identifiers, those IDs will be part of the FHIR resource path recorded in Cloud audit logs and Cloud Pub/Sub
     * notifications.
     */
    public readonly enableUpdateCreate!: pulumi.Output<boolean | undefined>;
    /**
     * User-supplied key-value pairs used to organize FHIR stores.
     * Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
     * conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
     * Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
     * bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
     * No more than 64 labels can be associated with a given store.
     * An object containing a list of "key": value pairs.
     * Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The resource name for the FhirStore.
     * ** Changing this property may recreate the FHIR store (removing all data) **
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A nested object resource
     * Structure is documented below.
     */
    public readonly notificationConfig!: pulumi.Output<outputs.healthcare.FhirStoreNotificationConfig | undefined>;
    /**
     * A list of notifcation configs that configure the notification for every resource mutation in this FHIR store.
     * Structure is documented below.
     */
    public readonly notificationConfigs!: pulumi.Output<outputs.healthcare.FhirStoreNotificationConfig[] | undefined>;
    /**
     * The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     */
    public /*out*/ readonly pulumiLabels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The fully qualified name of this dataset
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * A list of streaming configs that configure the destinations of streaming export for every resource mutation in
     * this FHIR store. Each store is allowed to have up to 10 streaming configs. After a new config is added, the next
     * resource mutation is streamed to the new location in addition to the existing ones. When a location is removed
     * from the list, the server stops streaming to that location. Before adding a new config, you must add the required
     * bigquery.dataEditor role to your project's Cloud Healthcare Service Agent service account. Some lag (typically on
     * the order of dozens of seconds) is expected before the results show up in the streaming destination.
     * Structure is documented below.
     */
    public readonly streamConfigs!: pulumi.Output<outputs.healthcare.FhirStoreStreamConfig[] | undefined>;
    /**
     * The FHIR specification version.
     * Default value is `STU3`.
     * Possible values are: `DSTU2`, `STU3`, `R4`.
     */
    public readonly version!: pulumi.Output<string | undefined>;

    /**
     * Create a FhirStore resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FhirStoreArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FhirStoreArgs | FhirStoreState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FhirStoreState | undefined;
            resourceInputs["complexDataTypeReferenceParsing"] = state ? state.complexDataTypeReferenceParsing : undefined;
            resourceInputs["dataset"] = state ? state.dataset : undefined;
            resourceInputs["defaultSearchHandlingStrict"] = state ? state.defaultSearchHandlingStrict : undefined;
            resourceInputs["disableReferentialIntegrity"] = state ? state.disableReferentialIntegrity : undefined;
            resourceInputs["disableResourceVersioning"] = state ? state.disableResourceVersioning : undefined;
            resourceInputs["effectiveLabels"] = state ? state.effectiveLabels : undefined;
            resourceInputs["enableHistoryImport"] = state ? state.enableHistoryImport : undefined;
            resourceInputs["enableHistoryModifications"] = state ? state.enableHistoryModifications : undefined;
            resourceInputs["enableUpdateCreate"] = state ? state.enableUpdateCreate : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["notificationConfig"] = state ? state.notificationConfig : undefined;
            resourceInputs["notificationConfigs"] = state ? state.notificationConfigs : undefined;
            resourceInputs["pulumiLabels"] = state ? state.pulumiLabels : undefined;
            resourceInputs["selfLink"] = state ? state.selfLink : undefined;
            resourceInputs["streamConfigs"] = state ? state.streamConfigs : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as FhirStoreArgs | undefined;
            if ((!args || args.dataset === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataset'");
            }
            resourceInputs["complexDataTypeReferenceParsing"] = args ? args.complexDataTypeReferenceParsing : undefined;
            resourceInputs["dataset"] = args ? args.dataset : undefined;
            resourceInputs["defaultSearchHandlingStrict"] = args ? args.defaultSearchHandlingStrict : undefined;
            resourceInputs["disableReferentialIntegrity"] = args ? args.disableReferentialIntegrity : undefined;
            resourceInputs["disableResourceVersioning"] = args ? args.disableResourceVersioning : undefined;
            resourceInputs["enableHistoryImport"] = args ? args.enableHistoryImport : undefined;
            resourceInputs["enableHistoryModifications"] = args ? args.enableHistoryModifications : undefined;
            resourceInputs["enableUpdateCreate"] = args ? args.enableUpdateCreate : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["notificationConfig"] = args ? args.notificationConfig : undefined;
            resourceInputs["notificationConfigs"] = args ? args.notificationConfigs : undefined;
            resourceInputs["streamConfigs"] = args ? args.streamConfigs : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["effectiveLabels"] = undefined /*out*/;
            resourceInputs["pulumiLabels"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["effectiveLabels", "pulumiLabels"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(FhirStore.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FhirStore resources.
 */
export interface FhirStoreState {
    /**
     * Enable parsing of references within complex FHIR data types such as Extensions. If this value is set to ENABLED, then features like referential integrity and Bundle reference rewriting apply to all references. If this flag has not been specified the behavior of the FHIR store will not change, references in complex data types will not be parsed. New stores will have this value set to ENABLED by default after a notification period. Warning: turning on this flag causes processing existing resources to fail if they contain references to non-existent resources.
     * Possible values are: `COMPLEX_DATA_TYPE_REFERENCE_PARSING_UNSPECIFIED`, `DISABLED`, `ENABLED`.
     */
    complexDataTypeReferenceParsing?: pulumi.Input<string>;
    /**
     * Identifies the dataset addressed by this request. Must be in the format
     * 'projects/{project}/locations/{location}/datasets/{dataset}'
     *
     *
     * - - -
     */
    dataset?: pulumi.Input<string>;
    /**
     * If true, overrides the default search behavior for this FHIR store to handling=strict which returns an error for unrecognized search parameters.
     * If false, uses the FHIR specification default handling=lenient which ignores unrecognized search parameters.
     * The handling can always be changed from the default on an individual API call by setting the HTTP header Prefer: handling=strict or Prefer: handling=lenient.
     */
    defaultSearchHandlingStrict?: pulumi.Input<boolean>;
    /**
     * Whether to disable referential integrity in this FHIR store. This field is immutable after FHIR store
     * creation. The default value is false, meaning that the API will enforce referential integrity and fail the
     * requests that will result in inconsistent state in the FHIR store. When this field is set to true, the API
     * will skip referential integrity check. Consequently, operations that rely on references, such as
     * Patient.get$everything, will not return all the results if broken references exist.
     * ** Changing this property may recreate the FHIR store (removing all data) **
     */
    disableReferentialIntegrity?: pulumi.Input<boolean>;
    /**
     * Whether to disable resource versioning for this FHIR store. This field can not be changed after the creation
     * of FHIR store. If set to false, which is the default behavior, all write operations will cause historical
     * versions to be recorded automatically. The historical versions can be fetched through the history APIs, but
     * cannot be updated. If set to true, no historical versions will be kept. The server will send back errors for
     * attempts to read the historical versions.
     * ** Changing this property may recreate the FHIR store (removing all data) **
     */
    disableResourceVersioning?: pulumi.Input<boolean>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     */
    effectiveLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Whether to allow the bulk import API to accept history bundles and directly insert historical resource
     * versions into the FHIR store. Importing resource histories creates resource interactions that appear to have
     * occurred in the past, which clients may not want to allow. If set to false, history bundles within an import
     * will fail with an error.
     * ** Changing this property may recreate the FHIR store (removing all data) **
     * ** This property can be changed manually in the Google Cloud Healthcare admin console without recreating the FHIR store **
     */
    enableHistoryImport?: pulumi.Input<boolean>;
    /**
     * Whether to allow the ExecuteBundle API to accept history bundles, and directly insert and overwrite historical
     * resource versions into the FHIR store. If set to false, using history bundles fails with an error.
     */
    enableHistoryModifications?: pulumi.Input<boolean>;
    /**
     * Whether this FHIR store has the updateCreate capability. This determines if the client can use an Update
     * operation to create a new resource with a client-specified ID. If false, all IDs are server-assigned through
     * the Create operation and attempts to Update a non-existent resource will return errors. Please treat the audit
     * logs with appropriate levels of care if client-specified resource IDs contain sensitive data such as patient
     * identifiers, those IDs will be part of the FHIR resource path recorded in Cloud audit logs and Cloud Pub/Sub
     * notifications.
     */
    enableUpdateCreate?: pulumi.Input<boolean>;
    /**
     * User-supplied key-value pairs used to organize FHIR stores.
     * Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
     * conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
     * Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
     * bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
     * No more than 64 labels can be associated with a given store.
     * An object containing a list of "key": value pairs.
     * Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The resource name for the FhirStore.
     * ** Changing this property may recreate the FHIR store (removing all data) **
     */
    name?: pulumi.Input<string>;
    /**
     * A nested object resource
     * Structure is documented below.
     */
    notificationConfig?: pulumi.Input<inputs.healthcare.FhirStoreNotificationConfig>;
    /**
     * A list of notifcation configs that configure the notification for every resource mutation in this FHIR store.
     * Structure is documented below.
     */
    notificationConfigs?: pulumi.Input<pulumi.Input<inputs.healthcare.FhirStoreNotificationConfig>[]>;
    /**
     * The combination of labels configured directly on the resource
     * and default labels configured on the provider.
     */
    pulumiLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The fully qualified name of this dataset
     */
    selfLink?: pulumi.Input<string>;
    /**
     * A list of streaming configs that configure the destinations of streaming export for every resource mutation in
     * this FHIR store. Each store is allowed to have up to 10 streaming configs. After a new config is added, the next
     * resource mutation is streamed to the new location in addition to the existing ones. When a location is removed
     * from the list, the server stops streaming to that location. Before adding a new config, you must add the required
     * bigquery.dataEditor role to your project's Cloud Healthcare Service Agent service account. Some lag (typically on
     * the order of dozens of seconds) is expected before the results show up in the streaming destination.
     * Structure is documented below.
     */
    streamConfigs?: pulumi.Input<pulumi.Input<inputs.healthcare.FhirStoreStreamConfig>[]>;
    /**
     * The FHIR specification version.
     * Default value is `STU3`.
     * Possible values are: `DSTU2`, `STU3`, `R4`.
     */
    version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a FhirStore resource.
 */
export interface FhirStoreArgs {
    /**
     * Enable parsing of references within complex FHIR data types such as Extensions. If this value is set to ENABLED, then features like referential integrity and Bundle reference rewriting apply to all references. If this flag has not been specified the behavior of the FHIR store will not change, references in complex data types will not be parsed. New stores will have this value set to ENABLED by default after a notification period. Warning: turning on this flag causes processing existing resources to fail if they contain references to non-existent resources.
     * Possible values are: `COMPLEX_DATA_TYPE_REFERENCE_PARSING_UNSPECIFIED`, `DISABLED`, `ENABLED`.
     */
    complexDataTypeReferenceParsing?: pulumi.Input<string>;
    /**
     * Identifies the dataset addressed by this request. Must be in the format
     * 'projects/{project}/locations/{location}/datasets/{dataset}'
     *
     *
     * - - -
     */
    dataset: pulumi.Input<string>;
    /**
     * If true, overrides the default search behavior for this FHIR store to handling=strict which returns an error for unrecognized search parameters.
     * If false, uses the FHIR specification default handling=lenient which ignores unrecognized search parameters.
     * The handling can always be changed from the default on an individual API call by setting the HTTP header Prefer: handling=strict or Prefer: handling=lenient.
     */
    defaultSearchHandlingStrict?: pulumi.Input<boolean>;
    /**
     * Whether to disable referential integrity in this FHIR store. This field is immutable after FHIR store
     * creation. The default value is false, meaning that the API will enforce referential integrity and fail the
     * requests that will result in inconsistent state in the FHIR store. When this field is set to true, the API
     * will skip referential integrity check. Consequently, operations that rely on references, such as
     * Patient.get$everything, will not return all the results if broken references exist.
     * ** Changing this property may recreate the FHIR store (removing all data) **
     */
    disableReferentialIntegrity?: pulumi.Input<boolean>;
    /**
     * Whether to disable resource versioning for this FHIR store. This field can not be changed after the creation
     * of FHIR store. If set to false, which is the default behavior, all write operations will cause historical
     * versions to be recorded automatically. The historical versions can be fetched through the history APIs, but
     * cannot be updated. If set to true, no historical versions will be kept. The server will send back errors for
     * attempts to read the historical versions.
     * ** Changing this property may recreate the FHIR store (removing all data) **
     */
    disableResourceVersioning?: pulumi.Input<boolean>;
    /**
     * Whether to allow the bulk import API to accept history bundles and directly insert historical resource
     * versions into the FHIR store. Importing resource histories creates resource interactions that appear to have
     * occurred in the past, which clients may not want to allow. If set to false, history bundles within an import
     * will fail with an error.
     * ** Changing this property may recreate the FHIR store (removing all data) **
     * ** This property can be changed manually in the Google Cloud Healthcare admin console without recreating the FHIR store **
     */
    enableHistoryImport?: pulumi.Input<boolean>;
    /**
     * Whether to allow the ExecuteBundle API to accept history bundles, and directly insert and overwrite historical
     * resource versions into the FHIR store. If set to false, using history bundles fails with an error.
     */
    enableHistoryModifications?: pulumi.Input<boolean>;
    /**
     * Whether this FHIR store has the updateCreate capability. This determines if the client can use an Update
     * operation to create a new resource with a client-specified ID. If false, all IDs are server-assigned through
     * the Create operation and attempts to Update a non-existent resource will return errors. Please treat the audit
     * logs with appropriate levels of care if client-specified resource IDs contain sensitive data such as patient
     * identifiers, those IDs will be part of the FHIR resource path recorded in Cloud audit logs and Cloud Pub/Sub
     * notifications.
     */
    enableUpdateCreate?: pulumi.Input<boolean>;
    /**
     * User-supplied key-value pairs used to organize FHIR stores.
     * Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must
     * conform to the following PCRE regular expression: [\p{Ll}\p{Lo}][\p{Ll}\p{Lo}\p{N}_-]{0,62}
     * Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128
     * bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63}
     * No more than 64 labels can be associated with a given store.
     * An object containing a list of "key": value pairs.
     * Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The resource name for the FhirStore.
     * ** Changing this property may recreate the FHIR store (removing all data) **
     */
    name?: pulumi.Input<string>;
    /**
     * A nested object resource
     * Structure is documented below.
     */
    notificationConfig?: pulumi.Input<inputs.healthcare.FhirStoreNotificationConfig>;
    /**
     * A list of notifcation configs that configure the notification for every resource mutation in this FHIR store.
     * Structure is documented below.
     */
    notificationConfigs?: pulumi.Input<pulumi.Input<inputs.healthcare.FhirStoreNotificationConfig>[]>;
    /**
     * A list of streaming configs that configure the destinations of streaming export for every resource mutation in
     * this FHIR store. Each store is allowed to have up to 10 streaming configs. After a new config is added, the next
     * resource mutation is streamed to the new location in addition to the existing ones. When a location is removed
     * from the list, the server stops streaming to that location. Before adding a new config, you must add the required
     * bigquery.dataEditor role to your project's Cloud Healthcare Service Agent service account. Some lag (typically on
     * the order of dozens of seconds) is expected before the results show up in the streaming destination.
     * Structure is documented below.
     */
    streamConfigs?: pulumi.Input<pulumi.Input<inputs.healthcare.FhirStoreStreamConfig>[]>;
    /**
     * The FHIR specification version.
     * Default value is `STU3`.
     * Possible values are: `DSTU2`, `STU3`, `R4`.
     */
    version?: pulumi.Input<string>;
}
