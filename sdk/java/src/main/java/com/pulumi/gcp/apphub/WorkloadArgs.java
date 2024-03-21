// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.apphub;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.gcp.apphub.inputs.WorkloadAttributesArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class WorkloadArgs extends com.pulumi.resources.ResourceArgs {

    public static final WorkloadArgs Empty = new WorkloadArgs();

    /**
     * Part of `parent`.  Full resource name of a parent Application. Example: projects/{HOST_PROJECT_ID}/locations/{LOCATION}/applications/{APPLICATION_ID}
     * 
     */
    @Import(name="applicationId", required=true)
    private Output<String> applicationId;

    /**
     * @return Part of `parent`.  Full resource name of a parent Application. Example: projects/{HOST_PROJECT_ID}/locations/{LOCATION}/applications/{APPLICATION_ID}
     * 
     */
    public Output<String> applicationId() {
        return this.applicationId;
    }

    /**
     * Consumer provided attributes.
     * Structure is documented below.
     * 
     */
    @Import(name="attributes")
    private @Nullable Output<WorkloadAttributesArgs> attributes;

    /**
     * @return Consumer provided attributes.
     * Structure is documented below.
     * 
     */
    public Optional<Output<WorkloadAttributesArgs>> attributes() {
        return Optional.ofNullable(this.attributes);
    }

    /**
     * User-defined description of a Workload.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return User-defined description of a Workload.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * Immutable. The resource name of the original discovered workload.
     * 
     */
    @Import(name="discoveredWorkload", required=true)
    private Output<String> discoveredWorkload;

    /**
     * @return Immutable. The resource name of the original discovered workload.
     * 
     */
    public Output<String> discoveredWorkload() {
        return this.discoveredWorkload;
    }

    /**
     * User-defined name for the Workload.
     * 
     */
    @Import(name="displayName")
    private @Nullable Output<String> displayName;

    /**
     * @return User-defined name for the Workload.
     * 
     */
    public Optional<Output<String>> displayName() {
        return Optional.ofNullable(this.displayName);
    }

    /**
     * Part of `parent`.  Full resource name of a parent Application. Example: projects/{HOST_PROJECT_ID}/locations/{LOCATION}/applications/{APPLICATION_ID}
     * 
     */
    @Import(name="location", required=true)
    private Output<String> location;

    /**
     * @return Part of `parent`.  Full resource name of a parent Application. Example: projects/{HOST_PROJECT_ID}/locations/{LOCATION}/applications/{APPLICATION_ID}
     * 
     */
    public Output<String> location() {
        return this.location;
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
     * The Workload identifier.
     * 
     * ***
     * 
     */
    @Import(name="workloadId", required=true)
    private Output<String> workloadId;

    /**
     * @return The Workload identifier.
     * 
     * ***
     * 
     */
    public Output<String> workloadId() {
        return this.workloadId;
    }

    private WorkloadArgs() {}

    private WorkloadArgs(WorkloadArgs $) {
        this.applicationId = $.applicationId;
        this.attributes = $.attributes;
        this.description = $.description;
        this.discoveredWorkload = $.discoveredWorkload;
        this.displayName = $.displayName;
        this.location = $.location;
        this.project = $.project;
        this.workloadId = $.workloadId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(WorkloadArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private WorkloadArgs $;

        public Builder() {
            $ = new WorkloadArgs();
        }

        public Builder(WorkloadArgs defaults) {
            $ = new WorkloadArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param applicationId Part of `parent`.  Full resource name of a parent Application. Example: projects/{HOST_PROJECT_ID}/locations/{LOCATION}/applications/{APPLICATION_ID}
         * 
         * @return builder
         * 
         */
        public Builder applicationId(Output<String> applicationId) {
            $.applicationId = applicationId;
            return this;
        }

        /**
         * @param applicationId Part of `parent`.  Full resource name of a parent Application. Example: projects/{HOST_PROJECT_ID}/locations/{LOCATION}/applications/{APPLICATION_ID}
         * 
         * @return builder
         * 
         */
        public Builder applicationId(String applicationId) {
            return applicationId(Output.of(applicationId));
        }

        /**
         * @param attributes Consumer provided attributes.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder attributes(@Nullable Output<WorkloadAttributesArgs> attributes) {
            $.attributes = attributes;
            return this;
        }

        /**
         * @param attributes Consumer provided attributes.
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder attributes(WorkloadAttributesArgs attributes) {
            return attributes(Output.of(attributes));
        }

        /**
         * @param description User-defined description of a Workload.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description User-defined description of a Workload.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param discoveredWorkload Immutable. The resource name of the original discovered workload.
         * 
         * @return builder
         * 
         */
        public Builder discoveredWorkload(Output<String> discoveredWorkload) {
            $.discoveredWorkload = discoveredWorkload;
            return this;
        }

        /**
         * @param discoveredWorkload Immutable. The resource name of the original discovered workload.
         * 
         * @return builder
         * 
         */
        public Builder discoveredWorkload(String discoveredWorkload) {
            return discoveredWorkload(Output.of(discoveredWorkload));
        }

        /**
         * @param displayName User-defined name for the Workload.
         * 
         * @return builder
         * 
         */
        public Builder displayName(@Nullable Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName User-defined name for the Workload.
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
        }

        /**
         * @param location Part of `parent`.  Full resource name of a parent Application. Example: projects/{HOST_PROJECT_ID}/locations/{LOCATION}/applications/{APPLICATION_ID}
         * 
         * @return builder
         * 
         */
        public Builder location(Output<String> location) {
            $.location = location;
            return this;
        }

        /**
         * @param location Part of `parent`.  Full resource name of a parent Application. Example: projects/{HOST_PROJECT_ID}/locations/{LOCATION}/applications/{APPLICATION_ID}
         * 
         * @return builder
         * 
         */
        public Builder location(String location) {
            return location(Output.of(location));
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
         * @param workloadId The Workload identifier.
         * 
         * ***
         * 
         * @return builder
         * 
         */
        public Builder workloadId(Output<String> workloadId) {
            $.workloadId = workloadId;
            return this;
        }

        /**
         * @param workloadId The Workload identifier.
         * 
         * ***
         * 
         * @return builder
         * 
         */
        public Builder workloadId(String workloadId) {
            return workloadId(Output.of(workloadId));
        }

        public WorkloadArgs build() {
            if ($.applicationId == null) {
                throw new MissingRequiredPropertyException("WorkloadArgs", "applicationId");
            }
            if ($.discoveredWorkload == null) {
                throw new MissingRequiredPropertyException("WorkloadArgs", "discoveredWorkload");
            }
            if ($.location == null) {
                throw new MissingRequiredPropertyException("WorkloadArgs", "location");
            }
            if ($.workloadId == null) {
                throw new MissingRequiredPropertyException("WorkloadArgs", "workloadId");
            }
            return $;
        }
    }

}
