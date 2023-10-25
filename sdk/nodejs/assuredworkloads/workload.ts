// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * The AssuredWorkloads Workload resource
 *
 * ## Example Usage
 * ### Basic_workload
 * A basic test of a assuredworkloads api
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const primary = new gcp.assuredworkloads.Workload("primary", {
 *     billingAccount: "billingAccounts/000000-0000000-0000000-000000",
 *     complianceRegime: "FEDRAMP_MODERATE",
 *     displayName: "Workload Example",
 *     kmsSettings: {
 *         nextRotationTime: "9999-10-02T15:01:23Z",
 *         rotationPeriod: "10368000s",
 *     },
 *     labels: {
 *         "label-one": "value-one",
 *     },
 *     location: "us-west1",
 *     organization: "123456789",
 *     provisionedResourcesParent: "folders/519620126891",
 *     resourceSettings: [
 *         {
 *             resourceType: "CONSUMER_PROJECT",
 *         },
 *         {
 *             resourceType: "ENCRYPTION_KEYS_PROJECT",
 *         },
 *         {
 *             resourceId: "ring",
 *             resourceType: "KEYRING",
 *         },
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Workload can be imported using any of these accepted formats
 *
 * ```sh
 *  $ pulumi import gcp:assuredworkloads/workload:Workload default organizations/{{organization}}/locations/{{location}}/workloads/{{name}}
 * ```
 *
 * ```sh
 *  $ pulumi import gcp:assuredworkloads/workload:Workload default {{organization}}/{{location}}/{{name}}
 * ```
 */
export class Workload extends pulumi.CustomResource {
    /**
     * Get an existing Workload resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WorkloadState, opts?: pulumi.CustomResourceOptions): Workload {
        return new Workload(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:assuredworkloads/workload:Workload';

    /**
     * Returns true if the given object is an instance of Workload.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Workload {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Workload.__pulumiType;
    }

    /**
     * Required. Input only. The billing account used for the resources which are direct children of workload. This billing account is initially associated with the resources created as part of Workload creation. After the initial creation of these resources, the customer can change the assigned billing account. The resource name has the form `billingAccounts/{billing_account_id}`. For example, 'billingAccounts/012345-567890-ABCDEF`.
     */
    public readonly billingAccount!: pulumi.Output<string>;
    /**
     * Required. Immutable. Compliance Regime associated with this workload. Possible values: COMPLIANCE_REGIME_UNSPECIFIED, IL4, CJIS, FEDRAMP_HIGH, FEDRAMP_MODERATE, US_REGIONAL_ACCESS, HIPAA, EU_REGIONS_AND_SUPPORT, CA_REGIONS_AND_SUPPORT, ITAR, AU_REGIONS_AND_US_SUPPORT, ASSURED_WORKLOADS_FOR_PARTNERS
     */
    public readonly complianceRegime!: pulumi.Output<string>;
    /**
     * Output only. Immutable. The Workload creation timestamp.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Required. The user-assigned display name of the Workload. When present it must be between 4 to 30 characters. Allowed characters are: lowercase and uppercase letters, numbers, hyphen, and spaces. Example: My Workload
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Input only. Settings used to create a CMEK crypto key. When set a project with a KMS CMEK key is provisioned. This field is mandatory for a subset of Compliance Regimes.
     */
    public readonly kmsSettings!: pulumi.Output<outputs.assuredworkloads.WorkloadKmsSettings | undefined>;
    /**
     * Optional. Labels applied to the workload.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The location for the resource
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Output only. The resource name of the workload.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The organization for the resource
     *
     *
     *
     * - - -
     */
    public readonly organization!: pulumi.Output<string>;
    /**
     * Input only. The parent resource for the resources managed by this Assured Workload. May be either an organization or a folder. Must be the same or a child of the Workload parent. If not specified all resources are created under the Workload parent. Formats: folders/{folder_id}, organizations/{organization_id}
     */
    public readonly provisionedResourcesParent!: pulumi.Output<string | undefined>;
    /**
     * Input only. Resource properties that are used to customize workload resources. These properties (such as custom project id) will be used to create workload resources if possible. This field is optional.
     */
    public readonly resourceSettings!: pulumi.Output<outputs.assuredworkloads.WorkloadResourceSetting[] | undefined>;
    /**
     * Output only. The resources associated with this workload. These resources will be created when creating the workload. If any of the projects already exist, the workload creation will fail. Always read only.
     */
    public /*out*/ readonly resources!: pulumi.Output<outputs.assuredworkloads.WorkloadResource[]>;

    /**
     * Create a Workload resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WorkloadArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WorkloadArgs | WorkloadState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WorkloadState | undefined;
            resourceInputs["billingAccount"] = state ? state.billingAccount : undefined;
            resourceInputs["complianceRegime"] = state ? state.complianceRegime : undefined;
            resourceInputs["createTime"] = state ? state.createTime : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["kmsSettings"] = state ? state.kmsSettings : undefined;
            resourceInputs["labels"] = state ? state.labels : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["organization"] = state ? state.organization : undefined;
            resourceInputs["provisionedResourcesParent"] = state ? state.provisionedResourcesParent : undefined;
            resourceInputs["resourceSettings"] = state ? state.resourceSettings : undefined;
            resourceInputs["resources"] = state ? state.resources : undefined;
        } else {
            const args = argsOrState as WorkloadArgs | undefined;
            if ((!args || args.billingAccount === undefined) && !opts.urn) {
                throw new Error("Missing required property 'billingAccount'");
            }
            if ((!args || args.complianceRegime === undefined) && !opts.urn) {
                throw new Error("Missing required property 'complianceRegime'");
            }
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.organization === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organization'");
            }
            resourceInputs["billingAccount"] = args ? args.billingAccount : undefined;
            resourceInputs["complianceRegime"] = args ? args.complianceRegime : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["kmsSettings"] = args ? args.kmsSettings : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["organization"] = args ? args.organization : undefined;
            resourceInputs["provisionedResourcesParent"] = args ? args.provisionedResourcesParent : undefined;
            resourceInputs["resourceSettings"] = args ? args.resourceSettings : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["resources"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Workload.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Workload resources.
 */
export interface WorkloadState {
    /**
     * Required. Input only. The billing account used for the resources which are direct children of workload. This billing account is initially associated with the resources created as part of Workload creation. After the initial creation of these resources, the customer can change the assigned billing account. The resource name has the form `billingAccounts/{billing_account_id}`. For example, 'billingAccounts/012345-567890-ABCDEF`.
     */
    billingAccount?: pulumi.Input<string>;
    /**
     * Required. Immutable. Compliance Regime associated with this workload. Possible values: COMPLIANCE_REGIME_UNSPECIFIED, IL4, CJIS, FEDRAMP_HIGH, FEDRAMP_MODERATE, US_REGIONAL_ACCESS, HIPAA, EU_REGIONS_AND_SUPPORT, CA_REGIONS_AND_SUPPORT, ITAR, AU_REGIONS_AND_US_SUPPORT, ASSURED_WORKLOADS_FOR_PARTNERS
     */
    complianceRegime?: pulumi.Input<string>;
    /**
     * Output only. Immutable. The Workload creation timestamp.
     */
    createTime?: pulumi.Input<string>;
    /**
     * Required. The user-assigned display name of the Workload. When present it must be between 4 to 30 characters. Allowed characters are: lowercase and uppercase letters, numbers, hyphen, and spaces. Example: My Workload
     */
    displayName?: pulumi.Input<string>;
    /**
     * Input only. Settings used to create a CMEK crypto key. When set a project with a KMS CMEK key is provisioned. This field is mandatory for a subset of Compliance Regimes.
     */
    kmsSettings?: pulumi.Input<inputs.assuredworkloads.WorkloadKmsSettings>;
    /**
     * Optional. Labels applied to the workload.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The location for the resource
     */
    location?: pulumi.Input<string>;
    /**
     * Output only. The resource name of the workload.
     */
    name?: pulumi.Input<string>;
    /**
     * The organization for the resource
     *
     *
     *
     * - - -
     */
    organization?: pulumi.Input<string>;
    /**
     * Input only. The parent resource for the resources managed by this Assured Workload. May be either an organization or a folder. Must be the same or a child of the Workload parent. If not specified all resources are created under the Workload parent. Formats: folders/{folder_id}, organizations/{organization_id}
     */
    provisionedResourcesParent?: pulumi.Input<string>;
    /**
     * Input only. Resource properties that are used to customize workload resources. These properties (such as custom project id) will be used to create workload resources if possible. This field is optional.
     */
    resourceSettings?: pulumi.Input<pulumi.Input<inputs.assuredworkloads.WorkloadResourceSetting>[]>;
    /**
     * Output only. The resources associated with this workload. These resources will be created when creating the workload. If any of the projects already exist, the workload creation will fail. Always read only.
     */
    resources?: pulumi.Input<pulumi.Input<inputs.assuredworkloads.WorkloadResource>[]>;
}

/**
 * The set of arguments for constructing a Workload resource.
 */
export interface WorkloadArgs {
    /**
     * Required. Input only. The billing account used for the resources which are direct children of workload. This billing account is initially associated with the resources created as part of Workload creation. After the initial creation of these resources, the customer can change the assigned billing account. The resource name has the form `billingAccounts/{billing_account_id}`. For example, 'billingAccounts/012345-567890-ABCDEF`.
     */
    billingAccount: pulumi.Input<string>;
    /**
     * Required. Immutable. Compliance Regime associated with this workload. Possible values: COMPLIANCE_REGIME_UNSPECIFIED, IL4, CJIS, FEDRAMP_HIGH, FEDRAMP_MODERATE, US_REGIONAL_ACCESS, HIPAA, EU_REGIONS_AND_SUPPORT, CA_REGIONS_AND_SUPPORT, ITAR, AU_REGIONS_AND_US_SUPPORT, ASSURED_WORKLOADS_FOR_PARTNERS
     */
    complianceRegime: pulumi.Input<string>;
    /**
     * Required. The user-assigned display name of the Workload. When present it must be between 4 to 30 characters. Allowed characters are: lowercase and uppercase letters, numbers, hyphen, and spaces. Example: My Workload
     */
    displayName: pulumi.Input<string>;
    /**
     * Input only. Settings used to create a CMEK crypto key. When set a project with a KMS CMEK key is provisioned. This field is mandatory for a subset of Compliance Regimes.
     */
    kmsSettings?: pulumi.Input<inputs.assuredworkloads.WorkloadKmsSettings>;
    /**
     * Optional. Labels applied to the workload.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The location for the resource
     */
    location: pulumi.Input<string>;
    /**
     * The organization for the resource
     *
     *
     *
     * - - -
     */
    organization: pulumi.Input<string>;
    /**
     * Input only. The parent resource for the resources managed by this Assured Workload. May be either an organization or a folder. Must be the same or a child of the Workload parent. If not specified all resources are created under the Workload parent. Formats: folders/{folder_id}, organizations/{organization_id}
     */
    provisionedResourcesParent?: pulumi.Input<string>;
    /**
     * Input only. Resource properties that are used to customize workload resources. These properties (such as custom project id) will be used to create workload resources if possible. This field is optional.
     */
    resourceSettings?: pulumi.Input<pulumi.Input<inputs.assuredworkloads.WorkloadResourceSetting>[]>;
}
