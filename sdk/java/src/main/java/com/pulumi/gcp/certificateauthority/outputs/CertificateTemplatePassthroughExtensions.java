// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.certificateauthority.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.gcp.certificateauthority.outputs.CertificateTemplatePassthroughExtensionsAdditionalExtension;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class CertificateTemplatePassthroughExtensions {
    /**
     * @return Optional. A set of ObjectIds identifying custom X.509 extensions. Will be combined with known_extensions to determine the full set of X.509 extensions.
     * Structure is documented below.
     * 
     */
    private @Nullable List<CertificateTemplatePassthroughExtensionsAdditionalExtension> additionalExtensions;
    /**
     * @return Optional. A set of named X.509 extensions. Will be combined with additional_extensions to determine the full set of X.509 extensions.
     * 
     */
    private @Nullable List<String> knownExtensions;

    private CertificateTemplatePassthroughExtensions() {}
    /**
     * @return Optional. A set of ObjectIds identifying custom X.509 extensions. Will be combined with known_extensions to determine the full set of X.509 extensions.
     * Structure is documented below.
     * 
     */
    public List<CertificateTemplatePassthroughExtensionsAdditionalExtension> additionalExtensions() {
        return this.additionalExtensions == null ? List.of() : this.additionalExtensions;
    }
    /**
     * @return Optional. A set of named X.509 extensions. Will be combined with additional_extensions to determine the full set of X.509 extensions.
     * 
     */
    public List<String> knownExtensions() {
        return this.knownExtensions == null ? List.of() : this.knownExtensions;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(CertificateTemplatePassthroughExtensions defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<CertificateTemplatePassthroughExtensionsAdditionalExtension> additionalExtensions;
        private @Nullable List<String> knownExtensions;
        public Builder() {}
        public Builder(CertificateTemplatePassthroughExtensions defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.additionalExtensions = defaults.additionalExtensions;
    	      this.knownExtensions = defaults.knownExtensions;
        }

        @CustomType.Setter
        public Builder additionalExtensions(@Nullable List<CertificateTemplatePassthroughExtensionsAdditionalExtension> additionalExtensions) {

            this.additionalExtensions = additionalExtensions;
            return this;
        }
        public Builder additionalExtensions(CertificateTemplatePassthroughExtensionsAdditionalExtension... additionalExtensions) {
            return additionalExtensions(List.of(additionalExtensions));
        }
        @CustomType.Setter
        public Builder knownExtensions(@Nullable List<String> knownExtensions) {

            this.knownExtensions = knownExtensions;
            return this;
        }
        public Builder knownExtensions(String... knownExtensions) {
            return knownExtensions(List.of(knownExtensions));
        }
        public CertificateTemplatePassthroughExtensions build() {
            final var _resultValue = new CertificateTemplatePassthroughExtensions();
            _resultValue.additionalExtensions = additionalExtensions;
            _resultValue.knownExtensions = knownExtensions;
            return _resultValue;
        }
    }
}
