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
    /// Three different resources help you manage your IAM policy for Healthcare FHIR store. Each of these resources serves a different use case:
    /// 
    /// * `gcp.healthcare.FhirStoreIamPolicy`: Authoritative. Sets the IAM policy for the FHIR store and replaces any existing policy already attached.
    /// * `gcp.healthcare.FhirStoreIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the FHIR store are preserved.
    /// * `gcp.healthcare.FhirStoreIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the FHIR store are preserved.
    /// 
    /// &gt; **Note:** `gcp.healthcare.FhirStoreIamPolicy` **cannot** be used in conjunction with `gcp.healthcare.FhirStoreIamBinding` and `gcp.healthcare.FhirStoreIamMember` or they will fight over what your policy should be.
    /// 
    /// &gt; **Note:** `gcp.healthcare.FhirStoreIamBinding` resources **can be** used in conjunction with `gcp.healthcare.FhirStoreIamMember` resources **only if** they do not grant privilege to the same role.
    /// 
    /// ## google\_healthcare\_fhir\_store\_iam\_policy
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
    ///     var fhirStore = new Gcp.Healthcare.FhirStoreIamPolicy("fhirStore", new()
    ///     {
    ///         FhirStoreId = "your-fhir-store-id",
    ///         PolicyData = admin.Apply(getIAMPolicyResult =&gt; getIAMPolicyResult.PolicyData),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## google\_healthcare\_fhir\_store\_iam\_binding
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var fhirStore = new Gcp.Healthcare.FhirStoreIamBinding("fhirStore", new()
    ///     {
    ///         FhirStoreId = "your-fhir-store-id",
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
    /// ## google\_healthcare\_fhir\_store\_iam\_member
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var fhirStore = new Gcp.Healthcare.FhirStoreIamMember("fhirStore", new()
    ///     {
    ///         FhirStoreId = "your-fhir-store-id",
    ///         Member = "user:jane@example.com",
    ///         Role = "roles/editor",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// IAM member imports use space-delimited identifiers; the resource in question, the role, and the account.
    /// 
    /// This member resource can be imported using the `fhir_store_id`, role, and account e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:healthcare/fhirStoreIamPolicy:FhirStoreIamPolicy fhir_store_iam "your-project-id/location-name/dataset-name/fhir-store-name roles/viewer user:foo@example.com"
    /// ```
    /// 
    ///  IAM binding imports use space-delimited identifiers; the resource in question and the role.
    /// 
    /// This binding resource can be imported using the `fhir_store_id` and role, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:healthcare/fhirStoreIamPolicy:FhirStoreIamPolicy fhir_store_iam "your-project-id/location-name/dataset-name/fhir-store-name roles/viewer"
    /// ```
    /// 
    ///  IAM policy imports use the identifier of the resource in question.
    /// 
    /// This policy resource can be imported using the `fhir_store_id`, role, and account e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:healthcare/fhirStoreIamPolicy:FhirStoreIamPolicy fhir_store_iam your-project-id/location-name/dataset-name/fhir-store-name
    /// ```
    /// </summary>
    [GcpResourceType("gcp:healthcare/fhirStoreIamPolicy:FhirStoreIamPolicy")]
    public partial class FhirStoreIamPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// (Computed) The etag of the FHIR store's IAM policy.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The FHIR store ID, in the form
        /// `{project_id}/{location_name}/{dataset_name}/{fhir_store_name}` or
        /// `{location_name}/{dataset_name}/{fhir_store_name}`. In the second form, the provider's
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
        [Output("fhirStoreId")]
        public Output<string> FhirStoreId { get; private set; } = null!;

        /// <summary>
        /// The policy data generated by
        /// a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        [Output("policyData")]
        public Output<string> PolicyData { get; private set; } = null!;


        /// <summary>
        /// Create a FhirStoreIamPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FhirStoreIamPolicy(string name, FhirStoreIamPolicyArgs args, CustomResourceOptions? options = null)
            : base("gcp:healthcare/fhirStoreIamPolicy:FhirStoreIamPolicy", name, args ?? new FhirStoreIamPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FhirStoreIamPolicy(string name, Input<string> id, FhirStoreIamPolicyState? state = null, CustomResourceOptions? options = null)
            : base("gcp:healthcare/fhirStoreIamPolicy:FhirStoreIamPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FhirStoreIamPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FhirStoreIamPolicy Get(string name, Input<string> id, FhirStoreIamPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new FhirStoreIamPolicy(name, id, state, options);
        }
    }

    public sealed class FhirStoreIamPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The FHIR store ID, in the form
        /// `{project_id}/{location_name}/{dataset_name}/{fhir_store_name}` or
        /// `{location_name}/{dataset_name}/{fhir_store_name}`. In the second form, the provider's
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
        [Input("fhirStoreId", required: true)]
        public Input<string> FhirStoreId { get; set; } = null!;

        /// <summary>
        /// The policy data generated by
        /// a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        [Input("policyData", required: true)]
        public Input<string> PolicyData { get; set; } = null!;

        public FhirStoreIamPolicyArgs()
        {
        }
        public static new FhirStoreIamPolicyArgs Empty => new FhirStoreIamPolicyArgs();
    }

    public sealed class FhirStoreIamPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Computed) The etag of the FHIR store's IAM policy.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The FHIR store ID, in the form
        /// `{project_id}/{location_name}/{dataset_name}/{fhir_store_name}` or
        /// `{location_name}/{dataset_name}/{fhir_store_name}`. In the second form, the provider's
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
        [Input("fhirStoreId")]
        public Input<string>? FhirStoreId { get; set; }

        /// <summary>
        /// The policy data generated by
        /// a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        [Input("policyData")]
        public Input<string>? PolicyData { get; set; }

        public FhirStoreIamPolicyState()
        {
        }
        public static new FhirStoreIamPolicyState Empty => new FhirStoreIamPolicyState();
    }
}
