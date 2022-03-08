// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Dataproc.Inputs
{

    public sealed class WorkflowTemplatePlacementManagedClusterConfigInitializationActionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. Cloud Storage URI of executable file.
        /// </summary>
        [Input("executableFile", required: true)]
        public Input<string> ExecutableFile { get; set; } = null!;

        /// <summary>
        /// Optional. Amount of time executable has to complete. Default is 10 minutes (see JSON representation of (https://developers.google.com/protocol-buffers/docs/proto3#json)). Cluster creation fails with an explanatory error message (the name of the executable that caused the error and the exceeded timeout period) if the executable is not completed at end of the timeout period.
        /// </summary>
        [Input("executionTimeout")]
        public Input<string>? ExecutionTimeout { get; set; }

        public WorkflowTemplatePlacementManagedClusterConfigInitializationActionArgs()
        {
        }
    }
}
