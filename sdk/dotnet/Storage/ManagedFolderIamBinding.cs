// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Storage
{
    /// <summary>
    /// Three different resources help you manage your IAM policy for Cloud Storage ManagedFolder. Each of these resources serves a different use case:
    /// 
    /// * `gcp.storage.ManagedFolderIamPolicy`: Authoritative. Sets the IAM policy for the managedfolder and replaces any existing policy already attached.
    /// * `gcp.storage.ManagedFolderIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the managedfolder are preserved.
    /// * `gcp.storage.ManagedFolderIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the managedfolder are preserved.
    /// 
    /// A data source can be used to retrieve policy data in advent you do not need creation
    /// 
    /// * `gcp.storage.ManagedFolderIamPolicy`: Retrieves the IAM policy for the managedfolder
    /// 
    /// &gt; **Note:** `gcp.storage.ManagedFolderIamPolicy` **cannot** be used in conjunction with `gcp.storage.ManagedFolderIamBinding` and `gcp.storage.ManagedFolderIamMember` or they will fight over what your policy should be.
    /// 
    /// &gt; **Note:** `gcp.storage.ManagedFolderIamBinding` resources **can be** used in conjunction with `gcp.storage.ManagedFolderIamMember` resources **only if** they do not grant privilege to the same role.
    /// 
    /// &gt; **Note:**  This resource supports IAM Conditions but they have some known limitations which can be found [here](https://cloud.google.com/iam/docs/conditions-overview#limitations). Please review this article if you are having issues with IAM Conditions.
    /// 
    /// ## gcp.storage.ManagedFolderIamPolicy
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
    ///                 Role = "roles/storage.admin",
    ///                 Members = new[]
    ///                 {
    ///                     "user:jane@example.com",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var policy = new Gcp.Storage.ManagedFolderIamPolicy("policy", new()
    ///     {
    ///         Bucket = folder.Bucket,
    ///         ManagedFolder = folder.Name,
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
    ///                 Role = "roles/storage.admin",
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
    ///     var policy = new Gcp.Storage.ManagedFolderIamPolicy("policy", new()
    ///     {
    ///         Bucket = folder.Bucket,
    ///         ManagedFolder = folder.Name,
    ///         PolicyData = admin.Apply(getIAMPolicyResult =&gt; getIAMPolicyResult.PolicyData),
    ///     });
    /// 
    /// });
    /// ```
    /// ## gcp.storage.ManagedFolderIamBinding
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var binding = new Gcp.Storage.ManagedFolderIamBinding("binding", new()
    ///     {
    ///         Bucket = folder.Bucket,
    ///         ManagedFolder = folder.Name,
    ///         Role = "roles/storage.admin",
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
    ///     var binding = new Gcp.Storage.ManagedFolderIamBinding("binding", new()
    ///     {
    ///         Bucket = folder.Bucket,
    ///         ManagedFolder = folder.Name,
    ///         Role = "roles/storage.admin",
    ///         Members = new[]
    ///         {
    ///             "user:jane@example.com",
    ///         },
    ///         Condition = new Gcp.Storage.Inputs.ManagedFolderIamBindingConditionArgs
    ///         {
    ///             Title = "expires_after_2019_12_31",
    ///             Description = "Expiring at midnight of 2019-12-31",
    ///             Expression = "request.time &lt; timestamp(\"2020-01-01T00:00:00Z\")",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ## gcp.storage.ManagedFolderIamMember
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var member = new Gcp.Storage.ManagedFolderIamMember("member", new()
    ///     {
    ///         Bucket = folder.Bucket,
    ///         ManagedFolder = folder.Name,
    ///         Role = "roles/storage.admin",
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
    ///     var member = new Gcp.Storage.ManagedFolderIamMember("member", new()
    ///     {
    ///         Bucket = folder.Bucket,
    ///         ManagedFolder = folder.Name,
    ///         Role = "roles/storage.admin",
    ///         Member = "user:jane@example.com",
    ///         Condition = new Gcp.Storage.Inputs.ManagedFolderIamMemberConditionArgs
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
    /// ## &gt; **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
    /// 
    /// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
    /// ---
    /// 
    /// # IAM policy for Cloud Storage ManagedFolder
    /// Three different resources help you manage your IAM policy for Cloud Storage ManagedFolder. Each of these resources serves a different use case:
    /// 
    /// * `gcp.storage.ManagedFolderIamPolicy`: Authoritative. Sets the IAM policy for the managedfolder and replaces any existing policy already attached.
    /// * `gcp.storage.ManagedFolderIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the managedfolder are preserved.
    /// * `gcp.storage.ManagedFolderIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the managedfolder are preserved.
    /// 
    /// A data source can be used to retrieve policy data in advent you do not need creation
    /// 
    /// * `gcp.storage.ManagedFolderIamPolicy`: Retrieves the IAM policy for the managedfolder
    /// 
    /// &gt; **Note:** `gcp.storage.ManagedFolderIamPolicy` **cannot** be used in conjunction with `gcp.storage.ManagedFolderIamBinding` and `gcp.storage.ManagedFolderIamMember` or they will fight over what your policy should be.
    /// 
    /// &gt; **Note:** `gcp.storage.ManagedFolderIamBinding` resources **can be** used in conjunction with `gcp.storage.ManagedFolderIamMember` resources **only if** they do not grant privilege to the same role.
    /// 
    /// &gt; **Note:**  This resource supports IAM Conditions but they have some known limitations which can be found [here](https://cloud.google.com/iam/docs/conditions-overview#limitations). Please review this article if you are having issues with IAM Conditions.
    /// 
    /// ## gcp.storage.ManagedFolderIamPolicy
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
    ///                 Role = "roles/storage.admin",
    ///                 Members = new[]
    ///                 {
    ///                     "user:jane@example.com",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var policy = new Gcp.Storage.ManagedFolderIamPolicy("policy", new()
    ///     {
    ///         Bucket = folder.Bucket,
    ///         ManagedFolder = folder.Name,
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
    ///                 Role = "roles/storage.admin",
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
    ///     var policy = new Gcp.Storage.ManagedFolderIamPolicy("policy", new()
    ///     {
    ///         Bucket = folder.Bucket,
    ///         ManagedFolder = folder.Name,
    ///         PolicyData = admin.Apply(getIAMPolicyResult =&gt; getIAMPolicyResult.PolicyData),
    ///     });
    /// 
    /// });
    /// ```
    /// ## gcp.storage.ManagedFolderIamBinding
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var binding = new Gcp.Storage.ManagedFolderIamBinding("binding", new()
    ///     {
    ///         Bucket = folder.Bucket,
    ///         ManagedFolder = folder.Name,
    ///         Role = "roles/storage.admin",
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
    ///     var binding = new Gcp.Storage.ManagedFolderIamBinding("binding", new()
    ///     {
    ///         Bucket = folder.Bucket,
    ///         ManagedFolder = folder.Name,
    ///         Role = "roles/storage.admin",
    ///         Members = new[]
    ///         {
    ///             "user:jane@example.com",
    ///         },
    ///         Condition = new Gcp.Storage.Inputs.ManagedFolderIamBindingConditionArgs
    ///         {
    ///             Title = "expires_after_2019_12_31",
    ///             Description = "Expiring at midnight of 2019-12-31",
    ///             Expression = "request.time &lt; timestamp(\"2020-01-01T00:00:00Z\")",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ## gcp.storage.ManagedFolderIamMember
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var member = new Gcp.Storage.ManagedFolderIamMember("member", new()
    ///     {
    ///         Bucket = folder.Bucket,
    ///         ManagedFolder = folder.Name,
    ///         Role = "roles/storage.admin",
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
    ///     var member = new Gcp.Storage.ManagedFolderIamMember("member", new()
    ///     {
    ///         Bucket = folder.Bucket,
    ///         ManagedFolder = folder.Name,
    ///         Role = "roles/storage.admin",
    ///         Member = "user:jane@example.com",
    ///         Condition = new Gcp.Storage.Inputs.ManagedFolderIamMemberConditionArgs
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
    /// * b/{{bucket}}/managedFolders/{{managed_folder}}
    /// 
    /// * {{bucket}}/{{managed_folder}}
    /// 
    /// Any variables not passed in the import command will be taken from the provider configuration.
    /// 
    /// Cloud Storage managedfolder IAM resources can be imported using the resource identifiers, role, and member.
    /// 
    /// IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
    /// 
    /// ```sh
    /// $ pulumi import gcp:storage/managedFolderIamBinding:ManagedFolderIamBinding editor "b/{{bucket}}/managedFolders/{{managed_folder}} roles/storage.objectViewer user:jane@example.com"
    /// ```
    /// 
    /// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
    /// 
    /// ```sh
    /// $ pulumi import gcp:storage/managedFolderIamBinding:ManagedFolderIamBinding editor "b/{{bucket}}/managedFolders/{{managed_folder}} roles/storage.objectViewer"
    /// ```
    /// 
    /// IAM policy imports use the identifier of the resource in question, e.g.
    /// 
    /// ```sh
    /// $ pulumi import gcp:storage/managedFolderIamBinding:ManagedFolderIamBinding editor b/{{bucket}}/managedFolders/{{managed_folder}}
    /// ```
    /// 
    /// -&gt; **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
    /// 
    ///  full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
    /// </summary>
    [GcpResourceType("gcp:storage/managedFolderIamBinding:ManagedFolderIamBinding")]
    public partial class ManagedFolderIamBinding : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the bucket that contains the managed folder. Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        /// <summary>
        /// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
        /// Structure is documented below.
        /// </summary>
        [Output("condition")]
        public Output<Outputs.ManagedFolderIamBindingCondition?> Condition { get; private set; } = null!;

        /// <summary>
        /// (Computed) The etag of the IAM policy.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Output("managedFolder")]
        public Output<string> ManagedFolder { get; private set; } = null!;

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
        /// The role that should be applied. Only one
        /// `gcp.storage.ManagedFolderIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;


        /// <summary>
        /// Create a ManagedFolderIamBinding resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ManagedFolderIamBinding(string name, ManagedFolderIamBindingArgs args, CustomResourceOptions? options = null)
            : base("gcp:storage/managedFolderIamBinding:ManagedFolderIamBinding", name, args ?? new ManagedFolderIamBindingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ManagedFolderIamBinding(string name, Input<string> id, ManagedFolderIamBindingState? state = null, CustomResourceOptions? options = null)
            : base("gcp:storage/managedFolderIamBinding:ManagedFolderIamBinding", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ManagedFolderIamBinding resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ManagedFolderIamBinding Get(string name, Input<string> id, ManagedFolderIamBindingState? state = null, CustomResourceOptions? options = null)
        {
            return new ManagedFolderIamBinding(name, id, state, options);
        }
    }

    public sealed class ManagedFolderIamBindingArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the bucket that contains the managed folder. Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
        /// Structure is documented below.
        /// </summary>
        [Input("condition")]
        public Input<Inputs.ManagedFolderIamBindingConditionArgs>? Condition { get; set; }

        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("managedFolder", required: true)]
        public Input<string> ManagedFolder { get; set; } = null!;

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
        /// The role that should be applied. Only one
        /// `gcp.storage.ManagedFolderIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        public ManagedFolderIamBindingArgs()
        {
        }
        public static new ManagedFolderIamBindingArgs Empty => new ManagedFolderIamBindingArgs();
    }

    public sealed class ManagedFolderIamBindingState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the bucket that contains the managed folder. Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        /// <summary>
        /// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
        /// Structure is documented below.
        /// </summary>
        [Input("condition")]
        public Input<Inputs.ManagedFolderIamBindingConditionGetArgs>? Condition { get; set; }

        /// <summary>
        /// (Computed) The etag of the IAM policy.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("managedFolder")]
        public Input<string>? ManagedFolder { get; set; }

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
        /// The role that should be applied. Only one
        /// `gcp.storage.ManagedFolderIamBinding` can be used per role. Note that custom roles must be of the format
        /// `[projects|organizations]/{parent-name}/roles/{role-name}`.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        public ManagedFolderIamBindingState()
        {
        }
        public static new ManagedFolderIamBindingState Empty => new ManagedFolderIamBindingState();
    }
}
