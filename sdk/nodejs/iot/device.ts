// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * > **Warning:** `gcp.iot.Device` is deprecated in the API. This resource will be removed in the next major release of the provider.
 *
 * A Google Cloud IoT Core device.
 *
 * To get more information about Device, see:
 *
 * * [API documentation](https://cloud.google.com/iot/docs/reference/cloudiot/rest/)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/iot/docs/)
 *
 * ## Example Usage
 * ### Cloudiot Device Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const registry = new gcp.iot.Registry("registry", {});
 * const test_device = new gcp.iot.Device("test-device", {registry: registry.id});
 * ```
 * ### Cloudiot Device Full
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fs from "fs";
 * import * as gcp from "@pulumi/gcp";
 *
 * const registry = new gcp.iot.Registry("registry", {});
 * const test_device = new gcp.iot.Device("test-device", {
 *     registry: registry.id,
 *     credentials: [{
 *         publicKey: {
 *             format: "RSA_PEM",
 *             key: fs.readFileSync("test-fixtures/rsa_public.pem"),
 *         },
 *     }],
 *     blocked: false,
 *     logLevel: "INFO",
 *     metadata: {
 *         test_key_1: "test_value_1",
 *     },
 *     gatewayConfig: {
 *         gatewayType: "NON_GATEWAY",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Device can be imported using any of these accepted formats:
 *
 * ```sh
 *  $ pulumi import gcp:iot/device:Device default {{registry}}/devices/{{name}}
 * ```
 */
export class Device extends pulumi.CustomResource {
    /**
     * Get an existing Device resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DeviceState, opts?: pulumi.CustomResourceOptions): Device {
        return new Device(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:iot/device:Device';

    /**
     * Returns true if the given object is an instance of Device.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Device {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Device.__pulumiType;
    }

    /**
     * If a device is blocked, connections or requests from this device will fail.
     */
    public readonly blocked!: pulumi.Output<boolean | undefined>;
    /**
     * The most recent device configuration, which is eventually sent from Cloud IoT Core to the device.
     * Structure is documented below.
     */
    public /*out*/ readonly configs!: pulumi.Output<outputs.iot.DeviceConfig[]>;
    /**
     * The credentials used to authenticate this device.
     * Structure is documented below.
     */
    public readonly credentials!: pulumi.Output<outputs.iot.DeviceCredential[] | undefined>;
    /**
     * Gateway-related configuration and state.
     * Structure is documented below.
     */
    public readonly gatewayConfig!: pulumi.Output<outputs.iot.DeviceGatewayConfig | undefined>;
    /**
     * The last time a cloud-to-device config version acknowledgment was received from the device.
     */
    public /*out*/ readonly lastConfigAckTime!: pulumi.Output<string>;
    /**
     * The last time a cloud-to-device config version was sent to the device.
     */
    public /*out*/ readonly lastConfigSendTime!: pulumi.Output<string>;
    /**
     * The error message of the most recent error, such as a failure to publish to Cloud Pub/Sub.
     * Structure is documented below.
     */
    public /*out*/ readonly lastErrorStatuses!: pulumi.Output<outputs.iot.DeviceLastErrorStatus[]>;
    /**
     * The time the most recent error occurred, such as a failure to publish to Cloud Pub/Sub.
     */
    public /*out*/ readonly lastErrorTime!: pulumi.Output<string>;
    /**
     * The last time a telemetry event was received.
     */
    public /*out*/ readonly lastEventTime!: pulumi.Output<string>;
    /**
     * The last time an MQTT PINGREQ was received.
     */
    public /*out*/ readonly lastHeartbeatTime!: pulumi.Output<string>;
    /**
     * The last time a state event was received.
     */
    public /*out*/ readonly lastStateTime!: pulumi.Output<string>;
    /**
     * The logging verbosity for device activity.
     * Possible values are: `NONE`, `ERROR`, `INFO`, `DEBUG`.
     */
    public readonly logLevel!: pulumi.Output<string | undefined>;
    /**
     * The metadata key-value pairs assigned to the device.
     */
    public readonly metadata!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A unique name for the resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * A server-defined unique numeric ID for the device.
     * This is a more compact way to identify devices, and it is globally unique.
     */
    public /*out*/ readonly numId!: pulumi.Output<string>;
    /**
     * The name of the device registry where this device should be created.
     *
     *
     * - - -
     */
    public readonly registry!: pulumi.Output<string>;
    /**
     * The state most recently received from the device.
     * Structure is documented below.
     */
    public /*out*/ readonly states!: pulumi.Output<outputs.iot.DeviceState[]>;

    /**
     * Create a Device resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DeviceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DeviceArgs | DeviceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DeviceState | undefined;
            resourceInputs["blocked"] = state ? state.blocked : undefined;
            resourceInputs["configs"] = state ? state.configs : undefined;
            resourceInputs["credentials"] = state ? state.credentials : undefined;
            resourceInputs["gatewayConfig"] = state ? state.gatewayConfig : undefined;
            resourceInputs["lastConfigAckTime"] = state ? state.lastConfigAckTime : undefined;
            resourceInputs["lastConfigSendTime"] = state ? state.lastConfigSendTime : undefined;
            resourceInputs["lastErrorStatuses"] = state ? state.lastErrorStatuses : undefined;
            resourceInputs["lastErrorTime"] = state ? state.lastErrorTime : undefined;
            resourceInputs["lastEventTime"] = state ? state.lastEventTime : undefined;
            resourceInputs["lastHeartbeatTime"] = state ? state.lastHeartbeatTime : undefined;
            resourceInputs["lastStateTime"] = state ? state.lastStateTime : undefined;
            resourceInputs["logLevel"] = state ? state.logLevel : undefined;
            resourceInputs["metadata"] = state ? state.metadata : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["numId"] = state ? state.numId : undefined;
            resourceInputs["registry"] = state ? state.registry : undefined;
            resourceInputs["states"] = state ? state.states : undefined;
        } else {
            const args = argsOrState as DeviceArgs | undefined;
            if ((!args || args.registry === undefined) && !opts.urn) {
                throw new Error("Missing required property 'registry'");
            }
            resourceInputs["blocked"] = args ? args.blocked : undefined;
            resourceInputs["credentials"] = args ? args.credentials : undefined;
            resourceInputs["gatewayConfig"] = args ? args.gatewayConfig : undefined;
            resourceInputs["logLevel"] = args ? args.logLevel : undefined;
            resourceInputs["metadata"] = args ? args.metadata : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["registry"] = args ? args.registry : undefined;
            resourceInputs["configs"] = undefined /*out*/;
            resourceInputs["lastConfigAckTime"] = undefined /*out*/;
            resourceInputs["lastConfigSendTime"] = undefined /*out*/;
            resourceInputs["lastErrorStatuses"] = undefined /*out*/;
            resourceInputs["lastErrorTime"] = undefined /*out*/;
            resourceInputs["lastEventTime"] = undefined /*out*/;
            resourceInputs["lastHeartbeatTime"] = undefined /*out*/;
            resourceInputs["lastStateTime"] = undefined /*out*/;
            resourceInputs["numId"] = undefined /*out*/;
            resourceInputs["states"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Device.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Device resources.
 */
export interface DeviceState {
    /**
     * If a device is blocked, connections or requests from this device will fail.
     */
    blocked?: pulumi.Input<boolean>;
    /**
     * The most recent device configuration, which is eventually sent from Cloud IoT Core to the device.
     * Structure is documented below.
     */
    configs?: pulumi.Input<pulumi.Input<inputs.iot.DeviceConfig>[]>;
    /**
     * The credentials used to authenticate this device.
     * Structure is documented below.
     */
    credentials?: pulumi.Input<pulumi.Input<inputs.iot.DeviceCredential>[]>;
    /**
     * Gateway-related configuration and state.
     * Structure is documented below.
     */
    gatewayConfig?: pulumi.Input<inputs.iot.DeviceGatewayConfig>;
    /**
     * The last time a cloud-to-device config version acknowledgment was received from the device.
     */
    lastConfigAckTime?: pulumi.Input<string>;
    /**
     * The last time a cloud-to-device config version was sent to the device.
     */
    lastConfigSendTime?: pulumi.Input<string>;
    /**
     * The error message of the most recent error, such as a failure to publish to Cloud Pub/Sub.
     * Structure is documented below.
     */
    lastErrorStatuses?: pulumi.Input<pulumi.Input<inputs.iot.DeviceLastErrorStatus>[]>;
    /**
     * The time the most recent error occurred, such as a failure to publish to Cloud Pub/Sub.
     */
    lastErrorTime?: pulumi.Input<string>;
    /**
     * The last time a telemetry event was received.
     */
    lastEventTime?: pulumi.Input<string>;
    /**
     * The last time an MQTT PINGREQ was received.
     */
    lastHeartbeatTime?: pulumi.Input<string>;
    /**
     * The last time a state event was received.
     */
    lastStateTime?: pulumi.Input<string>;
    /**
     * The logging verbosity for device activity.
     * Possible values are: `NONE`, `ERROR`, `INFO`, `DEBUG`.
     */
    logLevel?: pulumi.Input<string>;
    /**
     * The metadata key-value pairs assigned to the device.
     */
    metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A unique name for the resource.
     */
    name?: pulumi.Input<string>;
    /**
     * A server-defined unique numeric ID for the device.
     * This is a more compact way to identify devices, and it is globally unique.
     */
    numId?: pulumi.Input<string>;
    /**
     * The name of the device registry where this device should be created.
     *
     *
     * - - -
     */
    registry?: pulumi.Input<string>;
    /**
     * The state most recently received from the device.
     * Structure is documented below.
     */
    states?: pulumi.Input<pulumi.Input<inputs.iot.DeviceState>[]>;
}

/**
 * The set of arguments for constructing a Device resource.
 */
export interface DeviceArgs {
    /**
     * If a device is blocked, connections or requests from this device will fail.
     */
    blocked?: pulumi.Input<boolean>;
    /**
     * The credentials used to authenticate this device.
     * Structure is documented below.
     */
    credentials?: pulumi.Input<pulumi.Input<inputs.iot.DeviceCredential>[]>;
    /**
     * Gateway-related configuration and state.
     * Structure is documented below.
     */
    gatewayConfig?: pulumi.Input<inputs.iot.DeviceGatewayConfig>;
    /**
     * The logging verbosity for device activity.
     * Possible values are: `NONE`, `ERROR`, `INFO`, `DEBUG`.
     */
    logLevel?: pulumi.Input<string>;
    /**
     * The metadata key-value pairs assigned to the device.
     */
    metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A unique name for the resource.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the device registry where this device should be created.
     *
     *
     * - - -
     */
    registry: pulumi.Input<string>;
}
