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
    /// Three different resources help you manage your IAM policy for Healthcare DICOM store. Each of these resources serves a different use case:
    /// 
    /// * `gcp.healthcare.DicomStoreIamPolicy`: Authoritative. Sets the IAM policy for the DICOM store and replaces any existing policy already attached.
    /// * `gcp.healthcare.DicomStoreIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the DICOM store are preserved.
    /// * `gcp.healthcare.DicomStoreIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the DICOM store are preserved.
    /// 
    /// &gt; **Note:** `gcp.healthcare.DicomStoreIamPolicy` **cannot** be used in conjunction with `gcp.healthcare.DicomStoreIamBinding` and `gcp.healthcare.DicomStoreIamMember` or they will fight over what your policy should be.
    /// 
    /// &gt; **Note:** `gcp.healthcare.DicomStoreIamBinding` resources **can be** used in conjunction with `gcp.healthcare.DicomStoreIamMember` resources **only if** they do not grant privilege to the same role.
    /// 
    /// ## google\_healthcare\_dicom\_store\_iam\_policy
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var admin = Gcp.Organizations.GetIAMPolicy.Invoke(new()
    ///     {
    ///         Bindings = new[]
    ///         {
    ///             new Gcp.Organizations.Inputs.GetIAMPolicyBindingInputArgs
    ///             {
    ///                 Role = "roles/editor",
    ///                 Members = new[]
    ///                 {
    ///                     "user:jane@example.com",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var dicomStore = new Gcp.Healthcare.DicomStoreIamPolicy("dicomStore", new()
    ///     {
    ///         DicomStoreId = "your-dicom-store-id",
    ///         PolicyData = admin.Apply(getIAMPolicyResult =&gt; getIAMPolicyResult.PolicyData),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## google\_healthcare\_dicom\_store\_iam\_binding
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var dicomStore = new Gcp.Healthcare.DicomStoreIamBinding("dicomStore", new()
    ///     {
    ///         DicomStoreId = "your-dicom-store-id",
    ///         Members = new[]
    ///         {
    ///             "user:jane@example.com",
    ///         },
    ///         Role = "roles/editor",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## google\_healthcare\_dicom\_store\_iam\_member
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var dicomStore = new Gcp.Healthcare.DicomStoreIamMember("dicomStore", new()
    ///     {
    ///         DicomStoreId = "your-dicom-store-id",
    ///         Member = "user:jane@example.com",
    ///         Role = "roles/editor",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ### Importing IAM policies IAM policy imports use the identifier of the Healthcare DICOM store resource. For example* `"{{project_id}}/{{location}}/{{dataset}}/{{dicom_store}}"` An `import` block (Terraform v1.5.0 and later) can be used to import IAM policiestf import {
    /// 
    ///  id = "{{project_id}}/{{location}}/{{dataset}}/{{dicom_store}}"
    /// 
    ///  to = google_healthcare_dicom_store_iam_policy.default } The `pulumi import` command can also be used
    /// 
    /// ```sh
    ///  $ pulumi import gcp:healthcare/dicomStoreIamMember:DicomStoreIamMember default {{project_id}}/{{location}}/{{dataset}}/{{dicom_store}}
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
        /// 
        /// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
        /// Each entry can have one of the following values:
        /// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
        /// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
        /// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
        /// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
        /// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
        /// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
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
        /// 
        /// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
        /// Each entry can have one of the following values:
        /// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
        /// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
        /// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
        /// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
        /// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
        /// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
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
        /// 
        /// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
        /// Each entry can have one of the following values:
        /// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
        /// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
        /// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
        /// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
        /// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
        /// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
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
