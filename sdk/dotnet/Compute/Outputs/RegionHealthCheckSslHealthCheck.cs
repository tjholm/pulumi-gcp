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
    public sealed class RegionHealthCheckSslHealthCheck
    {
        /// <summary>
        /// The TCP port number for the HTTP2 health check request.
        /// The default value is 443.
        /// </summary>
        public readonly int? Port;
        /// <summary>
        /// Port name as defined in InstanceGroup#NamedPort#name. If both port and
        /// port_name are defined, port takes precedence.
        /// </summary>
        public readonly string? PortName;
        /// <summary>
        /// Specifies how port is selected for health checking, can be one of the
        /// following values:
        /// * `USE_FIXED_PORT`: The port number in `port` is used for health checking.
        /// * `USE_NAMED_PORT`: The `portName` is used for health checking.
        /// * `USE_SERVING_PORT`: For NetworkEndpointGroup, the port specified for each
        /// network endpoint is used for health checking. For other backends, the
        /// port or named port specified in the Backend Service is used for health
        /// checking.
        /// If not specified, HTTP2 health check follows behavior specified in `port` and
        /// `portName` fields.
        /// Possible values are: `USE_FIXED_PORT`, `USE_NAMED_PORT`, `USE_SERVING_PORT`.
        /// </summary>
        public readonly string? PortSpecification;
        /// <summary>
        /// Specifies the type of proxy header to append before sending data to the
        /// backend.
        /// Default value is `NONE`.
        /// Possible values are: `NONE`, `PROXY_V1`.
        /// </summary>
        public readonly string? ProxyHeader;
        /// <summary>
        /// The application data to send once the SSL connection has been
        /// established (default value is empty). If both request and response are
        /// empty, the connection establishment alone will indicate health. The request
        /// data can only be ASCII.
        /// </summary>
        public readonly string? Request;
        /// <summary>
        /// The bytes to match against the beginning of the response data. If left empty
        /// (the default value), any response will indicate health. The response data
        /// can only be ASCII.
        /// </summary>
        public readonly string? Response;

        [OutputConstructor]
        private RegionHealthCheckSslHealthCheck(
            int? port,

            string? portName,

            string? portSpecification,

            string? proxyHeader,

            string? request,

            string? response)
        {
            Port = port;
            PortName = portName;
            PortSpecification = portSpecification;
            ProxyHeader = proxyHeader;
            Request = request;
            Response = response;
        }
    }
}
