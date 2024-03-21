// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Workstations.Inputs
{

    public sealed class WorkstationConfigEphemeralDirectoryGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An EphemeralDirectory backed by a Compute Engine persistent disk.
        /// Structure is documented below.
        /// </summary>
        [Input("gcePd")]
        public Input<Inputs.WorkstationConfigEphemeralDirectoryGcePdGetArgs>? GcePd { get; set; }

        /// <summary>
        /// Location of this directory in the running workstation.
        /// </summary>
        [Input("mountPath")]
        public Input<string>? MountPath { get; set; }

        public WorkstationConfigEphemeralDirectoryGetArgs()
        {
        }
        public static new WorkstationConfigEphemeralDirectoryGetArgs Empty => new WorkstationConfigEphemeralDirectoryGetArgs();
    }
}
