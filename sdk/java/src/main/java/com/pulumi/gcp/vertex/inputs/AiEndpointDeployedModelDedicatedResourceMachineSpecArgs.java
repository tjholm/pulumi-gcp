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


public final class AiEndpointDeployedModelDedicatedResourceMachineSpecArgs extends com.pulumi.resources.ResourceArgs {

    public static final AiEndpointDeployedModelDedicatedResourceMachineSpecArgs Empty = new AiEndpointDeployedModelDedicatedResourceMachineSpecArgs();

    @Import(name="acceleratorCount")
    private @Nullable Output<Integer> acceleratorCount;

    public Optional<Output<Integer>> acceleratorCount() {
        return Optional.ofNullable(this.acceleratorCount);
    }

    @Import(name="acceleratorType")
    private @Nullable Output<String> acceleratorType;

    public Optional<Output<String>> acceleratorType() {
        return Optional.ofNullable(this.acceleratorType);
    }

    @Import(name="machineType")
    private @Nullable Output<String> machineType;

    public Optional<Output<String>> machineType() {
        return Optional.ofNullable(this.machineType);
    }

    private AiEndpointDeployedModelDedicatedResourceMachineSpecArgs() {}

    private AiEndpointDeployedModelDedicatedResourceMachineSpecArgs(AiEndpointDeployedModelDedicatedResourceMachineSpecArgs $) {
        this.acceleratorCount = $.acceleratorCount;
        this.acceleratorType = $.acceleratorType;
        this.machineType = $.machineType;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AiEndpointDeployedModelDedicatedResourceMachineSpecArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AiEndpointDeployedModelDedicatedResourceMachineSpecArgs $;

        public Builder() {
            $ = new AiEndpointDeployedModelDedicatedResourceMachineSpecArgs();
        }

        public Builder(AiEndpointDeployedModelDedicatedResourceMachineSpecArgs defaults) {
            $ = new AiEndpointDeployedModelDedicatedResourceMachineSpecArgs(Objects.requireNonNull(defaults));
        }

        public Builder acceleratorCount(@Nullable Output<Integer> acceleratorCount) {
            $.acceleratorCount = acceleratorCount;
            return this;
        }

        public Builder acceleratorCount(Integer acceleratorCount) {
            return acceleratorCount(Output.of(acceleratorCount));
        }

        public Builder acceleratorType(@Nullable Output<String> acceleratorType) {
            $.acceleratorType = acceleratorType;
            return this;
        }

        public Builder acceleratorType(String acceleratorType) {
            return acceleratorType(Output.of(acceleratorType));
        }

        public Builder machineType(@Nullable Output<String> machineType) {
            $.machineType = machineType;
            return this;
        }

        public Builder machineType(String machineType) {
            return machineType(Output.of(machineType));
        }

        public AiEndpointDeployedModelDedicatedResourceMachineSpecArgs build() {
            return $;
        }
    }

}
