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
    /// ## gcp.spanner.DatabaseIAMPolicy
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
    /// With IAM Conditions:
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
    ///                 Condition = new Gcp.Organizations.Inputs.GetIAMPolicyBindingConditionInputArgs
    ///                 {
    ///                     Title = "My Role",
    ///                     Description = "Grant permissions on my_role",
    ///                     Expression = "(resource.type == \"spanner.googleapis.com/DatabaseRole\" &amp;&amp; (resource.name.endsWith(\"/myrole\")))",
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
    /// ## gcp.spanner.DatabaseIAMBinding
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
    ///         Instance = "your-instance-name",
    ///         Database = "your-database-name",
    ///         Role = "roles/compute.networkUser",
    ///         Members = new[]
    ///         {
    ///             "user:jane@example.com",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// With IAM Conditions:
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
    ///         Instance = "your-instance-name",
    ///         Database = "your-database-name",
    ///         Role = "roles/compute.networkUser",
    ///         Members = new[]
    ///         {
    ///             "user:jane@example.com",
    ///         },
    ///         Condition = new Gcp.Spanner.Inputs.DatabaseIAMBindingConditionArgs
    ///         {
    ///             Title = "My Role",
    ///             Description = "Grant permissions on my_role",
    ///             Expression = "(resource.type == \"spanner.googleapis.com/DatabaseRole\" &amp;&amp; (resource.name.endsWith(\"/myrole\")))",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## gcp.spanner.DatabaseIAMMember
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
    ///         Instance = "your-instance-name",
    ///         Database = "your-database-name",
    ///         Role = "roles/compute.networkUser",
    ///         Member = "user:jane@example.com",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// With IAM Conditions:
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
    ///         Instance = "your-instance-name",
    ///         Database = "your-database-name",
    ///         Role = "roles/compute.networkUser",
    ///         Member = "user:jane@example.com",
    ///         Condition = new Gcp.Spanner.Inputs.DatabaseIAMMemberConditionArgs
    ///         {
    ///             Title = "My Role",
    ///             Description = "Grant permissions on my_role",
    ///             Expression = "(resource.type == \"spanner.googleapis.com/DatabaseRole\" &amp;&amp; (resource.name.endsWith(\"/myrole\")))",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## gcp.spanner.DatabaseIAMBinding
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
    ///         Instance = "your-instance-name",
    ///         Database = "your-database-name",
    ///         Role = "roles/compute.networkUser",
    ///         Members = new[]
    ///         {
    ///             "user:jane@example.com",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// With IAM Conditions:
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
    ///         Instance = "your-instance-name",
    ///         Database = "your-database-name",
    ///         Role = "roles/compute.networkUser",
    ///         Members = new[]
    ///         {
    ///             "user:jane@example.com",
    ///         },
    ///         Condition = new Gcp.Spanner.Inputs.DatabaseIAMBindingConditionArgs
    ///         {
    ///             Title = "My Role",
    ///             Description = "Grant permissions on my_role",
    ///             Expression = "(resource.type == \"spanner.googleapis.com/DatabaseRole\" &amp;&amp; (resource.name.endsWith(\"/myrole\")))",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## gcp.spanner.DatabaseIAMMember
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
    ///         Instance = "your-instance-name",
    ///         Database = "your-database-name",
    ///         Role = "roles/compute.networkUser",
    ///         Member = "user:jane@example.com",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// With IAM Conditions:
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
    ///         Instance = "your-instance-name",
    ///         Database = "your-database-name",
    ///         Role = "roles/compute.networkUser",
    ///         Member = "user:jane@example.com",
    ///         Condition = new Gcp.Spanner.Inputs.DatabaseIAMMemberConditionArgs
    ///         {
    ///             Title = "My Role",
    ///             Description = "Grant permissions on my_role",
    ///             Expression = "(resource.type == \"spanner.googleapis.com/DatabaseRole\" &amp;&amp; (resource.name.endsWith(\"/myrole\")))",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ### Importing IAM policies
    /// 
    /// IAM policy imports use the identifier of the Spanner Database resource in question. For example:
    /// 
    /// * `{{project}}/{{instance}}/{{database}}`
    /// 
    /// An `import` block (Terraform v1.5.0 and later) can be used to import IAM policies:
    /// 
    /// tf
    /// 
    /// import {
    /// 
    ///   id = {{project}}/{{instance}}/{{database}}
    /// 
    ///   to = google_spanner_database_iam_policy.default
    /// 
    /// }
    /// 
    /// The `pulumi import` command can also be used:
    /// 
    /// ```sh
    /// $ pulumi import gcp:spanner/databaseIAMBinding:DatabaseIAMBinding default {{project}}/{{instance}}/{{database}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:spanner/databaseIAMBinding:DatabaseIAMBinding")]
    public partial class DatabaseIAMBinding : global::Pulumi.CustomResource
    {
        /// <summary>
        /// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
        /// Structure is documented below.
        /// </summary>
        [Output("condition")]
        public Output<Outputs.DatabaseIAMBindingCondition?> Condition { get; private set; } = null!;

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
        /// </summary>
        [Output("instance")]
        public Output<string> Instance { get; private set; } = null!;

        /// <summary>
        /// Identities that will be granted the privilege in `role`.
        /// Each entry can have one of the following values:
        /// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
        /// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
        /// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
        /// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
        /// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
        /// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
        /// </summary>
        [Output("members")]
        public Output<ImmutableArray<string>> Members { get; private set; } = null!;

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
        /// Create a DatabaseIAMBinding resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DatabaseIAMBinding(string name, DatabaseIAMBindingArgs args, CustomResourceOptions? options = null)
            : base("gcp:spanner/databaseIAMBinding:DatabaseIAMBinding", name, args ?? new DatabaseIAMBindingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DatabaseIAMBinding(string name, Input<string> id, DatabaseIAMBindingState? state = null, CustomResourceOptions? options = null)
            : base("gcp:spanner/databaseIAMBinding:DatabaseIAMBinding", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DatabaseIAMBinding resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DatabaseIAMBinding Get(string name, Input<string> id, DatabaseIAMBindingState? state = null, CustomResourceOptions? options = null)
        {
            return new DatabaseIAMBinding(name, id, state, options);
        }
    }

    public sealed class DatabaseIAMBindingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
        /// Structure is documented below.
        /// </summary>
        [Input("condition")]
        public Input<Inputs.DatabaseIAMBindingConditionArgs>? Condition { get; set; }

        /// <summary>
        /// The name of the Spanner database.
        /// </summary>
        [Input("database", required: true)]
        public Input<string> Database { get; set; } = null!;

        /// <summary>
        /// The name of the Spanner instance the database belongs to.
        /// </summary>
        [Input("instance", required: true)]
        public Input<string> Instance { get; set; } = null!;

        [Input("members", required: true)]
        private InputList<string>? _members;

        /// <summary>
        /// Identities that will be granted the privilege in `role`.
        /// Each entry can have one of the following values:
        /// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
        /// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
        /// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
        /// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
        /// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
        /// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
        /// </summary>
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

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

        public DatabaseIAMBindingArgs()
        {
        }
        public static new DatabaseIAMBindingArgs Empty => new DatabaseIAMBindingArgs();
    }

    public sealed class DatabaseIAMBindingState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
        /// Structure is documented below.
        /// </summary>
        [Input("condition")]
        public Input<Inputs.DatabaseIAMBindingConditionGetArgs>? Condition { get; set; }

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
        /// </summary>
        [Input("instance")]
        public Input<string>? Instance { get; set; }

        [Input("members")]
        private InputList<string>? _members;

        /// <summary>
        /// Identities that will be granted the privilege in `role`.
        /// Each entry can have one of the following values:
        /// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
        /// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
        /// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
        /// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
        /// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
        /// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
        /// </summary>
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

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

        public DatabaseIAMBindingState()
        {
        }
        public static new DatabaseIAMBindingState Empty => new DatabaseIAMBindingState();
    }
}
