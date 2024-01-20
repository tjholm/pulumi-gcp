// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudDomains.Inputs
{

    public sealed class RegistrationYearlyPriceGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The three-letter currency code defined in ISO 4217.
        /// </summary>
        [Input("currencyCode")]
        public Input<string>? CurrencyCode { get; set; }

        /// <summary>
        /// The whole units of the amount. For example if currencyCode is "USD", then 1 unit is one US dollar.
        /// </summary>
        [Input("units")]
        public Input<string>? Units { get; set; }

        public RegistrationYearlyPriceGetArgs()
        {
        }
        public static new RegistrationYearlyPriceGetArgs Empty => new RegistrationYearlyPriceGetArgs();
    }
}
