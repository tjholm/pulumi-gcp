// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * A BigQuery Data Policy
 *
 * To get more information about DataPolicy, see:
 *
 * * [API documentation](https://cloud.google.com/bigquery/docs/reference/bigquerydatapolicy/rest/v1beta1/projects.locations.dataPolicies/create)
 * * How-to Guides
 *     * [Official Documentation](https://cloud.google.com/bigquery/docs/column-data-masking-intro)
 *
 * ## Example Usage
 * ### Bigquery Datapolicy Data Policy Basic
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const taxonomy = new gcp.datacatalog.Taxonomy("taxonomy", {
 *     region: "us-central1",
 *     displayName: "taxonomy",
 *     description: "A collection of policy tags",
 *     activatedPolicyTypes: ["FINE_GRAINED_ACCESS_CONTROL"],
 * });
 * const policyTag = new gcp.datacatalog.PolicyTag("policyTag", {
 *     taxonomy: taxonomy.id,
 *     displayName: "Low security",
 *     description: "A policy tag normally associated with low security items",
 * });
 * const dataPolicy = new gcp.bigquerydatapolicy.DataPolicy("dataPolicy", {
 *     location: "us-central1",
 *     dataPolicyId: "data_policy",
 *     policyTag: policyTag.name,
 *     dataPolicyType: "COLUMN_LEVEL_SECURITY_POLICY",
 * });
 * ```
 *
 * ## Import
 *
 * DataPolicy can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:bigquerydatapolicy/dataPolicy:DataPolicy default projects/{{project}}/locations/{{location}}/dataPolicies/{{data_policy_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:bigquerydatapolicy/dataPolicy:DataPolicy default {{project}}/{{location}}/{{data_policy_id}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:bigquerydatapolicy/dataPolicy:DataPolicy default {{location}}/{{data_policy_id}}
 * ```
 */
export class DataPolicy extends pulumi.CustomResource {
    /**
     * Get an existing DataPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DataPolicyState, opts?: pulumi.CustomResourceOptions): DataPolicy {
        return new DataPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:bigquerydatapolicy/dataPolicy:DataPolicy';

    /**
     * Returns true if the given object is an instance of DataPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DataPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DataPolicy.__pulumiType;
    }

    /**
     * The data masking policy that specifies the data masking rule to use.
     * Structure is documented below.
     */
    public readonly dataMaskingPolicy!: pulumi.Output<outputs.bigquerydatapolicy.DataPolicyDataMaskingPolicy | undefined>;
    /**
     * User-assigned (human readable) ID of the data policy that needs to be unique within a project. Used as {dataPolicyId} in part of the resource name.
     */
    public readonly dataPolicyId!: pulumi.Output<string>;
    /**
     * The enrollment level of the service.
     * Possible values are `COLUMN_LEVEL_SECURITY_POLICY` and `DATA_MASKING_POLICY`.
     */
    public readonly dataPolicyType!: pulumi.Output<string>;
    /**
     * The name of the location of the data policy.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource name of this data policy, in the format of projects/{project_number}/locations/{locationId}/dataPolicies/{dataPolicyId}.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Policy tag resource name, in the format of projects/{project_number}/locations/{locationId}/taxonomies/{taxonomyId}/policyTags/{policyTag_id}.
     */
    public readonly policyTag!: pulumi.Output<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    public readonly project!: pulumi.Output<string>;

    /**
     * Create a DataPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DataPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DataPolicyArgs | DataPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DataPolicyState | undefined;
            resourceInputs["dataMaskingPolicy"] = state ? state.dataMaskingPolicy : undefined;
            resourceInputs["dataPolicyId"] = state ? state.dataPolicyId : undefined;
            resourceInputs["dataPolicyType"] = state ? state.dataPolicyType : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["policyTag"] = state ? state.policyTag : undefined;
            resourceInputs["project"] = state ? state.project : undefined;
        } else {
            const args = argsOrState as DataPolicyArgs | undefined;
            if ((!args || args.dataPolicyId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataPolicyId'");
            }
            if ((!args || args.dataPolicyType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataPolicyType'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.policyTag === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyTag'");
            }
            resourceInputs["dataMaskingPolicy"] = args ? args.dataMaskingPolicy : undefined;
            resourceInputs["dataPolicyId"] = args ? args.dataPolicyId : undefined;
            resourceInputs["dataPolicyType"] = args ? args.dataPolicyType : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["policyTag"] = args ? args.policyTag : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["name"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DataPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DataPolicy resources.
 */
export interface DataPolicyState {
    /**
     * The data masking policy that specifies the data masking rule to use.
     * Structure is documented below.
     */
    dataMaskingPolicy?: pulumi.Input<inputs.bigquerydatapolicy.DataPolicyDataMaskingPolicy>;
    /**
     * User-assigned (human readable) ID of the data policy that needs to be unique within a project. Used as {dataPolicyId} in part of the resource name.
     */
    dataPolicyId?: pulumi.Input<string>;
    /**
     * The enrollment level of the service.
     * Possible values are `COLUMN_LEVEL_SECURITY_POLICY` and `DATA_MASKING_POLICY`.
     */
    dataPolicyType?: pulumi.Input<string>;
    /**
     * The name of the location of the data policy.
     */
    location?: pulumi.Input<string>;
    /**
     * Resource name of this data policy, in the format of projects/{project_number}/locations/{locationId}/dataPolicies/{dataPolicyId}.
     */
    name?: pulumi.Input<string>;
    /**
     * Policy tag resource name, in the format of projects/{project_number}/locations/{locationId}/taxonomies/{taxonomyId}/policyTags/{policyTag_id}.
     */
    policyTag?: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DataPolicy resource.
 */
export interface DataPolicyArgs {
    /**
     * The data masking policy that specifies the data masking rule to use.
     * Structure is documented below.
     */
    dataMaskingPolicy?: pulumi.Input<inputs.bigquerydatapolicy.DataPolicyDataMaskingPolicy>;
    /**
     * User-assigned (human readable) ID of the data policy that needs to be unique within a project. Used as {dataPolicyId} in part of the resource name.
     */
    dataPolicyId: pulumi.Input<string>;
    /**
     * The enrollment level of the service.
     * Possible values are `COLUMN_LEVEL_SECURITY_POLICY` and `DATA_MASKING_POLICY`.
     */
    dataPolicyType: pulumi.Input<string>;
    /**
     * The name of the location of the data policy.
     */
    location: pulumi.Input<string>;
    /**
     * Policy tag resource name, in the format of projects/{project_number}/locations/{locationId}/taxonomies/{taxonomyId}/policyTags/{policyTag_id}.
     */
    policyTag: pulumi.Input<string>;
    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     */
    project?: pulumi.Input<string>;
}
