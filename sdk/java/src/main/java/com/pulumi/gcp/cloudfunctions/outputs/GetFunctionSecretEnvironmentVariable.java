// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.cloudfunctions.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetFunctionSecretEnvironmentVariable {
    private final String key;
    private final String projectId;
    private final String secret;
    private final String version;

    @CustomType.Constructor
    private GetFunctionSecretEnvironmentVariable(
        @CustomType.Parameter("key") String key,
        @CustomType.Parameter("projectId") String projectId,
        @CustomType.Parameter("secret") String secret,
        @CustomType.Parameter("version") String version) {
        this.key = key;
        this.projectId = projectId;
        this.secret = secret;
        this.version = version;
    }

    public String key() {
        return this.key;
    }
    public String projectId() {
        return this.projectId;
    }
    public String secret() {
        return this.secret;
    }
    public String version() {
        return this.version;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetFunctionSecretEnvironmentVariable defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private String key;
        private String projectId;
        private String secret;
        private String version;

        public Builder() {
    	      // Empty
        }

        public Builder(GetFunctionSecretEnvironmentVariable defaults) {
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
        public Builder projectId(String projectId) {
            this.projectId = Objects.requireNonNull(projectId);
            return this;
        }
        public Builder secret(String secret) {
            this.secret = Objects.requireNonNull(secret);
            return this;
        }
        public Builder version(String version) {
            this.version = Objects.requireNonNull(version);
            return this;
        }        public GetFunctionSecretEnvironmentVariable build() {
            return new GetFunctionSecretEnvironmentVariable(key, projectId, secret, version);
        }
    }
}
