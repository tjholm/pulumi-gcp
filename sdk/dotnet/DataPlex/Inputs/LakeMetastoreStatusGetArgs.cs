// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.DataPlex.Inputs
{

    public sealed class LakeMetastoreStatusGetArgs : Pulumi.ResourceArgs
    {
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        [Input("message")]
        public Input<string>? Message { get; set; }

        [Input("state")]
        public Input<string>? State { get; set; }

        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        public LakeMetastoreStatusGetArgs()
        {
        }
    }
}
