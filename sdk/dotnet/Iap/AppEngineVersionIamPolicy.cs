// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Iap
{
    /// <summary>
    /// Three different resources help you manage your IAM policy for Identity-Aware Proxy AppEngineVersion. Each of these resources serves a different use case:
    /// 
    /// * `gcp.iap.AppEngineVersionIamPolicy`: Authoritative. Sets the IAM policy for the appengineversion and replaces any existing policy already attached.
    /// * `gcp.iap.AppEngineVersionIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the appengineversion are preserved.
    /// * `gcp.iap.AppEngineVersionIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the appengineversion are preserved.
    /// 
    /// &gt; **Note:** `gcp.iap.AppEngineVersionIamPolicy` **cannot** be used in conjunction with `gcp.iap.AppEngineVersionIamBinding` and `gcp.iap.AppEngineVersionIamMember` or they will fight over what your policy should be.
    /// 
    /// &gt; **Note:** `gcp.iap.AppEngineVersionIamBinding` resources **can be** used in conjunction with `gcp.iap.AppEngineVersionIamMember` resources **only if** they do not grant privilege to the same role.
    /// 
    /// &gt; **Note:**  This resource supports IAM Conditions but they have some known limitations which can be found [here](https://cloud.google.com/iam/docs/conditions-overview#limitations). Please review this article if you are having issues with IAM Conditions.
    /// 
    /// ## google\_iap\_app\_engine\_version\_iam\_policy
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
    ///                 Role = "roles/iap.httpsResourceAccessor",
    ///                 Members = new[]
    ///                 {
    ///                     "user:jane@example.com",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    ///     var policy = new Gcp.Iap.AppEngineVersionIamPolicy("policy", new()
    ///     {
    ///         Project = google_app_engine_standard_app_version.Version.Project,
    ///         AppId = google_app_engine_standard_app_version.Version.Project,
    ///         Service = google_app_engine_standard_app_version.Version.Service,
    ///         VersionId = google_app_engine_standard_app_version.Version.Version_id,
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
    ///                 Role = "roles/iap.httpsResourceAccessor",
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
    ///     var policy = new Gcp.Iap.AppEngineVersionIamPolicy("policy", new()
    ///     {
    ///         Project = google_app_engine_standard_app_version.Version.Project,
    ///         AppId = google_app_engine_standard_app_version.Version.Project,
    ///         Service = google_app_engine_standard_app_version.Version.Service,
    ///         VersionId = google_app_engine_standard_app_version.Version.Version_id,
    ///         PolicyData = admin.Apply(getIAMPolicyResult =&gt; getIAMPolicyResult.PolicyData),
    ///     });
    /// 
    /// });
    /// ```
    /// ## google\_iap\_app\_engine\_version\_iam\_binding
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var binding = new Gcp.Iap.AppEngineVersionIamBinding("binding", new()
    ///     {
    ///         AppId = google_app_engine_standard_app_version.Version.Project,
    ///         Members = new[]
    ///         {
    ///             "user:jane@example.com",
    ///         },
    ///         Project = google_app_engine_standard_app_version.Version.Project,
    ///         Role = "roles/iap.httpsResourceAccessor",
    ///         Service = google_app_engine_standard_app_version.Version.Service,
    ///         VersionId = google_app_engine_standard_app_version.Version.Version_id,
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
    ///     var binding = new Gcp.Iap.AppEngineVersionIamBinding("binding", new()
    ///     {
    ///         AppId = google_app_engine_standard_app_version.Version.Project,
    ///         Condition = new Gcp.Iap.Inputs.AppEngineVersionIamBindingConditionArgs
    ///         {
    ///             Description = "Expiring at midnight of 2019-12-31",
    ///             Expression = "request.time &lt; timestamp(\"2020-01-01T00:00:00Z\")",
    ///             Title = "expires_after_2019_12_31",
    ///         },
    ///         Members = new[]
    ///         {
    ///             "user:jane@example.com",
    ///         },
    ///         Project = google_app_engine_standard_app_version.Version.Project,
    ///         Role = "roles/iap.httpsResourceAccessor",
    ///         Service = google_app_engine_standard_app_version.Version.Service,
    ///         VersionId = google_app_engine_standard_app_version.Version.Version_id,
    ///     });
    /// 
    /// });
    /// ```
    /// ## google\_iap\_app\_engine\_version\_iam\_member
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var member = new Gcp.Iap.AppEngineVersionIamMember("member", new()
    ///     {
    ///         AppId = google_app_engine_standard_app_version.Version.Project,
    ///         Member = "user:jane@example.com",
    ///         Project = google_app_engine_standard_app_version.Version.Project,
    ///         Role = "roles/iap.httpsResourceAccessor",
    ///         Service = google_app_engine_standard_app_version.Version.Service,
    ///         VersionId = google_app_engine_standard_app_version.Version.Version_id,
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
    ///     var member = new Gcp.Iap.AppEngineVersionIamMember("member", new()
    ///     {
    ///         AppId = google_app_engine_standard_app_version.Version.Project,
    ///         Condition = new Gcp.Iap.Inputs.AppEngineVersionIamMemberConditionArgs
    ///         {
    ///             Description = "Expiring at midnight of 2019-12-31",
    ///             Expression = "request.time &lt; timestamp(\"2020-01-01T00:00:00Z\")",
    ///             Title = "expires_after_2019_12_31",
    ///         },
    ///         Member = "user:jane@example.com",
    ///         Project = google_app_engine_standard_app_version.Version.Project,
    ///         Role = "roles/iap.httpsResourceAccessor",
    ///         Service = google_app_engine_standard_app_version.Version.Service,
    ///         VersionId = google_app_engine_standard_app_version.Version.Version_id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/iap_web/appengine-{{appId}}/services/{{service}}/versions/{{versionId}} * {{project}}/{{appId}}/{{service}}/{{versionId}} * {{appId}}/{{service}}/{{versionId}} * {{version}} Any variables not passed in the import command will be taken from the provider configuration. Identity-Aware Proxy appengineversion IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:iap/appEngineVersionIamPolicy:AppEngineVersionIamPolicy editor "projects/{{project}}/iap_web/appengine-{{appId}}/services/{{service}}/versions/{{versionId}} roles/iap.httpsResourceAccessor user:jane@example.com"
    /// ```
    /// 
    ///  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:iap/appEngineVersionIamPolicy:AppEngineVersionIamPolicy editor "projects/{{project}}/iap_web/appengine-{{appId}}/services/{{service}}/versions/{{versionId}} roles/iap.httpsResourceAccessor"
    /// ```
    /// 
    ///  IAM policy imports use the identifier of the resource in question, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import gcp:iap/appEngineVersionIamPolicy:AppEngineVersionIamPolicy editor projects/{{project}}/iap_web/appengine-{{appId}}/services/{{service}}/versions/{{versionId}}
    /// ```
    /// 
    ///  -&gt; **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
    /// 
    /// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
    /// </summary>
    [GcpResourceType("gcp:iap/appEngineVersionIamPolicy:AppEngineVersionIamPolicy")]
    public partial class AppEngineVersionIamPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Output("appId")]
        public Output<string> AppId { get; private set; } = null!;

        /// <summary>
        /// (Computed) The etag of the IAM policy.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

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
        /// Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Output("service")]
        public Output<string> Service { get; private set; } = null!;

        /// <summary>
        /// Version id of the App Engine application Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Output("versionId")]
        public Output<string> VersionId { get; private set; } = null!;


        /// <summary>
        /// Create a AppEngineVersionIamPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AppEngineVersionIamPolicy(string name, AppEngineVersionIamPolicyArgs args, CustomResourceOptions? options = null)
            : base("gcp:iap/appEngineVersionIamPolicy:AppEngineVersionIamPolicy", name, args ?? new AppEngineVersionIamPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AppEngineVersionIamPolicy(string name, Input<string> id, AppEngineVersionIamPolicyState? state = null, CustomResourceOptions? options = null)
            : base("gcp:iap/appEngineVersionIamPolicy:AppEngineVersionIamPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AppEngineVersionIamPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AppEngineVersionIamPolicy Get(string name, Input<string> id, AppEngineVersionIamPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new AppEngineVersionIamPolicy(name, id, state, options);
        }
    }

    public sealed class AppEngineVersionIamPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("appId", required: true)]
        public Input<string> AppId { get; set; } = null!;

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
        /// Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("service", required: true)]
        public Input<string> Service { get; set; } = null!;

        /// <summary>
        /// Version id of the App Engine application Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("versionId", required: true)]
        public Input<string> VersionId { get; set; } = null!;

        public AppEngineVersionIamPolicyArgs()
        {
        }
        public static new AppEngineVersionIamPolicyArgs Empty => new AppEngineVersionIamPolicyArgs();
    }

    public sealed class AppEngineVersionIamPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("appId")]
        public Input<string>? AppId { get; set; }

        /// <summary>
        /// (Computed) The etag of the IAM policy.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

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
        /// Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("service")]
        public Input<string>? Service { get; set; }

        /// <summary>
        /// Version id of the App Engine application Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("versionId")]
        public Input<string>? VersionId { get; set; }

        public AppEngineVersionIamPolicyState()
        {
        }
        public static new AppEngineVersionIamPolicyState Empty => new AppEngineVersionIamPolicyState();
    }
}
