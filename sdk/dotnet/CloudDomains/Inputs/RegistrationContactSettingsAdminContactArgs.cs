// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudDomains.Inputs
{

    public sealed class RegistrationContactSettingsAdminContactArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. Email address of the contact.
        /// </summary>
        [Input("email", required: true)]
        public Input<string> Email { get; set; } = null!;

        /// <summary>
        /// Fax number of the contact in international format. For example, "+1-800-555-0123".
        /// </summary>
        [Input("faxNumber")]
        public Input<string>? FaxNumber { get; set; }

        /// <summary>
        /// Required. Phone number of the contact in international format. For example, "+1-800-555-0123".
        /// </summary>
        [Input("phoneNumber", required: true)]
        public Input<string> PhoneNumber { get; set; } = null!;

        /// <summary>
        /// Required. Postal address of the contact.
        /// Structure is documented below.
        /// </summary>
        [Input("postalAddress", required: true)]
        public Input<Inputs.RegistrationContactSettingsAdminContactPostalAddressArgs> PostalAddress { get; set; } = null!;

        public RegistrationContactSettingsAdminContactArgs()
        {
        }
        public static new RegistrationContactSettingsAdminContactArgs Empty => new RegistrationContactSettingsAdminContactArgs();
    }
}
