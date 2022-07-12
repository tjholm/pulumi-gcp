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
    public sealed class RegionBackendServiceBackend
    {
        /// <summary>
        /// Specifies the balancing mode for this backend.
        /// See the [Backend Services Overview](https://cloud.google.com/load-balancing/docs/backend-service#balancing-mode)
        /// for an explanation of load balancing modes.
        /// Default value is `CONNECTION`.
        /// Possible values are `UTILIZATION`, `RATE`, and `CONNECTION`.
        /// </summary>
        public readonly string? BalancingMode;
        /// <summary>
        /// A multiplier applied to the group's maximum servicing capacity
        /// (based on UTILIZATION, RATE or CONNECTION).
        /// ~&gt;**NOTE**: This field cannot be set for
        /// INTERNAL region backend services (default loadBalancingScheme),
        /// but is required for non-INTERNAL backend service. The total
        /// capacity_scaler for all backends must be non-zero.
        /// A setting of 0 means the group is completely drained, offering
        /// 0% of its available Capacity. Valid range is [0.0,1.0].
        /// </summary>
        public readonly double? CapacityScaler;
        /// <summary>
        /// An optional description of this resource.
        /// Provide this property when you create the resource.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// This field designates whether this is a failover backend. More
        /// than one failover backend can be configured for a given RegionBackendService.
        /// </summary>
        public readonly bool? Failover;
        /// <summary>
        /// The fully-qualified URL of an Instance Group or Network Endpoint
        /// Group resource. In case of instance group this defines the list
        /// of instances that serve traffic. Member virtual machine
        /// instances from each instance group must live in the same zone as
        /// the instance group itself. No two backends in a backend service
        /// are allowed to use same Instance Group resource.
        /// For Network Endpoint Groups this defines list of endpoints. All
        /// endpoints of Network Endpoint Group must be hosted on instances
        /// located in the same zone as the Network Endpoint Group.
        /// Backend services cannot mix Instance Group and
        /// Network Endpoint Group backends.
        /// When the `load_balancing_scheme` is INTERNAL, only instance groups
        /// are supported.
        /// Note that you must specify an Instance Group or Network Endpoint
        /// Group resource using the fully-qualified URL, rather than a
        /// partial URL.
        /// </summary>
        public readonly string Group;
        /// <summary>
        /// The maximum number of connections to the backend cluster.
        /// Defaults to 1024.
        /// </summary>
        public readonly int? MaxConnections;
        /// <summary>
        /// The max number of simultaneous connections that a single backend
        /// network endpoint can handle. Cannot be set
        /// for INTERNAL backend services.
        /// This is used to calculate the capacity of the group. Can be
        /// used in either CONNECTION or UTILIZATION balancing modes. For
        /// CONNECTION mode, either maxConnections or
        /// maxConnectionsPerEndpoint must be set.
        /// </summary>
        public readonly int? MaxConnectionsPerEndpoint;
        /// <summary>
        /// The max number of simultaneous connections that a single
        /// backend instance can handle. Cannot be set for INTERNAL backend
        /// services.
        /// This is used to calculate the capacity of the group.
        /// Can be used in either CONNECTION or UTILIZATION balancing modes.
        /// For CONNECTION mode, either maxConnections or
        /// maxConnectionsPerInstance must be set.
        /// </summary>
        public readonly int? MaxConnectionsPerInstance;
        /// <summary>
        /// The max requests per second (RPS) of the group. Cannot be set
        /// for INTERNAL backend services.
        /// Can be used with either RATE or UTILIZATION balancing modes,
        /// but required if RATE mode. Either maxRate or one
        /// of maxRatePerInstance or maxRatePerEndpoint, as appropriate for
        /// group type, must be set.
        /// </summary>
        public readonly int? MaxRate;
        /// <summary>
        /// The max requests per second (RPS) that a single backend network
        /// endpoint can handle. This is used to calculate the capacity of
        /// the group. Can be used in either balancing mode. For RATE mode,
        /// either maxRate or maxRatePerEndpoint must be set. Cannot be set
        /// for INTERNAL backend services.
        /// </summary>
        public readonly double? MaxRatePerEndpoint;
        /// <summary>
        /// The max requests per second (RPS) that a single backend
        /// instance can handle. This is used to calculate the capacity of
        /// the group. Can be used in either balancing mode. For RATE mode,
        /// either maxRate or maxRatePerInstance must be set. Cannot be set
        /// for INTERNAL backend services.
        /// </summary>
        public readonly double? MaxRatePerInstance;
        /// <summary>
        /// Used when balancingMode is UTILIZATION. This ratio defines the
        /// CPU utilization target for the group. Valid range is [0.0, 1.0].
        /// Cannot be set for INTERNAL backend services.
        /// </summary>
        public readonly double? MaxUtilization;

        [OutputConstructor]
        private RegionBackendServiceBackend(
            string? balancingMode,

            double? capacityScaler,

            string? description,

            bool? failover,

            string group,

            int? maxConnections,

            int? maxConnectionsPerEndpoint,

            int? maxConnectionsPerInstance,

            int? maxRate,

            double? maxRatePerEndpoint,

            double? maxRatePerInstance,

            double? maxUtilization)
        {
            BalancingMode = balancingMode;
            CapacityScaler = capacityScaler;
            Description = description;
            Failover = failover;
            Group = group;
            MaxConnections = maxConnections;
            MaxConnectionsPerEndpoint = maxConnectionsPerEndpoint;
            MaxConnectionsPerInstance = maxConnectionsPerInstance;
            MaxRate = maxRate;
            MaxRatePerEndpoint = maxRatePerEndpoint;
            MaxRatePerInstance = maxRatePerInstance;
            MaxUtilization = maxUtilization;
        }
    }
}
