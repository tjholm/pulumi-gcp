// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Workstations.Inputs
{

    public sealed class WorkstationConfigPersistentDirectoryGcePdGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of the disk to use.
        /// </summary>
        [Input("diskType")]
        public Input<string>? DiskType { get; set; }

        /// <summary>
        /// Type of file system that the disk should be formatted with. The workstation image must support this file system type. Must be empty if sourceSnapshot is set.
        /// </summary>
        [Input("fsType")]
        public Input<string>? FsType { get; set; }

        /// <summary>
        /// What should happen to the disk after the workstation is deleted. Defaults to DELETE.
        /// Possible values are: `DELETE`, `RETAIN`.
        /// </summary>
        [Input("reclaimPolicy")]
        public Input<string>? ReclaimPolicy { get; set; }

        /// <summary>
        /// Size of the disk in GB. Must be empty if sourceSnapshot is set.
        /// </summary>
        [Input("sizeGb")]
        public Input<int>? SizeGb { get; set; }

        public WorkstationConfigPersistentDirectoryGcePdGetArgs()
        {
        }
        public static new WorkstationConfigPersistentDirectoryGcePdGetArgs Empty => new WorkstationConfigPersistentDirectoryGcePdGetArgs();
    }
}
