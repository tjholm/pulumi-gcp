// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
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
 *     runtime: "nodejs16",
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
 *     runtime: "nodejs16",
 *     availableMemoryMb: 128,
 *     sourceArchiveBucket: bucket.name,
 *     sourceArchiveObject: archive.name,
 *     triggerHttp: true,
 *     httpsTriggerSecurityLevel: "SECURE_ALWAYS",
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
     * Name of the Cloud Build Custom Worker Pool that should be used to build the function.
     */
    public readonly buildWorkerPool!: pulumi.Output<string | undefined>;
    /**
     * Description of the function.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Docker Registry to use for storing the function's Docker images. Allowed values are CONTAINER_REGISTRY (default) and ARTIFACT_REGISTRY.
     */
    public readonly dockerRegistry!: pulumi.Output<string>;
    /**
     * User managed repository created in Artifact Registry optionally with a customer managed encryption key. If specified, deployments will use Artifact Registry. This is the repository to which the function docker image will be pushed after it is built by Cloud Build. If unspecified, Container Registry will be used by default, unless specified otherwise by other means.
     */
    public readonly dockerRepository!: pulumi.Output<string | undefined>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
     * clients and services.
     */
    public /*out*/ readonly effectiveLabels!: pulumi.Output<{[key: string]: string}>;
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
     * The security level for the function. The following options are available:
     */
    public readonly httpsTriggerSecurityLevel!: pulumi.Output<string>;
    /**
     * URL which triggers function execution. Returned only if `triggerHttp` is used.
     */
    public readonly httpsTriggerUrl!: pulumi.Output<string>;
    /**
     * String value that controls what traffic can reach the function. Allowed values are `ALLOW_ALL`, `ALLOW_INTERNAL_AND_GCLB` and `ALLOW_INTERNAL_ONLY`. Check [ingress documentation](https://cloud.google.com/functions/docs/networking/network-settings#ingress_settings) to see the impact of each settings value. Changes to this field will recreate the cloud function.
     */
    public readonly ingressSettings!: pulumi.Output<string | undefined>;
    /**
     * Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt function resources. It must match the pattern `projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}`.
     * If specified, you must also provide an artifact registry repository using the `dockerRepository` field that was created with the same KMS crypto key. Before deploying, please complete all pre-requisites described in https://cloud.google.com/functions/docs/securing/cmek#granting_service_accounts_access_to_the_key
     */
    public readonly kmsKeyName!: pulumi.Output<string | undefined>;
    /**
     * A set of key/value label pairs to assign to the function. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field 'effective_labels' for all of the labels present on the resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The limit on the maximum number of function instances that may coexist at a given time.
     */
    public readonly maxInstances!: pulumi.Output<number>;
    /**
     * The limit on the minimum number of function instances that may coexist at a given time.
     */
    public readonly minInstances!: pulumi.Output<number | undefined>;
    /**
     * A user-defined name of the function. Function names must be unique globally.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Project of the function. If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The combination of labels configured directly on the resource and default labels configured on the provider.
     */
    public /*out*/ readonly pulumiLabels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Region of function. If it is not provided, the provider region is used.
     */
    public readonly region!: pulumi.Output<string>;
    /**
     * The runtime in which the function is going to run.
     * Eg. `"nodejs16"`, `"python39"`, `"dotnet3"`, `"go116"`, `"java11"`, `"ruby30"`, `"php74"`, etc. Check the [official doc](https://cloud.google.com/functions/docs/concepts/exec#runtimes) for the up-to-date list.
     *
     * - - -
     */
    public readonly runtime!: pulumi.Output<string>;
    /**
     * Secret environment variables configuration. Structure is documented below.
     */
    public readonly secretEnvironmentVariables!: pulumi.Output<outputs.cloudfunctions.FunctionSecretEnvironmentVariable[] | undefined>;
    /**
     * Secret volumes configuration. Structure is documented below.
     */
    public readonly secretVolumes!: pulumi.Output<outputs.cloudfunctions.FunctionSecretVolume[] | undefined>;
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
     * Cannot be set alongside `sourceArchiveBucket` or `sourceArchiveObject`. Structure is documented below. It must match the pattern `projects/{project}/locations/{location}/repositories/{repository}`.*
     */
    public readonly sourceRepository!: pulumi.Output<outputs.cloudfunctions.FunctionSourceRepository | undefined>;
    /**
     * Describes the current stage of a deployment.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
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
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FunctionState | undefined;
            resourceInputs["availableMemoryMb"] = state ? state.availableMemoryMb : undefined;
            resourceInputs["buildEnvironmentVariables"] = state ? state.buildEnvironmentVariables : undefined;
            resourceInputs["buildWorkerPool"] = state ? state.buildWorkerPool : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["dockerRegistry"] = state ? state.dockerRegistry : undefined;
            resourceInputs["dockerRepository"] = state ? state.dockerRepository : undefined;
            resourceInputs["effectiveLabels"] = state ? state.effectiveLabels : undefined;
            resourceInputs["entryPoint"] = state ? state.entryPoint : undefined;
            resourceInputs["environmentVariables"] = state ? state.environmentVariables : undefined;
            resourceInputs["eventTrigger"] = state ? state.eventTrigger : undefined;
            resourceInputs["httpsTriggerSecurityLevel"] = state ? state.httpsTriggerSecurityLevel : undefined;
            resourceInputs["httpsTriggerUrl"] = state ? state.httpsTriggerUrl : undefined;
            resourceInputs["ingressSettings"] = state ? state.ingressSettings : undefined;
            resourceInputs["kmsKeyName"] = state ? state.kmsKeyName : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["maxInstances"] = state ? state.maxInstances : undefined;
            resourceInputs["minInstances"] = state ? state.minInstances : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["pulumiLabels"] = state ? state.pulumiLabels : undefined;
            resourceInputs["region"] = state ? state.region : undefined;
            resourceInputs["runtime"] = state ? state.runtime : undefined;
            resourceInputs["secretEnvironmentVariables"] = state ? state.secretEnvironmentVariables : undefined;
            resourceInputs["secretVolumes"] = state ? state.secretVolumes : undefined;
            resourceInputs["serviceAccountEmail"] = state ? state.serviceAccountEmail : undefined;
            resourceInputs["sourceArchiveBucket"] = state ? state.sourceArchiveBucket : undefined;
            resourceInputs["sourceArchiveObject"] = state ? state.sourceArchiveObject : undefined;
            resourceInputs["sourceRepository"] = state ? state.sourceRepository : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["timeout"] = state ? state.timeout : undefined;
            resourceInputs["triggerHttp"] = state ? state.triggerHttp : undefined;
            resourceInputs["vpcConnector"] = state ? state.vpcConnector : undefined;
            resourceInputs["vpcConnectorEgressSettings"] = state ? state.vpcConnectorEgressSettings : undefined;
        } else {
            const args = argsOrState as FunctionArgs | undefined;
            if ((!args || args.runtime === undefined) && !opts.urn) {
                throw new Error("Missing required property 'runtime'");
            }
            resourceInputs["availableMemoryMb"] = args ? args.availableMemoryMb : undefined;
            resourceInputs["buildEnvironmentVariables"] = args ? args.buildEnvironmentVariables : undefined;
            resourceInputs["buildWorkerPool"] = args ? args.buildWorkerPool : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["dockerRegistry"] = args ? args.dockerRegistry : undefined;
            resourceInputs["dockerRepository"] = args ? args.dockerRepository : undefined;
            resourceInputs["entryPoint"] = args ? args.entryPoint : undefined;
            resourceInputs["environmentVariables"] = args ? args.environmentVariables : undefined;
            resourceInputs["eventTrigger"] = args ? args.eventTrigger : undefined;
            resourceInputs["httpsTriggerSecurityLevel"] = args ? args.httpsTriggerSecurityLevel : undefined;
            resourceInputs["httpsTriggerUrl"] = args ? args.httpsTriggerUrl : undefined;
            resourceInputs["ingressSettings"] = args ? args.ingressSettings : undefined;
            resourceInputs["kmsKeyName"] = args ? args.kmsKeyName : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["maxInstances"] = args ? args.maxInstances : undefined;
            resourceInputs["minInstances"] = args ? args.minInstances : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["runtime"] = args ? args.runtime : undefined;
            resourceInputs["secretEnvironmentVariables"] = args ? args.secretEnvironmentVariables : undefined;
            resourceInputs["secretVolumes"] = args ? args.secretVolumes : undefined;
            resourceInputs["serviceAccountEmail"] = args ? args.serviceAccountEmail : undefined;
            resourceInputs["sourceArchiveBucket"] = args ? args.sourceArchiveBucket : undefined;
            resourceInputs["sourceArchiveObject"] = args ? args.sourceArchiveObject : undefined;
            resourceInputs["sourceRepository"] = args ? args.sourceRepository : undefined;
            resourceInputs["timeout"] = args ? args.timeout : undefined;
            resourceInputs["triggerHttp"] = args ? args.triggerHttp : undefined;
            resourceInputs["vpcConnector"] = args ? args.vpcConnector : undefined;
            resourceInputs["vpcConnectorEgressSettings"] = args ? args.vpcConnectorEgressSettings : undefined;
            resourceInputs["effectiveLabels"] = undefined /*out*/;
            resourceInputs["pulumiLabels"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["effectiveLabels", "pulumiLabels"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Function.__pulumiType, name, resourceInputs, opts);
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
     * Name of the Cloud Build Custom Worker Pool that should be used to build the function.
     */
    buildWorkerPool?: pulumi.Input<string>;
    /**
     * Description of the function.
     */
    description?: pulumi.Input<string>;
    /**
     * Docker Registry to use for storing the function's Docker images. Allowed values are CONTAINER_REGISTRY (default) and ARTIFACT_REGISTRY.
     */
    dockerRegistry?: pulumi.Input<string>;
    /**
     * User managed repository created in Artifact Registry optionally with a customer managed encryption key. If specified, deployments will use Artifact Registry. This is the repository to which the function docker image will be pushed after it is built by Cloud Build. If unspecified, Container Registry will be used by default, unless specified otherwise by other means.
     */
    dockerRepository?: pulumi.Input<string>;
    /**
     * All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other
     * clients and services.
     */
    effectiveLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
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
     * The security level for the function. The following options are available:
     */
    httpsTriggerSecurityLevel?: pulumi.Input<string>;
    /**
     * URL which triggers function execution. Returned only if `triggerHttp` is used.
     */
    httpsTriggerUrl?: pulumi.Input<string>;
    /**
     * String value that controls what traffic can reach the function. Allowed values are `ALLOW_ALL`, `ALLOW_INTERNAL_AND_GCLB` and `ALLOW_INTERNAL_ONLY`. Check [ingress documentation](https://cloud.google.com/functions/docs/networking/network-settings#ingress_settings) to see the impact of each settings value. Changes to this field will recreate the cloud function.
     */
    ingressSettings?: pulumi.Input<string>;
    /**
     * Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt function resources. It must match the pattern `projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}`.
     * If specified, you must also provide an artifact registry repository using the `dockerRepository` field that was created with the same KMS crypto key. Before deploying, please complete all pre-requisites described in https://cloud.google.com/functions/docs/securing/cmek#granting_service_accounts_access_to_the_key
     */
    kmsKeyName?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs to assign to the function. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field 'effective_labels' for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: any}>;
    /**
     * The limit on the maximum number of function instances that may coexist at a given time.
     */
    maxInstances?: pulumi.Input<number>;
    /**
     * The limit on the minimum number of function instances that may coexist at a given time.
     */
    minInstances?: pulumi.Input<number>;
    /**
     * A user-defined name of the function. Function names must be unique globally.
     */
    name?: pulumi.Input<string>;
    /**
     * Project of the function. If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The combination of labels configured directly on the resource and default labels configured on the provider.
     */
    pulumiLabels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Region of function. If it is not provided, the provider region is used.
     */
    region?: pulumi.Input<string>;
    /**
     * The runtime in which the function is going to run.
     * Eg. `"nodejs16"`, `"python39"`, `"dotnet3"`, `"go116"`, `"java11"`, `"ruby30"`, `"php74"`, etc. Check the [official doc](https://cloud.google.com/functions/docs/concepts/exec#runtimes) for the up-to-date list.
     *
     * - - -
     */
    runtime?: pulumi.Input<string>;
    /**
     * Secret environment variables configuration. Structure is documented below.
     */
    secretEnvironmentVariables?: pulumi.Input<pulumi.Input<inputs.cloudfunctions.FunctionSecretEnvironmentVariable>[]>;
    /**
     * Secret volumes configuration. Structure is documented below.
     */
    secretVolumes?: pulumi.Input<pulumi.Input<inputs.cloudfunctions.FunctionSecretVolume>[]>;
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
     * Cannot be set alongside `sourceArchiveBucket` or `sourceArchiveObject`. Structure is documented below. It must match the pattern `projects/{project}/locations/{location}/repositories/{repository}`.*
     */
    sourceRepository?: pulumi.Input<inputs.cloudfunctions.FunctionSourceRepository>;
    /**
     * Describes the current stage of a deployment.
     */
    status?: pulumi.Input<string>;
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
     * Name of the Cloud Build Custom Worker Pool that should be used to build the function.
     */
    buildWorkerPool?: pulumi.Input<string>;
    /**
     * Description of the function.
     */
    description?: pulumi.Input<string>;
    /**
     * Docker Registry to use for storing the function's Docker images. Allowed values are CONTAINER_REGISTRY (default) and ARTIFACT_REGISTRY.
     */
    dockerRegistry?: pulumi.Input<string>;
    /**
     * User managed repository created in Artifact Registry optionally with a customer managed encryption key. If specified, deployments will use Artifact Registry. This is the repository to which the function docker image will be pushed after it is built by Cloud Build. If unspecified, Container Registry will be used by default, unless specified otherwise by other means.
     */
    dockerRepository?: pulumi.Input<string>;
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
     * The security level for the function. The following options are available:
     */
    httpsTriggerSecurityLevel?: pulumi.Input<string>;
    /**
     * URL which triggers function execution. Returned only if `triggerHttp` is used.
     */
    httpsTriggerUrl?: pulumi.Input<string>;
    /**
     * String value that controls what traffic can reach the function. Allowed values are `ALLOW_ALL`, `ALLOW_INTERNAL_AND_GCLB` and `ALLOW_INTERNAL_ONLY`. Check [ingress documentation](https://cloud.google.com/functions/docs/networking/network-settings#ingress_settings) to see the impact of each settings value. Changes to this field will recreate the cloud function.
     */
    ingressSettings?: pulumi.Input<string>;
    /**
     * Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt function resources. It must match the pattern `projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}`.
     * If specified, you must also provide an artifact registry repository using the `dockerRepository` field that was created with the same KMS crypto key. Before deploying, please complete all pre-requisites described in https://cloud.google.com/functions/docs/securing/cmek#granting_service_accounts_access_to_the_key
     */
    kmsKeyName?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs to assign to the function. Label keys must follow the requirements at https://cloud.google.com/resource-manager/docs/creating-managing-labels#requirements.
     *
     * **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
     * Please refer to the field 'effective_labels' for all of the labels present on the resource.
     */
    labels?: pulumi.Input<{[key: string]: any}>;
    /**
     * The limit on the maximum number of function instances that may coexist at a given time.
     */
    maxInstances?: pulumi.Input<number>;
    /**
     * The limit on the minimum number of function instances that may coexist at a given time.
     */
    minInstances?: pulumi.Input<number>;
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
     * Eg. `"nodejs16"`, `"python39"`, `"dotnet3"`, `"go116"`, `"java11"`, `"ruby30"`, `"php74"`, etc. Check the [official doc](https://cloud.google.com/functions/docs/concepts/exec#runtimes) for the up-to-date list.
     *
     * - - -
     */
    runtime: pulumi.Input<string>;
    /**
     * Secret environment variables configuration. Structure is documented below.
     */
    secretEnvironmentVariables?: pulumi.Input<pulumi.Input<inputs.cloudfunctions.FunctionSecretEnvironmentVariable>[]>;
    /**
     * Secret volumes configuration. Structure is documented below.
     */
    secretVolumes?: pulumi.Input<pulumi.Input<inputs.cloudfunctions.FunctionSecretVolume>[]>;
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
     * Cannot be set alongside `sourceArchiveBucket` or `sourceArchiveObject`. Structure is documented below. It must match the pattern `projects/{project}/locations/{location}/repositories/{repository}`.*
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
