// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 * ### Cloudfunctions2 Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * // [START functions_v2_basic]
 * const bucket = new gcp.storage.Bucket("bucket", {
 *     location: "US",
 *     uniformBucketLevelAccess: true,
 * }, {
 *     provider: google_beta,
 * });
 * const object = new gcp.storage.BucketObject("object", {
 *     bucket: bucket.name,
 *     source: new pulumi.asset.FileAsset("path/to/index.zip"),
 * }, {
 *     provider: google_beta,
 * });
 * // Add path to the zipped function source code
 * const terraform_test2 = new gcp.cloudfunctionsv2.Function("terraform-test2", {
 *     location: "us-central1",
 *     description: "a new function",
 *     buildConfig: {
 *         runtime: "nodejs16",
 *         entryPoint: "helloHttp",
 *         source: {
 *             storageSource: {
 *                 bucket: bucket.name,
 *                 object: object.name,
 *             },
 *         },
 *     },
 *     serviceConfig: {
 *         maxInstanceCount: 1,
 *         availableMemory: "256M",
 *         timeoutSeconds: 60,
 *     },
 * }, {
 *     provider: google_beta,
 * });
 * ```
 * ### Cloudfunctions2 Full
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * // [START functions_v2_full]
 * const account = new gcp.serviceaccount.Account("account", {
 *     accountId: "s-a",
 *     displayName: "Test Service Account",
 * }, {
 *     provider: google_beta,
 * });
 * const sub = new gcp.pubsub.Topic("sub", {}, {
 *     provider: google_beta,
 * });
 * const bucket = new gcp.storage.Bucket("bucket", {
 *     location: "US",
 *     uniformBucketLevelAccess: true,
 * }, {
 *     provider: google_beta,
 * });
 * const object = new gcp.storage.BucketObject("object", {
 *     bucket: bucket.name,
 *     source: new pulumi.asset.FileAsset("path/to/index.zip"),
 * }, {
 *     provider: google_beta,
 * });
 * // Add path to the zipped function source code
 * const terraform_test = new gcp.cloudfunctionsv2.Function("terraform-test", {
 *     location: "us-central1",
 *     description: "a new function",
 *     buildConfig: {
 *         runtime: "nodejs16",
 *         entryPoint: "helloPubSub",
 *         environmentVariables: {
 *             BUILD_CONFIG_TEST: "build_test",
 *         },
 *         source: {
 *             storageSource: {
 *                 bucket: bucket.name,
 *                 object: object.name,
 *             },
 *         },
 *     },
 *     serviceConfig: {
 *         maxInstanceCount: 3,
 *         minInstanceCount: 1,
 *         availableMemory: "256M",
 *         timeoutSeconds: 60,
 *         environmentVariables: {
 *             SERVICE_CONFIG_TEST: "config_test",
 *         },
 *         ingressSettings: "ALLOW_INTERNAL_ONLY",
 *         allTrafficOnLatestRevision: true,
 *         serviceAccountEmail: account.email,
 *     },
 *     eventTrigger: {
 *         triggerRegion: "us-central1",
 *         eventType: "google.cloud.pubsub.topic.v1.messagePublished",
 *         pubsubTopic: sub.id,
 *         retryPolicy: "RETRY_POLICY_RETRY",
 *         serviceAccountEmail: account.email,
 *     },
 * }, {
 *     provider: google_beta,
 * });
 * ```
 *
 * ## Import
 *
 * function can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:cloudfunctionsv2/function:Function default projects/{{project}}/locations/{{location}}/functions/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:cloudfunctionsv2/function:Function default {{project}}/{{location}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:cloudfunctionsv2/function:Function default {{location}}/{{name}}
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
    public static readonly __pulumiType = 'gcp:cloudfunctionsv2/function:Function';

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
     * Describes the Build step of the function that builds a container
     * from the given source.
     * Structure is documented below.
     */
    public readonly buildConfig!: pulumi.Output<outputs.cloudfunctionsv2.FunctionBuildConfig | undefined>;
    /**
     * User-provided description of a function.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The environment the function is hosted on.
     */
    public /*out*/ readonly environment!: pulumi.Output<string>;
    /**
     * An Eventarc trigger managed by Google Cloud Functions that fires events in
     * response to a condition in another service.
     * Structure is documented below.
     */
    public readonly eventTrigger!: pulumi.Output<outputs.cloudfunctionsv2.FunctionEventTrigger | undefined>;
    /**
     * A set of key/value label pairs associated with this Cloud Function.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The location of this cloud function.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * A user-defined name of the function. Function names must
     * be unique globally and match pattern `projects/*&#47;locations/*&#47;functions/*`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Describes the Service being deployed.
     * Structure is documented below.
     */
    public readonly serviceConfig!: pulumi.Output<outputs.cloudfunctionsv2.FunctionServiceConfig | undefined>;
    /**
     * Describes the current state of the function.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The last update timestamp of a Cloud Function.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a Function resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: FunctionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FunctionArgs | FunctionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FunctionState | undefined;
            resourceInputs["buildConfig"] = state ? state.buildConfig : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["environment"] = state ? state.environment : undefined;
            resourceInputs["eventTrigger"] = state ? state.eventTrigger : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["serviceConfig"] = state ? state.serviceConfig : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as FunctionArgs | undefined;
            resourceInputs["buildConfig"] = args ? args.buildConfig : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["eventTrigger"] = args ? args.eventTrigger : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["serviceConfig"] = args ? args.serviceConfig : undefined;
            resourceInputs["environment"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Function.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Function resources.
 */
export interface FunctionState {
    /**
     * Describes the Build step of the function that builds a container
     * from the given source.
     * Structure is documented below.
     */
    buildConfig?: pulumi.Input<inputs.cloudfunctionsv2.FunctionBuildConfig>;
    /**
     * User-provided description of a function.
     */
    description?: pulumi.Input<string>;
    /**
     * The environment the function is hosted on.
     */
    environment?: pulumi.Input<string>;
    /**
     * An Eventarc trigger managed by Google Cloud Functions that fires events in
     * response to a condition in another service.
     * Structure is documented below.
     */
    eventTrigger?: pulumi.Input<inputs.cloudfunctionsv2.FunctionEventTrigger>;
    /**
     * A set of key/value label pairs associated with this Cloud Function.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The location of this cloud function.
     */
    location?: pulumi.Input<string>;
    /**
     * A user-defined name of the function. Function names must
     * be unique globally and match pattern `projects/*&#47;locations/*&#47;functions/*`.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Describes the Service being deployed.
     * Structure is documented below.
     */
    serviceConfig?: pulumi.Input<inputs.cloudfunctionsv2.FunctionServiceConfig>;
    /**
     * Describes the current state of the function.
     */
    state?: pulumi.Input<string>;
    /**
     * The last update timestamp of a Cloud Function.
     */
    updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Function resource.
 */
export interface FunctionArgs {
    /**
     * Describes the Build step of the function that builds a container
     * from the given source.
     * Structure is documented below.
     */
    buildConfig?: pulumi.Input<inputs.cloudfunctionsv2.FunctionBuildConfig>;
    /**
     * User-provided description of a function.
     */
    description?: pulumi.Input<string>;
    /**
     * An Eventarc trigger managed by Google Cloud Functions that fires events in
     * response to a condition in another service.
     * Structure is documented below.
     */
    eventTrigger?: pulumi.Input<inputs.cloudfunctionsv2.FunctionEventTrigger>;
    /**
     * A set of key/value label pairs associated with this Cloud Function.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The location of this cloud function.
     */
    location?: pulumi.Input<string>;
    /**
     * A user-defined name of the function. Function names must
     * be unique globally and match pattern `projects/*&#47;locations/*&#47;functions/*`.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Describes the Service being deployed.
     * Structure is documented below.
     */
    serviceConfig?: pulumi.Input<inputs.cloudfunctionsv2.FunctionServiceConfig>;
}
