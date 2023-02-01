// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.cloudbuild.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetTriggerGitFileSource {
    private String githubEnterpriseConfig;
    private String path;
    private String repoType;
    private String revision;
    private String uri;

    private GetTriggerGitFileSource() {}
    public String githubEnterpriseConfig() {
        return this.githubEnterpriseConfig;
    }
    public String path() {
        return this.path;
    }
    public String repoType() {
        return this.repoType;
    }
    public String revision() {
        return this.revision;
    }
    public String uri() {
        return this.uri;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetTriggerGitFileSource defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String githubEnterpriseConfig;
        private String path;
        private String repoType;
        private String revision;
        private String uri;
        public Builder() {}
        public Builder(GetTriggerGitFileSource defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.githubEnterpriseConfig = defaults.githubEnterpriseConfig;
    	      this.path = defaults.path;
    	      this.repoType = defaults.repoType;
    	      this.revision = defaults.revision;
    	      this.uri = defaults.uri;
        }

        @CustomType.Setter
        public Builder githubEnterpriseConfig(String githubEnterpriseConfig) {
            this.githubEnterpriseConfig = Objects.requireNonNull(githubEnterpriseConfig);
            return this;
        }
        @CustomType.Setter
        public Builder path(String path) {
            this.path = Objects.requireNonNull(path);
            return this;
        }
        @CustomType.Setter
        public Builder repoType(String repoType) {
            this.repoType = Objects.requireNonNull(repoType);
            return this;
        }
        @CustomType.Setter
        public Builder revision(String revision) {
            this.revision = Objects.requireNonNull(revision);
            return this;
        }
        @CustomType.Setter
        public Builder uri(String uri) {
            this.uri = Objects.requireNonNull(uri);
            return this;
        }
        public GetTriggerGitFileSource build() {
            final var o = new GetTriggerGitFileSource();
            o.githubEnterpriseConfig = githubEnterpriseConfig;
            o.path = path;
            o.repoType = repoType;
            o.revision = revision;
            o.uri = uri;
            return o;
        }
    }
}
