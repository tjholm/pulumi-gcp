// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.NetworkSecurity
{
    /// <summary>
    /// The GatewaySecurityPolicy resource contains a collection of GatewaySecurityPolicyRules and associated metadata.
    /// 
    /// To get more information about GatewaySecurityPolicy, see:
    /// 
    /// * [API documentation](https://cloud.google.com/secure-web-proxy/docs/reference/network-security/rest/v1/projects.locations.gatewaySecurityPolicies)
    /// 
    /// ## Example Usage
    /// ### Network Security Gateway Security Policy Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = new Gcp.NetworkSecurity.GatewaySecurityPolicy("default", new()
    ///     {
    ///         Description = "my description",
    ///         Location = "us-central1",
    ///     });
    /// 
    /// });
    /// ```
    /// ### Network Security Gateway Security Policy Tls Inspection Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var defaultCaPool = new Gcp.CertificateAuthority.CaPool("defaultCaPool", new()
    ///     {
    ///         Location = "us-central1",
    ///         Tier = "DEVOPS",
    ///         PublishingOptions = new Gcp.CertificateAuthority.Inputs.CaPoolPublishingOptionsArgs
    ///         {
    ///             PublishCaCert = false,
    ///             PublishCrl = false,
    ///         },
    ///         IssuancePolicy = new Gcp.CertificateAuthority.Inputs.CaPoolIssuancePolicyArgs
    ///         {
    ///             MaximumLifetime = "1209600s",
    ///             BaselineValues = new Gcp.CertificateAuthority.Inputs.CaPoolIssuancePolicyBaselineValuesArgs
    ///             {
    ///                 CaOptions = new Gcp.CertificateAuthority.Inputs.CaPoolIssuancePolicyBaselineValuesCaOptionsArgs
    ///                 {
    ///                     IsCa = false,
    ///                 },
    ///                 KeyUsage = new Gcp.CertificateAuthority.Inputs.CaPoolIssuancePolicyBaselineValuesKeyUsageArgs
    ///                 {
    ///                     BaseKeyUsage = null,
    ///                     ExtendedKeyUsage = new Gcp.CertificateAuthority.Inputs.CaPoolIssuancePolicyBaselineValuesKeyUsageExtendedKeyUsageArgs
    ///                     {
    ///                         ServerAuth = true,
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    ///     var defaultAuthority = new Gcp.CertificateAuthority.Authority("defaultAuthority", new()
    ///     {
    ///         Pool = defaultCaPool.Name,
    ///         CertificateAuthorityId = "my-basic-certificate-authority",
    ///         Location = "us-central1",
    ///         Lifetime = "86400s",
    ///         Type = "SELF_SIGNED",
    ///         DeletionProtection = false,
    ///         SkipGracePeriod = true,
    ///         IgnoreActiveCertificatesOnDeletion = true,
    ///         Config = new Gcp.CertificateAuthority.Inputs.AuthorityConfigArgs
    ///         {
    ///             SubjectConfig = new Gcp.CertificateAuthority.Inputs.AuthorityConfigSubjectConfigArgs
    ///             {
    ///                 Subject = new Gcp.CertificateAuthority.Inputs.AuthorityConfigSubjectConfigSubjectArgs
    ///                 {
    ///                     Organization = "Test LLC",
    ///                     CommonName = "my-ca",
    ///                 },
    ///             },
    ///             X509Config = new Gcp.CertificateAuthority.Inputs.AuthorityConfigX509ConfigArgs
    ///             {
    ///                 CaOptions = new Gcp.CertificateAuthority.Inputs.AuthorityConfigX509ConfigCaOptionsArgs
    ///                 {
    ///                     IsCa = true,
    ///                 },
    ///                 KeyUsage = new Gcp.CertificateAuthority.Inputs.AuthorityConfigX509ConfigKeyUsageArgs
    ///                 {
    ///                     BaseKeyUsage = new Gcp.CertificateAuthority.Inputs.AuthorityConfigX509ConfigKeyUsageBaseKeyUsageArgs
    ///                     {
    ///                         CertSign = true,
    ///                         CrlSign = true,
    ///                     },
    ///                     ExtendedKeyUsage = new Gcp.CertificateAuthority.Inputs.AuthorityConfigX509ConfigKeyUsageExtendedKeyUsageArgs
    ///                     {
    ///                         ServerAuth = false,
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///         KeySpec = new Gcp.CertificateAuthority.Inputs.AuthorityKeySpecArgs
    ///         {
    ///             Algorithm = "RSA_PKCS1_4096_SHA256",
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    ///     var nsSa = new Gcp.Projects.ServiceIdentity("nsSa", new()
    ///     {
    ///         Service = "networksecurity.googleapis.com",
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    ///     var tlsInspectionPermission = new Gcp.CertificateAuthority.CaPoolIamMember("tlsInspectionPermission", new()
    ///     {
    ///         CaPool = defaultCaPool.Id,
    ///         Role = "roles/privateca.certificateManager",
    ///         Member = nsSa.Email.Apply(email =&gt; $"serviceAccount:{email}"),
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    ///     var defaultTlsInspectionPolicy = new Gcp.NetworkSecurity.TlsInspectionPolicy("defaultTlsInspectionPolicy", new()
    ///     {
    ///         Location = "us-central1",
    ///         CaPool = defaultCaPool.Id,
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///         DependsOn = new[]
    ///         {
    ///             defaultCaPool,
    ///             defaultAuthority,
    ///             tlsInspectionPermission,
    ///         },
    ///     });
    /// 
    ///     var defaultGatewaySecurityPolicy = new Gcp.NetworkSecurity.GatewaySecurityPolicy("defaultGatewaySecurityPolicy", new()
    ///     {
    ///         Location = "us-central1",
    ///         Description = "my description",
    ///         TlsInspectionPolicy = defaultTlsInspectionPolicy.Id,
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///         DependsOn = new[]
    ///         {
    ///             defaultTlsInspectionPolicy,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// GatewaySecurityPolicy can be imported using any of these accepted formats* `projects/{{project}}/locations/{{location}}/gatewaySecurityPolicies/{{name}}` * `{{project}}/{{location}}/{{name}}` * `{{location}}/{{name}}` When using the `pulumi import` command, GatewaySecurityPolicy can be imported using one of the formats above. For example
    /// 
    /// ```sh
    ///  $ pulumi import gcp:networksecurity/gatewaySecurityPolicy:GatewaySecurityPolicy default projects/{{project}}/locations/{{location}}/gatewaySecurityPolicies/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:networksecurity/gatewaySecurityPolicy:GatewaySecurityPolicy default {{project}}/{{location}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:networksecurity/gatewaySecurityPolicy:GatewaySecurityPolicy default {{location}}/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:networksecurity/gatewaySecurityPolicy:GatewaySecurityPolicy")]
    public partial class GatewaySecurityPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The timestamp when the resource was created.
        /// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
        /// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z"
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// A free-text description of the resource. Max length 1024 characters.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The location of the gateway security policy.
        /// The default value is `global`.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Name of the resource. Name is of the form projects/{project}/locations/{location}/gatewaySecurityPolicies/{gatewaySecurityPolicy}
        /// gatewaySecurityPolicy should match the pattern:(^a-z?$).
        /// 
        /// 
        /// - - -
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// Server-defined URL of this resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// (Optional, Beta)
        /// Name of a TlsInspectionPolicy resource that defines how TLS inspection is performed for any rule that enables it.
        /// </summary>
        [Output("tlsInspectionPolicy")]
        public Output<string?> TlsInspectionPolicy { get; private set; } = null!;

        /// <summary>
        /// The timestamp when the resource was updated.
        /// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
        /// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a GatewaySecurityPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GatewaySecurityPolicy(string name, GatewaySecurityPolicyArgs? args = null, CustomResourceOptions? options = null)
            : base("gcp:networksecurity/gatewaySecurityPolicy:GatewaySecurityPolicy", name, args ?? new GatewaySecurityPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GatewaySecurityPolicy(string name, Input<string> id, GatewaySecurityPolicyState? state = null, CustomResourceOptions? options = null)
            : base("gcp:networksecurity/gatewaySecurityPolicy:GatewaySecurityPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GatewaySecurityPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GatewaySecurityPolicy Get(string name, Input<string> id, GatewaySecurityPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new GatewaySecurityPolicy(name, id, state, options);
        }
    }

    public sealed class GatewaySecurityPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A free-text description of the resource. Max length 1024 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The location of the gateway security policy.
        /// The default value is `global`.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Name of the resource. Name is of the form projects/{project}/locations/{location}/gatewaySecurityPolicies/{gatewaySecurityPolicy}
        /// gatewaySecurityPolicy should match the pattern:(^a-z?$).
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// (Optional, Beta)
        /// Name of a TlsInspectionPolicy resource that defines how TLS inspection is performed for any rule that enables it.
        /// </summary>
        [Input("tlsInspectionPolicy")]
        public Input<string>? TlsInspectionPolicy { get; set; }

        public GatewaySecurityPolicyArgs()
        {
        }
        public static new GatewaySecurityPolicyArgs Empty => new GatewaySecurityPolicyArgs();
    }

    public sealed class GatewaySecurityPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The timestamp when the resource was created.
        /// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
        /// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z"
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// A free-text description of the resource. Max length 1024 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The location of the gateway security policy.
        /// The default value is `global`.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Name of the resource. Name is of the form projects/{project}/locations/{location}/gatewaySecurityPolicies/{gatewaySecurityPolicy}
        /// gatewaySecurityPolicy should match the pattern:(^a-z?$).
        /// 
        /// 
        /// - - -
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// Server-defined URL of this resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// (Optional, Beta)
        /// Name of a TlsInspectionPolicy resource that defines how TLS inspection is performed for any rule that enables it.
        /// </summary>
        [Input("tlsInspectionPolicy")]
        public Input<string>? TlsInspectionPolicy { get; set; }

        /// <summary>
        /// The timestamp when the resource was updated.
        /// A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
        /// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public GatewaySecurityPolicyState()
        {
        }
        public static new GatewaySecurityPolicyState Empty => new GatewaySecurityPolicyState();
    }
}
