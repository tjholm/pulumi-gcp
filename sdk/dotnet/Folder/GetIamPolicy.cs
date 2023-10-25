// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Folder
{
    public static class GetIamPolicy
    {
        /// <summary>
        /// Retrieves the current IAM policy data for a folder.
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
        ///     var test = Gcp.Folder.GetIamPolicy.Invoke(new()
        ///     {
        ///         Folder = google_folder.Permissiontest.Name,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetIamPolicyResult> InvokeAsync(GetIamPolicyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetIamPolicyResult>("gcp:folder/getIamPolicy:getIamPolicy", args ?? new GetIamPolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the current IAM policy data for a folder.
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
        ///     var test = Gcp.Folder.GetIamPolicy.Invoke(new()
        ///     {
        ///         Folder = google_folder.Permissiontest.Name,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetIamPolicyResult> Invoke(GetIamPolicyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetIamPolicyResult>("gcp:folder/getIamPolicy:getIamPolicy", args ?? new GetIamPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetIamPolicyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
        /// </summary>
        [Input("folder", required: true)]
        public string Folder { get; set; } = null!;

        public GetIamPolicyArgs()
        {
        }
        public static new GetIamPolicyArgs Empty => new GetIamPolicyArgs();
    }

    public sealed class GetIamPolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The resource name of the folder the policy is attached to. Its format is folders/{folder_id}.
        /// </summary>
        [Input("folder", required: true)]
        public Input<string> Folder { get; set; } = null!;

        public GetIamPolicyInvokeArgs()
        {
        }
        public static new GetIamPolicyInvokeArgs Empty => new GetIamPolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetIamPolicyResult
    {
        /// <summary>
        /// (Computed) The etag of the IAM policy.
        /// </summary>
        public readonly string Etag;
        public readonly string Folder;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// (Computed) The policy data
        /// </summary>
        public readonly string PolicyData;

        [OutputConstructor]
        private GetIamPolicyResult(
            string etag,

            string folder,

            string id,

            string policyData)
        {
            Etag = etag;
            Folder = folder;
            Id = id;
            PolicyData = policyData;
        }
    }
}
