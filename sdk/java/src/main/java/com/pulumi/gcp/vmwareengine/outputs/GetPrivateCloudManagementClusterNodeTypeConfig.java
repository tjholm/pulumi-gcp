// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.vmwareengine.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetPrivateCloudManagementClusterNodeTypeConfig {
    private Integer customCoreCount;
    private Integer nodeCount;
    private String nodeTypeId;

    private GetPrivateCloudManagementClusterNodeTypeConfig() {}
    public Integer customCoreCount() {
        return this.customCoreCount;
    }
    public Integer nodeCount() {
        return this.nodeCount;
    }
    public String nodeTypeId() {
        return this.nodeTypeId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetPrivateCloudManagementClusterNodeTypeConfig defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Integer customCoreCount;
        private Integer nodeCount;
        private String nodeTypeId;
        public Builder() {}
        public Builder(GetPrivateCloudManagementClusterNodeTypeConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.customCoreCount = defaults.customCoreCount;
    	      this.nodeCount = defaults.nodeCount;
    	      this.nodeTypeId = defaults.nodeTypeId;
        }

        @CustomType.Setter
        public Builder customCoreCount(Integer customCoreCount) {
            this.customCoreCount = Objects.requireNonNull(customCoreCount);
            return this;
        }
        @CustomType.Setter
        public Builder nodeCount(Integer nodeCount) {
            this.nodeCount = Objects.requireNonNull(nodeCount);
            return this;
        }
        @CustomType.Setter
        public Builder nodeTypeId(String nodeTypeId) {
            this.nodeTypeId = Objects.requireNonNull(nodeTypeId);
            return this;
        }
        public GetPrivateCloudManagementClusterNodeTypeConfig build() {
            final var o = new GetPrivateCloudManagementClusterNodeTypeConfig();
            o.customCoreCount = customCoreCount;
            o.nodeCount = nodeCount;
            o.nodeTypeId = nodeTypeId;
            return o;
        }
    }
}
