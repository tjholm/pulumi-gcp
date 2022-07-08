// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.monitoring.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetSecretVersionResult {
    /**
     * @return The time at which the Secret was created.
     * 
     */
    private final String createTime;
    /**
     * @return The time at which the Secret was destroyed. Only present if state is DESTROYED.
     * 
     */
    private final String destroyTime;
    /**
     * @return True if the current state of the SecretVersion is enabled.
     * 
     */
    private final Boolean enabled;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private final String id;
    /**
     * @return The resource name of the SecretVersion. Format:
     * `projects/{{project}}/secrets/{{secret_id}}/versions/{{version}}`
     * 
     */
    private final String name;
    private final String project;
    private final String secret;
    /**
     * @return The secret data. No larger than 64KiB.
     * 
     */
    private final String secretData;
    private final String version;

    @CustomType.Constructor
    private GetSecretVersionResult(
        @CustomType.Parameter("createTime") String createTime,
        @CustomType.Parameter("destroyTime") String destroyTime,
        @CustomType.Parameter("enabled") Boolean enabled,
        @CustomType.Parameter("id") String id,
        @CustomType.Parameter("name") String name,
        @CustomType.Parameter("project") String project,
        @CustomType.Parameter("secret") String secret,
        @CustomType.Parameter("secretData") String secretData,
        @CustomType.Parameter("version") String version) {
        this.createTime = createTime;
        this.destroyTime = destroyTime;
        this.enabled = enabled;
        this.id = id;
        this.name = name;
        this.project = project;
        this.secret = secret;
        this.secretData = secretData;
        this.version = version;
    }

    /**
     * @return The time at which the Secret was created.
     * 
     */
    public String createTime() {
        return this.createTime;
    }
    /**
     * @return The time at which the Secret was destroyed. Only present if state is DESTROYED.
     * 
     */
    public String destroyTime() {
        return this.destroyTime;
    }
    /**
     * @return True if the current state of the SecretVersion is enabled.
     * 
     */
    public Boolean enabled() {
        return this.enabled;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The resource name of the SecretVersion. Format:
     * `projects/{{project}}/secrets/{{secret_id}}/versions/{{version}}`
     * 
     */
    public String name() {
        return this.name;
    }
    public String project() {
        return this.project;
    }
    public String secret() {
        return this.secret;
    }
    /**
     * @return The secret data. No larger than 64KiB.
     * 
     */
    public String secretData() {
        return this.secretData;
    }
    public String version() {
        return this.version;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetSecretVersionResult defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private String createTime;
        private String destroyTime;
        private Boolean enabled;
        private String id;
        private String name;
        private String project;
        private String secret;
        private String secretData;
        private String version;

        public Builder() {
    	      // Empty
        }

        public Builder(GetSecretVersionResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.createTime = defaults.createTime;
    	      this.destroyTime = defaults.destroyTime;
    	      this.enabled = defaults.enabled;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.project = defaults.project;
    	      this.secret = defaults.secret;
    	      this.secretData = defaults.secretData;
    	      this.version = defaults.version;
        }

        public Builder createTime(String createTime) {
            this.createTime = Objects.requireNonNull(createTime);
            return this;
        }
        public Builder destroyTime(String destroyTime) {
            this.destroyTime = Objects.requireNonNull(destroyTime);
            return this;
        }
        public Builder enabled(Boolean enabled) {
            this.enabled = Objects.requireNonNull(enabled);
            return this;
        }
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        public Builder project(String project) {
            this.project = Objects.requireNonNull(project);
            return this;
        }
        public Builder secret(String secret) {
            this.secret = Objects.requireNonNull(secret);
            return this;
        }
        public Builder secretData(String secretData) {
            this.secretData = Objects.requireNonNull(secretData);
            return this;
        }
        public Builder version(String version) {
            this.version = Objects.requireNonNull(version);
            return this;
        }        public GetSecretVersionResult build() {
            return new GetSecretVersionResult(createTime, destroyTime, enabled, id, name, project, secret, secretData, version);
        }
    }
}
