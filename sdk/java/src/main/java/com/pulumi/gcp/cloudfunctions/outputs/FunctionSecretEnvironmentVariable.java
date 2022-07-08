// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.cloudfunctions.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class FunctionSecretEnvironmentVariable {
    /**
     * @return Name of the environment variable.
     * 
     */
    private final String key;
    /**
     * @return Project identifier (due to a known limitation, only project number is supported by this field) of the project that contains the secret. If not set, it will be populated with the function&#39;s project, assuming that the secret exists in the same project as of the function.
     * 
     */
    private final @Nullable String projectId;
    /**
     * @return ID of the secret in secret manager (not the full resource name).
     * 
     */
    private final String secret;
    /**
     * @return Version of the secret (version number or the string &#34;latest&#34;). It is preferable to use &#34;latest&#34; version with secret volumes as secret value changes are reflected immediately.
     * 
     */
    private final String version;

    @CustomType.Constructor
    private FunctionSecretEnvironmentVariable(
        @CustomType.Parameter("key") String key,
        @CustomType.Parameter("projectId") @Nullable String projectId,
        @CustomType.Parameter("secret") String secret,
        @CustomType.Parameter("version") String version) {
        this.key = key;
        this.projectId = projectId;
        this.secret = secret;
        this.version = version;
    }

    /**
     * @return Name of the environment variable.
     * 
     */
    public String key() {
        return this.key;
    }
    /**
     * @return Project identifier (due to a known limitation, only project number is supported by this field) of the project that contains the secret. If not set, it will be populated with the function&#39;s project, assuming that the secret exists in the same project as of the function.
     * 
     */
    public Optional<String> projectId() {
        return Optional.ofNullable(this.projectId);
    }
    /**
     * @return ID of the secret in secret manager (not the full resource name).
     * 
     */
    public String secret() {
        return this.secret;
    }
    /**
     * @return Version of the secret (version number or the string &#34;latest&#34;). It is preferable to use &#34;latest&#34; version with secret volumes as secret value changes are reflected immediately.
     * 
     */
    public String version() {
        return this.version;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(FunctionSecretEnvironmentVariable defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private String key;
        private @Nullable String projectId;
        private String secret;
        private String version;

        public Builder() {
    	      // Empty
        }

        public Builder(FunctionSecretEnvironmentVariable defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.key = defaults.key;
    	      this.projectId = defaults.projectId;
    	      this.secret = defaults.secret;
    	      this.version = defaults.version;
        }

        public Builder key(String key) {
            this.key = Objects.requireNonNull(key);
            return this;
        }
        public Builder projectId(@Nullable String projectId) {
            this.projectId = projectId;
            return this;
        }
        public Builder secret(String secret) {
            this.secret = Objects.requireNonNull(secret);
            return this;
        }
        public Builder version(String version) {
            this.version = Objects.requireNonNull(version);
            return this;
        }        public FunctionSecretEnvironmentVariable build() {
            return new FunctionSecretEnvironmentVariable(key, projectId, secret, version);
        }
    }
}
