// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CertificateAuthority
{
    public static class GetAuthority
    {
        /// <summary>
        /// Get info about a Google Cloud IAP Client.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var @default = Output.Create(Gcp.CertificateAuthority.GetAuthority.InvokeAsync(new Gcp.CertificateAuthority.GetAuthorityArgs
        ///         {
        ///             Location = "us-west1",
        ///             Pool = "pool-name",
        ///             CertificateAuthorityId = "ca-id",
        ///         }));
        ///         this.Csr = @default.Apply(@default =&gt; @default.PemCsr);
        ///     }
        /// 
        ///     [Output("csr")]
        ///     public Output&lt;string&gt; Csr { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAuthorityResult> InvokeAsync(GetAuthorityArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAuthorityResult>("gcp:certificateauthority/getAuthority:getAuthority", args ?? new GetAuthorityArgs(), options.WithDefaults());

        /// <summary>
        /// Get info about a Google Cloud IAP Client.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Gcp = Pulumi.Gcp;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var @default = Output.Create(Gcp.CertificateAuthority.GetAuthority.InvokeAsync(new Gcp.CertificateAuthority.GetAuthorityArgs
        ///         {
        ///             Location = "us-west1",
        ///             Pool = "pool-name",
        ///             CertificateAuthorityId = "ca-id",
        ///         }));
        ///         this.Csr = @default.Apply(@default =&gt; @default.PemCsr);
        ///     }
        /// 
        ///     [Output("csr")]
        ///     public Output&lt;string&gt; Csr { get; set; }
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetAuthorityResult> Invoke(GetAuthorityInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAuthorityResult>("gcp:certificateauthority/getAuthority:getAuthority", args ?? new GetAuthorityInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAuthorityArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the certificate authority.
        /// </summary>
        [Input("certificateAuthorityId")]
        public string? CertificateAuthorityId { get; set; }

        /// <summary>
        /// The location the certificate authority exists in.
        /// </summary>
        [Input("location")]
        public string? Location { get; set; }

        /// <summary>
        /// The name of the pool the certificate authority belongs to.
        /// </summary>
        [Input("pool")]
        public string? Pool { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        public GetAuthorityArgs()
        {
        }
    }

    public sealed class GetAuthorityInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the certificate authority.
        /// </summary>
        [Input("certificateAuthorityId")]
        public Input<string>? CertificateAuthorityId { get; set; }

        /// <summary>
        /// The location the certificate authority exists in.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the pool the certificate authority belongs to.
        /// </summary>
        [Input("pool")]
        public Input<string>? Pool { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetAuthorityInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAuthorityResult
    {
        public readonly ImmutableArray<Outputs.GetAuthorityAccessUrlResult> AccessUrls;
        public readonly string? CertificateAuthorityId;
        public readonly ImmutableArray<Outputs.GetAuthorityConfigResult> Configs;
        public readonly string CreateTime;
        public readonly string GcsBucket;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool IgnoreActiveCertificatesOnDeletion;
        public readonly ImmutableArray<Outputs.GetAuthorityKeySpecResult> KeySpecs;
        public readonly ImmutableDictionary<string, string> Labels;
        public readonly string Lifetime;
        public readonly string? Location;
        public readonly string Name;
        public readonly ImmutableArray<string> PemCaCertificates;
        /// <summary>
        /// The PEM-encoded signed certificate signing request (CSR). This is only set on subordinate certificate authorities.
        /// </summary>
        public readonly string PemCsr;
        public readonly string? Pool;
        public readonly string? Project;
        public readonly string State;
        public readonly string Type;
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetAuthorityResult(
            ImmutableArray<Outputs.GetAuthorityAccessUrlResult> accessUrls,

            string? certificateAuthorityId,

            ImmutableArray<Outputs.GetAuthorityConfigResult> configs,

            string createTime,

            string gcsBucket,

            string id,

            bool ignoreActiveCertificatesOnDeletion,

            ImmutableArray<Outputs.GetAuthorityKeySpecResult> keySpecs,

            ImmutableDictionary<string, string> labels,

            string lifetime,

            string? location,

            string name,

            ImmutableArray<string> pemCaCertificates,

            string pemCsr,

            string? pool,

            string? project,

            string state,

            string type,

            string updateTime)
        {
            AccessUrls = accessUrls;
            CertificateAuthorityId = certificateAuthorityId;
            Configs = configs;
            CreateTime = createTime;
            GcsBucket = gcsBucket;
            Id = id;
            IgnoreActiveCertificatesOnDeletion = ignoreActiveCertificatesOnDeletion;
            KeySpecs = keySpecs;
            Labels = labels;
            Lifetime = lifetime;
            Location = location;
            Name = name;
            PemCaCertificates = pemCaCertificates;
            PemCsr = pemCsr;
            Pool = pool;
            Project = project;
            State = state;
            Type = type;
            UpdateTime = updateTime;
        }
    }
}
