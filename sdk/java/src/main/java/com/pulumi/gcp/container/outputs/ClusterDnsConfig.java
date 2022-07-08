// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.container.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ClusterDnsConfig {
    /**
     * @return Which in-cluster DNS provider should be used. `PROVIDER_UNSPECIFIED` (default) or `PLATFORM_DEFAULT` or `CLOUD_DNS`.
     * 
     */
    private final @Nullable String clusterDns;
    /**
     * @return The suffix used for all cluster service records.
     * 
     */
    private final @Nullable String clusterDnsDomain;
    /**
     * @return The scope of access to cluster DNS records. `DNS_SCOPE_UNSPECIFIED` (default) or `CLUSTER_SCOPE` or `VPC_SCOPE`.
     * 
     */
    private final @Nullable String clusterDnsScope;

    @CustomType.Constructor
    private ClusterDnsConfig(
        @CustomType.Parameter("clusterDns") @Nullable String clusterDns,
        @CustomType.Parameter("clusterDnsDomain") @Nullable String clusterDnsDomain,
        @CustomType.Parameter("clusterDnsScope") @Nullable String clusterDnsScope) {
        this.clusterDns = clusterDns;
        this.clusterDnsDomain = clusterDnsDomain;
        this.clusterDnsScope = clusterDnsScope;
    }

    /**
     * @return Which in-cluster DNS provider should be used. `PROVIDER_UNSPECIFIED` (default) or `PLATFORM_DEFAULT` or `CLOUD_DNS`.
     * 
     */
    public Optional<String> clusterDns() {
        return Optional.ofNullable(this.clusterDns);
    }
    /**
     * @return The suffix used for all cluster service records.
     * 
     */
    public Optional<String> clusterDnsDomain() {
        return Optional.ofNullable(this.clusterDnsDomain);
    }
    /**
     * @return The scope of access to cluster DNS records. `DNS_SCOPE_UNSPECIFIED` (default) or `CLUSTER_SCOPE` or `VPC_SCOPE`.
     * 
     */
    public Optional<String> clusterDnsScope() {
        return Optional.ofNullable(this.clusterDnsScope);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ClusterDnsConfig defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable String clusterDns;
        private @Nullable String clusterDnsDomain;
        private @Nullable String clusterDnsScope;

        public Builder() {
    	      // Empty
        }

        public Builder(ClusterDnsConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.clusterDns = defaults.clusterDns;
    	      this.clusterDnsDomain = defaults.clusterDnsDomain;
    	      this.clusterDnsScope = defaults.clusterDnsScope;
        }

        public Builder clusterDns(@Nullable String clusterDns) {
            this.clusterDns = clusterDns;
            return this;
        }
        public Builder clusterDnsDomain(@Nullable String clusterDnsDomain) {
            this.clusterDnsDomain = clusterDnsDomain;
            return this;
        }
        public Builder clusterDnsScope(@Nullable String clusterDnsScope) {
            this.clusterDnsScope = clusterDnsScope;
            return this;
        }        public ClusterDnsConfig build() {
            return new ClusterDnsConfig(clusterDns, clusterDnsDomain, clusterDnsScope);
        }
    }
}
