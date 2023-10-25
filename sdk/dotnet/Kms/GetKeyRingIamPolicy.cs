// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Kms
{
    public static class GetKeyRingIamPolicy
    {
        /// <summary>
        /// Retrieves the current IAM policy data for a Google Cloud KMS key ring.
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
        ///     var testKeyRingIamPolicy = Gcp.Kms.GetKeyRingIamPolicy.Invoke(new()
        ///     {
        ///         KeyRingId = "{project_id}/{location_name}/{key_ring_name}",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetKeyRingIamPolicyResult> InvokeAsync(GetKeyRingIamPolicyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetKeyRingIamPolicyResult>("gcp:kms/getKeyRingIamPolicy:getKeyRingIamPolicy", args ?? new GetKeyRingIamPolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the current IAM policy data for a Google Cloud KMS key ring.
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
        ///     var testKeyRingIamPolicy = Gcp.Kms.GetKeyRingIamPolicy.Invoke(new()
        ///     {
        ///         KeyRingId = "{project_id}/{location_name}/{key_ring_name}",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetKeyRingIamPolicyResult> Invoke(GetKeyRingIamPolicyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetKeyRingIamPolicyResult>("gcp:kms/getKeyRingIamPolicy:getKeyRingIamPolicy", args ?? new GetKeyRingIamPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetKeyRingIamPolicyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The key ring ID, in the form
        /// `{project_id}/{location_name}/{key_ring_name}` or
        /// `{location_name}/{key_ring_name}`. In the second form, the provider's
        /// project setting will be used as a fallback.
        /// </summary>
        [Input("keyRingId", required: true)]
        public string KeyRingId { get; set; } = null!;

        public GetKeyRingIamPolicyArgs()
        {
        }
        public static new GetKeyRingIamPolicyArgs Empty => new GetKeyRingIamPolicyArgs();
    }

    public sealed class GetKeyRingIamPolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The key ring ID, in the form
        /// `{project_id}/{location_name}/{key_ring_name}` or
        /// `{location_name}/{key_ring_name}`. In the second form, the provider's
        /// project setting will be used as a fallback.
        /// </summary>
        [Input("keyRingId", required: true)]
        public Input<string> KeyRingId { get; set; } = null!;

        public GetKeyRingIamPolicyInvokeArgs()
        {
        }
        public static new GetKeyRingIamPolicyInvokeArgs Empty => new GetKeyRingIamPolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetKeyRingIamPolicyResult
    {
        /// <summary>
        /// (Computed) The etag of the IAM policy.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string KeyRingId;
        /// <summary>
        /// (Computed) The policy data
        /// </summary>
        public readonly string PolicyData;

        [OutputConstructor]
        private GetKeyRingIamPolicyResult(
            string etag,

            string id,

            string keyRingId,

            string policyData)
        {
            Etag = etag;
            Id = id;
            KeyRingId = keyRingId;
            PolicyData = policyData;
        }
    }
}
