// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.container.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.util.Objects;

@CustomType
public final class AwsNodePoolAutoscaling {
    /**
     * @return Maximum number of nodes in the NodePool. Must be &gt;= min_node_count.
     * 
     */
    private final Integer maxNodeCount;
    /**
     * @return Minimum number of nodes in the NodePool. Must be &gt;= 1 and &lt;= max_node_count.
     * 
     */
    private final Integer minNodeCount;

    @CustomType.Constructor
    private AwsNodePoolAutoscaling(
        @CustomType.Parameter("maxNodeCount") Integer maxNodeCount,
        @CustomType.Parameter("minNodeCount") Integer minNodeCount) {
        this.maxNodeCount = maxNodeCount;
        this.minNodeCount = minNodeCount;
    }

    /**
     * @return Maximum number of nodes in the NodePool. Must be &gt;= min_node_count.
     * 
     */
    public Integer maxNodeCount() {
        return this.maxNodeCount;
    }
    /**
     * @return Minimum number of nodes in the NodePool. Must be &gt;= 1 and &lt;= max_node_count.
     * 
     */
    public Integer minNodeCount() {
        return this.minNodeCount;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(AwsNodePoolAutoscaling defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private Integer maxNodeCount;
        private Integer minNodeCount;

        public Builder() {
    	      // Empty
        }

        public Builder(AwsNodePoolAutoscaling defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.maxNodeCount = defaults.maxNodeCount;
    	      this.minNodeCount = defaults.minNodeCount;
        }

        public Builder maxNodeCount(Integer maxNodeCount) {
            this.maxNodeCount = Objects.requireNonNull(maxNodeCount);
            return this;
        }
        public Builder minNodeCount(Integer minNodeCount) {
            this.minNodeCount = Objects.requireNonNull(minNodeCount);
            return this;
        }        public AwsNodePoolAutoscaling build() {
            return new AwsNodePoolAutoscaling(maxNodeCount, minNodeCount);
        }
    }
}
