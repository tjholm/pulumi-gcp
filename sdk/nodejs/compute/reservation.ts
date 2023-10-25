// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Represents a reservation resource. A reservation ensures that capacity is
 * held in a specific zone even if the reserved VMs are not running.
 *
 * Reservations apply only to Compute Engine, Cloud Dataproc, and Google
 * Kubernetes Engine VM usage.Reservations do not apply to `f1-micro` or
 * `g1-small` machine types, preemptible VMs, sole tenant nodes, or other
 * services not listed above
 * like Cloud SQL and Dataflow.
 *
 * To get more information about Reservation, see:
 *
 * * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/reservations)
 * * How-to Guides
 *     * [Reserving zonal resources](https://cloud.google.com/compute/docs/instances/reserving-zonal-resources)
 *
 * ## Example Usage
 * ### Reservation Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const gceReservation = new gcp.compute.Reservation("gceReservation", {
 *     specificReservation: {
 *         count: 1,
 *         instanceProperties: {
 *             machineType: "n2-standard-2",
 *             minCpuPlatform: "Intel Cascade Lake",
 *         },
 *     },
 *     zone: "us-central1-a",
 * });
 * ```
 *
 * ## Import
 *
 * Reservation can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:compute/reservation:Reservation default projects/{{project}}/zones/{{zone}}/reservations/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/reservation:Reservation default {{project}}/{{zone}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/reservation:Reservation default {{zone}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/reservation:Reservation default {{name}}
 * ```
 */
export class Reservation extends pulumi.CustomResource {
    /**
     * Get an existing Reservation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ReservationState, opts?: pulumi.CustomResourceOptions): Reservation {
        return new Reservation(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/reservation:Reservation';

    /**
     * Returns true if the given object is an instance of Reservation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Reservation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Reservation.__pulumiType;
    }

    /**
     * Full or partial URL to a parent commitment. This field displays for
     * reservations that are tied to a commitment.
     */
    public /*out*/ readonly commitment!: pulumi.Output<string>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * The share setting for reservations.
     * Structure is documented below.
     */
    public readonly shareSettings!: pulumi.Output<outputs.compute.ReservationShareSettings>;
    /**
     * Reservation for instances with specific machine shapes.
     * Structure is documented below.
     */
    public readonly specificReservation!: pulumi.Output<outputs.compute.ReservationSpecificReservation>;
    /**
     * When set to true, only VMs that target this reservation by name can
     * consume this reservation. Otherwise, it can be consumed by VMs with
     * affinity for any reservation. Defaults to false.
     */
    public readonly specificReservationRequired!: pulumi.Output<boolean | undefined>;
    /**
     * The status of the reservation.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The zone where the reservation is made.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a Reservation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReservationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ReservationArgs | ReservationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ReservationState | undefined;
            resourceInputs["commitment"] = state ? state.commitment : undefined;
            resourceInputs["creationTimestamp"] = state ? state.creationTimestamp : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["selfLink"] = state ? state.selfLink : undefined;
            resourceInputs["shareSettings"] = state ? state.shareSettings : undefined;
            resourceInputs["specificReservation"] = state ? state.specificReservation : undefined;
            resourceInputs["specificReservationRequired"] = state ? state.specificReservationRequired : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as ReservationArgs | undefined;
            if ((!args || args.specificReservation === undefined) && !opts.urn) {
                throw new Error("Missing required property 'specificReservation'");
            }
            if ((!args || args.zone === undefined) && !opts.urn) {
                throw new Error("Missing required property 'zone'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["shareSettings"] = args ? args.shareSettings : undefined;
            resourceInputs["specificReservation"] = args ? args.specificReservation : undefined;
            resourceInputs["specificReservationRequired"] = args ? args.specificReservationRequired : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["commitment"] = undefined /*out*/;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Reservation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Reservation resources.
 */
export interface ReservationState {
    /**
     * Full or partial URL to a parent commitment. This field displays for
     * reservations that are tied to a commitment.
     */
    commitment?: pulumi.Input<string>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    creationTimestamp?: pulumi.Input<string>;
    /**
     * An optional description of this resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    selfLink?: pulumi.Input<string>;
    /**
     * The share setting for reservations.
     * Structure is documented below.
     */
    shareSettings?: pulumi.Input<inputs.compute.ReservationShareSettings>;
    /**
     * Reservation for instances with specific machine shapes.
     * Structure is documented below.
     */
    specificReservation?: pulumi.Input<inputs.compute.ReservationSpecificReservation>;
    /**
     * When set to true, only VMs that target this reservation by name can
     * consume this reservation. Otherwise, it can be consumed by VMs with
     * affinity for any reservation. Defaults to false.
     */
    specificReservationRequired?: pulumi.Input<boolean>;
    /**
     * The status of the reservation.
     */
    status?: pulumi.Input<string>;
    /**
     * The zone where the reservation is made.
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Reservation resource.
 */
export interface ReservationArgs {
    /**
     * An optional description of this resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the resource. Provided by the client when the resource is
     * created. The name must be 1-63 characters long, and comply with
     * RFC1035. Specifically, the name must be 1-63 characters long and match
     * the regular expression `a-z?` which means the
     * first character must be a lowercase letter, and all following
     * characters must be a dash, lowercase letter, or digit, except the last
     * character, which cannot be a dash.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The share setting for reservations.
     * Structure is documented below.
     */
    shareSettings?: pulumi.Input<inputs.compute.ReservationShareSettings>;
    /**
     * Reservation for instances with specific machine shapes.
     * Structure is documented below.
     */
    specificReservation: pulumi.Input<inputs.compute.ReservationSpecificReservation>;
    /**
     * When set to true, only VMs that target this reservation by name can
     * consume this reservation. Otherwise, it can be consumed by VMs with
     * affinity for any reservation. Defaults to false.
     */
    specificReservationRequired?: pulumi.Input<boolean>;
    /**
     * The zone where the reservation is made.
     */
    zone: pulumi.Input<string>;
}
