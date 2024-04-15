// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.ServiceAccount
{
    /// <summary>
    /// When managing IAM roles, you can treat a service account either as a resource or as an identity. This resource is to add iam policy bindings to a service account resource, such as allowing the members to run operations as or modify the service account. To configure permissions for a service account on other GCP resources, use the google_project_iam set of resources.
    /// 
    /// Three different resources help you manage your IAM policy for a service account. Each of these resources serves a different use case:
    /// 
    /// * `gcp.serviceaccount.IAMPolicy`: Authoritative. Sets the IAM policy for the service account and replaces any existing policy already attached.
    /// * `gcp.serviceaccount.IAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the service account are preserved.
    /// * `gcp.serviceaccount.IAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the service account are preserved.
    /// 
    /// &gt; **Note:** `gcp.serviceaccount.IAMPolicy` **cannot** be used in conjunction with `gcp.serviceaccount.IAMBinding` and `gcp.serviceaccount.IAMMember` or they will fight over what your policy should be.
    /// 
    /// &gt; **Note:** `gcp.serviceaccount.IAMBinding` resources **can be** used in conjunction with `gcp.serviceaccount.IAMMember` resources **only if** they do not grant privilege to the same role.
    /// 
    /// ## Example Usage
    /// 
    /// ### Service Account IAM Policy
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
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
    ///                 Role = "roles/iam.serviceAccountUser",
    ///                 Members = new[]
    ///                 {
    ///                     "user:jane@example.com",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var sa = new Gcp.ServiceAccount.Account("sa", new()
    ///     {
    ///         AccountId = "my-service-account",
    ///         DisplayName = "A service account that only Jane can interact with",
    ///     });
    /// 
    ///     var admin_account_iam = new Gcp.ServiceAccount.IAMPolicy("admin-account-iam", new()
    ///     {
    ///         ServiceAccountId = sa.Name,
    ///         PolicyData = admin.Apply(getIAMPolicyResult =&gt; getIAMPolicyResult.PolicyData),
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ### Service Account IAM Binding
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var sa = new Gcp.ServiceAccount.Account("sa", new()
    ///     {
    ///         AccountId = "my-service-account",
    ///         DisplayName = "A service account that only Jane can use",
    ///     });
    /// 
    ///     var admin_account_iam = new Gcp.ServiceAccount.IAMBinding("admin-account-iam", new()
    ///     {
    ///         ServiceAccountId = sa.Name,
    ///         Role = "roles/iam.serviceAccountUser",
    ///         Members = new[]
    ///         {
    ///             "user:jane@example.com",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ### Service Account IAM Binding With IAM Conditions:
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var sa = new Gcp.ServiceAccount.Account("sa", new()
    ///     {
    ///         AccountId = "my-service-account",
    ///         DisplayName = "A service account that only Jane can use",
    ///     });
    /// 
    ///     var admin_account_iam = new Gcp.ServiceAccount.IAMBinding("admin-account-iam", new()
    ///     {
    ///         ServiceAccountId = sa.Name,
    ///         Role = "roles/iam.serviceAccountUser",
    ///         Members = new[]
    ///         {
    ///             "user:jane@example.com",
    ///         },
    ///         Condition = new Gcp.ServiceAccount.Inputs.IAMBindingConditionArgs
    ///         {
    ///             Title = "expires_after_2019_12_31",
    ///             Description = "Expiring at midnight of 2019-12-31",
    ///             Expression = "request.time &lt; timestamp(\"2020-01-01T00:00:00Z\")",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ### Service Account IAM Member
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = Gcp.Compute.GetDefaultServiceAccount.Invoke();
    /// 
    ///     var sa = new Gcp.ServiceAccount.Account("sa", new()
    ///     {
    ///         AccountId = "my-service-account",
    ///         DisplayName = "A service account that Jane can use",
    ///     });
    /// 
    ///     var admin_account_iam = new Gcp.ServiceAccount.IAMMember("admin-account-iam", new()
    ///     {
    ///         ServiceAccountId = sa.Name,
    ///         Role = "roles/iam.serviceAccountUser",
    ///         Member = "user:jane@example.com",
    ///     });
    /// 
    ///     // Allow SA service account use the default GCE account
    ///     var gce_default_account_iam = new Gcp.ServiceAccount.IAMMember("gce-default-account-iam", new()
    ///     {
    ///         ServiceAccountId = @default.Apply(@default =&gt; @default.Apply(getDefaultServiceAccountResult =&gt; getDefaultServiceAccountResult.Name)),
    ///         Role = "roles/iam.serviceAccountUser",
    ///         Member = sa.Email.Apply(email =&gt; $"serviceAccount:{email}"),
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ### Service Account IAM Member With IAM Conditions:
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var sa = new Gcp.ServiceAccount.Account("sa", new()
    ///     {
    ///         AccountId = "my-service-account",
    ///         DisplayName = "A service account that Jane can use",
    ///     });
    /// 
    ///     var admin_account_iam = new Gcp.ServiceAccount.IAMMember("admin-account-iam", new()
    ///     {
    ///         ServiceAccountId = sa.Name,
    ///         Role = "roles/iam.serviceAccountUser",
    ///         Member = "user:jane@example.com",
    ///         Condition = new Gcp.ServiceAccount.Inputs.IAMMemberConditionArgs
    ///         {
    ///             Title = "expires_after_2019_12_31",
    ///             Description = "Expiring at midnight of 2019-12-31",
    ///             Expression = "request.time &lt; timestamp(\"2020-01-01T00:00:00Z\")",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ### Additional Examples
    /// 
    /// ### Service Account IAM Policy
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
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
    ///                 Role = "roles/iam.serviceAccountUser",
    ///                 Members = new[]
    ///                 {
    ///                     "user:jane@example.com",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var sa = new Gcp.ServiceAccount.Account("sa", new()
    ///     {
    ///         AccountId = "my-service-account",
    ///         DisplayName = "A service account that only Jane can interact with",
    ///     });
    /// 
    ///     var admin_account_iam = new Gcp.ServiceAccount.IAMPolicy("admin-account-iam", new()
    ///     {
    ///         ServiceAccountId = sa.Name,
    ///         PolicyData = admin.Apply(getIAMPolicyResult =&gt; getIAMPolicyResult.PolicyData),
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ### Service Account IAM Binding
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var sa = new Gcp.ServiceAccount.Account("sa", new()
    ///     {
    ///         AccountId = "my-service-account",
    ///         DisplayName = "A service account that only Jane can use",
    ///     });
    /// 
    ///     var admin_account_iam = new Gcp.ServiceAccount.IAMBinding("admin-account-iam", new()
    ///     {
    ///         ServiceAccountId = sa.Name,
    ///         Role = "roles/iam.serviceAccountUser",
    ///         Members = new[]
    ///         {
    ///             "user:jane@example.com",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ### Service Account IAM Binding With IAM Conditions:
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var sa = new Gcp.ServiceAccount.Account("sa", new()
    ///     {
    ///         AccountId = "my-service-account",
    ///         DisplayName = "A service account that only Jane can use",
    ///     });
    /// 
    ///     var admin_account_iam = new Gcp.ServiceAccount.IAMBinding("admin-account-iam", new()
    ///     {
    ///         ServiceAccountId = sa.Name,
    ///         Role = "roles/iam.serviceAccountUser",
    ///         Members = new[]
    ///         {
    ///             "user:jane@example.com",
    ///         },
    ///         Condition = new Gcp.ServiceAccount.Inputs.IAMBindingConditionArgs
    ///         {
    ///             Title = "expires_after_2019_12_31",
    ///             Description = "Expiring at midnight of 2019-12-31",
    ///             Expression = "request.time &lt; timestamp(\"2020-01-01T00:00:00Z\")",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ### Service Account IAM Member
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = Gcp.Compute.GetDefaultServiceAccount.Invoke();
    /// 
    ///     var sa = new Gcp.ServiceAccount.Account("sa", new()
    ///     {
    ///         AccountId = "my-service-account",
    ///         DisplayName = "A service account that Jane can use",
    ///     });
    /// 
    ///     var admin_account_iam = new Gcp.ServiceAccount.IAMMember("admin-account-iam", new()
    ///     {
    ///         ServiceAccountId = sa.Name,
    ///         Role = "roles/iam.serviceAccountUser",
    ///         Member = "user:jane@example.com",
    ///     });
    /// 
    ///     // Allow SA service account use the default GCE account
    ///     var gce_default_account_iam = new Gcp.ServiceAccount.IAMMember("gce-default-account-iam", new()
    ///     {
    ///         ServiceAccountId = @default.Apply(@default =&gt; @default.Apply(getDefaultServiceAccountResult =&gt; getDefaultServiceAccountResult.Name)),
    ///         Role = "roles/iam.serviceAccountUser",
    ///         Member = sa.Email.Apply(email =&gt; $"serviceAccount:{email}"),
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ### Service Account IAM Member With IAM Conditions:
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var sa = new Gcp.ServiceAccount.Account("sa", new()
    ///     {
    ///         AccountId = "my-service-account",
    ///         DisplayName = "A service account that Jane can use",
    ///     });
    /// 
    ///     var admin_account_iam = new Gcp.ServiceAccount.IAMMember("admin-account-iam", new()
    ///     {
    ///         ServiceAccountId = sa.Name,
    ///         Role = "roles/iam.serviceAccountUser",
    ///         Member = "user:jane@example.com",
    ///         Condition = new Gcp.ServiceAccount.Inputs.IAMMemberConditionArgs
    ///         {
    ///             Title = "expires_after_2019_12_31",
    ///             Description = "Expiring at midnight of 2019-12-31",
    ///             Expression = "request.time &lt; timestamp(\"2020-01-01T00:00:00Z\")",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// ### Importing with conditions:
    /// 
    /// Here are examples of importing IAM memberships and bindings that include conditions:
    /// 
    /// ```sh
    /// $ pulumi import gcp:serviceaccount/iAMMember:IAMMember admin-account-iam "projects/{your-project-id}/serviceAccounts/{your-service-account-email} roles/iam.serviceAccountUser expires_after_2019_12_31"
    /// ```
    /// 
    /// ```sh
    /// $ pulumi import gcp:serviceaccount/iAMMember:IAMMember admin-account-iam "projects/{your-project-id}/serviceAccounts/{your-service-account-email} roles/iam.serviceAccountUser user:foo@example.com expires_after_2019_12_31"
    /// ```
    /// </summary>
    [GcpResourceType("gcp:serviceaccount/iAMMember:IAMMember")]
    public partial class IAMMember : global::Pulumi.CustomResource
    {
        /// <summary>
        /// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
        /// Structure is documented below.
        /// </summary>
        [Output("condition")]
        public Output<Outputs.IAMMemberCondition?> Condition { get; private set; } = null!;

        /// <summary>
        /// (Computed) The etag of the service account IAM policy.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

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
        [Output("member")]
        public Output<string> Member { get; private set; } = null!;

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.serviceaccount.IAMBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;

        /// <summary>
        /// The fully-qualified name of the service account to apply policy to.
        /// </summary>
        [Output("serviceAccountId")]
        public Output<string> ServiceAccountId { get; private set; } = null!;


        /// <summary>
        /// Create a IAMMember resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IAMMember(string name, IAMMemberArgs args, CustomResourceOptions? options = null)
            : base("gcp:serviceaccount/iAMMember:IAMMember", name, args ?? new IAMMemberArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IAMMember(string name, Input<string> id, IAMMemberState? state = null, CustomResourceOptions? options = null)
            : base("gcp:serviceaccount/iAMMember:IAMMember", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "gcp:serviceAccount/iAMMember:IAMMember" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing IAMMember resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IAMMember Get(string name, Input<string> id, IAMMemberState? state = null, CustomResourceOptions? options = null)
        {
            return new IAMMember(name, id, state, options);
        }
    }

    public sealed class IAMMemberArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
        /// Structure is documented below.
        /// </summary>
        [Input("condition")]
        public Input<Inputs.IAMMemberConditionArgs>? Condition { get; set; }

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
        [Input("member", required: true)]
        public Input<string> Member { get; set; } = null!;

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.serviceaccount.IAMBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        /// <summary>
        /// The fully-qualified name of the service account to apply policy to.
        /// </summary>
        [Input("serviceAccountId", required: true)]
        public Input<string> ServiceAccountId { get; set; } = null!;

        public IAMMemberArgs()
        {
        }
        public static new IAMMemberArgs Empty => new IAMMemberArgs();
    }

    public sealed class IAMMemberState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
        /// Structure is documented below.
        /// </summary>
        [Input("condition")]
        public Input<Inputs.IAMMemberConditionGetArgs>? Condition { get; set; }

        /// <summary>
        /// (Computed) The etag of the service account IAM policy.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

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
        [Input("member")]
        public Input<string>? Member { get; set; }

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.serviceaccount.IAMBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// The fully-qualified name of the service account to apply policy to.
        /// </summary>
        [Input("serviceAccountId")]
        public Input<string>? ServiceAccountId { get; set; }

        public IAMMemberState()
        {
        }
        public static new IAMMemberState Empty => new IAMMemberState();
    }
}
