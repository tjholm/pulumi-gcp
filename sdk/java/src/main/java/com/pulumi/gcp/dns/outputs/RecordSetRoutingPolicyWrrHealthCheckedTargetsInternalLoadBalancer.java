// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.dns.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class RecordSetRoutingPolicyWrrHealthCheckedTargetsInternalLoadBalancer {
    /**
     * @return The frontend IP address of the load balancer.
     * 
     */
    private String ipAddress;
    /**
     * @return The configured IP protocol of the load balancer. This value is case-sensitive. Possible values: [&#34;tcp&#34;, &#34;udp&#34;]
     * 
     */
    private String ipProtocol;
    /**
     * @return The type of load balancer. This value is case-sensitive. Possible values: [&#34;regionalL4ilb&#34;]
     * 
     */
    private String loadBalancerType;
    /**
     * @return The fully qualified url of the network in which the load balancer belongs. This should be formatted like `projects/{project}/global/networks/{network}` or `https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}`.
     * 
     */
    private String networkUrl;
    /**
     * @return The configured port of the load balancer.
     * 
     */
    private String port;
    /**
     * @return The ID of the project in which the load balancer belongs.
     * 
     */
    private String project;
    /**
     * @return The region of the load balancer. Only needed for regional load balancers.
     * 
     */
    private @Nullable String region;

    private RecordSetRoutingPolicyWrrHealthCheckedTargetsInternalLoadBalancer() {}
    /**
     * @return The frontend IP address of the load balancer.
     * 
     */
    public String ipAddress() {
        return this.ipAddress;
    }
    /**
     * @return The configured IP protocol of the load balancer. This value is case-sensitive. Possible values: [&#34;tcp&#34;, &#34;udp&#34;]
     * 
     */
    public String ipProtocol() {
        return this.ipProtocol;
    }
    /**
     * @return The type of load balancer. This value is case-sensitive. Possible values: [&#34;regionalL4ilb&#34;]
     * 
     */
    public String loadBalancerType() {
        return this.loadBalancerType;
    }
    /**
     * @return The fully qualified url of the network in which the load balancer belongs. This should be formatted like `projects/{project}/global/networks/{network}` or `https://www.googleapis.com/compute/v1/projects/{project}/global/networks/{network}`.
     * 
     */
    public String networkUrl() {
        return this.networkUrl;
    }
    /**
     * @return The configured port of the load balancer.
     * 
     */
    public String port() {
        return this.port;
    }
    /**
     * @return The ID of the project in which the load balancer belongs.
     * 
     */
    public String project() {
        return this.project;
    }
    /**
     * @return The region of the load balancer. Only needed for regional load balancers.
     * 
     */
    public Optional<String> region() {
        return Optional.ofNullable(this.region);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RecordSetRoutingPolicyWrrHealthCheckedTargetsInternalLoadBalancer defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String ipAddress;
        private String ipProtocol;
        private String loadBalancerType;
        private String networkUrl;
        private String port;
        private String project;
        private @Nullable String region;
        public Builder() {}
        public Builder(RecordSetRoutingPolicyWrrHealthCheckedTargetsInternalLoadBalancer defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.ipAddress = defaults.ipAddress;
    	      this.ipProtocol = defaults.ipProtocol;
    	      this.loadBalancerType = defaults.loadBalancerType;
    	      this.networkUrl = defaults.networkUrl;
    	      this.port = defaults.port;
    	      this.project = defaults.project;
    	      this.region = defaults.region;
        }

        @CustomType.Setter
        public Builder ipAddress(String ipAddress) {
            this.ipAddress = Objects.requireNonNull(ipAddress);
            return this;
        }
        @CustomType.Setter
        public Builder ipProtocol(String ipProtocol) {
            this.ipProtocol = Objects.requireNonNull(ipProtocol);
            return this;
        }
        @CustomType.Setter
        public Builder loadBalancerType(String loadBalancerType) {
            this.loadBalancerType = Objects.requireNonNull(loadBalancerType);
            return this;
        }
        @CustomType.Setter
        public Builder networkUrl(String networkUrl) {
            this.networkUrl = Objects.requireNonNull(networkUrl);
            return this;
        }
        @CustomType.Setter
        public Builder port(String port) {
            this.port = Objects.requireNonNull(port);
            return this;
        }
        @CustomType.Setter
        public Builder project(String project) {
            this.project = Objects.requireNonNull(project);
            return this;
        }
        @CustomType.Setter
        public Builder region(@Nullable String region) {
            this.region = region;
            return this;
        }
        public RecordSetRoutingPolicyWrrHealthCheckedTargetsInternalLoadBalancer build() {
            final var o = new RecordSetRoutingPolicyWrrHealthCheckedTargetsInternalLoadBalancer();
            o.ipAddress = ipAddress;
            o.ipProtocol = ipProtocol;
            o.loadBalancerType = loadBalancerType;
            o.networkUrl = networkUrl;
            o.port = port;
            o.project = project;
            o.region = region;
            return o;
        }
    }
}
