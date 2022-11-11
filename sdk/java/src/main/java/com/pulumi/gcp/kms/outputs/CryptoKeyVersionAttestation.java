// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.kms.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.gcp.kms.outputs.CryptoKeyVersionAttestationCertChains;
import com.pulumi.gcp.kms.outputs.CryptoKeyVersionAttestationExternalProtectionLevelOptions;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class CryptoKeyVersionAttestation {
    private @Nullable CryptoKeyVersionAttestationCertChains certChains;
    private @Nullable String content;
    private @Nullable CryptoKeyVersionAttestationExternalProtectionLevelOptions externalProtectionLevelOptions;
    private @Nullable String format;

    private CryptoKeyVersionAttestation() {}
    public Optional<CryptoKeyVersionAttestationCertChains> certChains() {
        return Optional.ofNullable(this.certChains);
    }
    public Optional<String> content() {
        return Optional.ofNullable(this.content);
    }
    public Optional<CryptoKeyVersionAttestationExternalProtectionLevelOptions> externalProtectionLevelOptions() {
        return Optional.ofNullable(this.externalProtectionLevelOptions);
    }
    public Optional<String> format() {
        return Optional.ofNullable(this.format);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(CryptoKeyVersionAttestation defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable CryptoKeyVersionAttestationCertChains certChains;
        private @Nullable String content;
        private @Nullable CryptoKeyVersionAttestationExternalProtectionLevelOptions externalProtectionLevelOptions;
        private @Nullable String format;
        public Builder() {}
        public Builder(CryptoKeyVersionAttestation defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.certChains = defaults.certChains;
    	      this.content = defaults.content;
    	      this.externalProtectionLevelOptions = defaults.externalProtectionLevelOptions;
    	      this.format = defaults.format;
        }

        @CustomType.Setter
        public Builder certChains(@Nullable CryptoKeyVersionAttestationCertChains certChains) {
            this.certChains = certChains;
            return this;
        }
        @CustomType.Setter
        public Builder content(@Nullable String content) {
            this.content = content;
            return this;
        }
        @CustomType.Setter
        public Builder externalProtectionLevelOptions(@Nullable CryptoKeyVersionAttestationExternalProtectionLevelOptions externalProtectionLevelOptions) {
            this.externalProtectionLevelOptions = externalProtectionLevelOptions;
            return this;
        }
        @CustomType.Setter
        public Builder format(@Nullable String format) {
            this.format = format;
            return this;
        }
        public CryptoKeyVersionAttestation build() {
            final var o = new CryptoKeyVersionAttestation();
            o.certChains = certChains;
            o.content = content;
            o.externalProtectionLevelOptions = externalProtectionLevelOptions;
            o.format = format;
            return o;
        }
    }
}
