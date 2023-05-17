// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudRun.Outputs
{

    [OutputType]
    public sealed class ServiceTemplateSpecContainerStartupProbeHttpGet
    {
        /// <summary>
        /// Custom headers to set in the request. HTTP allows repeated headers.
        /// Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.ServiceTemplateSpecContainerStartupProbeHttpGetHttpHeader> HttpHeaders;
        /// <summary>
        /// Path to access on the HTTP server. If set, it should not be empty string.
        /// </summary>
        public readonly string? Path;
        /// <summary>
        /// Port number to access on the container. Number must be in the range 1 to 65535.
        /// If not specified, defaults to the same value as container.ports[0].containerPort.
        /// </summary>
        public readonly int? Port;

        [OutputConstructor]
        private ServiceTemplateSpecContainerStartupProbeHttpGet(
            ImmutableArray<Outputs.ServiceTemplateSpecContainerStartupProbeHttpGetHttpHeader> httpHeaders,

            string? path,

            int? port)
        {
            HttpHeaders = httpHeaders;
            Path = path;
            Port = port;
        }
    }
}
