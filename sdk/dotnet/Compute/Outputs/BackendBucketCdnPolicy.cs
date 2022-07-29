// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Outputs
{

    [OutputType]
    public sealed class BackendBucketCdnPolicy
    {
        /// <summary>
        /// Bypass the cache when the specified request headers are matched - e.g. Pragma or Authorization headers. Up to 5 headers can be specified. The cache is bypassed for all cdnPolicy.cacheMode settings.
        /// Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.BackendBucketCdnPolicyBypassCacheOnRequestHeader> BypassCacheOnRequestHeaders;
        /// <summary>
        /// The CacheKeyPolicy for this CdnPolicy.
        /// Structure is documented below.
        /// </summary>
        public readonly Outputs.BackendBucketCdnPolicyCacheKeyPolicy? CacheKeyPolicy;
        /// <summary>
        /// Specifies the cache setting for all responses from this backend.
        /// The possible values are: USE_ORIGIN_HEADERS, FORCE_CACHE_ALL and CACHE_ALL_STATIC
        /// Possible values are `USE_ORIGIN_HEADERS`, `FORCE_CACHE_ALL`, and `CACHE_ALL_STATIC`.
        /// </summary>
        public readonly string? CacheMode;
        /// <summary>
        /// Specifies the maximum allowed TTL for cached content served by this origin.
        /// </summary>
        public readonly int? ClientTtl;
        /// <summary>
        /// Specifies the default TTL for cached content served by this origin for responses
        /// that do not have an existing valid TTL (max-age or s-max-age).
        /// </summary>
        public readonly int? DefaultTtl;
        /// <summary>
        /// Specifies the maximum allowed TTL for cached content served by this origin.
        /// </summary>
        public readonly int? MaxTtl;
        /// <summary>
        /// Negative caching allows per-status code TTLs to be set, in order to apply fine-grained caching for common errors or redirects.
        /// </summary>
        public readonly bool? NegativeCaching;
        /// <summary>
        /// Sets a cache TTL for the specified HTTP status code. negativeCaching must be enabled to configure negativeCachingPolicy.
        /// Omitting the policy and leaving negativeCaching enabled will use Cloud CDN's default cache TTLs.
        /// Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.BackendBucketCdnPolicyNegativeCachingPolicy> NegativeCachingPolicies;
        /// <summary>
        /// If true then Cloud CDN will combine multiple concurrent cache fill requests into a small number of requests to the origin.
        /// </summary>
        public readonly bool? RequestCoalescing;
        /// <summary>
        /// Serve existing content from the cache (if available) when revalidating content with the origin, or when an error is encountered when refreshing the cache.
        /// </summary>
        public readonly int? ServeWhileStale;
        /// <summary>
        /// Maximum number of seconds the response to a signed URL request will
        /// be considered fresh. After this time period,
        /// the response will be revalidated before being served.
        /// When serving responses to signed URL requests,
        /// Cloud CDN will internally behave as though
        /// all responses from this backend had a "Cache-Control: public,
        /// max-age=[TTL]" header, regardless of any existing Cache-Control
        /// header. The actual headers served in responses will not be altered.
        /// </summary>
        public readonly int? SignedUrlCacheMaxAgeSec;

        [OutputConstructor]
        private BackendBucketCdnPolicy(
            ImmutableArray<Outputs.BackendBucketCdnPolicyBypassCacheOnRequestHeader> bypassCacheOnRequestHeaders,

            Outputs.BackendBucketCdnPolicyCacheKeyPolicy? cacheKeyPolicy,

            string? cacheMode,

            int? clientTtl,

            int? defaultTtl,

            int? maxTtl,

            bool? negativeCaching,

            ImmutableArray<Outputs.BackendBucketCdnPolicyNegativeCachingPolicy> negativeCachingPolicies,

            bool? requestCoalescing,

            int? serveWhileStale,

            int? signedUrlCacheMaxAgeSec)
        {
            BypassCacheOnRequestHeaders = bypassCacheOnRequestHeaders;
            CacheKeyPolicy = cacheKeyPolicy;
            CacheMode = cacheMode;
            ClientTtl = clientTtl;
            DefaultTtl = defaultTtl;
            MaxTtl = maxTtl;
            NegativeCaching = negativeCaching;
            NegativeCachingPolicies = negativeCachingPolicies;
            RequestCoalescing = requestCoalescing;
            ServeWhileStale = serveWhileStale;
            SignedUrlCacheMaxAgeSec = signedUrlCacheMaxAgeSec;
        }
    }
}
