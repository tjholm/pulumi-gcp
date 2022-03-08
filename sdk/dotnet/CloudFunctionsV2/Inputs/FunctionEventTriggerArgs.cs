// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudFunctionsV2.Inputs
{

    public sealed class FunctionEventTriggerArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. The type of event to observe.
        /// </summary>
        [Input("eventType")]
        public Input<string>? EventType { get; set; }

        /// <summary>
        /// The name of a Pub/Sub topic in the same project that will be used
        /// as the transport topic for the event delivery.
        /// </summary>
        [Input("pubsubTopic")]
        public Input<string>? PubsubTopic { get; set; }

        /// <summary>
        /// Describes the retry policy in case of function's execution failure.
        /// Retried execution is charged as any other execution.
        /// Possible values are `RETRY_POLICY_UNSPECIFIED`, `RETRY_POLICY_DO_NOT_RETRY`, and `RETRY_POLICY_RETRY`.
        /// </summary>
        [Input("retryPolicy")]
        public Input<string>? RetryPolicy { get; set; }

        /// <summary>
        /// -
        /// The email of the service account for this function.
        /// </summary>
        [Input("serviceAccountEmail")]
        public Input<string>? ServiceAccountEmail { get; set; }

        /// <summary>
        /// -
        /// The resource name of the Eventarc trigger.
        /// </summary>
        [Input("trigger")]
        public Input<string>? Trigger { get; set; }

        /// <summary>
        /// The region that the trigger will be in. The trigger will only receive
        /// events originating in this region. It can be the same
        /// region as the function, a different region or multi-region, or the global
        /// region. If not provided, defaults to the same region as the function.
        /// </summary>
        [Input("triggerRegion")]
        public Input<string>? TriggerRegion { get; set; }

        public FunctionEventTriggerArgs()
        {
        }
    }
}
