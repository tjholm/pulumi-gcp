// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Spanner
{
    /// <summary>
    /// Three different resources help you manage your IAM policy for a Spanner database. Each of these resources serves a different use case:
    /// 
    /// * `gcp.spanner.DatabaseIAMPolicy`: Authoritative. Sets the IAM policy for the database and replaces any existing policy already attached.
    /// 
    /// &gt; **Warning:** It's entirely possibly to lock yourself out of your database using `gcp.spanner.DatabaseIAMPolicy`. Any permissions granted by default will be removed unless you include them in your config.
    /// 
    /// * `gcp.spanner.DatabaseIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the database are preserved.
    /// * `gcp.spanner.DatabaseIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the database are preserved.
    /// 
    /// &gt; **Note:** `gcp.spanner.DatabaseIAMPolicy` **cannot** be used in conjunction with `gcp.spanner.DatabaseIAMBinding` and `gcp.spanner.DatabaseIAMMember` or they will fight over what your policy should be.
    /// 
    /// &gt; **Note:** `gcp.spanner.DatabaseIAMBinding` resources **can be** used in conjunction with `gcp.spanner.DatabaseIAMMember` resources **only if** they do not grant privilege to the same role.
    /// 
    /// ## google\_spanner\_database\_iam\_policy
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
    ///     var database = new Gcp.Spanner.DatabaseIAMPolicy("database", new()
    ///     {
    ///         Instance = "your-instance-name",
    ///         Database = "your-database-name",
    ///         PolicyData = admin.Apply(getIAMPolicyResult =&gt; getIAMPolicyResult.PolicyData),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## google\_spanner\_database\_iam\_binding
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var database = new Gcp.Spanner.DatabaseIAMBinding("database", new()
    ///     {
    ///         Database = "your-database-name",
    ///         Instance = "your-instance-name",
    ///         Members = new[]
    ///         {
    ///             "user:jane@example.com",
    ///         },
    ///         Role = "roles/compute.networkUser",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## google\_spanner\_database\_iam\_member
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var database = new Gcp.Spanner.DatabaseIAMMember("database", new()
    ///     {
    ///         Database = "your-database-name",
    ///         Instance = "your-instance-name",
    ///         Member = "user:jane@example.com",
    ///         Role = "roles/compute.networkUser",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// For all import syntaxes, the "resource in question" can take any of the following forms* {{project}}/{{instance}}/{{database}} * {{instance}}/{{database}} (project is taken from provider project) IAM member imports use space-delimited identifiers; the resource in question, the role, and the member identity, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:spanner/databaseIAMMember:DatabaseIAMMember database "project-name/instance-name/database-name roles/viewer user:foo@example.com"
    /// ```
    /// 
    ///  IAM binding imports use space-delimited identifiers; the resource in question and the role, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:spanner/databaseIAMMember:DatabaseIAMMember database "project-name/instance-name/database-name roles/viewer"
    /// ```
    /// 
    ///  IAM policy imports use the identifier of the resource in question, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:spanner/databaseIAMMember:DatabaseIAMMember database project-name/instance-name/database-name
    /// ```
    /// 
    ///  -&gt; **Custom Roles:** If you're importing a IAM resource with a custom role, make sure to use the
    /// 
    /// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
    /// </summary>
    [GcpResourceType("gcp:spanner/databaseIAMMember:DatabaseIAMMember")]
    public partial class DatabaseIAMMember : global::Pulumi.CustomResource
    {
        [Output("condition")]
        public Output<Outputs.DatabaseIAMMemberCondition?> Condition { get; private set; } = null!;

        /// <summary>
        /// The name of the Spanner database.
        /// </summary>
        [Output("database")]
        public Output<string> Database { get; private set; } = null!;

        /// <summary>
        /// (Computed) The etag of the database's IAM policy.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The name of the Spanner instance the database belongs to.
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
        [Output("instance")]
        public Output<string> Instance { get; private set; } = null!;

        [Output("member")]
        public Output<string> Member { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.spanner.DatabaseIAMBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;


        /// <summary>
        /// Create a DatabaseIAMMember resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DatabaseIAMMember(string name, DatabaseIAMMemberArgs args, CustomResourceOptions? options = null)
            : base("gcp:spanner/databaseIAMMember:DatabaseIAMMember", name, args ?? new DatabaseIAMMemberArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DatabaseIAMMember(string name, Input<string> id, DatabaseIAMMemberState? state = null, CustomResourceOptions? options = null)
            : base("gcp:spanner/databaseIAMMember:DatabaseIAMMember", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DatabaseIAMMember resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DatabaseIAMMember Get(string name, Input<string> id, DatabaseIAMMemberState? state = null, CustomResourceOptions? options = null)
        {
            return new DatabaseIAMMember(name, id, state, options);
        }
    }

    public sealed class DatabaseIAMMemberArgs : global::Pulumi.ResourceArgs
    {
        [Input("condition")]
        public Input<Inputs.DatabaseIAMMemberConditionArgs>? Condition { get; set; }

        /// <summary>
        /// The name of the Spanner database.
        /// </summary>
        [Input("database", required: true)]
        public Input<string> Database { get; set; } = null!;

        /// <summary>
        /// The name of the Spanner instance the database belongs to.
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
        [Input("instance", required: true)]
        public Input<string> Instance { get; set; } = null!;

        [Input("member", required: true)]
        public Input<string> Member { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.spanner.DatabaseIAMBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        public DatabaseIAMMemberArgs()
        {
        }
        public static new DatabaseIAMMemberArgs Empty => new DatabaseIAMMemberArgs();
    }

    public sealed class DatabaseIAMMemberState : global::Pulumi.ResourceArgs
    {
        [Input("condition")]
        public Input<Inputs.DatabaseIAMMemberConditionGetArgs>? Condition { get; set; }

        /// <summary>
        /// The name of the Spanner database.
        /// </summary>
        [Input("database")]
        public Input<string>? Database { get; set; }

        /// <summary>
        /// (Computed) The etag of the database's IAM policy.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The name of the Spanner instance the database belongs to.
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
        [Input("instance")]
        public Input<string>? Instance { get; set; }

        [Input("member")]
        public Input<string>? Member { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.spanner.DatabaseIAMBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        public DatabaseIAMMemberState()
        {
        }
        public static new DatabaseIAMMemberState Empty => new DatabaseIAMMemberState();
    }
}
