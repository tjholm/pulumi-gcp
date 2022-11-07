// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Dns.Inputs
{

    public sealed class RecordSetRoutingPolicyGeoArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// For A and AAAA types only. The list of targets to be health checked. These can be specified along with `rrdatas` within this item.
        /// Structure is document below.
        /// </summary>
        [Input("healthCheckedTargets")]
        public Input<Inputs.RecordSetRoutingPolicyGeoHealthCheckedTargetsArgs>? HealthCheckedTargets { get; set; }

        /// <summary>
        /// The location name defined in Google Cloud.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("rrdatas")]
        private InputList<string>? _rrdatas;

        /// <summary>
        /// Same as `rrdatas` above.
        /// </summary>
        public InputList<string> Rrdatas
        {
            get => _rrdatas ?? (_rrdatas = new InputList<string>());
            set => _rrdatas = value;
        }

        public RecordSetRoutingPolicyGeoArgs()
        {
        }
        public static new RecordSetRoutingPolicyGeoArgs Empty => new RecordSetRoutingPolicyGeoArgs();
    }
}
