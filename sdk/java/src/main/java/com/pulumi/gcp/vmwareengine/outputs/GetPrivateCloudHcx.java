// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.vmwareengine.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetPrivateCloudHcx {
    private String fqdn;
    private String internalIp;
    private String state;
    private String version;

    private GetPrivateCloudHcx() {}
    public String fqdn() {
        return this.fqdn;
    }
    public String internalIp() {
        return this.internalIp;
    }
    public String state() {
        return this.state;
    }
    public String version() {
        return this.version;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetPrivateCloudHcx defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String fqdn;
        private String internalIp;
        private String state;
        private String version;
        public Builder() {}
        public Builder(GetPrivateCloudHcx defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.fqdn = defaults.fqdn;
    	      this.internalIp = defaults.internalIp;
    	      this.state = defaults.state;
    	      this.version = defaults.version;
        }

        @CustomType.Setter
        public Builder fqdn(String fqdn) {
            this.fqdn = Objects.requireNonNull(fqdn);
            return this;
        }
        @CustomType.Setter
        public Builder internalIp(String internalIp) {
            this.internalIp = Objects.requireNonNull(internalIp);
            return this;
        }
        @CustomType.Setter
        public Builder state(String state) {
            this.state = Objects.requireNonNull(state);
            return this;
        }
        @CustomType.Setter
        public Builder version(String version) {
            this.version = Objects.requireNonNull(version);
            return this;
        }
        public GetPrivateCloudHcx build() {
            final var o = new GetPrivateCloudHcx();
            o.fqdn = fqdn;
            o.internalIp = internalIp;
            o.state = state;
            o.version = version;
            return o;
        }
    }
}
