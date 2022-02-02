// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Three different resources help you manage your IAM policy for Healthcare DICOM store. Each of these resources serves a different use case:
 *
 * * `gcp.healthcare.DicomStoreIamPolicy`: Authoritative. Sets the IAM policy for the DICOM store and replaces any existing policy already attached.
 * * `gcp.healthcare.DicomStoreIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the DICOM store are preserved.
 * * `gcp.healthcare.DicomStoreIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the DICOM store are preserved.
 *
 * > **Note:** `gcp.healthcare.DicomStoreIamPolicy` **cannot** be used in conjunction with `gcp.healthcare.DicomStoreIamBinding` and `gcp.healthcare.DicomStoreIamMember` or they will fight over what your policy should be.
 *
 * > **Note:** `gcp.healthcare.DicomStoreIamBinding` resources **can be** used in conjunction with `gcp.healthcare.DicomStoreIamMember` resources **only if** they do not grant privilege to the same role.
 *
 * ## google\_healthcare\_dicom\_store\_iam\_policy
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const admin = gcp.organizations.getIAMPolicy({
 *     bindings: [{
 *         role: "roles/editor",
 *         members: ["user:jane@example.com"],
 *     }],
 * });
 * const dicomStore = new gcp.healthcare.DicomStoreIamPolicy("dicomStore", {
 *     dicomStoreId: "your-dicom-store-id",
 *     policyData: admin.then(admin => admin.policyData),
 * });
 * ```
 *
 * ## google\_healthcare\_dicom\_store\_iam\_binding
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const dicomStore = new gcp.healthcare.DicomStoreIamBinding("dicom_store", {
 *     dicomStoreId: "your-dicom-store-id",
 *     members: ["user:jane@example.com"],
 *     role: "roles/editor",
 * });
 * ```
 *
 * ## google\_healthcare\_dicom\_store\_iam\_member
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as gcp from "@pulumi/gcp";
 *
 * const dicomStore = new gcp.healthcare.DicomStoreIamMember("dicom_store", {
 *     dicomStoreId: "your-dicom-store-id",
 *     member: "user:jane@example.com",
 *     role: "roles/editor",
 * });
 * ```
 *
 * ## Import
 *
 * IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.
 *
 * This member resource can be imported using the `dicom_store_id`, role, and account e.g.
 *
 * ```sh
 *  $ pulumi import gcp:healthcare/dicomStoreIamPolicy:DicomStoreIamPolicy dicom_store_iam "your-project-id/location-name/dataset-name/dicom-store-name roles/viewer user:foo@example.com"
 * ```
 *
 *  IAM binding imports use space-delimited identifiers; the resource in question and the role.
 *
 * This binding resource can be imported using the `dicom_store_id` and role, e.g.
 *
 * ```sh
 *  $ pulumi import gcp:healthcare/dicomStoreIamPolicy:DicomStoreIamPolicy dicom_store_iam "your-project-id/location-name/dataset-name/dicom-store-name roles/viewer"
 * ```
 *
 *  IAM policy imports use the identifier of the resource in question.
 *
 * This policy resource can be imported using the `dicom_store_id`, role, and account e.g.
 *
 * ```sh
 *  $ pulumi import gcp:healthcare/dicomStoreIamPolicy:DicomStoreIamPolicy dicom_store_iam your-project-id/location-name/dataset-name/dicom-store-name
 * ```
 */
export class DicomStoreIamPolicy extends pulumi.CustomResource {
    /**
     * Get an existing DicomStoreIamPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DicomStoreIamPolicyState, opts?: pulumi.CustomResourceOptions): DicomStoreIamPolicy {
        return new DicomStoreIamPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'gcp:healthcare/dicomStoreIamPolicy:DicomStoreIamPolicy';

    /**
     * Returns true if the given object is an instance of DicomStoreIamPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DicomStoreIamPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DicomStoreIamPolicy.__pulumiType;
    }

    /**
     * The DICOM store ID, in the form
     * `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
     * `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
     * project setting will be used as a fallback.
     */
    public readonly dicomStoreId!: pulumi.Output<string>;
    /**
     * (Computed) The etag of the DICOM store's IAM policy.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     */
    public readonly policyData!: pulumi.Output<string>;

    /**
     * Create a DicomStoreIamPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DicomStoreIamPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DicomStoreIamPolicyArgs | DicomStoreIamPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DicomStoreIamPolicyState | undefined;
            resourceInputs["dicomStoreId"] = state ? state.dicomStoreId : undefined;
            resourceInputs["etag"] = state ? state.etag : undefined;
            resourceInputs["policyData"] = state ? state.policyData : undefined;
        } else {
            const args = argsOrState as DicomStoreIamPolicyArgs | undefined;
            if ((!args || args.dicomStoreId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dicomStoreId'");
            }
            if ((!args || args.policyData === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policyData'");
            }
            resourceInputs["dicomStoreId"] = args ? args.dicomStoreId : undefined;
            resourceInputs["policyData"] = args ? args.policyData : undefined;
            resourceInputs["etag"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DicomStoreIamPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DicomStoreIamPolicy resources.
 */
export interface DicomStoreIamPolicyState {
    /**
     * The DICOM store ID, in the form
     * `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
     * `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
     * project setting will be used as a fallback.
     */
    dicomStoreId?: pulumi.Input<string>;
    /**
     * (Computed) The etag of the DICOM store's IAM policy.
     */
    etag?: pulumi.Input<string>;
    /**
     * The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     */
    policyData?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DicomStoreIamPolicy resource.
 */
export interface DicomStoreIamPolicyArgs {
    /**
     * The DICOM store ID, in the form
     * `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
     * `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
     * project setting will be used as a fallback.
     */
    dicomStoreId: pulumi.Input<string>;
    /**
     * The policy data generated by
     * a `gcp.organizations.getIAMPolicy` data source.
     */
    policyData: pulumi.Input<string>;
}
