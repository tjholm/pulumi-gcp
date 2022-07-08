// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.firebaserules;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.firebaserules.inputs.RulesetSourceArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RulesetArgs extends com.pulumi.resources.ResourceArgs {

    public static final RulesetArgs Empty = new RulesetArgs();

    /**
     * The project for the resource
     * 
     */
    @Import(name="project")
    private @Nullable Output<String> project;

    /**
     * @return The project for the resource
     * 
     */
    public Optional<Output<String>> project() {
        return Optional.ofNullable(this.project);
    }

    /**
     * `Source` for the `Ruleset`.
     * 
     */
    @Import(name="source", required=true)
    private Output<RulesetSourceArgs> source;

    /**
     * @return `Source` for the `Ruleset`.
     * 
     */
    public Output<RulesetSourceArgs> source() {
        return this.source;
    }

    private RulesetArgs() {}

    private RulesetArgs(RulesetArgs $) {
        this.project = $.project;
        this.source = $.source;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RulesetArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RulesetArgs $;

        public Builder() {
            $ = new RulesetArgs();
        }

        public Builder(RulesetArgs defaults) {
            $ = new RulesetArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param project The project for the resource
         * 
         * @return builder
         * 
         */
        public Builder project(@Nullable Output<String> project) {
            $.project = project;
            return this;
        }

        /**
         * @param project The project for the resource
         * 
         * @return builder
         * 
         */
        public Builder project(String project) {
            return project(Output.of(project));
        }

        /**
         * @param source `Source` for the `Ruleset`.
         * 
         * @return builder
         * 
         */
        public Builder source(Output<RulesetSourceArgs> source) {
            $.source = source;
            return this;
        }

        /**
         * @param source `Source` for the `Ruleset`.
         * 
         * @return builder
         * 
         */
        public Builder source(RulesetSourceArgs source) {
            return source(Output.of(source));
        }

        public RulesetArgs build() {
            $.source = Objects.requireNonNull($.source, "expected parameter 'source' to be non-null");
            return $;
        }
    }

}
