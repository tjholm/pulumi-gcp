// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Vertex.Inputs
{

    public sealed class AiEndpointDeployedModelPrivateEndpointArgs : global::Pulumi.ResourceArgs
    {
        [Input("explainHttpUri")]
        public Input<string>? ExplainHttpUri { get; set; }

        [Input("healthHttpUri")]
        public Input<string>? HealthHttpUri { get; set; }

        [Input("predictHttpUri")]
        public Input<string>? PredictHttpUri { get; set; }

        [Input("serviceAttachment")]
        public Input<string>? ServiceAttachment { get; set; }

        public AiEndpointDeployedModelPrivateEndpointArgs()
        {
        }
        public static new AiEndpointDeployedModelPrivateEndpointArgs Empty => new AiEndpointDeployedModelPrivateEndpointArgs();
    }
}
