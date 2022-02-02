// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configures the Google Compute Engine
 * [Default Network Tier](https://cloud.google.com/network-tiers/docs/using-network-service-tiers#setting_the_tier_for_all_resources_in_a_project)
 * for a project.
 *
 * For more information, see,
 * [the Project API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/projects/setDefaultNetworkTier).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const defaultProjectDefaultNetworkTier = new gcp.compute.ProjectDefaultNetworkTier("default", {
 *     networkTier: "PREMIUM",
 * });
 * ```
 *
 * ## Import
 *
 * This resource can be imported using the project ID
 *
 * ```sh
 *  $ pulumi import gcp:compute/projectDefaultNetworkTier:ProjectDefaultNetworkTier default project-id`
 * ```
 */
export class ProjectDefaultNetworkTier extends pulumi.CustomResource {
    /**
     * Get an existing ProjectDefaultNetworkTier resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectDefaultNetworkTierState, opts?: pulumi.CustomResourceOptions): ProjectDefaultNetworkTier {
        return new ProjectDefaultNetworkTier(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:compute/projectDefaultNetworkTier:ProjectDefaultNetworkTier';

    /**
     * Returns true if the given object is an instance of ProjectDefaultNetworkTier.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ProjectDefaultNetworkTier {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ProjectDefaultNetworkTier.__pulumiType;
    }

    /**
     * The default network tier to be configured for the project.
     * This field can take the following values: `PREMIUM` or `STANDARD`.
     */
    public readonly networkTier!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a ProjectDefaultNetworkTier resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectDefaultNetworkTierArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectDefaultNetworkTierArgs | ProjectDefaultNetworkTierState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectDefaultNetworkTierState | undefined;
            resourceInputs["networkTier"] = state ? state.networkTier : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as ProjectDefaultNetworkTierArgs | undefined;
            if ((!args || args.networkTier === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkTier'");
            }
            resourceInputs["networkTier"] = args ? args.networkTier : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ProjectDefaultNetworkTier.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ProjectDefaultNetworkTier resources.
 */
export interface ProjectDefaultNetworkTierState {
    /**
     * The default network tier to be configured for the project.
     * This field can take the following values: `PREMIUM` or `STANDARD`.
     */
    networkTier?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ProjectDefaultNetworkTier resource.
 */
export interface ProjectDefaultNetworkTierArgs {
    /**
     * The default network tier to be configured for the project.
     * This field can take the following values: `PREMIUM` or `STANDARD`.
     */
    networkTier: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs. If it
     * is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
}
