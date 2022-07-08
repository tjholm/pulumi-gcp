// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.storage.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.gcp.storage.outputs.GetBucketLifecycleRuleAction;
import com.pulumi.gcp.storage.outputs.GetBucketLifecycleRuleCondition;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetBucketLifecycleRule {
    private final List<GetBucketLifecycleRuleAction> actions;
    private final List<GetBucketLifecycleRuleCondition> conditions;

    @CustomType.Constructor
    private GetBucketLifecycleRule(
        @CustomType.Parameter("actions") List<GetBucketLifecycleRuleAction> actions,
        @CustomType.Parameter("conditions") List<GetBucketLifecycleRuleCondition> conditions) {
        this.actions = actions;
        this.conditions = conditions;
    }

    public List<GetBucketLifecycleRuleAction> actions() {
        return this.actions;
    }
    public List<GetBucketLifecycleRuleCondition> conditions() {
        return this.conditions;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetBucketLifecycleRule defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private List<GetBucketLifecycleRuleAction> actions;
        private List<GetBucketLifecycleRuleCondition> conditions;

        public Builder() {
    	      // Empty
        }

        public Builder(GetBucketLifecycleRule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.actions = defaults.actions;
    	      this.conditions = defaults.conditions;
        }

        public Builder actions(List<GetBucketLifecycleRuleAction> actions) {
            this.actions = Objects.requireNonNull(actions);
            return this;
        }
        public Builder actions(GetBucketLifecycleRuleAction... actions) {
            return actions(List.of(actions));
        }
        public Builder conditions(List<GetBucketLifecycleRuleCondition> conditions) {
            this.conditions = Objects.requireNonNull(conditions);
            return this;
        }
        public Builder conditions(GetBucketLifecycleRuleCondition... conditions) {
            return conditions(List.of(conditions));
        }        public GetBucketLifecycleRule build() {
            return new GetBucketLifecycleRule(actions, conditions);
        }
    }
}
