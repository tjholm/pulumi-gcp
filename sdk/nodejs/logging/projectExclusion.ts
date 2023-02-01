// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const my_exclusion = new gcp.logging.ProjectExclusion("my-exclusion", {
 *     description: "Exclude GCE instance debug logs",
 *     filter: "resource.type = gce_instance AND severity <= DEBUG",
 * });
 * ```
 *
 * ## Import
 *
 * Project-level logging exclusions can be imported using their URI, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:logging/projectExclusion:ProjectExclusion my_exclusion projects/my-project/exclusions/my-exclusion
 * ```
 */
export class ProjectExclusion extends pulumi.CustomResource {
    /**
     * Get an existing ProjectExclusion resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectExclusionState, opts?: pulumi.CustomResourceOptions): ProjectExclusion {
        return new ProjectExclusion(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:logging/projectExclusion:ProjectExclusion';

    /**
     * Returns true if the given object is an instance of ProjectExclusion.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProjectExclusion {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProjectExclusion.__pulumiType;
    }

    /**
     * A human-readable description.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Whether this exclusion rule should be disabled or not. This defaults to
     * false.
     */
    public readonly disabled!: pulumi.Output<boolean | undefined>;
    /**
     * The filter to apply when excluding logs. Only log entries that match the filter are excluded.
     * See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
     * write a filter.
     */
    public readonly filter!: pulumi.Output<string>;
    /**
     * The name of the logging exclusion.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The project to create the exclusion in. If omitted, the project associated with the provider is
     * used.
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a ProjectExclusion resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectExclusionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectExclusionArgs | ProjectExclusionState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectExclusionState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["disabled"] = state ? state.disabled : undefined;
            resourceInputs["filter"] = state ? state.filter : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as ProjectExclusionArgs | undefined;
            if ((!args || args.filter === undefined) && !opts.urn) {
                throw new Error("Missing required property 'filter'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["disabled"] = args ? args.disabled : undefined;
            resourceInputs["filter"] = args ? args.filter : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ProjectExclusion.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProjectExclusion resources.
 */
export interface ProjectExclusionState {
    /**
     * A human-readable description.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether this exclusion rule should be disabled or not. This defaults to
     * false.
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * The filter to apply when excluding logs. Only log entries that match the filter are excluded.
     * See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
     * write a filter.
     */
    filter?: pulumi.Input<string>;
    /**
     * The name of the logging exclusion.
     */
    name?: pulumi.Input<string>;
    /**
     * The project to create the exclusion in. If omitted, the project associated with the provider is
     * used.
     */
    project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProjectExclusion resource.
 */
export interface ProjectExclusionArgs {
    /**
     * A human-readable description.
     */
    description?: pulumi.Input<string>;
    /**
     * Whether this exclusion rule should be disabled or not. This defaults to
     * false.
     */
    disabled?: pulumi.Input<boolean>;
    /**
     * The filter to apply when excluding logs. Only log entries that match the filter are excluded.
     * See [Advanced Log Filters](https://cloud.google.com/logging/docs/view/advanced-filters) for information on how to
     * write a filter.
     */
    filter: pulumi.Input<string>;
    /**
     * The name of the logging exclusion.
     */
    name?: pulumi.Input<string>;
    /**
     * The project to create the exclusion in. If omitted, the project associated with the provider is
     * used.
     */
    project?: pulumi.Input<string>;
}
