// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Billing.Inputs
{

    public sealed class BudgetBudgetFilterCustomPeriodStartDateGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Day of a month. Must be from 1 to 31 and valid for the year and month.
        /// </summary>
        [Input("day", required: true)]
        public Input<int> Day { get; set; } = null!;

        /// <summary>
        /// Month of a year. Must be from 1 to 12.
        /// </summary>
        [Input("month", required: true)]
        public Input<int> Month { get; set; } = null!;

        /// <summary>
        /// Year of the date. Must be from 1 to 9999.
        /// </summary>
        [Input("year", required: true)]
        public Input<int> Year { get; set; } = null!;

        public BudgetBudgetFilterCustomPeriodStartDateGetArgs()
        {
        }
    }
}
