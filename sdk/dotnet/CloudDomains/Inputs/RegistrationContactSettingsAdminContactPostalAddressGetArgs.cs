// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudDomains.Inputs
{

    public sealed class RegistrationContactSettingsAdminContactPostalAddressGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("addressLines")]
        private InputList<string>? _addressLines;

        /// <summary>
        /// Unstructured address lines describing the lower levels of an address.
        /// Because values in addressLines do not have type information and may sometimes contain multiple values in a single
        /// field (e.g. "Austin, TX"), it is important that the line order is clear. The order of address lines should be
        /// "envelope order" for the country/region of the address. In places where this can vary (e.g. Japan), address_language
        /// is used to make it explicit (e.g. "ja" for large-to-small ordering and "ja-Latn" or "en" for small-to-large). This way,
        /// the most specific line of an address can be selected based on the language.
        /// </summary>
        public InputList<string> AddressLines
        {
            get => _addressLines ?? (_addressLines = new InputList<string>());
            set => _addressLines = value;
        }

        /// <summary>
        /// Highest administrative subdivision which is used for postal addresses of a country or region. For example, this can be a state,
        /// a province, an oblast, or a prefecture. Specifically, for Spain this is the province and not the autonomous community
        /// (e.g. "Barcelona" and not "Catalonia"). Many countries don't use an administrative area in postal addresses. E.g. in Switzerland
        /// this should be left unpopulated.
        /// </summary>
        [Input("administrativeArea")]
        public Input<string>? AdministrativeArea { get; set; }

        /// <summary>
        /// Generally refers to the city/town portion of the address. Examples: US city, IT comune, UK post town. In regions of the world
        /// where localities are not well defined or do not fit into this structure well, leave locality empty and use addressLines.
        /// </summary>
        [Input("locality")]
        public Input<string>? Locality { get; set; }

        /// <summary>
        /// The name of the organization at the address.
        /// </summary>
        [Input("organization")]
        public Input<string>? Organization { get; set; }

        /// <summary>
        /// Postal code of the address. Not all countries use or require postal codes to be present, but where they are used,
        /// they may trigger additional validation with other parts of the address (e.g. state/zip validation in the U.S.A.).
        /// </summary>
        [Input("postalCode")]
        public Input<string>? PostalCode { get; set; }

        [Input("recipients")]
        private InputList<string>? _recipients;

        /// <summary>
        /// The recipient at the address. This field may, under certain circumstances, contain multiline information. For example,
        /// it might contain "care of" information.
        /// 
        /// - - -
        /// </summary>
        public InputList<string> Recipients
        {
            get => _recipients ?? (_recipients = new InputList<string>());
            set => _recipients = value;
        }

        /// <summary>
        /// Required. CLDR region code of the country/region of the address. This is never inferred and it is up to the user to
        /// ensure the value is correct. See https://cldr.unicode.org/ and
        /// https://www.unicode.org/cldr/charts/30/supplemental/territory_information.html for details. Example: "CH" for Switzerland.
        /// </summary>
        [Input("regionCode", required: true)]
        public Input<string> RegionCode { get; set; } = null!;

        public RegistrationContactSettingsAdminContactPostalAddressGetArgs()
        {
        }
        public static new RegistrationContactSettingsAdminContactPostalAddressGetArgs Empty => new RegistrationContactSettingsAdminContactPostalAddressGetArgs();
    }
}
