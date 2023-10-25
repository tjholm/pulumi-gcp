// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.BigQuery
{
    public static class GetDefaultServiceAccount
    {
        /// <summary>
        /// Get the email address of a project's unique BigQuery service account.
        /// 
        /// Each Google Cloud project has a unique service account used by BigQuery. When using
        /// BigQuery with [customer-managed encryption keys](https://cloud.google.com/bigquery/docs/customer-managed-encryption),
        /// this account needs to be granted the
        /// `cloudkms.cryptoKeyEncrypterDecrypter` IAM role on the customer-managed Cloud KMS key used to protect the data.
        /// 
        /// For more information see
        /// [the API reference](https://cloud.google.com/bigquery/docs/reference/rest/v2/projects/getServiceAccount).
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var bqSa = Gcp.BigQuery.GetDefaultServiceAccount.Invoke();
        /// 
        ///     var keySaUser = new Gcp.Kms.CryptoKeyIAMMember("keySaUser", new()
        ///     {
        ///         CryptoKeyId = google_kms_crypto_key.Key.Id,
        ///         Role = "roles/cloudkms.cryptoKeyEncrypterDecrypter",
        ///         Member = $"serviceAccount:{bqSa.Apply(getDefaultServiceAccountResult =&gt; getDefaultServiceAccountResult.Email)}",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDefaultServiceAccountResult> InvokeAsync(GetDefaultServiceAccountArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDefaultServiceAccountResult>("gcp:bigquery/getDefaultServiceAccount:getDefaultServiceAccount", args ?? new GetDefaultServiceAccountArgs(), options.WithDefaults());

        /// <summary>
        /// Get the email address of a project's unique BigQuery service account.
        /// 
        /// Each Google Cloud project has a unique service account used by BigQuery. When using
        /// BigQuery with [customer-managed encryption keys](https://cloud.google.com/bigquery/docs/customer-managed-encryption),
        /// this account needs to be granted the
        /// `cloudkms.cryptoKeyEncrypterDecrypter` IAM role on the customer-managed Cloud KMS key used to protect the data.
        /// 
        /// For more information see
        /// [the API reference](https://cloud.google.com/bigquery/docs/reference/rest/v2/projects/getServiceAccount).
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var bqSa = Gcp.BigQuery.GetDefaultServiceAccount.Invoke();
        /// 
        ///     var keySaUser = new Gcp.Kms.CryptoKeyIAMMember("keySaUser", new()
        ///     {
        ///         CryptoKeyId = google_kms_crypto_key.Key.Id,
        ///         Role = "roles/cloudkms.cryptoKeyEncrypterDecrypter",
        ///         Member = $"serviceAccount:{bqSa.Apply(getDefaultServiceAccountResult =&gt; getDefaultServiceAccountResult.Email)}",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDefaultServiceAccountResult> Invoke(GetDefaultServiceAccountInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDefaultServiceAccountResult>("gcp:bigquery/getDefaultServiceAccount:getDefaultServiceAccount", args ?? new GetDefaultServiceAccountInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDefaultServiceAccountArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The project the unique service account was created for. If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        public GetDefaultServiceAccountArgs()
        {
        }
        public static new GetDefaultServiceAccountArgs Empty => new GetDefaultServiceAccountArgs();
    }

    public sealed class GetDefaultServiceAccountInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The project the unique service account was created for. If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetDefaultServiceAccountInvokeArgs()
        {
        }
        public static new GetDefaultServiceAccountInvokeArgs Empty => new GetDefaultServiceAccountInvokeArgs();
    }


    [OutputType]
    public sealed class GetDefaultServiceAccountResult
    {
        /// <summary>
        /// The email address of the service account. This value is often used to refer to the service account
        /// in order to grant IAM permissions.
        /// </summary>
        public readonly string Email;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The Identity of the service account in the form `serviceAccount:{email}`. This value is often used to refer to the service account in order to grant IAM permissions.
        /// </summary>
        public readonly string Member;
        public readonly string Project;

        [OutputConstructor]
        private GetDefaultServiceAccountResult(
            string email,

            string id,

            string member,

            string project)
        {
            Email = email;
            Id = id;
            Member = member;
            Project = project;
        }
    }
}
