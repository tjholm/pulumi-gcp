// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.container.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AzureClusterLoggingConfigComponentConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final AzureClusterLoggingConfigComponentConfigArgs Empty = new AzureClusterLoggingConfigComponentConfigArgs();

    /**
     * Components of the logging configuration to be enabled.
     * 
     */
    @Import(name="enableComponents")
    private @Nullable Output<List<String>> enableComponents;

    /**
     * @return Components of the logging configuration to be enabled.
     * 
     */
    public Optional<Output<List<String>>> enableComponents() {
        return Optional.ofNullable(this.enableComponents);
    }

    private AzureClusterLoggingConfigComponentConfigArgs() {}

    private AzureClusterLoggingConfigComponentConfigArgs(AzureClusterLoggingConfigComponentConfigArgs $) {
        this.enableComponents = $.enableComponents;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AzureClusterLoggingConfigComponentConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AzureClusterLoggingConfigComponentConfigArgs $;

        public Builder() {
            $ = new AzureClusterLoggingConfigComponentConfigArgs();
        }

        public Builder(AzureClusterLoggingConfigComponentConfigArgs defaults) {
            $ = new AzureClusterLoggingConfigComponentConfigArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param enableComponents Components of the logging configuration to be enabled.
         * 
         * @return builder
         * 
         */
        public Builder enableComponents(@Nullable Output<List<String>> enableComponents) {
            $.enableComponents = enableComponents;
            return this;
        }

        /**
         * @param enableComponents Components of the logging configuration to be enabled.
         * 
         * @return builder
         * 
         */
        public Builder enableComponents(List<String> enableComponents) {
            return enableComponents(Output.of(enableComponents));
        }

        /**
         * @param enableComponents Components of the logging configuration to be enabled.
         * 
         * @return builder
         * 
         */
        public Builder enableComponents(String... enableComponents) {
            return enableComponents(List.of(enableComponents));
        }

        public AzureClusterLoggingConfigComponentConfigArgs build() {
            return $;
        }
    }

}
