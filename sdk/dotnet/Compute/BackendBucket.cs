// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    /// <summary>
    /// Backend buckets allow you to use Google Cloud Storage buckets with HTTP(S)
    /// load balancing.
    /// 
    /// An HTTP(S) load balancer can direct traffic to specified URLs to a
    /// backend bucket rather than a backend service. It can send requests for
    /// static content to a Cloud Storage bucket and requests for dynamic content
    /// to a virtual machine instance.
    /// 
    /// To get more information about BackendBucket, see:
    /// 
    /// * [API documentation](https://cloud.google.com/compute/docs/reference/v1/backendBuckets)
    /// * How-to Guides
    ///     * [Using a Cloud Storage bucket as a load balancer backend](https://cloud.google.com/compute/docs/load-balancing/http/backend-bucket)
    /// 
    /// ## Example Usage
    /// ### Backend Bucket Basic
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var imageBucket = new Gcp.Storage.Bucket("imageBucket", new()
    ///     {
    ///         Location = "EU",
    ///     });
    /// 
    ///     var imageBackend = new Gcp.Compute.BackendBucket("imageBackend", new()
    ///     {
    ///         Description = "Contains beautiful images",
    ///         BucketName = imageBucket.Name,
    ///         EnableCdn = true,
    ///     });
    /// 
    /// });
    /// ```
    /// ### Backend Bucket Security Policy
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var imageBackendBucket = new Gcp.Storage.Bucket("imageBackendBucket", new()
    ///     {
    ///         Location = "EU",
    ///     });
    /// 
    ///     var policy = new Gcp.Compute.SecurityPolicy("policy", new()
    ///     {
    ///         Description = "basic security policy",
    ///         Type = "CLOUD_ARMOR_EDGE",
    ///     });
    /// 
    ///     var imageBackendBackendBucket = new Gcp.Compute.BackendBucket("imageBackendBackendBucket", new()
    ///     {
    ///         Description = "Contains beautiful images",
    ///         BucketName = imageBackendBucket.Name,
    ///         EnableCdn = true,
    ///         EdgeSecurityPolicy = policy.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// ### Backend Bucket Query String Whitelist
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var imageBucket = new Gcp.Storage.Bucket("imageBucket", new()
    ///     {
    ///         Location = "EU",
    ///     });
    /// 
    ///     var imageBackend = new Gcp.Compute.BackendBucket("imageBackend", new()
    ///     {
    ///         Description = "Contains beautiful images",
    ///         BucketName = imageBucket.Name,
    ///         EnableCdn = true,
    ///         CdnPolicy = new Gcp.Compute.Inputs.BackendBucketCdnPolicyArgs
    ///         {
    ///             CacheKeyPolicy = new Gcp.Compute.Inputs.BackendBucketCdnPolicyCacheKeyPolicyArgs
    ///             {
    ///                 QueryStringWhitelists = new[]
    ///                 {
    ///                     "image-version",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Backend Bucket Include Http Headers
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var imageBucket = new Gcp.Storage.Bucket("imageBucket", new()
    ///     {
    ///         Location = "EU",
    ///     });
    /// 
    ///     var imageBackend = new Gcp.Compute.BackendBucket("imageBackend", new()
    ///     {
    ///         Description = "Contains beautiful images",
    ///         BucketName = imageBucket.Name,
    ///         EnableCdn = true,
    ///         CdnPolicy = new Gcp.Compute.Inputs.BackendBucketCdnPolicyArgs
    ///         {
    ///             CacheKeyPolicy = new Gcp.Compute.Inputs.BackendBucketCdnPolicyCacheKeyPolicyArgs
    ///             {
    ///                 IncludeHttpHeaders = new[]
    ///                 {
    ///                     "X-My-Header-Field",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// BackendBucket can be imported using any of these accepted formats
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/backendBucket:BackendBucket default projects/{{project}}/global/backendBuckets/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/backendBucket:BackendBucket default {{project}}/{{name}}
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import gcp:compute/backendBucket:BackendBucket default {{name}}
    /// ```
    /// </summary>
    [GcpResourceType("gcp:compute/backendBucket:BackendBucket")]
    public partial class BackendBucket : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Cloud Storage bucket name.
        /// </summary>
        [Output("bucketName")]
        public Output<string> BucketName { get; private set; } = null!;

        /// <summary>
        /// Cloud CDN configuration for this Backend Bucket.
        /// Structure is documented below.
        /// </summary>
        [Output("cdnPolicy")]
        public Output<Outputs.BackendBucketCdnPolicy> CdnPolicy { get; private set; } = null!;

        /// <summary>
        /// Compress text responses using Brotli or gzip compression, based on the client's Accept-Encoding header.
        /// Possible values are: `AUTOMATIC`, `DISABLED`.
        /// </summary>
        [Output("compressionMode")]
        public Output<string?> CompressionMode { get; private set; } = null!;

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// Headers that the HTTP/S load balancer should add to proxied responses.
        /// </summary>
        [Output("customResponseHeaders")]
        public Output<ImmutableArray<string>> CustomResponseHeaders { get; private set; } = null!;

        /// <summary>
        /// An optional textual description of the resource; provided by the
        /// client when the resource is created.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The security policy associated with this backend bucket.
        /// </summary>
        [Output("edgeSecurityPolicy")]
        public Output<string?> EdgeSecurityPolicy { get; private set; } = null!;

        /// <summary>
        /// If true, enable Cloud CDN for this BackendBucket.
        /// </summary>
        [Output("enableCdn")]
        public Output<bool?> EnableCdn { get; private set; } = null!;

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is
        /// created. The name must be 1-63 characters long, and comply with
        /// RFC1035.  Specifically, the name must be 1-63 characters long and
        /// match the regular expression `a-z?` which means
        /// the first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the
        /// last character, which cannot be a dash.
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
        /// The URI of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;


        /// <summary>
        /// Create a BackendBucket resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BackendBucket(string name, BackendBucketArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/backendBucket:BackendBucket", name, args ?? new BackendBucketArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BackendBucket(string name, Input<string> id, BackendBucketState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/backendBucket:BackendBucket", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing BackendBucket resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BackendBucket Get(string name, Input<string> id, BackendBucketState? state = null, CustomResourceOptions? options = null)
        {
            return new BackendBucket(name, id, state, options);
        }
    }

    public sealed class BackendBucketArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cloud Storage bucket name.
        /// </summary>
        [Input("bucketName", required: true)]
        public Input<string> BucketName { get; set; } = null!;

        /// <summary>
        /// Cloud CDN configuration for this Backend Bucket.
        /// Structure is documented below.
        /// </summary>
        [Input("cdnPolicy")]
        public Input<Inputs.BackendBucketCdnPolicyArgs>? CdnPolicy { get; set; }

        /// <summary>
        /// Compress text responses using Brotli or gzip compression, based on the client's Accept-Encoding header.
        /// Possible values are: `AUTOMATIC`, `DISABLED`.
        /// </summary>
        [Input("compressionMode")]
        public Input<string>? CompressionMode { get; set; }

        [Input("customResponseHeaders")]
        private InputList<string>? _customResponseHeaders;

        /// <summary>
        /// Headers that the HTTP/S load balancer should add to proxied responses.
        /// </summary>
        public InputList<string> CustomResponseHeaders
        {
            get => _customResponseHeaders ?? (_customResponseHeaders = new InputList<string>());
            set => _customResponseHeaders = value;
        }

        /// <summary>
        /// An optional textual description of the resource; provided by the
        /// client when the resource is created.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The security policy associated with this backend bucket.
        /// </summary>
        [Input("edgeSecurityPolicy")]
        public Input<string>? EdgeSecurityPolicy { get; set; }

        /// <summary>
        /// If true, enable Cloud CDN for this BackendBucket.
        /// </summary>
        [Input("enableCdn")]
        public Input<bool>? EnableCdn { get; set; }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is
        /// created. The name must be 1-63 characters long, and comply with
        /// RFC1035.  Specifically, the name must be 1-63 characters long and
        /// match the regular expression `a-z?` which means
        /// the first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the
        /// last character, which cannot be a dash.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        public BackendBucketArgs()
        {
        }
        public static new BackendBucketArgs Empty => new BackendBucketArgs();
    }

    public sealed class BackendBucketState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cloud Storage bucket name.
        /// </summary>
        [Input("bucketName")]
        public Input<string>? BucketName { get; set; }

        /// <summary>
        /// Cloud CDN configuration for this Backend Bucket.
        /// Structure is documented below.
        /// </summary>
        [Input("cdnPolicy")]
        public Input<Inputs.BackendBucketCdnPolicyGetArgs>? CdnPolicy { get; set; }

        /// <summary>
        /// Compress text responses using Brotli or gzip compression, based on the client's Accept-Encoding header.
        /// Possible values are: `AUTOMATIC`, `DISABLED`.
        /// </summary>
        [Input("compressionMode")]
        public Input<string>? CompressionMode { get; set; }

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

        [Input("customResponseHeaders")]
        private InputList<string>? _customResponseHeaders;

        /// <summary>
        /// Headers that the HTTP/S load balancer should add to proxied responses.
        /// </summary>
        public InputList<string> CustomResponseHeaders
        {
            get => _customResponseHeaders ?? (_customResponseHeaders = new InputList<string>());
            set => _customResponseHeaders = value;
        }

        /// <summary>
        /// An optional textual description of the resource; provided by the
        /// client when the resource is created.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The security policy associated with this backend bucket.
        /// </summary>
        [Input("edgeSecurityPolicy")]
        public Input<string>? EdgeSecurityPolicy { get; set; }

        /// <summary>
        /// If true, enable Cloud CDN for this BackendBucket.
        /// </summary>
        [Input("enableCdn")]
        public Input<bool>? EnableCdn { get; set; }

        /// <summary>
        /// Name of the resource. Provided by the client when the resource is
        /// created. The name must be 1-63 characters long, and comply with
        /// RFC1035.  Specifically, the name must be 1-63 characters long and
        /// match the regular expression `a-z?` which means
        /// the first character must be a lowercase letter, and all following
        /// characters must be a dash, lowercase letter, or digit, except the
        /// last character, which cannot be a dash.
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
        /// The URI of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        public BackendBucketState()
        {
        }
        public static new BackendBucketState Empty => new BackendBucketState();
    }
}
