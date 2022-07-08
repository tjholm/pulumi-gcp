// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.certificateauthority.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class CertificateCertificateDescriptionSubjectKeyId {
    private final @Nullable String keyId;

    @CustomType.Constructor
    private CertificateCertificateDescriptionSubjectKeyId(@CustomType.Parameter("keyId") @Nullable String keyId) {
        this.keyId = keyId;
    }

    public Optional<String> keyId() {
        return Optional.ofNullable(this.keyId);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(CertificateCertificateDescriptionSubjectKeyId defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable String keyId;

        public Builder() {
    	      // Empty
        }

        public Builder(CertificateCertificateDescriptionSubjectKeyId defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.keyId = defaults.keyId;
        }

        public Builder keyId(@Nullable String keyId) {
            this.keyId = keyId;
            return this;
        }        public CertificateCertificateDescriptionSubjectKeyId build() {
            return new CertificateCertificateDescriptionSubjectKeyId(keyId);
        }
    }
}
