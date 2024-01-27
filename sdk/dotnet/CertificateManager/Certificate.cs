// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CertificateManager
{
    /// <summary>
    /// Certificate represents a HTTP-reachable backend for a Certificate.
    /// 
    /// ## Example Usage
    /// ### Certificate Manager Google Managed Certificate Dns
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var instance = new Gcp.CertificateManager.DnsAuthorization("instance", new()
    ///     {
    ///         Description = "The default dnss",
    ///         Domain = "subdomain.hashicorptest.com",
    ///     });
    /// 
    ///     var instance2 = new Gcp.CertificateManager.DnsAuthorization("instance2", new()
    ///     {
    ///         Description = "The default dnss",
    ///         Domain = "subdomain2.hashicorptest.com",
    ///     });
    /// 
    ///     var @default = new Gcp.CertificateManager.Certificate("default", new()
    ///     {
    ///         Description = "The default cert",
    ///         Scope = "EDGE_CACHE",
    ///         Labels = 
    ///         {
    ///             { "env", "test" },
    ///         },
    ///         Managed = new Gcp.CertificateManager.Inputs.CertificateManagedArgs
    ///         {
    ///             Domains = new[]
    ///             {
    ///                 instance.Domain,
    ///                 instance2.Domain,
    ///             },
    ///             DnsAuthorizations = new[]
    ///             {
    ///                 instance.Id,
    ///                 instance2.Id,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Certificate Manager Google Managed Certificate Issuance Config
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var pool = new Gcp.CertificateAuthority.CaPool("pool", new()
    ///     {
    ///         Location = "us-central1",
    ///         Tier = "ENTERPRISE",
    ///     });
    /// 
    ///     var caAuthority = new Gcp.CertificateAuthority.Authority("caAuthority", new()
    ///     {
    ///         Location = "us-central1",
    ///         Pool = pool.Name,
    ///         CertificateAuthorityId = "ca-authority",
    ///         Config = new Gcp.CertificateAuthority.Inputs.AuthorityConfigArgs
    ///         {
    ///             SubjectConfig = new Gcp.CertificateAuthority.Inputs.AuthorityConfigSubjectConfigArgs
    ///             {
    ///                 Subject = new Gcp.CertificateAuthority.Inputs.AuthorityConfigSubjectConfigSubjectArgs
    ///                 {
    ///                     Organization = "HashiCorp",
    ///                     CommonName = "my-certificate-authority",
    ///                 },
    ///                 SubjectAltName = new Gcp.CertificateAuthority.Inputs.AuthorityConfigSubjectConfigSubjectAltNameArgs
    ///                 {
    ///                     DnsNames = new[]
    ///                     {
    ///                         "hashicorp.com",
    ///                     },
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
    ///                         ServerAuth = true,
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///         KeySpec = new Gcp.CertificateAuthority.Inputs.AuthorityKeySpecArgs
    ///         {
    ///             Algorithm = "RSA_PKCS1_4096_SHA256",
    ///         },
    ///         DeletionProtection = false,
    ///         SkipGracePeriod = true,
    ///         IgnoreActiveCertificatesOnDeletion = true,
    ///     });
    /// 
    ///     // creating certificate_issuance_config to use it in the managed certificate
    ///     var issuanceconfig = new Gcp.CertificateManager.CertificateIssuanceConfig("issuanceconfig", new()
    ///     {
    ///         Description = "sample description for the certificate issuanceConfigs",
    ///         CertificateAuthorityConfig = new Gcp.CertificateManager.Inputs.CertificateIssuanceConfigCertificateAuthorityConfigArgs
    ///         {
    ///             CertificateAuthorityServiceConfig = new Gcp.CertificateManager.Inputs.CertificateIssuanceConfigCertificateAuthorityConfigCertificateAuthorityServiceConfigArgs
    ///             {
    ///                 CaPool = pool.Id,
    ///             },
    ///         },
    ///         Lifetime = "1814400s",
    ///         RotationWindowPercentage = 34,
    ///         KeyAlgorithm = "ECDSA_P256",
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             caAuthority,
    ///         },
    ///     });
    /// 
    ///     var @default = new Gcp.CertificateManager.Certificate("default", new()
    ///     {
    ///         Description = "The default cert",
    ///         Scope = "EDGE_CACHE",
    ///         Managed = new Gcp.CertificateManager.Inputs.CertificateManagedArgs
    ///         {
    ///             Domains = new[]
    ///             {
    ///                 "terraform.subdomain1.com",
    ///             },
    ///             IssuanceConfig = issuanceconfig.Id,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Certificate Manager Certificate Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var instance = new Gcp.CertificateManager.DnsAuthorization("instance", new()
    ///     {
    ///         Description = "The default dnss",
    ///         Domain = "subdomain.hashicorptest.com",
    ///     });
    /// 
    ///     var instance2 = new Gcp.CertificateManager.DnsAuthorization("instance2", new()
    ///     {
    ///         Description = "The default dnss",
    ///         Domain = "subdomain2.hashicorptest.com",
    ///     });
    /// 
    ///     var @default = new Gcp.CertificateManager.Certificate("default", new()
    ///     {
    ///         Description = "Global cert",
    ///         Scope = "EDGE_CACHE",
    ///         Managed = new Gcp.CertificateManager.Inputs.CertificateManagedArgs
    ///         {
    ///             Domains = new[]
    ///             {
    ///                 instance.Domain,
    ///                 instance2.Domain,
    ///             },
    ///             DnsAuthorizations = new[]
    ///             {
    ///                 instance.Id,
    ///                 instance2.Id,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Certificate Manager Self Managed Certificate Regional
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.IO;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = new Gcp.CertificateManager.Certificate("default", new()
    ///     {
    ///         Description = "Regional cert",
    ///         Location = "us-central1",
    ///         SelfManaged = new Gcp.CertificateManager.Inputs.CertificateSelfManagedArgs
    ///         {
    ///             PemCertificate = File.ReadAllText("test-fixtures/cert.pem"),
    ///             PemPrivateKey = File.ReadAllText("test-fixtures/private-key.pem"),
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Certificate Manager Google Managed Certificate Issuance Config All Regions
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var pool = new Gcp.CertificateAuthority.CaPool("pool", new()
    ///     {
    ///         Location = "us-central1",
    ///         Tier = "ENTERPRISE",
    ///     });
    /// 
    ///     var caAuthority = new Gcp.CertificateAuthority.Authority("caAuthority", new()
    ///     {
    ///         Location = "us-central1",
    ///         Pool = pool.Name,
    ///         CertificateAuthorityId = "ca-authority",
    ///         Config = new Gcp.CertificateAuthority.Inputs.AuthorityConfigArgs
    ///         {
    ///             SubjectConfig = new Gcp.CertificateAuthority.Inputs.AuthorityConfigSubjectConfigArgs
    ///             {
    ///                 Subject = new Gcp.CertificateAuthority.Inputs.AuthorityConfigSubjectConfigSubjectArgs
    ///                 {
    ///                     Organization = "HashiCorp",
    ///                     CommonName = "my-certificate-authority",
    ///                 },
    ///                 SubjectAltName = new Gcp.CertificateAuthority.Inputs.AuthorityConfigSubjectConfigSubjectAltNameArgs
    ///                 {
    ///                     DnsNames = new[]
    ///                     {
    ///                         "hashicorp.com",
    ///                     },
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
    ///                         ServerAuth = true,
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///         KeySpec = new Gcp.CertificateAuthority.Inputs.AuthorityKeySpecArgs
    ///         {
    ///             Algorithm = "RSA_PKCS1_4096_SHA256",
    ///         },
    ///         DeletionProtection = false,
    ///         SkipGracePeriod = true,
    ///         IgnoreActiveCertificatesOnDeletion = true,
    ///     });
    /// 
    ///     // creating certificate_issuance_config to use it in the managed certificate
    ///     var issuanceconfig = new Gcp.CertificateManager.CertificateIssuanceConfig("issuanceconfig", new()
    ///     {
    ///         Description = "sample description for the certificate issuanceConfigs",
    ///         CertificateAuthorityConfig = new Gcp.CertificateManager.Inputs.CertificateIssuanceConfigCertificateAuthorityConfigArgs
    ///         {
    ///             CertificateAuthorityServiceConfig = new Gcp.CertificateManager.Inputs.CertificateIssuanceConfigCertificateAuthorityConfigCertificateAuthorityServiceConfigArgs
    ///             {
    ///                 CaPool = pool.Id,
    ///             },
    ///         },
    ///         Lifetime = "1814400s",
    ///         RotationWindowPercentage = 34,
    ///         KeyAlgorithm = "ECDSA_P256",
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             caAuthority,
    ///         },
    ///     });
    /// 
    ///     var @default = new Gcp.CertificateManager.Certificate("default", new()
    ///     {
    ///         Description = "sample google managed all_regions certificate with issuance config for terraform",
    ///         Scope = "ALL_REGIONS",
    ///         Managed = new Gcp.CertificateManager.Inputs.CertificateManagedArgs
    ///         {
    ///             Domains = new[]
    ///             {
    ///                 "terraform.subdomain1.com",
    ///             },
    ///             IssuanceConfig = issuanceconfig.Id,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Certificate Manager Google Managed Certificate Dns All Regions
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var instance = new Gcp.CertificateManager.DnsAuthorization("instance", new()
    ///     {
    ///         Description = "The default dnss",
    ///         Domain = "subdomain.hashicorptest.com",
    ///     });
    /// 
    ///     var instance2 = new Gcp.CertificateManager.DnsAuthorization("instance2", new()
    ///     {
    ///         Description = "The default dnss",
    ///         Domain = "subdomain2.hashicorptest.com",
    ///     });
    /// 
    ///     var @default = new Gcp.CertificateManager.Certificate("default", new()
    ///     {
    ///         Description = "The default cert",
    ///         Scope = "ALL_REGIONS",
    ///         Managed = new Gcp.CertificateManager.Inputs.CertificateManagedArgs
    ///         {
    ///             Domains = new[]
    ///             {
    ///                 instance.Domain,
    ///                 instance2.Domain,
    ///             },
    ///             DnsAuthorizations = new[]
    ///             {
    ///                 instance.Id,
    ///                 instance2.Id,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Certificate can be imported using any of these accepted formats* `projects/{{project}}/locations/{{location}}/certificates/{{name}}` * `{{project}}/{{location}}/{{name}}` * `{{location}}/{{name}}` When using the `pulumi import` command, Certificate can be imported using one of the formats above. For example
    /// 
    /// ```sh
    ///  $ pulumi import gcp:certificatemanager/certificate:Certificate default projects/{{project}}/locations/{{location}}/certificates/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:certificatemanager/certificate:Certificate default {{project}}/{{location}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:certificatemanager/certificate:Certificate default {{location}}/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:certificatemanager/certificate:Certificate")]
    public partial class Certificate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A human-readable description of the resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
        /// </summary>
        [Output("effectiveLabels")]
        public Output<ImmutableDictionary<string, string>> EffectiveLabels { get; private set; } = null!;

        /// <summary>
        /// Set of label tags associated with the Certificate resource.
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// The Certificate Manager location. If not specified, "global" is used.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Configuration and state of a Managed Certificate.
        /// Certificate Manager provisions and renews Managed Certificates
        /// automatically, for as long as it's authorized to do so.
        /// Structure is documented below.
        /// </summary>
        [Output("managed")]
        public Output<Outputs.CertificateManaged?> Managed { get; private set; } = null!;

        /// <summary>
        /// A user-defined name of the certificate. Certificate names must be unique
        /// The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
        /// and all following characters must be a dash, underscore, letter or digit.
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
        /// The combination of labels configured directly on the resource
        /// and default labels configured on the provider.
        /// </summary>
        [Output("pulumiLabels")]
        public Output<ImmutableDictionary<string, string>> PulumiLabels { get; private set; } = null!;

        /// <summary>
        /// The scope of the certificate.
        /// DEFAULT: Certificates with default scope are served from core Google data centers.
        /// If unsure, choose this option.
        /// EDGE_CACHE: Certificates with scope EDGE_CACHE are special-purposed certificates, served from Edge Points of Presence.
        /// See https://cloud.google.com/vpc/docs/edge-locations.
        /// ALL_REGIONS: Certificates with ALL_REGIONS scope are served from all GCP regions (You can only use ALL_REGIONS with global certs).
        /// See https://cloud.google.com/compute/docs/regions-zones
        /// </summary>
        [Output("scope")]
        public Output<string?> Scope { get; private set; } = null!;

        /// <summary>
        /// Certificate data for a SelfManaged Certificate.
        /// SelfManaged Certificates are uploaded by the user. Updating such
        /// certificates before they expire remains the user's responsibility.
        /// Structure is documented below.
        /// </summary>
        [Output("selfManaged")]
        public Output<Outputs.CertificateSelfManaged?> SelfManaged { get; private set; } = null!;


        /// <summary>
        /// Create a Certificate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Certificate(string name, CertificateArgs? args = null, CustomResourceOptions? options = null)
            : base("gcp:certificatemanager/certificate:Certificate", name, args ?? new CertificateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Certificate(string name, Input<string> id, CertificateState? state = null, CustomResourceOptions? options = null)
            : base("gcp:certificatemanager/certificate:Certificate", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "effectiveLabels",
                    "pulumiLabels",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Certificate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Certificate Get(string name, Input<string> id, CertificateState? state = null, CustomResourceOptions? options = null)
        {
            return new Certificate(name, id, state, options);
        }
    }

    public sealed class CertificateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A human-readable description of the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Set of label tags associated with the Certificate resource.
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The Certificate Manager location. If not specified, "global" is used.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Configuration and state of a Managed Certificate.
        /// Certificate Manager provisions and renews Managed Certificates
        /// automatically, for as long as it's authorized to do so.
        /// Structure is documented below.
        /// </summary>
        [Input("managed")]
        public Input<Inputs.CertificateManagedArgs>? Managed { get; set; }

        /// <summary>
        /// A user-defined name of the certificate. Certificate names must be unique
        /// The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
        /// and all following characters must be a dash, underscore, letter or digit.
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
        /// The scope of the certificate.
        /// DEFAULT: Certificates with default scope are served from core Google data centers.
        /// If unsure, choose this option.
        /// EDGE_CACHE: Certificates with scope EDGE_CACHE are special-purposed certificates, served from Edge Points of Presence.
        /// See https://cloud.google.com/vpc/docs/edge-locations.
        /// ALL_REGIONS: Certificates with ALL_REGIONS scope are served from all GCP regions (You can only use ALL_REGIONS with global certs).
        /// See https://cloud.google.com/compute/docs/regions-zones
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        /// <summary>
        /// Certificate data for a SelfManaged Certificate.
        /// SelfManaged Certificates are uploaded by the user. Updating such
        /// certificates before they expire remains the user's responsibility.
        /// Structure is documented below.
        /// </summary>
        [Input("selfManaged")]
        public Input<Inputs.CertificateSelfManagedArgs>? SelfManaged { get; set; }

        public CertificateArgs()
        {
        }
        public static new CertificateArgs Empty => new CertificateArgs();
    }

    public sealed class CertificateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A human-readable description of the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("effectiveLabels")]
        private InputMap<string>? _effectiveLabels;

        /// <summary>
        /// All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Pulumi, other clients and services.
        /// </summary>
        public InputMap<string> EffectiveLabels
        {
            get => _effectiveLabels ?? (_effectiveLabels = new InputMap<string>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, string>());
                _effectiveLabels = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Set of label tags associated with the Certificate resource.
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// The Certificate Manager location. If not specified, "global" is used.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Configuration and state of a Managed Certificate.
        /// Certificate Manager provisions and renews Managed Certificates
        /// automatically, for as long as it's authorized to do so.
        /// Structure is documented below.
        /// </summary>
        [Input("managed")]
        public Input<Inputs.CertificateManagedGetArgs>? Managed { get; set; }

        /// <summary>
        /// A user-defined name of the certificate. Certificate names must be unique
        /// The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]* which means the first character must be a letter,
        /// and all following characters must be a dash, underscore, letter or digit.
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

        [Input("pulumiLabels")]
        private InputMap<string>? _pulumiLabels;

        /// <summary>
        /// The combination of labels configured directly on the resource
        /// and default labels configured on the provider.
        /// </summary>
        public InputMap<string> PulumiLabels
        {
            get => _pulumiLabels ?? (_pulumiLabels = new InputMap<string>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, string>());
                _pulumiLabels = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        /// <summary>
        /// The scope of the certificate.
        /// DEFAULT: Certificates with default scope are served from core Google data centers.
        /// If unsure, choose this option.
        /// EDGE_CACHE: Certificates with scope EDGE_CACHE are special-purposed certificates, served from Edge Points of Presence.
        /// See https://cloud.google.com/vpc/docs/edge-locations.
        /// ALL_REGIONS: Certificates with ALL_REGIONS scope are served from all GCP regions (You can only use ALL_REGIONS with global certs).
        /// See https://cloud.google.com/compute/docs/regions-zones
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        /// <summary>
        /// Certificate data for a SelfManaged Certificate.
        /// SelfManaged Certificates are uploaded by the user. Updating such
        /// certificates before they expire remains the user's responsibility.
        /// Structure is documented below.
        /// </summary>
        [Input("selfManaged")]
        public Input<Inputs.CertificateSelfManagedGetArgs>? SelfManaged { get; set; }

        public CertificateState()
        {
        }
        public static new CertificateState Empty => new CertificateState();
    }
}
