// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.apphub.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class ServiceAttributesCriticalityArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServiceAttributesCriticalityArgs Empty = new ServiceAttributesCriticalityArgs();

    /**
     * Criticality type.
     * Possible values are: `MISSION_CRITICAL`, `HIGH`, `MEDIUM`, `LOW`.
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return Criticality type.
     * Possible values are: `MISSION_CRITICAL`, `HIGH`, `MEDIUM`, `LOW`.
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    private ServiceAttributesCriticalityArgs() {}

    private ServiceAttributesCriticalityArgs(ServiceAttributesCriticalityArgs $) {
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServiceAttributesCriticalityArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServiceAttributesCriticalityArgs $;

        public Builder() {
            $ = new ServiceAttributesCriticalityArgs();
        }

        public Builder(ServiceAttributesCriticalityArgs defaults) {
            $ = new ServiceAttributesCriticalityArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param type Criticality type.
         * Possible values are: `MISSION_CRITICAL`, `HIGH`, `MEDIUM`, `LOW`.
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Criticality type.
         * Possible values are: `MISSION_CRITICAL`, `HIGH`, `MEDIUM`, `LOW`.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public ServiceAttributesCriticalityArgs build() {
            if ($.type == null) {
                throw new MissingRequiredPropertyException("ServiceAttributesCriticalityArgs", "type");
            }
            return $;
        }
    }

}
