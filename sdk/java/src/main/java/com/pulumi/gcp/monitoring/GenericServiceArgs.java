// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.monitoring;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.gcp.monitoring.inputs.GenericServiceBasicServiceArgs;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GenericServiceArgs extends com.pulumi.resources.ResourceArgs {

    public static final GenericServiceArgs Empty = new GenericServiceArgs();

    /**
     * A well-known service type, defined by its service type and service labels.
     * Valid values are described at
     * https://cloud.google.com/stackdriver/docs/solutions/slo-monitoring/api/api-structures#basic-svc-w-basic-sli
     * Structure is documented below.
     * 
     */
    @Import(name="basicService")
    private @Nullable Output<GenericServiceBasicServiceArgs> basicService;

    /**
     * @return A well-known service type, defined by its service type and service labels.
     * Valid values are described at
     * https://cloud.google.com/stackdriver/docs/solutions/slo-monitoring/api/api-structures#basic-svc-w-basic-sli
     * Structure is documented below.
     * 
     */
    public Optional<Output<GenericServiceBasicServiceArgs>> basicService() {
        return Optional.ofNullable(this.basicService);
    }

    /**
     * Name used for UI elements listing this Service.
     * 
     */
    @Import(name="displayName")
    private @Nullable Output<String> displayName;

    /**
     * @return Name used for UI elements listing this Service.
     * 
     */
    public Optional<Output<String>> displayName() {
        return Optional.ofNullable(this.displayName);
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
     * An optional service ID to use. If not given, the server will generate a
     * service ID.
     * 
     */
    @Import(name="serviceId", required=true)
    private Output<String> serviceId;

    /**
     * @return An optional service ID to use. If not given, the server will generate a
     * service ID.
     * 
     */
    public Output<String> serviceId() {
        return this.serviceId;
    }

    /**
     * Labels which have been used to annotate the service. Label keys must start
     * with a letter. Label keys and values may contain lowercase letters,
     * numbers, underscores, and dashes. Label keys and values have a maximum
     * length of 63 characters, and must be less than 128 bytes in size. Up to 64
     * label entries may be stored. For labels which do not have a semantic value,
     * the empty string may be supplied for the label value.
     * 
     */
    @Import(name="userLabels")
    private @Nullable Output<Map<String,String>> userLabels;

    /**
     * @return Labels which have been used to annotate the service. Label keys must start
     * with a letter. Label keys and values may contain lowercase letters,
     * numbers, underscores, and dashes. Label keys and values have a maximum
     * length of 63 characters, and must be less than 128 bytes in size. Up to 64
     * label entries may be stored. For labels which do not have a semantic value,
     * the empty string may be supplied for the label value.
     * 
     */
    public Optional<Output<Map<String,String>>> userLabels() {
        return Optional.ofNullable(this.userLabels);
    }

    private GenericServiceArgs() {}

    private GenericServiceArgs(GenericServiceArgs $) {
        this.basicService = $.basicService;
        this.displayName = $.displayName;
        this.project = $.project;
        this.serviceId = $.serviceId;
        this.userLabels = $.userLabels;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GenericServiceArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GenericServiceArgs $;

        public Builder() {
            $ = new GenericServiceArgs();
        }

        public Builder(GenericServiceArgs defaults) {
            $ = new GenericServiceArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param basicService A well-known service type, defined by its service type and service labels.
         * Valid values are described at
         * https://cloud.google.com/stackdriver/docs/solutions/slo-monitoring/api/api-structures#basic-svc-w-basic-sli
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder basicService(@Nullable Output<GenericServiceBasicServiceArgs> basicService) {
            $.basicService = basicService;
            return this;
        }

        /**
         * @param basicService A well-known service type, defined by its service type and service labels.
         * Valid values are described at
         * https://cloud.google.com/stackdriver/docs/solutions/slo-monitoring/api/api-structures#basic-svc-w-basic-sli
         * Structure is documented below.
         * 
         * @return builder
         * 
         */
        public Builder basicService(GenericServiceBasicServiceArgs basicService) {
            return basicService(Output.of(basicService));
        }

        /**
         * @param displayName Name used for UI elements listing this Service.
         * 
         * @return builder
         * 
         */
        public Builder displayName(@Nullable Output<String> displayName) {
            $.displayName = displayName;
            return this;
        }

        /**
         * @param displayName Name used for UI elements listing this Service.
         * 
         * @return builder
         * 
         */
        public Builder displayName(String displayName) {
            return displayName(Output.of(displayName));
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
         * @param serviceId An optional service ID to use. If not given, the server will generate a
         * service ID.
         * 
         * @return builder
         * 
         */
        public Builder serviceId(Output<String> serviceId) {
            $.serviceId = serviceId;
            return this;
        }

        /**
         * @param serviceId An optional service ID to use. If not given, the server will generate a
         * service ID.
         * 
         * @return builder
         * 
         */
        public Builder serviceId(String serviceId) {
            return serviceId(Output.of(serviceId));
        }

        /**
         * @param userLabels Labels which have been used to annotate the service. Label keys must start
         * with a letter. Label keys and values may contain lowercase letters,
         * numbers, underscores, and dashes. Label keys and values have a maximum
         * length of 63 characters, and must be less than 128 bytes in size. Up to 64
         * label entries may be stored. For labels which do not have a semantic value,
         * the empty string may be supplied for the label value.
         * 
         * @return builder
         * 
         */
        public Builder userLabels(@Nullable Output<Map<String,String>> userLabels) {
            $.userLabels = userLabels;
            return this;
        }

        /**
         * @param userLabels Labels which have been used to annotate the service. Label keys must start
         * with a letter. Label keys and values may contain lowercase letters,
         * numbers, underscores, and dashes. Label keys and values have a maximum
         * length of 63 characters, and must be less than 128 bytes in size. Up to 64
         * label entries may be stored. For labels which do not have a semantic value,
         * the empty string may be supplied for the label value.
         * 
         * @return builder
         * 
         */
        public Builder userLabels(Map<String,String> userLabels) {
            return userLabels(Output.of(userLabels));
        }

        public GenericServiceArgs build() {
            $.serviceId = Objects.requireNonNull($.serviceId, "expected parameter 'serviceId' to be non-null");
            return $;
        }
    }

}
