// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Dataproc.Inputs
{

    public sealed class JobPrestoConfigLoggingConfigGetArgs : Pulumi.ResourceArgs
    {
        [Input("driverLogLevels", required: true)]
        private InputMap<string>? _driverLogLevels;
        public InputMap<string> DriverLogLevels
        {
            get => _driverLogLevels ?? (_driverLogLevels = new InputMap<string>());
            set => _driverLogLevels = value;
        }

        public JobPrestoConfigLoggingConfigGetArgs()
        {
        }
    }
}
