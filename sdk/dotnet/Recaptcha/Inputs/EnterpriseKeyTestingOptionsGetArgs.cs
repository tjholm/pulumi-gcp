// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Recaptcha.Inputs
{

    public sealed class EnterpriseKeyTestingOptionsGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// For challenge-based keys only (CHECKBOX, INVISIBLE), all challenge requests for this site will return nocaptcha if NOCAPTCHA, or an unsolvable challenge if UNSOLVABLE_CHALLENGE. Possible values: TESTING_CHALLENGE_UNSPECIFIED, NOCAPTCHA, UNSOLVABLE_CHALLENGE
        /// </summary>
        [Input("testingChallenge")]
        public Input<string>? TestingChallenge { get; set; }

        /// <summary>
        /// All assessments for this Key will return this score. Must be between 0 (likely not legitimate) and 1 (likely legitimate) inclusive.
        /// </summary>
        [Input("testingScore")]
        public Input<double>? TestingScore { get; set; }

        public EnterpriseKeyTestingOptionsGetArgs()
        {
        }
    }
}
