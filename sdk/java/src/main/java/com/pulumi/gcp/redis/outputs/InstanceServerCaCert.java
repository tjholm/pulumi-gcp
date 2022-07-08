// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.redis.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class InstanceServerCaCert {
    private final @Nullable String cert;
    /**
     * @return -
     * Output only. The time when the policy was created.
     * A timestamp in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond
     * resolution and up to nine fractional digits.
     * 
     */
    private final @Nullable String createTime;
    private final @Nullable String expireTime;
    private final @Nullable String serialNumber;
    private final @Nullable String sha1Fingerprint;

    @CustomType.Constructor
    private InstanceServerCaCert(
        @CustomType.Parameter("cert") @Nullable String cert,
        @CustomType.Parameter("createTime") @Nullable String createTime,
        @CustomType.Parameter("expireTime") @Nullable String expireTime,
        @CustomType.Parameter("serialNumber") @Nullable String serialNumber,
        @CustomType.Parameter("sha1Fingerprint") @Nullable String sha1Fingerprint) {
        this.cert = cert;
        this.createTime = createTime;
        this.expireTime = expireTime;
        this.serialNumber = serialNumber;
        this.sha1Fingerprint = sha1Fingerprint;
    }

    public Optional<String> cert() {
        return Optional.ofNullable(this.cert);
    }
    /**
     * @return -
     * Output only. The time when the policy was created.
     * A timestamp in RFC3339 UTC &#34;Zulu&#34; format, with nanosecond
     * resolution and up to nine fractional digits.
     * 
     */
    public Optional<String> createTime() {
        return Optional.ofNullable(this.createTime);
    }
    public Optional<String> expireTime() {
        return Optional.ofNullable(this.expireTime);
    }
    public Optional<String> serialNumber() {
        return Optional.ofNullable(this.serialNumber);
    }
    public Optional<String> sha1Fingerprint() {
        return Optional.ofNullable(this.sha1Fingerprint);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(InstanceServerCaCert defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable String cert;
        private @Nullable String createTime;
        private @Nullable String expireTime;
        private @Nullable String serialNumber;
        private @Nullable String sha1Fingerprint;

        public Builder() {
    	      // Empty
        }

        public Builder(InstanceServerCaCert defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.cert = defaults.cert;
    	      this.createTime = defaults.createTime;
    	      this.expireTime = defaults.expireTime;
    	      this.serialNumber = defaults.serialNumber;
    	      this.sha1Fingerprint = defaults.sha1Fingerprint;
        }

        public Builder cert(@Nullable String cert) {
            this.cert = cert;
            return this;
        }
        public Builder createTime(@Nullable String createTime) {
            this.createTime = createTime;
            return this;
        }
        public Builder expireTime(@Nullable String expireTime) {
            this.expireTime = expireTime;
            return this;
        }
        public Builder serialNumber(@Nullable String serialNumber) {
            this.serialNumber = serialNumber;
            return this;
        }
        public Builder sha1Fingerprint(@Nullable String sha1Fingerprint) {
            this.sha1Fingerprint = sha1Fingerprint;
            return this;
        }        public InstanceServerCaCert build() {
            return new InstanceServerCaCert(cert, createTime, expireTime, serialNumber, sha1Fingerprint);
        }
    }
}
