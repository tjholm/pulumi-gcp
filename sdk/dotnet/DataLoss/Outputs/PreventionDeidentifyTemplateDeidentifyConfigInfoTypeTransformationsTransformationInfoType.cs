// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.DataLoss.Outputs
{

    [OutputType]
    public sealed class PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationInfoType
    {
        /// <summary>
        /// Name of the information type.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Version name for this InfoType.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private PreventionDeidentifyTemplateDeidentifyConfigInfoTypeTransformationsTransformationInfoType(
            string name,

            string? version)
        {
            Name = name;
            Version = version;
        }
    }
}
