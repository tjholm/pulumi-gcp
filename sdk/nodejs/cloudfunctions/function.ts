// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Creates a new Cloud Function. For more information see:
 *
 * * [API documentation](https://cloud.google.com/functions/docs/reference/rest/v1/projects.locations.functions)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/functions/docs)
 *
 * > **Warning:** As of November 1, 2019, newly created Functions are
 * private-by-default and will require [appropriate IAM permissions](https://cloud.google.com/functions/docs/reference/iam/roles)
 * to be invoked. See below examples for how to set up the appropriate permissions,
 * or view the [Cloud Functions IAM resources](https://www.terraform.io/docs/providers/google/r/cloudfunctions_cloud_function_iam.html)
 * for Cloud Functions.
 *
 * ## Example Usage
 * ### Public Function
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const bucket = new gcp.storage.Bucket("bucket", {location: "US"});
 * const archive = new gcp.storage.BucketObject("archive", {
 *     bucket: bucket.name,
 *     source: new pulumi.asset.FileAsset("./path/to/zip/file/which/contains/code"),
 * });
 * const _function = new gcp.cloudfunctions.Function("function", {
 *     description: "My function",
 *     runtime: "nodejs14",
 *     availableMemoryMb: 128,
 *     sourceArchiveBucket: bucket.name,
 *     sourceArchiveObject: archive.name,
 *     triggerHttp: true,
 *     entryPoint: "helloGET",
 * });
 * // IAM entry for all users to invoke the function
 * const invoker = new gcp.cloudfunctions.FunctionIamMember("invoker", {
 *     project: _function.project,
 *     region: _function.region,
 *     cloudFunction: _function.name,
 *     role: "roles/cloudfunctions.invoker",
 *     member: "allUsers",
 * });
 * ```
 * ### Single User
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const bucket = new gcp.storage.Bucket("bucket", {location: "US"});
 * const archive = new gcp.storage.BucketObject("archive", {
 *     bucket: bucket.name,
 *     source: new pulumi.asset.FileAsset("./path/to/zip/file/which/contains/code"),
 * });
 * const _function = new gcp.cloudfunctions.Function("function", {
 *     description: "My function",
 *     runtime: "nodejs14",
 *     availableMemoryMb: 128,
 *     sourceArchiveBucket: bucket.name,
 *     sourceArchiveObject: archive.name,
 *     triggerHttp: true,
 *     timeout: 60,
 *     entryPoint: "helloGET",
 *     labels: {
 *         "my-label": "my-label-value",
 *     },
 *     environmentVariables: {
 *         MY_ENV_VAR: "my-env-var-value",
 *     },
 * });
 * // IAM entry for a single user to invoke the function
 * const invoker = new gcp.cloudfunctions.FunctionIamMember("invoker", {
 *     project: _function.project,
 *     region: _function.region,
 *     cloudFunction: _function.name,
 *     role: "roles/cloudfunctions.invoker",
 *     member: "user:myFunctionInvoker@example.com",
 * });
 * ```
 *
 * ## Import
 *
 * Functions can be imported using the `name` or `{{project}}/{{region}}/name`, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:cloudfunctions/function:Function default function-test
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:cloudfunctions/function:Function default {{project}}/{{region}}/function-test
 * ```
 */
export class Function extends pulumi.CustomResource {
    /**
     * Get an existing Function resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FunctionState, opts?: pulumi.CustomResourceOptions): Function {
        return new Function(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:cloudfunctions/function:Function';

    /**
     * Returns true if the given object is an instance of Function.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Function {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Function.__pulumiType;
    }

    /**
     * Memory (in MB), available to the function. Default value is `256`. Possible values include `128`, `256`, `512`, `1024`, etc.
     */
    public readonly availableMemoryMb!: pulumi.Output<number | undefined>;
    /**
     * A set of key/value environment variable pairs available during build time.
     */
    public readonly buildEnvironmentVariables!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Description of the function.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Name of the function that will be executed when the Google Cloud Function is triggered.
     */
    public readonly entryPoint!: pulumi.Output<string | undefined>;
    /**
     * A set of key/value environment variable pairs to assign to the function.
     */
    public readonly environmentVariables!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * A source that fires events in response to a condition in another service. Structure is documented below. Cannot be used with `triggerHttp`.
     */
    public readonly eventTrigger!: pulumi.Output<outputs.cloudfunctions.FunctionEventTrigger>;
    /**
     * URL which triggers function execution. Returned only if `triggerHttp` is used.
     */
    public readonly httpsTriggerUrl!: pulumi.Output<string>;
    /**
     * String value that controls what traffic can reach the function. Allowed values are `ALLOW_ALL`, `ALLOW_INTERNAL_AND_GCLB` and `ALLOW_INTERNAL_ONLY`. Check [ingress documentation](https://cloud.google.com/functions/docs/networking/network-settings#ingress_settings) to see the impact of each settings value. Changes to this field will recreate the cloud function.
     */
    public readonly ingressSettings!: pulumi.Output<string | undefined>;
    /**
     * A set of key/value label pairs to assign to the function. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
     */
    public readonly labels!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The limit on the maximum number of function instances that may coexist at a given time.
     */
    public readonly maxInstances!: pulumi.Output<number | undefined>;
    /**
     * A user-defined name of the function. Function names must be unique globally.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Project of the function. If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Region of function. If it is not provided, the provider region is used.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The runtime in which the function is going to run.
     * Eg. `"nodejs10"`, `"nodejs12"`, `"nodejs14"`, `"python37"`, `"python38"`, `"python39"`, `"dotnet3"`, `"go113"`, `"java11"`, `"ruby27"`, etc. Check the [official doc](https://cloud.google.com/functions/docs/concepts/exec#runtimes) for the up-to-date list.
     */
    public readonly runtime!: pulumi.Output<string>;
    /**
     * If provided, the self-provided service account to run the function with.
     */
    public readonly serviceAccountEmail!: pulumi.Output<string>;
    /**
     * The GCS bucket containing the zip archive which contains the function.
     */
    public readonly sourceArchiveBucket!: pulumi.Output<string | undefined>;
    /**
     * The source archive object (file) in archive bucket.
     */
    public readonly sourceArchiveObject!: pulumi.Output<string | undefined>;
    /**
     * Represents parameters related to source repository where a function is hosted.
     * Cannot be set alongside `sourceArchiveBucket` or `sourceArchiveObject`. Structure is documented below.
     */
    public readonly sourceRepository!: pulumi.Output<outputs.cloudfunctions.FunctionSourceRepository | undefined>;
    /**
     * Timeout (in seconds) for the function. Default value is 60 seconds. Cannot be more than 540 seconds.
     */
    public readonly timeout!: pulumi.Output<number | undefined>;
    /**
     * Boolean variable. Any HTTP request (of a supported type) to the endpoint will trigger function execution. Supported HTTP request types are: POST, PUT, GET, DELETE, and OPTIONS. Endpoint is returned as `httpsTriggerUrl`. Cannot be used with `eventTrigger`.
     */
    public readonly triggerHttp!: pulumi.Output<boolean | undefined>;
    /**
     * The VPC Network Connector that this cloud function can connect to. It should be set up as fully-qualified URI. The format of this field is `projects/*&#47;locations/*&#47;connectors/*`.
     */
    public readonly vpcConnector!: pulumi.Output<string | undefined>;
    /**
     * The egress settings for the connector, controlling what traffic is diverted through it. Allowed values are `ALL_TRAFFIC` and `PRIVATE_RANGES_ONLY`. Defaults to `PRIVATE_RANGES_ONLY`. If unset, this field preserves the previously set value.
     */
    public readonly vpcConnectorEgressSettings!: pulumi.Output<string>;

    /**
     * Create a Function resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FunctionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FunctionArgs | FunctionState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FunctionState | undefined;
            inputs["availableMemoryMb"] = state ? state.availableMemoryMb : undefined;
            inputs["buildEnvironmentVariables"] = state ? state.buildEnvironmentVariables : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["entryPoint"] = state ? state.entryPoint : undefined;
            inputs["environmentVariables"] = state ? state.environmentVariables : undefined;
            inputs["eventTrigger"] = state ? state.eventTrigger : undefined;
            inputs["httpsTriggerUrl"] = state ? state.httpsTriggerUrl : undefined;
            inputs["ingressSettings"] = state ? state.ingressSettings : undefined;
            inputs["labels"] = state ? state.labels : undefined;
            inputs["maxInstances"] = state ? state.maxInstances : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["region"] = state ? state.region : undefined;
            inputs["runtime"] = state ? state.runtime : undefined;
            inputs["serviceAccountEmail"] = state ? state.serviceAccountEmail : undefined;
            inputs["sourceArchiveBucket"] = state ? state.sourceArchiveBucket : undefined;
            inputs["sourceArchiveObject"] = state ? state.sourceArchiveObject : undefined;
            inputs["sourceRepository"] = state ? state.sourceRepository : undefined;
            inputs["timeout"] = state ? state.timeout : undefined;
            inputs["triggerHttp"] = state ? state.triggerHttp : undefined;
            inputs["vpcConnector"] = state ? state.vpcConnector : undefined;
            inputs["vpcConnectorEgressSettings"] = state ? state.vpcConnectorEgressSettings : undefined;
        } else {
            const args = argsOrState as FunctionArgs | undefined;
            if ((!args || args.runtime === undefined) && !opts.urn) {
                throw new Error("Missing required property 'runtime'");
            }
            inputs["availableMemoryMb"] = args ? args.availableMemoryMb : undefined;
            inputs["buildEnvironmentVariables"] = args ? args.buildEnvironmentVariables : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["entryPoint"] = args ? args.entryPoint : undefined;
            inputs["environmentVariables"] = args ? args.environmentVariables : undefined;
            inputs["eventTrigger"] = args ? args.eventTrigger : undefined;
            inputs["httpsTriggerUrl"] = args ? args.httpsTriggerUrl : undefined;
            inputs["ingressSettings"] = args ? args.ingressSettings : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["maxInstances"] = args ? args.maxInstances : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["runtime"] = args ? args.runtime : undefined;
            inputs["serviceAccountEmail"] = args ? args.serviceAccountEmail : undefined;
            inputs["sourceArchiveBucket"] = args ? args.sourceArchiveBucket : undefined;
            inputs["sourceArchiveObject"] = args ? args.sourceArchiveObject : undefined;
            inputs["sourceRepository"] = args ? args.sourceRepository : undefined;
            inputs["timeout"] = args ? args.timeout : undefined;
            inputs["triggerHttp"] = args ? args.triggerHttp : undefined;
            inputs["vpcConnector"] = args ? args.vpcConnector : undefined;
            inputs["vpcConnectorEgressSettings"] = args ? args.vpcConnectorEgressSettings : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Function.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Function resources.
 */
export interface FunctionState {
    /**
     * Memory (in MB), available to the function. Default value is `256`. Possible values include `128`, `256`, `512`, `1024`, etc.
     */
    availableMemoryMb?: pulumi.Input<number>;
    /**
     * A set of key/value environment variable pairs available during build time.
     */
    buildEnvironmentVariables?: pulumi.Input<{[key: string]: any}>;
    /**
     * Description of the function.
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the function that will be executed when the Google Cloud Function is triggered.
     */
    entryPoint?: pulumi.Input<string>;
    /**
     * A set of key/value environment variable pairs to assign to the function.
     */
    environmentVariables?: pulumi.Input<{[key: string]: any}>;
    /**
     * A source that fires events in response to a condition in another service. Structure is documented below. Cannot be used with `triggerHttp`.
     */
    eventTrigger?: pulumi.Input<inputs.cloudfunctions.FunctionEventTrigger>;
    /**
     * URL which triggers function execution. Returned only if `triggerHttp` is used.
     */
    httpsTriggerUrl?: pulumi.Input<string>;
    /**
     * String value that controls what traffic can reach the function. Allowed values are `ALLOW_ALL`, `ALLOW_INTERNAL_AND_GCLB` and `ALLOW_INTERNAL_ONLY`. Check [ingress documentation](https://cloud.google.com/functions/docs/networking/network-settings#ingress_settings) to see the impact of each settings value. Changes to this field will recreate the cloud function.
     */
    ingressSettings?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs to assign to the function. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
     */
    labels?: pulumi.Input<{[key: string]: any}>;
    /**
     * The limit on the maximum number of function instances that may coexist at a given time.
     */
    maxInstances?: pulumi.Input<number>;
    /**
     * A user-defined name of the function. Function names must be unique globally.
     */
    name?: pulumi.Input<string>;
    /**
     * Project of the function. If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Region of function. If it is not provided, the provider region is used.
     */
    region?: pulumi.Input<string>;
    /**
     * The runtime in which the function is going to run.
     * Eg. `"nodejs10"`, `"nodejs12"`, `"nodejs14"`, `"python37"`, `"python38"`, `"python39"`, `"dotnet3"`, `"go113"`, `"java11"`, `"ruby27"`, etc. Check the [official doc](https://cloud.google.com/functions/docs/concepts/exec#runtimes) for the up-to-date list.
     */
    runtime?: pulumi.Input<string>;
    /**
     * If provided, the self-provided service account to run the function with.
     */
    serviceAccountEmail?: pulumi.Input<string>;
    /**
     * The GCS bucket containing the zip archive which contains the function.
     */
    sourceArchiveBucket?: pulumi.Input<string>;
    /**
     * The source archive object (file) in archive bucket.
     */
    sourceArchiveObject?: pulumi.Input<string>;
    /**
     * Represents parameters related to source repository where a function is hosted.
     * Cannot be set alongside `sourceArchiveBucket` or `sourceArchiveObject`. Structure is documented below.
     */
    sourceRepository?: pulumi.Input<inputs.cloudfunctions.FunctionSourceRepository>;
    /**
     * Timeout (in seconds) for the function. Default value is 60 seconds. Cannot be more than 540 seconds.
     */
    timeout?: pulumi.Input<number>;
    /**
     * Boolean variable. Any HTTP request (of a supported type) to the endpoint will trigger function execution. Supported HTTP request types are: POST, PUT, GET, DELETE, and OPTIONS. Endpoint is returned as `httpsTriggerUrl`. Cannot be used with `eventTrigger`.
     */
    triggerHttp?: pulumi.Input<boolean>;
    /**
     * The VPC Network Connector that this cloud function can connect to. It should be set up as fully-qualified URI. The format of this field is `projects/*&#47;locations/*&#47;connectors/*`.
     */
    vpcConnector?: pulumi.Input<string>;
    /**
     * The egress settings for the connector, controlling what traffic is diverted through it. Allowed values are `ALL_TRAFFIC` and `PRIVATE_RANGES_ONLY`. Defaults to `PRIVATE_RANGES_ONLY`. If unset, this field preserves the previously set value.
     */
    vpcConnectorEgressSettings?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Function resource.
 */
export interface FunctionArgs {
    /**
     * Memory (in MB), available to the function. Default value is `256`. Possible values include `128`, `256`, `512`, `1024`, etc.
     */
    availableMemoryMb?: pulumi.Input<number>;
    /**
     * A set of key/value environment variable pairs available during build time.
     */
    buildEnvironmentVariables?: pulumi.Input<{[key: string]: any}>;
    /**
     * Description of the function.
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the function that will be executed when the Google Cloud Function is triggered.
     */
    entryPoint?: pulumi.Input<string>;
    /**
     * A set of key/value environment variable pairs to assign to the function.
     */
    environmentVariables?: pulumi.Input<{[key: string]: any}>;
    /**
     * A source that fires events in response to a condition in another service. Structure is documented below. Cannot be used with `triggerHttp`.
     */
    eventTrigger?: pulumi.Input<inputs.cloudfunctions.FunctionEventTrigger>;
    /**
     * URL which triggers function execution. Returned only if `triggerHttp` is used.
     */
    httpsTriggerUrl?: pulumi.Input<string>;
    /**
     * String value that controls what traffic can reach the function. Allowed values are `ALLOW_ALL`, `ALLOW_INTERNAL_AND_GCLB` and `ALLOW_INTERNAL_ONLY`. Check [ingress documentation](https://cloud.google.com/functions/docs/networking/network-settings#ingress_settings) to see the impact of each settings value. Changes to this field will recreate the cloud function.
     */
    ingressSettings?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs to assign to the function. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
     */
    labels?: pulumi.Input<{[key: string]: any}>;
    /**
     * The limit on the maximum number of function instances that may coexist at a given time.
     */
    maxInstances?: pulumi.Input<number>;
    /**
     * A user-defined name of the function. Function names must be unique globally.
     */
    name?: pulumi.Input<string>;
    /**
     * Project of the function. If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Region of function. If it is not provided, the provider region is used.
     */
    region?: pulumi.Input<string>;
    /**
     * The runtime in which the function is going to run.
     * Eg. `"nodejs10"`, `"nodejs12"`, `"nodejs14"`, `"python37"`, `"python38"`, `"python39"`, `"dotnet3"`, `"go113"`, `"java11"`, `"ruby27"`, etc. Check the [official doc](https://cloud.google.com/functions/docs/concepts/exec#runtimes) for the up-to-date list.
     */
    runtime: pulumi.Input<string>;
    /**
     * If provided, the self-provided service account to run the function with.
     */
    serviceAccountEmail?: pulumi.Input<string>;
    /**
     * The GCS bucket containing the zip archive which contains the function.
     */
    sourceArchiveBucket?: pulumi.Input<string>;
    /**
     * The source archive object (file) in archive bucket.
     */
    sourceArchiveObject?: pulumi.Input<string>;
    /**
     * Represents parameters related to source repository where a function is hosted.
     * Cannot be set alongside `sourceArchiveBucket` or `sourceArchiveObject`. Structure is documented below.
     */
    sourceRepository?: pulumi.Input<inputs.cloudfunctions.FunctionSourceRepository>;
    /**
     * Timeout (in seconds) for the function. Default value is 60 seconds. Cannot be more than 540 seconds.
     */
    timeout?: pulumi.Input<number>;
    /**
     * Boolean variable. Any HTTP request (of a supported type) to the endpoint will trigger function execution. Supported HTTP request types are: POST, PUT, GET, DELETE, and OPTIONS. Endpoint is returned as `httpsTriggerUrl`. Cannot be used with `eventTrigger`.
     */
    triggerHttp?: pulumi.Input<boolean>;
    /**
     * The VPC Network Connector that this cloud function can connect to. It should be set up as fully-qualified URI. The format of this field is `projects/*&#47;locations/*&#47;connectors/*`.
     */
    vpcConnector?: pulumi.Input<string>;
    /**
     * The egress settings for the connector, controlling what traffic is diverted through it. Allowed values are `ALL_TRAFFIC` and `PRIVATE_RANGES_ONLY`. Defaults to `PRIVATE_RANGES_ONLY`. If unset, this field preserves the previously set value.
     */
    vpcConnectorEgressSettings?: pulumi.Input<string>;
}
