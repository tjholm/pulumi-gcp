// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.alloydb.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ClusterEncryptionInfo {
    /**
     * @return (Output)
     * Output only. Type of encryption.
     * 
     */
    private @Nullable String encryptionType;
    /**
     * @return (Output)
     * Output only. Cloud KMS key versions that are being used to protect the database or the backup.
     * 
     */
    private @Nullable List<String> kmsKeyVersions;

    private ClusterEncryptionInfo() {}
    /**
     * @return (Output)
     * Output only. Type of encryption.
     * 
     */
    public Optional<String> encryptionType() {
        return Optional.ofNullable(this.encryptionType);
    }
    /**
     * @return (Output)
     * Output only. Cloud KMS key versions that are being used to protect the database or the backup.
     * 
     */
    public List<String> kmsKeyVersions() {
        return this.kmsKeyVersions == null ? List.of() : this.kmsKeyVersions;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ClusterEncryptionInfo defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String encryptionType;
        private @Nullable List<String> kmsKeyVersions;
        public Builder() {}
        public Builder(ClusterEncryptionInfo defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.encryptionType = defaults.encryptionType;
    	      this.kmsKeyVersions = defaults.kmsKeyVersions;
        }

        @CustomType.Setter
        public Builder encryptionType(@Nullable String encryptionType) {
            this.encryptionType = encryptionType;
            return this;
        }
        @CustomType.Setter
        public Builder kmsKeyVersions(@Nullable List<String> kmsKeyVersions) {
            this.kmsKeyVersions = kmsKeyVersions;
            return this;
        }
        public Builder kmsKeyVersions(String... kmsKeyVersions) {
            return kmsKeyVersions(List.of(kmsKeyVersions));
        }
        public ClusterEncryptionInfo build() {
            final var o = new ClusterEncryptionInfo();
            o.encryptionType = encryptionType;
            o.kmsKeyVersions = kmsKeyVersions;
            return o;
        }
    }
}
