// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudDomains.Outputs
{

    [OutputType]
    public sealed class RegistrationYearlyPrice
    {
        /// <summary>
        /// The three-letter currency code defined in ISO 4217.
        /// </summary>
        public readonly string? CurrencyCode;
        /// <summary>
        /// The whole units of the amount. For example if currencyCode is "USD", then 1 unit is one US dollar.
        /// </summary>
        public readonly string? Units;

        [OutputConstructor]
        private RegistrationYearlyPrice(
            string? currencyCode,

            string? units)
        {
            CurrencyCode = currencyCode;
            Units = units;
        }
    }
}
