// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Kms
{
    public static class GetCryptoKeyIamPolicy
    {
        /// <summary>
        /// Retrieves the current IAM policy data for a Google Cloud KMS crypto key.
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
        ///     var foo = Gcp.Kms.GetCryptoKeyIamPolicy.Invoke(new()
        ///     {
        ///         CryptoKeyId = google_kms_crypto_key.Crypto_key.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetCryptoKeyIamPolicyResult> InvokeAsync(GetCryptoKeyIamPolicyArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCryptoKeyIamPolicyResult>("gcp:kms/getCryptoKeyIamPolicy:getCryptoKeyIamPolicy", args ?? new GetCryptoKeyIamPolicyArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves the current IAM policy data for a Google Cloud KMS crypto key.
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
        ///     var foo = Gcp.Kms.GetCryptoKeyIamPolicy.Invoke(new()
        ///     {
        ///         CryptoKeyId = google_kms_crypto_key.Crypto_key.Id,
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetCryptoKeyIamPolicyResult> Invoke(GetCryptoKeyIamPolicyInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCryptoKeyIamPolicyResult>("gcp:kms/getCryptoKeyIamPolicy:getCryptoKeyIamPolicy", args ?? new GetCryptoKeyIamPolicyInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCryptoKeyIamPolicyArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The crypto key ID, in the form
        /// </summary>
        [Input("cryptoKeyId", required: true)]
        public string CryptoKeyId { get; set; } = null!;

        public GetCryptoKeyIamPolicyArgs()
        {
        }
        public static new GetCryptoKeyIamPolicyArgs Empty => new GetCryptoKeyIamPolicyArgs();
    }

    public sealed class GetCryptoKeyIamPolicyInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The crypto key ID, in the form
        /// </summary>
        [Input("cryptoKeyId", required: true)]
        public Input<string> CryptoKeyId { get; set; } = null!;

        public GetCryptoKeyIamPolicyInvokeArgs()
        {
        }
        public static new GetCryptoKeyIamPolicyInvokeArgs Empty => new GetCryptoKeyIamPolicyInvokeArgs();
    }


    [OutputType]
    public sealed class GetCryptoKeyIamPolicyResult
    {
        public readonly string CryptoKeyId;
        /// <summary>
        /// (Computed) The etag of the IAM policy.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// (Computed) The policy data
        /// </summary>
        public readonly string PolicyData;

        [OutputConstructor]
        private GetCryptoKeyIamPolicyResult(
            string cryptoKeyId,

            string etag,

            string id,

            string policyData)
        {
            CryptoKeyId = cryptoKeyId;
            Etag = etag;
            Id = id;
            PolicyData = policyData;
        }
    }
}
