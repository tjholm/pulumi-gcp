// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.BigQueryDataPolicy
{
    /// <summary>
    /// Three different resources help you manage your IAM policy for BigQuery Data Policy DataPolicy. Each of these resources serves a different use case:
    /// 
    /// * `gcp.bigquerydatapolicy.DataPolicyIamPolicy`: Authoritative. Sets the IAM policy for the datapolicy and replaces any existing policy already attached.
    /// * `gcp.bigquerydatapolicy.DataPolicyIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the datapolicy are preserved.
    /// * `gcp.bigquerydatapolicy.DataPolicyIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the datapolicy are preserved.
    /// 
    /// A data source can be used to retrieve policy data in advent you do not need creation
    /// 
    /// * `gcp.bigquerydatapolicy.DataPolicyIamPolicy`: Retrieves the IAM policy for the datapolicy
    /// 
    /// &gt; **Note:** `gcp.bigquerydatapolicy.DataPolicyIamPolicy` **cannot** be used in conjunction with `gcp.bigquerydatapolicy.DataPolicyIamBinding` and `gcp.bigquerydatapolicy.DataPolicyIamMember` or they will fight over what your policy should be.
    /// 
    /// &gt; **Note:** `gcp.bigquerydatapolicy.DataPolicyIamBinding` resources **can be** used in conjunction with `gcp.bigquerydatapolicy.DataPolicyIamMember` resources **only if** they do not grant privilege to the same role.
    /// 
    /// ## google\_bigquery\_datapolicy\_data\_policy\_iam\_policy
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
    ///                 Role = "roles/viewer",
    ///                 Members = new[]
    ///                 {
    ///                     "user:jane@example.com",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var policy = new Gcp.BigQueryDataPolicy.DataPolicyIamPolicy("policy", new()
    ///     {
    ///         Project = google_bigquery_datapolicy_data_policy.Data_policy.Project,
    ///         Location = google_bigquery_datapolicy_data_policy.Data_policy.Location,
    ///         DataPolicyId = google_bigquery_datapolicy_data_policy.Data_policy.Data_policy_id,
    ///         PolicyData = admin.Apply(getIAMPolicyResult =&gt; getIAMPolicyResult.PolicyData),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## google\_bigquery\_datapolicy\_data\_policy\_iam\_binding
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var binding = new Gcp.BigQueryDataPolicy.DataPolicyIamBinding("binding", new()
    ///     {
    ///         Project = google_bigquery_datapolicy_data_policy.Data_policy.Project,
    ///         Location = google_bigquery_datapolicy_data_policy.Data_policy.Location,
    ///         DataPolicyId = google_bigquery_datapolicy_data_policy.Data_policy.Data_policy_id,
    ///         Role = "roles/viewer",
    ///         Members = new[]
    ///         {
    ///             "user:jane@example.com",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## google\_bigquery\_datapolicy\_data\_policy\_iam\_member
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var member = new Gcp.BigQueryDataPolicy.DataPolicyIamMember("member", new()
    ///     {
    ///         Project = google_bigquery_datapolicy_data_policy.Data_policy.Project,
    ///         Location = google_bigquery_datapolicy_data_policy.Data_policy.Location,
    ///         DataPolicyId = google_bigquery_datapolicy_data_policy.Data_policy.Data_policy_id,
    ///         Role = "roles/viewer",
    ///         Member = "user:jane@example.com",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{location}}/dataPolicies/{{data_policy_id}} * {{project}}/{{location}}/{{data_policy_id}} * {{location}}/{{data_policy_id}} * {{data_policy_id}} Any variables not passed in the import command will be taken from the provider configuration. BigQuery Data Policy datapolicy IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:bigquerydatapolicy/dataPolicyIamBinding:DataPolicyIamBinding editor "projects/{{project}}/locations/{{location}}/dataPolicies/{{data_policy_id}} roles/viewer user:jane@example.com"
    /// ```
    /// 
    ///  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:bigquerydatapolicy/dataPolicyIamBinding:DataPolicyIamBinding editor "projects/{{project}}/locations/{{location}}/dataPolicies/{{data_policy_id}} roles/viewer"
    /// ```
    /// 
    ///  IAM policy imports use the identifier of the resource in question, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:bigquerydatapolicy/dataPolicyIamBinding:DataPolicyIamBinding editor projects/{{project}}/locations/{{location}}/dataPolicies/{{data_policy_id}}
    /// ```
    /// 
    ///  -&gt; **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
    /// 
    /// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
    /// </summary>
    [GcpResourceType("gcp:bigquerydatapolicy/dataPolicyIamBinding:DataPolicyIamBinding")]
    public partial class DataPolicyIamBinding : global::Pulumi.CustomResource
    {
        [Output("condition")]
        public Output<Outputs.DataPolicyIamBindingCondition?> Condition { get; private set; } = null!;

        [Output("dataPolicyId")]
        public Output<string> DataPolicyId { get; private set; } = null!;

        /// <summary>
        /// (Computed) The etag of the IAM policy.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The name of the location of the data policy.
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        [Output("members")]
        public Output<ImmutableArray<string>> Members { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        /// 
        /// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
        /// Each entry can have one of the following values:
        /// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
        /// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
        /// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
        /// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
        /// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
        /// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
        /// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
        /// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
        /// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.bigquerydatapolicy.DataPolicyIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;


        /// <summary>
        /// Create a DataPolicyIamBinding resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DataPolicyIamBinding(string name, DataPolicyIamBindingArgs args, CustomResourceOptions? options = null)
            : base("gcp:bigquerydatapolicy/dataPolicyIamBinding:DataPolicyIamBinding", name, args ?? new DataPolicyIamBindingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DataPolicyIamBinding(string name, Input<string> id, DataPolicyIamBindingState? state = null, CustomResourceOptions? options = null)
            : base("gcp:bigquerydatapolicy/dataPolicyIamBinding:DataPolicyIamBinding", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DataPolicyIamBinding resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DataPolicyIamBinding Get(string name, Input<string> id, DataPolicyIamBindingState? state = null, CustomResourceOptions? options = null)
        {
            return new DataPolicyIamBinding(name, id, state, options);
        }
    }

    public sealed class DataPolicyIamBindingArgs : global::Pulumi.ResourceArgs
    {
        [Input("condition")]
        public Input<Inputs.DataPolicyIamBindingConditionArgs>? Condition { get; set; }

        [Input("dataPolicyId", required: true)]
        public Input<string> DataPolicyId { get; set; } = null!;

        /// <summary>
        /// The name of the location of the data policy.
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("members", required: true)]
        private InputList<string>? _members;
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        /// 
        /// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
        /// Each entry can have one of the following values:
        /// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
        /// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
        /// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
        /// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
        /// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
        /// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
        /// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
        /// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
        /// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.bigquerydatapolicy.DataPolicyIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        public DataPolicyIamBindingArgs()
        {
        }
        public static new DataPolicyIamBindingArgs Empty => new DataPolicyIamBindingArgs();
    }

    public sealed class DataPolicyIamBindingState : global::Pulumi.ResourceArgs
    {
        [Input("condition")]
        public Input<Inputs.DataPolicyIamBindingConditionGetArgs>? Condition { get; set; }

        [Input("dataPolicyId")]
        public Input<string>? DataPolicyId { get; set; }

        /// <summary>
        /// (Computed) The etag of the IAM policy.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The name of the location of the data policy.
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("members")]
        private InputList<string>? _members;
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        /// 
        /// * `member/members` - (Required) Identities that will be granted the privilege in `role`.
        /// Each entry can have one of the following values:
        /// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
        /// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
        /// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
        /// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
        /// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
        /// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
        /// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
        /// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
        /// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.bigquerydatapolicy.DataPolicyIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        public DataPolicyIamBindingState()
        {
        }
        public static new DataPolicyIamBindingState Empty => new DataPolicyIamBindingState();
    }
}
