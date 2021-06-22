// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Dataproc.Inputs
{

    public sealed class WorkflowTemplateJobSparkJobLoggingConfigGetArgs : Pulumi.ResourceArgs
    {
        [Input("driverLogLevels")]
        private InputMap<string>? _driverLogLevels;

        /// <summary>
        /// The per-package log levels for the driver. This may include "root" package name to configure rootLogger. Examples: 'com.google = FATAL', 'root = INFO', 'org.apache = DEBUG'
        /// </summary>
        public InputMap<string> DriverLogLevels
        {
            get => _driverLogLevels ?? (_driverLogLevels = new InputMap<string>());
            set => _driverLogLevels = value;
        }

        public WorkflowTemplateJobSparkJobLoggingConfigGetArgs()
        {
        }
    }
}
