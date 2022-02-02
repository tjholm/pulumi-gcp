// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Manages a VM instance resource within GCE. For more information see
 * [the official documentation](https://cloud.google.com/compute/docs/instances)
 * and
 * [API](https://cloud.google.com/compute/docs/reference/latest/instances).
 *
 * This resource is specifically to create a compute instance from a given
 * `sourceInstanceTemplate`. To create an instance without a template, use the
 * `gcp.compute.Instance` resource.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const tplInstanceTemplate = new gcp.compute.InstanceTemplate("tplInstanceTemplate", {
 *     machineType: "e2-medium",
 *     disks: [{
 *         sourceImage: "debian-cloud/debian-9",
 *         autoDelete: true,
 *         diskSizeGb: 100,
 *         boot: true,
 *     }],
 *     networkInterfaces: [{
 *         network: "default",
 *     }],
 *     metadata: {
 *         foo: "bar",
 *     },
 *     canIpForward: true,
 * });
 * const tplInstanceFromTemplate = new gcp.compute.InstanceFromTemplate("tplInstanceFromTemplate", {
 *     zone: "us-central1-a",
 *     sourceInstanceTemplate: tplInstanceTemplate.id,
 *     canIpForward: false,
 *     labels: {
 *         my_key: "my_value",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * This resource does not support import.
 */
export class InstanceFromTemplate extends pulumi.CustomResource {
    /**
     * Get an existing InstanceFromTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceFromTemplateState, opts?: pulumi.CustomResourceOptions): InstanceFromTemplate {
        return new InstanceFromTemplate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/instanceFromTemplate:InstanceFromTemplate';

    /**
     * Returns true if the given object is an instance of InstanceFromTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstanceFromTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstanceFromTemplate.__pulumiType;
    }

    /**
     * Controls for advanced machine-related behavior features.
     */
    public readonly advancedMachineFeatures!: pulumi.Output<outputs.compute.InstanceFromTemplateAdvancedMachineFeatures>;
    /**
     * If true, allows Terraform to stop the instance to update its properties. If you try to update a property that requires
     * stopping the instance without setting this field, the update will fail.
     */
    public readonly allowStoppingForUpdate!: pulumi.Output<boolean>;
    /**
     * List of disks attached to the instance
     */
    public readonly attachedDisks!: pulumi.Output<outputs.compute.InstanceFromTemplateAttachedDisk[]>;
    /**
     * The boot disk for the instance.
     */
    public readonly bootDisk!: pulumi.Output<outputs.compute.InstanceFromTemplateBootDisk>;
    /**
     * Whether sending and receiving of packets with non-matching source or destination IPs is allowed.
     */
    public readonly canIpForward!: pulumi.Output<boolean>;
    /**
     * The Confidential VM config being used by the instance. on_host_maintenance has to be set to TERMINATE or this will fail
     * to create.
     */
    public readonly confidentialInstanceConfig!: pulumi.Output<outputs.compute.InstanceFromTemplateConfidentialInstanceConfig>;
    /**
     * The CPU platform used by this instance.
     */
    public /*out*/ readonly cpuPlatform!: pulumi.Output<string>;
    /**
     * Current status of the instance.
     */
    public /*out*/ readonly currentStatus!: pulumi.Output<string>;
    /**
     * Whether deletion protection is enabled on this instance.
     */
    public readonly deletionProtection!: pulumi.Output<boolean>;
    /**
     * A brief description of the resource.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Desired status of the instance. Either "RUNNING" or "TERMINATED".
     */
    public readonly desiredStatus!: pulumi.Output<string>;
    /**
     * Whether the instance has virtual displays enabled.
     */
    public readonly enableDisplay!: pulumi.Output<boolean>;
    /**
     * List of the type and count of accelerator cards attached to the instance.
     */
    public readonly guestAccelerators!: pulumi.Output<outputs.compute.InstanceFromTemplateGuestAccelerator[]>;
    /**
     * A custom hostname for the instance. Must be a fully qualified DNS name and RFC-1035-valid. Valid format is a series of
     * labels 1-63 characters long matching the regular expression [a-z]([-a-z0-9]*[a-z0-9]), concatenated with periods. The
     * entire hostname must not exceed 253 characters. Changing this forces a new resource to be created.
     */
    public readonly hostname!: pulumi.Output<string>;
    /**
     * The server-assigned unique identifier of this instance.
     */
    public /*out*/ readonly instanceId!: pulumi.Output<string>;
    /**
     * The unique fingerprint of the labels.
     */
    public /*out*/ readonly labelFingerprint!: pulumi.Output<string>;
    /**
     * A set of key/value label pairs assigned to the instance.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The machine type to create.
     */
    public readonly machineType!: pulumi.Output<string>;
    /**
     * Metadata key/value pairs made available within the instance.
     */
    public readonly metadata!: pulumi.Output<{[key: string]: string}>;
    /**
     * The unique fingerprint of the metadata.
     */
    public /*out*/ readonly metadataFingerprint!: pulumi.Output<string>;
    /**
     * Metadata startup scripts made available within the instance.
     */
    public readonly metadataStartupScript!: pulumi.Output<string>;
    /**
     * The minimum CPU platform specified for the VM instance.
     */
    public readonly minCpuPlatform!: pulumi.Output<string>;
    /**
     * A unique name for the resource, required by GCE.
     * Changing this forces a new resource to be created.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The networks attached to the instance.
     */
    public readonly networkInterfaces!: pulumi.Output<outputs.compute.InstanceFromTemplateNetworkInterface[]>;
    /**
     * Configures network performance settings for the instance. If not specified, the instance will be created with its
     * default network performance configuration.
     */
    public readonly networkPerformanceConfig!: pulumi.Output<outputs.compute.InstanceFromTemplateNetworkPerformanceConfig>;
    /**
     * The ID of the project in which the resource belongs. If self_link is provided, this value is ignored. If neither
     * self_link nor project are provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Specifies the reservations that this instance can consume from.
     */
    public readonly reservationAffinity!: pulumi.Output<outputs.compute.InstanceFromTemplateReservationAffinity>;
    /**
     * A list of short names or self_links of resource policies to attach to the instance. Currently a max of 1 resource policy
     * is supported.
     */
    public readonly resourcePolicies!: pulumi.Output<string>;
    /**
     * The scheduling strategy being used by the instance.
     */
    public readonly scheduling!: pulumi.Output<outputs.compute.InstanceFromTemplateScheduling>;
    /**
     * The scratch disks attached to the instance.
     */
    public readonly scratchDisks!: pulumi.Output<outputs.compute.InstanceFromTemplateScratchDisk[]>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * The service account to attach to the instance.
     */
    public readonly serviceAccount!: pulumi.Output<outputs.compute.InstanceFromTemplateServiceAccount>;
    /**
     * The shielded vm config being used by the instance.
     */
    public readonly shieldedInstanceConfig!: pulumi.Output<outputs.compute.InstanceFromTemplateShieldedInstanceConfig>;
    /**
     * Name or self link of an instance
     * template to create the instance based on.
     */
    public readonly sourceInstanceTemplate!: pulumi.Output<string>;
    /**
     * The list of tags attached to the instance.
     */
    public readonly tags!: pulumi.Output<string[]>;
    /**
     * The unique fingerprint of the tags.
     */
    public /*out*/ readonly tagsFingerprint!: pulumi.Output<string>;
    /**
     * The zone that the machine should be created in. If not
     * set, the provider zone is used.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a InstanceFromTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceFromTemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceFromTemplateArgs | InstanceFromTemplateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceFromTemplateState | undefined;
            resourceInputs["advancedMachineFeatures"] = state ? state.advancedMachineFeatures : undefined;
            resourceInputs["allowStoppingForUpdate"] = state ? state.allowStoppingForUpdate : undefined;
            resourceInputs["attachedDisks"] = state ? state.attachedDisks : undefined;
            resourceInputs["bootDisk"] = state ? state.bootDisk : undefined;
            resourceInputs["canIpForward"] = state ? state.canIpForward : undefined;
            resourceInputs["confidentialInstanceConfig"] = state ? state.confidentialInstanceConfig : undefined;
            resourceInputs["cpuPlatform"] = state ? state.cpuPlatform : undefined;
            resourceInputs["currentStatus"] = state ? state.currentStatus : undefined;
            resourceInputs["deletionProtection"] = state ? state.deletionProtection : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["desiredStatus"] = state ? state.desiredStatus : undefined;
            resourceInputs["enableDisplay"] = state ? state.enableDisplay : undefined;
            resourceInputs["guestAccelerators"] = state ? state.guestAccelerators : undefined;
            resourceInputs["hostname"] = state ? state.hostname : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["labelFingerprint"] = state ? state.labelFingerprint : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["machineType"] = state ? state.machineType : undefined;
            resourceInputs["metadata"] = state ? state.metadata : undefined;
            resourceInputs["metadataFingerprint"] = state ? state.metadataFingerprint : undefined;
            resourceInputs["metadataStartupScript"] = state ? state.metadataStartupScript : undefined;
            resourceInputs["minCpuPlatform"] = state ? state.minCpuPlatform : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networkInterfaces"] = state ? state.networkInterfaces : undefined;
            resourceInputs["networkPerformanceConfig"] = state ? state.networkPerformanceConfig : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["reservationAffinity"] = state ? state.reservationAffinity : undefined;
            resourceInputs["resourcePolicies"] = state ? state.resourcePolicies : undefined;
            resourceInputs["scheduling"] = state ? state.scheduling : undefined;
            resourceInputs["scratchDisks"] = state ? state.scratchDisks : undefined;
            resourceInputs["selfLink"] = state ? state.selfLink : undefined;
            resourceInputs["serviceAccount"] = state ? state.serviceAccount : undefined;
            resourceInputs["shieldedInstanceConfig"] = state ? state.shieldedInstanceConfig : undefined;
            resourceInputs["sourceInstanceTemplate"] = state ? state.sourceInstanceTemplate : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsFingerprint"] = state ? state.tagsFingerprint : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as InstanceFromTemplateArgs | undefined;
            if ((!args || args.sourceInstanceTemplate === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceInstanceTemplate'");
            }
            resourceInputs["advancedMachineFeatures"] = args ? args.advancedMachineFeatures : undefined;
            resourceInputs["allowStoppingForUpdate"] = args ? args.allowStoppingForUpdate : undefined;
            resourceInputs["attachedDisks"] = args ? args.attachedDisks : undefined;
            resourceInputs["bootDisk"] = args ? args.bootDisk : undefined;
            resourceInputs["canIpForward"] = args ? args.canIpForward : undefined;
            resourceInputs["confidentialInstanceConfig"] = args ? args.confidentialInstanceConfig : undefined;
            resourceInputs["deletionProtection"] = args ? args.deletionProtection : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["desiredStatus"] = args ? args.desiredStatus : undefined;
            resourceInputs["enableDisplay"] = args ? args.enableDisplay : undefined;
            resourceInputs["guestAccelerators"] = args ? args.guestAccelerators : undefined;
            resourceInputs["hostname"] = args ? args.hostname : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["machineType"] = args ? args.machineType : undefined;
            resourceInputs["metadata"] = args ? args.metadata : undefined;
            resourceInputs["metadataStartupScript"] = args ? args.metadataStartupScript : undefined;
            resourceInputs["minCpuPlatform"] = args ? args.minCpuPlatform : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networkInterfaces"] = args ? args.networkInterfaces : undefined;
            resourceInputs["networkPerformanceConfig"] = args ? args.networkPerformanceConfig : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["reservationAffinity"] = args ? args.reservationAffinity : undefined;
            resourceInputs["resourcePolicies"] = args ? args.resourcePolicies : undefined;
            resourceInputs["scheduling"] = args ? args.scheduling : undefined;
            resourceInputs["scratchDisks"] = args ? args.scratchDisks : undefined;
            resourceInputs["serviceAccount"] = args ? args.serviceAccount : undefined;
            resourceInputs["shieldedInstanceConfig"] = args ? args.shieldedInstanceConfig : undefined;
            resourceInputs["sourceInstanceTemplate"] = args ? args.sourceInstanceTemplate : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["cpuPlatform"] = undefined /*out*/;
            resourceInputs["currentStatus"] = undefined /*out*/;
            resourceInputs["instanceId"] = undefined /*out*/;
            resourceInputs["labelFingerprint"] = undefined /*out*/;
            resourceInputs["metadataFingerprint"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["tagsFingerprint"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(InstanceFromTemplate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InstanceFromTemplate resources.
 */
export interface InstanceFromTemplateState {
    /**
     * Controls for advanced machine-related behavior features.
     */
    advancedMachineFeatures?: pulumi.Input<inputs.compute.InstanceFromTemplateAdvancedMachineFeatures>;
    /**
     * If true, allows Terraform to stop the instance to update its properties. If you try to update a property that requires
     * stopping the instance without setting this field, the update will fail.
     */
    allowStoppingForUpdate?: pulumi.Input<boolean>;
    /**
     * List of disks attached to the instance
     */
    attachedDisks?: pulumi.Input<pulumi.Input<inputs.compute.InstanceFromTemplateAttachedDisk>[]>;
    /**
     * The boot disk for the instance.
     */
    bootDisk?: pulumi.Input<inputs.compute.InstanceFromTemplateBootDisk>;
    /**
     * Whether sending and receiving of packets with non-matching source or destination IPs is allowed.
     */
    canIpForward?: pulumi.Input<boolean>;
    /**
     * The Confidential VM config being used by the instance. on_host_maintenance has to be set to TERMINATE or this will fail
     * to create.
     */
    confidentialInstanceConfig?: pulumi.Input<inputs.compute.InstanceFromTemplateConfidentialInstanceConfig>;
    /**
     * The CPU platform used by this instance.
     */
    cpuPlatform?: pulumi.Input<string>;
    /**
     * Current status of the instance.
     */
    currentStatus?: pulumi.Input<string>;
    /**
     * Whether deletion protection is enabled on this instance.
     */
    deletionProtection?: pulumi.Input<boolean>;
    /**
     * A brief description of the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Desired status of the instance. Either "RUNNING" or "TERMINATED".
     */
    desiredStatus?: pulumi.Input<string>;
    /**
     * Whether the instance has virtual displays enabled.
     */
    enableDisplay?: pulumi.Input<boolean>;
    /**
     * List of the type and count of accelerator cards attached to the instance.
     */
    guestAccelerators?: pulumi.Input<pulumi.Input<inputs.compute.InstanceFromTemplateGuestAccelerator>[]>;
    /**
     * A custom hostname for the instance. Must be a fully qualified DNS name and RFC-1035-valid. Valid format is a series of
     * labels 1-63 characters long matching the regular expression [a-z]([-a-z0-9]*[a-z0-9]), concatenated with periods. The
     * entire hostname must not exceed 253 characters. Changing this forces a new resource to be created.
     */
    hostname?: pulumi.Input<string>;
    /**
     * The server-assigned unique identifier of this instance.
     */
    instanceId?: pulumi.Input<string>;
    /**
     * The unique fingerprint of the labels.
     */
    labelFingerprint?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs assigned to the instance.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The machine type to create.
     */
    machineType?: pulumi.Input<string>;
    /**
     * Metadata key/value pairs made available within the instance.
     */
    metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The unique fingerprint of the metadata.
     */
    metadataFingerprint?: pulumi.Input<string>;
    /**
     * Metadata startup scripts made available within the instance.
     */
    metadataStartupScript?: pulumi.Input<string>;
    /**
     * The minimum CPU platform specified for the VM instance.
     */
    minCpuPlatform?: pulumi.Input<string>;
    /**
     * A unique name for the resource, required by GCE.
     * Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The networks attached to the instance.
     */
    networkInterfaces?: pulumi.Input<pulumi.Input<inputs.compute.InstanceFromTemplateNetworkInterface>[]>;
    /**
     * Configures network performance settings for the instance. If not specified, the instance will be created with its
     * default network performance configuration.
     */
    networkPerformanceConfig?: pulumi.Input<inputs.compute.InstanceFromTemplateNetworkPerformanceConfig>;
    /**
     * The ID of the project in which the resource belongs. If self_link is provided, this value is ignored. If neither
     * self_link nor project are provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Specifies the reservations that this instance can consume from.
     */
    reservationAffinity?: pulumi.Input<inputs.compute.InstanceFromTemplateReservationAffinity>;
    /**
     * A list of short names or self_links of resource policies to attach to the instance. Currently a max of 1 resource policy
     * is supported.
     */
    resourcePolicies?: pulumi.Input<string>;
    /**
     * The scheduling strategy being used by the instance.
     */
    scheduling?: pulumi.Input<inputs.compute.InstanceFromTemplateScheduling>;
    /**
     * The scratch disks attached to the instance.
     */
    scratchDisks?: pulumi.Input<pulumi.Input<inputs.compute.InstanceFromTemplateScratchDisk>[]>;
    /**
     * The URI of the created resource.
     */
    selfLink?: pulumi.Input<string>;
    /**
     * The service account to attach to the instance.
     */
    serviceAccount?: pulumi.Input<inputs.compute.InstanceFromTemplateServiceAccount>;
    /**
     * The shielded vm config being used by the instance.
     */
    shieldedInstanceConfig?: pulumi.Input<inputs.compute.InstanceFromTemplateShieldedInstanceConfig>;
    /**
     * Name or self link of an instance
     * template to create the instance based on.
     */
    sourceInstanceTemplate?: pulumi.Input<string>;
    /**
     * The list of tags attached to the instance.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The unique fingerprint of the tags.
     */
    tagsFingerprint?: pulumi.Input<string>;
    /**
     * The zone that the machine should be created in. If not
     * set, the provider zone is used.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InstanceFromTemplate resource.
 */
export interface InstanceFromTemplateArgs {
    /**
     * Controls for advanced machine-related behavior features.
     */
    advancedMachineFeatures?: pulumi.Input<inputs.compute.InstanceFromTemplateAdvancedMachineFeatures>;
    /**
     * If true, allows Terraform to stop the instance to update its properties. If you try to update a property that requires
     * stopping the instance without setting this field, the update will fail.
     */
    allowStoppingForUpdate?: pulumi.Input<boolean>;
    /**
     * List of disks attached to the instance
     */
    attachedDisks?: pulumi.Input<pulumi.Input<inputs.compute.InstanceFromTemplateAttachedDisk>[]>;
    /**
     * The boot disk for the instance.
     */
    bootDisk?: pulumi.Input<inputs.compute.InstanceFromTemplateBootDisk>;
    /**
     * Whether sending and receiving of packets with non-matching source or destination IPs is allowed.
     */
    canIpForward?: pulumi.Input<boolean>;
    /**
     * The Confidential VM config being used by the instance. on_host_maintenance has to be set to TERMINATE or this will fail
     * to create.
     */
    confidentialInstanceConfig?: pulumi.Input<inputs.compute.InstanceFromTemplateConfidentialInstanceConfig>;
    /**
     * Whether deletion protection is enabled on this instance.
     */
    deletionProtection?: pulumi.Input<boolean>;
    /**
     * A brief description of the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Desired status of the instance. Either "RUNNING" or "TERMINATED".
     */
    desiredStatus?: pulumi.Input<string>;
    /**
     * Whether the instance has virtual displays enabled.
     */
    enableDisplay?: pulumi.Input<boolean>;
    /**
     * List of the type and count of accelerator cards attached to the instance.
     */
    guestAccelerators?: pulumi.Input<pulumi.Input<inputs.compute.InstanceFromTemplateGuestAccelerator>[]>;
    /**
     * A custom hostname for the instance. Must be a fully qualified DNS name and RFC-1035-valid. Valid format is a series of
     * labels 1-63 characters long matching the regular expression [a-z]([-a-z0-9]*[a-z0-9]), concatenated with periods. The
     * entire hostname must not exceed 253 characters. Changing this forces a new resource to be created.
     */
    hostname?: pulumi.Input<string>;
    /**
     * A set of key/value label pairs assigned to the instance.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The machine type to create.
     */
    machineType?: pulumi.Input<string>;
    /**
     * Metadata key/value pairs made available within the instance.
     */
    metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Metadata startup scripts made available within the instance.
     */
    metadataStartupScript?: pulumi.Input<string>;
    /**
     * The minimum CPU platform specified for the VM instance.
     */
    minCpuPlatform?: pulumi.Input<string>;
    /**
     * A unique name for the resource, required by GCE.
     * Changing this forces a new resource to be created.
     */
    name?: pulumi.Input<string>;
    /**
     * The networks attached to the instance.
     */
    networkInterfaces?: pulumi.Input<pulumi.Input<inputs.compute.InstanceFromTemplateNetworkInterface>[]>;
    /**
     * Configures network performance settings for the instance. If not specified, the instance will be created with its
     * default network performance configuration.
     */
    networkPerformanceConfig?: pulumi.Input<inputs.compute.InstanceFromTemplateNetworkPerformanceConfig>;
    /**
     * The ID of the project in which the resource belongs. If self_link is provided, this value is ignored. If neither
     * self_link nor project are provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Specifies the reservations that this instance can consume from.
     */
    reservationAffinity?: pulumi.Input<inputs.compute.InstanceFromTemplateReservationAffinity>;
    /**
     * A list of short names or self_links of resource policies to attach to the instance. Currently a max of 1 resource policy
     * is supported.
     */
    resourcePolicies?: pulumi.Input<string>;
    /**
     * The scheduling strategy being used by the instance.
     */
    scheduling?: pulumi.Input<inputs.compute.InstanceFromTemplateScheduling>;
    /**
     * The scratch disks attached to the instance.
     */
    scratchDisks?: pulumi.Input<pulumi.Input<inputs.compute.InstanceFromTemplateScratchDisk>[]>;
    /**
     * The service account to attach to the instance.
     */
    serviceAccount?: pulumi.Input<inputs.compute.InstanceFromTemplateServiceAccount>;
    /**
     * The shielded vm config being used by the instance.
     */
    shieldedInstanceConfig?: pulumi.Input<inputs.compute.InstanceFromTemplateShieldedInstanceConfig>;
    /**
     * Name or self link of an instance
     * template to create the instance based on.
     */
    sourceInstanceTemplate: pulumi.Input<string>;
    /**
     * The list of tags attached to the instance.
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The zone that the machine should be created in. If not
     * set, the provider zone is used.
     */
    zone?: pulumi.Input<string>;
}
