// Copyright 2019 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * OpenAPI spec version: 1.0.0
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 */


package io.swagger.client.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

/**
 * V1RandomSearch
 */

public class V1RandomSearch {
  @SerializedName("kind")
  private String kind = null;

  @SerializedName("matrix")
  private String matrix = null;

  @SerializedName("n_runs")
  private String nRuns = null;

  @SerializedName("seed")
  private String seed = null;

  @SerializedName("concurrency")
  private String concurrency = null;

  @SerializedName("early_stopping")
  private List<Object> earlyStopping = null;

  public V1RandomSearch kind(String kind) {
    this.kind = kind;
    return this;
  }

   /**
   * Get kind
   * @return kind
  **/
  @ApiModelProperty(value = "")
  public String getKind() {
    return kind;
  }

  public void setKind(String kind) {
    this.kind = kind;
  }

  public V1RandomSearch matrix(String matrix) {
    this.matrix = matrix;
    return this;
  }

   /**
   * Get matrix
   * @return matrix
  **/
  @ApiModelProperty(value = "")
  public String getMatrix() {
    return matrix;
  }

  public void setMatrix(String matrix) {
    this.matrix = matrix;
  }

  public V1RandomSearch nRuns(String nRuns) {
    this.nRuns = nRuns;
    return this;
  }

   /**
   * Get nRuns
   * @return nRuns
  **/
  @ApiModelProperty(value = "")
  public String getNRuns() {
    return nRuns;
  }

  public void setNRuns(String nRuns) {
    this.nRuns = nRuns;
  }

  public V1RandomSearch seed(String seed) {
    this.seed = seed;
    return this;
  }

   /**
   * Get seed
   * @return seed
  **/
  @ApiModelProperty(value = "")
  public String getSeed() {
    return seed;
  }

  public void setSeed(String seed) {
    this.seed = seed;
  }

  public V1RandomSearch concurrency(String concurrency) {
    this.concurrency = concurrency;
    return this;
  }

   /**
   * Get concurrency
   * @return concurrency
  **/
  @ApiModelProperty(value = "")
  public String getConcurrency() {
    return concurrency;
  }

  public void setConcurrency(String concurrency) {
    this.concurrency = concurrency;
  }

  public V1RandomSearch earlyStopping(List<Object> earlyStopping) {
    this.earlyStopping = earlyStopping;
    return this;
  }

  public V1RandomSearch addEarlyStoppingItem(Object earlyStoppingItem) {
    if (this.earlyStopping == null) {
      this.earlyStopping = new ArrayList<Object>();
    }
    this.earlyStopping.add(earlyStoppingItem);
    return this;
  }

   /**
   * Get earlyStopping
   * @return earlyStopping
  **/
  @ApiModelProperty(value = "")
  public List<Object> getEarlyStopping() {
    return earlyStopping;
  }

  public void setEarlyStopping(List<Object> earlyStopping) {
    this.earlyStopping = earlyStopping;
  }


  @Override
  public boolean equals(java.lang.Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    V1RandomSearch v1RandomSearch = (V1RandomSearch) o;
    return Objects.equals(this.kind, v1RandomSearch.kind) &&
        Objects.equals(this.matrix, v1RandomSearch.matrix) &&
        Objects.equals(this.nRuns, v1RandomSearch.nRuns) &&
        Objects.equals(this.seed, v1RandomSearch.seed) &&
        Objects.equals(this.concurrency, v1RandomSearch.concurrency) &&
        Objects.equals(this.earlyStopping, v1RandomSearch.earlyStopping);
  }

  @Override
  public int hashCode() {
    return Objects.hash(kind, matrix, nRuns, seed, concurrency, earlyStopping);
  }


  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class V1RandomSearch {\n");
    
    sb.append("    kind: ").append(toIndentedString(kind)).append("\n");
    sb.append("    matrix: ").append(toIndentedString(matrix)).append("\n");
    sb.append("    nRuns: ").append(toIndentedString(nRuns)).append("\n");
    sb.append("    seed: ").append(toIndentedString(seed)).append("\n");
    sb.append("    concurrency: ").append(toIndentedString(concurrency)).append("\n");
    sb.append("    earlyStopping: ").append(toIndentedString(earlyStopping)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(java.lang.Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }

}

