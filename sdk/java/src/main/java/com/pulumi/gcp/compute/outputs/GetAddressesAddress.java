// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.compute.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetAddressesAddress {
    /**
     * @return The IP address (for example `1.2.3.4`).
     * 
     */
    private String address;
    /**
     * @return The IP address type, can be `EXTERNAL` or `INTERNAL`.
     * 
     */
    private String addressType;
    /**
     * @return The IP address description.
     * 
     */
    private String description;
    /**
     * @return A map containing IP labels.
     * 
     */
    private Map<String,String> labels;
    /**
     * @return The IP address name.
     * 
     */
    private String name;
    /**
     * @return Region that should be considered to search addresses.
     * All regions are considered if missing.
     * 
     */
    private String region;
    /**
     * @return The URI of the created resource.
     * 
     */
    private String selfLink;
    /**
     * @return Indicates if the address is used. Possible values are: RESERVED or IN_USE.
     * 
     */
    private String status;

    private GetAddressesAddress() {}
    /**
     * @return The IP address (for example `1.2.3.4`).
     * 
     */
    public String address() {
        return this.address;
    }
    /**
     * @return The IP address type, can be `EXTERNAL` or `INTERNAL`.
     * 
     */
    public String addressType() {
        return this.addressType;
    }
    /**
     * @return The IP address description.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return A map containing IP labels.
     * 
     */
    public Map<String,String> labels() {
        return this.labels;
    }
    /**
     * @return The IP address name.
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Region that should be considered to search addresses.
     * All regions are considered if missing.
     * 
     */
    public String region() {
        return this.region;
    }
    /**
     * @return The URI of the created resource.
     * 
     */
    public String selfLink() {
        return this.selfLink;
    }
    /**
     * @return Indicates if the address is used. Possible values are: RESERVED or IN_USE.
     * 
     */
    public String status() {
        return this.status;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAddressesAddress defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String address;
        private String addressType;
        private String description;
        private Map<String,String> labels;
        private String name;
        private String region;
        private String selfLink;
        private String status;
        public Builder() {}
        public Builder(GetAddressesAddress defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.address = defaults.address;
    	      this.addressType = defaults.addressType;
    	      this.description = defaults.description;
    	      this.labels = defaults.labels;
    	      this.name = defaults.name;
    	      this.region = defaults.region;
    	      this.selfLink = defaults.selfLink;
    	      this.status = defaults.status;
        }

        @CustomType.Setter
        public Builder address(String address) {
            if (address == null) {
              throw new MissingRequiredPropertyException("GetAddressesAddress", "address");
            }
            this.address = address;
            return this;
        }
        @CustomType.Setter
        public Builder addressType(String addressType) {
            if (addressType == null) {
              throw new MissingRequiredPropertyException("GetAddressesAddress", "addressType");
            }
            this.addressType = addressType;
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            if (description == null) {
              throw new MissingRequiredPropertyException("GetAddressesAddress", "description");
            }
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder labels(Map<String,String> labels) {
            if (labels == null) {
              throw new MissingRequiredPropertyException("GetAddressesAddress", "labels");
            }
            this.labels = labels;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetAddressesAddress", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder region(String region) {
            if (region == null) {
              throw new MissingRequiredPropertyException("GetAddressesAddress", "region");
            }
            this.region = region;
            return this;
        }
        @CustomType.Setter
        public Builder selfLink(String selfLink) {
            if (selfLink == null) {
              throw new MissingRequiredPropertyException("GetAddressesAddress", "selfLink");
            }
            this.selfLink = selfLink;
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            if (status == null) {
              throw new MissingRequiredPropertyException("GetAddressesAddress", "status");
            }
            this.status = status;
            return this;
        }
        public GetAddressesAddress build() {
            final var _resultValue = new GetAddressesAddress();
            _resultValue.address = address;
            _resultValue.addressType = addressType;
            _resultValue.description = description;
            _resultValue.labels = labels;
            _resultValue.name = name;
            _resultValue.region = region;
            _resultValue.selfLink = selfLink;
            _resultValue.status = status;
            return _resultValue;
        }
    }
}
