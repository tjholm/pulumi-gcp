// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Creates a Google Cloud Bigtable GC Policy inside a family. For more information see
 * [the official documentation](https://cloud.google.com/bigtable/) and
 * [API](https://cloud.google.com/bigtable/docs/go/reference).
 *
 * > **Warning**: We don't recommend having multiple GC policies for the same column
 * family as it may result in unexpected behavior.
 *
 * > **Note**: GC policies associated with a replicated table cannot be destroyed directly.
 * Destroying a GC policy is translated into never perform garbage collection, this is
 * considered relaxing from pure age-based or version-based GC policy, hence not allowed.
 * The workaround is unreplicating the instance first by updating the instance to have one
 * cluster.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const instance = new gcp.bigtable.Instance("instance", {clusters: [{
 *     clusterId: "tf-instance-cluster",
 *     numNodes: 3,
 *     storageType: "HDD",
 * }]});
 * const table = new gcp.bigtable.Table("table", {
 *     instanceName: instance.name,
 *     columnFamilies: [{
 *         family: "name",
 *     }],
 * });
 * const policy = new gcp.bigtable.GCPolicy("policy", {
 *     instanceName: instance.name,
 *     table: table.name,
 *     columnFamily: "name",
 *     maxAge: {
 *         duration: "168h",
 *     },
 * });
 * ```
 *
 * Multiple conditions is also supported. `UNION` when any of its sub-policies apply (OR). `INTERSECTION` when all its sub-policies apply (AND)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const policy = new gcp.bigtable.GCPolicy("policy", {
 *     instanceName: google_bigtable_instance.instance.name,
 *     table: google_bigtable_table.table.name,
 *     columnFamily: "name",
 *     mode: "UNION",
 *     maxAge: {
 *         duration: "168h",
 *     },
 *     maxVersions: [{
 *         number: 10,
 *     }],
 * });
 * ```
 *
 * For complex, nested policies, an optional `gcRules` field are supported. This field
 * conflicts with `mode`, `maxAge` and `maxVersion`. This field is a serialized JSON
 * string. Example:
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const instance = new gcp.bigtable.Instance("instance", {
 *     clusters: [{
 *         clusterId: "cid",
 *         zone: "us-central1-b",
 *     }],
 *     instanceType: "DEVELOPMENT",
 *     deletionProtection: false,
 * });
 * const table = new gcp.bigtable.Table("table", {
 *     instanceName: instance.id,
 *     columnFamilies: [{
 *         family: "cf1",
 *     }],
 * });
 * const policy = new gcp.bigtable.GCPolicy("policy", {
 *     instanceName: instance.id,
 *     table: table.name,
 *     columnFamily: "cf1",
 *     gcRules: `{
 *   "mode": "union",
 *   "rules": [
 *     {
 *       "max_age": "10h"
 *     },
 *     {
 *       "mode": "intersection",
 *       "rules": [
 *         {
 *           "max_age": "2h"
 *         },
 *         {
 *           "max_version": 2
 *         }
 *       ]
 *     }
 *   ]
 * }
 * `,
 * });
 * ```
 * This is equivalent to running the following `cbt` command:
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * ```
 *
 * ## Import
 *
 * This resource does not support import.
 */
export class GCPolicy extends pulumi.CustomResource {
    /**
     * Get an existing GCPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GCPolicyState, opts?: pulumi.CustomResourceOptions): GCPolicy {
        return new GCPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:bigtable/gCPolicy:GCPolicy';

    /**
     * Returns true if the given object is an instance of GCPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GCPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GCPolicy.__pulumiType;
    }

    /**
     * The name of the column family.
     */
    public readonly columnFamily!: pulumi.Output<string>;
    /**
     * Serialized JSON object to represent a more complex GC policy. Conflicts with `mode`, `maxAge` and `maxVersion`. Conflicts with `mode`, `maxAge` and `maxVersion`.
     */
    public readonly gcRules!: pulumi.Output<string | undefined>;
    /**
     * The name of the Bigtable instance.
     */
    public readonly instanceName!: pulumi.Output<string>;
    /**
     * GC policy that applies to all cells older than the given age.
     */
    public readonly maxAge!: pulumi.Output<outputs.bigtable.GCPolicyMaxAge | undefined>;
    /**
     * GC policy that applies to all versions of a cell except for the most recent.
     */
    public readonly maxVersions!: pulumi.Output<outputs.bigtable.GCPolicyMaxVersion[] | undefined>;
    /**
     * If multiple policies are set, you should choose between `UNION` OR `INTERSECTION`.
     */
    public readonly mode!: pulumi.Output<string | undefined>;
    /**
     * The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The name of the table.
     */
    public readonly table!: pulumi.Output<string>;

    /**
     * Create a GCPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GCPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GCPolicyArgs | GCPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GCPolicyState | undefined;
            resourceInputs["columnFamily"] = state ? state.columnFamily : undefined;
            resourceInputs["gcRules"] = state ? state.gcRules : undefined;
            resourceInputs["instanceName"] = state ? state.instanceName : undefined;
            resourceInputs["maxAge"] = state ? state.maxAge : undefined;
            resourceInputs["maxVersions"] = state ? state.maxVersions : undefined;
            resourceInputs["mode"] = state ? state.mode : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["table"] = state ? state.table : undefined;
        } else {
            const args = argsOrState as GCPolicyArgs | undefined;
            if ((!args || args.columnFamily === undefined) && !opts.urn) {
                throw new Error("Missing required property 'columnFamily'");
            }
            if ((!args || args.instanceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceName'");
            }
            if ((!args || args.table === undefined) && !opts.urn) {
                throw new Error("Missing required property 'table'");
            }
            resourceInputs["columnFamily"] = args ? args.columnFamily : undefined;
            resourceInputs["gcRules"] = args ? args.gcRules : undefined;
            resourceInputs["instanceName"] = args ? args.instanceName : undefined;
            resourceInputs["maxAge"] = args ? args.maxAge : undefined;
            resourceInputs["maxVersions"] = args ? args.maxVersions : undefined;
            resourceInputs["mode"] = args ? args.mode : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["table"] = args ? args.table : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GCPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GCPolicy resources.
 */
export interface GCPolicyState {
    /**
     * The name of the column family.
     */
    columnFamily?: pulumi.Input<string>;
    /**
     * Serialized JSON object to represent a more complex GC policy. Conflicts with `mode`, `maxAge` and `maxVersion`. Conflicts with `mode`, `maxAge` and `maxVersion`.
     */
    gcRules?: pulumi.Input<string>;
    /**
     * The name of the Bigtable instance.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * GC policy that applies to all cells older than the given age.
     */
    maxAge?: pulumi.Input<inputs.bigtable.GCPolicyMaxAge>;
    /**
     * GC policy that applies to all versions of a cell except for the most recent.
     */
    maxVersions?: pulumi.Input<pulumi.Input<inputs.bigtable.GCPolicyMaxVersion>[]>;
    /**
     * If multiple policies are set, you should choose between `UNION` OR `INTERSECTION`.
     */
    mode?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The name of the table.
     */
    table?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GCPolicy resource.
 */
export interface GCPolicyArgs {
    /**
     * The name of the column family.
     */
    columnFamily: pulumi.Input<string>;
    /**
     * Serialized JSON object to represent a more complex GC policy. Conflicts with `mode`, `maxAge` and `maxVersion`. Conflicts with `mode`, `maxAge` and `maxVersion`.
     */
    gcRules?: pulumi.Input<string>;
    /**
     * The name of the Bigtable instance.
     */
    instanceName: pulumi.Input<string>;
    /**
     * GC policy that applies to all cells older than the given age.
     */
    maxAge?: pulumi.Input<inputs.bigtable.GCPolicyMaxAge>;
    /**
     * GC policy that applies to all versions of a cell except for the most recent.
     */
    maxVersions?: pulumi.Input<pulumi.Input<inputs.bigtable.GCPolicyMaxVersion>[]>;
    /**
     * If multiple policies are set, you should choose between `UNION` OR `INTERSECTION`.
     */
    mode?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * The name of the table.
     */
    table: pulumi.Input<string>;
}
