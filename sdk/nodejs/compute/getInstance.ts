// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Get information about a VM instance resource within GCE. For more information see
 * [the official documentation](https://cloud.google.com/compute/docs/instances)
 * and
 * [API](https://cloud.google.com/compute/docs/reference/latest/instances).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const appserver = gcp.compute.getInstance({
 *     name: "primary-application-server",
 *     zone: "us-central1-a",
 * });
 * ```
 */
export function getInstance(args?: GetInstanceArgs, opts?: pulumi.InvokeOptions): Promise<GetInstanceResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("gcp:compute/getInstance:getInstance", {
        "name": args.name,
        "project": args.project,
        "selfLink": args.selfLink,
        "zone": args.zone,
    }, opts);
}

/**
 * A collection of arguments for invoking getInstance.
 */
export interface GetInstanceArgs {
    /**
     * The name of the instance. One of `name` or `selfLink` must be provided.
     */
    name?: string;
    /**
     * The ID of the project in which the resource belongs.
     * If `selfLink` is provided, this value is ignored.  If neither `selfLink`
     * nor `project` are provided, the provider project is used.
     */
    project?: string;
    /**
     * The self link of the instance. One of `name` or `selfLink` must be provided.
     */
    selfLink?: string;
    /**
     * The zone of the instance. If `selfLink` is provided, this
     * value is ignored.  If neither `selfLink` nor `zone` are provided, the
     * provider zone is used.
     */
    zone?: string;
}

/**
 * A collection of values returned by getInstance.
 */
export interface GetInstanceResult {
    readonly advancedMachineFeatures: outputs.compute.GetInstanceAdvancedMachineFeature[];
    readonly allowStoppingForUpdate: boolean;
    /**
     * List of disks attached to the instance. Structure is documented below.
     */
    readonly attachedDisks: outputs.compute.GetInstanceAttachedDisk[];
    /**
     * The boot disk for the instance. Structure is documented below.
     */
    readonly bootDisks: outputs.compute.GetInstanceBootDisk[];
    /**
     * Whether sending and receiving of packets with non-matching source or destination IPs is allowed.
     */
    readonly canIpForward: boolean;
    readonly confidentialInstanceConfigs: outputs.compute.GetInstanceConfidentialInstanceConfig[];
    /**
     * The CPU platform used by this instance.
     */
    readonly cpuPlatform: string;
    /**
     * The current status of the instance. This could be one of the following values: PROVISIONING, STAGING, RUNNING, STOPPING, SUSPENDING, SUSPENDED, REPAIRING, and TERMINATED. For more information about the status of the instance, see [Instance life cycle](https://cloud.google.com/compute/docs/instances/instance-life-cycle).`,
     */
    readonly currentStatus: string;
    /**
     * Whether deletion protection is enabled on this instance.
     */
    readonly deletionProtection: boolean;
    /**
     * A brief description of the resource.
     */
    readonly description: string;
    readonly desiredStatus: string;
    /**
     * Whether the instance has virtual displays enabled.
     */
    readonly enableDisplay: boolean;
    /**
     * List of the type and count of accelerator cards attached to the instance. Structure is documented below.
     */
    readonly guestAccelerators: outputs.compute.GetInstanceGuestAccelerator[];
    readonly hostname: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The server-assigned unique identifier of this instance.
     */
    readonly instanceId: string;
    /**
     * The unique fingerprint of the labels.
     */
    readonly labelFingerprint: string;
    /**
     * A set of key/value label pairs assigned to the disk.
     */
    readonly labels: {[key: string]: string};
    /**
     * The machine type to create.
     */
    readonly machineType: string;
    /**
     * Metadata key/value pairs made available within the instance.
     */
    readonly metadata: {[key: string]: string};
    /**
     * The unique fingerprint of the metadata.
     */
    readonly metadataFingerprint: string;
    readonly metadataStartupScript: string;
    /**
     * The minimum CPU platform specified for the VM instance.
     */
    readonly minCpuPlatform: string;
    readonly name?: string;
    /**
     * The networks attached to the instance. Structure is documented below.
     */
    readonly networkInterfaces: outputs.compute.GetInstanceNetworkInterface[];
    /**
     * The network performance configuration setting for the instance, if set. Structure is documented below.
     */
    readonly networkPerformanceConfigs: outputs.compute.GetInstanceNetworkPerformanceConfig[];
    readonly params: outputs.compute.GetInstanceParam[];
    readonly project?: string;
    readonly reservationAffinities: outputs.compute.GetInstanceReservationAffinity[];
    readonly resourcePolicies: string[];
    /**
     * The scheduling strategy being used by the instance. Structure is documented below
     */
    readonly schedulings: outputs.compute.GetInstanceScheduling[];
    /**
     * The scratch disks attached to the instance. Structure is documented below.
     */
    readonly scratchDisks: outputs.compute.GetInstanceScratchDisk[];
    /**
     * The URI of the created resource.
     */
    readonly selfLink?: string;
    /**
     * The service account to attach to the instance. Structure is documented below.
     */
    readonly serviceAccounts: outputs.compute.GetInstanceServiceAccount[];
    /**
     * The shielded vm config being used by the instance. Structure is documented below.
     */
    readonly shieldedInstanceConfigs: outputs.compute.GetInstanceShieldedInstanceConfig[];
    /**
     * The list of tags attached to the instance.
     */
    readonly tags: string[];
    /**
     * The unique fingerprint of the tags.
     */
    readonly tagsFingerprint: string;
    readonly zone?: string;
}
/**
 * Get information about a VM instance resource within GCE. For more information see
 * [the official documentation](https://cloud.google.com/compute/docs/instances)
 * and
 * [API](https://cloud.google.com/compute/docs/reference/latest/instances).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const appserver = gcp.compute.getInstance({
 *     name: "primary-application-server",
 *     zone: "us-central1-a",
 * });
 * ```
 */
export function getInstanceOutput(args?: GetInstanceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInstanceResult> {
    return pulumi.output(args).apply((a: any) => getInstance(a, opts))
}

/**
 * A collection of arguments for invoking getInstance.
 */
export interface GetInstanceOutputArgs {
    /**
     * The name of the instance. One of `name` or `selfLink` must be provided.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If `selfLink` is provided, this value is ignored.  If neither `selfLink`
     * nor `project` are provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The self link of the instance. One of `name` or `selfLink` must be provided.
     */
    selfLink?: pulumi.Input<string>;
    /**
     * The zone of the instance. If `selfLink` is provided, this
     * value is ignored.  If neither `selfLink` nor `zone` are provided, the
     * provider zone is used.
     */
    zone?: pulumi.Input<string>;
}
