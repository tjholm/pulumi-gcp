// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Apphub.Inputs
{

    public sealed class ServiceAttributesArgs : global::Pulumi.ResourceArgs
    {
        [Input("businessOwners")]
        private InputList<Inputs.ServiceAttributesBusinessOwnerArgs>? _businessOwners;

        /// <summary>
        /// Business team that ensures user needs are met and value is delivered
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.ServiceAttributesBusinessOwnerArgs> BusinessOwners
        {
            get => _businessOwners ?? (_businessOwners = new InputList<Inputs.ServiceAttributesBusinessOwnerArgs>());
            set => _businessOwners = value;
        }

        /// <summary>
        /// Criticality of the Application, Service, or Workload
        /// Structure is documented below.
        /// </summary>
        [Input("criticality")]
        public Input<Inputs.ServiceAttributesCriticalityArgs>? Criticality { get; set; }

        [Input("developerOwners")]
        private InputList<Inputs.ServiceAttributesDeveloperOwnerArgs>? _developerOwners;

        /// <summary>
        /// Developer team that owns development and coding.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.ServiceAttributesDeveloperOwnerArgs> DeveloperOwners
        {
            get => _developerOwners ?? (_developerOwners = new InputList<Inputs.ServiceAttributesDeveloperOwnerArgs>());
            set => _developerOwners = value;
        }

        /// <summary>
        /// Environment of the Application, Service, or Workload
        /// Structure is documented below.
        /// </summary>
        [Input("environment")]
        public Input<Inputs.ServiceAttributesEnvironmentArgs>? Environment { get; set; }

        [Input("operatorOwners")]
        private InputList<Inputs.ServiceAttributesOperatorOwnerArgs>? _operatorOwners;

        /// <summary>
        /// Operator team that ensures runtime and operations.
        /// Structure is documented below.
        /// </summary>
        public InputList<Inputs.ServiceAttributesOperatorOwnerArgs> OperatorOwners
        {
            get => _operatorOwners ?? (_operatorOwners = new InputList<Inputs.ServiceAttributesOperatorOwnerArgs>());
            set => _operatorOwners = value;
        }

        public ServiceAttributesArgs()
        {
        }
        public static new ServiceAttributesArgs Empty => new ServiceAttributesArgs();
    }
}
