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
    /// CertificateMap defines a collection of certificate configurations,
    /// which are usable by any associated target proxies
    /// 
    /// ## Example Usage
    /// ### Certificate Manager Certificate Map Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = new Gcp.CertificateManager.CertificateMap("default", new()
    ///     {
    ///         Description = "My acceptance test certificate map",
    ///         Labels = 
    ///         {
    ///             { "terraform", "true" },
    ///             { "acc-test", "true" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// CertificateMap can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:certificatemanager/certificateMap:CertificateMap default projects/{{project}}/locations/global/certificateMaps/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:certificatemanager/certificateMap:CertificateMap default {{project}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:certificatemanager/certificateMap:CertificateMap default {{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:certificatemanager/certificateMap:CertificateMap")]
    public partial class CertificateMap : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Creation timestamp of a Certificate Map. Timestamp is in RFC3339 UTC "Zulu" format,
        /// accurate to nanoseconds with up to nine fractional digits.
        /// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

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
        /// A list of target proxies that use this Certificate Map
        /// Structure is documented below.
        /// </summary>
        [Output("gclbTargets")]
        public Output<ImmutableArray<Outputs.CertificateMapGclbTarget>> GclbTargets { get; private set; } = null!;

        /// <summary>
        /// Set of labels associated with a Certificate Map resource.
        /// 
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>?> Labels { get; private set; } = null!;

        /// <summary>
        /// A user-defined name of the Certificate Map. Certificate Map names must be unique
        /// globally and match the pattern `projects/*/locations/*/certificateMaps/*`.
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
        /// Update timestamp of a Certificate Map. Timestamp is in RFC3339 UTC "Zulu" format,
        /// accurate to nanoseconds with up to nine fractional digits.
        /// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a CertificateMap resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CertificateMap(string name, CertificateMapArgs? args = null, CustomResourceOptions? options = null)
            : base("gcp:certificatemanager/certificateMap:CertificateMap", name, args ?? new CertificateMapArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CertificateMap(string name, Input<string> id, CertificateMapState? state = null, CustomResourceOptions? options = null)
            : base("gcp:certificatemanager/certificateMap:CertificateMap", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CertificateMap resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CertificateMap Get(string name, Input<string> id, CertificateMapState? state = null, CustomResourceOptions? options = null)
        {
            return new CertificateMap(name, id, state, options);
        }
    }

    public sealed class CertificateMapArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A human-readable description of the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Set of labels associated with a Certificate Map resource.
        /// 
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// A user-defined name of the Certificate Map. Certificate Map names must be unique
        /// globally and match the pattern `projects/*/locations/*/certificateMaps/*`.
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

        public CertificateMapArgs()
        {
        }
        public static new CertificateMapArgs Empty => new CertificateMapArgs();
    }

    public sealed class CertificateMapState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Creation timestamp of a Certificate Map. Timestamp is in RFC3339 UTC "Zulu" format,
        /// accurate to nanoseconds with up to nine fractional digits.
        /// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

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

        [Input("gclbTargets")]
        private InputList<Inputs.CertificateMapGclbTargetGetArgs>? _gclbTargets;

        /// <summary>
        /// A list of target proxies that use this Certificate Map
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.CertificateMapGclbTargetGetArgs> GclbTargets
        {
            get => _gclbTargets ?? (_gclbTargets = new InputList<Inputs.CertificateMapGclbTargetGetArgs>());
            set => _gclbTargets = value;
        }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// Set of labels associated with a Certificate Map resource.
        /// 
        /// **Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
        /// Please refer to the field `effective_labels` for all of the labels present on the resource.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        /// <summary>
        /// A user-defined name of the Certificate Map. Certificate Map names must be unique
        /// globally and match the pattern `projects/*/locations/*/certificateMaps/*`.
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
        /// Update timestamp of a Certificate Map. Timestamp is in RFC3339 UTC "Zulu" format,
        /// accurate to nanoseconds with up to nine fractional digits.
        /// Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public CertificateMapState()
        {
        }
        public static new CertificateMapState Empty => new CertificateMapState();
    }
}
