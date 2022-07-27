// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CertificateAuthority
{
    /// <summary>
    /// A CaPool represents a group of CertificateAuthorities that form a trust anchor. A CaPool can be used to manage
    /// issuance policies for one or more CertificateAuthority resources and to rotate CA certificates in and out of the
    /// trust anchor.
    /// 
    /// ## Example Usage
    /// ### Privateca Capool Basic
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var @default = new Gcp.CertificateAuthority.CaPool("default", new Gcp.CertificateAuthority.CaPoolArgs
    ///         {
    ///             Labels = 
    ///             {
    ///                 { "foo", "bar" },
    ///             },
    ///             Location = "us-central1",
    ///             PublishingOptions = new Gcp.CertificateAuthority.Inputs.CaPoolPublishingOptionsArgs
    ///             {
    ///                 PublishCaCert = true,
    ///                 PublishCrl = true,
    ///             },
    ///             Tier = "ENTERPRISE",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// CaPool can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:certificateauthority/caPool:CaPool default projects/{{project}}/locations/{{location}}/caPools/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:certificateauthority/caPool:CaPool default {{project}}/{{location}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:certificateauthority/caPool:CaPool default {{location}}/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:certificateauthority/caPool:CaPool")]
    public partial class CaPool : Pulumi.CustomResource
    {
        /// <summary>
        /// The IssuancePolicy to control how Certificates will be issued from this CaPool.
        /// Structure is documented below.
        /// </summary>
        [Output("issuancePolicy")]
        public Output<Outputs.CaPoolIssuancePolicy?> IssuancePolicy { get; private set; } = null!;

        /// <summary>
        /// Labels with user-defined metadata.
        /// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass":
        /// "1.3kg", "count": "3" }.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name for this CaPool.
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
        /// The PublishingOptions to follow when issuing Certificates from any CertificateAuthority in this CaPool.
        /// Structure is documented below.
        /// </summary>
        [Output("publishingOptions")]
        public Output<Outputs.CaPoolPublishingOptions?> PublishingOptions { get; private set; } = null!;

        /// <summary>
        /// The Tier of this CaPool.
        /// Possible values are `ENTERPRISE` and `DEVOPS`.
        /// </summary>
        [Output("tier")]
        public Output<string> Tier { get; private set; } = null!;


        /// <summary>
        /// Create a CaPool resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CaPool(string name, CaPoolArgs args, CustomResourceOptions? options = null)
            : base("gcp:certificateauthority/caPool:CaPool", name, args ?? new CaPoolArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CaPool(string name, Input<string> id, CaPoolState? state = null, CustomResourceOptions? options = null)
            : base("gcp:certificateauthority/caPool:CaPool", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CaPool resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CaPool Get(string name, Input<string> id, CaPoolState? state = null, CustomResourceOptions? options = null)
        {
            return new CaPool(name, id, state, options);
        }
    }

    public sealed class CaPoolArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IssuancePolicy to control how Certificates will be issued from this CaPool.
        /// Structure is documented below.
        /// </summary>
        [Input("issuancePolicy")]
        public Input<Inputs.CaPoolIssuancePolicyArgs>? IssuancePolicy { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels with user-defined metadata.
        /// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass":
        /// "1.3kg", "count": "3" }.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The name for this CaPool.
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
        /// The PublishingOptions to follow when issuing Certificates from any CertificateAuthority in this CaPool.
        /// Structure is documented below.
        /// </summary>
        [Input("publishingOptions")]
        public Input<Inputs.CaPoolPublishingOptionsArgs>? PublishingOptions { get; set; }

        /// <summary>
        /// The Tier of this CaPool.
        /// Possible values are `ENTERPRISE` and `DEVOPS`.
        /// </summary>
        [Input("tier", required: true)]
        public Input<string> Tier { get; set; } = null!;

        public CaPoolArgs()
        {
        }
    }

    public sealed class CaPoolState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IssuancePolicy to control how Certificates will be issued from this CaPool.
        /// Structure is documented below.
        /// </summary>
        [Input("issuancePolicy")]
        public Input<Inputs.CaPoolIssuancePolicyGetArgs>? IssuancePolicy { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Labels with user-defined metadata.
        /// An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass":
        /// "1.3kg", "count": "3" }.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name for this CaPool.
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
        /// The PublishingOptions to follow when issuing Certificates from any CertificateAuthority in this CaPool.
        /// Structure is documented below.
        /// </summary>
        [Input("publishingOptions")]
        public Input<Inputs.CaPoolPublishingOptionsGetArgs>? PublishingOptions { get; set; }

        /// <summary>
        /// The Tier of this CaPool.
        /// Possible values are `ENTERPRISE` and `DEVOPS`.
        /// </summary>
        [Input("tier")]
        public Input<string>? Tier { get; set; }

        public CaPoolState()
        {
        }
    }
}
