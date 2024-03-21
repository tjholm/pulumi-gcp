// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Apphub.Inputs
{

    public sealed class ApplicationAttributesEnvironmentGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Environment type.
        /// Possible values are: `PRODUCTION`, `STAGING`, `TEST`, `DEVELOPMENT`.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ApplicationAttributesEnvironmentGetArgs()
        {
        }
        public static new ApplicationAttributesEnvironmentGetArgs Empty => new ApplicationAttributesEnvironmentGetArgs();
    }
}
