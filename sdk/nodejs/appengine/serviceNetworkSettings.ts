// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * A NetworkSettings resource is a container for ingress settings for a version or service.
 *
 * To get more information about ServiceNetworkSettings, see:
 *
 * * [API documentation](https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services)
 *
 * ## Example Usage
 * ### App Engine Service Network Settings
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
 * const liveappV1 = new gcp.appengine.StandardAppVersion("liveappV1", {
 *     versionId: "v1",
 *     service: "liveapp",
 *     deleteServiceOnDestroy: true,
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
 * });
 * const liveapp = new gcp.appengine.ServiceNetworkSettings("liveapp", {
 *     service: liveappV1.service,
 *     networkSettings: {
 *         ingressTrafficAllowed: "INGRESS_TRAFFIC_ALLOWED_INTERNAL_ONLY",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * ServiceNetworkSettings can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:appengine/serviceNetworkSettings:ServiceNetworkSettings default apps/{{project}}/services/{{service}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:appengine/serviceNetworkSettings:ServiceNetworkSettings default {{project}}/{{service}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:appengine/serviceNetworkSettings:ServiceNetworkSettings default {{service}}
 * ```
 */
export class ServiceNetworkSettings extends pulumi.CustomResource {
    /**
     * Get an existing ServiceNetworkSettings resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceNetworkSettingsState, opts?: pulumi.CustomResourceOptions): ServiceNetworkSettings {
        return new ServiceNetworkSettings(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:appengine/serviceNetworkSettings:ServiceNetworkSettings';

    /**
     * Returns true if the given object is an instance of ServiceNetworkSettings.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceNetworkSettings {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceNetworkSettings.__pulumiType;
    }

    /**
     * Ingress settings for this service. Will apply to all versions.
     * Structure is documented below.
     */
    public readonly networkSettings!: pulumi.Output<outputs.appengine.ServiceNetworkSettingsNetworkSettings>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The name of the service these settings apply to.
     */
    public readonly service!: pulumi.Output<string>;

    /**
     * Create a ServiceNetworkSettings resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceNetworkSettingsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceNetworkSettingsArgs | ServiceNetworkSettingsState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceNetworkSettingsState | undefined;
            inputs["networkSettings"] = state ? state.networkSettings : undefined;
            inputs["project"] = state ? state.project : undefined;
            inputs["service"] = state ? state.service : undefined;
        } else {
            const args = argsOrState as ServiceNetworkSettingsArgs | undefined;
            if ((!args || args.networkSettings === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkSettings'");
            }
            if ((!args || args.service === undefined) && !opts.urn) {
                throw new Error("Missing required property 'service'");
            }
            inputs["networkSettings"] = args ? args.networkSettings : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["service"] = args ? args.service : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ServiceNetworkSettings.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceNetworkSettings resources.
 */
export interface ServiceNetworkSettingsState {
    /**
     * Ingress settings for this service. Will apply to all versions.
     * Structure is documented below.
     */
    networkSettings?: pulumi.Input<inputs.appengine.ServiceNetworkSettingsNetworkSettings>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The name of the service these settings apply to.
     */
    service?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceNetworkSettings resource.
 */
export interface ServiceNetworkSettingsArgs {
    /**
     * Ingress settings for this service. Will apply to all versions.
     * Structure is documented below.
     */
    networkSettings: pulumi.Input<inputs.appengine.ServiceNetworkSettingsNetworkSettings>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The name of the service these settings apply to.
     */
    service: pulumi.Input<string>;
}
