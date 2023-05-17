// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.databasemigrationservice.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.util.Objects;


public final class ConnectionProfileAlloydbSettingsPrimaryInstanceSettingsMachineConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final ConnectionProfileAlloydbSettingsPrimaryInstanceSettingsMachineConfigArgs Empty = new ConnectionProfileAlloydbSettingsPrimaryInstanceSettingsMachineConfigArgs();

    /**
     * The number of CPU&#39;s in the VM instance.
     * 
     */
    @Import(name="cpuCount", required=true)
    private Output<Integer> cpuCount;

    /**
     * @return The number of CPU&#39;s in the VM instance.
     * 
     */
    public Output<Integer> cpuCount() {
        return this.cpuCount;
    }

    private ConnectionProfileAlloydbSettingsPrimaryInstanceSettingsMachineConfigArgs() {}

    private ConnectionProfileAlloydbSettingsPrimaryInstanceSettingsMachineConfigArgs(ConnectionProfileAlloydbSettingsPrimaryInstanceSettingsMachineConfigArgs $) {
        this.cpuCount = $.cpuCount;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ConnectionProfileAlloydbSettingsPrimaryInstanceSettingsMachineConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ConnectionProfileAlloydbSettingsPrimaryInstanceSettingsMachineConfigArgs $;

        public Builder() {
            $ = new ConnectionProfileAlloydbSettingsPrimaryInstanceSettingsMachineConfigArgs();
        }

        public Builder(ConnectionProfileAlloydbSettingsPrimaryInstanceSettingsMachineConfigArgs defaults) {
            $ = new ConnectionProfileAlloydbSettingsPrimaryInstanceSettingsMachineConfigArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param cpuCount The number of CPU&#39;s in the VM instance.
         * 
         * @return builder
         * 
         */
        public Builder cpuCount(Output<Integer> cpuCount) {
            $.cpuCount = cpuCount;
            return this;
        }

        /**
         * @param cpuCount The number of CPU&#39;s in the VM instance.
         * 
         * @return builder
         * 
         */
        public Builder cpuCount(Integer cpuCount) {
            return cpuCount(Output.of(cpuCount));
        }

        public ConnectionProfileAlloydbSettingsPrimaryInstanceSettingsMachineConfigArgs build() {
            $.cpuCount = Objects.requireNonNull($.cpuCount, "expected parameter 'cpuCount' to be non-null");
            return $;
        }
    }

}
