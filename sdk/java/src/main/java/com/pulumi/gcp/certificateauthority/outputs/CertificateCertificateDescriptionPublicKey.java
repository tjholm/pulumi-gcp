// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.certificateauthority.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class CertificateCertificateDescriptionPublicKey {
    /**
     * @return The format of the public key. Currently, only PEM format is supported.
     * Possible values are `KEY_TYPE_UNSPECIFIED` and `PEM`.
     * 
     */
    private final @Nullable String format;
    /**
     * @return Required. A public key. When this is specified in a request, the padding and encoding can be any of the options described by the respective &#39;KeyType&#39; value. When this is generated by the service, it will always be an RFC 5280 SubjectPublicKeyInfo structure containing an algorithm identifier and a key. A base64-encoded string.
     * 
     */
    private final @Nullable String key;

    @CustomType.Constructor
    private CertificateCertificateDescriptionPublicKey(
        @CustomType.Parameter("format") @Nullable String format,
        @CustomType.Parameter("key") @Nullable String key) {
        this.format = format;
        this.key = key;
    }

    /**
     * @return The format of the public key. Currently, only PEM format is supported.
     * Possible values are `KEY_TYPE_UNSPECIFIED` and `PEM`.
     * 
     */
    public Optional<String> format() {
        return Optional.ofNullable(this.format);
    }
    /**
     * @return Required. A public key. When this is specified in a request, the padding and encoding can be any of the options described by the respective &#39;KeyType&#39; value. When this is generated by the service, it will always be an RFC 5280 SubjectPublicKeyInfo structure containing an algorithm identifier and a key. A base64-encoded string.
     * 
     */
    public Optional<String> key() {
        return Optional.ofNullable(this.key);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(CertificateCertificateDescriptionPublicKey defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable String format;
        private @Nullable String key;

        public Builder() {
    	      // Empty
        }

        public Builder(CertificateCertificateDescriptionPublicKey defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.format = defaults.format;
    	      this.key = defaults.key;
        }

        public Builder format(@Nullable String format) {
            this.format = format;
            return this;
        }
        public Builder key(@Nullable String key) {
            this.key = key;
            return this;
        }        public CertificateCertificateDescriptionPublicKey build() {
            return new CertificateCertificateDescriptionPublicKey(format, key);
        }
    }
}
