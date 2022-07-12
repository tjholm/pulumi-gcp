// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.gcp.billing.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.util.Objects;

@CustomType
public final class BudgetBudgetFilterCustomPeriodStartDate {
    /**
     * @return Day of a month. Must be from 1 to 31 and valid for the year and month.
     * 
     */
    private final Integer day;
    /**
     * @return Month of a year. Must be from 1 to 12.
     * 
     */
    private final Integer month;
    /**
     * @return Year of the date. Must be from 1 to 9999.
     * 
     */
    private final Integer year;

    @CustomType.Constructor
    private BudgetBudgetFilterCustomPeriodStartDate(
        @CustomType.Parameter("day") Integer day,
        @CustomType.Parameter("month") Integer month,
        @CustomType.Parameter("year") Integer year) {
        this.day = day;
        this.month = month;
        this.year = year;
    }

    /**
     * @return Day of a month. Must be from 1 to 31 and valid for the year and month.
     * 
     */
    public Integer day() {
        return this.day;
    }
    /**
     * @return Month of a year. Must be from 1 to 12.
     * 
     */
    public Integer month() {
        return this.month;
    }
    /**
     * @return Year of the date. Must be from 1 to 9999.
     * 
     */
    public Integer year() {
        return this.year;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(BudgetBudgetFilterCustomPeriodStartDate defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private Integer day;
        private Integer month;
        private Integer year;

        public Builder() {
    	      // Empty
        }

        public Builder(BudgetBudgetFilterCustomPeriodStartDate defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.day = defaults.day;
    	      this.month = defaults.month;
    	      this.year = defaults.year;
        }

        public Builder day(Integer day) {
            this.day = Objects.requireNonNull(day);
            return this;
        }
        public Builder month(Integer month) {
            this.month = Objects.requireNonNull(month);
            return this;
        }
        public Builder year(Integer year) {
            this.year = Objects.requireNonNull(year);
            return this;
        }        public BudgetBudgetFilterCustomPeriodStartDate build() {
            return new BudgetBudgetFilterCustomPeriodStartDate(day, month, year);
        }
    }
}
