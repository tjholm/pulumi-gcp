// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Workstations.Outputs
{

    [OutputType]
    public sealed class WorkstationConfigReadinessCheck
    {
        /// <summary>
        /// Path to which the request should be sent.
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// Port to which the request should be sent.
        /// </summary>
        public readonly int Port;

        [OutputConstructor]
        private WorkstationConfigReadinessCheck(
            string path,

            int port)
        {
            Path = path;
            Port = port;
        }
    }
}
