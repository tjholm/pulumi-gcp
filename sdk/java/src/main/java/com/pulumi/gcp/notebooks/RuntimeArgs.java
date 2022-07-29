// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.notebooks;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.notebooks.inputs.RuntimeAccessConfigArgs;
import com.pulumi.gcp.notebooks.inputs.RuntimeSoftwareConfigArgs;
import com.pulumi.gcp.notebooks.inputs.RuntimeVirtualMachineArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RuntimeArgs extends com.pulumi.resources.ResourceArgs {

    public static final RuntimeArgs Empty = new RuntimeArgs();

    /**
     * The config settings for accessing runtime.
     * Structure is documented below.
     * 
     */
    @Import(name="accessConfig")
    private @Nullable Output<RuntimeAccessConfigArgs> accessConfig;

    /**
     * @return The config settings for accessing runtime.
     * Structure is documented below.
     * 
     */
    public Optional<Output<RuntimeAccessConfigArgs>> accessConfig() {
        return Optional.ofNullable(this.accessConfig);
    }

    /**
     * A reference to the zone where the machine resides.
     * 
     */
    @Import(name="location", required=true)
    private Output<String> location;

    /**
     * @return A reference to the zone where the machine resides.
     * 
     */
    public Output<String> location() {
        return this.location;
    }

    /**
     * The name specified for the Notebook runtime.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name specified for the Notebook runtime.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return The ID of the project in which the resource belongs.
     * If it is not provided, the provider project is used.
     * 
     */
    public Optional<Output<String>> project() {
        return Optional.ofNullable(this.project);
    }

    /**
     * The config settings for software inside the runtime.
     * Structure is documented below.
     * 
     */
    @Import(name="softwareConfig")
    private @Nullable Output<RuntimeSoftwareConfigArgs> softwareConfig;

    /**
     * @return The config settings for software inside the runtime.
     * Structure is documented below.
     * 
     */
    public Optional<Output<RuntimeSoftwareConfigArgs>> softwareConfig() {
        return Optional.ofNullable(this.softwareConfig);
    }

    /**
     * Use a Compute Engine VM image to start the managed notebook instance.
     * Structure is documented below.
     * 
     */
    @Import(name="virtualMachine")
    private @Nullable Output<RuntimeVirtualMachineArgs> virtualMachine;

    /**
     * @return Use a Compute Engine VM image to start the managed notebook instance.
     * Structure is documented below.
     * 
     */
    public Optional<Output<RuntimeVirtualMachineArgs>> virtualMachine() {
        return Optional.ofNullable(this.virtualMachine);
    }

    private RuntimeArgs() {}

    private RuntimeArgs(RuntimeArgs $) {
        this.accessConfig = $.accessConfig;
        this.location = $.location;
        this.name = $.name;
        this.project = $.project;
        this.softwareConfig = $.softwareConfig;
        this.virtualMachine = $.virtualMachine;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RuntimeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RuntimeArgs $;

        public Builder() {
            $ = new RuntimeArgs();
        }

        public Builder(RuntimeArgs defaults) {
            $ = new RuntimeArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param accessConfig The config settings for accessing runtime.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder accessConfig(@Nullable Output<RuntimeAccessConfigArgs> accessConfig) {
            $.accessConfig = accessConfig;
            return this;
        }

        /**
         * @param accessConfig The config settings for accessing runtime.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder accessConfig(RuntimeAccessConfigArgs accessConfig) {
            return accessConfig(Output.of(accessConfig));
        }

        /**
         * @param location A reference to the zone where the machine resides.
         * 
         * @return builder
         * 
         */
        public Builder location(Output<String> location) {
            $.location = location;
            return this;
        }

        /**
         * @param location A reference to the zone where the machine resides.
         * 
         * @return builder
         * 
         */
        public Builder location(String location) {
            return location(Output.of(location));
        }

        /**
         * @param name The name specified for the Notebook runtime.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name specified for the Notebook runtime.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param project The ID of the project in which the resource belongs.
         * If it is not provided, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The ID of the project in which the resource belongs.
         * If it is not provided, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        /**
         * @param softwareConfig The config settings for software inside the runtime.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder softwareConfig(@Nullable Output<RuntimeSoftwareConfigArgs> softwareConfig) {
            $.softwareConfig = softwareConfig;
            return this;
        }

        /**
         * @param softwareConfig The config settings for software inside the runtime.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder softwareConfig(RuntimeSoftwareConfigArgs softwareConfig) {
            return softwareConfig(Output.of(softwareConfig));
        }

        /**
         * @param virtualMachine Use a Compute Engine VM image to start the managed notebook instance.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder virtualMachine(@Nullable Output<RuntimeVirtualMachineArgs> virtualMachine) {
            $.virtualMachine = virtualMachine;
            return this;
        }

        /**
         * @param virtualMachine Use a Compute Engine VM image to start the managed notebook instance.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder virtualMachine(RuntimeVirtualMachineArgs virtualMachine) {
            return virtualMachine(Output.of(virtualMachine));
        }

        public RuntimeArgs build() {
            $.location = Objects.requireNonNull($.location, "expected parameter 'location' to be non-null");
            return $;
        }
    }

}
