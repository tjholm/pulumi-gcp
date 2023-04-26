// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    public static class GetCertificate
    {
        /// <summary>
        /// Get info about a Google Compute SSL Certificate from its name.
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
        ///     var myCert = Gcp.Compute.GetCertificate.Invoke(new()
        ///     {
        ///         Name = "my-cert",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["certificate"] = myCert.Apply(getCertificateResult =&gt; getCertificateResult.Certificate),
        ///         ["certificateId"] = myCert.Apply(getCertificateResult =&gt; getCertificateResult.CertificateId),
        ///         ["selfLink"] = myCert.Apply(getCertificateResult =&gt; getCertificateResult.SelfLink),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetCertificateResult> InvokeAsync(GetCertificateArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCertificateResult>("gcp:compute/getCertificate:getCertificate", args ?? new GetCertificateArgs(), options.WithDefaults());

        /// <summary>
        /// Get info about a Google Compute SSL Certificate from its name.
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
        ///     var myCert = Gcp.Compute.GetCertificate.Invoke(new()
        ///     {
        ///         Name = "my-cert",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["certificate"] = myCert.Apply(getCertificateResult =&gt; getCertificateResult.Certificate),
        ///         ["certificateId"] = myCert.Apply(getCertificateResult =&gt; getCertificateResult.CertificateId),
        ///         ["selfLink"] = myCert.Apply(getCertificateResult =&gt; getCertificateResult.SelfLink),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetCertificateResult> Invoke(GetCertificateInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCertificateResult>("gcp:compute/getCertificate:getCertificate", args ?? new GetCertificateInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCertificateArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the certificate.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        public GetCertificateArgs()
        {
        }
        public static new GetCertificateArgs Empty => new GetCertificateArgs();
    }

    public sealed class GetCertificateInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the certificate.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetCertificateInvokeArgs()
        {
        }
        public static new GetCertificateInvokeArgs Empty => new GetCertificateInvokeArgs();
    }


    [OutputType]
    public sealed class GetCertificateResult
    {
        public readonly string Certificate;
        public readonly int CertificateId;
        public readonly string CreationTimestamp;
        public readonly string Description;
        public readonly string ExpireTime;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string NamePrefix;
        public readonly string PrivateKey;
        public readonly string? Project;
        public readonly string SelfLink;

        [OutputConstructor]
        private GetCertificateResult(
            string certificate,

            int certificateId,

            string creationTimestamp,

            string description,

            string expireTime,

            string id,

            string name,

            string namePrefix,

            string privateKey,

            string? project,

            string selfLink)
        {
            Certificate = certificate;
            CertificateId = certificateId;
            CreationTimestamp = creationTimestamp;
            Description = description;
            ExpireTime = expireTime;
            Id = id;
            Name = name;
            NamePrefix = namePrefix;
            PrivateKey = privateKey;
            Project = project;
            SelfLink = selfLink;
        }
    }
}
