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
    /// Three different resources help you manage your IAM policy for Healthcare HL7v2 store. Each of these resources serves a different use case:
    /// 
    /// * `gcp.healthcare.Hl7StoreIamPolicy`: Authoritative. Sets the IAM policy for the HL7v2 store and replaces any existing policy already attached.
    /// * `gcp.healthcare.Hl7StoreIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the HL7v2 store are preserved.
    /// * `gcp.healthcare.Hl7StoreIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the HL7v2 store are preserved.
    /// 
    /// &gt; **Note:** `gcp.healthcare.Hl7StoreIamPolicy` **cannot** be used in conjunction with `gcp.healthcare.Hl7StoreIamBinding` and `gcp.healthcare.Hl7StoreIamMember` or they will fight over what your policy should be.
    /// 
    /// &gt; **Note:** `gcp.healthcare.Hl7StoreIamBinding` resources **can be** used in conjunction with `gcp.healthcare.Hl7StoreIamMember` resources **only if** they do not grant privilege to the same role.
    /// 
    /// ## gcp.healthcare.Hl7StoreIamPolicy
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
    ///     var hl7V2Store = new Gcp.Healthcare.Hl7StoreIamPolicy("hl7_v2_store", new()
    ///     {
    ///         Hl7V2StoreId = "your-hl7-v2-store-id",
    ///         PolicyData = admin.Apply(getIAMPolicyResult =&gt; getIAMPolicyResult.PolicyData),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## gcp.healthcare.Hl7StoreIamBinding
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var hl7V2Store = new Gcp.Healthcare.Hl7StoreIamBinding("hl7_v2_store", new()
    ///     {
    ///         Hl7V2StoreId = "your-hl7-v2-store-id",
    ///         Role = "roles/editor",
    ///         Members = new[]
    ///         {
    ///             "user:jane@example.com",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## gcp.healthcare.Hl7StoreIamMember
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var hl7V2Store = new Gcp.Healthcare.Hl7StoreIamMember("hl7_v2_store", new()
    ///     {
    ///         Hl7V2StoreId = "your-hl7-v2-store-id",
    ///         Role = "roles/editor",
    ///         Member = "user:jane@example.com",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## gcp.healthcare.Hl7StoreIamBinding
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var hl7V2Store = new Gcp.Healthcare.Hl7StoreIamBinding("hl7_v2_store", new()
    ///     {
    ///         Hl7V2StoreId = "your-hl7-v2-store-id",
    ///         Role = "roles/editor",
    ///         Members = new[]
    ///         {
    ///             "user:jane@example.com",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## gcp.healthcare.Hl7StoreIamMember
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var hl7V2Store = new Gcp.Healthcare.Hl7StoreIamMember("hl7_v2_store", new()
    ///     {
    ///         Hl7V2StoreId = "your-hl7-v2-store-id",
    ///         Role = "roles/editor",
    ///         Member = "user:jane@example.com",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ### Importing IAM policies
    /// 
    /// IAM policy imports use the identifier of the Google Cloud Healthcare HL7v2 store resource. For example:
    /// 
    /// * `"{{project_id}}/{{location}}/{{dataset}}/{{hl7_v2_store}}"`
    /// 
    /// An `import` block (Terraform v1.5.0 and later) can be used to import IAM policies:
    /// 
    /// tf
    /// 
    /// import {
    /// 
    ///   id = "{{project_id}}/{{location}}/{{dataset}}/{{hl7_v2_store}}"
    /// 
    ///   to = google_healthcare_hl7_v2_store_iam_policy.default
    /// 
    /// }
    /// 
    /// The `pulumi import` command can also be used:
    /// 
    /// ```sh
    /// $ pulumi import gcp:healthcare/hl7StoreIamPolicy:Hl7StoreIamPolicy default {{project_id}}/{{location}}/{{dataset}}/{{hl7_v2_store}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:healthcare/hl7StoreIamPolicy:Hl7StoreIamPolicy")]
    public partial class Hl7StoreIamPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// (Computed) The etag of the HL7v2 store's IAM policy.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The HL7v2 store ID, in the form
        /// `{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}` or
        /// `{location_name}/{dataset_name}/{hl7_v2_store_name}`. In the second form, the provider's
        /// project setting will be used as a fallback.
        /// </summary>
        [Output("hl7V2StoreId")]
        public Output<string> Hl7V2StoreId { get; private set; } = null!;

        /// <summary>
        /// The policy data generated by
        /// a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        [Output("policyData")]
        public Output<string> PolicyData { get; private set; } = null!;


        /// <summary>
        /// Create a Hl7StoreIamPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Hl7StoreIamPolicy(string name, Hl7StoreIamPolicyArgs args, CustomResourceOptions? options = null)
            : base("gcp:healthcare/hl7StoreIamPolicy:Hl7StoreIamPolicy", name, args ?? new Hl7StoreIamPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Hl7StoreIamPolicy(string name, Input<string> id, Hl7StoreIamPolicyState? state = null, CustomResourceOptions? options = null)
            : base("gcp:healthcare/hl7StoreIamPolicy:Hl7StoreIamPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Hl7StoreIamPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Hl7StoreIamPolicy Get(string name, Input<string> id, Hl7StoreIamPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new Hl7StoreIamPolicy(name, id, state, options);
        }
    }

    public sealed class Hl7StoreIamPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The HL7v2 store ID, in the form
        /// `{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}` or
        /// `{location_name}/{dataset_name}/{hl7_v2_store_name}`. In the second form, the provider's
        /// project setting will be used as a fallback.
        /// </summary>
        [Input("hl7V2StoreId", required: true)]
        public Input<string> Hl7V2StoreId { get; set; } = null!;

        /// <summary>
        /// The policy data generated by
        /// a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        [Input("policyData", required: true)]
        public Input<string> PolicyData { get; set; } = null!;

        public Hl7StoreIamPolicyArgs()
        {
        }
        public static new Hl7StoreIamPolicyArgs Empty => new Hl7StoreIamPolicyArgs();
    }

    public sealed class Hl7StoreIamPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Computed) The etag of the HL7v2 store's IAM policy.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The HL7v2 store ID, in the form
        /// `{project_id}/{location_name}/{dataset_name}/{hl7_v2_store_name}` or
        /// `{location_name}/{dataset_name}/{hl7_v2_store_name}`. In the second form, the provider's
        /// project setting will be used as a fallback.
        /// </summary>
        [Input("hl7V2StoreId")]
        public Input<string>? Hl7V2StoreId { get; set; }

        /// <summary>
        /// The policy data generated by
        /// a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        [Input("policyData")]
        public Input<string>? PolicyData { get; set; }

        public Hl7StoreIamPolicyState()
        {
        }
        public static new Hl7StoreIamPolicyState Empty => new Hl7StoreIamPolicyState();
    }
}
