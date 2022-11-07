// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.vertex.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AiEndpointDeployedModelDedicatedResourceAutoscalingMetricSpecArgs extends com.pulumi.resources.ResourceArgs {

    public static final AiEndpointDeployedModelDedicatedResourceAutoscalingMetricSpecArgs Empty = new AiEndpointDeployedModelDedicatedResourceAutoscalingMetricSpecArgs();

    @Import(name="metricName")
    private @Nullable Output<String> metricName;

    public Optional<Output<String>> metricName() {
        return Optional.ofNullable(this.metricName);
    }

    @Import(name="target")
    private @Nullable Output<Integer> target;

    public Optional<Output<Integer>> target() {
        return Optional.ofNullable(this.target);
    }

    private AiEndpointDeployedModelDedicatedResourceAutoscalingMetricSpecArgs() {}

    private AiEndpointDeployedModelDedicatedResourceAutoscalingMetricSpecArgs(AiEndpointDeployedModelDedicatedResourceAutoscalingMetricSpecArgs $) {
        this.metricName = $.metricName;
        this.target = $.target;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AiEndpointDeployedModelDedicatedResourceAutoscalingMetricSpecArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AiEndpointDeployedModelDedicatedResourceAutoscalingMetricSpecArgs $;

        public Builder() {
            $ = new AiEndpointDeployedModelDedicatedResourceAutoscalingMetricSpecArgs();
        }

        public Builder(AiEndpointDeployedModelDedicatedResourceAutoscalingMetricSpecArgs defaults) {
            $ = new AiEndpointDeployedModelDedicatedResourceAutoscalingMetricSpecArgs(Objects.requireNonNull(defaults));
        }

        public Builder metricName(@Nullable Output<String> metricName) {
            $.metricName = metricName;
            return this;
        }

        public Builder metricName(String metricName) {
            return metricName(Output.of(metricName));
        }

        public Builder target(@Nullable Output<Integer> target) {
            $.target = target;
            return this;
        }

        public Builder target(Integer target) {
            return target(Output.of(target));
        }

        public AiEndpointDeployedModelDedicatedResourceAutoscalingMetricSpecArgs build() {
            return $;
        }
    }

}
