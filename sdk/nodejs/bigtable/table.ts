// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Creates a Google Cloud Bigtable table inside an instance. For more information see
 * [the official documentation](https://cloud.google.com/bigtable/) and
 * [API](https://cloud.google.com/bigtable/docs/go/reference).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const instance = new gcp.bigtable.Instance("instance", {clusters: [{
 *     clusterId: "tf-instance-cluster",
 *     zone: "us-central1-b",
 *     numNodes: 3,
 *     storageType: "HDD",
 * }]});
 * const table = new gcp.bigtable.Table("table", {
 *     instanceName: instance.name,
 *     splitKeys: [
 *         "a",
 *         "b",
 *         "c",
 *     ],
 *     columnFamilies: [
 *         {
 *             family: "family-first",
 *         },
 *         {
 *             family: "family-second",
 *         },
 *     ],
 *     changeStreamRetention: "24h0m0s",
 * });
 * ```
 *
 * ## Import
 *
 * Bigtable Tables can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:bigtable/table:Table default projects/{{project}}/instances/{{instance_name}}/tables/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:bigtable/table:Table default {{project}}/{{instance_name}}/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:bigtable/table:Table default {{instance_name}}/{{name}}
 * ```
 *
 *  The following fields can't be read and will show diffs if set in config when imported- `split_keys`
 */
export class Table extends pulumi.CustomResource {
    /**
     * Get an existing Table resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TableState, opts?: pulumi.CustomResourceOptions): Table {
        return new Table(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:bigtable/table:Table';

    /**
     * Returns true if the given object is an instance of Table.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Table {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Table.__pulumiType;
    }

    /**
     * Duration to retain change stream data for the table. Set to 0 to disable. Must be between 1 and 7 days.
     *
     * -----
     */
    public readonly changeStreamRetention!: pulumi.Output<string>;
    /**
     * A group of columns within a table which share a common configuration. This can be specified multiple times. Structure is documented below.
     */
    public readonly columnFamilies!: pulumi.Output<outputs.bigtable.TableColumnFamily[] | undefined>;
    /**
     * A field to make the table protected against data loss i.e. when set to PROTECTED, deleting the table, the column families in the table, and the instance containing the table would be prohibited. If not provided, deletion protection will be set to UNPROTECTED.
     */
    public readonly deletionProtection!: pulumi.Output<string>;
    /**
     * The name of the Bigtable instance.
     */
    public readonly instanceName!: pulumi.Output<string>;
    /**
     * The name of the table. Must be 1-50 characters and must only contain hyphens, underscores, periods, letters and numbers.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * A list of predefined keys to split the table on.
     * !> **Warning:** Modifying the `splitKeys` of an existing table will cause the provider
     * to delete/recreate the entire `gcp.bigtable.Table` resource.
     */
    public readonly splitKeys!: pulumi.Output<string[] | undefined>;

    /**
     * Create a Table resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TableArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TableArgs | TableState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TableState | undefined;
            resourceInputs["changeStreamRetention"] = state ? state.changeStreamRetention : undefined;
            resourceInputs["columnFamilies"] = state ? state.columnFamilies : undefined;
            resourceInputs["deletionProtection"] = state ? state.deletionProtection : undefined;
            resourceInputs["instanceName"] = state ? state.instanceName : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["splitKeys"] = state ? state.splitKeys : undefined;
        } else {
            const args = argsOrState as TableArgs | undefined;
            if ((!args || args.instanceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceName'");
            }
            resourceInputs["changeStreamRetention"] = args ? args.changeStreamRetention : undefined;
            resourceInputs["columnFamilies"] = args ? args.columnFamilies : undefined;
            resourceInputs["deletionProtection"] = args ? args.deletionProtection : undefined;
            resourceInputs["instanceName"] = args ? args.instanceName : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["splitKeys"] = args ? args.splitKeys : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Table.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Table resources.
 */
export interface TableState {
    /**
     * Duration to retain change stream data for the table. Set to 0 to disable. Must be between 1 and 7 days.
     *
     * -----
     */
    changeStreamRetention?: pulumi.Input<string>;
    /**
     * A group of columns within a table which share a common configuration. This can be specified multiple times. Structure is documented below.
     */
    columnFamilies?: pulumi.Input<pulumi.Input<inputs.bigtable.TableColumnFamily>[]>;
    /**
     * A field to make the table protected against data loss i.e. when set to PROTECTED, deleting the table, the column families in the table, and the instance containing the table would be prohibited. If not provided, deletion protection will be set to UNPROTECTED.
     */
    deletionProtection?: pulumi.Input<string>;
    /**
     * The name of the Bigtable instance.
     */
    instanceName?: pulumi.Input<string>;
    /**
     * The name of the table. Must be 1-50 characters and must only contain hyphens, underscores, periods, letters and numbers.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * A list of predefined keys to split the table on.
     * !> **Warning:** Modifying the `splitKeys` of an existing table will cause the provider
     * to delete/recreate the entire `gcp.bigtable.Table` resource.
     */
    splitKeys?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Table resource.
 */
export interface TableArgs {
    /**
     * Duration to retain change stream data for the table. Set to 0 to disable. Must be between 1 and 7 days.
     *
     * -----
     */
    changeStreamRetention?: pulumi.Input<string>;
    /**
     * A group of columns within a table which share a common configuration. This can be specified multiple times. Structure is documented below.
     */
    columnFamilies?: pulumi.Input<pulumi.Input<inputs.bigtable.TableColumnFamily>[]>;
    /**
     * A field to make the table protected against data loss i.e. when set to PROTECTED, deleting the table, the column families in the table, and the instance containing the table would be prohibited. If not provided, deletion protection will be set to UNPROTECTED.
     */
    deletionProtection?: pulumi.Input<string>;
    /**
     * The name of the Bigtable instance.
     */
    instanceName: pulumi.Input<string>;
    /**
     * The name of the table. Must be 1-50 characters and must only contain hyphens, underscores, periods, letters and numbers.
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * A list of predefined keys to split the table on.
     * !> **Warning:** Modifying the `splitKeys` of an existing table will cause the provider
     * to delete/recreate the entire `gcp.bigtable.Table` resource.
     */
    splitKeys?: pulumi.Input<pulumi.Input<string>[]>;
}
