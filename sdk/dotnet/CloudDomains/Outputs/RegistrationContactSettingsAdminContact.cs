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
    public sealed class RegistrationContactSettingsAdminContact
    {
        /// <summary>
        /// Required. Email address of the contact.
        /// </summary>
        public readonly string Email;
        /// <summary>
        /// Fax number of the contact in international format. For example, "+1-800-555-0123".
        /// </summary>
        public readonly string? FaxNumber;
        /// <summary>
        /// Required. Phone number of the contact in international format. For example, "+1-800-555-0123".
        /// </summary>
        public readonly string PhoneNumber;
        /// <summary>
        /// Required. Postal address of the contact.
        /// Structure is documented below.
        /// </summary>
        public readonly Outputs.RegistrationContactSettingsAdminContactPostalAddress PostalAddress;

        [OutputConstructor]
        private RegistrationContactSettingsAdminContact(
            string email,

            string? faxNumber,

            string phoneNumber,

            Outputs.RegistrationContactSettingsAdminContactPostalAddress postalAddress)
        {
            Email = email;
            FaxNumber = faxNumber;
            PhoneNumber = phoneNumber;
            PostalAddress = postalAddress;
        }
    }
}
