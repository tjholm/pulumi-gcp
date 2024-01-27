// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Describes an autoscaling policy for Dataproc cluster autoscaler.
 *
 * ## Example Usage
 * ### Dataproc Autoscaling Policy
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const asp = new gcp.dataproc.AutoscalingPolicy("asp", {
 *     policyId: "dataproc-policy",
 *     location: "us-central1",
 *     workerConfig: {
 *         maxInstances: 3,
 *     },
 *     basicAlgorithm: {
 *         yarnConfig: {
 *             gracefulDecommissionTimeout: "30s",
 *             scaleUpFactor: 0.5,
 *             scaleDownFactor: 0.5,
 *         },
 *     },
 * });
 * const basic = new gcp.dataproc.Cluster("basic", {
 *     region: "us-central1",
 *     clusterConfig: {
 *         autoscalingConfig: {
 *             policyUri: asp.name,
 *         },
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * AutoscalingPolicy can be imported using any of these accepted formats* `projects/{{project}}/locations/{{location}}/autoscalingPolicies/{{policy_id}}` * `{{project}}/{{location}}/{{policy_id}}` * `{{location}}/{{policy_id}}` When using the `pulumi import` command, AutoscalingPolicy can be imported using one of the formats above. For example
 *
 * ```sh
 *  $ pulumi import gcp:dataproc/autoscalingPolicy:AutoscalingPolicy default projects/{{project}}/locations/{{location}}/autoscalingPolicies/{{policy_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:dataproc/autoscalingPolicy:AutoscalingPolicy default {{project}}/{{location}}/{{policy_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:dataproc/autoscalingPolicy:AutoscalingPolicy default {{location}}/{{policy_id}}
 * ```
 */
export class AutoscalingPolicy extends pulumi.CustomResource {
    /**
     * Get an existing AutoscalingPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AutoscalingPolicyState, opts?: pulumi.CustomResourceOptions): AutoscalingPolicy {
        return new AutoscalingPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:dataproc/autoscalingPolicy:AutoscalingPolicy';

    /**
     * Returns true if the given object is an instance of AutoscalingPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AutoscalingPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AutoscalingPolicy.__pulumiType;
    }

    /**
     * Basic algorithm for autoscaling.
     * Structure is documented below.
     */
    public readonly basicAlgorithm!: pulumi.Output<outputs.dataproc.AutoscalingPolicyBasicAlgorithm | undefined>;
    /**
     * The  location where the autoscaling policy should reside.
     * The default value is `global`.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The "resource name" of the autoscaling policy.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The policy id. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
     * and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between
     * 3 and 50 characters.
     *
     *
     * - - -
     */
    public readonly policyId!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Describes how the autoscaler will operate for secondary workers.
     * Structure is documented below.
     */
    public readonly secondaryWorkerConfig!: pulumi.Output<outputs.dataproc.AutoscalingPolicySecondaryWorkerConfig | undefined>;
    /**
     * Describes how the autoscaler will operate for primary workers.
     * Structure is documented below.
     */
    public readonly workerConfig!: pulumi.Output<outputs.dataproc.AutoscalingPolicyWorkerConfig | undefined>;

    /**
     * Create a AutoscalingPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AutoscalingPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AutoscalingPolicyArgs | AutoscalingPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AutoscalingPolicyState | undefined;
            resourceInputs["basicAlgorithm"] = state ? state.basicAlgorithm : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["policyId"] = state ? state.policyId : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
            resourceInputs["secondaryWorkerConfig"] = state ? state.secondaryWorkerConfig : undefined;
            resourceInputs["workerConfig"] = state ? state.workerConfig : undefined;
        } else {
            const args = argsOrState as AutoscalingPolicyArgs | undefined;
            if ((!args || args.policyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyId'");
            }
            resourceInputs["basicAlgorithm"] = args ? args.basicAlgorithm : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["policyId"] = args ? args.policyId : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["secondaryWorkerConfig"] = args ? args.secondaryWorkerConfig : undefined;
            resourceInputs["workerConfig"] = args ? args.workerConfig : undefined;
            resourceInputs["name"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AutoscalingPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AutoscalingPolicy resources.
 */
export interface AutoscalingPolicyState {
    /**
     * Basic algorithm for autoscaling.
     * Structure is documented below.
     */
    basicAlgorithm?: pulumi.Input<inputs.dataproc.AutoscalingPolicyBasicAlgorithm>;
    /**
     * The  location where the autoscaling policy should reside.
     * The default value is `global`.
     */
    location?: pulumi.Input<string>;
    /**
     * The "resource name" of the autoscaling policy.
     */
    name?: pulumi.Input<string>;
    /**
     * The policy id. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
     * and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between
     * 3 and 50 characters.
     *
     *
     * - - -
     */
    policyId?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Describes how the autoscaler will operate for secondary workers.
     * Structure is documented below.
     */
    secondaryWorkerConfig?: pulumi.Input<inputs.dataproc.AutoscalingPolicySecondaryWorkerConfig>;
    /**
     * Describes how the autoscaler will operate for primary workers.
     * Structure is documented below.
     */
    workerConfig?: pulumi.Input<inputs.dataproc.AutoscalingPolicyWorkerConfig>;
}

/**
 * The set of arguments for constructing a AutoscalingPolicy resource.
 */
export interface AutoscalingPolicyArgs {
    /**
     * Basic algorithm for autoscaling.
     * Structure is documented below.
     */
    basicAlgorithm?: pulumi.Input<inputs.dataproc.AutoscalingPolicyBasicAlgorithm>;
    /**
     * The  location where the autoscaling policy should reside.
     * The default value is `global`.
     */
    location?: pulumi.Input<string>;
    /**
     * The policy id. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_),
     * and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between
     * 3 and 50 characters.
     *
     *
     * - - -
     */
    policyId: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
    /**
     * Describes how the autoscaler will operate for secondary workers.
     * Structure is documented below.
     */
    secondaryWorkerConfig?: pulumi.Input<inputs.dataproc.AutoscalingPolicySecondaryWorkerConfig>;
    /**
     * Describes how the autoscaler will operate for primary workers.
     * Structure is documented below.
     */
    workerConfig?: pulumi.Input<inputs.dataproc.AutoscalingPolicyWorkerConfig>;
}
