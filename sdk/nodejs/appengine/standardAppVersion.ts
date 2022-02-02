// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Standard App Version resource to create a new version of standard GAE Application.
 * Learn about the differences between the standard environment and the flexible environment
 * at https://cloud.google.com/appengine/docs/the-appengine-environments.
 * Currently supporting Zip and File Containers.
 *
 * To get more information about StandardAppVersion, see:
 *
 * * [API documentation](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/appengine/docs/standard)
 *
 * ## Example Usage
 * ### App Engine Standard App Version
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const bucket = new gcp.storage.Bucket("bucket", {location: "US"});
 * const object = new gcp.storage.BucketObject("object", {
 *     bucket: bucket.name,
 *     source: new pulumi.asset.FileAsset("./test-fixtures/appengine/hello-world.zip"),
 * });
 * const myappV1 = new gcp.appengine.StandardAppVersion("myappV1", {
 *     versionId: "v1",
 *     service: "myapp",
 *     runtime: "nodejs10",
 *     entrypoint: {
 *         shell: "node ./app.js",
 *     },
 *     deployment: {
 *         zip: {
 *             sourceUrl: pulumi.interpolate`https://storage.googleapis.com/${bucket.name}/${object.name}`,
 *         },
 *     },
 *     envVariables: {
 *         port: "8080",
 *     },
 *     automaticScaling: {
 *         maxConcurrentRequests: 10,
 *         minIdleInstances: 1,
 *         maxIdleInstances: 3,
 *         minPendingLatency: "1s",
 *         maxPendingLatency: "5s",
 *         standardSchedulerSettings: {
 *             targetCpuUtilization: 0.5,
 *             targetThroughputUtilization: 0.75,
 *             minInstances: 2,
 *             maxInstances: 10,
 *         },
 *     },
 *     deleteServiceOnDestroy: true,
 * });
 * const myappV2 = new gcp.appengine.StandardAppVersion("myappV2", {
 *     versionId: "v2",
 *     service: "myapp",
 *     runtime: "nodejs10",
 *     entrypoint: {
 *         shell: "node ./app.js",
 *     },
 *     deployment: {
 *         zip: {
 *             sourceUrl: pulumi.interpolate`https://storage.googleapis.com/${bucket.name}/${object.name}`,
 *         },
 *     },
 *     envVariables: {
 *         port: "8080",
 *     },
 *     basicScaling: {
 *         maxInstances: 5,
 *     },
 *     noopOnDestroy: true,
 * });
 * ```
 *
 * ## Import
 *
 * StandardAppVersion can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:appengine/standardAppVersion:StandardAppVersion default apps/{{project}}/services/{{service}}/versions/{{version_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:appengine/standardAppVersion:StandardAppVersion default {{project}}/{{service}}/{{version_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:appengine/standardAppVersion:StandardAppVersion default {{service}}/{{version_id}}
 * ```
 */
export class StandardAppVersion extends pulumi.CustomResource {
    /**
     * Get an existing StandardAppVersion resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: StandardAppVersionState, opts?: pulumi.CustomResourceOptions): StandardAppVersion {
        return new StandardAppVersion(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:appengine/standardAppVersion:StandardAppVersion';

    /**
     * Returns true if the given object is an instance of StandardAppVersion.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is StandardAppVersion {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === StandardAppVersion.__pulumiType;
    }

    /**
     * Automatic scaling is based on request rate, response latencies, and other application metrics.
     * Structure is documented below.
     */
    public readonly automaticScaling!: pulumi.Output<outputs.appengine.StandardAppVersionAutomaticScaling | undefined>;
    /**
     * Basic scaling creates instances when your application receives requests. Each instance will be shut down when the application becomes idle. Basic scaling is ideal for work that is intermittent or driven by user activity.
     * Structure is documented below.
     */
    public readonly basicScaling!: pulumi.Output<outputs.appengine.StandardAppVersionBasicScaling | undefined>;
    /**
     * If set to `true`, the service will be deleted if it is the last version.
     */
    public readonly deleteServiceOnDestroy!: pulumi.Output<boolean | undefined>;
    /**
     * Code and application artifacts that make up this version.
     * Structure is documented below.
     */
    public readonly deployment!: pulumi.Output<outputs.appengine.StandardAppVersionDeployment>;
    /**
     * The entrypoint for the application.
     * Structure is documented below.
     */
    public readonly entrypoint!: pulumi.Output<outputs.appengine.StandardAppVersionEntrypoint>;
    /**
     * Environment variables available to the application.
     */
    public readonly envVariables!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * An ordered list of URL-matching patterns that should be applied to incoming requests.
     * The first matching URL handles the request and other request handlers are not attempted.
     * Structure is documented below.
     */
    public readonly handlers!: pulumi.Output<outputs.appengine.StandardAppVersionHandler[]>;
    /**
     * A list of the types of messages that this application is able to receive.
     * Each value may be one of `INBOUND_SERVICE_MAIL`, `INBOUND_SERVICE_MAIL_BOUNCE`, `INBOUND_SERVICE_XMPP_ERROR`, `INBOUND_SERVICE_XMPP_MESSAGE`, `INBOUND_SERVICE_XMPP_SUBSCRIBE`, `INBOUND_SERVICE_XMPP_PRESENCE`, `INBOUND_SERVICE_CHANNEL_PRESENCE`, and `INBOUND_SERVICE_WARMUP`.
     */
    public readonly inboundServices!: pulumi.Output<string[] | undefined>;
    /**
     * Instance class that is used to run this version. Valid values are
     * AutomaticScaling: F1, F2, F4, F4_1G
     * BasicScaling or ManualScaling: B1, B2, B4, B4_1G, B8
     * Defaults to F1 for AutomaticScaling and B2 for ManualScaling and BasicScaling. If no scaling is specified, AutomaticScaling is chosen.
     */
    public readonly instanceClass!: pulumi.Output<string>;
    /**
     * Configuration for third-party Python runtime libraries that are required by the application.
     * Structure is documented below.
     */
    public readonly libraries!: pulumi.Output<outputs.appengine.StandardAppVersionLibrary[] | undefined>;
    /**
     * A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of its memory over time.
     * Structure is documented below.
     */
    public readonly manualScaling!: pulumi.Output<outputs.appengine.StandardAppVersionManualScaling | undefined>;
    /**
     * Full Serverless VPC Access Connector name e.g. /projects/my-project/locations/us-central1/connectors/c1.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * If set to `true`, the application version will not be deleted.
     */
    public readonly noopOnDestroy!: pulumi.Output<boolean | undefined>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Desired runtime. Example python27.
     */
    public readonly runtime!: pulumi.Output<string>;
    /**
     * The version of the API in the given runtime environment.
     * Please see the app.yaml reference for valid values at `https://cloud.google.com/appengine/docs/standard/<language>/config/appref`\
     * Substitute `<language>` with `python`, `java`, `php`, `ruby`, `go` or `nodejs`.
     */
    public readonly runtimeApiVersion!: pulumi.Output<string | undefined>;
    /**
     * AppEngine service resource
     */
    public readonly service!: pulumi.Output<string>;
    /**
     * Whether multiple requests can be dispatched to this version at once.
     */
    public readonly threadsafe!: pulumi.Output<boolean | undefined>;
    /**
     * Relative name of the version within the service. For example, `v1`. Version names can contain only lowercase letters, numbers, or hyphens. Reserved names,"default", "latest", and any name with the prefix "ah-".
     */
    public readonly versionId!: pulumi.Output<string | undefined>;
    /**
     * Enables VPC connectivity for standard apps.
     * Structure is documented below.
     */
    public readonly vpcAccessConnector!: pulumi.Output<outputs.appengine.StandardAppVersionVpcAccessConnector | undefined>;

    /**
     * Create a StandardAppVersion resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: StandardAppVersionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: StandardAppVersionArgs | StandardAppVersionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as StandardAppVersionState | undefined;
            resourceInputs["automaticScaling"] = state ? state.automaticScaling : undefined;
            resourceInputs["basicScaling"] = state ? state.basicScaling : undefined;
            resourceInputs["deleteServiceOnDestroy"] = state ? state.deleteServiceOnDestroy : undefined;
            resourceInputs["deployment"] = state ? state.deployment : undefined;
            resourceInputs["entrypoint"] = state ? state.entrypoint : undefined;
            resourceInputs["envVariables"] = state ? state.envVariables : undefined;
            resourceInputs["handlers"] = state ? state.handlers : undefined;
            resourceInputs["inboundServices"] = state ? state.inboundServices : undefined;
            resourceInputs["instanceClass"] = state ? state.instanceClass : undefined;
            resourceInputs["libraries"] = state ? state.libraries : undefined;
            resourceInputs["manualScaling"] = state ? state.manualScaling : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["noopOnDestroy"] = state ? state.noopOnDestroy : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["runtime"] = state ? state.runtime : undefined;
            resourceInputs["runtimeApiVersion"] = state ? state.runtimeApiVersion : undefined;
            resourceInputs["service"] = state ? state.service : undefined;
            resourceInputs["threadsafe"] = state ? state.threadsafe : undefined;
            resourceInputs["versionId"] = state ? state.versionId : undefined;
            resourceInputs["vpcAccessConnector"] = state ? state.vpcAccessConnector : undefined;
        } else {
            const args = argsOrState as StandardAppVersionArgs | undefined;
            if ((!args || args.deployment === undefined) && !opts.urn) {
                throw new Error("Missing required property 'deployment'");
            }
            if ((!args || args.entrypoint === undefined) && !opts.urn) {
                throw new Error("Missing required property 'entrypoint'");
            }
            if ((!args || args.runtime === undefined) && !opts.urn) {
                throw new Error("Missing required property 'runtime'");
            }
            if ((!args || args.service === undefined) && !opts.urn) {
                throw new Error("Missing required property 'service'");
            }
            resourceInputs["automaticScaling"] = args ? args.automaticScaling : undefined;
            resourceInputs["basicScaling"] = args ? args.basicScaling : undefined;
            resourceInputs["deleteServiceOnDestroy"] = args ? args.deleteServiceOnDestroy : undefined;
            resourceInputs["deployment"] = args ? args.deployment : undefined;
            resourceInputs["entrypoint"] = args ? args.entrypoint : undefined;
            resourceInputs["envVariables"] = args ? args.envVariables : undefined;
            resourceInputs["handlers"] = args ? args.handlers : undefined;
            resourceInputs["inboundServices"] = args ? args.inboundServices : undefined;
            resourceInputs["instanceClass"] = args ? args.instanceClass : undefined;
            resourceInputs["libraries"] = args ? args.libraries : undefined;
            resourceInputs["manualScaling"] = args ? args.manualScaling : undefined;
            resourceInputs["noopOnDestroy"] = args ? args.noopOnDestroy : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["runtime"] = args ? args.runtime : undefined;
            resourceInputs["runtimeApiVersion"] = args ? args.runtimeApiVersion : undefined;
            resourceInputs["service"] = args ? args.service : undefined;
            resourceInputs["threadsafe"] = args ? args.threadsafe : undefined;
            resourceInputs["versionId"] = args ? args.versionId : undefined;
            resourceInputs["vpcAccessConnector"] = args ? args.vpcAccessConnector : undefined;
            resourceInputs["name"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(StandardAppVersion.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering StandardAppVersion resources.
 */
export interface StandardAppVersionState {
    /**
     * Automatic scaling is based on request rate, response latencies, and other application metrics.
     * Structure is documented below.
     */
    automaticScaling?: pulumi.Input<inputs.appengine.StandardAppVersionAutomaticScaling>;
    /**
     * Basic scaling creates instances when your application receives requests. Each instance will be shut down when the application becomes idle. Basic scaling is ideal for work that is intermittent or driven by user activity.
     * Structure is documented below.
     */
    basicScaling?: pulumi.Input<inputs.appengine.StandardAppVersionBasicScaling>;
    /**
     * If set to `true`, the service will be deleted if it is the last version.
     */
    deleteServiceOnDestroy?: pulumi.Input<boolean>;
    /**
     * Code and application artifacts that make up this version.
     * Structure is documented below.
     */
    deployment?: pulumi.Input<inputs.appengine.StandardAppVersionDeployment>;
    /**
     * The entrypoint for the application.
     * Structure is documented below.
     */
    entrypoint?: pulumi.Input<inputs.appengine.StandardAppVersionEntrypoint>;
    /**
     * Environment variables available to the application.
     */
    envVariables?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * An ordered list of URL-matching patterns that should be applied to incoming requests.
     * The first matching URL handles the request and other request handlers are not attempted.
     * Structure is documented below.
     */
    handlers?: pulumi.Input<pulumi.Input<inputs.appengine.StandardAppVersionHandler>[]>;
    /**
     * A list of the types of messages that this application is able to receive.
     * Each value may be one of `INBOUND_SERVICE_MAIL`, `INBOUND_SERVICE_MAIL_BOUNCE`, `INBOUND_SERVICE_XMPP_ERROR`, `INBOUND_SERVICE_XMPP_MESSAGE`, `INBOUND_SERVICE_XMPP_SUBSCRIBE`, `INBOUND_SERVICE_XMPP_PRESENCE`, `INBOUND_SERVICE_CHANNEL_PRESENCE`, and `INBOUND_SERVICE_WARMUP`.
     */
    inboundServices?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Instance class that is used to run this version. Valid values are
     * AutomaticScaling: F1, F2, F4, F4_1G
     * BasicScaling or ManualScaling: B1, B2, B4, B4_1G, B8
     * Defaults to F1 for AutomaticScaling and B2 for ManualScaling and BasicScaling. If no scaling is specified, AutomaticScaling is chosen.
     */
    instanceClass?: pulumi.Input<string>;
    /**
     * Configuration for third-party Python runtime libraries that are required by the application.
     * Structure is documented below.
     */
    libraries?: pulumi.Input<pulumi.Input<inputs.appengine.StandardAppVersionLibrary>[]>;
    /**
     * A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of its memory over time.
     * Structure is documented below.
     */
    manualScaling?: pulumi.Input<inputs.appengine.StandardAppVersionManualScaling>;
    /**
     * Full Serverless VPC Access Connector name e.g. /projects/my-project/locations/us-central1/connectors/c1.
     */
    name?: pulumi.Input<string>;
    /**
     * If set to `true`, the application version will not be deleted.
     */
    noopOnDestroy?: pulumi.Input<boolean>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Desired runtime. Example python27.
     */
    runtime?: pulumi.Input<string>;
    /**
     * The version of the API in the given runtime environment.
     * Please see the app.yaml reference for valid values at `https://cloud.google.com/appengine/docs/standard/<language>/config/appref`\
     * Substitute `<language>` with `python`, `java`, `php`, `ruby`, `go` or `nodejs`.
     */
    runtimeApiVersion?: pulumi.Input<string>;
    /**
     * AppEngine service resource
     */
    service?: pulumi.Input<string>;
    /**
     * Whether multiple requests can be dispatched to this version at once.
     */
    threadsafe?: pulumi.Input<boolean>;
    /**
     * Relative name of the version within the service. For example, `v1`. Version names can contain only lowercase letters, numbers, or hyphens. Reserved names,"default", "latest", and any name with the prefix "ah-".
     */
    versionId?: pulumi.Input<string>;
    /**
     * Enables VPC connectivity for standard apps.
     * Structure is documented below.
     */
    vpcAccessConnector?: pulumi.Input<inputs.appengine.StandardAppVersionVpcAccessConnector>;
}

/**
 * The set of arguments for constructing a StandardAppVersion resource.
 */
export interface StandardAppVersionArgs {
    /**
     * Automatic scaling is based on request rate, response latencies, and other application metrics.
     * Structure is documented below.
     */
    automaticScaling?: pulumi.Input<inputs.appengine.StandardAppVersionAutomaticScaling>;
    /**
     * Basic scaling creates instances when your application receives requests. Each instance will be shut down when the application becomes idle. Basic scaling is ideal for work that is intermittent or driven by user activity.
     * Structure is documented below.
     */
    basicScaling?: pulumi.Input<inputs.appengine.StandardAppVersionBasicScaling>;
    /**
     * If set to `true`, the service will be deleted if it is the last version.
     */
    deleteServiceOnDestroy?: pulumi.Input<boolean>;
    /**
     * Code and application artifacts that make up this version.
     * Structure is documented below.
     */
    deployment: pulumi.Input<inputs.appengine.StandardAppVersionDeployment>;
    /**
     * The entrypoint for the application.
     * Structure is documented below.
     */
    entrypoint: pulumi.Input<inputs.appengine.StandardAppVersionEntrypoint>;
    /**
     * Environment variables available to the application.
     */
    envVariables?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * An ordered list of URL-matching patterns that should be applied to incoming requests.
     * The first matching URL handles the request and other request handlers are not attempted.
     * Structure is documented below.
     */
    handlers?: pulumi.Input<pulumi.Input<inputs.appengine.StandardAppVersionHandler>[]>;
    /**
     * A list of the types of messages that this application is able to receive.
     * Each value may be one of `INBOUND_SERVICE_MAIL`, `INBOUND_SERVICE_MAIL_BOUNCE`, `INBOUND_SERVICE_XMPP_ERROR`, `INBOUND_SERVICE_XMPP_MESSAGE`, `INBOUND_SERVICE_XMPP_SUBSCRIBE`, `INBOUND_SERVICE_XMPP_PRESENCE`, `INBOUND_SERVICE_CHANNEL_PRESENCE`, and `INBOUND_SERVICE_WARMUP`.
     */
    inboundServices?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Instance class that is used to run this version. Valid values are
     * AutomaticScaling: F1, F2, F4, F4_1G
     * BasicScaling or ManualScaling: B1, B2, B4, B4_1G, B8
     * Defaults to F1 for AutomaticScaling and B2 for ManualScaling and BasicScaling. If no scaling is specified, AutomaticScaling is chosen.
     */
    instanceClass?: pulumi.Input<string>;
    /**
     * Configuration for third-party Python runtime libraries that are required by the application.
     * Structure is documented below.
     */
    libraries?: pulumi.Input<pulumi.Input<inputs.appengine.StandardAppVersionLibrary>[]>;
    /**
     * A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of its memory over time.
     * Structure is documented below.
     */
    manualScaling?: pulumi.Input<inputs.appengine.StandardAppVersionManualScaling>;
    /**
     * If set to `true`, the application version will not be deleted.
     */
    noopOnDestroy?: pulumi.Input<boolean>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Desired runtime. Example python27.
     */
    runtime: pulumi.Input<string>;
    /**
     * The version of the API in the given runtime environment.
     * Please see the app.yaml reference for valid values at `https://cloud.google.com/appengine/docs/standard/<language>/config/appref`\
     * Substitute `<language>` with `python`, `java`, `php`, `ruby`, `go` or `nodejs`.
     */
    runtimeApiVersion?: pulumi.Input<string>;
    /**
     * AppEngine service resource
     */
    service: pulumi.Input<string>;
    /**
     * Whether multiple requests can be dispatched to this version at once.
     */
    threadsafe?: pulumi.Input<boolean>;
    /**
     * Relative name of the version within the service. For example, `v1`. Version names can contain only lowercase letters, numbers, or hyphens. Reserved names,"default", "latest", and any name with the prefix "ah-".
     */
    versionId?: pulumi.Input<string>;
    /**
     * Enables VPC connectivity for standard apps.
     * Structure is documented below.
     */
    vpcAccessConnector?: pulumi.Input<inputs.appengine.StandardAppVersionVpcAccessConnector>;
}
