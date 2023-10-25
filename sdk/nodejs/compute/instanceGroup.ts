// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Creates a group of dissimilar Compute Engine virtual machine instances.
 * For more information, see [the official documentation](https://cloud.google.com/compute/docs/instance-groups/#unmanaged_instance_groups)
 * and [API](https://cloud.google.com/compute/docs/reference/latest/instanceGroups)
 *
 * ## Example Usage
 * ### Empty Instance Group
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const test = new gcp.compute.InstanceGroup("test", {
 *     description: "Test instance group",
 *     zone: "us-central1-a",
 *     network: google_compute_network["default"].id,
 * });
 * ```
 * ### Example Usage - With instances and named ports
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const webservers = new gcp.compute.InstanceGroup("webservers", {
 *     description: "Test instance group",
 *     instances: [
 *         google_compute_instance.test.id,
 *         google_compute_instance.test2.id,
 *     ],
 *     namedPorts: [
 *         {
 *             name: "http",
 *             port: 8080,
 *         },
 *         {
 *             name: "https",
 *             port: 8443,
 *         },
 *     ],
 *     zone: "us-central1-a",
 * });
 * ```
 *
 * ## Import
 *
 * Instance group can be imported using the `zone` and `name` with an optional `project`, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:compute/instanceGroup:InstanceGroup webservers us-central1-a/terraform-webservers
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/instanceGroup:InstanceGroup webservers big-project/us-central1-a/terraform-webservers
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:compute/instanceGroup:InstanceGroup webservers projects/big-project/zones/us-central1-a/instanceGroups/terraform-webservers
 * ```
 */
export class InstanceGroup extends pulumi.CustomResource {
    /**
     * Get an existing InstanceGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: InstanceGroupState, opts?: pulumi.CustomResourceOptions): InstanceGroup {
        return new InstanceGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/instanceGroup:InstanceGroup';

    /**
     * Returns true if the given object is an instance of InstanceGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstanceGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstanceGroup.__pulumiType;
    }

    /**
     * An optional textual description of the instance
     * group.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The list of instances in the group, in `selfLink` format.
     * When adding instances they must all be in the same network and zone as the instance group.
     */
    public readonly instances!: pulumi.Output<string[]>;
    /**
     * The name of the instance group. Must be 1-63
     * characters long and comply with
     * [RFC1035](https://www.ietf.org/rfc/rfc1035.txt). Supported characters
     * include lowercase letters, numbers, and hyphens.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The named port configuration. See the section below
     * for details on configuration. Structure is documented below.
     */
    public readonly namedPorts!: pulumi.Output<outputs.compute.InstanceGroupNamedPort[] | undefined>;
    /**
     * The URL of the network the instance group is in. If
     * this is different from the network where the instances are in, the creation
     * fails. Defaults to the network where the instances are in (if neither
     * `network` nor `instances` is specified, this field will be blank).
     */
    public readonly network!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The URI of the created resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * The number of instances in the group.
     */
    public /*out*/ readonly size!: pulumi.Output<number>;
    /**
     * The zone that this instance group should be created in.
     *
     * - - -
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a InstanceGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: InstanceGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: InstanceGroupArgs | InstanceGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as InstanceGroupState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["instances"] = state ? state.instances : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namedPorts"] = state ? state.namedPorts : undefined;
            resourceInputs["network"] = state ? state.network : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["selfLink"] = state ? state.selfLink : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
            resourceInputs["zone"] = state ? state.zone : undefined;
        } else {
            const args = argsOrState as InstanceGroupArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["instances"] = args ? args.instances : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namedPorts"] = args ? args.namedPorts : undefined;
            resourceInputs["network"] = args ? args.network : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["zone"] = args ? args.zone : undefined;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["size"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(InstanceGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering InstanceGroup resources.
 */
export interface InstanceGroupState {
    /**
     * An optional textual description of the instance
     * group.
     */
    description?: pulumi.Input<string>;
    /**
     * The list of instances in the group, in `selfLink` format.
     * When adding instances they must all be in the same network and zone as the instance group.
     */
    instances?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the instance group. Must be 1-63
     * characters long and comply with
     * [RFC1035](https://www.ietf.org/rfc/rfc1035.txt). Supported characters
     * include lowercase letters, numbers, and hyphens.
     */
    name?: pulumi.Input<string>;
    /**
     * The named port configuration. See the section below
     * for details on configuration. Structure is documented below.
     */
    namedPorts?: pulumi.Input<pulumi.Input<inputs.compute.InstanceGroupNamedPort>[]>;
    /**
     * The URL of the network the instance group is in. If
     * this is different from the network where the instances are in, the creation
     * fails. Defaults to the network where the instances are in (if neither
     * `network` nor `instances` is specified, this field will be blank).
     */
    network?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The URI of the created resource.
     */
    selfLink?: pulumi.Input<string>;
    /**
     * The number of instances in the group.
     */
    size?: pulumi.Input<number>;
    /**
     * The zone that this instance group should be created in.
     *
     * - - -
     */
    zone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a InstanceGroup resource.
 */
export interface InstanceGroupArgs {
    /**
     * An optional textual description of the instance
     * group.
     */
    description?: pulumi.Input<string>;
    /**
     * The list of instances in the group, in `selfLink` format.
     * When adding instances they must all be in the same network and zone as the instance group.
     */
    instances?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the instance group. Must be 1-63
     * characters long and comply with
     * [RFC1035](https://www.ietf.org/rfc/rfc1035.txt). Supported characters
     * include lowercase letters, numbers, and hyphens.
     */
    name?: pulumi.Input<string>;
    /**
     * The named port configuration. See the section below
     * for details on configuration. Structure is documented below.
     */
    namedPorts?: pulumi.Input<pulumi.Input<inputs.compute.InstanceGroupNamedPort>[]>;
    /**
     * The URL of the network the instance group is in. If
     * this is different from the network where the instances are in, the creation
     * fails. Defaults to the network where the instances are in (if neither
     * `network` nor `instances` is specified, this field will be blank).
     */
    network?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The zone that this instance group should be created in.
     *
     * - - -
     */
    zone?: pulumi.Input<string>;
}
