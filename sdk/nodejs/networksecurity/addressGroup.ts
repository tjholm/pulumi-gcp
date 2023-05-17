// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 * ### Network Security Address Groups Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const _default = new gcp.networksecurity.AddressGroup("default", {
 *     parent: "projects/my-project-name",
 *     location: "us-central1",
 *     type: "IPV4",
 *     capacity: 100,
 *     items: ["208.80.154.224/32"],
 * }, {
 *     provider: google_beta,
 * });
 * ```
 * ### Network Security Address Groups Organization Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const _default = new gcp.networksecurity.AddressGroup("default", {
 *     parent: "organizations/123456789",
 *     location: "us-central1",
 *     type: "IPV4",
 *     capacity: 100,
 *     items: ["208.80.154.224/32"],
 * }, {
 *     provider: google_beta,
 * });
 * ```
 * ### Network Security Address Groups Advanced
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const _default = new gcp.networksecurity.AddressGroup("default", {
 *     parent: "projects/my-project-name",
 *     location: "us-central1",
 *     description: "my description",
 *     type: "IPV4",
 *     capacity: 100,
 *     items: ["208.80.154.224/32"],
 * }, {
 *     provider: google_beta,
 * });
 * ```
 *
 * ## Import
 *
 * AddressGroup can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:networksecurity/addressGroup:AddressGroup default {{parent}}/locations/{{location}}/addressGroups/{{name}}
 * ```
 */
export class AddressGroup extends pulumi.CustomResource {
    /**
     * Get an existing AddressGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AddressGroupState, opts?: pulumi.CustomResourceOptions): AddressGroup {
        return new AddressGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:networksecurity/addressGroup:AddressGroup';

    /**
     * Returns true if the given object is an instance of AddressGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AddressGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AddressGroup.__pulumiType;
    }

    /**
     * Capacity of the Address Group.
     */
    public readonly capacity!: pulumi.Output<number>;
    /**
     * The timestamp when the resource was created.
     * A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
     * Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z"
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Free-text description of the resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * List of items.
     */
    public readonly items!: pulumi.Output<string[] | undefined>;
    /**
     * Set of label tags associated with the AddressGroup resource.
     * An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The location of the gateway security policy.
     * The default value is `global`.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Name of the AddressGroup resource.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the parent this address group belongs to. Format: organizations/{organization_id} or projects/{project_id}.
     */
    public readonly parent!: pulumi.Output<string | undefined>;
    /**
     * The type of the Address Group. Possible values are "IPV4" or "IPV6".
     * Possible values are: `IPV4`, `IPV6`.
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The timestamp when the resource was updated.
     * A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
     * Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a AddressGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AddressGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AddressGroupArgs | AddressGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AddressGroupState | undefined;
            resourceInputs["capacity"] = state ? state.capacity : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["items"] = state ? state.items : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["parent"] = state ? state.parent : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["updateTime"] = state ? state.updateTime : undefined;
        } else {
            const args = argsOrState as AddressGroupArgs | undefined;
            if ((!args || args.capacity === undefined) && !opts.urn) {
                throw new Error("Missing required property 'capacity'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["capacity"] = args ? args.capacity : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["items"] = args ? args.items : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["parent"] = args ? args.parent : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AddressGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AddressGroup resources.
 */
export interface AddressGroupState {
    /**
     * Capacity of the Address Group.
     */
    capacity?: pulumi.Input<number>;
    /**
     * The timestamp when the resource was created.
     * A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
     * Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z"
     */
    createTime?: pulumi.Input<string>;
    /**
     * Free-text description of the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * List of items.
     */
    items?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Set of label tags associated with the AddressGroup resource.
     * An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The location of the gateway security policy.
     * The default value is `global`.
     */
    location?: pulumi.Input<string>;
    /**
     * Name of the AddressGroup resource.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the parent this address group belongs to. Format: organizations/{organization_id} or projects/{project_id}.
     */
    parent?: pulumi.Input<string>;
    /**
     * The type of the Address Group. Possible values are "IPV4" or "IPV6".
     * Possible values are: `IPV4`, `IPV6`.
     */
    type?: pulumi.Input<string>;
    /**
     * The timestamp when the resource was updated.
     * A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
     * Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
     */
    updateTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AddressGroup resource.
 */
export interface AddressGroupArgs {
    /**
     * Capacity of the Address Group.
     */
    capacity: pulumi.Input<number>;
    /**
     * Free-text description of the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * List of items.
     */
    items?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Set of label tags associated with the AddressGroup resource.
     * An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The location of the gateway security policy.
     * The default value is `global`.
     */
    location: pulumi.Input<string>;
    /**
     * Name of the AddressGroup resource.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the parent this address group belongs to. Format: organizations/{organization_id} or projects/{project_id}.
     */
    parent?: pulumi.Input<string>;
    /**
     * The type of the Address Group. Possible values are "IPV4" or "IPV6".
     * Possible values are: `IPV4`, `IPV6`.
     */
    type: pulumi.Input<string>;
}
