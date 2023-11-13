// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Represents a user-visible job which provides the insights for the related data source.
 *
 * To get more information about Datascan, see:
 *
 * * [API documentation](https://cloud.google.com/dataplex/docs/reference/rest)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/dataplex/docs)
 *
 * ## Example Usage
 * ### Dataplex Datascan Basic Profile
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const basicProfile = new gcp.dataplex.Datascan("basicProfile", {
 *     data: {
 *         resource: "//bigquery.googleapis.com/projects/bigquery-public-data/datasets/samples/tables/shakespeare",
 *     },
 *     dataProfileSpec: {},
 *     dataScanId: "dataprofile-basic",
 *     executionSpec: {
 *         trigger: {
 *             onDemand: {},
 *         },
 *     },
 *     location: "us-central1",
 *     project: "my-project-name",
 * });
 * ```
 * ### Dataplex Datascan Full Profile
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const source = new gcp.bigquery.Dataset("source", {
 *     datasetId: "dataplex_dataset",
 *     friendlyName: "test",
 *     description: "This is a test description",
 *     location: "US",
 *     deleteContentsOnDestroy: true,
 * });
 * const fullProfile = new gcp.dataplex.Datascan("fullProfile", {
 *     location: "us-central1",
 *     displayName: "Full Datascan Profile",
 *     dataScanId: "dataprofile-full",
 *     description: "Example resource - Full Datascan Profile",
 *     labels: {
 *         author: "billing",
 *     },
 *     data: {
 *         resource: "//bigquery.googleapis.com/projects/bigquery-public-data/datasets/samples/tables/shakespeare",
 *     },
 *     executionSpec: {
 *         trigger: {
 *             schedule: {
 *                 cron: "TZ=America/New_York 1 1 * * *",
 *             },
 *         },
 *     },
 *     dataProfileSpec: {
 *         samplingPercent: 80,
 *         rowFilter: "word_count > 10",
 *         includeFields: {
 *             fieldNames: ["word_count"],
 *         },
 *         excludeFields: {
 *             fieldNames: ["property_type"],
 *         },
 *         postScanActions: {
 *             bigqueryExport: {
 *                 resultsTable: "//bigquery.googleapis.com/projects/my-project-name/datasets/dataplex_dataset/tables/profile_export",
 *             },
 *         },
 *     },
 *     project: "my-project-name",
 * }, {
 *     dependsOn: [source],
 * });
 * ```
 * ### Dataplex Datascan Basic Quality
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const basicQuality = new gcp.dataplex.Datascan("basicQuality", {
 *     data: {
 *         resource: "//bigquery.googleapis.com/projects/bigquery-public-data/datasets/samples/tables/shakespeare",
 *     },
 *     dataQualitySpec: {
 *         rules: [{
 *             description: "rule 1 for validity dimension",
 *             dimension: "VALIDITY",
 *             name: "rule1",
 *             tableConditionExpectation: {
 *                 sqlExpression: "COUNT(*) > 0",
 *             },
 *         }],
 *     },
 *     dataScanId: "dataquality-basic",
 *     executionSpec: {
 *         trigger: {
 *             onDemand: {},
 *         },
 *     },
 *     location: "us-central1",
 *     project: "my-project-name",
 * });
 * ```
 * ### Dataplex Datascan Full Quality
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const fullQuality = new gcp.dataplex.Datascan("fullQuality", {
 *     data: {
 *         resource: "//bigquery.googleapis.com/projects/bigquery-public-data/datasets/austin_bikeshare/tables/bikeshare_stations",
 *     },
 *     dataQualitySpec: {
 *         rowFilter: "station_id > 1000",
 *         rules: [
 *             {
 *                 column: "address",
 *                 dimension: "VALIDITY",
 *                 nonNullExpectation: {},
 *                 threshold: 0.99,
 *             },
 *             {
 *                 column: "council_district",
 *                 dimension: "VALIDITY",
 *                 ignoreNull: true,
 *                 rangeExpectation: {
 *                     maxValue: "10",
 *                     minValue: "1",
 *                     strictMaxEnabled: false,
 *                     strictMinEnabled: true,
 *                 },
 *                 threshold: 0.9,
 *             },
 *             {
 *                 column: "power_type",
 *                 dimension: "VALIDITY",
 *                 ignoreNull: false,
 *                 regexExpectation: {
 *                     regex: ".*solar.*",
 *                 },
 *             },
 *             {
 *                 column: "property_type",
 *                 dimension: "VALIDITY",
 *                 ignoreNull: false,
 *                 setExpectation: {
 *                     values: [
 *                         "sidewalk",
 *                         "parkland",
 *                     ],
 *                 },
 *             },
 *             {
 *                 column: "address",
 *                 dimension: "UNIQUENESS",
 *                 uniquenessExpectation: {},
 *             },
 *             {
 *                 column: "number_of_docks",
 *                 dimension: "VALIDITY",
 *                 statisticRangeExpectation: {
 *                     maxValue: "15",
 *                     minValue: "5",
 *                     statistic: "MEAN",
 *                     strictMaxEnabled: true,
 *                     strictMinEnabled: true,
 *                 },
 *             },
 *             {
 *                 column: "footprint_length",
 *                 dimension: "VALIDITY",
 *                 rowConditionExpectation: {
 *                     sqlExpression: "footprint_length > 0 AND footprint_length <= 10",
 *                 },
 *             },
 *             {
 *                 dimension: "VALIDITY",
 *                 tableConditionExpectation: {
 *                     sqlExpression: "COUNT(*) > 0",
 *                 },
 *             },
 *         ],
 *         samplingPercent: 5,
 *     },
 *     dataScanId: "dataquality-full",
 *     description: "Example resource - Full Datascan Quality",
 *     displayName: "Full Datascan Quality",
 *     executionSpec: {
 *         field: "modified_date",
 *         trigger: {
 *             schedule: {
 *                 cron: "TZ=America/New_York 1 1 * * *",
 *             },
 *         },
 *     },
 *     labels: {
 *         author: "billing",
 *     },
 *     location: "us-central1",
 *     project: "my-project-name",
 * });
 * ```
 *
 * ## Import
 *
 * Datascan can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:dataplex/datascan:Datascan default projects/{{project}}/locations/{{location}}/dataScans/{{data_scan_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:dataplex/datascan:Datascan default {{project}}/{{location}}/{{data_scan_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:dataplex/datascan:Datascan default {{location}}/{{data_scan_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:dataplex/datascan:Datascan default {{data_scan_id}}
 * ```
 */
export class Datascan extends pulumi.CustomResource {
    /**
     * Get an existing Datascan resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DatascanState, opts?: pulumi.CustomResourceOptions): Datascan {
        return new Datascan(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:dataplex/datascan:Datascan';

    /**
     * Returns true if the given object is an instance of Datascan.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Datascan {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Datascan.__pulumiType;
    }

    /**
     * The time when the scan was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The data source for DataScan.
     * Structure is documented below.
     */
    public readonly data!: pulumi.Output<outputs.dataplex.DatascanData>;
    /**
     * DataProfileScan related setting.
     * Structure is documented below.
     */
    public readonly dataProfileSpec!: pulumi.Output<outputs.dataplex.DatascanDataProfileSpec | undefined>;
    /**
     * DataQualityScan related setting.
     * Structure is documented below.
     */
    public readonly dataQualitySpec!: pulumi.Output<outputs.dataplex.DatascanDataQualitySpec | undefined>;
    /**
     * DataScan identifier. Must contain only lowercase letters, numbers and hyphens. Must start with a letter. Must end with a number or a letter.
     */
    public readonly dataScanId!: pulumi.Output<string>;
    /**
     * Description of the rule.
     * The maximum length is 1,024 characters.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * User friendly display name.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     */
    public /*out*/ readonly effectiveLabels!: pulumi.Output<{[key: string]: string}>;
    /**
     * DataScan execution settings.
     * Structure is documented below.
     */
    public readonly executionSpec!: pulumi.Output<outputs.dataplex.DatascanExecutionSpec>;
    /**
     * Status of the data scan execution.
     * Structure is documented below.
     */
    public /*out*/ readonly executionStatuses!: pulumi.Output<outputs.dataplex.DatascanExecutionStatus[]>;
    /**
     * User-defined labels for the scan. A list of key->value pairs.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The location where the data scan should reside.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * A mutable name for the rule.
     * The name must contain only letters (a-z, A-Z), numbers (0-9), or hyphens (-).
     * The maximum length is 63 characters.
     * Must start with a letter.
     * Must end with a number or a letter.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
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
     * Current state of the DataScan.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The type of DataScan.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * System generated globally unique ID for the scan. This ID will be different if the scan is deleted and re-created with the same name.
     */
    public /*out*/ readonly uid!: pulumi.Output<string>;
    /**
     * The time when the scan was last updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Datascan resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DatascanArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DatascanArgs | DatascanState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DatascanState | undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["data"] = state ? state.data : undefined;
            resourceInputs["dataProfileSpec"] = state ? state.dataProfileSpec : undefined;
            resourceInputs["dataQualitySpec"] = state ? state.dataQualitySpec : undefined;
            resourceInputs["dataScanId"] = state ? state.dataScanId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["effectiveLabels"] = state ? state.effectiveLabels : undefined;
            resourceInputs["executionSpec"] = state ? state.executionSpec : undefined;
            resourceInputs["executionStatuses"] = state ? state.executionStatuses : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["pulumiLabels"] = state ? state.pulumiLabels : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["uid"] = state ? state.uid : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as DatascanArgs | undefined;
            if ((!args || args.data === undefined) && !opts.urn) {
                throw new Error("Missing required property 'data'");
            }
            if ((!args || args.dataScanId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataScanId'");
            }
            if ((!args || args.executionSpec === undefined) && !opts.urn) {
                throw new Error("Missing required property 'executionSpec'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            resourceInputs["data"] = args ? args.data : undefined;
            resourceInputs["dataProfileSpec"] = args ? args.dataProfileSpec : undefined;
            resourceInputs["dataQualitySpec"] = args ? args.dataQualitySpec : undefined;
            resourceInputs["dataScanId"] = args ? args.dataScanId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["executionSpec"] = args ? args.executionSpec : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["effectiveLabels"] = undefined /*out*/;
            resourceInputs["executionStatuses"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["pulumiLabels"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["effectiveLabels", "pulumiLabels"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Datascan.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Datascan resources.
 */
export interface DatascanState {
    /**
     * The time when the scan was created.
     */
    createTime?: pulumi.Input<string>;
    /**
     * The data source for DataScan.
     * Structure is documented below.
     */
    data?: pulumi.Input<inputs.dataplex.DatascanData>;
    /**
     * DataProfileScan related setting.
     * Structure is documented below.
     */
    dataProfileSpec?: pulumi.Input<inputs.dataplex.DatascanDataProfileSpec>;
    /**
     * DataQualityScan related setting.
     * Structure is documented below.
     */
    dataQualitySpec?: pulumi.Input<inputs.dataplex.DatascanDataQualitySpec>;
    /**
     * DataScan identifier. Must contain only lowercase letters, numbers and hyphens. Must start with a letter. Must end with a number or a letter.
     */
    dataScanId?: pulumi.Input<string>;
    /**
     * Description of the rule.
     * The maximum length is 1,024 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * User friendly display name.
     */
    displayName?: pulumi.Input<string>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
     */
    effectiveLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * DataScan execution settings.
     * Structure is documented below.
     */
    executionSpec?: pulumi.Input<inputs.dataplex.DatascanExecutionSpec>;
    /**
     * Status of the data scan execution.
     * Structure is documented below.
     */
    executionStatuses?: pulumi.Input<pulumi.Input<inputs.dataplex.DatascanExecutionStatus>[]>;
    /**
     * User-defined labels for the scan. A list of key->value pairs.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The location where the data scan should reside.
     */
    location?: pulumi.Input<string>;
    /**
     * A mutable name for the rule.
     * The name must contain only letters (a-z, A-Z), numbers (0-9), or hyphens (-).
     * The maximum length is 63 characters.
     * Must start with a letter.
     * Must end with a number or a letter.
     */
    name?: pulumi.Input<string>;
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
     * Current state of the DataScan.
     */
    state?: pulumi.Input<string>;
    /**
     * The type of DataScan.
     */
    type?: pulumi.Input<string>;
    /**
     * System generated globally unique ID for the scan. This ID will be different if the scan is deleted and re-created with the same name.
     */
    uid?: pulumi.Input<string>;
    /**
     * The time when the scan was last updated.
     */
    updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Datascan resource.
 */
export interface DatascanArgs {
    /**
     * The data source for DataScan.
     * Structure is documented below.
     */
    data: pulumi.Input<inputs.dataplex.DatascanData>;
    /**
     * DataProfileScan related setting.
     * Structure is documented below.
     */
    dataProfileSpec?: pulumi.Input<inputs.dataplex.DatascanDataProfileSpec>;
    /**
     * DataQualityScan related setting.
     * Structure is documented below.
     */
    dataQualitySpec?: pulumi.Input<inputs.dataplex.DatascanDataQualitySpec>;
    /**
     * DataScan identifier. Must contain only lowercase letters, numbers and hyphens. Must start with a letter. Must end with a number or a letter.
     */
    dataScanId: pulumi.Input<string>;
    /**
     * Description of the rule.
     * The maximum length is 1,024 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * User friendly display name.
     */
    displayName?: pulumi.Input<string>;
    /**
     * DataScan execution settings.
     * Structure is documented below.
     */
    executionSpec: pulumi.Input<inputs.dataplex.DatascanExecutionSpec>;
    /**
     * User-defined labels for the scan. A list of key->value pairs.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field `effectiveLabels` for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The location where the data scan should reside.
     */
    location: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
}
