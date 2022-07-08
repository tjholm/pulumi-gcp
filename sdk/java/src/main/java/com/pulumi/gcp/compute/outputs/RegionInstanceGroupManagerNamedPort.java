// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class RegionInstanceGroupManagerNamedPort {
    /**
     * @return - Version name.
     * 
     */
    private final String name;
    /**
     * @return The port number.
     * ***
     * 
     */
    private final Integer port;

    @CustomType.Constructor
    private RegionInstanceGroupManagerNamedPort(
        @CustomType.Parameter("name") String name,
        @CustomType.Parameter("port") Integer port) {
        this.name = name;
        this.port = port;
    }

    /**
     * @return - Version name.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return The port number.
     * ***
     * 
     */
    public Integer port() {
        return this.port;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RegionInstanceGroupManagerNamedPort defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private String name;
        private Integer port;

        public Builder() {
    	      // Empty
        }

        public Builder(RegionInstanceGroupManagerNamedPort defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.name = defaults.name;
    	      this.port = defaults.port;
        }

        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        public Builder port(Integer port) {
            this.port = Objects.requireNonNull(port);
            return this;
        }        public RegionInstanceGroupManagerNamedPort build() {
            return new RegionInstanceGroupManagerNamedPort(name, port);
        }
    }
}
