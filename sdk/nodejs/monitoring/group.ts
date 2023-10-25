// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * The description of a dynamic collection of monitored resources. Each group
 * has a filter that is matched against monitored resources and their
 * associated metadata. If a group's filter matches an available monitored
 * resource, then that resource is a member of that group.
 *
 * To get more information about Group, see:
 *
 * * [API documentation](https://cloud.google.com/monitoring/api/ref_v3/rest/v3/projects.groups)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/monitoring/groups/)
 *
 * ## Example Usage
 * ### Monitoring Group Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const basic = new gcp.monitoring.Group("basic", {
 *     displayName: "tf-test MonitoringGroup",
 *     filter: "resource.metadata.region=\"europe-west2\"",
 * });
 * ```
 * ### Monitoring Group Subgroup
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const parent = new gcp.monitoring.Group("parent", {
 *     displayName: "tf-test MonitoringParentGroup",
 *     filter: "resource.metadata.region=\"europe-west2\"",
 * });
 * const subgroup = new gcp.monitoring.Group("subgroup", {
 *     displayName: "tf-test MonitoringSubGroup",
 *     filter: "resource.metadata.region=\"europe-west2\"",
 *     parentName: parent.name,
 * });
 * ```
 *
 * ## Import
 *
 * Group can be imported using any of these accepted formats:
 *
 * ```sh
 *  $ pulumi import gcp:monitoring/group:Group default {{name}}
 * ```
 */
export class Group extends pulumi.CustomResource {
    /**
     * Get an existing Group resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupState, opts?: pulumi.CustomResourceOptions): Group {
        return new Group(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:monitoring/group:Group';

    /**
     * Returns true if the given object is an instance of Group.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Group {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Group.__pulumiType;
    }

    /**
     * A user-assigned name for this group, used only for display
     * purposes.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * The filter used to determine which monitored resources
     * belong to this group.
     *
     *
     * - - -
     */
    public readonly filter!: pulumi.Output<string>;
    /**
     * If true, the members of this group are considered to be a
     * cluster. The system can perform additional analysis on
     * groups that are clusters.
     */
    public readonly isCluster!: pulumi.Output<boolean | undefined>;
    /**
     * A unique identifier for this group. The format is
     * "projects/{project_id_or_number}/groups/{group_id}".
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The name of the group's parent, if it has one. The format is
     * "projects/{project_id_or_number}/groups/{group_id}". For
     * groups with no parent, parentName is the empty string, "".
     */
    public readonly parentName!: pulumi.Output<string | undefined>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a Group resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupArgs | GroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GroupState | undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["filter"] = state ? state.filter : undefined;
            resourceInputs["isCluster"] = state ? state.isCluster : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["parentName"] = state ? state.parentName : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as GroupArgs | undefined;
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.filter === undefined) && !opts.urn) {
                throw new Error("Missing required property 'filter'");
            }
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["filter"] = args ? args.filter : undefined;
            resourceInputs["isCluster"] = args ? args.isCluster : undefined;
            resourceInputs["parentName"] = args ? args.parentName : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["name"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Group.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Group resources.
 */
export interface GroupState {
    /**
     * A user-assigned name for this group, used only for display
     * purposes.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The filter used to determine which monitored resources
     * belong to this group.
     *
     *
     * - - -
     */
    filter?: pulumi.Input<string>;
    /**
     * If true, the members of this group are considered to be a
     * cluster. The system can perform additional analysis on
     * groups that are clusters.
     */
    isCluster?: pulumi.Input<boolean>;
    /**
     * A unique identifier for this group. The format is
     * "projects/{project_id_or_number}/groups/{group_id}".
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the group's parent, if it has one. The format is
     * "projects/{project_id_or_number}/groups/{group_id}". For
     * groups with no parent, parentName is the empty string, "".
     */
    parentName?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Group resource.
 */
export interface GroupArgs {
    /**
     * A user-assigned name for this group, used only for display
     * purposes.
     */
    displayName: pulumi.Input<string>;
    /**
     * The filter used to determine which monitored resources
     * belong to this group.
     *
     *
     * - - -
     */
    filter: pulumi.Input<string>;
    /**
     * If true, the members of this group are considered to be a
     * cluster. The system can perform additional analysis on
     * groups that are clusters.
     */
    isCluster?: pulumi.Input<boolean>;
    /**
     * The name of the group's parent, if it has one. The format is
     * "projects/{project_id_or_number}/groups/{group_id}". For
     * groups with no parent, parentName is the empty string, "".
     */
    parentName?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
}
