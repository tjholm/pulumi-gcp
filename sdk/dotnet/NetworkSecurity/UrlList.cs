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
    /// ## Example Usage
    /// ### Network Security Url Lists Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = new Gcp.NetworkSecurity.UrlList("default", new()
    ///     {
    ///         Location = "us-central1",
    ///         Values = new[]
    ///         {
    ///             "www.example.com",
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    /// });
    /// ```
    /// ### Network Security Url Lists Advanced
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = new Gcp.NetworkSecurity.UrlList("default", new()
    ///     {
    ///         Location = "us-central1",
    ///         Description = "my description",
    ///         Values = new[]
    ///         {
    ///             "www.example.com",
    ///             "about.example.com",
    ///             "github.com/example-org/*",
    ///         },
    ///     }, new CustomResourceOptions
    ///     {
    ///         Provider = google_beta,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// UrlLists can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:networksecurity/urlList:UrlList default projects/{{project}}/locations/{{location}}/urlLists/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:networksecurity/urlList:UrlList default {{project}}/{{location}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:networksecurity/urlList:UrlList default {{location}}/{{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:networksecurity/urlList:UrlList")]
    public partial class UrlList : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Output only. Time when the security policy was created.
        /// A timestamp in RFC3339 UTC 'Zulu' format, with nanosecond resolution and up to nine fractional digits.
        /// Examples: '2014-10-02T15:01:23Z' and '2014-10-02T15:01:23.045123456Z'
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// Free-text description of the resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The location of the url lists.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Short name of the UrlList resource to be created.
        /// This value should be 1-63 characters long, containing only letters, numbers, hyphens, and underscores, and should not start with a number. E.g. 'urlList'.
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
        /// Output only. Time when the security policy was updated.
        /// A timestamp in RFC3339 UTC 'Zulu' format, with nanosecond resolution and up to nine fractional digits.
        /// Examples: '2014-10-02T15:01:23Z' and '2014-10-02T15:01:23.045123456Z'.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// FQDNs and URLs.
        /// </summary>
        [Output("values")]
        public Output<ImmutableArray<string>> Values { get; private set; } = null!;


        /// <summary>
        /// Create a UrlList resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UrlList(string name, UrlListArgs args, CustomResourceOptions? options = null)
            : base("gcp:networksecurity/urlList:UrlList", name, args ?? new UrlListArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UrlList(string name, Input<string> id, UrlListState? state = null, CustomResourceOptions? options = null)
            : base("gcp:networksecurity/urlList:UrlList", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UrlList resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UrlList Get(string name, Input<string> id, UrlListState? state = null, CustomResourceOptions? options = null)
        {
            return new UrlList(name, id, state, options);
        }
    }

    public sealed class UrlListArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Free-text description of the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The location of the url lists.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Short name of the UrlList resource to be created.
        /// This value should be 1-63 characters long, containing only letters, numbers, hyphens, and underscores, and should not start with a number. E.g. 'urlList'.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("values", required: true)]
        private InputList<string>? _values;

        /// <summary>
        /// FQDNs and URLs.
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public UrlListArgs()
        {
        }
        public static new UrlListArgs Empty => new UrlListArgs();
    }

    public sealed class UrlListState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Output only. Time when the security policy was created.
        /// A timestamp in RFC3339 UTC 'Zulu' format, with nanosecond resolution and up to nine fractional digits.
        /// Examples: '2014-10-02T15:01:23Z' and '2014-10-02T15:01:23.045123456Z'
        /// </summary>
        [Input("createTime")]
        public Input<string>? CreateTime { get; set; }

        /// <summary>
        /// Free-text description of the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The location of the url lists.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Short name of the UrlList resource to be created.
        /// This value should be 1-63 characters long, containing only letters, numbers, hyphens, and underscores, and should not start with a number. E.g. 'urlList'.
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
        /// Output only. Time when the security policy was updated.
        /// A timestamp in RFC3339 UTC 'Zulu' format, with nanosecond resolution and up to nine fractional digits.
        /// Examples: '2014-10-02T15:01:23Z' and '2014-10-02T15:01:23.045123456Z'.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        [Input("values")]
        private InputList<string>? _values;

        /// <summary>
        /// FQDNs and URLs.
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public UrlListState()
        {
        }
        public static new UrlListState Empty => new UrlListState();
    }
}
