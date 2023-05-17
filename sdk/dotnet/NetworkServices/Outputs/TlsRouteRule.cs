// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.NetworkServices.Outputs
{

    [OutputType]
    public sealed class TlsRouteRule
    {
        /// <summary>
        /// Required. A detailed rule defining how to route traffic.
        /// Structure is documented below.
        /// </summary>
        public readonly Outputs.TlsRouteRuleAction Action;
        /// <summary>
        /// Matches define the predicate used to match requests to a given action.
        /// Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.TlsRouteRuleMatch> Matches;

        [OutputConstructor]
        private TlsRouteRule(
            Outputs.TlsRouteRuleAction action,

            ImmutableArray<Outputs.TlsRouteRuleMatch> matches)
        {
            Action = action;
            Matches = matches;
        }
    }
}
