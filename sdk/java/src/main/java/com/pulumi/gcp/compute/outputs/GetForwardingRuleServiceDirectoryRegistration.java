// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetForwardingRuleServiceDirectoryRegistration {
    private final String namespace;
    private final String service;

    @CustomType.Constructor
    private GetForwardingRuleServiceDirectoryRegistration(
        @CustomType.Parameter("namespace") String namespace,
        @CustomType.Parameter("service") String service) {
        this.namespace = namespace;
        this.service = service;
    }

    public String namespace() {
        return this.namespace;
    }
    public String service() {
        return this.service;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetForwardingRuleServiceDirectoryRegistration defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private String namespace;
        private String service;

        public Builder() {
    	      // Empty
        }

        public Builder(GetForwardingRuleServiceDirectoryRegistration defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.namespace = defaults.namespace;
    	      this.service = defaults.service;
        }

        public Builder namespace(String namespace) {
            this.namespace = Objects.requireNonNull(namespace);
            return this;
        }
        public Builder service(String service) {
            this.service = Objects.requireNonNull(service);
            return this;
        }        public GetForwardingRuleServiceDirectoryRegistration build() {
            return new GetForwardingRuleServiceDirectoryRegistration(namespace, service);
        }
    }
}
