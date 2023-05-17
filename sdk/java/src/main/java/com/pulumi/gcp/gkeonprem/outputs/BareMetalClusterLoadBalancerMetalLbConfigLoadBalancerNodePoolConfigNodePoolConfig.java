// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.gkeonprem.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.gcp.gkeonprem.outputs.BareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigNodeConfig;
import com.pulumi.gcp.gkeonprem.outputs.BareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigTaint;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class BareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfig {
    /**
     * @return The map of Kubernetes labels (key/value pairs) to be applied to
     * each node. These will added in addition to any default label(s)
     * that Kubernetes may apply to the node. In case of conflict in
     * label keys, the applied set may differ depending on the Kubernetes
     * version -- it&#39;s best to assume the behavior is undefined and
     * conflicts should be avoided. For more information, including usage
     * and the valid values, see:
     * http://kubernetes.io/v1.1/docs/user-guide/labels.html
     * An object containing a list of &#34;key&#34;: value pairs.
     * Example: { &#34;name&#34;: &#34;wrench&#34;, &#34;mass&#34;: &#34;1.3kg&#34;, &#34;count&#34;: &#34;3&#34; }.
     * 
     */
    private @Nullable Map<String,String> labels;
    /**
     * @return The list of machine addresses in the Bare Metal Node Pool.
     * Structure is documented below.
     * 
     */
    private @Nullable List<BareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigNodeConfig> nodeConfigs;
    /**
     * @return Specifies the nodes operating system (default: LINUX).
     * 
     */
    private @Nullable String operatingSystem;
    /**
     * @return The initial taints assigned to nodes of this node pool.
     * Structure is documented below.
     * 
     */
    private @Nullable List<BareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigTaint> taints;

    private BareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfig() {}
    /**
     * @return The map of Kubernetes labels (key/value pairs) to be applied to
     * each node. These will added in addition to any default label(s)
     * that Kubernetes may apply to the node. In case of conflict in
     * label keys, the applied set may differ depending on the Kubernetes
     * version -- it&#39;s best to assume the behavior is undefined and
     * conflicts should be avoided. For more information, including usage
     * and the valid values, see:
     * http://kubernetes.io/v1.1/docs/user-guide/labels.html
     * An object containing a list of &#34;key&#34;: value pairs.
     * Example: { &#34;name&#34;: &#34;wrench&#34;, &#34;mass&#34;: &#34;1.3kg&#34;, &#34;count&#34;: &#34;3&#34; }.
     * 
     */
    public Map<String,String> labels() {
        return this.labels == null ? Map.of() : this.labels;
    }
    /**
     * @return The list of machine addresses in the Bare Metal Node Pool.
     * Structure is documented below.
     * 
     */
    public List<BareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigNodeConfig> nodeConfigs() {
        return this.nodeConfigs == null ? List.of() : this.nodeConfigs;
    }
    /**
     * @return Specifies the nodes operating system (default: LINUX).
     * 
     */
    public Optional<String> operatingSystem() {
        return Optional.ofNullable(this.operatingSystem);
    }
    /**
     * @return The initial taints assigned to nodes of this node pool.
     * Structure is documented below.
     * 
     */
    public List<BareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigTaint> taints() {
        return this.taints == null ? List.of() : this.taints;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(BareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfig defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Map<String,String> labels;
        private @Nullable List<BareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigNodeConfig> nodeConfigs;
        private @Nullable String operatingSystem;
        private @Nullable List<BareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigTaint> taints;
        public Builder() {}
        public Builder(BareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.labels = defaults.labels;
    	      this.nodeConfigs = defaults.nodeConfigs;
    	      this.operatingSystem = defaults.operatingSystem;
    	      this.taints = defaults.taints;
        }

        @CustomType.Setter
        public Builder labels(@Nullable Map<String,String> labels) {
            this.labels = labels;
            return this;
        }
        @CustomType.Setter
        public Builder nodeConfigs(@Nullable List<BareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigNodeConfig> nodeConfigs) {
            this.nodeConfigs = nodeConfigs;
            return this;
        }
        public Builder nodeConfigs(BareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigNodeConfig... nodeConfigs) {
            return nodeConfigs(List.of(nodeConfigs));
        }
        @CustomType.Setter
        public Builder operatingSystem(@Nullable String operatingSystem) {
            this.operatingSystem = operatingSystem;
            return this;
        }
        @CustomType.Setter
        public Builder taints(@Nullable List<BareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigTaint> taints) {
            this.taints = taints;
            return this;
        }
        public Builder taints(BareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfigTaint... taints) {
            return taints(List.of(taints));
        }
        public BareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfig build() {
            final var o = new BareMetalClusterLoadBalancerMetalLbConfigLoadBalancerNodePoolConfigNodePoolConfig();
            o.labels = labels;
            o.nodeConfigs = nodeConfigs;
            o.operatingSystem = operatingSystem;
            o.taints = taints;
            return o;
        }
    }
}
