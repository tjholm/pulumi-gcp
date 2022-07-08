// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.gcp.compute.outputs.GetBackendBucketCdnPolicy;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetBackendBucketResult {
    private final String bucketName;
    private final List<GetBackendBucketCdnPolicy> cdnPolicies;
    private final String creationTimestamp;
    private final List<String> customResponseHeaders;
    private final String description;
    private final String edgeSecurityPolicy;
    private final Boolean enableCdn;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private final String id;
    private final String name;
    private final @Nullable String project;
    private final String selfLink;

    @CustomType.Constructor
    private GetBackendBucketResult(
        @CustomType.Parameter("bucketName") String bucketName,
        @CustomType.Parameter("cdnPolicies") List<GetBackendBucketCdnPolicy> cdnPolicies,
        @CustomType.Parameter("creationTimestamp") String creationTimestamp,
        @CustomType.Parameter("customResponseHeaders") List<String> customResponseHeaders,
        @CustomType.Parameter("description") String description,
        @CustomType.Parameter("edgeSecurityPolicy") String edgeSecurityPolicy,
        @CustomType.Parameter("enableCdn") Boolean enableCdn,
        @CustomType.Parameter("id") String id,
        @CustomType.Parameter("name") String name,
        @CustomType.Parameter("project") @Nullable String project,
        @CustomType.Parameter("selfLink") String selfLink) {
        this.bucketName = bucketName;
        this.cdnPolicies = cdnPolicies;
        this.creationTimestamp = creationTimestamp;
        this.customResponseHeaders = customResponseHeaders;
        this.description = description;
        this.edgeSecurityPolicy = edgeSecurityPolicy;
        this.enableCdn = enableCdn;
        this.id = id;
        this.name = name;
        this.project = project;
        this.selfLink = selfLink;
    }

    public String bucketName() {
        return this.bucketName;
    }
    public List<GetBackendBucketCdnPolicy> cdnPolicies() {
        return this.cdnPolicies;
    }
    public String creationTimestamp() {
        return this.creationTimestamp;
    }
    public List<String> customResponseHeaders() {
        return this.customResponseHeaders;
    }
    public String description() {
        return this.description;
    }
    public String edgeSecurityPolicy() {
        return this.edgeSecurityPolicy;
    }
    public Boolean enableCdn() {
        return this.enableCdn;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String name() {
        return this.name;
    }
    public Optional<String> project() {
        return Optional.ofNullable(this.project);
    }
    public String selfLink() {
        return this.selfLink;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBackendBucketResult defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private String bucketName;
        private List<GetBackendBucketCdnPolicy> cdnPolicies;
        private String creationTimestamp;
        private List<String> customResponseHeaders;
        private String description;
        private String edgeSecurityPolicy;
        private Boolean enableCdn;
        private String id;
        private String name;
        private @Nullable String project;
        private String selfLink;

        public Builder() {
    	      // Empty
        }

        public Builder(GetBackendBucketResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.bucketName = defaults.bucketName;
    	      this.cdnPolicies = defaults.cdnPolicies;
    	      this.creationTimestamp = defaults.creationTimestamp;
    	      this.customResponseHeaders = defaults.customResponseHeaders;
    	      this.description = defaults.description;
    	      this.edgeSecurityPolicy = defaults.edgeSecurityPolicy;
    	      this.enableCdn = defaults.enableCdn;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.project = defaults.project;
    	      this.selfLink = defaults.selfLink;
        }

        public Builder bucketName(String bucketName) {
            this.bucketName = Objects.requireNonNull(bucketName);
            return this;
        }
        public Builder cdnPolicies(List<GetBackendBucketCdnPolicy> cdnPolicies) {
            this.cdnPolicies = Objects.requireNonNull(cdnPolicies);
            return this;
        }
        public Builder cdnPolicies(GetBackendBucketCdnPolicy... cdnPolicies) {
            return cdnPolicies(List.of(cdnPolicies));
        }
        public Builder creationTimestamp(String creationTimestamp) {
            this.creationTimestamp = Objects.requireNonNull(creationTimestamp);
            return this;
        }
        public Builder customResponseHeaders(List<String> customResponseHeaders) {
            this.customResponseHeaders = Objects.requireNonNull(customResponseHeaders);
            return this;
        }
        public Builder customResponseHeaders(String... customResponseHeaders) {
            return customResponseHeaders(List.of(customResponseHeaders));
        }
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
            return this;
        }
        public Builder edgeSecurityPolicy(String edgeSecurityPolicy) {
            this.edgeSecurityPolicy = Objects.requireNonNull(edgeSecurityPolicy);
            return this;
        }
        public Builder enableCdn(Boolean enableCdn) {
            this.enableCdn = Objects.requireNonNull(enableCdn);
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
        public Builder project(@Nullable String project) {
            this.project = project;
            return this;
        }
        public Builder selfLink(String selfLink) {
            this.selfLink = Objects.requireNonNull(selfLink);
            return this;
        }        public GetBackendBucketResult build() {
            return new GetBackendBucketResult(bucketName, cdnPolicies, creationTimestamp, customResponseHeaders, description, edgeSecurityPolicy, enableCdn, id, name, project, selfLink);
        }
    }
}
