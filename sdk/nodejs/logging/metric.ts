// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Logs-based metric can also be used to extract values from logs and create a a distribution
 * of the values. The distribution records the statistics of the extracted values along with
 * an optional histogram of the values as specified by the bucket options.
 *
 * To get more information about Metric, see:
 *
 * * [API documentation](https://cloud.google.com/logging/docs/reference/v2/rest/v2/projects.metrics/create)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/logging/docs/apis)
 *
 * ## Example Usage
 * ### Logging Metric Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const loggingMetric = new gcp.logging.Metric("loggingMetric", {
 *     bucketOptions: {
 *         linearBuckets: {
 *             numFiniteBuckets: 3,
 *             offset: 1,
 *             width: 1,
 *         },
 *     },
 *     filter: "resource.type=gae_app AND severity>=ERROR",
 *     labelExtractors: {
 *         mass: "EXTRACT(jsonPayload.request)",
 *         sku: "EXTRACT(jsonPayload.id)",
 *     },
 *     metricDescriptor: {
 *         displayName: "My metric",
 *         labels: [
 *             {
 *                 description: "amount of matter",
 *                 key: "mass",
 *                 valueType: "STRING",
 *             },
 *             {
 *                 description: "Identifying number for item",
 *                 key: "sku",
 *                 valueType: "INT64",
 *             },
 *         ],
 *         metricKind: "DELTA",
 *         unit: "1",
 *         valueType: "DISTRIBUTION",
 *     },
 *     valueExtractor: "EXTRACT(jsonPayload.request)",
 * });
 * ```
 * ### Logging Metric Counter Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const loggingMetric = new gcp.logging.Metric("loggingMetric", {
 *     filter: "resource.type=gae_app AND severity>=ERROR",
 *     metricDescriptor: {
 *         metricKind: "DELTA",
 *         valueType: "INT64",
 *     },
 * });
 * ```
 * ### Logging Metric Counter Labels
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const loggingMetric = new gcp.logging.Metric("loggingMetric", {
 *     filter: "resource.type=gae_app AND severity>=ERROR",
 *     labelExtractors: {
 *         mass: "EXTRACT(jsonPayload.request)",
 *     },
 *     metricDescriptor: {
 *         labels: [{
 *             description: "amount of matter",
 *             key: "mass",
 *             valueType: "STRING",
 *         }],
 *         metricKind: "DELTA",
 *         valueType: "INT64",
 *     },
 * });
 * ```
 * ### Logging Metric Logging Bucket
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const loggingMetricProjectBucketConfig = new gcp.logging.ProjectBucketConfig("loggingMetricProjectBucketConfig", {
 *     location: "global",
 *     project: "my-project-name",
 *     bucketId: "_Default",
 * });
 * const loggingMetricMetric = new gcp.logging.Metric("loggingMetricMetric", {
 *     filter: "resource.type=gae_app AND severity>=ERROR",
 *     bucketName: loggingMetricProjectBucketConfig.id,
 * });
 * ```
 * ### Logging Metric Disabled
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const loggingMetric = new gcp.logging.Metric("loggingMetric", {
 *     disabled: true,
 *     filter: "resource.type=gae_app AND severity>=ERROR",
 *     metricDescriptor: {
 *         metricKind: "DELTA",
 *         valueType: "INT64",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Metric can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:logging/metric:Metric default {{project}} {{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:logging/metric:Metric default {{name}}
 * ```
 */
export class Metric extends pulumi.CustomResource {
    /**
     * Get an existing Metric resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MetricState, opts?: pulumi.CustomResourceOptions): Metric {
        return new Metric(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:logging/metric:Metric';

    /**
     * Returns true if the given object is an instance of Metric.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Metric {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Metric.__pulumiType;
    }

    /**
     * The resource name of the Log Bucket that owns the Log Metric. Only Log Buckets in projects
     * are supported. The bucket has to be in the same project as the metric.
     */
    public readonly bucketName!: pulumi.Output<string | undefined>;
    /**
     * The bucketOptions are required when the logs-based metric is using a DISTRIBUTION value type and it
     * describes the bucket boundaries used to create a histogram of the extracted values.
     * Structure is documented below.
     */
    public readonly bucketOptions!: pulumi.Output<outputs.logging.MetricBucketOptions | undefined>;
    /**
     * A description of this metric, which is used in documentation. The maximum length of the
     * description is 8000 characters.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * If set to True, then this metric is disabled and it does not generate any points.
     */
    public readonly disabled!: pulumi.Output<boolean | undefined>;
    /**
     * An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-filters) which
     * is used to match log entries.
     *
     *
     * - - -
     */
    public readonly filter!: pulumi.Output<string>;
    /**
     * A map from a label key string to an extractor expression which is used to extract data from a log
     * entry field and assign as the label value. Each label key specified in the LabelDescriptor must
     * have an associated extractor expression in this map. The syntax of the extractor expression is
     * the same as for the valueExtractor field.
     */
    public readonly labelExtractors!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The optional metric descriptor associated with the logs-based metric.
     * If unspecified, it uses a default metric descriptor with a DELTA metric kind,
     * INT64 value type, with no labels and a unit of "1". Such a metric counts the
     * number of log entries matching the filter expression.
     * Structure is documented below.
     */
    public readonly metricDescriptor!: pulumi.Output<outputs.logging.MetricMetricDescriptor>;
    /**
     * The client-assigned metric identifier. Examples - "errorCount", "nginx/requests".
     * Metric identifiers are limited to 100 characters and can include only the following
     * characters A-Z, a-z, 0-9, and the special characters _-.,+!*',()%/. The forward-slash
     * character (/) denotes a hierarchy of name pieces, and it cannot be the first character
     * of the name.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * A valueExtractor is required when using a distribution logs-based metric to extract the values to
     * record from a log entry. Two functions are supported for value extraction - EXTRACT(field) or
     * REGEXP_EXTRACT(field, regex). The argument are 1. field - The name of the log entry field from which
     * the value is to be extracted. 2. regex - A regular expression using the Google RE2 syntax
     * (https://github.com/google/re2/wiki/Syntax) with a single capture group to extract data from the specified
     * log entry field. The value of the field is converted to a string before applying the regex. It is an
     * error to specify a regex that does not include exactly one capture group.
     */
    public readonly valueExtractor!: pulumi.Output<string | undefined>;

    /**
     * Create a Metric resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MetricArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MetricArgs | MetricState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MetricState | undefined;
            resourceInputs["bucketName"] = state ? state.bucketName : undefined;
            resourceInputs["bucketOptions"] = state ? state.bucketOptions : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["disabled"] = state ? state.disabled : undefined;
            resourceInputs["filter"] = state ? state.filter : undefined;
            resourceInputs["labelExtractors"] = state ? state.labelExtractors : undefined;
            resourceInputs["metricDescriptor"] = state ? state.metricDescriptor : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["valueExtractor"] = state ? state.valueExtractor : undefined;
        } else {
            const args = argsOrState as MetricArgs | undefined;
            if ((!args || args.filter === undefined) && !opts.urn) {
                throw new Error("Missing required property 'filter'");
            }
            resourceInputs["bucketName"] = args ? args.bucketName : undefined;
            resourceInputs["bucketOptions"] = args ? args.bucketOptions : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["disabled"] = args ? args.disabled : undefined;
            resourceInputs["filter"] = args ? args.filter : undefined;
            resourceInputs["labelExtractors"] = args ? args.labelExtractors : undefined;
            resourceInputs["metricDescriptor"] = args ? args.metricDescriptor : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["valueExtractor"] = args ? args.valueExtractor : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Metric.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Metric resources.
 */
export interface MetricState {
    /**
     * The resource name of the Log Bucket that owns the Log Metric. Only Log Buckets in projects
     * are supported. The bucket has to be in the same project as the metric.
     */
    bucketName?: pulumi.Input<string>;
    /**
     * The bucketOptions are required when the logs-based metric is using a DISTRIBUTION value type and it
     * describes the bucket boundaries used to create a histogram of the extracted values.
     * Structure is documented below.
     */
    bucketOptions?: pulumi.Input<inputs.logging.MetricBucketOptions>;
    /**
     * A description of this metric, which is used in documentation. The maximum length of the
     * description is 8000 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * If set to True, then this metric is disabled and it does not generate any points.
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-filters) which
     * is used to match log entries.
     *
     *
     * - - -
     */
    filter?: pulumi.Input<string>;
    /**
     * A map from a label key string to an extractor expression which is used to extract data from a log
     * entry field and assign as the label value. Each label key specified in the LabelDescriptor must
     * have an associated extractor expression in this map. The syntax of the extractor expression is
     * the same as for the valueExtractor field.
     */
    labelExtractors?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The optional metric descriptor associated with the logs-based metric.
     * If unspecified, it uses a default metric descriptor with a DELTA metric kind,
     * INT64 value type, with no labels and a unit of "1". Such a metric counts the
     * number of log entries matching the filter expression.
     * Structure is documented below.
     */
    metricDescriptor?: pulumi.Input<inputs.logging.MetricMetricDescriptor>;
    /**
     * The client-assigned metric identifier. Examples - "errorCount", "nginx/requests".
     * Metric identifiers are limited to 100 characters and can include only the following
     * characters A-Z, a-z, 0-9, and the special characters _-.,+!*',()%/. The forward-slash
     * character (/) denotes a hierarchy of name pieces, and it cannot be the first character
     * of the name.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * A valueExtractor is required when using a distribution logs-based metric to extract the values to
     * record from a log entry. Two functions are supported for value extraction - EXTRACT(field) or
     * REGEXP_EXTRACT(field, regex). The argument are 1. field - The name of the log entry field from which
     * the value is to be extracted. 2. regex - A regular expression using the Google RE2 syntax
     * (https://github.com/google/re2/wiki/Syntax) with a single capture group to extract data from the specified
     * log entry field. The value of the field is converted to a string before applying the regex. It is an
     * error to specify a regex that does not include exactly one capture group.
     */
    valueExtractor?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Metric resource.
 */
export interface MetricArgs {
    /**
     * The resource name of the Log Bucket that owns the Log Metric. Only Log Buckets in projects
     * are supported. The bucket has to be in the same project as the metric.
     */
    bucketName?: pulumi.Input<string>;
    /**
     * The bucketOptions are required when the logs-based metric is using a DISTRIBUTION value type and it
     * describes the bucket boundaries used to create a histogram of the extracted values.
     * Structure is documented below.
     */
    bucketOptions?: pulumi.Input<inputs.logging.MetricBucketOptions>;
    /**
     * A description of this metric, which is used in documentation. The maximum length of the
     * description is 8000 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * If set to True, then this metric is disabled and it does not generate any points.
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * An advanced logs filter (https://cloud.google.com/logging/docs/view/advanced-filters) which
     * is used to match log entries.
     *
     *
     * - - -
     */
    filter: pulumi.Input<string>;
    /**
     * A map from a label key string to an extractor expression which is used to extract data from a log
     * entry field and assign as the label value. Each label key specified in the LabelDescriptor must
     * have an associated extractor expression in this map. The syntax of the extractor expression is
     * the same as for the valueExtractor field.
     */
    labelExtractors?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The optional metric descriptor associated with the logs-based metric.
     * If unspecified, it uses a default metric descriptor with a DELTA metric kind,
     * INT64 value type, with no labels and a unit of "1". Such a metric counts the
     * number of log entries matching the filter expression.
     * Structure is documented below.
     */
    metricDescriptor?: pulumi.Input<inputs.logging.MetricMetricDescriptor>;
    /**
     * The client-assigned metric identifier. Examples - "errorCount", "nginx/requests".
     * Metric identifiers are limited to 100 characters and can include only the following
     * characters A-Z, a-z, 0-9, and the special characters _-.,+!*',()%/. The forward-slash
     * character (/) denotes a hierarchy of name pieces, and it cannot be the first character
     * of the name.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * A valueExtractor is required when using a distribution logs-based metric to extract the values to
     * record from a log entry. Two functions are supported for value extraction - EXTRACT(field) or
     * REGEXP_EXTRACT(field, regex). The argument are 1. field - The name of the log entry field from which
     * the value is to be extracted. 2. regex - A regular expression using the Google RE2 syntax
     * (https://github.com/google/re2/wiki/Syntax) with a single capture group to extract data from the specified
     * log entry field. The value of the field is converted to a string before applying the regex. It is an
     * error to specify a regex that does not include exactly one capture group.
     */
    valueExtractor?: pulumi.Input<string>;
}
