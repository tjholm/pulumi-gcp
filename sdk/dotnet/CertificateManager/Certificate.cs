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
    /// ## Example Usage
    /// ### Certificate Manager Google Managed Certificate
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
    ///         Description = "The default cert",
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
    /// 
    /// ## Import
    /// 
    /// Certificate can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:certificatemanager/certificate:Certificate default projects/{{project}}/locations/global/certificates/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:certificatemanager/certificate:Certificate default {{project}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:certificatemanager/certificate:Certificate default {{name}}
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
        /// Set of label tags associated with the Certificate resource.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

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
        /// The scope of the certificate.
        /// DEFAULT: Certificates with default scope are served from core Google data centers.
        /// If unsure, choose this option.
        /// EDGE_CACHE: Certificates with scope EDGE_CACHE are special-purposed certificates,
        /// served from non-core Google data centers.
        /// Currently allowed only for managed certificates.
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
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

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
        /// EDGE_CACHE: Certificates with scope EDGE_CACHE are special-purposed certificates,
        /// served from non-core Google data centers.
        /// Currently allowed only for managed certificates.
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

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Set of label tags associated with the Certificate resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

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
        /// EDGE_CACHE: Certificates with scope EDGE_CACHE are special-purposed certificates,
        /// served from non-core Google data centers.
        /// Currently allowed only for managed certificates.
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
