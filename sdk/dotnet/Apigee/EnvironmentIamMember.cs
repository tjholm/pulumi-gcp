// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Apigee
{
    /// <summary>
    /// Three different resources help you manage your IAM policy for Apigee Environment. Each of these resources serves a different use case:
    /// 
    /// * `gcp.apigee.EnvironmentIamPolicy`: Authoritative. Sets the IAM policy for the environment and replaces any existing policy already attached.
    /// * `gcp.apigee.EnvironmentIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the environment are preserved.
    /// * `gcp.apigee.EnvironmentIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the environment are preserved.
    /// 
    /// A data source can be used to retrieve policy data in advent you do not need creation
    /// 
    /// * `gcp.apigee.EnvironmentIamPolicy`: Retrieves the IAM policy for the environment
    /// 
    /// &gt; **Note:** `gcp.apigee.EnvironmentIamPolicy` **cannot** be used in conjunction with `gcp.apigee.EnvironmentIamBinding` and `gcp.apigee.EnvironmentIamMember` or they will fight over what your policy should be.
    /// 
    /// &gt; **Note:** `gcp.apigee.EnvironmentIamBinding` resources **can be** used in conjunction with `gcp.apigee.EnvironmentIamMember` resources **only if** they do not grant privilege to the same role.
    /// 
    /// ## google\_apigee\_environment\_iam\_policy
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
    ///                 Role = "roles/viewer",
    ///                 Members = new[]
    ///                 {
    ///                     "user:jane@example.com",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var policy = new Gcp.Apigee.EnvironmentIamPolicy("policy", new()
    ///     {
    ///         OrgId = apigeeEnvironment.OrgId,
    ///         EnvId = apigeeEnvironment.Name,
    ///         PolicyData = admin.Apply(getIAMPolicyResult =&gt; getIAMPolicyResult.PolicyData),
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## google\_apigee\_environment\_iam\_binding
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
    ///     var binding = new Gcp.Apigee.EnvironmentIamBinding("binding", new()
    ///     {
    ///         OrgId = apigeeEnvironment.OrgId,
    ///         EnvId = apigeeEnvironment.Name,
    ///         Role = "roles/viewer",
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
    /// ## google\_apigee\_environment\_iam\_member
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
    ///     var member = new Gcp.Apigee.EnvironmentIamMember("member", new()
    ///     {
    ///         OrgId = apigeeEnvironment.OrgId,
    ///         EnvId = apigeeEnvironment.Name,
    ///         Role = "roles/viewer",
    ///         Member = "user:jane@example.com",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## google\_apigee\_environment\_iam\_policy
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
    ///                 Role = "roles/viewer",
    ///                 Members = new[]
    ///                 {
    ///                     "user:jane@example.com",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var policy = new Gcp.Apigee.EnvironmentIamPolicy("policy", new()
    ///     {
    ///         OrgId = apigeeEnvironment.OrgId,
    ///         EnvId = apigeeEnvironment.Name,
    ///         PolicyData = admin.Apply(getIAMPolicyResult =&gt; getIAMPolicyResult.PolicyData),
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## google\_apigee\_environment\_iam\_binding
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
    ///     var binding = new Gcp.Apigee.EnvironmentIamBinding("binding", new()
    ///     {
    ///         OrgId = apigeeEnvironment.OrgId,
    ///         EnvId = apigeeEnvironment.Name,
    ///         Role = "roles/viewer",
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
    /// ## google\_apigee\_environment\_iam\_member
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
    ///     var member = new Gcp.Apigee.EnvironmentIamMember("member", new()
    ///     {
    ///         OrgId = apigeeEnvironment.OrgId,
    ///         EnvId = apigeeEnvironment.Name,
    ///         Role = "roles/viewer",
    ///         Member = "user:jane@example.com",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// For all import syntaxes, the "resource in question" can take any of the following forms:
    /// 
    /// * {{org_id}}/environments/{{name}}
    /// 
    /// * {{name}}
    /// 
    /// Any variables not passed in the import command will be taken from the provider configuration.
    /// 
    /// Apigee environment IAM resources can be imported using the resource identifiers, role, and member.
    /// 
    /// IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
    /// 
    /// ```sh
    /// $ pulumi import gcp:apigee/environmentIamMember:EnvironmentIamMember editor "{{org_id}}/environments/{{environment}} roles/viewer user:jane@example.com"
    /// ```
    /// 
    /// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
    /// 
    /// ```sh
    /// $ pulumi import gcp:apigee/environmentIamMember:EnvironmentIamMember editor "{{org_id}}/environments/{{environment}} roles/viewer"
    /// ```
    /// 
    /// IAM policy imports use the identifier of the resource in question, e.g.
    /// 
    /// ```sh
    /// $ pulumi import gcp:apigee/environmentIamMember:EnvironmentIamMember editor {{org_id}}/environments/{{environment}}
    /// ```
    /// 
    /// -&gt; **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
    /// 
    ///  full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
    /// </summary>
    [GcpResourceType("gcp:apigee/environmentIamMember:EnvironmentIamMember")]
    public partial class EnvironmentIamMember : global::Pulumi.CustomResource
    {
        [Output("condition")]
        public Output<Outputs.EnvironmentIamMemberCondition?> Condition { get; private set; } = null!;

        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Output("envId")]
        public Output<string> EnvId { get; private set; } = null!;

        /// <summary>
        /// (Computed) The etag of the IAM policy.
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
        /// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
        /// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
        /// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
        /// </summary>
        [Output("member")]
        public Output<string> Member { get; private set; } = null!;

        [Output("orgId")]
        public Output<string> OrgId { get; private set; } = null!;

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.apigee.EnvironmentIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;


        /// <summary>
        /// Create a EnvironmentIamMember resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EnvironmentIamMember(string name, EnvironmentIamMemberArgs args, CustomResourceOptions? options = null)
            : base("gcp:apigee/environmentIamMember:EnvironmentIamMember", name, args ?? new EnvironmentIamMemberArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EnvironmentIamMember(string name, Input<string> id, EnvironmentIamMemberState? state = null, CustomResourceOptions? options = null)
            : base("gcp:apigee/environmentIamMember:EnvironmentIamMember", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EnvironmentIamMember resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EnvironmentIamMember Get(string name, Input<string> id, EnvironmentIamMemberState? state = null, CustomResourceOptions? options = null)
        {
            return new EnvironmentIamMember(name, id, state, options);
        }
    }

    public sealed class EnvironmentIamMemberArgs : global::Pulumi.ResourceArgs
    {
        [Input("condition")]
        public Input<Inputs.EnvironmentIamMemberConditionArgs>? Condition { get; set; }

        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("envId", required: true)]
        public Input<string> EnvId { get; set; } = null!;

        /// <summary>
        /// Identities that will be granted the privilege in `role`.
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
        [Input("member", required: true)]
        public Input<string> Member { get; set; } = null!;

        [Input("orgId", required: true)]
        public Input<string> OrgId { get; set; } = null!;

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.apigee.EnvironmentIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        public EnvironmentIamMemberArgs()
        {
        }
        public static new EnvironmentIamMemberArgs Empty => new EnvironmentIamMemberArgs();
    }

    public sealed class EnvironmentIamMemberState : global::Pulumi.ResourceArgs
    {
        [Input("condition")]
        public Input<Inputs.EnvironmentIamMemberConditionGetArgs>? Condition { get; set; }

        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("envId")]
        public Input<string>? EnvId { get; set; }

        /// <summary>
        /// (Computed) The etag of the IAM policy.
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
        /// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
        /// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
        /// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
        /// </summary>
        [Input("member")]
        public Input<string>? Member { get; set; }

        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.apigee.EnvironmentIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        public EnvironmentIamMemberState()
        {
        }
        public static new EnvironmentIamMemberState Empty => new EnvironmentIamMemberState();
    }
}
