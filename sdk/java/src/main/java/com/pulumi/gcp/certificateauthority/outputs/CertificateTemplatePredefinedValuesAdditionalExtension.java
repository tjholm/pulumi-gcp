// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.certificateauthority.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.gcp.certificateauthority.outputs.CertificateTemplatePredefinedValuesAdditionalExtensionObjectId;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class CertificateTemplatePredefinedValuesAdditionalExtension {
    /**
     * @return Optional. Indicates whether or not this extension is critical (i.e., if the client does not know how to handle this extension, the client should consider this to be an error).
     * 
     */
    private @Nullable Boolean critical;
    /**
     * @return Required. The OID for this X.509 extension.
     * Structure is documented below.
     * 
     */
    private CertificateTemplatePredefinedValuesAdditionalExtensionObjectId objectId;
    /**
     * @return Required. The value of this X.509 extension.
     * 
     */
    private String value;

    private CertificateTemplatePredefinedValuesAdditionalExtension() {}
    /**
     * @return Optional. Indicates whether or not this extension is critical (i.e., if the client does not know how to handle this extension, the client should consider this to be an error).
     * 
     */
    public Optional<Boolean> critical() {
        return Optional.ofNullable(this.critical);
    }
    /**
     * @return Required. The OID for this X.509 extension.
     * Structure is documented below.
     * 
     */
    public CertificateTemplatePredefinedValuesAdditionalExtensionObjectId objectId() {
        return this.objectId;
    }
    /**
     * @return Required. The value of this X.509 extension.
     * 
     */
    public String value() {
        return this.value;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(CertificateTemplatePredefinedValuesAdditionalExtension defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Boolean critical;
        private CertificateTemplatePredefinedValuesAdditionalExtensionObjectId objectId;
        private String value;
        public Builder() {}
        public Builder(CertificateTemplatePredefinedValuesAdditionalExtension defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.critical = defaults.critical;
    	      this.objectId = defaults.objectId;
    	      this.value = defaults.value;
        }

        @CustomType.Setter
        public Builder critical(@Nullable Boolean critical) {

            this.critical = critical;
            return this;
        }
        @CustomType.Setter
        public Builder objectId(CertificateTemplatePredefinedValuesAdditionalExtensionObjectId objectId) {
            if (objectId == null) {
              throw new MissingRequiredPropertyException("CertificateTemplatePredefinedValuesAdditionalExtension", "objectId");
            }
            this.objectId = objectId;
            return this;
        }
        @CustomType.Setter
        public Builder value(String value) {
            if (value == null) {
              throw new MissingRequiredPropertyException("CertificateTemplatePredefinedValuesAdditionalExtension", "value");
            }
            this.value = value;
            return this;
        }
        public CertificateTemplatePredefinedValuesAdditionalExtension build() {
            final var _resultValue = new CertificateTemplatePredefinedValuesAdditionalExtension();
            _resultValue.critical = critical;
            _resultValue.objectId = objectId;
            _resultValue.value = value;
            return _resultValue;
        }
    }
}
