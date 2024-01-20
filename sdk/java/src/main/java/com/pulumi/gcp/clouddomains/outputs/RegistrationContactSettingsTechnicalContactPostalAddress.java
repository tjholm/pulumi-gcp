// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.clouddomains.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class RegistrationContactSettingsTechnicalContactPostalAddress {
    /**
     * @return Unstructured address lines describing the lower levels of an address.
     * Because values in addressLines do not have type information and may sometimes contain multiple values in a single
     * field (e.g. &#34;Austin, TX&#34;), it is important that the line order is clear. The order of address lines should be
     * &#34;envelope order&#34; for the country/region of the address. In places where this can vary (e.g. Japan), address_language
     * is used to make it explicit (e.g. &#34;ja&#34; for large-to-small ordering and &#34;ja-Latn&#34; or &#34;en&#34; for small-to-large). This way,
     * the most specific line of an address can be selected based on the language.
     * 
     */
    private @Nullable List<String> addressLines;
    /**
     * @return Highest administrative subdivision which is used for postal addresses of a country or region. For example, this can be a state,
     * a province, an oblast, or a prefecture. Specifically, for Spain this is the province and not the autonomous community
     * (e.g. &#34;Barcelona&#34; and not &#34;Catalonia&#34;). Many countries don&#39;t use an administrative area in postal addresses. E.g. in Switzerland
     * this should be left unpopulated.
     * 
     */
    private @Nullable String administrativeArea;
    /**
     * @return Generally refers to the city/town portion of the address. Examples: US city, IT comune, UK post town. In regions of the world
     * where localities are not well defined or do not fit into this structure well, leave locality empty and use addressLines.
     * 
     */
    private @Nullable String locality;
    /**
     * @return The name of the organization at the address.
     * 
     */
    private @Nullable String organization;
    /**
     * @return Postal code of the address. Not all countries use or require postal codes to be present, but where they are used,
     * they may trigger additional validation with other parts of the address (e.g. state/zip validation in the U.S.A.).
     * 
     */
    private @Nullable String postalCode;
    /**
     * @return The recipient at the address. This field may, under certain circumstances, contain multiline information. For example,
     * it might contain &#34;care of&#34; information.
     * 
     * ***
     * 
     */
    private @Nullable List<String> recipients;
    /**
     * @return Required. CLDR region code of the country/region of the address. This is never inferred and it is up to the user to
     * ensure the value is correct. See https://cldr.unicode.org/ and
     * https://www.unicode.org/cldr/charts/30/supplemental/territory_information.html for details. Example: &#34;CH&#34; for Switzerland.
     * 
     */
    private String regionCode;

    private RegistrationContactSettingsTechnicalContactPostalAddress() {}
    /**
     * @return Unstructured address lines describing the lower levels of an address.
     * Because values in addressLines do not have type information and may sometimes contain multiple values in a single
     * field (e.g. &#34;Austin, TX&#34;), it is important that the line order is clear. The order of address lines should be
     * &#34;envelope order&#34; for the country/region of the address. In places where this can vary (e.g. Japan), address_language
     * is used to make it explicit (e.g. &#34;ja&#34; for large-to-small ordering and &#34;ja-Latn&#34; or &#34;en&#34; for small-to-large). This way,
     * the most specific line of an address can be selected based on the language.
     * 
     */
    public List<String> addressLines() {
        return this.addressLines == null ? List.of() : this.addressLines;
    }
    /**
     * @return Highest administrative subdivision which is used for postal addresses of a country or region. For example, this can be a state,
     * a province, an oblast, or a prefecture. Specifically, for Spain this is the province and not the autonomous community
     * (e.g. &#34;Barcelona&#34; and not &#34;Catalonia&#34;). Many countries don&#39;t use an administrative area in postal addresses. E.g. in Switzerland
     * this should be left unpopulated.
     * 
     */
    public Optional<String> administrativeArea() {
        return Optional.ofNullable(this.administrativeArea);
    }
    /**
     * @return Generally refers to the city/town portion of the address. Examples: US city, IT comune, UK post town. In regions of the world
     * where localities are not well defined or do not fit into this structure well, leave locality empty and use addressLines.
     * 
     */
    public Optional<String> locality() {
        return Optional.ofNullable(this.locality);
    }
    /**
     * @return The name of the organization at the address.
     * 
     */
    public Optional<String> organization() {
        return Optional.ofNullable(this.organization);
    }
    /**
     * @return Postal code of the address. Not all countries use or require postal codes to be present, but where they are used,
     * they may trigger additional validation with other parts of the address (e.g. state/zip validation in the U.S.A.).
     * 
     */
    public Optional<String> postalCode() {
        return Optional.ofNullable(this.postalCode);
    }
    /**
     * @return The recipient at the address. This field may, under certain circumstances, contain multiline information. For example,
     * it might contain &#34;care of&#34; information.
     * 
     * ***
     * 
     */
    public List<String> recipients() {
        return this.recipients == null ? List.of() : this.recipients;
    }
    /**
     * @return Required. CLDR region code of the country/region of the address. This is never inferred and it is up to the user to
     * ensure the value is correct. See https://cldr.unicode.org/ and
     * https://www.unicode.org/cldr/charts/30/supplemental/territory_information.html for details. Example: &#34;CH&#34; for Switzerland.
     * 
     */
    public String regionCode() {
        return this.regionCode;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RegistrationContactSettingsTechnicalContactPostalAddress defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> addressLines;
        private @Nullable String administrativeArea;
        private @Nullable String locality;
        private @Nullable String organization;
        private @Nullable String postalCode;
        private @Nullable List<String> recipients;
        private String regionCode;
        public Builder() {}
        public Builder(RegistrationContactSettingsTechnicalContactPostalAddress defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.addressLines = defaults.addressLines;
    	      this.administrativeArea = defaults.administrativeArea;
    	      this.locality = defaults.locality;
    	      this.organization = defaults.organization;
    	      this.postalCode = defaults.postalCode;
    	      this.recipients = defaults.recipients;
    	      this.regionCode = defaults.regionCode;
        }

        @CustomType.Setter
        public Builder addressLines(@Nullable List<String> addressLines) {

            this.addressLines = addressLines;
            return this;
        }
        public Builder addressLines(String... addressLines) {
            return addressLines(List.of(addressLines));
        }
        @CustomType.Setter
        public Builder administrativeArea(@Nullable String administrativeArea) {

            this.administrativeArea = administrativeArea;
            return this;
        }
        @CustomType.Setter
        public Builder locality(@Nullable String locality) {

            this.locality = locality;
            return this;
        }
        @CustomType.Setter
        public Builder organization(@Nullable String organization) {

            this.organization = organization;
            return this;
        }
        @CustomType.Setter
        public Builder postalCode(@Nullable String postalCode) {

            this.postalCode = postalCode;
            return this;
        }
        @CustomType.Setter
        public Builder recipients(@Nullable List<String> recipients) {

            this.recipients = recipients;
            return this;
        }
        public Builder recipients(String... recipients) {
            return recipients(List.of(recipients));
        }
        @CustomType.Setter
        public Builder regionCode(String regionCode) {
            if (regionCode == null) {
              throw new MissingRequiredPropertyException("RegistrationContactSettingsTechnicalContactPostalAddress", "regionCode");
            }
            this.regionCode = regionCode;
            return this;
        }
        public RegistrationContactSettingsTechnicalContactPostalAddress build() {
            final var _resultValue = new RegistrationContactSettingsTechnicalContactPostalAddress();
            _resultValue.addressLines = addressLines;
            _resultValue.administrativeArea = administrativeArea;
            _resultValue.locality = locality;
            _resultValue.organization = organization;
            _resultValue.postalCode = postalCode;
            _resultValue.recipients = recipients;
            _resultValue.regionCode = regionCode;
            return _resultValue;
        }
    }
}
