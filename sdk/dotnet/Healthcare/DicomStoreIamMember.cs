// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Healthcare
{
    /// <summary>
    /// ## Import
    /// 
    /// IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.
    /// 
    /// This member resource can be imported using the `dicom_store_id`, role, and account e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:healthcare/dicomStoreIamMember:DicomStoreIamMember dicom_store_iam "your-project-id/location-name/dataset-name/dicom-store-name roles/viewer user:foo@example.com"
    /// ```
    /// 
    ///  IAM binding imports use space-delimited identifiers; the resource in question and the role.
    /// 
    /// This binding resource can be imported using the `dicom_store_id` and role, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:healthcare/dicomStoreIamMember:DicomStoreIamMember dicom_store_iam "your-project-id/location-name/dataset-name/dicom-store-name roles/viewer"
    /// ```
    /// 
    ///  IAM policy imports use the identifier of the resource in question.
    /// 
    /// This policy resource can be imported using the `dicom_store_id`, role, and account e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:healthcare/dicomStoreIamMember:DicomStoreIamMember dicom_store_iam your-project-id/location-name/dataset-name/dicom-store-name
    /// ```
    /// </summary>
    [GcpResourceType("gcp:healthcare/dicomStoreIamMember:DicomStoreIamMember")]
    public partial class DicomStoreIamMember : global::Pulumi.CustomResource
    {
        [Output("condition")]
        public Output<Outputs.DicomStoreIamMemberCondition?> Condition { get; private set; } = null!;

        /// <summary>
        /// The DICOM store ID, in the form
        /// `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
        /// `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
        /// project setting will be used as a fallback.
        /// </summary>
        [Output("dicomStoreId")]
        public Output<string> DicomStoreId { get; private set; } = null!;

        /// <summary>
        /// (Computed) The etag of the DICOM store's IAM policy.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        [Output("member")]
        public Output<string> Member { get; private set; } = null!;

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.healthcare.DicomStoreIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;


        /// <summary>
        /// Create a DicomStoreIamMember resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DicomStoreIamMember(string name, DicomStoreIamMemberArgs args, CustomResourceOptions? options = null)
            : base("gcp:healthcare/dicomStoreIamMember:DicomStoreIamMember", name, args ?? new DicomStoreIamMemberArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DicomStoreIamMember(string name, Input<string> id, DicomStoreIamMemberState? state = null, CustomResourceOptions? options = null)
            : base("gcp:healthcare/dicomStoreIamMember:DicomStoreIamMember", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DicomStoreIamMember resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DicomStoreIamMember Get(string name, Input<string> id, DicomStoreIamMemberState? state = null, CustomResourceOptions? options = null)
        {
            return new DicomStoreIamMember(name, id, state, options);
        }
    }

    public sealed class DicomStoreIamMemberArgs : global::Pulumi.ResourceArgs
    {
        [Input("condition")]
        public Input<Inputs.DicomStoreIamMemberConditionArgs>? Condition { get; set; }

        /// <summary>
        /// The DICOM store ID, in the form
        /// `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
        /// `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
        /// project setting will be used as a fallback.
        /// </summary>
        [Input("dicomStoreId", required: true)]
        public Input<string> DicomStoreId { get; set; } = null!;

        [Input("member", required: true)]
        public Input<string> Member { get; set; } = null!;

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.healthcare.DicomStoreIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        public DicomStoreIamMemberArgs()
        {
        }
        public static new DicomStoreIamMemberArgs Empty => new DicomStoreIamMemberArgs();
    }

    public sealed class DicomStoreIamMemberState : global::Pulumi.ResourceArgs
    {
        [Input("condition")]
        public Input<Inputs.DicomStoreIamMemberConditionGetArgs>? Condition { get; set; }

        /// <summary>
        /// The DICOM store ID, in the form
        /// `{project_id}/{location_name}/{dataset_name}/{dicom_store_name}` or
        /// `{location_name}/{dataset_name}/{dicom_store_name}`. In the second form, the provider's
        /// project setting will be used as a fallback.
        /// </summary>
        [Input("dicomStoreId")]
        public Input<string>? DicomStoreId { get; set; }

        /// <summary>
        /// (Computed) The etag of the DICOM store's IAM policy.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        [Input("member")]
        public Input<string>? Member { get; set; }

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.healthcare.DicomStoreIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        public DicomStoreIamMemberState()
        {
        }
        public static new DicomStoreIamMemberState Empty => new DicomStoreIamMemberState();
    }
}
