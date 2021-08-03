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
    public sealed class RegionUrlMapPathMatcherPathRule
    {
        /// <summary>
        /// The list of path patterns to match. Each must start with / and the only place a
        /// \* is allowed is at the end following a /. The string fed to the path matcher
        /// does not include any text after the first ? or #, and those chars are not
        /// allowed here.
        /// </summary>
        public readonly ImmutableArray<string> Paths;
        /// <summary>
        /// In response to a matching path, the load balancer performs advanced routing
        /// actions like URL rewrites, header transformations, etc. prior to forwarding the
        /// request to the selected backend. If routeAction specifies any
        /// weightedBackendServices, service must not be set. Conversely if service is set,
        /// routeAction cannot contain any  weightedBackendServices. Only one of routeAction
        /// or urlRedirect must be set.
        /// Structure is documented below.
        /// </summary>
        public readonly Outputs.RegionUrlMapPathMatcherPathRuleRouteAction? RouteAction;
        /// <summary>
        /// A reference to expected RegionBackendService resource the given URL should be mapped to.
        /// </summary>
        public readonly string? Service;
        /// <summary>
        /// When a path pattern is matched, the request is redirected to a URL specified
        /// by urlRedirect. If urlRedirect is specified, service or routeAction must not
        /// be set.
        /// Structure is documented below.
        /// </summary>
        public readonly Outputs.RegionUrlMapPathMatcherPathRuleUrlRedirect? UrlRedirect;

        [OutputConstructor]
        private RegionUrlMapPathMatcherPathRule(
            ImmutableArray<string> paths,

            Outputs.RegionUrlMapPathMatcherPathRuleRouteAction? routeAction,

            string? service,

            Outputs.RegionUrlMapPathMatcherPathRuleUrlRedirect? urlRedirect)
        {
            Paths = paths;
            RouteAction = routeAction;
            Service = service;
            UrlRedirect = urlRedirect;
        }
    }
}
