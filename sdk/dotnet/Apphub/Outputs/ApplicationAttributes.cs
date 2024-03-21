// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Apphub.Outputs
{

    [OutputType]
    public sealed class ApplicationAttributes
    {
        /// <summary>
        /// Optional. Business team that ensures user needs are met and value is delivered
        /// Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationAttributesBusinessOwner> BusinessOwners;
        /// <summary>
        /// Criticality of the Application, Service, or Workload
        /// Structure is documented below.
        /// </summary>
        public readonly Outputs.ApplicationAttributesCriticality? Criticality;
        /// <summary>
        /// Optional. Developer team that owns development and coding.
        /// Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationAttributesDeveloperOwner> DeveloperOwners;
        /// <summary>
        /// Environment of the Application, Service, or Workload
        /// Structure is documented below.
        /// </summary>
        public readonly Outputs.ApplicationAttributesEnvironment? Environment;
        /// <summary>
        /// Optional. Operator team that ensures runtime and operations.
        /// Structure is documented below.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApplicationAttributesOperatorOwner> OperatorOwners;

        [OutputConstructor]
        private ApplicationAttributes(
            ImmutableArray<Outputs.ApplicationAttributesBusinessOwner> businessOwners,

            Outputs.ApplicationAttributesCriticality? criticality,

            ImmutableArray<Outputs.ApplicationAttributesDeveloperOwner> developerOwners,

            Outputs.ApplicationAttributesEnvironment? environment,

            ImmutableArray<Outputs.ApplicationAttributesOperatorOwner> operatorOwners)
        {
            BusinessOwners = businessOwners;
            Criticality = criticality;
            DeveloperOwners = developerOwners;
            Environment = environment;
            OperatorOwners = operatorOwners;
        }
    }
}
