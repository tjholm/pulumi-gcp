// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CertificateAuthority
{
    /// <summary>
    /// Three different resources help you manage your IAM policy for Certificate Authority Service CertificateTemplate. Each of these resources serves a different use case:
    /// 
    /// * `gcp.certificateauthority.CertificateTemplateIamPolicy`: Authoritative. Sets the IAM policy for the certificatetemplate and replaces any existing policy already attached.
    /// * `gcp.certificateauthority.CertificateTemplateIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the certificatetemplate are preserved.
    /// * `gcp.certificateauthority.CertificateTemplateIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the certificatetemplate are preserved.
    /// 
    /// A data source can be used to retrieve policy data in advent you do not need creation
    /// 
    /// * `gcp.certificateauthority.CertificateTemplateIamPolicy`: Retrieves the IAM policy for the certificatetemplate
    /// 
    /// &gt; **Note:** `gcp.certificateauthority.CertificateTemplateIamPolicy` **cannot** be used in conjunction with `gcp.certificateauthority.CertificateTemplateIamBinding` and `gcp.certificateauthority.CertificateTemplateIamMember` or they will fight over what your policy should be.
    /// 
    /// &gt; **Note:** `gcp.certificateauthority.CertificateTemplateIamBinding` resources **can be** used in conjunction with `gcp.certificateauthority.CertificateTemplateIamMember` resources **only if** they do not grant privilege to the same role.
    /// 
    /// &gt; **Note:**  This resource supports IAM Conditions but they have some known limitations which can be found [here](https://cloud.google.com/iam/docs/conditions-overview#limitations). Please review this article if you are having issues with IAM Conditions.
    /// 
    /// ## gcp.certificateauthority.CertificateTemplateIamPolicy
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
    ///                 Role = "roles/privateca.templateUser",
    ///                 Members = new[]
    ///                 {
    ///                     "user:jane@example.com",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var policy = new Gcp.CertificateAuthority.CertificateTemplateIamPolicy("policy", new()
    ///     {
    ///         CertificateTemplate = @default.Id,
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
    ///                 Role = "roles/privateca.templateUser",
    ///                 Members = new[]
    ///                 {
    ///                     "user:jane@example.com",
    ///                 },
    ///                 Condition = new Gcp.Organizations.Inputs.GetIAMPolicyBindingConditionInputArgs
    ///                 {
    ///                     Title = "expires_after_2019_12_31",
    ///                     Description = "Expiring at midnight of 2019-12-31",
    ///                     Expression = "request.time &lt; timestamp(\"2020-01-01T00:00:00Z\")",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var policy = new Gcp.CertificateAuthority.CertificateTemplateIamPolicy("policy", new()
    ///     {
    ///         CertificateTemplate = @default.Id,
    ///         PolicyData = admin.Apply(getIAMPolicyResult =&gt; getIAMPolicyResult.PolicyData),
    ///     });
    /// 
    /// });
    /// ```
    /// ## gcp.certificateauthority.CertificateTemplateIamBinding
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var binding = new Gcp.CertificateAuthority.CertificateTemplateIamBinding("binding", new()
    ///     {
    ///         CertificateTemplate = @default.Id,
    ///         Role = "roles/privateca.templateUser",
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
    ///     var binding = new Gcp.CertificateAuthority.CertificateTemplateIamBinding("binding", new()
    ///     {
    ///         CertificateTemplate = @default.Id,
    ///         Role = "roles/privateca.templateUser",
    ///         Members = new[]
    ///         {
    ///             "user:jane@example.com",
    ///         },
    ///         Condition = new Gcp.CertificateAuthority.Inputs.CertificateTemplateIamBindingConditionArgs
    ///         {
    ///             Title = "expires_after_2019_12_31",
    ///             Description = "Expiring at midnight of 2019-12-31",
    ///             Expression = "request.time &lt; timestamp(\"2020-01-01T00:00:00Z\")",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ## gcp.certificateauthority.CertificateTemplateIamMember
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var member = new Gcp.CertificateAuthority.CertificateTemplateIamMember("member", new()
    ///     {
    ///         CertificateTemplate = @default.Id,
    ///         Role = "roles/privateca.templateUser",
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
    ///     var member = new Gcp.CertificateAuthority.CertificateTemplateIamMember("member", new()
    ///     {
    ///         CertificateTemplate = @default.Id,
    ///         Role = "roles/privateca.templateUser",
    ///         Member = "user:jane@example.com",
    ///         Condition = new Gcp.CertificateAuthority.Inputs.CertificateTemplateIamMemberConditionArgs
    ///         {
    ///             Title = "expires_after_2019_12_31",
    ///             Description = "Expiring at midnight of 2019-12-31",
    ///             Expression = "request.time &lt; timestamp(\"2020-01-01T00:00:00Z\")",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## gcp.certificateauthority.CertificateTemplateIamPolicy
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
    ///                 Role = "roles/privateca.templateUser",
    ///                 Members = new[]
    ///                 {
    ///                     "user:jane@example.com",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var policy = new Gcp.CertificateAuthority.CertificateTemplateIamPolicy("policy", new()
    ///     {
    ///         CertificateTemplate = @default.Id,
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
    ///                 Role = "roles/privateca.templateUser",
    ///                 Members = new[]
    ///                 {
    ///                     "user:jane@example.com",
    ///                 },
    ///                 Condition = new Gcp.Organizations.Inputs.GetIAMPolicyBindingConditionInputArgs
    ///                 {
    ///                     Title = "expires_after_2019_12_31",
    ///                     Description = "Expiring at midnight of 2019-12-31",
    ///                     Expression = "request.time &lt; timestamp(\"2020-01-01T00:00:00Z\")",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var policy = new Gcp.CertificateAuthority.CertificateTemplateIamPolicy("policy", new()
    ///     {
    ///         CertificateTemplate = @default.Id,
    ///         PolicyData = admin.Apply(getIAMPolicyResult =&gt; getIAMPolicyResult.PolicyData),
    ///     });
    /// 
    /// });
    /// ```
    /// ## gcp.certificateauthority.CertificateTemplateIamBinding
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var binding = new Gcp.CertificateAuthority.CertificateTemplateIamBinding("binding", new()
    ///     {
    ///         CertificateTemplate = @default.Id,
    ///         Role = "roles/privateca.templateUser",
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
    ///     var binding = new Gcp.CertificateAuthority.CertificateTemplateIamBinding("binding", new()
    ///     {
    ///         CertificateTemplate = @default.Id,
    ///         Role = "roles/privateca.templateUser",
    ///         Members = new[]
    ///         {
    ///             "user:jane@example.com",
    ///         },
    ///         Condition = new Gcp.CertificateAuthority.Inputs.CertificateTemplateIamBindingConditionArgs
    ///         {
    ///             Title = "expires_after_2019_12_31",
    ///             Description = "Expiring at midnight of 2019-12-31",
    ///             Expression = "request.time &lt; timestamp(\"2020-01-01T00:00:00Z\")",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ## gcp.certificateauthority.CertificateTemplateIamMember
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var member = new Gcp.CertificateAuthority.CertificateTemplateIamMember("member", new()
    ///     {
    ///         CertificateTemplate = @default.Id,
    ///         Role = "roles/privateca.templateUser",
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
    ///     var member = new Gcp.CertificateAuthority.CertificateTemplateIamMember("member", new()
    ///     {
    ///         CertificateTemplate = @default.Id,
    ///         Role = "roles/privateca.templateUser",
    ///         Member = "user:jane@example.com",
    ///         Condition = new Gcp.CertificateAuthority.Inputs.CertificateTemplateIamMemberConditionArgs
    ///         {
    ///             Title = "expires_after_2019_12_31",
    ///             Description = "Expiring at midnight of 2019-12-31",
    ///             Expression = "request.time &lt; timestamp(\"2020-01-01T00:00:00Z\")",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// For all import syntaxes, the "resource in question" can take any of the following forms:
    /// 
    /// * projects/{{project}}/locations/{{location}}/certificateTemplates/{{name}}
    /// 
    /// * {{project}}/{{location}}/{{name}}
    /// 
    /// * {{location}}/{{name}}
    /// 
    /// Any variables not passed in the import command will be taken from the provider configuration.
    /// 
    /// Certificate Authority Service certificatetemplate IAM resources can be imported using the resource identifiers, role, and member.
    /// 
    /// IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
    /// 
    /// ```sh
    /// $ pulumi import gcp:certificateauthority/certificateTemplateIamBinding:CertificateTemplateIamBinding editor "projects/{{project}}/locations/{{location}}/certificateTemplates/{{certificate_template}} roles/privateca.templateUser user:jane@example.com"
    /// ```
    /// 
    /// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
    /// 
    /// ```sh
    /// $ pulumi import gcp:certificateauthority/certificateTemplateIamBinding:CertificateTemplateIamBinding editor "projects/{{project}}/locations/{{location}}/certificateTemplates/{{certificate_template}} roles/privateca.templateUser"
    /// ```
    /// 
    /// IAM policy imports use the identifier of the resource in question, e.g.
    /// 
    /// ```sh
    /// $ pulumi import gcp:certificateauthority/certificateTemplateIamBinding:CertificateTemplateIamBinding editor projects/{{project}}/locations/{{location}}/certificateTemplates/{{certificate_template}}
    /// ```
    /// 
    /// -&gt; **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
    /// 
    ///  full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
    /// </summary>
    [GcpResourceType("gcp:certificateauthority/certificateTemplateIamBinding:CertificateTemplateIamBinding")]
    public partial class CertificateTemplateIamBinding : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Output("certificateTemplate")]
        public Output<string> CertificateTemplate { get; private set; } = null!;

        /// <summary>
        /// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
        /// Structure is documented below.
        /// </summary>
        [Output("condition")]
        public Output<Outputs.CertificateTemplateIamBindingCondition?> Condition { get; private set; } = null!;

        /// <summary>
        /// (Computed) The etag of the IAM policy.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The location for the resource Used to find the parent resource to bind the IAM policy to. If not specified,
        /// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
        /// location is specified, it is taken from the provider configuration.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

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
        [Output("members")]
        public Output<ImmutableArray<string>> Members { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.certificateauthority.CertificateTemplateIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;


        /// <summary>
        /// Create a CertificateTemplateIamBinding resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CertificateTemplateIamBinding(string name, CertificateTemplateIamBindingArgs args, CustomResourceOptions? options = null)
            : base("gcp:certificateauthority/certificateTemplateIamBinding:CertificateTemplateIamBinding", name, args ?? new CertificateTemplateIamBindingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CertificateTemplateIamBinding(string name, Input<string> id, CertificateTemplateIamBindingState? state = null, CustomResourceOptions? options = null)
            : base("gcp:certificateauthority/certificateTemplateIamBinding:CertificateTemplateIamBinding", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CertificateTemplateIamBinding resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CertificateTemplateIamBinding Get(string name, Input<string> id, CertificateTemplateIamBindingState? state = null, CustomResourceOptions? options = null)
        {
            return new CertificateTemplateIamBinding(name, id, state, options);
        }
    }

    public sealed class CertificateTemplateIamBindingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("certificateTemplate", required: true)]
        public Input<string> CertificateTemplate { get; set; } = null!;

        /// <summary>
        /// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
        /// Structure is documented below.
        /// </summary>
        [Input("condition")]
        public Input<Inputs.CertificateTemplateIamBindingConditionArgs>? Condition { get; set; }

        /// <summary>
        /// The location for the resource Used to find the parent resource to bind the IAM policy to. If not specified,
        /// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
        /// location is specified, it is taken from the provider configuration.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

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
        /// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
        /// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
        /// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
        /// </summary>
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.certificateauthority.CertificateTemplateIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        public CertificateTemplateIamBindingArgs()
        {
        }
        public static new CertificateTemplateIamBindingArgs Empty => new CertificateTemplateIamBindingArgs();
    }

    public sealed class CertificateTemplateIamBindingState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("certificateTemplate")]
        public Input<string>? CertificateTemplate { get; set; }

        /// <summary>
        /// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
        /// Structure is documented below.
        /// </summary>
        [Input("condition")]
        public Input<Inputs.CertificateTemplateIamBindingConditionGetArgs>? Condition { get; set; }

        /// <summary>
        /// (Computed) The etag of the IAM policy.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The location for the resource Used to find the parent resource to bind the IAM policy to. If not specified,
        /// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
        /// location is specified, it is taken from the provider configuration.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

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
        /// * **projectOwner:projectid**: Owners of the given project. For example, "projectOwner:my-example-project"
        /// * **projectEditor:projectid**: Editors of the given project. For example, "projectEditor:my-example-project"
        /// * **projectViewer:projectid**: Viewers of the given project. For example, "projectViewer:my-example-project"
        /// </summary>
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The role that should be applied. Only one
        /// `gcp.certificateauthority.CertificateTemplateIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        public CertificateTemplateIamBindingState()
        {
        }
        public static new CertificateTemplateIamBindingState Empty => new CertificateTemplateIamBindingState();
    }
}
