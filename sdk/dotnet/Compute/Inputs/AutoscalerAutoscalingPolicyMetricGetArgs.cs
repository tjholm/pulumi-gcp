// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class AutoscalerAutoscalingPolicyMetricGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("filter")]
        public Input<string>? Filter { get; set; }

        /// <summary>
        /// The identifier (type) of the Stackdriver Monitoring metric.
        /// The metric cannot have negative values.
        /// The metric must have a value type of INT64 or DOUBLE.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("singleInstanceAssignment")]
        public Input<double>? SingleInstanceAssignment { get; set; }

        /// <summary>
        /// The target value of the metric that autoscaler should
        /// maintain. This must be a positive value. A utilization
        /// metric scales number of virtual machines handling requests
        /// to increase or decrease proportionally to the metric.
        /// For example, a good metric to use as a utilizationTarget is
        /// www.googleapis.com/compute/instance/network/received_bytes_count.
        /// The autoscaler will work to keep this value constant for each
        /// of the instances.
        /// </summary>
        [Input("target")]
        public Input<double>? Target { get; set; }

        /// <summary>
        /// Defines how target utilization value is expressed for a
        /// Stackdriver Monitoring metric.
        /// Possible values are `GAUGE`, `DELTA_PER_SECOND`, and `DELTA_PER_MINUTE`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public AutoscalerAutoscalingPolicyMetricGetArgs()
        {
        }
        public static new AutoscalerAutoscalingPolicyMetricGetArgs Empty => new AutoscalerAutoscalingPolicyMetricGetArgs();
    }
}
