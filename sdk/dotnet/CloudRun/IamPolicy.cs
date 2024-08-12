// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudRun
{
    /// <summary>
    /// Three different resources help you manage your IAM policy for Cloud Run Service. Each of these resources serves a different use case:
    /// 
    /// * `gcp.cloudrun.IamPolicy`: Authoritative. Sets the IAM policy for the service and replaces any existing policy already attached.
    /// * `gcp.cloudrun.IamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the service are preserved.
    /// * `gcp.cloudrun.IamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the service are preserved.
    /// 
    /// A data source can be used to retrieve policy data in advent you do not need creation
    /// 
    /// * `gcp.cloudrun.IamPolicy`: Retrieves the IAM policy for the service
    /// 
    /// &gt; **Note:** `gcp.cloudrun.IamPolicy` **cannot** be used in conjunction with `gcp.cloudrun.IamBinding` and `gcp.cloudrun.IamMember` or they will fight over what your policy should be.
    /// 
    /// &gt; **Note:** `gcp.cloudrun.IamBinding` resources **can be** used in conjunction with `gcp.cloudrun.IamMember` resources **only if** they do not grant privilege to the same role.
    /// 
    /// ## gcp.cloudrun.IamPolicy
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
    ///     var policy = new Gcp.CloudRun.IamPolicy("policy", new()
    ///     {
    ///         Location = @default.Location,
    ///         Project = @default.Project,
    ///         Service = @default.Name,
    ///         PolicyData = admin.Apply(getIAMPolicyResult =&gt; getIAMPolicyResult.PolicyData),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## gcp.cloudrun.IamBinding
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var binding = new Gcp.CloudRun.IamBinding("binding", new()
    ///     {
    ///         Location = @default.Location,
    ///         Project = @default.Project,
    ///         Service = @default.Name,
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
    /// ## gcp.cloudrun.IamMember
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var member = new Gcp.CloudRun.IamMember("member", new()
    ///     {
    ///         Location = @default.Location,
    ///         Project = @default.Project,
    ///         Service = @default.Name,
    ///         Role = "roles/viewer",
    ///         Member = "user:jane@example.com",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## This resource supports User Project Overrides.
    /// 
    /// - 
    /// 
    /// # IAM policy for Cloud Run Service
    /// Three different resources help you manage your IAM policy for Cloud Run Service. Each of these resources serves a different use case:
    /// 
    /// * `gcp.cloudrun.IamPolicy`: Authoritative. Sets the IAM policy for the service and replaces any existing policy already attached.
    /// * `gcp.cloudrun.IamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the service are preserved.
    /// * `gcp.cloudrun.IamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the service are preserved.
    /// 
    /// A data source can be used to retrieve policy data in advent you do not need creation
    /// 
    /// * `gcp.cloudrun.IamPolicy`: Retrieves the IAM policy for the service
    /// 
    /// &gt; **Note:** `gcp.cloudrun.IamPolicy` **cannot** be used in conjunction with `gcp.cloudrun.IamBinding` and `gcp.cloudrun.IamMember` or they will fight over what your policy should be.
    /// 
    /// &gt; **Note:** `gcp.cloudrun.IamBinding` resources **can be** used in conjunction with `gcp.cloudrun.IamMember` resources **only if** they do not grant privilege to the same role.
    /// 
    /// ## gcp.cloudrun.IamPolicy
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
    ///     var policy = new Gcp.CloudRun.IamPolicy("policy", new()
    ///     {
    ///         Location = @default.Location,
    ///         Project = @default.Project,
    ///         Service = @default.Name,
    ///         PolicyData = admin.Apply(getIAMPolicyResult =&gt; getIAMPolicyResult.PolicyData),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## gcp.cloudrun.IamBinding
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var binding = new Gcp.CloudRun.IamBinding("binding", new()
    ///     {
    ///         Location = @default.Location,
    ///         Project = @default.Project,
    ///         Service = @default.Name,
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
    /// ## gcp.cloudrun.IamMember
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var member = new Gcp.CloudRun.IamMember("member", new()
    ///     {
    ///         Location = @default.Location,
    ///         Project = @default.Project,
    ///         Service = @default.Name,
    ///         Role = "roles/viewer",
    ///         Member = "user:jane@example.com",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// For all import syntaxes, the "resource in question" can take any of the following forms:
    /// 
    /// * projects/{{project}}/locations/{{location}}/services/{{service}}
    /// 
    /// * {{project}}/{{location}}/{{service}}
    /// 
    /// * {{location}}/{{service}}
    /// 
    /// * {{service}}
    /// 
    /// Any variables not passed in the import command will be taken from the provider configuration.
    /// 
    /// Cloud Run service IAM resources can be imported using the resource identifiers, role, and member.
    /// 
    /// IAM member imports use space-delimited identifiers: the resource in question, the role, and the member identity, e.g.
    /// 
    /// ```sh
    /// $ pulumi import gcp:cloudrun/iamPolicy:IamPolicy editor "projects/{{project}}/locations/{{location}}/services/{{service}} roles/viewer user:jane@example.com"
    /// ```
    /// 
    /// IAM binding imports use space-delimited identifiers: the resource in question and the role, e.g.
    /// 
    /// ```sh
    /// $ pulumi import gcp:cloudrun/iamPolicy:IamPolicy editor "projects/{{project}}/locations/{{location}}/services/{{service}} roles/viewer"
    /// ```
    /// 
    /// IAM policy imports use the identifier of the resource in question, e.g.
    /// 
    /// ```sh
    /// $ pulumi import gcp:cloudrun/iamPolicy:IamPolicy editor projects/{{project}}/locations/{{location}}/services/{{service}}
    /// ```
    /// 
    /// -&gt; **Custom Roles**: If you're importing a IAM resource with a custom role, make sure to use the
    /// 
    ///  full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
    /// </summary>
    [GcpResourceType("gcp:cloudrun/iamPolicy:IamPolicy")]
    public partial class IamPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// (Computed) The etag of the IAM policy.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The location of the cloud run instance. eg us-central1 Used to find the parent resource to bind the IAM policy to. If not specified,
        /// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
        /// location is specified, it is taken from the provider configuration.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The policy data generated by
        /// a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        [Output("policyData")]
        public Output<string> PolicyData { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Output("service")]
        public Output<string> Service { get; private set; } = null!;


        /// <summary>
        /// Create a IamPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IamPolicy(string name, IamPolicyArgs args, CustomResourceOptions? options = null)
            : base("gcp:cloudrun/iamPolicy:IamPolicy", name, args ?? new IamPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IamPolicy(string name, Input<string> id, IamPolicyState? state = null, CustomResourceOptions? options = null)
            : base("gcp:cloudrun/iamPolicy:IamPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IamPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IamPolicy Get(string name, Input<string> id, IamPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new IamPolicy(name, id, state, options);
        }
    }

    public sealed class IamPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The location of the cloud run instance. eg us-central1 Used to find the parent resource to bind the IAM policy to. If not specified,
        /// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
        /// location is specified, it is taken from the provider configuration.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The policy data generated by
        /// a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        [Input("policyData", required: true)]
        public Input<string> PolicyData { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("service", required: true)]
        public Input<string> Service { get; set; } = null!;

        public IamPolicyArgs()
        {
        }
        public static new IamPolicyArgs Empty => new IamPolicyArgs();
    }

    public sealed class IamPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Computed) The etag of the IAM policy.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The location of the cloud run instance. eg us-central1 Used to find the parent resource to bind the IAM policy to. If not specified,
        /// the value will be parsed from the identifier of the parent resource. If no location is provided in the parent identifier and no
        /// location is specified, it is taken from the provider configuration.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The policy data generated by
        /// a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        [Input("policyData")]
        public Input<string>? PolicyData { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("service")]
        public Input<string>? Service { get; set; }

        public IamPolicyState()
        {
        }
        public static new IamPolicyState Empty => new IamPolicyState();
    }
}
