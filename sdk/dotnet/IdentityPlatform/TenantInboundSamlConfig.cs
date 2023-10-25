// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.IdentityPlatform
{
    /// <summary>
    /// Inbound SAML configuration for a Identity Toolkit tenant.
    /// 
    /// You must enable the
    /// [Google Identity Platform](https://console.cloud.google.com/marketplace/details/google-cloud-platform/customer-identity) in
    /// the marketplace prior to using this resource.
    /// 
    /// ## Example Usage
    /// ### Identity Platform Tenant Inbound Saml Config Basic
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
    ///     var tenant = new Gcp.IdentityPlatform.Tenant("tenant", new()
    ///     {
    ///         DisplayName = "tenant",
    ///     });
    /// 
    ///     var tenantSamlConfig = new Gcp.IdentityPlatform.TenantInboundSamlConfig("tenantSamlConfig", new()
    ///     {
    ///         DisplayName = "Display Name",
    ///         Tenant = tenant.Name,
    ///         IdpConfig = new Gcp.IdentityPlatform.Inputs.TenantInboundSamlConfigIdpConfigArgs
    ///         {
    ///             IdpEntityId = "tf-idp",
    ///             SignRequest = true,
    ///             SsoUrl = "https://example.com",
    ///             IdpCertificates = new[]
    ///             {
    ///                 new Gcp.IdentityPlatform.Inputs.TenantInboundSamlConfigIdpConfigIdpCertificateArgs
    ///                 {
    ///                     X509Certificate = File.ReadAllText("test-fixtures/rsa_cert.pem"),
    ///                 },
    ///             },
    ///         },
    ///         SpConfig = new Gcp.IdentityPlatform.Inputs.TenantInboundSamlConfigSpConfigArgs
    ///         {
    ///             SpEntityId = "tf-sp",
    ///             CallbackUri = "https://example.com",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// TenantInboundSamlConfig can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:identityplatform/tenantInboundSamlConfig:TenantInboundSamlConfig default projects/{{project}}/tenants/{{tenant}}/inboundSamlConfigs/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:identityplatform/tenantInboundSamlConfig:TenantInboundSamlConfig default {{project}}/{{tenant}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:identityplatform/tenantInboundSamlConfig:TenantInboundSamlConfig default {{tenant}}/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:identityplatform/tenantInboundSamlConfig:TenantInboundSamlConfig")]
    public partial class TenantInboundSamlConfig : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Human friendly display name.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// If this config allows users to sign in with the provider.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// SAML IdP configuration when the project acts as the relying party
        /// Structure is documented below.
        /// </summary>
        [Output("idpConfig")]
        public Output<Outputs.TenantInboundSamlConfigIdpConfig> IdpConfig { get; private set; } = null!;

        /// <summary>
        /// The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters,
        /// hyphens, underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an
        /// alphanumeric character, and have at least 2 characters.
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
        /// SAML SP (Service Provider) configuration when the project acts as the relying party to receive
        /// and accept an authentication assertion issued by a SAML identity provider.
        /// Structure is documented below.
        /// </summary>
        [Output("spConfig")]
        public Output<Outputs.TenantInboundSamlConfigSpConfig> SpConfig { get; private set; } = null!;

        /// <summary>
        /// The name of the tenant where this inbound SAML config resource exists
        /// </summary>
        [Output("tenant")]
        public Output<string> Tenant { get; private set; } = null!;


        /// <summary>
        /// Create a TenantInboundSamlConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TenantInboundSamlConfig(string name, TenantInboundSamlConfigArgs args, CustomResourceOptions? options = null)
            : base("gcp:identityplatform/tenantInboundSamlConfig:TenantInboundSamlConfig", name, args ?? new TenantInboundSamlConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TenantInboundSamlConfig(string name, Input<string> id, TenantInboundSamlConfigState? state = null, CustomResourceOptions? options = null)
            : base("gcp:identityplatform/tenantInboundSamlConfig:TenantInboundSamlConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing TenantInboundSamlConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TenantInboundSamlConfig Get(string name, Input<string> id, TenantInboundSamlConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new TenantInboundSamlConfig(name, id, state, options);
        }
    }

    public sealed class TenantInboundSamlConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Human friendly display name.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// If this config allows users to sign in with the provider.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// SAML IdP configuration when the project acts as the relying party
        /// Structure is documented below.
        /// </summary>
        [Input("idpConfig", required: true)]
        public Input<Inputs.TenantInboundSamlConfigIdpConfigArgs> IdpConfig { get; set; } = null!;

        /// <summary>
        /// The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters,
        /// hyphens, underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an
        /// alphanumeric character, and have at least 2 characters.
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
        /// SAML SP (Service Provider) configuration when the project acts as the relying party to receive
        /// and accept an authentication assertion issued by a SAML identity provider.
        /// Structure is documented below.
        /// </summary>
        [Input("spConfig", required: true)]
        public Input<Inputs.TenantInboundSamlConfigSpConfigArgs> SpConfig { get; set; } = null!;

        /// <summary>
        /// The name of the tenant where this inbound SAML config resource exists
        /// </summary>
        [Input("tenant", required: true)]
        public Input<string> Tenant { get; set; } = null!;

        public TenantInboundSamlConfigArgs()
        {
        }
        public static new TenantInboundSamlConfigArgs Empty => new TenantInboundSamlConfigArgs();
    }

    public sealed class TenantInboundSamlConfigState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Human friendly display name.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// If this config allows users to sign in with the provider.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// SAML IdP configuration when the project acts as the relying party
        /// Structure is documented below.
        /// </summary>
        [Input("idpConfig")]
        public Input<Inputs.TenantInboundSamlConfigIdpConfigGetArgs>? IdpConfig { get; set; }

        /// <summary>
        /// The name of the InboundSamlConfig resource. Must start with 'saml.' and can only have alphanumeric characters,
        /// hyphens, underscores or periods. The part after 'saml.' must also start with a lowercase letter, end with an
        /// alphanumeric character, and have at least 2 characters.
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
        /// SAML SP (Service Provider) configuration when the project acts as the relying party to receive
        /// and accept an authentication assertion issued by a SAML identity provider.
        /// Structure is documented below.
        /// </summary>
        [Input("spConfig")]
        public Input<Inputs.TenantInboundSamlConfigSpConfigGetArgs>? SpConfig { get; set; }

        /// <summary>
        /// The name of the tenant where this inbound SAML config resource exists
        /// </summary>
        [Input("tenant")]
        public Input<string>? Tenant { get; set; }

        public TenantInboundSamlConfigState()
        {
        }
        public static new TenantInboundSamlConfigState Empty => new TenantInboundSamlConfigState();
    }
}
