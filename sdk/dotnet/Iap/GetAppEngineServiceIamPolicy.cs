// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Iap
{
    public static class GetAppEngineServiceIamPolicy
    {
        /// <summary>
        /// Retrieves the current IAM policy data for appengineservice
        /// 
        /// 
        /// 
        /// ## example
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var policy = Gcp.Iap.GetAppEngineServiceIamPolicy.Invoke(new()
        ///     {
        ///         AppId = google_app_engine_standard_app_version.Version.Project,
        ///         Project = google_app_engine_standard_app_version.Version.Project,
        ///         Service = google_app_engine_standard_app_version.Version.Service,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetAppEngineServiceIamPolicyResult> InvokeAsync(GetAppEngineServiceIamPolicyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAppEngineServiceIamPolicyResult>("gcp:iap/getAppEngineServiceIamPolicy:getAppEngineServiceIamPolicy", args ?? new GetAppEngineServiceIamPolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the current IAM policy data for appengineservice
        /// 
        /// 
        /// 
        /// ## example
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var policy = Gcp.Iap.GetAppEngineServiceIamPolicy.Invoke(new()
        ///     {
        ///         AppId = google_app_engine_standard_app_version.Version.Project,
        ///         Project = google_app_engine_standard_app_version.Version.Project,
        ///         Service = google_app_engine_standard_app_version.Version.Service,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetAppEngineServiceIamPolicyResult> Invoke(GetAppEngineServiceIamPolicyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAppEngineServiceIamPolicyResult>("gcp:iap/getAppEngineServiceIamPolicy:getAppEngineServiceIamPolicy", args ?? new GetAppEngineServiceIamPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAppEngineServiceIamPolicyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("appId", required: true)]
        public string AppId { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        /// <summary>
        /// Service id of the App Engine application Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("service", required: true)]
        public string Service { get; set; } = null!;

        public GetAppEngineServiceIamPolicyArgs()
        {
        }
        public static new GetAppEngineServiceIamPolicyArgs Empty => new GetAppEngineServiceIamPolicyArgs();
    }

    public sealed class GetAppEngineServiceIamPolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
        /// </summary>
        [Input("appId", required: true)]
        public Input<string> AppId { get; set; } = null!;

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

        public GetAppEngineServiceIamPolicyInvokeArgs()
        {
        }
        public static new GetAppEngineServiceIamPolicyInvokeArgs Empty => new GetAppEngineServiceIamPolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetAppEngineServiceIamPolicyResult
    {
        public readonly string AppId;
        /// <summary>
        /// (Computed) The etag of the IAM policy.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// (Required only by `gcp.iap.AppEngineServiceIamPolicy`) The policy data generated by
        /// a `gcp.organizations.getIAMPolicy` data source.
        /// </summary>
        public readonly string PolicyData;
        public readonly string Project;
        public readonly string Service;

        [OutputConstructor]
        private GetAppEngineServiceIamPolicyResult(
            string appId,

            string etag,

            string id,

            string policyData,

            string project,

            string service)
        {
            AppId = appId;
            Etag = etag;
            Id = id;
            PolicyData = policyData;
            Project = project;
            Service = service;
        }
    }
}
