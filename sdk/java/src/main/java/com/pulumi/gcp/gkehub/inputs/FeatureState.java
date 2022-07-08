// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.gkehub.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.gkehub.inputs.FeatureResourceStateArgs;
import com.pulumi.gcp.gkehub.inputs.FeatureSpecArgs;
import com.pulumi.gcp.gkehub.inputs.FeatureStateArgs;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class FeatureState extends com.pulumi.resources.ResourceArgs {

    public static final FeatureState Empty = new FeatureState();

    /**
     * Output only. When the Feature resource was created.
     * 
     */
    @Import(name="createTime")
    private @Nullable Output<String> createTime;

    /**
     * @return Output only. When the Feature resource was created.
     * 
     */
    public Optional<Output<String>> createTime() {
        return Optional.ofNullable(this.createTime);
    }

    /**
     * Output only. When the Feature resource was deleted.
     * 
     */
    @Import(name="deleteTime")
    private @Nullable Output<String> deleteTime;

    /**
     * @return Output only. When the Feature resource was deleted.
     * 
     */
    public Optional<Output<String>> deleteTime() {
        return Optional.ofNullable(this.deleteTime);
    }

    /**
     * GCP labels for this Feature.
     * 
     */
    @Import(name="labels")
    private @Nullable Output<Map<String,String>> labels;

    /**
     * @return GCP labels for this Feature.
     * 
     */
    public Optional<Output<Map<String,String>>> labels() {
        return Optional.ofNullable(this.labels);
    }

    /**
     * The location for the resource
     * 
     */
    @Import(name="location")
    private @Nullable Output<String> location;

    /**
     * @return The location for the resource
     * 
     */
    public Optional<Output<String>> location() {
        return Optional.ofNullable(this.location);
    }

    /**
     * The full, unique name of this Feature resource
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The full, unique name of this Feature resource
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

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
     * State of the Feature resource itself.
     * 
     */
    @Import(name="resourceStates")
    private @Nullable Output<List<FeatureResourceStateArgs>> resourceStates;

    /**
     * @return State of the Feature resource itself.
     * 
     */
    public Optional<Output<List<FeatureResourceStateArgs>>> resourceStates() {
        return Optional.ofNullable(this.resourceStates);
    }

    /**
     * Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.
     * 
     */
    @Import(name="spec")
    private @Nullable Output<FeatureSpecArgs> spec;

    /**
     * @return Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.
     * 
     */
    public Optional<Output<FeatureSpecArgs>> spec() {
        return Optional.ofNullable(this.spec);
    }

    /**
     * Output only. The Hub-wide Feature state
     * 
     */
    @Import(name="states")
    private @Nullable Output<List<FeatureStateArgs>> states;

    /**
     * @return Output only. The Hub-wide Feature state
     * 
     */
    public Optional<Output<List<FeatureStateArgs>>> states() {
        return Optional.ofNullable(this.states);
    }

    /**
     * Output only. When the Feature resource was last updated.
     * 
     */
    @Import(name="updateTime")
    private @Nullable Output<String> updateTime;

    /**
     * @return Output only. When the Feature resource was last updated.
     * 
     */
    public Optional<Output<String>> updateTime() {
        return Optional.ofNullable(this.updateTime);
    }

    private FeatureState() {}

    private FeatureState(FeatureState $) {
        this.createTime = $.createTime;
        this.deleteTime = $.deleteTime;
        this.labels = $.labels;
        this.location = $.location;
        this.name = $.name;
        this.project = $.project;
        this.resourceStates = $.resourceStates;
        this.spec = $.spec;
        this.states = $.states;
        this.updateTime = $.updateTime;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FeatureState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FeatureState $;

        public Builder() {
            $ = new FeatureState();
        }

        public Builder(FeatureState defaults) {
            $ = new FeatureState(Objects.requireNonNull(defaults));
        }

        /**
         * @param createTime Output only. When the Feature resource was created.
         * 
         * @return builder
         * 
         */
        public Builder createTime(@Nullable Output<String> createTime) {
            $.createTime = createTime;
            return this;
        }

        /**
         * @param createTime Output only. When the Feature resource was created.
         * 
         * @return builder
         * 
         */
        public Builder createTime(String createTime) {
            return createTime(Output.of(createTime));
        }

        /**
         * @param deleteTime Output only. When the Feature resource was deleted.
         * 
         * @return builder
         * 
         */
        public Builder deleteTime(@Nullable Output<String> deleteTime) {
            $.deleteTime = deleteTime;
            return this;
        }

        /**
         * @param deleteTime Output only. When the Feature resource was deleted.
         * 
         * @return builder
         * 
         */
        public Builder deleteTime(String deleteTime) {
            return deleteTime(Output.of(deleteTime));
        }

        /**
         * @param labels GCP labels for this Feature.
         * 
         * @return builder
         * 
         */
        public Builder labels(@Nullable Output<Map<String,String>> labels) {
            $.labels = labels;
            return this;
        }

        /**
         * @param labels GCP labels for this Feature.
         * 
         * @return builder
         * 
         */
        public Builder labels(Map<String,String> labels) {
            return labels(Output.of(labels));
        }

        /**
         * @param location The location for the resource
         * 
         * @return builder
         * 
         */
        public Builder location(@Nullable Output<String> location) {
            $.location = location;
            return this;
        }

        /**
         * @param location The location for the resource
         * 
         * @return builder
         * 
         */
        public Builder location(String location) {
            return location(Output.of(location));
        }

        /**
         * @param name The full, unique name of this Feature resource
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The full, unique name of this Feature resource
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
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
         * @param resourceStates State of the Feature resource itself.
         * 
         * @return builder
         * 
         */
        public Builder resourceStates(@Nullable Output<List<FeatureResourceStateArgs>> resourceStates) {
            $.resourceStates = resourceStates;
            return this;
        }

        /**
         * @param resourceStates State of the Feature resource itself.
         * 
         * @return builder
         * 
         */
        public Builder resourceStates(List<FeatureResourceStateArgs> resourceStates) {
            return resourceStates(Output.of(resourceStates));
        }

        /**
         * @param resourceStates State of the Feature resource itself.
         * 
         * @return builder
         * 
         */
        public Builder resourceStates(FeatureResourceStateArgs... resourceStates) {
            return resourceStates(List.of(resourceStates));
        }

        /**
         * @param spec Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.
         * 
         * @return builder
         * 
         */
        public Builder spec(@Nullable Output<FeatureSpecArgs> spec) {
            $.spec = spec;
            return this;
        }

        /**
         * @param spec Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.
         * 
         * @return builder
         * 
         */
        public Builder spec(FeatureSpecArgs spec) {
            return spec(Output.of(spec));
        }

        /**
         * @param states Output only. The Hub-wide Feature state
         * 
         * @return builder
         * 
         */
        public Builder states(@Nullable Output<List<FeatureStateArgs>> states) {
            $.states = states;
            return this;
        }

        /**
         * @param states Output only. The Hub-wide Feature state
         * 
         * @return builder
         * 
         */
        public Builder states(List<FeatureStateArgs> states) {
            return states(Output.of(states));
        }

        /**
         * @param states Output only. The Hub-wide Feature state
         * 
         * @return builder
         * 
         */
        public Builder states(FeatureStateArgs... states) {
            return states(List.of(states));
        }

        /**
         * @param updateTime Output only. When the Feature resource was last updated.
         * 
         * @return builder
         * 
         */
        public Builder updateTime(@Nullable Output<String> updateTime) {
            $.updateTime = updateTime;
            return this;
        }

        /**
         * @param updateTime Output only. When the Feature resource was last updated.
         * 
         * @return builder
         * 
         */
        public Builder updateTime(String updateTime) {
            return updateTime(Output.of(updateTime));
        }

        public FeatureState build() {
            return $;
        }
    }

}
