// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.runtimeconfig.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetVariableArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetVariableArgs Empty = new GetVariableArgs();

    /**
     * The name of the Runtime Configurator configuration.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return The name of the Runtime Configurator configuration.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * The name of the RuntimeConfig resource containing this variable.
     * 
     */
    @Import(name="parent", required=true)
    private Output<String> parent;

    /**
     * @return The name of the RuntimeConfig resource containing this variable.
     * 
     */
    public Output<String> parent() {
        return this.parent;
    }

    /**
     * The project in which the resource belongs. If it
     * is not provided, the provider project is used.
     * 
     */
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return The project in which the resource belongs. If it
     * is not provided, the provider project is used.
     * 
     */
    public Optional<Output<String>> project() {
        return Optional.ofNullable(this.project);
    }

    private GetVariableArgs() {}

    private GetVariableArgs(GetVariableArgs $) {
        this.name = $.name;
        this.parent = $.parent;
        this.project = $.project;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetVariableArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetVariableArgs $;

        public Builder() {
            $ = new GetVariableArgs();
        }

        public Builder(GetVariableArgs defaults) {
            $ = new GetVariableArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The name of the Runtime Configurator configuration.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the Runtime Configurator configuration.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param parent The name of the RuntimeConfig resource containing this variable.
         * 
         * @return builder
         * 
         */
        public Builder parent(Output<String> parent) {
            $.parent = parent;
            return this;
        }

        /**
         * @param parent The name of the RuntimeConfig resource containing this variable.
         * 
         * @return builder
         * 
         */
        public Builder parent(String parent) {
            return parent(Output.of(parent));
        }

        /**
         * @param project The project in which the resource belongs. If it
         * is not provided, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The project in which the resource belongs. If it
         * is not provided, the provider project is used.
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        public GetVariableArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            $.parent = Objects.requireNonNull($.parent, "expected parameter 'parent' to be non-null");
            return $;
        }
    }

}
