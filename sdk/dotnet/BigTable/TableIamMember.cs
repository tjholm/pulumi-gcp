// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.BigTable
{
    /// <summary>
    /// Three different resources help you manage IAM policies on bigtable tables. Each of these resources serves a different use case:
    /// 
    /// * `gcp.bigtable.TableIamPolicy`: Authoritative. Sets the IAM policy for the tables and replaces any existing policy already attached.
    /// * `gcp.bigtable.TableIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the table are preserved.
    /// * `gcp.bigtable.TableIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the table are preserved.
    /// 
    /// &gt; **Note:** `gcp.bigtable.TableIamPolicy` **cannot** be used in conjunction with `gcp.bigtable.TableIamBinding` and `gcp.bigtable.TableIamMember` or they will fight over what your policy should be. In addition, be careful not to accidentally unset ownership of the table as `gcp.bigtable.TableIamPolicy` replaces the entire policy.
    /// 
    /// &gt; **Note:** `gcp.bigtable.TableIamBinding` resources **can be** used in conjunction with `gcp.bigtable.TableIamMember` resources **only if** they do not grant privilege to the same role.
    /// 
    /// ## google\_bigtable\_table\_iam\_policy
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
    ///                 Role = "roles/bigtable.user",
    ///                 Members = new[]
    ///                 {
    ///                     "user:jane@example.com",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var editor = new Gcp.BigTable.TableIamPolicy("editor", new()
    ///     {
    ///         Project = "your-project",
    ///         Instance = "your-bigtable-instance",
    ///         Table = "your-bigtable-table",
    ///         PolicyData = admin.Apply(getIAMPolicyResult =&gt; getIAMPolicyResult.PolicyData),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## google\_bigtable\_table\_iam\_binding
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var editor = new Gcp.BigTable.TableIamBinding("editor", new()
    ///     {
    ///         Instance = "your-bigtable-instance",
    ///         Members = new[]
    ///         {
    ///             "user:jane@example.com",
    ///         },
    ///         Role = "roles/bigtable.user",
    ///         Table = "your-bigtable-table",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## google\_bigtable\_table\_iam\_member
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var editor = new Gcp.BigTable.TableIamMember("editor", new()
    ///     {
    ///         Instance = "your-bigtable-instance",
    ///         Member = "user:jane@example.com",
    ///         Role = "roles/bigtable.user",
    ///         Table = "your-bigtable-table",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ### Importing IAM policies IAM policy imports use the `table` identifier of the Bigtable Table resource only. For example* `"projects/{project}/tables/{table}"` An `import` block (Terraform v1.5.0 and later) can be used to import IAM policiestf import {
    /// 
    ///  id = "projects/{project}/tables/{table}"
    /// 
    ///  to = google_bigtable_table_iam_policy.default } The `pulumi import` command can also be used
    /// 
    /// ```sh
    ///  $ pulumi import gcp:bigtable/tableIamMember:TableIamMember default projects/{project}/tables/{table}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:bigtable/tableIamMember:TableIamMember")]
    public partial class TableIamMember : global::Pulumi.CustomResource
    {
        [Output("condition")]
        public Output<Outputs.TableIamMemberCondition?> Condition { get; private set; } = null!;

        /// <summary>
        /// (Computed) The etag of the tables's IAM policy.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The name or relative resource id of the instance that owns the table.
        /// </summary>
        [Output("instance")]
        public Output<string> Instance { get; private set; } = null!;

        [Output("member")]
        public Output<string> Member { get; private set; } = null!;

        /// <summary>
        /// The project in which the table belongs. If it
        /// is not provided, this provider will use the provider default.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.bigtable.TableIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
        /// 
        /// `gcp.bigtable.TableIamPolicy` only:
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;

        /// <summary>
        /// The name or relative resource id of the table to manage IAM policies for.
        /// 
        /// For `gcp.bigtable.TableIamMember` or `gcp.bigtable.TableIamBinding`:
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
        [Output("table")]
        public Output<string> Table { get; private set; } = null!;


        /// <summary>
        /// Create a TableIamMember resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TableIamMember(string name, TableIamMemberArgs args, CustomResourceOptions? options = null)
            : base("gcp:bigtable/tableIamMember:TableIamMember", name, args ?? new TableIamMemberArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TableIamMember(string name, Input<string> id, TableIamMemberState? state = null, CustomResourceOptions? options = null)
            : base("gcp:bigtable/tableIamMember:TableIamMember", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TableIamMember resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TableIamMember Get(string name, Input<string> id, TableIamMemberState? state = null, CustomResourceOptions? options = null)
        {
            return new TableIamMember(name, id, state, options);
        }
    }

    public sealed class TableIamMemberArgs : global::Pulumi.ResourceArgs
    {
        [Input("condition")]
        public Input<Inputs.TableIamMemberConditionArgs>? Condition { get; set; }

        /// <summary>
        /// The name or relative resource id of the instance that owns the table.
        /// </summary>
        [Input("instance", required: true)]
        public Input<string> Instance { get; set; } = null!;

        [Input("member", required: true)]
        public Input<string> Member { get; set; } = null!;

        /// <summary>
        /// The project in which the table belongs. If it
        /// is not provided, this provider will use the provider default.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.bigtable.TableIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
        /// 
        /// `gcp.bigtable.TableIamPolicy` only:
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        /// <summary>
        /// The name or relative resource id of the table to manage IAM policies for.
        /// 
        /// For `gcp.bigtable.TableIamMember` or `gcp.bigtable.TableIamBinding`:
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
        [Input("table", required: true)]
        public Input<string> Table { get; set; } = null!;

        public TableIamMemberArgs()
        {
        }
        public static new TableIamMemberArgs Empty => new TableIamMemberArgs();
    }

    public sealed class TableIamMemberState : global::Pulumi.ResourceArgs
    {
        [Input("condition")]
        public Input<Inputs.TableIamMemberConditionGetArgs>? Condition { get; set; }

        /// <summary>
        /// (Computed) The etag of the tables's IAM policy.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The name or relative resource id of the instance that owns the table.
        /// </summary>
        [Input("instance")]
        public Input<string>? Instance { get; set; }

        [Input("member")]
        public Input<string>? Member { get; set; }

        /// <summary>
        /// The project in which the table belongs. If it
        /// is not provided, this provider will use the provider default.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.bigtable.TableIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`. Read more about roles [here](https://cloud.google.com/bigtable/docs/access-control#roles).
        /// 
        /// `gcp.bigtable.TableIamPolicy` only:
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// The name or relative resource id of the table to manage IAM policies for.
        /// 
        /// For `gcp.bigtable.TableIamMember` or `gcp.bigtable.TableIamBinding`:
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
        [Input("table")]
        public Input<string>? Table { get; set; }

        public TableIamMemberState()
        {
        }
        public static new TableIamMemberState Empty => new TableIamMemberState();
    }
}
